package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 uint8 }
func genFn0(p0 int16) int32 {
    for fi6 := 0; fi6 < 16; fi6++ {
        fmt.Println(p0, fi6)
        for fi7 := 0; fi7 < 14; fi7++ {
            p0 = p0
            _ = fi7
        }
        _ = fi6
    }
    switch 128 {
    case 9223372036854775807:
        var v8 uint = 65535
        _ = v8
    default:
        var v9 int64 = -1
        _ = v9
    }
    for fi10 := 0; fi10 < 10; fi10++ {
        p0 += p0
        _ = fi10
    }
    v11 := GenS1{F0: 152, F1: 1, F2: 128}
    _ = p0
    _ = v11
    return 2147483647
}
func genFn1(p0 uint64, p1 uint16, p2 bool) int16 {
    var v12 int32 = -923
    var v13 int8 = int8(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return -533
}

func main() {
    var v1 int = 1
    var v2 uint64 = 18446744073709551615
    v3 := []int{(v1 ^ (-(v1)))}
    v4 := make(map[uint32]string)
    var v5 uint = 955
    var v14 int8 = 42
    fmt.Println(v5)
    var v15 uint64 = 18446744073709551615
    var v16 int = v1
    var v17 bool = false
    v14 += (v14 & v14)
    fmt.Println(v14)
    v15 = 4294967295
    _ = len(v4)
    v1 += (v16 + v16)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
