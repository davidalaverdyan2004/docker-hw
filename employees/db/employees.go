package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrNotFound = errors.New("employee not found")

func CreateTable(ctx context.Context, pool *pgxpool.Pool) error {

	sqlQuery := `CREATE TABLE IF NOT EXISTS employees 
	(id SERIAL PRIMARY KEY,
	full_name VARCHAR(100) NOT NULL,
	position VARCHAR(100) NOT NULL)`

	_, err := pool.Exec(ctx, sqlQuery)
	if err != nil {
		return fmt.Errorf("create employees table: %w", err)
	}

	return nil

}

func AddEmployee(ctx context.Context, pool *pgxpool.Pool, emp Employee) (Employee, error) {

	sqlQuery := `INSERT INTO employees (full_name, position) VALUES ($1, $2) RETURNING id`

	err := pool.QueryRow(ctx, sqlQuery, emp.FullName, emp.Position).Scan(&emp.ID)
	if err != nil {
		return Employee{}, fmt.Errorf("insert employee: %w", err)
	}

	return emp, nil
}

func GetEmployees(ctx context.Context, pool *pgxpool.Pool) ([]Employee, error) {
	sqlQuery := `SELECT id, full_name, position 
	FROM employees
	ORDER by id`

	rows, err := pool.Query(ctx, sqlQuery)
	if err != nil {
		return nil, fmt.Errorf("select employees: %w", err)
	}

	defer rows.Close()

	employees := []Employee{}

	for rows.Next() {
		var emp Employee

		err := rows.Scan(&emp.ID, &emp.FullName, &emp.Position)
		if err != nil {
			return nil, fmt.Errorf("scan employee: %w", err)
		}

		employees = append(employees, emp)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate employees: %w", err)
	}

	return employees, nil

}

func DeleteEmployee(ctx context.Context, pool *pgxpool.Pool, id int) error {

	sqlQuery := `DELETE FROM employees WHERE id = $1`

	tag, err := pool.Exec(ctx, sqlQuery, id)

	if err != nil {
		return fmt.Errorf("delete employee: %w", err)
	}

	if tag.RowsAffected() == 0 {
		return fmt.Errorf("delete employee id=%d: %w", id, ErrNotFound)
	}

	return nil
}
