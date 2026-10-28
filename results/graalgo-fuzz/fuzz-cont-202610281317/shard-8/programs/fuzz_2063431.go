package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 string }
type GenS2 struct { F0 int; F1 int32; F2 string }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint8; F1 int64 }
type GenS5 struct { F0 uint8; F1 uint16 }
type GenS6 struct { F0 string; F1 string }
type GenS7 struct { F0 int8; F1 uint16 }
type GenS8 struct { F0 int16; F1 int32 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 uint64; F1 int32 }
type GenS11 struct { F0 bool; F1 int }
type GenS12 struct { F0 int; F1 uint64 }
type GenS13 struct { F0 uint }
type GenS14 struct { F0 bool; F1 int64 }
type GenS15 struct { F0 int8 }
type GenS16 struct { F0 int16; F1 uint16 }

func main() {
    var v1 int64 = -9223372036854775808
    v2 := []int{89}
    v3 := make(map[uint32]int)
    delete(v3, 42)
    var v4 float32 = 1.5
    v4 += v4
    v4 = v4
    _ = len(v3)
    if false {
        var v5 uint64 = 18446744073709551615
        _ = v5
    } else {
        v6 := GenS1{F0: 4294967295, F1: (v1 * v1), F2: "a"}
        fmt.Println(v1, v4)
        _ = v6
    }
    v7 := GenS2{F0: int(v1), F1: -258, F2: ""}
    v1 = 501
    v3[65535] = -1
    if (!false) {
        if (!(false || (-65 > 127))) {
            switch -1 {
            case 32768:
                switch "\n\t" {
                case "\n\t":
                    v1 = v1
                case "a":
                    switch ("\n\t" + ("\n\t" + "foo bar")) {
                    case "":
                        var v8 int8 = 100
                        _ = v8
                    case "\n\t":
                        fmt.Println(v4, v1)
                    case "hello":
                        var v9 uint8 = 245
                        _ = v9
                    }
                default:
                    if (!(65535 > 0)) {
                        v10 := GenS2{F0: -2147483648, F1: 453, F2: (("foo bar" + "foo bar") + (("a" + "foo bar") + "\n\t"))}
                        v2 = append(v2, 1)
                        _ = v10
                    } else {
                        if true {
                            var v11 bool = (42 != -81)
                            var v12 int32 = -2147483648
                            if ((-2.25 <= ((((1e10 + (float64(v4) + (float64(3.14159) + (1e10 / 1.0)))) * float64(((v4 / v4) * 3.14159))) / 1e10) - -2.25)) || (("a" + ((("a" + ("a" + "")) + "a") + (("" + "a") + "foo bar"))) <= "foo bar")) {
                                for fi13 := 0; fi13 < 11; fi13++ {
                                    if (32767 > 439) {
                                        v12 = 1
                                        v14 := GenS2{F0: fi13, F1: v12, F2: "a"}
                                        var v15 uint32 = 2147483648
                                        _ = v14
                                        _ = v15
                                    } else {
                                        var v16 bool = (!false)
                                        _ = v16
                                    }
                                    switch "" {
                                    case "hello":
                                        fmt.Println(v12, v4, v11)
                                    case "\n\t":
                                        v11 = ((v4 <= ((v4 - v4) * float32(0.0))) || v11)
                                    default:
                                        if (false || (293 == 255)) {
                                            fmt.Println(v12, v1, v12)
                                            switch (fi13 % ((fi13) | 1)) {
                                            case 42:
                                                if len(v2) > 0 {
                                                    v2[0] = (^(fi13))
                                                }
                                            case -309:
                                                if ((v11 || (2 < 0)) && (((v4 * v4) * -2.25) >= 0.1)) {
                                                    v1 -= 42
                                                    for fi17 := 0; fi17 < 4; fi17++ {
                                                        if ((v12 >> ((42) & 63)) > (v12 | (v12 * v12))) {
                                                            _ = len(v3)
                                                            if (42 == 71) {
                                                                if ((42 >= -105) && v11) {
                                                                    if ((65535 < 32767) || ((v1 % (((v1 | 9223372036854775807)) | 1)) < (^(v1)))) {
                                                                        v2 = append(v2, (fi13 / (((^(fi17))) | 1)))
                                                                    }
                                                                    v1 = (v1 & 301)
                                                                    v18 := GenS3{F0: 2}
                                                                    _ = v18
                                                                } else {
                                                                    if (v11 && false) {
                                                                        v19 := GenS3{F0: 42}
                                                                        switch ("hello" + "hello") {
                                                                        case "a":
                                                                            var v20 uint = 0
                                                                            _ = v20
                                                                        default:
                                                                            switch -1 {
                                                                            case 491:
                                                                                if len(v2) > 0 {
                                                                                    v2[0] = fi13
                                                                                }
                                                                            case 4294967295:
                                                                                switch fi13 {
                                                                                case 1:
                                                                                    fmt.Println(v12)
                                                                                case 127:
                                                                                    var v21 uint32 = 0
                                                                                    _ = v21
                                                                                case -1:
                                                                                    switch "\n\t" {
                                                                                    case "hello":
                                                                                        var v22 float32 = v4
                                                                                        _ = v22
                                                                                    }
                                                                                }
                                                                            default:
                                                                                v23 := GenS4{F0: 155, F1: v1}
                                                                                _ = v23
                                                                            }
                                                                        }
                                                                        v3[403] = int(v1)
                                                                        var v24 int = (fi13 * 0)
                                                                        v25 := GenS5{F0: 199, F1: 42}
                                                                        switch fi13 {
                                                                        case 42:
                                                                            switch -9223372036854775808 {
                                                                            case -435:
                                                                                var v26 int = (fi13 - (fi17 | (v24 - int(v1))))
                                                                                _ = v26
                                                                            case 32767:
                                                                                if (true || v11) {
                                                                                    var v27 int8 = 42
                                                                                    v11 = (2 == 4294967295)
                                                                                    _ = v27
                                                                                } else {
                                                                                    var v28 uint16 = 2
                                                                                    _ = v28
                                                                                }
                                                                            case 42:
                                                                                var v29 uint8 = 88
                                                                                _ = v29
                                                                            default:
                                                                                if (int16(v1) <= 0) {
                                                                                    fmt.Println(v12)
                                                                                    v11 = (338 <= 18446744073709551615)
                                                                                }
                                                                            }
                                                                        default:
                                                                            v30 := GenS6{F0: "a", F1: (("foo bar" + "a") + "hello")}
                                                                            var v31 int32 = (v12 - (v12 / ((v12) | 1)))
                                                                            _ = v30
                                                                            _ = v31
                                                                        }
                                                                        _ = v19
                                                                        _ = v24
                                                                        _ = v25
                                                                    }
                                                                    switch ("a" + ("a" + "hello")) {
                                                                    case "":
                                                                        var v32 uint64 = 0
                                                                        _ = v32
                                                                    case "hello":
                                                                        v33 := GenS7{F0: -128, F1: 255}
                                                                        _ = v33
                                                                    }
                                                                    switch "foo bar" {
                                                                    case "hello":
                                                                        v11 = v11
                                                                    case "":
                                                                        v34 := GenS5{F0: 232, F1: 255}
                                                                        _ = v34
                                                                    case "\n\t":
                                                                        v1 = (-(v1))
                                                                    default:
                                                                        v4 -= (v4 / 1e10)
                                                                    }
                                                                }
                                                                if (0 >= 1) {
                                                                    v4 = (v4 * 1.5)
                                                                    if v11 {
                                                                        v11 = (v11 && v11)
                                                                        v12 -= v12
                                                                    } else {
                                                                        var v35 float64 = -2.25
                                                                        v36 := GenS8{F0: -235, F1: v12}
                                                                        _ = v35
                                                                        _ = v36
                                                                    }
                                                                }
                                                            } else {
                                                                _ = len(v3)
                                                                v37 := GenS1{F0: 911, F1: v1, F2: (((("" + "\n\t") + "hello") + ("" + "\n\t")) + ("hello" + "hello"))}
                                                                _ = v37
                                                            }
                                                        }
                                                        var v38 uint = 4294967295
                                                        var v39 int = (-(fi13))
                                                        _ = fi17
                                                        _ = v38
                                                        _ = v39
                                                    }
                                                } else {
                                                    var v40 bool = (18446744073709551615 >= 65535)
                                                    var v41 int16 = -279
                                                    _ = v40
                                                    _ = v41
                                                }
                                            default:
                                                var v42 float64 = (-2.25 / ((-2.25 + 1.5) + 0.0))
                                                _ = v42
                                            }
                                            switch -543 {
                                            case 1:
                                                var v43 uint16 = 2
                                                _ = v43
                                            default:
                                                var v44 int32 = (v12 + (v12 & 0))
                                                _ = v44
                                            }
                                        } else {
                                            switch "" {
                                            case "foo bar":
                                                for fi45 := 0; fi45 < 10; fi45++ {
                                                    fmt.Println(fi13, v4)
                                                    _ = fi45
                                                }
                                            default:
                                                v46 := GenS9{F0: 520}
                                                _ = v46
                                            }
                                            if (514 >= -435) {
                                                switch "\n\t" {
                                                case "":
                                                    switch "" {
                                                    case "\n\t":
                                                        switch 9223372036854775807 {
                                                        case 0:
                                                            if (v4 > (v4 * ((float32(3.14159) / (0.0 - v4)) * (3.14159 - (0.0 - 0.1))))) {
                                                                switch "" {
                                                                case "\n\t":
                                                                    var v47 int = fi13
                                                                    _ = v47
                                                                case "a":
                                                                    var v48 int8 = 0
                                                                    _ = v48
                                                                }
                                                                for fi49 := 0; fi49 < 17; fi49++ {
                                                                    v11 = (true || (4294967295 <= 65535))
                                                                    _ = fi49
                                                                }
                                                            } else {
                                                                v12 = v12
                                                                fmt.Println(v12, v11)
                                                            }
                                                        case 9223372036854775807:
                                                            v11 = v11
                                                        }
                                                    case "a":
                                                        var v50 float32 = 1.5
                                                        _ = v50
                                                    default:
                                                        for fi51 := 0; fi51 < 10; fi51++ {
                                                            v12 = v12
                                                            v52 := v3[2147483648]
                                                            if (3.14159 == (0.1 - 0.1)) {
                                                                v53 := GenS10{F0: 181, F1: -1}
                                                                _ = v53
                                                            }
                                                            if (1e10 != float64(float32((1e10 - 1.5)))) {
                                                                _ = len(v2)
                                                                if true {
                                                                    if ((!(128 > 19)) && (0 > 856)) {
                                                                        if (0.1 != (float64(float32((3.14159 / -2.25))) / -2.25)) {
                                                                            if v11 {
                                                                                v4 -= (v4 - (float32((0.0 / 1e10)) - (v4 / 3.14159)))
                                                                                fmt.Println(fi13, v1)
                                                                                v11 = true
                                                                            }
                                                                            v54 := GenS11{F0: v11, F1: (fi13 & 2147483647)}
                                                                            v12 = -1
                                                                            _ = v54
                                                                        } else {
                                                                            v55 := GenS4{F0: 177, F1: 1}
                                                                            v3[42] = (v52 | v52)
                                                                            var v56 int8 = 7
                                                                            v57 := GenS12{F0: fi13, F1: 4294967295}
                                                                            v12 -= (v12 ^ (v12 % (((v12 / (((v12 - v12)) | 1))) | 1)))
                                                                            _ = v55
                                                                            _ = v56
                                                                            _ = v57
                                                                        }
                                                                    }
                                                                } else {
                                                                    v12 -= (v12 >> ((1) & 63))
                                                                    if ((-2.25 + ((3.14159 * float64(v4)) / 3.14159)) > ((0.1 / ((3.14159 + (-2.25 - float64(float32((-2.25 - 3.14159))))) + 1.5)) / float64(float32(3.14159)))) {
                                                                        v11 = (v1 >= 488)
                                                                        var v58 bool = ((float32(1e10) / (1e10 * v4)) == 0.0)
                                                                        switch -9223372036854775808 {
                                                                        case -800:
                                                                            v59 := GenS8{F0: 42, F1: (v12 << ((255) & 63))}
                                                                            if (v58 && true) {
                                                                                v11 = v58
                                                                            } else {
                                                                                switch fi51 {
                                                                                case 0:
                                                                                    v60 := GenS12{F0: (fi51 & (fi13 & (-(fi13)))), F1: 2}
                                                                                    _ = v60
                                                                                case 32767:
                                                                                    if (127 == 0) {
                                                                                        if (!(42 != -1)) {
                                                                                            fmt.Println(v58, fi13, v52)
                                                                                            if v11 {
                                                                                                switch v52 {
                                                                                                case -130:
                                                                                                    switch "" {
                                                                                                    case "\n\t":
                                                                                                        var v61 int16 = -360
                                                                                                        _ = v61
                                                                                                    case "hello":
                                                                                                        var v62 int16 = 42
                                                                                                        _ = v62
                                                                                                    case "a":
                                                                                                        v58 = (!(42 == 255))
                                                                                                    default:
                                                                                                        var v63 uint8 = 1
                                                                                                        _ = v63
                                                                                                    }
                                                                                                case 290:
                                                                                                    fmt.Println(fi13, v58)
                                                                                                case 4294967295:
                                                                                                    var v64 uint16 = 881
                                                                                                    _ = v64
                                                                                                default:
                                                                                                    if (65535 >= 42) {
                                                                                                        if v11 {
                                                                                                            fmt.Println(v52, v58, v11)
                                                                                                            var v65 string = "foo bar"
                                                                                                            _ = v65
                                                                                                        }
                                                                                                        var v66 uint = 1
                                                                                                        v4 += (0.1 * 0.0)
                                                                                                        _ = v66
                                                                                                    } else {
                                                                                                        v58 = v11
                                                                                                        if v11 {
                                                                                                            if v11 {
                                                                                                                v11 = (!false)
                                                                                                            }
                                                                                                            fmt.Println(v1, v1)
                                                                                                            if v58 {
                                                                                                                fmt.Println(v52, fi13, v52)
                                                                                                                if len(v2) >= 2 {
                                                                                                                    v2 = v2[1:len(v2)]
                                                                                                                }
                                                                                                                fmt.Println(v58)
                                                                                                            }
                                                                                                        } else {
                                                                                                            v4 = (v4 + ((v4 / ((0.1 - v4) / (v4 - v4))) + v4))
                                                                                                        }
                                                                                                        if false {
                                                                                                            switch 127 {
                                                                                                            case 4294967295:
                                                                                                                var v67 uint = 18446744073709551615
                                                                                                                _ = v67
                                                                                                            case -1:
                                                                                                                v4 -= (v4 - ((1.5 + v4) + v4))
                                                                                                            default:
                                                                                                                var v68 float64 = 0.0
                                                                                                                _ = v68
                                                                                                            }
                                                                                                        } else {
                                                                                                            var v69 uint = 0
                                                                                                            _ = v69
                                                                                                        }
                                                                                                        delete(v3, 977)
                                                                                                    }
                                                                                                }
                                                                                                switch v52 {
                                                                                                case 2147483647:
                                                                                                    switch "a" {
                                                                                                    case "a":
                                                                                                        _ = len(v2)
                                                                                                    case "hello":
                                                                                                        v4 = v4
                                                                                                    case "":
                                                                                                        v4 = v4
                                                                                                    default:
                                                                                                        v52 -= (fi51 ^ 4294967295)
                                                                                                    }
                                                                                                    _ = len(v3)
                                                                                                    var v70 int64 = (v1 - (v1 << ((1) & 63)))
                                                                                                    _ = v70
                                                                                                case -897:
                                                                                                    var v71 string = "hello"
                                                                                                    _ = v71
                                                                                                case 127:
                                                                                                    if v58 {
                                                                                                        var v72 int64 = v1
                                                                                                        _ = v72
                                                                                                    }
                                                                                                }
                                                                                                switch fi51 {
                                                                                                case -1:
                                                                                                    var v73 int16 = 42
                                                                                                    _ = v73
                                                                                                }
                                                                                                var v74 float32 = 1e10
                                                                                                var v75 uint8 = 85
                                                                                                _ = v74
                                                                                                _ = v75
                                                                                            }
                                                                                        }
                                                                                    } else {
                                                                                        v58 = false
                                                                                        var v76 int = fi13
                                                                                        _ = v76
                                                                                    }
                                                                                case 281:
                                                                                    fmt.Println(v1, v4, v58)
                                                                                }
                                                                                switch -851 {
                                                                                case -1:
                                                                                    _ = len(v2)
                                                                                case 2147483647:
                                                                                    v77 := GenS8{F0: -32768, F1: v12}
                                                                                    _ = v77
                                                                                }
                                                                            }
                                                                            if (2 >= 0) {
                                                                                v11 = v58
                                                                            }
                                                                            _ = v59
                                                                        case 1:
                                                                            v78 := GenS2{F0: (^(fi51)), F1: int32(v1), F2: "\n\t"}
                                                                            _ = v78
                                                                        }
                                                                        fmt.Println(v12, v52, v1)
                                                                        if ((905 >= -32768) && v11) {
                                                                            var v79 float32 = (-2.25 + (1e10 - -2.25))
                                                                            switch -804 {
                                                                            case 176:
                                                                                v1 -= (^(v1))
                                                                            }
                                                                            _ = v79
                                                                        }
                                                                        if false {
                                                                            v58 = (true || true)
                                                                        } else {
                                                                            if ((v58 || true) && v58) {
                                                                                v80 := GenS11{F0: true, F1: (v52 << ((372) & 63))}
                                                                                v4 += (3.14159 + ((float32(0.1) * float32(1.5)) / v4))
                                                                                v11 = (v58 || (("" == "\n\t") || (v58 && v58)))
                                                                                _ = v80
                                                                            }
                                                                            if (!false) {
                                                                                fmt.Println(v4, v4, v12)
                                                                                v4 = 1.5
                                                                                v81 := GenS13{F0: 1}
                                                                                _ = v81
                                                                            }
                                                                        }
                                                                        v12 = (v12 >> ((906) & 63))
                                                                        _ = v58
                                                                    } else {
                                                                        var v82 string = (("" + ("foo bar" + "hello")) + "")
                                                                        v12 -= (v12 + v12)
                                                                        _ = v82
                                                                    }
                                                                    v12 += 42
                                                                }
                                                            } else {
                                                                var v83 string = "a"
                                                                v3[2147483648] = (fi51 >> ((42) & 63))
                                                                if true {
                                                                    v84 := v3[252]
                                                                    v83 = (v83 + ((v83 + v83) + v83))
                                                                    _ = v84
                                                                }
                                                                switch -523 {
                                                                case 128:
                                                                    if v11 {
                                                                        v83 = v83
                                                                        switch 128 {
                                                                        case 2147483647:
                                                                            var v85 uint = 255
                                                                            _ = v85
                                                                        }
                                                                    } else {
                                                                        v52 -= (v52 << ((4294967295) & 63))
                                                                    }
                                                                default:
                                                                    var v86 bool = true
                                                                    _ = v86
                                                                }
                                                                var v87 uint64 = 18446744073709551615
                                                                var v88 int = (v52 * (fi51 | fi13))
                                                                _ = v83
                                                                _ = v87
                                                                _ = v88
                                                            }
                                                            _ = fi51
                                                            _ = v52
                                                        }
                                                    }
                                                }
                                                if (2147483648 == 2147483648) {
                                                    fmt.Println(v4, v4, v1)
                                                    v11 = (v1 >= (v1 % (((v1 * v1)) | 1)))
                                                } else {
                                                    switch (("hello" + ("\n\t" + ("" + "\n\t"))) + "foo bar") {
                                                    case "foo bar":
                                                        for fi89 := 0; fi89 < 18; fi89++ {
                                                            var v90 uint64 = 0
                                                            _ = fi89
                                                            _ = v90
                                                        }
                                                    case "\n\t":
                                                        var v91 int64 = 1
                                                        _ = v91
                                                    }
                                                    for fi92 := 0; fi92 < 6; fi92++ {
                                                        if v11 {
                                                            var v93 uint8 = 42
                                                            v94 := v3[4294967295]
                                                            _ = v93
                                                            _ = v94
                                                        }
                                                        _ = len(v3)
                                                        _ = fi92
                                                    }
                                                    for fi95 := 0; fi95 < 14; fi95++ {
                                                        fmt.Println(fi95)
                                                        if ((false || true) || v11) {
                                                            if v11 {
                                                                var v96 int64 = v1
                                                                _ = v96
                                                            }
                                                            switch "" {
                                                            case "":
                                                                var v97 float64 = 0.1
                                                                _ = v97
                                                            case "a":
                                                                var v98 uint16 = 32767
                                                                _ = v98
                                                            case "hello":
                                                                var v99 float32 = v4
                                                                _ = v99
                                                            }
                                                            v11 = (938 > v1)
                                                            v12 += 2147483647
                                                            fmt.Println(v1, fi95)
                                                        } else {
                                                            if len(v2) > 0 {
                                                                v2[0] = (fi13 / ((-96) | 1))
                                                            }
                                                            v4 = (v4 - float32(0.1))
                                                            var v100 uint32 = 2147483648
                                                            _ = v100
                                                        }
                                                        _ = fi95
                                                    }
                                                }
                                                switch (fi13 % ((42) | 1)) {
                                                case -9223372036854775808:
                                                    if v11 {
                                                        v4 = v4
                                                    } else {
                                                        v101 := GenS14{F0: (32767 < 255), F1: (v1 * -534)}
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        _ = v101
                                                    }
                                                case 128:
                                                    v102 := GenS13{F0: 18446744073709551615}
                                                    _ = v102
                                                case 4294967295:
                                                    if v11 {
                                                        v103 := GenS15{F0: int8(v12)}
                                                        if (!v11) {
                                                            v4 -= 3.14159
                                                            v1 = (v1 - -492)
                                                            v104 := GenS15{F0: 1}
                                                            _ = v104
                                                        } else {
                                                            for fi105 := 0; fi105 < 10; fi105++ {
                                                                var v106 uint64 = 42
                                                                _ = fi105
                                                                _ = v106
                                                            }
                                                            if (!true) {
                                                                var v107 bool = (true || (!(((366 != -32768) || v11) || v11)))
                                                                _ = v107
                                                            }
                                                        }
                                                        switch 32768 {
                                                        case -1:
                                                            switch "" {
                                                            case "\n\t":
                                                                for fi108 := 0; fi108 < 11; fi108++ {
                                                                    v1 += 9223372036854775807
                                                                    v4 = float32(0.1)
                                                                    fmt.Println(fi13, v1, fi108)
                                                                    fmt.Println(v1, v12)
                                                                    _ = fi108
                                                                }
                                                            default:
                                                                var v109 uint64 = 2
                                                                _ = v109
                                                            }
                                                        default:
                                                            v11 = v11
                                                        }
                                                        _ = v103
                                                    }
                                                }
                                            }
                                        }
                                    }
                                    v110 := GenS6{F0: "foo bar", F1: ""}
                                    _ = fi13
                                    _ = v110
                                }
                                if (!(1 != 4294967295)) {
                                    if len(v2) > 0 {
                                        v2[0] = 128
                                    }
                                    for fi111 := 0; fi111 < 11; fi111++ {
                                        var v112 bool = v11
                                        var v113 float64 = ((((0.1 / 1.0) / float64(v4)) / (-2.25 / (0.1 / (1e10 * float64((0.0 - (0.0 - v4))))))) / 3.14159)
                                        fmt.Println(v12)
                                        _ = fi111
                                        _ = v112
                                        _ = v113
                                    }
                                }
                            }
                            _ = v11
                            _ = v12
                        } else {
                            v1 += (v1 + 9223372036854775807)
                        }
                        fmt.Println(v4, v1)
                        v3[2147483648] = 4294967295
                    }
                }
            default:
                switch -1 {
                case -148:
                    v114 := GenS14{F0: (((v1 == v1) || (!true)) || (!false)), F1: (^(v1))}
                    _ = v114
                case 933:
                    var v115 string = ""
                    _ = v115
                }
            }
            fmt.Println(v4, v4)
        } else {
            if false {
                for fi116 := 0; fi116 < 9; fi116++ {
                    for fi117 := 0; fi117 < 1; fi117++ {
                        v4 = (v4 / v4)
                        _ = fi117
                    }
                    if true {
                        if ((!(100 <= 1)) || true) {
                            if (!((2 >= 186) && (!true))) {
                                v4 += 0.1
                            }
                        } else {
                            delete(v3, 2)
                            var v118 int64 = v1
                            v1 -= (v1 + v1)
                            _ = v118
                        }
                    }
                    _ = fi116
                }
                var v119 int = 32768
                _ = v119
            }
            fmt.Println(v4)
            v1 = -514
        }
    } else {
        v120 := GenS16{F0: 1000, F1: 299}
        _ = v120
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
