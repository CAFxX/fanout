package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
func genFn0() uint8 {
    var v6 int16 = 553
    if false {
        v7 := []int{368}
        _ = v7
    }
    _ = v6
    return 42
}
func genFn1(p0 uint64) int32 {
    var v8 int32 = 810
    var v9 bool = (!false)
    if true {
        var v10 float32 = (1.5 / 1.5)
        _ = v10
    } else {
        switch ("\n\t" + ("\n\t" + "hello")) {
        case "foo bar":
            var v11 uint16 = 753
            _ = v11
        }
        for fi12 := 0; fi12 < 11; fi12++ {
            var v13 uint32 = 261
            v13 = (v13 * 65535)
            _ = fi12
            _ = v13
        }
        var v14 string = ("\n\t" + "foo bar")
        _ = v14
    }
    _ = p0
    _ = v8
    _ = v9
    return (v8 % ((-1) | 1))
}

func main() {
    var v1 uint8 = 42
    v2 := []int{-9223372036854775808, 32768, 671, 102}
    v3 := make(map[uint32]uint32)
    var v4 float64 = 1e10
    var v5 int8 = -1
    v5 -= v5
    v15 := GenS1{F0: -128, F1: v1}
    for fi16 := 0; fi16 < 20; fi16++ {
        var v17 int = 99
        fmt.Println(v5, v17, fi16)
        v1 -= v1
        _ = fi16
        _ = v17
    }
    v18 := GenS1{F0: -1, F1: v1}
    fmt.Println(v5, v5, v5)
    switch "foo bar" {
    case "\n\t":
        delete(v3, 1)
    case "a":
        fmt.Println(v4, v4)
    case "foo bar":
        switch "" {
        case "\n\t":
            _ = len(v2)
        default:
            for fi19 := 0; fi19 < 3; fi19++ {
                var v20 uint16 = 255
                v5 = (-(v5))
                v20 += (v20 | v20)
                _ = fi19
                _ = v20
            }
        }
    }
    for fi21 := 0; fi21 < 14; fi21++ {
        for fi22 := 0; fi22 < 9; fi22++ {
            fmt.Println(v1, fi21, v4)
            _ = fi22
        }
        for fi23 := 0; fi23 < 20; fi23++ {
            var v24 int = fi21
            _ = fi23
            _ = v24
        }
        _ = fi21
    }
    v25 := GenS1{F0: 52, F1: v1}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
