package models

// Comment represents the comment
type Comment struct {
	ID         int
	Username   string `json:"username"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	CreateDate string `json:"create_date"`
	ArticleID  string `json:"article_id"`
	Likes      []int  `json:"likes"`
	Dislike    []int  `json:"dislike"`
}

// CommentForm is the form for the request CreateArticle
type CommentForm struct {
	Username  string `json:"username"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	ArticleID string `json:"article_id"`
}
