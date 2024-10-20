mkdir -p /var/keys/clients

echo $HOSTNAME >> "/var/hosts/nodes"

sleep 2

acc=""

while IFS= read -r line
do
  if [ -z "$acc" ]; then
    acc="DNS:$line"
  else
    acc="$acc,DNS:$line"
  fi
done < "/var/hosts/nodes"
acc="$acc,DNS:$(head -n 1 /var/hosts/hospital)"
echo $acc

echo "subjectAltName=$acc" >> "/var/certs/conf/$HOSTNAME-ext.cnf"


# Create private key and CSR
openssl req -newkey rsa:4096 \
    -nodes -keyout /var/keys/clients/$HOSTNAME-key.pem \
    -out /var/certs/$HOSTNAME-req.pem \
    -subj "/C=DK/L=Copenhagen/O=ITU/OU=Education/CN=$HOSTNAME/emailAddress=alrj@itu.dk" > /dev/null 2>&1

# Sign the cert
openssl x509 \
    -req -in /var/certs/$HOSTNAME-req.pem \
    -days 60 -CA /var/certs/ca/ca-cert.pem \
    -CAkey /var/certs/ca/ca-key.pem \
    -extfile /var/certs/conf/$HOSTNAME-ext.cnf \
    -CAcreateserial -out /var/certs/$HOSTNAME-cert.pem > /dev/null 2>&1

go run ./client/
