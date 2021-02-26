# apache-pulsar-poc

Add pulsar-manager user `admin` with password `apachepulsar` with this command:

```bash
CSRF_TOKEN=$(curl http://localhost:7750/pulsar-manager/csrf-token)
curl -H 'X-XSRF-TOKEN: $CSRF_TOKEN' -H 'Cookie: XSRF-TOKEN=$CSRF_TOKEN;' -H "Content-Type: application/json" -X PUT http://localhost:7750/pulsar-manager/users/superuser -d '{"name": "admin", "password": "apachepulsar", "description": "test", "email": "username@test.org"}'
```

Access the pulsar-manager site at http://localhost:9527. Add a new environment with service url http://standalone:8080

Access pulsar dashboard at http://localhost:8081 
