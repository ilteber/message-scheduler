# Message Scheduler

An automatic message sending system built with Go that processes and sends messages from a database to a webhook endpoint at configurable intervals.

## Features

- Automatic message sending: 2 messages every 2 minutes (configurable)
- Custom scheduler implementation without external cron packages
- RESTful API for scheduler control and message retrieval
- Redis caching for sent message metadata
- PostgreSQL database with status tracking
- Standard library only (net/http, database/sql)
- Docker Compose setup with all dependencies
- Comprehensive seed data for testing
- Complete documentation (README, Swagger, Configuration guide)

## Quick Start

### Prerequisites

- Docker and Docker Compose installed
- Ports 8080, 5432, and 6379 available

### Installation

**Important:** Before starting, configure your webhook URL (see `CONFIGURATION.md` or Configuration section below).

1. Clone the repository:
```bash
git clone https://github.com/ilteber/message-scheduler.git
cd message-scheduler
```

2. Configure webhook URL:
   - Visit https://webhook.site and copy your unique URL
   - Edit `docker-compose.yml` line 80 with your webhook URL
   - OR set environment variable: `export WEBHOOK_URL="https://webhook.site/YOUR-ID"`

3. Start all services:
```bash
docker-compose up -d
```

This automatically:
- Starts PostgreSQL and Redis
- Creates database schema
- Inserts seed data (122 messages across 6 phone numbers)
- Starts the API server with scheduler enabled

4. Verify the system is running:
```bash
curl http://localhost:8080/health
```

5. Access API documentation:
```
Swagger UI: http://localhost:8081
API: http://localhost:8080
```

### Configuration

#### Setup Your Webhook URL (Required)

Before starting the system, configure your own webhook endpoint:

1. **Get your webhook URL:**
   - Visit https://webhook.site
   - Copy your unique URL (e.g., `https://webhook.site/YOUR-UNIQUE-ID`)

2. **Configure the system:**
   - Open `docker-compose.yml`
   - Find line 80: `WEBHOOK_URL: ${WEBHOOK_URL:-https://webhook.site/...}`
   - Replace the default URL with your webhook URL:
     ```yaml
     WEBHOOK_URL: ${WEBHOOK_URL:-https://webhook.site/YOUR-UNIQUE-ID}
     ```

3. **Alternative: Use environment variable**
   ```bash
   export WEBHOOK_URL="https://webhook.site/YOUR-UNIQUE-ID"
   docker-compose up -d
   ```

#### Configuration Variables

Configure via environment variables in `docker-compose.yml`:

| Variable | Description | Default | Location |
|----------|-------------|---------|----------|
| `WEBHOOK_URL` | Webhook endpoint | https://webhook.site/... | Line 80 |
| `WEBHOOK_AUTH_KEY` | Webhook auth header | INS.me1x9u... | Line 81 |
| `SCHEDULER_INTERVAL_SECONDS` | Send interval (seconds) | 120 | Line 83 |
| `MESSAGES_PER_BATCH` | Messages per batch | 2 | Line 84 |
| `DB_HOST` | PostgreSQL host | postgres | Line 70 |
| `REDIS_HOST` | Redis host | redis | Line 76 |

## Architecture

```
message-scheduler/
├── cmd/api/              # Application entry point
├── internal/
│   ├── config/           # Configuration
│   ├── models/           # Data models
│   ├── database/         # Database connection (database/sql)
│   ├── cache/            # Redis caching
│   ├── scheduler/        # Custom scheduler
│   ├── service/          # Business logic
│   ├── handler/          # HTTP handlers (net/http)
│   └── router/           # Route definitions
├── scripts/              # SQL schema and seed files
└── docker-compose.yml    # Container orchestration
```

## API Documentation

The API is fully documented using OpenAPI 3.0 (Swagger) specification.

### View API Documentation

**Swagger UI (Interactive):**
```
http://localhost:8081
```

The Swagger UI provides:
- Complete API reference
- Interactive "Try it out" functionality
- Request/response examples
- Schema definitions

**Swagger File:** `swagger.yaml` (OpenAPI 3.0 specification)

### Alternative Viewing Options

1. **Online Swagger Editor:**
   - Visit https://editor.swagger.io
   - Import the `swagger.yaml` file

2. **VS Code:**
   - Install "OpenAPI (Swagger) Editor" extension
   - Open `swagger.yaml`

## API Endpoints

### 1. Scheduler Control

**Endpoint:** `POST /api/scheduler`

**Start Scheduler:**
```bash
curl -X POST http://localhost:8080/api/scheduler \
  -H 'Content-Type: application/json' \
  -d '{"command": "start"}'
```

**Stop Scheduler:**
```bash
curl -X POST http://localhost:8080/api/scheduler \
  -H 'Content-Type: application/json' \
  -d '{"command": "stop"}'
```

