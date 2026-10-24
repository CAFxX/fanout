package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 int32; F1 uint64; F2 uint32 }
func genFn0() uint32 {
    for fi5 := 0; fi5 < 11; fi5++ {
        for fi6 := 0; fi6 < 11; fi6++ {
            v7 := GenS1{F0: 2, F1: 316}
            var v8 uint32 = 4294967295
            _ = fi6
            _ = v7
            _ = v8
        }
        for fi9 := 0; fi9 < 12; fi9++ {
            v10 := GenS1{F0: 4294967295, F1: 541}
            _ = fi9
            _ = v10
        }
        var v11 uint16 = 863
        _ = fi5
        _ = v11
    }
    var v12 float64 = (3.14159 - 0.1)
    v13 := []int{}
    v12 -= v12
    _ = v12
    _ = v13
    return 501
}

func main() {
    v1 := []uint32{814}
    v2 := []uint32{1, 4294967295, 2147483647, 42}
    var v3 uint = 42
    var v4 uint16 = 65535
    if len(v2) > 0 {
        v2[0] = 1
    }
    v14 := GenS2{F0: -492, F1: 18446744073709551615, F2: 1}
    var v15 float32 = 3.14159
    var v16 float32 = v15
    fmt.Println(v3)
    var v17 int64 = 42
    v1 = append(v1, 65535)
    v17 += -467
    for fi18 := 0; fi18 < 5; fi18++ {
        fmt.Println(v15)
        _ = fi18
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
