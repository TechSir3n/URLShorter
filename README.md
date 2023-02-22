# Shortener URL 

![image](https://github.com/Kl1ck9r/ShortenerURL/blob/main/screenshots/url-shorter.png)

# Requirements
* **Go:** `20.0`
* **Redis:** `6.0.17`


## docker-compose
Server is ready immediately after containers are up
```shell
 make compose-up        #docker-compose up
```

## Dockerfile 
```shell
make docker-build   #docker build -t url-shortener .
make docker-run     #docker run url-shortener
```

## Build / Run

```shell
git clone https://github.com/Kl1ck9r/ShortenerURL.git
cd url-shortener
make build
make run 
```
