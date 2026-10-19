package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 int16 = 1
    var v2 uint8 = 127
    var v3 uint32 = 296
    v4 := []uint32{uint32(v2), (v3 % (((^(v3))) | 1)), 827, v3}
    v5 := []int{}
    var v6 float32 = float32((0.0 / ((0.1 + (((((0.0 / 3.14159) / 1e10) * (3.14159 * 0.1)) - 0.0) / -2.25)) / -2.25)))
    if true {
        fmt.Println(v6)
    }
    v3 += v3
    v3 += 65535
    v7 := GenS1{F0: 339}
    var v8 float64 = 3.14159
    v3 = (v3 / (((^(v3))) | 1))
    v8 += v8
    switch -2147483648 {
    case -2147483648:
        v9 := GenS1{F0: 255}
        _ = v9
    case 42:
        v8 = 3.14159
    }
    if len(v4) > 0 {
        v4[0] = 704
    }
    if (9 < 4294967295) {
        if (-9223372036854775808 <= -140) {
            v10 := GenS1{F0: 1}
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = v10
        }
        fmt.Println(v1, v6)
    }
    var v11 uint16 = 705
    var v12 float32 = v6
    v13 := GenS1{F0: uint(v3)}
    switch -630 {
    case 0:
        v14 := GenS1{F0: 1}
        _ = v14
    case -322:
        for fi15 := 0; fi15 < 5; fi15++ {
            fmt.Println(fi15)
            _ = fi15
        }
    case 2147483647:
        fmt.Println(v2, v1)
    default:
        for fi16 := 0; fi16 < 12; fi16++ {
            var v17 uint = 4294967295
            v11 += v11
            _ = fi16
            _ = v17
        }
    }
    v12 += v6
    fmt.Println(v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
