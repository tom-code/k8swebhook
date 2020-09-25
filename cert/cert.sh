openssl genrsa  -out ca.key 2048
openssl req -x509 -new -nodes -key ca.key -sha256 -days 1825 -out ca.pem

openssl genrsa -out hook.key 2048
openssl req -new -key hook.key -out hook.csr

openssl x509 -req -in hook.csr -CA ca.pem -CAkey ca.key -CAcreateserial  -out hook.crt -days 825 -sha256
