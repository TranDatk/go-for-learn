social/
│
├── bin/ 
├── cmd/
│   ├── api/
│   └── migrate/
│       └── migrations/
│
├── api/
│   ├── handlers/
│   ├── middleware/
│   └── router.go
│
├── internal/
│   ├── storage/
│   ├── email/
│   ├── ratelimiter/
│   ├── services/
│   └── mocks/
│
├── docs/
├── scripts/
└── web/ (optional)

/bin
Chứa binary sau khi build (Go build → output vào đây).


/cmd
Chứa entrypoints của ứng dụng.
Ví dụ:
/cmd/api – server API chính
/cmd/migrate – nơi chạy migration tool custom (nếu có)

/cmd/migrate/migrations
Nơi chứa file migration .sql cho Postgres.
Dùng cho tool Golang Migrate.

/api
Chứa tất cả thứ liên quan HTTP:
+ handlers
+ controllers
+ routers
+ middlewares
+ transport layer
→ Đây là nơi request “bước vào” hệ thống.

/internal
Một thư mục cực kỳ quan trọng trong thế giới Go.
Ý nghĩa:
Code trong thư mục internal KHÔNG được import từ bên ngoài.
Nghĩa là:
→ nội bộ hệ thống dùng với nhau
→ tránh lộ logic không cần thiết
Bên trong /internal sẽ có:
✔ storage/
Tương tác database (Postgres), repository pattern.
✔ email/
Gửi email (SendGrid hoặc nhà cung cấp khác).
✔ ratelimiter/
Cài đặt rate limiter custom.
✔ validation/
Validate dữ liệu.
✔ mocks/
Mock interfaces để test.
✔ services/
Business logic chính của hệ thống.

/docs
Chứa file Swagger auto-generated.

/scripts
Script hỗ trợ:
+ start server
+ setup env
+ seed database
+ deploy scripts (nếu cần)

/web (tùy chọn)
Nếu bạn muốn xây dựng full-stack trong monorepo:
React
Svelte
HTML static
Hoặc một Go server serve HTML. Tùy ý đặt ở đây.
Giảng viên thường:
Deploy /web như static site
Deploy backend trong container