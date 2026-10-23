package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int32; F1 int32 }
func genFn0(p0 int8, p1 int16, p2 string) uint {
    v6 := GenS1{F0: -451, F1: 871}
    v7 := []int{259, -9223372036854775808}
    v8 := GenS1{F0: 1, F1: 191}
    if (18446744073709551615 == 4294967295) {
        if (!false) {
            var v9 string = ("\n\t" + ("foo bar" + p2))
            _ = v9
        } else {
            fmt.Println(p2, p2)
            p0 += p0
            if false {
                fmt.Println(p0)
                v10 := GenS2{F0: 542}
                p0 += (p0 | p0)
                _ = v10
            }
        }
        var v11 uint64 = 1
        v12 := GenS1{F0: 546, F1: 32767}
        _ = v11
        _ = v12
    }
    var v13 float32 = ((0.0 / -2.25) * float32((-2.25 * (0.0 + 1.5))))
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v6
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 int16 = 1
    v2 := make(map[uint32]uint32)
    var v3 uint8 = 32
    var v4 int = 9223372036854775807
    v5 := make(map[uint32]bool)
    if (!true) {
        switch 32767 {
        case 4294967295:
            v3 = (^(v3))
        case -9223372036854775808:
            for fi14 := 0; fi14 < 17; fi14++ {
                v15 := []int{}
                _ = fi14
                _ = v15
            }
        }
        v3 = (^(v3))
    }
    v4 = (v4 | v4)
    switch "" {
    case "foo bar":
        v1 = -32768
    default:
        switch "" {
        case "a":
            v16 := GenS1{F0: -1, F1: 0}
            _ = v16
        default:
            v1 += v1
        }
    }
    v17 := []int{}
    v18 := GenS3{F0: -2147483648, F1: 0}
    var v19 uint64 = uint64(v3)
    var v20 uint32 = 42
    var v21 uint8 = (^(v3))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
