# Configuration Guide

## Webhook URL Setup (Required)

Before running the system, you must configure your own webhook endpoint to receive messages.

### Step 1: Get Your Webhook URL

1. Visit **https://webhook.site**
2. You will automatically be assigned a unique URL
3. Copy your URL (format: `https://webhook.site/YOUR-UNIQUE-ID`)
4. Keep this browser tab open to see incoming messages

**Example:**
```
https://webhook.site/9e83d1aa-0046-4f29-a2d8-deb68ce93ec0
```

### Step 2: Configure the System

**Option A: Edit docker-compose.yml (Recommended)**

1. Open `docker-compose.yml`
2. Find line 84 (look for `WEBHOOK_URL:`)
3. Replace the default URL with your webhook URL:

```yaml
WEBHOOK_URL: ${WEBHOOK_URL:-https://webhook.site/YOUR-UNIQUE-ID}
```

**Option B: Use Environment Variable**

```bash
export WEBHOOK_URL="https://webhook.site/YOUR-UNIQUE-ID"
docker-compose up -d
```

### Step 3: Start the System

```bash
docker-compose up -d
```

### Step 4: Verify Messages Are Being Sent

1. Open your webhook.site URL in a browser
2. The scheduler will send 2 messages every 2 minutes (or 3 seconds if testing)
3. You should see POST requests appearing with message content

**Example webhook payload:**
```json
{
  "to": "+905551234567",
  "content": "This message is supposed to be pending (1/31) - waiting for scheduler"
}
```

## Configuration Variables Reference

All configuration is in `docker-compose.yml` under the `app` service environment section:

| Variable | Line | Description | Default | Change? |
|----------|------|-------------|---------|---------|
| `WEBHOOK_URL` | 84 | Your webhook endpoint | webhook.site URL | **YES - Required** |
| `WEBHOOK_AUTH_KEY` | 85 | Authorization header | INS.me1x9u... | Optional |
| `SCHEDULER_INTERVAL_SECONDS` | 87 | Send interval | 3 (testing) / 120 (production) | Optional |
| `MESSAGES_PER_BATCH` | 88 | Messages per batch | 2 | No |
| `DB_HOST` | 70 | Database host | postgres | No |
| `DB_PORT` | 71 | Database port | 5432 | No |
| `DB_USER` | 72 | Database user | insider | No |
| `DB_PASSWORD` | 73 | Database password | insider123 | No |
| `REDIS_HOST` | 76 | Redis host | redis | No |
| `SERVER_PORT` | 89 | API server port | 8080 | No |

## Testing Configuration

### For Quick Testing (3-second intervals)

Already configured in `docker-compose.yml`:
```yaml
SCHEDULER_INTERVAL_SECONDS: 3
```

This will send messages every 3 seconds for quick demonstration.

### For Production (2-minute intervals)

Change line 87:
```yaml
SCHEDULER_INTERVAL_SECONDS: 120
```

Then restart:
```bash
docker-compose restart app
```

## Troubleshooting

### Messages not appearing on webhook.site

1. **Check webhook URL is correct:**
   ```bash
   docker-compose logs message_scheduler_app | grep "webhook"
   ```

2. **Verify scheduler is running:**
   ```bash
   curl http://localhost:8080/health
   ```
   Should show `"scheduler_status": true`

3. **Check for pending messages:**
   ```bash
   docker exec message_scheduler_postgres psql -U insider -d insider_db -c \
     "SELECT COUNT(*) FROM messages WHERE status='pending';"
   ```

4. **View scheduler logs:**
   ```bash
   docker-compose logs -f app | grep "Batch completed"
   ```

### Webhook.site not receiving messages

- Ensure your webhook.site URL is still valid (they expire after some time)
- Get a fresh URL from https://webhook.site
- Update `docker-compose.yml` and restart: `docker-compose restart app`

## Quick Configuration Summary

**Minimum steps to run:**
1. Get webhook URL: https://webhook.site
2. Edit `docker-compose.yml` line 84
3. Run: `docker-compose up -d`
4. Watch: Your webhook.site URL in browser

That's it! The system will automatically:
- Create database and tables
- Load 122 test messages
- Start sending pending messages (31 messages)
- Process 2 messages every 3 seconds

