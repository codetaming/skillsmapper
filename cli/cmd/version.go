package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/codetaming/skillsmapper/cli/version"
	"github.com/codetaming/skillsmapper/internal/model"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of cli",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("CLI version:\t", version.Version)
		client := &http.Client{}
		req, err := http.NewRequest("GET", fmt.Sprint(ServerURI, "/version"), strings.NewReader("{}"))
		if err != nil {
			log.Printf("Error making request to %s", ServerURI)
		}
		req.Header.Set("Content-Type", "application/json")
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("Error with request: %s", err.Error())
		}
		body, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			fmt.Printf("Error: %d\n%s", resp.StatusCode, body)
		} else {
			var v model.Version
			json.Unmarshal(body, &v)
			fmt.Printf("Server version:\t %s\n", v.Version)
		}
	},
}
