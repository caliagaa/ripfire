package firebase_client

import (
	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"log"
	"os"
)

func GetClient(ctx context.Context) *firestore.Client {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	keyFile := homedir + "/.ripfire/key.json"
	if _, err := os.Stat(keyFile); err != nil {
		fmt.Println("No key file defined:", keyFile)
		return nil
	}

	sa := option.WithCredentialsFile(keyFile)
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	return client
}
