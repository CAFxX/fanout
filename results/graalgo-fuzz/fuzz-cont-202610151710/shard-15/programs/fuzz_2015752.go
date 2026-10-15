package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 string }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int16; F1 string; F2 bool }
type GenS6 struct { F0 bool; F1 uint32; F2 uint8 }
type GenS7 struct { F0 int8; F1 uint32; F2 uint }
type GenS8 struct { F0 uint8; F1 int32; F2 int64 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 uint16; F1 uint8; F2 int32 }
type GenS12 struct { F0 int8; F1 uint32 }
type GenS13 struct { F0 int16; F1 int64 }
type GenS14 struct { F0 uint64; F1 uint8 }
type GenS15 struct { F0 uint16; F1 int64; F2 uint16 }
type GenS16 struct { F0 uint8 }
type GenS17 struct { F0 uint32; F1 int8 }
type GenS18 struct { F0 uint; F1 int64; F2 int8 }
type GenS19 struct { F0 int16; F1 uint64; F2 int16 }
type GenS20 struct { F0 uint32 }
type GenS21 struct { F0 uint64 }
type GenS22 struct { F0 uint8; F1 int32 }
type GenS23 struct { F0 int16; F1 int16 }

func main() {
    v1 := make(map[string]string)
    var v2 int = 32767
    var v3 int32 = -416
    v4 := []uint32{2, 42}
    var v5 int16 = 32767
    var v6 uint8 = 42
    v6 = v6
    fmt.Println(v6, v2)
    v6 -= v6
    v7 := GenS1{F0: v2}
    fmt.Println(v5)
    var v8 bool = (v6 > 222)
    if (!true) {
        for fi9 := 0; fi9 < 7; fi9++ {
            var v10 string = ("\n\t" + "foo bar")
            var v11 int = fi9
            _ = fi9
            _ = v10
            _ = v11
        }
    } else {
        var v12 int32 = (v3 >> ((4294967295) & 63))
        for fi13 := 0; fi13 < 14; fi13++ {
            for fi14 := 0; fi14 < 2; fi14++ {
                if v8 {
                    var v15 uint = 2
                    _ = v15
                } else {
                    switch ("" + ("" + ("a" + ("\n\t" + "")))) {
                    case "":
                        fmt.Println(fi14)
                    }
                    if v8 {
                        v6 -= v6
                        if false {
                            if (!(32767 < -2147483648)) {
                                var v16 uint64 = 276
                                var v17 uint8 = v6
                                _ = len(v1)
                                var v18 uint32 = 910
                                var v19 uint64 = (v16 << ((1) & 63))
                                var v20 int16 = 1
                                v4 = append(v4, v18)
                                switch (fi13 & (v2 ^ (^(fi13)))) {
                                case -9223372036854775808:
                                    v5 = v20
                                case -1:
                                    if false {
                                        v6 = v6
                                    }
                                case 9223372036854775807:
                                    v21 := GenS1{F0: fi13}
                                    if true {
                                        v6 = v6
                                        delete(v1, (("" + "") + ("\n\t" + ("\n\t" + ""))))
                                    }
                                    _ = v21
                                }
                                _ = v16
                                _ = v17
                                _ = v18
                                _ = v19
                                _ = v20
                            }
                            switch ("foo bar" + ("" + "hello")) {
                            case "a":
                                var v22 uint = 65535
                                _ = v22
                            case "":
                                if v8 {
                                    delete(v1, "foo bar")
                                    _ = len(v1)
                                    if false {
                                        if (255 > 2) {
                                            v23 := GenS1{F0: fi13}
                                            fmt.Println(v6, fi14)
                                            var v24 uint64 = 2
                                            _ = v23
                                            _ = v24
                                        }
                                    }
                                } else {
                                    var v25 bool = v8
                                    _ = v25
                                }
                            default:
                                v26 := GenS1{F0: 32767}
                                _ = v26
                            }
                            switch (^(fi14)) {
                            case 128:
                                v27 := GenS2{F0: "foo bar", F1: "hello"}
                                _ = v27
                            case 297:
                                if (0.1 < ((0.0 - 1.5) * 1e10)) {
                                    if len(v4) > 0 {
                                        v4[0] = 4294967295
                                    }
                                    v3 = v12
                                } else {
                                    if false {
                                        v6 -= 42
                                        var v28 string = "a"
                                        _ = v28
                                    }
                                    if ((579 < 2) || v8) {
                                        var v29 int32 = v12
                                        switch "hello" {
                                        case "":
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        default:
                                            _ = len(v1)
                                        }
                                        if true {
                                            v5 += int16(v6)
                                        } else {
                                            v3 -= int32(v5)
                                            if v8 {
                                                var v30 uint = 965
                                                v31 := v1[("a" + ("\n\t" + "foo bar"))]
                                                var v32 uint64 = 1
                                                _ = v30
                                                _ = v31
                                                _ = v32
                                            }
                                            if (!v8) {
                                                if ((v12 / ((v3) | 1)) < (v12 / ((v29) | 1))) {
                                                    if (true || ((v6 | (v6 / ((v6) | 1))) == v6)) {
                                                        if false {
                                                            v1["\n\t"] = (("foo bar" + "foo bar") + "")
                                                            v33 := GenS3{F0: 575}
                                                            _ = v33
                                                        }
                                                    }
                                                    if len(v4) > 0 {
                                                        v4[0] = 0
                                                    }
                                                    if false {
                                                        if (!v8) {
                                                            var v34 uint32 = 2
                                                            _ = len(v1)
                                                            switch "\n\t" {
                                                            case "a":
                                                                v2 = (-(v2))
                                                            case "foo bar":
                                                                var v35 int = v2
                                                                _ = v35
                                                            case "":
                                                                v2 += fi14
                                                            }
                                                            _ = len(v4)
                                                            _ = v34
                                                        } else {
                                                            if false {
                                                                var v36 uint32 = 65535
                                                                switch ("" + (("foo bar" + "") + "foo bar")) {
                                                                case "\n\t":
                                                                    switch 128 {
                                                                    case -1:
                                                                        v5 = v5
                                                                        var v37 int32 = -1
                                                                        delete(v1, "\n\t")
                                                                        _ = v37
                                                                    }
                                                                }
                                                                var v38 bool = (!(0 <= 819))
                                                                v2 += fi14
                                                                _ = v36
                                                                _ = v38
                                                            }
                                                            fmt.Println(v8, v6)
                                                            fmt.Println(fi14)
                                                        }
                                                        if (fi13 < (^(v2))) {
                                                            fmt.Println(fi14, v5, v5)
                                                            delete(v1, "a")
                                                            if false {
                                                                if false {
                                                                    if len(v4) > 0 {
                                                                        v4[0] = 411
                                                                    }
                                                                    var v39 uint8 = v6
                                                                    _ = v39
                                                                } else {
                                                                    switch fi13 {
                                                                    case -2147483648:
                                                                        var v40 bool = (((("a" + "\n\t") + "foo bar") + "hello") == "")
                                                                        var v41 int64 = 0
                                                                        if (!(v40 || (!(((3.14159 * 3.14159) * 3.14159) <= 0.0)))) {
                                                                            fmt.Println(fi13)
                                                                            v29 = v3
                                                                            var v42 uint64 = 18446744073709551615
                                                                            _ = v42
                                                                        }
                                                                        _ = v40
                                                                        _ = v41
                                                                    default:
                                                                        if true {
                                                                            if (int64(v6) >= int64(v6)) {
                                                                                v43 := GenS1{F0: (fi14 % ((859) | 1))}
                                                                                _ = v43
                                                                            } else {
                                                                                v44 := GenS1{F0: -1}
                                                                                v8 = (!false)
                                                                                var v45 int64 = -9223372036854775808
                                                                                _ = v44
                                                                                _ = v45
                                                                            }
                                                                            v1["foo bar"] = "a"
                                                                            var v46 string = ("\n\t" + (((("" + (("hello" + ("hello" + "")) + ("a" + ""))) + "foo bar") + "foo bar") + "hello"))
                                                                            _ = v46
                                                                        } else {
                                                                            fmt.Println(v6, v5, v29)
                                                                            if v8 {
                                                                                v1[("" + "\n\t")] = ""
                                                                            } else {
                                                                                var v47 int32 = v3
                                                                                _ = len(v1)
                                                                                delete(v1, ("foo bar" + (("hello" + "\n\t") + "\n\t")))
                                                                                _ = v47
                                                                            }
                                                                        }
                                                                    }
                                                                    v48 := GenS2{F0: "hello", F1: "a"}
                                                                    v49 := v1[(((("" + "foo bar") + "hello") + "a") + "hello")]
                                                                    v50 := GenS3{F0: 0}
                                                                    _ = v48
                                                                    _ = v49
                                                                    _ = v50
                                                                }
                                                                v8 = (!((!v8) && (127 <= int8(v5))))
                                                                var v51 float64 = 1e10
                                                                _ = v51
                                                            }
                                                            var v52 int16 = -32768
                                                            _ = v52
                                                        }
                                                        switch "foo bar" {
                                                        case "hello":
                                                            if ("" < ("" + ("a" + ((("a" + ("hello" + (("foo bar" + ("" + ("\n\t" + ("" + "\n\t")))) + ("hello" + ("hello" + "foo bar"))))) + ("\n\t" + ("" + "a"))) + "hello")))) {
                                                                switch "\n\t" {
                                                                case "foo bar":
                                                                    v1["\n\t"] = (("hello" + "foo bar") + "hello")
                                                                }
                                                            } else {
                                                                switch (fi14 / ((fi13) | 1)) {
                                                                case 297:
                                                                    if true {
                                                                        fmt.Println(fi13, v8, fi14)
                                                                        if ((v3 % (((v29 << ((2) & 63))) | 1)) != 427) {
                                                                            var v53 int32 = -1
                                                                            if (!true) {
                                                                                if true {
                                                                                    v6 = 255
                                                                                    v29 = (v53 ^ v12)
                                                                                    v12 = v29
                                                                                } else {
                                                                                    fmt.Println(v6)
                                                                                    var v54 int32 = (^(v12))
                                                                                    _ = v54
                                                                                }
                                                                            }
                                                                            if ((uint32(v6) >= 2) && (127 != v6)) {
                                                                                switch "hello" {
                                                                                case "\n\t":
                                                                                    v8 = false
                                                                                default:
                                                                                    var v55 uint8 = (v6 | v6)
                                                                                    _ = v55
                                                                                }
                                                                                switch (fi14 | (v2 + (fi14 << ((927) & 63)))) {
                                                                                case -908:
                                                                                    v56 := GenS1{F0: (fi14 | 128)}
                                                                                    _ = v56
                                                                                case 987:
                                                                                    fmt.Println(v12)
                                                                                }
                                                                            } else {
                                                                                v29 += (v53 << ((691) & 63))
                                                                            }
                                                                            fmt.Println(v8)
                                                                            if (255 <= 1) {
                                                                                v3 -= (v29 & (^(v12)))
                                                                                if v8 {
                                                                                    var v57 uint8 = v6
                                                                                    _ = v57
                                                                                }
                                                                                switch (fi14 & fi13) {
                                                                                case 2147483647:
                                                                                    var v58 uint8 = v6
                                                                                    _ = v58
                                                                                default:
                                                                                    v59 := v1["a"]
                                                                                    _ = v59
                                                                                }
                                                                            } else {
                                                                                fmt.Println(v29, v12, v8)
                                                                            }
                                                                            v60 := GenS4{F0: 0}
                                                                            if false {
                                                                                var v61 uint16 = 1
                                                                                if false {
                                                                                    v62 := GenS5{F0: (v5 * v5), F1: "a", F2: v8}
                                                                                    _ = v62
                                                                                }
                                                                                _ = v61
                                                                            } else {
                                                                                var v63 int8 = 42
                                                                                var v64 float32 = float32(0.0)
                                                                                v65 := GenS6{F0: v8, F1: 42, F2: (^(v6))}
                                                                                _ = v63
                                                                                _ = v64
                                                                                _ = v65
                                                                            }
                                                                            _ = v53
                                                                            _ = v60
                                                                        } else {
                                                                            var v66 uint16 = 525
                                                                            _ = v66
                                                                        }
                                                                        v12 += v29
                                                                    } else {
                                                                        var v67 int = (v2 | v2)
                                                                        fmt.Println(fi13, v3)
                                                                        if v8 {
                                                                            v68 := GenS7{F0: 0, F1: 219, F2: 65535}
                                                                            _ = v68
                                                                        } else {
                                                                            var v69 int8 = 0
                                                                            if len(v4) >= 2 {
                                                                                v4 = v4[1:len(v4)]
                                                                            }
                                                                            v70 := GenS1{F0: fi13}
                                                                            _ = v69
                                                                            _ = v70
                                                                        }
                                                                        _ = v67
                                                                    }
                                                                case 127:
                                                                    if false {
                                                                        v71 := GenS8{F0: (^(v6)), F1: -1, F2: -9223372036854775808}
                                                                        _ = v71
                                                                    } else {
                                                                        var v72 int = (^(fi14))
                                                                        if (!v8) {
                                                                            switch ((("\n\t" + "\n\t") + "") + "\n\t") {
                                                                            case "hello":
                                                                                fmt.Println(v5, v72, v2)
                                                                            case "foo bar":
                                                                                switch fi14 {
                                                                                case 9223372036854775807:
                                                                                    var v73 uint64 = 1
                                                                                    _ = v73
                                                                                default:
                                                                                    var v74 int16 = v5
                                                                                    _ = v74
                                                                                }
                                                                            case "a":
                                                                                v3 = -2147483648
                                                                            }
                                                                        } else {
                                                                            if (42 > v6) {
                                                                                var v75 int8 = 66
                                                                                v5 = v5
                                                                                if v8 {
                                                                                    v75 -= v75
                                                                                    var v76 int32 = v12
                                                                                    var v77 int64 = -488
                                                                                    if (!true) {
                                                                                        var v78 uint = 399
                                                                                        if (!false) {
                                                                                            var v79 int8 = (v75 ^ (v75 & 42))
                                                                                            var v80 int = 32767
                                                                                            v81 := v1["\n\t"]
                                                                                            _ = v79
                                                                                            _ = v80
                                                                                            _ = v81
                                                                                        }
                                                                                        v82 := GenS9{F0: v75}
                                                                                        _ = v78
                                                                                        _ = v82
                                                                                    } else {
                                                                                        if false {
                                                                                            v3 += (v3 + (v29 * (v76 << ((2) & 63))))
                                                                                        }
                                                                                        v72 = -2147483648
                                                                                    }
                                                                                    switch fi14 {
                                                                                    case 32768:
                                                                                        v83 := GenS10{F0: 581}
                                                                                        _ = len(v4)
                                                                                        fmt.Println(v3)
                                                                                        _ = v83
                                                                                    default:
                                                                                        var v84 bool = ((v75 / ((-128) | 1)) < v75)
                                                                                        _ = v84
                                                                                    }
                                                                                    _ = v76
                                                                                    _ = v77
                                                                                }
                                                                                _ = v75
                                                                            } else {
                                                                                if len(v4) >= 2 {
                                                                                    v4 = v4[1:len(v4)]
                                                                                }
                                                                                fmt.Println(v5)
                                                                            }
                                                                            var v85 float32 = 1e10
                                                                            if true {
                                                                                var v86 int = 2147483647
                                                                                fmt.Println(v3, v29, v6)
                                                                                if false {
                                                                                    if (!v8) {
                                                                                        v87 := GenS1{F0: v2}
                                                                                        fmt.Println(fi13)
                                                                                        _ = v87
                                                                                    } else {
                                                                                        var v88 uint = 783
                                                                                        var v89 int8 = 0
                                                                                        var v90 int8 = 2
                                                                                        v2 -= (v72 - 128)
                                                                                        v91 := GenS3{F0: -406}
                                                                                        if len(v4) >= 2 {
                                                                                            v4 = v4[1:len(v4)]
                                                                                        }
                                                                                        _ = v88
                                                                                        _ = v89
                                                                                        _ = v90
                                                                                        _ = v91
                                                                                    }
                                                                                    v12 += (v3 / (((v3 * v12)) | 1))
                                                                                }
                                                                                _ = v86
                                                                            } else {
                                                                                v12 -= -2147483648
                                                                            }
                                                                            _ = v85
                                                                        }
                                                                        v92 := GenS3{F0: -9223372036854775808}
                                                                        _ = v72
                                                                        _ = v92
                                                                    }
                                                                default:
                                                                    v5 -= 1000
                                                                }
                                                            }
                                                        default:
                                                            if (true && v8) {
                                                                v93 := GenS1{F0: fi14}
                                                                var v94 uint64 = 576
                                                                v95 := GenS11{F0: 76, F1: v6, F2: -785}
                                                                _ = v93
                                                                _ = v94
                                                                _ = v95
                                                            }
                                                        }
                                                        if (0 == 18446744073709551615) {
                                                            v1[(("hello" + (("\n\t" + "") + "foo bar")) + "")] = "foo bar"
                                                            var v96 float32 = float32(float64(1.5))
                                                            var v97 uint = 331
                                                            _ = v96
                                                            _ = v97
                                                        }
                                                    }
                                                }
                                                var v98 int = fi13
                                                _ = v98
                                            } else {
                                                switch "" {
                                                case "a":
                                                    switch fi14 {
                                                    case -1:
                                                        v8 = v8
                                                    default:
                                                        v99 := GenS12{F0: 42, F1: 65535}
                                                        switch "\n\t" {
                                                        case "":
                                                            var v100 int32 = (v3 % ((189) | 1))
                                                            _ = v100
                                                        }
                                                        if v8 {
                                                            var v101 uint32 = 422
                                                            v1["foo bar"] = (("" + "a") + "hello")
                                                            _ = v101
                                                        } else {
                                                            var v102 int32 = v29
                                                            var v103 uint16 = 65535
                                                            _ = v102
                                                            _ = v103
                                                        }
                                                        if v8 {
                                                            v3 -= (-(v12))
                                                            v5 -= v5
                                                            var v104 uint16 = 1
                                                            _ = v104
                                                        } else {
                                                            var v105 int32 = (v29 << ((0) & 63))
                                                            var v106 int16 = 1000
                                                            v106 += v106
                                                            if ((!(!((true && (2 != v6)) || false))) && v8) {
                                                                _ = len(v4)
                                                                fmt.Println(v105, v105)
                                                            } else {
                                                                var v107 int16 = 32767
                                                                _ = v107
                                                            }
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                            _ = v105
                                                            _ = v106
                                                        }
                                                        v8 = true
                                                        _ = v99
                                                    }
                                                case "hello":
                                                    var v108 bool = true
                                                    _ = v108
                                                default:
                                                    var v109 uint = 652
                                                    var v110 uint64 = 18446744073709551615
                                                    v2 = fi14
                                                    _ = v109
                                                    _ = v110
                                                }
                                            }
                                        }
                                        _ = v29
                                    }
                                    if v8 {
                                        switch (("foo bar" + ("a" + "\n\t")) + "hello") {
                                        case "":
                                            _ = len(v1)
                                        default:
                                            var v111 float32 = 3.14159
                                            _ = v111
                                        }
                                        if (int8(v6) <= -128) {
                                            if (true && ((18446744073709551615 >= uint64(v6)) && ((^(v12)) != (v12 << ((378) & 63))))) {
                                                v2 += (^(v2))
                                                v5 = int16(v6)
                                            } else {
                                                v112 := v1["a"]
                                                v113 := GenS9{F0: 1}
                                                _ = v112
                                                _ = v113
                                            }
                                            switch "" {
                                            case "\n\t":
                                                delete(v1, "")
                                            case "":
                                                v8 = v8
                                            case "foo bar":
                                                if v8 {
                                                    var v114 int16 = (v5 | 1000)
                                                    var v115 int = (v2 & 32768)
                                                    var v116 float32 = 0.1
                                                    _ = v114
                                                    _ = v115
                                                    _ = v116
                                                }
                                                v4 = append(v4, 1)
                                            default:
                                                v117 := GenS13{F0: -451, F1: -282}
                                                _ = v117
                                            }
                                            fmt.Println(v2)
                                        }
                                        v118 := GenS9{F0: -8}
                                        _ = v118
                                    } else {
                                        var v119 uint16 = 1
                                        if (!("" != "a")) {
                                            var v120 int8 = -67
                                            var v121 uint8 = (v6 << ((18446744073709551615) & 63))
                                            _ = v120
                                            _ = v121
                                        } else {
                                            v122 := GenS14{F0: uint64(v6), F1: v6}
                                            if v8 {
                                                if false {
                                                    v6 -= v6
                                                } else {
                                                    fmt.Println(fi14, v6, v8)
                                                }
                                                delete(v1, "foo bar")
                                                _ = len(v4)
                                            }
                                            _ = v122
                                        }
                                        switch -683 {
                                        case 9223372036854775807:
                                            v2 = (^(fi13))
                                        case 4294967295:
                                            switch "hello" {
                                            case "":
                                                v123 := v1["hello"]
                                                _ = v123
                                            case "hello":
                                                if (!v8) {
                                                    switch ("hello" + (("\n\t" + "\n\t") + "")) {
                                                    case "":
                                                        if true {
                                                            var v124 int32 = (v3 - v3)
                                                            var v125 int16 = (-(v5))
                                                            _ = v124
                                                            _ = v125
                                                        } else {
                                                            if (true || (!((v3 ^ v12) >= (v3 - (v3 / ((v3) | 1)))))) {
                                                                _ = len(v4)
                                                                if (0 > 1) {
                                                                    var v126 int16 = -671
                                                                    v127 := GenS14{F0: 4294967295, F1: (v6 * (^(v6)))}
                                                                    _ = v126
                                                                    _ = v127
                                                                } else {
                                                                    v3 += (v12 & v12)
                                                                    if true {
                                                                        v12 = v12
                                                                    }
                                                                    v8 = (v8 || false)
                                                                }
                                                            }
                                                            switch 9223372036854775807 {
                                                            case -295:
                                                                var v128 uint64 = 16
                                                                var v129 bool = (42 >= v6)
                                                                var v130 uint = 42
                                                                v131 := GenS12{F0: -1, F1: 2147483647}
                                                                fmt.Println(fi13)
                                                                _ = v128
                                                                _ = v129
                                                                _ = v130
                                                                _ = v131
                                                            case -9223372036854775808:
                                                                switch (fi14 ^ 32767) {
                                                                case 1:
                                                                    if len(v4) > 0 {
                                                                        v4[0] = 2
                                                                    }
                                                                case 128:
                                                                    v132 := GenS14{F0: 18446744073709551615, F1: (v6 / ((v6) | 1))}
                                                                    _ = v132
                                                                case -9223372036854775808:
                                                                    fmt.Println(v6, fi13)
                                                                }
                                                            case 248:
                                                                var v133 float32 = 0.0
                                                                v134 := GenS15{F0: 2, F1: 1, F2: (v119 / ((0) | 1))}
                                                                if true {
                                                                    if len(v4) > 0 {
                                                                        v4[0] = 1
                                                                    }
                                                                    if len(v4) > 0 {
                                                                        v4[0] = 65535
                                                                    }
                                                                    switch (("\n\t" + "hello") + ("" + "foo bar")) {
                                                                    case "foo bar":
                                                                        var v135 int64 = -9223372036854775808
                                                                        _ = v135
                                                                    case "a":
                                                                        v2 -= fi13
                                                                    default:
                                                                        fmt.Println(v12, fi13)
                                                                    }
                                                                    switch ("\n\t" + "a") {
                                                                    case "\n\t":
                                                                        var v136 string = "hello"
                                                                        v137 := GenS16{F0: v6}
                                                                        v138 := GenS7{F0: 1, F1: 65535, F2: 1}
                                                                        v139 := GenS17{F0: 96, F1: 1}
                                                                        _ = v136
                                                                        _ = v137
                                                                        _ = v138
                                                                        _ = v139
                                                                    case "foo bar":
                                                                        var v140 int32 = (^(v3))
                                                                        _ = v140
                                                                    default:
                                                                        var v141 uint32 = 2147483647
                                                                        _ = v141
                                                                    }
                                                                    fmt.Println(v6)
                                                                    if false {
                                                                        if (!(v8 && false)) {
                                                                            var v142 string = "a"
                                                                            v143 := GenS18{F0: uint(v119), F1: 15, F2: int8(v5)}
                                                                            _ = v142
                                                                            _ = v143
                                                                        }
                                                                    }
                                                                    v144 := GenS19{F0: (^(v5)), F1: 178, F2: v5}
                                                                    _ = v144
                                                                }
                                                                _ = v133
                                                                _ = v134
                                                            }
                                                        }
                                                    case "\n\t":
                                                        switch ("" + "a") {
                                                        case "hello":
                                                            v145 := v1["a"]
                                                            _ = v145
                                                        case "\n\t":
                                                            var v146 uint8 = (v6 & v6)
                                                            _ = v146
                                                        case "":
                                                            fmt.Println(v119, v3, v6)
                                                        }
                                                    default:
                                                        v4 = append(v4, 65535)
                                                    }
                                                    if true {
                                                        if len(v4) >= 2 {
                                                            v4 = v4[1:len(v4)]
                                                        }
                                                        var v147 int32 = int32(fi13)
                                                        _ = v147
                                                    } else {
                                                        v12 = (v3 - 0)
                                                    }
                                                    var v148 int16 = (v5 + v5)
                                                    var v149 float32 = -2.25
                                                    _ = v148
                                                    _ = v149
                                                }
                                            case "\n\t":
                                                if (("foo bar" + "a") == "hello") {
                                                    if (42 != -89) {
                                                        v150 := GenS20{F0: 4294967295}
                                                        if ((v5 - v5) == v5) {
                                                            switch "a" {
                                                            case "\n\t":
                                                                v151 := v1["foo bar"]
                                                                _ = v151
                                                            case "hello":
                                                                fmt.Println(v6)
                                                            }
                                                            fmt.Println(v2)
                                                            v2 -= v2
                                                        }
                                                        v119 += v119
                                                        _ = v150
                                                    }
                                                    v1["\n\t"] = ("" + "\n\t")
                                                }
                                            default:
                                                v2 = fi14
                                            }
                                        default:
                                            v152 := GenS21{F0: 2}
                                            v2 = -2147483648
                                            var v153 int = (-(fi13))
                                            _ = v152
                                            _ = v153
                                        }
                                        _ = v119
                                    }
                                }
                            case 0:
                                var v154 uint64 = 4294967295
                                _ = v154
                            default:
                                switch "a" {
                                case "a":
                                    v12 -= v3
                                default:
                                    switch "\n\t" {
                                    case "a":
                                        v8 = ((!v8) && v8)
                                    case "hello":
                                        v4 = append(v4, 4294967295)
                                    case "":
                                        if (2 != 2147483647) {
                                            v155 := GenS19{F0: -1, F1: 2, F2: (v5 << ((1) & 63))}
                                            _ = v155
                                        } else {
                                            v5 = (v5 & 444)
                                            var v156 uint = 2
                                            _ = v156
                                        }
                                    default:
                                        v157 := GenS22{F0: v6, F1: (v3 | (v3 | 151))}
                                        _ = len(v4)
                                        _ = v157
                                    }
                                }
                            }
                        }
                    }
                }
                v8 = v8
                var v158 uint8 = 127
                _ = fi14
                _ = v158
            }
            v6 = 91
            var v159 bool = v8
            _ = fi13
            _ = v159
        }
        _ = v12
    }
    switch int(v3) {
    case 128:
        switch ("" + "foo bar") {
        case "\n\t":
            switch "" {
            case "foo bar":
                if len(v4) > 0 {
                    v4[0] = 4294967295
                }
            default:
                v160 := GenS23{F0: (v5 % (((v5 | 1000)) | 1)), F1: v5}
                _ = v160
            }
        }
    }
    fmt.Println(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
