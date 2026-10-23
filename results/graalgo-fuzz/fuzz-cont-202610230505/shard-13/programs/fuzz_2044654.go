package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }

func main() {
    v1 := []string{}
    var v2 int64 = -763
    var v3 uint16 = 42
    var v4 uint = 4294967295
    fmt.Println(v2)
    switch -740 {
    case -479:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case 121:
        var v5 int16 = 585
        _ = v5
    case 9223372036854775807:
        v4 -= (v4 * (^(v4)))
    default:
        if len(v1) > 0 {
            v1[0] = "a"
        }
    }
    v4 -= v4
    if false {
        var v6 int8 = -110
        _ = v6
    }
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 uint32 = 65535
        _ = fi7
        _ = v8
    }
    v2 = v2
    var v9 bool = false
    for fi10 := 0; fi10 < 16; fi10++ {
        v3 = 65535
        v11 := GenS1{F0: 233, F1: 2147483648}
        var v12 uint64 = 42
        _ = fi10
        _ = v11
        _ = v12
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
