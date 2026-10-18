package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 uint32; F1 uint8; F2 int32 }
type GenS3 struct { F0 uint16; F1 int64; F2 uint64 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int; F1 int; F2 int }
type GenS6 struct { F0 int16; F1 uint16; F2 uint64 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 bool; F1 uint; F2 uint64 }
type GenS9 struct { F0 int8; F1 int16; F2 uint16 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 uint; F1 string; F2 uint64 }
type GenS12 struct { F0 int32; F1 string; F2 uint64 }
type GenS13 struct { F0 string; F1 int32; F2 uint8 }
type GenS14 struct { F0 int8; F1 uint32; F2 int64 }
type GenS15 struct { F0 uint32; F1 int32 }
func genFn0() bool {
    var v5 int = 32767
    if true {
        v5 += 32767
    } else {
        var v6 uint16 = 255
        for fi7 := 0; fi7 < 13; fi7++ {
            v6 += v6
            v6 -= (^(v6))
            v8 := []int{42, (fi7 % (((fi7 & (-(fi7)))) | 1)), (v5 % (((v5 - fi7)) | 1)), fi7}
            _ = fi7
            _ = v8
        }
        var v9 uint16 = 208
        _ = v6
        _ = v9
    }
    _ = v5
    return (4294967295 != 18446744073709551615)
}

