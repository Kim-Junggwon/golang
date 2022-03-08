# Go lang - 게시판
![image](https://img.shields.io/badge/-html-red)
![image](https://img.shields.io/badge/-javascript-yellow)
![image](https://img.shields.io/badge/-css-blue)
![image](https://img.shields.io/badge/-bootstrap-blueviolet)
![image](https://img.shields.io/badge/-go-green)
![image](https://img.shields.io/badge/-sqlite3-9cf)


## 코드 구조
![image](https://user-images.githubusercontent.com/94525599/157183333-1a330d99-c36d-46f4-9db9-d15814b562f2.png)

- main.go : 구현된 핸들러를 호출하여 서버 실행
- test.db : 데이터베이스 파일
- db
  - db.go : sqlite3 연동과 기능에 따른 쿼리문 실행
  - db_test.go : db 쿼리문 호출 테스트 코드
- hd
  - hd.go : 라우터 설정과 핸들러 함수 구현
  - hd_test.go : db 연동 테스트 코드
- templates
  - index.html : 인덱스 페이지 "/" 템플릿
  - modify.html : 게시물 수정 페이지 "/modify" 템플릿
  - page.html : 게시물 페이지 "/page" 템플릿
  - write.html : 게시물 작성 페이지 "/write" 템플릿


