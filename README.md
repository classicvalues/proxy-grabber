# Proxy Grabber

#### for now , it just Grabbing HTTP proxies

> Grabbing 5000 proxies & Find active proxies less than 3 minutes (depends on your network and hardware)

### Release
> you can access to latest release version by going to Release section or [Go To Release](https://github.com/amirvalhalla/proxy-grabber/releases)

### Usage

* Install GO >= 1.13
* Build the project
```
go build
```
> depending on your OS (win/linux) it will create proxy-grabber.exe or proxy-grabber

* Run the proxy-grabber (which you already built) by ./proxy-grabber.exe or ./proxy-grabber (in command-line)


### Todo
* add HTTPS proxies (save in different file)
* add SOCKS 5 proxies (save in different file)
* adding Dynamic Host to Reverse Proxy for checking grabbed proxies
* using channel to produce grabbed proxies to find active proxies (consumer)

### Predict
> In spite of the license, I PREDICT that all the examples here are for reference only, and not for criminal (or malicious) purposes.
