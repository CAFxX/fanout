package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8; F1 uint64 }
type GenS4 struct { F0 uint64; F1 int16 }
type GenS5 struct { F0 int; F1 uint32 }
type GenS6 struct { F0 uint32 }
func genFn0(p0 uint32) uint8 {
    if false {
        fmt.Println(p0)
    } else {
        fmt.Println(p0)
        p0 = p0
    }
    for fi7 := 0; fi7 < 12; fi7++ {
        var v8 uint32 = p0
        _ = fi7
        _ = v8
    }
    _ = p0
    return 42
}

func main() {
    v1 := []uint32{2147483648}
    v2 := make(map[int]bool)
    var v3 int = -719
    var v4 uint32 = 1
    var v5 uint = 498
    var v6 int16 = 0
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (100 == -128) {
        switch v3 {
        case 32768:
            var v9 uint16 = 32767
            _ = v9
        case 127:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        var v10 uint = v5
        for fi11 := 0; fi11 < 13; fi11++ {
            v12 := GenS1{F0: 108, F1: -2147483648}
            v13 := GenS2{F0: (true && (false && (!(-9223372036854775808 < 1))))}
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = v10
    }
    v6 = -26
    for fi14 := 0; fi14 < 11; fi14++ {
        if false {
            v15 := GenS3{F0: 1, F1: 1}
            v16 := GenS4{F0: 358, F1: v6}
            v4 = 1
            _ = v15
            _ = v16
        }
        v2[(fi14 % (((v3 >> ((0) & 63))) | 1))] = true
        _ = fi14
    }
    v5 = v5
    v17 := GenS5{F0: (v3 | v3), F1: v4}
    for fi18 := 0; fi18 < 8; fi18++ {
        switch ("a" + ((("" + (("foo bar" + "\n\t") + "hello")) + "") + "foo bar")) {
        case "a":
            fmt.Println(v6, v3)
        case "\n\t":
            _ = len(v1)
        case "foo bar":
            var v19 int8 = -78
            _ = v19
        }
        _ = fi18
    }
    if len(v1) > 0 {
        v1[0] = v4
    }
    switch "hello" {
    case "foo bar":
        if (-9223372036854775808 < -9223372036854775808) {
            for fi20 := 0; fi20 < 17; fi20++ {
                var v21 float64 = float64(3.14159)
                _ = fi20
                _ = v21
            }
            var v22 uint16 = 0
            _ = v22
        }
    case "\n\t":
        for fi23 := 0; fi23 < 20; fi23++ {
            var v24 uint64 = 741
            var v25 float64 = 0.0
            var v26 uint64 = (v24 >> (((^(v5))) & 63))
            _ = fi23
            _ = v24
            _ = v25
            _ = v26
        }
    default:
        v27 := GenS3{F0: -37, F1: 536}
        _ = v27
    }
    v4 = (v4 + (^(v4)))
    v28 := GenS6{F0: (^(v4))}
    v29 := v2[v3]
    var v30 float64 = (-2.25 / -2.25)
    v30 = v30
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
