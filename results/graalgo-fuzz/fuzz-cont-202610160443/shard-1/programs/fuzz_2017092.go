package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 int, p1 int16) int8 {
    var v7 uint = 1
    v8 := GenS1{F0: 42}
    for fi9 := 0; fi9 < 15; fi9++ {
        var v10 int8 = -96
        var v11 float64 = ((0.1 - 0.0) * 1e10)
        v12 := []int{}
        _ = fi9
        _ = v10
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 0
}
func genFn1() int64 {
    var v13 int = 377
    var v14 int32 = -658
    _ = v13
    _ = v14
    return -1
}

func main() {
    v1 := []uint32{463, 238}
    var v2 float64 = -2.25
    var v3 uint = 18446744073709551615
    var v4 int = 2147483647
    v5 := []int{v4, (v4 + 32768)}
    var v6 uint32 = 2147483647
    v15 := GenS1{F0: 2}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v16 := GenS1{F0: 1}
    switch "hello" {
    case "foo bar":
        v2 += (1e10 / v2)
    default:
        if true {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            for fi17 := 0; fi17 < 19; fi17++ {
                _ = len(v1)
                _ = fi17
            }
            v18 := GenS1{F0: 224}
            _ = v18
        }
    }
    var v19 uint16 = 65535
    v4 = (v4 ^ v4)
    _ = len(v5)
    v3 = (v3 % ((65535) | 1))
    for fi20 := 0; fi20 < 18; fi20++ {
        v21 := GenS1{F0: uint8(v3)}
        _ = fi20
        _ = v21
    }
    v19 = v19
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
