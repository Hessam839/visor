package server

type Items struct {
	Name    string
	Created string
	Size    string
	Status  string
	ID      string
}
type ViewData struct {
	Name  string
	Items []Items
}

type Stats struct {
	Name    string
	ID      string
	CPU     string
	RAM     string
	Network string
}

type ViewStat struct {
	Name string
	Item Stats
}
