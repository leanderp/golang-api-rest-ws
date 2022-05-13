package database

import (
	"context"
	"database/sql"
	"log"

	"github.com/leanderp/golang_rest_web/models"
	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{db}, nil
}

func (r *PostgresRepository) InsertUser(ctx context.Context, user *models.User) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO users (id, email, password) VALUES ($1, $2, $3)", user.Id, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgresRepository) GetUserById(ctx context.Context, id string) (*models.User, error) {
	var user models.User
	rows, err := r.db.QueryContext(ctx, "SELECT id, email FROM users WHERE id = $1", id)

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	for rows.Next() {
		if err = rows.Scan(&user.Id, &user.Email); err == nil {
			return &user, nil
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *PostgresRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	rows, err := r.db.QueryContext(ctx, "SELECT id, email, password FROM users WHERE email = $1", email)

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	for rows.Next() {
		if err = rows.Scan(&user.Id, &user.Email, &user.Password); err == nil {
			return &user, nil
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *PostgresRepository) InsertPost(ctx context.Context, post *models.Post) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO posts (id, post_content, user_id) VALUES ($1, $2, $3)", post.Id, post.PostContent, post.UserId)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgresRepository) GetPostById(ctx context.Context, id string) (*models.Post, error) {
	var post models.Post

	rows, err := r.db.QueryContext(ctx, "SELECT id, post_content, created_at, updated_at, user_id FROM posts WHERE id = $1", id)

	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	for rows.Next() {
		if err = rows.Scan(&post.Id, &post.PostContent, &post.CreatedAt, &post.UpdatedPost, &post.UserId); err == nil {
			return &post, nil
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &post, nil
}

func (r *PostgresRepository) UpdatePost(ctx context.Context, post *models.Post) error {
	_, err := r.db.ExecContext(ctx, "UPDATE posts SET post_content = $1, updated_at = $2 WHERE id = $3 and user_id = $4", post.PostContent, post.UpdatedPost, post.Id, post.UserId)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgresRepository) DeletePost(ctx context.Context, id string, userId string) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM posts WHERE id = $1 and user_id = $2", id, userId)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgresRepository) ListPosts(ctx context.Context, page, quantity uint64) ([]*models.Post, error) {
	var posts []*models.Post
	rows, err := r.db.QueryContext(ctx, "SELECT id, post_content, created_at, updated_at, user_id FROM posts ORDER BY created_at DESC LIMIT $1 OFFSET $2", quantity, page*quantity)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	for rows.Next() {
		var post models.Post
		if err = rows.Scan(&post.Id, &post.PostContent, &post.CreatedAt, &post.UpdatedPost, &post.UserId); err == nil {
			posts = append(posts, &post)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *PostgresRepository) Close() error {
	return r.db.Close()
}
