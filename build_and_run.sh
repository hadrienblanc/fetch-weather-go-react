echo "# Running Build and run script"
echo "[!] Think about stoping the current hello go app"
docker build -t "hadrienblanc/hello-go" .
docker run -p 8080:8080 -d hadrienblanc/hello-go
docker ps 