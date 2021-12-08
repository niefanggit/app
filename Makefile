build:
	go build -o bin/main main/main.go
	go build -o bin/app app/app.go
	go build -o bin/web web/web.go

run:
	nohup ./bin/main &
	nohup ./bin/app &
	PORT=80 nohup ./bin/web &