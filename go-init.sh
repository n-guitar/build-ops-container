echo "init start"
echo "####################"

echo "> goenv local 1.16.4"
goenv local 1.16.4

echo "> go version"
go version

echo "go mod init"
go mod init github.com/n-guitar/build-ops-container

echo "go mod tidy"
go mod tidy

echo "####################"
echo "finish!!"
