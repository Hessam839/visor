package structs

type Container struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Tags    string `json:"tags"`
	Created int64  `json:"created"`
	Status  string `json:"status"`
}
