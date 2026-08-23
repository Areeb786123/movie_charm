package repository

import (
	"backend/repository"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupMockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock, func()) {

	sqlDB, mock, err := sqlmock.New()

	if err != nil {
		t.Fatal(err)
	}

	db, err := gorm.Open(
		postgres.New(postgres.Config{
			Conn: sqlDB,
		}),
		&gorm.Config{},
	)

	if err != nil {
		t.Fatal(err)
	}

	cleanup := func() {
		sqlDB.Close()
	}

	return db, mock, cleanup
}

func TestGetAllComment(t *testing.T) {

	// Arrange
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()

	repo := &repository.MovieRep{
		Db: db,
	}

	rows := sqlmock.NewRows([]string{
		"comment_id",
		"movie_id",
		"comment",
	}).
		AddRow(1, 10, "Great movie").
		AddRow(2, 10, "Nice movie")

	mock.ExpectQuery(
		`SELECT .*FROM.*comments`,
	).
		WithArgs(10, 10).
		WillReturnRows(rows)

	// Act
	comments, err := repo.GetAllComment(10, 0, 10)

	// Assert
	if err != nil {
		t.Fatal(err)
	}

	if len(comments) != 2 {
		t.Fatalf(
			"expected 2 comments, got %d",
			len(comments),
		)
	}

	if comments[0].Comment != "Great movie" {
		t.Errorf(
			"expected Great movie, got %s",
			comments[0].Comment,
		)
	}

	if comments[1].Comment != "Nice movie" {
		t.Errorf(
			"expected Nice movie, got %s",
			comments[1].Comment,
		)

	}

	// Make sure expected SQL was executed
	err = mock.ExpectationsWereMet()

	if err != nil {
		t.Fatal(err)
	}
}

func TestGetAllComment_Error(t *testing.T) {

	// Arrange
	db, mock, cleanup := setupMockDB(t)
	defer cleanup()

	repo := &repository.MovieRep{
		Db: db,
	}

	mock.ExpectQuery(
		`SELECT .*FROM.*comments`,
	).
		WithArgs(10, 10).
		WillReturnError(
			errors.New("database error"),
		)

	// Act
	comments, err := repo.GetAllComment(10, 0, 10)

	// Assert
	if err == nil {
		t.Fatal("expected error but got nil")
	}

	if comments != nil {
		t.Fatal("expected comments to be nil")
	}

	err = mock.ExpectationsWereMet()

	if err != nil {
		t.Fatal(err)
	}
}
