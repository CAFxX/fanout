package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 int64; F1 uint; F2 int8 }
type GenS4 struct { F0 uint8; F1 uint64 }
type GenS5 struct { F0 int8; F1 uint16 }
type GenS6 struct { F0 uint16; F1 uint8 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 uint64; F1 int }
type GenS10 struct { F0 uint; F1 uint64; F2 uint8 }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 bool; F1 int32 }
type GenS14 struct { F0 uint32 }
type GenS15 struct { F0 bool }
type GenS16 struct { F0 int8; F1 uint16; F2 bool }
type GenS17 struct { F0 uint; F1 string; F2 int }
type GenS18 struct { F0 int64; F1 uint64; F2 uint32 }
type GenS19 struct { F0 uint16; F1 int16 }
func genFn0(p0 uint32, p1 int32) string {
    fmt.Println(p0)
    v4 := GenS1{F0: 65535, F1: 20}
    var v5 uint64 = 1
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return ("" + "hello")
}
func genFn1() uint32 {
    var v6 uint32 = 857
    var v7 float32 = 1e10
    fmt.Println(v7, v7)
    _ = v6
    _ = v7
    return v6
}

func main() {
    var v1 int8 = 0
    var v2 int16 = -1
    var v3 string = (("hello" + "\n\t") + ("" + "foo bar"))
    v1 = (v1 / ((42) | 1))
    for fi8 := 0; fi8 < 20; fi8++ {
        for fi9 := 0; fi9 < 10; fi9++ {
            switch v3 {
            case "":
                v2 = -32768
            }
            _ = fi9
        }
        v1 = (v1 % ((v1) | 1))
        _ = fi8
    }
    for fi10 := 0; fi10 < 2; fi10++ {
        v2 -= 0
        var v11 uint16 = 0
        if (!false) {
            var v12 uint64 = 42
            var v13 int16 = (v2 % ((v2) | 1))
            var v14 uint = 1
            _ = v12
            _ = v13
            _ = v14
        }
        _ = fi10
        _ = v11
    }
    v2 -= -641
    for fi15 := 0; fi15 < 14; fi15++ {
        v3 = v3
        _ = fi15
    }
    var v16 bool = (false || true)
    v17 := []int{801}
    v3 = v3
    v3 = "\n\t"
    if len(v17) >= 2 {
        v17 = v17[1:len(v17)]
    }
    v1 = v1
    if (255 == 163) {
        fmt.Println(v3, v1)
    } else {
        if len(v17) >= 2 {
            v17 = v17[1:len(v17)]
        }
        var v18 uint32 = 2147483648
        if (4294967295 == 1) {
            if true {
                v19 := GenS2{F0: 2147483647, F1: 599}
                for fi20 := 0; fi20 < 11; fi20++ {
                    fmt.Println(fi20)
                    v16 = true
                    v18 = genFn1()
                    _ = fi20
                }
                if v16 {
                    v2 -= (v2 - (v2 >> ((4294967295) & 63)))
                } else {
                    for fi21 := 0; fi21 < 13; fi21++ {
                        v1 += (-(v1))
                        v3 = (v3 + v3)
                        _ = fi21
                    }
                    if (18446744073709551615 == 42) {
                        for fi22 := 0; fi22 < 19; fi22++ {
                            v23 := GenS2{F0: 539, F1: -2147483648}
                            _ = fi22
                            _ = v23
                        }
                        switch 32767 {
                        case 337:
                            if v16 {
                                v1 = (v1 + (v1 | v1))
                            } else {
                                v18 += v18
                                var v24 int16 = v2
                                var v25 uint16 = 0
                                _ = v24
                                _ = v25
                            }
                        case 127:
                            v16 = v16
                        default:
                            for fi26 := 0; fi26 < 3; fi26++ {
                                if len(v17) > 0 {
                                    v17[0] = (-(fi26))
                                }
                                _ = fi26
                            }
                        }
                        var v27 int32 = 0
                        _ = v27
                    }
                }
                _ = v19
            }
            var v28 float64 = 1e10
            if (v1 != v1) {
                v1 -= (v1 % ((-69) | 1))
            }
            _ = v28
        } else {
            switch 128 {
            case 2147483647:
                if v16 {
                    v29 := GenS3{F0: -9223372036854775808, F1: 779, F2: v1}
                    v17 = append(v17, -589)
                    _ = v29
                }
            case 128:
                if v16 {
                    if (-2147483648 >= 0) {
                        v30 := GenS1{F0: 936, F1: 985}
                        for fi31 := 0; fi31 < 8; fi31++ {
                            for fi32 := 0; fi32 < 2; fi32++ {
                                fmt.Println(v16)
                                switch fi32 {
                                case -1:
                                    if true {
                                        v33 := GenS4{F0: 219, F1: 1}
                                        v2 = (v2 | (v2 + 1000))
                                        _ = v33
                                    }
                                case -125:
                                    if (!(v1 == -128)) {
                                        v2 = v2
                                        fmt.Println(v3, v18)
                                        switch ("" + v3) {
                                        case "foo bar":
                                            if v16 {
                                                switch (^(fi32)) {
                                                case 963:
                                                    switch ((v3 + ((((v3 + v3) + v3) + ("hello" + v3)) + (v3 + "a"))) + v3) {
                                                    case "":
                                                        switch ((((((v3 + ("" + (v3 + v3))) + v3) + "hello") + "foo bar") + ((v3 + v3) + v3)) + "\n\t") {
                                                        case "\n\t":
                                                            var v34 uint = 2
                                                            var v35 float64 = 3.14159
                                                            if false {
                                                                if v16 {
                                                                    switch ((v3 + "foo bar") + "foo bar") {
                                                                    case "a":
                                                                        var v36 int8 = 0
                                                                        _ = v36
                                                                    case "foo bar":
                                                                        _ = len(v17)
                                                                    default:
                                                                        var v37 float64 = (v35 / float64(1.5))
                                                                        _ = v37
                                                                    }
                                                                    if len(v17) > 0 {
                                                                        v17[0] = fi31
                                                                    }
                                                                }
                                                                if len(v17) >= 2 {
                                                                    v17 = v17[1:len(v17)]
                                                                }
                                                                v16 = (32767 == 42)
                                                            } else {
                                                                if v16 {
                                                                    v38 := GenS3{F0: 42, F1: (v34 % ((v34) | 1)), F2: (^(v1))}
                                                                    var v39 uint = (v34 * v34)
                                                                    v16 = v16
                                                                    _ = v38
                                                                    _ = v39
                                                                }
                                                                v1 += v1
                                                                v2 = -32768
                                                            }
                                                            _ = v34
                                                            _ = v35
                                                        default:
                                                            v40 := GenS4{F0: 1, F1: 18446744073709551615}
                                                            _ = v40
                                                        }
                                                    case "\n\t":
                                                        if (24 < v1) {
                                                            v18 -= v18
                                                        } else {
                                                            if v16 {
                                                                v3 = "foo bar"
                                                                var v41 int16 = v2
                                                                fmt.Println(fi31)
                                                                _ = v41
                                                            } else {
                                                                if v16 {
                                                                    var v42 uint8 = 127
                                                                    var v43 float32 = 1.5
                                                                    _ = v42
                                                                    _ = v43
                                                                } else {
                                                                    fmt.Println(v3)
                                                                    v44 := GenS2{F0: 1, F1: 181}
                                                                    switch fi31 {
                                                                    case 42:
                                                                        switch fi32 {
                                                                        case -100:
                                                                            v16 = (false && v16)
                                                                        case 32768:
                                                                            fmt.Println(v1)
                                                                        case -1:
                                                                            switch v3 {
                                                                            case "hello":
                                                                                fmt.Println(fi31, v16)
                                                                                var v45 int16 = (v2 | (^(v2)))
                                                                                _ = v45
                                                                            case "foo bar":
                                                                                v1 += (v1 >> ((300) & 63))
                                                                            }
                                                                        }
                                                                    case 1:
                                                                        v2 = 0
                                                                    case 0:
                                                                        var v46 float32 = 1.5
                                                                        _ = v46
                                                                    default:
                                                                        if (!false) {
                                                                            if v16 {
                                                                                switch v3 {
                                                                                case "hello":
                                                                                    var v47 uint8 = 42
                                                                                    _ = v47
                                                                                case "\n\t":
                                                                                    if (v18 <= (v18 - (v18 & genFn1()))) {
                                                                                        if v16 {
                                                                                            fmt.Println(fi32, v16)
                                                                                            var v48 uint16 = 65535
                                                                                            _ = v48
                                                                                        } else {
                                                                                            if v16 {
                                                                                                v3 = "hello"
                                                                                            }
                                                                                            switch v3 {
                                                                                            case "foo bar":
                                                                                                fmt.Println(v1, fi31)
                                                                                            case "":
                                                                                                var v49 int = fi31
                                                                                                _ = v49
                                                                                            case "a":
                                                                                                v16 = (!(0 == 497))
                                                                                            default:
                                                                                                v3 = v3
                                                                                                _ = len(v17)
                                                                                            }
                                                                                            var v50 uint16 = 735
                                                                                            _ = v50
                                                                                        }
                                                                                    } else {
                                                                                        switch (fi31 & fi32) {
                                                                                        case 9223372036854775807:
                                                                                            v17 = append(v17, fi32)
                                                                                        default:
                                                                                            var v51 uint = 509
                                                                                            _ = v51
                                                                                        }
                                                                                    }
                                                                                default:
                                                                                    switch ("" + v3) {
                                                                                    case "":
                                                                                        var v52 int8 = (^(v1))
                                                                                        _ = v52
                                                                                    case "\n\t":
                                                                                        var v53 string = v3
                                                                                        _ = v53
                                                                                    case "a":
                                                                                        var v54 int8 = v1
                                                                                        _ = v54
                                                                                    default:
                                                                                        v55 := GenS5{F0: 100, F1: 1}
                                                                                        _ = v55
                                                                                    }
                                                                                }
                                                                            }
                                                                            if (v2 > v2) {
                                                                                v18 -= (v18 << ((2) & 63))
                                                                            }
                                                                        } else {
                                                                            v16 = (-12 >= (^(fi32)))
                                                                            fmt.Println(v3)
                                                                        }
                                                                    }
                                                                    _ = v44
                                                                }
                                                                v17 = append(v17, (fi32 >> ((42) & 63)))
                                                                fmt.Println(v3, v18, v2)
                                                            }
                                                            v16 = v16
                                                            v16 = v16
                                                        }
                                                    }
                                                case 42:
                                                    v1 -= (v1 * v1)
                                                }
                                            } else {
                                                v56 := GenS4{F0: 0, F1: 4294967295}
                                                if v16 {
                                                    v57 := GenS6{F0: 42, F1: 42}
                                                    switch v3 {
                                                    case "foo bar":
                                                        if ((v2 | -129) == (v2 / ((687) | 1))) {
                                                            var v58 int64 = 1
                                                            _ = len(v17)
                                                            v59 := GenS7{F0: (v18 / ((v18) | 1))}
                                                            _ = v58
                                                            _ = v59
                                                        }
                                                    }
                                                    if (!((0.1 <= float64(0.0)) || (((((0.0 + -2.25) - ((0.0 * 1.5) * float64((0.0 / ((0.0 / -2.25) + -2.25))))) * 3.14159) * (-2.25 / -2.25)) < ((1.5 * 3.14159) - 3.14159)))) {
                                                        if v16 {
                                                            v17 = append(v17, (fi31 >> ((255) & 63)))
                                                            if len(v17) >= 2 {
                                                                v17 = v17[1:len(v17)]
                                                            }
                                                        }
                                                    } else {
                                                        var v60 int64 = -9223372036854775808
                                                        _ = len(v17)
                                                        _ = v60
                                                    }
                                                    _ = v57
                                                }
                                                var v61 uint64 = 492
                                                _ = v56
                                                _ = v61
                                            }
                                        case "a":
                                            if (!v16) {
                                                v1 = (-(v1))
                                                var v62 float32 = 0.0
                                                v63 := GenS2{F0: 2147483647, F1: 0}
                                                v64 := GenS8{F0: ((true && (!((float32(1.5) * 0.1) <= (float32(1.5) / ((v62 / v62) * v62))))) && v16)}
                                                _ = v62
                                                _ = v63
                                                _ = v64
                                            } else {
                                                v1 = -26
                                            }
                                        case "\n\t":
                                            if len(v17) > 0 {
                                                v17[0] = (-(fi31))
                                            }
                                        default:
                                            switch (v3 + "") {
                                            case "\n\t":
                                                v65 := GenS9{F0: 18446744073709551615, F1: (fi31 << ((0) & 63))}
                                                var v66 int16 = v2
                                                _ = v65
                                                _ = v66
                                            default:
                                                v18 = v18
                                            }
                                        }
                                    } else {
                                        var v67 uint32 = (v18 & genFn1())
                                        v2 -= (^(v2))
                                        if false {
                                            fmt.Println(v18, v67, v67)
                                            var v68 int64 = 0
                                            _ = v68
                                        } else {
                                            v1 += (v1 | v1)
                                            var v69 uint64 = 18446744073709551615
                                            var v70 int16 = (v2 % (((v2 | (v2 - (v2 * (^(v2)))))) | 1))
                                            _ = v69
                                            _ = v70
                                        }
                                        _ = v67
                                    }
                                case 4294967295:
                                    var v71 string = (v3 + (((v3 + v3) + "foo bar") + (("" + (((v3 + "foo bar") + v3) + v3)) + v3)))
                                    _ = v71
                                default:
                                    if (42 > -19) {
                                        if len(v17) > 0 {
                                            v17[0] = fi31
                                        }
                                        v16 = (!v16)
                                        switch fi32 {
                                        case 128:
                                            switch ("hello" + (v3 + "\n\t")) {
                                            case "\n\t":
                                                _ = len(v17)
                                                if len(v17) >= 2 {
                                                    v17 = v17[1:len(v17)]
                                                }
                                                switch (fi32 * 128) {
                                                case -2147483648:
                                                    var v72 bool = true
                                                    _ = v72
                                                }
                                            }
                                        case 9223372036854775807:
                                            if (-2147483648 == 969) {
                                                if len(v17) >= 2 {
                                                    v17 = v17[1:len(v17)]
                                                }
                                                if (v16 && (4294967295 == 4294967295)) {
                                                    if len(v17) > 0 {
                                                        v17[0] = fi31
                                                    }
                                                } else {
                                                    var v73 float32 = 0.1
                                                    var v74 uint64 = 70
                                                    v18 += v18
                                                    _ = v73
                                                    _ = v74
                                                }
                                                v17 = append(v17, 127)
                                            }
                                        case 0:
                                            fmt.Println(v3, v18)
                                        default:
                                            var v75 int64 = 1
                                            _ = v75
                                        }
                                    } else {
                                        v76 := GenS3{F0: int64(v1), F1: 271, F2: -1}
                                        v77 := GenS10{F0: 1, F1: 1, F2: 0}
                                        _ = v76
                                        _ = v77
                                    }
                                }
                                if (!(uint(v18) != 42)) {
                                    if v16 {
                                        var v78 int16 = (v2 * (v2 + 1))
                                        var v79 uint16 = 762
                                        _ = v78
                                        _ = v79
                                    } else {
                                        switch fi31 {
                                        case 1:
                                            switch (("\n\t" + v3) + "") {
                                            case "foo bar":
                                                switch fi31 {
                                                case 0:
                                                    if len(v17) > 0 {
                                                        v17[0] = (fi31 ^ (fi31 ^ (fi32 & fi31)))
                                                    }
                                                case 559:
                                                    var v80 uint64 = 1
                                                    _ = v80
                                                case -307:
                                                    var v81 int64 = 1
                                                    _ = v81
                                                }
                                            }
                                            switch fi31 {
                                            case -2147483648:
                                                if v16 {
                                                    v2 = v2
                                                    switch fi31 {
                                                    case 42:
                                                        v82 := GenS4{F0: 128, F1: 0}
                                                        _ = v82
                                                    case 9223372036854775807:
                                                        v2 += 42
                                                    case 127:
                                                        if ((3.14159 * float64(float32(float64(((-2.25 / -2.25) + 0.0))))) <= 0.1) {
                                                            fmt.Println(fi31, v16)
                                                            switch ("" + ((v3 + v3) + ((v3 + ((("foo bar" + (v3 + "a")) + v3) + v3)) + ((v3 + v3) + "a")))) {
                                                            case "hello":
                                                                var v83 int32 = 1
                                                                _ = v83
                                                            case "":
                                                                if (false || true) {
                                                                    v84 := GenS5{F0: -1, F1: 255}
                                                                    _ = v84
                                                                }
                                                            case "foo bar":
                                                                v1 += v1
                                                            }
                                                        } else {
                                                            v1 += v1
                                                            var v85 uint8 = 3
                                                            v17 = append(v17, 42)
                                                            _ = v85
                                                        }
                                                    default:
                                                        var v86 uint = 42
                                                        _ = v86
                                                    }
                                                    v87 := GenS3{F0: 0, F1: 903, F2: 42}
                                                    _ = v87
                                                }
                                            }
                                            var v88 bool = (42 <= 2)
                                            _ = v88
                                        case 0:
                                            var v89 uint16 = 42
                                            _ = v89
                                        case 2147483647:
                                            switch fi31 {
                                            case 127:
                                                var v90 uint64 = 2
                                                _ = v90
                                            }
                                        }
                                        var v91 uint32 = genFn1()
                                        v3 = ("hello" + ("\n\t" + v3))
                                        _ = v91
                                    }
                                } else {
                                    v16 = true
                                    v92 := GenS5{F0: v1, F1: 65535}
                                    _ = v92
                                }
                                _ = fi32
                            }
                            for fi93 := 0; fi93 < 10; fi93++ {
                                v17 = append(v17, (-(fi93)))
                                var v94 uint32 = v18
                                _ = fi93
                                _ = v94
                            }
                            if false {
                                if (v16 || (true && (!(!(0 <= 255))))) {
                                    var v95 int = (fi31 % (((fi31 >> ((560) & 63))) | 1))
                                    if (v16 || true) {
                                        v3 = v3
                                        if (v3 < ("a" + (((v3 + ("hello" + v3)) + (v3 + v3)) + "a"))) {
                                            v3 = (((v3 + v3) + "\n\t") + v3)
                                            v16 = (3.14159 == (0.1 / 1.0))
                                            v96 := GenS1{F0: 1, F1: 0}
                                            _ = v96
                                        } else {
                                            switch "hello" {
                                            case "\n\t":
                                                v97 := GenS9{F0: 227, F1: int(v2)}
                                                _ = v97
                                            case "":
                                                var v98 uint8 = 127
                                                _ = v98
                                            case "a":
                                                for fi99 := 0; fi99 < 6; fi99++ {
                                                    v2 = v2
                                                    _ = fi99
                                                }
                                            default:
                                                fmt.Println(fi31)
                                            }
                                        }
                                    } else {
                                        v1 = (^(v1))
                                        v100 := GenS11{F0: 42}
                                        if v16 {
                                            switch fi31 {
                                            case -439:
                                                var v101 uint = 42
                                                _ = v101
                                            case -1:
                                                v102 := GenS6{F0: 32767, F1: 131}
                                                _ = v102
                                            default:
                                                v17 = append(v17, fi31)
                                            }
                                            var v103 int = int(v2)
                                            _ = v103
                                        } else {
                                            var v104 int = -169
                                            _ = v104
                                        }
                                        _ = v100
                                    }
                                    _ = v95
                                }
                                if v16 {
                                    v105 := GenS3{F0: 1, F1: 286, F2: -49}
                                    _ = v105
                                } else {
                                    if (0 <= 75) {
                                        if len(v17) > 0 {
                                            v17[0] = fi31
                                        }
                                        v106 := GenS12{F0: (v18 + v18)}
                                        _ = v106
                                    } else {
                                        var v107 int = fi31
                                        _ = v107
                                    }
                                    if len(v17) > 0 {
                                        v17[0] = 1
                                    }
                                    if (!v16) {
                                        v16 = true
                                        v1 = 127
                                        v16 = false
                                    }
                                }
                            } else {
                                if (1.5 >= 1e10) {
                                    if false {
                                        if false {
                                            fmt.Println(v16)
                                        } else {
                                            var v108 bool = false
                                            if len(v17) > 0 {
                                                v17[0] = fi31
                                            }
                                            _ = v108
                                        }
                                        switch 128 {
                                        case -9223372036854775808:
                                            var v109 int64 = 0
                                            _ = v109
                                        default:
                                            v110 := GenS13{F0: v16, F1: -917}
                                            _ = v110
                                        }
                                        v111 := GenS7{F0: 73}
                                        _ = v111
                                    }
                                    v3 = v3
                                } else {
                                    var v112 uint8 = 255
                                    _ = v112
                                }
                                for fi113 := 0; fi113 < 8; fi113++ {
                                    v114 := GenS9{F0: 603, F1: fi113}
                                    _ = len(v17)
                                    if v16 {
                                        var v115 float32 = (1e10 / (3.14159 - 1e10))
                                        v116 := GenS14{F0: 0}
                                        _ = v115
                                        _ = v116
                                    }
                                    _ = fi113
                                    _ = v114
                                }
                            }
                            _ = fi31
                        }
                        _ = len(v17)
                        _ = v30
                    } else {
                        v117 := GenS5{F0: v1, F1: 566}
                        if (-1 == 2147483647) {
                            var v118 int32 = 1
                            for fi119 := 0; fi119 < 18; fi119++ {
                                var v120 uint8 = 109
                                _ = fi119
                                _ = v120
                            }
                            var v121 float32 = 1.5
                            _ = v118
                            _ = v121
                        } else {
                            for fi122 := 0; fi122 < 14; fi122++ {
                                var v123 int8 = 42
                                for fi124 := 0; fi124 < 20; fi124++ {
                                    fmt.Println(v18, v16)
                                    if len(v17) > 0 {
                                        v17[0] = fi124
                                    }
                                    var v125 uint = 4294967295
                                    _ = fi124
                                    _ = v125
                                }
                                v18 = genFn1()
                                _ = fi122
                                _ = v123
                            }
                        }
                        v126 := GenS7{F0: genFn1()}
                        _ = v117
                        _ = v126
                    }
                } else {
                    for fi127 := 0; fi127 < 20; fi127++ {
                        v17 = append(v17, 127)
                        _ = fi127
                    }
                    v3 = v3
                    v1 = 1
                }
            case -146:
                if (1 <= 0) {
                    fmt.Println(v18, v18)
                    switch int(v2) {
                    case 9223372036854775807:
                        v16 = (!false)
                    case -191:
                        v17 = append(v17, 585)
                    }
                    var v128 int = -674
                    _ = v128
                } else {
                    v18 -= v18
                    v129 := GenS15{F0: ((1000 <= 0) && v16)}
                    _ = v129
                }
            }
            for fi130 := 0; fi130 < 19; fi130++ {
                if (1.5 > 0.0) {
                    for fi131 := 0; fi131 < 9; fi131++ {
                        v132 := GenS9{F0: 771, F1: -9223372036854775808}
                        if true {
                            var v133 uint = 42
                            v134 := GenS13{F0: true, F1: -2147483648}
                            fmt.Println(v3, v2, v2)
                            _ = v133
                            _ = v134
                        }
                        v18 += (v18 + v18)
                        var v135 uint16 = 830
                        if false {
                            v135 -= v135
                            var v136 float32 = (3.14159 / 1.0)
                            var v137 int32 = 0
                            _ = v136
                            _ = v137
                        }
                        _ = fi131
                        _ = v132
                        _ = v135
                    }
                    if v16 {
                        v138 := GenS2{F0: -459, F1: 946}
                        if false {
                            v1 -= (-(v1))
                        } else {
                            for fi139 := 0; fi139 < 17; fi139++ {
                                v2 -= 1000
                                _ = fi139
                            }
                            v1 = 127
                        }
                        v3 = ("" + v3)
                        _ = v138
                    } else {
                        if true {
                            var v140 int = (^(fi130))
                            if (42 == v2) {
                                for fi141 := 0; fi141 < 14; fi141++ {
                                    var v142 int = 1
                                    _ = fi141
                                    _ = v142
                                }
                            } else {
                                var v143 float64 = (1e10 - 1e10)
                                switch (v140 * (fi130 | (v140 << ((398) & 63)))) {
                                case 127:
                                    v3 = ("a" + (v3 + "a"))
                                case 42:
                                    v144 := GenS16{F0: 104, F1: 2, F2: ((v1 << ((1) & 63)) > v1)}
                                    _ = v144
                                default:
                                    var v145 uint = 65535
                                    _ = v145
                                }
                                _ = v143
                            }
                            v16 = v16
                            _ = v140
                        } else {
                            v1 += (v1 + (v1 / (((v1 + (v1 << ((2) & 63)))) | 1)))
                        }
                        if ((!(!false)) || v16) {
                            v146 := GenS8{F0: v16}
                            if true {
                                if len(v17) >= 2 {
                                    v17 = v17[1:len(v17)]
                                }
                                v3 = (v3 + v3)
                                v3 = v3
                            } else {
                                switch fi130 {
                                case 4294967295:
                                    if ((!(1 == 1)) || v16) {
                                        v147 := GenS17{F0: 672, F1: v3, F2: 32768}
                                        for fi148 := 0; fi148 < 5; fi148++ {
                                            fmt.Println(v18)
                                            v149 := GenS11{F0: 320}
                                            _ = fi148
                                            _ = v149
                                        }
                                        _ = v147
                                    }
                                case 293:
                                    var v150 int64 = 0
                                    _ = v150
                                default:
                                    switch fi130 {
                                    case -2147483648:
                                        v2 = v2
                                    case -1:
                                        var v151 uint32 = genFn1()
                                        _ = v151
                                    default:
                                        if len(v17) >= 2 {
                                            v17 = v17[1:len(v17)]
                                        }
                                    }
                                }
                            }
                            _ = v146
                        } else {
                            var v152 int64 = -69
                            if (v2 != (v2 & v2)) {
                                var v153 float32 = (1.5 + -2.25)
                                if true {
                                    var v154 int16 = (v2 * v2)
                                    _ = v154
                                }
                                _ = v153
                            } else {
                                if len(v17) > 0 {
                                    v17[0] = fi130
                                }
                                var v155 uint64 = 18446744073709551615
                                v16 = true
                                _ = v155
                            }
                            switch (fi130 / (((fi130 >> ((446) & 63))) | 1)) {
                            case 2147483647:
                                v2 += (v2 | int16(v152))
                            case 0:
                                v156 := GenS18{F0: v152, F1: 410, F2: (^(v18))}
                                _ = v156
                            default:
                                switch fi130 {
                                case 42:
                                    for fi157 := 0; fi157 < 5; fi157++ {
                                        var v158 uint8 = 255
                                        v152 -= 42
                                        var v159 uint = 906
                                        _ = fi157
                                        _ = v158
                                        _ = v159
                                    }
                                default:
                                    for fi160 := 0; fi160 < 9; fi160++ {
                                        if (v16 && (2 < 712)) {
                                            v161 := GenS19{F0: 255, F1: (v2 - (v2 << ((65535) & 63)))}
                                            _ = v161
                                        } else {
                                            var v162 uint64 = 1
                                            switch (fi130 >> ((4294967295) & 63)) {
                                            case 128:
                                                fmt.Println(v3, v2, v152)
                                            case 32768:
                                                v17 = append(v17, (fi130 / ((32768) | 1)))
                                            default:
                                                if len(v17) > 0 {
                                                    v17[0] = fi160
                                                }
                                            }
                                            _ = v162
                                        }
                                        v163 := GenS12{F0: (v18 ^ v18)}
                                        v16 = (((^(v18)) > genFn1()) || true)
                                        _ = fi160
                                        _ = v163
                                    }
                                }
                            }
                            _ = v152
                        }
                    }
                    if len(v17) >= 2 {
                        v17 = v17[1:len(v17)]
                    }
                }
                _ = fi130
            }
            var v164 float32 = 0.0
            _ = v164
        }
        _ = v18
    }
    if false {
        fmt.Println(v16, v2)
        for fi165 := 0; fi165 < 11; fi165++ {
            switch fi165 {
            case -1:
                var v166 uint = 0
                _ = v166
            }
            fmt.Println(v2, v1, fi165)
            v16 = (!(int32(v1) == -1))
            _ = fi165
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
