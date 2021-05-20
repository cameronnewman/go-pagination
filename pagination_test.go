package pagination

import (
	"reflect"
	"testing"
)

func TestToken_Decode(t *testing.T) {
	tests := []struct {
		name string
		tr   Token
		want Page
	}{
		{
			name: "Pass/Empty",
			tr:   "",
			want: Page{},
		},
		{
			name: "Pass/RandomString",
			tr:   "askdjhaskdjhaskdh",
			want: Page{},
		},
		{
			name: "Pass/UUID",
			tr:   "eyJpZCI6IjhmZWQxYjMyLTMwNTktNGY4OS05NmVjLTNhZjE1YjNkMmJjOCIsIm9wZW5fYXRfdXRjIjoxNjIxNDgwMTc2LCJzaXplIjoxfQ==",
			want: Page{
				OpenedAtUTC: 1621480176,
				ID:          "8fed1b32-3059-4f89-96ec-3af15b3d2bc8",
				Size:        1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.Decode(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Token.Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPage_Encode(t *testing.T) {
	type fields struct {
		OpenedAtUTC int64
		ID          string
		Size        int64
	}
	tests := []struct {
		name   string
		fields fields
		want   Token
	}{
		{
			name: "Pass/Empty",
			fields: fields{
				OpenedAtUTC: 0,
				ID:          "",
				Size:        0,
			},
			want: "eyJpZCI6IiIsIm9wZW5fYXRfdXRjIjowLCJzaXplIjowfQ==",
		},
		{
			name: "Pass/UUID",
			fields: fields{
				OpenedAtUTC: 1621480176,
				ID:          "8fed1b32-3059-4f89-96ec-3af15b3d2bc8",
				Size:        1,
			},
			want: "eyJpZCI6IjhmZWQxYjMyLTMwNTktNGY4OS05NmVjLTNhZjE1YjNkMmJjOCIsIm9wZW5fYXRfdXRjIjoxNjIxNDgwMTc2LCJzaXplIjoxfQ==",
		},
		{
			name: "Pass/Email",
			fields: fields{
				OpenedAtUTC: 1621480176,
				ID:          "larry@google.com",
				Size:        1203948203498029384,
			},
			want: "eyJpZCI6ImxhcnJ5QGdvb2dsZS5jb20iLCJvcGVuX2F0X3V0YyI6MTYyMTQ4MDE3Niwic2l6ZSI6MTIwMzk0ODIwMzQ5ODAyOTM4NH0=",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Page{
				OpenedAtUTC: tt.fields.OpenedAtUTC,
				ID:          tt.fields.ID,
				Size:        tt.fields.Size,
			}
			if got := p.Encode(); got != tt.want {
				t.Errorf("Page.Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}
