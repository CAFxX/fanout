package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 int8; F1 int16; F2 string }

func main() {
    var v1 uint8 = 253
    v2 := []uint32{4294967295, 2147483648}
    var v3 int8 = 0
    var v4 int16 = 1000
    v5 := make(map[int]int)
    var v6 uint8 = (v1 - 128)
    v5[-1] = 9223372036854775807
    var v7 int8 = (v3 * v3)
    switch "foo bar" {
    case "a":
        for fi8 := 0; fi8 < 19; fi8++ {
            var v9 float64 = -2.25
            for fi10 := 0; fi10 < 2; fi10++ {
                var v11 float64 = v9
                if (false || false) {
                    v12 := GenS1{F0: (v3 << ((4294967295) & 63)), F1: 0}
                    v1 += (^(v1))
                    var v13 bool = (true && (!false))
                    if len(v2) > 0 {
                        v2[0] = 423
                    }
                    _ = v12
                    _ = v13
                }
                v5[42] = fi8
                _ = fi10
                _ = v11
            }
            switch "hello" {
            case "\n\t":
                v2 = append(v2, 65535)
            }
            _ = fi8
            _ = v9
        }
    case "hello":
        var v14 int8 = v3
        _ = v14
    }
    var v15 uint16 = 1
    for fi16 := 0; fi16 < 10; fi16++ {
        if true {
            v17 := GenS2{F0: (v3 & (-(v7))), F1: v4, F2: "\n\t"}
            v3 = (^(v7))
            _ = v17
        }
        _ = fi16
    }
    var v18 float64 = (0.0 / 1.5)
    v4 = -1
    v18 += float64(float32(v18))
    for fi19 := 0; fi19 < 2; fi19++ {
        v18 = (v18 + (0.1 - v18))
        v20 := GenS1{F0: v7, F1: -878}
        var v21 float32 = 1.5
        _ = fi19
        _ = v20
        _ = v21
    }
    if (0.0 <= -2.25) {
        var v22 uint8 = (v1 * 127)
        v22 = (v22 & v22)
        _ = v22
    } else {
        if (false && (-295 <= 42)) {
            var v23 uint64 = 18446744073709551615
            v24 := v5[-242]
            _ = v23
            _ = v24
        }
        v18 -= (float64(0.0) * (v18 / v18))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
