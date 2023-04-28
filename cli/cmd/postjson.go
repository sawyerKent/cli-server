package cmd

import (
	"fmt"

	"github.com/sawyerKent/cli-server/server/handlers"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var postJsonCmd = &cobra.Command{
	Use:   "postjson",
	Short: "Post data to the specified endpoint as JSON",
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetString("port")
		endpoint, _ := cmd.Flags().GetString("endpoint")
		baseurl, _ := cmd.Flags().GetString("baseurl")
		frvrid, _ := cmd.Flags().GetString("frvrid")
		language, _ := cmd.Flags().GetString("language")

		// Check if any flags are provided
		if port == "" || endpoint == "" || baseurl == "" {
			// Print command usage
			fmt.Println("No flags provided. Usage:")
			cmd.Usage()
			return
		}

		url := baseurl + ":" + port + endpoint
		fmt.Println("URL:", url)

		data := handlers.HappyLangResponse{
			FRVRID:   frvrid,
			Language: language,
		}
		
		result, err := handlers.PostJsonEndpoint(url, data)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		switch endpoint {
		case "/HappyLang":
			fmt.Println(result.Display())
		}
	},
}

func init() {
	postJsonCmd.Flags().StringP("port", "p", viper.GetString("port"), "Port number of the server")
	postJsonCmd.Flags().StringP("endpoint", "e", viper.GetString("endpoint"), "Endpoint to send POST request")
	postJsonCmd.Flags().StringP("baseurl", "b", viper.GetString("baseurl"), "Base URL of the server")
	postJsonCmd.Flags().StringP("frvrid", "f", "-1", "FRVRID value")
	postJsonCmd.Flags().StringP("language", "l", "Sample", "Language value")
	rootCmd.AddCommand(postJsonCmd)
}
