package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
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

	now := time.Now().In(time.Local)
	tradeDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	priceFetcher := fetcher.NewPriceAPIFetcher(cfg.PriceAPI.BaseURL)
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	successCount := 0
	failureCount := 0

	for i, code := range codes {
		log.Printf("[INFO] stock_code=%s start", code)

		price, err := priceFetcher.FetchCurrentPrice(ctx, code)
		if err != nil {
			failureCount++
			var fetchErr *fetcher.FetchError
			if errors.As(err, &fetchErr) {
				switch fetchErr.Kind {
				case fetcher.FetchErrorAPI:
					log.Printf("[ERROR] stock_code=%s api_fetch_failed err=%v", code, err)
				case fetcher.FetchErrorParse:
					log.Printf("[ERROR] stock_code=%s current_price_parse_failed err=%v", code, err)
				default:
					log.Printf("[ERROR] stock_code=%s fetch_failed err=%v", code, err)
				}
			} else {
				log.Printf("[ERROR] stock_code=%s fetch_failed err=%v", code, err)
			}
		} else {
			log.Printf("[INFO] stock_code=%s api_fetch_success price=%.2f", code, price)

			if err := priceRepo.Upsert(ctx, tradeDate, code, price); err != nil {
				failureCount++
				log.Printf("[ERROR] stock_code=%s db_save_failed err=%v", code, err)
			} else {
				successCount++
				log.Printf("[INFO] stock_code=%s db_save_success trade_date=%s price=%.2f", code, tradeDate.Format("2006-01-02"), price)
			}
		}

		if i < len(codes)-1 {
			sleepMillis := 1000 + rng.Intn(1001)
			sleepDuration := time.Duration(sleepMillis) * time.Millisecond
			log.Printf("[INFO] stock_code=%s sleep_before_next=%s", code, sleepDuration)
			time.Sleep(sleepDuration)
		}
	}

	fmt.Printf("[SUMMARY] total=%d success=%d failure=%d\n", len(codes), successCount, failureCount)

	if failureCount > 0 {
		os.Exit(1)
	}
}
