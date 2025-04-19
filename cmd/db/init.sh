#!/bin/sh
set -e

echo "Running schema.sql to initialize the database schema..."
psql -U "$POSTGRES_USER" -d "$POSTGRES_DB" -f /docker-entrypoint-initdb.d/01-schema.sql

echo "Running Go parser to populate the database with data..."
cd /app
go run cmd/db/init.go

echo "Database initialization complete!"