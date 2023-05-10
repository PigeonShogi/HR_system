# 2023/5/10
## 新增 Makefile
## 安裝 golang-migrate 的 migrate CLI
 $ brew install golang-migrate
安裝後檢視版本號
 $ migrate -version
建立第一個 migration
 $ migrate create -ext sql -dir db/migration -seq init_schema
Makefile 加入 migrate 腳本
 $ migrate -path db/migration -database "postgresql://root:secret@localhost:5432/hr_system?sslmode=disable" -verbose up
 $ migrate -path db/migration -database "postgresql://root:secret@localhost:5432/hr_system?sslmode=disable" -verbose down

# 2023/5/9
開始編寫專案
新增 db schema
