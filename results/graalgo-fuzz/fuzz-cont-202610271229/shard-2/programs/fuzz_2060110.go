package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 uint16; F1 uint64; F2 uint64 }

func main() {
    var v1 int = -1
    var v2 uint8 = 180
    v3 := make(map[int]int)
    var v4 uint64 = 18446744073709551615
    var v5 uint = 57
    v6 := []uint32{65535, 933, 4294967295, 0}
    v7 := GenS1{F0: uint(v4), F1: -99}
    if len(v6) > 0 {
        v6[0] = 2
    }
    v8 := GenS1{F0: 691, F1: -2147483648}
    var v9 bool = true
    for fi10 := 0; fi10 < 9; fi10++ {
        v11 := GenS1{F0: 971, F1: -2147483648}
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 15; fi12++ {
        var v13 bool = (0 >= -110)
        switch "\n\t" {
        case "foo bar":
            var v14 int16 = 1
            _ = v14
        case "":
            v15 := GenS2{F0: 42, F1: 0, F2: v4}
            _ = v15
        }
        _ = fi12
        _ = v13
    }
    v1 += -231
    if v9 {
        fmt.Println(v9)
        _ = len(v6)
    }
    for fi16 := 0; fi16 < 8; fi16++ {
        for fi17 := 0; fi17 < 1; fi17++ {
            v4 = v4
            v18 := v3[fi17]
            _ = fi17
            _ = v18
        }
        v4 -= (v4 - (v4 + 42))
        v6 = append(v6, 221)
        _ = fi16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
