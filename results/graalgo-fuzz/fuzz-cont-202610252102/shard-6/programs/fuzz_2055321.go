package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 uint; F2 int32 }
type GenS3 struct { F0 uint64; F1 uint32; F2 int32 }
type GenS4 struct { F0 uint8; F1 int; F2 uint16 }
func genFn0(p0 uint32, p1 uint16) bool {
    p1 -= p1
    if (127 <= 98) {
        for fi5 := 0; fi5 < 17; fi5++ {
            v6 := GenS1{F0: "foo bar"}
            _ = fi5
            _ = v6
        }
    } else {
        var v7 uint32 = (p0 / ((p0) | 1))
        _ = v7
    }
    if (-1 <= -109) {
        p1 += p1
    }
    var v8 string = ""
    var v9 int = 9223372036854775807
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return true
}
func genFn1(p0 bool) int64 {
    switch "\n\t" {
    case "\n\t":
        v10 := GenS1{F0: ""}
        _ = v10
    case "a":
        var v11 string = ("hello" + (("foo bar" + "\n\t") + ""))
        _ = v11
    }
    for fi12 := 0; fi12 < 7; fi12++ {
        v13 := GenS2{F0: 42, F1: 4294967295, F2: -523}
        _ = fi12
        _ = v13
    }
    if p0 {
        p0 = false
    } else {
        var v14 uint32 = 1
        for fi15 := 0; fi15 < 14; fi15++ {
            var v16 bool = (p0 || (3.14159 > (((0.1 * 1e10) - 1.5) + 0.1)))
            for fi17 := 0; fi17 < 15; fi17++ {
                p0 = false
                v16 = p0
                p0 = (65535 >= 255)
                _ = fi17
            }
            _ = fi15
            _ = v16
        }
        p0 = (-746 < 885)
        _ = v14
    }
    v18 := []int{}
    _ = p0
    _ = v18
    return -9223372036854775808
}

func main() {
    var v1 int8 = 37
    v2 := []uint32{2147483648, 295, 0}
    v3 := make(map[int]int)
    v4 := []uint32{}
    _ = len(v2)
    fmt.Println(v1)
    switch 0 {
    case 128:
        fmt.Println(v1)
    case 604:
        for fi19 := 0; fi19 < 10; fi19++ {
            var v20 int64 = -560
            v1 -= v1
            v21 := GenS3{F0: 2, F1: 455, F2: -2147483648}
            _ = fi19
            _ = v20
            _ = v21
        }
    case 0:
        for fi22 := 0; fi22 < 5; fi22++ {
            fmt.Println(fi22, v1)
            for fi23 := 0; fi23 < 15; fi23++ {
                v24 := GenS2{F0: -9223372036854775808, F1: 65535, F2: -27}
                v4 = append(v4, 0)
                _ = fi23
                _ = v24
            }
            _ = fi22
        }
    }
    v25 := GenS2{F0: 0, F1: 255, F2: 1}
    v1 += 127
    v1 -= (v1 / ((127) | 1))
    if true {
        v1 -= v1
    } else {
        for fi26 := 0; fi26 < 8; fi26++ {
            switch 32768 {
            case 32767:
                v1 -= v1
            case -638:
                v1 = -51
            }
            _ = fi26
        }
        var v27 int = -2147483648
        v28 := GenS4{F0: 174, F1: -2147483648, F2: 42}
        _ = v27
        _ = v28
    }
    var v29 int = 1
    v29 = v29
    fmt.Println(v29, v29)
    for fi30 := 0; fi30 < 4; fi30++ {
        var v31 int32 = -1
        if ("a" < "hello") {
            var v32 uint = 1
            _ = v32
        }
        _ = fi30
        _ = v31
    }
    v29 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v4
}
