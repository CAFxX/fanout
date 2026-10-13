package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint; F1 uint; F2 int32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint; F1 uint32; F2 uint16 }
func genFn0() uint {
    var v4 uint32 = 42
    v4 -= 2147483647
    _ = v4
    return 0
}
func genFn1(p0 int16, p1 int) uint64 {
    if false {
        for fi5 := 0; fi5 < 14; fi5++ {
            if true {
                if true {
                    v6 := GenS1{F0: -1, F1: -192, F2: 0}
                    v7 := GenS2{F0: -104}
                    var v8 int64 = 146
                    _ = v6
                    _ = v7
                    _ = v8
                }
                p1 += fi5
            } else {
                if false {
                    p1 = p1
                } else {
                    v9 := GenS3{F0: 65535, F1: 2, F2: 0}
                    _ = v9
                }
                p1 -= 0
            }
            _ = fi5
        }
        var v10 int32 = 1
        _ = v10
    }
    var v11 bool = true
    v12 := []int{128, (p1 << ((4294967295) & 63))}
    var v13 uint = 340
    v14 := GenS1{F0: 42, F1: p1, F2: -1}
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    return 104
}

func main() {
    var v1 uint8 = 128
    v2 := make(map[uint32]int)
    v3 := make(map[int]string)
    v1 = (v1 * 186)
    for fi15 := 0; fi15 < 4; fi15++ {
        v1 = (^(v1))
        _ = fi15
    }
    v16 := []int{}
    switch 0 {
    case 367:
        v1 += (v1 - (v1 * v1))
    }
    var v17 float32 = 0.1
    switch "a" {
    case "\n\t":
        var v18 uint = 18446744073709551615
        _ = v18
    case "hello":
        fmt.Println(v1, v17)
    case "a":
        fmt.Println(v17)
    }
    switch (((("a" + "foo bar") + "a") + (("\n\t" + "\n\t") + "")) + "") {
    case "hello":
        if len(v16) > 0 {
            v16[0] = 128
        }
    case "a":
        v19 := GenS4{F0: 9223372036854775807}
        _ = v19
    case "":
        switch "\n\t" {
        case "":
            v20 := GenS3{F0: 1, F1: genFn0(), F2: -943}
            _ = v20
        default:
            var v21 int64 = 0
            _ = v21
        }
    default:
        v22 := GenS5{F0: 42, F1: 65535, F2: uint16(v1)}
        _ = v22
    }
    var v23 float32 = (v17 * (v17 - v17))
    var v24 uint32 = 2147483647
    if true {
        switch "foo bar" {
        case "a":
            v1 = (v1 << ((280) & 63))
        case "hello":
            var v25 string = ("\n\t" + "foo bar")
            _ = v25
        case "\n\t":
            fmt.Println(v1, v23)
        }
    }
    if (1 > 0) {
        v24 += v24
    } else {
        if false {
            v2[721] = -2147483648
            _ = len(v16)
        } else {
            v17 = 3.14159
            v26 := GenS3{F0: genFn0(), F1: genFn0(), F2: -1}
            _ = v26
        }
    }
    if ((42 == 363) && false) {
        var v27 float64 = 1e10
        var v28 int16 = 1
        var v29 uint32 = v24
        _ = v27
        _ = v28
        _ = v29
    }
    fmt.Println(v24)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v24
    _ = v3
}
