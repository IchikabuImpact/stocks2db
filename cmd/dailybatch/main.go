package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"stocks2db/internal/config"
	"stocks2db/internal/db"
	"stocks2db/internal/fetcher"
	"stocks2db/internal/repository"
)

func main() {
	ctx := context.Background()

	cfg, err := config.Load("config.json")
	if err != nil {
		log.Fatalf("[FATAL] %v", err)
	}

	database, err := db.NewMySQL(cfg.DB)
	if err != nil {
		log.Fatalf("[FATAL] %v", err)
	}
	defer database.Close()

	masterRepo := repository.NewStockMasterRepository(database)
	priceRepo, err := repository.NewStockPriceDailyRepository(database)
	if err != nil {
		log.Fatalf("[FATAL] %v", err)
	}
	defer priceRepo.Close()

	codes, err := masterRepo.ListStockCodes(ctx)
	if err != nil {
		log.Fatalf("[FATAL] %v", err)
	}

	if len(codes) == 0 {
		log.Println("[INFO] no stocks found in stock_master")
		os.Exit(0)
	}

	tradeDate := time.Now().Local().Truncate(24 * time.Hour)
	fetcher := fetcher.NewPriceAPIFetcher(cfg.PriceAPI.BaseURL)

	successCount := 0
	failureCount := 0

	for _, code := range codes {
		price, err := fetcher.FetchCurrentPrice(ctx, code)
		if err != nil {
			failureCount++
			log.Printf("[ERROR] stock_code=%s failed to fetch price: %v", code, err)
			continue
		}

		if err := priceRepo.Upsert(ctx, tradeDate, code, price); err != nil {
			failureCount++
			log.Printf("[ERROR] stock_code=%s failed to save price: %v", code, err)
			continue
		}

		successCount++
		log.Printf("[INFO] stock_code=%s trade_date=%s price=%.2f saved", code, tradeDate.Format("2006-01-02"), price)
	}

	fmt.Printf("[SUMMARY] total=%d success=%d failure=%d\n", len(codes), successCount, failureCount)

	if successCount == 0 && failureCount > 0 {
		os.Exit(1)
	}
}
