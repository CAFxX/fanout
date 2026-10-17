package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 string = ""
    var v2 uint16 = 141
    var v3 uint = 540
    var v4 bool = true
    v5 := []uint32{238, 334, 2, 2}
    var v6 string = "a"
    var v7 int16 = 545
    var v8 int16 = v7
    var v9 uint = (v3 + (v3 + (v3 / ((v3) | 1))))
    var v10 int32 = 42
    switch ("" + v1) {
    case "\n\t":
        if len(v5) > 0 {
            v5[0] = 0
        }
    case "foo bar":
        var v11 int8 = -128
        _ = v11
    case "hello":
        switch "foo bar" {
        case "\n\t":
            var v12 uint16 = (v2 ^ (v2 * (^(v2))))
            _ = v12
        case "hello":
            v13 := GenS1{F0: int64(v10), F1: false}
            _ = v13
        default:
            for fi14 := 0; fi14 < 8; fi14++ {
                v3 -= 457
                v5 = append(v5, uint32(v9))
                _ = fi14
            }
        }
    }
    var v15 int8 = -128
    v4 = false
    var v16 int64 = 474
    switch "" {
    case "\n\t":
        for fi17 := 0; fi17 < 17; fi17++ {
            v18 := GenS2{F0: 0}
            if v4 {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                if (!v4) {
                    var v19 int32 = (v10 & (v10 * (^(v10))))
                    _ = v19
                }
                v10 += -1
            } else {
                fmt.Println(fi17, v15)
                var v20 uint16 = (v2 / (((^(v2))) | 1))
                _ = v20
            }
            _ = fi17
            _ = v18
        }
    case "foo bar":
        var v21 uint8 = 0
        _ = v21
    }
    var v22 uint64 = 348
    v23 := GenS2{F0: v2}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
