package fetcher

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
	"unicode"
)

type FetchErrorKind string

const (
	FetchErrorAPI   FetchErrorKind = "api"
	FetchErrorParse FetchErrorKind = "parse"
)

type FetchError struct {
	Kind   FetchErrorKind
	Ticker string
	Err    error
}

func (e *FetchError) Error() string {
	return fmt.Sprintf("fetch %s error for ticker %s: %v", e.Kind, e.Ticker, e.Err)
}

func (e *FetchError) Unwrap() error {
	return e.Err
}

type PriceAPIFetcher struct {
	baseURL string
	client  *http.Client
}

type priceAPIResponse struct {
	Ticker       string `json:"ticker"`
	CurrentPrice string `json:"currentPrice"`
}

func NewPriceAPIFetcher(baseURL string) *PriceAPIFetcher {
	return &PriceAPIFetcher{
		baseURL: strings.TrimRight(baseURL, "/"),
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (f *PriceAPIFetcher) FetchCurrentPrice(ctx context.Context, ticker string) (float64, error) {
	apiURL := fmt.Sprintf("%s/scrape?ticker=%s", f.baseURL, url.QueryEscape(ticker))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiURL, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := f.client.Do(req)
	if err != nil {
		return 0, &FetchError{Kind: FetchErrorAPI, Ticker: ticker, Err: fmt.Errorf("failed to call price api: %w", err)}
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 512))
		return 0, &FetchError{Kind: FetchErrorAPI, Ticker: ticker, Err: fmt.Errorf("price api returned status %d: %s", resp.StatusCode, strings.TrimSpace(string(body)))}
	}

	var payload priceAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return 0, &FetchError{Kind: FetchErrorAPI, Ticker: ticker, Err: fmt.Errorf("failed to decode price api response: %w", err)}
	}

	price, err := ParsePrice(payload.CurrentPrice)
	if err != nil {
		return 0, &FetchError{Kind: FetchErrorParse, Ticker: ticker, Err: fmt.Errorf("invalid currentPrice: %w", err)}
	}

	return price, nil
}

func ParsePrice(raw string) (float64, error) {
	raw = strings.TrimSpace(raw)
	if raw == "" || raw == "-" || raw == "－" {
		return 0, fmt.Errorf("empty or unavailable price: %q", raw)
	}

	cleaned := strings.Map(func(r rune) rune {
		switch {
		case r == ',' || r == '円' || unicode.IsSpace(r):
			return -1
		default:
			return r
		}
	}, raw)

	if cleaned == "" {
		return 0, fmt.Errorf("price became empty after cleanup: %q", raw)
	}

	price, err := strconv.ParseFloat(cleaned, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse price %q: %w", cleaned, err)
	}

	return price, nil
}
