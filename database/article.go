package database

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type PostgresRepo struct {
	DB *sql.DB
}

func (m *PostgresRepo) GetSelectedArticles() ([]Article, error) {
	var articles []Article
	ctx := context.Background()
	query := ` SELECT * FROM selected`
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var article Article
		err = rows.Scan(&article.Name)

		if err != nil {
			articles = append(articles, article)
		} else {
			return articles, err
		}

	}
	return articles, nil

}

func (m *PostgresRepo) GetAllArticles() ([]Article, error) {
	var articles []Article
	ctx := context.Background()
	query := ` SELECT * FROM article`
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var article Article
		err = rows.Scan(&article.Name)

		if err != nil {
			articles = append(articles, article)
		} else {
			return articles, err
		}

	}
	return articles, nil

}
func (m *PostgresRepo) InitTables() error {
	ctx := context.Background()
	_, err := m.DB.ExecContext(ctx, `
	CREATE TABLE IF NOT EXISTS article (name VARCHAR); 
	CREATE TABLE IF NOT EXISTS selected (name VARCHAR); 
	`)
	return err
}
