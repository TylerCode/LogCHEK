# LogCHEK

A simple tool for checking log files. More functionality to come but wanted to make a small application cause I'm lazy. 

Also wanted to try GO. 

v0.1.0.0



# Building

This application works on Windows, Mac, and Linux (should work in ARM environments but I have not tested it there)

## Windows

You will need the `mingw` compiler tools which can be procured through the chocolatey package manager (the easiest way, just do it)

```bash
choco install mingw
```

You may need to add gcc to path but iirc this was done for me. 

```bash
gcc --version
```

That command will tell you if you're good to go. 

Finally, you can use:

```bash
go build -ldflags -H=windowsgui -o LogCHEK.exe main.go
```

to build the application and give yourself an exe. You can use RCEdit to customize it further with an icon and such.



## Linux

I just usually use 

```bash
go run main.go
```

but it can be built using

```bash
go build -ldflags -o LogCHEK main.go
```