package firebase_admin

import (
	"context"

	firebase "firebase.google.com/go/v4"
)

func InitFirebaseAdmin() (*firebase.App, error) {
	app, err := firebase.NewApp(context.Background(), nil)

	if err != nil {
		return nil, err
	}

	return app, nil
}
