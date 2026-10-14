package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint16; F1 bool }
type GenS5 struct { F0 int; F1 int32 }
type GenS6 struct { F0 uint32; F1 uint }
func genFn0(p0 uint64, p1 int, p2 int16) string {
    var v4 uint32 = 1
    switch "a" {
    case "hello":
        v5 := GenS1{F0: 42, F1: (18446744073709551615 == 65535)}
        _ = v5
    default:
        v6 := GenS2{F0: (p1 % (((-(p1))) | 1))}
        _ = v6
    }
    p2 -= p2
    v7 := GenS3{F0: (("hello" + ("" + "a")) + "a")}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v7
    return "foo bar"
}
func genFn1() int64 {
    var v8 float64 = float64((1e10 + 3.14159))
    v9 := GenS2{F0: 1}
    v10 := GenS1{F0: 255, F1: true}
    v11 := GenS3{F0: "a"}
    var v12 int = -1
    _ = v10
    _ = v11
    _ = v12
    _ = v8
    _ = v9
    return 0
}

func main() {
    v1 := []uint32{2147483648, 673}
    v2 := make(map[int]int)
    var v3 int32 = 42
    var v13 uint64 = 602
    v14 := GenS4{F0: 64, F1: false}
    var v15 float32 = 0.1
    v13 += v13
    for fi16 := 0; fi16 < 7; fi16++ {
        fmt.Println(v15, v15, v13)
        v17 := GenS2{F0: (fi16 << ((269) & 63))}
        v18 := GenS1{F0: 1, F1: false}
        _ = fi16
        _ = v17
        _ = v18
    }
    v19 := GenS5{F0: 2147483647, F1: -115}
    var v20 string = "foo bar"
    var v21 uint16 = 255
    v22 := GenS4{F0: 65535, F1: true}
    switch "a" {
    case "\n\t":
        var v23 float32 = float32(3.14159)
        _ = v23
    }
    var v24 int32 = 1
    v25 := GenS6{F0: 117, F1: 1}
    if true {
        for fi26 := 0; fi26 < 1; fi26++ {
            v24 += -953
            _ = fi26
        }
    }
    if (v3 <= v3) {
        if false {
            v27 := GenS5{F0: 2147483647, F1: (v3 / (((v24 * (v3 - v3))) | 1))}
            _ = v27
        } else {
            var v28 uint8 = 255
            var v29 int = -2147483648
            fmt.Println(v21, v28, v15)
            _ = v28
            _ = v29
        }
    } else {
        v15 = v15
    }
    var v30 uint32 = 2147483648
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v24
    _ = v25
    _ = v3
    _ = v30
}
