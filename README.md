# NAME

veth-peer -- Show the peer (pair) interface name of the specified veth interface

# Synopsys

veth-peer interface-name

# Description

The veth-peer command shows the peer interface name of the
specified veth interface.

# Example

```
$ sudo ip link add foo-bar type veth peer name bar-foo
[sudo] password for XXX:
$ veth-peer foor-bar
bar-foo
$ veth-peer xxx
xxx: Link not found
$ veth-peer lo
lo is not veth.
$ sudo ip link del foo-bar
           $ veth-peer bar-foo
           bar-foo: Link not found
$
```

