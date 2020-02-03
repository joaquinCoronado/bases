# Numeric Bases Converter

Convert any number with base between 2 to 30 to another number at the same bases range. 

## Getting Started

Only install Go 1.13 or lates version and clone out of the GOPATH.

### Prerequisites

* [GO](https://golang.org/doc/install) - Make a typical Installation of Go

### Installing

Install the modules of the project

```
-> go install
```

Create build at root directory of the project

```
-> go build
```

Test the convert command executing the done build.

```
-> ./bases convert 101 -b 2 --to 10
```

#### Optional ( I prefer this )`

Add of the path the location of the done build; in the file `.bash_profile` or `.zshrc` paste
the next line with the path of the done build

```
export PATH=$PATH:{ the path of your build }/bases

## Example
export PATH=$PATH:/Users/JoaquinCoronado/GoProjects/bases
```

Then run the `base` command any place at your terminal
```
-> bases convert 101 -b 2 --to 10
```


## Built With

* [Go](https://golang.org/) - The principal language of the project
* [Go Modules](https://golang.org/) - Dependency Management
* [Cobra](https://github.com/spf13/cobra) - Library to create powerful modern CLI similar GIT.

## Authors

* **Joaquin Coronado** - *Initial work* - [joaquinCoronado](https://github.com/joaquinCoronado)


## License

This project is licensed under the Apache License - see the [LICENSE.md](LICENSE.md) file for details
