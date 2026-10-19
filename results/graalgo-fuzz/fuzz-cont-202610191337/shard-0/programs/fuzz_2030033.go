package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 int; F1 bool }
type GenS4 struct { F0 uint64; F1 uint16 }
type GenS5 struct { F0 uint8; F1 uint16 }
type GenS6 struct { F0 uint8; F1 uint32 }
type GenS7 struct { F0 int32; F1 uint16; F2 uint16 }
type GenS8 struct { F0 int8; F1 uint16; F2 string }
type GenS9 struct { F0 uint8; F1 uint; F2 uint8 }
type GenS10 struct { F0 int8; F1 int64; F2 string }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 int16 }
type GenS13 struct { F0 string; F1 int64 }
type GenS14 struct { F0 string }
type GenS15 struct { F0 uint16; F1 int8; F2 int8 }
func genFn0(p0 string, p1 string, p2 int64) int {
    var v6 bool = (35 < 42)
    var v7 uint8 = 128
    if v6 {
        fmt.Println(p0, v6, v7)
        p0 = (p0 + "foo bar")
    }
    var v8 string = "foo bar"
    if v6 {
        for fi9 := 0; fi9 < 10; fi9++ {
            switch (fi9 >> ((0) & 63)) {
            case 128:
                v8 = "hello"
            }
            for fi10 := 0; fi10 < 15; fi10++ {
                if (!((1000 != -463) || (65535 <= 255))) {
                    var v11 float64 = ((1e10 - (3.14159 + 0.0)) + 0.0)
                    fmt.Println(v11, v6, p0)
                    fmt.Println(v11, fi9, v7)
                    _ = v11
                } else {
                    switch (v8 + v8) {
                    case "a":
                        var v12 float32 = 0.1
                        _ = v12
                    case "hello":
                        if (1.5 == -2.25) {
                            var v13 bool = (-32768 < -32768)
                            v14 := GenS1{F0: 127, F1: 2147483648}
                            _ = v13
                            _ = v14
                        } else {
                            v6 = (((!(!true)) && true) || ((("a" + "\n\t") > "\n\t") || (fi9 <= (-(fi9)))))
                            var v15 float32 = float32(float64((0.1 + 3.14159)))
                            var v16 uint8 = v7
                            _ = v15
                            _ = v16
                        }
                    }
                    v17 := []int{fi10, fi10, 32767, fi9}
                    if (v6 && (v6 || v6)) {
                        var v18 float32 = (1.5 * (3.14159 + -2.25))
                        if (255 == v7) {
                            var v19 uint16 = uint16(v7)
                            var v20 bool = (!(0 <= -1))
                            _ = v19
                            _ = v20
                        } else {
                            p0 = v8
                            fmt.Println(v8)
                            if (664 < 1000) {
                                v8 = (p0 + ("\n\t" + p0))
                                v8 = (p0 + "a")
                                switch "foo bar" {
                                case "":
                                    v21 := GenS2{F0: v7, F1: 0}
                                    _ = v21
                                case "a":
                                    switch "\n\t" {
                                    case "hello":
                                        switch fi9 {
                                        case 4294967295:
                                            var v22 uint8 = v7
                                            _ = v22
                                        default:
                                            switch "foo bar" {
                                            case "foo bar":
                                                v7 += v7
                                                if true {
                                                    switch p0 {
                                                    case "\n\t":
                                                        var v23 int32 = 1
                                                        _ = v23
                                                    case "foo bar":
                                                        if len(v17) >= 2 {
                                                            v17 = v17[1:len(v17)]
                                                        }
                                                    case "hello":
                                                        p0 = v8
                                                    default:
                                                        if (42 >= 4294967295) {
                                                            if len(v17) > 0 {
                                                                v17[0] = fi10
                                                            }
                                                            var v24 bool = (255 == 911)
                                                            _ = v24
                                                        } else {
                                                            fmt.Println(p2, p1)
                                                            var v25 int8 = -1
                                                            _ = v25
                                                        }
                                                    }
                                                }
                                                v18 = (v18 + -2.25)
                                            case "\n\t":
                                                if v6 {
                                                    if v6 {
                                                        if len(v17) > 0 {
                                                            v17[0] = 128
                                                        }
                                                    }
                                                    if (18446744073709551615 < 615) {
                                                        var v26 int32 = -1
                                                        switch (fi10 ^ fi9) {
                                                        case -9223372036854775808:
                                                            var v27 string = v8
                                                            _ = v27
                                                        }
                                                        _ = v26
                                                    }
                                                } else {
                                                    fmt.Println(p1)
                                                    p1 = p0
                                                    if v6 {
                                                        switch (fi10 % ((-542) | 1)) {
                                                        case -9223372036854775808:
                                                            switch (((v8 + "a") + (p1 + (p0 + (("a" + p0) + p1)))) + (p0 + ("foo bar" + p0))) {
                                                            case "\n\t":
                                                                if v6 {
                                                                    v28 := GenS3{F0: 32767, F1: (((!(float32(((((float64(-2.25) + ((((3.14159 * 1.5) * float64(3.14159)) + 1.5) - 0.1)) * 0.1) - 0.1) + 0.0)) <= (((v18 * 1e10) - 0.0) + (1.5 * v18)))) || v6) && false)}
                                                                    v29 := GenS3{F0: (fi9 / (((fi10 - (fi9 * (-(fi10))))) | 1)), F1: (!(true || false))}
                                                                    _ = v28
                                                                    _ = v29
                                                                }
                                                            case "a":
                                                                p2 -= (p2 / (((p2 >> ((42) & 63))) | 1))
                                                            case "":
                                                                if len(v17) > 0 {
                                                                    v17[0] = -452
                                                                }
                                                            }
                                                        }
                                                    } else {
                                                        var v30 bool = false
                                                        _ = v30
                                                    }
                                                }
                                                v31 := GenS3{F0: 32768, F1: (!v6)}
                                                if v6 {
                                                    if true {
                                                        if v6 {
                                                            var v32 string = ""
                                                            _ = v32
                                                        } else {
                                                            var v33 bool = (fi10 <= 2147483647)
                                                            if ((!v6) && (!(!true))) {
                                                                v7 += 255
                                                                v34 := GenS4{F0: 2, F1: 1}
                                                                _ = v34
                                                            } else {
                                                                fmt.Println(v6, p1)
                                                                fmt.Println(p0, v18, v18)
                                                                var v35 uint64 = 4294967295
                                                                v7 += (v7 % (((v7 % ((v7) | 1))) | 1))
                                                                _ = v35
                                                            }
                                                            _ = v33
                                                        }
                                                    }
                                                }
                                                p2 += (p2 - (p2 / (((p2 % (((p2 / (((p2 - p2)) | 1))) | 1))) | 1)))
                                                if v6 {
                                                    fmt.Println(p1, p1, fi9)
                                                    if (!(false && (false && (!false)))) {
                                                        if (!v6) {
                                                            switch 42 {
                                                            case -692:
                                                                p2 += 42
                                                            case 32768:
                                                                fmt.Println(v6, fi10)
                                                            }
                                                        } else {
                                                            if len(v17) > 0 {
                                                                v17[0] = (^(fi10))
                                                            }
                                                            v7 = (v7 << ((1) & 63))
                                                            v36 := GenS4{F0: 2, F1: 2}
                                                            _ = v36
                                                        }
                                                        var v37 float64 = (3.14159 * float64(v18))
                                                        _ = v37
                                                    }
                                                    var v38 int32 = -2147483648
                                                    _ = v38
                                                } else {
                                                    p1 = ""
                                                    var v39 int8 = -128
                                                    _ = v39
                                                }
                                                var v40 uint64 = 0
                                                _ = v31
                                                _ = v40
                                            case "hello":
                                                fmt.Println(p2, v8, v6)
                                            }
                                        }
                                    case "":
                                        _ = len(v17)
                                    default:
                                        v6 = ((v18 - v18) > 1e10)
                                    }
                                    if v6 {
                                        if v6 {
                                            if true {
                                                v18 = v18
                                            } else {
                                                var v41 float64 = (0.1 / (1.5 - float64(v18)))
                                                fmt.Println(v8)
                                                var v42 uint64 = 2
                                                _ = v41
                                                _ = v42
                                            }
                                            var v43 int8 = -1
                                            fmt.Println(v8, v7, v43)
                                            _ = v43
                                        } else {
                                            var v44 float64 = 0.0
                                            v8 = v8
                                            switch "\n\t" {
                                            case "hello":
                                                v44 -= (v44 + v44)
                                            case "foo bar":
                                                if len(v17) > 0 {
                                                    v17[0] = (-(fi10))
                                                }
                                            default:
                                                if (v7 < v7) {
                                                    p1 = v8
                                                    if len(v17) >= 2 {
                                                        v17 = v17[1:len(v17)]
                                                    }
                                                    if len(v17) > 0 {
                                                        v17[0] = (^(fi9))
                                                    }
                                                } else {
                                                    if (65535 > 65535) {
                                                        v45 := GenS5{F0: 127, F1: 42}
                                                        _ = v45
                                                    }
                                                    if (2 == 65535) {
                                                        var v46 float32 = float32(v44)
                                                        var v47 bool = ((583 == -1) && (4294967295 > 42))
                                                        _ = v46
                                                        _ = v47
                                                    }
                                                }
                                            }
                                            _ = v44
                                        }
                                    } else {
                                        var v48 float64 = 3.14159
                                        fmt.Println(v8, v18)
                                        fmt.Println(v18, p1)
                                        v6 = (1 == 42)
                                        _ = v48
                                    }
                                    if (4294967295 >= 908) {
                                        v6 = (1.5 != -2.25)
                                        fmt.Println(fi9, v8)
                                    }
                                case "\n\t":
                                    switch (fi10 % ((fi10) | 1)) {
                                    case 709:
                                        _ = len(v17)
                                    case 4294967295:
                                        var v49 int64 = (^(p2))
                                        _ = v49
                                    case 863:
                                        if ((!(!true)) && (42 != 2)) {
                                            fmt.Println(v6, v6)
                                        } else {
                                            if v6 {
                                                var v50 int16 = -1
                                                var v51 bool = false
                                                switch int(v50) {
                                                case -432:
                                                    if (-9223372036854775808 < p2) {
                                                        v52 := GenS6{F0: (v7 | v7), F1: 616}
                                                        if v6 {
                                                            fmt.Println(v50, p2, v51)
                                                        }
                                                        p2 = p2
                                                        _ = v52
                                                    }
                                                case -957:
                                                    if (42 < -127) {
                                                        switch fi10 {
                                                        case 0:
                                                            var v53 bool = false
                                                            _ = v53
                                                        case -230:
                                                            v17 = append(v17, fi10)
                                                        }
                                                        if (-2147483648 < 353) {
                                                            v54 := GenS2{F0: v7, F1: 4294967295}
                                                            var v55 float32 = v18
                                                            var v56 int32 = int32(v50)
                                                            _ = v54
                                                            _ = v55
                                                            _ = v56
                                                        } else {
                                                            p0 = p1
                                                            v7 += (v7 | (^(v7)))
                                                            var v57 int = fi10
                                                            _ = v57
                                                        }
                                                        fmt.Println(fi10)
                                                    } else {
                                                        v58 := GenS7{F0: 2147483647, F1: 255, F2: 849}
                                                        switch ("hello" + p1) {
                                                        case "\n\t":
                                                            v51 = false
                                                        default:
                                                            if v6 {
                                                                v7 -= 42
                                                                if len(v17) > 0 {
                                                                    v17[0] = 668
                                                                }
                                                            } else {
                                                                v59 := GenS3{F0: 180, F1: ((^(v7)) <= 128)}
                                                                var v60 int8 = 65
                                                                _ = v59
                                                                _ = v60
                                                            }
                                                        }
                                                        var v61 int32 = 42
                                                        _ = v58
                                                        _ = v61
                                                    }
                                                case 32767:
                                                    switch (fi9 - int(v50)) {
                                                    case 4294967295:
                                                        if ((539 > 978) && (v7 <= (v7 & (v7 & v7)))) {
                                                            v17 = append(v17, (fi10 << ((0) & 63)))
                                                        } else {
                                                            fmt.Println(p0, v6, fi9)
                                                            _ = len(v17)
                                                            v62 := GenS8{F0: -128, F1: 1, F2: p0}
                                                            _ = v62
                                                        }
                                                    case 32768:
                                                        if len(v17) > 0 {
                                                            v17[0] = fi10
                                                        }
                                                    default:
                                                        var v63 bool = ((p2 >> ((42) & 63)) > -907)
                                                        _ = v63
                                                    }
                                                default:
                                                    var v64 int64 = (p2 & (^(p2)))
                                                    _ = v64
                                                }
                                                _ = v50
                                                _ = v51
                                            } else {
                                                if true {
                                                    p0 = (p1 + "foo bar")
                                                    p1 = v8
                                                    var v65 float64 = -2.25
                                                    _ = v65
                                                }
                                                var v66 uint16 = 1
                                                _ = v66
                                            }
                                            var v67 string = v8
                                            if false {
                                                fmt.Println(v18, v18)
                                                var v68 float32 = 0.1
                                                p2 = p2
                                                v69 := GenS9{F0: 4, F1: 625, F2: v7}
                                                var v70 uint32 = 65535
                                                _ = v68
                                                _ = v69
                                                _ = v70
                                            }
                                            _ = v67
                                        }
                                    }
                                }
                            } else {
                                if (-32768 <= 1000) {
                                    if v6 {
                                        v71 := GenS10{F0: 100, F1: 1, F2: "hello"}
                                        fmt.Println(v7)
                                        fmt.Println(v6, v18)
                                        _ = v71
                                    } else {
                                        fmt.Println(v8, fi10, p1)
                                        switch fi10 {
                                        case 1:
                                            if (42 >= -1) {
                                                v7 += v7
                                                var v72 int32 = 1
                                                v18 = 3.14159
                                                _ = v72
                                            }
                                        case 0:
                                            var v73 string = "\n\t"
                                            _ = v73
                                        case 353:
                                            var v74 uint8 = 127
                                            _ = v74
                                        }
                                        switch "foo bar" {
                                        case "\n\t":
                                            switch p1 {
                                            case "foo bar":
                                                var v75 int32 = -1
                                                _ = v75
                                            case "a":
                                                var v76 int64 = (-(p2))
                                                _ = v76
                                            default:
                                                if v6 {
                                                    var v77 int8 = 123
                                                    if v6 {
                                                        var v78 string = (p1 + (("hello" + ("a" + ("foo bar" + ((("a" + p1) + p0) + "hello")))) + "hello"))
                                                        _ = len(v17)
                                                        var v79 float32 = v18
                                                        _ = v78
                                                        _ = v79
                                                    } else {
                                                        var v80 uint16 = 0
                                                        _ = v80
                                                    }
                                                    _ = len(v17)
                                                    v77 = v77
                                                    _ = v77
                                                }
                                            }
                                        }
                                    }
                                } else {
                                    if len(v17) >= 2 {
                                        v17 = v17[1:len(v17)]
                                    }
                                    var v81 int = fi10
                                    p1 = (p0 + "a")
                                    _ = v81
                                }
                                switch (fi9 + fi9) {
                                case -484:
                                    if (!(1 < (^(v7)))) {
                                        _ = len(v17)
                                    }
                                case 9223372036854775807:
                                    if len(v17) >= 2 {
                                        v17 = v17[1:len(v17)]
                                    }
                                }
                                v18 += v18
                            }
                        }
                        _ = v18
                    }
                    _ = v17
                }
                if v6 {
                    var v82 uint16 = 255
                    var v83 uint = 1
                    v6 = false
                    _ = v82
                    _ = v83
                }
                switch -99 {
                case 32768:
                    v7 += (^(v7))
                case -2147483648:
                    if false {
                        if ((42 == 42) && (!(1 >= 1))) {
                            v7 -= v7
                            var v84 int64 = (p2 << ((4294967295) & 63))
                            fmt.Println(fi9, v6)
                            v85 := []int{int(v7), (-(fi10)), (fi9 ^ (fi9 / ((fi10) | 1)))}
                            _ = len(v85)
                            _ = v84
                            _ = v85
                        }
                        var v86 int64 = p2
                        fmt.Println(v86)
                        _ = v86
                    } else {
                        fmt.Println(fi10, v7)
                        v8 = "foo bar"
                        v7 -= (v7 << ((65535) & 63))
                    }
                default:
                    var v87 int = (^(fi10))
                    _ = v87
                }
                _ = fi10
            }
            v8 = p0
            _ = fi9
        }
        if (42 < 0) {
            v88 := GenS5{F0: (^(v7)), F1: 255}
            v89 := GenS11{F0: 266}
            if (v6 && true) {
                if (12 != 42) {
                    v90 := []int{9223372036854775807, -9223372036854775808}
                    fmt.Println(p0, v6, v6)
                    for fi91 := 0; fi91 < 10; fi91++ {
                        switch "a" {
                        case "":
                            v6 = (!v6)
                        case "hello":
                            var v92 int16 = -32768
                            _ = v92
                        case "foo bar":
                            if len(v90) > 0 {
                                v90[0] = (fi91 >> ((844) & 63))
                            }
                        default:
                            for fi93 := 0; fi93 < 6; fi93++ {
                                p0 = (("foo bar" + p1) + v8)
                                var v94 uint32 = 928
                                _ = fi93
                                _ = v94
                            }
                        }
                        _ = fi91
                    }
                    _ = v90
                } else {
                    v95 := GenS6{F0: 23, F1: 162}
                    _ = v95
                }
                v96 := GenS8{F0: 0, F1: 32767, F2: ("foo bar" + (v8 + ""))}
                _ = v96
            }
            _ = v88
            _ = v89
        } else {
            switch -1000 {
            case 4294967295:
                switch 128 {
                case 367:
                    for fi97 := 0; fi97 < 8; fi97++ {
                        for fi98 := 0; fi98 < 17; fi98++ {
                            v6 = v6
                            var v99 bool = (42 > 981)
                            _ = fi98
                            _ = v99
                        }
                        if false {
                            v100 := GenS9{F0: (v7 ^ (v7 << ((1) & 63))), F1: 4294967295, F2: v7}
                            var v101 float64 = (1.5 - 1e10)
                            for fi102 := 0; fi102 < 16; fi102++ {
                                var v103 int8 = -110
                                if (-1 <= -1) {
                                    v104 := []int{fi97, 821, fi97}
                                    _ = v104
                                } else {
                                    var v105 int32 = 105
                                    v106 := []int{153, (fi102 % ((fi102) | 1))}
                                    v107 := GenS12{F0: -596}
                                    _ = v105
                                    _ = v106
                                    _ = v107
                                }
                                _ = fi102
                                _ = v103
                            }
                            _ = v100
                            _ = v101
                        } else {
                            fmt.Println(v6, v8)
                            p2 = (p2 + (^(p2)))
                            fmt.Println(p1, p1)
                        }
                        if true {
                            for fi108 := 0; fi108 < 8; fi108++ {
                                var v109 int8 = -95
                                switch v8 {
                                case "a":
                                    v110 := GenS6{F0: 69, F1: 863}
                                    _ = v110
                                }
                                switch 0 {
                                case 131:
                                    var v111 string = (p1 + (v8 + "hello"))
                                    _ = v111
                                }
                                _ = fi108
                                _ = v109
                            }
                            var v112 uint8 = 127
                            v8 = p0
                            _ = v112
                        }
                        _ = fi97
                    }
                default:
                    switch int(v7) {
                    case 389:
                        p0 = "foo bar"
                    case 4294967295:
                        for fi113 := 0; fi113 < 17; fi113++ {
                            var v114 uint16 = 1
                            switch int(v7) {
                            case 42:
                                v115 := GenS9{F0: (v7 ^ v7), F1: 0, F2: 240}
                                _ = v115
                            case -775:
                                var v116 uint8 = (v7 % (((^(v7))) | 1))
                                _ = v116
                            case 128:
                                p2 -= (p2 / ((-79) | 1))
                            default:
                                var v117 uint8 = (^(v7))
                                _ = v117
                            }
                            p0 = (p1 + (p0 + (v8 + (((v8 + v8) + p1) + v8))))
                            _ = fi113
                            _ = v114
                        }
                    }
                }
            case 228:
                var v118 float32 = float32(1e10)
                _ = v118
            case 126:
                for fi119 := 0; fi119 < 2; fi119++ {
                    var v120 uint32 = 107
                    if (false && ((v7 - v7) == (v7 * 21))) {
                        switch (("\n\t" + p1) + "hello") {
                        case "foo bar":
                            switch v8 {
                            case "hello":
                                if (2 >= (v7 | (v7 << ((255) & 63)))) {
                                    v6 = (-89 == -755)
                                    for fi121 := 0; fi121 < 3; fi121++ {
                                        var v122 uint = 65535
                                        _ = fi121
                                        _ = v122
                                    }
                                } else {
                                    var v123 int32 = -2147483648
                                    for fi124 := 0; fi124 < 15; fi124++ {
                                        fmt.Println(p1, v120, v6)
                                        fmt.Println(v7, p2)
                                        p0 = "a"
                                        _ = fi124
                                    }
                                    _ = v123
                                }
                            default:
                                v6 = v6
                            }
                        }
                        v6 = v6
                    }
                    _ = fi119
                    _ = v120
                }
            }
            v8 = v8
            v125 := []int{}
            _ = v125
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 uint64) int {
    p0 = p0
    if (false && false) {
        v126 := GenS13{F0: "hello", F1: -9223372036854775808}
        for fi127 := 0; fi127 < 7; fi127++ {
            fmt.Println(fi127, p0)
            _ = fi127
        }
        for fi128 := 0; fi128 < 10; fi128++ {
            for fi129 := 0; fi129 < 8; fi129++ {
                p0 = p0
                if (!false) {
                    v130 := GenS8{F0: 127, F1: 1, F2: "\n\t"}
                    _ = v130
                }
                v131 := GenS14{F0: ("hello" + "\n\t")}
                v132 := GenS12{F0: -32768}
                fmt.Println(fi128, fi129)
                _ = fi129
                _ = v131
                _ = v132
            }
            v133 := GenS4{F0: p0, F1: 32767}
            _ = fi128
            _ = v133
        }
        _ = v126
    }
    p0 = (p0 >> ((1) & 63))
    v134 := []int{127, -60, 9223372036854775807}
    _ = p0
    _ = v134
    return 4294967295
}

func main() {
    var v1 float32 = -2.25
    var v2 uint = 0
    var v3 float64 = ((0.0 / (1.5 + (float64(float32((1.5 - float64(float32(0.0))))) - 3.14159))) - 0.0)
    var v4 string = "hello"
    v5 := []string{}
    v1 = (v1 / (v1 - (3.14159 - float32(float64(v1)))))
    var v135 int64 = 9223372036854775807
    if (false || true) {
        fmt.Println(v135, v4, v135)
        if (2147483647 < 712) {
            switch "a" {
            case "foo bar":
                if false {
                    for fi136 := 0; fi136 < 20; fi136++ {
                        for fi137 := 0; fi137 < 9; fi137++ {
                            if (4294967295 > 972) {
                                switch v4 {
                                case "a":
                                    v138 := GenS2{F0: 103, F1: 42}
                                    _ = v138
                                }
                                fmt.Println(v1, v4, v2)
                                v4 = ""
                            } else {
                                switch fi136 {
                                case 127:
                                    var v139 int = (fi137 & (fi136 * fi137))
                                    _ = v139
                                case 42:
                                    v140 := GenS5{F0: 99, F1: 113}
                                    _ = v140
                                default:
                                    var v141 uint = v2
                                    _ = v141
                                }
                            }
                            _ = fi137
                        }
                        var v142 uint16 = 2
                        var v143 bool = false
                        _ = fi136
                        _ = v142
                        _ = v143
                    }
                } else {
                    var v144 uint64 = 1
                    var v145 bool = (-563 <= -623)
                    _ = v144
                    _ = v145
                }
            case "\n\t":
                switch 2147483647 {
                case -1:
                    for fi146 := 0; fi146 < 13; fi146++ {
                        if false {
                            var v147 uint64 = 54
                            var v148 int8 = 32
                            if (uint16(v2) <= 2) {
                                v2 = (v2 ^ 18446744073709551615)
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                if false {
                                    for fi149 := 0; fi149 < 8; fi149++ {
                                        v135 += v135
                                        v150 := GenS2{F0: 211, F1: v147}
                                        v1 += v1
                                        fmt.Println(v4, fi149, fi149)
                                        v1 = v1
                                        _ = fi149
                                        _ = v150
                                    }
                                }
                            }
                            _ = v147
                            _ = v148
                        } else {
                            v1 = (v1 / (float32(v3) - float32(v3)))
                            v151 := GenS1{F0: 1, F1: 2147483647}
                            _ = v151
                        }
                        var v152 uint8 = 255
                        _ = fi146
                        _ = v152
                    }
                }
            }
            var v153 uint64 = 18446744073709551615
            _ = v153
        }
    }
    for fi154 := 0; fi154 < 17; fi154++ {
        v155 := GenS6{F0: 0, F1: 42}
        if len(v5) > 0 {
            v5[0] = "hello"
        }
        for fi156 := 0; fi156 < 11; fi156++ {
            var v157 int64 = -1
            _ = fi156
            _ = v157
        }
        _ = fi154
        _ = v155
    }
    for fi158 := 0; fi158 < 17; fi158++ {
        if true {
            var v159 uint8 = 42
            for fi160 := 0; fi160 < 9; fi160++ {
                v161 := GenS6{F0: 0, F1: uint32(v159)}
                var v162 string = "\n\t"
                var v163 bool = true
                _ = fi160
                _ = v161
                _ = v162
                _ = v163
            }
            if true {
                for fi164 := 0; fi164 < 8; fi164++ {
                    if (false && true) {
                        var v165 uint16 = 974
                        v166 := GenS9{F0: 127, F1: (v2 / ((813) | 1)), F2: v159}
                        v167 := GenS15{F0: (v165 ^ (v165 + 70)), F1: 42, F2: -110}
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = v165
                        _ = v166
                        _ = v167
                    } else {
                        v159 += 0
                        var v168 string = v4
                        _ = v168
                    }
                    _ = fi164
                }
                v159 -= v159
                v169 := GenS14{F0: v4}
                _ = v169
            }
            _ = v159
        } else {
            v170 := GenS7{F0: 2147483647, F1: 2, F2: 0}
            for fi171 := 0; fi171 < 7; fi171++ {
                v2 = v2
                _ = fi171
            }
            fmt.Println(v3)
            _ = v170
        }
        _ = fi158
    }
    fmt.Println(v1, v1)
    v135 -= v135
    var v172 int64 = v135
    var v173 bool = true
    var v174 uint8 = 127
    v173 = v173
    v175 := GenS11{F0: (^(v2))}
    fmt.Println("v1:", v1)
    fmt.Println("v135:", v135)
    fmt.Println("v172:", v172)
    fmt.Println("v173:", v173)
    fmt.Println("v174:", v174)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v135
    _ = v172
    _ = v173
    _ = v174
    _ = v175
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
