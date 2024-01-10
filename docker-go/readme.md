## docker-compose
```bash
docker-compose up -d

docker ps -a
```


## docker build & docker run
```bash
docker build -t dcoer-go-web:v1 .
docker save dcoer-go-web:v1 | gzip > dcoer-go-web_v1.tgz
docker load dcoer-go-web_v1.tgz
```