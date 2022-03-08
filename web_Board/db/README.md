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

## 주요 Function
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
- SELECT를 제외한 DML(INSERT. UPDATE, DELETEE) 명령을 하기 위해서 sql.DB 객체의 Exec() 메소드를 사용
- 리턴되는 데이터가 있는 경우 Exec() 메소드를 사용해야 함


```go
func (db *DB) Prepare(query string) (*Stmt, error)
```
- 데이터베이스 서버에 Placeholder를 가진 SQL 문을 미리 준비시키는 메소드
- 해당 Statement를 호출 할 때 준비된 SQL문을 빠르게 실행하도록 하는 기법
- sql.Stmt 객체를 리턴받은 후, sql.Stmt 객체의 Exec or Query/QueryRow 메소드를 사용하여 준비된 SQL문을 실행함


  
참고  
https://brownbears.tistory.com/186
