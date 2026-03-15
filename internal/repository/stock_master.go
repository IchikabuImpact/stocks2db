package repository

import (
	"context"
	"database/sql"
	"fmt"
)

type StockMasterRepository struct {
	db *sql.DB
}

func NewStockMasterRepository(db *sql.DB) *StockMasterRepository {
	return &StockMasterRepository{db: db}
}

func (r *StockMasterRepository) ListStockCodes(ctx context.Context) ([]string, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT stock_code FROM stock_master ORDER BY stock_code`)
	if err != nil {
		return nil, fmt.Errorf("failed to query stock_master: %w", err)
	}
	defer rows.Close()

	codes := make([]string, 0)
	for rows.Next() {
		var stockCode string
		if err := rows.Scan(&stockCode); err != nil {
			return nil, fmt.Errorf("failed to scan stock_code: %w", err)
		}
		codes = append(codes, stockCode)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed while iterating stock_master rows: %w", err)
	}

	return codes, nil
}
