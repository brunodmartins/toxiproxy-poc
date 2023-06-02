# toxiproxy-poc
A POC to test toxiproxy targetting an external web server

## How to run it?

First, we should start the toxiproxy. We can use a Docker container for it

```bash
docker run -p 8474:8474 -p 6379:6379 -p 1234:1234 --rm -it ghcr.io/shopify/toxiproxy
```
- Port 8474 is the default one used by toxiproxy to operate
- Port 1234 should be used as our main proxy to an external webserver

At our application, we define the toxiproxy usage:

```go
tx_client := toxiproxy.NewClient("localhost:8474")
```

And then, add as many proxies as we need

```
proxy, err := tx_client.CreateProxy("viacep", "0.0.0.0:1234", "viacep.com.br:80")
proxy.AddToxic("", "latency", "", 1, toxiproxy.Attributes{
		"latency": 1000,
	})
```

Change the toxic values to play!

## Doubts?

Open an Issue
