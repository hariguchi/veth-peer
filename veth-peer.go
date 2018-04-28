package main

import (
    "fmt"
    "os"
    "github.com/vishvananda/netlink"
)

func usage() {
    fmt.Fprintf(os.Stderr, "Usage: veth-peer veth-name\n")
    os.Exit(1)
}

func main() {
    var rc int = 1

    if len(os.Args) < 2 {
        usage()
    }
    ifName := os.Args[1]
    if l, err := netlink.LinkByName(ifName); err == nil {
        switch l := l.(type) {
        case *netlink.Veth:
            if idx, err := netlink.VethPeerIndex(l); err == nil {
                if l, err := netlink.LinkByIndex(idx); err == nil {
                    rc = 0
                    fmt.Printf("%s\n", l.(*netlink.Veth).Attrs().Name)
                } else {
                    fmt.Fprintf(os.Stderr,
                        "failed to get link: ifindex %d\n", idx)
                }
            } else {
                fmt.Fprintf(os.Stderr, "failed to get peer: %v\n", err)
            }
        default:
            fmt.Fprintf(os.Stderr, "%v is not veth.\n", ifName)
        }
    } else {
        fmt.Fprintf(os.Stderr, "%v: %v\n", ifName, err)
    }
    os.Exit(rc)
}
