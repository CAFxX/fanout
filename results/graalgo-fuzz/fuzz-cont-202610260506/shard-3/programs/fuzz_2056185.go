package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }

func main() {
    var v1 uint64 = 2
    var v2 int32 = -2147483648
    var v3 bool = (true || false)
    var v4 string = "a"
    var v5 int64 = 650
    v1 -= v1
    if (v3 || v3) {
        if (!true) {
            v6 := GenS1{F0: 134, F1: v1}
            _ = v6
        }
        for fi7 := 0; fi7 < 8; fi7++ {
            var v8 uint64 = v1
            _ = fi7
            _ = v8
        }
    }
    var v9 uint16 = 652
    v1 += 18446744073709551615
    switch 1 {
    case 32767:
        var v10 uint64 = 4294967295
        _ = v10
    }
    fmt.Println(v3)
    v2 -= 0
    v11 := []int{}
    if (true && true) {
        fmt.Println(v5, v1, v4)
        v9 = (v9 ^ 1)
    }
    fmt.Println(v4, v4, v5)
    for fi12 := 0; fi12 < 16; fi12++ {
        for fi13 := 0; fi13 < 16; fi13++ {
            v2 = v2
            _ = fi13
        }
        _ = fi12
    }
    switch 390 {
    case -9223372036854775808:
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
    default:
        v14 := GenS1{F0: 128, F1: 18446744073709551615}
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
