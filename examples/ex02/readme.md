 ### Redirect 
 ```
 go run main.go > out_full.txt
 ```
```
$ cat out_full.txt 
Message One: is written to StdOut.
Message Two: is written to StdOut.
```

 ### Pipe 
 ```
 go run main.go | grep One
 ```
```
Message One: is written to StdOut.
```
 ### Pipe & Redirect
 ```
 go run main.go | grep Two > out_two.txt
 ```
```
$ cat out_two.txt 
Message Two: is written to StdOut.
```