# mkrndf
mkrndf — Make random data file


## Usage

- If you want to create a 100 KiB file.

```
$ mkrndf -k 100
=> 100KiB.dat
```

- If you want to create a 100 MiB file.

```
$ mkrndf -m 100
=> 100MiB.dat
```

- If you want to create a 1 GiB file.

```
$ mkrndf -g 1
=> 1GiB.dat
```

- If you want to create a 1000 bytes file.

```
$ mkrndf -b 1000
=> 1000.dat
```


- If you want to create a 100 KiB file with filename.

```
$ mkrndf -k 100 foobar.dat
=> foobar.dat
```
