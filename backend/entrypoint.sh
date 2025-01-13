#!/bin/bash

set -e

# マイグレーションを実行
echo "Running migrations..."
migrate -path ./migrations -database "mysql://${DB_USER}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}" up

# アプリケーションを起動
echo "Starting the application..."
air