### 2. Retrieve Sent Messages

**Endpoint:** `POST /api/messages/sent`

**Get all sent messages:**
```bash
curl -X POST http://localhost:8080/api/messages/sent \
  -H 'Content-Type: application/json' \
  -d '{
    "phone_number": "all",
    "limit": 10,
    "offset": 0
  }'
```

**Get messages for specific phone:**
```bash
curl -X POST http://localhost:8080/api/messages/sent \
  -H 'Content-Type: application/json' \
  -d '{
    "phone_number": "+905551234567",
    "limit": 10,
    "offset": 0
  }'
```

**With date filtering:**
```bash
curl -X POST http://localhost:8080/api/messages/sent \
  -H 'Content-Type: application/json' \
  -d '{
    "phone_number": "all",
    "limit": 10,
    "offset": 0,
    "sent_after": "2025-11-14T00:00:00Z",
    "sent_before": "2025-11-14T23:59:59Z"
  }'
```

**Request Parameters:**
- `phone_number` (required): Phone number or "all"
- `limit` (optional): Number of results, max 50, default 50
- `offset` (optional): Pagination offset, default 0
- `sent_after` (optional): RFC3339 timestamp
- `sent_before` (optional): RFC3339 timestamp

### 3. Health Check

**Endpoint:** `GET /health`

```bash
curl http://localhost:8080/health
```

## Seed Data

The system includes 122 pre-configured messages:

- **81 sent messages**: Pre-seeded as already sent with timestamps
- **31 pending messages**: Will be processed by scheduler
- **10 failed messages**: For error handling testing
- **6 phone numbers**: Various distribution patterns

### Message Processing Order

Messages are processed by `created_at` timestamp (oldest first). When starting fresh, the scheduler will send pending messages in the following order:

**Batch 1 (0 min):**
This message is supposed to be pending (1/31) - +905551234567
This message is supposed to be pending (23/31) - +905556667777

**Batch 2 (2 min):**
This message is supposed to be pending (2/31) - +905551234567
This message is supposed to be pending (3/31) - +905551234567

**Batch 3 (4 min):**
This message is supposed to be pending (12/31) - +905559876543
This message is supposed to be pending (21/31) - +905553334444

**Batch 4 (6 min):**
This message is supposed to be pending (28/31) - +905552221111
This message is supposed to be pending (4/31) - +905551234567

**Batch 5 (8 min):**
This message is supposed to be pending (5/31) - +905551234567
This message is supposed to be pending (13/31) - +905559876543

**Batch 6 (10 min):**
This message is supposed to be pending (19/31) - +905557778888
This message is supposed to be pending (24/31) - +905556667777

**Batch 7 (12 min):**
This message is supposed to be pending (6/31) - +905551234567
This message is supposed to be pending (7/31) - +905551234567

**Batch 8 (14 min):**
This message is supposed to be pending (14/31) - +905559876543
This message is supposed to be pending (20/31) - +905557778888

**Batch 9 (16 min):**
This message is supposed to be pending (8/31) - +905551234567
This message is supposed to be pending (9/31) - +905551234567

**Batch 10 (18 min):**
This message is supposed to be pending (15/31) - +905559876543
This message is supposed to be pending (22/31) - +905553334444

**Batch 11 (20 min):**
This message is supposed to be pending (25/31) - +905556667777
This message is supposed to be pending (10/31) - +905551234567

**Batch 12 (22 min):**
This message is supposed to be pending (16/31) - +905559876543
This message is supposed to be pending (29/31) - +905552221111

**Batch 13 (24 min):**
This message is supposed to be pending (17/31) - +905559876543
This message is supposed to be pending (26/31) - +905556667777

**Batch 14 (26 min):**
This message is supposed to be pending (18/31) - +905559876543
This message is supposed to be pending (30/31) - +905552221111

**Batch 15 (28 min):**
This message is supposed to be pending (27/31) - +905556667777
This message is supposed to be pending (31/31) - +905552221111

**Batch 16 (30 min):**
(Last message if no new pending messages added)

**Processing Details:**
- Interval: 2 messages every 120 seconds (2 minutes)
- Order: By created_at timestamp (oldest first)
- Tie-breaker: When timestamps match, ordered by ID (lowest first)
- Total time: Approximately 31 minutes to process all 31 pending messages

**Check Processing Status:**
```bash
# View pending messages
docker exec message_scheduler_postgres psql -U insider -d insider_db -c \
  "SELECT COUNT(*) FROM messages WHERE status='pending';"

# View sent messages
docker exec message_scheduler_postgres psql -U insider -d insider_db -c \
  "SELECT COUNT(*) FROM messages WHERE status='sent';"
```

## Database Schema

