package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func main() {
	if len(os.Args) == 2{
		if os.Args[1] == "help"{
			fmt.Println("Command: ./proglimit 1 2 3\n1 - target program\n2 - number of RAM (in percents)\n3 - number of CPU (in percents)\nIf the limit is not needed, set it to -1\nExamples:\n./proglimit firefox 10 20 - limit 10% RAM and 20% CPU for firefox\n./proglimit java 30 -1 - limit 30% RAM for java")
			return
		}
	}
	if len(os.Args) < 4{
		fmt.Println("Not enough arguments\nEnter \"./proglimit help\" for help")
		return
	}
	target := os.Args[1]
	mem_limit, e := strconv.Atoi(os.Args[2]); // in percents
	if e != nil{
		fmt.Println("Error. Memory must be number!")
		return
	}
	cpu_limit, e := strconv.Atoi(os.Args[3]); // in percents
	if e != nil{
		fmt.Println("Error. CPU must be number!")
		return
	}
	if mem_limit == -1{
		mem_limit = 9999;
	}
	if cpu_limit == -1{
		cpu_limit = 9999;
	}
	for{
		cmd := exec.Command("ps", "aux", "--sort", "-rss")
		res, _ := cmd.Output()
		strs := strings.Split(strings.Replace(string(res), "\r\n", "\n", -1), "\n")
		for _, str := range strs{
			for strings.Contains(str, "  "){
				str = strings.Replace(str, "  ", " ", -1)
			}
			args := strings.Split(str, " ")
			if(len(args) < 10){
				continue
			}
			exe := strings.Split(args[10], "/")
			if exe[len(exe)-1] == target{
				mem, _ := strconv.Atoi(strings.Split(args[3], ".")[0])
				cpu, _ := strconv.Atoi(strings.Split(args[2], ".")[0])
				if (mem > mem_limit) || (cpu > cpu_limit){
					cmd = exec.Command("killall", target)
					res, _ = cmd.Output()
					if string(res) != ""{
						fmt.Println(string(res))
					}else{
						fmt.Println("Killed: "+args[10])
					}
				}
			}
		}
		time.Sleep(10*time.Second)
	}
}
