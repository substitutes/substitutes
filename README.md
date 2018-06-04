# Substitutes
Substitutes is an API wrapper and frontend for UNTIS 2018 web.

## Badges

[![codebeat badge](https://codebeat.co/badges/3b86030a-201a-4777-aff6-a5095d4c5958)](https://codebeat.co/projects/github-com-fronbasal-substitutes-master)
[![Go Report Card](https://goreportcard.com/badge/github.com/fronbasal/substitutes)](https://goreportcard.com/report/github.com/fronbasal/substitutes)
[![Maintainability](https://api.codeclimate.com/v1/badges/009d317bc648fadaf7ec/maintainability)](https://codeclimate.com/github/fronbasal/substitutes/maintainability)

## About

The default interface of UNTIS is outdated and should be deprecated. This program utilizes goquery to parse the HTML table to turn it into a JSON API.

## API Endpoints

| Endpoint		| Description				|
| --------		| -----------				|
| /api			| List availible classes		|
| /api/c/{class}	| Show substitutes for a specific class	|

## Installation

You may utilize the Dockerfile to run this program. Please make sure to create/mount the credentials.json file with the credentials of the school.
There is a prebuilt docker image on [Docker Hub](https://hub.docker.com/r/fronbasal/substitutes).

You can also use the go toolchain to run this application using `go get`.

## License

GPL

## Maintainers

- Daniel Malik ([mail@fronbasal.de](mailto:mail@fronbasal.de))

