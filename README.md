# dcard-demo

## 題目描述 : 

設計縮短網址的功能，此開發需開發兩隻api，分別是上傳需要縮短的網址，及輸入短網址後導到原網址。

* 第一支api為上傳原網址及過期時間的資料，response則為回傳一個短網址

* 第二支api則是使用短網址，重新定向至原網址

* 如果短網址已過期回覆 http status 404

* 可使用各種資料庫去實作

* 需對這兩隻api去實作合理的錯誤處理

* 無需考慮身分驗證

* 需考慮客戶端可能會大量訪問短網址，或是已不存在的的短網址進行訪問

## 使用語言及技術

語言框架 : GO-Gin 

使用Gin來建構RESTful API

使用套件 : 

* redis : 因有可能會有大量用戶來對服務進行大量訪問，因此我採用redis這種快取資料庫來做為資料上的存取

* rs/cors : 網頁跨域連線處理

* satori/go.uuid : 生成uuid

* spf13/viper : 取得yaml檔裡的連線資訊

## 實作

### 第一支api : 用來新增短網址

生成一組uuid取前8碼，來當短網址的唯一key，總數有16^8的可能，要重複的機率相當的低。因此會將key=uuid, value=originalUrl存入redis，且會設定此key的到期時間。

因uuid的總數十分龐大，目前是不會檢查redis裡是否已經有重複的key了，若有大量的使用者來做新增的操作，可加上檢核的機制。

### 第二支api : 用來導向原url

將短網址上的param取出，至redis取出原url，重新定向至原網址。

### error handling

我定義了兩種錯誤分別為AppError、NotFoundError，以及實作了一個ErrorHandler。

ErrorHandler主要會去判斷controller所回傳的error種類的不同，會回傳不同的http status給到訪問端。

* AppError : http status 400, 在程式碼中若是本身自己檢核出來的錯誤，或是自己能掌握的錯誤，會以此拋出

* NotFoundError : http status 404，在此專案中，當我從redis中，用訪問端給的urlId查不出資料時會以此拋出

* others : http status 500，我無法主動得知的錯誤，預設回傳500

### test

我將兩支api的測試用TestWrapper再次封裝起來，目的是將兩者的關聯性給串起來。可使用第一支api回傳的urlId做為第二支api的測試參數，並檢查第二支api所給的原url是否為當初第一支api所給的原url。

## 延伸

因為資料庫是使用快取資料庫，若server或VM掛掉後，資料是會全部遺失的，若要將資料持久化，可再建一個postgres做為終端，第一支api在寫入redis時可再寫一份到postgres，
之後可以在server出狀況後有機會將資料重新倒回redis。


