#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
CONFIG_FILE="$SCRIPT_DIR/config.json"
SCHEMA_FILE="$SCRIPT_DIR/schema.sql"

read_config() {
  python3 - <<'PY' "$CONFIG_FILE"
import json, sys
with open(sys.argv[1], "r", encoding="utf-8") as f:
    cfg = json.load(f)["db"]
print(cfg["host"])
print(cfg["port"])
print(cfg["name"])
print(cfg["user"])
print(cfg["password"])
PY
}

mapfile -t DBCONF < <(read_config)

DB_HOST="${DBCONF[0]}"
DB_PORT="${DBCONF[1]}"
DB_NAME="${DBCONF[2]}"
DB_USER="${DBCONF[3]}"
DB_PASSWORD="${DBCONF[4]}"

MYSQL_PWD="$DB_PASSWORD" mysqldump \
  -h "$DB_HOST" \
  -P "$DB_PORT" \
  -u "$DB_USER" \
  --single-transaction \
  --skip-lock-tables \
  --no-data \
  "$DB_NAME" \
  stock_master stock_price_daily > "$SCHEMA_FILE"

echo "schema exported: $SCHEMA_FILE"

"$SCRIPT_DIR/exportdata.sh"
