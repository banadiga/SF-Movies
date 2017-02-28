# SF Movies

## Setup local/dev environment
Setup virtual environment using vagrant.

The following software is required:

* vagrant (version > **TBD** )
* ansible (version > **TBD** for windows is not needed)
* Virtual box (version > **TBD** )


The following vagrant plugins is required:

* nugrant (version > **TBD** )
* vagrant-hostmanager (version > **TBD** )
* vagrant-docker-compose (version > **TBD** )
* vagrant-guest_ansible (version > **TBD** )

To install plugins run:

```bash
vagrant plugin install <plugin-name>
```

To start local environment run:

```bash
vagrant up
``` 



## Build

**TBD** 

### Build local

**TBD** 

```bash
goleng build
```


### Quick build without tests

**TBD** 

```bash
```


### Build docker image

**TBD** 

```bash
docer build
```



## Deploy

**TBD** 

### Deploy in to local/dev environment

**TBD** 

```bash
```

### Deploy in to AWS/demo environment

**TBD** 


```bash

```



## Public REST API

**TBD** 

### Hello message

Request:

```
GET /welcome/[<user-name>]
```

Responce:

```
status: 200 OK
{
"message": "Hi <user-name>!"
}
```


### Get list

**TBD** 

### Get item

**TBD** 

### Search

**TBD** 



## Links

### Local/dev

* (cAdvisor)[http://sfmovies:8080/] without user/pass
* (kibana)[http://sfmovies:5601/] without user/pass
* (SF Movies)[http://sfmovies:8888/] without user/pass

###  AWS/demo

* **TBD** without user/pass
