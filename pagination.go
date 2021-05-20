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
	ID          string `json:"id"`
	OpenedAtUTC int64  `json:"open_at_utc"`
	Size        int64  `json:"size"`
}

// Encode returns the Token type for a page
func (p Page) Encode() Token {
	b, err := json.Marshal(p)
	if err != nil {
		return Token("")
	}
	return Token(base64.StdEncoding.EncodeToString(b))
}