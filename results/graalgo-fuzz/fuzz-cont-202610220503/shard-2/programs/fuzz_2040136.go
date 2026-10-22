package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int32; F1 uint32; F2 bool }
func genFn0(p0 uint16, p1 bool, p2 int8) bool {
    v4 := GenS1{F0: 215, F1: 2, F2: 32767}
    v5 := GenS1{F0: 42, F1: 210, F2: 255}
    var v6 uint8 = 1
    p1 = p1
    if p1 {
        var v7 int8 = p2
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return (uint64(p0) >= 2)
}
func genFn1(p0 int64, p1 uint) int64 {
    p0 -= (p0 + p0)
    p1 += 1
    var v8 int32 = -1
    _ = p0
    _ = p1
    _ = v8
    return (p0 + p0)
}

func main() {
    v1 := make(map[int]uint32)
    var v2 bool = false
    v3 := []int{9223372036854775807, 650}
    v9 := GenS1{F0: 2, F1: 144, F2: 378}
    fmt.Println(v2)
    v2 = (128 <= 0)
    v10 := GenS2{F0: 522}
    v2 = v2
    delete(v1, -2147483648)
    for fi11 := 0; fi11 < 17; fi11++ {
        if v2 {
            v12 := GenS3{F0: "foo bar"}
            _ = v12
        }
        var v13 bool = (v2 || (1e10 >= 0.1))
        _ = len(v1)
        _ = fi11
        _ = v13
    }
    for fi14 := 0; fi14 < 9; fi14++ {
        var v15 int64 = -1
        _ = fi14
        _ = v15
    }
    switch "" {
    case "hello":
        for fi16 := 0; fi16 < 1; fi16++ {
            v2 = (963 != 0)
            fmt.Println(v2, fi16)
            _ = fi16
        }
    case "foo bar":
        v2 = (-1 != -979)
    default:
        switch "foo bar" {
        case "a":
            v2 = ((0.0 < 0.0) || (!(322 <= 81)))
        }
    }
    if true {
        for fi17 := 0; fi17 < 9; fi17++ {
            v2 = v2
            _ = fi17
        }
    }
    for fi18 := 0; fi18 < 7; fi18++ {
        v2 = true
        var v19 float64 = 0.0
        _ = fi18
        _ = v19
    }
    fmt.Println(v2)
    v20 := GenS3{F0: ("" + "\n\t")}
    for fi21 := 0; fi21 < 16; fi21++ {
        fmt.Println(fi21, fi21)
        _ = fi21
    }
    if (453 == 953) {
        var v22 string = ("\n\t" + "\n\t")
        _ = v22
    } else {
        var v23 uint32 = 65535
        _ = v23
    }
    v24 := GenS4{F0: 42, F1: 361, F2: ((4294967295 < 128) && (v2 && v2))}
    v2 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v9
}
