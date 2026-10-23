package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8; F1 uint16 }
type GenS4 struct { F0 int; F1 int32 }
type GenS5 struct { F0 uint32; F1 uint }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint }
func genFn0(p0 string, p1 int32, p2 string) uint32 {
    p0 = "\n\t"
    for fi5 := 0; fi5 < 9; fi5++ {
        p2 = p2
        _ = fi5
    }
    p1 -= (p1 + p1)
    _ = p0
    _ = p1
    _ = p2
    return 2147483648
}

func main() {
    v1 := []uint32{108, 65535, 2147483647}
    var v2 float64 = 0.1
    var v3 uint32 = 722
    var v4 float32 = ((1.5 * (((1e10 / 0.1) / -2.25) - float32(float64((0.0 - ((1.5 * 0.0) * 0.0)))))) + (0.0 / (1.5 / float32(v2))))
    if (!((-1 <= -914) && false)) {
        fmt.Println(v4, v4)
        if (true && (v3 > genFn0(("a" + ""), -690, "hello"))) {
            v4 -= (float32(v2) + v4)
            var v6 string = "a"
            switch "hello" {
            case "\n\t":
                switch v6 {
                case "foo bar":
                    v2 = float64(float32(((1e10 - 0.1) - (-2.25 * v2))))
                case "":
                    for fi7 := 0; fi7 < 12; fi7++ {
                        v4 = v4
                        _ = fi7
                    }
                default:
                    if (0 < 370) {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        if (v6 >= "foo bar") {
                            if true {
                                fmt.Println(v6, v2, v6)
                            }
                            if false {
                                switch -131 {
                                case 232:
                                    var v8 int = -9223372036854775808
                                    _ = v8
                                case -994:
                                    fmt.Println(v2, v6)
                                }
                            } else {
                                if (0.1 < ((v2 / v2) / ((v2 / float64(1e10)) * v2))) {
                                    v9 := GenS1{F0: 1, F1: -417, F2: true}
                                    for fi10 := 0; fi10 < 13; fi10++ {
                                        var v11 float32 = (0.1 - v4)
                                        var v12 int32 = 0
                                        v13 := GenS2{F0: 0}
                                        _ = fi10
                                        _ = v11
                                        _ = v12
                                        _ = v13
                                    }
                                    _ = v9
                                }
                                v14 := GenS3{F0: -128, F1: 812}
                                _ = len(v1)
                                _ = v14
                            }
                            v15 := GenS4{F0: 1, F1: 1}
                            _ = v15
                        }
                    }
                }
            }
            _ = v6
        } else {
            for fi16 := 0; fi16 < 15; fi16++ {
                fmt.Println(v3)
                var v17 float64 = -2.25
                switch -2147483648 {
                case 484:
                    for fi18 := 0; fi18 < 1; fi18++ {
                        var v19 int16 = 42
                        _ = fi18
                        _ = v19
                    }
                case 976:
                    switch ("hello" + ("" + "foo bar")) {
                    case "":
                        fmt.Println(v2, v4)
                    case "foo bar":
                        v20 := GenS4{F0: (fi16 << ((831) & 63)), F1: 1}
                        _ = v20
                    }
                case -716:
                    v2 = 1.5
                }
                _ = fi16
                _ = v17
            }
            fmt.Println(v3, v2)
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v3 -= 1
    _ = len(v1)
    for fi21 := 0; fi21 < 17; fi21++ {
        fmt.Println(v4, fi21, fi21)
        var v22 int = (^(fi21))
        _ = fi21
        _ = v22
    }
    var v23 int64 = -9223372036854775808
    fmt.Println(v3)
    var v24 int32 = -2147483648
    if false {
        v25 := GenS5{F0: (v3 & 897), F1: 273}
        fmt.Println(v23, v24)
        _ = v25
    } else {
        v2 = float64(v4)
        v26 := GenS2{F0: v24}
        v23 = (v23 | (v23 - (v23 << ((0) & 63))))
        _ = v26
    }
    if (("hello" + "\n\t") <= ("hello" + (("a" + (("a" + (((("a" + (("hello" + "a") + ("a" + "foo bar"))) + ("\n\t" + "hello")) + "foo bar") + "a")) + "a")) + "foo bar"))) {
        v27 := GenS6{F0: v24}
        v24 += -1
        v24 = v24
        _ = v27
    }
    v28 := GenS7{F0: 0}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v28
    _ = v3
    _ = v4
}
