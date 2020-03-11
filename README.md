# Potatosync
This is a custom API for supporting the PotatoNotes app, which can be found [here](https://github.com/HrX03/PotatoNotes/) on GitHub, and can be downloaded from the play store. If you feel like contributing, feel free to send a PR our way.


[![Actions Status](https://github.com/ATechnoHazard/potatosync/workflows/CI/badge.svg)](https://github.com/ATechnoHazard/potatosync/actions) ![](https://img.shields.io/docker/pulls/atechnohazard/potatosync)
## Building
The project has a Makefile conveniently set up for building the project and its docker image. We use scratch containers and add only the binary, instead of building the binary inside the container. This is done to reduce image size.

To build the project, you must first have the following dependencies -
* `make`
* `Go version 1.11 and onwards`
* `docker`

To install all package dependencies, run `make dep`.

To build the binary, run `make build`

To build the docker image with the binary, run `make img-build`

*Note: You MUST have built the binary to be able to build the docker image.*

## Self-Hosting
In order to host the api yourself you need to follow these steps:

* Install Docker https://docs.docker.com/install/ 
* Install Docker-Compose https://docs.docker.com/compose/install/
* Download the docker-compose.yml
```
wget https://raw.githubusercontent.com/ATechnoHazard/potatosync/master/docker-compose.yml
```
* Run docker-compose in the directory where you downloaded the file
```
sudo docker-compose up
```
* You should now be up and running!

## Contributing
We are grateful for any and all contributions, so feel free to send a PR our way! Just remember to `gofmt` often, and document as much as you can.

## Links
<a href='https://play.google.com/store/apps/details?id=com.potatoproject.notes&hl=en&pcampaignid=pcampaignidMKT-Other-global-all-co-prtnr-py-PartBadge-Mar2515-1'><img alt='Get it on Google Play' src='https://play.google.com/intl/en_us/badges/static/images/badges/en_badge_web_generic.png'/></a>