package domain

// Category : category domain
type Category struct {
	CategoryID string `json:"category_id"`
	Name       string `json:"name"`
	Goal       string `json:"goal"`
}

// CategoryMap : onmemory global elem
var CategoryMap = map[string]Category{}
