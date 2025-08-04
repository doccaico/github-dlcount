# github_dlcount

## Dependencies
- [jq](https://github.com/jqlang/jq)
- curl

## Usage
```sh
$ github_dlcount vim/vim-win32-installer

gvim_9.1.1591_arm64.exe           588
gvim_9.1.1591_arm64.zip           427
gvim_9.1.1591_x64.exe            3964
gvim_9.1.1591_x64.zip             979
gvim_9.1.1591_x64_pdb.zip         334
gvim_9.1.1591_x86.exe            1430
gvim_9.1.1591_x86.zip             598
gvim_9.1.1591_x86_pdb.zip         331
```

## Build
```sh
# Windows
$ build.bat

# WSL2 (or Linux Native?)
$ ./build.sh
```
