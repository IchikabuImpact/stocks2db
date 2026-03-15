package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type StockPriceDailyRepository struct {
	db   *sql.DB
	stmt *sql.Stmt
}

func NewStockPriceDailyRepository(db *sql.DB) (*StockPriceDailyRepository, error) {
	stmt, err := db.Prepare(`
		INSERT INTO stock_price_daily (trade_date, stock_code, price)
		VALUES (?, ?, ?)
		ON DUPLICATE KEY UPDATE price = VALUES(price)
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare stock_price_daily upsert statement: %w", err)
	}
	return &StockPriceDailyRepository{db: db, stmt: stmt}, nil
}

func (r *StockPriceDailyRepository) Close() error {
	if r.stmt != nil {
		return r.stmt.Close()
	}
	return nil
}

func (r *StockPriceDailyRepository) Upsert(ctx context.Context, tradeDate time.Time, stockCode string, price float64) error {
	if _, err := r.stmt.ExecContext(ctx, tradeDate, stockCode, price); err != nil {
		return fmt.Errorf("failed to upsert stock_price_daily for %s: %w", stockCode, err)
	}
	return nil
}
