# go-vaccel-docs
Documentation for vAccel Golang Bindings

## vAccel Installation
First of all, a vAccelRT installation is required before proceeding to the next sections. 

### Build from source
In Ubuntu-based systems, you need to have the following packages to build vaccelrt:
1. cmake
2. build-essential

You can install them using the following command:
```
sudo apt-get install -y cmake build-essential
```
Get the source code for **vaccelrt**:
```
git clone https://github.com/cloudkernels/vaccelrt --recursive
```
Prepare the build directory:
```
cd vaccelrt
mkdir build
cd build
```
### Building the core runtime library
```
# This sets the installation path to /usr/local, and the current build
# type to 'Release'. The other option is the 'Debug' build
cmake ../ -DCMAKE_INSTALL_PREFIX=/usr/local -DCMAKE_BUILD_TYPE=Release -DBUILD_EXAMPLES=ON -DBUILD_PLUGIN_EXEC=ON -DBUILD_PLUGIN_NOOP=ON
```
```
make
sudo make install
```
## vAccel-Go Bindings Installation

### Go Installation
Of cource, prior to installing the bindings, we have to make sure that Golang 1.20 or newer is installed in our system. We can check this using the following command:
```
go version
```
Otherwise, `go 1.20` needs to be installed. You can find instructions on how to install Go [here.](https://go.dev/doc/install) 

### Build the Bindings from source
Download the source code:
```
git clone https://github.com/nubificus/go-vaccel.git
```
And you can build the examples. 
```
# Set vaccel location
export PKG_CONFIG_PATH=/usr/local/share/

cd go-vaccel
make all
```
Now you have succesfully built some vaccel programs using Go. The executables are located in go-vaccel/bin. You can run the `noop` example:
```
export VACCEL_BACKENDS=/usr/local/lib/libvaccel-noop.so
./bin/noop
```
Or the `exec` example, providing a path for the shared object and an integer:
```
export VACCEL_BACKENDS=/usr/local/lib/libvaccel-exec.so
./bin/exec /usr/local/lib/libmytestlib.so 100
# if everuthing go as expected, the
# plugin will probably double the integer 
```
## Write your own programs
Of course, now you can use the vaccel package in your Go programs:
```go
import "github.com/nubificus/go-vaccel/vaccel"
```
### Create a session
```go
var session vaccel.Session

err = vaccel.SessionInit(&session, 0)

if err != 0 {
  [...]
}
```

### Create argument lists
```go
var read, write vaccel.ArgList

numArgs := ...;
read  = vaccel.ArgsInit(numArgs)
write = vaccel.ArgsInit(numArgs)

if read == nil || write == nil {
  [...]
}
```

### Add an integer as input
```go
input := 10
buf   := unsafe.Pointer(&input)
size  := unsafe.Sizeof(input)

err := read.AddSerialArg(buf, size)

if err != 0 {
  [...]
}
```
### Define an Expected Argument
```go
var output int
buf  = unsafe.Pointer(&output)
size = unsafe.Sizeof(output)

err = write.ExpectSerialArg(buf, size)

if error != 0 {
  [...]
}
```

### Run a vAccel Operation
```go
err = vaccel.SomeOp(&session, ... , read, write)

if err != 0 {
  [...]
}
```

### Extract the Arguments
```go
/* Extract an argument (eg an integer) */
idx := 0
outputBuf := write.ExtractSerialArg(idx)

/*
   vacccel plugins are implemented
   in C, therefore you should convert
   the result into a Go integer.
*/
val := C.uint(*((*int)(outbuf)))
fmt.Println("Output: ", val)
```

### Delete the Lists


```go
if write.Delete() != 0 || read.Delete() != 0 {
  [...]
}
```

### Delete the Session
```go
if vaccel.SessionFree(&session) != 0 {
  [...]
}
```

## Working with Non-Serialized Arguments
In vaccel, you can work with non-serialized arguments, too. When working with this type of arguments (ie structures that contain pointers), you just need to pass a function that handles the data. When adding an argument in a list, make sure you provide a function that serializes the data, so that the argument can be safely transferred remotely. Correspondingly, when you extract a non-serialized argument, you should provide a function that deserializes the data, in order to retrieve the "actual" structure from the serialized sequence. Those functions must follow a specific signature, which is presented below:
```go
/* Let's say you want to handle the following structure */
type MyData struct {
	Size uint32
	Arr []uint32
}

/*
 * Function that serializes an instance of MyData:
 *
 * nonSerial := MyData{Size: 3, Arr: {7,8,9} }
 * serialize(nonSerial) -> (3 7 8 9 , 16)	
*/
func Serialize(buf unsafe.Pointer) (unsafe.Pointer, uint32) {

	mydata := (*MyData)(buf)
	serialBuf := make([]uint32, mydata.Size + 1)
	serialBuf[0] = uint32(mydata.Size)

	var i uint32
	for i=0; i<mydata.Size; i++ {
		serialBuf[i + 1] = mydata.Arr[i]
	}

	retBuf := unsafe.Pointer(&serialBuf[0])
	bytes  := (mydata.Size + 1) * 4

	return retBuf, bytes
}

/* Function that constructs an instance of MyData out of serialized data */
func Deserialize(buf unsafe.Pointer) unsafe.Pointer {

	sizeExtr := *((*uint32)(buf))

	/* Convert unsafe.Pointer to Slice */
	var slice []uint32
	header := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	header.Data = uintptr(buf)
	header.Len  = int(sizeExtr + 1)
	header.Cap  = int(sizeExtr + 1)
	
	/* Reconstruct the structure */
	mydatabuf := new(MyData)
	mydatabuf.Size = sizeExtr
	mydatabuf.Arr = make([]uint32, sizeExtr)

	var i uint32
	for i=0; i<sizeExtr; i++ {
		mydatabuf.Arr[i] = slice[i + 1]
	}
	
	return unsafe.Pointer(mydatabuf)
}
```
## Usage with vAccel Functions
### Add a Non-Serialized Argument
```go
var myDataInput MyData = ... 

err = read.AddNonSerialArg(&myDataInput, 0, Serialize)

if err != 0 {
	[...]
}
```
### Extract a Non-Serialized Argument
```go
outbuf := write.ExtractNonSerialArg(0, Deserialize)

if outbuf == nil {
	[...]
}

mydata_out := (*MyData)(outbuf)
```
Therefore, non-serialized arguments can be used with vaccel functions in a similar manner as the serialized ones. The user needs just to provide proper serializer and deserializer functions. 
