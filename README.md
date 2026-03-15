# stocks2db

## Daily batch 概要

`stock_master` に登録されている全銘柄を **1件ずつ順番に** 処理し、価格API (`GET /scrape?ticker=...`) から `currentPrice` を取得して `stock_price_daily` へ保存するCLIバッチです。

- 保存対象: `trade_date`, `stock_code`, `price`
- `trade_date` はローカル日付の 00:00:00 を明示生成して保存
- 同日・同銘柄の再実行は `INSERT ... ON DUPLICATE KEY UPDATE` により冪等
- API負荷軽減のため、各銘柄処理後に 1〜2 秒（1000〜2000ms）のランダム待機を実施
- 失敗銘柄があっても他銘柄の処理を継続
- 最後に成功件数・失敗件数・総件数を表示

## 設定ファイル

`config.json` をプロジェクトルートに配置してください。`config.json.sample` をコピーして使えます。

```json
{
  "db": {
    "host": "127.0.0.1",
    "port": 3306,
    "name": "stocks2db",
    "user": "stocksapp",
    "password": "your-db-password"
  },
  "priceApi": {
    "baseUrl": "http://127.0.0.1:8085"
  }
}
```

- `priceApi.baseUrl` は価格取得APIのベースURLです（例: `http://127.0.0.1:8085`）。

## 実行方法

### スクリプトから実行

```bash
./scripts/run_daily_batch.sh
```

### 直接実行

```bash
go run ./cmd/dailybatch
```

## 実行ログ例

```text
[INFO] stock_code=7203 start
[INFO] stock_code=7203 api_fetch_success price=1364.00
[INFO] stock_code=7203 db_save_success trade_date=2026-01-10 price=1364.00
[INFO] stock_code=7203 sleep_before_next=1.742s
[SUMMARY] total=100 success=99 failure=1
```

失敗時は `api_fetch_failed` / `current_price_parse_failed` / `db_save_failed` のログを確認してください。

## cron 設定例

平日 15:30 に実行する例:

```cron
30 15 * * 1-5 cd /path/to/stocks2db && ./scripts/run_daily_batch.sh >> ./dailybatch.log 2>&1
```
