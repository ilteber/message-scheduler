#!/bin/bash

# Database initialization script
# This script initializes the database with schema and seed data

set -e

# Configuration
DB_HOST="${DB_HOST:-localhost}"
DB_PORT="${DB_PORT:-5432}"
DB_USER="${DB_USER:-insider}"
DB_PASSWORD="${DB_PASSWORD:-insider123}"
DB_NAME="${DB_NAME:-insider_db}"

echo "Initializing database..."
echo "Host: $DB_HOST:$DB_PORT"
echo "Database: $DB_NAME"
echo "User: $DB_USER"
echo ""

# Wait for PostgreSQL to be ready
echo "Waiting for PostgreSQL to be ready..."
until PGPASSWORD=$DB_PASSWORD psql -h "$DB_HOST" -U "$DB_USER" -d "$DB_NAME" -c '\q' 2>/dev/null; do
  echo "PostgreSQL is unavailable - sleeping"
  sleep 2
done

echo "PostgreSQL is ready!"
echo ""

# Run schema
echo "Creating schema..."
PGPASSWORD=$DB_PASSWORD psql -h "$DB_HOST" -U "$DB_USER" -d "$DB_NAME" -f /scripts/schema.sql
echo "✓ Schema created successfully"
echo ""

# Run seed data
echo "Inserting seed data..."
PGPASSWORD=$DB_PASSWORD psql -h "$DB_HOST" -U "$DB_USER" -d "$DB_NAME" -f /scripts/seed.sql
echo "✓ Seed data inserted successfully"
echo ""

# Verify data
echo "Verifying data..."
RECORD_COUNT=$(PGPASSWORD=$DB_PASSWORD psql -h "$DB_HOST" -U "$DB_USER" -d "$DB_NAME" -t -c "SELECT COUNT(*) FROM messages;")
echo "✓ Total messages in database: $RECORD_COUNT"
echo ""

echo "Database initialization complete!"

