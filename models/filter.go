package models

type Filter struct {
	Search string `json:"search"`
	SortBy string `json:"sortBy"`
	Sort   string `json:"sort"`
	Limit  string `json:"limit"`
}
