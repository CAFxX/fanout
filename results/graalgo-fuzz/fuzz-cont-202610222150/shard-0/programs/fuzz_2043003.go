package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 uint }
type GenS2 struct { F0 int8; F1 int64; F2 uint16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 bool; F1 uint16; F2 uint }
type GenS5 struct { F0 string; F1 int; F2 int32 }
type GenS6 struct { F0 uint8; F1 string }
type GenS7 struct { F0 int; F1 int; F2 uint16 }
type GenS8 struct { F0 uint64; F1 bool; F2 string }

func main() {
    var v1 int = 4294967295
    var v2 int8 = 1
    var v3 uint8 = 255
    v4 := []string{"foo bar", "foo bar", "foo bar", (("" + "foo bar") + "foo bar")}
    var v5 float32 = float32(1.5)
    var v6 float32 = 3.14159
    var v7 float32 = float32(((0.1 * 3.14159) * 1.5))
    v4 = append(v4, "hello")
    var v8 string = ""
    if false {
        if true {
            fmt.Println(v6, v2)
            for fi9 := 0; fi9 < 15; fi9++ {
                for fi10 := 0; fi10 < 16; fi10++ {
                    switch (fi10 ^ fi9) {
                    case 32767:
                        v5 = 0.1
                    case 1:
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    case -29:
                        var v11 uint16 = 65535
                        _ = v11
                    }
                    v6 = (v7 - 1e10)
                    _ = fi10
                }
                if true {
                    v3 = v3
                    v4 = append(v4, ("" + "a"))
                    v12 := GenS1{F0: 4294967295, F1: -1, F2: 1}
                    _ = v12
                } else {
                    fmt.Println(v7)
                }
                fmt.Println(v3, v2)
                _ = fi9
            }
            if len(v4) > 0 {
                v4[0] = "foo bar"
            }
        } else {
            v13 := GenS2{F0: (-(v2)), F1: 48, F2: 251}
            _ = v13
        }
        if false {
            if false {
                v8 = v8
                if (42 == 341) {
                    var v14 float32 = (v7 + (v6 - (v6 * 1.5)))
                    for fi15 := 0; fi15 < 12; fi15++ {
                        if len(v4) > 0 {
                            v4[0] = ("" + v8)
                        }
                        if (!false) {
                            v16 := GenS3{F0: 2}
                            _ = v16
                        } else {
                            for fi17 := 0; fi17 < 5; fi17++ {
                                v7 = (3.14159 * v5)
                                var v18 string = ("" + ((v8 + "hello") + ("a" + v8)))
                                v19 := GenS4{F0: true, F1: 0, F2: 922}
                                _ = fi17
                                _ = v18
                                _ = v19
                            }
                        }
                        if (32767 != 0) {
                            v7 = ((float32((1e10 - (3.14159 + (-2.25 * ((0.1 + 3.14159) - 1e10))))) + 1.5) - 0.0)
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        } else {
                            for fi20 := 0; fi20 < 3; fi20++ {
                                var v21 int = int(v3)
                                v22 := GenS3{F0: 2}
                                var v23 uint = 403
                                _ = fi20
                                _ = v21
                                _ = v22
                                _ = v23
                            }
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            var v24 int32 = 0
                            _ = v24
                        }
                        _ = fi15
                    }
                    _ = v14
                }
            }
            v25 := GenS4{F0: ((3.14159 / 1.0) > (((1e10 - 0.1) * 1e10) * 1.5)), F1: 42, F2: 255}
            _ = v25
        }
        v26 := GenS5{F0: v8, F1: (v1 ^ v1), F2: -1}
        _ = v26
    }
    v27 := GenS1{F0: 421, F1: 10, F2: 1}
    var v28 uint = 1
    v29 := GenS2{F0: v2, F1: 793, F2: 531}
    for fi30 := 0; fi30 < 18; fi30++ {
        if (1 == 139) {
            if false {
                var v31 uint64 = 1
                _ = v31
            }
            v1 -= (fi30 + (fi30 << ((v28) & 63)))
            switch v8 {
            case "a":
                v5 = (v7 + (-2.25 * 3.14159))
            case "\n\t":
                for fi32 := 0; fi32 < 4; fi32++ {
                    if false {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        v33 := GenS6{F0: v3, F1: v8}
                        v34 := GenS5{F0: v8, F1: int(v2), F2: 2147483647}
                        _ = v33
                        _ = v34
                    }
                    v2 -= (^(v2))
                    switch v8 {
                    case "foo bar":
                        v5 -= (float32(-2.25) - float32(1.5))
                    case "hello":
                        var v35 uint32 = 65535
                        _ = v35
                    case "":
                        if true {
                            switch "\n\t" {
                            case "":
                                if len(v4) > 0 {
                                    v4[0] = v8
                                }
                            case "\n\t":
                                _ = len(v4)
                            case "foo bar":
                                v2 = v2
                            }
                        } else {
                            var v36 uint8 = (v3 | (v3 / ((v3) | 1)))
                            var v37 uint16 = 761
                            v38 := GenS7{F0: (fi32 + fi32), F1: fi30, F2: (^(v37))}
                            var v39 uint16 = (^(v37))
                            v3 += v36
                            _ = v36
                            _ = v37
                            _ = v38
                            _ = v39
                        }
                    }
                    _ = fi32
                }
            case "hello":
                switch ("" + (v8 + v8)) {
                case "hello":
                    for fi40 := 0; fi40 < 5; fi40++ {
                        if (1 == 944) {
                            fmt.Println(v1, v7, v3)
                            var v41 int8 = -63
                            v41 += 52
                            switch ((v8 + v8) + ((((v8 + ((v8 + "\n\t") + "hello")) + "") + (v8 + v8)) + "\n\t")) {
                            case "\n\t":
                                switch "" {
                                case "\n\t":
                                    v42 := GenS5{F0: v8, F1: 128, F2: -35}
                                    _ = v42
                                default:
                                    switch v8 {
                                    case "":
                                        fmt.Println(fi30, v3)
                                    }
                                }
                            case "":
                                switch (fi30 ^ fi30) {
                                case 2147483647:
                                    fmt.Println(v1, v41, v41)
                                case -543:
                                    var v43 uint64 = 143
                                    _ = v43
                                case 9223372036854775807:
                                    if len(v4) > 0 {
                                        v4[0] = v8
                                    }
                                default:
                                    var v44 int8 = (v41 * 42)
                                    _ = v44
                                }
                            default:
                                v4 = append(v4, v8)
                            }
                            if (65535 < 0) {
                                if false {
                                    switch ("" + v8) {
                                    case "a":
                                        var v45 uint32 = 42
                                        _ = v45
                                    case "foo bar":
                                        if len(v4) > 0 {
                                            v4[0] = v8
                                        }
                                    case "hello":
                                        v4 = append(v4, (v8 + v8))
                                        v41 = (^(v41))
                                        v46 := GenS8{F0: 79, F1: true, F2: "\n\t"}
                                        var v47 uint = (v28 | v28)
                                        var v48 int8 = v2
                                        _ = v46
                                        _ = v47
                                        _ = v48
                                    default:
                                        var v49 int = (fi30 * int(v28))
                                        _ = v49
                                    }
                                    v50 := GenS1{F0: 0, F1: -32768, F2: 1}
                                    _ = v50
                                } else {
                                    v6 = float32((1.5 / 1.0))
                                }
                                v2 += (v2 | -5)
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            } else {
                                v6 = (v6 + -2.25)
                                var v51 string = v8
                                _ = v51
                            }
                            _ = v41
                        } else {
                            v5 = (0.1 * 1.5)
                        }
                        v4 = append(v4, (v8 + "foo bar"))
                        var v52 int = v1
                        _ = fi40
                        _ = v52
                    }
                case "a":
                    v1 = 9223372036854775807
                }
            }
        }
        var v53 float64 = 0.1
        _ = fi30
        _ = v53
    }
    var v54 uint8 = v3
    if ((^(v28)) < (v28 ^ 42)) {
        for fi55 := 0; fi55 < 2; fi55++ {
            v54 = (^(v3))
            v56 := GenS1{F0: 2, F1: 1, F2: v28}
            v57 := GenS8{F0: 18446744073709551615, F1: ((0 <= 4294967295) || (float32(float64(1.5)) != ((v7 / 1e10) / (v6 / v6)))), F2: v8}
            _ = fi55
            _ = v56
            _ = v57
        }
        fmt.Println(v2, v6)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    } else {
        _ = len(v4)
    }
    for fi58 := 0; fi58 < 9; fi58++ {
        switch fi58 {
        case 127:
            v7 = v5
        }
        var v59 string = v8
        if true {
            for fi60 := 0; fi60 < 2; fi60++ {
                v5 = 1e10
                var v61 uint64 = 185
                _ = len(v4)
                _ = fi60
                _ = v61
            }
            if (0 != 205) {
                var v62 int8 = v2
                v54 = v54
                var v63 uint64 = 4294967295
                _ = v62
                _ = v63
            }
        }
        _ = fi58
        _ = v59
    }
    var v64 int = v1
    v65 := GenS4{F0: false, F1: 42, F2: (v28 * 2)}
    v66 := GenS8{F0: 551, F1: true, F2: v8}
    if true {
        switch ((("\n\t" + ("\n\t" + v8)) + (v8 + v8)) + ("foo bar" + v8)) {
        case "hello":
            fmt.Println(v1)
        case "\n\t":
            _ = len(v4)
        case "":
            var v67 string = "hello"
            _ = v67
        default:
            for fi68 := 0; fi68 < 6; fi68++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi68
            }
        }
    } else {
        v3 += (^(v54))
    }
    fmt.Println(v1)
    fmt.Println(v64, v8)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v6:", v6)
    fmt.Println("v64:", v64)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v6
    _ = v64
    _ = v65
    _ = v66
    _ = v7
    _ = v8
}
