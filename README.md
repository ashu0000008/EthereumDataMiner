# EthereumDataMiner


CGO_ENABLED=0;GOOS=linux;GOARCH=amd64
GOPROXY=https://goproxy.cn,direct;CGO_ENABLED=0;GOOS=linux;GOARCH=amd64
GOPROXY=https://goproxy.cn,direct;GOOS=linux;GOARCH=amd64;CGO_ENABLED=0

go build ./main.go
mv main dataMiner
scp ./dataMiner root@47.98.54.147:/home/dataMiner/dataMiner



GOPROXY=https://goproxy.cn,direct;GOOS=linux;GOARCH=amd64;CGO_ENABLED=0


