package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint }
func genFn0(p0 int64, p1 int8, p2 int16) int64 {
    switch (("a" + "") + "hello") {
    case "":
        var v6 uint8 = 42
        _ = v6
    }
    fmt.Println(p0, p0)
    p2 -= (-(p2))
    if true {
        v7 := []int{32768, 2147483647, -336}
        _ = v7
    } else {
        v8 := []int{127, 128, 127, 280}
        fmt.Println(p2, p1)
        if false {
            p1 += -1
            var v9 uint16 = 181
            _ = v9
        } else {
            p0 = (p0 / ((p0) | 1))
            var v10 int = 0
            var v11 uint8 = 42
            _ = v10
            _ = v11
        }
        _ = v8
    }
    switch ("hello" + "") {
    case "foo bar":
        switch "\n\t" {
        case "hello":
            var v12 uint64 = 283
            _ = v12
        }
    case "hello":
        var v13 uint16 = 410
        _ = v13
    default:
        fmt.Println(p0, p1, p1)
    }
    _ = p0
    _ = p1
    _ = p2
    return -1
}

func main() {
    var v1 uint16 = 179
    v2 := make(map[uint32]uint32)
    var v3 uint8 = 42
    var v4 uint32 = 42
    v5 := make(map[string]bool)
    v1 += 2
    if (!false) {
        delete(v2, v4)
        for fi14 := 0; fi14 < 17; fi14++ {
            v15 := GenS1{F0: -1, F1: -305, F2: fi14}
            _ = fi14
            _ = v15
        }
    } else {
        var v16 int8 = -128
        _ = v16
    }
    fmt.Println(v4, v3)
    v17 := GenS1{F0: 0, F1: 1000, F2: 0}
    if (true || false) {
        switch -866 {
        case 987:
            v1 -= 756
        case 4294967295:
            var v18 uint = 0
            _ = v18
        case 32767:
            v3 -= (^(v3))
        default:
            if (false && (((0.1 / float64(1e10)) + 0.1) >= 3.14159)) {
                v19 := v5["\n\t"]
                for fi20 := 0; fi20 < 4; fi20++ {
                    v1 = 23
                    _ = fi20
                }
                _ = v19
            }
        }
        v21 := GenS2{F0: -128}
        switch "a" {
        case "\n\t":
            var v22 int = 217
            _ = v22
        case "a":
            if true {
                v1 = v1
            }
        }
        _ = v21
    } else {
        v4 = (v4 - 4294967295)
    }
    var v23 float64 = ((((1e10 + 1.5) + 0.0) + (-2.25 * 0.1)) / 1e10)
    var v24 int = 4294967295
    v25 := GenS1{F0: 0, F1: -32768, F2: (^(v24))}
    for fi26 := 0; fi26 < 15; fi26++ {
        v27 := GenS3{F0: 255}
        _ = fi26
        _ = v27
    }
    fmt.Println(v4)
    if true {
        var v28 int64 = 827
        fmt.Println(v24)
        _ = v28
    }
    var v29 int64 = -1
    v30 := GenS2{F0: 53}
    var v31 uint32 = (v4 << ((uint(v4)) & 63))
    var v32 int16 = -595
    var v33 int64 = v29
    var v34 uint8 = (^(v3))
    v31 = v31
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
