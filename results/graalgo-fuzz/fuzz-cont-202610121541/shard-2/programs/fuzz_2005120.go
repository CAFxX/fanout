package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 int8 }
func genFn0(p0 int64) string {
    p0 += (p0 & (p0 % ((p0) | 1)))
    fmt.Println(p0)
    for fi6 := 0; fi6 < 15; fi6++ {
        if (80 < 128) {
            var v7 uint16 = 1
            switch ("" + "foo bar") {
            case "":
                fmt.Println(fi6)
            default:
                for fi8 := 0; fi8 < 12; fi8++ {
                    p0 = p0
                    switch fi8 {
                    case 427:
                        var v9 float64 = 0.1
                        _ = v9
                    default:
                        var v10 uint64 = 94
                        _ = v10
                    }
                    fmt.Println(v7, v7, fi8)
                    _ = fi8
                }
            }
            v7 += 878
            _ = v7
        }
        _ = fi6
    }
    v11 := []int{-541}
    _ = p0
    _ = v11
    return ("foo bar" + "hello")
}
func genFn1(p0 uint, p1 uint32, p2 bool) uint64 {
    fmt.Println(p0)
    v12 := []int{32767, 2147483647, 42, 0}
    p1 -= p1
    p2 = true
    if p2 {
        for fi13 := 0; fi13 < 1; fi13++ {
            p1 -= p1
            if len(v12) > 0 {
                v12[0] = (-(fi13))
            }
            fmt.Println(p2)
            _ = fi13
        }
        for fi14 := 0; fi14 < 6; fi14++ {
            var v15 int = fi14
            var v16 int16 = 32767
            _ = fi14
            _ = v15
            _ = v16
        }
        for fi17 := 0; fi17 < 19; fi17++ {
            var v18 uint64 = 1
            fmt.Println(p2)
            _ = fi17
            _ = v18
        }
    } else {
        for fi19 := 0; fi19 < 10; fi19++ {
            _ = len(v12)
            _ = fi19
        }
        fmt.Println(p1)
        var v20 int64 = -9223372036854775808
        _ = v20
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 18446744073709551615
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]int)
    var v3 bool = true
    v4 := []string{}
    var v5 float32 = (((float32(0.1) - 0.0) - 0.0) / 1.5)
    v5 -= float32(1.5)
    v21 := GenS1{F0: (((false || v3) || (false && v3)) || false), F1: 883, F2: 1}
    if (!((9223372036854775807 < -402) || false)) {
        switch -2147483648 {
        case 729:
            v2[4294967295] = 2147483647
        }
    }
    v2[2147483647] = -1
    v5 -= float32(-2.25)
    var v22 uint8 = 127
    var v23 float64 = 1e10
    switch (("a" + ("\n\t" + "foo bar")) + "foo bar") {
    case "":
        v24 := v2[-193]
        _ = v24
    case "\n\t":
        v25 := GenS1{F0: (0 > 42), F1: 0, F2: 0}
        _ = v25
    }
    for fi26 := 0; fi26 < 2; fi26++ {
        fmt.Println(v22)
        _ = fi26
    }
    var v27 uint = 507
    switch -721 {
    case -949:
        if v3 {
            v23 = v23
        }
    case 120:
        v22 -= (v22 >> ((v27) & 63))
    default:
        v3 = v3
    }
    if (32767 > 32767) {
        delete(v1, "\n\t")
        var v28 int16 = -32768
        v23 += v23
        _ = v28
    } else {
        v23 = v23
        v29 := GenS1{F0: (((v3 || false) && (!(!((148 < (v22 ^ 146)) || v3)))) || v3), F1: 42, F2: -128}
        v22 -= v22
        _ = v29
    }
    v3 = v3
    fmt.Println(v27)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
