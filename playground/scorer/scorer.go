package scorer

import (
	"context"
	"log"
	"strconv"
	"time"

	firestore "cloud.google.com/go/firestore"
)

// FirestoreEvent is the payload of a Firestore event.
type FirestoreEvent struct {
	OldValue   FirestoreValue `json:"oldValue"`
	Value      FirestoreValue `json:"value"`
	UpdateMask struct {
		FieldPaths []string `json:"fieldPaths"`
	} `json:"updateMask"`
}

// FirestoreValue holds Firestore fields.
type FirestoreValue struct {
	CreateTime time.Time `json:"createTime"`
	// Fields is the data for this value. The type depends on the format of your
	// database. Log an interface{} value and inspect the result to see a JSON
	// representation of your database fields.
	Fields     Review    `json:"fields"`
	Name       string    `json:"name"`
	UpdateTime time.Time `json:"updateTime"`
}

// Review represents the Firestore schema of a movie review.
type Review struct {
	Author struct {
		Value string `json:"stringValue"`
	} `json:"author"`
	Text struct {
		Value string `json:"stringValue"`
	} `json:"text"`
}

var client *firestore.Client

func init() {
	ctx := context.Background()
	// conf := &firebase.Config{
	// 	DatabaseURL: "https://crystal-factory.firebaseio.com/",
	// }
	// app, err := firebase.NewApp(ctx, conf)
	// if err != nil {
	// 	log.Fatalf("firebase.NewApp: %v", err)
	// }
	// client, err = app.Database(ctx)
	// if err != nil {
	// 	log.Fatalf("app.Firestore: %v", err)
	// }
	var err error
	client, err = firestore.NewClient(ctx, "crystal-factory")
	if err != nil {
		log.Fatalf("app.Firestore: %v", err)
	}
}

// MovieReview is runtime store repr object
type MovieReview struct {
	Author string `firestore:"author"`
	Text   string `firestore:"text"`
}

// ScoreReview generates the scores for movie reviews and transactionally writes them to the
// Firebase Realtime Database.
func ScoreReview(ctx context.Context, e FirestoreEvent) error {
	review := e.Value.Fields
	reviweScore := score(review.Text.Value)

	// ref := client.NewRef("scores").Child(review.Author.Value)
	// updateTxn := func(node db.TransactionNode) (interface{}, error) {
	// 	var currentScore int
	// 	if err := node.Unmarshal(&currentScore); err != nil {
	// 		return nil, err
	// 	}
	// 	return currentScore + reviweScore, nil
	// }
	// return ref.Transaction(ctx, updateTxn)

	ref := client.Collection("scores").Doc(review.Author.Value)
	var err error
	_, err = ref.Set(ctx, MovieReview{Author: review.Author.Value, Text: "score: " + strconv.Itoa(reviweScore)})
	// var err error
	// _, err = ref.Update(ctx, []firestore.Update{{Value: "Sacramento-" + strconv.Itoa(reviweScore)}})
	// return err
	// var movie MovieReview
	// if err := ref.Get(ctx).Data().DataT
	return err
}

// score computes the score for a review text.
//
// This is currently just the length of the text.
func score(text string) int {
	return len(text)
}
