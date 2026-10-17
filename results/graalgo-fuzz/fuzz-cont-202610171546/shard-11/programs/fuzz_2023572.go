package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 int16 }

func main() {
    var v1 int8 = -66
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]int)
    var v4 int64 = 571
    var v5 uint = 442
    fmt.Println(v1)
    for fi6 := 0; fi6 < 11; fi6++ {
        v4 = -9223372036854775808
        for fi7 := 0; fi7 < 4; fi7++ {
            var v8 uint32 = 42
            v4 = 1
            var v9 int32 = 0
            _ = fi7
            _ = v8
            _ = v9
        }
        fmt.Println(fi6)
        _ = fi6
    }
    v10 := GenS1{F0: (1e10 >= (1.5 * (float64(0.1) - ((-2.25 * ((float64(1.5) / (3.14159 - (float64(1.5) + 1e10))) * ((0.0 / float64(3.14159)) - 0.0))) + ((1e10 / -2.25) * 0.1)))))}
    fmt.Println(v5, v5, v4)
    v11 := GenS2{F0: -32768, F1: 0}
    delete(v2, 2147483647)
    var v12 float64 = 3.14159
    v12 = ((v12 - v12) - (v12 / v12))
    var v13 uint = 576
    switch (("foo bar" + "foo bar") + "\n\t") {
    case "":
        v5 = 255
    case "foo bar":
        switch 127 {
        case 32768:
            v14 := []int{-2147483648, 42}
            _ = v14
        case 408:
            var v15 uint32 = 0
            _ = v15
        case 32767:
            var v16 int32 = -951
            _ = v16
        default:
            v4 = 0
        }
    }
    for fi17 := 0; fi17 < 19; fi17++ {
        if (v13 == 654) {
            var v18 uint16 = 1000
            var v19 float32 = -2.25
            _ = len(v3)
            _ = v18
            _ = v19
        } else {
            var v20 bool = (!(534 <= 4294967295))
            _ = v20
        }
        _ = fi17
    }
    for fi21 := 0; fi21 < 15; fi21++ {
        for fi22 := 0; fi22 < 10; fi22++ {
            if (2 != 452) {
                v23 := []int{(fi22 + (-(fi21))), (^(fi21)), 9223372036854775807}
                _ = v23
            }
            var v24 int = fi21
            _ = fi22
            _ = v24
        }
        v4 = 0
        _ = fi21
    }
    v13 = (v13 | 255)
    v1 = v1
    v25 := GenS2{F0: 1000, F1: -1}
    for fi26 := 0; fi26 < 20; fi26++ {
        var v27 float64 = float64(1.5)
        _ = fi26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
