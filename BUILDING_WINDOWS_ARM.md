# Building LogCHEK on Windows ARM

This guide provides instructions for building LogCHEK on Windows ARM64 devices (like the Samsung Galaxy Book GO).

## The Problem

When building on Windows ARM, you may encounter errors like:

```
gcc_arm64.S: Assembler messages:
gcc_arm64.S:30: Error: no such instruction: `stp x29,x30,[sp,'
```

This happens because:
- Go correctly detects your ARM64 architecture
- But the standard MinGW from Chocolatey is x86/x64 only
- The x86 assembler cannot understand ARM64 assembly instructions

## Solutions

### Option 1: Use llvm-mingw (Recommended)

llvm-mingw is a MinGW distribution that supports ARM64 Windows.

#### Installation Steps:

1. Download llvm-mingw from: https://github.com/mstorsjo/llvm-mingw/releases
   - Get the latest `llvm-mingw-<version>-ucrt-aarch64.zip` file

2. Extract to a location like `C:\llvm-mingw`

3. Add to PATH (PowerShell):
   ```powershell
   $env:PATH = "C:\llvm-mingw\bin;$env:PATH"
   ```

4. Verify the compiler:
   ```powershell
   aarch64-w64-mingw32-gcc --version
   ```

5. Set the C compiler for Go:
   ```powershell
   $env:CC = "aarch64-w64-mingw32-gcc"
   $env:CXX = "aarch64-w64-mingw32-g++"
   ```

6. Build LogCHEK:
   ```powershell
   cd src
   go build -ldflags -H=windowsgui -o LogCHEK-GUI.exe main.go
   go build -o LogCHEK.exe cli.go
   ```

#### Using the build script:

A build script is provided that automates this process:

```powershell
cd src
.\build_windows_arm.ps1 -LlvmMingwPath "C:\llvm-mingw"
```

### Option 2: Use MSVC (Visual Studio Build Tools)

If you have Visual Studio or Build Tools for Visual Studio installed with ARM64 support:

1. Open "ARM64 Native Tools Command Prompt for VS"

2. Set up Go to use MSVC:
   ```cmd
   set CGO_ENABLED=1
   set CC=cl
   set CXX=cl
   ```

3. Build:
   ```cmd
   cd src
   go build -ldflags -H=windowsgui -o LogCHEK-GUI.exe main.go
   go build -o LogCHEK.exe cli.go
   ```

Note: MSVC integration with Go can be tricky and may require additional configuration.

### Option 3: Cross-compile from x64 Windows

If you have access to an x64 Windows machine or can run x64 emulation:

1. Install standard MinGW on x64 Windows:
   ```powershell
   choco install mingw
   ```

2. Set up for ARM64 cross-compilation:
   ```powershell
   $env:GOOS = "windows"
   $env:GOARCH = "arm64"
   $env:CGO_ENABLED = "1"
   ```

3. Build with the standard build script

Note: Cross-compilation with CGO requires matching cross-compiler tools, which can be complex.

## Troubleshooting

### "gcc: command not found"
- Make sure the llvm-mingw `bin` directory is in your PATH
- Or set `$env:CC` to the full path: `$env:CC = "C:\llvm-mingw\bin\aarch64-w64-mingw32-gcc"`

### "unsupported GOOS/GOARCH pair"
- This shouldn't happen on native ARM64 Windows, but if it does, ensure you're using Go 1.19 or later

### Build is very slow
- The first build downloads dependencies and may take several minutes
- Subsequent builds will be faster due to caching

### Still having issues?
- Verify your Go version: `go version` (should show windows/arm64)
- Verify your compiler: `aarch64-w64-mingw32-gcc --version`
- Check that CGO is enabled: `go env CGO_ENABLED` (should be "1")

## Additional Resources

- [llvm-mingw GitHub](https://github.com/mstorsjo/llvm-mingw)
- [Go CGO documentation](https://pkg.go.dev/cmd/cgo)
- [Fyne cross-compilation guide](https://developer.fyne.io/started/cross-compiling)
