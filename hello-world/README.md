# Hello World with Golang & Docker


## Simple Run
Build:
```
docker build -t my-golang-app .
```

Run:
```
docker run -it --rm my-golang-app
```


## Remote debug
Build:
```
docker build -t delve-into-docker-app -f Dockerfile.remote .
```

Run:
```
docker run --rm --publish 40000:40000 --publish 1541:1541 --security-opt=seccomp:unconfined --name delve-into-docker delve-into-docker-app
```