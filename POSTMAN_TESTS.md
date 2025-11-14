# 🚀 Postman API Test Cases

## Base URL
```
http://localhost:8080
```

---

## 1️⃣ Health Check

**Endpoint:** `GET /health`

**Method:** GET

**Headers:** None required

**Expected Response:**
```json
{
  "success": true,
  "message": "API is running",
  "data": {
    "status": "healthy",
    "scheduler_status": true
  }
}
```

---

## 2️⃣ Scheduler Control - Start

**Endpoint:** `POST /api/scheduler`

**Method:** POST

**Headers:**
```
Content-Type: application/json
```

**Body (raw JSON):**
```json
{
  "command": "start"
}
```

**Expected Response:**
```json
{
  "success": true,
  "message": "Scheduler started"
}
```

---

## 3️⃣ Scheduler Control - Stop

**Endpoint:** `POST /api/scheduler`

**Method:** POST

**Headers:**
```
Content-Type: application/json
```

**Body (raw JSON):**
```json
{
  "command": "stop"
}
```

**Expected Response:**
```json
{
  "success": true,
  "message": "Scheduler stopped"
}
```

---

## 4️⃣ Scheduler Control - Invalid Command (Error Test)

**Endpoint:** `POST /api/scheduler`

**Method:** POST

**Headers:**
```
Content-Type: application/json
```

**Body (raw JSON):**
```json
{
  "command": "pause"
}
```

**Expected Response:**
```json
{
  "success": false,
  "error": "Invalid command. Use 'start' or 'stop'"
}
```

---

## 5️⃣ Get All Sent Messages (Default)

**Endpoint:** `POST /api/messages/sent`

**Method:** POST

**Headers:**
```
Content-Type: application/json
```

**Body (raw JSON):**
```json
{
  "phone_number": "all",
  "limit": 10,
  "offset": 0
}
```

**Expected Response:**
```json
{
  "success": true,
  "data": {
    "total": 114,
    "messages": [
      {
        "id": 124,
        "phone_number": "+905559876543",
        "content": "NEW TEST MESSAGE: Second fresh message...",
        "status": "sent",
        "message_id": "67f2f8a8-ea58-4ed0-a6f9-ff217df4d849",
        "sent_at": "2025-11-14T02:18:22.624131Z",
        "created_at": "2025-11-14T02:18:19.479291Z",
        "updated_at": "2025-11-14T02:18:22.624131Z"
      }
      // ... 9 more messages
    ]
  }
}
```

---

## 6️⃣ Get Sent Messages for Specific Phone Number

**Endpoint:** `POST /api/messages/sent`

**Method:** POST

**Headers:**
```
Content-Type: application/json
```

**Body (raw JSON):**
```json
{
  "phone_number": "+905551234567",
  "limit": 10,
  "offset": 0
}
```

**Expected Response:**
```json
{
  "success": true,
  "data": {
    "total": 66,
    "messages": [
      // 10 messages for +905551234567
    ]
  }
}
```

---

## 7️⃣ Get Sent Messages - Pagination (Page 1)

**Endpoint:** `POST /api/messages/sent`

**Method:** POST

**Headers:**
```
Content-Type: application/json
```

**Body (raw JSON):**
```json
{
  "phone_number": "+905551234567",
  "limit": 50,
  "offset": 0
}
```

**Expected Response:**
```json
{
  "success": true,
  "data": {
    "total": 66,
    "messages": [
      // First 50 messages
    ]
  }
}
```

---

## 8️⃣ Get Sent Messages - Pagination (Page 2)

**Endpoint:** `POST /api/messages/sent`

**Method:** POST

**Headers:**
```
Content-Type: application/json
```

**Body (raw JSON):**
```json
{
  "phone_number": "+905551234567",
  "limit": 50,
  "offset": 50
}
```

**Expected Response:**
```json
{
  "success": true,
  "data": {
    "total": 66,
    "messages": [
      // Remaining 16 messages
    ]
  }
}
```

---

## 9️⃣ Get Sent Messages - With Date Filter (After)

**Endpoint:** `POST /api/messages/sent`

**Method:** POST

**Headers:**
```
Content-Type: application/json
```

**Body (raw JSON):**
```json
{
  "phone_number": "all",
  "limit": 10,
  "offset": 0,
  "sent_after": "2025-11-14T00:00:00Z"
}
```

**Expected Response:**
```json
{
  "success": true,
  "data": {
    "total": 33,
    "messages": [
      // Messages sent after Nov 14, 2025 00:00:00
    ]
  }
}
```

---

## 🔟 Get Sent Messages - With Date Filter (Before)

**Endpoint:** `POST /api/messages/sent`

**Method:** POST

**Headers:**
```
Content-Type: application/json
```

**Body (raw JSON):**
```json
{
  "phone_number": "all",
  "limit": 10,
  "offset": 0,
  "sent_before": "2025-11-13T00:00:00Z"
}
```

**Expected Response:**
```json
{
  "success": true,
  "data": {
    "total": 81,
    "messages": [
      // Messages sent before Nov 13, 2025 00:00:00
    ]
  }
}
```

