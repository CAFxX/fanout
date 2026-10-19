package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 int; F1 uint }
func genFn0() uint {
    v5 := GenS1{F0: 42, F1: 551}
    var v6 int64 = 9223372036854775807
    switch 42 {
    case 2147483647:
        v6 -= (v6 << ((255) & 63))
    case -9223372036854775808:
        if true {
            v6 = v6
            v6 = 9223372036854775807
            fmt.Println(v6)
        } else {
            v7 := []int{}
            _ = v7
        }
    case 32768:
        v6 -= v6
    }
    var v8 int32 = -1
    if true {
        var v9 int8 = -99
        var v10 uint16 = 0
        _ = v10
        _ = v9
    } else {
        v11 := []int{802, 608}
        if (true || (65535 >= 173)) {
            v8 = (v8 | v8)
            for fi12 := 0; fi12 < 20; fi12++ {
                for fi13 := 0; fi13 < 13; fi13++ {
                    var v14 int16 = int16(v8)
                    var v15 string = "a"
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                for fi16 := 0; fi16 < 12; fi16++ {
                    if ((146 != 0) && true) {
                        var v17 float64 = 1.5
                        _ = v17
                    }
                    var v18 int8 = -10
                    fmt.Println(v18, fi12)
                    _ = fi16
                    _ = v18
                }
                v6 -= int64(v8)
                _ = fi12
            }
            var v19 float64 = float64(0.0)
            _ = v19
        }
        _ = v11
    }
    _ = v5
    _ = v6
    _ = v8
    return 4294967295
}
func genFn1(p0 uint32, p1 uint8, p2 int16) string {
    var v20 bool = false
    switch int(p2) {
    case 32767:
        var v21 uint = 18446744073709551615
        _ = v21
    }
    for fi22 := 0; fi22 < 17; fi22++ {
        v23 := GenS1{F0: (p2 * p2), F1: (p0 - (p0 << ((uint(p1)) & 63)))}
        _ = fi22
        _ = v23
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    return "\n\t"
}

func main() {
    v1 := []int{128, 821}
    var v2 int32 = -848
    var v3 int8 = -128
    var v4 bool = (!false)
    v3 -= 68
    v24 := GenS2{F0: (v2 ^ 1), F1: 65535}
    if len(v1) > 0 {
        v1[0] = 270
    }
    if (!v4) {
        var v25 uint64 = 130
        _ = v25
    } else {
        for fi26 := 0; fi26 < 11; fi26++ {
            v27 := GenS3{F0: (fi26 << ((65535) & 63)), F1: 18446744073709551615}
            var v28 bool = (!(-520 >= 1))
            _ = fi26
            _ = v27
            _ = v28
        }
    }
    v4 = (220 <= 739)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v29 int16 = 32767
    v3 = (^(v3))
    fmt.Println(v29, v4, v3)
    var v30 int = 128
    var v31 uint16 = 142
    var v32 uint8 = 1
    v32 = 47
    if len(v1) > 0 {
        v1[0] = (v30 % ((2147483647) | 1))
    }
    if false {
        fmt.Println(v31, v31)
        for fi33 := 0; fi33 < 1; fi33++ {
            var v34 string = ""
            fmt.Println(v4, v31, v34)
            var v35 string = v34
            _ = fi33
            _ = v34
            _ = v35
        }
    } else {
        if len(v1) > 0 {
            v1[0] = v30
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
}
