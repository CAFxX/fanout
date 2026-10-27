package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 uint; F1 string; F2 int }
type GenS3 struct { F0 uint64; F1 int16 }
type GenS4 struct { F0 bool; F1 uint }
type GenS5 struct { F0 bool; F1 uint64 }
type GenS6 struct { F0 int64; F1 uint32; F2 bool }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 uint32; F1 bool }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 uint; F1 uint64; F2 int8 }
type GenS12 struct { F0 int64; F1 bool; F2 uint32 }
type GenS13 struct { F0 uint; F1 string }
type GenS14 struct { F0 string; F1 uint16 }
type GenS15 struct { F0 int; F1 int; F2 string }
type GenS16 struct { F0 bool; F1 int8; F2 bool }
type GenS17 struct { F0 int8 }
type GenS18 struct { F0 uint8; F1 uint }
type GenS19 struct { F0 int64; F1 int16 }
type GenS20 struct { F0 uint; F1 uint16 }
type GenS21 struct { F0 string; F1 string }
func genFn0() uint {
    var v6 int64 = 1
    var v7 int16 = -1
    _ = v6
    _ = v7
    return 1
}

func main() {
    var v1 int16 = -422
    var v2 bool = (255 != 65535)
    v3 := make(map[uint32]int)
    var v4 float32 = float32(0.1)
    v5 := make(map[string]int)
    switch "a" {
    case "foo bar":
        var v8 float64 = 1.5
        _ = v8
    case "hello":
        for fi9 := 0; fi9 < 4; fi9++ {
            v10 := GenS1{F0: 63, F1: 4294967295}
            v11 := []int{(fi9 * (fi9 >> ((4294967295) & 63))), fi9, fi9}
            _ = fi9
            _ = v10
            _ = v11
        }
    }
    for fi12 := 0; fi12 < 10; fi12++ {
        if (2 != 42) {
            v1 -= v1
            switch "hello" {
            case "":
                fmt.Println(fi12, v1)
            case "hello":
                for fi13 := 0; fi13 < 18; fi13++ {
                    switch "foo bar" {
                    case "foo bar":
                        v14 := GenS1{F0: 241, F1: genFn0()}
                        _ = v14
                    }
                    v1 = 1
                    _ = fi13
                }
            default:
                v15 := v5[""]
                _ = v15
            }
        } else {
            v16 := GenS2{F0: 0, F1: "a", F2: fi12}
            v17 := GenS3{F0: 1, F1: (v1 % (((v1 % (((v1 ^ (v1 * v1))) | 1))) | 1))}
            _ = v16
            _ = v17
        }
        _ = len(v3)
        v1 = (v1 | v1)
        _ = fi12
    }
    var v18 uint64 = 4294967295
    v4 += (float32(0.1) + v4)
    switch -9223372036854775808 {
    case 9223372036854775807:
        if (("foo bar" >= ("" + ("foo bar" + "a"))) && (!(-984 < -1))) {
            var v19 int = 32767
            if ((-318 == -1) && v2) {
                fmt.Println(v4)
            } else {
                v20 := GenS4{F0: (((-9223372036854775808 < 270) || v2) && (v18 != 42)), F1: 0}
                fmt.Println(v1, v1, v4)
                _ = v20
            }
            v1 = (v1 + (v1 >> ((65535) & 63)))
            _ = v19
        } else {
            if (2 > 0) {
                if v2 {
                    v21 := []int{-666, -9223372036854775808, -9223372036854775808, 9223372036854775807}
                    _ = v21
                }
                if (!true) {
                    var v22 int8 = -77
                    if v2 {
                        v18 = 4294967295
                        for fi23 := 0; fi23 < 20; fi23++ {
                            v1 = v1
                            fmt.Println(v18, fi23)
                            if true {
                                v24 := []int{}
                                _ = v24
                            } else {
                                for fi25 := 0; fi25 < 2; fi25++ {
                                    fmt.Println(v18, v4, v4)
                                    if false {
                                        fmt.Println(v4, v22)
                                        var v26 int64 = 42
                                        switch ("\n\t" + ("hello" + ("foo bar" + ""))) {
                                        case "\n\t":
                                            var v27 uint64 = (v18 * v18)
                                            _ = v27
                                        case "":
                                            switch (fi25 >> ((193) & 63)) {
                                            case 42:
                                                if v2 {
                                                    var v28 float64 = (3.14159 / 3.14159)
                                                    if ("\n\t" != (("" + "\n\t") + "")) {
                                                        var v29 int = fi23
                                                        v22 = (^(v22))
                                                        _ = v29
                                                    } else {
                                                        var v30 float32 = -2.25
                                                        v31 := v5["foo bar"]
                                                        var v32 float32 = (0.1 - 1e10)
                                                        _ = v30
                                                        _ = v31
                                                        _ = v32
                                                    }
                                                    _ = len(v5)
                                                    var v33 int8 = -128
                                                    var v34 int = (-(fi23))
                                                    _ = v28
                                                    _ = v33
                                                    _ = v34
                                                }
                                            case -601:
                                                var v35 int = fi25
                                                _ = v35
                                            case 539:
                                                switch (fi23 / ((32767) | 1)) {
                                                case 0:
                                                    v36 := GenS3{F0: (v18 ^ 0), F1: (v1 * (-(v1)))}
                                                    _ = v36
                                                case -9223372036854775808:
                                                    v37 := GenS5{F0: v2, F1: v18}
                                                    _ = v37
                                                case 2147483647:
                                                    v18 -= (v18 ^ v18)
                                                    v22 += v22
                                                    v38 := GenS4{F0: v2, F1: 18446744073709551615}
                                                    _ = v38
                                                }
                                                delete(v5, "foo bar")
                                            default:
                                                var v39 int16 = -982
                                                _ = v39
                                            }
                                        default:
                                            v3[uint32(v18)] = fi23
                                        }
                                        switch ("foo bar" + "") {
                                        case "foo bar":
                                            if (!true) {
                                                v40 := GenS4{F0: v2, F1: 42}
                                                v41 := GenS2{F0: 843, F1: ("a" + ""), F2: fi23}
                                                if v2 {
                                                    switch "" {
                                                    case "\n\t":
                                                        v22 = v22
                                                    }
                                                    var v42 uint64 = (^(v18))
                                                    if v2 {
                                                        v2 = (!v2)
                                                        fmt.Println(fi23, v1)
                                                    } else {
                                                        v43 := []int{-1, 9223372036854775807, (^(fi25))}
                                                        v22 -= 0
                                                        switch fi25 {
                                                        case -408:
                                                            if true {
                                                                switch 1 {
                                                                case 32767:
                                                                    switch 640 {
                                                                    case 32767:
                                                                        v44 := GenS2{F0: 18446744073709551615, F1: "foo bar", F2: (fi23 | (fi25 / ((fi23) | 1)))}
                                                                        _ = v44
                                                                    default:
                                                                        var v45 uint = 18446744073709551615
                                                                        _ = v45
                                                                    }
                                                                case 1:
                                                                    _ = len(v43)
                                                                default:
                                                                    var v46 int = fi23
                                                                    _ = v46
                                                                }
                                                                var v47 int64 = v26
                                                                if (true && v2) {
                                                                    switch (fi25 >> ((2) & 63)) {
                                                                    case -877:
                                                                        var v48 uint16 = 860
                                                                        _ = v48
                                                                    case -9223372036854775808:
                                                                        var v49 uint = 2
                                                                        _ = v49
                                                                    case -1:
                                                                        if len(v43) > 0 {
                                                                            v43[0] = fi23
                                                                        }
                                                                    }
                                                                }
                                                                _ = v47
                                                            } else {
                                                                v50 := GenS6{F0: v26, F1: 2147483647, F2: (!v2)}
                                                                v51 := GenS4{F0: v2, F1: uint(v18)}
                                                                _ = v50
                                                                _ = v51
                                                            }
                                                        case 128:
                                                            v52 := GenS2{F0: 583, F1: "\n\t", F2: fi23}
                                                            _ = v52
                                                        case -418:
                                                            var v53 int8 = (v22 % (((v22 * v22)) | 1))
                                                            _ = v53
                                                        }
                                                        var v54 int = (fi25 << ((501) & 63))
                                                        switch v54 {
                                                        case 2147483647:
                                                            delete(v3, 0)
                                                        case -9223372036854775808:
                                                            var v55 uint16 = 2
                                                            _ = v55
                                                        }
                                                        if false {
                                                            v56 := GenS7{F0: (v1 << ((55) & 63))}
                                                            var v57 int32 = 42
                                                            var v58 bool = false
                                                            _ = v56
                                                            _ = v57
                                                            _ = v58
                                                        } else {
                                                            if v2 {
                                                                if v2 {
                                                                    switch "foo bar" {
                                                                    case "a":
                                                                        v1 = 1
                                                                    default:
                                                                        if v2 {
                                                                            delete(v3, 42)
                                                                        }
                                                                    }
                                                                    v59 := GenS8{F0: 1}
                                                                    var v60 uint16 = 1
                                                                    var v61 int8 = v22
                                                                    delete(v3, 2)
                                                                    var v62 int16 = (v1 - v1)
                                                                    _ = v59
                                                                    _ = v60
                                                                    _ = v61
                                                                    _ = v62
                                                                } else {
                                                                    v18 = 978
                                                                    if len(v43) > 0 {
                                                                        v43[0] = (fi25 + (fi23 >> ((42) & 63)))
                                                                    }
                                                                }
                                                                switch 785 {
                                                                case 567:
                                                                    v63 := v5["a"]
                                                                    _ = v63
                                                                case 42:
                                                                    if len(v43) >= 2 {
                                                                        v43 = v43[1:len(v43)]
                                                                    }
                                                                case 128:
                                                                    v64 := GenS9{F0: 937, F1: v2}
                                                                    _ = v64
                                                                default:
                                                                    switch "foo bar" {
                                                                    case "":
                                                                        fmt.Println(v54)
                                                                    case "hello":
                                                                        delete(v3, 574)
                                                                    case "a":
                                                                        if (922 <= 42) {
                                                                            v42 = v42
                                                                            v65 := GenS10{F0: v42}
                                                                            _ = v65
                                                                        }
                                                                    default:
                                                                        switch fi23 {
                                                                        case 970:
                                                                            var v66 float32 = (((v4 * (float32(float64((v4 - v4))) / (0.1 - 3.14159))) * 0.1) + 3.14159)
                                                                            _ = v66
                                                                        case 442:
                                                                            _ = len(v43)
                                                                        case -2147483648:
                                                                            fmt.Println(v1, v1, fi25)
                                                                        default:
                                                                            v26 += v26
                                                                        }
                                                                    }
                                                                }
                                                                var v67 int32 = 1
                                                                _ = v67
                                                            } else {
                                                                if v2 {
                                                                    v68 := v3[42]
                                                                    var v69 int8 = (^(v22))
                                                                    if (0 > genFn0()) {
                                                                        var v70 uint8 = 91
                                                                        _ = v70
                                                                    } else {
                                                                        if (false && (!(v2 && v2))) {
                                                                            fmt.Println(v1, v68, v68)
                                                                        } else {
                                                                            var v71 uint64 = v42
                                                                            var v72 float32 = 1.5
                                                                            _ = v71
                                                                            _ = v72
                                                                        }
                                                                    }
                                                                    if v2 {
                                                                        v73 := GenS11{F0: 961, F1: (v18 & (v42 | v42)), F2: (-(v69))}
                                                                        v74 := GenS7{F0: v1}
                                                                        _ = v73
                                                                        _ = v74
                                                                    } else {
                                                                        fmt.Println(v68, v22, v18)
                                                                    }
                                                                    v18 += 18446744073709551615
                                                                    v54 = int(v26)
                                                                    _ = v68
                                                                    _ = v69
                                                                }
                                                                if ("\n\t" <= "\n\t") {
                                                                    v75 := GenS7{F0: v1}
                                                                    var v76 int16 = (v1 ^ v1)
                                                                    _ = v75
                                                                    _ = v76
                                                                } else {
                                                                    switch (v54 + (fi25 / ((32768) | 1))) {
                                                                    case -285:
                                                                        if (!((42 == 255) || false)) {
                                                                            v77 := GenS12{F0: 51, F1: (!((1e10 / -2.25) >= (-2.25 + 1.5))), F2: 2147483648}
                                                                            _ = v77
                                                                        }
                                                                    case 32768:
                                                                        var v78 int64 = v26
                                                                        _ = v78
                                                                    default:
                                                                        var v79 uint64 = (v42 - v18)
                                                                        switch (fi23 % ((v54) | 1)) {
                                                                        case 9223372036854775807:
                                                                            v80 := GenS13{F0: genFn0(), F1: "foo bar"}
                                                                            _ = v80
                                                                        case -565:
                                                                            switch 0 {
                                                                            case -11:
                                                                                var v81 float32 = 1.5
                                                                                _ = v81
                                                                            case -605:
                                                                                if len(v43) > 0 {
                                                                                    v43[0] = (fi23 | v54)
                                                                                }
                                                                            }
                                                                        default:
                                                                            fmt.Println(v1, v4)
                                                                        }
                                                                        v82 := GenS14{F0: "\n\t", F1: 32767}
                                                                        _ = v79
                                                                        _ = v82
                                                                    }
                                                                    if false {
                                                                        v1 = v1
                                                                        fmt.Println(fi25)
                                                                    }
                                                                    fmt.Println(v26, v54)
                                                                }
                                                                var v83 uint16 = 970
                                                                _ = v83
                                                            }
                                                            if (v54 <= fi25) {
                                                                switch int(v22) {
                                                                case -1:
                                                                    v84 := GenS1{F0: 179, F1: 65535}
                                                                    _ = v84
                                                                case 545:
                                                                    v85 := GenS12{F0: 499, F1: (v2 && v2), F2: 1}
                                                                    var v86 int32 = 0
                                                                    _ = len(v5)
                                                                    _ = v85
                                                                    _ = v86
                                                                case 208:
                                                                    if v2 {
                                                                        fmt.Println(v54, v54)
                                                                        v2 = false
                                                                    }
                                                                default:
                                                                    switch ("hello" + "\n\t") {
                                                                    case "":
                                                                        v18 += (v18 ^ (v18 % ((v18) | 1)))
                                                                    case "foo bar":
                                                                        var v87 int = (fi25 / ((v54) | 1))
                                                                        _ = v87
                                                                    default:
                                                                        var v88 int64 = 1
                                                                        _ = v88
                                                                    }
                                                                }
                                                            } else {
                                                                v42 -= v42
                                                                if v2 {
                                                                    v89 := GenS15{F0: 32768, F1: 558, F2: "hello"}
                                                                    v54 = (fi25 * (fi23 - -1))
                                                                    v90 := GenS11{F0: genFn0(), F1: 42, F2: (^(v22))}
                                                                    delete(v3, 422)
                                                                    var v91 uint16 = 1
                                                                    _ = v89
                                                                    _ = v90
                                                                    _ = v91
                                                                } else {
                                                                    var v92 string = "\n\t"
                                                                    _ = v92
                                                                }
                                                            }
                                                        }
                                                        if v2 {
                                                            if (((v26 << ((18446744073709551615) & 63)) >= v26) && v2) {
                                                                fmt.Println(v4)
                                                                var v93 uint = 1
                                                                fmt.Println(v4, v4)
                                                                _ = v93
                                                            }
                                                            v22 = 1
                                                        }
                                                        v2 = v2
                                                        _ = v43
                                                        _ = v54
                                                    }
                                                    _ = v42
                                                }
                                                _ = v40
                                                _ = v41
                                            } else {
                                                if v2 {
                                                    fmt.Println(v1)
                                                } else {
                                                    var v94 uint16 = 42
                                                    v1 -= v1
                                                    v95 := GenS10{F0: v18}
                                                    _ = v94
                                                    _ = v95
                                                }
                                                var v96 int16 = v1
                                                _ = v96
                                            }
                                        case "a":
                                            switch (fi25 - (fi25 & (fi23 % ((fi25) | 1)))) {
                                            case 866:
                                                if (675 > 535) {
                                                    v97 := GenS16{F0: v2, F1: v22, F2: false}
                                                    var v98 float32 = v4
                                                    _ = v97
                                                    _ = v98
                                                }
                                            }
                                        default:
                                            switch 32767 {
                                            case 128:
                                                var v99 uint32 = 2147483647
                                                _ = v99
                                            }
                                        }
                                        v2 = (false || false)
                                        fmt.Println(fi23, v18)
                                        v22 = 100
                                        _ = v26
                                    } else {
                                        v100 := v3[2147483647]
                                        v101 := GenS17{F0: -80}
                                        var v102 uint8 = 2
                                        var v103 int64 = 0
                                        v104 := []int{v100, (^(v100))}
                                        _ = v100
                                        _ = v101
                                        _ = v102
                                        _ = v103
                                        _ = v104
                                    }
                                    v2 = v2
                                    _ = fi25
                                }
                                var v105 uint32 = 2147483647
                                if v2 {
                                    v106 := GenS18{F0: 127, F1: 42}
                                    v105 = (v105 + (v105 ^ (^(v105))))
                                    _ = v106
                                } else {
                                    v107 := v3[(^(v105))]
                                    _ = len(v5)
                                    v108 := []int{9223372036854775807}
                                    _ = v107
                                    _ = v108
                                }
                                _ = v105
                            }
                            _ = fi23
                        }
                        var v109 int16 = (v1 * v1)
                        _ = v109
                    } else {
                        var v110 int8 = (v22 * v22)
                        _ = v110
                    }
                    _ = v22
                } else {
                    v5[""] = 42
                    v2 = (!(-48 >= 42))
                    switch -1 {
                    case -2147483648:
                        v4 -= 3.14159
                    case -890:
                        fmt.Println(v4)
                    case 9223372036854775807:
                        v111 := v5[("hello" + "foo bar")]
                        _ = v111
                    default:
                        _ = len(v3)
                    }
                }
                v112 := GenS1{F0: 0, F1: 42}
                _ = v112
            } else {
                if true {
                    for fi113 := 0; fi113 < 8; fi113++ {
                        for fi114 := 0; fi114 < 10; fi114++ {
                            if v2 {
                                v18 = (v18 % ((v18) | 1))
                            }
                            v18 = (v18 >> ((255) & 63))
                            _ = fi114
                        }
                        for fi115 := 0; fi115 < 9; fi115++ {
                            fmt.Println(v2, fi115, fi113)
                            var v116 int16 = v1
                            v1 -= (v1 & (v116 ^ v116))
                            _ = fi115
                            _ = v116
                        }
                        _ = fi113
                    }
                }
                v117 := []int{128}
                v4 += 1.5
                _ = v117
            }
            for fi118 := 0; fi118 < 10; fi118++ {
                switch 0 {
                case -108:
                    v18 += (v18 + v18)
                case 32767:
                    for fi119 := 0; fi119 < 3; fi119++ {
                        v2 = false
                        var v120 uint8 = 2
                        switch 32768 {
                        case 993:
                            var v121 uint32 = 65535
                            _ = v121
                        case 128:
                            var v122 uint = 2
                            _ = v122
                        case 127:
                            v123 := []int{(fi119 ^ 127), fi118, (-(fi119)), (fi118 + (fi119 << ((65535) & 63)))}
                            _ = v123
                        }
                        var v124 uint = uint(v120)
                        _ = fi119
                        _ = v120
                        _ = v124
                    }
                case 721:
                    switch fi118 {
                    case 2147483647:
                        v2 = true
                    case 708:
                        var v125 int32 = -267
                        _ = v125
                    case 128:
                        v2 = (genFn0() > uint(v18))
                    default:
                        for fi126 := 0; fi126 < 6; fi126++ {
                            var v127 int64 = 31
                            v1 -= v1
                            _ = fi126
                            _ = v127
                        }
                    }
                default:
                    for fi128 := 0; fi128 < 16; fi128++ {
                        v2 = (false || v2)
                        _ = fi128
                    }
                }
                switch "" {
                case "hello":
                    for fi129 := 0; fi129 < 10; fi129++ {
                        var v130 uint16 = 177
                        v131 := GenS11{F0: 61, F1: v18, F2: -76}
                        fmt.Println(fi118, fi118, v18)
                        var v132 uint8 = 0
                        if (-2147483648 > 2147483647) {
                            v18 = (v18 | (v18 * (v18 | v18)))
                            v18 += v18
                            var v133 float64 = 0.1
                            _ = v133
                        }
                        _ = fi129
                        _ = v130
                        _ = v131
                        _ = v132
                    }
                case "":
                    v18 = v18
                case "foo bar":
                    fmt.Println(v18)
                default:
                    for fi134 := 0; fi134 < 5; fi134++ {
                        if false {
                            var v135 float32 = v4
                            v135 -= 0.1
                            _ = v135
                        } else {
                            v18 += v18
                            _ = len(v5)
                        }
                        if v2 {
                            v2 = (true && true)
                        } else {
                            v136 := GenS10{F0: 2}
                            if v2 {
                                switch "foo bar" {
                                case "\n\t":
                                    v4 += (v4 - 0.0)
                                }
                            }
                            v137 := []int{int(v1), int(v1)}
                            _ = v136
                            _ = v137
                        }
                        _ = fi134
                    }
                }
                if v2 {
                    v4 -= v4
                    if ((!(42 == 0)) || (v2 && (!v2))) {
                        fmt.Println(v2, v2)
                    } else {
                        var v138 uint64 = v18
                        _ = v138
                    }
                } else {
                    switch "foo bar" {
                    case "":
                        var v139 uint8 = 128
                        _ = v139
                    default:
                        v140 := []int{}
                        _ = v140
                    }
                }
                _ = fi118
            }
        }
    case 2147483647:
        v18 += v18
    case 687:
        switch "hello" {
        case "\n\t":
            var v141 uint = 822
            _ = v141
        case "foo bar":
            var v142 uint8 = 1
            _ = v142
        case "hello":
            v4 = -2.25
        }
    }
    if false {
        v143 := []int{0, 42, 289, 128}
        if v2 {
            v144 := GenS17{F0: -1}
            for fi145 := 0; fi145 < 3; fi145++ {
                var v146 int64 = -658
                _ = fi145
                _ = v146
            }
            _ = v144
        } else {
            if (378 == 2) {
                if false {
                    fmt.Println(v18)
                }
            } else {
                v4 -= v4
                var v147 uint8 = 1
                if (640 != v18) {
                    fmt.Println(v18, v4)
                    v18 = (v18 >> ((321) & 63))
                    if (v2 || v2) {
                        v148 := GenS15{F0: 0, F1: -762, F2: "a"}
                        _ = v148
                    }
                }
                _ = v147
            }
        }
        for fi149 := 0; fi149 < 11; fi149++ {
            var v150 uint = 42
            _ = fi149
            _ = v150
        }
        _ = v143
    } else {
        if v2 {
            var v151 uint16 = 658
            v2 = v2
            _ = v151
        }
        var v152 uint16 = 32767
        v4 -= (v4 - v4)
        _ = v152
    }
    for fi153 := 0; fi153 < 19; fi153++ {
        v1 -= v1
        _ = fi153
    }
    for fi154 := 0; fi154 < 12; fi154++ {
        v155 := []int{fi154, fi154, fi154}
        switch ("foo bar" + "\n\t") {
        case "":
            if v2 {
                var v156 uint64 = (v18 ^ v18)
                v157 := GenS19{F0: 9223372036854775807, F1: 32767}
                _ = v156
                _ = v157
            } else {
                v18 = (v18 << ((481) & 63))
                for fi158 := 0; fi158 < 13; fi158++ {
                    var v159 uint = 1
                    v18 = (v18 & 261)
                    v160 := GenS16{F0: ((fi158 >> (((v159 + (v159 - (v159 << (((v159 - genFn0())) & 63))))) & 63)) >= fi154), F1: -60, F2: false}
                    _ = fi158
                    _ = v159
                    _ = v160
                }
                v4 = 1e10
            }
        case "foo bar":
            if (18446744073709551615 <= 4294967295) {
                fmt.Println(fi154)
            }
        case "a":
            var v161 uint64 = (v18 + v18)
            _ = v161
        default:
            v162 := GenS20{F0: 18446744073709551615, F1: 487}
            _ = v162
        }
        _ = fi154
        _ = v155
    }
    delete(v5, ("foo bar" + ("a" + "foo bar")))
    fmt.Println(v18, v1)
    for fi163 := 0; fi163 < 14; fi163++ {
        var v164 uint64 = 4294967295
        for fi165 := 0; fi165 < 10; fi165++ {
            var v166 int16 = v1
            _ = fi165
            _ = v166
        }
        _ = fi163
        _ = v164
    }
    for fi167 := 0; fi167 < 14; fi167++ {
        v4 = v4
        for fi168 := 0; fi168 < 6; fi168++ {
            v169 := GenS19{F0: -396, F1: -1}
            _ = fi168
            _ = v169
        }
        v170 := GenS9{F0: 179, F1: v2}
        _ = fi167
        _ = v170
    }
    v1 -= 0
    if (!true) {
        v171 := GenS21{F0: ((("a" + "\n\t") + (("a" + "hello") + "foo bar")) + "a"), F1: (("hello" + "hello") + "foo bar")}
        v172 := GenS19{F0: 42, F1: -948}
        _ = v171
        _ = v172
    }
    fmt.Println(v1, v1)
    v173 := GenS14{F0: "", F1: 255}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v173
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
