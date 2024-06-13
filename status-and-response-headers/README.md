This code WRITES to the RESPONSE the STATUS LINE and some RESPONSE HEADERS:

REQUEST LINE
GET / HTTP/1.1
method SP request-target SP HTTP-version CRLF
https://tools.ietf.org/html/rfc7230#section-3.1.1

RESPONSE (STATUS) LINE
HTTP/1.1 302 Found
HTTP-version SP status-code SP reason-phrase CRLF
https://tools.ietf.org/html/rfc7230#section-3.1.2

