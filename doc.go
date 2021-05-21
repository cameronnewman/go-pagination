/*
Package pagination is a simple pagination token generator
*/
package pagination

// Simple example:
//
//	page := Page{
//		OffsetID: "8fed1b32-3059-4f89-96ec-3af15b3d2bc8",
//		OffsetTimeAtUTC: 1621480176,
//		PageSize: 1,
//	}
//
//	token := page.Encode()
//
//	fmt.Println(token)
//	eyJpZCI6IjhmZWQxYjMyLTMwNTktNGY4OS05NmVjLTNhZjE1YjNkMmJjOCIsIm9wZW5fYXRfdXRjIjoxNjIxNDgwMTc2LCJzaXplIjoxfQ==
//
