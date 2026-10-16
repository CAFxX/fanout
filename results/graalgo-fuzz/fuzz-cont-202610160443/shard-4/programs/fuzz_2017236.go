package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int32; F1 int16 }
type GenS5 struct { F0 uint64; F1 string; F2 int }
type GenS6 struct { F0 int; F1 uint16; F2 uint64 }
type GenS7 struct { F0 int; F1 string; F2 bool }
type GenS8 struct { F0 uint16; F1 string; F2 uint8 }
type GenS9 struct { F0 int64; F1 int }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 bool; F1 int16; F2 int }
type GenS12 struct { F0 int }
type GenS13 struct { F0 int; F1 uint8 }
type GenS14 struct { F0 string; F1 int16 }
type GenS15 struct { F0 bool }
type GenS16 struct { F0 int16 }

func main() {
    var v1 string = ("foo bar" + "a")
    v2 := make(map[int]uint32)
    var v3 float32 = 1e10
    var v4 string = v1
    v5 := []string{(v1 + (v1 + "a")), "\n\t"}
    v6 := GenS1{F0: -280, F1: v4}
    v7 := GenS2{F0: (("hello" + v4) + "foo bar"), F1: 919}
    var v8 uint = 820
    v4 = "foo bar"
    if len(v5) > 0 {
        v5[0] = "\n\t"
    }
    delete(v2, 816)
    for fi9 := 0; fi9 < 20; fi9++ {
        if (0.0 >= 0.0) {
            if false {
                if (-1 < 42) {
                    v1 = ((((v4 + v1) + v4) + "a") + "foo bar")
                } else {
                    for fi10 := 0; fi10 < 15; fi10++ {
                        v3 = (1e10 - v3)
                        if true {
                            v11 := GenS2{F0: ("foo bar" + ("foo bar" + v4)), F1: 378}
                            var v12 uint = v8
                            fmt.Println(fi9, v8)
                            _ = v11
                            _ = v12
                        }
                        _ = fi10
                    }
                    if (0.0 != float32((float64(0.1) * -2.25))) {
                        for fi13 := 0; fi13 < 6; fi13++ {
                            if false {
                                var v14 string = "foo bar"
                                v15 := GenS2{F0: v4, F1: 1}
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                _ = v14
                                _ = v15
                            } else {
                                v16 := GenS2{F0: v4, F1: 826}
                                if false {
                                    if ((v3 * (v3 * 1.5)) >= 3.14159) {
                                        v17 := GenS1{F0: -1, F1: "foo bar"}
                                        v8 = 4294967295
                                        v18 := GenS1{F0: (fi13 % ((-1) | 1)), F1: v4}
                                        _ = v17
                                        _ = v18
                                    }
                                } else {
                                    v1 = v1
                                    fmt.Println(v3)
                                }
                                if true {
                                    v1 = (("" + ("\n\t" + ("" + ("foo bar" + ((v4 + v4) + ""))))) + "\n\t")
                                }
                                _ = v16
                            }
                            var v19 uint64 = uint64(v8)
                            if false {
                                _ = len(v5)
                                v20 := GenS2{F0: "a", F1: 2147483647}
                                var v21 string = "\n\t"
                                _ = v20
                                _ = v21
                            } else {
                                if (954 <= 0) {
                                    fmt.Println(v8)
                                    v4 = v1
                                } else {
                                    v19 += v19
                                    fmt.Println(v8, v19)
                                    fmt.Println(fi13, v4)
                                }
                                fmt.Println(v19)
                            }
                            _ = fi13
                            _ = v19
                        }
                        if (true || false) {
                            var v22 int8 = -73
                            if ((v8 << ((v8) & 63)) >= (^(v8))) {
                                for fi23 := 0; fi23 < 10; fi23++ {
                                    if (9223372036854775807 < 0) {
                                        if true {
                                            var v24 string = v4
                                            if false {
                                                fmt.Println(fi9, v24)
                                            } else {
                                                switch ("a" + "foo bar") {
                                                case "foo bar":
                                                    v4 = v1
                                                    v22 = (v22 << (((v8 >> ((v8) & 63))) & 63))
                                                    fmt.Println(v3, v8)
                                                case "":
                                                    v24 = v24
                                                case "a":
                                                    var v25 float64 = 1.5
                                                    _ = v25
                                                }
                                                v26 := v2[fi9]
                                                _ = v26
                                            }
                                            if false {
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                                v27 := GenS3{F0: -32768}
                                                v28 := GenS4{F0: -2147483648, F1: 579}
                                                _ = v27
                                                _ = v28
                                            }
                                            var v29 uint64 = 458
                                            _ = v24
                                            _ = v29
                                        }
                                        switch 351 {
                                        case 4294967295:
                                            v4 = v1
                                        case 32768:
                                            var v30 int64 = -559
                                            _ = v30
                                        case -134:
                                            v1 = v1
                                        default:
                                            fmt.Println(v8)
                                        }
                                    } else {
                                        v22 += 42
                                        if len(v5) > 0 {
                                            v5[0] = v4
                                        }
                                        var v31 bool = (!(float64(float32((0.1 / 0.1))) > (float64(v3) - 1e10)))
                                        _ = v31
                                    }
                                    _ = fi23
                                }
                                v32 := GenS3{F0: 1000}
                                v33 := GenS5{F0: 1, F1: ("\n\t" + ("a" + v4)), F2: fi9}
                                _ = v32
                                _ = v33
                            }
                            _ = v22
                        }
                    }
                    var v34 string = "a"
                    _ = v34
                }
            }
            var v35 uint32 = 601
            _ = v35
        }
        for fi36 := 0; fi36 < 10; fi36++ {
            switch (fi36 % ((2147483647) | 1)) {
            case 32768:
                if (-658 == 2147483647) {
                    fmt.Println(v1)
                } else {
                    v37 := GenS4{F0: -204, F1: 1}
                    _ = v37
                }
            case 42:
                v38 := GenS2{F0: v1, F1: 386}
                _ = v38
            case 735:
                var v39 uint32 = 19
                _ = v39
            default:
                fmt.Println(fi9)
            }
            if true {
                delete(v2, (fi36 % ((fi36) | 1)))
                fmt.Println(v1, v1)
                switch 991 {
                case 0:
                    switch v1 {
                    case "a":
                        v3 -= v3
                    case "hello":
                        v40 := GenS2{F0: v1, F1: 65535}
                        _ = v40
                    case "\n\t":
                        switch fi36 {
                        case -1:
                            var v41 uint32 = 1
                            _ = v41
                        }
                    default:
                        switch "" {
                        case "foo bar":
                            if (!(0 < 42)) {
                                v8 = v8
                                switch (^(fi9)) {
                                case -9223372036854775808:
                                    if (42 > 553) {
                                        switch (v1 + v4) {
                                        case "":
                                            switch v4 {
                                            case "hello":
                                                if true {
                                                    v42 := GenS6{F0: (fi9 | fi9), F1: 1, F2: 982}
                                                    v4 = (v4 + "a")
                                                    _ = v42
                                                }
                                            case "foo bar":
                                                fmt.Println(v4, fi36, v3)
                                            case "":
                                                var v43 int32 = 203
                                                var v44 uint32 = 0
                                                v45 := GenS7{F0: int(v44), F1: "a", F2: true}
                                                _ = v43
                                                _ = v44
                                                _ = v45
                                            }
                                        case "hello":
                                            switch fi9 {
                                            case -364:
                                                var v46 bool = false
                                                _ = v46
                                            case 9223372036854775807:
                                                v3 += v3
                                                _ = len(v2)
                                                var v47 int16 = 42
                                                _ = v47
                                            }
                                        default:
                                            var v48 uint = v8
                                            _ = v48
                                        }
                                        fmt.Println(v3)
                                        if (0 < 696) {
                                            var v49 int32 = -840
                                            var v50 int64 = 9223372036854775807
                                            _ = v49
                                            _ = v50
                                        }
                                    }
                                case -2147483648:
                                    v1 = "a"
                                default:
                                    v2[fi9] = 886
                                }
                                v8 += v8
                            } else {
                                _ = len(v5)
                                v8 = (v8 >> ((42) & 63))
                            }
                        case "":
                            if ((v3 - (float32((3.14159 + 1e10)) + v3)) <= -2.25) {
                                var v51 uint64 = 841
                                v52 := GenS3{F0: 1000}
                                var v53 uint64 = (v51 / (((v51 & v51)) | 1))
                                _ = v51
                                _ = v52
                                _ = v53
                            }
                        default:
                            if true {
                                switch 128 {
                                case -701:
                                    fmt.Println(v8, v3, v3)
                                }
                                v54 := GenS6{F0: fi9, F1: uint16(v8), F2: uint64(v8)}
                                var v55 uint = (v8 & 65535)
                                _ = v54
                                _ = v55
                            } else {
                                switch "a" {
                                case "\n\t":
                                    var v56 int8 = 49
                                    _ = v56
                                default:
                                    if true {
                                        var v57 uint16 = 255
                                        if true {
                                            v1 = ""
                                            if false {
                                                if false {
                                                    switch v4 {
                                                    case "":
                                                        v1 = "hello"
                                                    }
                                                    if true {
                                                        v58 := GenS1{F0: 0, F1: "hello"}
                                                        _ = v58
                                                    } else {
                                                        var v59 float32 = ((0.1 * v3) + v3)
                                                        v8 = (v8 % ((v8) | 1))
                                                        var v60 float32 = v3
                                                        fmt.Println(v57, v60, v8)
                                                        _ = v59
                                                        _ = v60
                                                    }
                                                } else {
                                                    var v61 uint8 = 2
                                                    v57 += v57
                                                    _ = v61
                                                }
                                                if (-128 < 127) {
                                                    if (787 <= 568) {
                                                        var v62 uint64 = 346
                                                        if ((-64 > -128) || true) {
                                                            switch v4 {
                                                            case "":
                                                                var v63 int = (fi9 - fi36)
                                                                _ = v63
                                                            }
                                                            if len(v5) >= 2 {
                                                                v5 = v5[1:len(v5)]
                                                            }
                                                            v8 += 377
                                                        } else {
                                                            switch (fi9 - 32767) {
                                                            case 42:
                                                                fmt.Println(v4, fi36)
                                                            case 1:
                                                                var v64 int8 = -81
                                                                _ = v64
                                                            default:
                                                                fmt.Println(v62, v57, v3)
                                                            }
                                                            fmt.Println(fi9, fi36, v62)
                                                            switch "foo bar" {
                                                            case "":
                                                                if true {
                                                                    fmt.Println(v4, v3, fi36)
                                                                    var v65 float64 = float64(float32((0.0 / (0.0 * (1e10 - (-2.25 + (-2.25 * float64(v3))))))))
                                                                    v3 = ((float32((v65 / -2.25)) + v3) * v3)
                                                                    _ = v65
                                                                }
                                                            }
                                                        }
                                                        var v66 uint = v8
                                                        _ = v62
                                                        _ = v66
                                                    }
                                                }
                                            }
                                            v67 := GenS3{F0: 1000}
                                            _ = v67
                                        } else {
                                            if len(v5) > 0 {
                                                v5[0] = (("foo bar" + "foo bar") + ("a" + v4))
                                            }
                                            if false {
                                                v1 = ("foo bar" + v4)
                                                var v68 uint8 = 255
                                                if false {
                                                    if true {
                                                        var v69 uint32 = uint32(v57)
                                                        _ = v69
                                                    }
                                                    v70 := GenS2{F0: ("foo bar" + v4), F1: 1}
                                                    var v71 uint = 18446744073709551615
                                                    fmt.Println(v3, fi36)
                                                    v72 := GenS6{F0: (fi36 % (((fi9 ^ fi36)) | 1)), F1: v57, F2: 1}
                                                    v2[(fi9 & fi36)] = 4294967295
                                                    _ = v70
                                                    _ = v71
                                                    _ = v72
                                                } else {
                                                    if false {
                                                        if true {
                                                            if true {
                                                                if len(v5) >= 2 {
                                                                    v5 = v5[1:len(v5)]
                                                                }
                                                            } else {
                                                                var v73 float32 = (v3 * -2.25)
                                                                if false {
                                                                    fmt.Println(v68)
                                                                    if true {
                                                                        v73 = v3
                                                                        var v74 uint64 = 42
                                                                        fmt.Println(v4, fi36)
                                                                        _ = v74
                                                                    }
                                                                    if ((v8 >> (((^(v8))) & 63)) <= v8) {
                                                                        if true {
                                                                            var v75 int = 127
                                                                            if (((v57 - 255) <= v57) || true) {
                                                                                switch 128 {
                                                                                case 427:
                                                                                    if len(v5) > 0 {
                                                                                        v5[0] = v1
                                                                                    }
                                                                                case 9223372036854775807:
                                                                                    switch (v4 + v1) {
                                                                                    case "":
                                                                                        var v76 float32 = v73
                                                                                        _ = v76
                                                                                    case "foo bar":
                                                                                        var v77 uint32 = 42
                                                                                        _ = v77
                                                                                    case "\n\t":
                                                                                        v78 := GenS8{F0: 65535, F1: v1, F2: v68}
                                                                                        _ = v78
                                                                                    }
                                                                                case 127:
                                                                                    if ("a" < ((("a" + v4) + v1) + v4)) {
                                                                                        v5 = append(v5, "hello")
                                                                                    }
                                                                                default:
                                                                                    fmt.Println(v75)
                                                                                }
                                                                                v57 -= (v57 << (((v8 * v8)) & 63))
                                                                                v79 := GenS9{F0: 0, F1: v75}
                                                                                var v80 float32 = ((0.0 + float32(0.1)) + (v3 * ((v73 - 1.5) - float32(((3.14159 - (0.1 - 0.1)) + 1.5)))))
                                                                                _ = v79
                                                                                _ = v80
                                                                            }
                                                                            _ = v75
                                                                        }
                                                                        var v81 uint64 = 67
                                                                        switch ("\n\t" + (v1 + v4)) {
                                                                        case "\n\t":
                                                                            var v82 uint16 = 382
                                                                            _ = v82
                                                                        case "":
                                                                            var v83 int64 = -9223372036854775808
                                                                            _ = v83
                                                                        case "a":
                                                                            if (int16(v81) >= 930) {
                                                                                var v84 string = v4
                                                                                _ = v84
                                                                            } else {
                                                                                v57 += (v57 - (v57 | v57))
                                                                                v3 = 1.5
                                                                            }
                                                                            delete(v2, fi36)
                                                                            fmt.Println(fi9, v57, v73)
                                                                        default:
                                                                            switch fi9 {
                                                                            case 32767:
                                                                                var v85 uint64 = (v81 ^ (v81 + (v81 | v81)))
                                                                                _ = v85
                                                                            case 9223372036854775807:
                                                                                var v86 uint64 = (v81 & v81)
                                                                                _ = v86
                                                                            }
                                                                        }
                                                                        _ = v81
                                                                    } else {
                                                                        var v87 uint16 = v57
                                                                        v4 = ""
                                                                        _ = v87
                                                                    }
                                                                }
                                                                fmt.Println(v8, v4)
                                                                _ = v73
                                                            }
                                                            v88 := GenS3{F0: 42}
                                                            _ = v88
                                                        }
                                                        if (!true) {
                                                            v4 = ("\n\t" + v4)
                                                            var v89 int8 = 0
                                                            _ = v89
                                                        }
                                                        var v90 uint64 = 2
                                                        if true {
                                                            var v91 int64 = 458
                                                            fmt.Println(v4)
                                                            _ = v91
                                                        } else {
                                                            fmt.Println(fi36)
                                                        }
                                                        _ = v90
                                                    } else {
                                                        v92 := GenS4{F0: 42, F1: 739}
                                                        _ = v92
                                                    }
                                                    v3 += (v3 - 0.0)
                                                }
                                                v5 = append(v5, v4)
                                                _ = v68
                                            } else {
                                                v8 += (v8 | (v8 & 4294967295))
                                                v1 = v1
                                                fmt.Println(v1)
                                            }
                                        }
                                        var v93 float32 = v3
                                        if (!true) {
                                            if len(v5) > 0 {
                                                v5[0] = v1
                                            }
                                        } else {
                                            if (!(true && true)) {
                                                var v94 uint8 = 48
                                                _ = v94
                                            }
                                            if (false && ("foo bar" <= "a")) {
                                                v93 -= ((((v93 / ((v3 * v3) * 3.14159)) + (3.14159 / 3.14159)) * float32(float64((float32(1e10) + (v3 + -2.25))))) - 0.1)
                                                var v95 string = "a"
                                                _ = v95
                                            } else {
                                                v96 := GenS7{F0: fi36, F1: "hello", F2: true}
                                                _ = v96
                                            }
                                            fmt.Println(fi9)
                                            var v97 string = (v1 + "\n\t")
                                            if (983 <= 1) {
                                                var v98 uint8 = 1
                                                v99 := GenS5{F0: 0, F1: v4, F2: (^(fi9))}
                                                var v100 int64 = 0
                                                _ = v100
                                                _ = v98
                                                _ = v99
                                            }
                                            _ = len(v5)
                                            _ = v97
                                        }
                                        _ = v57
                                        _ = v93
                                    } else {
                                        delete(v2, fi9)
                                        fmt.Println(fi36)
                                    }
                                }
                                if (1 == 42) {
                                    v2[(fi36 ^ (-(fi9)))] = 886
                                    if (((((v3 * v3) + 1e10) - 3.14159) * 1e10) < v3) {
                                        switch ("hello" + v4) {
                                        case "a":
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                        default:
                                            var v101 uint16 = 1
                                            _ = v101
                                        }
                                        if (!false) {
                                            fmt.Println(v4)
                                        }
                                        var v102 string = v1
                                        _ = v102
                                    } else {
                                        var v103 uint16 = uint16(v8)
                                        v103 = v103
                                        _ = v103
                                    }
                                }
                                v2[(fi9 & fi9)] = 42
                                _ = len(v5)
                            }
                        }
                    }
                }
            }
            var v104 uint64 = 0
            _ = fi36
            _ = v104
        }
        _ = fi9
    }
    v3 = v3
    var v105 bool = true
    fmt.Println(v1)
    for fi106 := 0; fi106 < 20; fi106++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v2[(fi106 >> ((62) & 63))] = 2147483648
        var v107 bool = v105
        _ = fi106
        _ = v107
    }
    if false {
        v108 := GenS5{F0: 573, F1: "\n\t", F2: 0}
        v109 := GenS10{F0: -445}
        if v105 {
            v110 := v2[1]
            v3 = 0.1
            _ = v110
        } else {
            v111 := GenS7{F0: -799, F1: ("hello" + v4), F2: (v105 || true)}
            switch v1 {
            case "hello":
                var v112 float32 = 1.5
                _ = v112
            default:
                for fi113 := 0; fi113 < 8; fi113++ {
                    if (!v105) {
                        switch (^(fi113)) {
                        case -608:
                            for fi114 := 0; fi114 < 7; fi114++ {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                if len(v5) > 0 {
                                    v5[0] = ""
                                }
                                _ = fi114
                            }
                        case -246:
                            v2[1] = 2
                        case 66:
                            for fi115 := 0; fi115 < 3; fi115++ {
                                v3 = v3
                                var v116 int32 = -943
                                v117 := GenS11{F0: true, F1: -32768, F2: (fi113 - (-(fi113)))}
                                if (42 <= 967) {
                                    v105 = v105
                                    var v118 uint = v8
                                    switch (fi113 * (-(fi115))) {
                                    case -434:
                                        switch 0 {
                                        case 4294967295:
                                            if v105 {
                                                v4 = (v4 + (("" + v4) + "hello"))
                                            } else {
                                                if true {
                                                    var v119 float64 = 0.0
                                                    v118 -= (^(v8))
                                                    fmt.Println(v119, v8)
                                                    _ = v119
                                                } else {
                                                    switch fi113 {
                                                    case -500:
                                                        v4 = (v4 + ((((("foo bar" + ((("\n\t" + v1) + v4) + v4)) + (v1 + (v4 + "hello"))) + "foo bar") + (("" + "foo bar") + "hello")) + "a"))
                                                    }
                                                    if (370 <= 1) {
                                                        if v105 {
                                                            var v120 bool = true
                                                            _ = v120
                                                        } else {
                                                            if (false || (!false)) {
                                                                _ = len(v2)
                                                                if len(v5) > 0 {
                                                                    v5[0] = (v1 + (v4 + (v4 + (v4 + (((v1 + "foo bar") + (((("\n\t" + v1) + v1) + (v4 + v1)) + ("" + ("hello" + ((v4 + (((v4 + v1) + v4) + "foo bar")) + (v1 + v4)))))) + v4)))))
                                                                }
                                                                v3 = 1e10
                                                            }
                                                            v121 := GenS12{F0: fi115}
                                                            if (!(fi115 == (^(fi115)))) {
                                                                v3 += (3.14159 - v3)
                                                            }
                                                            _ = v121
                                                        }
                                                        switch fi115 {
                                                        case 785:
                                                            switch fi115 {
                                                            case 9223372036854775807:
                                                                v3 -= 3.14159
                                                                if (!v105) {
                                                                    switch v1 {
                                                                    case "a":
                                                                        fmt.Println(fi115)
                                                                    case "\n\t":
                                                                        v105 = v105
                                                                    case "foo bar":
                                                                        switch fi115 {
                                                                        case 127:
                                                                            if v105 {
                                                                                var v122 bool = (false || false)
                                                                                _ = v122
                                                                            } else {
                                                                                if true {
                                                                                    v123 := GenS12{F0: 42}
                                                                                    _ = v123
                                                                                }
                                                                                fmt.Println(v4)
                                                                                _ = len(v5)
                                                                            }
                                                                        default:
                                                                            if v105 {
                                                                                v4 = ("\n\t" + ("a" + (v4 + (v4 + ((((v4 + ("" + "foo bar")) + "hello") + v4) + v1)))))
                                                                            } else {
                                                                                var v124 int = -863
                                                                                var v125 int64 = -1
                                                                                v124 -= fi113
                                                                                v3 -= (float32(float64(1.5)) / v3)
                                                                                _ = v124
                                                                                _ = v125
                                                                            }
                                                                        }
                                                                    }
                                                                    fmt.Println(v4)
                                                                    v3 += (v3 - v3)
                                                                    v8 = v118
                                                                    v1 = (("hello" + ((v1 + v4) + ("a" + v4))) + v1)
                                                                }
                                                            case -1:
                                                                fmt.Println(v105, v8, fi113)
                                                            default:
                                                                switch v1 {
                                                                case "a":
                                                                    var v126 int16 = 156
                                                                    _ = v126
                                                                case "foo bar":
                                                                    switch fi113 {
                                                                    case 1:
                                                                        v127 := GenS3{F0: -92}
                                                                        _ = v127
                                                                    case 2147483647:
                                                                        var v128 int32 = 0
                                                                        _ = v128
                                                                    case -9223372036854775808:
                                                                        v1 = v4
                                                                    default:
                                                                        var v129 uint64 = 18446744073709551615
                                                                        _ = v129
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        if ("a" < v4) {
                                                            v105 = true
                                                        }
                                                        if (100 >= 59) {
                                                            var v130 uint = v118
                                                            switch (fi115 >> (((v8 | v118)) & 63)) {
                                                            case 42:
                                                                if (127 == 109) {
                                                                    var v131 bool = v105
                                                                    _ = v131
                                                                }
                                                            case 127:
                                                                if ((1.5 - (float32((-2.25 - 0.1)) - v3)) > float32(float64(((0.0 - float32(3.14159)) / (((v3 - v3) * 1.5) + 3.14159))))) {
                                                                    v132 := GenS13{F0: fi115, F1: 2}
                                                                    v1 = v4
                                                                    v118 = v130
                                                                    var v133 int64 = 42
                                                                    v4 = "a"
                                                                    _ = v132
                                                                    _ = v133
                                                                }
                                                            default:
                                                                switch "a" {
                                                                case "hello":
                                                                    var v134 float64 = 0.0
                                                                    _ = v134
                                                                case "a":
                                                                    v3 = v3
                                                                }
                                                            }
                                                            _ = v130
                                                        }
                                                        if (!v105) {
                                                            if (-32768 <= -32768) {
                                                                delete(v2, 0)
                                                            } else {
                                                                var v135 float64 = -2.25
                                                                if v105 {
                                                                    if len(v5) >= 2 {
                                                                        v5 = v5[1:len(v5)]
                                                                    }
                                                                    var v136 uint8 = 128
                                                                    if (!(float64(0.0) == v135)) {
                                                                        if v105 {
                                                                            var v137 string = ((v4 + v4) + (("\n\t" + (("\n\t" + (v1 + v1)) + "")) + v4))
                                                                            _ = v137
                                                                        } else {
                                                                            switch (v1 + ((("\n\t" + (("foo bar" + "\n\t") + (v1 + v4))) + v1) + (v1 + v1))) {
                                                                            case "a":
                                                                                v4 = v4
                                                                                if (-2147483648 == v116) {
                                                                                    v138 := GenS2{F0: "\n\t", F1: 0}
                                                                                    if (!(v105 && ((v136 - (^(v136))) < v136))) {
                                                                                        v5 = append(v5, ((v4 + "\n\t") + (v1 + v1)))
                                                                                        var v139 int8 = 127
                                                                                        _ = v139
                                                                                    } else {
                                                                                        if (!v105) {
                                                                                            v140 := GenS14{F0: v1, F1: 32767}
                                                                                            var v141 string = v4
                                                                                            if len(v5) > 0 {
                                                                                                v5[0] = (("foo bar" + v1) + v4)
                                                                                            }
                                                                                            _ = v140
                                                                                            _ = v141
                                                                                        } else {
                                                                                            v142 := GenS15{F0: (v135 < (v135 + v135))}
                                                                                            v143 := GenS13{F0: (fi115 | fi115), F1: (v136 ^ v136)}
                                                                                            v8 += 2
                                                                                            _ = v142
                                                                                            _ = v143
                                                                                        }
                                                                                        v136 = v136
                                                                                        if false {
                                                                                            fmt.Println(v8)
                                                                                            switch v1 {
                                                                                            case "foo bar":
                                                                                                var v144 int64 = -9223372036854775808
                                                                                                _ = v144
                                                                                            case "\n\t":
                                                                                                var v145 int8 = 45
                                                                                                _ = v145
                                                                                            case "a":
                                                                                                v135 = v135
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    if (!(2 > 2)) {
                                                                                        fmt.Println(v135, v105, v3)
                                                                                        fmt.Println(v3)
                                                                                        v118 += v118
                                                                                    }
                                                                                    _ = v138
                                                                                } else {
                                                                                    var v146 uint16 = 697
                                                                                    var v147 float32 = float32(1.5)
                                                                                    v146 = (^(v146))
                                                                                    _ = v146
                                                                                    _ = v147
                                                                                }
                                                                                var v148 int32 = 711
                                                                                fmt.Println(v135, fi113)
                                                                                _ = v148
                                                                            case "\n\t":
                                                                                var v149 uint64 = 42
                                                                                _ = v149
                                                                            case "":
                                                                                _ = len(v5)
                                                                            default:
                                                                                var v150 int16 = 1000
                                                                                _ = v150
                                                                            }
                                                                        }
                                                                    }
                                                                    _ = v136
                                                                }
                                                                _ = v135
                                                            }
                                                            var v151 uint8 = 255
                                                            _ = v151
                                                        } else {
                                                            v152 := GenS14{F0: "", F1: 42}
                                                            _ = v152
                                                        }
                                                    } else {
                                                        var v153 int64 = -9223372036854775808
                                                        v154 := GenS16{F0: -32768}
                                                        _ = v153
                                                        _ = v154
                                                    }
                                                }
                                                fmt.Println(v116, fi113, v105)
                                                var v155 uint32 = 42
                                                _ = v155
                                            }
                                        }
                                    default:
                                        var v156 uint32 = 65535
                                        _ = v156
                                    }
                                    _ = v118
                                } else {
                                    v157 := GenS11{F0: true, F1: -529, F2: (fi113 << ((255) & 63))}
                                    _ = v157
                                }
                                _ = fi115
                                _ = v116
                                _ = v117
                            }
                        default:
                            v158 := GenS7{F0: 4294967295, F1: v4, F2: (1e10 < 0.0)}
                            _ = v158
                        }
                        var v159 uint = 18446744073709551615
                        v160 := GenS16{F0: 32767}
                        _ = v159
                        _ = v160
                    }
                    v161 := GenS15{F0: (true && v105)}
                    _ = fi113
                    _ = v161
                }
            }
            _ = v111
        }
        _ = v108
        _ = v109
    }
    fmt.Println(v3, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v105:", v105)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v105
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
