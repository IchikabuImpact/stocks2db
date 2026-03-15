# stocks2db

## Daily batch 概要

`stock_master` に登録されている全銘柄を1件ずつ処理し、価格API (`GET /scrape?ticker=...`) から `currentPrice` を取得して `stock_price_daily` へ保存するCLIバッチです。

- 保存対象: `trade_date`, `stock_code`, `price`
- `trade_date` は実行日のシステム日付
- 同日・同銘柄の再実行は `INSERT ... ON DUPLICATE KEY UPDATE` により冪等
- 失敗銘柄があっても他銘柄の処理を継続
- 最後に成功件数・失敗件数を標準出力へ表示

## 設定ファイル

`config.json` をプロジェクトルートに配置してください。`config.json.sample` をコピーして使えます。

```json
{
  "db": {
    "host": "127.0.0.1",
    "port": 3306,
    "name": "stocks2db",
    "user": "stocksapp",
    "password": "331155"
  },
  "priceApi": {
    "baseUrl": "http://127.0.0.1:8085"
  }
}
```

## 実行方法

### スクリプトから実行

```bash
./scripts/run_daily_batch.sh
```

### 直接実行

```bash
go run ./cmd/dailybatch
```

## cron 設定例

平日 15:30 に実行する例:

```cron
30 15 * * 1-5 cd /path/to/stocks2db && ./scripts/run_daily_batch.sh >> ./dailybatch.log 2>&1
```
