// # dotxy
// DNS over TLS Proxy
//
// dotxy is a simple DNS over TLS Proxy. It runs as a daemon and listens for DNS queries on on :53. It then redirects these DNS qeuries over TLS to the Secure DNS server(Cloudflare for example).
//
// you can choose to send the DNS queries over TCP or UDP.
//
// over UDP ```dig google.com.com @localhost```
// over TCP ```dig +tcp google.com.com @localhost```

package main
