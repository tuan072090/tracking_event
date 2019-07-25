# Tracking_event service

**Thư viện sử dụng**
1. Gin gonic cho restAPI
2. go-playground để validate
3. và 1 đống thư viện để gen UID

**Run in debug mode**
-     `PORT=1234 go run main`
- Default port is 8080


**Dockerize**
- 1)      `docker build -t go-tracking .`
- 2)      `docker run -d -p 8080:8080 --name tracking go-tracking`
 
