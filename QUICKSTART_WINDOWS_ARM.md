# Quick Start: Windows ARM Build

If you're getting assembly errors when building on Windows ARM, follow these steps:

## 1. Install llvm-mingw

Download from: https://github.com/mstorsjo/llvm-mingw/releases

Get: `llvm-mingw-*-ucrt-aarch64.zip` (latest version)

Extract to: `C:\llvm-mingw`

## 2. Run the build script

Open PowerShell in the `src` directory:

```powershell
.\build_windows_arm.ps1 -LlvmMingwPath "C:\llvm-mingw"
```

## 3. Done!

You should now have:
- `LogCHEK-GUI.exe` - GUI version
- `LogCHEK.exe` - CLI version

## Troubleshooting

If the script fails, make sure:
- Go is installed (`go version`)
- llvm-mingw is extracted to the correct location
- You're running PowerShell (not CMD)

For more details, see [BUILDING_WINDOWS_ARM.md](BUILDING_WINDOWS_ARM.md)