func main() {
    var v1 uint32 = 1
    v2 := []string{((("foo bar" + "a") + ("a" + "foo bar")) + "foo bar"), "foo bar", "\n\t", "a"}
    var v3 uint16 = 45
    var v4 uint = 18446744073709551615
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if (-721 == 44) {
        if (2147483648 != 1) {
            if false {
                if false {
                    v1 = 65535
                    v3 -= v3
                } else {
                    switch "\n\t" {
                    case "":
                        fmt.Println(v4)
                    case "a":
                        var v10 int64 = -145
                        _ = v10
                    case "hello":
                        fmt.Println(v4, v3)
                    }
                    fmt.Println(v1, v4)
                }
                v2 = append(v2, ("foo bar" + "a"))
                var v11 uint = (v4 / ((uint(v3)) | 1))
                _ = v11
            } else {
                switch "\n\t" {
                case "":
                    v1 += (v1 - v1)
                }
                switch -9223372036854775808 {
                case 127:
                    fmt.Println(v1, v1, v1)
                case 406:
                    v12 := GenS1{F0: -9223372036854775808, F1: v3}
                    _ = v12
                default:
                    v1 += v1
                }
                fmt.Println(v1, v3, v4)
            }
            for fi13 := 0; fi13 < 5; fi13++ {
                if (!true) {
                    if (true && false) {
                        v3 -= v3
                        _ = len(v2)
                    } else {
                        for fi14 := 0; fi14 < 8; fi14++ {
                            var v15 int = fi14
                            fmt.Println(v1, v1)
                            var v16 uint8 = 0
                            if len(v2) > 0 {
                                v2[0] = "\n\t"
                            }
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = fi14
                            _ = v15
                            _ = v16
                        }
                    }
                    v4 -= v4
                    v3 = 32767
                }
                _ = fi13
            }
            switch "a" {
            case "a":
                for fi17 := 0; fi17 < 16; fi17++ {
                    if true {
                        for fi18 := 0; fi18 < 14; fi18++ {
                            v19 := GenS1{F0: (fi17 + fi17), F1: 32767}
                            _ = fi18
                            _ = v19
                        }
                    }
                    v1 = 0
                    for fi20 := 0; fi20 < 17; fi20++ {
                        var v21 uint64 = 42
                        v22 := GenS1{F0: 127, F1: (^(v3))}
                        v3 += (v3 / ((51) | 1))
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                    _ = fi17
                }
            case "hello":
                fmt.Println(v3, v1, v1)
            default:
                if true {
                    if false {
                        if true {
                            v2 = append(v2, ("a" + "foo bar"))
                            for fi23 := 0; fi23 < 5; fi23++ {
                                var v24 uint = v4
                                for fi25 := 0; fi25 < 15; fi25++ {
                                    if true {
                                        if (-49 < -49) {
                                            var v26 string = "\n\t"
                                            if (v26 < v26) {
                                                if true {
                                                    var v27 string = ""
                                                    if ((fi25 | 42) != fi23) {
                                                        v27 = (v27 + v27)
                                                        var v28 bool = false
                                                        _ = v28
                                                    }
                                                    _ = v27
                                                }
                                            }
                                            _ = v26
                                        }
                                        if true {
                                            if (!false) {
                                                var v29 int8 = -1
                                                _ = v29
                                            } else {
                                                v1 += v1
                                            }
                                            if false {
                                                if true {
                                                    var v30 uint8 = 123
                                                    _ = v30
                                                } else {
                                                    var v31 int8 = 127
                                                    v32 := GenS2{F0: 2147483648, F1: 0, F2: 0}
                                                    if true {
                                                        if true {
                                                            switch "" {
                                                            case "":
                                                                if false {
                                                                    v24 -= v4
                                                                    _ = len(v2)
                                                                } else {
                                                                    v33 := GenS3{F0: (v3 & (^(v3))), F1: -1, F2: 1}
                                                                    _ = len(v2)
                                                                    var v34 int8 = (v31 + (v31 | (v31 * 127)))
                                                                    _ = v33
                                                                    _ = v34
                                                                }
                                                            case "foo bar":
                                                                var v35 uint = 424
                                                                v4 += 712
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                                _ = v35
                                                            }
                                                            v2 = append(v2, "hello")
                                                            var v36 int64 = -953
                                                            _ = v36
                                                        }
                                                        var v37 int32 = 586
                                                        _ = v37
                                                    } else {
                                                        if ("foo bar" <= "\n\t") {
                                                            var v38 uint = 0
                                                            _ = v38
                                                        }
                                                    }
                                                    _ = v31
                                                    _ = v32
                                                }
                                            } else {
                                                if ("a" > (("a" + (("a" + "foo bar") + "hello")) + "")) {
                                                    _ = len(v2)
                                                    var v39 uint16 = v3
                                                    _ = v39
                                                } else {
                                                    v1 -= v1
                                                    v4 += (^(v24))
                                                }
                                                v40 := GenS4{F0: -2147483648}
                                                if true {
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        if true {
                                                            var v41 float64 = (3.14159 - 3.14159)
                                                            v1 += v1
                                                            var v42 int64 = 886
                                                            _ = v41
                                                            _ = v42
                                                        }
                                                    case "a":
                                                        v43 := GenS2{F0: 2147483648, F1: 190, F2: -599}
                                                        _ = v43
                                                    }
                                                } else {
                                                    fmt.Println(fi25, v24)
                                                    if false {
                                                        v2 = append(v2, "")
                                                        var v44 float64 = (0.1 - 3.14159)
                                                        if (false && (fi23 <= -1)) {
                                                            switch fi25 {
                                                            case 1:
                                                                v24 = v24
                                                            case 128:
                                                                v44 = float64(-2.25)
                                                            case -302:
                                                                var v45 float32 = 3.14159
                                                                _ = v45
                                                            default:
                                                                v2 = append(v2, (("hello" + "foo bar") + ("foo bar" + "foo bar")))
                                                            }
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                            v24 += v24
                                                        } else {
                                                            if false {
                                                                v46 := GenS5{F0: 57, F1: 32767, F2: (-(fi23))}
                                                                _ = v46
                                                            }
                                                            var v47 string = ""
                                                            switch (v47 + (v47 + v47)) {
                                                            case "\n\t":
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                            case "a":
                                                                v48 := GenS6{F0: 0, F1: 446, F2: 2}
                                                                _ = v48
                                                            }
                                                            _ = v47
                                                        }
                                                        _ = len(v2)
                                                        v3 -= (v3 * (v3 & (v3 ^ v3)))
                                                        _ = v44
                                                    }
                                                    var v49 string = "\n\t"
                                                    var v50 uint8 = uint8(v24)
                                                    _ = v49
                                                    _ = v50
                                                }
                                                if len(v2) > 0 {
                                                    v2[0] = "\n\t"
                                                }
                                                _ = v40
                                            }
                                            var v51 int16 = 191
                                            _ = v51
                                        } else {
                                            fmt.Println(v24, v1, v24)
                                            var v52 uint = v24
                                            _ = v52
                                        }
                                        var v53 float32 = (1e10 / 1.5)
                                        _ = v53
                                    } else {
                                        if len(v2) > 0 {
                                            v2[0] = "a"
                                        }
                                        v54 := GenS6{F0: 42, F1: 32767, F2: 0}
                                        v3 -= (v3 * (^(v3)))
                                        _ = v54
                                    }
                                    if ((!((true && false) && (("foo bar" + "foo bar") == "foo bar"))) && true) {
                                        if len(v2) > 0 {
                                            v2[0] = ("a" + "foo bar")
                                        }
                                        var v55 int64 = -431
                                        fmt.Println(v55, v24)
                                        var v56 int32 = 204
                                        fmt.Println(v3, v56, fi23)
                                        var v57 string = ("a" + "hello")
                                        v24 = v4
                                        var v58 string = v57
                                        _ = v55
                                        _ = v56
                                        _ = v57
                                        _ = v58
                                    }
                                    _ = fi25
                                }
                                v59 := GenS7{F0: (v3 * v3)}
                                _ = fi23
                                _ = v24
                                _ = v59
                            }
                        } else {
                            if (true || false) {
                                if (-32768 == 1) {
                                    v3 = v3
                                    var v60 float64 = 3.14159
                                    _ = v60
                                } else {
                                    for fi61 := 0; fi61 < 3; fi61++ {
                                        v62 := GenS8{F0: (!false), F1: (^(v4)), F2: 956}
                                        _ = fi61
                                        _ = v62
                                    }
                                    switch "foo bar" {
                                    case "a":
                                        var v63 bool = false
                                        _ = v63
                                    }
                                    switch (("\n\t" + ("\n\t" + "")) + (((((("foo bar" + "hello") + "") + "a") + "foo bar") + "hello") + ("\n\t" + "hello"))) {
                                    case "a":
                                        var v64 int64 = -9223372036854775808
                                        _ = v64
                                    case "":
                                        v65 := GenS4{F0: -128}
                                        _ = v65
                                    case "hello":
                                        for fi66 := 0; fi66 < 19; fi66++ {
                                            var v67 float32 = 3.14159
                                            fmt.Println(v3, fi66, v4)
                                            v4 += (v4 | v4)
                                            _ = fi66
                                            _ = v67
                                        }
                                    default:
                                        v68 := GenS9{F0: 100, F1: 42, F2: v3}
                                        _ = v68
                                    }
                                }
                                v3 = (v3 + (v3 & 73))
                                var v69 int8 = -128
                                _ = v69
                            }
                        }
                        switch "" {
                        case "":
                            for fi70 := 0; fi70 < 19; fi70++ {
                                var v71 uint16 = v3
                                _ = len(v2)
                                v72 := GenS2{F0: v1, F1: 42, F2: 2147483647}
                                _ = fi70
                                _ = v71
                                _ = v72
                            }
                        case "a":
                            v3 += v3
                        case "foo bar":
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        default:
                            switch ("\n\t" + ("" + "\n\t")) {
                            case "\n\t":
                                for fi73 := 0; fi73 < 4; fi73++ {
                                    v1 = 251
                                    _ = fi73
                                }
                            case "hello":
                                if true {
                                    if ((true && true) || (1 >= -489)) {
                                        switch 147 {
                                        case 146:
                                            v1 = (v1 * 2)
                                        case 632:
                                            if (!true) {
                                                if (1e10 < 3.14159) {
                                                    fmt.Println(v4)
                                                }
                                                if (((v1 - (^(v1))) == (v1 + v1)) && true) {
                                                    fmt.Println(v4, v3, v3)
                                                }
                                            } else {
                                                v3 += v3
                                                var v74 string = "\n\t"
                                                var v75 float32 = (1.5 * 0.0)
                                                _ = v74
                                                _ = v75
                                            }
                                        case -9223372036854775808:
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        default:
                                            fmt.Println(v3, v4, v3)
                                        }
                                        fmt.Println(v3, v3)
                                    }
                                    for fi76 := 0; fi76 < 1; fi76++ {
                                        if ((^(v1)) <= (v1 * v1)) {
                                            v77 := GenS10{F0: v1}
                                            fmt.Println(v4, v4)
                                            _ = v77
                                        } else {
                                            fmt.Println(v3)
                                            if false {
                                                var v78 bool = false
                                                var v79 int8 = -12
                                                _ = v78
                                                _ = v79
                                            } else {
                                                fmt.Println(v1)
                                                v80 := GenS11{F0: v4, F1: ("a" + "hello"), F2: 0}
                                                if (!false) {
                                                    var v81 uint16 = v3
                                                    _ = v81
                                                }
                                                _ = v80
                                            }
                                            var v82 int64 = -1
                                            _ = v82
                                        }
                                        v3 = (v3 ^ (^(v3)))
                                        _ = fi76
                                    }
                                    v83 := GenS12{F0: 0, F1: ("foo bar" + ""), F2: 430}
                                    _ = v83
                                }
                            }
                        }
                        switch "foo bar" {
                        case "":
                            var v84 uint8 = 244
                            _ = v84
                        case "hello":
                            if len(v2) > 0 {
                                v2[0] = ""
                            }
                        default:
                            var v85 float32 = 0.1
                            _ = v85
                        }
                    }
                    v86 := GenS13{F0: "foo bar", F1: -2147483648, F2: 79}
                    _ = v86
                }
            }
        } else {
            if (9223372036854775807 > -9223372036854775808) {
                for fi87 := 0; fi87 < 14; fi87++ {
                    var v88 float32 = 0.0
                    _ = fi87
                    _ = v88
                }
            } else {
                if ("hello" >= "") {
                    for fi89 := 0; fi89 < 11; fi89++ {
                        v90 := GenS14{F0: -128, F1: 42, F2: 9223372036854775807}
                        _ = fi89
                        _ = v90
                    }
                    fmt.Println(v1, v4)
                }
                switch 127 {
                case -777:
                    fmt.Println(v4)
                }
                v4 = 1
            }
        }
        fmt.Println(v3, v1, v3)
        for fi91 := 0; fi91 < 17; fi91++ {
            v92 := GenS3{F0: v3, F1: -343, F2: 2}
            v4 += 551
            _ = fi91
            _ = v92
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v93 := GenS12{F0: 1, F1: "hello", F2: 2}
    for fi94 := 0; fi94 < 11; fi94++ {
        v3 = uint16(v4)
        if false {
            v95 := GenS15{F0: v1, F1: 2147483647}
            _ = v95
        } else {
            fmt.Println(v1)
            var v96 int16 = 1000
            v4 = 0
            _ = v96
        }
        var v97 uint32 = (v1 & (^(v1)))
        _ = fi94
        _ = v97
    }
    if (2 <= 127) {
        v4 += v4
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v2 = append(v2, "hello")
    if false {
        v4 = v4
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v93
}
