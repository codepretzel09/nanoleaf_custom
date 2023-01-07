# nanoleaf_custom

This code is making an HTTP GET request to the URL "http://192.168.X.X:16021/api/v1/AUTHTOKEN/panelLayout/layout", and then unmarshalling the response body into a struct of type "Response". The "Response" struct has three fields: NumPanels, SideLength, and PositionData. The "PositionData" field is a slice of "PositionData" structs.

The "PositionData" struct has four fields: PanelID, X, Y, and O. The "for" loop iterates over each "PositionData" struct in the "PositionData" slice, and generates a random color for each panel. It then converts the random color values and the PanelID to strings, and makes an HTTP PUT request to the URL "http://192.168.X.X:16021/api/v1/AUTHTOKEN/effects" with a payload containing the panel ID, color values, and some other data.

The function "makeHTTPCall" is used to make the HTTP request with the specified method, URL, and payload. It creates a new HTTP client, creates a new HTTP request with the specified method, URL, and payload, sends the request, and reads the response body. If there are any errors, it will print the error and return.






Compiling and Running Go Code

To compile and run Go code, follow these steps:

Install Go on your machine by following the instructions at https://golang.org/doc/install.

Create a new directory for your Go code and open a terminal in that directory.

Create a go.mod file that defines the module name and the dependencies for your Go code. The go.mod file should look something like this:

``` 
module [module name]

go 1.16

require (
    [dependency] v1.2.3
)
```

Replace "[module name]" with the name of your module and "[dependency]" with the import path of any dependencies your code uses.

Write your Go code in a file with a .go extension. For example, you could create a file called "hello.go" with the following code:

``` package main

import "fmt"

func main() {
    fmt.Println("Hello, world!")
}
```

To compile the code, run the following command in the terminal:

```
go build
```

This will create an executable file in the same directory as your Go code. The name of the file will be the same as the name of the directory that contains your Go code.

To run the compiled code, enter the following command in the terminal:

``` 
./[executable file]
 ```

Replace "[executable file]" with the name of the file that was created when you ran the go build command. For example, if your Go code is in a directory called "hello", the executable file will be named "hello", so you would run the following command:

 ``` 
 ./hello
```

This will execute the Go code and you should see the output "Hello, world!" printed to the terminal.

That's it! You have successfully compiled and run Go code that includes a go.mod file.