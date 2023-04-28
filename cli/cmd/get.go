package cmd

import (
	"fmt"

	"github.com/sawyerKent/cli-server/server/handlers"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get data from the specified endpoint",
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetString("port")
		endpoint, _ := cmd.Flags().GetString("endpoint")
		baseurl, _ := cmd.Flags().GetString("baseurl")
	
		// Check if any flags are provided
		if port == "" || endpoint == "" || baseurl == "" {
			// Print command usage
			fmt.Println("No flags provided. Usage:")
			cmd.Usage()
			return
		}
	
		url := baseurl + ":" + port + endpoint
		fmt.Println(url)
		result, err := handlers.GetEndpoint(url)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		switch endpoint {
		case "/":
			fmt.Println(result.Display())
		case "/heartbeat":
			fmt.Println(result.Display())
		case "/HappyLang":
			fmt.Println(result.Display())
		}
	},	
}

func init() {
	getCmd.Flags().StringP("port", "p", viper.GetString("port"), "Port number of the server")
	getCmd.Flags().StringP("endpoint", "e", viper.GetString("endpoint"), "Endpoint to send GET request")
	getCmd.Flags().StringP("baseurl", "b", viper.GetString("baseurl"), "Base URL of the server")
	rootCmd.AddCommand(getCmd)
}
