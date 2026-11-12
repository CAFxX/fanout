package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 uint32 }
func genFn0(p0 uint32, p1 int16) int32 {
    p0 -= (p0 + p0)
    if false {
        for fi5 := 0; fi5 < 12; fi5++ {
            fmt.Println(fi5, fi5, p1)
            v6 := GenS1{F0: 42}
            switch ("hello" + "a") {
            case "a":
                v7 := []int{38}
                _ = v7
            }
            _ = fi5
            _ = v6
        }
    }
    v8 := GenS2{F0: 128, F1: p0}
    _ = p0
    _ = p1
    _ = v8
    return -1
}
func genFn1() int32 {
    var v9 int8 = 0
    var v10 int64 = 0
    var v11 uint64 = 0
    _ = v10
    _ = v11
    _ = v9
    return -2147483648
}

func main() {
    v1 := make(map[string]int)
    v2 := []uint32{}
    v3 := []uint32{950, 42, 2147483647}
    var v4 uint8 = 42
    var v12 uint8 = v4
    v4 = 39
    var v13 uint8 = (^(v12))
    for fi14 := 0; fi14 < 19; fi14++ {
        for fi15 := 0; fi15 < 19; fi15++ {
            v4 -= (v12 & (^(v12)))
            _ = len(v1)
            _ = fi15
        }
        switch fi14 {
        case 410:
            for fi16 := 0; fi16 < 20; fi16++ {
                v4 = (^(v13))
                v4 += v4
                _ = fi16
            }
        default:
            v17 := GenS1{F0: 1}
            _ = v17
        }
        v1["a"] = fi14
        _ = fi14
    }
    var v18 uint16 = 32767
    for fi19 := 0; fi19 < 8; fi19++ {
        for fi20 := 0; fi20 < 12; fi20++ {
            var v21 int8 = 35
            v22 := GenS2{F0: -2147483648, F1: 1}
            var v23 uint8 = v13
            _ = fi20
            _ = v21
            _ = v22
            _ = v23
        }
        _ = fi19
    }
    v13 -= v13
    for fi24 := 0; fi24 < 14; fi24++ {
        var v25 string = "foo bar"
        _ = fi24
        _ = v25
    }
    v4 = (v13 ^ v13)
    var v26 uint64 = 0
    if (-2147483648 < 0) {
        fmt.Println(v12)
        var v27 float64 = ((1.5 + (((0.1 * ((1.5 * 1.5) - ((1.5 / float64((float32(1e10) - -2.25))) / 3.14159))) / (0.0 - 0.1)) - 0.1)) / (float64(-2.25) * (0.1 - (1.5 * (((0.1 * ((0.0 * (-2.25 / float64((float32(1.5) * 3.14159)))) / (1.5 - (1.5 * -2.25)))) + (float64(3.14159) - 1.5)) + 0.1)))))
        _ = v27
    } else {
        fmt.Println(v26)
    }
    var v28 uint8 = (^(v12))
    var v29 uint16 = (v18 | v18)
    fmt.Println(v4, v13, v18)
    var v30 uint = uint(v18)
    v12 = (v13 + v12)
    _ = len(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v26
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
