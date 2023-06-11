package flag

type BuildOptions struct {
	Architectre string
	OS          string
}

var (
	ArchitectreFlag = Flag{
		Name:       "arch",
		ConfigName: "build.arch",
		Value:      "amd64",
		Usage:      "Target architecture",
	}
	OsFlag = Flag{
		Name:       "os",
		ConfigName: "build.os",
		Value:      "linux",
		Usage:      "Target OS",
	}
)

func NewBuildFlagGroup() *BuildFlagGroup {
	return &BuildFlagGroup{
		Architectre: &ArchitectreFlag,
		OS:          &OsFlag,
	}
}

type BuildFlagGroup struct {
	Architectre *Flag
	OS          *Flag
}

func (f *BuildFlagGroup) Name() string {
	return "build"
}

func (f *BuildFlagGroup) Flags() []*Flag {
	return []*Flag{
		f.OS,
		f.Architectre,
	}
}

func (f *BuildFlagGroup) ToOptions() BuildOptions {
	return BuildOptions{
		OS:          f.OS.getString(),
		Architectre: f.Architectre.getString(),
	}
}
