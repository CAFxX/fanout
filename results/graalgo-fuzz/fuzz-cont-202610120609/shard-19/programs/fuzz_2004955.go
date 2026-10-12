package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 string; F1 int16 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 string; F1 int32; F2 bool }
type GenS8 struct { F0 int8; F1 int; F2 uint32 }
type GenS9 struct { F0 uint16; F1 uint32; F2 uint }
type GenS10 struct { F0 uint; F1 string }
type GenS11 struct { F0 int; F1 uint16 }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 int16 }
type GenS14 struct { F0 bool; F1 int32 }
type GenS15 struct { F0 uint16; F1 uint16; F2 uint64 }
type GenS16 struct { F0 int64; F1 string; F2 uint64 }
type GenS17 struct { F0 uint64; F1 bool }

func main() {
    var v1 float64 = 1.5
    v2 := []int{-160, 284, -1, 127}
    v3 := []string{("" + ((("\n\t" + "") + "foo bar") + "a")), "a", "hello", "foo bar"}
    v4 := []string{"foo bar", (("hello" + "\n\t") + "a"), "foo bar", "foo bar"}
    var v5 uint8 = 47
    var v6 uint32 = 0
    for fi7 := 0; fi7 < 4; fi7++ {
        switch "foo bar" {
        case "foo bar":
            if false {
                v8 := GenS1{F0: true, F1: -11, F2: 695}
                _ = v8
            }
        case "hello":
            if (true && (false && (true || (!false)))) {
                if (42 <= 124) {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v9 bool = true
                    var v10 uint8 = 128
                    _ = v10
                    _ = v9
                } else {
                    if len(v2) > 0 {
                        v2[0] = 9223372036854775807
                    }
                    var v11 int = fi7
                    v5 += (^(v5))
                    _ = v11
                }
            }
        }
        v12 := GenS1{F0: ((fi7 * fi7) >= (fi7 & fi7)), F1: fi7, F2: 0}
        v1 -= -2.25
        _ = fi7
        _ = v12
    }
    var v13 int64 = -9223372036854775808
    fmt.Println(v1)
    v4 = append(v4, "a")
    for fi14 := 0; fi14 < 4; fi14++ {
        v6 = v6
        for fi15 := 0; fi15 < 20; fi15++ {
            fmt.Println(fi15, fi15, fi15)
            var v16 uint = 4294967295
            fmt.Println(v5, v13)
            _ = len(v4)
            v4 = append(v4, "")
            _ = fi15
            _ = v16
        }
        var v17 float64 = 1.5
        _ = fi14
        _ = v17
    }
    v6 -= 848
    v5 = (v5 | (v5 - 24))
    fmt.Println(v6, v1, v6)
    if ((false || ((976 > -773) && ((0.1 / 1.0) <= 1e10))) && false) {
        fmt.Println(v1)
    } else {
        v18 := GenS2{F0: 0}
        v19 := GenS1{F0: (42 != int8(v13)), F1: -593, F2: 4294967295}
        v1 -= (v1 * v1)
        _ = v18
        _ = v19
    }
    if (-671 > 1) {
        for fi20 := 0; fi20 < 13; fi20++ {
            v21 := GenS1{F0: false, F1: (^(fi20)), F2: 0}
            var v22 uint64 = 1
            if (false && false) {
                fmt.Println(v22)
            } else {
                if false {
                    var v23 int8 = -117
                    if ((v23 >> ((4294967295) & 63)) >= v23) {
                        switch (fi20 + fi20) {
                        case 2147483647:
                            if (fi20 < fi20) {
                                fmt.Println(v6)
                                if (!false) {
                                    fmt.Println(v5)
                                    v1 = v1
                                }
                            } else {
                                fmt.Println(v23, v13, v6)
                                switch fi20 {
                                case -504:
                                    if false {
                                        v4 = append(v4, "hello")
                                    } else {
                                        _ = len(v4)
                                        var v24 string = "hello"
                                        _ = v24
                                    }
                                case -706:
                                    for fi25 := 0; fi25 < 1; fi25++ {
                                        switch fi25 {
                                        case 2147483647:
                                            v1 -= v1
                                        default:
                                            switch "hello" {
                                            case "hello":
                                                if (1 == -2147483648) {
                                                    v23 -= (v23 & v23)
                                                    _ = len(v4)
                                                    switch (fi25 - (fi20 | fi25)) {
                                                    case 0:
                                                        var v26 int16 = 1000
                                                        _ = v26
                                                    case 2147483647:
                                                        v6 -= (v6 % ((v6) | 1))
                                                    case 127:
                                                        if (1 > 0) {
                                                            if (42 < 423) {
                                                                var v27 uint64 = (v22 ^ (v22 + (^(v22))))
                                                                switch (fi20 | fi25) {
                                                                case 1:
                                                                    v23 = v23
                                                                }
                                                                _ = v27
                                                            }
                                                            switch "" {
                                                            case "a":
                                                                _ = len(v2)
                                                            case "\n\t":
                                                                var v28 string = (("" + ("hello" + "hello")) + "foo bar")
                                                                _ = v28
                                                            case "foo bar":
                                                                var v29 uint64 = v22
                                                                _ = v29
                                                            }
                                                        }
                                                    }
                                                    v30 := GenS3{F0: (v5 % (((v5 - 1)) | 1))}
                                                    switch (fi25 - 619) {
                                                    case 42:
                                                        if (true || false) {
                                                            var v31 int16 = 963
                                                            if true {
                                                                v31 -= v31
                                                                var v32 int = (fi25 + (fi25 << ((255) & 63)))
                                                                if true {
                                                                    v33 := GenS4{F0: ("foo bar" + (("\n\t" + ("" + "\n\t")) + "foo bar")), F1: v31}
                                                                    v6 = (v6 << ((2) & 63))
                                                                    _ = v33
                                                                } else {
                                                                    if (v6 != 328) {
                                                                        if (!false) {
                                                                            if (0.0 <= v1) {
                                                                                v34 := GenS1{F0: true, F1: 9223372036854775807, F2: (v22 - 42)}
                                                                                if false {
                                                                                    v35 := GenS5{F0: 4294967295}
                                                                                    var v36 int64 = (v13 + v13)
                                                                                    _ = v35
                                                                                    _ = v36
                                                                                }
                                                                                _ = v34
                                                                            } else {
                                                                                var v37 int32 = 2147483647
                                                                                if (3.14159 < 3.14159) {
                                                                                    v2 = append(v2, (fi25 * 127))
                                                                                    var v38 float32 = 1e10
                                                                                    fmt.Println(v32, v32, v23)
                                                                                    v39 := GenS6{F0: fi20}
                                                                                    _ = v38
                                                                                    _ = v39
                                                                                }
                                                                                _ = v37
                                                                            }
                                                                            switch "foo bar" {
                                                                            case "\n\t":
                                                                                fmt.Println(v23, v5, v1)
                                                                            case "foo bar":
                                                                                if true {
                                                                                    v40 := GenS7{F0: ("foo bar" + ("\n\t" + (("\n\t" + (("hello" + "") + ("a" + ""))) + "a"))), F1: 0, F2: false}
                                                                                    v41 := GenS3{F0: v5}
                                                                                    v2 = append(v2, v32)
                                                                                    _ = v40
                                                                                    _ = v41
                                                                                } else {
                                                                                    v31 -= (^(v31))
                                                                                }
                                                                            case "":
                                                                                var v42 bool = (((v5 >> ((2) & 63)) <= v5) && false)
                                                                                _ = v42
                                                                            default:
                                                                                fmt.Println(v6, fi20, v32)
                                                                            }
                                                                            v43 := GenS5{F0: (^(v22))}
                                                                            v1 = ((v1 - (v1 - v1)) + float64(-2.25))
                                                                            v22 += (v22 & 32)
                                                                            if len(v4) > 0 {
                                                                                v4[0] = ""
                                                                            }
                                                                            switch ("\n\t" + "") {
                                                                            case "\n\t":
                                                                                if (127 < -560) {
                                                                                    var v44 uint8 = 128
                                                                                    var v45 int64 = int64(v32)
                                                                                    _ = v44
                                                                                    _ = v45
                                                                                }
                                                                            case "foo bar":
                                                                                var v46 float32 = float32(-2.25)
                                                                                _ = v46
                                                                            case "a":
                                                                                fmt.Println(v31, fi20)
                                                                            default:
                                                                                _ = len(v3)
                                                                            }
                                                                            fmt.Println(fi20, v1, v13)
                                                                            _ = v43
                                                                        } else {
                                                                            if false {
                                                                                v31 = -1
                                                                                fmt.Println(v13)
                                                                                _ = len(v4)
                                                                            }
                                                                            var v47 uint = 255
                                                                            v48 := GenS8{F0: int8(v6), F1: (v32 | (fi25 / (((v32 & 9223372036854775807)) | 1))), F2: (v6 & (v6 % ((0) | 1)))}
                                                                            v1 = 1e10
                                                                            v49 := GenS9{F0: 42, F1: v6, F2: (v47 + 2)}
                                                                            v50 := GenS2{F0: 1}
                                                                            v51 := GenS7{F0: "hello", F1: 132, F2: (uint8(v6) <= (^(v5)))}
                                                                            if false {
                                                                                v6 += 1
                                                                                var v52 bool = (!true)
                                                                                fmt.Println(v32, v22)
                                                                                if (false || false) {
                                                                                    v13 = v13
                                                                                    v53 := GenS3{F0: (v5 ^ 135)}
                                                                                    var v54 int64 = (^(v13))
                                                                                    fmt.Println(fi20)
                                                                                    v55 := GenS1{F0: true, F1: fi20, F2: v22}
                                                                                    _ = v53
                                                                                    _ = v54
                                                                                    _ = v55
                                                                                } else {
                                                                                    switch ("" + "a") {
                                                                                    case "hello":
                                                                                        var v56 uint64 = v22
                                                                                        _ = v56
                                                                                    }
                                                                                    v5 -= (v5 + 255)
                                                                                    var v57 int64 = (v13 - (v13 + (v13 / (((v13 >> (((v47 << ((0) & 63))) & 63))) | 1))))
                                                                                    _ = v57
                                                                                }
                                                                                v58 := GenS2{F0: 65535}
                                                                                _ = v52
                                                                                _ = v58
                                                                            }
                                                                            _ = v47
                                                                            _ = v48
                                                                            _ = v49
                                                                            _ = v50
                                                                            _ = v51
                                                                        }
                                                                        var v59 int = 128
                                                                        if false {
                                                                            if false {
                                                                                var v60 uint32 = 239
                                                                                v59 = (fi25 | (fi25 - v59))
                                                                                _ = v60
                                                                            } else {
                                                                                fmt.Println(v32, fi25)
                                                                            }
                                                                        } else {
                                                                            var v61 uint = 255
                                                                            fmt.Println(v59)
                                                                            if len(v3) >= 2 {
                                                                                v3 = v3[1:len(v3)]
                                                                            }
                                                                            v62 := GenS2{F0: 42}
                                                                            if len(v2) > 0 {
                                                                                v2[0] = (^(fi25))
                                                                            }
                                                                            _ = v61
                                                                            _ = v62
                                                                        }
                                                                        _ = v59
                                                                    } else {
                                                                        fmt.Println(v6, fi25, fi25)
                                                                        if true {
                                                                            v63 := GenS4{F0: "", F1: v31}
                                                                            _ = v63
                                                                        }
                                                                        if ((1e10 * -2.25) != ((1.5 / 1.0) + 1e10)) {
                                                                            v22 = uint64(v6)
                                                                            if (false && false) {
                                                                                var v64 float32 = (3.14159 / (float32(v1) / (float32(1.5) + 3.14159)))
                                                                                _ = v64
                                                                            }
                                                                            v13 = (v13 ^ 9223372036854775807)
                                                                        } else {
                                                                            if (true && (true && (v1 == (((3.14159 * v1) - v1) / float64(-2.25))))) {
                                                                                if (!false) {
                                                                                    fmt.Println(v23, v22)
                                                                                    v2 = append(v2, (fi25 / ((int(v13)) | 1)))
                                                                                    if true {
                                                                                        if false {
                                                                                            if (true && (((-(v13)) == (v13 + 996)) && true)) {
                                                                                                v22 = 0
                                                                                                if len(v4) >= 2 {
                                                                                                    v4 = v4[1:len(v4)]
                                                                                                }
                                                                                                fmt.Println(v31, v23, v6)
                                                                                            }
                                                                                            v65 := GenS5{F0: v22}
                                                                                            _ = v65
                                                                                        } else {
                                                                                            if len(v2) > 0 {
                                                                                                v2[0] = (fi20 / ((v32) | 1))
                                                                                            }
                                                                                            switch ("" + "") {
                                                                                            case "hello":
                                                                                                v66 := GenS10{F0: uint(v6), F1: (("foo bar" + (("hello" + "\n\t") + "\n\t")) + ("" + "\n\t"))}
                                                                                                _ = v66
                                                                                            case "a":
                                                                                                v23 = 100
                                                                                            default:
                                                                                                v67 := GenS11{F0: (^(v32)), F1: 32767}
                                                                                                _ = v67
                                                                                            }
                                                                                            if ((((((3.14159 * float32(v1)) - 1e10) + 3.14159) / 1.0) * 0.0) <= (3.14159 * (0.0 / 1.5))) {
                                                                                                switch "\n\t" {
                                                                                                case "hello":
                                                                                                    var v68 int16 = (v31 ^ v31)
                                                                                                    _ = v68
                                                                                                }
                                                                                                var v69 bool = false
                                                                                                _ = v69
                                                                                            } else {
                                                                                                v13 += (^(v13))
                                                                                                v23 = int8(v6)
                                                                                                fmt.Println(v13, v1, v13)
                                                                                                fmt.Println(v1, fi25)
                                                                                                if len(v2) > 0 {
                                                                                                    v2[0] = v32
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        var v70 bool = false
                                                                                        _ = v70
                                                                                    }
                                                                                } else {
                                                                                    v3 = append(v3, ("\n\t" + "hello"))
                                                                                }
                                                                            } else {
                                                                                v4 = append(v4, "")
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                                _ = v32
                                                            } else {
                                                                var v71 int32 = -940
                                                                var v72 uint = 65535
                                                                var v73 bool = false
                                                                v74 := GenS7{F0: (("" + "hello") + ""), F1: (v71 + 283), F2: false}
                                                                v71 += (v71 ^ v71)
                                                                _ = v71
                                                                _ = v72
                                                                _ = v73
                                                                _ = v74
                                                            }
                                                            _ = v31
                                                        }
                                                        v22 = v22
                                                        if true {
                                                            var v75 uint32 = 1
                                                            var v76 int = (fi25 | (fi25 | 32768))
                                                            if true {
                                                                switch (fi20 / ((v76) | 1)) {
                                                                case 0:
                                                                    var v77 uint8 = (^(v5))
                                                                    _ = v77
                                                                default:
                                                                    var v78 int16 = 48
                                                                    _ = v78
                                                                }
                                                                var v79 int32 = 2147483647
                                                                var v80 float32 = 1e10
                                                                _ = v79
                                                                _ = v80
                                                            } else {
                                                                switch "a" {
                                                                case "foo bar":
                                                                    v81 := GenS12{F0: 42}
                                                                    _ = v81
                                                                default:
                                                                    switch "\n\t" {
                                                                    case "foo bar":
                                                                        var v82 int8 = (v23 >> ((2) & 63))
                                                                        _ = v82
                                                                    case "\n\t":
                                                                        _ = len(v2)
                                                                    case "":
                                                                        var v83 float32 = (1e10 - (0.1 + (3.14159 * 1.5)))
                                                                        _ = v83
                                                                    default:
                                                                        var v84 int8 = (v23 + v23)
                                                                        v84 = (v23 * 100)
                                                                        var v85 uint64 = 0
                                                                        _ = v84
                                                                        _ = v85
                                                                    }
                                                                }
                                                                var v86 uint8 = v5
                                                                _ = v86
                                                            }
                                                            _ = v75
                                                            _ = v76
                                                        } else {
                                                            if true {
                                                                v22 += v22
                                                                fmt.Println(fi25, fi25, fi20)
                                                                var v87 string = ("" + "a")
                                                                _ = v87
                                                            } else {
                                                                switch fi20 {
                                                                case 32767:
                                                                    v22 = (v22 - (v22 + v22))
                                                                case 128:
                                                                    var v88 int8 = (v23 >> ((0) & 63))
                                                                    _ = v88
                                                                }
                                                                var v89 int16 = 32767
                                                                v90 := GenS13{F0: (v89 ^ int16(v22))}
                                                                _ = v89
                                                                _ = v90
                                                            }
                                                            var v91 uint32 = v6
                                                            v92 := GenS14{F0: false, F1: 470}
                                                            _ = v91
                                                            _ = v92
                                                        }
                                                    case -501:
                                                        if false {
                                                            var v93 int = (fi20 << ((4294967295) & 63))
                                                            v6 = (v6 >> ((0) & 63))
                                                            if true {
                                                                var v94 int64 = (v13 / ((v13) | 1))
                                                                v94 = v13
                                                                _ = v94
                                                            } else {
                                                                v2 = append(v2, int(v22))
                                                                var v95 string = ("\n\t" + "a")
                                                                var v96 int64 = v13
                                                                _ = v95
                                                                _ = v96
                                                            }
                                                            _ = v93
                                                        }
                                                    default:
                                                        v13 += (v13 / ((9223372036854775807) | 1))
                                                    }
                                                    _ = v30
                                                } else {
                                                    v13 = v13
                                                    var v97 float64 = ((v1 / (1e10 / (0.0 / (v1 - v1)))) - float64(0.1))
                                                    fmt.Println(v13)
                                                    _ = v97
                                                }
                                            case "a":
                                                fmt.Println(fi20, fi20, fi25)
                                            default:
                                                v22 = (v22 + 1)
                                            }
                                        }
                                        _ = fi25
                                    }
                                }
                                var v98 uint64 = (v22 << ((255) & 63))
                                _ = v98
                            }
                        default:
                            for fi99 := 0; fi99 < 9; fi99++ {
                                v100 := GenS12{F0: uint32(v22)}
                                if (((!((v13 << ((18446744073709551615) & 63)) <= v13)) && false) && (!true)) {
                                    fmt.Println(v23, fi99, v1)
                                    v23 += 100
                                    fmt.Println(v22, fi20, v22)
                                } else {
                                    if false {
                                        v3 = append(v3, ("foo bar" + (("hello" + ("hello" + ("hello" + "foo bar"))) + "")))
                                    }
                                    var v101 int16 = -975
                                    _ = v101
                                }
                                _ = fi99
                                _ = v100
                            }
                        }
                        v102 := GenS10{F0: 764, F1: ("hello" + "\n\t")}
                        v103 := GenS15{F0: 32767, F1: 255, F2: (^(v22))}
                        _ = v102
                        _ = v103
                    } else {
                        v23 = (v23 * -128)
                        if ((127 != v23) || false) {
                            if false {
                                _ = len(v4)
                                for fi104 := 0; fi104 < 8; fi104++ {
                                    v5 += v5
                                    if len(v2) > 0 {
                                        v2[0] = fi20
                                    }
                                    _ = fi104
                                }
                            } else {
                                switch fi20 {
                                case 131:
                                    v6 -= (v6 >> ((42) & 63))
                                default:
                                    switch "\n\t" {
                                    case "":
                                        fmt.Println(v23, v13, v1)
                                    case "hello":
                                        for fi105 := 0; fi105 < 9; fi105++ {
                                            v106 := GenS16{F0: int64(fi105), F1: (("foo bar" + "hello") + (("foo bar" + "a") + "foo bar")), F2: (v22 | uint64(v6))}
                                            _ = fi105
                                            _ = v106
                                        }
                                    case "a":
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    default:
                                        v107 := GenS1{F0: (((("a" + "") + "\n\t") + "a") <= ("foo bar" + ("\n\t" + "hello"))), F1: fi20, F2: v22}
                                        _ = v107
                                    }
                                }
                                if (!false) {
                                    fmt.Println(v22, fi20, v1)
                                } else {
                                    var v108 int8 = (-(v23))
                                    if true {
                                        var v109 uint = 65535
                                        _ = v109
                                    } else {
                                        if false {
                                            for fi110 := 0; fi110 < 13; fi110++ {
                                                if false {
                                                    if true {
                                                        var v111 string = ("hello" + ((("\n\t" + "hello") + (("foo bar" + "a") + ("foo bar" + ""))) + (("\n\t" + ("" + "hello")) + "")))
                                                        _ = v111
                                                    }
                                                    var v112 int8 = (-(v23))
                                                    _ = v112
                                                }
                                                switch 32767 {
                                                case -394:
                                                    if (((fi20 & fi20) >= (fi110 | fi20)) && (-9223372036854775808 == v13)) {
                                                        fmt.Println(v22, v13)
                                                    } else {
                                                        var v113 float32 = 0.1
                                                        _ = v113
                                                    }
                                                }
                                                v108 = v108
                                                _ = fi110
                                            }
                                            switch "a" {
                                            case "\n\t":
                                                if len(v2) > 0 {
                                                    v2[0] = -758
                                                }
                                            case "foo bar":
                                                v4 = append(v4, "hello")
                                            case "hello":
                                                v108 = 127
                                            default:
                                                v114 := GenS8{F0: v108, F1: (fi20 >> ((65535) & 63)), F2: v6}
                                                _ = v114
                                            }
                                        } else {
                                            v115 := GenS6{F0: fi20}
                                            _ = v115
                                        }
                                    }
                                    _ = v108
                                }
                            }
                            v23 = (v23 - (v23 << ((813) & 63)))
                            v116 := GenS17{F0: 4294967295, F1: false}
                            _ = v116
                        }
                    }
                    _ = v23
                }
                if len(v2) > 0 {
                    v2[0] = (fi20 - fi20)
                }
                for fi117 := 0; fi117 < 17; fi117++ {
                    v118 := GenS15{F0: 539, F1: 65535, F2: (v22 & 335)}
                    fmt.Println(fi117)
                    _ = fi117
                    _ = v118
                }
            }
            _ = fi20
            _ = v21
            _ = v22
        }
        v119 := GenS12{F0: v6}
        _ = v119
    }
    switch "a" {
    case "foo bar":
        var v120 uint8 = 42
        _ = v120
    case "":
        for fi121 := 0; fi121 < 8; fi121++ {
            v1 += 1e10
            fmt.Println(v13)
            _ = fi121
        }
    case "\n\t":
        var v122 uint16 = 0
        _ = v122
    }
    var v123 int8 = 127
    v124 := GenS13{F0: 1000}
    var v125 int32 = 2147483647
    v126 := GenS7{F0: ("" + ("hello" + "foo bar")), F1: -205, F2: ((3.14159 / 1e10) < (float32(1.5) * (((float32((1e10 / (v1 + v1))) * 0.1) / (1.5 / 1.0)) + 3.14159)))}
    v125 = v125
    fmt.Println("v1:", v1)
    fmt.Println("v123:", v123)
    fmt.Println("v125:", v125)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v123
    _ = v124
    _ = v125
    _ = v126
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
