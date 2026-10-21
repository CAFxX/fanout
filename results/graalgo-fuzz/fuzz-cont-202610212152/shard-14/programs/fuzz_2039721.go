package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 int32 }
type GenS2 struct { F0 uint16; F1 int32; F2 int64 }
type GenS3 struct { F0 uint32; F1 string }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int16; F1 int16; F2 uint32 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 uint32; F1 uint8 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 uint64; F1 int }
type GenS13 struct { F0 int; F1 int; F2 uint16 }
type GenS14 struct { F0 bool }
type GenS15 struct { F0 int8; F1 int8 }
type GenS16 struct { F0 int16; F1 bool }
type GenS17 struct { F0 uint32 }
type GenS18 struct { F0 uint64; F1 bool }
type GenS19 struct { F0 uint }
type GenS20 struct { F0 uint32; F1 uint }
type GenS21 struct { F0 int16; F1 uint8 }
func genFn0(p0 uint16, p1 int64) bool {
    v5 := GenS1{F0: -128, F1: "\n\t", F2: -209}
    var v6 int32 = 0
    p0 -= (p0 + (p0 % ((p0) | 1)))
    switch -9223372036854775808 {
    case 477:
        var v7 uint16 = (p0 | (p0 & p0))
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return (-1 != p1)
}
func genFn1() int64 {
    v8 := []int{0, 42, -2147483648}
    var v9 int8 = 1
    _ = v8
    _ = v9
    return -9223372036854775808
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]int)
    v3 := []uint32{2}
    var v4 uint64 = 2
    v4 += (v4 << ((65535) & 63))
    fmt.Println(v4)
    for fi10 := 0; fi10 < 7; fi10++ {
        v11 := GenS1{F0: 0, F1: "\n\t", F2: 0}
        _ = fi10
        _ = v11
    }
    fmt.Println(v4)
    if (-472 != genFn1()) {
        v4 = v4
        v2[-895] = 4294967295
        v4 -= (v4 / (((^(v4))) | 1))
    } else {
        var v12 uint8 = 128
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi13 := 0; fi13 < 7; fi13++ {
            if ((fi13 < fi13) && true) {
                var v14 bool = ((0 != 171) || false)
                if v14 {
                    var v15 int64 = -329
                    v16 := GenS1{F0: 127, F1: ("hello" + (("\n\t" + "hello") + ("" + (((("foo bar" + (("" + "") + "a")) + ("" + ((("\n\t" + "foo bar") + ("hello" + "a")) + "a"))) + ("foo bar" + ("hello" + "foo bar"))) + ("\n\t" + ("a" + ("a" + "a"))))))), F2: 1}
                    for fi17 := 0; fi17 < 6; fi17++ {
                        v14 = (v14 && v14)
                        if (("a" + "\n\t") != "") {
                            var v18 float32 = 1.5
                            _ = v18
                        } else {
                            v19 := GenS2{F0: 42, F1: -2147483648, F2: v15}
                            if false {
                                switch fi17 {
                                case 42:
                                    var v20 float32 = 0.0
                                    _ = v20
                                case -9223372036854775808:
                                    fmt.Println(fi17, v12, v4)
                                    v12 += 2
                                default:
                                    v21 := GenS2{F0: 32767, F1: -594, F2: int64(v12)}
                                    _ = v21
                                }
                                var v22 int = fi17
                                if (v14 || v14) {
                                    if (genFn1() > v15) {
                                        switch "a" {
                                        case "foo bar":
                                            if false {
                                                v4 = (v4 ^ 2)
                                                var v23 string = "foo bar"
                                                v24 := GenS3{F0: 42, F1: "foo bar"}
                                                _ = v23
                                                _ = v24
                                            } else {
                                                if v14 {
                                                    fmt.Println(v4)
                                                    fmt.Println(fi17)
                                                }
                                                var v25 uint16 = 1
                                                fmt.Println(v14, fi13)
                                                switch (^(fi13)) {
                                                case 32767:
                                                    if (1 > 2147483647) {
                                                        var v26 uint64 = 2
                                                        var v27 uint64 = v26
                                                        v28 := GenS4{F0: 65535}
                                                        _ = v26
                                                        _ = v27
                                                        _ = v28
                                                    }
                                                case -79:
                                                    if (false || (!v14)) {
                                                        switch fi17 {
                                                        case 2147483647:
                                                            fmt.Println(v12, v15, v22)
                                                        }
                                                        if ((32767 != 32767) || true) {
                                                            fmt.Println(v12, v22)
                                                        }
                                                        v12 = 127
                                                        v12 = 127
                                                    } else {
                                                        fmt.Println(v12, v12)
                                                        v29 := GenS2{F0: 32767, F1: 42, F2: (v15 | v15)}
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        v12 -= 2
                                                        _ = v29
                                                    }
                                                }
                                                var v30 uint = 0
                                                _ = v25
                                                _ = v30
                                            }
                                        case "hello":
                                            fmt.Println(fi17, fi17)
                                        }
                                        var v31 int64 = 42
                                        _ = v31
                                    }
                                    if true {
                                        v32 := GenS3{F0: 2147483647, F1: "hello"}
                                        _ = v32
                                    } else {
                                        fmt.Println(v14)
                                        if (!(!v14)) {
                                            var v33 int = -569
                                            var v34 uint16 = 32767
                                            v35 := GenS2{F0: (^(v34)), F1: 431, F2: v15}
                                            _ = v33
                                            _ = v34
                                            _ = v35
                                        }
                                        _ = len(v1)
                                    }
                                    switch fi13 {
                                    case 32768:
                                        fmt.Println(fi17, v22)
                                    }
                                    if (0 <= 0) {
                                        delete(v2, fi13)
                                        v36 := v1[fi17]
                                        _ = v36
                                    }
                                    switch "\n\t" {
                                    case "hello":
                                        var v37 int8 = -1
                                        _ = v37
                                    }
                                } else {
                                    fmt.Println(fi13, v12, v12)
                                }
                                _ = v22
                            }
                            v14 = (!v14)
                            _ = v19
                        }
                        v38 := GenS2{F0: 656, F1: 42, F2: (-(v15))}
                        if v14 {
                            _ = len(v3)
                            if v14 {
                                v39 := GenS5{F0: v12}
                                _ = v39
                            } else {
                                v4 = (v4 * (v4 + v4))
                            }
                            if (!(-2.25 == float32(1e10))) {
                                v15 = (v15 % (((v15 * (v15 + v15))) | 1))
                                v40 := GenS3{F0: 969, F1: "hello"}
                                switch -31 {
                                case 32768:
                                    var v41 uint16 = 32767
                                    _ = v41
                                case 973:
                                    v12 -= (v12 >> ((546) & 63))
                                default:
                                    var v42 uint8 = (v12 | (v12 & v12))
                                    _ = v42
                                }
                                _ = v40
                            } else {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                        }
                        switch fi17 {
                        case 4294967295:
                            v4 = 1
                        case 2147483647:
                            v4 = (v4 % ((v4) | 1))
                        case 32767:
                            var v43 string = "a"
                            _ = v43
                        default:
                            if true {
                                var v44 int8 = 0
                                _ = v44
                            }
                        }
                        var v45 int = (fi13 >> ((4294967295) & 63))
                        _ = fi17
                        _ = v38
                        _ = v45
                    }
                    _ = v15
                    _ = v16
                } else {
                    if (float32(0.1) != 0.0) {
                        switch "hello" {
                        case "":
                            v46 := GenS6{F0: 1000, F1: 1000, F2: 1}
                            _ = v46
                        case "foo bar":
                            if v14 {
                                if false {
                                    delete(v1, fi13)
                                    v14 = true
                                    _ = len(v3)
                                }
                                v12 = 2
                            }
                        case "a":
                            for fi47 := 0; fi47 < 11; fi47++ {
                                switch "" {
                                case "":
                                    var v48 uint = 0
                                    _ = v48
                                case "hello":
                                    v49 := GenS3{F0: 0, F1: ""}
                                    _ = v49
                                case "\n\t":
                                    switch (^(fi13)) {
                                    case 0:
                                        if false {
                                            fmt.Println(v4)
                                            v12 = uint8(v4)
                                            v14 = v14
                                            v50 := GenS6{F0: -32768, F1: -32768, F2: 2}
                                            _ = v50
                                        } else {
                                            v4 += 0
                                            if v14 {
                                                if false {
                                                    v12 += 1
                                                    switch 128 {
                                                    case -1:
                                                        switch (fi47 - (fi13 + (^(fi13)))) {
                                                        case -346:
                                                            var v51 float32 = 0.0
                                                            if ((!(0 > 893)) || (969 < 1)) {
                                                                switch 32767 {
                                                                case 2147483647:
                                                                    v52 := GenS7{F0: 4294967295}
                                                                    _ = v52
                                                                case 470:
                                                                    if (42 != 2147483648) {
                                                                        if true {
                                                                            v12 = 1
                                                                            if len(v3) > 0 {
                                                                                v3[0] = 744
                                                                            }
                                                                            v4 -= v4
                                                                            var v53 float64 = (float64((v51 * float32(1.5))) * (0.1 + 1e10))
                                                                            _ = len(v2)
                                                                            v14 = (!(!(1 > 127)))
                                                                            _ = v53
                                                                        }
                                                                        var v54 int = fi47
                                                                        v4 += 122
                                                                        _ = v54
                                                                    }
                                                                case 128:
                                                                    v14 = true
                                                                }
                                                                if (42 < 4294967295) {
                                                                    v12 = v12
                                                                    if ((!v14) && v14) {
                                                                        if (v14 || v14) {
                                                                            v12 += (v12 % ((v12) | 1))
                                                                        } else {
                                                                            v55 := GenS5{F0: (^(v12))}
                                                                            v4 = v4
                                                                            var v56 uint64 = v4
                                                                            if true {
                                                                                fmt.Println(v12, v51)
                                                                                var v57 float32 = 0.1
                                                                                switch "a" {
                                                                                case "hello":
                                                                                    v51 -= (float32(((0.1 / 0.1) / 1e10)) / (v51 * v57))
                                                                                case "a":
                                                                                    v58 := GenS7{F0: 17}
                                                                                    _ = v58
                                                                                case "foo bar":
                                                                                    fmt.Println(v14, v4)
                                                                                }
                                                                                var v59 int16 = 1000
                                                                                _ = v57
                                                                                _ = v59
                                                                            } else {
                                                                                v60 := GenS6{F0: -848, F1: -368, F2: uint32(v4)}
                                                                                if (!v14) {
                                                                                    if false {
                                                                                        var v61 uint = 0
                                                                                        if v14 {
                                                                                            var v62 uint8 = v12
                                                                                            var v63 int8 = int8(v61)
                                                                                            _ = v62
                                                                                            _ = v63
                                                                                        } else {
                                                                                            fmt.Println(v4, fi13, v61)
                                                                                            var v64 float64 = (0.1 / 1e10)
                                                                                            _ = v64
                                                                                        }
                                                                                        _ = v61
                                                                                    } else {
                                                                                        v1[fi13] = ""
                                                                                    }
                                                                                    fmt.Println(fi13, fi47, v12)
                                                                                    v4 -= 42
                                                                                } else {
                                                                                    switch ("hello" + "foo bar") {
                                                                                    case "":
                                                                                        v2[42] = (fi47 * 738)
                                                                                    case "hello":
                                                                                        switch 413 {
                                                                                        case 723:
                                                                                            var v65 int = 1
                                                                                            _ = v65
                                                                                        }
                                                                                    }
                                                                                }
                                                                                _ = v60
                                                                            }
                                                                            var v66 uint64 = v4
                                                                            _ = v55
                                                                            _ = v56
                                                                            _ = v66
                                                                        }
                                                                        fmt.Println(fi13, fi47, v14)
                                                                    } else {
                                                                        v4 -= (v4 * (v4 % ((v4) | 1)))
                                                                        var v67 int8 = 15
                                                                        v68 := GenS8{F0: uint(v4)}
                                                                        _ = v67
                                                                        _ = v68
                                                                    }
                                                                    var v69 float64 = ((3.14159 - -2.25) - 0.1)
                                                                    _ = v69
                                                                } else {
                                                                    v14 = true
                                                                    delete(v2, fi13)
                                                                    v4 += (v4 ^ 0)
                                                                }
                                                            } else {
                                                                if v14 {
                                                                    v12 = (v12 % ((v12) | 1))
                                                                    var v70 int32 = -2147483648
                                                                    delete(v1, 128)
                                                                    if v14 {
                                                                        if (("foo bar" <= ("" + (("" + (((("hello" + "\n\t") + "") + "") + "")) + "hello"))) || true) {
                                                                            var v71 uint16 = 845
                                                                            _ = v71
                                                                        }
                                                                    }
                                                                    _ = v70
                                                                } else {
                                                                    v4 -= (^(v4))
                                                                    if false {
                                                                        var v72 int8 = 100
                                                                        v12 += 127
                                                                        v73 := GenS2{F0: 0, F1: -1, F2: 42}
                                                                        fmt.Println(v72, fi47, v51)
                                                                        switch "a" {
                                                                        case "hello":
                                                                            _ = len(v3)
                                                                        }
                                                                        _ = v72
                                                                        _ = v73
                                                                    } else {
                                                                        v51 -= (v51 * (float32(0.1) + v51))
                                                                        v74 := GenS1{F0: 100, F1: "\n\t", F2: 1}
                                                                        v4 = (v4 + v4)
                                                                        _ = v74
                                                                    }
                                                                    if ("\n\t" < "\n\t") {
                                                                        fmt.Println(v51, v51)
                                                                        v51 += (v51 / (v51 / v51))
                                                                    }
                                                                }
                                                                v3 = append(v3, 42)
                                                            }
                                                            v75 := GenS9{F0: 2147483648, F1: (v12 * (v12 << ((18446744073709551615) & 63)))}
                                                            _ = v51
                                                            _ = v75
                                                        case 2147483647:
                                                            v14 = (!((-2.25 > 3.14159) && v14))
                                                        case 869:
                                                            v76 := GenS3{F0: 65535, F1: "foo bar"}
                                                            _ = v76
                                                        }
                                                    }
                                                    if len(v3) > 0 {
                                                        v3[0] = 656
                                                    }
                                                }
                                                fmt.Println(v4, fi47)
                                                v3 = append(v3, 65535)
                                            }
                                            if (fi47 >= -942) {
                                                var v77 uint32 = 144
                                                if v14 {
                                                    if (("\n\t" + "hello") != "foo bar") {
                                                        var v78 int32 = 2147483647
                                                        v77 += v77
                                                        _ = v78
                                                    }
                                                }
                                                _ = v77
                                            } else {
                                                if true {
                                                    v79 := GenS10{F0: int32(v4)}
                                                    if len(v3) > 0 {
                                                        v3[0] = 330
                                                    }
                                                    if ("" < (("foo bar" + ("a" + "\n\t")) + "\n\t")) {
                                                        switch ("" + "") {
                                                        case "a":
                                                            switch 9223372036854775807 {
                                                            case 32767:
                                                                _ = len(v2)
                                                            }
                                                        case "hello":
                                                            if (32767 == -454) {
                                                                v12 += 4
                                                                var v80 bool = (!v14)
                                                                _ = v80
                                                            }
                                                        case "\n\t":
                                                            switch fi47 {
                                                            case -9223372036854775808:
                                                                fmt.Println(v4, fi47, v4)
                                                            case 4294967295:
                                                                var v81 uint64 = (v4 + v4)
                                                                _ = v81
                                                            case 127:
                                                                var v82 int = fi47
                                                                _ = v82
                                                            }
                                                        default:
                                                            v4 += (v4 & v4)
                                                        }
                                                        switch ("a" + ("hello" + "\n\t")) {
                                                        case "foo bar":
                                                            fmt.Println(fi47)
                                                        default:
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                        }
                                                        if (127 <= v12) {
                                                            v83 := GenS11{F0: int64(v4)}
                                                            var v84 string = ("hello" + "hello")
                                                            v12 -= 255
                                                            v4 += 2
                                                            var v85 uint = 42
                                                            _ = v83
                                                            _ = v84
                                                            _ = v85
                                                        } else {
                                                            var v86 string = ("\n\t" + (("a" + "\n\t") + "a"))
                                                            _ = v86
                                                        }
                                                        v87 := GenS1{F0: 1, F1: "hello", F2: -2147483648}
                                                        _ = v87
                                                    }
                                                    _ = v79
                                                } else {
                                                    v14 = ((!v14) && true)
                                                    v14 = v14
                                                    switch "hello" {
                                                    case "a":
                                                        switch fi47 {
                                                        case 32768:
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                        case 32767:
                                                            if v14 {
                                                                v88 := GenS9{F0: 937, F1: (v12 | 2)}
                                                                switch 128 {
                                                                case 32767:
                                                                    var v89 int8 = 100
                                                                    _ = v89
                                                                case 0:
                                                                    if len(v3) > 0 {
                                                                        v3[0] = 2
                                                                    }
                                                                case 128:
                                                                    v1[2147483647] = "\n\t"
                                                                default:
                                                                    v90 := GenS1{F0: 96, F1: "a", F2: -170}
                                                                    _ = v90
                                                                }
                                                                v91 := GenS3{F0: 4294967295, F1: (("hello" + "hello") + "foo bar")}
                                                                _ = v88
                                                                _ = v91
                                                            }
                                                        default:
                                                            var v92 uint32 = 2147483647
                                                            _ = v92
                                                        }
                                                    }
                                                }
                                                if (-2.25 < 0.0) {
                                                    fmt.Println(v4, v14, fi13)
                                                }
                                                if false {
                                                    var v93 uint64 = 4294967295
                                                    fmt.Println(v93, v14, v4)
                                                    switch 4294967295 {
                                                    case 42:
                                                        _ = len(v3)
                                                    case 507:
                                                        v94 := GenS1{F0: -128, F1: "", F2: 0}
                                                        _ = v94
                                                    default:
                                                        var v95 int64 = -9223372036854775808
                                                        _ = v95
                                                    }
                                                    _ = v93
                                                } else {
                                                    fmt.Println(v12)
                                                }
                                            }
                                        }
                                    case 9223372036854775807:
                                        fmt.Println(v14)
                                    case -716:
                                        v96 := GenS7{F0: 830}
                                        _ = v96
                                    }
                                default:
                                    v12 -= 255
                                }
                                _ = fi47
                            }
                        default:
                            v14 = v14
                        }
                        for fi97 := 0; fi97 < 18; fi97++ {
                            fmt.Println(v14, v12)
                            v98 := GenS5{F0: v12}
                            _ = fi97
                            _ = v98
                        }
                    } else {
                        v12 -= (v12 % ((v12) | 1))
                    }
                    switch (fi13 << ((uint(v4)) & 63)) {
                    case 434:
                        v14 = true
                    case 9223372036854775807:
                        for fi99 := 0; fi99 < 12; fi99++ {
                            var v100 float32 = 3.14159
                            v101 := GenS12{F0: 996, F1: fi13}
                            _ = fi99
                            _ = v100
                            _ = v101
                        }
                    }
                    v102 := GenS13{F0: -9223372036854775808, F1: (fi13 * (fi13 - -9223372036854775808)), F2: 793}
                    _ = v102
                }
                switch "a" {
                case "foo bar":
                    v12 = (v12 * v12)
                case "a":
                    var v103 uint16 = 255
                    _ = v103
                case "hello":
                    var v104 uint64 = 0
                    _ = v104
                default:
                    v105 := GenS1{F0: 127, F1: (("foo bar" + "\n\t") + ""), F2: int32(fi13)}
                    _ = v105
                }
                _ = v14
            } else {
                for fi106 := 0; fi106 < 19; fi106++ {
                    var v107 uint32 = 42
                    if false {
                        v108 := GenS11{F0: -6}
                        _ = len(v3)
                        var v109 uint8 = uint8(v107)
                        _ = v108
                        _ = v109
                    } else {
                        _ = len(v2)
                        v3 = append(v3, v107)
                        if (!false) {
                            v110 := GenS11{F0: -9223372036854775808}
                            switch "foo bar" {
                            case "\n\t":
                                if (!((v107 | v107) >= v107)) {
                                    if (65535 <= 708) {
                                        switch (("a" + (("a" + "\n\t") + "")) + ("" + "hello")) {
                                        case "":
                                            if false {
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                if true {
                                                    v12 = 249
                                                    var v111 float32 = 1e10
                                                    var v112 uint = 18446744073709551615
                                                    _ = v111
                                                    _ = v112
                                                }
                                            } else {
                                                v107 -= v107
                                                v107 = v107
                                                v113 := GenS5{F0: (v12 % ((v12) | 1))}
                                                v114 := GenS1{F0: 23, F1: (("foo bar" + "") + ""), F2: 42}
                                                _ = v113
                                                _ = v114
                                            }
                                        }
                                    }
                                    var v115 string = "\n\t"
                                    _ = v115
                                }
                            default:
                                v116 := GenS14{F0: false}
                                _ = v116
                            }
                            _ = v110
                        }
                    }
                    _ = fi106
                    _ = v107
                }
                v117 := GenS2{F0: 291, F1: -1, F2: genFn1()}
                var v118 uint64 = v4
                _ = v117
                _ = v118
            }
            v4 += (v4 * (^(v4)))
            _ = fi13
        }
        _ = v12
    }
    switch "foo bar" {
    case "":
        for fi119 := 0; fi119 < 4; fi119++ {
            var v120 float64 = 0.1
            for fi121 := 0; fi121 < 2; fi121++ {
                v122 := GenS15{F0: 1, F1: -89}
                _ = fi121
                _ = v122
            }
            v123 := v2[fi119]
            _ = fi119
            _ = v120
            _ = v123
        }
    case "\n\t":
        v4 = (v4 / ((v4) | 1))
    default:
        fmt.Println(v4)
    }
    v124 := GenS16{F0: 32767, F1: true}
    _ = len(v1)
    v125 := GenS11{F0: 9223372036854775807}
    for fi126 := 0; fi126 < 13; fi126++ {
        var v127 int32 = -364
        _ = fi126
        _ = v127
    }
    if (!(!(-9223372036854775808 >= 265))) {
        for fi128 := 0; fi128 < 20; fi128++ {
            if (-2.25 <= 0.0) {
                var v129 float64 = (1e10 / -2.25)
                switch fi128 {
                case 32768:
                    v130 := GenS6{F0: 1000, F1: 0, F2: 805}
                    _ = v130
                case 292:
                    var v131 uint16 = 42
                    _ = v131
                case 4294967295:
                    switch 32767 {
                    case 869:
                        var v132 int8 = -1
                        _ = v132
                    case -2147483648:
                        for fi133 := 0; fi133 < 14; fi133++ {
                            if (("" + ((("\n\t" + ("\n\t" + "\n\t")) + ("a" + ("hello" + ((("hello" + "a") + "hello") + ("a" + (("a" + ("\n\t" + "hello")) + "foo bar")))))) + "a")) <= (("" + ("\n\t" + "foo bar")) + "")) {
                                if false {
                                    var v134 float64 = 0.1
                                    _ = v134
                                } else {
                                    v135 := GenS17{F0: 897}
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    _ = v135
                                }
                            }
                            switch (("" + "hello") + "foo bar") {
                            case "a":
                                v2[fi128] = 514
                            case "hello":
                                if false {
                                    if (2 == 387) {
                                        var v136 int8 = -48
                                        var v137 int = (fi133 << ((0) & 63))
                                        _ = v136
                                        _ = v137
                                    }
                                    delete(v1, fi133)
                                    if false {
                                        if true {
                                            if false {
                                                v4 = v4
                                                if true {
                                                    switch ("a" + (("hello" + "a") + "\n\t")) {
                                                    case "a":
                                                        fmt.Println(v4)
                                                    case "hello":
                                                        v138 := GenS18{F0: v4, F1: false}
                                                        _ = v138
                                                    case "\n\t":
                                                        switch (-(fi133)) {
                                                        case 42:
                                                            fmt.Println(fi133, v129)
                                                        case 2147483647:
                                                            v4 -= 18446744073709551615
                                                        case 9223372036854775807:
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                        }
                                                    }
                                                    if ((v4 - v4) != (v4 * (v4 + (^(v4))))) {
                                                        var v139 int32 = 0
                                                        var v140 int16 = 32767
                                                        v141 := v1[fi128]
                                                        _ = v139
                                                        _ = v140
                                                        _ = v141
                                                    } else {
                                                        v4 = 1
                                                        v142 := GenS8{F0: 42}
                                                        _ = v142
                                                    }
                                                    var v143 float32 = 3.14159
                                                    _ = v143
                                                } else {
                                                    fmt.Println(v4, fi128, fi133)
                                                    v4 -= (v4 | 744)
                                                    if (-17 >= -9223372036854775808) {
                                                        v129 = float64(0.0)
                                                        v4 = (v4 & v4)
                                                        var v144 int64 = 0
                                                        _ = v144
                                                    }
                                                }
                                            }
                                            v145 := GenS8{F0: 729}
                                            _ = v145
                                        }
                                        v1[fi128] = ("a" + "foo bar")
                                        v146 := GenS3{F0: 65535, F1: ("\n\t" + "\n\t")}
                                        _ = v146
                                    } else {
                                        v3 = append(v3, 384)
                                        if (1 != -125) {
                                            v147 := GenS14{F0: (0 == 889)}
                                            var v148 float32 = (float32(((0.1 - (v129 / float64(((3.14159 * (-2.25 + -2.25)) * ((1e10 + 0.0) / 0.1))))) / 0.1)) - (1.5 * 3.14159))
                                            _ = v147
                                            _ = v148
                                        } else {
                                            var v149 int64 = 42
                                            v2[622] = 32767
                                            var v150 uint32 = 2147483647
                                            v129 = 0.0
                                            v129 -= v129
                                            switch fi133 {
                                            case 2147483647:
                                                fmt.Println(fi133, v150)
                                            }
                                            _ = v149
                                            _ = v150
                                        }
                                    }
                                } else {
                                    v151 := GenS4{F0: 454}
                                    switch "a" {
                                    case "foo bar":
                                        v4 = 2
                                    default:
                                        var v152 uint64 = v4
                                        _ = v152
                                    }
                                    v153 := v2[1]
                                    _ = v151
                                    _ = v153
                                }
                            case "":
                                v4 -= 0
                            default:
                                if (-106 < 100) {
                                    v4 += v4
                                }
                            }
                            fmt.Println(fi133, v4)
                            _ = fi133
                        }
                    case -188:
                        var v154 uint32 = 879
                        _ = v154
                    }
                }
                _ = v129
            }
            v155 := GenS19{F0: 0}
            if (!(!true)) {
                var v156 uint = 255
                v3 = append(v3, 4294967295)
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = v156
            } else {
                v157 := v2[(fi128 + 32767)]
                v158 := GenS20{F0: 590, F1: 313}
                _ = v157
                _ = v158
            }
            _ = fi128
            _ = v155
        }
        var v159 uint = 18446744073709551615
        v4 -= v4
        _ = v159
    } else {
        delete(v1, 9223372036854775807)
        if (!(true && (3.14159 == 1e10))) {
            switch "\n\t" {
            case "foo bar":
                switch 1 {
                case 1:
                    var v160 uint = 18446744073709551615
                    _ = v160
                case -2147483648:
                    var v161 int8 = 97
                    _ = v161
                case 32768:
                    v4 += (v4 & (v4 / ((2) | 1)))
                default:
                    if len(v3) > 0 {
                        v3[0] = 42
                    }
                }
            }
            var v162 uint32 = 65535
            _ = v162
        } else {
            if (-2147483648 < 0) {
                for fi163 := 0; fi163 < 11; fi163++ {
                    for fi164 := 0; fi164 < 17; fi164++ {
                        var v165 float64 = 0.0
                        _ = fi164
                        _ = v165
                    }
                    _ = fi163
                }
            }
        }
    }
    v166 := GenS12{F0: 2, F1: 554}
    switch "a" {
    case "hello":
        if false {
            v4 += 1
        } else {
            v167 := GenS21{F0: 32767, F1: 1}
            _ = v167
        }
    case "foo bar":
        var v168 int8 = -128
        _ = v168
    }
    v169 := GenS19{F0: 255}
    var v170 uint8 = 42
    v171 := GenS21{F0: -297, F1: v170}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v170:", v170)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v124
    _ = v125
    _ = v166
    _ = v169
    _ = v170
    _ = v171
    _ = v2
    _ = v3
    _ = v4
}
