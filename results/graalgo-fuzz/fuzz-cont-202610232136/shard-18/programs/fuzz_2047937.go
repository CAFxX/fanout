package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int8; F1 int32; F2 bool }
type GenS5 struct { F0 int32; F1 int32; F2 int }
func genFn0(p0 uint16, p1 int, p2 int16) int8 {
    for fi7 := 0; fi7 < 15; fi7++ {
        p1 = fi7
        _ = fi7
    }
    var v8 string = ""
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 74
}
func genFn1(p0 int32, p1 uint16) uint16 {
    p0 = p0
    p0 = p0
    v9 := GenS1{F0: true, F1: -32768}
    p1 += 32767
    if (!(-2.25 > float32((0.0 / ((-2.25 + 1e10) * 3.14159))))) {
        fmt.Println(p0, p0)
        p0 -= (p0 + p0)
    } else {
        p0 = (p0 % (((^(p0))) | 1))
        p0 += p0
    }
    _ = p0
    _ = p1
    _ = v9
    return (p1 % (((p1 / ((p1) | 1))) | 1))
}

func main() {
    var v1 string = "\n\t"
    var v2 bool = (65535 >= 1)
    var v3 int32 = 1
    v4 := []uint32{2, 0, 2147483648, 42}
    var v5 int64 = 0
    v6 := []string{}
    v5 += (v5 - 1)
    var v10 float64 = 0.0
    v11 := GenS2{F0: 42, F1: 406}
    for fi12 := 0; fi12 < 4; fi12++ {
        _ = len(v4)
        _ = fi12
    }
    fmt.Println(v10)
    v5 = v5
    if (v1 > "") {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v2 = v2
        v1 = v1
    } else {
        switch v1 {
        case "\n\t":
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        case "hello":
            for fi13 := 0; fi13 < 9; fi13++ {
                v5 += v5
                if v2 {
                    var v14 bool = v2
                    fmt.Println(v5, v3, v14)
                    _ = v14
                } else {
                    var v15 int = (fi13 / ((fi13) | 1))
                    fmt.Println(v15, v2)
                    _ = len(v4)
                    _ = v15
                }
                _ = fi13
            }
        case "foo bar":
            var v16 float32 = 0.1
            _ = v16
        default:
            v10 = 3.14159
        }
        v17 := GenS2{F0: 32767, F1: 255}
        var v18 float64 = (float64(0.1) - float64(-2.25))
        _ = v17
        _ = v18
    }
    _ = len(v6)
    var v19 uint8 = 255
    fmt.Println(v2, v5)
    switch 9223372036854775807 {
    case 127:
        v19 += v19
    case 128:
        v20 := GenS3{F0: 32767}
        _ = v20
    default:
        var v21 uint = 0
        _ = v21
    }
    for fi22 := 0; fi22 < 6; fi22++ {
        v23 := GenS4{F0: 127, F1: 42, F2: v2}
        v24 := GenS5{F0: (v3 * 2147483647), F1: (-(v3)), F2: fi22}
        v3 -= v3
        _ = fi22
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 13; fi25++ {
        v6 = append(v6, ((((("hello" + "\n\t") + "\n\t") + v1) + "hello") + v1))
        var v26 int64 = v5
        _ = fi25
        _ = v26
    }
    for fi27 := 0; fi27 < 14; fi27++ {
        var v28 int = (fi27 << ((18446744073709551615) & 63))
        switch (v1 + v1) {
        case "foo bar":
            var v29 int = 459
            _ = v29
        default:
            for fi30 := 0; fi30 < 3; fi30++ {
                v19 -= v19
                _ = fi30
            }
        }
        switch "a" {
        case "\n\t":
            if (!v2) {
                var v31 uint = 1
                _ = v31
            }
        }
        _ = fi27
        _ = v28
    }
    fmt.Println(v10, v3)
    v3 -= v3
    switch 32768 {
    case 9223372036854775807:
        v5 -= 42
    default:
        v3 -= (^(v3))
    }
    fmt.Println(v19, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
