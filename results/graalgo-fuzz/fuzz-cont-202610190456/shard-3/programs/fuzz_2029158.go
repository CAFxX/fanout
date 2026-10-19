package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 int32; F2 uint8 }
type GenS4 struct { F0 string; F1 bool; F2 uint16 }
func genFn0(p0 int64, p1 bool) uint64 {
    for fi6 := 0; fi6 < 13; fi6++ {
        fmt.Println(fi6, p0, p1)
        fmt.Println(p0, fi6, p1)
        var v7 int8 = 1
        _ = fi6
        _ = v7
    }
    fmt.Println(p0)
    var v8 uint32 = 1
    v9 := []int{32768}
    _ = len(v9)
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return uint64(v8)
}

func main() {
    v1 := []int{1, -500}
    var v2 bool = (1e10 > 1e10)
    var v3 uint16 = 1
    var v4 uint16 = (v3 * v3)
    var v5 uint8 = 127
    v10 := GenS1{F0: v2}
    for fi11 := 0; fi11 < 8; fi11++ {
        v12 := GenS1{F0: ((1.5 >= 0.1) && true)}
        _ = fi11
        _ = v12
    }
    switch 2147483647 {
    case 273:
        v3 -= v3
    case -502:
        var v13 int32 = -2147483648
        _ = v13
    case 2147483647:
        v14 := GenS2{F0: 32767}
        _ = v14
    }
    var v15 int = 4294967295
    v16 := GenS3{F0: v15, F1: -2147483648, F2: 1}
    for fi17 := 0; fi17 < 14; fi17++ {
        var v18 string = "\n\t"
        _ = fi17
        _ = v18
    }
    v3 = (^(v4))
    for fi19 := 0; fi19 < 13; fi19++ {
        var v20 string = ("\n\t" + "foo bar")
        for fi21 := 0; fi21 < 18; fi21++ {
            v22 := GenS4{F0: "\n\t", F1: v2, F2: (v3 % ((v4) | 1))}
            _ = fi21
            _ = v22
        }
        v5 = (v5 << ((uint(v3)) & 63))
        _ = fi19
        _ = v20
    }
    v3 = 42
    v1 = append(v1, v15)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v5 += v5
    v15 -= (-(v15))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
