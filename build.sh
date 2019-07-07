CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/macos64/lianjia lianjia.go
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/macos64/zhilian zhilian.go
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/macos64/clean_status clean_status.go
cp ./config.yaml ./bin/macos64/config.yaml

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o bin/linux64/lianjia  lianjia.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/linux64/zhilian zhilian.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/linux64/clean_status clean_status.go
cp ./config.yaml ./bin/linux64/config.yaml


CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/windows64/lianjia.exe lianjia.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/windows64/zhilian.exe zhilian.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/windows64/clean_status.exe clean_status.go
cp ./config.yaml ./bin/windows64/config.yaml
