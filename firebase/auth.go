package firebase

import (
	"context"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func InitFirebase(projectID string) (*firebase.App, *firestore.Client, error) {
	opt := option.WithCredentialsJSON([]byte(os.Getenv("FIREBASE_CRED")))
	config := &firebase.Config{ProjectID: projectID}
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		return nil, nil, err
	}
	firestore, err := app.Firestore(context.Background())
	return app, firestore, err
}