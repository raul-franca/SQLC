package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/raul-franca/SQLC/internal/db"
)

type CourseDB struct {
	dbConn *sql.DB
	*db.Queries
}

func NewCourseDB(dbConn *sql.DB) *CourseDB {
	return &CourseDB{
		dbConn:  dbConn,
		Queries: db.New(dbConn),
	}
}

// recebe uma funcao a execulta como param se tudo der certo da commit senao rollback
func (c *CourseDB) callTx(ctx context.Context, fn func(*db.Queries) error) error {
	//inicia um Transação "a abri um commit"
	tx, err := c.dbConn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	//injetata o commit a query
	q := db.New(tx)
	// execulta a função com o commit
	err = fn(q)
	if err != nil {
		if errRb := tx.Rollback(); errRb != nil {
			return fmt.Errorf("error on rollback: %v, original error: %w", errRb, err)
		}
		return err
	}
	return tx.Commit()
}
func (c *CourseDB) CreateCourseAndCategory(ctx context.Context, argsCategory CategoryParams, argsCourse CourseParams) error {
	err := c.callTx(ctx, func(q *db.Queries) error {
		var err error
		err = q.CreateCategory(ctx, db.CreateCategoryParams{
			ID:          argsCategory.ID,
			Name:        argsCategory.Name,
			Description: argsCategory.Description,
		})
		if err != nil {
			return err
		}
		err = q.CreateCourse(ctx, db.CreateCourseParams{
			ID:          argsCourse.ID,
			Name:        argsCourse.Name,
			Description: argsCourse.Description,
			CategoryID:  argsCategory.ID,
			Price:       argsCourse.Price,
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

type CourseParams struct {
	ID          string
	Name        string
	Description sql.NullString
	Price       float64
}

type CategoryParams struct {
	ID          string
	Name        string
	Description sql.NullString
}

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	//queries := db.New(dbConn)

	//courses, err := queries.ListCourses(ctx)
	//if err != nil {
	//	panic(err)
	//}
	//for _, course := range courses {
	//	fmt.Printf("Category: %s, Course ID: %s, Course Name: %s, Course Description: %s, Course Price: %f",
	//		course.CategoryName, course.ID, course.Name, course.Description.String, course.Price)
	//}

	courseArgs := CourseParams{
		ID:          uuid.New().String(),
		Name:        "Go",
		Description: sql.NullString{String: "Go Course", Valid: true},
		Price:       10.95,
	}

	categoryArgs := CategoryParams{
		ID:          uuid.New().String(),
		Name:        "Backend",
		Description: sql.NullString{String: "Backend Course", Valid: true},
	}

	courseDB := NewCourseDB(dbConn)
	err = courseDB.CreateCourseAndCategory(ctx, categoryArgs, courseArgs)
	if err != nil {
		panic(err)
	}
}
