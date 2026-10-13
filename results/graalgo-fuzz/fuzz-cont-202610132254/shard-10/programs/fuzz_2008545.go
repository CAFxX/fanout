package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint16 }
type GenS2 struct { F0 int16; F1 int; F2 uint64 }

func main() {
    var v1 int8 = 1
    var v2 string = ((((("foo bar" + "\n\t") + (("foo bar" + "hello") + "hello")) + "foo bar") + "\n\t") + "")
    var v3 float64 = (((3.14159 * 1e10) * (0.0 * float64(float32(0.1)))) + float64(float32(float64(3.14159))))
    var v4 uint8 = 127
    var v5 int = 32767
    fmt.Println(v4, v5, v4)
    v6 := GenS1{F0: v4, F1: 0, F2: 32767}
    var v7 uint64 = 2
    var v8 float32 = float32(v3)
    if ((^(v7)) >= (^(v7))) {
        fmt.Println(v4, v7, v1)
        v1 -= (v1 & (v1 * (v1 + (v1 / (((v1 << ((2) & 63))) | 1)))))
    }
    v9 := []int{}
    v1 = (v1 % ((v1) | 1))
    fmt.Println(v7, v1)
    if len(v9) > 0 {
        v9[0] = v5
    }
    if ((v1 % ((v1) | 1)) <= 42) {
        if ((v4 | v4) >= (v4 - (v4 << ((4294967295) & 63)))) {
            var v10 int8 = v1
            v5 -= (v5 >> ((42) & 63))
            _ = v10
        }
        var v11 int16 = 32767
        _ = v11
    }
    var v12 string = v2
    v13 := GenS2{F0: 0, F1: (-(v5)), F2: 18446744073709551615}
    var v14 bool = ((228 > v4) || (false && true))
    v7 -= v7
    if (0 != (v5 << ((255) & 63))) {
        switch v5 {
        case -2147483648:
            for fi15 := 0; fi15 < 8; fi15++ {
                if v14 {
                    var v16 int16 = 1000
                    _ = v16
                }
                v1 -= v1
                _ = fi15
            }
        case 715:
            v1 += -62
        }
        var v17 float32 = v8
        _ = v17
    } else {
        for fi18 := 0; fi18 < 17; fi18++ {
            for fi19 := 0; fi19 < 7; fi19++ {
                var v20 uint16 = 0
                var v21 uint8 = 127
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
                var v22 int64 = 0
                _ = fi19
                _ = v20
                _ = v21
                _ = v22
            }
            v23 := GenS2{F0: 1, F1: 42, F2: 18446744073709551615}
            _ = fi18
            _ = v23
        }
        var v24 uint16 = 2
        v4 += v4
        _ = v24
    }
    v8 += (v8 * (v8 - 1e10))
    _ = len(v9)
    if false {
        var v25 int16 = 42
        if ((v4 / (((v4 + (^(v4)))) | 1)) < (v4 / (((v4 * (v4 / ((v4) | 1)))) | 1))) {
            var v26 uint8 = (v4 ^ v4)
            _ = v26
        }
        _ = v25
    } else {
        var v27 uint32 = 1
        var v28 bool = (!(!((v7 & v7) <= (v7 * 2))))
        var v29 bool = false
        _ = v27
        _ = v28
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
