package cmd

import (
	"encoding/json"
	"fmt"
	fc "github.com/caliagaa/ripfire/cmd/firebase_client"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	_ "google.golang.org/api/option"
	"io/ioutil"
	"log"
)

var (
	documentJsonFileName string
	collectionName string
	documentName string
)

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Uploads a document based on JSON file for a collection to Firestore project",
	Long: `Uploads a document based on JSON file for a collection to Firestore project`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Uploading to collection: ", collectionName)
		uploadCollection(collectionName, documentJsonFileName, documentName)
	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)
	uploadCmd.Flags().StringVarP(&documentJsonFileName, "document", "d", "", "JSON file (as document) to upload")
	uploadCmd.Flags().StringVarP(&documentName, "name", "n", "", "Document ID")
	uploadCmd.Flags().StringVarP(&collectionName, "collection", "c", "", "Collection to upload")
	uploadCmd.MarkFlagRequired("document")
	uploadCmd.MarkFlagRequired("collection")
}

func uploadCollection(collectionName string, documentJsonFileName string, documentName string) {
	ctx := context.Background()

	client := fc.GetClient(ctx)
	defer client.Close()

	file, err0 := ioutil.ReadFile(documentJsonFileName)
	if err0 != nil {
		log.Fatalf("%v", err0)
		return
	}

	jsonObject := make(map[string]interface{})
	err1 := json.Unmarshal(file, &jsonObject)
	if err1 != nil {
		log.Fatalf("%v", err1)
		return
	}
	fmt.Println(len(jsonObject))

	if documentName == "" {
		documentName = fmt.Sprintf("%v", jsonObject["code"])
	}

	_, err := client.Collection(collectionName).Doc(documentName).Set(ctx, jsonObject)
	if err != nil {
		log.Fatalf("Failed to upload document: %v", err)
		return
	}
	log.Println("Document created")
}