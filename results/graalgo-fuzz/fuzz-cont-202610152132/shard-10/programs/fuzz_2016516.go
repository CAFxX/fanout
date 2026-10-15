package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 int16 }
type GenS2 struct { F0 int64; F1 int64; F2 int8 }
type GenS3 struct { F0 uint16; F1 string; F2 uint64 }
type GenS4 struct { F0 int64; F1 string }
func genFn0(p0 uint16, p1 bool, p2 int) int {
    fmt.Println(p0, p2)
    p2 = (p2 - (p2 * (p2 | p2)))
    _ = p0
    _ = p1
    _ = p2
    return 0
}
func genFn1(p0 int16) uint64 {
    for fi6 := 0; fi6 < 12; fi6++ {
        v7 := []int{}
        fmt.Println(p0, p0)
        _ = fi6
        _ = v7
    }
    p0 = p0
    for fi8 := 0; fi8 < 12; fi8++ {
        var v9 string = ""
        v10 := []int{(fi8 - fi8), fi8, 127, 0}
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    return 1
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]string)
    var v3 uint64 = 0
    var v4 string = ("a" + "foo bar")
    var v5 int64 = 1
    if ("a" <= (("hello" + (v4 + v4)) + v4)) {
        v5 += 42
        var v11 uint16 = 255
        _ = v11
    }
    var v12 uint32 = 42
    v12 = (v12 % ((v12) | 1))
    for fi13 := 0; fi13 < 19; fi13++ {
        for fi14 := 0; fi14 < 18; fi14++ {
            fmt.Println(v4, v12)
            v15 := GenS1{F0: 1000, F1: 1, F2: 1000}
            v16 := GenS1{F0: -47, F1: 255, F2: 1}
            if false {
                var v17 uint16 = 673
                v3 += (v3 + 4294967295)
                _ = v17
            }
            _ = fi14
            _ = v15
            _ = v16
        }
        v1["\n\t"] = (v12 ^ (v12 / ((2147483648) | 1)))
        for fi18 := 0; fi18 < 6; fi18++ {
            var v19 float32 = float32((0.0 * float64(0.1)))
            _ = fi18
            _ = v19
        }
        _ = fi13
    }
    var v20 uint16 = 896
    switch v4 {
    case "":
        switch v4 {
        case "hello":
            var v21 int64 = (v5 >> ((102) & 63))
            _ = v21
        case "":
            v22 := GenS2{F0: -1, F1: v5, F2: -1}
            _ = v22
        }
    case "a":
        var v23 uint16 = v20
        _ = v23
    }
    v24 := v1[(v4 + "a")]
    v25 := []int{2147483647}
    var v26 int32 = -653
    v27 := GenS3{F0: (v20 + v20), F1: "a", F2: v3}
    switch 309 {
    case 284:
        var v28 int8 = 100
        _ = v28
    case 32768:
        for fi29 := 0; fi29 < 12; fi29++ {
            var v30 uint = 132
            var v31 int64 = (v5 >> (((^(v30))) & 63))
            _ = fi29
            _ = v30
            _ = v31
        }
    case -9223372036854775808:
        if len(v25) > 0 {
            v25[0] = -9223372036854775808
        }
    default:
        v32 := GenS4{F0: (v5 % ((int64(v3)) | 1)), F1: (((("a" + ("hello" + v4)) + "") + v4) + "hello")}
        _ = v32
    }
    v26 -= v26
    fmt.Println(v26)
    for fi33 := 0; fi33 < 8; fi33++ {
        v26 = (v26 / ((v26) | 1))
        switch 32768 {
        case 128:
            v34 := GenS2{F0: (v5 << ((255) & 63)), F1: (v5 + (v5 - v5)), F2: 0}
            _ = v34
        default:
            var v35 int8 = 100
            _ = v35
        }
        _ = fi33
    }
    var v36 bool = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v36
    _ = v4
    _ = v5
}
