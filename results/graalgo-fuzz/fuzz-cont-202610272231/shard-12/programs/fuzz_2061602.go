package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 int }
type GenS2 struct { F0 int; F1 int8; F2 int8 }
func genFn0(p0 int8) int {
    var v6 int32 = 1
    var v7 uint32 = 0
    _ = p0
    _ = v6
    _ = v7
    return -2147483648
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := []int{2147483647, 9223372036854775807, 724}
    v3 := []uint32{42, 65535, 1, 65535}
    var v4 uint = 370
    var v5 int64 = -9223372036854775808
    v5 = 879
    if true {
        for fi8 := 0; fi8 < 14; fi8++ {
            v1 += v5
            v1 -= v5
            _ = fi8
        }
    }
    v5 -= v5
    fmt.Println(v5)
    for fi9 := 0; fi9 < 3; fi9++ {
        var v10 float64 = 0.1
        for fi11 := 0; fi11 < 15; fi11++ {
            if (2 >= 128) {
                switch fi9 {
                case 128:
                    v1 -= v5
                }
                v5 = -1
                if true {
                    v1 += 615
                    v12 := GenS1{F0: 423, F1: (v4 ^ 1), F2: fi11}
                    _ = v12
                }
            }
            v5 = v1
            fmt.Println(v10)
            _ = fi11
        }
        _ = fi9
        _ = v10
    }
    v1 = -444
    for fi13 := 0; fi13 < 3; fi13++ {
        v14 := GenS1{F0: 1, F1: 241, F2: 685}
        fmt.Println(v5, v5, fi13)
        _ = fi13
        _ = v14
    }
    v1 -= (v5 / (((^(v5))) | 1))
    fmt.Println(v5, v1, v1)
    var v15 int16 = -1
    var v16 uint32 = 2147483648
    for fi17 := 0; fi17 < 14; fi17++ {
        var v18 uint = v4
        _ = fi17
        _ = v18
    }
    v19 := GenS2{F0: 128, F1: -81, F2: -128}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
