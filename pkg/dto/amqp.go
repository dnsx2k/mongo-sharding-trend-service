package dto

type TrendOperation struct {
	Direction string   `json:"direction"`
	Domain    string   `json:"domain,omitempty"`
	IDs       []string `json:"IDs"`
}
