mkdir -p /var/keys/clients
# 2. Generate web server's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout /var/keys/clients/$HOSTNAME-key.pem -out /var/certs/$HOSTNAME-req.pem -subj "/C=DK/L=Copenhagen/O=ITU/OU=Education/CN=*.itu.dk/emailAddress=alrj@itu.dk" > /dev/null 2>&1

# 3. Use CA's private key to sign web server's CSR and get back the signed certificate
openssl x509 -req -in /var/certs/$HOSTNAME-req.pem -days 60 -CA /var/certs/ca/ca-cert.pem -CAkey /var/certs/ca/ca-key.pem -CAcreateserial -out /var/certs/$HOSTNAME-cert.pem > /dev/null 2>&1
#-extfile /var/certs/clients/$HOSTNAME-ext.cnf

go run ./server/
#echo "Server's signed certificate"
#openssl x509 -in server-cert.pem -noout -text
