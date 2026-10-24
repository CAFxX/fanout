package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 int; F2 uint }
type GenS3 struct { F0 uint64; F1 uint8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint32; F1 int64; F2 int32 }
type GenS6 struct { F0 uint16; F1 string }
type GenS7 struct { F0 int16; F1 string }
func genFn0(p0 string, p1 int8) uint8 {
    fmt.Println(p0)
    fmt.Println(p1, p1)
    p0 = p0
    var v6 uint8 = 1
    _ = p0
    _ = p1
    _ = v6
    return 162
}
func genFn1(p0 int) bool {
    v7 := GenS1{F0: 0}
    for fi8 := 0; fi8 < 9; fi8++ {
        p0 -= (p0 << ((42) & 63))
        var v9 string = "\n\t"
        v10 := []int{(p0 / ((4294967295) | 1)), (fi8 >> ((255) & 63)), -312, fi8}
        _ = fi8
        _ = v10
        _ = v9
    }
    p0 = p0
    v11 := GenS2{F0: 65535, F1: 42, F2: 2}
    var v12 bool = false
    _ = p0
    _ = v11
    _ = v12
    _ = v7
    return v12
}

func main() {
    var v1 string = "a"
    v2 := []uint32{}
    v3 := make(map[int]int)
    var v4 int32 = -2147483648
    var v5 int32 = -2147483648
    fmt.Println(v5, v4)
    v5 = -2147483648
    if (false || true) {
        fmt.Println(v1, v1)
    }
    v4 = -584
    for fi13 := 0; fi13 < 2; fi13++ {
        v4 = 0
        var v14 uint = 278
        for fi15 := 0; fi15 < 6; fi15++ {
            v16 := GenS1{F0: 255}
            var v17 int32 = -733
            if true {
                v18 := GenS3{F0: 1, F1: 1}
                _ = v18
            }
            v19 := GenS4{F0: 2}
            fmt.Println(v5, v4)
            fmt.Println(v5, fi15)
            _ = fi15
            _ = v16
            _ = v17
            _ = v19
        }
        _ = fi13
        _ = v14
    }
    if (!(!(true || true))) {
        for fi20 := 0; fi20 < 19; fi20++ {
            v21 := v3[-993]
            _ = fi20
            _ = v21
        }
        v1 = (v1 + v1)
        v22 := GenS5{F0: 42, F1: -192, F2: (v5 | 1)}
        _ = v22
    }
    fmt.Println(v1, v1)
    v1 = (("" + v1) + "")
    v23 := GenS6{F0: 32767, F1: "hello"}
    v24 := GenS6{F0: 2, F1: v1}
    var v25 int32 = 0
    fmt.Println(v4, v1)
    v4 = (v25 | v5)
    v1 = v1
    if (false && false) {
        switch 316 {
        case -119:
            for fi26 := 0; fi26 < 19; fi26++ {
                v5 += (v4 + -2147483648)
                _ = fi26
            }
        default:
            v1 = (v1 + ("a" + v1))
        }
        fmt.Println(v5, v4, v5)
    } else {
        switch 9223372036854775807 {
        case 2147483647:
            v27 := GenS4{F0: 613}
            _ = v27
        case 128:
            _ = len(v2)
        }
        v28 := GenS7{F0: 32767, F1: ""}
        _ = v28
    }
    switch (("foo bar" + "hello") + ((v1 + "hello") + "foo bar")) {
    case "foo bar":
        v4 += v5
    case "\n\t":
        var v29 uint = 255
        _ = v29
    }
    var v30 uint32 = 2147483647
    v25 = (v25 / ((v4) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
