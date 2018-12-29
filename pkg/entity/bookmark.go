package entity

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type ID bson.ObjectId

type Bookmark struct {
	ID          ID        `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Link        string    `json:"link"`
	Tags        []string  `json:"tags"`
	Favorite    bool      `json:"favorite"`
	CreatedAt   time.Time `json:"created_at"`
}

// ToString Convert an ID in a string
func (i ID) String() string {
	return bson.ObjectId(i).Hex()
}

// MarshalJSON will marshal ID to JSON
func (i ID) MarshalJSON() ([]byte, error) {
	return bson.ObjectId(i).MarshalJSON()
}

func (i *ID) unMarshalJSON(data []byte) error {
	s := string(data)
	s = s[1 : len(s)-1]
	if bson.IsObjectIdHex(s) {
		*i = ID(bson.ObjectIdHex(s))
	}
	return nil
}

// GetBSON implements bson.Getter
func (i ID) GetBSON() (interface{}, error) {
	if i == "" {
		return "", nil
	}
	return bson.ObjectId(i), nil
}
