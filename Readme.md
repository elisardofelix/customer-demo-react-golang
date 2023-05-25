# Customer Demo App

## Requirements

Is required to run the application the following parts:

* **Docker**
* **docker-compose**
* **make** Is required to runs the one line commands available.

**NOTE: This documentation supposes that you already have the installation of docker and docker-compose**

## Repository Structure

There are only two folders at the root of the repository which are the following:

```
Repository
│
└───api
└───docker-compose
└───frontend-app
```

The **api folder** has the laravel API with all the required files and the **frontend-app folder** contains the NodeJS client app in React that consume the Golang API. The **docker-compose** folder contains any required file for the initialization or configuration of the containers.


### Make Installation

If you do not have **make** installed in you computer you can use the following command depending of the OS.

**Windows**

In case of Windows I recommend [chocolatey](https://community.chocolatey.org/) to install just run the following command in Administrator access of PowerShell:

`Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))`

And then just install make with the following command:

`choco install make`

**MAC OS**

In case of Mac just install [brew](https://brew.sh/index_es) with this command:

`/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"`

And then just install make with the following command:

`brew install make`

**Linux**

For Linux is just install make with the package manager installed, here are some examples:

In case of Debian or Ubuntu:

`sudo apt-get install make`

In case of Arch linux:
`sudo pacman -Sy make`

## How to run the application

There are multiple commands available for running the multiple options when you are located in the root folder of the project:

* `make local`: This will allow you to set up the local environment with all the required dependencies.

* `make test-front-end`: To run all the unit testing for the front-end.

* `make test-back-end`: To run all the unit testing for the back-end.

## How to use the application

After both parts of the application are running you just need to go with the internet browser to the address http://127.0.0.1:3000/ which is the one pointing at the front-end.

The Front-End App is just a simple list with all the customers with a hyperlink in each Id in the Id column that goes to the details screen that will show all the information stored in the API of the specific customer clicked.
