# shortener

- POST /              // store url for general people
- GET /:code          // redirect url by code
- POST /signup        // sign up: require name, email, and password
- POST /login         // login: require email and password
- POST /auth/url      // add custom url by authorized user: require url and token
- GET /auth/urls      // get all urls by authorized user: require token

### Guild
- web 的 env DEBUG=true 用於 docker-compose，專案開發中的時候，自訂 .env DEBUG 為 true 時，即可 go run main.go port 4000
- 正式站重 build web image，開發時可在本機 run go 連接 docker-compose 的 db 和 redis

```
docker-compose up -d
```
- localhost:443

### Future add-on
- [Rate Limiter with Redis](https://github.com/ulule/limiter)
- Envoy and Frontend
- [Monitor System](https://blog.techbridge.cc/2019/08/26/how-to-use-prometheus-grafana-in-flask-app/)

### References
- [go-clean-arch](https://github.com/bxcodec/go-clean-arch)

### Question
- 為什麼 go mod tidy 會引入其他不認識的套件
- logrus 是否每一個 err 都要印出來？比較好找 bug？
- 在 login 的地方，必須想辦法區分是 sql 取資料的錯誤，還是 service 層的錯誤，還是真的密碼沒過
- 在 UserURL - AddURL 的 interface 應該設計成製作詳細填入必要的參數如 userID URL Code 還是直接塞入 *UserURL 就好？
- 在 UserURL - FetchAll 當 handler 不想要回傳 UserID，但 domain.UserURL 有 UserID 的欄位，要怎麼隱藏？

### Issues
- 通用的 redirect 有重複的 url 就用舊的就好，不必再新增檔案
- url 最後面多一撇也要可以使用
- db 的錯誤訊息要處理掉
- logout 後 token 要 revoked
- config 自成一個 package 處理，要支援 .env 檔，有 env 用 env，沒有的話就用 config.json
- add Testing
- 客製化 user_url