package cmd

import (
	"log"

	"github.com/SMR39/dotxy/tcp"
	"github.com/SMR39/dotxy/udp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var verbose bool
var network string

// listenCmd represents the command to start listening
var listenCmd = &cobra.Command{
	Use:   "listen",
	Short: "listens to the DNS queries on the client",
	Run: func(cmd *cobra.Command, args []string) {
		validateEnv()
		listenAddr := viper.GetString("LISTEN_ADDR")
		resolverAddr := viper.GetString("DNS_RESOLVER_ADDR")
		switch network {
		case "udp":
			err := udp.ServeUDP(listenAddr, resolverAddr, verbose)
			if err != nil {
				log.Fatal("Couldn't serve udp ", err)
			}
		case "tcp":
			err2 := tcp.ServeTCP(listenAddr, resolverAddr, verbose)
			if err2 != nil {
				log.Fatal("Couldn't serve tcp ", err2)
			}
		default:
			log.Fatal("enter network flag as either udp or tcp")
		}
	}}

func init() {
	RootCmd.AddCommand(listenCmd)
	cobra.OnInitialize(viper.AutomaticEnv)
	listenCmd.Flags().BoolVarP(&verbose, "verbose", "v", true, "verbose logging")
	listenCmd.Flags().StringVarP(&network, "network", "n", "", "netowrk protocol")
}
