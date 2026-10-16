package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string; F1 uint64; F2 int32 }
type GenS3 struct { F0 int32; F1 int; F2 int16 }
func genFn0(p0 int64) uint64 {
    v5 := []int{277, 9223372036854775807}
    p0 = p0
    p0 = p0
    switch "hello" {
    case "a":
        if true {
            p0 = p0
            for fi6 := 0; fi6 < 2; fi6++ {
                if len(v5) > 0 {
                    v5[0] = fi6
                }
                fmt.Println(fi6)
                var v7 int64 = (-(p0))
                _ = fi6
                _ = v7
            }
            p0 += (p0 % ((p0) | 1))
        } else {
            p0 -= p0
            v5 = append(v5, -878)
        }
    }
    p0 = p0
    _ = p0
    _ = v5
    return 0
}
func genFn1(p0 int, p1 uint) int64 {
    v8 := GenS1{F0: 2}
    p0 = (^(p0))
    v9 := GenS1{F0: 18446744073709551615}
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return -9223372036854775808
}

func main() {
    var v1 int16 = 1000
    var v2 int64 = 748
    var v3 uint8 = 1
    v4 := make(map[string]uint32)
    v1 -= (v1 | (-(v1)))
    v10 := []int{128, 128}
    var v11 int32 = 710
    var v12 int = 2147483647
    if (!((1 == 958) || (false || false))) {
        v2 += (v2 >> ((18446744073709551615) & 63))
        for fi13 := 0; fi13 < 5; fi13++ {
            fmt.Println(v11, v3, v1)
            switch (("hello" + "") + "foo bar") {
            case "foo bar":
                if len(v10) > 0 {
                    v10[0] = (-(fi13))
                }
            case "":
                fmt.Println(v2, v2, v1)
            }
            _ = fi13
        }
        for fi14 := 0; fi14 < 4; fi14++ {
            fmt.Println(v2, fi14, v2)
            _ = fi14
        }
    } else {
        switch "\n\t" {
        case "hello":
            var v15 int8 = -1
            _ = v15
        }
        var v16 float32 = (0.1 * float32(0.0))
        fmt.Println(v1)
        _ = v16
    }
    v4["\n\t"] = 4294967295
    var v17 bool = true
    v18 := GenS1{F0: genFn0((v2 / ((v2) | 1)))}
    for fi19 := 0; fi19 < 10; fi19++ {
        v1 = int16(fi19)
        _ = fi19
    }
    if v17 {
        v20 := GenS1{F0: 18446744073709551615}
        switch v12 {
        case 9223372036854775807:
            v21 := GenS2{F0: ("\n\t" + ("" + (("hello" + ("a" + "foo bar")) + "a"))), F1: 710, F2: 0}
            _ = v21
        case -1:
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
        case -427:
            v22 := GenS3{F0: (v11 << ((255) & 63)), F1: 1, F2: (v1 & (v1 ^ (v1 - v1)))}
            _ = v22
        default:
            if v17 {
                v2 = (v2 >> ((382) & 63))
            }
        }
        _ = v20
    } else {
        var v23 int = v12
        for fi24 := 0; fi24 < 7; fi24++ {
            for fi25 := 0; fi25 < 5; fi25++ {
                v11 += (v11 % ((v11) | 1))
                fmt.Println(v3, fi25, v3)
                v26 := GenS1{F0: 18446744073709551615}
                _ = fi25
                _ = v26
            }
            var v27 int = (v12 + 32767)
            var v28 uint8 = (v3 + (v3 * v3))
            _ = fi24
            _ = v27
            _ = v28
        }
        v3 = (v3 >> ((2) & 63))
        _ = v23
    }
    if len(v10) > 0 {
        v10[0] = 0
    }
    var v29 uint8 = 196
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v17
    _ = v18
    _ = v2
    _ = v29
    _ = v3
    _ = v4
}
