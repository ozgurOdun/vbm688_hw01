# vbm688_hw01

Golang implementation of astar and bfs solution of 8-puzzle.
!!! Needs optimization.
Copying state workaround needs to be fixed. 


## Usage
### Linux & OS X 
`./vbm688_hw01 bfs` or `./vbm688_hw01 astar`

### Windows
`vbm688_hw01.exe bfs` or `vbm688_hw01.exe astar`


## Example Output
for input 
```
7 2 4
5 0 6
8 3 1
```

bfs output is:
```
Alloc = 0 MiB	TotalAlloc = 0 MiB	Sys = 66 MiB	NumGC = 0
bfs
Starting... 2019-04-03 09:18:55.66285 +0300 +03 m=+0.000619205
Raw input is:  7 2 4
5 0 6
8 3 1
_______
[7 2 4]
[5 0 6]
[8 3 1]
-------
_______
[7 2 4]
[0 5 6]
[8 3 1]
-------
_______
[0 2 4]
[7 5 6]
[8 3 1]
-------
...
...
_______
[3 1 2]
[6 4 5]
[0 7 8]
-------
_______
[3 1 2]
[0 4 5]
[6 7 8]
-------
_______
[0 1 2]
[3 4 5]
[6 7 8]
-------
Çözüm Maliyeti: 26
Frontier'e Giren Düğüm Sayısı: 177808
Frontier'den Çıkan Düğüm Sayısı: 171712
End... 2019-04-03 09:18:56.66522 +0300 +03 m=+1.003014568
Process took: 1.002302591s
Alloc = 97 MiB	TotalAlloc = 255 MiB	Sys = 137 MiB	NumGC = 14
```

Astar output is:

```
Alloc = 0 MiB	TotalAlloc = 0 MiB	Sys = 66 MiB	NumGC = 0
astar
Starting... 2019-04-03 09:18:47.367321 +0300 +03 m=+0.000675843
Raw input is:  7 2 4
5 0 6
8 3 1
_______
[7 2 4]
[5 0 6]
[8 3 1]
-------
_______
[7 2 4]
[0 5 6]
[8 3 1]
-------
_______
[0 2 4]
[7 5 6]
[8 3 1]
-------
...
...
_______
[3 1 2]
[4 0 5]
[6 7 8]
-------
_______
[3 1 2]
[0 4 5]
[6 7 8]
-------
_______
[0 1 2]
[3 4 5]
[6 7 8]
-------
Çözüm Maliyeti: 26
Frontier'e Giren Düğüm Sayısı: 2873
Frontier'den Çıkan Düğüm Sayısı: 1809
End... 2019-04-03 09:18:47.376776 +0300 +03 m=+0.010130957
Process took: 9.357782ms
Alloc = 3 MiB	TotalAlloc = 3 MiB	Sys = 68 MiB	NumGC = 0

```
