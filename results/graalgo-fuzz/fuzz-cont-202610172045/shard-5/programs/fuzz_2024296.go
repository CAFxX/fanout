package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 int }
type GenS2 struct { F0 int64; F1 bool; F2 string }
type GenS3 struct { F0 uint64; F1 uint32; F2 int8 }

func main() {
    var v1 int8 = 127
    var v2 uint16 = 42
    var v3 int32 = -2147483648
    fmt.Println(v2)
    fmt.Println(v1, v3)
    var v4 string = ""
    v5 := GenS1{F0: (v2 ^ 255), F1: 91, F2: -2147483648}
    v6 := []int{}
    var v7 uint = 65535
    var v8 int16 = 624
    fmt.Println(v3, v1, v3)
    switch 650 {
    case 9223372036854775807:
        var v9 int32 = (v3 + (v3 % ((v3) | 1)))
        _ = v9
    default:
        var v10 uint16 = 32767
        _ = v10
    }
    switch 32768 {
    case 32768:
        fmt.Println(v8)
    case -57:
        var v11 uint32 = 1
        _ = v11
    }
    fmt.Println(v7, v1)
    var v12 float32 = float32((1e10 + 1.5))
    for fi13 := 0; fi13 < 12; fi13++ {
        v6 = append(v6, (^(fi13)))
        _ = fi13
    }
    v14 := GenS2{F0: 1, F1: true, F2: v4}
    for fi15 := 0; fi15 < 13; fi15++ {
        var v16 float32 = v12
        _ = fi15
        _ = v16
    }
    var v17 float64 = ((0.1 * 0.1) / 1.5)
    switch "" {
    case "a":
        for fi18 := 0; fi18 < 17; fi18++ {
            if (!(true && true)) {
                switch (fi18 >> ((uint(v2)) & 63)) {
                case -1:
                    switch (fi18 * fi18) {
                    case 32767:
                        v19 := GenS3{F0: 4294967295, F1: 643, F2: 1}
                        _ = v19
                    case 162:
                        var v20 bool = true
                        _ = v20
                    case 42:
                        var v21 int8 = (v1 - (v1 | v1))
                        _ = v21
                    }
                case 127:
                    var v22 int16 = v8
                    _ = v22
                default:
                    var v23 uint = v7
                    _ = v23
                }
                var v24 uint8 = 127
                fmt.Println(v8, v3, v8)
                _ = v24
            } else {
                if len(v6) > 0 {
                    v6[0] = fi18
                }
                fmt.Println(v1, v1)
                _ = len(v6)
            }
            _ = fi18
        }
    default:
        switch 128 {
        case -2147483648:
            var v25 float32 = 0.1
            _ = v25
        case -486:
            for fi26 := 0; fi26 < 5; fi26++ {
                var v27 float64 = ((1.5 - v17) / (v17 + v17))
                _ = fi26
                _ = v27
            }
        case 757:
            fmt.Println(v7)
        }
    }
    for fi28 := 0; fi28 < 9; fi28++ {
        v6 = append(v6, (fi28 * fi28))
        if (2 == 0) {
            if len(v6) > 0 {
                v6[0] = (fi28 << (((v7 << (((v7 & 939)) & 63))) & 63))
            }
        } else {
            if true {
                fmt.Println(v1, v8, v12)
                v4 = v4
            } else {
                var v29 int16 = 1000
                _ = v29
            }
        }
        for fi30 := 0; fi30 < 4; fi30++ {
            var v31 float32 = float32(v17)
            _ = fi30
            _ = v31
        }
        _ = fi28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
