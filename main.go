package main

import (
  "io"
  "net/http"
  "encoding/hex"
)

func d(x []byte) string {
  r := make([]byte, hex.DecodedLen(len(x)))
  n, err := hex.Decode(r, x)
  if err != nil {
    return ("Error")
  }
  return (string(r[:n]))
}

func f(w http.ResponseWriter, r *http.Request) {
  io.WriteString(w, d([]byte("5374756172742c2064656c697665727920796f752063616e2063656c6562726174652e")))
}

func main() {
  http.HandleFunc("/", f)
  http.ListenAndServe(d([]byte("3a38393832")), nil)
}
