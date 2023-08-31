package generate

import (
	"github.com/fededp/dbg-go/pkg/generate"
	"github.com/fededp/dbg-go/pkg/root"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewGenerateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate new dbg configs",
		Long: `
In auto mode, configs will be generated starting from kernel-crawler output. 
In this scenario, target-{distro,kernelrelease,kernelversion} are available to filter to-be-generated configs. Regexes are allowed.
Moreover, you can pass special value "load" as target-distro to make the tool automatically fetch latest distro kernel-crawler ran against.
Instead, when auto mode is disabled, the tool is able to generate a single config (for each driver version).
In this scenario, target-{distro,kernelrelease,kernelversion} CANNOT be regexes but must be exact values.
Also, in non-automatic mode, kernelurls driverkit config key will be constructed using driverkit libraries.
`,
		RunE: execute,
	}
	flags := cmd.Flags()
	flags.Bool("auto", false, "automatically generate configs from kernel-crawler output")
	flags.String("driver-name", "falco", "driver name to be used")
	return cmd
}

func execute(c *cobra.Command, args []string) error {
	options := generate.Options{
		Options:    root.LoadRootOptions(),
		Auto:       viper.GetBool("auto"),
		DriverName: viper.GetString("driver-name"),
	}
	return generate.Run(options)
}
