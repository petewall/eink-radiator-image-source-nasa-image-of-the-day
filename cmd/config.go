package cmd

import (
	"encoding/json"

	"github.com/spf13/cobra"

	"github.com/petewall/eink-radiator-image-source-nasa-image-of-the-day/pkg"
)

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Print a blank config for the " + ImageTypeName + " image type",
	Run: func(cmd *cobra.Command, args []string) {
		encoded, _ := json.Marshal(pkg.Config{
			APIKey: "",
		})
		cmd.Println(string(encoded))
	},
}

func init() {
	rootCmd.AddCommand(ConfigCmd)
	ConfigCmd.SetOut(ConfigCmd.OutOrStdout())
}
