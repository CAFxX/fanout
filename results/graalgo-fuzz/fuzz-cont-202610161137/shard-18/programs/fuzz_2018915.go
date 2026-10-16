package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint; F1 uint8; F2 string }
type GenS4 struct { F0 uint16; F1 int64; F2 int32 }
type GenS5 struct { F0 uint32; F1 uint16 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int; F1 uint16; F2 uint }
type GenS8 struct { F0 int; F1 uint16 }
type GenS9 struct { F0 int8; F1 bool }
type GenS10 struct { F0 int32; F1 bool; F2 uint8 }
type GenS11 struct { F0 bool; F1 uint16; F2 uint32 }
type GenS12 struct { F0 int8; F1 int8 }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 bool }
type GenS15 struct { F0 uint }
type GenS16 struct { F0 string }
type GenS17 struct { F0 int32; F1 uint32 }
type GenS18 struct { F0 uint16 }
type GenS19 struct { F0 bool }
type GenS20 struct { F0 int8; F1 int64; F2 uint64 }
type GenS21 struct { F0 int; F1 uint8; F2 uint64 }
type GenS22 struct { F0 bool; F1 uint8 }
type GenS23 struct { F0 int; F1 uint32; F2 string }
type GenS24 struct { F0 int16; F1 int16; F2 int8 }
type GenS25 struct { F0 int }
type GenS26 struct { F0 int; F1 uint16; F2 uint64 }
type GenS27 struct { F0 uint16; F1 uint64 }
type GenS28 struct { F0 bool; F1 int64; F2 int8 }
type GenS29 struct { F0 int16; F1 uint64; F2 uint }
type GenS30 struct { F0 int8; F1 string; F2 uint16 }

