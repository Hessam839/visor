package utils

type Items struct {
	Image   string
	Name    string
	Command string
	Created string
	Status  string
	Ports   string
	Paused  bool
	Exited  bool
	Size    string
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
	Item []Stats
}

type Search struct {
	Name        string
	Description string
	IsOfficial  bool
	Stars       int
}

type ViewSearch struct {
	Name  string
	Items []Search
}
