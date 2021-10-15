
# Robot Apocalypse App
Follow instruction to run this app based on docker.

## Run App

**1. Run Postgres Database**

Application used postgres DB ORM. Following command enable to run postgres in docker container.
`make composer-up`

**2. Migration**

App depends on goose migration tool.

Install goose: `go get -u github.com/pressly/goose/v3/cmd/goose`

From project root directory run:`make migration`

or change DB credentials:

`make DBUSER=postgres  DBPASS=postgres  DBNAME=robot  DBPORT=5432 migration `

**3. Start**

`make start`

**4. Stop**

Stop app: app receives interrupt signal: `ctrl+c`

Stop DB:`make composer-down`

## API Docs

(import postman collection :Robot.postman_collection.json)

**1. Add survivors**

POST http://localhost:8085/api/survivor

Body: [ {    
"name":"Test",   
"gender":"m",    
"longitude":"231241.23432",
"latitude":"34235423.23423",   
"resources":[     
{   
"resource_name":"water",   
"quantity":"10",   
"unit":"litre"   
},         
{     
"resource_name":"medicine",   
"quantity":"10",   
"unit":"count"    
}  
]   
} ]

**2. Update Location**

PATCH http://localhost:8085/api/survivor

Body: {
"id":3,
"longitude":"231241.00001",
"latitude":"34235423.0002"
}

**3. Mark Infected**

POST http://localhost:8085/api/survivor/infected

Body:  {
"id":3,
"infected":true
}

**4. Get Survivors List**

GET http://localhost:8085/api/survivor


**5. Get  Survivor**

GET http://localhost:8085/api/survivor/{id}

**6. Get Robot List**

GET http://localhost:8085/api/robot

**7. Get Total Report**

GET http://localhost:8085/api/report


