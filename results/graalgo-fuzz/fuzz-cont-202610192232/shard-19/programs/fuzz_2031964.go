package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }

func main() {
    v1 := []int{9223372036854775807}
    var v2 int = 668
    var v3 uint32 = 2
    switch 127 {
    case 42:
        if true {
            v3 += (v3 & 2)
        }
    case 32767:
        v2 += (v2 + v2)
    default:
        for fi4 := 0; fi4 < 3; fi4++ {
            v3 -= (^(v3))
            var v5 bool = false
            var v6 int8 = 127
            _ = fi4
            _ = v5
            _ = v6
        }
    }
    if true {
        if false {
            v3 = v3
        }
        v7 := GenS1{F0: v3, F1: v2}
        v2 = v2
        _ = v7
    }
    for fi8 := 0; fi8 < 10; fi8++ {
        if false {
            var v9 int64 = -1
            v10 := GenS1{F0: (v3 / ((65535) | 1)), F1: fi8}
            _ = v10
            _ = v9
        }
        _ = fi8
    }
    for fi11 := 0; fi11 < 16; fi11++ {
        var v12 int = 127
        v3 = v3
        _ = fi11
        _ = v12
    }
    v2 = v2
    switch (-(v2)) {
    case 2147483647:
        v2 -= v2
    default:
        v3 = (^(v3))
    }
    var v13 int = -860
    if true {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v14 uint = 18446744073709551615
        _ = v14
    }
    v13 = (v13 / (((v2 ^ 4294967295)) | 1))
    if len(v1) > 0 {
        v1[0] = v2
    }
    v15 := GenS1{F0: (^(v3)), F1: (v2 ^ int(v3))}
    fmt.Println(v3, v13, v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v15
    _ = v2
    _ = v3
}
