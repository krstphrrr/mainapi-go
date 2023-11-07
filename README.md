## Restful API in GO for tall data

# todo
- finish complete implementation of gap endpoint with all the logic currently in nest. ( parsing params, like operators, pagination etc.)
- aws
- create clean pg tenant switch with no orm bloat
- follow repository pattern for app structure
- implement goroutines for concurrent reqs that may take a while 
- handling uncommon types like postgis geometries and dates

# currently 
- created postgres interface ( requires localhost postgres instance with tall data to be running)
- created gap interface

- `docker compose build` will build a docker image with all the external packages into a container

- `docker compose up` will serve the go server on `8081`

- store currently connects to a containerized pg instance on a attachable overlay network. the host on the connection string uses the overlay network alias to find the db. 

- get request to `/gap` will currently return empty json with the correct fields. 

