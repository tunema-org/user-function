package model

// User represents a user.
type User struct {
	ID            int    `json:"id"`
	Username      string `json:"username"`
	Email         string `json:"email"`
	Password      string `json:"-"`
	ProfileImgURL string `json:"profile_img_url"`
}
