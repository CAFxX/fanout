package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 int32 }
type GenS2 struct { F0 string; F1 uint8; F2 int }
type GenS3 struct { F0 int; F1 uint32; F2 uint8 }
func genFn0(p0 uint32, p1 uint32, p2 int) uint32 {
    for fi6 := 0; fi6 < 20; fi6++ {
        p2 = fi6
        var v7 int = 32767
        _ = fi6
        _ = v7
    }
    if (false || true) {
        var v8 string = "\n\t"
        _ = v8
    }
    p1 -= (p0 & p1)
    switch p2 {
    case 128:
        if (!(-1 != int16(p1))) {
            var v9 int32 = 42
            _ = v9
        }
    default:
        fmt.Println(p0, p2, p0)
    }
    _ = p0
    _ = p1
    _ = p2
    return 2
}

func main() {
    v1 := []uint32{}
    var v2 int64 = 1
    var v3 float64 = float64(1e10)
    var v4 float32 = ((float32(0.0) + 0.0) + 0.0)
    var v5 int16 = 32767
    var v10 int16 = v5
    var v11 float32 = (-2.25 * float32(-2.25))
    if false {
        fmt.Println(v5)
    } else {
        for fi12 := 0; fi12 < 10; fi12++ {
            fmt.Println(v11)
            var v13 uint32 = genFn0(65535, 134, -494)
            v14 := GenS1{F0: 0, F1: 4294967295, F2: 1}
            _ = fi12
            _ = v13
            _ = v14
        }
        v15 := GenS2{F0: "foo bar", F1: 99, F2: -952}
        if true {
            var v16 uint32 = 967
            _ = v16
        }
        _ = v15
    }
    if (!false) {
        switch "foo bar" {
        case "":
            fmt.Println(v2, v11, v4)
        case "foo bar":
            for fi17 := 0; fi17 < 19; fi17++ {
                v1 = append(v1, genFn0(344, 662, 42))
                _ = fi17
            }
        case "\n\t":
            v18 := GenS2{F0: "hello", F1: 54, F2: 350}
            _ = v18
        default:
            v3 = (1e10 * v3)
        }
        fmt.Println(v11)
        if false {
            v19 := GenS3{F0: -1, F1: 42, F2: 127}
            _ = v19
        }
    }
    v2 = (v2 | v2)
    var v20 int64 = (v2 >> ((1) & 63))
    if (32767 <= 255) {
        var v21 uint64 = 0
        _ = v21
    }
    var v22 float64 = -2.25
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
