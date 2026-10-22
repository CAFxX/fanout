package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
func genFn0(p0 int32, p1 uint64) int8 {
    fmt.Println(p1, p0)
    for fi4 := 0; fi4 < 15; fi4++ {
        if (-2.25 > ((0.1 / 1.5) / (0.1 / 1.0))) {
            p0 = (p0 ^ -1)
            if (-1 == 176) {
                p0 = p0
            }
        }
        var v5 float64 = ((((0.1 - 1e10) + -2.25) / -2.25) * 1.5)
        _ = fi4
        _ = v5
    }
    v6 := []int{-2147483648}
    var v7 uint32 = 625
    v8 := GenS1{F0: -9223372036854775808, F1: 0}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return -128
}

func main() {
    v1 := []string{}
    var v2 float32 = float32(0.0)
    var v3 int32 = -193
    v2 = -2.25
    for fi9 := 0; fi9 < 11; fi9++ {
        v2 = v2
        _ = fi9
    }
    switch ("a" + "\n\t") {
    case "foo bar":
        v3 += 42
    default:
        var v10 uint64 = 111
        _ = v10
    }
    if (-249 == 2147483647) {
        for fi11 := 0; fi11 < 17; fi11++ {
            v3 = (v3 * 1)
            v12 := GenS1{F0: (fi11 / ((fi11) | 1)), F1: fi11}
            _ = fi11
            _ = v12
        }
    } else {
        v1 = append(v1, "hello")
        var v13 int16 = 1000
        v3 = -1
        _ = v13
    }
    v2 = v2
    switch 32767 {
    case 32768:
        var v14 string = (("" + ("hello" + "a")) + ("hello" + "foo bar"))
        _ = v14
    case -1:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    if (true || (true || (((1 != 42) || false) || ("foo bar" <= "\n\t")))) {
        if false {
            v15 := GenS1{F0: 0, F1: -424}
            if false {
                var v16 uint32 = 2147483647
                _ = v16
            }
            for fi17 := 0; fi17 < 4; fi17++ {
                v3 += (v3 * -644)
                _ = fi17
            }
            _ = v15
        }
    } else {
        v18 := GenS1{F0: 9223372036854775807, F1: 32768}
        for fi19 := 0; fi19 < 9; fi19++ {
            fmt.Println(v2, v2, v3)
            v3 = (v3 ^ (v3 << ((864) & 63)))
            if false {
                switch "foo bar" {
                case "":
                    var v20 string = ("\n\t" + "")
                    _ = v20
                default:
                    v2 -= ((v2 - v2) / v2)
                }
                v21 := GenS1{F0: 0, F1: (fi19 / ((fi19) | 1))}
                if false {
                    v2 += v2
                    var v22 float32 = (v2 * v2)
                    _ = v22
                }
                _ = v21
            }
            _ = fi19
        }
        if len(v1) > 0 {
            v1[0] = ((("" + ("\n\t" + "a")) + "foo bar") + "a")
        }
        _ = v18
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi23 := 0; fi23 < 4; fi23++ {
        v24 := GenS1{F0: (-(fi23)), F1: (fi23 & 4294967295)}
        v3 += (v3 + (v3 % (((v3 ^ (v3 % ((v3) | 1)))) | 1)))
        v25 := GenS1{F0: fi23, F1: (-(fi23))}
        _ = fi23
        _ = v24
        _ = v25
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
}
