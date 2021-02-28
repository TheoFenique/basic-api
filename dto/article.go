package models

// ArticleForm is the form for the request CreateArticle
type ArticleForm struct {
	Username string `json:"username"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}
