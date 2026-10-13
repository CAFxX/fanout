package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 string }
type GenS2 struct { F0 string; F1 uint32; F2 int64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint16; F1 int16; F2 int }
type GenS5 struct { F0 uint; F1 uint16; F2 string }
type GenS6 struct { F0 uint64; F1 int64; F2 string }
func genFn0(p0 bool, p1 string, p2 int32) int16 {
    switch 462 {
    case 9223372036854775807:
        if p0 {
            var v7 uint32 = 381
            _ = v7
        }
    }
    var v8 uint64 = 42
    for fi9 := 0; fi9 < 18; fi9++ {
        fmt.Println(p0)
        _ = fi9
    }
    fmt.Println(v8)
    var v10 int8 = -68
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    return -492
}
func genFn1() int64 {
    var v11 int32 = 2147483647
    v11 -= 0
    v11 += -2147483648
    _ = v11
    return 0
}

func main() {
    v1 := make(map[string]int)
    var v2 int32 = 108
    var v3 uint = 18446744073709551615
    var v4 uint = (^(v3))
    var v5 uint = v3
    v6 := make(map[int]int)
    for fi12 := 0; fi12 < 19; fi12++ {
        for fi13 := 0; fi13 < 2; fi13++ {
            var v14 uint16 = 2
            fmt.Println(fi13, v2, v3)
            if true {
                if (!((v3 == v3) || true)) {
                    v3 += (^(v5))
                }
                var v15 string = "hello"
                fmt.Println(v4, v5, v4)
                _ = v15
            } else {
                _ = len(v1)
                if false {
                    if true {
                        var v16 float64 = (1.5 * float64(((0.0 * 1e10) - 1.5)))
                        v14 = 0
                        v5 = (v3 + v4)
                        _ = v16
                    }
                    fmt.Println(v3, v3, v5)
                    var v17 uint32 = 0
                    _ = v17
                } else {
                    fmt.Println(fi13, v2)
                }
                fmt.Println(v3, v14, v14)
            }
            _ = fi13
            _ = v14
        }
        switch fi12 {
        case 32768:
            v5 += v5
        case 42:
            v2 = v2
        case -9223372036854775808:
            v2 = (v2 & 42)
        default:
            v5 = (v5 & 18446744073709551615)
        }
        _ = fi12
    }
    if false {
        if (!true) {
            v18 := GenS1{F0: 2, F1: 237, F2: ("hello" + "")}
            _ = len(v6)
            _ = v18
        }
        v19 := GenS2{F0: ("" + ""), F1: 2, F2: 0}
        _ = v19
    }
    v20 := GenS2{F0: "foo bar", F1: 4294967295, F2: 1}
    for fi21 := 0; fi21 < 20; fi21++ {
        var v22 uint16 = 930
        v23 := GenS1{F0: 900, F1: 65535, F2: "\n\t"}
        v3 = (v5 / ((v5) | 1))
        _ = fi21
        _ = v22
        _ = v23
    }
    var v24 uint64 = 4294967295
    v6[128] = 4294967295
    var v25 float64 = ((1.5 / (1.5 / 1e10)) - 1.5)
    var v26 float64 = v25
    v26 -= -2.25
    fmt.Println(v3, v3, v25)
    switch -463 {
    case 1:
        v27 := GenS3{F0: 32767}
        _ = v27
    default:
        v28 := GenS3{F0: 0}
        _ = v28
    }
    fmt.Println(v24, v25)
    if (2 < 32767) {
        if true {
            v29 := GenS4{F0: 412, F1: 264, F2: 0}
            v30 := GenS5{F0: (v5 ^ (v5 + 0)), F1: uint16(v3), F2: "\n\t"}
            v31 := []int{4294967295}
            _ = v29
            _ = v30
            _ = v31
        }
        v26 = v26
    }
    v2 += 42
    v32 := GenS6{F0: (^(v24)), F1: -1, F2: ("foo bar" + ("" + "\n\t"))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
