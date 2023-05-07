package domain

// Action : action domain
type Action struct {
	ActionID     string `json:"action_id"`
	CategoryID   string `json:"category_id"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	StandardTime string `json:"standard_time"`
	ActionType   int    `json:"action_type"`
	// RelatedActionID string `json:"related_action_id"`
	URL   string `json:"url"`
	After string `json:"after"`
}

const (
  // TUTORIAL : initial add action num
	TUTORIAL = 0
	// DAILY : daily action num
	DAILY = 1
	// NOMAL : normal action num
	NOMAL = 2
	// SEASONAL : seasonal action num
	SEASONAL = 3
)
