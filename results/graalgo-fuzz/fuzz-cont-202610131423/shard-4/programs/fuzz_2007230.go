package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 uint16 }
type GenS2 struct { F0 uint8; F1 uint16; F2 int64 }
func genFn0() int16 {
    for fi6 := 0; fi6 < 17; fi6++ {
        switch -643 {
        case -265:
            v7 := GenS1{F0: 2, F1: ("foo bar" + "a"), F2: 403}
            _ = v7
        default:
            if ((3.14159 + 0.0) != float32(3.14159)) {
                v8 := GenS2{F0: 146, F1: 255, F2: -1}
                _ = v8
            }
        }
        _ = fi6
    }
    if false {
        var v9 int8 = -128
        v9 += v9
        var v10 uint = 407
        _ = v10
        _ = v9
    }
    return 42
}

func main() {
    var v1 int8 = 84
    var v2 int = 9223372036854775807
    var v3 int32 = 287
    v4 := []string{"foo bar"}
    v5 := make(map[int]uint32)
    var v11 float32 = -2.25
    v12 := GenS2{F0: 49, F1: 834, F2: 223}
    v13 := GenS2{F0: 0, F1: 1, F2: 1}
    delete(v5, v2)
    for fi14 := 0; fi14 < 14; fi14++ {
        v3 -= (v3 * -2147483648)
        v11 -= (float32(1e10) / (float32(1e10) + (v11 / ((v11 - (v11 - v11)) / v11))))
        for fi15 := 0; fi15 < 6; fi15++ {
            v1 += v1
            _ = fi15
        }
        _ = fi14
    }
    var v16 uint16 = 1
    var v17 int64 = 1
    v11 = (v11 / (v11 * (v11 + (v11 * (float32(1.5) / 1.0)))))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
