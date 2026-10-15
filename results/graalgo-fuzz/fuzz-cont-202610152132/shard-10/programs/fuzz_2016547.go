package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 string; F1 bool; F2 uint32 }

func main() {
    var v1 int16 = -729
    v2 := make(map[uint32]uint32)
    var v3 uint32 = 647
    var v4 int16 = v1
    v5 := make(map[uint32]bool)
    v6 := []string{""}
    v7 := GenS1{F0: 867, F1: 42}
    var v8 int = -9223372036854775808
    for fi9 := 0; fi9 < 13; fi9++ {
        var v10 uint = 402
        if (-721 > v4) {
            if false {
                var v11 float32 = (float32((0.1 * -2.25)) / 1.5)
                for fi12 := 0; fi12 < 7; fi12++ {
                    v13 := GenS2{F0: "a", F1: (!true), F2: (v3 ^ (v3 % ((529) | 1)))}
                    v10 = v10
                    _ = len(v2)
                    _ = fi12
                    _ = v13
                }
                var v14 int = 2147483647
                _ = v11
                _ = v14
            }
            if false {
                v15 := GenS1{F0: 252, F1: 447}
                _ = v15
            }
            fmt.Println(v10, v10, v3)
        }
        var v16 int32 = 802
        _ = fi9
        _ = v10
        _ = v16
    }
    switch ("foo bar" + (("" + "hello") + "")) {
    case "\n\t":
        v8 = (^(v8))
    case "":
        v1 += v1
    default:
        var v17 bool = false
        _ = v17
    }
    for fi18 := 0; fi18 < 13; fi18++ {
        v8 = v8
        v8 += (v8 * (v8 % ((v8) | 1)))
        switch ("foo bar" + "") {
        case "foo bar":
            for fi19 := 0; fi19 < 12; fi19++ {
                if false {
                    switch "\n\t" {
                    case "\n\t":
                        v8 += int(v4)
                    }
                }
                v4 += v1
                _ = fi19
            }
        default:
            fmt.Println(fi18, v8)
        }
        _ = fi18
    }
    var v20 uint16 = 358
    var v21 float64 = -2.25
    v3 += 613
    if (((v8 * (v8 / (((v8 << ((42) & 63))) | 1))) >= -454) || true) {
        for fi22 := 0; fi22 < 1; fi22++ {
            var v23 int64 = 9223372036854775807
            v23 -= (^(v23))
            v20 = (v20 >> ((18446744073709551615) & 63))
            _ = fi22
            _ = v23
        }
    }
    var v24 float64 = (((1e10 + (((float64(((((1.5 + 0.1) + 0.0) / -2.25) - (float32(0.0) * float32((0.1 - v21))))) + 1e10) + v21) / 3.14159)) * v21) / (-2.25 * 1.5))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