func main() {
    v1 := []string{"foo bar", ("a" + ("" + ""))}
    v2 := make(map[int]string)
    var v3 bool = (true && ("" > (((("foo bar" + "\n\t") + ("foo bar" + "\n\t")) + "\n\t") + (("\n\t" + "\n\t") + "hello"))))
    var v4 int = 880
    v5 := GenS1{F0: 0, F1: -1}
    v1 = append(v1, "hello")
    for fi6 := 0; fi6 < 7; fi6++ {
        if false {
            v7 := GenS1{F0: 1, F1: 614}
            v3 = (("hello" + "foo bar") != "a")
            fmt.Println(fi6, fi6)
            _ = v7
        }
        if (-521 == -9223372036854775808) {
            var v8 int64 = 776
            v4 = v4
            var v9 int64 = 502
            _ = v8
            _ = v9
        }
        for fi10 := 0; fi10 < 10; fi10++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi10
        }
        _ = fi6
    }
    v1 = append(v1, "foo bar")
    for fi11 := 0; fi11 < 9; fi11++ {
        switch ("a" + (((("foo bar" + (("" + ("hello" + (("\n\t" + "a") + "a"))) + "")) + (("\n\t" + "hello") + "foo bar")) + "a") + ((("a" + "foo bar") + "hello") + ("\n\t" + "foo bar")))) {
        case "":
            fmt.Println(v4)
        case "foo bar":
            _ = len(v2)
        case "hello":
            if false {
                if v3 {
                    var v12 float32 = (-2.25 / 3.14159)
                    for fi13 := 0; fi13 < 10; fi13++ {
                        if v3 {
                            fmt.Println(v3)
                            if ("foo bar" > ((("hello" + "\n\t") + "foo bar") + "foo bar")) {
                                v14 := v2[fi11]
                                v3 = ((fi11 >> ((876) & 63)) > fi13)
                                switch ("foo bar" + "a") {
                                case "":
                                    var v15 int8 = 0
                                    _ = v15
                                case "foo bar":
                                    if ((609 < 18446744073709551615) || (-770 != 0)) {
                                        if v3 {
                                            v14 = v14
                                            _ = len(v2)
                                            v4 += 1
                                        } else {
                                            v3 = v3
                                            v16 := GenS2{F0: 0}
                                            _ = v16
                                        }
                                        if (("hello" + v14) > ((v14 + v14) + v14)) {
                                            v17 := GenS2{F0: 49}
                                            _ = v17
                                        }
                                        v18 := GenS2{F0: 1}
                                        v19 := GenS3{F0: 42, F1: 127, F2: "\n\t"}
                                        _ = v18
                                        _ = v19
                                    } else {
                                        if true {
                                            v12 -= 3.14159
                                            switch (-(v4)) {
                                            case 32767:
                                                v3 = ((!v3) || false)
                                            case 877:
                                                fmt.Println(v12)
                                            default:
                                                var v20 uint8 = 147
                                                _ = v20
                                            }
                                        } else {
                                            v14 = "foo bar"
                                            if (32767 >= 411) {
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                v12 -= v12
                                            } else {
                                                v4 = fi13
                                                if ((v14 > v14) || (!(false && (v3 || (!v3))))) {
                                                    var v21 string = (v14 + v14)
                                                    v21 = (("foo bar" + v21) + (v21 + (v21 + ("hello" + ((v21 + (((v14 + "foo bar") + (v14 + v14)) + (((("foo bar" + (v14 + ("a" + "a"))) + ("" + ("hello" + "hello"))) + "hello") + ""))) + (((v21 + ("foo bar" + "")) + v14) + (v14 + v21)))))))
                                                    if v3 {
                                                        v22 := GenS2{F0: 18446744073709551615}
                                                        v1 = append(v1, ("hello" + "hello"))
                                                        _ = v22
                                                    } else {
                                                        if (!(v3 && v3)) {
                                                            var v23 string = v21
                                                            var v24 string = v23
                                                            var v25 int32 = 1
                                                            _ = v23
                                                            _ = v24
                                                            _ = v25
                                                        }
                                                        v26 := GenS1{F0: -32768, F1: 0}
                                                        switch "" {
                                                        case "\n\t":
                                                            v4 -= fi11
                                                        case "a":
                                                            if true {
                                                                if (-1 == 0) {
                                                                    if (-359 >= -2147483648) {
                                                                        v27 := GenS4{F0: 255, F1: 9223372036854775807, F2: -2147483648}
                                                                        if v3 {
                                                                            if (42 <= 18446744073709551615) {
                                                                                if (!true) {
                                                                                    if true {
                                                                                        var v28 bool = v3
                                                                                        _ = v28
                                                                                    } else {
                                                                                        if false {
                                                                                            v21 = (((v21 + "hello") + (v14 + v14)) + v21)
                                                                                            v3 = (0 < 9223372036854775807)
                                                                                        } else {
                                                                                            fmt.Println(v12)
                                                                                            var v29 uint16 = 2
                                                                                            _ = v29
                                                                                        }
                                                                                        if true {
                                                                                            v21 = ("foo bar" + (v21 + v21))
                                                                                            if v3 {
                                                                                                var v30 uint64 = 2
                                                                                                _ = v30
                                                                                            } else {
                                                                                                fmt.Println(fi11, v14, v3)
                                                                                            }
                                                                                            if len(v1) > 0 {
                                                                                                v1[0] = v21
                                                                                            }
                                                                                            v14 = v14
                                                                                        } else {
                                                                                            var v31 float32 = v12
                                                                                            _ = v31
                                                                                        }
                                                                                        switch ((("foo bar" + "foo bar") + v21) + ((v14 + ("foo bar" + v21)) + "foo bar")) {
                                                                                        case "foo bar":
                                                                                            v14 = "\n\t"
                                                                                        case "a":
                                                                                            var v32 bool = true
                                                                                            var v33 float32 = v12
                                                                                            var v34 int8 = 127
                                                                                            _ = v32
                                                                                            _ = v33
                                                                                            _ = v34
                                                                                        case "\n\t":
                                                                                            if v3 {
                                                                                                if (v3 || false) {
                                                                                                    var v35 uint32 = 308
                                                                                                    if v3 {
                                                                                                        fmt.Println(fi13, fi11, v35)
                                                                                                        if (!true) {
                                                                                                            v3 = true
                                                                                                            v4 = (fi13 * -9223372036854775808)
                                                                                                            var v36 string = ("foo bar" + (v14 + "\n\t"))
                                                                                                            _ = v36
                                                                                                        } else {
                                                                                                            v37 := GenS1{F0: 28, F1: 0}
                                                                                                            switch fi11 {
                                                                                                            case -978:
                                                                                                                v2[975] = v14
                                                                                                            case 9223372036854775807:
                                                                                                                var v38 int32 = 2147483647
                                                                                                                _ = v38
                                                                                                            }
                                                                                                            v39 := GenS2{F0: 0}
                                                                                                            _ = v37
                                                                                                            _ = v39
                                                                                                        }
                                                                                                        var v40 int64 = 383
                                                                                                        _ = v40
                                                                                                    } else {
                                                                                                        var v41 uint64 = 18446744073709551615
                                                                                                        if v3 {
                                                                                                            switch -737 {
                                                                                                            case 1:
                                                                                                                v4 += fi11
                                                                                                            case 838:
                                                                                                                fmt.Println(fi11)
                                                                                                            case 337:
                                                                                                                var v42 float32 = -2.25
                                                                                                                _ = v42
                                                                                                            }
                                                                                                            switch v4 {
                                                                                                            case -614:
                                                                                                                var v43 int64 = -774
                                                                                                                _ = v43
                                                                                                            case -887:
                                                                                                                v44 := GenS4{F0: 1, F1: -1, F2: 42}
                                                                                                                _ = v44
                                                                                                            default:
                                                                                                                _ = len(v1)
                                                                                                            }
                                                                                                            _ = len(v2)
                                                                                                            v45 := v2[(-(fi13))]
                                                                                                            _ = v45
                                                                                                        } else {
                                                                                                            var v46 uint = 255
                                                                                                            _ = v46
                                                                                                        }
                                                                                                        var v47 int16 = 1
                                                                                                        _ = v41
                                                                                                        _ = v47
                                                                                                    }
                                                                                                    v21 = ("foo bar" + "\n\t")
                                                                                                    _ = v35
                                                                                                } else {
                                                                                                    v48 := GenS3{F0: 340, F1: 0, F2: (v21 + ((((v14 + v21) + v21) + "\n\t") + v14))}
                                                                                                    _ = v48
                                                                                                }
                                                                                                v4 += (fi13 & fi13)
                                                                                                var v49 int32 = 42
                                                                                                _ = v49
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    v1 = append(v1, "\n\t")
                                                                                }
                                                                                var v50 uint = 18446744073709551615
                                                                                switch (fi13 / ((v4) | 1)) {
                                                                                case -247:
                                                                                    fmt.Println(fi13, v21, v12)
                                                                                case 4294967295:
                                                                                    if (fi13 <= v4) {
                                                                                        if true {
                                                                                            v21 = (v21 + v14)
                                                                                        }
                                                                                        if v3 {
                                                                                            var v51 float32 = (v12 + (v12 / v12))
                                                                                            if false {
                                                                                                var v52 int16 = -645
                                                                                                _ = len(v2)
                                                                                                if v3 {
                                                                                                    var v53 bool = true
                                                                                                    _ = v53
                                                                                                } else {
                                                                                                    v54 := GenS5{F0: 229, F1: 2}
                                                                                                    v3 = ((v51 < v12) && v3)
                                                                                                    switch v4 {
                                                                                                    case 127:
                                                                                                        var v55 uint8 = 87
                                                                                                        _ = v55
                                                                                                    case 32768:
                                                                                                        v52 -= int16(v50)
                                                                                                    case 1:
                                                                                                        v56 := GenS3{F0: 579, F1: 201, F2: v14}
                                                                                                        _ = v56
                                                                                                    }
                                                                                                    if v3 {
                                                                                                        if (fi11 <= fi13) {
                                                                                                            var v57 uint = v50
                                                                                                            var v58 uint16 = 42
                                                                                                            if (!v3) {
                                                                                                                if true {
                                                                                                                    if (-9223372036854775808 == int64(v58)) {
                                                                                                                        fmt.Println(fi11)
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    v1 = append(v1, "\n\t")
                                                                                                                }
                                                                                                                v59 := GenS3{F0: 42, F1: 203, F2: (("a" + v21) + (v21 + "\n\t"))}
                                                                                                                if v3 {
                                                                                                                    fmt.Println(v14, v12, fi13)
                                                                                                                } else {
                                                                                                                    if (v3 && (-2.25 < 0.1)) {
                                                                                                                        v60 := GenS6{F0: v14}
                                                                                                                        delete(v2, -294)
                                                                                                                        _ = v60
                                                                                                                    }
                                                                                                                    v4 += (^(fi11))
                                                                                                                    switch v4 {
                                                                                                                    case -192:
                                                                                                                        v61 := GenS6{F0: "a"}
                                                                                                                        _ = v61
                                                                                                                    default:
                                                                                                                        v62 := GenS7{F0: (v4 | (^(fi13))), F1: v58, F2: (v50 << (((v50 + (^(v50)))) & 63))}
                                                                                                                        _ = v62
                                                                                                                    }
                                                                                                                }
                                                                                                                fmt.Println(v21)
                                                                                                                fmt.Println(v21)
                                                                                                                _ = v59
                                                                                                            } else {
                                                                                                                switch (fi11 & (fi11 & 32768)) {
                                                                                                                case 4294967295:
                                                                                                                    switch "" {
                                                                                                                    case "foo bar":
                                                                                                                        v3 = (-208 >= (fi13 * fi11))
                                                                                                                    case "hello":
                                                                                                                        v63 := GenS1{F0: (v52 >> ((606) & 63)), F1: 1}
                                                                                                                        _ = v63
                                                                                                                    }
                                                                                                                    var v64 uint32 = 42
                                                                                                                    _ = v64
                                                                                                                default:
                                                                                                                    fmt.Println(v12)
                                                                                                                }
                                                                                                            }
                                                                                                            _ = v57
                                                                                                            _ = v58
                                                                                                        } else {
                                                                                                            v4 = 127
                                                                                                            var v65 int = (fi13 % (((fi13 % ((v4) | 1))) | 1))
                                                                                                            _ = v65
                                                                                                        }
                                                                                                    } else {
                                                                                                        v66 := v2[fi11]
                                                                                                        switch (fi11 * v4) {
                                                                                                        case 32767:
                                                                                                            switch v14 {
                                                                                                            case "\n\t":
                                                                                                                switch v4 {
                                                                                                                case -447:
                                                                                                                    var v67 bool = (9223372036854775807 <= (fi13 ^ fi13))
                                                                                                                    _ = v67
                                                                                                                case 0:
                                                                                                                    v3 = (!v3)
                                                                                                                case 56:
                                                                                                                    var v68 float64 = 0.0
                                                                                                                    _ = v68
                                                                                                                }
                                                                                                                v69 := GenS8{F0: v4, F1: 65535}
                                                                                                                v14 = ("hello" + "foo bar")
                                                                                                                if v3 {
                                                                                                                    switch "\n\t" {
                                                                                                                    case "hello":
                                                                                                                        var v70 int = fi11
                                                                                                                        _ = v70
                                                                                                                    case "":
                                                                                                                        v52 = (v52 << ((v50) & 63))
                                                                                                                    default:
                                                                                                                        if (4294967295 < 42) {
                                                                                                                            var v71 uint64 = 1
                                                                                                                            _ = len(v2)
                                                                                                                            _ = v71
                                                                                                                        } else {
                                                                                                                            var v72 int16 = 1
                                                                                                                            switch v21 {
                                                                                                                            case "":
                                                                                                                                var v73 string = "hello"
                                                                                                                                _ = v73
                                                                                                                            case "a":
                                                                                                                                if true {
                                                                                                                                    if v3 {
                                                                                                                                        if v3 {
                                                                                                                                            var v74 string = v21
                                                                                                                                            _ = len(v1)
                                                                                                                                            _ = v74
                                                                                                                                        } else {
                                                                                                                                            v75 := GenS9{F0: 16, F1: v3}
                                                                                                                                            switch "" {
                                                                                                                                            case "hello":
                                                                                                                                                delete(v2, fi13)
                                                                                                                                            case "":
                                                                                                                                                var v76 int16 = v72
                                                                                                                                                _ = v76
                                                                                                                                            }
                                                                                                                                            _ = v75
                                                                                                                                        }
                                                                                                                                        v77 := GenS3{F0: v50, F1: 42, F2: v14}
                                                                                                                                        var v78 string = "hello"
                                                                                                                                        _ = v77
                                                                                                                                        _ = v78
                                                                                                                                    } else {
                                                                                                                                        v52 += (-(v52))
                                                                                                                                        var v79 string = ("hello" + v66)
                                                                                                                                        switch v4 {
                                                                                                                                        case 127:
                                                                                                                                            fmt.Println(v12, v51)
                                                                                                                                            var v80 int8 = -128
                                                                                                                                            v12 += v51
                                                                                                                                            _ = v80
                                                                                                                                        default:
                                                                                                                                            v81 := GenS7{F0: 128, F1: 255, F2: (^(v50))}
                                                                                                                                            _ = v81
                                                                                                                                        }
                                                                                                                                        _ = v79
                                                                                                                                    }
                                                                                                                                    var v82 float64 = ((((((3.14159 + (1.5 * 3.14159)) - 0.1) * (0.0 - 1e10)) - (3.14159 / 1e10)) + (3.14159 - 1e10)) - -2.25)
                                                                                                                                    if false {
                                                                                                                                        v21 = v14
                                                                                                                                    }
                                                                                                                                    _ = v82
                                                                                                                                }
                                                                                                                            default:
                                                                                                                                var v83 uint32 = 2147483647
                                                                                                                                _ = v83
                                                                                                                            }
                                                                                                                            v84 := GenS5{F0: 592, F1: 26}
                                                                                                                            v51 += v12
                                                                                                                            _ = v72
                                                                                                                            _ = v84
                                                                                                                        }
                                                                                                                    }
                                                                                                                    var v85 int64 = 42
                                                                                                                    v50 = 525
                                                                                                                    _ = v85
                                                                                                                } else {
                                                                                                                    v86 := GenS10{F0: 42, F1: false, F2: 127}
                                                                                                                    delete(v2, fi11)
                                                                                                                    _ = v86
                                                                                                                }
                                                                                                                _ = v69
                                                                                                            case "hello":
                                                                                                                if (false || v3) {
                                                                                                                    v66 = (v14 + (("\n\t" + v66) + v14))
                                                                                                                } else {
                                                                                                                    v87 := GenS11{F0: (100 > 0), F1: 255, F2: 2147483647}
                                                                                                                    _ = v87
                                                                                                                }
                                                                                                            case "a":
                                                                                                                var v88 uint32 = 2147483648
                                                                                                                _ = v88
                                                                                                            default:
                                                                                                                var v89 int8 = int8(fi11)
                                                                                                                _ = v89
                                                                                                            }
                                                                                                            switch fi11 {
                                                                                                            case 9223372036854775807:
                                                                                                                fmt.Println(v51)
                                                                                                            case -776:
                                                                                                                var v90 uint8 = 1
                                                                                                                _ = v90
                                                                                                            case -856:
                                                                                                                fmt.Println(v3, v51, fi13)
                                                                                                            }
                                                                                                        case 4294967295:
                                                                                                            if len(v1) > 0 {
                                                                                                                v1[0] = v66
                                                                                                            }
                                                                                                        case 128:
                                                                                                            fmt.Println(v4, v66, v3)
                                                                                                        }
                                                                                                        if ((false || v3) && ((v12 - v51) < -2.25)) {
                                                                                                            v52 -= (v52 & -403)
                                                                                                            var v91 float32 = (0.1 * (0.0 + 1e10))
                                                                                                            v14 = v66
                                                                                                            fmt.Println(v52)
                                                                                                            var v92 uint = (v50 / (((v50 + v50)) | 1))
                                                                                                            _ = v91
                                                                                                            _ = v92
                                                                                                        }
                                                                                                        _ = v66
                                                                                                    }
                                                                                                    _ = v54
                                                                                                }
                                                                                                _ = v52
                                                                                            } else {
                                                                                                var v93 uint8 = 59
                                                                                                v94 := GenS12{F0: -128, F1: 21}
                                                                                                if true {
                                                                                                    fmt.Println(v4, v14)
                                                                                                }
                                                                                                _ = v93
                                                                                                _ = v94
                                                                                            }
                                                                                            var v95 bool = (1 >= 1)
                                                                                            fmt.Println(v95, v4, v95)
                                                                                            _ = v51
                                                                                            _ = v95
                                                                                        } else {
                                                                                            v21 = ((v14 + "\n\t") + (v21 + ("hello" + v21)))
                                                                                            var v96 uint = v50
                                                                                            v12 += 1e10
                                                                                            _ = v96
                                                                                        }
                                                                                    } else {
                                                                                        if (130 <= 16) {
                                                                                            v97 := GenS13{F0: 65535}
                                                                                            if true {
                                                                                                v98 := GenS8{F0: fi11, F1: 323}
                                                                                                _ = v98
                                                                                            }
                                                                                            v4 = (v4 + 32768)
                                                                                            _ = v97
                                                                                        }
                                                                                        v3 = (597 >= -349)
                                                                                        fmt.Println(fi13, v3, v14)
                                                                                    }
                                                                                }
                                                                                _ = v50
                                                                            } else {
                                                                                var v99 uint32 = 2
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = ""
                                                                                }
                                                                                v3 = (v12 <= float32(-2.25))
                                                                                if ((!v3) && true) {
                                                                                    if true {
                                                                                        v21 = "a"
                                                                                        delete(v2, (fi11 ^ (-(fi13))))
                                                                                    } else {
                                                                                        v3 = false
                                                                                    }
                                                                                } else {
                                                                                    v12 -= ((v12 + (v12 * v12)) - 0.0)
                                                                                }
                                                                                _ = v99
                                                                            }
                                                                            v14 = ((v14 + "hello") + (v14 + "foo bar"))
                                                                            var v100 float32 = 0.1
                                                                            _ = v100
                                                                        } else {
                                                                            v12 = (((v12 - (1.5 / v12)) / v12) * (float32(1e10) * v12))
                                                                            switch (v14 + v14) {
                                                                            case "a":
                                                                                var v101 float64 = (1.5 - 0.0)
                                                                                _ = v101
                                                                            case "foo bar":
                                                                                var v102 bool = (96 > 4294967295)
                                                                                _ = v102
                                                                            case "\n\t":
                                                                                _ = len(v1)
                                                                            default:
                                                                                var v103 int8 = 34
                                                                                _ = v103
                                                                            }
                                                                        }
                                                                        var v104 int64 = 9223372036854775807
                                                                        fmt.Println(fi11)
                                                                        if v3 {
                                                                            fmt.Println(v4)
                                                                        } else {
                                                                            v3 = true
                                                                        }
                                                                        _ = v104
                                                                        _ = v27
                                                                    }
                                                                    v105 := GenS14{F0: v3}
                                                                    switch fi11 {
                                                                    case 127:
                                                                        v21 = "\n\t"
                                                                    case 1:
                                                                        var v106 string = "foo bar"
                                                                        var v107 uint = 2
                                                                        if v3 {
                                                                            if true {
                                                                                v14 = "foo bar"
                                                                                if v3 {
                                                                                    v4 = fi11
                                                                                    if v3 {
                                                                                        fmt.Println(v3, fi11, v3)
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                fmt.Println(v21)
                                                                                fmt.Println(v3)
                                                                                v108 := GenS3{F0: v107, F1: 42, F2: ""}
                                                                                var v109 float32 = v12
                                                                                _ = v108
                                                                                _ = v109
                                                                            }
                                                                            v14 = "hello"
                                                                            var v110 uint = (^(v107))
                                                                            _ = v110
                                                                        } else {
                                                                            v111 := GenS5{F0: 4294967295, F1: 499}
                                                                            v1 = append(v1, (v106 + "foo bar"))
                                                                            _ = v111
                                                                        }
                                                                        _ = v106
                                                                        _ = v107
                                                                    case 42:
                                                                        v112 := GenS7{F0: 2147483647, F1: 227, F2: 255}
                                                                        _ = v112
                                                                    default:
                                                                        if len(v1) >= 2 {
                                                                            v1 = v1[1:len(v1)]
                                                                        }
                                                                    }
                                                                    v113 := GenS15{F0: 975}
                                                                    if v3 {
                                                                        switch "hello" {
                                                                        case "\n\t":
                                                                            _ = len(v2)
                                                                        case "hello":
                                                                            var v114 uint32 = 2147483647
                                                                            _ = v114
                                                                        case "":
                                                                            var v115 int32 = 1
                                                                            _ = v115
                                                                        }
                                                                        if (105 >= 42) {
                                                                            fmt.Println(v3, fi13, v3)
                                                                            fmt.Println(v3, fi13, v14)
                                                                            switch v21 {
                                                                            case "a":
                                                                                var v116 bool = v3
                                                                                _ = v116
                                                                            case "foo bar":
                                                                                var v117 int = v4
                                                                                _ = v117
                                                                            case "\n\t":
                                                                                v3 = (!(("foo bar" + v14) >= v14))
                                                                            }
                                                                        } else {
                                                                            v21 = v21
                                                                        }
                                                                    }
                                                                    fmt.Println(v3, fi13, v12)
                                                                    var v118 uint16 = 32767
                                                                    _ = v105
                                                                    _ = v113
                                                                    _ = v118
                                                                }
                                                                v119 := GenS16{F0: "foo bar"}
                                                                _ = len(v1)
                                                                var v120 bool = true
                                                                var v121 int16 = 32767
                                                                _ = v119
                                                                _ = v120
                                                                _ = v121
                                                            }
                                                        case "foo bar":
                                                            v3 = false
                                                        }
                                                        _ = v26
                                                    }
                                                    _ = v21
                                                } else {
                                                    switch 1 {
                                                    case 2147483647:
                                                        v122 := GenS7{F0: fi13, F1: 325, F2: 255}
                                                        _ = v122
                                                    case -2147483648:
                                                        var v123 int8 = -128
                                                        _ = v123
                                                    case 32768:
                                                        if v3 {
                                                            fmt.Println(v3, fi11, v4)
                                                            v3 = (v12 <= 1e10)
                                                            var v124 uint8 = 255
                                                            _ = v124
                                                        }
                                                    }
                                                }
                                                var v125 uint16 = 255
                                                _ = v125
                                            }
                                        }
                                        switch ("a" + "a") {
                                        case "hello":
                                            if (-32768 != 639) {
                                                var v126 uint8 = 1
                                                v14 = (v14 + "\n\t")
                                                v127 := GenS17{F0: 1, F1: 67}
                                                _ = v126
                                                _ = v127
                                            } else {
                                                if (1 == 0) {
                                                    if v3 {
                                                        v128 := GenS2{F0: 18446744073709551615}
                                                        _ = v128
                                                    }
                                                    switch fi11 {
                                                    case -2147483648:
                                                        v129 := v2[(fi13 % ((fi11) | 1))]
                                                        _ = v129
                                                    case 42:
                                                        v130 := GenS16{F0: v14}
                                                        _ = v130
                                                    case 0:
                                                        v14 = ""
                                                        v4 += (v4 ^ 42)
                                                        v131 := GenS18{F0: 671}
                                                        _ = v131
                                                    }
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                }
                                                v4 = (-(v4))
                                                switch -9223372036854775808 {
                                                case 42:
                                                    fmt.Println(fi13, fi13, fi11)
                                                }
                                                if false {
                                                    fmt.Println(fi11)
                                                } else {
                                                    v12 = (0.0 - v12)
                                                    if (42 >= 1) {
                                                        if v3 {
                                                            if true {
                                                                v14 = "foo bar"
                                                            } else {
                                                                fmt.Println(fi11, v12, fi13)
                                                                if false {
                                                                    if false {
                                                                        fmt.Println(fi13)
                                                                        v12 -= 1.5
                                                                        v12 = 3.14159
                                                                    }
                                                                    var v132 string = v14
                                                                    v12 = 1.5
                                                                    _ = v132
                                                                }
                                                                fmt.Println(v14, v12, v12)
                                                                v133 := v2[(^(fi11))]
                                                                _ = v133
                                                            }
                                                        } else {
                                                            v134 := GenS19{F0: false}
                                                            if (-470 != fi11) {
                                                                v4 -= (^(fi11))
                                                                switch (fi11 - fi13) {
                                                                case 216:
                                                                    v2[(fi13 + (fi13 - 2147483647))] = ("" + "a")
                                                                case -585:
                                                                    var v135 int32 = -1
                                                                    v136 := GenS20{F0: 100, F1: 0, F2: 0}
                                                                    _ = v135
                                                                    _ = v136
                                                                }
                                                                if false {
                                                                    v137 := GenS19{F0: true}
                                                                    if v3 {
                                                                        if (2 == 87) {
                                                                            switch 10 {
                                                                            case -485:
                                                                                switch (fi13 * (v4 - (v4 + (fi13 << ((65535) & 63))))) {
                                                                                case 32767:
                                                                                    var v138 int16 = -32768
                                                                                    _ = v138
                                                                                }
                                                                            case -753:
                                                                                var v139 bool = false
                                                                                _ = v139
                                                                            case -1:
                                                                                v12 = (float32(float64((float32(1e10) / v12))) + (((v12 / 1.5) / v12) * v12))
                                                                            }
                                                                            switch "" {
                                                                            case "hello":
                                                                                fmt.Println(fi13)
                                                                                v3 = (v12 < 0.0)
                                                                                switch "\n\t" {
                                                                                case "":
                                                                                    v4 = (-(fi11))
                                                                                case "foo bar":
                                                                                    var v140 uint64 = 2
                                                                                    _ = v140
                                                                                case "hello":
                                                                                    v12 = 0.1
                                                                                }
                                                                                var v141 int = (fi13 / (((fi11 + fi11)) | 1))
                                                                                if v3 {
                                                                                    var v142 uint = 42
                                                                                    _ = len(v1)
                                                                                    _ = v142
                                                                                } else {
                                                                                    if len(v1) > 0 {
                                                                                        v1[0] = (v14 + v14)
                                                                                    }
                                                                                    v14 = "foo bar"
                                                                                    _ = len(v2)
                                                                                }
                                                                                _ = v141
                                                                            }
                                                                        } else {
                                                                            v3 = ((v3 && v3) && v3)
                                                                            if (4294967295 < 0) {
                                                                                v14 = v14
                                                                            }
                                                                            v12 -= float32(3.14159)
                                                                        }
                                                                        v14 = v14
                                                                        var v143 int64 = 9223372036854775807
                                                                        v144 := GenS21{F0: fi11, F1: 128, F2: 683}
                                                                        _ = v143
                                                                        _ = v144
                                                                    } else {
                                                                        v145 := GenS22{F0: (v3 && v3), F1: 149}
                                                                        v2[(fi13 << ((118) & 63))] = ((v14 + v14) + "a")
                                                                        _ = v145
                                                                    }
                                                                    var v146 int32 = -1
                                                                    _ = v137
                                                                    _ = v146
                                                                }
                                                            }
                                                            _ = v134
                                                        }
                                                        if v3 {
                                                            fmt.Println(v12, v4)
                                                            fmt.Println(v3)
                                                        } else {
                                                            var v147 int = 32767
                                                            fmt.Println(fi11)
                                                            var v148 int8 = 94
                                                            _ = v147
                                                            _ = v148
                                                        }
                                                        if (false && false) {
                                                            v149 := GenS1{F0: -1, F1: 42}
                                                            _ = v149
                                                        }
                                                        if true {
                                                            var v150 int8 = -128
                                                            _ = v150
                                                        } else {
                                                            var v151 uint16 = 1
                                                            _ = v151
                                                        }
                                                        if v3 {
                                                            v152 := GenS22{F0: (((324 <= 1) && v3) || true), F1: 189}
                                                            _ = v152
                                                        }
                                                        v153 := GenS23{F0: (v4 - 1), F1: 2147483647, F2: ("a" + v14)}
                                                        var v154 uint32 = 477
                                                        _ = v153
                                                        _ = v154
                                                    } else {
                                                        v155 := GenS24{F0: -32768, F1: -32768, F2: 100}
                                                        switch (v14 + v14) {
                                                        case "":
                                                            delete(v2, v4)
                                                        case "\n\t":
                                                            if v3 {
                                                                v14 = v14
                                                            }
                                                        default:
                                                            var v156 uint = 947
                                                            _ = v156
                                                        }
                                                        _ = v155
                                                    }
                                                    v14 = ((v14 + (v14 + (v14 + "a"))) + v14)
                                                }
                                            }
                                        case "\n\t":
                                            v157 := GenS25{F0: fi13}
                                            _ = v157
                                        case "a":
                                            v14 = v14
                                        }
                                        v3 = ((fi11 << ((200) & 63)) != (fi13 & 32767))
                                    }
                                case "hello":
                                    if (v3 && (-1 < -32768)) {
                                        v3 = false
                                        fmt.Println(v12, fi13, v14)
                                        var v158 uint32 = 65535
                                        _ = v158
                                    } else {
                                        var v159 uint8 = 1
                                        _ = v159
                                    }
                                default:
                                    v160 := GenS6{F0: v14}
                                    _ = v160
                                }
                                if v3 {
                                    if false {
                                        v3 = (!((2147483647 == 2147483647) && (2 >= 255)))
                                        switch v14 {
                                        case "":
                                            if (!(555 != 255)) {
                                                v1 = append(v1, ((((("\n\t" + ("\n\t" + v14)) + "\n\t") + v14) + v14) + "hello"))
                                                var v161 int8 = 127
                                                var v162 int8 = 100
                                                _ = v161
                                                _ = v162
                                            }
                                            switch (fi11 ^ (v4 | (fi13 - 42))) {
                                            case -9223372036854775808:
                                                v4 = v4
                                            case 128:
                                                var v163 int32 = 0
                                                _ = v163
                                            case 9223372036854775807:
                                                v3 = (v3 && true)
                                            default:
                                                v14 = (v14 + "")
                                            }
                                            var v164 int8 = 77
                                            _ = v164
                                        default:
                                            if (!v3) {
                                                var v165 uint = 1
                                                v166 := GenS26{F0: 365, F1: 65535, F2: 0}
                                                _ = v165
                                                _ = v166
                                            } else {
                                                v14 = v14
                                                var v167 uint32 = 2
                                                _ = v167
                                            }
                                        }
                                        var v168 uint32 = 681
                                        _ = v168
                                    }
                                    if v3 {
                                        switch -817 {
                                        case 624:
                                            var v169 int8 = -128
                                            _ = v169
                                        case 128:
                                            var v170 uint32 = 1
                                            _ = v170
                                        case 32768:
                                            var v171 float64 = 1e10
                                            v172 := GenS6{F0: v14}
                                            _ = v171
                                            _ = v172
                                        default:
                                            var v173 int = fi11
                                            _ = v173
                                        }
                                        v12 -= -2.25
                                        var v174 uint8 = 42
                                        fmt.Println(fi13, v12)
                                        fmt.Println(v12, fi13, v174)
                                        _ = v174
                                    } else {
                                        if (v3 || v3) {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            v175 := GenS15{F0: 0}
                                            var v176 int64 = -918
                                            v177 := GenS8{F0: -1, F1: 1}
                                            if true {
                                                if false {
                                                    if true {
                                                        v14 = v14
                                                    }
                                                    if (false && (false && (((-2147483648 >= 128) && (42 == -1)) || v3))) {
                                                        v2[fi11] = v14
                                                        var v178 int32 = 2147483647
                                                        v179 := GenS4{F0: 13, F1: (v176 * v176), F2: (v178 | (v178 - v178))}
                                                        v4 += (fi13 / ((v4) | 1))
                                                        switch fi13 {
                                                        case -2147483648:
                                                            if (!(267 >= 450)) {
                                                                v180 := GenS27{F0: 42, F1: 75}
                                                                _ = v180
                                                            } else {
                                                                if false {
                                                                    if v3 {
                                                                        _ = len(v2)
                                                                        v181 := GenS21{F0: (v4 << ((791) & 63)), F1: 25, F2: 140}
                                                                        v4 += (fi13 - (fi11 - v4))
                                                                        _ = v181
                                                                    }
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                    fmt.Println(v4, v4)
                                                                    v4 = (fi13 & fi11)
                                                                } else {
                                                                    var v182 uint8 = 42
                                                                    if v3 {
                                                                        var v183 float64 = (1.5 / float64((v12 * v12)))
                                                                        _ = v183
                                                                    }
                                                                    v178 = v178
                                                                    _ = v182
                                                                }
                                                                if false {
                                                                    if ((fi13 + fi11) <= (v4 ^ int(v178))) {
                                                                        switch (v14 + v14) {
                                                                        case "hello":
                                                                            fmt.Println(v14)
                                                                        case "":
                                                                            var v184 int64 = (v176 | 111)
                                                                            _ = v184
                                                                        default:
                                                                            fmt.Println(fi13, v12)
                                                                        }
                                                                        v185 := GenS17{F0: (v178 >> ((42) & 63)), F1: 2147483648}
                                                                        switch (^(fi11)) {
                                                                        case 42:
                                                                            if v3 {
                                                                                v12 = v12
                                                                                if ((1000 <= -32768) || (true || v3)) {
                                                                                    v176 = (v176 + (^(v176)))
                                                                                    if (((1.5 * 0.1) / -2.25) <= (1.5 * (-2.25 + (0.0 + ((0.1 * 0.0) - 1e10))))) {
                                                                                        var v186 uint16 = 1
                                                                                        _ = v186
                                                                                    } else {
                                                                                        v187 := v2[fi11]
                                                                                        v3 = v3
                                                                                        _ = v187
                                                                                    }
                                                                                }
                                                                            }
                                                                        default:
                                                                            v12 = v12
                                                                        }
                                                                        _ = v185
                                                                    }
                                                                }
                                                                var v188 uint8 = 0
                                                                _ = v188
                                                            }
                                                        case 32768:
                                                            if len(v1) > 0 {
                                                                v1[0] = ""
                                                            }
                                                        default:
                                                            if len(v1) > 0 {
                                                                v1[0] = v14
                                                            }
                                                            if (-9223372036854775808 != (v176 / ((v176) | 1))) {
                                                                var v189 int = 1
                                                                v3 = (42 > 0)
                                                                if false {
                                                                    if (((-(v178)) != v178) && v3) {
                                                                        switch "a" {
                                                                        case "a":
                                                                            v12 -= 3.14159
                                                                            v2[v4] = v14
                                                                            if len(v1) >= 2 {
                                                                                v1 = v1[1:len(v1)]
                                                                            }
                                                                            var v190 uint16 = 65535
                                                                            _ = v190
                                                                        case "\n\t":
                                                                            var v191 int64 = v176
                                                                            _ = v191
                                                                        default:
                                                                            v178 += 42
                                                                        }
                                                                        v178 += v178
                                                                    }
                                                                    if (!v3) {
                                                                        var v192 uint32 = 42
                                                                        _ = v192
                                                                    }
                                                                    var v193 int64 = v176
                                                                    _ = v193
                                                                } else {
                                                                    var v194 int64 = v176
                                                                    var v195 int8 = -128
                                                                    v4 += (v4 % (((v4 | fi11)) | 1))
                                                                    _ = v194
                                                                    _ = v195
                                                                }
                                                                _ = v189
                                                            }
                                                            var v196 float64 = -2.25
                                                            _ = v196
                                                        }
                                                        v197 := v2[(^(fi13))]
                                                        fmt.Println(v197)
                                                        if v3 {
                                                            v3 = v3
                                                            delete(v2, (-(v4)))
                                                            v198 := GenS1{F0: 1000, F1: v178}
                                                            _ = v198
                                                        } else {
                                                            var v199 int16 = -1
                                                            _ = v199
                                                        }
                                                        var v200 bool = (2 < 972)
                                                        _ = v178
                                                        _ = v179
                                                        _ = v197
                                                        _ = v200
                                                    }
                                                    if v3 {
                                                        v201 := GenS17{F0: -802, F1: 0}
                                                        _ = v201
                                                    }
                                                }
                                                var v202 string = v14
                                                switch (fi11 & (^(fi11))) {
                                                case 32767:
                                                    if true {
                                                        if false {
                                                            fmt.Println(v176, v14, v14)
                                                            v3 = (v3 || (v3 && (((952 > (v176 << ((18446744073709551615) & 63))) || (!v3)) || v3)))
                                                            if (424 != 1) {
                                                                v203 := v2[v4]
                                                                _ = v203
                                                            }
                                                        }
                                                    }
                                                    v176 = (v176 % ((v176) | 1))
                                                    var v204 float64 = (0.0 * 0.0)
                                                    _ = v204
                                                default:
                                                    v205 := GenS28{F0: false, F1: v176, F2: -3}
                                                    _ = v205
                                                }
                                                _ = v202
                                            }
                                            _ = v175
                                            _ = v176
                                            _ = v177
                                        }
                                        v206 := GenS29{F0: 32767, F1: 0, F2: 2}
                                        _ = v206
                                    }
                                    v207 := GenS21{F0: (fi11 & 32767), F1: 0, F2: 2}
                                    _ = v207
                                } else {
                                    switch v14 {
                                    case "foo bar":
                                        var v208 uint16 = 1
                                        _ = v208
                                    case "":
                                        _ = len(v1)
                                    default:
                                        var v209 uint16 = 2
                                        _ = v209
                                    }
                                    fmt.Println(v12, v14, fi11)
                                    if v3 {
                                        v3 = (2 < 255)
                                        v4 = (fi11 ^ 244)
                                        var v210 uint8 = 1
                                        _ = v210
                                    } else {
                                        _ = len(v1)
                                    }
                                }
                                _ = v14
                            }
                        } else {
                            v2[(^(fi11))] = (("foo bar" + "") + "")
                        }
                        v4 = fi13
                        if v3 {
                            if len(v1) > 0 {
                                v1[0] = "a"
                            }
                            if (250 > -251) {
                                delete(v2, (v4 & (fi11 - 128)))
                                var v211 float64 = 0.0
                                if ((v211 * (((v211 - (v211 / (0.0 + 3.14159))) + v211) / (float64((v12 + v12)) - float64(float32(v211))))) <= 3.14159) {
                                    if v3 {
                                        v1 = append(v1, "\n\t")
                                        var v212 bool = (!(v3 && v3))
                                        if v3 {
                                            var v213 int64 = 42
                                            v212 = v212
                                            var v214 float64 = v211
                                            _ = v213
                                            _ = v214
                                        }
                                        _ = v212
                                    }
                                }
                                _ = v211
                            } else {
                                var v215 float32 = v12
                                _ = v215
                            }
                            v216 := GenS30{F0: 32, F1: "a", F2: 2}
                            _ = v216
                        }
                        v3 = true
                        _ = fi13
                    }
                    v217 := GenS27{F0: 871, F1: 197}
                    _ = v12
                    _ = v217
                }
                fmt.Println(v3, v3, v4)
                var v218 uint = 732
                _ = v218
            }
        default:
            v219 := GenS26{F0: 909, F1: 519, F2: 0}
            _ = v219
        }
        _ = fi11
    }
    switch 0 {
    case -1:
        v220 := GenS20{F0: 127, F1: 0, F2: 42}
        _ = v220
    }
    switch v4 {
    case -822:
        var v221 uint64 = 42
        _ = v221
    }
    if (2 > 968) {
        v3 = v3
    }
    _ = len(v1)
    for fi222 := 0; fi222 < 4; fi222++ {
        var v223 int8 = 1
        if false {
            v223 -= (v223 % ((v223) | 1))
        } else {
            if (289 != -1) {
                var v224 int32 = 2147483647
                delete(v2, 942)
                v4 -= (fi222 + 42)
                _ = v224
            } else {
                switch "hello" {
                case "foo bar":
                    v1 = append(v1, ("hello" + "foo bar"))
                case "a":
                    var v225 bool = ((v3 || (-9223372036854775808 == 1)) && (-104 > 1))
                    _ = v225
                case "hello":
                    if len(v1) > 0 {
                        v1[0] = ""
                    }
                default:
                    v3 = v3
                }
                v223 += 1
            }
        }
        _ = fi222
        _ = v223
    }
    if v3 {
        fmt.Println(v4, v3)
        v3 = (v3 || (318 != 9223372036854775807))
        var v226 bool = v3
        _ = v226
    } else {
        switch "" {
        case "\n\t":
            v227 := GenS29{F0: 32767, F1: 1, F2: 255}
            _ = v227
        }
        v3 = v3
    }
    v3 = v3
    for fi228 := 0; fi228 < 13; fi228++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v229 uint32 = 2147483648
        _ = fi228
        _ = v229
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
