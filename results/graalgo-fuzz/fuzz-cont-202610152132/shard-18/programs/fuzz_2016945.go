package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 string }
func genFn0() int8 {
    v5 := GenS1{F0: 0, F1: 790, F2: "a"}
    switch "a" {
    case "foo bar":
        var v6 int16 = 1
        _ = v6
    default:
        if (0 != 42) {
            var v7 int = 447
            fmt.Println(v7)
            _ = v7
        }
    }
    _ = v5
    return -75
}

func main() {
    var v1 string = ("\n\t" + "")
    v2 := []string{v1, ("foo bar" + v1), v1}
    var v3 int16 = 0
    var v4 bool = true
    v4 = v4
    v4 = false
    fmt.Println(v1, v4)
    var v8 uint64 = 4294967295
    var v9 int8 = 71
    var v10 uint8 = 127
    v3 += (v3 << ((255) & 63))
    switch v1 {
    case "foo bar":
        v11 := GenS1{F0: 65535, F1: 710, F2: ("a" + v1)}
        _ = v11
    case "\n\t":
        v3 = 518
    default:
        fmt.Println(v3)
    }
    v3 = v3
    var v12 int = -9223372036854775808
    var v13 int16 = (v3 & (v3 - v3))
    v13 += (v3 >> ((0) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
