openssl req -new -nodes -keyout cert.key -out cert.csr -days 3650
openssl x509 -req -days 3650 -in cert.csr -signkey cert.key -out cert.crt

openssl ecparam -list_curves
openssl ecparam -name secp521r1 -genkey -param_enc explicit -out private-key.pem
openssl req -new -x509 -key private-key.pem -out server.pem -days 730

cat private-key.pem server.pem > server-private.pem

