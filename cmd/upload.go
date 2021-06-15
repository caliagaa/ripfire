package cmd

import (
	"encoding/json"
	"fmt"
	fc "github.com/caliagaa/ripfire/cmd/firebase_client"
	definition "github.com/caliagaa/ripfire/cmd/models/definition"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	_ "google.golang.org/api/option"
	"io/ioutil"
	"log"
)

var (
	documentName string
	collectionName string
)

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Uploads a document or collection to Firestore project",
	Long: `Uploads a document or collection to Firestore project`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Uploading to collection: ", collectionName)
		uploadCollection(collectionName)
	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)
	uploadCmd.Flags().StringVarP(&documentName, "document", "d", "", "Document to upload")
	uploadCmd.Flags().StringVarP(&collectionName, "collection", "c", "", "Collection to upload")
}

func uploadCollection(collectionName string) {
	ctx := context.Background()

	client := fc.GetClient(ctx)
	defer client.Close()

	world := definition.WorldDef{}
	file, _ := ioutil.ReadFile("test.json")
	_ = json.Unmarshal(file, &world)

	_, err := client.Collection(collectionName).Doc(world.Code).Set(ctx, world)
	if err != nil {
		log.Fatalf("Failed to upload document: %v", err)
		return
	}
}