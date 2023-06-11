package flag

import (
	"github.com/spf13/cobra"
	"golang.org/x/xerrors"
)

type Flags struct {
	BuildFlagGroup *BuildFlagGroup
}

type Options struct {
	BuildOptions

	AppVersion string
}

func (f *Flags) Bind(cmd *cobra.Command) error {
	for _, group := range f.groups() {
		if group == nil {
			continue
		}
		for _, flag := range group.Flags() {
			if err := bind(cmd, flag); err != nil {
				return xerrors.Errorf("flag groups: %w", err)
			}
		}
	}
	return nil
}

func (f *Flags) groups() []FlagGroup {
	var groups []FlagGroup
	// This order affects the usage message, so they are sorted by frequency of use.
	if f.BuildFlagGroup != nil {
		groups = append(groups, f.BuildFlagGroup)
	}

	return groups
}

func (f *Flags) ToOptions(appVersion string, args []string) (Options, error) {
	opts := Options{
		AppVersion: appVersion,
	}

	if f.BuildFlagGroup != nil {
		opts.BuildOptions = f.BuildFlagGroup.ToOptions()
	}

	// opts.Align()

	return opts, nil
}
