# HTTP Availibility Test

### How to use it

1. Insure [git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git) is installed on your system
2. Clone this repo. From the command line, `git clone https://github.com/markhendrix79/availability-test.git` 
3. Change directories into the newly downloaded repo. `cd availability-test`
4. Compiled binary files are provided for the most common use cases. Run the applicable binary for your 
system. For example, if you're using Windows, `.\availabilitytest-win.exe` will launch the program. If 
you're using a Mac with a newer M-series processor, `./availabilitytest-arm-mac` is what you want to use. 
Also provided are binaries for older Macs with Intel processors and for Linux systems with Intel/AMD 
processors.
5. You'll be prompted to enter the path of a yaml file that contains HTTP request data. If you don't have
one handy and want to test the program out, type `testdata.yml` to use the test data included in the repo.
6. The availability percentages of the appropriate domains will be provided and will update every 15 seconds.
7. Once you're done, press `CTRL+C` to exit.

### What it does

This application requests the path to a YAML file from the user. The file should be formatted as a list
with each entry containing a **name** that describes the endpoint to be tested and a **url** to test.
Optionally, the file may also include the **method** of the request (**GET** will be assumed if method
is not provided), **headers**, and **body**.

Once provided with the full path to the YAML file, the program will parse the information and execute 
the HTTP requests described every 15 seconds until the user exits by pressing **CTRL+C**. After each 
iteration, the console will log the cumulative availability percentage of each domain provided. An 
endpoint is considered available if the responses includes a 2xx response code and the response latency
is under 500 ms.
