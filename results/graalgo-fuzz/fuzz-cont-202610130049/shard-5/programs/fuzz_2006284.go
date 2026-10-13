package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 uint16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int64; F1 uint16 }
type GenS6 struct { F0 int; F1 int8; F2 int16 }
type GenS7 struct { F0 uint; F1 string; F2 int32 }
type GenS8 struct { F0 bool; F1 uint32 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 uint16; F1 uint16 }
type GenS11 struct { F0 string; F1 uint; F2 uint16 }
type GenS12 struct { F0 uint; F1 uint32 }
type GenS13 struct { F0 int64; F1 bool; F2 uint32 }
type GenS14 struct { F0 int64; F1 int8; F2 bool }
type GenS15 struct { F0 int32; F1 int64; F2 int64 }
type GenS16 struct { F0 bool; F1 int32; F2 uint8 }
type GenS17 struct { F0 uint32 }
type GenS18 struct { F0 int8 }
type GenS19 struct { F0 int16; F1 uint }
type GenS20 struct { F0 string }
type GenS21 struct { F0 int }
type GenS22 struct { F0 int16; F1 uint8; F2 int16 }
type GenS23 struct { F0 bool; F1 uint32; F2 int64 }
type GenS24 struct { F0 int32; F1 bool; F2 int64 }

