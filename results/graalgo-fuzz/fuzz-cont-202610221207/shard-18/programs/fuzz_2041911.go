package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool; F1 int8; F2 uint }
func genFn0(p0 int, p1 uint64) uint16 {
    p0 = (p0 - (p0 | p0))
    var v6 uint8 = 46
    p0 = 319
    _ = p0
    _ = p1
    _ = v6
    return 65535
}
func genFn1(p0 int8, p1 uint16, p2 int16) uint16 {
    var v7 uint16 = 255
    v8 := GenS1{F0: -286, F1: 127, F2: ("a" + ("" + ""))}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (v7 ^ genFn0(42, 4294967295))
}

func main() {
    v1 := []int{0, 32768, -522}
    var v2 float64 = ((1.5 + float64(-2.25)) / 1.0)
    v3 := make(map[string]string)
    var v4 uint64 = 18446744073709551615
    var v5 int8 = 100
    fmt.Println(v2)
    switch int(v4) {
    case -2147483648:
        switch "hello" {
        case "foo bar":
            if true {
                if (true && false) {
                    switch "\n\t" {
                    case "\n\t":
                        switch 32767 {
                        case -706:
                            for fi9 := 0; fi9 < 8; fi9++ {
                                delete(v3, "\n\t")
                                v10 := v3[("a" + "a")]
                                v2 = float64(3.14159)
                                _ = fi9
                                _ = v10
                            }
                        }
                    default:
                        var v11 uint = 18446744073709551615
                        _ = v11
                    }
                } else {
                    fmt.Println(v4, v2, v5)
                    v12 := v3["a"]
                    _ = v12
                }
            }
        case "hello":
            v4 = 948
        default:
            delete(v3, "hello")
        }
    case 0:
        if false {
            v13 := GenS1{F0: -29, F1: 128, F2: ("" + "")}
            _ = v13
        }
    case 226:
        for fi14 := 0; fi14 < 8; fi14++ {
            fmt.Println(v4)
            _ = fi14
        }
    default:
        if (32767 < genFn1(v5, uint16(v4), 1)) {
            v2 -= 1.5
        }
    }
    if true {
        v4 -= v4
        if (!false) {
            v5 = v5
        } else {
            v15 := GenS2{F0: ("\n\t" + "foo bar")}
            if (65535 == 652) {
                for fi16 := 0; fi16 < 13; fi16++ {
                    switch "a" {
                    case "\n\t":
                        var v17 bool = true
                        _ = v17
                    case "a":
                        v18 := GenS2{F0: ("\n\t" + ("foo bar" + "a"))}
                        _ = v18
                    }
                    fmt.Println(v5)
                    var v19 uint = 331
                    _ = fi16
                    _ = v19
                }
            }
            _ = v15
        }
    }
    fmt.Println(v2, v2)
    var v20 uint32 = 1
    v21 := GenS3{F0: (("\n\t" + "\n\t") == (("hello" + "") + "\n\t")), F1: v5, F2: 65535}
    v2 = 1.5
    v5 = v5
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
