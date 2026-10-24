package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 int8 }
type GenS2 struct { F0 int16; F1 uint32; F2 uint32 }

func main() {
    var v1 int64 = -9223372036854775808
    v2 := []int{}
    var v3 uint64 = 0
    var v4 float64 = -2.25
    v5 := make(map[uint32]uint32)
    var v6 uint = 65535
    if ((127 != 2147483647) || true) {
        v2 = append(v2, -183)
    } else {
        var v7 int = -351
        v8 := GenS1{F0: 42, F1: (-(v1)), F2: 30}
        fmt.Println(v1, v6)
        _ = v7
        _ = v8
    }
    var v9 uint32 = 2147483648
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch 543 {
    case -137:
        if false {
            fmt.Println(v9, v4)
        } else {
            v6 += 1
            for fi10 := 0; fi10 < 20; fi10++ {
                var v11 uint8 = 2
                if true {
                    v3 = (v3 + (v3 + v3))
                } else {
                    v11 += 2
                    var v12 int32 = -217
                    _ = v12
                }
                v4 -= v4
                _ = fi10
                _ = v11
            }
            for fi13 := 0; fi13 < 6; fi13++ {
                for fi14 := 0; fi14 < 9; fi14++ {
                    if len(v2) > 0 {
                        v2[0] = (fi14 * fi14)
                    }
                    _ = fi14
                }
                _ = fi13
            }
        }
    case 127:
        var v15 int = int(v1)
        _ = v15
    default:
        delete(v5, 42)
    }
    for fi16 := 0; fi16 < 2; fi16++ {
        if len(v2) > 0 {
            v2[0] = 0
        }
        var v17 float32 = 3.14159
        _ = fi16
        _ = v17
    }
    for fi18 := 0; fi18 < 17; fi18++ {
        v19 := GenS1{F0: 2, F1: int64(fi18), F2: 100}
        v3 = (v3 * (v3 / (((^(v3))) | 1)))
        _ = fi18
        _ = v19
    }
    var v20 float32 = 0.1
    for fi21 := 0; fi21 < 1; fi21++ {
        v22 := GenS2{F0: 1000, F1: (^(v9)), F2: (v9 * (v9 | 42))}
        var v23 int8 = 100
        fmt.Println(v6)
        _ = fi21
        _ = v22
        _ = v23
    }
    var v24 bool = (false || false)
    fmt.Println(v24, v6)
    v25 := GenS2{F0: -599, F1: (^(v9)), F2: v9}
    v4 = (0.1 + 1e10)
    var v26 int64 = (v1 / ((-1) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
