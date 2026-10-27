package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 bool }

func main() {
    var v1 float64 = float64(0.0)
    var v2 int16 = -36
    v3 := []int{428, 32768}
    var v4 float32 = 1e10
    var v5 uint32 = 1
    var v6 int64 = 0
    if (v4 != (0.1 - 3.14159)) {
        v2 = 32767
        v6 -= (^(v6))
        v7 := GenS1{F0: 0, F1: 0}
        _ = v7
    }
    v1 = v1
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi8 := 0; fi8 < 20; fi8++ {
        v9 := GenS1{F0: 291, F1: int32(v6)}
        _ = fi8
        _ = v9
    }
    var v10 int64 = (v6 * v6)
    for fi11 := 0; fi11 < 12; fi11++ {
        var v12 uint = 758
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 14; fi13++ {
        v3 = append(v3, -480)
        _ = fi13
    }
    var v14 int = -171
    fmt.Println(v6, v1)
    v1 -= v1
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi15 := 0; fi15 < 17; fi15++ {
        fmt.Println(fi15, v2, v2)
        _ = fi15
    }
    v3 = append(v3, v14)
    if (v2 > (v2 | 32767)) {
        if len(v3) > 0 {
            v3[0] = v14
        }
        if len(v3) > 0 {
            v3[0] = v14
        }
    } else {
        if (-9223372036854775808 <= 9223372036854775807) {
            v3 = append(v3, v14)
            v16 := GenS1{F0: 1, F1: 881}
            if true {
                fmt.Println(v14)
                switch v14 {
                case 128:
                    if (v6 > (v6 * (^(v6)))) {
                        fmt.Println(v6)
                        v17 := GenS2{F0: ((!(false && ((0 == 65535) && false))) || true)}
                        var v18 int32 = -1
                        _ = v17
                        _ = v18
                    }
                default:
                    var v19 int32 = 42
                    _ = v19
                }
                _ = len(v3)
            } else {
                v20 := GenS1{F0: 1, F1: 42}
                _ = v20
            }
            _ = v16
        }
        var v21 int8 = -1
        fmt.Println(v1)
        _ = v21
    }
    var v22 int = (v14 & v14)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
