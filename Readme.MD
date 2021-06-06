<h2>Docker way</h2>
For starting application please type 
`docker-compose up --build`

Afterwards => you need to restore database schema. SQL file is in folder mysql-dump -> nba.sql

After performing this actions navigate to `localhost:9000` and slick start button

<h2>No Docker Way</h2>
If you want to test if without Docker please do:

`go get github.com/satori/go.uuid` <br>
`go get github.com/go-sql-driver/mysql`

Change mysql host in main.go file<br>
FROM:<br>
`conn, err := sql.Open("mysql", "root:@tcp(db:3306)/nba")`<br>
TO:<br>
`conn, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/nba")`<br>

Where `127.0.0.1` is host of your local DB and `3306` its port

And run `go run cmd/*.go`

But before do not forget to import dump file from the mysql-dump folder
