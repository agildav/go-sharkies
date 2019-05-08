# .env File

Enviroment variables file

## Content

The .env file located at the root of the project must have:

```
# APP Config
APP_ENV = "<environment(development, test, production, etc)>"
PORT = "<port>"

# DB Config - Dev
DB_USER = "<user>"
DB_PASSWORD = "<password>"
DB_HOST = "<host>"
DB_PORT = "<port>"
DB_NAME = "<name>"

# DB Config - Test
TEST_DB_USER = "<user>"
TEST_DB_PASSWORD = "<password>"
TEST_DB_HOST = "<host>"
TEST_DB_PORT = "<port>"
TEST_DB_NAME = "<name>"
```
