package domain

type Node struct {
	ID    string  `json:"id"`
	Left  *string `json:"left"`
	Right *string `json:"right"`
	Value int     `json:"value"`
}
