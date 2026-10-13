package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
func genFn0(p0 int16, p1 string) string {
    p1 = "a"
    var v5 bool = true
    var v6 int64 = 9223372036854775807
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return "a"
}
func genFn1(p0 uint32, p1 uint64, p2 bool) uint8 {
    v7 := []int{}
    if p2 {
        if len(v7) > 0 {
            v7[0] = -9223372036854775808
        }
        for fi8 := 0; fi8 < 12; fi8++ {
            p2 = (42 != 1)
            p2 = p2
            _ = fi8
        }
    } else {
        v9 := GenS1{F0: p0, F1: 127}
        _ = v9
    }
    p1 -= 458
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 42
}

func main() {
    var v1 bool = (1 != 0)
    v2 := []string{((("hello" + "a") + "a") + "foo bar"), ("" + "\n\t")}
    var v3 uint8 = 1
    v4 := make(map[string]uint32)
    if len(v2) > 0 {
        v2[0] = ("" + ("hello" + "hello"))
    }
    if (!(!(!false))) {
        _ = len(v2)
    } else {
        fmt.Println(v1)
        delete(v4, (("" + "") + "hello"))
    }
    if v1 {
        var v10 int32 = -2147483648
        var v11 int64 = -1
        v11 -= (v11 | (v11 % (((v11 & (v11 >> ((1) & 63)))) | 1)))
        _ = v10
        _ = v11
    }
    switch -9223372036854775808 {
    case 128:
        v3 += v3
    default:
        var v12 string = "\n\t"
        _ = v12
    }
    var v13 int32 = -1
    for fi14 := 0; fi14 < 8; fi14++ {
        var v15 float64 = -2.25
        _ = fi14
        _ = v15
    }
    var v16 uint16 = 255
    if v1 {
        var v17 int64 = 42
        var v18 uint64 = 11
        for fi19 := 0; fi19 < 9; fi19++ {
            _ = len(v4)
            switch (fi19 | 0) {
            case 32768:
                v16 += 32767
            case -810:
                for fi20 := 0; fi20 < 5; fi20++ {
                    v21 := v4["\n\t"]
                    var v22 float64 = (0.0 - (3.14159 - 3.14159))
                    _ = fi20
                    _ = v21
                    _ = v22
                }
            case 0:
                v17 -= -1
            }
            v18 += (v18 | (v18 % ((2) | 1)))
            _ = fi19
        }
        _ = v17
        _ = v18
    }
    v23 := GenS1{F0: 0, F1: genFn1(2, 0, ((0.1 >= float32((float64(1e10) / 1.0))) || (42 >= v16)))}
    _ = len(v4)
    var v24 float32 = float32(0.0)
    v24 -= v24
    v16 -= 403
    if (!(!(!v1))) {
        if v1 {
            if v1 {
                v13 = (v13 - -791)
                fmt.Println(v3)
            } else {
                fmt.Println(v13)
                v25 := GenS1{F0: 2147483648, F1: v3}
                _ = v25
            }
        }
        var v26 float32 = v24
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v16
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
