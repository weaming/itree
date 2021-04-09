# iTree

Find the biggest file / directory on your disk

Same as `tree --sort=size -s/-h`, but more pretty 😉

## Usage

```
Usage of itree:
  -L int
    	level in tree mode (default 1024)
  -h	human readable size (default true)
  -md5
    	with MD5
  -sha256
    	with SHA256
  -t	tree mode (default true)
```

```
> itree
├── .git 39KB
├── filetree 5KB
│   ├── print.go 2KB
│   ├── tree.go 2KB
│   └── hash.go 289B
├── main.go 879B
├── README.md 478B
└── go.mod 41B

> itree -md5
├── .git 39KB
├── filetree 5KB
│   ├── print.go 2KB 980c0231de0539926ee9417a5689488b
│   ├── tree.go 2KB e594fb1930979aba390ce4735dc81ce7
│   └── hash.go 289B 3f72c3556b9323cd1d22d262b0df7130
├── main.go 879B 073ad902e4da6b37820ef908ffea434c
├── README.md 478B 1221c2e389b21a77e35145f24f320220
└── go.mod 41B a072bc77668ba6a6c95f0a272097d2a0
```
