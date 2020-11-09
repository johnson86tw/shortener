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
- url 最後面多一撇也要可以使用
- 錯誤訊息要處理掉: ex. 尚未建立 table, login 的 bcrypt 錯誤
- logout 後 token 要 revoked
- config 自成一個 package 處理，要支援 .env 檔，有 env 用 env，沒有的話就用 config.json
- add Testing
- 客製化 user_url

### 一些寫作原則
- log 盡量都寫在 service 層
- 將 model 隱藏於 json 的方法： `json:"-"` 或 小寫的 struct field
- 檔名寫清楚： xxx_handler, pg_xxx, xxx_service
- db 的錯誤訊息要在 repo 層處理
- api 層與 service 層可以用指標的方式賦值，db 層的 input 則盡量單純；但模仿的專案是讓 service 層同 db 層一樣單純，盡量不傳 struct 作為參數
- struct如果是pointer可以使用nil，但nil的壞處是在取field的值時如果是nil就會噴錯，因此，service層能回傳值就回傳值就好，除非input指標進來output指標出去，否則選擇傳值。如果是傳指標，也不要用nil回傳，回傳該struct的預設值的指標比較好。