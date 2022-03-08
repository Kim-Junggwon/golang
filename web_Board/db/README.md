# Database 연동 코드

## Package
sql 데이터베이스를 사용하기 위해 표준 패키지 database/sql을 사용
관계형 데이터베이스들에게 공통적으로 사용되는 인터페이스를 제공하고 있음

database/sql 패키지가 지원하는 SQL 종류와 각각의 Driver
- MySQL: https://github.com/go-sql-driver/mysql
- MSSQL: https://github.com/denisenkom/go-mssqldb
- Oracle: https://github.com/rana/ora
- Postgres: https://github.com/lib/pq
- SQLite: https://github.com/mattn/go-sqlite3
- DB2: https://bitbucket.org/phiggins/db2cli

## Function
```go
func Open(driverName, dataSourceName string) (*DB, error)
```
: 오픈

참고  
https://brownbears.tistory.com/186
