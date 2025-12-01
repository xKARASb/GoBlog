package types

type Role string
type ContextKey string
type PostStatus string

const (
	Author    Role       = "author"
	Reader    Role       = "reader"
	CtxUser   ContextKey = "user"
	Draft     PostStatus = "draft"
	Published PostStatus = "published"
)
