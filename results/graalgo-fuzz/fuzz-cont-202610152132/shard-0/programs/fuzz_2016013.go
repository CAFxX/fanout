package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int; F1 int32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint8; F1 uint }
type GenS5 struct { F0 int; F1 uint64 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 uint64; F1 int32; F2 int8 }
type GenS10 struct { F0 uint8; F1 int16; F2 uint16 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 uint16; F1 int16; F2 uint8 }
type GenS13 struct { F0 int32; F1 int32 }
type GenS14 struct { F0 uint8 }
type GenS15 struct { F0 uint32; F1 string }
type GenS16 struct { F0 int8 }
type GenS17 struct { F0 int8 }
type GenS18 struct { F0 int64; F1 int16 }
type GenS19 struct { F0 bool }
type GenS20 struct { F0 int; F1 int32; F2 string }

func main() {
    var v1 uint16 = 255
    v2 := []int{-2147483648}
    v3 := make(map[string]string)
    for fi4 := 0; fi4 < 8; fi4++ {
        for fi5 := 0; fi5 < 18; fi5++ {
            if false {
                fmt.Println(fi4)
                v2 = append(v2, 0)
                var v6 uint32 = 65535
                _ = v6
            } else {
                fmt.Println(fi5)
                if true {
                    if false {
                        if (!(((1.5 / -2.25) - 1e10) < 1e10)) {
                            v2 = append(v2, (fi5 | fi4))
                        }
                    } else {
                        _ = len(v3)
                        if (true || (float64(0.0) > -2.25)) {
                            switch ("hello" + "a") {
                            case "foo bar":
                                switch fi4 {
                                case 9223372036854775807:
                                    var v7 uint = 876
                                    _ = v7
                                case 591:
                                    v1 = (v1 << ((143) & 63))
                                case 840:
                                    var v8 int8 = 0
                                    fmt.Println(fi5)
                                    var v9 uint16 = v1
                                    v9 -= 65535
                                    fmt.Println(fi4, fi4)
                                    if (!((1.5 + (-2.25 + (0.0 - 1.5))) != (0.0 + (-2.25 - (float32(-2.25) / 1.5))))) {
                                        fmt.Println(v1, v1, fi4)
                                    } else {
                                        if false {
                                            v3["\n\t"] = "foo bar"
                                            if false {
                                                var v10 bool = (false && (((v9 % (((v9 / ((65535) | 1))) | 1)) != (^(v9))) && (true && true)))
                                                var v11 int32 = 42
                                                if len(v2) > 0 {
                                                    v2[0] = (fi4 >> ((18446744073709551615) & 63))
                                                }
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                var v12 uint8 = 0
                                                _ = v10
                                                _ = v11
                                                _ = v12
                                            } else {
                                                v8 = -1
                                            }
                                            var v13 uint32 = 2147483648
                                            _ = v13
                                        }
                                        if false {
                                            var v14 int8 = (v8 & v8)
                                            _ = v14
                                        }
                                    }
                                    _ = v8
                                    _ = v9
                                default:
                                    var v15 float32 = 1.5
                                    _ = v15
                                }
                            case "hello":
                                var v16 uint64 = 4294967295
                                v17 := GenS1{F0: (!false)}
                                var v18 float64 = (0.0 - ((0.1 - ((-2.25 - 1.5) / 1.5)) + 3.14159))
                                _ = v16
                                _ = v17
                                _ = v18
                            case "\n\t":
                                v1 = 42
                            default:
                                v1 -= 32767
                            }
                        } else {
                            var v19 float32 = 1.5
                            _ = v19
                        }
                    }
                    var v20 int16 = 42
                    switch "hello" {
                    case "hello":
                        fmt.Println(fi4, fi5)
                    }
                    var v21 int16 = -1
                    var v22 int16 = 1
                    if true {
                        if false {
                            if ((v1 ^ v1) < 455) {
                                v23 := GenS1{F0: (false || (4294967295 == 136))}
                                _ = v23
                            } else {
                                _ = len(v2)
                                switch fi5 {
                                case 745:
                                    var v24 int16 = 1000
                                    _ = v24
                                case 30:
                                    var v25 float64 = 0.0
                                    _ = v25
                                case 128:
                                    v26 := GenS2{F0: (fi5 | fi5), F1: -2147483648}
                                    _ = v26
                                default:
                                    var v27 bool = false
                                    _ = v27
                                }
                                v28 := GenS2{F0: 4294967295, F1: 526}
                                v20 -= v22
                                _ = v28
                            }
                        }
                        v29 := GenS1{F0: false}
                        _ = v29
                    }
                    _ = v20
                    _ = v21
                    _ = v22
                }
            }
            if true {
                v30 := GenS1{F0: (508 == 2)}
                _ = v30
            } else {
                v31 := GenS3{F0: ("hello" + "hello")}
                v32 := GenS3{F0: "foo bar"}
                v1 -= (v1 << ((0) & 63))
                _ = v31
                _ = v32
            }
            _ = fi5
        }
        v1 = 32767
        _ = fi4
    }
    v2 = append(v2, 32768)
    if (1 < v1) {
        if (65535 >= uint(v1)) {
            v33 := GenS3{F0: "\n\t"}
            _ = v33
        } else {
            fmt.Println(v1)
            v2 = append(v2, 32768)
            var v34 float32 = ((3.14159 - 0.0) * float32(1.5))
            _ = v34
        }
        var v35 float64 = 1e10
        if (1.5 == -2.25) {
            if false {
                if false {
                    var v36 int64 = -835
                    _ = v36
                } else {
                    var v37 bool = false
                    _ = v37
                }
            } else {
                if true {
                    for fi38 := 0; fi38 < 17; fi38++ {
                        var v39 int32 = -2147483648
                        _ = fi38
                        _ = v39
                    }
                    if false {
                        v40 := GenS2{F0: 9223372036854775807, F1: 488}
                        v2 = append(v2, 42)
                        _ = v40
                    }
                    var v41 string = ""
                    _ = v41
                } else {
                    v35 = v35
                    v35 -= ((v35 / (0.0 / v35)) + 3.14159)
                    v2 = append(v2, 32768)
                }
                if false {
                    fmt.Println(v35, v35)
                    v3["a"] = "hello"
                } else {
                    if len(v2) > 0 {
                        v2[0] = 9223372036854775807
                    }
                    fmt.Println(v1)
                }
                if true {
                    if false {
                        v42 := v3["\n\t"]
                        _ = v42
                    } else {
                        var v43 uint64 = 42
                        var v44 int32 = 2147483647
                        var v45 int32 = v44
                        _ = v43
                        _ = v44
                        _ = v45
                    }
                    v1 = (v1 << ((203) & 63))
                    v46 := GenS3{F0: "foo bar"}
                    _ = v46
                }
            }
        } else {
            var v47 uint16 = v1
            _ = v47
        }
        _ = v35
    } else {
        switch (("hello" + "foo bar") + ("hello" + "")) {
        case "hello":
            v2 = append(v2, -2147483648)
        case "":
            v1 = (v1 % (((v1 >> ((uint(v1)) & 63))) | 1))
        case "\n\t":
            fmt.Println(v1)
        default:
            for fi48 := 0; fi48 < 19; fi48++ {
                var v49 int64 = 714
                var v50 uint16 = (v1 >> ((18446744073709551615) & 63))
                _ = fi48
                _ = v49
                _ = v50
            }
        }
        for fi51 := 0; fi51 < 10; fi51++ {
            if (553 != 42) {
                v1 = 2
                if false {
                    v1 = (v1 & v1)
                    v52 := GenS4{F0: 1, F1: 255}
                    var v53 int16 = 32767
                    _ = v52
                    _ = v53
                } else {
                    v54 := GenS4{F0: 108, F1: 4294967295}
                    switch "a" {
                    case "":
                        v2 = append(v2, (fi51 + (^(fi51))))
                    }
                    _ = v54
                }
            } else {
                var v55 int16 = -1
                v55 = v55
                _ = v55
            }
            if false {
                var v56 int8 = -90
                v3["hello"] = "hello"
                _ = v56
            } else {
                v1 = (^(v1))
                var v57 int16 = 0
                for fi58 := 0; fi58 < 16; fi58++ {
                    if (3.14159 <= 1.5) {
                        var v59 bool = false
                        fmt.Println(fi58)
                        if v59 {
                            switch (fi58 << ((1) & 63)) {
                            case 718:
                                fmt.Println(v57, v59)
                            }
                            if true {
                                if v59 {
                                    v1 -= (v1 << ((65535) & 63))
                                    if v59 {
                                        v59 = ((-2.25 * 0.1) < ((0.1 * 0.0) + 3.14159))
                                        _ = len(v3)
                                        v60 := GenS5{F0: fi51, F1: 18446744073709551615}
                                        _ = v60
                                    } else {
                                        v57 -= -1
                                        v61 := GenS6{F0: 172}
                                        _ = v61
                                    }
                                    var v62 int16 = (v57 - v57)
                                    _ = v62
                                } else {
                                    if v59 {
                                        var v63 uint64 = 1
                                        v57 -= (v57 + v57)
                                        _ = v63
                                    } else {
                                        if ((^(fi51)) > fi51) {
                                            var v64 bool = false
                                            var v65 uint8 = 2
                                            var v66 uint64 = 601
                                            if (((fi51 - (fi51 & -2147483648)) > (fi58 * (fi51 % (((fi58 / ((fi51) | 1))) | 1)))) || (((v1 != (^(v1))) || v64) && false)) {
                                                var v67 float32 = (0.1 / 1.0)
                                                if false {
                                                    switch (("\n\t" + ("\n\t" + "")) + ("foo bar" + "a")) {
                                                    case "a":
                                                        _ = len(v2)
                                                        var v68 int = (-(fi58))
                                                        fmt.Println(v66, v66)
                                                        if v64 {
                                                            fmt.Println(fi58, v66, v1)
                                                            switch ("foo bar" + ("hello" + "a")) {
                                                            case "foo bar":
                                                                v69 := GenS5{F0: 127, F1: uint64(v1)}
                                                                _ = v69
                                                            case "\n\t":
                                                                _ = len(v2)
                                                            case "hello":
                                                                if (!((!v59) || ((v66 + 922) > (^(v66))))) {
                                                                    v59 = (!true)
                                                                    v70 := GenS7{F0: "hello"}
                                                                    _ = v70
                                                                }
                                                            }
                                                        } else {
                                                            var v71 int8 = -73
                                                            if (("foo bar" + "foo bar") <= (("a" + "\n\t") + ("foo bar" + ("\n\t" + "hello")))) {
                                                                var v72 int8 = (v71 - 42)
                                                                v73 := GenS8{F0: v66}
                                                                if v59 {
                                                                    switch ("a" + ("\n\t" + ("" + ("hello" + "foo bar")))) {
                                                                    case "":
                                                                        v74 := GenS9{F0: 18446744073709551615, F1: -392, F2: v71}
                                                                        _ = v74
                                                                    case "foo bar":
                                                                        fmt.Println(v66, v65, v65)
                                                                    case "hello":
                                                                        if ((v1 + v1) > (v1 & 42)) {
                                                                            fmt.Println(v1, v72)
                                                                        } else {
                                                                            var v75 uint = 2
                                                                            v66 = (^(v66))
                                                                            switch int(v1) {
                                                                            case 127:
                                                                                if len(v2) > 0 {
                                                                                    v2[0] = (fi58 | int(v57))
                                                                                }
                                                                            }
                                                                            v59 = ("hello" > "\n\t")
                                                                            _ = v75
                                                                        }
                                                                    default:
                                                                        v2 = append(v2, (fi58 - fi51))
                                                                    }
                                                                } else {
                                                                    v76 := GenS10{F0: (^(v65)), F1: v57, F2: v1}
                                                                    _ = v76
                                                                }
                                                                if v64 {
                                                                    var v77 string = "hello"
                                                                    var v78 int = (fi58 ^ (fi58 ^ fi51))
                                                                    if v59 {
                                                                        v68 = (v68 >> ((uint(v66)) & 63))
                                                                        v67 = -2.25
                                                                    } else {
                                                                        var v79 int32 = 98
                                                                        switch fi51 {
                                                                        case 128:
                                                                            if v64 {
                                                                                v80 := v3["hello"]
                                                                                var v81 uint32 = 0
                                                                                _ = v80
                                                                                _ = v81
                                                                            }
                                                                        default:
                                                                            if (115 == 18) {
                                                                                var v82 uint8 = (v65 * (v65 % ((v65) | 1)))
                                                                                _ = v82
                                                                            } else {
                                                                                v2 = append(v2, -4)
                                                                                fmt.Println(v1, v71, fi51)
                                                                                fmt.Println(v77, v68)
                                                                            }
                                                                        }
                                                                        var v83 float32 = 0.0
                                                                        _ = v79
                                                                        _ = v83
                                                                    }
                                                                    _ = v77
                                                                    _ = v78
                                                                }
                                                                _ = v72
                                                                _ = v73
                                                            }
                                                            if ((v57 % (((^(v57))) | 1)) != v57) {
                                                                v84 := v3[("foo bar" + "a")]
                                                                v85 := GenS3{F0: v84}
                                                                var v86 int16 = (v57 << ((65535) & 63))
                                                                _ = v84
                                                                _ = v85
                                                                _ = v86
                                                            } else {
                                                                v87 := GenS3{F0: ("hello" + "hello")}
                                                                _ = len(v3)
                                                                v64 = true
                                                                switch "a" {
                                                                case "hello":
                                                                    if (!false) {
                                                                        v1 += v1
                                                                    }
                                                                case "":
                                                                    var v88 int8 = (v71 & -70)
                                                                    _ = v88
                                                                case "\n\t":
                                                                    if len(v2) > 0 {
                                                                        v2[0] = fi51
                                                                    }
                                                                default:
                                                                    switch (("foo bar" + "hello") + ("" + "hello")) {
                                                                    case "":
                                                                        v66 = 583
                                                                    case "a":
                                                                        v65 = v65
                                                                    case "hello":
                                                                        if v59 {
                                                                            switch fi58 {
                                                                            case 127:
                                                                                switch (-(fi58)) {
                                                                                case -961:
                                                                                    if (!v59) {
                                                                                        v64 = (true || (fi58 >= 423))
                                                                                    }
                                                                                default:
                                                                                    v89 := v3[(("foo bar" + "hello") + "\n\t")]
                                                                                    _ = v89
                                                                                }
                                                                            case 0:
                                                                                switch "\n\t" {
                                                                                case "a":
                                                                                    var v90 uint = uint(v65)
                                                                                    _ = v90
                                                                                case "hello":
                                                                                    v3["a"] = "\n\t"
                                                                                case "":
                                                                                    v57 = (v57 + -101)
                                                                                }
                                                                            default:
                                                                                delete(v3, "a")
                                                                            }
                                                                            v59 = (v71 < (v71 / (((v71 - (v71 - v71))) | 1)))
                                                                            v66 = 1
                                                                        }
                                                                    default:
                                                                        switch int(v57) {
                                                                        case 128:
                                                                            v68 -= 843
                                                                        case 2147483647:
                                                                            var v91 bool = (!false)
                                                                            v68 += (v68 >> ((42) & 63))
                                                                            if ("a" > "\n\t") {
                                                                                var v92 float64 = 3.14159
                                                                                v93 := GenS9{F0: (v66 - v66), F1: 239, F2: -89}
                                                                                _ = v92
                                                                                _ = v93
                                                                            } else {
                                                                                switch "" {
                                                                                case "":
                                                                                    if (uint32(v65) == 910) {
                                                                                        v67 = float32(1.5)
                                                                                    }
                                                                                    v94 := GenS6{F0: v65}
                                                                                    _ = v94
                                                                                }
                                                                                v95 := v3["foo bar"]
                                                                                _ = v95
                                                                            }
                                                                            _ = v91
                                                                        default:
                                                                            var v96 uint16 = 0
                                                                            _ = v96
                                                                        }
                                                                    }
                                                                }
                                                                if (((v67 == ((-2.25 - 1e10) / float32(-2.25))) && false) || ((^(v66)) < v66)) {
                                                                    v97 := GenS10{F0: (v65 + 119), F1: 0, F2: 2}
                                                                    var v98 uint16 = v1
                                                                    _ = v97
                                                                    _ = v98
                                                                } else {
                                                                    v67 = (1.5 / 3.14159)
                                                                    if ((1 < (^(v66))) && false) {
                                                                        if false {
                                                                            _ = len(v2)
                                                                            fmt.Println(v65, v65, v65)
                                                                        } else {
                                                                            switch (("" + "foo bar") + ("foo bar" + "foo bar")) {
                                                                            case "hello":
                                                                                if true {
                                                                                    v71 += v71
                                                                                    var v99 uint32 = 4294967295
                                                                                    v1 = (^(v1))
                                                                                    _ = v99
                                                                                } else {
                                                                                    fmt.Println(fi58, fi58, v59)
                                                                                    v59 = v64
                                                                                    var v100 uint16 = uint16(v65)
                                                                                    fmt.Println(v65)
                                                                                    _ = v100
                                                                                }
                                                                            case "":
                                                                                if false {
                                                                                    v66 += v66
                                                                                } else {
                                                                                    fmt.Println(v65, v1)
                                                                                    switch int(v1) {
                                                                                    case -2147483648:
                                                                                        fmt.Println(fi58)
                                                                                    }
                                                                                    switch ("foo bar" + ("a" + "hello")) {
                                                                                    case "\n\t":
                                                                                        v66 += 273
                                                                                        v101 := GenS10{F0: 234, F1: (v57 | (v57 << ((255) & 63))), F2: v1}
                                                                                        _ = v101
                                                                                    case "":
                                                                                        if v64 {
                                                                                            var v102 uint16 = (v1 << ((300) & 63))
                                                                                            v66 = 4294967295
                                                                                            var v103 uint = 4294967295
                                                                                            _ = v102
                                                                                            _ = v103
                                                                                        } else {
                                                                                            fmt.Println(v65, v57)
                                                                                            if true {
                                                                                                var v104 int32 = -2147483648
                                                                                                var v105 uint = 42
                                                                                                _ = v104
                                                                                                _ = v105
                                                                                            } else {
                                                                                                v65 += v65
                                                                                                var v106 int64 = -9223372036854775808
                                                                                                if ((-2.25 <= 1.5) && v59) {
                                                                                                    delete(v3, "a")
                                                                                                    fmt.Println(v64, v59)
                                                                                                }
                                                                                                _ = v106
                                                                                            }
                                                                                        }
                                                                                    case "hello":
                                                                                        var v107 uint32 = 662
                                                                                        _ = v107
                                                                                    }
                                                                                }
                                                                            case "foo bar":
                                                                                v64 = v64
                                                                            default:
                                                                                if len(v2) >= 2 {
                                                                                    v2 = v2[1:len(v2)]
                                                                                }
                                                                            }
                                                                            var v108 string = "hello"
                                                                            _ = v108
                                                                        }
                                                                        var v109 int32 = 42
                                                                        delete(v3, ("hello" + "\n\t"))
                                                                        _ = v109
                                                                    } else {
                                                                        if v59 {
                                                                            if ((4294967295 == 0) || v64) {
                                                                                if v59 {
                                                                                    v64 = v59
                                                                                    v110 := GenS1{F0: ("\n\t" >= "\n\t")}
                                                                                    v2 = append(v2, -1)
                                                                                    if (-2147483648 == -1) {
                                                                                        var v111 uint32 = 775
                                                                                        fmt.Println(v65, v71, v57)
                                                                                        var v112 int32 = 42
                                                                                        delete(v3, "\n\t")
                                                                                        _ = v111
                                                                                        _ = v112
                                                                                    }
                                                                                    v68 = (fi51 + (^(fi51)))
                                                                                    _ = v110
                                                                                }
                                                                                var v113 int64 = -965
                                                                                var v114 uint64 = v66
                                                                                _ = v113
                                                                                _ = v114
                                                                            }
                                                                            _ = len(v3)
                                                                            var v115 uint64 = v66
                                                                            fmt.Println(v64, v59, fi58)
                                                                            _ = len(v2)
                                                                            _ = v115
                                                                        } else {
                                                                            var v116 int = v68
                                                                            var v117 float32 = 1.5
                                                                            v64 = ((^(v71)) > 127)
                                                                            _ = v116
                                                                            _ = v117
                                                                        }
                                                                    }
                                                                    v68 -= (v68 / (((-(fi58))) | 1))
                                                                    if (-1 > 42) {
                                                                        if (!v64) {
                                                                            _ = len(v2)
                                                                            switch (^(fi58)) {
                                                                            case 491:
                                                                                v118 := GenS7{F0: (("a" + ((("" + "") + (("foo bar" + "") + "foo bar")) + "hello")) + "a")}
                                                                                _ = v118
                                                                            default:
                                                                                fmt.Println(v1, v67)
                                                                            }
                                                                            if (v64 || false) {
                                                                                var v119 int32 = -43
                                                                                var v120 uint16 = 65535
                                                                                v3["foo bar"] = (("hello" + "a") + (("a" + "\n\t") + ""))
                                                                                var v121 float64 = 0.0
                                                                                fmt.Println(fi58, v65, v68)
                                                                                _ = v119
                                                                                _ = v120
                                                                                _ = v121
                                                                            }
                                                                            switch ("hello" + "a") {
                                                                            case "\n\t":
                                                                                v59 = true
                                                                            case "":
                                                                                v68 += (^(fi58))
                                                                            default:
                                                                                v122 := v3[("foo bar" + (("" + "") + ("foo bar" + "a")))]
                                                                                _ = v122
                                                                            }
                                                                            switch v68 {
                                                                            case 2147483647:
                                                                                var v123 float32 = v67
                                                                                _ = v123
                                                                            default:
                                                                                var v124 uint8 = (v65 << ((2) & 63))
                                                                                _ = v124
                                                                            }
                                                                        } else {
                                                                            if (41 == 255) {
                                                                                if v64 {
                                                                                    if v59 {
                                                                                        v125 := GenS2{F0: 9223372036854775807, F1: 806}
                                                                                        fmt.Println(fi58, v64, v71)
                                                                                        v66 -= v66
                                                                                        _ = v125
                                                                                    }
                                                                                    var v126 uint16 = (^(v1))
                                                                                    switch "a" {
                                                                                    case "foo bar":
                                                                                        var v127 uint8 = (^(v65))
                                                                                        _ = v127
                                                                                    default:
                                                                                        fmt.Println(v71, v71)
                                                                                    }
                                                                                    _ = v126
                                                                                }
                                                                                v128 := GenS11{F0: 42}
                                                                                var v129 int32 = 0
                                                                                _ = v128
                                                                                _ = v129
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                                _ = v87
                                                            }
                                                            _ = v71
                                                        }
                                                        _ = v68
                                                    case "foo bar":
                                                        v67 = (v67 + -2.25)
                                                    case "":
                                                        v67 = v67
                                                        fmt.Println(v64, v1, v59)
                                                    default:
                                                        v1 = 760
                                                    }
                                                    fmt.Println(v1, fi51)
                                                    v130 := GenS7{F0: ("\n\t" + (("hello" + ("hello" + "foo bar")) + ""))}
                                                    _ = v130
                                                }
                                                _ = v67
                                            }
                                            _ = v64
                                            _ = v65
                                            _ = v66
                                        } else {
                                            v59 = v59
                                        }
                                        v131 := GenS12{F0: v1, F1: (v57 - v57), F2: 30}
                                        switch (("\n\t" + ("a" + "a")) + "") {
                                        case "":
                                            v1 = (v1 % ((v1) | 1))
                                        case "hello":
                                            var v132 int16 = 0
                                            _ = v132
                                        }
                                        v133 := GenS2{F0: (^(fi51)), F1: 93}
                                        var v134 uint8 = 127
                                        if (18446744073709551615 != 566) {
                                            v135 := v3["\n\t"]
                                            _ = v135
                                        }
                                        _ = v131
                                        _ = v133
                                        _ = v134
                                    }
                                    v136 := GenS13{F0: -446, F1: 70}
                                    fmt.Println(v57, fi58, v1)
                                    _ = v136
                                }
                            } else {
                                fmt.Println(v59, fi58)
                            }
                        } else {
                            if (!(42 >= 42)) {
                                var v137 float64 = 1.5
                                fmt.Println(fi58)
                                switch "foo bar" {
                                case "\n\t":
                                    fmt.Println(v57, fi51)
                                case "hello":
                                    var v138 int8 = 127
                                    _ = v138
                                case "a":
                                    if false {
                                        v137 += v137
                                    } else {
                                        fmt.Println(v57, v1)
                                    }
                                default:
                                    if v59 {
                                        v137 = -2.25
                                        v139 := GenS14{F0: 2}
                                        var v140 int32 = -1
                                        v141 := GenS10{F0: 128, F1: 42, F2: v1}
                                        _ = v139
                                        _ = v140
                                        _ = v141
                                    } else {
                                        if (2 != 1) {
                                            v137 = 0.1
                                            fmt.Println(v57)
                                        }
                                        v137 = v137
                                    }
                                }
                                _ = v137
                            }
                            if v59 {
                                v142 := GenS15{F0: 1, F1: "foo bar"}
                                var v143 uint64 = 1
                                _ = v142
                                _ = v143
                            } else {
                                v144 := GenS1{F0: (42 > 0)}
                                v1 += v1
                                _ = v144
                            }
                            v1 += v1
                        }
                        _ = v59
                    } else {
                        switch ("hello" + ("foo bar" + "")) {
                        case "hello":
                            if ("a" <= (("foo bar" + "foo bar") + "")) {
                                if (3.14159 != (1e10 + 1.5)) {
                                    v145 := GenS8{F0: 4294967295}
                                    if (33 <= -117) {
                                        var v146 int16 = v57
                                        switch fi51 {
                                        case 127:
                                            v147 := GenS5{F0: (fi58 << ((1) & 63)), F1: 752}
                                            _ = v147
                                        case 0:
                                            if (((!(4294967295 != 0)) || (true || (9223372036854775807 > -9223372036854775808))) || false) {
                                                var v148 int64 = -9223372036854775808
                                                _ = v148
                                            }
                                        case -205:
                                            switch 128 {
                                            case 32767:
                                                var v149 int8 = 42
                                                _ = v149
                                            case 2147483647:
                                                v150 := GenS13{F0: -2147483648, F1: 42}
                                                _ = v150
                                            case 127:
                                                delete(v3, (("\n\t" + "hello") + "hello"))
                                            default:
                                                if (!false) {
                                                    v146 = v57
                                                    v146 = v57
                                                } else {
                                                    v57 -= v146
                                                    var v151 int16 = (^(v57))
                                                    _ = v151
                                                }
                                            }
                                        default:
                                            _ = len(v3)
                                        }
                                        v1 = (v1 & v1)
                                        _ = v146
                                    }
                                    _ = v145
                                } else {
                                    switch "foo bar" {
                                    case "\n\t":
                                        var v152 int64 = -139
                                        _ = v152
                                    case "a":
                                        v153 := GenS11{F0: 9223372036854775807}
                                        _ = v153
                                    case "":
                                        if (false && (42 < 42)) {
                                            v57 = v57
                                            if (!true) {
                                                var v154 int64 = 1
                                                _ = len(v3)
                                                if ("hello" < "hello") {
                                                    v155 := GenS11{F0: (v154 + (v154 + (v154 & v154)))}
                                                    _ = v155
                                                }
                                                _ = v154
                                            } else {
                                                if false {
                                                    if ((1.5 / -2.25) != 1.5) {
                                                        v156 := v3["hello"]
                                                        _ = v156
                                                    }
                                                } else {
                                                    if (!(18446744073709551615 != 12)) {
                                                        var v157 float64 = 0.0
                                                        if true {
                                                            var v158 int = fi58
                                                            _ = v158
                                                        } else {
                                                            if (-1 != -128) {
                                                                _ = len(v2)
                                                            }
                                                            if false {
                                                                var v159 int32 = -609
                                                                fmt.Println(v57)
                                                                v160 := GenS16{F0: int8(v1)}
                                                                _ = v159
                                                                _ = v160
                                                            }
                                                            fmt.Println(v157, v1)
                                                            var v161 int8 = -1
                                                            var v162 uint = 4294967295
                                                            _ = v161
                                                            _ = v162
                                                        }
                                                        _ = v157
                                                    }
                                                    if true {
                                                        var v163 int64 = 862
                                                        if true {
                                                            v164 := GenS11{F0: v163}
                                                            delete(v3, (("a" + "foo bar") + "a"))
                                                            _ = v164
                                                        } else {
                                                            switch int(v57) {
                                                            case 32767:
                                                                v165 := GenS17{F0: 114}
                                                                _ = v165
                                                            case 111:
                                                                v57 = v57
                                                            }
                                                            fmt.Println(v57, fi58, v1)
                                                            switch "hello" {
                                                            case "hello":
                                                                v166 := GenS10{F0: 127, F1: 962, F2: 625}
                                                                _ = v166
                                                            case "":
                                                                if (!true) {
                                                                    v1 -= v1
                                                                    switch "" {
                                                                    case "\n\t":
                                                                        if true {
                                                                            v167 := GenS16{F0: -8}
                                                                            _ = v167
                                                                        } else {
                                                                            v163 = (v163 | (^(v163)))
                                                                            v168 := GenS17{F0: 127}
                                                                            _ = v168
                                                                        }
                                                                    }
                                                                    if false {
                                                                        var v169 int = (fi51 | fi51)
                                                                        v170 := GenS15{F0: 178, F1: ((("\n\t" + (("foo bar" + "hello") + "foo bar")) + ("a" + "hello")) + (("" + "\n\t") + "hello"))}
                                                                        _ = v169
                                                                        _ = v170
                                                                    }
                                                                }
                                                            case "a":
                                                                v171 := GenS15{F0: 2147483648, F1: "a"}
                                                                _ = v171
                                                            default:
                                                                if true {
                                                                    var v172 uint8 = 137
                                                                    if false {
                                                                        v173 := GenS18{F0: -1, F1: v57}
                                                                        _ = v173
                                                                    }
                                                                    v1 += (v1 + 1)
                                                                    _ = v172
                                                                }
                                                            }
                                                        }
                                                        v3["hello"] = ("" + "")
                                                        v57 = (v57 & -865)
                                                        var v174 int16 = (v57 % ((v57) | 1))
                                                        _ = v163
                                                        _ = v174
                                                    } else {
                                                        var v175 bool = (1 != 0)
                                                        v3[""] = ("\n\t" + "\n\t")
                                                        _ = v175
                                                    }
                                                    v3[""] = "a"
                                                }
                                                v1 -= v1
                                            }
                                            if false {
                                                fmt.Println(fi58, fi58)
                                                var v176 int = (^(fi51))
                                                var v177 string = "\n\t"
                                                var v178 float64 = 0.0
                                                _ = v176
                                                _ = v177
                                                _ = v178
                                            } else {
                                                v179 := GenS3{F0: "hello"}
                                                var v180 int8 = -1
                                                _ = v179
                                                _ = v180
                                            }
                                        } else {
                                            var v181 uint = 0
                                            if (42 >= 42) {
                                                fmt.Println(v181, v181)
                                                var v182 float32 = 1e10
                                                _ = v182
                                            }
                                            var v183 float32 = 1.5
                                            _ = v181
                                            _ = v183
                                        }
                                    default:
                                        if true {
                                            v184 := GenS11{F0: 0}
                                            _ = v184
                                        } else {
                                            var v185 int8 = -102
                                            v186 := GenS19{F0: ((fi58 / ((9223372036854775807) | 1)) != fi51)}
                                            _ = v185
                                            _ = v186
                                        }
                                    }
                                }
                            }
                        default:
                            v187 := v3[("\n\t" + "")]
                            if len(v2) > 0 {
                                v2[0] = fi51
                            }
                            _ = v187
                        }
                        v188 := GenS20{F0: fi51, F1: 867, F2: ""}
                        v1 += 2
                        _ = v188
                    }
                    v57 -= (-(v57))
                    _ = fi58
                }
                _ = v57
            }
            fmt.Println(v1, v1)
            _ = fi51
        }
        v1 -= (^(v1))
    }
    var v189 float32 = (1.5 - -2.25)
    if false {
        var v190 int32 = 1
        for fi191 := 0; fi191 < 17; fi191++ {
            if (true || (!((0 < -1) || false))) {
                v192 := GenS18{F0: 804, F1: 42}
                v190 += v190
                for fi193 := 0; fi193 < 12; fi193++ {
                    var v194 int64 = 239
                    var v195 uint8 = 1
                    _ = fi193
                    _ = v194
                    _ = v195
                }
                _ = v192
            } else {
                v1 = v1
                v196 := v3[(("\n\t" + "\n\t") + "foo bar")]
                _ = v196
            }
            v197 := v3["\n\t"]
            _ = fi191
            _ = v197
        }
        var v198 uint32 = 4294967295
        _ = v190
        _ = v198
    }
    v1 -= v1
    fmt.Println(v189)
    for fi199 := 0; fi199 < 14; fi199++ {
        if false {
            v189 -= v189
            for fi200 := 0; fi200 < 15; fi200++ {
                var v201 uint = 2
                var v202 uint = v201
                _ = fi200
                _ = v201
                _ = v202
            }
            switch "\n\t" {
            case "foo bar":
                var v203 string = "hello"
                _ = v203
            }
        }
        switch (fi199 << ((2) & 63)) {
        case 2147483647:
            v1 += 808
        case 4294967295:
            for fi204 := 0; fi204 < 17; fi204++ {
                fmt.Println(v189, fi199, fi199)
                v205 := GenS20{F0: (fi199 ^ 127), F1: 1, F2: ("foo bar" + (("\n\t" + "a") + "foo bar"))}
                var v206 uint32 = 1
                var v207 float64 = -2.25
                _ = fi204
                _ = v205
                _ = v206
                _ = v207
            }
        default:
            fmt.Println(v189, fi199, v189)
        }
        var v208 float64 = 0.0
        _ = fi199
        _ = v208
    }
    if (42 < 42) {
        v209 := GenS5{F0: 128, F1: 751}
        _ = v209
    }
    var v210 float64 = (1.5 + float64(3.14159))
    v189 -= ((v189 + ((float32(v210) - (0.0 - (((float32(v210) - v189) / v189) + v189))) - 1.5)) / 1.5)
    v1 = (v1 + 462)
    var v211 int32 = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v189:", v189)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v210:", v210)
    fmt.Println("v211:", v211)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v189
    _ = v2
    _ = v210
    _ = v211
    _ = v3
}
