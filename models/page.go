package models

type PaginationMeta struct {
	Total int
	Limit int
	Page  int
	Next  string
	Prev  string
}

type AfterOptions struct {
	ID string
}

type BeforeOptions struct {
	ID string
}

type PageOptions struct {
	Q      string
	Limit  int
	Page   int
	Cursor string
}

type GetPageOptionsConfig struct {
	DefaultLimit int
}
