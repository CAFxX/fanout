package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 bool }
type GenS2 struct { F0 int16; F1 uint32 }
func genFn0() uint8 {
    var v4 int16 = 32767
    fmt.Println(v4)
    fmt.Println(v4)
    for fi5 := 0; fi5 < 20; fi5++ {
        fmt.Println(fi5)
        _ = fi5
    }
    if false {
        v4 = v4
        fmt.Println(v4)
    }
    _ = v4
    return 155
}
func genFn1(p0 int64) int32 {
    var v6 string = (("foo bar" + (("foo bar" + "\n\t") + ("a" + "\n\t"))) + "foo bar")
    v7 := GenS1{F0: -32768, F1: (p0 | p0), F2: (65535 > 2)}
    switch 807 {
    case 4294967295:
        v8 := []int{}
        _ = v8
    case 1:
        var v9 uint64 = 0
        _ = v9
    case -2147483648:
        switch (v6 + v6) {
        case "hello":
            v10 := GenS1{F0: 1, F1: (p0 >> ((1) & 63)), F2: true}
            _ = v10
        default:
            switch 1 {
            case 127:
                v11 := []int{128, 1}
                _ = v11
            case 4294967295:
                v6 = ("\n\t" + ("foo bar" + (((v6 + "a") + (v6 + (v6 + v6))) + v6)))
            default:
                var v12 string = v6
                _ = v12
            }
        }
    default:
        v13 := GenS1{F0: 0, F1: (-(p0)), F2: false}
        _ = v13
    }
    _ = p0
    _ = v6
    _ = v7
    return 2147483647
}

func main() {
    v1 := []uint32{634, 2147483648, 2}
    var v2 int64 = 0
    var v3 float64 = (1e10 * 0.0)
    var v14 uint64 = 248
    for fi15 := 0; fi15 < 10; fi15++ {
        v1 = append(v1, 1)
        v16 := GenS2{F0: 1, F1: 170}
        _ = fi15
        _ = v16
    }
    v2 = v2
    var v17 int = -605
    if (42 == 0) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if (2 == 255) {
            var v18 uint64 = v14
            fmt.Println(v17, v18)
            _ = v18
        } else {
            if (((v3 / ((float64(((0.1 + 3.14159) - 1e10)) * (3.14159 + v3)) - (float64(float32(v3)) / ((((v3 + float64(1.5)) - v3) * float64(((1e10 + 1e10) / 0.1))) / (v3 + v3))))) - v3) >= v3) {
                v2 -= (v2 / (((v2 ^ (v2 & v2))) | 1))
                fmt.Println(v14)
                v14 = 4294967295
            } else {
                v14 -= (v14 % (((v14 << ((2) & 63))) | 1))
            }
        }
        if (667 == v2) {
            v1 = append(v1, 4294967295)
            if len(v1) > 0 {
                v1[0] = 0
            }
        }
    } else {
        switch (((("\n\t" + "a") + ("\n\t" + "foo bar")) + "foo bar") + ("foo bar" + (("a" + "") + "a"))) {
        case "":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        fmt.Println(v3)
    }
    for fi19 := 0; fi19 < 20; fi19++ {
        if len(v1) > 0 {
            v1[0] = 845
        }
        var v20 int8 = 1
        var v21 int16 = 1000
        _ = fi19
        _ = v20
        _ = v21
    }
    for fi22 := 0; fi22 < 12; fi22++ {
        v2 = v2
        _ = fi22
    }
    fmt.Println(v3)
    fmt.Println(v3)
    v3 += 0.0
    if len(v1) > 0 {
        v1[0] = 65535
    }
    if len(v1) > 0 {
        v1[0] = 1
    }
    var v23 uint32 = 2147483647
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v17
    _ = v2
    _ = v23
    _ = v3
}
