go build -ldflags -H=windowsgui -o LogCHEK-GUI.exe main.go
go build -o LogCHEK.exe cli.go
./rcedit.exe LogCHEK-GUI.exe --set-version-string "FileDescription" "Log file checker."
./rcedit.exe LogCHEK-GUI.exe --set-version-string "ProductName" "LogCHEK"
./rcedit.exe LogCHEK-GUI.exe --set-file-version "0.3.0.0"
./rcedit.exe LogCHEK-GUI.exe --set-product-version "0.3.0.0"
./rcedit.exe LogCHEK-GUI.exe --set-icon "logo.ico"
./rcedit.exe LogCHEK.exe --set-version-string "FileDescription" "Log file checker CLI."
./rcedit.exe LogCHEK.exe --set-version-string "ProductName" "LogCHEK"
./rcedit.exe LogCHEK.exe --set-file-version "0.3.0.0"
./rcedit.exe LogCHEK.exe --set-product-version "0.3.0.0"
./rcedit.exe LogCHEK.exe --set-icon "logo.ico"
