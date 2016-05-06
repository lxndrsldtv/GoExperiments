package main 

import (
    "flag"; "fmt"; "os"; "strings"; "time"; "bufio"
)

const APP_VERSION = "0.1"

// The flag package provides a default help printer via -h switch
var versionFlag *bool = flag.Bool("v", false, "Print the version number.")

func main() {
    flag.Parse() // Scan the arguments list 

    if *versionFlag {
        fmt.Println("Version:", APP_VERSION)
    }
    fmt.Println("Hello", sum(12, 56))
    fmt.Println(echo())
    
    dup_01()
}

func sum(a, b int) int {
	return a + b
}

func echo() string {
	startTime := time.Now()
	args := ""
	for _, arg :=  range os.Args[1:] {
		//fmt.Println(index, ": " + arg)
		args += arg + "; " 
		//fmt.Println(args)
	}
	//fmt.Println(os.Args[:])
	stopTime := time.Now()
	
	fmt.Println("First:", stopTime.Nanosecond() - startTime.Nanosecond())
	
	startTime = time.Now()
	args = ""
	strings.Join(os.Args[1:], "; ")
	stopTime = time.Now()
	
	fmt.Println("Second:", stopTime.Nanosecond() - startTime.Nanosecond())

	return args 
}

func dup_01 () {
	
	counts := make ( map[string]int )
	input := bufio.NewScanner(os.Stdin)
	
	for input.Scan() {
		counts[input.Text()]++
	}
	
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}