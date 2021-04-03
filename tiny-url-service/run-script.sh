docker build -t tiny-url-service
docker run -d --name tiny-url -p 8080:8080 tiny-url-service