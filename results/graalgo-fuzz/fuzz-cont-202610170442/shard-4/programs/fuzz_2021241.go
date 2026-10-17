package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16; F1 string }
type GenS4 struct { F0 uint16 }

func main() {
    var v1 float32 = (3.14159 / (1e10 - (0.0 / ((3.14159 - ((float32(1.5) * (1e10 * 1.5)) + 1.5)) - 0.1))))
    var v2 int16 = -1
    v3 := make(map[string]bool)
    var v4 uint8 = 220
    v5 := make(map[int]string)
    var v6 int = 219
    switch 128 {
    case -657:
        v2 += (-(v2))
    case 856:
        v7 := GenS1{F0: 1, F1: (v2 / ((v2) | 1))}
        _ = v7
    default:
        if ((v2 - v2) >= 1000) {
            v8 := GenS1{F0: 1, F1: int16(v4)}
            var v9 int64 = 9223372036854775807
            _ = v8
            _ = v9
        } else {
            v4 += 255
            v6 += v6
            for fi10 := 0; fi10 < 5; fi10++ {
                if (!(false || (-31 == 549))) {
                    v6 -= fi10
                } else {
                    v6 = (fi10 >> ((1) & 63))
                }
                if (false && (0 < 1)) {
                    v1 = float32((0.0 / (1.5 - 1e10)))
                }
                _ = fi10
            }
        }
    }
    v1 = (3.14159 - (v1 + -2.25))
    if true {
        v6 -= (v6 | v6)
    } else {
        if true {
            fmt.Println(v1)
            var v11 uint8 = v4
            v6 -= v6
            _ = v11
        } else {
            var v12 float32 = v1
            _ = v12
        }
    }
    switch ("hello" + "") {
    case "foo bar":
        v4 = v4
    default:
        v1 = v1
    }
    v13 := GenS2{F0: 60}
    if (!false) {
        var v14 uint = 833
        var v15 float32 = v1
        if (0 != 0) {
            switch v6 {
            case 1:
                v1 += ((v1 / (1.5 / ((float32(float64(v1)) * (((v15 + 0.1) - v1) + 3.14159)) * (0.1 / 1.0)))) * float32(0.1))
            case 581:
                var v16 bool = ((v2 * v2) == (^(v2)))
                _ = v16
            default:
                var v17 int = v6
                _ = v17
            }
            v15 = ((v1 / 0.1) + (float32(0.0) + float32(1.5)))
            v18 := GenS3{F0: 2, F1: "a"}
            _ = v18
        } else {
            v19 := []int{32768, 4294967295, 117}
            v15 = 3.14159
            if (193 > v4) {
                v1 = ((v1 + (float32(0.0) - v1)) - (3.14159 * 0.0))
                v20 := GenS4{F0: 32767}
                _ = v20
            }
            _ = v19
        }
        _ = v14
        _ = v15
    }
    _ = len(v3)
    v1 = v1
    var v21 uint32 = 1
    var v22 uint64 = 1
    var v23 int8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
