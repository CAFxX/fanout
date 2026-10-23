package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8; F1 int64; F2 uint64 }

func main() {
    v1 := []uint32{209, 2147483647, 1}
    var v2 int16 = 1
    v3 := make(map[uint32]bool)
    if (!(false || true)) {
        var v4 float64 = (float64(((((0.0 - -2.25) * -2.25) / (float32(-2.25) - 3.14159)) - 0.1)) - (1.5 / 1e10))
        switch 544 {
        case 32767:
            v4 -= (v4 * (v4 / (0.0 - 1e10)))
        }
        switch "" {
        case "hello":
            switch 1 {
            case 473:
                v5 := GenS1{F0: (v2 * (v2 | 1))}
                _ = v5
            default:
                for fi6 := 0; fi6 < 6; fi6++ {
                    var v7 int = (fi6 & (fi6 >> ((1) & 63)))
                    _ = fi6
                    _ = v7
                }
            }
        case "\n\t":
            for fi8 := 0; fi8 < 8; fi8++ {
                if (true && false) {
                    v2 = v2
                    _ = len(v1)
                } else {
                    v2 -= v2
                }
                var v9 uint64 = 0
                _ = fi8
                _ = v9
            }
        }
        _ = v4
    }
    var v10 uint8 = 255
    v10 = (v10 * v10)
    delete(v3, 493)
    _ = len(v1)
    v11 := GenS1{F0: (v2 << ((42) & 63))}
    if false {
        v2 = (v2 >> ((65535) & 63))
    } else {
        var v12 int = int(v2)
        v13 := GenS2{F0: 18446744073709551615}
        v10 += (v10 >> ((814) & 63))
        _ = v12
        _ = v13
    }
    switch 32768 {
    case -178:
        var v14 int = -2147483648
        _ = v14
    case 9223372036854775807:
        for fi15 := 0; fi15 < 8; fi15++ {
            _ = len(v1)
            var v16 int16 = 374
            _ = fi15
            _ = v16
        }
    default:
        v17 := GenS3{F0: -128, F1: -152, F2: 706}
        _ = v17
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
}
