# proglimit
Utility to limit the use of resources by programs for Linux

### Help
Type "./proglimit help" for help.
Command: ./proglimit 1 2 3
1 - target program
2 - number of RAM (in percents)
3 - number of CPU (in percents)
If the limit is not needed, set it to -1
Examples:
./proglimit firefox 10 20 - limit 10% RAM and 20% CPU for firefox
./proglimit java 30 -1 - limit 30% RAM for java

### Compilation
Go version - 1.11
Type "go build" in folder with main.go

### Files
main.go - source file in golang
proglimit - compiled file
