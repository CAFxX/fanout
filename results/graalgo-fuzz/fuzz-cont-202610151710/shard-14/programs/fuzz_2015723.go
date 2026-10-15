package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 int; F1 uint64; F2 int32 }
type GenS3 struct { F0 int8; F1 int64; F2 int32 }
type GenS4 struct { F0 uint16; F1 uint; F2 int8 }
type GenS5 struct { F0 int; F1 uint32 }
type GenS6 struct { F0 int8; F1 int16; F2 int16 }
func genFn0(p0 int8) uint16 {
    if (42 > 18446744073709551615) {
        var v7 uint8 = 255
        var v8 uint32 = 0
        _ = v7
        _ = v8
    } else {
        v9 := []int{-2147483648, 893, 127}
        fmt.Println(p0)
        _ = v9
    }
    var v10 uint = 2
    for fi11 := 0; fi11 < 7; fi11++ {
        var v12 uint64 = 2
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 5; fi13++ {
        switch (^(fi13)) {
        case -2147483648:
            if (0 != 9223372036854775807) {
                if ((^(v10)) != (v10 % (((v10 * 42)) | 1))) {
                    var v14 uint32 = 987
                    v14 = v14
                    if true {
                        v15 := []int{(fi13 & fi13), fi13}
                        if (((true || (((fi13 - (^(fi13))) > -1) && (float64((((((1e10 - 1e10) - (1.5 + (((0.0 + ((float32(((3.14159 + (1e10 * (0.1 + 0.1))) - 3.14159)) + 0.0) - 1.5)) * 1.5) / 1.0))) + ((-2.25 + 3.14159) - 3.14159)) * (1e10 - -2.25)) / -2.25)) > (-2.25 * 0.1)))) || false) && (!false)) {
                            v10 = (v10 << ((v10) & 63))
                        } else {
                            if len(v15) > 0 {
                                v15[0] = fi13
                            }
                        }
                        fmt.Println(fi13, v14)
                        _ = v15
                    }
                    _ = v14
                }
                if (float32((1e10 - 0.0)) != 3.14159) {
                    if false {
                        var v16 int16 = 1
                        v10 -= (v10 ^ (v10 & 4294967295))
                        v17 := []int{128, -9223372036854775808, int(p0), (fi13 ^ 42)}
                        _ = v16
                        _ = v17
                    } else {
                        var v18 uint8 = 0
                        _ = v18
                    }
                }
                if true {
                    switch (("\n\t" + ((("a" + ("foo bar" + "")) + (("hello" + "foo bar") + "a")) + (("a" + "hello") + (("\n\t" + ((("\n\t" + "") + "hello") + "a")) + "")))) + ("\n\t" + "a")) {
                    case "\n\t":
                        var v19 uint64 = 2
                        _ = v19
                    default:
                        var v20 float64 = (float64(-2.25) / (float64(0.1) - ((1e10 - (3.14159 * (1.5 * -2.25))) - -2.25)))
                        _ = v20
                    }
                    v10 += (v10 % ((1) | 1))
                } else {
                    var v21 float32 = -2.25
                    _ = v21
                }
            } else {
                if true {
                    v22 := []int{}
                    p0 -= p0
                    _ = v22
                }
            }
        case -1:
            fmt.Println(v10, fi13)
        default:
            v10 = (v10 & v10)
        }
        _ = fi13
    }
    fmt.Println(v10)
    _ = p0
    _ = v10
    return 42
}
func genFn1(p0 int32, p1 uint16) bool {
    v23 := []int{42}
    var v24 int64 = 0
    var v25 float64 = 0.1
    fmt.Println(p1, p0, v24)
    switch ("hello" + ("foo bar" + "a")) {
    case "a":
        for fi26 := 0; fi26 < 10; fi26++ {
            if true {
                v27 := GenS1{F0: 183, F1: (("a" + "\n\t") + "\n\t")}
                switch (fi26 * fi26) {
                case 4294967295:
                    p0 += (p0 >> ((255) & 63))
                case -264:
                    var v28 int16 = 497
                    _ = v28
                default:
                    p1 += 0
                }
                _ = v27
            } else {
                switch ("" + "") {
                case "":
                    p1 = (^(p1))
                case "foo bar":
                    for fi29 := 0; fi29 < 6; fi29++ {
                        switch "\n\t" {
                        case "":
                            var v30 int16 = 1
                            _ = v30
                        case "hello":
                            if true {
                                p0 -= 42
                            }
                        }
                        _ = fi29
                    }
                case "hello":
                    for fi31 := 0; fi31 < 1; fi31++ {
                        p0 = p0
                        if (4294967295 == 2147483648) {
                            _ = len(v23)
                            v32 := GenS1{F0: 127, F1: (("hello" + "a") + ("foo bar" + "\n\t"))}
                            var v33 int32 = (p0 - -1)
                            _ = v32
                            _ = v33
                        }
                        _ = fi31
                    }
                }
            }
            _ = fi26
        }
    default:
        v25 = float64(((1.5 + 0.0) + (0.0 * 1.5)))
    }
    _ = p0
    _ = p1
    _ = v23
    _ = v24
    _ = v25
    return false
}

func main() {
    var v1 int8 = 1
    var v2 int32 = -2147483648
    var v3 uint16 = 32767
    var v4 int16 = 1000
    v5 := make(map[string]string)
    v6 := make(map[string]string)
    var v34 int32 = (v2 + v2)
    v34 = (v2 * (v34 & -2147483648))
    if false {
        var v35 int32 = -433
        v36 := []int{-395, -200, 42, -1}
        _ = v35
        _ = v36
    }
    v2 += v34
    v37 := GenS2{F0: 886, F1: 42, F2: 0}
    v38 := GenS1{F0: 137, F1: ("hello" + ("a" + "foo bar"))}
    if false {
        v39 := []int{}
        _ = v39
    } else {
        var v40 uint = 568
        for fi41 := 0; fi41 < 12; fi41++ {
            v42 := GenS3{F0: v1, F1: 9223372036854775807, F2: (v2 << ((65535) & 63))}
            if true {
                v40 = v40
                v43 := GenS2{F0: fi41, F1: 538, F2: 0}
                v44 := GenS3{F0: (v1 % (((v1 >> (((v40 / ((uint(v3)) | 1))) & 63))) | 1)), F1: 9223372036854775807, F2: v2}
                _ = v43
                _ = v44
            } else {
                var v45 int16 = (v4 << (((v40 & v40)) & 63))
                v46 := GenS2{F0: fi41, F1: 202, F2: 42}
                _ = v45
                _ = v46
            }
            v47 := GenS2{F0: (fi41 + fi41), F1: 18446744073709551615, F2: 0}
            _ = fi41
            _ = v42
            _ = v47
        }
        for fi48 := 0; fi48 < 3; fi48++ {
            if false {
                var v49 int16 = int16(v2)
                _ = v49
            }
            v4 += 32767
            v6[((((("\n\t" + "") + "a") + "") + "hello") + "hello")] = "foo bar"
            _ = fi48
        }
        _ = v40
    }
    fmt.Println(v1)
    var v50 uint16 = (v3 + (v3 / ((v3) | 1)))
    v34 += v2
    fmt.Println(v50, v1)
    v51 := []int{}
    if false {
        fmt.Println(v50, v50)
    } else {
        var v52 uint16 = genFn0((v1 * v1))
        if false {
            v52 = 42
        } else {
            fmt.Println(v34)
            fmt.Println(v52)
            if ((v2 - v2) >= -802) {
                var v53 int8 = v1
                if true {
                    v50 -= genFn0((v1 & (v1 * 42)))
                } else {
                    var v54 uint32 = 328
                    switch -9223372036854775808 {
                    case 127:
                        switch "hello" {
                        case "hello":
                            if false {
                                v2 = -370
                                v52 -= (^(v3))
                                v54 = v54
                            } else {
                                var v55 int = 9223372036854775807
                                v56 := GenS4{F0: genFn0(v53), F1: 42, F2: v53}
                                _ = v55
                                _ = v56
                            }
                        case "\n\t":
                            if true {
                                for fi57 := 0; fi57 < 3; fi57++ {
                                    delete(v5, "a")
                                    _ = fi57
                                }
                            } else {
                                var v58 uint32 = (^(v54))
                                var v59 int16 = (v4 * v4)
                                v52 += 65535
                                _ = v58
                                _ = v59
                            }
                        }
                    case -573:
                        if true {
                            switch -493 {
                            case 423:
                                v60 := GenS3{F0: v1, F1: 1, F2: (v34 + int32(v4))}
                                _ = v60
                            }
                        }
                    case -1:
                        for fi61 := 0; fi61 < 3; fi61++ {
                            if (((18446744073709551615 <= 348) && (1.5 == float64(-2.25))) || (true && false)) {
                                var v62 uint = 255
                                v63 := GenS5{F0: (fi61 << ((0) & 63)), F1: 65535}
                                _ = v62
                                _ = v63
                            } else {
                                if false {
                                    fmt.Println(v4, v2)
                                } else {
                                    var v64 uint64 = 4294967295
                                    v65 := GenS6{F0: (v1 * (v1 ^ 42)), F1: (v4 | (v4 % (((v4 / (((v4 >> ((65535) & 63))) | 1))) | 1))), F2: v4}
                                    if false {
                                        if true {
                                            for fi66 := 0; fi66 < 14; fi66++ {
                                                switch (("a" + "foo bar") + "a") {
                                                case "a":
                                                    var v67 float32 = 0.0
                                                    _ = v67
                                                case "\n\t":
                                                    if true {
                                                        if (v64 == 4294967295) {
                                                            v1 = v53
                                                            if true {
                                                                var v68 int16 = (-(v4))
                                                                _ = v68
                                                            }
                                                            fmt.Println(v53)
                                                        } else {
                                                            v69 := GenS5{F0: fi66, F1: (v54 << ((18446744073709551615) & 63))}
                                                            var v70 float32 = ((((0.0 + -2.25) * (3.14159 + 3.14159)) / (0.0 - ((3.14159 / 1.0) + 0.0))) + 1e10)
                                                            v53 = (v1 % ((v1) | 1))
                                                            _ = v69
                                                            _ = v70
                                                        }
                                                        v4 = v4
                                                    }
                                                default:
                                                    var v71 int16 = v4
                                                    _ = v71
                                                }
                                                fmt.Println(v64, fi66, v53)
                                                _ = fi66
                                            }
                                        }
                                        if (0.1 != (0.1 - 1.5)) {
                                            fmt.Println(v34, v34)
                                            fmt.Println(v54, fi61)
                                        } else {
                                            switch "foo bar" {
                                            case "foo bar":
                                                for fi72 := 0; fi72 < 16; fi72++ {
                                                    _ = len(v6)
                                                    delete(v5, "\n\t")
                                                    _ = fi72
                                                }
                                            case "\n\t":
                                                v53 += v53
                                            case "hello":
                                                fmt.Println(fi61, v4)
                                            }
                                            if true {
                                                fmt.Println(v4, v2, v1)
                                                var v73 float64 = (3.14159 / 0.1)
                                                _ = v73
                                            }
                                            fmt.Println(v34, v52, v64)
                                        }
                                    } else {
                                        switch (fi61 | (fi61 + fi61)) {
                                        case -267:
                                            var v74 float64 = float64((-2.25 + 0.1))
                                            _ = v74
                                        case 2147483647:
                                            var v75 string = "\n\t"
                                            _ = v75
                                        case 32768:
                                            v50 = (v50 ^ v3)
                                        default:
                                            switch (fi61 + (fi61 & fi61)) {
                                            case 4294967295:
                                                v54 = v54
                                            default:
                                                v52 = 255
                                            }
                                        }
                                        v6["hello"] = "a"
                                    }
                                    _ = v64
                                    _ = v65
                                }
                                if len(v51) >= 2 {
                                    v51 = v51[1:len(v51)]
                                }
                            }
                            for fi76 := 0; fi76 < 11; fi76++ {
                                v1 += 100
                                _ = fi76
                            }
                            switch fi61 {
                            case 32768:
                                v54 += (v54 << ((0) & 63))
                            default:
                                switch (fi61 + (fi61 / ((fi61) | 1))) {
                                case 32768:
                                    for fi77 := 0; fi77 < 13; fi77++ {
                                        var v78 int32 = (v34 + v34)
                                        fmt.Println(fi77)
                                        _ = fi77
                                        _ = v78
                                    }
                                case 1:
                                    switch -967 {
                                    case 2147483647:
                                        if true {
                                            v1 = v1
                                            v53 = (v1 / (((v1 / (((v53 + (-(v53)))) | 1))) | 1))
                                        } else {
                                            for fi79 := 0; fi79 < 1; fi79++ {
                                                _ = len(v51)
                                                v80 := GenS2{F0: fi61, F1: uint64(v54), F2: -264}
                                                _ = fi79
                                                _ = v80
                                            }
                                        }
                                    case 293:
                                        v54 += v54
                                    case 951:
                                        v81 := GenS6{F0: int8(v3), F1: v4, F2: 32767}
                                        _ = v81
                                    default:
                                        v4 = v4
                                    }
                                case -9223372036854775808:
                                    if false {
                                        if false {
                                            fmt.Println(v54, fi61)
                                            if (fi61 != fi61) {
                                                v53 = 0
                                                for fi82 := 0; fi82 < 18; fi82++ {
                                                    fmt.Println(v54)
                                                    v83 := GenS4{F0: (v52 ^ (v3 & (v50 + (v52 - v52)))), F1: 492, F2: int8(v2)}
                                                    v34 = v34
                                                    _ = fi82
                                                    _ = v83
                                                }
                                                var v84 float64 = 1.5
                                                _ = v84
                                            }
                                        } else {
                                            if ((v54 % ((v54) | 1)) >= (^(v54))) {
                                                v52 = (v52 << ((255) & 63))
                                                for fi85 := 0; fi85 < 4; fi85++ {
                                                    v54 = uint32(v3)
                                                    v86 := GenS5{F0: 1, F1: (v54 ^ (v54 + v54))}
                                                    v4 = 1000
                                                    _ = fi85
                                                    _ = v86
                                                }
                                            }
                                            var v87 uint64 = 483
                                            _ = v87
                                        }
                                        for fi88 := 0; fi88 < 17; fi88++ {
                                            v89 := GenS1{F0: 1, F1: "foo bar"}
                                            _ = fi88
                                            _ = v89
                                        }
                                    } else {
                                        var v90 uint = 18446744073709551615
                                        var v91 int64 = 1
                                        _ = v90
                                        _ = v91
                                    }
                                default:
                                    v52 = genFn0((v1 & (v53 % (((v1 << ((960) & 63))) | 1))))
                                }
                            }
                            _ = fi61
                        }
                    default:
                        v1 = (v1 | -1)
                    }
                    _ = v54
                }
                _ = v53
            } else {
                v50 = v3
            }
        }
        if ((v34 - 2147483647) == (^(v34))) {
            switch 1 {
            case 1:
                if true {
                    var v92 float32 = (1.5 / -2.25)
                    _ = v92
                } else {
                    v4 = v4
                    v1 = int8(v50)
                }
            case -451:
                for fi93 := 0; fi93 < 7; fi93++ {
                    v94 := GenS5{F0: fi93, F1: 48}
                    v95 := GenS5{F0: fi93, F1: 0}
                    v96 := GenS2{F0: (fi93 * (^(fi93))), F1: 4294967295, F2: v2}
                    _ = fi93
                    _ = v94
                    _ = v95
                    _ = v96
                }
            case 42:
                var v97 uint = 4294967295
                _ = v97
            }
            var v98 string = ""
            for fi99 := 0; fi99 < 13; fi99++ {
                if (false && true) {
                    v34 = v34
                    fmt.Println(fi99)
                    v4 += v4
                }
                var v100 int = (fi99 << ((732) & 63))
                if true {
                    var v101 uint = 1
                    for fi102 := 0; fi102 < 15; fi102++ {
                        switch 2147483647 {
                        case -9223372036854775808:
                            v103 := GenS6{F0: v1, F1: (v4 * (v4 ^ v4)), F2: v4}
                            _ = v103
                        case 127:
                            var v104 int64 = -1
                            _ = v104
                        }
                        var v105 uint64 = 314
                        _ = fi102
                        _ = v105
                    }
                    _ = v101
                } else {
                    if (!(false && true)) {
                        var v106 string = (v98 + ("" + v98))
                        fmt.Println(fi99, v100, v4)
                        _ = v106
                    }
                    var v107 string = (v98 + v98)
                    _ = v107
                }
                _ = fi99
                _ = v100
            }
            _ = v98
        }
        _ = v52
    }
    for fi108 := 0; fi108 < 1; fi108++ {
        fmt.Println(v3)
        v50 = v50
        _ = fi108
    }
    v109 := GenS6{F0: (-(v1)), F1: 157, F2: (v4 * v4)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("slicelen v51:", len(v51), v51)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v109
    _ = v2
    _ = v3
    _ = v34
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v6
}
