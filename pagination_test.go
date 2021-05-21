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
			tr:   "eyJvZmZzZXRfaWQiOiI4ZmVkMWIzMi0zMDU5LTRmODktOTZlYy0zYWYxNWIzZDJiYzgiLCJvZmZzZXRfdGltZV9hdF91dGMiOjE2MjE0ODAxNzYsInBhZ2Vfc2l6ZSI6MX0=",
			want: Page{
				OffsetID:        "8fed1b32-3059-4f89-96ec-3af15b3d2bc8",
				OffsetTimeAtUTC: 1621480176,
				PageSize:        1,
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
		OffsetID        string
		OffsetTimeAtUTC int64
		PageSize        int64
	}
	tests := []struct {
		name   string
		fields fields
		want   Token
	}{
		{
			name: "Pass/Empty",
			fields: fields{
				OffsetID:        "",
				OffsetTimeAtUTC: 0,
				PageSize:        0,
			},
			want: "eyJvZmZzZXRfaWQiOiIiLCJvZmZzZXRfdGltZV9hdF91dGMiOjAsInBhZ2Vfc2l6ZSI6MH0=",
		},
		{
			name: "Pass/UUID",
			fields: fields{
				OffsetID:        "8fed1b32-3059-4f89-96ec-3af15b3d2bc8",
				OffsetTimeAtUTC: 1621480176,
				PageSize:        1,
			},
			want: "eyJvZmZzZXRfaWQiOiI4ZmVkMWIzMi0zMDU5LTRmODktOTZlYy0zYWYxNWIzZDJiYzgiLCJvZmZzZXRfdGltZV9hdF91dGMiOjE2MjE0ODAxNzYsInBhZ2Vfc2l6ZSI6MX0=",
		},
		{
			name: "Pass/Email",
			fields: fields{
				OffsetID:        "larry@google.com",
				OffsetTimeAtUTC: 1621480176,
				PageSize:        1203948203498029384,
			},
			want: "eyJvZmZzZXRfaWQiOiJsYXJyeUBnb29nbGUuY29tIiwib2Zmc2V0X3RpbWVfYXRfdXRjIjoxNjIxNDgwMTc2LCJwYWdlX3NpemUiOjEyMDM5NDgyMDM0OTgwMjkzODR9",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Page{
				OffsetID:        tt.fields.OffsetID,
				OffsetTimeAtUTC: tt.fields.OffsetTimeAtUTC,
				PageSize:        tt.fields.PageSize,
			}
			if got := p.Encode(); got != tt.want {
				t.Errorf("Page.Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}
