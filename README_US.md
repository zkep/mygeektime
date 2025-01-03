English | [中文](./README.md)

### my geektime
This is a tool to obtain the geektime  articles docs

---

### [Docs](https://zkep.github.io/mygeektime/)


### [Show Time](https://mygeektime.anyfun.tech)


#### Install

#### install with docker compose

```shell
git clone https://github.com/zkep/mygeektime.git

cd mygeektime/docker

# download default data
# wget https://github.com/zkep/mygeektime/releases/download/v0.0.1/tasks.sql -O mysql/init/tasks.sql

docker-compose up -d
```

browser web url:  http://127.0.0.1:8090

#### install with  docker
```shell
docker run  -d  --restart always \
--name mygeektime  \
-p 8090:8090 \
zkep/mygeektime:latest  server 
```
browser web url:  http://127.0.0.1:8090

#### docker with config
```shell
docker run  -d  --restart always \
--name mygeektime  \
-p 8090:8090 \
-v config.yml:/config.yml \
zkep/mygeektime:latest  server --config=config.yml
```
browser web url:  http://127.0.0.1:8090

##### docker with specify download directory
replace ${directory} with you local directory
```shell
docker run  -d  --restart always \
-p 8090:8090 \
--name mygeektime \
-v config.yml:/config.yml \
-v ${directory}:/repo  \
-v config.yml:/config.yml \
zkep/mygeektime:latest  server --config=config.yml 
```
browser web url:  http://127.0.0.1:8090



#### install with golang
```shell
go install github.com/zkep/mygeektime@latest
```
#### Start web service

##### Default configuration to start HTTP service
```shell
mygeektime server
```

##### Customize configuration to start HTTP service
```shell
# Use custom configuration templates
mygeektime server --config=config.yml
```

#### [FFmpeg](https://ffmpeg.org/download.html)

MacOS
```shell
brew install ffmpeg        
```
Linux
```shell
git clone https://github.com/FFmpeg/FFmpeg.git ffmpeg

cd ffmpeg

./configure --enable-gpl --enable-libx264

make && make install
```

### Simulate user login：

> Method 1: The browser developer tool retrieves a valid Geektime cookie


#### thanks
* [gin](https://github.com/gin-gonic/gin)
* [amis](https://github.com/baidu/amis)
* [gorm](https://github.com/go-gorm/gorm)
* [FFmpeg](https://ffmpeg.org/download.html)
