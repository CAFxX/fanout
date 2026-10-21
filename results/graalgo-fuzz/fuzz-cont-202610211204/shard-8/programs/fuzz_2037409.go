package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 uint }

func main() {
    var v1 uint8 = 2
    var v2 uint8 = (v1 ^ v1)
    var v3 int = 32767
    var v4 int16 = -1
    v5 := []int{v3}
    v5 = append(v5, v3)
    v4 = 42
    for fi6 := 0; fi6 < 15; fi6++ {
        var v7 uint8 = v2
        for fi8 := 0; fi8 < 15; fi8++ {
            _ = len(v5)
            v4 = (v4 + v4)
            _ = len(v5)
            _ = fi8
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi6
        _ = v7
    }
    var v9 int16 = v4
    v1 += v2
    if true {
        fmt.Println(v2)
    }
    fmt.Println(v3)
    if len(v5) > 0 {
        v5[0] = (-(v3))
    }
    if len(v5) > 0 {
        v5[0] = 0
    }
    _ = len(v5)
    if ((v3 + (-(v3))) != (v3 & v3)) {
        fmt.Println(v4)
        v9 = v4
        _ = len(v5)
    } else {
        fmt.Println(v2, v4, v9)
        v1 -= (v2 % (((v2 & (v1 ^ (v1 + (v1 | (v2 >> ((255) & 63))))))) | 1))
    }
    if ((100 <= 0) && (18446744073709551615 >= 4294967295)) {
        fmt.Println(v1, v9, v4)
        if false {
            _ = len(v5)
            var v10 uint64 = 1
            _ = v10
        }
    }
    if len(v5) > 0 {
        v5[0] = 42
    }
    if (v9 <= (-(v4))) {
        for fi11 := 0; fi11 < 18; fi11++ {
            for fi12 := 0; fi12 < 15; fi12++ {
                var v13 uint8 = (v2 + (^(v1)))
                switch fi12 {
                case -144:
                    var v14 bool = (false || (((1 == 0) || true) || true))
                    _ = v14
                case 127:
                    var v15 int8 = 56
                    _ = v15
                }
                _ = fi12
                _ = v13
            }
            _ = fi11
        }
        v4 = (v4 >> ((uint(v1)) & 63))
        v3 = 0
    } else {
        var v16 int = (-(v3))
        fmt.Println(v2)
        _ = v16
    }
    fmt.Println(v1, v2)
    v9 = v9
    for fi17 := 0; fi17 < 7; fi17++ {
        for fi18 := 0; fi18 < 17; fi18++ {
            v19 := GenS1{F0: 1, F1: (^(v1)), F2: 1}
            _ = fi18
            _ = v19
        }
        _ = fi17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
