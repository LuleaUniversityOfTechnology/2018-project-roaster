Project Roaster
===============
> Automatically grades _your_ code!

[Roast](https://roast.software) analyzes and grades your code. The user can
register and follow their personal programming progress and compare it with
friends! You could compare it to music scrobbling, but for code.

Personal statistics is generated by running static code analysis on the uploaded
code - data such as number of errors, what type of errors and even code style
problems is collected. The statistics can then be viewed in a feed where both
your own and your friends progress is listed.

Global statistics like most common errors and number of rows analyzed is
displayed on a page that can be viewed by everyone.

## Project status
| Build status | Test coverage |
|:------------:|:-------------:|
| [![Build Status](https://travis-ci.org/LuleaUniversityOfTechnology/2018-project-roaster.svg?branch=master)](https://travis-ci.org/LuleaUniversityOfTechnology/2018-project-roaster) | [![Coverage Status](https://coveralls.io/repos/github/LuleaUniversityOfTechnology/2018-project-roaster/badge.svg)](https://coveralls.io/github/LuleaUniversityOfTechnology/2018-project-roaster) |

## Set up Roaster (for developers)
### Prerequisites
 * Go >= v1.11
 * NodeJS w/npm
 * Docker

### Initial setup
We use PostgreSQL and Redis as our databases. The simplest way to set them up is
using the official Docker images.

Run the PostgreSQL Docker image with (exposed at port: `5432`):
```
docker run --name roaster-postgresql -e POSTGRES_PASSWORD=AReallyGreatPassword -d postgres
```

And run the Redis Docker image with (exposed at port: `6379`, non-persistant):
```
docker run --name roaster-redis -d redis
```

Clone the Roaster repository:
```
git@github.com:LuleaUniversityOfTechnology/2018-project-roaster.git
```

And enter the directory:
```
cd 2018-project-roaster
```

### Backend (roasterd)
The backend is written in Go and therefore requires that Go is installed
correctly on your machine. Also, a version >= v1.11 is required for Go modules
which is used in this project.

Make sure you enable Go modules if you are using Go v1.11:
```
export GO111MODULE=on
```

To run the `roasterd` server, simply run:
```
go run github.com/LuleaUniversityOfTechnology/2018-project-roaster/cmd/roasterd
```

### Frontend
The frontend is written in TypeScript, HTML and CSS. Everything is packaged and
compiled using webpack.

First you have to be located in the `www/` folder:
```
cd www/
```

Then, install the required dependencies with:
```
npm install
```

Finally, start the frontend autobuild with:
```
npm start
```

Now, everytime you make any change to the frontend, everything will
automatically recompile and can be accesses from: `http://localhost:5000`
(hosted by the `roasterd` backend).
