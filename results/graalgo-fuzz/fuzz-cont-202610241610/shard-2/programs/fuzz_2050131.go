package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64; F1 uint }
func genFn0() uint16 {
    var v5 uint8 = 61
    if (760 <= 606) {
        var v6 float32 = ((1.5 / 1.5) - (0.0 - -2.25))
        for fi7 := 0; fi7 < 11; fi7++ {
            var v8 int64 = -1
            v9 := GenS1{F0: uint(v5)}
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = v6
    }
    var v10 float32 = (0.1 / 0.1)
    fmt.Println(v5, v5)
    _ = v10
    _ = v5
    return 0
}
func genFn1() int32 {
    var v11 int = 0
    switch v11 {
    case 4294967295:
        var v12 int16 = -32768
        _ = v12
    case 32768:
        for fi13 := 0; fi13 < 1; fi13++ {
            fmt.Println(v11, v11)
            var v14 uint = 946
            _ = fi13
            _ = v14
        }
    default:
        v11 = (-(v11))
    }
    var v15 int8 = -1
    _ = v11
    _ = v15
    return -1
}

func main() {
    var v1 float64 = 0.1
    var v2 uint32 = 42
    v3 := []string{"", "", "hello", "foo bar"}
    var v4 int = 1
    var v16 uint8 = 128
    var v17 float64 = v1
    fmt.Println(v17, v16)
    for fi18 := 0; fi18 < 15; fi18++ {
        v19 := GenS2{F0: 2, F1: 548}
        v17 = (v17 / v17)
        _ = fi18
        _ = v19
    }
    for fi20 := 0; fi20 < 1; fi20++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi20
    }
    v3 = append(v3, "")
    var v21 float32 = ((1.5 - 1.5) / (-2.25 + 1.5))
    var v22 int8 = 0
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v1)
    if true {
        var v23 int32 = -862
        var v24 uint64 = uint64(v16)
        v25 := GenS2{F0: (v24 / ((v24) | 1)), F1: 4294967295}
        _ = v23
        _ = v24
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
