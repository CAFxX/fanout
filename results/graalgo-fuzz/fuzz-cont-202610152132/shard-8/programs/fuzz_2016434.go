package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32; F1 bool }
type GenS4 struct { F0 uint16; F1 uint64 }

func main() {
    var v1 uint16 = 65535
    v2 := []string{"\n\t", "hello", ""}
    v3 := []string{"hello"}
    if true {
        for fi4 := 0; fi4 < 13; fi4++ {
            var v5 float64 = 1.5
            _ = fi4
            _ = v5
        }
        if (true && false) {
            v6 := GenS1{F0: 0, F1: (false && (float64((0.1 * 3.14159)) != 0.0)), F2: 171}
            _ = v6
        }
    }
    if (2147483647 <= -1) {
        var v7 uint = 286
        _ = v7
    } else {
        for fi8 := 0; fi8 < 7; fi8++ {
            v1 = v1
            _ = fi8
        }
    }
    v1 = 930
    for fi9 := 0; fi9 < 13; fi9++ {
        v1 -= v1
        for fi10 := 0; fi10 < 4; fi10++ {
            v1 = (v1 / (((v1 << ((42) & 63))) | 1))
            v11 := GenS2{F0: "a"}
            var v12 int8 = 100
            _ = fi10
            _ = v11
            _ = v12
        }
        v1 = v1
        _ = fi9
    }
    v3 = append(v3, "a")
    var v13 int = -9223372036854775808
    switch (("\n\t" + "foo bar") + "hello") {
    case "\n\t":
        for fi14 := 0; fi14 < 5; fi14++ {
            v15 := GenS3{F0: -2147483648, F1: (true && (40 == 113))}
            v16 := GenS1{F0: (v13 / ((-9223372036854775808) | 1)), F1: (!(358 != 65535)), F2: 255}
            fmt.Println(v13)
            _ = fi14
            _ = v15
            _ = v16
        }
    case "foo bar":
        if ((0 <= 0) || true) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            var v17 int = (v13 - (v13 & (v13 / (((-(v13))) | 1))))
            v18 := GenS2{F0: "\n\t"}
            _ = v17
            _ = v18
        } else {
            if (-9223372036854775808 == 512) {
                for fi19 := 0; fi19 < 14; fi19++ {
                    fmt.Println(fi19)
                    _ = fi19
                }
                if len(v2) > 0 {
                    v2[0] = (("foo bar" + "\n\t") + ("" + ("a" + "foo bar")))
                }
                for fi20 := 0; fi20 < 20; fi20++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    switch "foo bar" {
                    case "":
                        var v21 string = "a"
                        _ = v21
                    case "foo bar":
                        var v22 float64 = 3.14159
                        _ = v22
                    case "hello":
                        if (false && false) {
                            if (0 != -62) {
                                _ = len(v3)
                            } else {
                                for fi23 := 0; fi23 < 8; fi23++ {
                                    switch ("foo bar" + (("a" + "") + ("hello" + "\n\t"))) {
                                    case "hello":
                                        fmt.Println(fi23, v1)
                                    default:
                                        var v24 float32 = 0.1
                                        _ = v24
                                    }
                                    fmt.Println(v1, v1)
                                    var v25 float32 = -2.25
                                    _ = fi23
                                    _ = v25
                                }
                                if true {
                                    if false {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        v13 = 32768
                                        v26 := GenS3{F0: 1, F1: false}
                                        _ = v26
                                    }
                                } else {
                                    switch (fi20 | 0) {
                                    case -1:
                                        var v27 uint32 = 924
                                        _ = v27
                                    }
                                    v28 := GenS1{F0: (v13 + 128), F1: true, F2: 42}
                                    switch ("\n\t" + ("\n\t" + "")) {
                                    case "foo bar":
                                        for fi29 := 0; fi29 < 16; fi29++ {
                                            if false {
                                                var v30 float32 = float32(-2.25)
                                                _ = v30
                                            } else {
                                                fmt.Println(fi29)
                                            }
                                            _ = fi29
                                        }
                                    }
                                    _ = v28
                                }
                            }
                            v31 := GenS4{F0: v1, F1: 42}
                            _ = v31
                        }
                    }
                    _ = fi20
                }
            }
            v1 = (v1 ^ v1)
        }
    case "a":
        v1 += v1
    }
    v1 -= (v1 | (v1 & (v1 / (((v1 % ((v1) | 1))) | 1))))
    v1 = v1
    var v32 uint32 = 992
    var v33 int16 = 1
    if (228 > 105) {
        if (!(((!(-9223372036854775808 != 42)) && ((false && false) && (255 <= (v1 & v1)))) || true)) {
            v13 = v13
            switch ("\n\t" + "") {
            case "":
                var v34 uint16 = v1
                _ = v34
            default:
                if len(v3) > 0 {
                    v3[0] = ""
                }
            }
            var v35 uint32 = 0
            _ = v35
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v32
    _ = v33
}
