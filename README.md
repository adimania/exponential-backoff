# exponential-backoff
Requests any URL, any given number of time with client timeout set to be exponentially high with each iteration. 

Use the binary from the release page or container image `adimania/exponential-backoff`

## How to build binary?
```
git clone https://github.com/adimania/exponential-backoff.git
cd exponential-backoff
go build -tags netgo -a -v exponential.go # note that tags are important for static linking of net package
```

## How to build docker image?
```
git clone https://github.com/adimania/exponential-backoff.git
cd exponential-backoff
docker build .
```

## How to run?
The exponential binary supports a couple of flags. If you miss the flags, appropriate defaults are in place so things should work smoothly. 
```
# ./exponential -help
Usage of ./exponential:
  -URL string
    	URL to open (default "https://httpbin.org/delay/3")
  -iterations int
    	Number of attempts (default 3)
```
To attempt a GET request 5 times for https://httpbin.org/delay/8, execute the following
```
# ./exponential -URL https://httpbin.org/delay/8 -iterations 5
```

For Docker, execute
```
# docker run adimania/exponential-backoff /usr/local/bin/exponential -URL https://httpbin.org/delay/8 -iterations 5
```

Note that both `-URL` and `-iterations` flags are optional. The binary on release page is build for Linux and might not work for other environments. Please use docker container. 

## Is there a circuit breaker?
Yes. If the request succeeds, then the subsequent iternations are not executed. To override this behaviour comment line 40 and rebuild binary (and container).

