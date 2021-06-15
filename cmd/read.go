package cmd

import (
	"fmt"
	fc "github.com/caliagaa/ripfire/cmd/firebase_client"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
	"log"
)

// readCmd represents the read command
var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read Firestore collection",
	Long: `Read Firestore collection`,
	Run: func(cmd *cobra.Command, args []string) {
		readCollection(collectionName)
	},
}

func init() {
	rootCmd.AddCommand(readCmd)
	readCmd.Flags().StringVarP(&collectionName, "collection", "c", "", "Collection to read")
}

func readCollection(collectionName string) {
	ctx := context.Background()

	client := fc.GetClient(ctx)
	defer client.Close()

	iter := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		print(doc.Data())
	}
}

func print(m map[string]interface{}) {
	for key, element := range m {
		fmt.Println(key, ":",element)
	}
}
