# quotes

Minimal Golang API that returns random quotes from JSON file

## Run in Docker

```
cd <../quotes>
docker build -t <dockername>/quotes:<version> .

docker run <dockername>/quotes:<version> -d -p 8080:8080
```

## Get a random quote

get container address

```
docker inspect <container_id>
```

```
...
 "Networks": {
                "bridge": {
                    "IPAMConfig": null,
                    "Links": null,
                    "Aliases": null,
                    "NetworkID": "b8e2246e80dcfbda18bc8265010d60f62da2900d60a0f7b2da0355191d638632",
                    "EndpointID": "01d3fa2202e9b3d3d812445a691031a16103e83beb64a619de490a42d18ea1e8",
                    "Gateway": "172.17.0.1",
                    "IPAddress": "172.17.0.2",
                    "IPPrefixLen": 16,
                    "IPv6Gateway": "",
                    "GlobalIPv6Address": "",
                    "GlobalIPv6PrefixLen": 0,
                    "MacAddress": "02:42:ac:11:00:02",
                    "DriverOpts": null
                }
            }
...
```

-> 172.17.0.2

```
curl 172.17.0.2:8080/quote
```

## Push image to Dockerhub

```
docker login

docker push <username>/quotes:version
```