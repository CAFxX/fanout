package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 string }
type GenS2 struct { F0 int }

func main() {
    var v1 uint64 = 1
    var v2 bool = (false || false)
    v3 := make(map[string]int)
    v1 = 0
    var v4 float64 = -2.25
    var v5 int32 = -1
    if v2 {
        v5 += 42
    } else {
        v1 += v1
        if ((v1 | (v1 + (v1 ^ v1))) >= (v1 + (v1 % ((v1) | 1)))) {
            v5 -= (v5 / ((-2147483648) | 1))
        }
    }
    var v6 uint = 255
    switch -901 {
    case -2147483648:
        v1 -= (v1 << (((v6 >> (((v6 / (((v6 << (((v6 % ((v6) | 1))) & 63))) | 1))) & 63))) & 63))
    case 257:
        switch 0 {
        case -9223372036854775808:
            v7 := GenS1{F0: (v1 | uint64(v6)), F1: -9223372036854775808, F2: "\n\t"}
            _ = v7
        case 32767:
            var v8 int8 = 0
            _ = v8
        case 127:
            v9 := GenS2{F0: -892}
            _ = v9
        }
    case 128:
        v10 := GenS1{F0: v1, F1: 1, F2: "a"}
        _ = v10
    }
    v11 := GenS1{F0: v1, F1: 42, F2: "a"}
    var v12 int32 = (-(v5))
    fmt.Println(v1, v12)
    for fi13 := 0; fi13 < 2; fi13++ {
        var v14 uint8 = 42
        for fi15 := 0; fi15 < 9; fi15++ {
            var v16 int8 = 0
            switch (fi15 % ((fi15) | 1)) {
            case -1:
                v5 += (v5 - (v12 / (((v12 << ((v6) & 63))) | 1)))
            case 342:
                fmt.Println(fi15)
            case 312:
                if (104 != v16) {
                    v17 := v3[""]
                    v5 -= (v12 & (v5 * 431))
                    v6 = v6
                    _ = v17
                } else {
                    var v18 float64 = 3.14159
                    _ = v18
                }
            default:
                fmt.Println(v12)
            }
            _ = fi15
            _ = v16
        }
        _ = fi13
        _ = v14
    }
    v1 = v1
    v12 = (v5 >> ((v6) & 63))
    var v19 int = -394
    var v20 uint32 = 740
    v12 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
