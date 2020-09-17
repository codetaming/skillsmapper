package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func init() {
	rootCmd.AddCommand(addSkillCmd)
}

var addSkillCmd = &cobra.Command{
	Use:   "add skill",
	Short: "Add a skill",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		client := &http.Client{}
		req, err := http.NewRequest("POST", fmt.Sprint(ServerURI, "/skill"), strings.NewReader(fmt.Sprintf("{\n  \"email\": \"dan@example.com\",\n  \"tag\": \"%s\",\n  \"level\": \"learning\"\n}", args[1])))
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
		if resp.StatusCode != http.StatusCreated {
			fmt.Printf("error: %d\n%s", resp.StatusCode, body)
		} else {
			fmt.Printf("added: %s", body)
		}
	},
}
