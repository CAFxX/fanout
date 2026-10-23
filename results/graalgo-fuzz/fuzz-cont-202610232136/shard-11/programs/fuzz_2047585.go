package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool; F1 uint32; F2 int64 }
func genFn0(p0 bool, p1 int8) uint32 {
    var v7 int64 = -9223372036854775808
    fmt.Println(v7, p0)
    var v8 uint32 = 2
    v9 := GenS1{F0: false}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 0
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int32 = 0
    v3 := []string{"", "a", ""}
    var v4 uint32 = 2147483648
    var v5 uint = 782
    var v6 float64 = -2.25
    v1[9223372036854775807] = (v4 + 2)
    for fi10 := 0; fi10 < 7; fi10++ {
        var v11 uint16 = 0
        var v12 float32 = 1e10
        _ = fi10
        _ = v11
        _ = v12
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v13 int32 = v2
    v14 := GenS1{F0: false}
    if false {
        var v15 int32 = v2
        for fi16 := 0; fi16 < 17; fi16++ {
            v17 := GenS2{F0: false, F1: v4, F2: 9223372036854775807}
            _ = fi16
            _ = v17
        }
        _ = v15
    }
    v13 += (v13 << (((v5 | v5)) & 63))
    var v18 int = 32767
    v2 -= (^(v2))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
