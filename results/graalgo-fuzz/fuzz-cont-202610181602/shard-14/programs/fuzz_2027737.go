package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 int64 }
type GenS2 struct { F0 uint16; F1 int16; F2 uint }
type GenS3 struct { F0 uint; F1 int16; F2 int8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int8; F1 bool; F2 uint64 }
func genFn0(p0 int16, p1 uint8) uint32 {
    if true {
        v4 := GenS1{F0: 2, F1: (p0 + (p0 | p0)), F2: 1}
        var v5 int = 2147483647
        _ = v4
        _ = v5
    }
    if false {
        v6 := []int{}
        v7 := GenS1{F0: 449, F1: -661, F2: -9223372036854775808}
        var v8 uint64 = 857
        _ = v6
        _ = v7
        _ = v8
    }
    v9 := GenS2{F0: 0, F1: p0, F2: 826}
    _ = p0
    _ = p1
    _ = v9
    return 0
}

func main() {
    v1 := make(map[string]string)
    var v2 float32 = 1.5
    var v3 int8 = -89
    v10 := []int{}
    v11 := GenS3{F0: 2, F1: -858, F2: (v3 >> ((2) & 63))}
    v2 = v2
    var v12 uint8 = 2
    fmt.Println(v3, v2, v2)
    var v13 int8 = (v3 % ((-124) | 1))
    v14 := GenS1{F0: 42, F1: 682, F2: 315}
    var v15 uint64 = 762
    for fi16 := 0; fi16 < 14; fi16++ {
        v17 := GenS4{F0: 0}
        _ = fi16
        _ = v17
    }
    var v18 int8 = int8(v15)
    fmt.Println(v13, v18, v15)
    v19 := GenS5{F0: -46, F1: true, F2: 0}
    fmt.Println(v15, v13, v13)
    _ = len(v1)
    for fi20 := 0; fi20 < 12; fi20++ {
        v10 = append(v10, (fi20 << ((2) & 63)))
        _ = fi20
    }
    _ = len(v10)
    v18 = -128
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
