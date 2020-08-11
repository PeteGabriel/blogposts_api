
### The API expects the following environment variables:

* _DATABASE_URI_ - postgres db connection string. 


### Run tests:

You can use the PSQL docker image made for this purpose [here](https://hub.docker.com/repository/docker/petegabriel/microblog_psql/general). 

Also, something like `postgresql://gopher:myscretpassword@localhost:5432/microblog` as de DB connection string.
  
`go test personalblog ENV_PATH=../.env`


 


