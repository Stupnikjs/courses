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
	query := ` SELECT * FROM selected; `
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
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
	query := ` SELECT name FROM article; `
	rows, err := m.DB.QueryContext(ctx, query)

	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var text string
		var article Article
		err = rows.Scan(&text)

		if err == nil {
			article.Name = text
			articles = append(articles, article)
		} else {
			return articles, err
		}

	}
	defer rows.Close()

	return articles, nil

}

func (m *PostgresRepo) PushSelectedArticles(articles []Article) error {
	fmt.Println(articles)
	ctx := context.Background()
	query := `INSERT INTO selected (name) VALUES ($1); `
	for _, a := range articles {
		_, err := m.DB.ExecContext(ctx, query, a.Name)
		if err != nil {
			return err
		}
	}
	return nil

}
func (m *PostgresRepo) InitTables(articles []Article) error {
	ctx := context.Background()
	_, err := m.DB.ExecContext(ctx, `
	CREATE TABLE IF NOT EXISTS article (name VARCHAR UNIQUE); 
	CREATE TABLE IF NOT EXISTS selected (name VARCHAR UNIQUE); 
	`)
	if err != nil {
		return err
	}

	query := `INSERT INTO article (name) VALUES ($1); `
	for _, a := range articles {
		_, err := m.DB.ExecContext(ctx, query, a.Name)
		if err != nil {
			return err
		}
	}
	return nil
}
