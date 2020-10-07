# shortener

### 加入 docker-compose 採的坑

- config.json redis address 必須用 redis:6379，而且 docker-compose 的 web 要加入 link，否則 web 會一直連不上 redis

### 待加入

- 如何不重新 build image 只更動 code 就能改變 docker-compose 的結果？
- init.sql 建立三個 table

table url

- id
- url
- code
- created_at

table user

- id
- name
- email
- password

table user_url

- user_id
- url_id
- created_at