func main() {
    var v1 uint16 = 772
    var v2 uint16 = v1
    v3 := []uint32{2147483648, 799, 2}
    var v4 int32 = 42
    var v5 int8 = 100
    v6 := []uint32{1, 42}
    v2 -= v2
    switch 5 {
    case 42:
        if (false && false) {
            var v7 bool = false
            for fi8 := 0; fi8 < 1; fi8++ {
                if v7 {
                    if v7 {
                        if (127 > 175) {
                            switch (fi8 << ((42) & 63)) {
                            case -144:
                                v5 -= v5
                            case 0:
                                for fi9 := 0; fi9 < 10; fi9++ {
                                    switch (("a" + ("" + "hello")) + "") {
                                    case "":
                                        switch fi8 {
                                        case 9223372036854775807:
                                            var v10 string = "foo bar"
                                            _ = v10
                                        case -9223372036854775808:
                                            v2 = v2
                                        }
                                    case "\n\t":
                                        v4 -= (v4 >> ((65535) & 63))
                                    default:
                                        if v7 {
                                            v6 = append(v6, 966)
                                            switch -2147483648 {
                                            case 1:
                                                switch 42 {
                                                case 1:
                                                    var v11 int16 = -32768
                                                    _ = v11
                                                case 0:
                                                    var v12 uint16 = v2
                                                    if (218 == 122) {
                                                        var v13 uint16 = (^(v2))
                                                        v7 = v7
                                                        var v14 bool = (556 == 0)
                                                        _ = v13
                                                        _ = v14
                                                    }
                                                    var v15 string = "hello"
                                                    v16 := GenS1{F0: 63}
                                                    _ = v12
                                                    _ = v15
                                                    _ = v16
                                                case -9223372036854775808:
                                                    var v17 int8 = v5
                                                    _ = v17
                                                }
                                            case 128:
                                                if ("\n\t" == ("" + ("a" + "a"))) {
                                                    v18 := GenS1{F0: 791}
                                                    if false {
                                                        v2 = 32767
                                                        var v19 float64 = (((-2.25 * float64(0.1)) / float64(1.5)) * (0.1 - 1.5))
                                                        v20 := GenS1{F0: 42}
                                                        _ = v19
                                                        _ = v20
                                                    } else {
                                                        if v7 {
                                                            var v21 uint8 = 255
                                                            _ = v21
                                                        }
                                                    }
                                                    var v22 float32 = (-2.25 * (-2.25 * -2.25))
                                                    _ = v18
                                                    _ = v22
                                                }
                                            case 0:
                                                if (false || v7) {
                                                    if v7 {
                                                        var v23 float32 = 3.14159
                                                        v2 += 66
                                                        _ = v23
                                                    } else {
                                                        v7 = v7
                                                        var v24 uint32 = 2147483648
                                                        var v25 uint8 = 157
                                                        if v7 {
                                                            var v26 uint32 = v24
                                                            _ = v26
                                                        } else {
                                                            v27 := GenS1{F0: (^(v24))}
                                                            if false {
                                                                switch ("foo bar" + "a") {
                                                                case "\n\t":
                                                                    if false {
                                                                        var v28 string = ""
                                                                        _ = v28
                                                                    } else {
                                                                        if ("\n\t" != "foo bar") {
                                                                            var v29 uint = 1
                                                                            var v30 uint = (^(v29))
                                                                            var v31 float64 = ((float64((1.5 * (((1e10 / 1e10) - ((0.0 / 0.1) + 3.14159)) * 0.0))) - 3.14159) * 0.1)
                                                                            v32 := GenS2{F0: 211, F1: 536}
                                                                            var v33 int32 = -2147483648
                                                                            if false {
                                                                                v34 := GenS1{F0: (v24 / (((^(v24))) | 1))}
                                                                                var v35 int32 = v4
                                                                                if len(v6) > 0 {
                                                                                    v6[0] = 2147483647
                                                                                }
                                                                                if v7 {
                                                                                    v7 = v7
                                                                                } else {
                                                                                    v36 := GenS3{F0: (v29 % ((628) | 1))}
                                                                                    if len(v6) >= 2 {
                                                                                        v6 = v6[1:len(v6)]
                                                                                    }
                                                                                    if true {
                                                                                        fmt.Println(v24)
                                                                                        v29 = v29
                                                                                    } else {
                                                                                        if len(v6) > 0 {
                                                                                            v6[0] = v24
                                                                                        }
                                                                                        if (0.0 >= 1e10) {
                                                                                            if (42 <= 0) {
                                                                                                switch fi8 {
                                                                                                case 127:
                                                                                                    fmt.Println(v25, v7, v2)
                                                                                                case -9223372036854775808:
                                                                                                    v29 += uint(v24)
                                                                                                }
                                                                                                var v37 uint16 = v1
                                                                                                if len(v3) >= 2 {
                                                                                                    v3 = v3[1:len(v3)]
                                                                                                }
                                                                                                v35 += (v4 >> (((v29 + 65535)) & 63))
                                                                                                _ = v37
                                                                                            } else {
                                                                                                v3 = append(v3, (v24 - (v24 - v24)))
                                                                                                v38 := GenS1{F0: v24}
                                                                                                var v39 uint = (v30 + (v30 - v30))
                                                                                                _ = v38
                                                                                                _ = v39
                                                                                            }
                                                                                            fmt.Println(v30, v24)
                                                                                            v40 := GenS2{F0: 799, F1: v2}
                                                                                            _ = v40
                                                                                        }
                                                                                        if true {
                                                                                            var v41 string = ("a" + "")
                                                                                            var v42 int32 = (v33 * -1)
                                                                                            _ = v41
                                                                                            _ = v42
                                                                                        } else {
                                                                                            _ = len(v6)
                                                                                            var v43 bool = (!false)
                                                                                            _ = v43
                                                                                        }
                                                                                    }
                                                                                    _ = v36
                                                                                }
                                                                                if ((v24 / (((v24 / ((v24) | 1))) | 1)) >= 176) {
                                                                                    var v44 int32 = (^(v35))
                                                                                    _ = v44
                                                                                }
                                                                                _ = v34
                                                                                _ = v35
                                                                            } else {
                                                                                v31 = v31
                                                                            }
                                                                            _ = v29
                                                                            _ = v30
                                                                            _ = v31
                                                                            _ = v32
                                                                            _ = v33
                                                                        } else {
                                                                            if v7 {
                                                                                if len(v6) >= 2 {
                                                                                    v6 = v6[1:len(v6)]
                                                                                }
                                                                                var v45 float64 = 1e10
                                                                                _ = v45
                                                                            }
                                                                        }
                                                                        var v46 int32 = (v4 * v4)
                                                                        _ = v46
                                                                    }
                                                                default:
                                                                    if v7 {
                                                                        var v47 uint = 18446744073709551615
                                                                        _ = v47
                                                                    } else {
                                                                        fmt.Println(v4, fi8)
                                                                        v48 := GenS4{F0: ((v7 && ((v5 | (v5 - (v5 << ((4294967295) & 63)))) <= (v5 - (v5 % ((v5) | 1))))) && true)}
                                                                        var v49 int = fi9
                                                                        _ = v48
                                                                        _ = v49
                                                                    }
                                                                }
                                                            }
                                                            _ = v27
                                                        }
                                                        var v50 float64 = 0.0
                                                        v2 = v1
                                                        var v51 int16 = 936
                                                        if v7 {
                                                            v5 = (-(v5))
                                                            switch ("\n\t" + "hello") {
                                                            case "":
                                                                switch "hello" {
                                                                case "foo bar":
                                                                    var v52 int64 = -49
                                                                    _ = v52
                                                                default:
                                                                    v50 -= v50
                                                                }
                                                            case "\n\t":
                                                                var v53 uint64 = 0
                                                                _ = v53
                                                            case "hello":
                                                                fmt.Println(v24, v2, v50)
                                                            default:
                                                                var v54 uint32 = v24
                                                                _ = v54
                                                            }
                                                            v1 -= 255
                                                        } else {
                                                            fmt.Println(fi9)
                                                        }
                                                        v3 = append(v3, 2)
                                                        v2 = 0
                                                        v55 := GenS2{F0: 0, F1: 255}
                                                        if ((!false) && v7) {
                                                            v51 += v51
                                                        } else {
                                                            v24 -= 134
                                                        }
                                                        _ = v24
                                                        _ = v25
                                                        _ = v50
                                                        _ = v51
                                                        _ = v55
                                                    }
                                                    fmt.Println(v5, v5)
                                                    if (0 > 1) {
                                                        if (v2 != 672) {
                                                            v56 := GenS5{F0: 1, F1: 32767}
                                                            _ = v56
                                                        } else {
                                                            if v7 {
                                                                fmt.Println(v5)
                                                            } else {
                                                                if len(v3) > 0 {
                                                                    v3[0] = 686
                                                                }
                                                                _ = len(v6)
                                                                if v7 {
                                                                    var v57 float32 = 0.0
                                                                    _ = v57
                                                                }
                                                            }
                                                        }
                                                        switch -976 {
                                                        case -431:
                                                            if v7 {
                                                                if ((v7 && (!v7)) || false) {
                                                                    if v7 {
                                                                        if v7 {
                                                                            fmt.Println(fi8)
                                                                            if false {
                                                                                var v58 int8 = v5
                                                                                _ = v58
                                                                            } else {
                                                                                fmt.Println(v7, fi9)
                                                                            }
                                                                            v5 -= 42
                                                                            fmt.Println(v2, v5, fi9)
                                                                            var v59 float64 = (3.14159 * -2.25)
                                                                            var v60 int64 = -645
                                                                            var v61 int8 = 1
                                                                            if ((v60 != v60) && v7) {
                                                                                var v62 int16 = -32768
                                                                                _ = v62
                                                                            } else {
                                                                                var v63 uint16 = v2
                                                                                var v64 uint = 926
                                                                                var v65 int64 = (v60 ^ 42)
                                                                                _ = v63
                                                                                _ = v64
                                                                                _ = v65
                                                                            }
                                                                            _ = v59
                                                                            _ = v60
                                                                            _ = v61
                                                                        } else {
                                                                            v66 := GenS3{F0: 351}
                                                                            v5 += v5
                                                                            if v7 {
                                                                                _ = len(v6)
                                                                                var v67 int8 = v5
                                                                                var v68 int16 = 512
                                                                                var v69 int64 = 9223372036854775807
                                                                                var v70 int8 = 78
                                                                                _ = v67
                                                                                _ = v68
                                                                                _ = v69
                                                                                _ = v70
                                                                            } else {
                                                                                v3 = append(v3, 401)
                                                                                var v71 int = fi9
                                                                                _ = v71
                                                                            }
                                                                            _ = v66
                                                                        }
                                                                    } else {
                                                                        v7 = v7
                                                                        v5 = -87
                                                                    }
                                                                    switch ("" + "a") {
                                                                    case "\n\t":
                                                                        switch (fi8 - (fi9 & fi8)) {
                                                                        case 32768:
                                                                            if len(v6) > 0 {
                                                                                v6[0] = 2
                                                                            }
                                                                        case 32767:
                                                                            if (702 <= 2) {
                                                                                if (fi8 <= fi8) {
                                                                                    if len(v3) >= 2 {
                                                                                        v3 = v3[1:len(v3)]
                                                                                    }
                                                                                }
                                                                                v2 += v2
                                                                            }
                                                                        }
                                                                        var v72 string = ("foo bar" + ("" + "\n\t"))
                                                                        _ = v72
                                                                    case "hello":
                                                                        switch "foo bar" {
                                                                        case "foo bar":
                                                                            switch (fi8 << ((0) & 63)) {
                                                                            case 942:
                                                                                v73 := GenS1{F0: 65535}
                                                                                _ = v73
                                                                            case -112:
                                                                                if (int16(v5) > -709) {
                                                                                    fmt.Println(v1, v2)
                                                                                    switch "" {
                                                                                    case "\n\t":
                                                                                        if len(v3) > 0 {
                                                                                            v3[0] = 1
                                                                                        }
                                                                                    case "hello":
                                                                                        v7 = v7
                                                                                    default:
                                                                                        var v74 uint8 = 123
                                                                                        _ = v74
                                                                                    }
                                                                                    if len(v6) > 0 {
                                                                                        v6[0] = 4294967295
                                                                                    }
                                                                                    fmt.Println(fi9)
                                                                                    var v75 uint32 = 315
                                                                                    _ = v75
                                                                                }
                                                                                v7 = v7
                                                                                if v7 {
                                                                                    v76 := GenS6{F0: fi9, F1: v5, F2: -32768}
                                                                                    if len(v6) >= 2 {
                                                                                        v6 = v6[1:len(v6)]
                                                                                    }
                                                                                    _ = v76
                                                                                }
                                                                            default:
                                                                                v77 := GenS7{F0: 0, F1: "\n\t", F2: 42}
                                                                                _ = v77
                                                                            }
                                                                            v1 += v2
                                                                            var v78 float64 = 1.5
                                                                            _ = v78
                                                                        case "hello":
                                                                            v79 := GenS3{F0: 42}
                                                                            if v7 {
                                                                                v80 := GenS5{F0: 9223372036854775807, F1: v2}
                                                                                v5 = (v5 ^ v5)
                                                                                fmt.Println(v2)
                                                                                _ = v80
                                                                            } else {
                                                                                var v81 string = "hello"
                                                                                var v82 int16 = 32767
                                                                                fmt.Println(v7, v1, v4)
                                                                                _ = v81
                                                                                _ = v82
                                                                            }
                                                                            v1 += v2
                                                                            _ = v79
                                                                        case "\n\t":
                                                                            fmt.Println(fi9, v1, v4)
                                                                        default:
                                                                            fmt.Println(v5, v5, v1)
                                                                        }
                                                                    case "foo bar":
                                                                        v2 += v1
                                                                    }
                                                                }
                                                                if true {
                                                                    v6 = append(v6, 888)
                                                                    var v83 uint8 = 1
                                                                    _ = v83
                                                                } else {
                                                                    fmt.Println(v4, fi9)
                                                                    _ = len(v6)
                                                                }
                                                            } else {
                                                                var v84 uint8 = 42
                                                                if v7 {
                                                                    v85 := GenS8{F0: (!((-32768 > 1) && ((v7 && false) || v7))), F1: 371}
                                                                    var v86 int16 = -112
                                                                    if (false && (0.1 <= float64(((3.14159 * -2.25) * -2.25)))) {
                                                                        v5 -= (v5 + v5)
                                                                        if (!((798 != 624) && false)) {
                                                                            switch ((("foo bar" + "foo bar") + ("a" + (((("" + "a") + "a") + "\n\t") + "hello"))) + "foo bar") {
                                                                            case "foo bar":
                                                                                if false {
                                                                                    if v7 {
                                                                                        fmt.Println(v2, v7)
                                                                                    } else {
                                                                                        if len(v3) > 0 {
                                                                                            v3[0] = 132
                                                                                        }
                                                                                        switch "a" {
                                                                                        case "\n\t":
                                                                                            if len(v6) > 0 {
                                                                                                v6[0] = 0
                                                                                            }
                                                                                        case "hello":
                                                                                            var v87 uint64 = 243
                                                                                            _ = v87
                                                                                        default:
                                                                                            var v88 string = ("" + ("hello" + "\n\t"))
                                                                                            _ = v88
                                                                                        }
                                                                                    }
                                                                                    if v7 {
                                                                                        var v89 uint = 2
                                                                                        if len(v3) >= 2 {
                                                                                            v3 = v3[1:len(v3)]
                                                                                        }
                                                                                        _ = v89
                                                                                    } else {
                                                                                        fmt.Println(v84)
                                                                                        var v90 int32 = v4
                                                                                        _ = v90
                                                                                    }
                                                                                    fmt.Println(fi9, v84, v2)
                                                                                    switch fi9 {
                                                                                    case 2147483647:
                                                                                        v84 = v84
                                                                                    case -9223372036854775808:
                                                                                        var v91 int8 = v5
                                                                                        _ = v91
                                                                                    }
                                                                                    v84 += v84
                                                                                    var v92 int32 = (v4 & v4)
                                                                                    v4 = v4
                                                                                    _ = v92
                                                                                } else {
                                                                                    if len(v3) >= 2 {
                                                                                        v3 = v3[1:len(v3)]
                                                                                    }
                                                                                }
                                                                            }
                                                                        }
                                                                        if (true && (int8(v4) > v5)) {
                                                                            if v7 {
                                                                                if v7 {
                                                                                    v3 = append(v3, 2147483648)
                                                                                    v93 := GenS9{F0: "\n\t"}
                                                                                    _ = v93
                                                                                }
                                                                                if (!v7) {
                                                                                    v5 -= (v5 ^ (-(v5)))
                                                                                }
                                                                                var v94 uint16 = v1
                                                                                _ = v94
                                                                            } else {
                                                                                v95 := GenS6{F0: fi9, F1: int8(fi9), F2: 0}
                                                                                if (v4 < (^(v4))) {
                                                                                    v3 = append(v3, 4294967295)
                                                                                    if ((v7 && (-2.25 != (-2.25 / 1e10))) && (!(fi8 < 42))) {
                                                                                        if (469 <= 272) {
                                                                                            v3 = append(v3, 729)
                                                                                            if v7 {
                                                                                                v3 = append(v3, 2147483647)
                                                                                                fmt.Println(fi9, v7)
                                                                                                v7 = v7
                                                                                            }
                                                                                            v96 := GenS5{F0: -329, F1: 934}
                                                                                            var v97 int64 = -932
                                                                                            _ = v96
                                                                                            _ = v97
                                                                                        }
                                                                                        var v98 int64 = -1
                                                                                        var v99 int16 = v86
                                                                                        _ = v98
                                                                                        _ = v99
                                                                                    }
                                                                                }
                                                                                if true {
                                                                                    v100 := GenS1{F0: 2}
                                                                                    v86 = int16(v84)
                                                                                    var v101 int8 = -128
                                                                                    _ = v100
                                                                                    _ = v101
                                                                                } else {
                                                                                    if false {
                                                                                        if (("hello" + "a") >= "a") {
                                                                                            if true {
                                                                                                if (2147483648 >= 2147483647) {
                                                                                                    v102 := GenS6{F0: (fi9 ^ (-(fi8))), F1: v5, F2: v86}
                                                                                                    fmt.Println(fi8)
                                                                                                    _ = v102
                                                                                                } else {
                                                                                                    switch "\n\t" {
                                                                                                    case "hello":
                                                                                                        v1 += v1
                                                                                                    case "a":
                                                                                                        if true {
                                                                                                            if len(v3) >= 2 {
                                                                                                                v3 = v3[1:len(v3)]
                                                                                                            }
                                                                                                            v103 := GenS5{F0: -604, F1: (v2 & 2)}
                                                                                                            _ = v103
                                                                                                        } else {
                                                                                                            _ = len(v6)
                                                                                                        }
                                                                                                    }
                                                                                                    fmt.Println(v1)
                                                                                                }
                                                                                                v104 := GenS2{F0: 887, F1: v2}
                                                                                                if len(v3) >= 2 {
                                                                                                    v3 = v3[1:len(v3)]
                                                                                                }
                                                                                                _ = v104
                                                                                            } else {
                                                                                                var v105 int = -9223372036854775808
                                                                                                _ = v105
                                                                                            }
                                                                                        } else {
                                                                                            fmt.Println(v86)
                                                                                            v106 := GenS10{F0: v2, F1: v1}
                                                                                            v107 := GenS11{F0: "", F1: 922, F2: (v1 ^ 141)}
                                                                                            _ = v106
                                                                                            _ = v107
                                                                                        }
                                                                                    }
                                                                                    switch "\n\t" {
                                                                                    case "foo bar":
                                                                                        v2 = v1
                                                                                    default:
                                                                                        switch "hello" {
                                                                                        case "\n\t":
                                                                                            switch int(v4) {
                                                                                            case -777:
                                                                                                switch fi8 {
                                                                                                case -58:
                                                                                                    v1 -= 0
                                                                                                case 127:
                                                                                                    v108 := GenS3{F0: 0}
                                                                                                    var v109 bool = true
                                                                                                    _ = v108
                                                                                                    _ = v109
                                                                                                case 128:
                                                                                                    switch 0 {
                                                                                                    case -1:
                                                                                                        switch ("\n\t" + "\n\t") {
                                                                                                        case "hello":
                                                                                                            fmt.Println(v5, v4, v86)
                                                                                                        case "\n\t":
                                                                                                            var v110 int32 = int32(v2)
                                                                                                            _ = v110
                                                                                                        case "foo bar":
                                                                                                            if (v7 || v7) {
                                                                                                                var v111 float64 = (3.14159 * 0.0)
                                                                                                                _ = v111
                                                                                                            }
                                                                                                        }
                                                                                                    case -777:
                                                                                                        fmt.Println(v84, v2)
                                                                                                    case 1:
                                                                                                        if v7 {
                                                                                                            var v112 float64 = 0.1
                                                                                                            v113 := GenS12{F0: 255, F1: 519}
                                                                                                            v114 := GenS11{F0: "a", F1: 4294967295, F2: v1}
                                                                                                            _ = v112
                                                                                                            _ = v113
                                                                                                            _ = v114
                                                                                                        } else {
                                                                                                            var v115 bool = true
                                                                                                            var v116 bool = v7
                                                                                                            v117 := GenS6{F0: fi8, F1: v5, F2: (^(v86))}
                                                                                                            _ = v115
                                                                                                            _ = v116
                                                                                                            _ = v117
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                        case "a":
                                                                                            v1 = (v1 + (v1 % ((v2) | 1)))
                                                                                            v4 += (v4 % ((-50) | 1))
                                                                                        default:
                                                                                            fmt.Println(v84, v7, v84)
                                                                                        }
                                                                                    }
                                                                                    var v118 int64 = 0
                                                                                    v86 = -32768
                                                                                    _ = v118
                                                                                }
                                                                                _ = v95
                                                                            }
                                                                            v7 = false
                                                                            switch "" {
                                                                            case "hello":
                                                                                if false {
                                                                                    if (float64(3.14159) < 3.14159) {
                                                                                        v3 = append(v3, 638)
                                                                                    }
                                                                                } else {
                                                                                    if true {
                                                                                        if v7 {
                                                                                            fmt.Println(v5)
                                                                                            _ = len(v6)
                                                                                            var v119 int8 = v5
                                                                                            _ = v119
                                                                                        }
                                                                                    } else {
                                                                                        var v120 float32 = (((1e10 * ((((float32(1e10) - ((1.5 / 1e10) - -2.25)) - -2.25) / float32(float64(1e10))) / 1.0)) / (1.5 - 0.0)) + (1.5 + (1.5 * -2.25)))
                                                                                        _ = v120
                                                                                    }
                                                                                    var v121 int = (fi8 >> ((74) & 63))
                                                                                    fmt.Println(v4, v1)
                                                                                    v5 = int8(v86)
                                                                                    _ = v121
                                                                                }
                                                                            case "\n\t":
                                                                                v3 = append(v3, 42)
                                                                            case "":
                                                                                v122 := GenS12{F0: 4294967295, F1: 65535}
                                                                                _ = v122
                                                                            default:
                                                                                fmt.Println(v5, v7)
                                                                            }
                                                                        } else {
                                                                            v123 := GenS2{F0: uint64(v1), F1: (^(v1))}
                                                                            v7 = (!((1e10 + 1e10) <= 3.14159))
                                                                            var v124 uint8 = (v84 + (v84 & v84))
                                                                            var v125 uint32 = 2
                                                                            if (!(v7 || true)) {
                                                                                if (false && false) {
                                                                                    v124 -= v84
                                                                                    var v126 uint32 = 42
                                                                                    _ = v126
                                                                                }
                                                                                if ((^(v86)) <= (v86 & v86)) {
                                                                                    switch -498 {
                                                                                    case 127:
                                                                                        v7 = (v86 != v86)
                                                                                    default:
                                                                                        switch 757 {
                                                                                        case 886:
                                                                                            _ = len(v3)
                                                                                        case -1:
                                                                                            fmt.Println(v125)
                                                                                        case 2147483647:
                                                                                            v127 := GenS8{F0: true, F1: (v125 & (v125 % ((2147483648) | 1)))}
                                                                                            _ = v127
                                                                                        }
                                                                                    }
                                                                                } else {
                                                                                    var v128 int8 = (-(v5))
                                                                                    v129 := GenS4{F0: v7}
                                                                                    var v130 uint16 = v2
                                                                                    var v131 uint = 482
                                                                                    fmt.Println(v131, v130)
                                                                                    var v132 uint = (v131 + (v131 ^ v131))
                                                                                    var v133 uint16 = v1
                                                                                    if (v2 == (v133 - v2)) {
                                                                                        switch (fi8 ^ (fi9 | 128)) {
                                                                                        case -1:
                                                                                            var v134 uint16 = 65535
                                                                                            _ = v134
                                                                                        }
                                                                                        v135 := GenS9{F0: ("\n\t" + "foo bar")}
                                                                                        _ = v135
                                                                                    } else {
                                                                                        if len(v6) >= 2 {
                                                                                            v6 = v6[1:len(v6)]
                                                                                        }
                                                                                    }
                                                                                    v86 = (v86 ^ (v86 + 42))
                                                                                    var v136 float32 = ((1.5 + 1e10) + 0.0)
                                                                                    if v7 {
                                                                                        if ((!(1.5 < (0.1 + -2.25))) || v7) {
                                                                                            fmt.Println(v133, v84, fi8)
                                                                                        }
                                                                                    } else {
                                                                                        v128 = v5
                                                                                        v128 -= v128
                                                                                    }
                                                                                    _ = v128
                                                                                    _ = v129
                                                                                    _ = v130
                                                                                    _ = v131
                                                                                    _ = v132
                                                                                    _ = v133
                                                                                    _ = v136
                                                                                }
                                                                            }
                                                                            _ = v123
                                                                            _ = v124
                                                                            _ = v125
                                                                        }
                                                                    }
                                                                    _ = v85
                                                                    _ = v86
                                                                }
                                                                _ = v84
                                                            }
                                                        case 9223372036854775807:
                                                            v137 := GenS4{F0: v7}
                                                            _ = v137
                                                        }
                                                        var v138 float64 = 0.0
                                                        _ = v138
                                                    }
                                                } else {
                                                    v4 += -2147483648
                                                    switch -367 {
                                                    case 128:
                                                        switch (^(fi9)) {
                                                        case 4294967295:
                                                            switch (fi9 * (-(fi8))) {
                                                            case 9223372036854775807:
                                                                var v139 int64 = 977
                                                                _ = v139
                                                            default:
                                                                var v140 float64 = 1e10
                                                                _ = v140
                                                            }
                                                            fmt.Println(v2, v4)
                                                        default:
                                                            v141 := GenS1{F0: 0}
                                                            _ = v141
                                                        }
                                                    case 42:
                                                        v142 := GenS6{F0: -2147483648, F1: (v5 ^ (v5 & v5)), F2: 42}
                                                        _ = v142
                                                    case -2147483648:
                                                        if v7 {
                                                            v2 += (v1 * v1)
                                                        }
                                                    }
                                                    fmt.Println(fi8, fi8)
                                                }
                                            default:
                                                var v143 int32 = 1
                                                _ = v143
                                            }
                                            switch "\n\t" {
                                            case "a":
                                                switch (fi8 << ((2) & 63)) {
                                                case -774:
                                                    var v144 uint32 = 2147483647
                                                    _ = v144
                                                }
                                            case "hello":
                                                var v145 int32 = (v4 & (v4 & 587))
                                                _ = v145
                                            }
                                        }
                                    }
                                    switch -770 {
                                    case 1:
                                        v4 += (^(v4))
                                    case 9223372036854775807:
                                        switch ("" + "") {
                                        case "":
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                        case "\n\t":
                                            v146 := GenS2{F0: 407, F1: (v1 & (v1 & v1))}
                                            v147 := GenS13{F0: 9223372036854775807, F1: true, F2: 2147483648}
                                            if v7 {
                                                var v148 int8 = (v5 >> ((2) & 63))
                                                v4 = 42
                                                _ = v148
                                            } else {
                                                var v149 int64 = int64(v5)
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                if len(v6) > 0 {
                                                    v6[0] = 187
                                                }
                                                v7 = true
                                                v150 := GenS14{F0: (v149 / ((int64(v1)) | 1)), F1: (-(v5)), F2: v7}
                                                _ = v149
                                                _ = v150
                                            }
                                            _ = v146
                                            _ = v147
                                        case "foo bar":
                                            v6 = append(v6, 507)
                                        default:
                                            v151 := GenS15{F0: (v4 * -234), F1: 1, F2: 9223372036854775807}
                                            _ = v151
                                        }
                                    }
                                    switch fi9 {
                                    case 32767:
                                        v1 = 83
                                    case 127:
                                        if ((v5 << ((4294967295) & 63)) == v5) {
                                            v2 = v1
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            if v7 {
                                                var v152 uint32 = 0
                                                v4 -= (v4 << ((4294967295) & 63))
                                                switch fi9 {
                                                case 0:
                                                    var v153 int16 = 42
                                                    _ = v153
                                                case 211:
                                                    if ((0 < -242) && ((-1 >= 0) && ((0.0 - ((1.5 + 0.0) * (0.1 / 3.14159))) != 3.14159))) {
                                                        switch (fi8 | 819) {
                                                        case 0:
                                                            switch (fi9 - (fi9 ^ 1)) {
                                                            case 4294967295:
                                                                if (1 >= 42) {
                                                                    switch (("\n\t" + (("hello" + "\n\t") + "a")) + "") {
                                                                    case "\n\t":
                                                                        var v154 int32 = (v4 >> ((4294967295) & 63))
                                                                        if len(v6) > 0 {
                                                                            v6[0] = (^(v152))
                                                                        }
                                                                        _ = v154
                                                                    default:
                                                                        var v155 uint32 = v152
                                                                        _ = v155
                                                                    }
                                                                    v2 = v1
                                                                    var v156 uint16 = (v2 & v2)
                                                                    _ = v156
                                                                }
                                                            case 0:
                                                                if (4294967295 >= 370) {
                                                                    switch (fi9 ^ 42) {
                                                                    case -155:
                                                                        var v157 uint16 = (^(v2))
                                                                        _ = v157
                                                                    case 32767:
                                                                        v6 = append(v6, v152)
                                                                    default:
                                                                        v158 := GenS16{F0: false, F1: -2147483648, F2: 2}
                                                                        _ = v158
                                                                    }
                                                                } else {
                                                                    if false {
                                                                        var v159 string = "\n\t"
                                                                        v1 = (v2 + (v1 - v1))
                                                                        _ = v159
                                                                    } else {
                                                                        v160 := GenS17{F0: (^(v152))}
                                                                        _ = v160
                                                                    }
                                                                    _ = len(v3)
                                                                    var v161 int8 = v5
                                                                    _ = v161
                                                                }
                                                            case 9223372036854775807:
                                                                if len(v6) > 0 {
                                                                    v6[0] = v152
                                                                }
                                                            default:
                                                                v7 = v7
                                                            }
                                                        case 2147483647:
                                                            v5 = v5
                                                        case -9223372036854775808:
                                                            var v162 int16 = -32768
                                                            _ = v162
                                                        default:
                                                            if (v7 || ((1e10 + ((-2.25 / (-2.25 + float32(0.1))) + 0.0)) < (float32(((1e10 * (float64(((-2.25 / 1.0) - 3.14159)) / 1e10)) / 1.5)) + 0.0))) {
                                                                if (int8(v152) < (v5 % (((v5 % ((v5) | 1))) | 1))) {
                                                                    v163 := GenS18{F0: (v5 % ((32) | 1))}
                                                                    v2 = 167
                                                                    _ = v163
                                                                }
                                                                v2 += 255
                                                                v164 := GenS19{F0: 620, F1: 44}
                                                                _ = v164
                                                            }
                                                        }
                                                        v152 = (v152 ^ 4294967295)
                                                    } else {
                                                        _ = len(v3)
                                                        v5 -= (v5 & -42)
                                                    }
                                                default:
                                                    switch fi9 {
                                                    case -53:
                                                        v4 = (v4 / ((-567) | 1))
                                                    }
                                                }
                                                v1 += v1
                                                v165 := GenS2{F0: 42, F1: 42}
                                                _ = v152
                                                _ = v165
                                            }
                                            var v166 uint8 = 69
                                            _ = v166
                                        }
                                    case -1:
                                        switch "" {
                                        case "foo bar":
                                            fmt.Println(v1)
                                        default:
                                            var v167 uint8 = 6
                                            _ = v167
                                        }
                                    }
                                    v168 := GenS19{F0: 1000, F1: 619}
                                    if (("\n\t" + "hello") == "foo bar") {
                                        var v169 uint32 = 2147483648
                                        v169 = v169
                                        v7 = false
                                        v4 -= v4
                                        _ = v169
                                    }
                                    _ = fi9
                                    _ = v168
                                }
                            case -9223372036854775808:
                                var v170 int8 = v5
                                _ = v170
                            default:
                                for fi171 := 0; fi171 < 4; fi171++ {
                                    v4 += (v4 % ((v4) | 1))
                                    _ = len(v3)
                                    v172 := GenS20{F0: "\n\t"}
                                    _ = fi171
                                    _ = v172
                                }
                            }
                        } else {
                            var v173 uint32 = 2147483648
                            v2 += (v2 << ((65535) & 63))
                            v174 := GenS21{F0: (fi8 / (((fi8 & (fi8 >> ((113) & 63)))) | 1))}
                            _ = v173
                            _ = v174
                        }
                        var v175 uint32 = 903
                        var v176 uint16 = v1
                        _ = v175
                        _ = v176
                    } else {
                        var v177 string = "\n\t"
                        _ = v177
                    }
                    fmt.Println(v1)
                }
                _ = fi8
            }
            _ = v7
        } else {
            v2 += (v1 % ((v2) | 1))
            var v178 bool = true
            if true {
                v6 = append(v6, 0)
                v5 -= 12
            }
            _ = v178
        }
    default:
        v179 := GenS22{F0: -616, F1: 42, F2: -32768}
        _ = v179
    }
    v1 = v2
    v180 := GenS23{F0: (255 != (v2 ^ v2)), F1: uint32(v2), F2: 293}
    switch "\n\t" {
    case "hello":
        if len(v6) > 0 {
            v6[0] = 65535
        }
    }
    v2 = (v1 >> ((0) & 63))
    fmt.Println(v5, v4)
    if true {
        switch 2147483647 {
        case 4294967295:
            var v181 uint32 = 42
            _ = v181
        default:
            v3 = append(v3, 2147483647)
        }
    } else {
        var v182 uint32 = 432
        if (1 > (v2 * 255)) {
            v4 = (v4 + v4)
            switch "hello" {
            case "a":
                switch ("hello" + "foo bar") {
                case "a":
                    v6 = append(v6, 4294967295)
                case "\n\t":
                    var v183 int = 293
                    _ = v183
                }
            case "hello":
                v184 := GenS24{F0: (^(v4)), F1: false, F2: 42}
                _ = v184
            }
        }
        var v185 string = "hello"
        _ = v182
        _ = v185
    }
    v1 = (v2 << ((4294967295) & 63))
    v4 = (v4 >> ((2) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v180
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
