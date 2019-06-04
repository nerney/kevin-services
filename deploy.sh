set -e 
rm -rf ./bin
export GO111MODULE=on
export GOOS=linux
go build -ldflags "-s -w" -o bin/externalQuotesEtl lambdas/external-quotes-etl/main.go
go build -ldflags "-s -w" -o bin/quotesCrud lambdas/quotes-crud/main.go
sls deploy
rm -rf ./bin

