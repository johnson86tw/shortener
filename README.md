# shortener

### 關於這份專案
這是一個縮短網址的服務，作為我學習 golang 網頁後端的 side project，這是第一份筆記，往後隨著功能的擴張，我會把舊的程式碼（包含這份筆記）存成分支，希望同樣在學習之路的開發者們也能從這份專案中學到些東西。作者是從前端轉進來的，也就是從 javascript 轉進 golang，js 之前沒有任何資訊工程的背景，學習 golang 後才開始慢慢去碰演算法，如果你也同樣沒什麼底子，那跟我懂得就差不多，更適合看我的程式碼去補足所需的知識。

### 架構
希望這份專案未來能朝以下的概念發展下去
- microservice
- domain driven design (DDD)
- test driven design
- clean code arch
- dockerize

### 學到什麼
- initial prototype 主要是學習怎麼從零開始架構一個服務，從中可以體會 interface 的奇妙之處，以及 golang 的 package 系統，如此架構起專案，感覺真的很像在玩積木一樣，從 DB 積木開始，repository 的介面，拼裝到 service，然後 service 的介面再拼裝到 API handler，這麼寫起來非常清晰，而且 model 統一放在 domain 裡頭，非常有助於聚焦服務的核心。在規劃一項產品時，PM 應該可以去理解甚至去架構 domain 裡頭的程式碼，在發想產品原型時將他們定義出來，再由工程師去實作 domain 裡頭的資料結構與介面，這應該會是不錯的產品發想方式。
