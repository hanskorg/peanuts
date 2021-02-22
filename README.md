## Peanutes 
apis for some web components
---

### components
 - [x] qrcode 二维码 http://localhost:9909/qrcode?i=xxxxx

### Usage

- Start with Docker  
    ```bash  
     docker run --name peanuts -d -p 9909:9909 hansk887/peanuts 
    ```

- Start with code
    ```bash
     go run api.go -log.alsoStdout=1
    ```
-- 