package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 int32; F1 uint; F2 int }
func genFn0(p0 int, p1 uint) int16 {
    fmt.Println(p0, p0)
    p1 = (^(p1))
    p1 = (^(p1))
    for fi7 := 0; fi7 < 10; fi7++ {
        v8 := []int{}
        var v9 int8 = 127
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    return 0
}
func genFn1(p0 uint64, p1 bool) uint {
    var v10 uint8 = 1
    var v11 uint16 = uint16(p0)
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return uint(v11)
}

func main() {
    v1 := make(map[uint32]string)
    var v2 float64 = float64(float32((3.14159 * 0.0)))
    v3 := make(map[uint32]string)
    v4 := []uint32{}
    var v5 bool = true
    var v6 bool = (v2 > (v2 * float64(3.14159)))
    for fi12 := 0; fi12 < 18; fi12++ {
        for fi13 := 0; fi13 < 10; fi13++ {
            v6 = true
            _ = fi13
        }
        fmt.Println(v5, fi12, v2)
        var v14 uint16 = 65535
        _ = fi12
        _ = v14
    }
    v2 = (0.1 * 0.1)
    var v15 string = "a"
    v16 := GenS1{F0: 0, F1: 1000}
    var v17 uint8 = 128
    switch -1 {
    case 0:
        switch 4294967295 {
        case -374:
            fmt.Println(v5, v2)
        case -9223372036854775808:
            var v18 int16 = -32768
            _ = v18
        }
    case -9223372036854775808:
        v19 := GenS2{F0: -841, F1: 4294967295, F2: -1}
        _ = v19
    case 42:
        _ = len(v3)
    default:
        var v20 bool = v6
        _ = v20
    }
    v15 = "hello"
    _ = len(v3)
    for fi21 := 0; fi21 < 13; fi21++ {
        switch (fi21 * (fi21 + fi21)) {
        case 128:
            _ = len(v4)
        case -509:
            v2 += v2
        }
        v15 = v15
        var v22 uint32 = 42
        _ = fi21
        _ = v22
    }
    if true {
        var v23 int = -1
        var v24 float32 = 1.5
        _ = v23
        _ = v24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