```sql
CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    phone_number VARCHAR(20) NOT NULL,
    content VARCHAR(500) NOT NULL,
    status VARCHAR(20) DEFAULT 'pending' NOT NULL,
    message_id VARCHAR(100),
    sent_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE INDEX idx_messages_status ON messages(status);
CREATE INDEX idx_messages_sent_at ON messages(sent_at);
CREATE INDEX idx_messages_created_at ON messages(created_at);
```

## How It Works

### Scheduler Implementation

The scheduler uses Go's standard library without external cron packages:

```go
type Scheduler struct {
    ticker   *time.Ticker    // Standard library timer
    stopChan chan bool       // Stop signal channel
    mu       sync.RWMutex    // Thread-safe operations
}
```

**Process Flow:**
1. Application starts and auto-starts scheduler
2. Every 2 minutes (configurable), scheduler:
   - Fetches 2 oldest pending messages (ORDER BY created_at ASC)
   - Sends each message to webhook endpoint
   - Updates status to "sent" or "failed"
   - Caches messageId and timestamp in Redis
3. Continues until stopped via API

## Redis Caching

After successfully sending a message, the system caches:

```json
{
  "message_id": "67f2f8a8-ea58-4ed0-a6f9-ff217df4d849",
  "sent_at": "2025-11-14T02:15:22Z",
  "phone_number": "+905551234567",
  "content": "Message content"
}
```

**Key format:** `message:sent:{messageId}`

**Access Redis:**
```bash
docker exec -it message_scheduler_redis redis-cli
KEYS message:sent:*
GET message:sent:{messageId}
```

## Monitoring

### View Logs

```bash
# Application logs
docker-compose logs -f app

# Database logs
docker-compose logs postgres

# Redis logs
docker-compose logs redis
```

### Database Queries

```bash
# Connect to database
docker exec -it message_scheduler_postgres psql -U insider -d insider_db

# View message counts by status
SELECT status, COUNT(*) FROM messages GROUP BY status;

# View recent sent messages
SELECT id, phone_number, content, sent_at 
FROM messages 
WHERE status='sent' 
ORDER BY sent_at DESC 
LIMIT 10;
```

### Check Pending Messages

```bash
# Count pending messages
docker exec message_scheduler_postgres psql -U insider -d insider_db -c \
  "SELECT COUNT(*) FROM messages WHERE status='pending';"

# View pending messages details
docker exec message_scheduler_postgres psql -U insider -d insider_db -c \
  "SELECT id, phone_number, LEFT(content, 50), created_at 
   FROM messages 
   WHERE status='pending' 
   ORDER BY created_at ASC, id ASC 
   LIMIT 10;"
```

## Testing

### Postman Tests

See `POSTMAN_TESTS.md` for complete test cases including:
- Health check
- Scheduler control (start/stop)
- Message retrieval (all, specific phone, pagination)
- Date filtering
- Error handling

### Add Test Messages

```bash
docker exec message_scheduler_postgres psql -U insider -d insider_db -c \
  "INSERT INTO messages (phone_number, content, status) VALUES 
   ('+905551234567', 'Test message content', 'pending');"
```

## Docker Commands

```bash
# Start services
docker-compose up -d

# View logs
docker-compose logs -f app

# Stop services
docker-compose down

# Restart with fresh data
docker-compose down -v && docker-compose up -d

# Rebuild after code changes
docker-compose up -d --build
```

## Technology Stack

| Component | Technology |
|-----------|-----------|
| Language | Go 1.21+ |
| HTTP Server | net/http (standard library) |
| Database | PostgreSQL 15 |
| Database Driver | lib/pq |
| Cache | Redis 7 |
| Redis Client | go-redis/v8 |

**Note:** No external frameworks used. Built with Go standard library for simplicity and learning.

## Requirements Met

- Auto-send 2 messages every 2 minutes
- Retrieve messages from database (no message creation API)
- Auto-start scheduler on deployment
- No duplicate sends
- Process newly added records automatically
- Redis caching for sent messages
- Two API endpoints: scheduler control and message retrieval
- Custom scheduler without cron packages
- Docker Compose for easy setup
- SQL-based seed data

## Troubleshooting

### Services won't start

Check port availability:
```bash
lsof -i :8080  # API
lsof -i :5432  # PostgreSQL
lsof -i :6379  # Redis
```

### Scheduler not processing messages

1. Check scheduler status:
```bash
curl http://localhost:8080/health
```

2. Verify pending messages exist:
```bash
docker exec message_scheduler_postgres psql -U insider -d insider_db -c \
  "SELECT COUNT(*) FROM messages WHERE status='pending';"
```

3. Check logs for errors:
```bash
docker-compose logs -f app
```

### Database connection issues

```bash
# Check PostgreSQL is running
docker-compose ps postgres

# Reinitialize database
docker-compose down -v
docker-compose up -d
```

## License

MIT