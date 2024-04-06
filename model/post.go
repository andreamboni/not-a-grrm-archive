package model

type Post struct {
	ID      int64
	Image   string
	Title   string
	URL     string
	Date    string
	Content Content
}
