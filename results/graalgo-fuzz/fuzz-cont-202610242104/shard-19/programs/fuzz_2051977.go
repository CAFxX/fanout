package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }

func main() {
    v1 := make(map[string]bool)
    var v2 uint8 = 248
    var v3 uint16 = 32767
    var v4 int = -5
    for fi5 := 0; fi5 < 19; fi5++ {
        for fi6 := 0; fi6 < 15; fi6++ {
            var v7 float32 = (0.0 / 1.5)
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    var v8 int64 = 0
    var v9 uint64 = 910
    v10 := GenS1{F0: 1, F1: (v8 & v8)}
    for fi11 := 0; fi11 < 7; fi11++ {
        v1["foo bar"] = (-1 >= 1)
        if true {
            var v12 float64 = float64(0.0)
            var v13 uint32 = 975
            _ = v12
            _ = v13
        } else {
            if (!true) {
                if (127 <= 127) {
                    switch v4 {
                    case -357:
                        v2 = (v2 >> ((255) & 63))
                    case 42:
                        v3 += (v3 + (v3 + (v3 & v3)))
                    default:
                        var v14 int = (fi11 / (((v4 ^ -9223372036854775808)) | 1))
                        _ = v14
                    }
                }
                var v15 uint8 = v2
                _ = v15
            } else {
                v3 += (v3 >> ((843) & 63))
            }
            v16 := GenS1{F0: 0, F1: (v8 ^ v8)}
            _ = v16
        }
        v3 -= (v3 + (v3 | v3))
        _ = fi11
    }
    switch v4 {
    case -199:
        var v17 uint32 = 315
        _ = v17
    default:
        _ = len(v1)
    }
    var v18 int = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
