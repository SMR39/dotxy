# dotxy
DNS over TLS Proxy

dotxy is a simple DNS over TLS Proxy. It runs as a daemon and listens for DNS queries on on :53. It then redirects these DNS queries over TLS to the Secure DNS server (Cloudflare for example).

you can choose to send the DNS queries over TCP or UDP.

`make build` will build the docker image \n
Run in either udp or tcp mode
`docker run -p 53:53 -p 53:53/udp --net=bridge -it dotxy:latest  dotxy listen -n udp`
`docker run -p 53:53 -p 53:53/udp --net=bridge -it dotxy:latest  dotxy listen -n tcp`

On your local machine run the client
over UDP `dig google.com.com @localhost` \n
over TCP `dig +tcp google.com.com @localhost`  \n

you can tail the logs from the container while running the client DNS queries using dig.
