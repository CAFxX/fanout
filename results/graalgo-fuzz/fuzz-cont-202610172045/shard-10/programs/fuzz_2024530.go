package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 int; F1 int; F2 int8 }
type GenS3 struct { F0 int16 }
func genFn0(p0 int32) int8 {
    var v4 uint8 = 0
    for fi5 := 0; fi5 < 9; fi5++ {
        switch (fi5 % (((fi5 | (fi5 >> ((1) & 63)))) | 1)) {
        case 0:
            fmt.Println(p0)
        case 42:
            fmt.Println(v4)
        default:
            var v6 int16 = -512
            _ = v6
        }
        v7 := GenS1{F0: 9223372036854775807, F1: true}
        v4 = 99
        _ = fi5
        _ = v7
    }
    for fi8 := 0; fi8 < 11; fi8++ {
        var v9 int8 = 127
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = v4
    return 1
}

func main() {
    var v1 uint = 0
    v2 := []int{}
    var v3 float64 = (((0.0 / 1.0) - ((((0.1 - 0.0) + 3.14159) + 0.0) - 1e10)) + 0.1)
    fmt.Println(v1, v1)
    v10 := GenS2{F0: -402, F1: 468, F2: -128}
    v1 = (v1 % ((4294967295) | 1))
    if false {
        v3 -= v3
        if true {
            fmt.Println(v1, v3)
        } else {
            var v11 uint64 = 474
            _ = v11
        }
        if true {
            v12 := GenS3{F0: 877}
            var v13 int8 = 42
            _ = v12
            _ = v13
        }
    }
    v2 = append(v2, 32768)
    var v14 uint16 = 32767
    var v15 uint16 = 65535
    var v16 int = 459
    switch "foo bar" {
    case "a":
        var v17 uint8 = 2
        _ = v17
    case "\n\t":
        var v18 uint8 = 127
        _ = v18
    default:
        var v19 int32 = 0
        _ = v19
    }
    if false {
        if true {
            var v20 string = "a"
            _ = v20
        }
        v2 = append(v2, v16)
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v21 uint8 = 2
    var v22 int16 = -218
    for fi23 := 0; fi23 < 1; fi23++ {
        v22 += v22
        _ = fi23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v3
}
