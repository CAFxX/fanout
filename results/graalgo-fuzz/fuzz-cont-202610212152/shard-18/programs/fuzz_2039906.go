package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 string }

func main() {
    v1 := []uint32{65535, 2147483648, 2, 398}
    v2 := make(map[uint32]string)
    var v3 int32 = 1
    var v4 int32 = 1
    var v5 uint = 126
    var v6 int16 = 1
    v6 = (v6 * v6)
    delete(v2, 2)
    v7 := GenS1{F0: 2, F1: uint8(v5), F2: ("a" + "foo bar")}
    switch "hello" {
    case "\n\t":
        fmt.Println(v5, v5, v5)
    default:
        fmt.Println(v5, v5, v3)
    }
    var v8 uint = v5
    for fi9 := 0; fi9 < 18; fi9++ {
        v3 = 1
        _ = fi9
    }
    v3 = 1
    for fi10 := 0; fi10 < 3; fi10++ {
        if false {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            var v11 uint64 = 543
            _ = v11
        }
        _ = fi10
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
