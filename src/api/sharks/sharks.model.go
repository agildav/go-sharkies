package sharks

// Shark table
type Shark struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Bname       string `json:"bname"`
	Description string `json:"description,omitempty"`
	Image       string `json:"image,omitempty"`
}
