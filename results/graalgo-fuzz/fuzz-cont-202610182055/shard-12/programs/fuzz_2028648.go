package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
func genFn0(p0 int32, p1 string) string {
    for fi4 := 0; fi4 < 14; fi4++ {
        p0 += p0
        var v5 string = ("\n\t" + (("a" + (p1 + (((p1 + (("foo bar" + (p1 + (p1 + (((("" + ("a" + p1)) + "foo bar") + "\n\t") + (p1 + (p1 + "hello")))))) + ("foo bar" + "hello"))) + p1) + "foo bar"))) + ("" + (((("foo bar" + (((("a" + (p1 + p1)) + p1) + p1) + "hello")) + p1) + (p1 + p1)) + p1))))
        _ = fi4
        _ = v5
    }
    v6 := GenS1{F0: "foo bar", F1: 255}
    switch -2147483648 {
    case 0:
        p0 = (-(p0))
    case -999:
        var v7 int = 2147483647
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v6
    return "hello"
}
func genFn1() uint16 {
    var v8 int = -569
    fmt.Println(v8)
    v8 = (v8 & v8)
    v8 = v8
    _ = v8
    return 9
}

func main() {
    var v1 int16 = -752
    var v2 float32 = -2.25
    var v3 int = 0
    for fi9 := 0; fi9 < 9; fi9++ {
        v3 += v3
        v10 := GenS1{F0: "hello", F1: 65535}
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        v12 := []int{(fi11 - (fi11 ^ 32767)), fi11, (fi11 / ((fi11) | 1))}
        var v13 uint32 = 2
        if len(v12) > 0 {
            v12[0] = (fi11 - fi11)
        }
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 int16 = (v1 + -1)
    v15 := []int{v3, 4294967295, -585}
    if (false && false) {
        v15 = append(v15, 385)
    } else {
        v15 = append(v15, (v3 + -916))
        switch (("a" + ("\n\t" + "foo bar")) + "a") {
        case "foo bar":
            for fi16 := 0; fi16 < 11; fi16++ {
                if (!(1 > 2147483647)) {
                    v1 += v1
                    var v17 uint = 255
                    _ = v17
                } else {
                    var v18 int64 = -1
                    _ = v18
                }
                _ = fi16
            }
        case "hello":
            var v19 int64 = 0
            _ = v19
        case "\n\t":
            for fi20 := 0; fi20 < 15; fi20++ {
                fmt.Println(v3)
                v14 = v1
                _ = fi20
            }
        default:
            var v21 float32 = -2.25
            _ = v21
        }
        v14 += (v14 / ((v1) | 1))
    }
    v3 += -812
    _ = len(v15)
    var v22 uint = 18446744073709551615
    var v23 int64 = -9223372036854775808
    if (((v14 | v14) != v1) && true) {
        if ((65535 != uint16(v22)) || true) {
            for fi24 := 0; fi24 < 10; fi24++ {
                fmt.Println(v1, v1)
                v14 = 1
                _ = fi24
            }
        }
        if false {
            var v25 int8 = 100
            for fi26 := 0; fi26 < 4; fi26++ {
                v22 -= v22
                _ = fi26
            }
            _ = v25
        } else {
            var v27 uint = v22
            var v28 float64 = 1e10
            _ = v27
            _ = v28
        }
    } else {
        v14 += (v1 | v14)
        v3 -= -90
    }
    fmt.Println(v23, v2)
    switch ("a" + "hello") {
    case "a":
        v23 = v23
    case "foo bar":
        if len(v15) >= 2 {
            v15 = v15[1:len(v15)]
        }
    default:
        if len(v15) > 0 {
            v15[0] = (v3 * v3)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v22
    _ = v23
    _ = v3
}
