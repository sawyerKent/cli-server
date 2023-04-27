package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"cli-server/server/handlers"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var sendJsonCmd = &cobra.Command{
	Use:   "sendjson",
	Short: "Send a JSON file to the server",
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetString("port")
		endpoint, _ := cmd.Flags().GetString("endpoint")
		baseurl, _ := cmd.Flags().GetString("baseurl")
		file, _ := cmd.Flags().GetString("file")

		if port == "" || endpoint == "" || baseurl == "" || file == "" {
			fmt.Println("No flags provided. Usage:")
			cmd.Usage()
			return
		}

		url := baseurl + ":" + port + endpoint
		fmt.Println("URL:", url)

		responseData, err := handlers.SendJson(url, file)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// Beautify JSON output
		var prettyJSON bytes.Buffer
		err = json.Indent(&prettyJSON, responseData, "", "  ")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// Save the beautified JSON to the response.json file
		err = os.WriteFile("response.json", prettyJSON.Bytes(), 0644)
		if err != nil {
			fmt.Println("Error writing response to file:", err)
			return
		}

		fmt.Println("Response saved to response.json")
	},
}

func init() {
	sendJsonCmd.Flags().StringP("port", "p", viper.GetString("port"), "Port number of the server")
	sendJsonCmd.Flags().StringP("endpoint", "e", viper.GetString("endpoint"), "Endpoint to send POST request")
	sendJsonCmd.Flags().StringP("baseurl", "b", viper.GetString("baseurl"), "Base URL of the server")
	sendJsonCmd.Flags().StringP("file", "f", "", "Path to the JSON file")
	rootCmd.AddCommand(sendJsonCmd)
}
