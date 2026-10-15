package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 uint16; F1 uint }
func genFn0(p0 uint) uint32 {
    var v7 string = ("a" + "\n\t")
    v8 := GenS1{F0: 2, F1: 662}
    var v9 uint8 = 1
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 459
}

func main() {
    var v1 int16 = -32768
    v2 := make(map[string]uint32)
    var v3 float64 = float64(1.5)
    var v4 bool = false
    v5 := make(map[int]uint32)
    v6 := []int{0}
    if v4 {
        v10 := GenS1{F0: 2, F1: 355}
        var v11 int = 129
        v12 := GenS1{F0: 2147483648, F1: 1}
        _ = v10
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 1; fi13++ {
        v1 -= v1
        v14 := v5[(fi13 << ((42) & 63))]
        _ = fi13
        _ = v14
    }
    fmt.Println(v3, v3, v4)
    var v15 string = (("" + "foo bar") + "a")
    switch v15 {
    case "foo bar":
        switch 0 {
        case 42:
            var v16 uint16 = 42
            _ = v16
        }
    case "hello":
        v17 := GenS2{F0: 32767, F1: 65535}
        _ = v17
    default:
        for fi18 := 0; fi18 < 11; fi18++ {
            fmt.Println(v15, v4, v15)
            fmt.Println(v4, v15)
            _ = fi18
        }
    }
    var v19 int = 930
    v1 += v1
    var v20 uint64 = 2
    if (!(v20 > (^(v20)))) {
        v1 = (^(v1))
    } else {
        var v21 uint64 = (v20 + v20)
        var v22 int64 = 0
        _ = v21
        _ = v22
    }
    for fi23 := 0; fi23 < 2; fi23++ {
        var v24 int8 = -128
        v25 := GenS2{F0: 155, F1: 1}
        _ = fi23
        _ = v24
        _ = v25
    }
    if len(v6) > 0 {
        v6[0] = (v19 + v19)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
