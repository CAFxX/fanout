package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 int64 }
func genFn0() bool {
    switch "hello" {
    case "foo bar":
        var v7 uint = 459
        _ = v7
    default:
        var v8 uint32 = 42
        _ = v8
    }
    for fi9 := 0; fi9 < 14; fi9++ {
        fmt.Println(fi9)
        v10 := []int{}
        fmt.Println(fi9)
        _ = fi9
        _ = v10
    }
    var v11 int = 4294967295
    _ = v11
    return false
}
func genFn1(p0 int) uint64 {
    var v12 float64 = ((0.0 / 0.1) * -2.25)
    var v13 int8 = -1
    fmt.Println(v12, v12, v12)
    var v14 int32 = -1
    _ = p0
    _ = v12
    _ = v13
    _ = v14
    return 335
}

func main() {
    var v1 uint = 992
    var v2 uint64 = 42
    v3 := make(map[string]uint32)
    v4 := []int{32768}
    v5 := []uint32{0, 459}
    var v6 string = "hello"
    var v15 uint = v1
    var v16 float64 = (0.0 / 3.14159)
    switch 127 {
    case 127:
        var v17 int16 = 32767
        _ = v17
    default:
        for fi18 := 0; fi18 < 13; fi18++ {
            v1 = (v15 ^ v1)
            fmt.Println(fi18, v15, v1)
            v2 = 183
            _ = fi18
        }
    }
    var v19 int64 = 49
    var v20 uint16 = 965
    for fi21 := 0; fi21 < 14; fi21++ {
        if (-2147483648 > -296) {
            v22 := GenS1{F0: "\n\t", F1: 4294967295}
            var v23 bool = false
            v6 = v6
            _ = v22
            _ = v23
        }
        _ = fi21
    }
    v24 := GenS2{F0: v19}
    var v25 bool = false
    var v26 uint8 = 151
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
