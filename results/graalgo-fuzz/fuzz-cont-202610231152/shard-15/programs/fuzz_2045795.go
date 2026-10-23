package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 uint16 }
type GenS2 struct { F0 uint64; F1 uint16; F2 int8 }
func genFn0(p0 uint32, p1 string, p2 int16) uint64 {
    p1 = ((p1 + ("hello" + p1)) + p1)
    var v7 float32 = 0.1
    v8 := GenS1{F0: -123, F1: 255, F2: 250}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 18446744073709551615
}
func genFn1() uint16 {
    v9 := GenS1{F0: -128, F1: 1, F2: 164}
    switch "" {
    case "hello":
        var v10 int = 2147483647
        _ = v10
    case "":
        var v11 string = "\n\t"
        _ = v11
    case "\n\t":
        var v12 uint16 = 320
        _ = v12
    default:
        var v13 string = ("\n\t" + "hello")
        _ = v13
    }
    switch ("a" + ("foo bar" + (("\n\t" + "") + (("foo bar" + "foo bar") + "hello")))) {
    case "\n\t":
        for fi14 := 0; fi14 < 7; fi14++ {
            var v15 int = (-(fi14))
            var v16 uint32 = 2147483648
            _ = fi14
            _ = v15
            _ = v16
        }
    case "a":
        if false {
            for fi17 := 0; fi17 < 20; fi17++ {
                v18 := GenS2{F0: 1, F1: 197, F2: 0}
                if ((!(-1 > 1)) && (float64(float32(float64(((1.5 * -2.25) / 3.14159)))) <= ((0.1 / (0.0 + -2.25)) - -2.25))) {
                    var v19 uint32 = 616
                    _ = v19
                }
                _ = fi17
                _ = v18
            }
            v20 := []int{127}
            _ = v20
        } else {
            var v21 int16 = -1
            _ = v21
        }
    case "foo bar":
        var v22 int8 = 1
        _ = v22
    default:
        var v23 uint = 18446744073709551615
        _ = v23
    }
    _ = v9
    return 2
}

func main() {
    var v1 int = 910
    var v2 bool = false
    v3 := []string{"hello", "", ("a" + "hello")}
    v4 := []string{"hello", "\n\t", "a", ("" + "\n\t")}
    var v5 string = ("hello" + ("foo bar" + "\n\t"))
    var v6 float64 = 0.0
    v1 -= (v1 - v1)
    if false {
        var v24 uint32 = 65535
        _ = v24
    } else {
        var v25 int16 = -426
        for fi26 := 0; fi26 < 3; fi26++ {
            v1 += (v1 | (v1 & 0))
            if len(v3) > 0 {
                v3[0] = ("a" + "a")
            }
            _ = fi26
        }
        v2 = (v2 && v2)
        _ = v25
    }
    if (130 >= 1) {
        if true {
            for fi27 := 0; fi27 < 8; fi27++ {
                v1 = v1
                _ = fi27
            }
        } else {
            fmt.Println(v2, v5)
        }
        v28 := GenS1{F0: 42, F1: 4294967295, F2: 32767}
        _ = v28
    }
    var v29 string = "a"
    v30 := GenS1{F0: 100, F1: 1, F2: 255}
    v31 := GenS2{F0: 18446744073709551615, F1: 42, F2: 0}
    fmt.Println(v29, v2, v5)
    if len(v4) > 0 {
        v4[0] = ("" + v5)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
