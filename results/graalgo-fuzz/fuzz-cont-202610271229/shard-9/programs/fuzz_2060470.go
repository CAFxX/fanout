package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 int64; F1 int32; F2 int64 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int16; F1 string; F2 uint8 }

func main() {
    v1 := []uint32{42, 262, 1, 2147483648}
    var v2 int64 = 0
    var v3 float32 = -2.25
    if (1 > 2147483647) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v4 bool = false
        switch 42 {
        case 127:
            if (!v4) {
                v5 := GenS1{F0: 42, F1: 2147483648}
                if (v4 && (!(!(v4 && (true || ((!v4) && true)))))) {
                    v1 = append(v1, 4294967295)
                    var v6 string = ((("\n\t" + "a") + "hello") + "\n\t")
                    v6 = "\n\t"
                    _ = v6
                }
                _ = v5
            }
        case -383:
            var v7 int32 = -1
            _ = v7
        default:
            v4 = (42 > 9223372036854775807)
        }
        _ = v4
    } else {
        v1 = append(v1, 322)
        _ = len(v1)
    }
    for fi8 := 0; fi8 < 12; fi8++ {
        fmt.Println(v3, v2, v3)
        _ = fi8
    }
    for fi9 := 0; fi9 < 6; fi9++ {
        if true {
            switch "" {
            case "":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case "\n\t":
                for fi10 := 0; fi10 < 7; fi10++ {
                    _ = len(v1)
                    _ = len(v1)
                    v3 = v3
                    _ = fi10
                }
            default:
                var v11 uint8 = 127
                _ = v11
            }
        } else {
            switch "\n\t" {
            case "foo bar":
                var v12 int16 = 0
                _ = v12
            case "a":
                var v13 uint8 = 42
                _ = v13
            case "\n\t":
                for fi14 := 0; fi14 < 14; fi14++ {
                    var v15 int64 = v2
                    v16 := GenS2{F0: 1, F1: 1}
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            default:
                for fi17 := 0; fi17 < 18; fi17++ {
                    var v18 uint32 = 760
                    fmt.Println(v3, v3)
                    var v19 uint16 = 83
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v20 := GenS3{F0: (("hello" + "a") + "hello"), F1: 0}
            _ = v20
        }
        var v21 string = ("" + (("foo bar" + "\n\t") + "a"))
        _ = fi9
        _ = v21
    }
    for fi22 := 0; fi22 < 1; fi22++ {
        switch "hello" {
        case "hello":
            var v23 uint16 = 255
            _ = v23
        case "foo bar":
            if true {
                var v24 float64 = 0.0
                _ = v24
            } else {
                _ = len(v1)
            }
        case "a":
            for fi25 := 0; fi25 < 8; fi25++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v26 int16 = 1000
                _ = fi25
                _ = v26
            }
        default:
            fmt.Println(v3, v2, v2)
        }
        v1 = append(v1, 2)
        _ = fi22
    }
    switch ("foo bar" + ("a" + "")) {
    case "":
        var v27 uint16 = 0
        _ = v27
    case "hello":
        var v28 uint = 270
        _ = v28
    case "\n\t":
        for fi29 := 0; fi29 < 7; fi29++ {
            if true {
                v30 := GenS3{F0: "hello", F1: 85}
                var v31 float64 = 0.0
                _ = v30
                _ = v31
            } else {
                v32 := GenS1{F0: 42, F1: 2147483647}
                v33 := GenS1{F0: 1, F1: 227}
                _ = v32
                _ = v33
            }
            if len(v1) > 0 {
                v1[0] = 42
            }
            _ = fi29
        }
    default:
        fmt.Println(v2, v3)
    }
    fmt.Println(v2)
    v34 := GenS3{F0: ((("hello" + "") + "hello") + ("hello" + "")), F1: 65535}
    v35 := GenS1{F0: 83, F1: 2}
    switch 965 {
    case 128:
        fmt.Println(v2, v3)
    default:
        v2 -= v2
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch "a" {
    case "hello":
        if (!false) {
            if len(v1) > 0 {
                v1[0] = 2
            }
            v3 += (v3 * 1e10)
            if (42 != 65535) {
                v3 += v3
            }
        } else {
            for fi36 := 0; fi36 < 9; fi36++ {
                var v37 int32 = 0
                for fi38 := 0; fi38 < 20; fi38++ {
                    v1 = append(v1, 0)
                    var v39 uint32 = 4294967295
                    _ = fi38
                    _ = v39
                }
                for fi40 := 0; fi40 < 12; fi40++ {
                    if (fi40 < 32768) {
                        v2 = v2
                    }
                    if len(v1) > 0 {
                        v1[0] = 42
                    }
                    if ((89 >= 197) || true) {
                        if false {
                            v41 := GenS4{F0: (v2 ^ v2), F1: 706, F2: -283}
                            var v42 uint8 = 1
                            _ = v41
                            _ = v42
                        } else {
                            v37 = (v37 / (((v37 + (v37 >> ((4294967295) & 63)))) | 1))
                            v43 := GenS2{F0: -1, F1: 32767}
                            _ = v43
                        }
                    }
                    if false {
                        var v44 string = ""
                        v45 := GenS5{F0: 255}
                        if false {
                            var v46 uint64 = 4294967295
                            var v47 uint32 = 65535
                            _ = v46
                            _ = v47
                        }
                        _ = v44
                        _ = v45
                    } else {
                        if (false || (968 != 42)) {
                            if (0 > int16(v37)) {
                                var v48 float32 = (v3 * v3)
                                switch "\n\t" {
                                case "":
                                    switch (-(fi36)) {
                                    case 9223372036854775807:
                                        v3 = 1e10
                                    case 32767:
                                        if false {
                                            if len(v1) > 0 {
                                                v1[0] = 2147483648
                                            }
                                            if false {
                                                var v49 int8 = -128
                                                if (fi36 >= fi36) {
                                                    v37 = v37
                                                    v48 = v3
                                                }
                                                v49 -= (v49 / ((100) | 1))
                                                v48 = v48
                                                _ = v49
                                            } else {
                                                switch fi40 {
                                                case 45:
                                                    fmt.Println(v3)
                                                case -9223372036854775808:
                                                    v50 := GenS3{F0: "\n\t", F1: 0}
                                                    if false {
                                                        v48 = 0.0
                                                        v37 = v37
                                                    } else {
                                                        var v51 float64 = 0.1
                                                        _ = v51
                                                    }
                                                    fmt.Println(v2)
                                                    _ = v50
                                                default:
                                                    fmt.Println(fi36, v48, v37)
                                                }
                                            }
                                        }
                                    }
                                default:
                                    var v52 int32 = -2147483648
                                    _ = v52
                                }
                                v53 := GenS4{F0: (-(v2)), F1: v37, F2: (v2 % ((v2) | 1))}
                                _ = v48
                                _ = v53
                            }
                        } else {
                            var v54 int64 = (^(v2))
                            var v55 int64 = 1
                            var v56 int32 = -91
                            _ = v54
                            _ = v55
                            _ = v56
                        }
                        v2 -= (v2 & -9223372036854775808)
                        v57 := GenS5{F0: 1}
                        _ = v57
                    }
                    var v58 uint = 18446744073709551615
                    fmt.Println(v2, v58)
                    var v59 uint16 = 32767
                    _ = fi40
                    _ = v58
                    _ = v59
                }
                _ = fi36
                _ = v37
            }
            v2 += v2
        }
    case "":
        if (false && (319 > 1000)) {
            switch -1 {
            case 127:
                if true {
                    var v60 string = "foo bar"
                    var v61 float64 = 1e10
                    _ = v60
                    _ = v61
                } else {
                    var v62 int = 0
                    if false {
                        v3 = ((v3 - 0.1) / v3)
                        for fi63 := 0; fi63 < 14; fi63++ {
                            switch (v62 >> ((829) & 63)) {
                            case -229:
                                v3 = v3
                            default:
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                            v64 := GenS4{F0: (v2 - 199), F1: 1, F2: (v2 >> ((0) & 63))}
                            _ = fi63
                            _ = v64
                        }
                    } else {
                        v65 := GenS5{F0: 2}
                        for fi66 := 0; fi66 < 13; fi66++ {
                            for fi67 := 0; fi67 < 15; fi67++ {
                                if false {
                                    switch "" {
                                    case "hello":
                                        v62 = (fi66 ^ (fi67 & -9223372036854775808))
                                    case "foo bar":
                                        fmt.Println(fi66, v3)
                                    default:
                                        switch "a" {
                                        case "a":
                                            v62 = fi67
                                        case "hello":
                                            var v68 uint16 = 32767
                                            _ = v68
                                        case "":
                                            var v69 uint32 = 243
                                            _ = v69
                                        default:
                                            var v70 bool = (!((-121 < 0) || true))
                                            _ = v70
                                        }
                                    }
                                }
                                var v71 uint16 = 1
                                var v72 uint32 = 2147483648
                                _ = fi67
                                _ = v71
                                _ = v72
                            }
                            _ = fi66
                        }
                        if false {
                            switch 1 {
                            case -1:
                                v62 = 127
                            case 4294967295:
                                v73 := GenS5{F0: 128}
                                _ = v73
                            default:
                                fmt.Println(v2)
                            }
                            _ = len(v1)
                        }
                        _ = v65
                    }
                    _ = v62
                }
            default:
                v74 := GenS6{F0: -1, F1: "", F2: 2}
                _ = v74
            }
            v75 := GenS3{F0: (("" + "a") + ("foo bar" + "foo bar")), F1: 252}
            v2 += v2
            _ = v75
        }
    case "foo bar":
        var v76 uint8 = 103
        _ = v76
    default:
        fmt.Println(v3, v2)
    }
    if len(v1) > 0 {
        v1[0] = 2
    }
    if true {
        var v77 int = 2147483647
        var v78 float32 = float32(1e10)
        v78 = -2.25
        _ = v77
        _ = v78
    }
    v3 = ((v3 * 1.5) / (float32(0.1) - v3))
    switch "hello" {
    case "hello":
        v3 += v3
    case "foo bar":
        switch "a" {
        case "foo bar":
            v2 = v2
        }
    default:
        v2 = v2
    }
    v3 += 0.1
    if (false && true) {
        v79 := GenS3{F0: ("foo bar" + ("" + "")), F1: 65535}
        v1 = append(v1, 2)
        _ = v79
    }
    var v80 uint64 = 176
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v80
}
