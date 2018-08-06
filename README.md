# iTree

Find the biggest file / directory on your disk

Same as `tree --sort=size -s/-h`, but more pretty 😉

## Usage

```
Usage of itree:
  -L int
    	level in tree mode (default 1024)
  -h	human readable size (default true)
  -t	tree mode (default true)
```

```
$ itree
├── .git 44KB
├── filetree 4KB
│   ├── tree.go 2KB
│   └── print.go 2KB
├── main.go 559B
└── README.md 118B
```
