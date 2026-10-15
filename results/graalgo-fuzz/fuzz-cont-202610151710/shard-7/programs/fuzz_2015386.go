package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32; F1 int64; F2 string }
type GenS4 struct { F0 int; F1 int8; F2 int32 }
type GenS5 struct { F0 bool; F1 uint16; F2 int16 }
type GenS6 struct { F0 uint32; F1 uint64; F2 int8 }
type GenS7 struct { F0 uint8; F1 uint64; F2 bool }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 int; F1 int8; F2 int64 }
type GenS10 struct { F0 bool; F1 bool; F2 int }
type GenS11 struct { F0 int8; F1 uint8; F2 int8 }
type GenS12 struct { F0 bool; F1 uint }
type GenS13 struct { F0 uint }
type GenS14 struct { F0 int; F1 uint64; F2 bool }
type GenS15 struct { F0 uint16; F1 uint8; F2 uint64 }
type GenS16 struct { F0 string }
type GenS17 struct { F0 int8; F1 uint64 }
type GenS18 struct { F0 string; F1 string; F2 int32 }
type GenS19 struct { F0 uint }
type GenS20 struct { F0 uint32 }
type GenS21 struct { F0 string }
type GenS22 struct { F0 int64; F1 int8; F2 int8 }
type GenS23 struct { F0 uint64; F1 uint32 }
type GenS24 struct { F0 int; F1 uint16 }
type GenS25 struct { F0 int; F1 uint8; F2 int64 }
func genFn0() int64 {
    var v7 uint = 255
    v8 := GenS1{F0: 0}
    v9 := []int{127, 9223372036854775807, 1, 4294967295}
    _ = v7
    _ = v8
    _ = v9
    return 581
}
func genFn1() int {
    switch 286 {
    case -495:
        var v10 bool = (true && (!(-9223372036854775808 >= 32768)))
        _ = v10
    default:
        switch 2147483647 {
        case -2147483648:
            switch -1 {
            case 32767:
                for fi11 := 0; fi11 < 9; fi11++ {
                    v12 := GenS1{F0: -708}
                    var v13 string = ("" + ("\n\t" + ""))
                    v13 = v13
                    _ = fi11
                    _ = v12
                    _ = v13
                }
            case 4294967295:
                v14 := GenS2{F0: 522}
                _ = v14
            case -1:
                if (3.14159 >= 0.0) {
                    v15 := GenS2{F0: 28}
                    if (1 == 1) {
                        var v16 bool = (2147483647 < 1)
                        for fi17 := 0; fi17 < 1; fi17++ {
                            v18 := []int{}
                            for fi19 := 0; fi19 < 19; fi19++ {
                                v20 := GenS3{F0: 42, F1: -9223372036854775808, F2: "hello"}
                                v21 := GenS4{F0: (fi17 - (fi19 * (fi19 >> ((65535) & 63)))), F1: 100, F2: 650}
                                if v16 {
                                    var v22 int = (fi19 / (((fi19 << ((65535) & 63))) | 1))
                                    v16 = (4294967295 != 882)
                                    _ = v22
                                } else {
                                    var v23 uint8 = 57
                                    if (!(!true)) {
                                        if true {
                                            if len(v18) >= 2 {
                                                v18 = v18[1:len(v18)]
                                            }
                                            var v24 string = ("hello" + "a")
                                            switch (-(fi17)) {
                                            case -506:
                                                fmt.Println(v23, v16)
                                            case 2147483647:
                                                var v25 int16 = 443
                                                _ = v25
                                            }
                                            _ = v24
                                        }
                                        switch "" {
                                        case "":
                                            if (295 < 42) {
                                                var v26 int8 = 76
                                                _ = v26
                                            } else {
                                                fmt.Println(v23, fi17, v16)
                                                switch "\n\t" {
                                                case "\n\t":
                                                    v16 = v16
                                                default:
                                                    v18 = append(v18, (fi19 % ((645) | 1)))
                                                }
                                            }
                                            v27 := GenS2{F0: 65535}
                                            v28 := GenS1{F0: 0}
                                            var v29 uint64 = 517
                                            v29 = v29
                                            _ = v27
                                            _ = v28
                                            _ = v29
                                        case "\n\t":
                                            var v30 uint64 = 2
                                            _ = v30
                                        case "a":
                                            v18 = append(v18, -179)
                                        default:
                                            var v31 float32 = 1e10
                                            _ = v31
                                        }
                                    }
                                    fmt.Println(fi19, v16, v23)
                                    _ = v23
                                }
                                fmt.Println(fi17, fi19)
                                if ((true && (-24 > -39)) || (462 < 18446744073709551615)) {
                                    fmt.Println(fi17, fi19, fi19)
                                    if (1e10 <= 1.5) {
                                        if (!((-207 > (fi17 | (^(fi17)))) && v16)) {
                                            if (243 > 128) {
                                                if v16 {
                                                    v16 = (1 >= 555)
                                                    switch ("foo bar" + (("a" + ("foo bar" + "hello")) + "\n\t")) {
                                                    case "":
                                                        v32 := GenS5{F0: (582 <= 1), F1: 255, F2: 0}
                                                        _ = v32
                                                    case "hello":
                                                        var v33 uint32 = 824
                                                        _ = v33
                                                    case "a":
                                                        v16 = false
                                                    default:
                                                        if false {
                                                            v16 = v16
                                                            v16 = (true && (1.5 >= 3.14159))
                                                            var v34 uint = 255
                                                            _ = v34
                                                        } else {
                                                            v16 = true
                                                        }
                                                    }
                                                    v35 := GenS6{F0: 2147483647, F1: 444, F2: 1}
                                                    _ = v35
                                                }
                                            } else {
                                                switch (("foo bar" + ("foo bar" + ((("" + "\n\t") + "hello") + ("" + "hello")))) + ("\n\t" + "hello")) {
                                                case "a":
                                                    if (!v16) {
                                                        v16 = (!(v16 || true))
                                                        if v16 {
                                                            if (-62 > 9223372036854775807) {
                                                                v16 = (!v16)
                                                                fmt.Println(fi17, fi19)
                                                                if v16 {
                                                                    if v16 {
                                                                        v36 := GenS3{F0: 0, F1: 903, F2: ("hello" + (("hello" + "hello") + ("" + ("foo bar" + ("foo bar" + "foo bar")))))}
                                                                        v16 = (0 <= 708)
                                                                        _ = v36
                                                                    } else {
                                                                        if true {
                                                                            fmt.Println(fi17)
                                                                        } else {
                                                                            var v37 uint8 = 217
                                                                            var v38 uint32 = 65535
                                                                            v39 := GenS7{F0: (^(v37)), F1: 157, F2: (false || (fi19 != (fi19 & (-(fi17)))))}
                                                                            v38 = v38
                                                                            _ = v37
                                                                            _ = v38
                                                                            _ = v39
                                                                        }
                                                                        v16 = (v16 && (v16 && v16))
                                                                    }
                                                                    var v40 float32 = 0.1
                                                                    if (!(((v16 || v16) && v16) && (809 <= 721))) {
                                                                        var v41 float64 = ((0.0 * (float64(0.0) * (((((1e10 * 1e10) + -2.25) / 3.14159) - 1.5) * 3.14159))) / (0.1 / (1e10 * 1e10)))
                                                                        if false {
                                                                            if len(v18) >= 2 {
                                                                                v18 = v18[1:len(v18)]
                                                                            }
                                                                        }
                                                                        _ = v41
                                                                    }
                                                                    _ = v40
                                                                }
                                                            } else {
                                                                switch (fi17 | fi19) {
                                                                case 128:
                                                                    v18 = append(v18, fi19)
                                                                case 42:
                                                                    v16 = (1 > 127)
                                                                default:
                                                                    if (0 >= -9223372036854775808) {
                                                                        if v16 {
                                                                            v16 = false
                                                                            fmt.Println(fi17, fi17)
                                                                        } else {
                                                                            v42 := GenS7{F0: 2, F1: 2, F2: (!false)}
                                                                            var v43 uint32 = 600
                                                                            _ = v42
                                                                            _ = v43
                                                                        }
                                                                        switch "" {
                                                                        case "":
                                                                            switch fi19 {
                                                                            case 4294967295:
                                                                                v16 = true
                                                                                v16 = ((^(fi19)) >= (fi19 % (((^(fi17))) | 1)))
                                                                                v16 = v16
                                                                            case -876:
                                                                                var v44 string = ""
                                                                                _ = v44
                                                                            case 2147483647:
                                                                                fmt.Println(fi17, fi19)
                                                                                var v45 string = "hello"
                                                                                var v46 float64 = 3.14159
                                                                                _ = v45
                                                                                _ = v46
                                                                            }
                                                                        case "hello":
                                                                            fmt.Println(fi17, fi17, fi19)
                                                                        }
                                                                        var v47 int16 = 1000
                                                                        _ = v47
                                                                    } else {
                                                                        v48 := GenS1{F0: 0}
                                                                        fmt.Println(fi19)
                                                                        var v49 float64 = 1.5
                                                                        _ = v48
                                                                        _ = v49
                                                                    }
                                                                }
                                                                var v50 int16 = -820
                                                                var v51 uint16 = 32767
                                                                switch (fi19 % ((1) | 1)) {
                                                                case 9223372036854775807:
                                                                    v52 := GenS8{F0: 1}
                                                                    var v53 int32 = 1
                                                                    _ = v52
                                                                    _ = v53
                                                                case 32768:
                                                                    v50 -= v50
                                                                case 0:
                                                                    v54 := GenS3{F0: 42, F1: 9223372036854775807, F2: ("" + ("foo bar" + "a"))}
                                                                    _ = v54
                                                                }
                                                                v55 := GenS6{F0: uint32(v51), F1: 1, F2: -67}
                                                                _ = v50
                                                                _ = v51
                                                                _ = v55
                                                            }
                                                            if len(v18) > 0 {
                                                                v18[0] = fi17
                                                            }
                                                        }
                                                        if len(v18) >= 2 {
                                                            v18 = v18[1:len(v18)]
                                                        }
                                                    } else {
                                                        switch fi17 {
                                                        case 2147483647:
                                                            if false {
                                                                if (!v16) {
                                                                    var v56 int16 = -898
                                                                    v16 = false
                                                                    v56 += 475
                                                                    _ = v56
                                                                } else {
                                                                    v57 := GenS6{F0: 0, F1: 836, F2: 42}
                                                                    var v58 bool = false
                                                                    fmt.Println(fi17)
                                                                    if true {
                                                                        if len(v18) >= 2 {
                                                                            v18 = v18[1:len(v18)]
                                                                        }
                                                                        if ((fi19 ^ fi19) >= 142) {
                                                                            v59 := GenS2{F0: 65535}
                                                                            v60 := GenS1{F0: genFn0()}
                                                                            v58 = false
                                                                            if v58 {
                                                                                v16 = (false || v58)
                                                                                if true {
                                                                                    if v16 {
                                                                                        switch fi19 {
                                                                                        case 42:
                                                                                            if v58 {
                                                                                                v16 = v58
                                                                                                switch (fi17 ^ (fi19 + (fi17 << ((255) & 63)))) {
                                                                                                case -647:
                                                                                                    v61 := GenS3{F0: 42, F1: genFn0(), F2: "\n\t"}
                                                                                                    _ = v61
                                                                                                case 4294967295:
                                                                                                    var v62 float64 = 3.14159
                                                                                                    _ = v62
                                                                                                default:
                                                                                                    if (-2147483648 > 482) {
                                                                                                        if false {
                                                                                                            var v63 int8 = 7
                                                                                                            _ = v63
                                                                                                        } else {
                                                                                                            switch fi17 {
                                                                                                            case -2147483648:
                                                                                                                v64 := GenS3{F0: 2147483648, F1: 1, F2: ""}
                                                                                                                _ = v64
                                                                                                            default:
                                                                                                                _ = len(v18)
                                                                                                            }
                                                                                                            if false {
                                                                                                                switch 32768 {
                                                                                                                case 0:
                                                                                                                    v16 = v16
                                                                                                                case 168:
                                                                                                                    var v65 int32 = -2147483648
                                                                                                                    _ = v65
                                                                                                                default:
                                                                                                                    if v16 {
                                                                                                                        v66 := GenS9{F0: fi19, F1: 100, F2: genFn0()}
                                                                                                                        if v58 {
                                                                                                                            var v67 string = "\n\t"
                                                                                                                            fmt.Println(v58, v16, v16)
                                                                                                                            _ = v67
                                                                                                                        }
                                                                                                                        _ = v66
                                                                                                                    } else {
                                                                                                                        var v68 bool = false
                                                                                                                        var v69 int = fi17
                                                                                                                        if v68 {
                                                                                                                            fmt.Println(fi19, fi17, v58)
                                                                                                                            var v70 int64 = -1
                                                                                                                            v71 := GenS10{F0: (!(!(33 != 0))), F1: v68, F2: v69}
                                                                                                                            _ = v70
                                                                                                                            _ = v71
                                                                                                                        }
                                                                                                                        _ = v68
                                                                                                                        _ = v69
                                                                                                                    }
                                                                                                                }
                                                                                                                if (v16 && v58) {
                                                                                                                    v72 := GenS1{F0: -557}
                                                                                                                    var v73 float32 = 3.14159
                                                                                                                    _ = v72
                                                                                                                    _ = v73
                                                                                                                }
                                                                                                            }
                                                                                                            if ((-2.25 * 0.1) != 1e10) {
                                                                                                                v74 := GenS1{F0: genFn0()}
                                                                                                                v58 = ((-1 > 1000) && (!(198 >= 127)))
                                                                                                                var v75 uint32 = 0
                                                                                                                _ = v74
                                                                                                                _ = v75
                                                                                                            } else {
                                                                                                                var v76 uint16 = 255
                                                                                                                if v58 {
                                                                                                                    v58 = false
                                                                                                                    var v77 int32 = -869
                                                                                                                    _ = v77
                                                                                                                } else {
                                                                                                                    v58 = v58
                                                                                                                    v78 := GenS4{F0: fi17, F1: 0, F2: -2147483648}
                                                                                                                    var v79 uint32 = 65535
                                                                                                                    _ = v78
                                                                                                                    _ = v79
                                                                                                                }
                                                                                                                _ = v76
                                                                                                            }
                                                                                                        }
                                                                                                    } else {
                                                                                                        v16 = v58
                                                                                                    }
                                                                                                }
                                                                                            }
                                                                                        case 2147483647:
                                                                                            v80 := GenS3{F0: 1, F1: -1, F2: "a"}
                                                                                            _ = v80
                                                                                        }
                                                                                        switch fi17 {
                                                                                        case 167:
                                                                                            if (-2.25 != 0.1) {
                                                                                                var v81 int8 = 0
                                                                                                v18 = append(v18, -9223372036854775808)
                                                                                                var v82 int8 = (v81 ^ v81)
                                                                                                v83 := GenS11{F0: (v81 ^ v81), F1: 2, F2: v82}
                                                                                                v84 := GenS6{F0: 2147483648, F1: 873, F2: 24}
                                                                                                _ = v81
                                                                                                _ = v82
                                                                                                _ = v83
                                                                                                _ = v84
                                                                                            }
                                                                                        case 1:
                                                                                            var v85 int32 = 1
                                                                                            _ = v85
                                                                                        case 79:
                                                                                            var v86 float64 = 1e10
                                                                                            _ = v86
                                                                                        default:
                                                                                            v18 = append(v18, fi17)
                                                                                        }
                                                                                        switch fi19 {
                                                                                        case 127:
                                                                                            switch 127 {
                                                                                            case 0:
                                                                                                fmt.Println(v16, fi19, fi17)
                                                                                            case 128:
                                                                                                var v87 int32 = 42
                                                                                                _ = v87
                                                                                            case 42:
                                                                                                if (42 < 42) {
                                                                                                    var v88 string = "foo bar"
                                                                                                    if len(v18) > 0 {
                                                                                                        v18[0] = fi17
                                                                                                    }
                                                                                                    v58 = true
                                                                                                    _ = v88
                                                                                                } else {
                                                                                                    var v89 string = "foo bar"
                                                                                                    fmt.Println(v89)
                                                                                                    v90 := GenS10{F0: ((genFn0() == genFn0()) && (((v89 + (("hello" + (v89 + v89)) + "foo bar")) >= "\n\t") || true)), F1: true, F2: fi17}
                                                                                                    _ = v89
                                                                                                    _ = v90
                                                                                                }
                                                                                            }
                                                                                        default:
                                                                                            v58 = v16
                                                                                        }
                                                                                    }
                                                                                    if len(v18) > 0 {
                                                                                        v18[0] = -2147483648
                                                                                    }
                                                                                } else {
                                                                                    switch (fi19 << ((65535) & 63)) {
                                                                                    case 32768:
                                                                                        if len(v18) >= 2 {
                                                                                            v18 = v18[1:len(v18)]
                                                                                        }
                                                                                    case 366:
                                                                                        var v91 uint16 = 963
                                                                                        _ = v91
                                                                                    case 4294967295:
                                                                                        var v92 int8 = 27
                                                                                        _ = v92
                                                                                    default:
                                                                                        fmt.Println(fi17, v58)
                                                                                    }
                                                                                    var v93 bool = (137 == 2)
                                                                                    v94 := GenS12{F0: v58, F1: 1}
                                                                                    _ = v93
                                                                                    _ = v94
                                                                                }
                                                                            }
                                                                            _ = v59
                                                                            _ = v60
                                                                        } else {
                                                                            switch "foo bar" {
                                                                            case "a":
                                                                                switch (fi17 / (((fi17 * 9223372036854775807)) | 1)) {
                                                                                case 966:
                                                                                    if ((347 != (fi17 << ((18446744073709551615) & 63))) && v58) {
                                                                                        v16 = (255 <= 2)
                                                                                    }
                                                                                case 1:
                                                                                    var v95 uint64 = 2
                                                                                    _ = v95
                                                                                case 4294967295:
                                                                                    var v96 int16 = 804
                                                                                    _ = v96
                                                                                }
                                                                            case "":
                                                                                v16 = true
                                                                            case "\n\t":
                                                                                if false {
                                                                                    if v16 {
                                                                                        if v16 {
                                                                                            if len(v18) > 0 {
                                                                                                v18[0] = fi17
                                                                                            }
                                                                                            v16 = false
                                                                                        }
                                                                                        var v97 uint64 = 98
                                                                                        _ = v97
                                                                                    } else {
                                                                                        var v98 bool = (v16 && v58)
                                                                                        if (!((((3.14159 + 3.14159) / 3.14159) / (0.1 / (-2.25 - (float64((0.1 - ((float32(3.14159) / 0.1) - 0.0))) * (1e10 * (1e10 * 0.0)))))) >= -2.25)) {
                                                                                            var v99 float32 = 0.1
                                                                                            v58 = (v58 && false)
                                                                                            var v100 float64 = (0.1 - float64((0.0 / (v99 + v99))))
                                                                                            _ = v100
                                                                                            _ = v99
                                                                                        } else {
                                                                                            switch fi17 {
                                                                                            case -2147483648:
                                                                                                if v58 {
                                                                                                    var v101 int = 127
                                                                                                    _ = v101
                                                                                                }
                                                                                            }
                                                                                            var v102 uint16 = 42
                                                                                            _ = v102
                                                                                        }
                                                                                        _ = v98
                                                                                    }
                                                                                }
                                                                            default:
                                                                                if (!(1 == -32768)) {
                                                                                    var v103 int64 = 9223372036854775807
                                                                                    var v104 int16 = -32768
                                                                                    _ = v103
                                                                                    _ = v104
                                                                                } else {
                                                                                    v105 := GenS12{F0: (!v58), F1: 18446744073709551615}
                                                                                    _ = v105
                                                                                }
                                                                                var v106 string = ("a" + "hello")
                                                                                var v107 uint8 = 2
                                                                                _ = v106
                                                                                _ = v107
                                                                            }
                                                                            var v108 uint16 = 1
                                                                            _ = v108
                                                                        }
                                                                        v109 := GenS6{F0: 996, F1: 4294967295, F2: -1}
                                                                        _ = v109
                                                                    } else {
                                                                        switch "hello" {
                                                                        case "foo bar":
                                                                            v16 = v16
                                                                        case "":
                                                                            v16 = v16
                                                                        case "a":
                                                                            v16 = v58
                                                                        }
                                                                        var v110 int64 = -1
                                                                        if (((911 > 65535) || (1e10 <= (3.14159 - -2.25))) || (!("foo bar" >= "\n\t"))) {
                                                                            v16 = v58
                                                                        }
                                                                        _ = v110
                                                                    }
                                                                    var v111 float32 = -2.25
                                                                    _ = v111
                                                                    _ = v57
                                                                    _ = v58
                                                                }
                                                                fmt.Println(fi17)
                                                            } else {
                                                                switch ("foo bar" + ("hello" + "")) {
                                                                case "hello":
                                                                    v16 = v16
                                                                    v16 = (0.1 < (0.1 * (-2.25 / (((float64(1e10) / 1e10) * 1.5) - 0.1))))
                                                                case "a":
                                                                    v112 := GenS13{F0: 811}
                                                                    _ = v112
                                                                }
                                                            }
                                                        case 127:
                                                            v16 = v16
                                                        case -596:
                                                            var v113 int16 = -32768
                                                            _ = v113
                                                        default:
                                                            if (4294967295 > fi17) {
                                                                switch (fi19 - fi17) {
                                                                case 32767:
                                                                    if len(v18) >= 2 {
                                                                        v18 = v18[1:len(v18)]
                                                                    }
                                                                }
                                                                v114 := GenS2{F0: 42}
                                                                _ = v114
                                                            } else {
                                                                var v115 int16 = -32768
                                                                if false {
                                                                    var v116 uint64 = 4294967295
                                                                    if (4294967295 != (^(v116))) {
                                                                        v115 = v115
                                                                        switch fi17 {
                                                                        case -887:
                                                                            v117 := GenS2{F0: 406}
                                                                            _ = v117
                                                                        default:
                                                                            var v118 float32 = 1e10
                                                                            _ = v118
                                                                        }
                                                                        switch fi19 {
                                                                        case 1:
                                                                            fmt.Println(v115, v116, v116)
                                                                        }
                                                                    }
                                                                    var v119 int64 = -1
                                                                    _ = v116
                                                                    _ = v119
                                                                } else {
                                                                    fmt.Println(fi17, fi17)
                                                                    v16 = (v16 || v16)
                                                                    v18 = append(v18, (fi19 >> ((390) & 63)))
                                                                }
                                                                _ = v115
                                                            }
                                                        }
                                                        v16 = v16
                                                        switch "" {
                                                        case "\n\t":
                                                            v16 = v16
                                                        default:
                                                            fmt.Println(fi19)
                                                        }
                                                    }
                                                case "":
                                                    var v120 uint = 923
                                                    _ = v120
                                                }
                                                if len(v18) > 0 {
                                                    v18[0] = 9223372036854775807
                                                }
                                                var v121 bool = ((0 <= -889) || (1e10 > 1.5))
                                                if true {
                                                    fmt.Println(fi19, fi17, fi17)
                                                    fmt.Println(v121, fi19)
                                                    if v121 {
                                                        switch "hello" {
                                                        case "":
                                                            switch "\n\t" {
                                                            case "\n\t":
                                                                if false {
                                                                    v122 := GenS14{F0: fi19, F1: 4294967295, F2: true}
                                                                    fmt.Println(v121, fi19, v121)
                                                                    _ = v122
                                                                }
                                                            case "foo bar":
                                                                var v123 string = (("a" + "\n\t") + ((("" + (("" + "a") + (("foo bar" + "") + "hello"))) + "hello") + ("" + ("foo bar" + (("\n\t" + "foo bar") + "foo bar")))))
                                                                _ = v123
                                                            }
                                                        default:
                                                            v121 = (-92 == -113)
                                                        }
                                                    }
                                                    v18 = append(v18, fi19)
                                                    v124 := GenS14{F0: fi19, F1: 42, F2: v121}
                                                    _ = v124
                                                }
                                                _ = v121
                                            }
                                            _ = len(v18)
                                        }
                                        switch "a" {
                                        case "hello":
                                            var v125 int8 = 69
                                            _ = v125
                                        case "foo bar":
                                            v126 := GenS15{F0: 0, F1: 42, F2: 1}
                                            _ = v126
                                        }
                                        v16 = true
                                    } else {
                                        v16 = (!(v16 && v16))
                                        var v127 uint8 = 1
                                        v16 = false
                                        var v128 uint64 = 1
                                        if len(v18) > 0 {
                                            v18[0] = -969
                                        }
                                        _ = v127
                                        _ = v128
                                    }
                                    v129 := GenS10{F0: v16, F1: v16, F2: 1}
                                    var v130 float32 = -2.25
                                    if (!v16) {
                                        if v16 {
                                            if (3.14159 < ((1.5 - v130) * v130)) {
                                                if v16 {
                                                    fmt.Println(v130)
                                                    _ = len(v18)
                                                    var v131 int32 = -2147483648
                                                    _ = v131
                                                } else {
                                                    v132 := GenS16{F0: "a"}
                                                    v130 = (((v130 + (1.5 * 3.14159)) - 1e10) - v130)
                                                    fmt.Println(v16)
                                                    _ = v132
                                                }
                                                switch fi19 {
                                                case 9223372036854775807:
                                                    if (0 < 1) {
                                                        if v16 {
                                                            var v133 int64 = 9223372036854775807
                                                            _ = v133
                                                        }
                                                    } else {
                                                        var v134 int8 = -82
                                                        v134 = 127
                                                        v134 = 1
                                                        _ = v134
                                                    }
                                                }
                                                v16 = (-1 != -1)
                                                fmt.Println(fi17, fi17, fi19)
                                            }
                                            v135 := GenS3{F0: 975, F1: 1, F2: "a"}
                                            if v16 {
                                                v16 = false
                                                if false {
                                                    switch (-(fi17)) {
                                                    case 9223372036854775807:
                                                        var v136 int = fi19
                                                        _ = v136
                                                    case 42:
                                                        fmt.Println(fi19, fi19, v130)
                                                    case 32767:
                                                        fmt.Println(fi19, v16, v16)
                                                    default:
                                                        var v137 uint = 18446744073709551615
                                                        _ = v137
                                                    }
                                                    var v138 uint8 = 1
                                                    if (v16 && false) {
                                                        v139 := GenS10{F0: false, F1: false, F2: fi17}
                                                        _ = v139
                                                    }
                                                    _ = v138
                                                } else {
                                                    var v140 float32 = (v130 + v130)
                                                    _ = v140
                                                }
                                            } else {
                                                var v141 float64 = 0.1
                                                v142 := GenS11{F0: 100, F1: 178, F2: 0}
                                                _ = v141
                                                _ = v142
                                            }
                                            var v143 float64 = 1e10
                                            var v144 int = (fi19 % (((fi19 & -740)) | 1))
                                            _ = v135
                                            _ = v143
                                            _ = v144
                                        }
                                        v130 += 0.0
                                        fmt.Println(fi19, fi17)
                                    }
                                    _ = v129
                                    _ = v130
                                }
                                _ = fi19
                                _ = v20
                                _ = v21
                            }
                            v16 = true
                            _ = fi17
                            _ = v18
                        }
                        if false {
                            for fi145 := 0; fi145 < 5; fi145++ {
                                var v146 int16 = 480
                                _ = fi145
                                _ = v146
                            }
                        } else {
                            v16 = v16
                            for fi147 := 0; fi147 < 9; fi147++ {
                                v148 := GenS17{F0: -128, F1: 228}
                                if (0 < genFn0()) {
                                    for fi149 := 0; fi149 < 16; fi149++ {
                                        v150 := []int{(fi149 - fi147), (fi149 >> ((1) & 63))}
                                        var v151 uint8 = 127
                                        switch (fi147 % ((42) | 1)) {
                                        case 9223372036854775807:
                                            if v16 {
                                                if true {
                                                    var v152 float32 = (0.1 * 3.14159)
                                                    _ = v152
                                                }
                                                v16 = false
                                            } else {
                                                v16 = v16
                                                if false {
                                                    switch fi147 {
                                                    case 2147483647:
                                                        v153 := GenS18{F0: "hello", F1: "", F2: -875}
                                                        _ = v153
                                                    default:
                                                        if (-9223372036854775808 <= 1) {
                                                            var v154 uint16 = 255
                                                            if (2147483647 != 2) {
                                                                fmt.Println(v151, v151, v16)
                                                                if v16 {
                                                                    fmt.Println(v154)
                                                                    _ = len(v150)
                                                                }
                                                                var v155 uint64 = 964
                                                                v150 = append(v150, fi147)
                                                                if len(v150) > 0 {
                                                                    v150[0] = fi149
                                                                }
                                                                v154 = 0
                                                                _ = v155
                                                            }
                                                            v16 = (uint64(v154) == 1)
                                                            _ = v154
                                                        } else {
                                                            switch "" {
                                                            case "":
                                                                var v156 string = "a"
                                                                _ = v156
                                                            }
                                                            v16 = (!(!(255 < 900)))
                                                            v151 = 201
                                                            if (-2147483648 != 42) {
                                                                if (0 != 2) {
                                                                    var v157 int32 = 42
                                                                    _ = v157
                                                                } else {
                                                                    v151 = 42
                                                                    if len(v150) > 0 {
                                                                        v150[0] = -9223372036854775808
                                                                    }
                                                                    v16 = (1000 <= 1000)
                                                                }
                                                                if true {
                                                                    switch -308 {
                                                                    case -2147483648:
                                                                        var v158 int16 = 0
                                                                        _ = v158
                                                                    }
                                                                    var v159 int8 = 0
                                                                    if ((255 <= 255) || ((v151 + 255) >= (v151 << ((18446744073709551615) & 63)))) {
                                                                        switch fi147 {
                                                                        case 32767:
                                                                            v160 := GenS1{F0: -1}
                                                                            _ = v160
                                                                        default:
                                                                            if len(v150) > 0 {
                                                                                v150[0] = (fi147 >> ((4294967295) & 63))
                                                                            }
                                                                        }
                                                                    } else {
                                                                        if len(v150) >= 2 {
                                                                            v150 = v150[1:len(v150)]
                                                                        }
                                                                        fmt.Println(fi149, fi147, v159)
                                                                    }
                                                                    _ = v159
                                                                } else {
                                                                    v16 = true
                                                                    v161 := GenS7{F0: 0, F1: 580, F2: (2 < 131)}
                                                                    v151 += 2
                                                                    _ = v161
                                                                }
                                                                fmt.Println(fi149, v16)
                                                            } else {
                                                                var v162 uint = 1
                                                                _ = v162
                                                            }
                                                        }
                                                        fmt.Println(fi149, fi147)
                                                    }
                                                    v16 = v16
                                                    _ = len(v150)
                                                } else {
                                                    if true {
                                                        fmt.Println(fi149)
                                                        v163 := GenS11{F0: 117, F1: v151, F2: 59}
                                                        _ = v163
                                                    }
                                                }
                                                fmt.Println(v151)
                                            }
                                        case 0:
                                            fmt.Println(v16)
                                        default:
                                            v164 := GenS19{F0: 102}
                                            _ = v164
                                        }
                                        _ = fi149
                                        _ = v150
                                        _ = v151
                                    }
                                    if (18446744073709551615 != 161) {
                                        if v16 {
                                            if true {
                                                for fi165 := 0; fi165 < 20; fi165++ {
                                                    if true {
                                                        fmt.Println(v16)
                                                        if (v16 && (139 > 127)) {
                                                            switch (("a" + ("hello" + "foo bar")) + "foo bar") {
                                                            case "foo bar":
                                                                var v166 int64 = 0
                                                                _ = v166
                                                            case "hello":
                                                                v16 = v16
                                                                v167 := []int{}
                                                                if ((!v16) && v16) {
                                                                    switch "" {
                                                                    case "":
                                                                        v168 := GenS20{F0: 843}
                                                                        _ = v168
                                                                    case "hello":
                                                                        v169 := GenS20{F0: 866}
                                                                        _ = v169
                                                                    default:
                                                                        fmt.Println(v16)
                                                                    }
                                                                }
                                                                _ = v167
                                                            }
                                                            v16 = v16
                                                            switch "" {
                                                            case "\n\t":
                                                                v16 = (!true)
                                                            }
                                                        }
                                                    } else {
                                                        if v16 {
                                                            v170 := []int{(fi147 - -1), fi165, (fi147 & fi147)}
                                                            switch "hello" {
                                                            case "a":
                                                                v171 := GenS2{F0: 0}
                                                                _ = v171
                                                            case "":
                                                                v16 = true
                                                            }
                                                            if v16 {
                                                                v16 = (520 >= 1000)
                                                                switch (("a" + "a") + ("hello" + (((("foo bar" + ("" + ("\n\t" + "a"))) + "a") + "") + "a"))) {
                                                                case "hello":
                                                                    v172 := GenS5{F0: (-898 >= 32767), F1: 900, F2: -1}
                                                                    _ = v172
                                                                }
                                                            }
                                                            v16 = (v16 && true)
                                                            _ = v170
                                                        }
                                                        v16 = (true || v16)
                                                        var v173 string = ("\n\t" + "a")
                                                        _ = v173
                                                    }
                                                    v174 := GenS21{F0: ("a" + "hello")}
                                                    v175 := []int{(^(fi165)), 0}
                                                    _ = fi165
                                                    _ = v174
                                                    _ = v175
                                                }
                                                for fi176 := 0; fi176 < 9; fi176++ {
                                                    var v177 int = fi176
                                                    v177 -= (fi176 << ((42) & 63))
                                                    v178 := GenS4{F0: (^(fi147)), F1: -17, F2: 42}
                                                    var v179 int = (fi176 - 32767)
                                                    _ = fi176
                                                    _ = v177
                                                    _ = v178
                                                    _ = v179
                                                }
                                            }
                                            v180 := GenS18{F0: ("\n\t" + ("foo bar" + "hello")), F1: ("foo bar" + "hello"), F2: 1}
                                            for fi181 := 0; fi181 < 2; fi181++ {
                                                v16 = v16
                                                var v182 uint32 = 1
                                                fmt.Println(fi181, v182)
                                                _ = fi181
                                                _ = v182
                                            }
                                            _ = v180
                                        } else {
                                            v16 = (false && true)
                                        }
                                        v16 = v16
                                    } else {
                                        var v183 float64 = float64((float32(1e10) / 1e10))
                                        v184 := GenS17{F0: 94, F1: 0}
                                        var v185 float32 = float32((float64(0.1) / v183))
                                        _ = v183
                                        _ = v184
                                        _ = v185
                                    }
                                } else {
                                    for fi186 := 0; fi186 < 15; fi186++ {
                                        if (42 == -1) {
                                            v187 := []int{-107, (^(fi147)), fi147}
                                            _ = v187
                                        } else {
                                            v16 = true
                                            var v188 uint = 2
                                            v189 := []int{(fi186 | (fi186 ^ (fi147 | fi186)))}
                                            _ = v188
                                            _ = v189
                                        }
                                        var v190 int32 = 2147483647
                                        _ = fi186
                                        _ = v190
                                    }
                                    v16 = (205 != 932)
                                }
                                for fi191 := 0; fi191 < 16; fi191++ {
                                    v192 := GenS13{F0: 4294967295}
                                    v193 := GenS10{F0: (-9223372036854775808 > 1), F1: (!((fi191 + (fi191 + (fi147 - (fi191 % (((^(fi147))) | 1))))) < (fi147 + 128))), F2: 32768}
                                    switch "foo bar" {
                                    case "\n\t":
                                        var v194 uint32 = 42
                                        _ = v194
                                    case "":
                                        fmt.Println(v16)
                                    }
                                    v16 = true
                                    _ = fi191
                                    _ = v192
                                    _ = v193
                                }
                                _ = fi147
                                _ = v148
                            }
                            v16 = (!(1 != 0))
                        }
                        _ = v16
                    } else {
                        for fi195 := 0; fi195 < 1; fi195++ {
                            for fi196 := 0; fi196 < 17; fi196++ {
                                var v197 int32 = -905
                                _ = fi196
                                _ = v197
                            }
                            _ = fi195
                        }
                        var v198 float32 = 1.5
                        if false {
                            switch ("hello" + (("" + ("hello" + "\n\t")) + ("\n\t" + ("foo bar" + "hello")))) {
                            case "hello":
                                var v199 uint64 = 0
                                _ = v199
                            }
                            var v200 int = 4294967295
                            _ = v200
                        }
                        _ = v198
                    }
                    var v201 int16 = -1
                    _ = v15
                    _ = v201
                }
            default:
                var v202 uint = 769
                _ = v202
            }
        default:
            var v203 uint32 = 4294967295
            _ = v203
        }
    }
    var v204 uint = 2
    var v205 uint = (v204 >> ((v204) & 63))
    switch ("a" + "a") {
    case "hello":
        v204 -= v205
    case "\n\t":
        if true {
            var v206 int16 = 107
            var v207 int32 = 2147483647
            _ = v206
            _ = v207
        }
    case "":
        var v208 uint8 = 205
        _ = v208
    default:
        var v209 int32 = -1
        _ = v209
    }
    _ = v204
    _ = v205
    return -440
}

