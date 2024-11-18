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
	query := ` SELECT id, name FROM selected; `
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var article Article
		err = rows.Scan(&article.Id, &article.Name)

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
	query := ` SELECT id, name FROM article; `
	rows, err := m.DB.QueryContext(ctx, query)

	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var article Article
		err = rows.Scan(&article.Id, &article.Name)

		if err == nil {
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
	CREATE TABLE IF NOT EXISTS article (id SERIAL PRIMARY KEY, name VARCHAR); 
	CREATE TABLE IF NOT EXISTS selected (id SERIAL PRIMARY KEY, name VARCHAR); 
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
func (m *PostgresRepo) DeleteOneArticle(Id int) error {
	ctx := context.Background()
	_, err := m.DB.ExecContext(ctx, `DELETE article WHERE id = $1;`, Id)
	if err != nil {
		return err
	}
	return nil
}
func (m *PostgresRepo) InsertOneArticle(Name string) error {
	ctx := context.Background()
	_, err := m.DB.ExecContext(ctx, `INSERT INTO article (name) VALUES ($1)`, Name)
	if err != nil {
		return err
	}
	return nil
}
