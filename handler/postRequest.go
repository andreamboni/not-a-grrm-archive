package handler

type CreatePostRequest struct {
	Image   string
	Title   string
	URL     string
	Date    string
	Content []string
}
