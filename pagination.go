package pagination

import (
	"encoding/base64"
	"encoding/json"
)

// Token represents the encoded pagination
type Token string

// Decode returns the Page for a given token
func (t Token) Decode() Page {
	var result Page
	if len(t) == 0 {
		return result
	}

	bytes, err := base64.StdEncoding.DecodeString(string(t))
	if err != nil {
		return result
	}

	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return result
	}

	return result
}

// Page represents the next page for the query
type Page struct {
	OffsetID        string `json:"offset_id"`
	OffsetTimeAtUTC int64  `json:"offset_time_at_utc"`
	PageSize        int64  `json:"page_size"`
}

// Encode returns the Token type for a page
func (p Page) Encode() Token {
	b, err := json.Marshal(p)
	if err != nil {
		return Token("")
	}
	return Token(base64.StdEncoding.EncodeToString(b))
}
