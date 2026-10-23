package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 uint16; F1 int32; F2 string }

func main() {
    v1 := []string{}
    v2 := make(map[int]bool)
    v3 := []uint32{42, 2147483648, 748}
    v4 := []string{}
    var v5 float64 = float64(3.14159)
    v6 := GenS1{F0: "foo bar", F1: 42}
    if len(v1) > 0 {
        v1[0] = "foo bar"
    }
    delete(v2, 60)
    switch "a" {
    case "":
        v5 += v5
    }
    var v7 uint = 324
    switch "" {
    case "":
        if (!(4294967295 >= 0)) {
            switch ("hello" + (("" + (("" + "") + "hello")) + "foo bar")) {
            case "hello":
                if true {
                    v5 += (float64(3.14159) - 1.5)
                    for fi8 := 0; fi8 < 18; fi8++ {
                        var v9 uint8 = 7
                        _ = fi8
                        _ = v9
                    }
                    var v10 uint = v7
                    _ = v10
                }
            case "foo bar":
                for fi11 := 0; fi11 < 11; fi11++ {
                    v12 := GenS1{F0: ("foo bar" + "hello"), F1: -92}
                    _ = fi11
                    _ = v12
                }
            case "a":
                if (false || false) {
                    v5 = v5
                    for fi13 := 0; fi13 < 9; fi13++ {
                        if (2 == 65535) {
                            var v14 int64 = 9223372036854775807
                            v5 -= (-2.25 - (((v5 * (v5 - v5)) + (-2.25 * (v5 * ((3.14159 - v5) / float64(1e10))))) * (v5 - float64(0.1))))
                            for fi15 := 0; fi15 < 20; fi15++ {
                                if ((2 <= 2) || ((!false) && false)) {
                                    switch (fi15 | (fi15 ^ (fi13 ^ fi15))) {
                                    case -913:
                                        v5 = 0.1
                                    case -228:
                                        v16 := GenS1{F0: "foo bar", F1: 0}
                                        _ = v16
                                    default:
                                        var v17 int8 = 0
                                        _ = v17
                                    }
                                    v7 += (v7 ^ v7)
                                    var v18 string = ""
                                    _ = v18
                                }
                                _ = fi15
                            }
                            _ = v14
                        }
                        _ = fi13
                    }
                } else {
                    if (1 <= 127) {
                        v19 := v2[128]
                        _ = len(v2)
                        _ = v19
                    }
                }
            default:
                for fi20 := 0; fi20 < 9; fi20++ {
                    switch 149 {
                    case -2147483648:
                        var v21 uint32 = 0
                        _ = v21
                    case 1:
                        if (486 == 42) {
                            if (true || (127 != -122)) {
                                if (21 != -78) {
                                    switch 237 {
                                    case 127:
                                        v7 = (v7 & (^(v7)))
                                    default:
                                        var v22 int16 = 42
                                        _ = v22
                                    }
                                    for fi23 := 0; fi23 < 2; fi23++ {
                                        if true {
                                            if (42 >= -832) {
                                                v7 = 2
                                            }
                                        }
                                        switch (("hello" + "foo bar") + "\n\t") {
                                        case "a":
                                            switch (fi23 >> ((v7) & 63)) {
                                            case 0:
                                                v7 = (^(v7))
                                            case 32768:
                                                v24 := GenS1{F0: ("\n\t" + "\n\t"), F1: -128}
                                                _ = v24
                                            }
                                        case "":
                                            var v25 uint8 = 42
                                            _ = v25
                                        case "foo bar":
                                            v7 += 196
                                        }
                                        if (fi23 < fi20) {
                                            v5 += v5
                                            fmt.Println(v5, fi20)
                                        }
                                        _ = fi23
                                    }
                                }
                                var v26 bool = (!(!true))
                                switch fi20 {
                                case 32768:
                                    for fi27 := 0; fi27 < 5; fi27++ {
                                        var v28 int8 = 41
                                        if len(v1) > 0 {
                                            v1[0] = ""
                                        }
                                        var v29 int16 = 0
                                        _ = fi27
                                        _ = v28
                                        _ = v29
                                    }
                                case 94:
                                    for fi30 := 0; fi30 < 9; fi30++ {
                                        v31 := GenS1{F0: ("\n\t" + ""), F1: -91}
                                        v5 -= v5
                                        v5 -= 0.0
                                        _ = fi30
                                        _ = v31
                                    }
                                default:
                                    v26 = v26
                                }
                                _ = v26
                            }
                            var v32 uint8 = 143
                            _ = v32
                        }
                    }
                    _ = fi20
                }
            }
        } else {
            fmt.Println(v7)
            var v33 bool = true
            _ = v33
        }
    }
    var v34 uint = (v7 + v7)
    for fi35 := 0; fi35 < 14; fi35++ {
        v36 := GenS2{F0: -833, F1: 65535}
        for fi37 := 0; fi37 < 18; fi37++ {
            v34 += v7
            var v38 int16 = -1
            _ = fi37
            _ = v38
        }
        _ = fi35
        _ = v36
    }
    if ((591 != 253) && true) {
        fmt.Println(v5, v5)
        var v39 int64 = 1
        v40 := GenS3{F0: 77, F1: 1, F2: ""}
        _ = v39
        _ = v40
    }
    var v41 float32 = ((-2.25 * (-2.25 + 0.0)) * 1.5)
    var v42 string = "\n\t"
    v41 -= -2.25
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v6
    _ = v7
}
