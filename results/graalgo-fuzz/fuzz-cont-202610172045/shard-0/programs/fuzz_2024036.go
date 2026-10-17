package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 bool) uint64 {
    switch "\n\t" {
    case "a":
        for fi4 := 0; fi4 < 14; fi4++ {
            var v5 uint32 = 0
            _ = fi4
            _ = v5
        }
    default:
        p0 = true
    }
    v6 := []int{9223372036854775807, 127}
    _ = p0
    _ = v6
    return 1
}

func main() {
    var v1 int = 42
    v2 := []string{"\n\t", ("" + "\n\t"), "hello"}
    var v3 uint64 = 1
    var v7 uint64 = (v3 * 1)
    fmt.Println(v3, v3)
    if (!true) {
        v8 := GenS1{F0: -82}
        for fi9 := 0; fi9 < 8; fi9++ {
            var v10 uint8 = 20
            switch (v1 & fi9) {
            case 4294967295:
                v2 = append(v2, "hello")
            case 128:
                var v11 float32 = 0.1
                _ = v11
            default:
                fmt.Println(v3, v7, v7)
            }
            _ = len(v2)
            _ = fi9
            _ = v10
        }
        _ = v8
    }
    switch v1 {
    case 127:
        if len(v2) > 0 {
            v2[0] = ""
        }
    case 2147483647:
        switch ("hello" + "") {
        case "a":
            v3 -= (v7 ^ (^(v7)))
        case "foo bar":
            v3 = 4294967295
        default:
            var v12 float64 = 3.14159
            _ = v12
        }
    default:
        if false {
            for fi13 := 0; fi13 < 12; fi13++ {
                for fi14 := 0; fi14 < 5; fi14++ {
                    var v15 uint16 = 1
                    if (genFn0((false && false)) >= v7) {
                        if (-128 >= 1) {
                            v15 = v15
                        }
                    }
                    v2 = append(v2, (("a" + "a") + "a"))
                    _ = fi14
                    _ = v15
                }
                var v16 uint32 = 464
                _ = fi13
                _ = v16
            }
        } else {
            if false {
                v7 = (v3 * 42)
                for fi17 := 0; fi17 < 5; fi17++ {
                    if false {
                        for fi18 := 0; fi18 < 2; fi18++ {
                            var v19 uint16 = 255
                            v3 += (^(v3))
                            _ = fi18
                            _ = v19
                        }
                    } else {
                        if true {
                            v20 := GenS1{F0: 118}
                            v21 := GenS2{F0: (v1 | v1), F1: 982}
                            _ = v20
                            _ = v21
                        }
                    }
                    v1 -= (-(v1))
                    v7 += (^(v3))
                    _ = fi17
                }
                v22 := GenS3{F0: uint8(v7)}
                _ = v22
            }
        }
    }
    if false {
        if (("" < ("hello" + ("hello" + ("a" + (((("" + ("hello" + "hello")) + "a") + "\n\t") + ("foo bar" + "")))))) && ((!((65535 != 562) && true)) || (0.0 >= (3.14159 * (1.5 / (3.14159 + 1.5)))))) {
            var v23 int8 = -128
            v7 -= v3
            _ = v23
        } else {
            if (true && (false && true)) {
                if ((("hello" + "hello") + "a") <= "foo bar") {
                    fmt.Println(v3, v1)
                    for fi24 := 0; fi24 < 12; fi24++ {
                        if (!(!false)) {
                            if len(v2) > 0 {
                                v2[0] = ((("a" + "\n\t") + "a") + "")
                            }
                            v25 := GenS1{F0: 1}
                            _ = v25
                        }
                        _ = fi24
                    }
                }
            }
            var v26 uint16 = 88
            _ = v26
        }
    }
    var v27 bool = (-1 <= 1)
    fmt.Println(v27, v1, v27)
    var v28 uint8 = 255
    _ = len(v2)
    var v29 float64 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v7
}
