package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 string; F1 int8; F2 bool }
type GenS4 struct { F0 uint64; F1 uint8 }
type GenS5 struct { F0 bool; F1 uint8 }
func genFn0(p0 int) bool {
    switch -205 {
    case -1:
        var v6 uint = 42
        _ = v6
    case 32768:
        v7 := GenS1{F0: 664}
        _ = v7
    }
    v8 := []int{p0}
    v9 := GenS2{F0: 1}
    _ = p0
    _ = v8
    _ = v9
    return (255 < 18446744073709551615)
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]uint32)
    var v3 uint8 = 255
    var v4 uint8 = v3
    v5 := make(map[uint32]int)
    var v10 int64 = 1
    var v11 float64 = 0.1
    for fi12 := 0; fi12 < 6; fi12++ {
        v13 := GenS3{F0: "a", F1: 127, F2: false}
        fmt.Println(v4, v3)
        _ = fi12
        _ = v13
    }
    v14 := GenS4{F0: 18446744073709551615, F1: v4}
    v4 -= v3
    if (((128 == 0) || true) || (152 >= 65535)) {
        if (((false && true) || true) && (950 != 794)) {
            _ = len(v1)
        }
    }
    for fi15 := 0; fi15 < 19; fi15++ {
        var v16 int = (fi15 << ((86) & 63))
        _ = fi15
        _ = v16
    }
    v3 += (^(v3))
    var v17 float32 = 0.0
    switch 32767 {
    case -467:
        if (1 < 32767) {
            v3 = 2
            fmt.Println(v10, v3)
            var v18 int8 = -1
            _ = v18
        }
    case -2147483648:
        var v19 bool = false
        _ = v19
    case 32768:
        var v20 int = 1
        _ = v20
    default:
        if ("hello" == ("" + "foo bar")) {
            v4 = (v3 / (((v4 & v4)) | 1))
        }
    }
    var v21 float64 = (((-2.25 * float64(v17)) / v11) * (0.0 - 1.5))
    var v22 float64 = (v11 * v21)
    switch "a" {
    case "foo bar":
        v21 = (((float64(v17) * v22) / v22) * v21)
    case "\n\t":
        var v23 int32 = 1
        _ = v23
    case "hello":
        fmt.Println(v22, v3)
    default:
        for fi24 := 0; fi24 < 6; fi24++ {
            for fi25 := 0; fi25 < 17; fi25++ {
                v26 := GenS5{F0: (-2147483648 < -1), F1: v4}
                var v27 int32 = 566
                _ = fi25
                _ = v26
                _ = v27
            }
            var v28 int8 = -1
            _ = fi24
            _ = v28
        }
    }
    var v29 float64 = (1e10 + 1.5)
    v29 -= v11
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
