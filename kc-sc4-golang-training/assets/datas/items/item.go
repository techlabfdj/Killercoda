package items

import (
	"encoding/json"
	"time"
)

// Envelope contains envelope information about the data
type Envelope struct {
	ID         string    `json:"id" binding:"omitempty,uuid4"`
	CreatedAt  time.Time `json:"created_at" time_format:"2006-01-02T15:04:05.999Z07:00"`
	ModifiedAt time.Time `json:"modified_at" time_format:"2006-01-02T15:04:05.999Z07:00"`
}

// Data is used to hold additional data
type Data map[string]interface{}

// Item is the holder for an item handled by system
type Item struct {
	Envelope
	Data
}

// UnmarshalJSON  converts JSON to object
func (item *Item) UnmarshalJSON(bs []byte) (err error) {
	readEnvelope := Envelope{}
	if err = json.Unmarshal(bs, &readEnvelope); err == nil {
		readData := make(map[string]interface{})
		if err = json.Unmarshal(bs, &readData); err == nil {
			envelopeProperties := [3]string{"id", "created_at", "modified_at"}
			for _, property := range envelopeProperties {
				_, ok := readData[property]
				if ok {
					delete(readData, property)
				}
			}
			item.Envelope = readEnvelope
			item.Data = readData
		}
	}
	return err
}

// MarshalJSON converts object to JSON
func (item Item) MarshalJSON() ([]byte, error) {
	eb, err := json.Marshal(item.Envelope)
	if err != nil {
		return nil, err
	}
	if item.Data == nil || len(item.Data) == 0 {
		return eb, nil
	}
	db, err := json.Marshal(item.Data)
	if err != nil {
		return nil, err
	}
	if len(db) == 2 {
		return eb, nil
	}
	eb[len(eb)-1] = ','
	return append(eb, db[1:]...), nil
}
