package database

type DBRepo interface {
	SelectAllArticles() ([]Article, error)
	GetSelectedArticles() ([]Article, error)
	DeleteAllArticles() error
	SelectSomeArticles([]Article) error
	InitTables() error
}
