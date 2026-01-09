package storage

import (
	"context"
	"taskapi/internal/models"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresStore struct {
	db *pgxpool.Pool
}

func NewPostgresStore(databaseURL string) (*PostgresStore, error) {
	ctx := context.Background()

	pool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		return nil, err
	}

	// Verify connection
	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	return &PostgresStore{db: pool}, nil
}

func (s *PostgresStore) Close(){
	s.db.Close()
}

func (s *PostgresStore) Create(task *models.Task) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		INSERT INTO tasks (title, description, completed, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, updated_at
	`

	now := time.Now()
	err := s.db.QueryRow(
		ctx,
		query,
		task.Title,
		task.Description,
		task.Completed,
		now,
		now,
	).Scan(&task.ID, &task.CreatedAt, &task.UpdatedAt)

	return err
}

func (s *PostgresStore) GetByID(id int) (*models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
	SELECT id, title, description, completed, created_at, updated_at
	FROM tasks
	WHERE id = $1
	`

	task := &models.Task{}
	err := s.db.QueryRow(ctx, query, id).Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Completed,
		&task.CreatedAt,
		&task.UpdatedAt,
	)
	
	if err != nil {
		return nil, models.ErrTaskNotFound
	}

	return task, nil
}

func (s *PostgresStore) GetAll() ([]*models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		SELECT id, title, description, completed, created_at, updated_at
		FROM tasks
		ORDER BY created_at DESC
	`
	
	rows, err := s.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*models.Task
	for rows.Next() {
		task := &models.Task{}
		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Completed,
			&task.CreatedAt,
			&task.UpdatedAt,			
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)		
	}

		if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *PostgresStore) Update(id int, updates *models.UpdateTaskRequest) (*models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// First, verify that it exists
	task, err := s.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Build query dinamically according to fields to update
	query := `
		UPDATE tasks
		SET title = COALESCE($1, title),
		    description = COALESCE($2, description),
		    completed = COALESCE($3, completed),
		    updated_at = $4
		WHERE id = $5
		RETURNING id, title, description, completed, created_at, updated_at
	`

	var title, description *string
	var completed *bool


	if updates.Title != nil {
		title = updates.Title		
	}
	if updates.Description != nil {
		description = updates.Description
	}
	if updates.Completed != nil {
		completed = updates.Completed
	}

	err = s.db.QueryRow(
		ctx,
		query,
		title,
		description,
		completed,
		time.Now(),
		id,		
	).Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Completed,
		&task.CreatedAt,
		&task.UpdatedAt,		
	)

	if err != nil {
		return nil, err
	}

	return task, nil
}

func (s *PostgresStore) Delete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `DELETE FROM tasks WHERE id = $1`

	result, err := s.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if (result.RowsAffected() == 0) {
		return models.ErrTaskNotFound
	}

	return nil
}