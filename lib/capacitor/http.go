package capacitor

import (
  "net"
  "net/http"
  "time"
)

// https://medium.com/@nate510/don-t-use-go-s-default-http-client-4804cb19f779

var netTransport = &http.Transport{
  Dial: (&net.Dialer{
    Timeout: 5 * time.Second,
  }).Dial,
  TLSHandshakeTimeout: 5 * time.Second,
}

var httpClient = &http.Client{
  Timeout: time.Second * 10,
  Transport: netTransport,
}
