#asset-go

`asset` helps golang developers to reference external assets inside binaries.

##The goal

 - Configurable base path for assets
 - Relative paths done right

###What wait! where is the trouble?
Golang does not have a built-in bundle manager like many popular languages do. ie. It does not generate a portable `.app` `.jar` container that includes everything the app needs (binaries and assets).

###Possible solution: Absolute paths
```go
file, err := os.Open("/var/www/image.png")
```

Disadvanges:

 - Your assets and binaries are spread around the system.
 - Tricky for version control
 - No cross platform

###Possible solution: Relative paths
```go
file, err := os.Open("image.png")
```

Disadvanages:

- The binary MUST by executed in the same working path.

```bash
$ /var/bin/myserver
```
does not work. You have to change the working path in order to get it run properly:

```bash
$ cd /var/bin/
$ ./myserver
```



##asset.go solution

```go
file, err := os.Open(asset.Path("image.png"))
```

`asset` knows where the executable is located, so there is not need to change the working path.


##Understanding asset

If the binary is located at `/go/bin/myserver `.

```go
func main() {
   fmt.Println(asset.Path("image.png"))
}
```
```bash
$ cd /root
$ /go/bin/myserver
/go/bin/myserver/image.png

$ myserver
/go/bin/myserver/image.png

$ cd test
$ ./symlink_myserver
/go/bin/myserver/image.png
```

`asset.Path() returns the same path, no matter how the binary is executed.

###Configuring asset
`asset.Configure(DEFAULT_VALUE, INPUT_FLAGS)`

Let's focus in **DEFAULT_VALUE** by the moment.

```go
func init() {
   asset.Configure("resources", asset.NoInput)
   fmt.Println(asset.Path("image.png"))
}
```
```
/go/bin/server/resources/image.png
```

How about an absolute path?

```go
asset.Configure("/var/www/resources", asset.NoInput)
fmt.Println(asset.Path("image.png"))
```
```
/var/www/resources/image.png
```

####INPUT_FLAGS
```go
func init() {
   asset.EnvVariableName = "SERVER_ASSETS"
   asset.Configure("", asset.Environment)
   fmt.Println(asset.Path("image.png"))
}
```
Now you can change the base path with an environment variables:

``` bash
$ export SERVER_ASSETS="/var/myserver/resources"
$ myserver
```
```
/var/myserver/resources/image.png
```

A flag can be also used:

```go
asset.Configure("", asset.Environment | asset.Flag)
```

```
myserver -apath="/myassets"
```
