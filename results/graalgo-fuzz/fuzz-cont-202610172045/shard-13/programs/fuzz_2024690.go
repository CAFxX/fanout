package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 bool }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int64; F1 int8 }
type GenS5 struct { F0 uint8; F1 int8 }
func genFn0(p0 uint) int32 {
    var v5 int32 = -12
    if (p0 == p0) {
        switch "\n\t" {
        case "":
            v6 := GenS1{F0: 42, F1: 98, F2: true}
            _ = v6
        case "\n\t":
            for fi7 := 0; fi7 < 16; fi7++ {
                v8 := GenS2{F0: 32767, F1: 18446744073709551615}
                _ = fi7
                _ = v8
            }
        case "foo bar":
            if (0 != -9) {
                v5 += -2147483648
            }
        default:
            fmt.Println(v5, v5)
        }
        var v9 uint = p0
        _ = v9
    }
    v5 = 479
    switch 556 {
    case 1:
        switch 70 {
        case 32768:
            v10 := []int{989, -9223372036854775808, 128}
            _ = v10
        case 127:
            v5 = -610
        default:
            var v11 int8 = 81
            _ = v11
        }
    case 4294967295:
        var v12 uint64 = 2
        _ = v12
    default:
        v13 := []int{-696}
        _ = v13
    }
    var v14 int = 1
    _ = p0
    _ = v14
    _ = v5
    return int32(p0)
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]bool)
    var v3 string = "\n\t"
    var v4 int8 = 42
    var v15 uint16 = 255
    v16 := []int{}
    for fi17 := 0; fi17 < 5; fi17++ {
        _ = len(v2)
        v3 = (v3 + "")
        _ = fi17
    }
    v4 = (v4 - 127)
    for fi18 := 0; fi18 < 7; fi18++ {
        var v19 float64 = 3.14159
        var v20 float32 = 0.0
        v21 := GenS3{F0: 741}
        _ = fi18
        _ = v19
        _ = v20
        _ = v21
    }
    _ = len(v16)
    var v22 int = 9223372036854775807
    switch v3 {
    case "a":
        delete(v1, uint32(v15))
    case "hello":
        for fi23 := 0; fi23 < 6; fi23++ {
            fmt.Println(v4, v3, v22)
            v4 = (v4 & v4)
            _ = fi23
        }
    case "foo bar":
        if true {
            var v24 bool = (42 <= int8(v22))
            v25 := GenS2{F0: 1, F1: 91}
            v3 = (v3 + (v3 + v3))
            _ = v24
            _ = v25
        }
    default:
        v3 = v3
    }
    if (((!false) || (!(18446744073709551615 < 1))) && true) {
        for fi26 := 0; fi26 < 20; fi26++ {
            if ((false && false) && true) {
                var v27 uint32 = 0
                v28 := GenS3{F0: 476}
                var v29 uint = 18446744073709551615
                _ = v27
                _ = v28
                _ = v29
            } else {
                v15 = (v15 >> ((193) & 63))
                fmt.Println(v15)
            }
            fmt.Println(v3, v22)
            v30 := GenS4{F0: 568, F1: -7}
            _ = fi26
            _ = v30
        }
        var v31 bool = true
        var v32 uint32 = 2
        _ = v31
        _ = v32
    } else {
        var v33 int = v22
        _ = v33
    }
    var v34 int64 = 42
    fmt.Println(v15)
    v35 := GenS4{F0: v34, F1: 100}
    v36 := GenS3{F0: 1}
    v37 := GenS5{F0: 128, F1: 42}
    fmt.Println(v3, v4, v15)
    delete(v2, -2147483648)
    v38 := GenS3{F0: v22}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
}
