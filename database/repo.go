package database

type DBRepo interface {
	SelectAllArticles() ([]Article, error)
	GetSelectedArticles() ([]Article, error)
	DeleteAllArticles() error
	PushSelectedArticle([]Article) error
	InitTables([]Article) error
}
