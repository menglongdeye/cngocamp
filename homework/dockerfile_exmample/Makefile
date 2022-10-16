export tag=v1.0
export servername=httpserver1
export repository=lcl1988
root:
	export ROOT=github.com/cncamp/golang

build:
	echo "building httpserver binary"
	# mkdir -p bin/amd64
	# CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64 .
	# CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o .

release: build
	echo "building httpserver container"
	docker build -t ${repository}/${servername}:${tag} .

push: release
	echo "pushing lcl1988/httpserver"
	docker push ${repository}/${servername}:${tag}