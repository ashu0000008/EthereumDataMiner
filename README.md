# EthereumDataMiner


CGO_ENABLED=0;GOOS=linux;GOARCH=amd64

go build ./main.go
mv main dataMiner
scp ./dataMiner root@47.98.54.147:/home/dataMiner/dataMiner

