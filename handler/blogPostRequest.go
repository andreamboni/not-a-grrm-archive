package handler

type CreateBlogPostRequest struct {
	PostImage string
	Title     string
	URL       string
	TheDate   string
	Content   []string
}
