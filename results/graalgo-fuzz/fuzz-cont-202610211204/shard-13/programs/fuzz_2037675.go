package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint32; F1 string }
type GenS4 struct { F0 uint16; F1 uint8 }
type GenS5 struct { F0 int16; F1 int; F2 uint64 }

func main() {
    var v1 int64 = 317
    var v2 float64 = ((0.0 / 0.1) / (3.14159 + (float64((0.1 * 1e10)) - (1.5 - float64(1.5)))))
    var v3 float64 = (v2 * v2)
    var v4 uint16 = 0
    var v5 float32 = 1e10
    v6 := []uint32{1, 0, 1}
    v5 += float32((v2 / v3))
    for fi7 := 0; fi7 < 9; fi7++ {
        v8 := GenS1{F0: (v1 | v1), F1: 2}
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: v1, F1: 255}
    switch 32767 {
    case -1:
        var v10 float32 = v5
        _ = v10
    }
    if (1e10 >= (((3.14159 + v2) * (v2 - 3.14159)) * (((v3 * v2) + (v2 - 0.0)) / 3.14159))) {
        if (18446744073709551615 > 0) {
            v5 += v5
            v11 := GenS1{F0: v1, F1: 76}
            _ = v11
        }
        for fi12 := 0; fi12 < 15; fi12++ {
            var v13 int8 = 66
            switch 259 {
            case 128:
                switch "\n\t" {
                case "":
                    v14 := GenS1{F0: v1, F1: uint8(v4)}
                    _ = v14
                case "a":
                    if false {
                        if (!false) {
                            v5 = 0.0
                            var v15 int32 = 42
                            _ = v15
                        } else {
                            for fi16 := 0; fi16 < 16; fi16++ {
                                v2 -= -2.25
                                if (-1 < int32(v4)) {
                                    v17 := GenS2{F0: 751}
                                    var v18 uint = 541
                                    _ = v17
                                    _ = v18
                                }
                                _ = fi16
                            }
                            fmt.Println(v3, v4, v4)
                            v3 = v3
                        }
                    }
                case "hello":
                    var v19 uint64 = 542
                    _ = v19
                }
            case -639:
                v3 += (v2 / (1e10 + (float64((v5 * (float32(float64(3.14159)) / v5))) + -2.25)))
            }
            var v20 int = (fi12 << ((219) & 63))
            _ = fi12
            _ = v13
            _ = v20
        }
        for fi21 := 0; fi21 < 11; fi21++ {
            v1 -= (v1 | (v1 & (v1 / (((v1 % ((-290) | 1))) | 1))))
            var v22 uint32 = uint32(v4)
            for fi23 := 0; fi23 < 1; fi23++ {
                switch "hello" {
                case "a":
                    var v24 int64 = -1
                    _ = v24
                }
                v5 += v5
                _ = fi23
            }
            _ = fi21
            _ = v22
        }
    } else {
        v25 := GenS3{F0: 1, F1: "foo bar"}
        _ = v25
    }
    var v26 uint16 = v4
    v27 := GenS1{F0: 9223372036854775807, F1: 144}
    fmt.Println(v5, v5, v5)
    if true {
        if false {
            v26 += v26
            v4 -= 255
        } else {
            for fi28 := 0; fi28 < 14; fi28++ {
                switch 127 {
                case -2147483648:
                    for fi29 := 0; fi29 < 18; fi29++ {
                        var v30 float32 = 3.14159
                        _ = fi29
                        _ = v30
                    }
                case 2147483647:
                    switch "foo bar" {
                    case "hello":
                        v26 = 34
                    case "\n\t":
                        if false {
                            if false {
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                            } else {
                                for fi31 := 0; fi31 < 5; fi31++ {
                                    if ((v4 & 997) >= 32767) {
                                        switch (fi31 >> ((4294967295) & 63)) {
                                        case 0:
                                            v32 := GenS4{F0: v26, F1: 0}
                                            _ = v32
                                        case 4294967295:
                                            switch ("" + (("a" + ("hello" + ("\n\t" + ("\n\t" + "a")))) + ("hello" + "foo bar"))) {
                                            case "":
                                                switch "a" {
                                                case "foo bar":
                                                    v26 += v4
                                                case "a":
                                                    if false {
                                                        var v33 uint = 65535
                                                        var v34 uint8 = 13
                                                        v33 = (v33 | 18446744073709551615)
                                                        _ = v33
                                                        _ = v34
                                                    } else {
                                                        v5 -= v5
                                                        _ = len(v6)
                                                    }
                                                case "\n\t":
                                                    v2 -= (float64(1e10) - 3.14159)
                                                }
                                            case "a":
                                                var v35 uint16 = 225
                                                _ = v35
                                            default:
                                                v26 += (v26 + (v26 ^ 32767))
                                            }
                                        }
                                        v36 := GenS5{F0: 32767, F1: fi28, F2: 304}
                                        var v37 float32 = v5
                                        v38 := GenS5{F0: 426, F1: fi28, F2: 732}
                                        if len(v6) > 0 {
                                            v6[0] = 840
                                        }
                                        var v39 uint16 = (v4 - 65535)
                                        v26 += v39
                                        _ = v36
                                        _ = v37
                                        _ = v38
                                        _ = v39
                                    }
                                    v40 := GenS1{F0: v1, F1: 122}
                                    _ = fi31
                                    _ = v40
                                }
                                if true {
                                    v6 = append(v6, 621)
                                }
                            }
                            _ = len(v6)
                        }
                    }
                case 433:
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                default:
                    if len(v6) > 0 {
                        v6[0] = 42
                    }
                }
                _ = fi28
            }
            v5 += (0.0 + v5)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
