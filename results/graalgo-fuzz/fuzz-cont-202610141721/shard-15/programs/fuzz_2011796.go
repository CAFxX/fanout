package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64; F1 int32 }
func genFn0(p0 int8, p1 uint8, p2 int64) uint {
    v6 := GenS1{F0: (p0 + (p0 << ((18446744073709551615) & 63)))}
    if true {
        var v7 float64 = -2.25
        _ = v7
    }
    if true {
        var v8 uint32 = 2147483648
        p0 += (p0 + 127)
        var v9 int64 = p2
        _ = v8
        _ = v9
    } else {
        v10 := GenS2{F0: p2, F1: -2147483648}
        _ = v10
    }
    for fi11 := 0; fi11 < 4; fi11++ {
        p0 -= (p0 | 124)
        fmt.Println(p1, p0)
        _ = fi11
    }
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 4294967295
}
func genFn1(p0 int16) uint64 {
    fmt.Println(p0)
    v12 := []int{}
    for fi13 := 0; fi13 < 2; fi13++ {
        p0 -= (^(p0))
        _ = fi13
    }
    switch (("hello" + ("\n\t" + "")) + (("\n\t" + ("a" + "")) + ("foo bar" + "foo bar"))) {
    case "foo bar":
        p0 = (-(p0))
    case "\n\t":
        if true {
            var v14 uint64 = 143
            for fi15 := 0; fi15 < 14; fi15++ {
                fmt.Println(fi15)
                _ = fi15
            }
            _ = v14
        }
    default:
        if true {
            for fi16 := 0; fi16 < 6; fi16++ {
                switch fi16 {
                case -1:
                    p0 = (p0 << ((293) & 63))
                case 794:
                    var v17 uint16 = 1
                    _ = v17
                }
                _ = fi16
            }
            p0 = (-(p0))
            if (255 != 255) {
                var v18 float64 = 0.1
                _ = len(v12)
                var v19 bool = false
                _ = v18
                _ = v19
            }
        }
    }
    var v20 uint32 = 693
    _ = p0
    _ = v12
    _ = v20
    return 2
}

func main() {
    var v1 uint8 = 94
    var v2 float64 = ((1e10 + float64((0.1 - 0.1))) + 3.14159)
    var v3 bool = (4294967295 >= 18446744073709551615)
    var v4 int = 392
    var v5 float64 = v2
    v3 = (0 < 0)
    var v21 uint64 = 458
    v22 := []int{v4, 9223372036854775807, (v4 ^ 101), v4}
    var v23 int64 = 89
    v24 := GenS2{F0: 243, F1: -1}
    var v25 int32 = 2147483647
    v4 = v4
    var v26 int16 = 32767
    switch v4 {
    case -9223372036854775808:
        v4 = -2147483648
    default:
        for fi27 := 0; fi27 < 18; fi27++ {
            _ = len(v22)
            v4 += v4
            _ = fi27
        }
    }
    v28 := GenS1{F0: -1}
    v29 := GenS2{F0: v23, F1: (v25 - (-(v25)))}
    for fi30 := 0; fi30 < 12; fi30++ {
        var v31 float32 = (-2.25 + 3.14159)
        fmt.Println(v1)
        v3 = (1 <= 0)
        _ = fi30
        _ = v31
    }
    var v32 bool = v3
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
