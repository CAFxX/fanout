package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }

func main() {
    var v1 bool = (!(!true))
    var v2 uint8 = 61
    var v3 uint8 = (v2 << ((844) & 63))
    v4 := []int{}
    v1 = (!v1)
    var v5 uint16 = 426
    for fi6 := 0; fi6 < 17; fi6++ {
        if len(v4) > 0 {
            v4[0] = fi6
        }
        _ = fi6
    }
    fmt.Println(v1)
    var v7 uint32 = 1
    var v8 float64 = 0.0
    v9 := GenS1{F0: 348, F1: 127}
    v10 := GenS1{F0: (v5 & 992), F1: 42}
    for fi11 := 0; fi11 < 15; fi11++ {
        var v12 int = fi11
        for fi13 := 0; fi13 < 8; fi13++ {
            v4 = append(v4, 2147483647)
            v8 += (v8 * v8)
            v2 -= v3
            _ = fi13
        }
        v5 += v5
        _ = fi11
        _ = v12
    }
    if v1 {
        var v14 uint = 65535
        v1 = (float32(float64((0.0 + (1.5 + 1.5)))) <= -2.25)
        if true {
            v14 = (v14 & (v14 >> ((uint(v7)) & 63)))
            fmt.Println(v5, v14)
            for fi15 := 0; fi15 < 16; fi15++ {
                v16 := GenS1{F0: 1, F1: 127}
                if len(v4) > 0 {
                    v4[0] = fi15
                }
                _ = fi15
                _ = v16
            }
        }
        _ = v14
    } else {
        fmt.Println(v1, v2, v7)
    }
    var v17 int = 1
    if (!(("foo bar" + "") > "a")) {
        switch v17 {
        case -891:
            fmt.Println(v3, v3)
        case 1:
            var v18 int16 = 592
            _ = v18
        case 32768:
            fmt.Println(v5)
        }
    } else {
        switch (v17 | v17) {
        case 127:
            v17 = v17
        case 9223372036854775807:
            var v19 int16 = 1000
            _ = v19
        }
        var v20 uint8 = 71
        v3 += (v2 ^ (v2 * 255))
        _ = v20
    }
    v21 := GenS1{F0: v5, F1: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
