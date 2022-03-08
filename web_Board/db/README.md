# Database 연동

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
- 사용 할 DB 드라이버와 해당 DB의 연결 정보를 매개변수로 입력하면, sql.DB 객체를 리턴 시킴
- 리턴 받은 sql.DB 객체를 통해 쿼리문을 실행할 수 있음
- 실제 DB Connection은 Query 등과 같이 실제 DB 연결이 필요한 시점에만 이루어지게 됨

```go
func (db *DB) QueryRow(query string, args ...interface{}) *Row
func (db *DB) Query(query string, args ...interface{}) (*Rows, error)
```
- 데이터베이스 조회(SELECT) 시 두 개의 함수를 사용 할 수 있음
- QueryRow()
  - 하나의 Row만 리턴할 경우 혹은 하나의 Row만 리턴이 될 것을 예상한 경우
- Query()
  - 복수 개의 Row를 리턴할 경우
- 하나의 Row에서 실제 데이터를 읽어 로컬 변수에 할당하기 위해선 Scan() 메소드를 사용
- 복수 Row에서 다음 Row로 이동하기 위해 Next() 메소드를 사용
  
```go
func (db *DB) Exec(query string, args ...interface{}) (Result, error)
```

  
  
참고  
https://brownbears.tistory.com/186
