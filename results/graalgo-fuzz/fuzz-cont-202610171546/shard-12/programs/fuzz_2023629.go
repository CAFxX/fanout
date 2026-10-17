package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 uint8 }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint }
func genFn0(p0 uint64, p1 int32, p2 int) int {
    var v4 int16 = 551
    var v5 float32 = (-2.25 / 3.14159)
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return p2
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint = 263
    var v3 uint = (v2 / (((v2 % ((18446744073709551615) | 1))) | 1))
    v2 = (v2 >> ((v2) & 63))
    if (!false) {
        var v6 float32 = 1e10
        _ = v6
    }
    v7 := GenS1{F0: 255, F1: 25, F2: 255}
    v8 := GenS1{F0: 366, F1: 128, F2: 86}
    switch "foo bar" {
    case "\n\t":
        v9 := GenS2{F0: "\n\t", F1: uint16(v2)}
        _ = v9
    case "":
        v10 := v1[""]
        _ = v10
    case "foo bar":
        v3 -= (v2 / ((v3) | 1))
    }
    v11 := GenS3{F0: ""}
    v12 := GenS2{F0: "\n\t", F1: 42}
    var v13 bool = false
    for fi14 := 0; fi14 < 13; fi14++ {
        _ = len(v1)
        _ = fi14
    }
    var v15 int64 = 42
    if v13 {
        v15 -= (v15 + (^(v15)))
    } else {
        v16 := GenS4{F0: (v3 * (v2 * v3))}
        v17 := []int{}
        switch genFn0(18446744073709551615, 0, -902) {
        case 32768:
            var v18 float32 = (3.14159 - 1e10)
            _ = v18
        case 2147483647:
            if (127 != 127) {
                var v19 uint16 = 2
                v15 -= (v15 / (((v15 % ((v15) | 1))) | 1))
                _ = v19
            } else {
                for fi20 := 0; fi20 < 7; fi20++ {
                    v21 := GenS4{F0: (v2 ^ v3)}
                    _ = fi20
                    _ = v21
                }
            }
        case 42:
            var v22 uint16 = 0
            _ = v22
        default:
            v15 -= (^(v15))
        }
        _ = v16
        _ = v17
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
