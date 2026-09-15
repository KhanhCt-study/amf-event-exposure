# AMF Event Exposure Service (Namf_EventExposure)

Triển khai phía **AMF server** của dịch vụ Namf_EventExposure theo 3GPP TS 29.518,
đi kèm NWDAF consumer đã có trong `nwdaf-event-subscription`.

Kiến trúc, coding conventions, error model và schema bám theo bản kế hoạch
`amf-event-exposure-plan.md`.

## API

| Method | URI | Status thành công |
|--------|-----|-------------------|
| POST | `/namf-evts/v1/subscriptions` | 201 + `Location` |
| PATCH | `/namf-evts/v1/subscriptions/{subscriptionId}` | 200 |
| DELETE | `/namf-evts/v1/subscriptions/{subscriptionId}` | 204 |
| POST | `{eventNotifyUri}` (AMF → consumer) | consumer trả 204 |

Ngoài spec, có thêm hai route phục vụ vận hành/prototype:

| Method | URI | Mục đích |
|--------|-----|----------|
| GET | `/healthz` | health check |
| POST | `/internal/v1/event-reports` | kích hoạt thủ công một event report để test luồng notification |

Mọi lỗi trả `application/problem+json` (ProblemDetails, RFC 9457):
`ErrInvalid→400`, `ErrUnsupported→403`, `ErrNotFound→404`, `ErrConflict→409`, còn lại `500`.

## Chạy thử

```bash
go mod tidy

export DATABASE_URL="postgres://postgres:postgres@localhost:5432/amf?sslmode=disable"
export AMF_LISTEN_ADDR=":8081"
export AMF_PUBLIC_API_URL="http://localhost:8081"

go run ./cmd/amf-event-exposure/server
```

Migration `001_create_amf_subscriptions` được nhúng vào binary và tự áp dụng khi
khởi động (bảng dùng `gen_random_uuid()` nên cần PostgreSQL 13+).

### Biến môi trường

| Biến | Mặc định | Ghi chú |
|------|----------|---------|
| `AMF_LISTEN_ADDR` | `:8081` | |
| `AMF_PUBLIC_API_URL` | `http://localhost:8081` | dùng cho `Location` và `subscriptionId` |
| `AMF_NF_INSTANCE_ID` | — | UUID, tùy chọn |
| `DATABASE_URL` | — | **bắt buộc** |
| `AMF_TIMEOUT` | `10s` | timeout HTTP client gửi notification |
| `DATABASE_TIMEOUT` | `5s` | |
| `AMF_SHUTDOWN_TIMEOUT` | `15s` | |
| `AMF_DEFAULT_MCC` / `AMF_DEFAULT_MNC` / `AMF_DEFAULT_TAC` / `AMF_DEFAULT_NR_CELL_ID` | `208` / `95` / `000001` / `000000001` | vị trí tĩnh dùng cho LOCATION_REPORT ở giai đoạn prototype |

### Ví dụ

Subscribe:

```bash
curl -i -X POST http://localhost:8081/namf-evts/v1/subscriptions \
  -H 'Content-Type: application/json' \
  -d '{
    "subscription": {
      "eventList": [{"type": "LOCATION_REPORT", "immediateFlag": true}],
      "eventNotifyUri": "http://nwdaf:8080/callbacks/amf-notify",
      "notifyCorrelationId": "corr-abc-123",
      "nfId": "550e8400-e29b-41d4-a716-446655440000",
      "supi": "imsi-208950000000010",
      "options": {"trigger": "CONTINUOUS", "maxReports": 100}
    }
  }'
```

Modify:

```bash
curl -i -X PATCH http://localhost:8081/namf-evts/v1/subscriptions/{id} \
  -H 'Content-Type: application/json' \
  -d '[{"op":"replace","path":"/subscription/eventList/0",
        "value":{"type":"LOCATION_REPORT","immediateFlag":false,"maxReports":50}}]'
```

Kích hoạt notification (hook nội bộ):

```bash
curl -i -X POST http://localhost:8081/internal/v1/event-reports \
  -H 'Content-Type: application/json' \
  -d '{"subscriptionId":"{id}","type":"LOCATION_REPORT","supi":"imsi-208950000000010"}'
```

Unsubscribe:

```bash
curl -i -X DELETE http://localhost:8081/namf-evts/v1/subscriptions/{id}
```

## Test

```bash
go test ./...
```

Hiện có unit test cho router, `api` response writer, subscription usecase và
notifier — tất cả dùng manual mocks, không mockgen.

## Trạng thái theo feature branch

| Branch | Trạng thái |
|--------|------------|
| 1. project-setup | ✅ go.mod, config, main.go, database.go |
| 2. domain-models | ✅ entity, interfaces, errors |
| 3. api-framework | ✅ router, api package, recovery middleware, router_test |
| 4. db-migration | ✅ migration + repository (CAS) — ⚠️ chưa có contract test testcontainers |
| 5. api-create-subscription | ✅ handler + usecase + validation |
| 6. api-modify-subscription | ✅ JSON Patch trong phạm vi prototype |
| 7. api-delete-subscription | ✅ CAS delete |
| 8. notification-sender | ✅ client + notifier + hook kích hoạt |
| 9. unit-tests | 🟡 đã có cho usecase/api/router; còn thiếu handler, validator, config, repo |
| 10. integration-tests | ❌ chưa làm (testcontainers PostgreSQL) |

Chi tiết giả định và điểm cần rà lại: xem `docs/implementation-notes.md`.
