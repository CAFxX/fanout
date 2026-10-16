package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 bool }
type GenS2 struct { F0 int16; F1 int8 }
func genFn0(p0 int8, p1 int8) uint {
    v5 := GenS1{F0: 42, F1: 310, F2: true}
    p1 -= (p0 % ((0) | 1))
    v6 := GenS2{F0: 1, F1: p1}
    var v7 uint32 = 4294967295
    v7 -= (v7 << ((4294967295) & 63))
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 252
}

func main() {
    var v1 string = ("hello" + ("\n\t" + "a"))
    var v2 bool = false
    var v3 uint16 = 42
    var v4 uint32 = 0
    for fi8 := 0; fi8 < 12; fi8++ {
        var v9 bool = true
        if (!(9223372036854775807 > 966)) {
            var v10 uint8 = 42
            _ = v10
        }
        _ = fi8
        _ = v9
    }
    if (0.0 == (1.5 + (3.14159 - 0.1))) {
        fmt.Println(v3, v1)
        var v11 uint64 = 241
        _ = v11
    }
    v3 += (v3 | (v3 % (((v3 + (v3 - v3))) | 1)))
    for fi12 := 0; fi12 < 8; fi12++ {
        for fi13 := 0; fi13 < 17; fi13++ {
            var v14 uint64 = 42
            var v15 string = (("hello" + (((v1 + (v1 + v1)) + v1) + "hello")) + "hello")
            _ = fi13
            _ = v14
            _ = v15
        }
        switch (fi12 + (fi12 << ((genFn0(42, 100)) & 63))) {
        case -984:
            fmt.Println(fi12, v3, fi12)
        case 568:
            v16 := GenS2{F0: 42, F1: -128}
            _ = v16
        }
        for fi17 := 0; fi17 < 3; fi17++ {
            if (fi17 <= (-(fi17))) {
                var v18 int8 = -105
                _ = v18
            }
            _ = fi17
        }
        _ = fi12
    }
    v19 := GenS1{F0: 4294967295, F1: -808, F2: true}
    for fi20 := 0; fi20 < 16; fi20++ {
        for fi21 := 0; fi21 < 17; fi21++ {
            v1 = "hello"
            _ = fi21
        }
        v22 := GenS1{F0: 866, F1: 0, F2: ((-9223372036854775808 == (fi20 & 0)) || true)}
        _ = fi20
        _ = v22
    }
    v4 = v4
    var v23 int8 = 1
    fmt.Println(v1, v2, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
