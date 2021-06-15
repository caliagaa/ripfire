package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
)

var (
	keyFileName string
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes Firebase authentication via JSON key file",
	Long: `Initializes Firestore connection with service account key file`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Initializing with: ",  keyFileName)
		addServiceAccount(keyFileName)
		fmt.Println("Done")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringVarP(&keyFileName, "file", "f", "", "JSON key file location")
}

func addServiceAccount(keyFileName string) {
	dest := "/.ripfire/key.json"

	input, err := ioutil.ReadFile(keyFileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal( err )
		return
	}

	destinationFile := homedir + dest
	if _, err := os.Stat(homedir + "/.ripfire"); err != nil {
		err = os.Mkdir(homedir + "/.ripfire", 0755)
		if err != nil {
			log.Fatal(err)
			return
		}
	}

	if _, err := os.Stat(destinationFile); err != nil {
		err = ioutil.WriteFile(destinationFile, input, 0644)
		if err != nil {
			fmt.Println("Error copying file to: ", destinationFile)
			fmt.Println(err)
			return
		}
	}
}