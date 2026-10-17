package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint; F1 uint; F2 bool }
func genFn0(p0 uint, p1 uint, p2 uint8) bool {
    for fi5 := 0; fi5 < 19; fi5++ {
        p2 += 42
        _ = fi5
    }
    if true {
        p2 -= 246
        p1 = (p0 - (p1 / ((p0) | 1)))
    }
    fmt.Println(p1, p2, p0)
    _ = p0
    _ = p1
    _ = p2
    return (!(-75 >= 964))
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]bool)
    v4 := make(map[int]uint32)
    var v6 float32 = (1e10 - (-2.25 / ((1e10 * 0.1) / 1.5)))
    fmt.Println(v6)
    fmt.Println(v6)
    v7 := GenS1{F0: -32768, F1: 100}
    for fi8 := 0; fi8 < 1; fi8++ {
        fmt.Println(v6)
        v6 = float32(-2.25)
        v6 -= (1e10 + (v6 * (float32(-2.25) + v6)))
        _ = fi8
    }
    v9 := GenS2{F0: (42 >= 0)}
    v10 := GenS1{F0: 32767, F1: 90}
    switch "hello" {
    case "foo bar":
        for fi11 := 0; fi11 < 2; fi11++ {
            v6 = float32((1.5 / (((3.14159 - 1.5) * 0.1) / 1e10)))
            v6 += v6
            v12 := GenS1{F0: 32767, F1: -1}
            _ = fi11
            _ = v12
        }
    case "\n\t":
        if (!true) {
            v6 = v6
        } else {
            var v13 int32 = 989
            fmt.Println(v6)
            var v14 float64 = ((1.5 - 1.5) / 1e10)
            _ = v13
            _ = v14
        }
    default:
        for fi15 := 0; fi15 < 6; fi15++ {
            var v16 uint = 4294967295
            v6 += v6
            _ = fi15
            _ = v16
        }
    }
    v17 := GenS2{F0: false}
    if (0 > 634) {
        v18 := GenS3{F0: 65535, F1: 892, F2: true}
        for fi19 := 0; fi19 < 19; fi19++ {
            v6 = v6
            var v20 bool = (-1 <= -32768)
            _ = fi19
            _ = v20
        }
        var v21 uint16 = 1
        _ = v18
        _ = v21
    } else {
        var v22 float32 = (v6 - (1.5 + ((0.0 + 0.1) - 1.5)))
        _ = v22
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v9
}
