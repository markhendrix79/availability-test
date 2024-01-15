# HTTP Availibility Test

### How to use it

I've compiled this application for Windows, MacOS and Linux. To use it, all you need to do is clone down 
this repo and run the applicable executable from the command line on your system. See examples below.

```
Ex. 
```

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
