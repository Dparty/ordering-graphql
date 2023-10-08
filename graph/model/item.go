package model

type Item struct {
	ID         string       `json:"id"`
	Name       string       `json:"name"`
	Pricing    int          `json:"pricing"`
	Attributes []*Attribute `json:"attributes"`
	Images     []string     `json:"images"`
	Printers   []string     `json:"printers"`
	Restaurant *Restaurant  `json:"restaurant"`
}

func (i Item) Owner() User {
	return User{
		Name: "owner",
	}
}
