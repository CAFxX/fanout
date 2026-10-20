package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 uint64 }
func genFn0() string {
    var v5 float64 = -2.25
    v6 := []int{128}
    _ = v5
    _ = v6
    return ("hello" + "foo bar")
}

func main() {
    v1 := []string{"hello", (("hello" + ("" + ("foo bar" + "a"))) + ("foo bar" + "")), ("\n\t" + "hello")}
    var v2 int32 = 42
    var v3 uint = 18446744073709551615
    var v4 int8 = 0
    v1 = append(v1, "hello")
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v1 = append(v1, "")
    v3 = (v3 + 65535)
    fmt.Println(v4, v3, v2)
    v1 = append(v1, "a")
    switch "foo bar" {
    case "foo bar":
        var v7 int32 = v2
        _ = v7
    case "hello":
        if (((v3 >> (((v3 & v3)) & 63)) <= 0) || false) {
            if (!true) {
                if len(v1) > 0 {
                    v1[0] = "a"
                }
                switch 2147483647 {
                case -1:
                    fmt.Println(v4, v4)
                case 851:
                    for fi8 := 0; fi8 < 19; fi8++ {
                        switch "foo bar" {
                        case "foo bar":
                            for fi9 := 0; fi9 < 3; fi9++ {
                                v10 := GenS1{F0: 255, F1: (v4 >> ((18446744073709551615) & 63)), F2: 48}
                                v4 = 1
                                fmt.Println(v3, fi8, v3)
                                _ = fi9
                                _ = v10
                            }
                        case "hello":
                            if (!false) {
                                v11 := GenS1{F0: 217, F1: -1, F2: 90}
                                _ = v11
                            } else {
                                if (!((307 == 32767) || (("hello" + "foo bar") >= "hello"))) {
                                    var v12 int8 = (-(v4))
                                    _ = v12
                                }
                            }
                        }
                        for fi13 := 0; fi13 < 20; fi13++ {
                            _ = len(v1)
                            fmt.Println(v4)
                            fmt.Println(fi8, v2)
                            var v14 int64 = -491
                            _ = fi13
                            _ = v14
                        }
                        if false {
                            if (!true) {
                                if true {
                                    var v15 float64 = (float64(3.14159) - (-2.25 + 0.1))
                                    _ = v15
                                }
                            }
                        } else {
                            var v16 uint = v3
                            if true {
                                if (!true) {
                                    var v17 float64 = 0.1
                                    _ = v17
                                }
                                _ = len(v1)
                                for fi18 := 0; fi18 < 13; fi18++ {
                                    v19 := GenS1{F0: 0, F1: (v4 - 127), F2: uint64(v3)}
                                    _ = fi18
                                    _ = v19
                                }
                            } else {
                                for fi20 := 0; fi20 < 10; fi20++ {
                                    v21 := GenS1{F0: 255, F1: 127, F2: 18446744073709551615}
                                    _ = fi20
                                    _ = v21
                                }
                            }
                            var v22 float64 = (-2.25 - (((0.0 + 1e10) / (((float64((1e10 - 1e10)) / 1.0) * 3.14159) + 1.5)) - (float64((1e10 + 0.0)) + (0.1 / 1e10))))
                            _ = v16
                            _ = v22
                        }
                        _ = fi8
                    }
                case 42:
                    fmt.Println(v2, v4)
                }
            }
            if (4294967295 <= 2147483648) {
                if len(v1) > 0 {
                    v1[0] = "hello"
                }
                if (false && false) {
                    if len(v1) > 0 {
                        v1[0] = "\n\t"
                    }
                    v3 += (^(v3))
                } else {
                    var v23 uint64 = 4294967295
                    _ = v23
                }
            } else {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            v1 = append(v1, "hello")
        }
    default:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    var v24 int64 = -402
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
}
