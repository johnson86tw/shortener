# shortener

### 加入使用者(暫時傾向先做一張表的)

- 需不需要將有持有者的和沒有持有者的 url 分成兩個 table？
- 分成兩個 table，server 要重導的時候，就得分別到這兩個 table 去找相對應的 code
- 分開來的好處是，普通版的 url 可以重複使用，因為不需要紀錄 total_clicks
- 分表要從 id 取資料的話要怎麼取，如果 id 可能重複？
- 感覺是用同一張表就可以了，# 可是共用一張表，大量的無持有者 url 根本不需要那麼多欄位
- 如何確保兩張表的 url code 是 unique 的

1. User CRUD (先做完這部分再說)
2. Redirect Binding User 

### 是否需要 user-urls 還是延伸 urls table

url_id serial not null unique,
url text
code varchar(12)
owner 有無持有者
total_clicks integer
created_at timestamp default now()
updated_at timestamp default now()
deleted_at

### UUID & password hash

- 應該都是在 service 上處理，database 只用於檢查 uuid 的型別
- uuid 可透過 go 產出後以 pgtype uuid 存進 db，或者在 dbeaver 透過 extension 去產 uuid insert 進去
- password hash 在 db 沒有檢查機制，因此可存放明碼，但需要在 service hash 過在存進去

### Redirect vs API
- Redirect 只是一組資料的結構，真正讓功能分成兩種的是api，可以只取得Redirect的URL讓router重新導向，或者取得Redirect整組資料然後秀出來