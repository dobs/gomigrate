package migrator

type InstallOptions struct {
	Help *bool
}

type UpDownOptions struct {
	PreDeployOnly  *bool
	PostDeployOnly *bool
	Version        *string
	Force          *bool
	Help           *bool
}

type NewOptions struct {
	Name *string
	Help *bool
}

type Options struct {
	Install InstallOptions
	New     NewOptions
	Up      UpDownOptions
	Down    UpDownOptions
}
