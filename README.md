# Instructions

## Requirements

  * Golang 1.21 or higher
  * Docker installed
  * Docker compose installed

## Keycloak

Run the dockerized keycloak application under the `kc` folder typing:

  docker-compose run -d

### Configuration

1 - Create the a realm
2 - inside the realm create a client

## Application

First you need to import the library gocloak typing the following command  in the root of the applcation project:

  go get github.com/Nerzal/gocloak/v13
