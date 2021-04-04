docker build --no-cache  -t tiny-url-service .
docker stop tiny-url
docker rm tiny-url
docker run --network=tinyurl-network -d --name tiny-url -p 8080:8080 tiny-url-service