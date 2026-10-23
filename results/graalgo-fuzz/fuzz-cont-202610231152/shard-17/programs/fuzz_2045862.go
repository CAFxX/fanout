package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 int }
type GenS2 struct { F0 int32; F1 int8 }
func genFn0(p0 int32, p1 uint, p2 uint64) int {
    switch ("foo bar" + ((("" + "a") + ("foo bar" + "")) + (("a" + "\n\t") + ""))) {
    case "":
        for fi7 := 0; fi7 < 4; fi7++ {
            p0 = (^(p0))
            _ = fi7
        }
    }
    for fi8 := 0; fi8 < 8; fi8++ {
        var v9 int = 173
        for fi10 := 0; fi10 < 7; fi10++ {
            if false {
                switch (fi10 / ((fi8) | 1)) {
                case 128:
                    var v11 string = ("a" + ("" + ((((("hello" + "") + "") + "foo bar") + ("a" + ("" + "\n\t"))) + "\n\t")))
                    _ = v11
                case 9223372036854775807:
                    var v12 int = -9223372036854775808
                    _ = v12
                default:
                    p1 = (^(p1))
                }
                var v13 float32 = 1e10
                fmt.Println(p2)
                _ = v13
            } else {
                var v14 bool = (!true)
                p0 = (p0 % ((p0) | 1))
                _ = v14
            }
            p2 = p2
            fmt.Println(p1)
            _ = fi10
        }
        var v15 uint = uint(p2)
        _ = fi8
        _ = v15
        _ = v9
    }
    if true {
        p2 = 2
        if (128 > -1) {
            p1 = 661
            p0 -= p0
        }
        fmt.Println(p1)
    } else {
        p0 -= 42
    }
    _ = p0
    _ = p1
    _ = p2
    return 287
}

func main() {
    var v1 uint8 = 1
    var v2 float32 = 1e10
    var v3 uint16 = 2
    v4 := make(map[uint32]int)
    var v5 int16 = 83
    var v6 int8 = 1
    var v16 int = 2147483647
    if true {
        v4[63] = v16
    } else {
        var v17 uint16 = v3
        var v18 uint16 = (v3 << ((42) & 63))
        v19 := v4[2147483647]
        _ = v17
        _ = v18
        _ = v19
    }
    if (genFn0(-614, 742, 2) >= v16) {
        v5 = (^(v5))
        v20 := v4[830]
        var v21 uint16 = (v3 & 798)
        _ = v20
        _ = v21
    }
    v22 := GenS1{F0: 818, F1: 1, F2: v16}
    _ = len(v4)
    if false {
        if (v1 == (^(v1))) {
            v23 := GenS2{F0: 250, F1: (v6 | (v6 | (-(v6))))}
            _ = v23
        }
        switch (v16 << ((2) & 63)) {
        case 1:
            v1 -= v1
        case 4294967295:
            for fi24 := 0; fi24 < 6; fi24++ {
                var v25 float64 = 0.1
                _ = fi24
                _ = v25
            }
        }
    }
    v5 -= v5
    var v26 uint32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v22
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
