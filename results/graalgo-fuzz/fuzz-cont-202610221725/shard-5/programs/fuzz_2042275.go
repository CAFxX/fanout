package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 uint64; F1 uint64 }
func genFn0(p0 uint8, p1 int64, p2 int) uint {
    fmt.Println(p2, p2, p1)
    p1 += p1
    fmt.Println(p0)
    v4 := []int{(-(p2))}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 4294967295
}

func main() {
    var v1 int32 = 2147483647
    v2 := []uint32{65535}
    var v3 uint16 = 42
    v5 := GenS1{F0: 311, F1: (v3 + 122)}
    var v6 int = 258
    if (v3 <= v3) {
        for fi7 := 0; fi7 < 6; fi7++ {
            switch v6 {
            case 2147483647:
                v3 = (v3 ^ 65535)
            case 1:
                v3 = v3
            }
            var v8 int8 = 101
            _ = fi7
            _ = v8
        }
    }
    switch 4294967295 {
    case 0:
        fmt.Println(v3, v3)
    case -2147483648:
        switch v6 {
        case 32767:
            fmt.Println(v6)
        }
    default:
        var v9 string = "\n\t"
        _ = v9
    }
    fmt.Println(v6, v1, v6)
    if false {
        var v10 uint32 = 65535
        fmt.Println(v10)
        var v11 uint32 = 4294967295
        _ = v10
        _ = v11
    } else {
        switch 127 {
        case 42:
            v3 += 0
        case 32767:
            var v12 uint = 1
            _ = v12
        }
        var v13 uint8 = 2
        var v14 uint8 = (v13 + (v13 + (v13 + v13)))
        _ = v13
        _ = v14
    }
    if (1 > v3) {
        v3 -= (v3 % ((v3) | 1))
    }
    var v15 int64 = -25
    if true {
        var v16 float32 = ((0.1 + 3.14159) / 1.5)
        if true {
            var v17 uint16 = 32767
            _ = v17
        }
        v18 := GenS2{F0: 4294967295, F1: 171}
        _ = v16
        _ = v18
    } else {
        v3 = (v3 ^ v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v5
    _ = v6
}
