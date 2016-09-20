package cmd

import (
	"fmt"
	"log"
	"path"

	"github.com/drud/drud-go/drudapi"
	"github.com/spf13/cobra"
)

// localStopCmd represents the stop command
var localRMCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove an application's local services.",
	Long:  `Remove will delete the local service containers from this machine..`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			log.Fatalln("app_name and deploy_name are expected as arguments.")
		}

		if appClient == "" {
			appClient = cfg.Client
		}

		basePath := path.Join(homedir, ".drud", appClient, args[0], args[1])
		err := drudapi.DockerCompose("-f", path.Join(basePath, "docker-compose.yaml"), "stop")
		if err != nil {
			log.Fatalln(err)
		}

		dcErr := drudapi.DockerCompose("-f", path.Join(basePath, "docker-compose.yaml"), "rm", "-f")
		if dcErr != nil {
			fmt.Println(fmt.Errorf("%s", dcErr.Error()))
		}

	},
}

func init() {
	localRMCmd.Flags().StringVarP(&appClient, "client", "c", "", "Client name")
	LocalCmd.AddCommand(localRMCmd)

}