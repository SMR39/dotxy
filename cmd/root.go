package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "dotxy",
	Short: "dotxy is a DNS over TLS proxy",
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func validateEnv() {
	requiredEnv := map[string]string{
		"DNS_RESOLVER_ADDR": viper.GetString("DNS_RESOLVER_ADDR"),
		"LISTEN_ADDR":       viper.GetString("LISTEN_ADDR"),
	}

	var missing bool
	for k, v := range requiredEnv {
		if len(v) == 0 {
			missing = true
			fmt.Printf("%s env variable must be set\n", k)
		}
	}
	if missing {
		log.Fatal("missing ENV vars")
	}
}
