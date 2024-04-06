package model

type BlogPost struct {
	ID        int64
	PostImage string
	Title     string
	URL       string
	TheDate   string
	Content   []string
}
