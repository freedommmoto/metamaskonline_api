# metamaskonline_api
api connect to bscscan and line api 

````
1) make poc line api work 

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

````
2) make poc bsc api work 
cd poc
vim bscscanapi.go
****
change tokenapi to you token 
you can get tokenapi for free by enter this link 
https://bscscan.com/myapikey
****
go run bscscanapi.go
you will see api connect on log
````

````
3) setup app database and basic data 
make pullpostgres
make installmigration
make migrationcheck
make postgres
make createdb
make dbup
````