package models

// CommentForm is the form for the request CreateArticle
type CommentForm struct {
	Username  string `json:"username"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	ArticleID string `json:"article_id"`
}
