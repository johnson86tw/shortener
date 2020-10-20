# shortener

### 加入使用者客製化短網址
- 建立新表 user_urls table
- 如何用 JWT 知道使用者的 uuid？發 claim 時把 uuid 塞在裡面，authRequired 函式將其解開

### Bugs to be fixed
- users table name need to be not null

### References
- [go-clean-arch](https://github.com/bxcodec/go-clean-arch)

### Question
- 為什麼 go mod tidy 會引入其他不認識的套件
- logrus是否每一個err都要印出來？比較好找bug？
