package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 int8 }
type GenS2 struct { F0 uint16; F1 uint8; F2 int8 }
type GenS3 struct { F0 uint64; F1 uint8; F2 int16 }
type GenS4 struct { F0 int }

func main() {
    var v1 uint16 = 91
    var v2 int16 = 965
    v3 := make(map[string]bool)
    v1 -= v1
    v1 -= v1
    for fi4 := 0; fi4 < 13; fi4++ {
        var v5 uint64 = 2
        v6 := GenS1{F0: 128, F1: (v1 + v1), F2: 0}
        _ = fi4
        _ = v5
        _ = v6
    }
    for fi7 := 0; fi7 < 4; fi7++ {
        for fi8 := 0; fi8 < 14; fi8++ {
            if (false || (4294967295 == 65535)) {
                v2 -= (v2 >> ((585) & 63))
                v9 := GenS1{F0: 0, F1: (^(v1)), F2: -71}
                v1 -= 255
                _ = v9
            } else {
                v10 := GenS1{F0: 2, F1: (v1 / ((v1) | 1)), F2: 100}
                v11 := GenS2{F0: 255, F1: 2, F2: 0}
                fmt.Println(fi8, fi7, v1)
                _ = v10
                _ = v11
            }
            _ = fi8
        }
        if false {
            v1 += v1
        } else {
            for fi12 := 0; fi12 < 10; fi12++ {
                v13 := []int{1, 9223372036854775807, (^(fi12)), -580}
                _ = fi12
                _ = v13
            }
        }
        _ = fi7
    }
    v14 := GenS3{F0: 2, F1: 0, F2: v2}
    var v15 uint64 = 42
    fmt.Println(v1, v2, v15)
    fmt.Println(v1, v2, v2)
    v15 = (v15 >> ((65535) & 63))
    for fi16 := 0; fi16 < 17; fi16++ {
        var v17 int16 = 0
        for fi18 := 0; fi18 < 9; fi18++ {
            fmt.Println(fi18)
            _ = fi18
        }
        _ = fi16
        _ = v17
    }
    switch 2147483647 {
    case -276:
        if (!(false && true)) {
            switch "\n\t" {
            case "\n\t":
                var v19 int32 = -564
                _ = v19
            default:
                v2 += -958
            }
            var v20 int16 = v2
            _ = v20
        } else {
            v21 := v3[("foo bar" + ("" + "a"))]
            var v22 int16 = (v2 + (v2 * 460))
            _ = v21
            _ = v22
        }
    case 1:
        v3["hello"] = (42 >= 42)
    default:
        var v23 float32 = -2.25
        _ = v23
    }
    var v24 float64 = 0.0
    v25 := GenS4{F0: 127}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v24
    _ = v25
    _ = v3
}
