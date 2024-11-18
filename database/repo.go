package database

type DBRepo interface {
	SelectAllArticles() ([]Article, error)
	GetSelectedArticles() ([]Article, error)
	DeleteAllArticles() error
	DeleteOneArticle(id int) error
	PushSelectedArticle([]Article) error
	InitTables([]Article) error
}