---

## 1️⃣1️⃣ Get Sent Messages - Date Range Filter

**Endpoint:** `POST /api/messages/sent`

**Method:** POST

**Headers:**
```
Content-Type: application/json
```

**Body (raw JSON):**
```json
{
  "phone_number": "all",
  "limit": 50,
  "offset": 0,
  "sent_after": "2025-11-11T00:00:00Z",
  "sent_before": "2025-11-11T23:59:59Z"
}
```

**Expected Response:**
```json
{
  "success": true,
  "data": {
    "total": 20,
    "messages": [
      // Messages sent on Nov 11, 2025
    ]
  }
}
```

---

## 1️⃣2️⃣ Get Sent Messages - Limit Validation (Over 50)

**Endpoint:** `POST /api/messages/sent`

**Method:** POST

**Headers:**
```
Content-Type: application/json
```

**Body (raw JSON):**
```json
{
  "phone_number": "all",
  "limit": 100,
  "offset": 0
}
```

**Expected Response:**
```json
{
  "success": true,
  "data": {
    "total": 114,
    "messages": [
      // Returns MAXIMUM 50 messages (limit capped at 50)
    ]
  }
}
```

**Note:** Even though we requested 100, the system caps it at 50.

---

## 1️⃣3️⃣ Get Sent Messages - Phone Number Not Found

**Endpoint:** `POST /api/messages/sent`

**Method:** POST

**Headers:**
```
Content-Type: application/json
```

**Body (raw JSON):**
```json
{
  "phone_number": "+905559999999",
  "limit": 10,
  "offset": 0
}
```

**Expected Response:**
```json
{
  "success": false,
  "error": "Phone number not found"
}
```

---

## 1️⃣4️⃣ Get Sent Messages - Missing Required Field

**Endpoint:** `POST /api/messages/sent`

**Method:** POST

**Headers:**
```
Content-Type: application/json
```

**Body (raw JSON):**
```json
{
  "limit": 10,
  "offset": 0
}
```

**Expected Response:**
```json
{
  "success": false,
  "error": "phone_number field is required"
}
```

---

## 1️⃣5️⃣ Get Sent Messages - Default Limit

**Endpoint:** `POST /api/messages/sent`

**Method:** POST

**Headers:**
```
Content-Type: application/json
```

**Body (raw JSON):**
```json
{
  "phone_number": "all"
}
```

**Expected Response:**
```json
{
  "success": true,
  "data": {
    "total": 114,
    "messages": [
      // Returns 50 messages (default limit)
    ]
  }
}
```

---

## 📊 Quick Test Summary

| Test Case | Endpoint | Expected Result |
|-----------|----------|-----------------|
| Health Check | GET /health | 200 OK |
| Start Scheduler | POST /api/scheduler | Started |
| Stop Scheduler | POST /api/scheduler | Stopped |
| Invalid Command | POST /api/scheduler | Error message |
| Get All Sent | POST /api/messages/sent | 114 total |
| Specific Phone | POST /api/messages/sent | Phone-specific count |
| Pagination Page 1 | POST /api/messages/sent | Max 50 messages |
| Pagination Page 2 | POST /api/messages/sent | Remaining messages |
| Date Filter (After) | POST /api/messages/sent | Filtered by date |
| Date Filter (Before) | POST /api/messages/sent | Filtered by date |
| Date Range | POST /api/messages/sent | Specific day only |
| Limit > 50 | POST /api/messages/sent | Capped at 50 |
| Phone Not Found | POST /api/messages/sent | 404 Error |
| Missing Field | POST /api/messages/sent | 400 Error |

---

## 🎯 Postman Collection Setup

### Import as Collection

You can copy each test case into Postman manually, or create a collection with these settings:

1. **Create New Collection**: "Insider Message API"
2. **Add Requests** for each test case above
3. **Set Base URL** as variable: `{{base_url}}` = `http://localhost:8080`
4. **Headers** preset: `Content-Type: application/json`

### Environment Variables

Create an environment with:
```
base_url = http://localhost:8080
phone_number = +905551234567
```

---

## 🔍 Expected Counts (Current State)

- **Total Sent Messages**: 114
- **Messages for +905551234567**: 66
- **Messages sent on Nov 14, 2025**: 33
- **Total Pending**: 0
- **Total Failed**: 10

---

## ⚡ Quick Copy-Paste for Postman

### Request 1: Get All Sent Messages
```
POST http://localhost:8080/api/messages/sent
Content-Type: application/json

{
  "phone_number": "all",
  "limit": 10,
  "offset": 0
}
```

### Request 2: Get Specific Phone
```
POST http://localhost:8080/api/messages/sent
Content-Type: application/json

{
  "phone_number": "+905551234567",
  "limit": 10,
  "offset": 0
}
```

### Request 3: Stop Scheduler
```
POST http://localhost:8080/api/scheduler
Content-Type: application/json

{
  "command": "stop"
}
```

### Request 4: Start Scheduler
```
POST http://localhost:8080/api/scheduler
Content-Type: application/json

{
  "command": "start"
}
```

---

Happy Testing! 🎉

