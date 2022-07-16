# metamaskonline_api
api connect to bscscan and line api 

````
1) make poc line api work 

brew install ngrok/ngrok/ngrok
ngrok http 8888
cd poc
cp lineapi.go_ lineapi.go 
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
cp bscscanapi.go_ bscscanapi.go 
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

````
4) run test 
make test
****
then check data in table event , line-event , block-event
****
````

````
5) run live app
go run cmd/chanapi/cron.go
go run cmd/line-webhook/server.go 
ngrok http 8888
````