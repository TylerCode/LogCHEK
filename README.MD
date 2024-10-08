# LogCHEK

<p align="center">
  <img width="256" height="256" src="https://raw.githubusercontent.com/TylerCode/LogCHEK/main/src/logo-large.png">
</p>


[![Go](https://github.com/TylerCode/LogCHEK/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/TylerCode/LogCHEK/actions/workflows/go.yml) << See #4

Release: `v0.2.0.0`

Dev(main): `v0.3.1.0`



### General

A simple tool for checking log files. Mostly just wanted to slap something together that saves me a bit of a time every week, figured if it saves someone else 5 minutes every week then that's a bonus!



# Usage

Just open up the loglist.csv and put the path to each log file on a separate line. Use the full path when possible. Should work with network paths such as `\\someserver\c$\windows`

Once you are satisfied with the list, hit save and run the application.

Once it finishes scanning you will either see "All Clear" or you will see a list of files with errors displayed. 

<p align="center">
  <img width="640" height="480" src="https://raw.githubusercontent.com/TylerCode/LogCHEK/main/.github/media/screenshot_0.2.0.0.png">
</p>


## What's next

These might get implemented but I really just needed something quick and dirty for the moment. 

#### Dynamic log filenames

For a simple example, you should be able to put `$date` in the name and have it be replaced with todays date.

#### In-App list editing

Pretty straightforward, just be able to manage the list in the application itself instead. 

It should also NOT fail if the csv doesn't exist, it should just create the csv instead and tell you to update it. 


#### Additional Releases

- MacOS Release (x86)
- Linux ARM Release (should be pretty straightforward) 
- Window ARM Release (need to do some research on getting GCC for ARM for Windows)
- MacOS ARM Release (Will be a while, I don't have an M1 mac)


# Building

This application works on Windows, Mac, and Linux (should work in ARM environments but I have not tested it there)

Also, if you just want to run it, this command should work on any machine that the build tools are installed on:
```bash
go run main.go
```

Also, there are build scripts which mostly work :D

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
go build -ldflags -H=windowsgui -o LogCHEK-GUI.exe main.go
go build -o LogCHEK.exe cli.go
```

to build the application and give yourself an exe. You can use RCEdit to customize it further with an icon and such.

If you DO want to use RCEdit, use the following commands:

```bash
./rcedit.exe LogCHEK-GUI.exe --set-version-string "FileDescription" "Log file checker."
./rcedit.exe LogCHEK-GUI.exe --set-version-string "ProductName" "LogCHEK"
./rcedit.exe LogCHEK-GUI.exe --set-file-version "0.3.0.0"
./rcedit.exe LogCHEK-GUI.exe --set-product-version "0.3.0.0"
./rcedit.exe LogCHEK-GUI.exe --set-icon "logo.ico"
./rcedit.exe LogCHEK.exe --set-version-string "FileDescription" "Log file checker."
./rcedit.exe LogCHEK.exe --set-version-string "ProductName" "LogCHEK"
./rcedit.exe LogCHEK.exe --set-file-version "0.3.0.0"
./rcedit.exe LogCHEK.exe --set-product-version "0.3.0.0"
./rcedit.exe LogCHEK.exe --set-icon "logo.ico"
```


## Linux

For all the dependencies:

```bash
sudo apt-get install -y libgl1-mesa-dev xorg-dev gcc
```

To build!

```bash
CGO_ENABLED=1 go build -o LogCHEK-GUI main.go
CGO_ENABLED=1 go build -o LogCHEK cli.go
```


