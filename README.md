# proglimit
Utility to limit the use of resources by programs for Linux

### Help
Type "./proglimit help" for help.<br>
Command: ./proglimit 1 2 3<br>
1 - target program<br>
2 - number of RAM (in percents)<br>
3 - number of CPU (in percents)<br>
If the limit is not needed, set it to -1<br>
Examples:<br>
./proglimit firefox 10 20 - limit 10% RAM and 20% CPU for firefox<br>
./proglimit java 30 -1 - limit 30% RAM for java<br>

### Compilation
Go version - 1.11<br>
Type "go build" in folder with main.go<br>

### Files
main.go - source file in golang<br>
proglimit - compiled file<br>
