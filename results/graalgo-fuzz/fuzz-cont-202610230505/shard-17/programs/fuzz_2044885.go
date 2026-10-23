package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 int64; F1 int; F2 uint16 }
type GenS3 struct { F0 int8; F1 uint64; F2 uint64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint64; F1 uint }
type GenS7 struct { F0 uint32; F1 bool; F2 uint }
type GenS8 struct { F0 uint8; F1 uint16; F2 string }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 uint; F1 uint; F2 int16 }
type GenS12 struct { F0 int8; F1 uint8 }
type GenS13 struct { F0 int32 }

func main() {
    var v1 string = ""
    var v2 uint = 441
    var v3 int16 = -748
    var v4 bool = false
    switch (v1 + "hello") {
    case "\n\t":
        if true {
            switch "\n\t" {
            case "hello":
                var v5 uint64 = 541
                _ = v5
            case "a":
                if v4 {
                    for fi6 := 0; fi6 < 9; fi6++ {
                        v3 = (^(v3))
                        _ = fi6
                    }
                    for fi7 := 0; fi7 < 1; fi7++ {
                        v1 = (v1 + v1)
                        v8 := GenS1{F0: 631, F1: (!v4)}
                        _ = fi7
                        _ = v8
                    }
                } else {
                    v9 := GenS1{F0: 32767, F1: v4}
                    v10 := []int{-9223372036854775808, -2147483648, 32767, -481}
                    _ = v10
                    _ = v9
                }
            case "foo bar":
                var v11 int = -1
                _ = v11
            default:
                v12 := GenS2{F0: -211, F1: 1, F2: 0}
                _ = v12
            }
            v13 := []int{-758}
            _ = v13
        } else {
            v14 := GenS3{F0: -1, F1: 42, F2: 986}
            v15 := GenS3{F0: -128, F1: 0, F2: 649}
            v16 := []int{32767, 42, 32767}
            _ = v14
            _ = v15
            _ = v16
        }
    default:
        v17 := GenS3{F0: -59, F1: 18446744073709551615, F2: uint64(v2)}
        _ = v17
    }
    if (!v4) {
        v1 = "hello"
        if false {
            v18 := GenS1{F0: 0, F1: v4}
            v2 = v2
            fmt.Println(v2)
            _ = v18
        }
        if (124 < 75) {
            for fi19 := 0; fi19 < 5; fi19++ {
                switch (fi19 % ((fi19) | 1)) {
                case 32767:
                    v20 := GenS4{F0: 0}
                    _ = v20
                }
                _ = fi19
            }
            for fi21 := 0; fi21 < 19; fi21++ {
                for fi22 := 0; fi22 < 10; fi22++ {
                    v23 := []int{fi22}
                    _ = fi22
                    _ = v23
                }
                v24 := GenS1{F0: 0, F1: (true || false)}
                v25 := GenS5{F0: 1}
                _ = fi21
                _ = v24
                _ = v25
            }
        } else {
            fmt.Println(v1, v3)
            fmt.Println(v1, v1)
        }
    } else {
        fmt.Println(v4, v1)
        var v26 uint64 = 87
        switch v1 {
        case "hello":
            v2 = 255
        case "\n\t":
            if ((v4 && v4) || v4) {
                for fi27 := 0; fi27 < 13; fi27++ {
                    fmt.Println(v3, v26, v2)
                    var v28 int16 = (v3 % ((6) | 1))
                    _ = fi27
                    _ = v28
                }
                for fi29 := 0; fi29 < 15; fi29++ {
                    if (v2 == 218) {
                        v4 = (false && true)
                        for fi30 := 0; fi30 < 3; fi30++ {
                            if (!((v26 % ((967) | 1)) < 429)) {
                                if v4 {
                                    if (("hello" + v1) <= ((v1 + ((("" + (v1 + ("\n\t" + v1))) + v1) + (v1 + (v1 + "a")))) + ("hello" + "foo bar"))) {
                                        if v4 {
                                            var v31 int8 = -128
                                            switch fi29 {
                                            case -1:
                                                var v32 bool = v4
                                                _ = v32
                                            case 4294967295:
                                                var v33 uint = 861
                                                _ = v33
                                            case 157:
                                                var v34 uint = v2
                                                if true {
                                                    switch ("" + "hello") {
                                                    case "hello":
                                                        if (v4 && ((v4 || v4) || ((!v4) || (v4 && (v4 && true))))) {
                                                            v35 := []int{fi29, fi30, -590}
                                                            _ = v35
                                                        } else {
                                                            v36 := GenS6{F0: 0, F1: (v34 | (^(v2)))}
                                                            var v37 int8 = 127
                                                            _ = v36
                                                            _ = v37
                                                        }
                                                    default:
                                                        if v4 {
                                                            if v4 {
                                                                v38 := []int{384, (^(fi29)), (fi29 << ((v2) & 63)), (-(fi29))}
                                                                _ = v38
                                                            }
                                                            var v39 int = 9223372036854775807
                                                            var v40 float64 = -2.25
                                                            _ = v39
                                                            _ = v40
                                                        }
                                                    }
                                                    var v41 uint16 = 284
                                                    v26 += (v26 * v26)
                                                    var v42 int64 = -1
                                                    _ = v41
                                                    _ = v42
                                                } else {
                                                    if v4 {
                                                        if v4 {
                                                            v43 := []int{-537, fi30, (fi29 % ((fi29) | 1))}
                                                            fmt.Println(v1)
                                                            if len(v43) >= 2 {
                                                                v43 = v43[1:len(v43)]
                                                            }
                                                            v3 = (v3 / ((32767) | 1))
                                                            _ = v43
                                                        } else {
                                                            v4 = (!(v31 <= v31))
                                                            v4 = v4
                                                        }
                                                        var v44 string = v1
                                                        if (!false) {
                                                            switch v44 {
                                                            case "hello":
                                                                v45 := []int{}
                                                                _ = v45
                                                            default:
                                                                fmt.Println(fi29, fi30, v4)
                                                            }
                                                            v46 := GenS7{F0: 270, F1: v4, F2: (v2 - (v34 & v2))}
                                                            _ = v46
                                                        }
                                                        _ = v44
                                                    }
                                                    if false {
                                                        v47 := GenS8{F0: 86, F1: 2, F2: v1}
                                                        var v48 uint16 = 42
                                                        _ = v47
                                                        _ = v48
                                                    } else {
                                                        if (971 == -9223372036854775808) {
                                                            v49 := GenS7{F0: 2, F1: ((699 <= 65535) && v4), F2: (v2 * (v34 | 65535))}
                                                            switch (fi30 >> (((v2 * (v34 + (v34 & v34)))) & 63)) {
                                                            case 32768:
                                                                switch fi29 {
                                                                case 128:
                                                                    switch (fi30 ^ fi29) {
                                                                    case -23:
                                                                        fmt.Println(fi30, v3)
                                                                    }
                                                                case 9223372036854775807:
                                                                    v3 -= (v3 + v3)
                                                                }
                                                            }
                                                            _ = v49
                                                        }
                                                        if (803 <= 1) {
                                                            if v4 {
                                                                var v50 int32 = 758
                                                                var v51 uint = 1
                                                                _ = v50
                                                                _ = v51
                                                            } else {
                                                                if (3.14159 > (1e10 + 1.5)) {
                                                                    var v52 int8 = (v31 / ((v31) | 1))
                                                                    v52 += (v52 % (((v52 + v31)) | 1))
                                                                    v53 := GenS3{F0: v31, F1: 734, F2: 2}
                                                                    fmt.Println(fi30, v4, v31)
                                                                    v54 := GenS9{F0: 0}
                                                                    _ = v52
                                                                    _ = v53
                                                                    _ = v54
                                                                } else {
                                                                    var v55 int64 = 890
                                                                    _ = v55
                                                                }
                                                            }
                                                        }
                                                        var v56 int8 = (v31 ^ v31)
                                                        var v57 float64 = float64(0.0)
                                                        v31 += (v31 % ((v56) | 1))
                                                        _ = v56
                                                        _ = v57
                                                    }
                                                }
                                                _ = v34
                                            default:
                                                var v58 int = (^(fi30))
                                                _ = v58
                                            }
                                            var v59 uint64 = (v26 | (v26 * v26))
                                            _ = v31
                                            _ = v59
                                        } else {
                                            switch 1 {
                                            case 32768:
                                                switch v1 {
                                                case "hello":
                                                    v60 := []int{fi29, 232, (fi30 - fi30), fi30}
                                                    _ = v60
                                                case "foo bar":
                                                    fmt.Println(v1)
                                                case "\n\t":
                                                    var v61 float64 = (1e10 / 1.0)
                                                    _ = v61
                                                default:
                                                    v2 = (v2 * v2)
                                                }
                                                if (127 > 1) {
                                                    v26 += v26
                                                    v26 = (v26 + (v26 + (^(v26))))
                                                }
                                                var v62 string = (v1 + "foo bar")
                                                _ = v62
                                            case -767:
                                                v4 = (594 >= v2)
                                            default:
                                                var v63 uint8 = 2
                                                _ = v63
                                            }
                                            var v64 uint8 = 2
                                            v26 = uint64(v2)
                                            v65 := []int{fi29, (fi29 ^ fi30), -206, 32767}
                                            fmt.Println(v64)
                                            _ = v64
                                            _ = v65
                                        }
                                        v66 := GenS4{F0: (^(v26))}
                                        _ = v66
                                    } else {
                                        v2 = (v2 % (((v2 | (^(v2)))) | 1))
                                    }
                                    fmt.Println(fi30, v4)
                                    v67 := GenS1{F0: uint16(v26), F1: v4}
                                    _ = v67
                                }
                                v68 := []int{(fi29 >> (((v2 ^ 240)) & 63)), (fi30 + (fi30 | fi29))}
                                _ = v68
                            }
                            v69 := GenS10{F0: 1}
                            v70 := []int{}
                            fmt.Println(v4)
                            _ = fi30
                            _ = v69
                            _ = v70
                        }
                    } else {
                        v71 := GenS10{F0: v26}
                        var v72 int16 = v3
                        for fi73 := 0; fi73 < 12; fi73++ {
                            fmt.Println(fi73)
                            _ = fi73
                        }
                        _ = v71
                        _ = v72
                    }
                    var v74 int64 = 1
                    _ = fi29
                    _ = v74
                }
                v4 = v4
            } else {
                v3 = (^(v3))
                for fi75 := 0; fi75 < 1; fi75++ {
                    fmt.Println(v1, v3, v4)
                    if ((v1 + (v1 + (v1 + ("a" + "")))) > "\n\t") {
                        v4 = v4
                        var v76 uint64 = v26
                        _ = v76
                    } else {
                        fmt.Println(v2, v4)
                    }
                    var v77 int16 = (v3 + v3)
                    _ = fi75
                    _ = v77
                }
            }
        case "foo bar":
            v78 := GenS4{F0: (^(v26))}
            _ = v78
        }
        _ = v26
    }
    v3 -= (v3 % ((v3) | 1))
    v1 = ((v1 + v1) + v1)
    if (189 <= 127) {
        if (112 == 1) {
            for fi79 := 0; fi79 < 19; fi79++ {
                v2 += (v2 << ((18446744073709551615) & 63))
                _ = fi79
            }
        }
    } else {
        var v80 uint16 = 347
        v4 = (!true)
        v80 -= v80
        _ = v80
    }
    switch (((v1 + ((((v1 + v1) + (v1 + v1)) + "foo bar") + "")) + (((v1 + "foo bar") + (v1 + v1)) + ("foo bar" + v1))) + ("" + v1)) {
    case "\n\t":
        v2 = (^(v2))
    case "":
        for fi81 := 0; fi81 < 7; fi81++ {
            if (v4 || false) {
                if ("\n\t" >= v1) {
                    v82 := GenS6{F0: 0, F1: v2}
                    fmt.Println(v1, fi81)
                    v3 -= v3
                    _ = v82
                } else {
                    var v83 bool = (v4 || ("a" >= v1))
                    v84 := []int{fi81}
                    for fi85 := 0; fi85 < 12; fi85++ {
                        var v86 bool = v4
                        v2 -= 255
                        var v87 bool = (!v86)
                        _ = fi85
                        _ = v86
                        _ = v87
                    }
                    _ = v83
                    _ = v84
                }
                var v88 float64 = (0.1 * -2.25)
                _ = v88
            }
            _ = fi81
        }
    default:
        v4 = v4
    }
    var v89 uint64 = 18446744073709551615
    if v4 {
        v4 = (!(-508 != 426))
        v3 = v3
        v3 += (v3 + v3)
    } else {
        if true {
            switch v1 {
            case "":
                for fi90 := 0; fi90 < 7; fi90++ {
                    v91 := []int{fi90}
                    fmt.Println(fi90, v4)
                    _ = fi90
                    _ = v91
                }
            default:
                v92 := []int{9223372036854775807, 128}
                _ = v92
            }
            if false {
                switch 670 {
                case -673:
                    var v93 float32 = ((1.5 / 3.14159) + float32(0.1))
                    _ = v93
                case 128:
                    v2 -= (v2 * v2)
                case 88:
                    if true {
                        v2 = v2
                    } else {
                        v94 := []int{42, int(v89), 2147483647, 4294967295}
                        var v95 uint16 = 42
                        var v96 uint64 = v89
                        _ = v94
                        _ = v95
                        _ = v96
                    }
                }
                switch v1 {
                case "foo bar":
                    for fi97 := 0; fi97 < 1; fi97++ {
                        v1 = (v1 + (((v1 + ((v1 + "\n\t") + v1)) + v1) + ("a" + "foo bar")))
                        _ = fi97
                    }
                case "\n\t":
                    if v4 {
                        if v4 {
                            v2 = 18446744073709551615
                        } else {
                            v4 = true
                            v98 := GenS9{F0: 32767}
                            v99 := GenS2{F0: -117, F1: 9223372036854775807, F2: 42}
                            _ = v98
                            _ = v99
                        }
                    } else {
                        v2 = 2
                    }
                default:
                    v89 = v89
                }
            } else {
                var v100 uint16 = 70
                var v101 uint = uint(v100)
                _ = v100
                _ = v101
            }
            for fi102 := 0; fi102 < 20; fi102++ {
                v103 := []int{(fi102 % (((fi102 + (fi102 | fi102))) | 1)), (fi102 >> ((v2) & 63)), -1, (fi102 & int(v3))}
                fmt.Println(v89, v3, v1)
                _ = fi102
                _ = v103
            }
        }
    }
    for fi104 := 0; fi104 < 1; fi104++ {
        v105 := []int{128}
        _ = fi104
        _ = v105
    }
    var v106 int = 1
    if (!v4) {
        switch 156 {
        case -2147483648:
            fmt.Println(v89)
        }
    } else {
        v107 := []int{}
        switch v106 {
        case 2147483647:
            v108 := GenS6{F0: 42, F1: v2}
            _ = v108
        }
        for fi109 := 0; fi109 < 1; fi109++ {
            var v110 int16 = (v3 / ((v3) | 1))
            _ = fi109
            _ = v110
        }
        _ = v107
    }
    v111 := []int{(v106 >> (((v2 * v2)) & 63))}
    switch (((v1 + "a") + "") + "a") {
    case "foo bar":
        var v112 uint32 = 2
        _ = v112
    case "a":
        if (false && v4) {
            var v113 float32 = 0.0
            var v114 int32 = -1
            _ = v113
            _ = v114
        } else {
            if (((!v4) || (!v4)) || v4) {
                v115 := GenS6{F0: 1, F1: v2}
                v116 := GenS11{F0: (v2 / (((v2 << ((v2) & 63))) | 1)), F1: (v2 * (v2 ^ v2)), F2: 42}
                _ = v115
                _ = v116
            } else {
                switch ((v1 + "\n\t") + "\n\t") {
                case "foo bar":
                    fmt.Println(v3)
                case "\n\t":
                    var v117 float32 = (-2.25 + (-2.25 * 1.5))
                    _ = v117
                case "a":
                    var v118 uint = (v2 ^ (v2 / ((1) | 1)))
                    _ = v118
                }
                if v4 {
                    if v4 {
                        switch (^(v106)) {
                        case 9223372036854775807:
                            var v119 float64 = (0.1 * (1e10 + 3.14159))
                            _ = v119
                        default:
                            if v4 {
                                var v120 uint8 = 1
                                _ = v120
                            }
                        }
                    } else {
                        switch -1 {
                        case 412:
                            if false {
                                var v121 uint = (v2 ^ v2)
                                _ = v121
                            }
                        default:
                            var v122 int = (v106 - v106)
                            _ = v122
                        }
                        switch "hello" {
                        case "":
                            var v123 int8 = -28
                            _ = v123
                        case "hello":
                            v124 := GenS8{F0: 42, F1: 42, F2: (("foo bar" + v1) + (v1 + v1))}
                            _ = v124
                        }
                        var v125 bool = v4
                        _ = v125
                    }
                    var v126 float32 = -2.25
                    _ = v126
                }
            }
            fmt.Println(v2, v3, v2)
        }
    case "\n\t":
        v111 = append(v111, (v106 >> (((^(v2))) & 63)))
    default:
        if len(v111) > 0 {
            v111[0] = (-(v106))
        }
    }
    if ((!(v106 < v106)) && v4) {
        var v127 int16 = 32767
        if len(v111) >= 2 {
            v111 = v111[1:len(v111)]
        }
        _ = v127
    } else {
        var v128 uint8 = 23
        fmt.Println(v4, v106, v89)
        for fi129 := 0; fi129 < 3; fi129++ {
            v130 := GenS11{F0: v2, F1: v2, F2: v3}
            switch v1 {
            case "\n\t":
                v131 := GenS12{F0: 3, F1: v128}
                _ = v131
            case "foo bar":
                v89 += 64
            }
            _ = fi129
            _ = v130
        }
        _ = v128
    }
    v132 := GenS5{F0: 4294967295}
    v1 = ((v1 + "a") + "hello")
    v133 := GenS13{F0: 376}
    if len(v111) > 0 {
        v111[0] = (^(v106))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v106:", v106)
    fmt.Println("slicelen v111:", len(v111), v111)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v106
    _ = v111
    _ = v132
    _ = v133
    _ = v2
    _ = v3
    _ = v4
    _ = v89
}
