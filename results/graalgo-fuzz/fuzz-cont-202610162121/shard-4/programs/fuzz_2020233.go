package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 string; F1 int16 }
func genFn0(p0 string) int32 {
    p0 = (p0 + (p0 + "\n\t"))
    v4 := GenS1{F0: 0, F1: -1}
    var v5 int32 = -2147483648
    p0 = ("" + (((p0 + (p0 + p0)) + p0) + "\n\t"))
    p0 = (p0 + ((p0 + "foo bar") + ("\n\t" + ("hello" + (p0 + p0)))))
    _ = p0
    _ = v4
    _ = v5
    return v5
}
func genFn1(p0 uint16, p1 int8) uint16 {
    if true {
        v6 := GenS1{F0: 42, F1: -128}
        for fi7 := 0; fi7 < 2; fi7++ {
            p1 = p1
            var v8 int32 = 0
            _ = fi7
            _ = v8
        }
        var v9 float64 = (3.14159 - ((1.5 / 1.0) - -2.25))
        _ = v6
        _ = v9
    }
    switch "hello" {
    case "hello":
        p1 = p1
    case "":
        v10 := GenS1{F0: 0, F1: p1}
        _ = v10
    }
    p1 -= (^(p1))
    p1 = -95
    v11 := GenS1{F0: p1, F1: (p1 / ((127) | 1))}
    _ = p0
    _ = p1
    _ = v11
    return (p0 - p0)
}

func main() {
    var v1 string = ("\n\t" + ("hello" + "a"))
    var v2 uint64 = 4294967295
    v3 := []string{v1, "a", "\n\t", (v1 + v1)}
    _ = len(v3)
    switch 42 {
    case 396:
        v2 = 2
    default:
        var v12 uint64 = v2
        _ = v12
    }
    switch (v1 + "hello") {
    case "a":
        if false {
            if (false && false) {
                switch -1000 {
                case 127:
                    v13 := GenS2{F0: ("foo bar" + ""), F1: 1000}
                    _ = v13
                case 32768:
                    v3 = append(v3, v1)
                case -9223372036854775808:
                    var v14 int8 = 1
                    _ = v14
                default:
                    var v15 uint = 4294967295
                    _ = v15
                }
            } else {
                var v16 uint16 = 42
                if true {
                    fmt.Println(v1)
                    switch -2147483648 {
                    case 128:
                        for fi17 := 0; fi17 < 8; fi17++ {
                            if len(v3) > 0 {
                                v3[0] = ((((v1 + ((v1 + "") + v1)) + "\n\t") + v1) + v1)
                            }
                            _ = fi17
                        }
                    default:
                        var v18 bool = (9223372036854775807 == -9223372036854775808)
                        _ = v18
                    }
                    v16 += (v16 >> ((1) & 63))
                } else {
                    if len(v3) > 0 {
                        v3[0] = ((v1 + ((((v1 + ("foo bar" + (("foo bar" + v1) + (v1 + ((v1 + (v1 + v1)) + v1))))) + v1) + v1) + (((v1 + v1) + v1) + ""))) + "foo bar")
                    }
                    var v19 int64 = -9223372036854775808
                    fmt.Println(v2)
                    _ = v19
                }
                v3 = append(v3, v1)
                _ = v16
            }
            v20 := GenS1{F0: 127, F1: 119}
            _ = v20
        }
    default:
        v1 = "foo bar"
    }
    if len(v3) > 0 {
        v3[0] = v1
    }
    v2 = (v2 << ((65535) & 63))
    _ = len(v3)
    if true {
        switch 9223372036854775807 {
        case -433:
            var v21 uint64 = (v2 % ((251) | 1))
            _ = v21
        case 2147483647:
            var v22 int8 = 20
            _ = v22
        }
        for fi23 := 0; fi23 < 17; fi23++ {
            var v24 float32 = (0.0 / (((0.1 - -2.25) - 0.0) + (0.1 - -2.25)))
            var v25 uint = 42
            _ = fi23
            _ = v24
            _ = v25
        }
        var v26 uint32 = 4294967295
        _ = v26
    }
    if (42 >= 251) {
        v27 := GenS1{F0: 127, F1: 127}
        _ = v27
    }
    v1 = v1
    if (-1 >= -1) {
        v1 = "a"
        var v28 int = 127
        _ = v28
    } else {
        v1 = "\n\t"
    }
    if (genFn0((v1 + "a")) >= 0) {
        for fi29 := 0; fi29 < 20; fi29++ {
            _ = len(v3)
            v30 := GenS2{F0: ("hello" + v1), F1: 1000}
            v1 = ""
            _ = fi29
            _ = v30
        }
        fmt.Println(v2, v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v3
}
