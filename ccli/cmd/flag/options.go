package flag

import (
	"os"
	"strings"
	"sync"
	"time"

	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"golang.org/x/xerrors"
)

type Flag struct {
	// Name is for CLI flag and environment variable.
	// If this field is empty, it will be available only in config file.
	Name string

	// ConfigName is a key in config file. It is also used as a key of viper.
	ConfigName string

	// Shorthand is a shorthand letter.
	Shorthand string

	// Value is the default value. It must be filled to determine the flag type.
	Value interface{}

	// Usage explains how to use the flag.
	Usage string

	// Persistent represents if the flag is persistent
	Persistent bool

	// Deprecated represents if the flag is deprecated
	Deprecated bool

	// Aliases represents aliases
	Aliases []Alias
}

type Alias struct {
	Name       string
	ConfigName string
	Deprecated bool
}

type FlagGroup interface {
	Name() string
	Flags() []*Flag
}

func (f *Flag) getValue() any {
	// First, looks for aliases in config file (trivy.yaml).
	// Note that viper.RegisterAlias cannot be used for this purpose.
	var v any
	for _, alias := range f.Aliases {
		if alias.ConfigName == "" {
			continue
		}
		v = viper.Get(alias.ConfigName)
		if v != nil {
			// log.Logger.Warnf("'%s' in config file is deprecated. Use '%s' instead.", alias.ConfigName, f.ConfigName)
			return v
		}
	}
	return viper.Get(f.ConfigName)
}

func (f *Flag) getString() string {
	return cast.ToString(f.getValue())
}

func bind(cmd *cobra.Command, flag *Flag) error {
	if flag == nil {
		return nil
	} else if flag.Name == "" {
		// This flag is available only in trivy.yaml
		viper.SetDefault(flag.ConfigName, flag.Value)
		return nil
	}

	// Bind CLI flags
	f := cmd.Flags().Lookup(flag.Name)
	if f == nil {
		// Lookup local persistent flags
		f = cmd.PersistentFlags().Lookup(flag.Name)
	}
	if err := viper.BindPFlag(flag.ConfigName, f); err != nil {
		return xerrors.Errorf("bind flag error: %w", err)
	}

	// Bind environmental variable
	if err := bindEnv(flag); err != nil {
		return err
	}

	return nil
}

func bindEnv(flag *Flag) error {
	// We don't use viper.AutomaticEnv, so we need to add a prefix manually here.
	envName := strings.ToUpper("ccli_" + strings.ReplaceAll(flag.Name, "-", "_"))
	if err := viper.BindEnv(flag.ConfigName, envName); err != nil {
		return xerrors.Errorf("bind env error: %w", err)
	}

	// Bind env aliases
	for _, alias := range flag.Aliases {
		envAlias := strings.ToUpper("ccli_" + strings.ReplaceAll(alias.Name, "-", "_"))
		if err := viper.BindEnv(flag.ConfigName, envAlias); err != nil {
			return xerrors.Errorf("bind env error: %w", err)
		}
		if alias.Deprecated {
			if _, ok := os.LookupEnv(envAlias); ok {
				// log.Logger.Warnf("'%s' is deprecated. Use '%s' instead.", envAlias, envName)
			}
		}
	}
	return nil
}

func (f *Flags) AddFlags(cmd *cobra.Command) {
	aliases := make(flagAliases)
	for _, group := range f.groups() {
		for _, flag := range group.Flags() {
			addFlag(cmd, flag)

			// Register flag aliases
			aliases.Add(flag)
		}
	}

	cmd.Flags().SetNormalizeFunc(aliases.NormalizeFunc())
}

type flagAlias struct {
	formalName string
	deprecated bool
	once       sync.Once
}

// flagAliases have aliases for CLI flags
type flagAliases map[string]*flagAlias

func (a flagAliases) Add(flag *Flag) {
	if flag == nil {
		return
	}
	for _, alias := range flag.Aliases {
		a[alias.Name] = &flagAlias{
			formalName: flag.Name,
			deprecated: alias.Deprecated,
		}
	}
}

func addFlag(cmd *cobra.Command, flag *Flag) {
	if flag == nil || flag.Name == "" {
		return
	}
	var flags *pflag.FlagSet
	if flag.Persistent {
		flags = cmd.PersistentFlags()
	} else {
		flags = cmd.Flags()
	}

	switch v := flag.Value.(type) {
	case int:
		flags.IntP(flag.Name, flag.Shorthand, v, flag.Usage)
	case string:
		flags.StringP(flag.Name, flag.Shorthand, v, flag.Usage)
	case []string:
		flags.StringSliceP(flag.Name, flag.Shorthand, v, flag.Usage)
	case bool:
		flags.BoolP(flag.Name, flag.Shorthand, v, flag.Usage)
	case time.Duration:
		flags.DurationP(flag.Name, flag.Shorthand, v, flag.Usage)
	case float64:
		flags.Float64P(flag.Name, flag.Shorthand, v, flag.Usage)
	}

	if flag.Deprecated {
		flags.MarkHidden(flag.Name) // nolint: gosec
	}
}

func (a flagAliases) NormalizeFunc() func(*pflag.FlagSet, string) pflag.NormalizedName {
	return func(_ *pflag.FlagSet, name string) pflag.NormalizedName {
		if alias, ok := a[name]; ok {
			if alias.deprecated {
				// NormalizeFunc is called several times
				alias.once.Do(func() {
					// log.Logger.Warnf("'--%s' is deprecated. Use '--%s' instead.", name, alias.formalName)
				})
			}
			name = alias.formalName
		}
		return pflag.NormalizedName(name)
	}
}
