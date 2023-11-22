Run DBs:
1. Create .mysql.env
   ```
   MYSQL_ROOT_PASSWORD='some_pass'
   MYSQL_ROOT_HOST='%'
   ```
2. Create .psql.env
   ```
   POSTGRES_PASSWORD='some_pass'
   ```
3. ```
   docker compose up -d
   ```