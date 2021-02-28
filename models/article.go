package models

// Article represent the article
type Article struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt string    `json:"created_at"`
	Comments  []Comment `json:"comments"`
	Likes     []int     `json:"likes"`
	Dislike   []int     `json:"dislike"`
}
