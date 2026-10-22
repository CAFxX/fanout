package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 int }
type GenS2 struct { F0 int64; F1 int64; F2 int64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint64; F1 uint16; F2 int8 }
type GenS5 struct { F0 string; F1 int16; F2 int16 }

func main() {
    v1 := make(map[int]int)
    var v2 uint8 = 0
    var v3 uint8 = v2
    var v4 int32 = -2147483648
    var v5 int = 111
    var v6 int16 = 1000
    v7 := GenS1{F0: -71, F1: 4294967295, F2: -831}
    for fi8 := 0; fi8 < 4; fi8++ {
        if (2147483647 < 2) {
            if (false || false) {
                v9 := GenS2{F0: 9223372036854775807, F1: 9223372036854775807, F2: -1}
                _ = v9
            } else {
                switch "foo bar" {
                case "a":
                    v10 := GenS3{F0: v6}
                    _ = v10
                }
            }
            var v11 int64 = 0
            _ = v11
        } else {
            v12 := GenS3{F0: v6}
            v13 := GenS4{F0: 2, F1: 65535, F2: -1}
            _ = v12
            _ = v13
        }
        _ = fi8
    }
    for fi14 := 0; fi14 < 13; fi14++ {
        var v15 uint32 = 62
        _ = fi14
        _ = v15
    }
    v5 -= v5
    var v16 uint16 = 1
    v17 := GenS5{F0: ("foo bar" + "hello"), F1: -32768, F2: (-(v6))}
    v6 = 42
    for fi18 := 0; fi18 < 4; fi18++ {
        switch (v5 % (((fi18 << ((936) & 63))) | 1)) {
        case 0:
            var v19 uint64 = 42
            _ = v19
        case -318:
            fmt.Println(v4)
        case 392:
            v20 := []int{1, v5, (fi18 - (v5 / ((-9223372036854775808) | 1)))}
            _ = v20
        default:
            var v21 float64 = float64(1e10)
            _ = v21
        }
        var v22 int = v5
        _ = fi18
        _ = v22
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
