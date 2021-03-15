## Description
dịch vụ này sẽ xử lý nghiệp vụ liên quan tới image

hỗ trợ:
- Upload, download hình
- Thêm xóa sửa thông tin của hình

## REQUIREMENT
```
- Rabbitmq
- Mongodb
```

## INSTALL
```bash
#install golang
sudo snap install go --classic
```

## RUN
```bash
go run .
```

## DOC
```bash
// generate doc
//go get -u github.com/swaggo/swag/cmd/swag
//swag init
// access doc
http://localhost/docs/index.html
```

## CONFIG
```.env
NO_SSL_PORT=
ENV=development|production

RABBITMQ_HOST=host
RABBITMQ_PORT=5672
RABBITMQ_USERNAME=
RABBITMQ_PASSWORD=
RABBITMQ_VHOST=/

MONGODB_HOST=host
MONGODB_PORT=27017
MONGODB_USERNAME=
MONGODB_PASSWORD=
```

## Docker
```bash
docker build -t ocr.service.backend:1.0.0 .
```

## TEST (not yet)
unit test
```bash
go test $(go list ./... | grep -v /vendor/ | grep -v /test)
```
e2e test
```bash
go test ./test
```