package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint16 = 80
    var v3 uint64 = 61
    v4 := GenS1{F0: v2}
    fmt.Println(v2, v2)
    var v5 int32 = -1
    var v6 int32 = (v5 * (v5 * v5))
    switch "foo bar" {
    case "\n\t":
        v3 = (v3 & (v3 << ((541) & 63)))
    }
    switch "" {
    case "a":
        v1[0] = 4294967295
    case "foo bar":
        v7 := GenS1{F0: v2}
        _ = v7
    case "hello":
        switch int(v3) {
        case 2147483647:
            switch "\n\t" {
            case "a":
                v3 = 0
            case "foo bar":
                v8 := []int{42, int(v3)}
                _ = v8
            }
        default:
            var v9 int64 = -782
            _ = v9
        }
    default:
        var v10 uint64 = v3
        _ = v10
    }
    v11 := GenS1{F0: 32767}
    var v12 uint32 = 112
    var v13 bool = (("" + "\n\t") == "hello")
    v14 := GenS1{F0: v2}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
