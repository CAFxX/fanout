package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 int16 }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 int64; F1 uint8; F2 uint64 }
type GenS4 struct { F0 int32; F1 int64; F2 uint64 }
func genFn0(p0 bool) int64 {
    v6 := []int{128, 9223372036854775807, -1}
    _ = len(v6)
    if p0 {
        fmt.Println(p0)
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = len(v6)
    }
    if len(v6) > 0 {
        v6[0] = 752
    }
    _ = p0
    _ = v6
    return -1
}

func main() {
    v1 := []int{306}
    v2 := make(map[string]bool)
    var v3 string = "\n\t"
    var v4 float64 = 3.14159
    v5 := []uint32{}
    var v7 int16 = -289
    v8 := GenS1{F0: 0, F1: -2147483648, F2: (-(v7))}
    v9 := GenS2{F0: 55, F1: 126}
    for fi10 := 0; fi10 < 3; fi10++ {
        _ = len(v1)
        v7 = v7
        var v11 uint16 = 255
        _ = fi10
        _ = v11
    }
    if true {
        for fi12 := 0; fi12 < 19; fi12++ {
            fmt.Println(v4)
            _ = fi12
        }
    } else {
        v13 := GenS1{F0: -1, F1: -1, F2: (v7 & v7)}
        _ = v13
    }
    switch 128 {
    case 2147483647:
        switch 156 {
        case -1:
            var v14 uint64 = 2
            _ = v14
        case 128:
            switch -629 {
            case -608:
                fmt.Println(v3)
            case 32768:
                for fi15 := 0; fi15 < 10; fi15++ {
                    v16 := GenS3{F0: 1, F1: 2, F2: 352}
                    _ = fi15
                    _ = v16
                }
            case 269:
                v17 := GenS3{F0: 42, F1: 90, F2: 524}
                _ = v17
            }
        case 32767:
            var v18 float64 = v4
            _ = v18
        default:
            for fi19 := 0; fi19 < 14; fi19++ {
                for fi20 := 0; fi20 < 20; fi20++ {
                    v4 = v4
                    if false {
                        var v21 int8 = -8
                        fmt.Println(v4)
                        _ = v21
                    }
                    _ = fi20
                }
                v22 := GenS1{F0: 42, F1: 60, F2: (v7 - v7)}
                _ = fi19
                _ = v22
            }
        }
    }
    var v23 float32 = 0.0
    fmt.Println(v23)
    var v24 uint8 = 0
    v25 := GenS3{F0: 0, F1: v24, F2: 718}
    v26 := GenS4{F0: 2147483647, F1: -9223372036854775808, F2: 1}
    fmt.Println(v23, v23, v4)
    if false {
        var v27 bool = true
        v3 = v3
        _ = v27
    }
    v28 := GenS1{F0: 1, F1: -1, F2: v7}
    for fi29 := 0; fi29 < 19; fi29++ {
        v24 += (^(v24))
        v30 := v2[v3]
        _ = fi29
        _ = v30
    }
    fmt.Println(v3)
    if (v4 != 0.0) {
        var v31 int = 1
        fmt.Println(v4)
        _ = v31
    } else {
        v7 += -32768
    }
    v3 = (v3 + v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
