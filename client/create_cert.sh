mkdir -p /var/keys/clients

echo $HOSTNAME >>"/var/hosts/nodes"

echo "subjectAltName=DNS:$HOSTNAME" >>"/var/certs/conf/$HOSTNAME-ext.cnf"

# Create private key and CSR for server
openssl req -newkey rsa:4096 \
  -nodes -keyout /var/keys/clients/$HOSTNAME-server-key.pem \
  -out /var/certs/$HOSTNAME-server-req.pem \
  -subj "/C=DK/L=Copenhagen/O=ITU/OU=Education/CN=$HOSTNAME/emailAddress=alrj@itu.dk" >/dev/null 2>&1

# Create private key and CSR for client
openssl req -newkey rsa:4096 \
  -nodes -keyout /var/keys/clients/$HOSTNAME-client-key.pem \
  -out /var/certs/$HOSTNAME-client-req.pem \
  -subj "/C=DK/L=Copenhagen/O=ITU/OU=Education/CN=$HOSTNAME/emailAddress=alrj@itu.dk" >/dev/null 2>&1

# Sign the server cert
openssl x509 \
  -req -in /var/certs/$HOSTNAME-server-req.pem \
  -days 60 -CA /var/certs/ca/ca-cert.pem \
  -CAkey /var/certs/ca/ca-key.pem \
  -extfile /var/certs/conf/$HOSTNAME-ext.cnf \
  -CAcreateserial -out /var/certs/$HOSTNAME-server-cert.pem >/dev/null 2>&1

# Sign the client cert
openssl x509 \
  -req -in /var/certs/$HOSTNAME-client-req.pem \
  -days 60 -CA /var/certs/ca/ca-cert.pem \
  -CAkey /var/certs/ca/ca-key.pem \
  -CAcreateserial -out /var/certs/$HOSTNAME-client-cert.pem \
  -extensions usr_cert #> /dev/null 2>&1

go run ./client/
