package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 int8 }
type GenS2 struct { F0 int32; F1 uint64; F2 int8 }
type GenS3 struct { F0 uint; F1 int; F2 int }
type GenS4 struct { F0 int8; F1 bool }
type GenS5 struct { F0 int64; F1 uint32 }
func genFn0() string {
    for fi5 := 0; fi5 < 20; fi5++ {
        v6 := GenS1{F0: 2, F1: 42, F2: 127}
        var v7 uint64 = 4294967295
        for fi8 := 0; fi8 < 14; fi8++ {
            var v9 uint64 = v7
            v7 = (v7 ^ v9)
            v10 := []int{fi5, 127}
            var v11 float32 = 1e10
            _ = fi8
            _ = v10
            _ = v11
            _ = v9
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    v12 := []int{993}
    var v13 int32 = 736
    fmt.Println(v13)
    _ = v12
    _ = v13
    return ""
}

func main() {
    var v1 uint32 = 806
    var v2 uint = 211
    v3 := []string{}
    var v4 int64 = 1
    switch -974 {
    case -954:
        v14 := GenS1{F0: 18446744073709551615, F1: -940, F2: 127}
        _ = v14
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v4 = 9223372036854775807
    v15 := GenS2{F0: 306, F1: 18446744073709551615, F2: -128}
    fmt.Println(v4, v2, v1)
    var v16 int8 = 42
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v16 = (v16 * (v16 & int8(v2)))
    switch 1 {
    case -572:
        v2 = (v2 ^ (v2 % (((v2 ^ (v2 + v2))) | 1)))
    case 2147483647:
        v17 := GenS3{F0: 933, F1: 127, F2: 396}
        _ = v17
    default:
        v18 := GenS4{F0: (v16 << (((v2 % ((200) | 1))) & 63)), F1: (701 > 42)}
        _ = v18
    }
    v2 = v2
    for fi19 := 0; fi19 < 12; fi19++ {
        if true {
            v20 := GenS5{F0: v4, F1: (v1 % ((v1) | 1))}
            var v21 int64 = (v4 & (v4 << ((v2) & 63)))
            _ = v20
            _ = v21
        } else {
            var v22 int = (fi19 + (fi19 / ((fi19) | 1)))
            var v23 int8 = (v16 | (^(v16)))
            _ = v22
            _ = v23
        }
        _ = fi19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
