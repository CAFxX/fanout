package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 int8 }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 string; F1 bool }
type GenS4 struct { F0 int8; F1 uint64 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int16; F1 string }
type GenS7 struct { F0 int; F1 uint16 }
type GenS8 struct { F0 uint32; F1 string }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 uint64; F1 uint; F2 int32 }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 uint; F1 bool; F2 uint64 }
type GenS13 struct { F0 uint8; F1 uint64 }
type GenS14 struct { F0 uint }
type GenS15 struct { F0 int8; F1 int64 }

func main() {
    var v1 uint32 = 0
    var v2 int16 = -32768
    var v3 int32 = -1
    v4 := make(map[uint32]bool)
    fmt.Println(v1, v1)
    v5 := GenS1{F0: -1, F1: v1, F2: 127}
    for fi6 := 0; fi6 < 5; fi6++ {
        var v7 int = 32768
        var v8 string = (((("a" + "\n\t") + ("a" + "\n\t")) + "") + (((("" + ("" + (("\n\t" + "hello") + "foo bar"))) + "hello") + "foo bar") + ("a" + "\n\t")))
        _ = fi6
        _ = v7
        _ = v8
    }
    v3 = (v3 - (v3 & v3))
    var v9 int8 = 106
    var v10 int16 = (v2 - -1)
    v4[1] = true
    v11 := []int{}
    v1 = 2147483647
    for fi12 := 0; fi12 < 14; fi12++ {
        v1 += 42
        switch "\n\t" {
        case "\n\t":
            if false {
                v2 = (v10 * v2)
                v13 := GenS2{F0: "hello", F1: 51}
                v4[(v1 ^ 42)] = ((!(!((v9 / (((^(v9))) | 1)) >= 1))) && true)
                _ = v13
            } else {
                if (true || (false || (!(true && true)))) {
                    var v14 uint = 4294967295
                    if ("foo bar" > ("\n\t" + ("hello" + ("hello" + "\n\t")))) {
                        v3 = (^(v3))
                        fmt.Println(v2, fi12, v2)
                    } else {
                        for fi15 := 0; fi15 < 14; fi15++ {
                            _ = len(v11)
                            fmt.Println(v2)
                            _ = fi15
                        }
                    }
                    _ = v14
                } else {
                    switch (fi12 * (fi12 % ((0) | 1))) {
                    case -804:
                        if (("" < "foo bar") && true) {
                            if false {
                                v16 := GenS3{F0: "\n\t", F1: (!false)}
                                var v17 int64 = 42
                                _ = v16
                                _ = v17
                            }
                            var v18 uint64 = 0
                            _ = v18
                        }
                    case -703:
                        fmt.Println(fi12, v10)
                    case -872:
                        for fi19 := 0; fi19 < 8; fi19++ {
                            if true {
                                if (118 < -128) {
                                    fmt.Println(fi19, v10, v2)
                                    v2 = v2
                                } else {
                                    if (17 < 205) {
                                        if true {
                                            var v20 string = ""
                                            _ = v20
                                        }
                                    } else {
                                        if false {
                                            switch "foo bar" {
                                            case "hello":
                                                if true {
                                                    v21 := GenS3{F0: "a", F1: true}
                                                    fmt.Println(v3, v9)
                                                    fmt.Println(v3)
                                                    if (88 >= 4294967295) {
                                                        v22 := GenS4{F0: v9, F1: 2}
                                                        delete(v4, v1)
                                                        v2 = 42
                                                        _ = v22
                                                    } else {
                                                        v23 := GenS5{F0: 65535}
                                                        v24 := GenS3{F0: "", F1: (true && false)}
                                                        v9 += (v9 / ((v9) | 1))
                                                        if (935 <= 2) {
                                                            v2 += v10
                                                            if (!((v1 + (v1 << ((4294967295) & 63))) != (v1 * (v1 - 0)))) {
                                                                var v25 float64 = (3.14159 + float64(3.14159))
                                                                _ = len(v4)
                                                                _ = v25
                                                            } else {
                                                                var v26 int8 = -90
                                                                _ = v26
                                                            }
                                                        } else {
                                                            var v27 int = (fi19 / ((fi19) | 1))
                                                            v3 += v3
                                                            _ = v27
                                                        }
                                                        _ = v23
                                                        _ = v24
                                                    }
                                                    var v28 uint16 = 0
                                                    switch 1 {
                                                    case 1:
                                                        if (true && true) {
                                                            if (false || ((^(v2)) <= v2)) {
                                                                var v29 uint = 739
                                                                switch -1 {
                                                                case -9223372036854775808:
                                                                    var v30 int = int(v1)
                                                                    switch ("foo bar" + (("" + "foo bar") + "foo bar")) {
                                                                    case "\n\t":
                                                                        v31 := GenS2{F0: "foo bar", F1: 128}
                                                                        _ = v31
                                                                    case "hello":
                                                                        if len(v11) >= 2 {
                                                                            v11 = v11[1:len(v11)]
                                                                        }
                                                                    default:
                                                                        switch (("" + "hello") + "foo bar") {
                                                                        case "foo bar":
                                                                            var v32 uint = (v29 + 18446744073709551615)
                                                                            _ = v32
                                                                        case "a":
                                                                            v33 := v4[v1]
                                                                            _ = v33
                                                                        }
                                                                    }
                                                                    _ = v30
                                                                case 32768:
                                                                    if len(v11) > 0 {
                                                                        v11[0] = fi19
                                                                    }
                                                                case 1:
                                                                    fmt.Println(v28, v3, fi12)
                                                                default:
                                                                    var v34 int8 = (v9 | v9)
                                                                    _ = v34
                                                                }
                                                                fmt.Println(fi19, v9, v28)
                                                                _ = v29
                                                            } else {
                                                                v35 := GenS6{F0: v2, F1: (((("a" + "a") + (("hello" + "\n\t") + ("hello" + "\n\t"))) + "") + "hello")}
                                                                var v36 uint32 = v1
                                                                if false {
                                                                    v37 := v4[v36]
                                                                    _ = v37
                                                                } else {
                                                                    if false {
                                                                        fmt.Println(fi19)
                                                                    } else {
                                                                        v38 := GenS7{F0: 32768, F1: v28}
                                                                        var v39 uint16 = 65535
                                                                        _ = v38
                                                                        _ = v39
                                                                    }
                                                                    switch (fi12 ^ fi19) {
                                                                    case 0:
                                                                        var v40 int32 = v3
                                                                        _ = v40
                                                                    case -1:
                                                                        v1 -= (v36 / ((v36) | 1))
                                                                    }
                                                                    v28 = v28
                                                                }
                                                                var v41 int8 = 0
                                                                _ = v35
                                                                _ = v36
                                                                _ = v41
                                                            }
                                                            switch fi12 {
                                                            case 0:
                                                                var v42 uint8 = 2
                                                                _ = v42
                                                            default:
                                                                var v43 uint32 = (v1 - v1)
                                                                v3 += (^(v3))
                                                                _ = v43
                                                            }
                                                        } else {
                                                            v44 := v4[(^(v1))]
                                                            switch fi12 {
                                                            case -2147483648:
                                                                v45 := GenS1{F0: fi19, F1: v1, F2: int8(v28)}
                                                                _ = v45
                                                            case 4294967295:
                                                                var v46 uint8 = 255
                                                                _ = v46
                                                            }
                                                            v47 := GenS4{F0: (v9 >> ((2) & 63)), F1: 4294967295}
                                                            _ = v44
                                                            _ = v47
                                                        }
                                                    case 0:
                                                        if ((v2 ^ (v10 * v2)) >= (v2 ^ (v2 / ((v10) | 1)))) {
                                                            if false {
                                                                v28 -= (v28 & (v28 % (((v28 + v28)) | 1)))
                                                                if (!(v2 < (v2 >> ((65535) & 63)))) {
                                                                    var v48 uint = 65535
                                                                    v49 := GenS2{F0: "", F1: 76}
                                                                    if (fi19 <= fi12) {
                                                                        if (128 == 194) {
                                                                            if false {
                                                                                switch "a" {
                                                                                case "foo bar":
                                                                                    v1 += (v1 % ((4294967295) | 1))
                                                                                default:
                                                                                    var v50 int32 = 733
                                                                                    _ = v50
                                                                                }
                                                                            } else {
                                                                                var v51 uint8 = 105
                                                                                var v52 uint16 = (v28 >> ((v48) & 63))
                                                                                _ = v51
                                                                                _ = v52
                                                                            }
                                                                            var v53 int8 = (v9 >> (((v48 * v48)) & 63))
                                                                            _ = v53
                                                                        }
                                                                        fmt.Println(v2, v9, v3)
                                                                        v54 := GenS8{F0: 65535, F1: ("" + ((("a" + (("" + "") + "\n\t")) + (("foo bar" + "") + ("\n\t" + "\n\t"))) + (((("hello" + ("hello" + "hello")) + "") + ((((("hello" + ("a" + "\n\t")) + "hello") + "foo bar") + (("hello" + "foo bar") + "foo bar")) + "a")) + "")))}
                                                                        if len(v11) > 0 {
                                                                            v11[0] = 1
                                                                        }
                                                                        v55 := GenS9{F0: ("" >= "\n\t")}
                                                                        _ = v54
                                                                        _ = v55
                                                                    } else {
                                                                        fmt.Println(v3)
                                                                    }
                                                                    _ = v48
                                                                    _ = v49
                                                                }
                                                                if (!(v1 >= (v1 + (^(v1))))) {
                                                                    switch (fi12 / ((8) | 1)) {
                                                                    case 127:
                                                                        _ = len(v4)
                                                                    case -929:
                                                                        var v56 bool = (false || true)
                                                                        _ = v56
                                                                    case -2147483648:
                                                                        v28 = (v28 & 576)
                                                                    }
                                                                    switch ((("hello" + "foo bar") + (((("hello" + "a") + "a") + ("foo bar" + "a")) + "\n\t")) + ("foo bar" + ((("foo bar" + ("hello" + "foo bar")) + "\n\t") + "hello"))) {
                                                                    case "hello":
                                                                        v57 := GenS3{F0: "\n\t", F1: false}
                                                                        _ = v57
                                                                    case "":
                                                                        var v58 uint32 = (v1 - 2147483648)
                                                                        _ = v58
                                                                    case "\n\t":
                                                                        v59 := GenS10{F0: 18446744073709551615, F1: 42, F2: (v3 << ((42) & 63))}
                                                                        _ = v59
                                                                    }
                                                                } else {
                                                                    if true {
                                                                        switch fi19 {
                                                                        case 9223372036854775807:
                                                                            var v60 float64 = 1e10
                                                                            switch fi12 {
                                                                            case 32767:
                                                                                switch fi12 {
                                                                                case 32767:
                                                                                    v28 = 2
                                                                                default:
                                                                                    var v61 float32 = (0.1 / ((1.5 - 0.0) / (1.5 / (1e10 / (1.5 - 3.14159)))))
                                                                                    _ = v61
                                                                                }
                                                                            default:
                                                                                if false {
                                                                                    var v62 uint64 = 2
                                                                                    v9 = v9
                                                                                    _ = v62
                                                                                }
                                                                            }
                                                                            _ = len(v4)
                                                                            _ = v60
                                                                        }
                                                                        v1 = v1
                                                                        var v63 float64 = (float64((((float32(1e10) * (1e10 / (float32(1.5) + (0.1 - ((0.0 - 0.1) / (3.14159 - 1.5)))))) / float32((-2.25 / 3.14159))) * -2.25)) + 1.5)
                                                                        if len(v11) >= 2 {
                                                                            v11 = v11[1:len(v11)]
                                                                        }
                                                                        switch (fi12 - fi19) {
                                                                        case -9223372036854775808:
                                                                            v64 := GenS11{F0: int8(v28)}
                                                                            var v65 int = fi12
                                                                            v9 += -1
                                                                            v66 := GenS12{F0: 486, F1: (-2.25 > -2.25), F2: 876}
                                                                            if (v28 > (^(v28))) {
                                                                                switch "a" {
                                                                                case "hello":
                                                                                    var v67 int16 = v2
                                                                                    _ = v67
                                                                                case "a":
                                                                                    var v68 bool = ((!false) && false)
                                                                                    _ = v68
                                                                                }
                                                                                v1 -= 2
                                                                                if ("" == "\n\t") {
                                                                                    fmt.Println(fi19)
                                                                                    v3 = v3
                                                                                }
                                                                            }
                                                                            _ = v64
                                                                            _ = v65
                                                                            _ = v66
                                                                        case 128:
                                                                            var v69 int16 = 0
                                                                            _ = v69
                                                                        }
                                                                        _ = v63
                                                                    } else {
                                                                        var v70 int = fi12
                                                                        var v71 uint32 = (v1 ^ v1)
                                                                        fmt.Println(v10, v3, v70)
                                                                        v71 = (v71 * (v71 % ((uint32(v28)) | 1)))
                                                                        if (false && true) {
                                                                            var v72 int = v70
                                                                            v1 = v71
                                                                            v73 := GenS8{F0: v71, F1: "a"}
                                                                            _ = v72
                                                                            _ = v73
                                                                        }
                                                                        _ = v70
                                                                        _ = v71
                                                                    }
                                                                    switch int(v28) {
                                                                    case -332:
                                                                        var v74 uint32 = (v1 >> ((4294967295) & 63))
                                                                        _ = v74
                                                                    case 127:
                                                                        if ((fi12 - fi19) > fi19) {
                                                                            v9 = (^(v9))
                                                                            var v75 uint32 = v1
                                                                            _ = v75
                                                                        } else {
                                                                            var v76 int64 = 9223372036854775807
                                                                            var v77 uint64 = 42
                                                                            v77 = v77
                                                                            _ = v76
                                                                            _ = v77
                                                                        }
                                                                    default:
                                                                        var v78 string = "foo bar"
                                                                        _ = v78
                                                                    }
                                                                }
                                                                if false {
                                                                    switch (fi12 / (((-(fi19))) | 1)) {
                                                                    case 127:
                                                                        var v79 uint16 = v28
                                                                        _ = v79
                                                                    case 42:
                                                                        if true {
                                                                            var v80 int = (fi12 & fi12)
                                                                            v2 = v10
                                                                            var v81 int = (v80 >> ((847) & 63))
                                                                            delete(v4, (v1 / (((^(v1))) | 1)))
                                                                            _ = v80
                                                                            _ = v81
                                                                        } else {
                                                                            v82 := GenS2{F0: "", F1: 48}
                                                                            _ = v82
                                                                        }
                                                                    case 9223372036854775807:
                                                                        if (v1 > 2) {
                                                                            var v83 int = int(v9)
                                                                            _ = len(v4)
                                                                            v1 -= (^(v1))
                                                                            v84 := GenS13{F0: 1, F1: 815}
                                                                            _ = v83
                                                                            _ = v84
                                                                        } else {
                                                                            v2 += v2
                                                                            var v85 int8 = (-(v9))
                                                                            switch "a" {
                                                                            case "":
                                                                                if true {
                                                                                    v4[4294967295] = (((v1 ^ (v1 % ((v1) | 1))) == 859) && (!(!(0 != -9223372036854775808))))
                                                                                    fmt.Println(v1, v28)
                                                                                    v86 := GenS7{F0: fi19, F1: v28}
                                                                                    v87 := GenS3{F0: "", F1: (18446744073709551615 > 4294967295)}
                                                                                    _ = v86
                                                                                    _ = v87
                                                                                }
                                                                            case "foo bar":
                                                                                fmt.Println(v1)
                                                                            default:
                                                                                var v88 float64 = 3.14159
                                                                                _ = v88
                                                                            }
                                                                            _ = v85
                                                                        }
                                                                    default:
                                                                        var v89 string = ("a" + "foo bar")
                                                                        _ = v89
                                                                    }
                                                                    if len(v11) > 0 {
                                                                        v11[0] = fi12
                                                                    }
                                                                } else {
                                                                    switch (^(fi19)) {
                                                                    case 768:
                                                                        var v90 float32 = float32(float64(1e10))
                                                                        _ = v90
                                                                    case 794:
                                                                        v91 := GenS14{F0: 2}
                                                                        _ = v91
                                                                    case 593:
                                                                        v3 = v3
                                                                    default:
                                                                        var v92 int = (fi12 * (fi19 | fi12))
                                                                        _ = v92
                                                                    }
                                                                }
                                                            }
                                                            v93 := GenS15{F0: (v9 - v9), F1: 1}
                                                            var v94 int32 = (v3 / (((v3 * (^(v3)))) | 1))
                                                            _ = v93
                                                            _ = v94
                                                        }
                                                    default:
                                                        var v95 int = 515
                                                        _ = v95
                                                    }
                                                    fmt.Println(v9, v2, v3)
                                                    _ = v21
                                                    _ = v28
                                                } else {
                                                    switch "" {
                                                    case "foo bar":
                                                        _ = len(v11)
                                                    }
                                                }
                                            case "a":
                                                var v96 bool = (!(v9 <= 127))
                                                _ = v96
                                            }
                                            v97 := GenS13{F0: 2, F1: 2}
                                            _ = v97
                                        }
                                    }
                                }
                                if false {
                                    var v98 uint16 = 786
                                    fmt.Println(fi12, v2)
                                    v98 += 42
                                    var v99 uint = 1
                                    _ = v98
                                    _ = v99
                                } else {
                                    v1 += v1
                                    var v100 uint = 0
                                    v101 := GenS8{F0: 2147483647, F1: "\n\t"}
                                    v9 -= v9
                                    switch ((("" + "\n\t") + "hello") + "") {
                                    case "":
                                        v1 = 0
                                    case "foo bar":
                                        v11 = append(v11, fi19)
                                    case "a":
                                        v3 = -347
                                    default:
                                        switch "hello" {
                                        case "":
                                            v100 = v100
                                        case "\n\t":
                                            v102 := GenS15{F0: v9, F1: 0}
                                            switch fi12 {
                                            case -746:
                                                v2 = (v10 - (v10 >> (((v100 & v100)) & 63)))
                                            case 42:
                                                fmt.Println(v1)
                                            case 856:
                                                var v103 int8 = v9
                                                _ = v103
                                            }
                                            var v104 int8 = v9
                                            _ = v102
                                            _ = v104
                                        case "a":
                                            if false {
                                                v100 = v100
                                                v1 -= v1
                                                switch (fi12 ^ (fi12 % (((fi12 / ((2147483647) | 1))) | 1))) {
                                                case 0:
                                                    fmt.Println(v10, v9)
                                                    v100 = v100
                                                    var v105 int8 = v9
                                                    var v106 uint8 = 1
                                                    v1 += 65535
                                                    var v107 uint8 = 2
                                                    var v108 uint32 = v1
                                                    _ = v105
                                                    _ = v106
                                                    _ = v107
                                                    _ = v108
                                                case 42:
                                                    var v109 int16 = v2
                                                    _ = v109
                                                case -1:
                                                    v110 := GenS10{F0: 712, F1: (v100 * v100), F2: v3}
                                                    _ = v110
                                                default:
                                                    fmt.Println(v100, v9)
                                                }
                                            } else {
                                                switch (-(fi19)) {
                                                case -234:
                                                    var v111 uint64 = 342
                                                    _ = v111
                                                case 2147483647:
                                                    if len(v11) > 0 {
                                                        v11[0] = fi12
                                                    }
                                                case 721:
                                                    switch ("a" + "hello") {
                                                    case "hello":
                                                        var v112 float32 = (0.1 / ((float32(3.14159) + 0.0) + (1e10 - ((-2.25 + (0.1 + 0.0)) + float32(1.5)))))
                                                        _ = v112
                                                    case "a":
                                                        var v113 uint32 = v1
                                                        _ = v113
                                                    default:
                                                        var v114 float32 = 0.1
                                                        _ = v114
                                                    }
                                                default:
                                                    switch "\n\t" {
                                                    case "hello":
                                                        var v115 uint64 = 2
                                                        _ = v115
                                                    case "foo bar":
                                                        var v116 float64 = ((3.14159 - 3.14159) / 0.1)
                                                        _ = v116
                                                    case "a":
                                                        v117 := GenS7{F0: -181, F1: 353}
                                                        _ = v117
                                                    }
                                                }
                                                fmt.Println(v2)
                                            }
                                        default:
                                            v118 := GenS15{F0: 1, F1: -103}
                                            _ = v118
                                        }
                                    }
                                    _ = v100
                                    _ = v101
                                }
                                v2 += (v10 & v10)
                            } else {
                                var v119 float64 = 1e10
                                if false {
                                    var v120 float64 = ((1e10 + -2.25) / 3.14159)
                                    var v121 float64 = v119
                                    v3 = (v3 ^ -1)
                                    _ = v120
                                    _ = v121
                                } else {
                                    var v122 uint32 = 1
                                    fmt.Println(fi19, v10)
                                    v9 = (v9 - v9)
                                    _ = v122
                                }
                                _ = v119
                            }
                            _ = fi19
                        }
                    default:
                        switch (fi12 ^ (-(fi12))) {
                        case -9223372036854775808:
                            fmt.Println(v9, v10)
                        case 32767:
                            var v123 uint = uint(v1)
                            _ = v123
                        }
                    }
                }
            }
        case "foo bar":
            v10 -= 42
        default:
            var v124 uint8 = 127
            _ = v124
        }
        _ = fi12
    }
    var v125 uint64 = 75
    if ((v3 * (v3 << ((42) & 63))) > v3) {
        v9 = (^(v9))
    } else {
        if true {
            v1 = (v1 << ((4294967295) & 63))
            for fi126 := 0; fi126 < 12; fi126++ {
                var v127 uint8 = 2
                var v128 int8 = (v9 | v9)
                var v129 uint = 530
                _ = fi126
                _ = v127
                _ = v128
                _ = v129
            }
        }
        v125 = (v125 * 2)
    }
    _ = len(v11)
    fmt.Println(v2, v10, v125)
    switch "" {
    case "":
        for fi130 := 0; fi130 < 2; fi130++ {
            fmt.Println(fi130, v10)
            if ((1.5 / 1.5) == ((0.0 * 1e10) * 0.1)) {
                v10 -= (v2 % ((int16(v1)) | 1))
                if false {
                    v3 -= (-(v3))
                } else {
                    for fi131 := 0; fi131 < 4; fi131++ {
                        var v132 uint8 = 193
                        fmt.Println(fi131, v2, v9)
                        v133 := GenS2{F0: "foo bar", F1: v132}
                        v134 := GenS6{F0: (v10 - v2), F1: "a"}
                        fmt.Println(v10, v9, v132)
                        _ = fi131
                        _ = v132
                        _ = v133
                        _ = v134
                    }
                }
            }
            var v135 float64 = (-2.25 + (0.1 * ((0.0 * 0.1) - 0.1)))
            _ = fi130
            _ = v135
        }
    default:
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
    }
    var v136 uint64 = v125
    v137 := GenS2{F0: "foo bar", F1: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v125:", v125)
    fmt.Println("v136:", v136)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v125
    _ = v136
    _ = v137
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