func main() {
    var v1 uint32 = 4294967295
    var v2 uint32 = 65535
    var v3 uint = 1
    var v4 uint = v3
    v5 := []string{"\n\t", "a"}
    var v6 int = 740
    if ((((-2.25 * -2.25) >= (0.0 / 3.14159)) || true) && (false || (42 == 32767))) {
        switch (v6 - v6) {
        case 2147483647:
            if (false && false) {
                for fi210 := 0; fi210 < 11; fi210++ {
                    for fi211 := 0; fi211 < 4; fi211++ {
                        v5 = append(v5, "a")
                        _ = fi211
                    }
                    _ = fi210
                }
                v4 -= (v3 << ((v4) & 63))
                if false {
                    v5 = append(v5, "")
                    for fi212 := 0; fi212 < 12; fi212++ {
                        switch 4294967295 {
                        case 9223372036854775807:
                            v213 := GenS22{F0: 1, F1: -128, F2: -60}
                            _ = v213
                        case 128:
                            v4 += (v3 >> ((2) & 63))
                        }
                        var v214 float32 = 0.0
                        if (255 > 2) {
                            v2 = 2147483648
                            v6 -= v6
                        }
                        _ = fi212
                        _ = v214
                    }
                }
            } else {
                for fi215 := 0; fi215 < 10; fi215++ {
                    var v216 int16 = 832
                    var v217 int32 = 1
                    var v218 string = "foo bar"
                    _ = fi215
                    _ = v216
                    _ = v217
                    _ = v218
                }
                var v219 uint32 = (v2 * v1)
                _ = v219
            }
        case 28:
            var v220 float32 = (0.0 * -2.25)
            _ = v220
        default:
            fmt.Println(v6, v3)
        }
        var v221 uint16 = 0
        _ = v221
    } else {
        v5 = append(v5, "\n\t")
        switch (-(v6)) {
        case 32768:
            v2 -= v1
        case 941:
            v6 += (v6 >> (((v4 | 255)) & 63))
        case 801:
            if true {
                if (false && true) {
                    var v222 uint16 = 32767
                    v222 = v222
                    _ = v222
                } else {
                    var v223 uint = v3
                    v3 += v3
                    _ = v223
                }
                if (255 < 1) {
                    if (100 != -1) {
                        var v224 string = "hello"
                        var v225 int16 = 846
                        _ = v224
                        _ = v225
                    }
                    if true {
                        if true {
                            v6 = -50
                        } else {
                            for fi226 := 0; fi226 < 1; fi226++ {
                                v1 = v1
                                _ = fi226
                            }
                            var v227 int32 = 42
                            _ = v227
                        }
                        if len(v5) > 0 {
                            v5[0] = "a"
                        }
                        v228 := GenS23{F0: 0, F1: (v1 ^ v2)}
                        _ = v228
                    } else {
                        if ((2 <= 0) || true) {
                            v229 := GenS4{F0: (v6 >> ((v3) & 63)), F1: 127, F2: -2147483648}
                            v230 := GenS24{F0: v6, F1: 692}
                            _ = v229
                            _ = v230
                        }
                        switch (-(v6)) {
                        case -9223372036854775808:
                            var v231 bool = true
                            _ = v231
                        case 516:
                            if ("a" > ((("hello" + "") + (("" + ((("foo bar" + "a") + (("hello" + "\n\t") + ("a" + ""))) + "hello")) + (("\n\t" + "hello") + ""))) + "a")) {
                                var v232 uint32 = v2
                                v3 += (^(v3))
                                _ = v232
                            }
                        case 9223372036854775807:
                            v5 = append(v5, (("\n\t" + ((("a" + "a") + "") + "\n\t")) + ("" + "")))
                        }
                        var v233 int8 = 0
                        _ = v233
                    }
                    for fi234 := 0; fi234 < 18; fi234++ {
                        fmt.Println(v4, v2, v2)
                        var v235 float64 = float64(3.14159)
                        _ = fi234
                        _ = v235
                    }
                } else {
                    if false {
                        var v236 uint32 = (v2 & v2)
                        var v237 uint64 = 4294967295
                        var v238 int32 = 586
                        _ = v236
                        _ = v237
                        _ = v238
                    } else {
                        switch genFn1() {
                        case -2147483648:
                            if true {
                                fmt.Println(v1, v3, v4)
                                var v239 float32 = 3.14159
                                _ = v239
                            }
                        case -9223372036854775808:
                            var v240 uint = (v3 ^ v4)
                            _ = v240
                        default:
                            v2 = v1
                        }
                        var v241 int16 = 42
                        v4 = 918
                        _ = v241
                    }
                    v1 = 2147483648
                    for fi242 := 0; fi242 < 10; fi242++ {
                        for fi243 := 0; fi243 < 20; fi243++ {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            _ = fi243
                        }
                        v244 := GenS8{F0: -1}
                        _ = fi242
                        _ = v244
                    }
                }
                if len(v5) > 0 {
                    v5[0] = ("\n\t" + "foo bar")
                }
            } else {
                v1 += 4294967295
                var v245 int = genFn1()
                _ = v245
            }
        default:
            v1 -= 65535
        }
    }
    v246 := GenS2{F0: 255}
    if (((true || false) && (("hello" != ("\n\t" + "\n\t")) && (843 != 2))) && true) {
        for fi247 := 0; fi247 < 6; fi247++ {
            switch "a" {
            case "hello":
                var v248 string = "hello"
                _ = v248
            case "":
                v6 = 127
            }
            v2 += v1
            _ = fi247
        }
        v6 += (v6 << ((4294967295) & 63))
        if true {
            _ = len(v5)
            v2 += 2147483647
            if (42 == 32767) {
                v1 = 437
                v1 = v2
            }
        }
    }
    for fi249 := 0; fi249 < 5; fi249++ {
        if (!(486 == v1)) {
            for fi250 := 0; fi250 < 1; fi250++ {
                switch -140 {
                case -9223372036854775808:
                    var v251 float64 = 1e10
                    _ = v251
                case 4294967295:
                    fmt.Println(v2, fi249, fi249)
                }
                _ = fi250
            }
            switch ("a" + "foo bar") {
            case "a":
                var v252 uint64 = 90
                _ = v252
            }
        } else {
            switch "hello" {
            case "foo bar":
                v253 := GenS2{F0: 65535}
                _ = v253
            }
            if ("foo bar" != "foo bar") {
                var v254 float32 = 0.0
                var v255 int8 = 0
                _ = v254
                _ = v255
            } else {
                v4 = v3
            }
        }
        v5 = append(v5, ("hello" + (("hello" + "") + ("a" + (("foo bar" + ("hello" + "hello")) + "hello")))))
        for fi256 := 0; fi256 < 18; fi256++ {
            var v257 uint16 = 65535
            _ = fi256
            _ = v257
        }
        _ = fi249
    }
    v258 := GenS5{F0: true, F1: 0, F2: 42}
    var v259 float32 = ((1e10 * (-2.25 / 0.1)) / 1.0)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println(v1)
    _ = len(v5)
    switch "a" {
    case "\n\t":
        switch "\n\t" {
        case "foo bar":
            for fi260 := 0; fi260 < 18; fi260++ {
                switch ("hello" + "hello") {
                case "":
                    if (true && false) {
                        for fi261 := 0; fi261 < 19; fi261++ {
                            var v262 int = 2147483647
                            var v263 uint = v3
                            v4 = (v4 + v4)
                            _ = fi261
                            _ = v262
                            _ = v263
                        }
                        for fi264 := 0; fi264 < 10; fi264++ {
                            v265 := GenS25{F0: fi260, F1: 42, F2: -1}
                            switch ("" + ("\n\t" + ("foo bar" + ("\n\t" + ("hello" + ((("hello" + "hello") + "a") + (("hello" + "") + "foo bar"))))))) {
                            case "hello":
                                switch (-(v6)) {
                                case -9223372036854775808:
                                    var v266 uint = v3
                                    _ = v266
                                }
                            case "foo bar":
                                v4 -= uint(v2)
                            case "a":
                                switch ("" + ("hello" + "foo bar")) {
                                case "a":
                                    var v267 bool = true
                                    _ = v267
                                case "foo bar":
                                    if (-1 == 71) {
                                        _ = len(v5)
                                        v2 = (^(v1))
                                        var v268 uint16 = 1
                                        _ = len(v5)
                                        v269 := GenS7{F0: 127, F1: 1, F2: (true && false)}
                                        _ = v268
                                        _ = v269
                                    }
                                case "":
                                    _ = len(v5)
                                }
                            default:
                                fmt.Println(fi260, v3, v2)
                            }
                            if len(v5) > 0 {
                                v5[0] = (("a" + "") + ("" + "hello"))
                            }
                            v6 += (fi264 * -1)
                            _ = fi264
                            _ = v265
                        }
                    } else {
                        v1 += (^(v1))
                        var v270 int = genFn1()
                        _ = v270
                    }
                case "foo bar":
                    if (2 == 42) {
                        var v271 uint = v4
                        v271 = (^(v4))
                        _ = v271
                    } else {
                        var v272 uint64 = 18446744073709551615
                        for fi273 := 0; fi273 < 12; fi273++ {
                            if (42 >= v4) {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            }
                            v1 -= (v1 ^ (v1 + (v1 * 689)))
                            if len(v5) > 0 {
                                v5[0] = "a"
                            }
                            _ = fi273
                        }
                        v274 := GenS17{F0: 120, F1: (^(v272))}
                        _ = v272
                        _ = v274
                    }
                case "hello":
                    v259 -= ((v259 - v259) / (float32(float64(v259)) * 1.5))
                }
                v2 += (v2 % (((v1 ^ 0)) | 1))
                for fi275 := 0; fi275 < 20; fi275++ {
                    var v276 uint8 = 128
                    _ = fi275
                    _ = v276
                }
                _ = fi260
            }
        case "a":
            fmt.Println(v6)
        case "":
            var v277 int8 = 113
            _ = v277
        }
    case "hello":
        fmt.Println(v4, v3, v1)
    case "":
        fmt.Println(v4, v3)
    }
    var v278 uint32 = (^(v1))
    v278 += (v1 << ((v3) & 63))
    var v279 uint = 0
    var v280 int8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v259:", v259)
    fmt.Println("v278:", v278)
    fmt.Println("v279:", v279)
    fmt.Println("v280:", v280)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v246
    _ = v258
    _ = v259
    _ = v278
    _ = v279
    _ = v280
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
