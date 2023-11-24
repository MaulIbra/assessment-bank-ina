package models

type Wrapper struct {
	Data     interface{} `json:"data"`
	MetaData *MetaData   `json:"meta_data,omitempty"`
}

type MetaData struct {
	Token       string `json:"token,omitempty"`
	CurrentPage int    `json:"current_page,omitempty"`
	FirstPage   int    `json:"first_page,omitempty"`
	LastPage    int    `json:"last_page,omitempty"`
	TotalData   int    `json:"total_data,omitempty"`
}
