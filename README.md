# shortener

### 最後兩個部份
- redirect service Find 同時搜尋 urls and user_urls table，如果是 user_urls 的要紀錄 totalClick
- redis for caching

### issues to be fixed
- users table name need to be not null

### References
- [go-clean-arch](https://github.com/bxcodec/go-clean-arch)

### Question
- 為什麼 go mod tidy 會引入其他不認識的套件
- logrus是否每一個err都要印出來？比較好找bug？
- 例如login的地方，必須想辦法區分是sql取資料的錯誤，還是service層的錯誤，還是真的密碼沒過

### 設計原則
- 要替換 DB 只需要動到 repository
- 要替換 web framework 只需要動到 api