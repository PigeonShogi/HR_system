# HR_system
---
本專案為學習 Golang 的副產品，完成後將是一個簡易的考勤系統網站，具備上下班打卡、檢視出勤記錄等功能。
## 專案應用技術
- PostgreSQL：建立資料庫
- Docker：建立 PostgreSQL 容器
- Makefile：建立 Makefile 檔案，方便重複執行特定指令。
- golang-migrate：以本套件落實資料遷移
- sqlc：以本套件將 SQL Queries 轉換為 Go 函式
- 單元測試：利用 Go 原生套件落實單元測試
- CI/CD：透過 GitHub Actions 實現
- Server：使用 Gin Gonic 框架建立伺服器
- gomock：以模擬資料庫輔助單元測試
## 專案發展目標
在此之前我曾用 Node.js 搭配 Vue3 開發考勤系統網站的 Side Project，其相關連結如下：

##### 後端 Repo
https://github.com/PigeonShogi/Ti-HR
##### 前端 Repo
https://github.com/PigeonShogi/Ti-HR-Client
##### 前後端專案整合 demo
* https://ti-hr-client.vercel.app/
##### Demo 帳號
* 管理者
  帳號：admin
  密碼：tiadmin
* 一般使用者
  帳號：user
  密碼：titaner

今後將視上述作品為本專案之前身，致力於以 Go 實做其各項功能及擴充新功能。