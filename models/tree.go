package domain

type Tree struct {
	Nodes []Node `json:"nodes"`
	Root  string `json:"root"`
}
