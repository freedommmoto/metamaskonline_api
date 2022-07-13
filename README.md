# metamaskonline_api
api connect to bscscan and line api 

````
on mac : 

brew install ngrok/ngrok/ngrok
ngrok http 8888
cd poc
vim lineapi.go 
****
change tokenapi to you token 
you can get tokenapi for free by enter this link 
https://account.line.biz/login
****
go run lineapi.go
enter 0.0.0.0:8080/testpush
enter any text on line channel
````