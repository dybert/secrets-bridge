package cmd

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/abourget/secrets-bridge/pkg/bridge"
	"github.com/abourget/secrets-bridge/pkg/client"
	"github.com/spf13/cobra"
)

// killCmd represents the kill command
var killCmd = &cobra.Command{
	Use:   "kill",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		bridgeConf, err := ioutil.ReadFile(bridgeConfFilename)
		if err != nil {
			log.Fatalln("reading %q: %s", bridgeConfFilename, err)
		}

		bridge, err := bridge.NewFromString(string(bridgeConf))
		if err != nil {
			log.Fatalln("--bridge-conf has an invalid value:", err)
		}

		c := client.NewClient(bridge)
		if err := c.ChooseEndpoint(); err != nil {
			log.Fatalln("error pinging server:", err)
		}

		err = c.Quit()
		if err != nil {
			log.Fatalln("error killing previous server:", err)
		}

		fmt.Println("bridge server terminated")
	},
}

func init() {
	RootCmd.AddCommand(killCmd)
}
