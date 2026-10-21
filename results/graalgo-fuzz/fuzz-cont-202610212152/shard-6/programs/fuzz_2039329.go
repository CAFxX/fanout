package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 int8 }

func main() {
    var v1 bool = true
    v2 := make(map[uint32]string)
    var v3 uint16 = 0
    var v4 int8 = 100
    var v5 int16 = 0
    var v6 string = "foo bar"
    v3 = (v3 / ((v3) | 1))
    v7 := GenS1{F0: 128, F1: -356}
    for fi8 := 0; fi8 < 4; fi8++ {
        if (96 <= 968) {
            switch (fi8 << ((0) & 63)) {
            case 315:
                fmt.Println(v5, v6, v1)
            case 127:
                v9 := v2[939]
                _ = v9
            case 2147483647:
                var v10 int64 = -389
                _ = v10
            }
            v6 = v6
        } else {
            var v11 int32 = 2147483647
            v12 := []int{fi8, (fi8 - (fi8 & (fi8 + 0)))}
            _ = v11
            _ = v12
        }
        var v13 uint16 = v3
        _ = fi8
        _ = v13
    }
    if false {
        var v14 uint64 = 42
        for fi15 := 0; fi15 < 13; fi15++ {
            v16 := GenS1{F0: 118, F1: (fi15 | 1)}
            _ = fi15
            _ = v16
        }
        if v1 {
            if false {
                v17 := GenS2{F0: v4}
                _ = v17
            }
            fmt.Println(v5)
            v18 := []int{-370}
            _ = v18
        } else {
            v19 := GenS1{F0: 127, F1: 4294967295}
            v2[2147483647] = "a"
            _ = v19
        }
        _ = v14
    }
    v20 := GenS1{F0: 2, F1: 988}
    v2[2147483647] = v6
    v4 = v4
    v6 = "hello"
    var v21 string = ("foo bar" + v6)
    v1 = (((1e10 * (1e10 - 1e10)) == ((((float64(1e10) + 1e10) * 0.0) + ((0.1 + 1e10) * 0.0)) * (0.1 + (1.5 / 0.1)))) || (3.14159 <= 3.14159))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
