# function-load-tester

function-load-tester is a Golang based command line utility to simulate HTTP load testing based on the Azure Functions Public Dataset.

## Setup

### Build and Compile

To build and compile, run `go build function-load-tester.go`.

Before you run the program, download the (Azure Functions Public Dataset file(~150Mb))[https://azurecloudpublicdataset2.blob.core.windows.net/azurepublicdatasetv2/azurefunctions_dataset2019/azurefunctions-dataset2019.tar.xz] and paste the `invocations_per_function_md.anon.d01.csv` file in this project directory. You can choose another dataset file, but you must specify it with the `-dataPath` flag.

## Usage

To run the program, after building it, use `./function-load-tester`. function-load-tester contains a variety of flags that can be used. The example flags represent default values.

- `-dataPath (string)` - Specifies the name of the dataset CSV file. Ex: `-dataPath=invocations_per_function_md.anon.d01.csv`
- `-functionsCount (int)` - Specifies the amount of functions to read. Ex: `-functionsCount=100`
- `-timeInterval (int)` - Specifies the time interval between batches in seconds. Ex: `-timeInterval=1`
- `-endpoint (string)` - Specifies the HTTP endpoint to hit. Ex: `-endpoint=http://localhost:8080/ping`
- `-timeout (int)` - Specifies the timeout for HTTP requests. Ex: `-timeout=10`

## Results

As function-load-tester runs, it will continually update a `results.csv` with stats from the load testing. This includes:
- `functions_executed`
- `time_elapsed` (in seconds)
- `average_functions_executed` (in functions per second)