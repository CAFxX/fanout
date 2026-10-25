package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 uint16; F2 uint32 }
type GenS3 struct { F0 int64; F1 uint8 }
func genFn0() int32 {
    var v4 int64 = 397
    switch 0 {
    case 9223372036854775807:
        v5 := []int{}
        _ = v5
    case 4294967295:
        v6 := GenS1{F0: 1}
        _ = v6
    default:
        v7 := GenS2{F0: 669, F1: 32767, F2: 2147483648}
        _ = v7
    }
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 uint8 = 148
        _ = fi8
        _ = v9
    }
    var v10 float64 = float64(((0.1 * 0.0) / 1.5))
    v11 := []int{-1}
    _ = v10
    _ = v11
    _ = v4
    return -1
}

func main() {
    var v1 int = 32767
    var v2 uint = 18446744073709551615
    v3 := make(map[string]uint32)
    v12 := GenS3{F0: 42, F1: 1}
    for fi13 := 0; fi13 < 1; fi13++ {
        var v14 int8 = -128
        var v15 int64 = 185
        _ = fi13
        _ = v14
        _ = v15
    }
    for fi16 := 0; fi16 < 19; fi16++ {
        var v17 bool = true
        fmt.Println(v2, v2)
        _ = fi16
        _ = v17
    }
    var v18 uint16 = 423
    var v19 int8 = -1
    v3["\n\t"] = 1
    for fi20 := 0; fi20 < 11; fi20++ {
        delete(v3, (("foo bar" + "") + "a"))
        _ = fi20
    }
    delete(v3, "")
    v3["hello"] = 0
    var v21 uint64 = 256
    var v22 int = v1
    fmt.Println(v2)
    var v23 float64 = -2.25
    v24 := []int{-2147483648, (v22 >> ((v2) & 63)), v22}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}
