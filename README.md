# NAME

veth-peer -- Show the peer interface name of the specified veth

# Synopsys

veth-peer interface-name

# Description

The veth-peer command shows the peer interface name of the
specified veth interface.

# Example

```
$ ip link add foo-bar type veth peer name bar-foo
$ veth-peer foor-bar
bar-foo
$ veth-peer xxx
xxx: Link not found
$ veth-peer lo
lo is not veth.
$
```

