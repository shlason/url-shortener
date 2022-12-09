# url-shortener

一個短網址服務，主要是用來練習看看 Gin 和 GORM 的練習項目。

API 文件：[點我](https://short.sidesideeffect.io/swagger/index.html)

線上 Demo: [點我](https://short.sidesideeffect.io)

![url-shortener demo image](https://raw.githubusercontent.com/shlason/url-shortener/docs/images/demo.png)

## 實現方式
在資料庫中有一個表是專門存放產出的短網址和相對應原始網址的一個關係表，短網址的部分目前都是使用 base62 的編碼格式，來減少網址的長度。

而要使用什麼東西做為短網址，一開始原本想直接用資料表的 primary key ID，但這樣即便編碼為 base62 還是能夠看出來每次有新網址的轉換短網址都會是一個遞增的關係，感覺會有一點安全疑慮 (雖然這個短網址服務本身就是公開沒有隱私性的)，因此最後還是沒有採用，目前是先暫時使用當下的 timestamp 取到毫秒，再將 timestamp 編碼為 base62 格式來作為短網址的對應 ID。

## TODO
使用 Redis 避免一直查詢資料庫

## 技術棧
- Go
- Gin (管理路由和 middleware)
- `gopkg.in/ini.v1` (用來讀取設定檔 configs.ini)
- MySQL
- GORM
- swaggo/swag (產 API 文件)
- autocert (產 SSL 憑證)
- github Actions (用以 build 專案和部署到 EC2 上)
- Cloudflare DNS
- AWS EC2