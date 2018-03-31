package grabutils

// SampleToken -- this works, but not sure how long
var SampleToken = `
{"access_token": "1732638360344261|1sDfyp0c-FjRqTs0JezG8XV-3D4",
 "groups":[{"group_id":"335787510131917","group_name":"Cheltenham Township Residents"}]}

`

// Group -- need to add comment
type Group struct {
	GroupId   string `json:"group_id"`
	GroupName string `json:"group_name"`
}

// Token -- need to add comment
type Token struct {
	AccessToken   string  `json:"access_token"`
	FacebookGroup []Group `json:"groups"`
}

// MemberData -- need to add detail
type MemberData struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// CursorData -- need to add detail
type CursorData struct {
	Before string `json:"before"`
	After  string `json:"after"`
}

// Paging -- need to add detail
type Paging struct {
	Cursors CursorData `json:"cursors"`
	Next    string     `json:"next"`
}

// FacebookMembers -- members in the group
type FacebookMembers struct {
	Data []MemberData `json:"data"`
	Page Paging       `json:"paging"`
}
