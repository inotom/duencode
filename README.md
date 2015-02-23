# duencode

data uri scheme encoding tool

## Installation

```
$ go get -u github.com/inotom/duencode
```

## Usage

```
$ duencode sample.jpg
data:image/jpg;base64,/9j/4AAQSkZJRgABAQ...
```

-p flag is outputting plain format. (without data uri scheme)

```
$ duencode -p sample.jpg
/9j/4AAQSkZJRgABAQAAAQABAAD//gA7Q1JFQVRP...
```

## License

duencode is licensed under the MIT

## Author
iNo (inotom)
