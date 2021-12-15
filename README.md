# secgen
Generate secure token

# Usage

```
$ go get github.com/dongri/secgen

$ secgen -h
Usage of 
  -e string
    	"e" encoding (ex: base64)
  -l int
    	"i" length (default 10)
  -q int
    	"q" quantity (default 1)

$ secgen
*v[h%~29(b

$ secgen -l 10 -q 5
E:(#y]@_2b
$WNvEcD#Bb
wUS_l~%iOr
dLX&(VOf89
(;E-YDuqCq

$ secgen -l 10 -q 5 -e base64
I0sxXVh1fj1eZQ==
Ok0xcF1KJCQkKQ==
YT1iZEpbMTd+Mw==
TlFTTWdpb0JkUQ==
WChENXJ+RE9DcA==
```
