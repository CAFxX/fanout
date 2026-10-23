package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 uint; F1 uint16; F2 int64 }
type GenS3 struct { F0 uint8; F1 int8; F2 int32 }
type GenS4 struct { F0 uint; F1 int16; F2 uint16 }
type GenS5 struct { F0 int8; F1 int; F2 uint32 }
type GenS6 struct { F0 int64; F1 int8; F2 uint16 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int32; F1 uint; F2 int8 }
type GenS9 struct { F0 int64; F1 uint16 }
type GenS10 struct { F0 uint16; F1 int16 }
type GenS11 struct { F0 int8; F1 uint64 }
type GenS12 struct { F0 int32 }
type GenS13 struct { F0 string }
type GenS14 struct { F0 int }
type GenS15 struct { F0 uint64; F1 int64; F2 int32 }
type GenS16 struct { F0 uint64 }
type GenS17 struct { F0 uint8; F1 uint64 }
type GenS18 struct { F0 uint32 }
type GenS19 struct { F0 int16; F1 bool; F2 uint }
type GenS20 struct { F0 uint64; F1 uint64 }
type GenS21 struct { F0 int; F1 uint64 }
type GenS22 struct { F0 int8; F1 int32 }
type GenS23 struct { F0 int; F1 uint64; F2 int }
type GenS24 struct { F0 uint16; F1 uint32 }
type GenS25 struct { F0 uint64; F1 uint }
type GenS26 struct { F0 uint16; F1 bool; F2 int16 }
type GenS27 struct { F0 bool; F1 uint64; F2 uint8 }
type GenS28 struct { F0 int32 }
type GenS29 struct { F0 bool; F1 uint16 }
type GenS30 struct { F0 uint; F1 bool }
type GenS31 struct { F0 int8; F1 uint32 }
type GenS32 struct { F0 bool; F1 int64; F2 uint64 }
type GenS33 struct { F0 int8; F1 uint32 }
type GenS34 struct { F0 int16; F1 int16 }
func genFn0(p0 int16, p1 int32, p2 uint32) int64 {
    p0 -= p0
    for fi6 := 0; fi6 < 7; fi6++ {
        for fi7 := 0; fi7 < 10; fi7++ {
            v8 := GenS1{F0: 2, F1: "a"}
            v9 := []int{-326, (fi6 ^ fi7), int(p1)}
            var v10 float64 = 1.5
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
        p2 -= 500
        _ = fi6
    }
    var v11 int32 = 42
    p1 -= (v11 ^ (-(v11)))
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 42
}

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[uint32]int)
    var v3 int = 32768
    var v4 uint8 = 2
    var v5 int = v3
    if (false && false) {
        for fi12 := 0; fi12 < 3; fi12++ {
            fmt.Println(fi12, fi12)
            if false {
                if false {
                    for fi13 := 0; fi13 < 17; fi13++ {
                        v4 = v4
                        if ("" != ("\n\t" + "\n\t")) {
                            fmt.Println(v3)
                            var v14 uint32 = 2147483647
                            if true {
                                v1 += (1e10 + ((0.0 * 1.5) - (v1 * ((0.1 - v1) * -2.25))))
                                var v15 int32 = 0
                                _ = v15
                            } else {
                                if false {
                                    var v16 float64 = 1e10
                                    switch (("" + "") + "hello") {
                                    case "hello":
                                        var v17 uint = 65535
                                        _ = v17
                                    case "foo bar":
                                        fmt.Println(v16, fi12, v4)
                                    }
                                    switch "a" {
                                    case "hello":
                                        var v18 float32 = (0.0 * v1)
                                        _ = v18
                                    case "":
                                        var v19 uint32 = (v14 >> ((65535) & 63))
                                        _ = v19
                                    case "\n\t":
                                        var v20 float32 = v1
                                        _ = v20
                                    }
                                    _ = v16
                                }
                                var v21 int32 = -51
                                _ = v21
                            }
                            var v22 uint16 = 2
                            v4 = (v4 | 127)
                            _ = v14
                            _ = v22
                        } else {
                            v4 += v4
                            delete(v2, 2147483647)
                            fmt.Println(fi13, v1, v4)
                            var v23 uint8 = v4
                            if true {
                                var v24 int32 = 66
                                v24 -= v24
                                if ("a" == "") {
                                    v23 = (^(v23))
                                } else {
                                    v25 := GenS1{F0: 42, F1: ("" + (("hello" + "hello") + ("\n\t" + "foo bar")))}
                                    v26 := GenS2{F0: 255, F1: 255, F2: genFn0(-1, v24, 42)}
                                    _ = v25
                                    _ = v26
                                }
                                v4 = v23
                                fmt.Println(v1, fi13)
                                fmt.Println(v4, fi12)
                                if true {
                                    if ((2147483647 == 65535) && false) {
                                        if true {
                                            switch ("a" + "hello") {
                                            case "\n\t":
                                                var v27 float32 = 0.0
                                                _ = v27
                                            case "a":
                                                if true {
                                                    if true {
                                                        var v28 int = v3
                                                        _ = v28
                                                    } else {
                                                        v29 := []int{(v5 ^ fi13), (v5 + v3), fi13, (fi13 ^ (v3 ^ (fi12 - (v5 | fi13))))}
                                                        _ = v29
                                                    }
                                                } else {
                                                    v30 := GenS2{F0: 18446744073709551615, F1: 855, F2: 42}
                                                    v31 := []int{fi12, v5, (fi13 >> ((873) & 63)), fi12}
                                                    _ = v30
                                                    _ = v31
                                                }
                                            case "hello":
                                                v32 := []int{}
                                                _ = v32
                                            default:
                                                v23 = (v23 << ((42) & 63))
                                                var v33 uint = 213
                                                v23 = 2
                                                _ = v33
                                            }
                                            var v34 int64 = -922
                                            _ = v34
                                        } else {
                                            if (3.14159 != 3.14159) {
                                                if (!true) {
                                                    if (21 == -69) {
                                                        var v35 int16 = 1000
                                                        var v36 float64 = 0.0
                                                        _ = v35
                                                        _ = v36
                                                    } else {
                                                        var v37 uint64 = 530
                                                        _ = v37
                                                    }
                                                    switch "" {
                                                    case "hello":
                                                        _ = len(v2)
                                                    default:
                                                        v38 := GenS3{F0: (v23 & (v23 * v23)), F1: -6, F2: (v24 << ((4294967295) & 63))}
                                                        _ = v38
                                                    }
                                                } else {
                                                    v39 := []int{fi12, fi12, v3, v3}
                                                    v5 = fi13
                                                    _ = v39
                                                }
                                                fmt.Println(v23)
                                            } else {
                                                v1 = v1
                                                var v40 uint16 = 975
                                                _ = v40
                                            }
                                        }
                                        if false {
                                            fmt.Println(fi12)
                                        } else {
                                            v41 := GenS4{F0: 43, F1: 42, F2: 65535}
                                            v24 = (v24 << ((790) & 63))
                                            var v42 bool = true
                                            _ = v41
                                            _ = v42
                                        }
                                    } else {
                                        if (false || (!(!false))) {
                                            v43 := GenS5{F0: -33, F1: fi13, F2: 1}
                                            v1 += (v1 - v1)
                                            if false {
                                                switch "a" {
                                                case "foo bar":
                                                    var v44 int16 = 42
                                                    _ = v44
                                                case "a":
                                                    var v45 string = ("foo bar" + ("\n\t" + "hello"))
                                                    _ = v45
                                                case "\n\t":
                                                    var v46 uint32 = 4294967295
                                                    v5 = (fi13 % ((v5) | 1))
                                                    _ = v46
                                                }
                                                var v47 uint64 = 4294967295
                                                _ = v47
                                            } else {
                                                fmt.Println(fi13, v24)
                                            }
                                            _ = v43
                                        } else {
                                            if true {
                                                v4 = v23
                                                switch v5 {
                                                case 127:
                                                    switch ((("a" + "hello") + ("a" + "a")) + "") {
                                                    case "a":
                                                        v48 := []int{42, -9223372036854775808}
                                                        _ = v48
                                                    default:
                                                        v5 = (^(fi12))
                                                    }
                                                case 42:
                                                    v4 -= v23
                                                default:
                                                    v49 := GenS1{F0: 18446744073709551615, F1: (("" + ("a" + "\n\t")) + "")}
                                                    _ = v49
                                                }
                                                var v50 float64 = ((3.14159 * 1e10) + -2.25)
                                                fmt.Println(v23, v50, v5)
                                                _ = v50
                                            } else {
                                                if true {
                                                    var v51 uint64 = 18446744073709551615
                                                    var v52 uint32 = 2147483648
                                                    v2[(^(v52))] = 4294967295
                                                    v23 += v4
                                                    _ = v51
                                                    _ = v52
                                                }
                                                switch "a" {
                                                case "\n\t":
                                                    v53 := GenS6{F0: -9223372036854775808, F1: -51, F2: 255}
                                                    _ = v53
                                                case "a":
                                                    if ((("\n\t" + ("" + "hello")) + "") >= "hello") {
                                                        var v54 int8 = -18
                                                        var v55 uint64 = uint64(v4)
                                                        v56 := GenS7{F0: ("a" + "foo bar")}
                                                        if (!false) {
                                                            if (0 == 849) {
                                                                v57 := v2[204]
                                                                if (!(!((true || (false || true)) && (32767 != -32768)))) {
                                                                    _ = len(v2)
                                                                    var v58 int64 = -539
                                                                    _ = v58
                                                                }
                                                                if true {
                                                                    var v59 uint8 = v4
                                                                    _ = v59
                                                                }
                                                                _ = v57
                                                            } else {
                                                                if ((!(2 >= 1)) || false) {
                                                                    fmt.Println(v23, fi12, v55)
                                                                    v60 := GenS2{F0: 65535, F1: 255, F2: 885}
                                                                    var v61 uint16 = 42
                                                                    _ = v60
                                                                    _ = v61
                                                                } else {
                                                                    if (-97 < (fi13 << ((403) & 63))) {
                                                                        if (true || ((((false && false) || (42 != 65535)) || false) || false)) {
                                                                            var v62 uint64 = 0
                                                                            v63 := []int{128, (^(fi12)), 474, 265}
                                                                            v62 += v55
                                                                            _ = v62
                                                                            _ = v63
                                                                        }
                                                                        var v64 uint = 2
                                                                        v65 := GenS4{F0: (v64 & 1), F1: 32767, F2: 201}
                                                                        var v66 int = 127
                                                                        _ = v64
                                                                        _ = v65
                                                                        _ = v66
                                                                    }
                                                                }
                                                                switch (fi13 >> ((42) & 63)) {
                                                                case 42:
                                                                    if true {
                                                                        v4 = (v4 * v23)
                                                                    }
                                                                case 32767:
                                                                    var v67 int16 = 626
                                                                    _ = v67
                                                                case -438:
                                                                    v68 := v2[807]
                                                                    _ = v68
                                                                }
                                                                v5 = 32767
                                                            }
                                                        }
                                                        v54 = (-(v54))
                                                        _ = v54
                                                        _ = v55
                                                        _ = v56
                                                    }
                                                case "":
                                                    v1 = 1e10
                                                default:
                                                    v1 = (float32((-2.25 - (3.14159 + 1e10))) / 1.0)
                                                }
                                                if (!true) {
                                                    if true {
                                                        if (1e10 >= 1e10) {
                                                            v24 -= 1
                                                            v5 += v3
                                                            var v69 int32 = (v24 >> ((65535) & 63))
                                                            var v70 uint32 = 1
                                                            var v71 int8 = 5
                                                            if ((v4 + v23) >= v4) {
                                                                if false {
                                                                    v71 += (v71 * (v71 << ((377) & 63)))
                                                                }
                                                                var v72 int8 = (^(v71))
                                                                v73 := GenS3{F0: v4, F1: v71, F2: 1}
                                                                _ = v72
                                                                _ = v73
                                                            }
                                                            if false {
                                                                var v74 uint64 = 807
                                                                v70 = v70
                                                                _ = v74
                                                            } else {
                                                                switch fi13 {
                                                                case 128:
                                                                    if (!(!false)) {
                                                                        v75 := GenS2{F0: 18446744073709551615, F1: 386, F2: 9223372036854775807}
                                                                        var v76 int = (fi12 << ((508) & 63))
                                                                        _ = v75
                                                                        _ = v76
                                                                    } else {
                                                                        if false {
                                                                            _ = len(v2)
                                                                            if false {
                                                                                delete(v2, v70)
                                                                            } else {
                                                                                if false {
                                                                                    v77 := GenS5{F0: v71, F1: (fi13 + 127), F2: v70}
                                                                                    if ((!(!true)) && (!false)) {
                                                                                        v2[v70] = (fi12 % (((fi12 - fi13)) | 1))
                                                                                        v2[(v70 & v70)] = (fi12 % ((fi13) | 1))
                                                                                        var v78 int64 = -1
                                                                                        _ = v78
                                                                                    } else {
                                                                                        v69 = 42
                                                                                        var v79 int8 = int8(v70)
                                                                                        _ = v79
                                                                                    }
                                                                                    v5 -= fi13
                                                                                    _ = v77
                                                                                }
                                                                                if true {
                                                                                    v80 := []int{v3, fi12, fi12, v5}
                                                                                    var v81 int8 = -7
                                                                                    switch (^(v3)) {
                                                                                    case 9223372036854775807:
                                                                                        var v82 uint16 = uint16(v70)
                                                                                        _ = v82
                                                                                    case 863:
                                                                                        v83 := GenS3{F0: (v23 * (v23 & v4)), F1: (v71 ^ v81), F2: 42}
                                                                                        _ = v83
                                                                                    }
                                                                                    _ = v80
                                                                                    _ = v81
                                                                                } else {
                                                                                    v84 := v2[v70]
                                                                                    _ = v84
                                                                                }
                                                                            }
                                                                        } else {
                                                                            var v85 string = (("a" + "a") + ("" + "a"))
                                                                            fmt.Println(fi13)
                                                                            fmt.Println(v3, v3, v5)
                                                                            var v86 uint16 = 42
                                                                            var v87 float64 = ((((0.1 / 1.5) / 1e10) - 0.1) * ((1.5 * 1e10) * 0.1))
                                                                            _ = v85
                                                                            _ = v86
                                                                            _ = v87
                                                                        }
                                                                    }
                                                                case 971:
                                                                    var v88 int16 = 1000
                                                                    _ = v88
                                                                default:
                                                                    v89 := GenS5{F0: (^(v71)), F1: (v5 & -247), F2: v70}
                                                                    _ = v89
                                                                }
                                                            }
                                                            if false {
                                                                var v90 uint = 65535
                                                                v91 := GenS8{F0: 1, F1: (v90 - 1), F2: v71}
                                                                switch 1 {
                                                                case -9223372036854775808:
                                                                    var v92 int16 = 0
                                                                    _ = v92
                                                                case 127:
                                                                    fmt.Println(v4, v3)
                                                                }
                                                                _ = v90
                                                                _ = v91
                                                            }
                                                            _ = v69
                                                            _ = v70
                                                            _ = v71
                                                        } else {
                                                            v93 := []int{(fi12 + v5), (v3 / ((-2147483648) | 1)), (v5 * v3)}
                                                            if (42 != 32767) {
                                                                _ = len(v2)
                                                            }
                                                            _ = v93
                                                        }
                                                        v23 += v4
                                                        fmt.Println(v1, v5)
                                                    }
                                                    switch fi12 {
                                                    case 127:
                                                        v94 := GenS5{F0: 127, F1: int(v23), F2: 42}
                                                        _ = v94
                                                    case 32768:
                                                        v4 += (v23 / ((v4) | 1))
                                                    }
                                                } else {
                                                    v95 := GenS9{F0: genFn0(-32768, (v24 + v24), uint32(v23)), F1: 42}
                                                    v3 = v5
                                                    v1 -= (0.0 * v1)
                                                    if (3.14159 > 1e10) {
                                                        fmt.Println(v4, v24)
                                                        if (("" + "\n\t") != "foo bar") {
                                                            if false {
                                                                if false {
                                                                    if ((false && false) || true) {
                                                                        if true {
                                                                            v3 -= v5
                                                                        } else {
                                                                            v23 = (v4 | v23)
                                                                        }
                                                                        v2[1] = (fi13 | 4294967295)
                                                                        v96 := GenS6{F0: 0, F1: -128, F2: 1}
                                                                        fmt.Println(v23, v5, fi12)
                                                                        v97 := v2[42]
                                                                        _ = v96
                                                                        _ = v97
                                                                    }
                                                                } else {
                                                                    if false {
                                                                        if false {
                                                                            fmt.Println(v4)
                                                                            if (42 >= genFn0(422, v24, 0)) {
                                                                                var v98 uint64 = 4294967295
                                                                                var v99 bool = false
                                                                                _ = v98
                                                                                _ = v99
                                                                            } else {
                                                                                fmt.Println(fi12, v1)
                                                                                v5 -= (v5 >> ((895) & 63))
                                                                            }
                                                                        }
                                                                        v100 := []int{}
                                                                        _ = v100
                                                                    }
                                                                    v101 := GenS6{F0: 0, F1: 42, F2: 42}
                                                                    v102 := GenS10{F0: 536, F1: -200}
                                                                    fmt.Println(v3, v1, v5)
                                                                    if ((42 <= (v4 & (^(v4)))) && ((v3 % ((128) | 1)) <= 730)) {
                                                                        v4 = (v23 >> ((18446744073709551615) & 63))
                                                                        var v103 int32 = 42
                                                                        _ = v103
                                                                    }
                                                                    _ = v101
                                                                    _ = v102
                                                                }
                                                                if (int8(v24) != 127) {
                                                                    v4 = v4
                                                                    var v104 uint = 1
                                                                    var v105 float32 = v1
                                                                    _ = v104
                                                                    _ = v105
                                                                }
                                                                fmt.Println(v24, v3, fi12)
                                                            }
                                                            switch ((("hello" + "hello") + ("" + "\n\t")) + "") {
                                                            case "":
                                                                var v106 uint32 = 65535
                                                                _ = v106
                                                            case "a":
                                                                switch 2147483647 {
                                                                case -980:
                                                                    fmt.Println(fi13, v5, v23)
                                                                case 625:
                                                                    if false {
                                                                        fmt.Println(fi13, v4, fi12)
                                                                        var v107 int32 = 42
                                                                        var v108 uint = 865
                                                                        _ = v107
                                                                        _ = v108
                                                                    }
                                                                case -762:
                                                                    v109 := GenS2{F0: 2, F1: 0, F2: 675}
                                                                    _ = v109
                                                                }
                                                            default:
                                                                v110 := GenS10{F0: 999, F1: -1}
                                                                _ = v110
                                                            }
                                                        }
                                                    } else {
                                                        var v111 int16 = 0
                                                        if true {
                                                            fmt.Println(fi13, fi13)
                                                            var v112 string = ((("foo bar" + "foo bar") + (("" + "hello") + "a")) + "hello")
                                                            _ = v112
                                                        } else {
                                                            var v113 int16 = 42
                                                            var v114 int32 = v24
                                                            _ = v113
                                                            _ = v114
                                                        }
                                                        _ = v111
                                                    }
                                                    delete(v2, 2147483648)
                                                    _ = v95
                                                }
                                            }
                                        }
                                        v5 += -151
                                        var v115 uint64 = 1
                                        var v116 uint = 222
                                        _ = v115
                                        _ = v116
                                    }
                                    v5 -= (v5 % ((fi12) | 1))
                                    if false {
                                        var v117 bool = false
                                        v4 += 42
                                        _ = v117
                                    } else {
                                        var v118 uint16 = 402
                                        v119 := GenS10{F0: (v118 << ((2) & 63)), F1: -32768}
                                        switch fi13 {
                                        case -570:
                                            v5 = v3
                                        case 4294967295:
                                            if (55 >= 73) {
                                                var v120 uint = 255
                                                switch ("hello" + "foo bar") {
                                                case "hello":
                                                    if false {
                                                        v3 += 9223372036854775807
                                                        v121 := []int{v3, v5, (fi13 << (((v120 % ((0) | 1))) & 63))}
                                                        _ = v121
                                                    } else {
                                                        v122 := GenS9{F0: 0, F1: v118}
                                                        var v123 uint64 = 989
                                                        switch (v3 * (^(v5))) {
                                                        case -900:
                                                            if ("hello" < (((("a" + "hello") + "hello") + "\n\t") + "a")) {
                                                                var v124 uint = (^(v120))
                                                                _ = len(v2)
                                                                v125 := GenS1{F0: (v124 / ((v124) | 1)), F1: ""}
                                                                _ = v124
                                                                _ = v125
                                                            }
                                                        case 328:
                                                            if (((((v1 + 1e10) * v1) / v1) * float32((0.1 + float64((v1 / 3.14159))))) > 0.0) {
                                                                switch (v3 << ((566) & 63)) {
                                                                case 32768:
                                                                    switch "hello" {
                                                                    case "\n\t":
                                                                        if false {
                                                                            v23 = v4
                                                                        } else {
                                                                            v126 := []int{v5}
                                                                            _ = v126
                                                                        }
                                                                    case "foo bar":
                                                                        var v127 uint8 = 2
                                                                        _ = v127
                                                                    default:
                                                                        v128 := GenS5{F0: -1, F1: v5, F2: 2147483648}
                                                                        _ = v128
                                                                    }
                                                                case 2147483647:
                                                                    var v129 float32 = float32(1.5)
                                                                    if (!true) {
                                                                        v130 := []int{-1, fi13, (v3 / ((fi12) | 1))}
                                                                        switch (((("a" + "\n\t") + ("foo bar" + "")) + "a") + "") {
                                                                        case "foo bar":
                                                                            var v131 int32 = v24
                                                                            _ = v131
                                                                        case "\n\t":
                                                                            var v132 uint64 = 42
                                                                            _ = v132
                                                                        }
                                                                        _ = v130
                                                                    } else {
                                                                        v23 = (v4 % ((v4) | 1))
                                                                        fmt.Println(fi13, v5)
                                                                    }
                                                                    if true {
                                                                        v133 := GenS2{F0: v120, F1: uint16(v123), F2: 42}
                                                                        _ = v133
                                                                    } else {
                                                                        v134 := GenS3{F0: 255, F1: -128, F2: -1}
                                                                        _ = v134
                                                                    }
                                                                    _ = v129
                                                                case 901:
                                                                    if (884 < -1) {
                                                                        if true {
                                                                            v135 := []int{fi13, (v3 & (fi12 - fi12)), v5, v5}
                                                                            switch "a" {
                                                                            case "\n\t":
                                                                                if false {
                                                                                    var v136 float64 = ((1.5 * float64(v1)) / (((1e10 * 3.14159) * (-2.25 * 1e10)) * 3.14159))
                                                                                    if true {
                                                                                        if true {
                                                                                            if (v123 < (v123 << (((^(v120))) & 63))) {
                                                                                                delete(v2, 141)
                                                                                                if false {
                                                                                                    if len(v135) >= 2 {
                                                                                                        v135 = v135[1:len(v135)]
                                                                                                    }
                                                                                                    if (!((v23 - v23) < (v4 * (^(v4))))) {
                                                                                                        v137 := GenS11{F0: 42, F1: v123}
                                                                                                        if (-1 <= -450) {
                                                                                                            if true {
                                                                                                                v118 -= 32767
                                                                                                                v3 -= v3
                                                                                                                if (v136 == float64(1.5)) {
                                                                                                                    switch "" {
                                                                                                                    case "hello":
                                                                                                                        v3 = v3
                                                                                                                        fmt.Println(fi12, v1)
                                                                                                                        v138 := v2[1]
                                                                                                                        _ = v138
                                                                                                                    case "foo bar":
                                                                                                                        v139 := GenS9{F0: -1, F1: (v118 / (((^(v118))) | 1))}
                                                                                                                        _ = v139
                                                                                                                    case "\n\t":
                                                                                                                        if (!(-1 < 227)) {
                                                                                                                            var v140 uint8 = v4
                                                                                                                            if len(v135) >= 2 {
                                                                                                                                v135 = v135[1:len(v135)]
                                                                                                                            }
                                                                                                                            _ = v140
                                                                                                                        } else {
                                                                                                                            v141 := GenS9{F0: -1, F1: (^(v118))}
                                                                                                                            _ = v141
                                                                                                                        }
                                                                                                                    }
                                                                                                                    v5 -= (v3 << ((2) & 63))
                                                                                                                }
                                                                                                            } else {
                                                                                                                var v142 float32 = float32(v136)
                                                                                                                fmt.Println(v120, v23)
                                                                                                                if false {
                                                                                                                    switch v5 {
                                                                                                                    case 382:
                                                                                                                        v143 := GenS5{F0: -128, F1: fi13, F2: 0}
                                                                                                                        v120 += (v120 ^ v120)
                                                                                                                        _ = v143
                                                                                                                    case -9223372036854775808:
                                                                                                                        v23 += (v4 - v23)
                                                                                                                    case -1:
                                                                                                                        var v144 uint32 = 2147483648
                                                                                                                        _ = v144
                                                                                                                    }
                                                                                                                    v135 = append(v135, fi12)
                                                                                                                }
                                                                                                                _ = v142
                                                                                                            }
                                                                                                            switch (v5 ^ (^(fi13))) {
                                                                                                            case 42:
                                                                                                                if false {
                                                                                                                    var v145 int8 = 1
                                                                                                                    v4 = 20
                                                                                                                    _ = v145
                                                                                                                } else {
                                                                                                                    v3 -= 32767
                                                                                                                    v23 += uint8(v118)
                                                                                                                    v146 := GenS9{F0: 333, F1: (v118 % (((^(v118))) | 1))}
                                                                                                                    _ = v146
                                                                                                                }
                                                                                                            }
                                                                                                            v118 -= (v118 | v118)
                                                                                                        }
                                                                                                        v147 := GenS5{F0: 100, F1: v3, F2: 65535}
                                                                                                        _ = v137
                                                                                                        _ = v147
                                                                                                    } else {
                                                                                                        v3 -= fi12
                                                                                                        v148 := GenS12{F0: v24}
                                                                                                        _ = v148
                                                                                                    }
                                                                                                    var v149 float64 = v136
                                                                                                    _ = v149
                                                                                                }
                                                                                                v120 -= (v120 | uint(v118))
                                                                                            }
                                                                                            if false {
                                                                                                switch v5 {
                                                                                                case 4294967295:
                                                                                                    v150 := GenS2{F0: (v120 | v120), F1: (v118 | uint16(v4)), F2: -203}
                                                                                                    _ = v150
                                                                                                default:
                                                                                                    v24 = v24
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        if ((!(true && false)) && false) {
                                                                                            fmt.Println(fi12, v118, v3)
                                                                                            var v151 uint32 = 1
                                                                                            var v152 float64 = 1e10
                                                                                            _ = v151
                                                                                            _ = v152
                                                                                        }
                                                                                        switch "\n\t" {
                                                                                        case "\n\t":
                                                                                            v2[2] = int(v4)
                                                                                        default:
                                                                                            if (!true) {
                                                                                                var v153 uint16 = (v118 ^ 255)
                                                                                                v23 += (v23 + 127)
                                                                                                v154 := GenS13{F0: ""}
                                                                                                fmt.Println(v118, v153, v120)
                                                                                                _ = v153
                                                                                                _ = v154
                                                                                            } else {
                                                                                                if (false && false) {
                                                                                                    v123 = 634
                                                                                                    fmt.Println(v4, v4)
                                                                                                    if false {
                                                                                                        var v155 int64 = int64(v24)
                                                                                                        _ = v155
                                                                                                    } else {
                                                                                                        v3 += (v3 - v5)
                                                                                                        var v156 int64 = 915
                                                                                                        fmt.Println(v120, v24, v1)
                                                                                                        _ = v156
                                                                                                    }
                                                                                                    switch 128 {
                                                                                                    case 2147483647:
                                                                                                        switch 127 {
                                                                                                        case 32767:
                                                                                                            if (true || false) {
                                                                                                                v2[796] = fi12
                                                                                                                v4 = (v4 ^ (v23 | v4))
                                                                                                            } else {
                                                                                                                v157 := GenS11{F0: 1, F1: (v123 * 42)}
                                                                                                                fmt.Println(v5)
                                                                                                                _ = v157
                                                                                                            }
                                                                                                        case 4294967295:
                                                                                                            if (v3 != (fi12 - v5)) {
                                                                                                                v118 = (^(v118))
                                                                                                                v123 = (v123 | (v123 & uint64(v120)))
                                                                                                                v158 := GenS9{F0: 1, F1: (v118 * uint16(v120))}
                                                                                                                _ = v158
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                                if true {
                                                                                                    v159 := GenS2{F0: 339, F1: 32767, F2: 1}
                                                                                                    v160 := GenS3{F0: v23, F1: -128, F2: (v24 % ((v24) | 1))}
                                                                                                    v120 += (v120 >> (((v120 << (((v120 ^ v120)) & 63))) & 63))
                                                                                                    _ = v159
                                                                                                    _ = v160
                                                                                                }
                                                                                            }
                                                                                            v161 := GenS4{F0: (v120 >> ((v120) & 63)), F1: -161, F2: (v118 + 32767)}
                                                                                            _ = v161
                                                                                        }
                                                                                    }
                                                                                    v123 = v123
                                                                                    _ = v136
                                                                                } else {
                                                                                    v162 := GenS14{F0: (v3 / ((2147483647) | 1))}
                                                                                    _ = v162
                                                                                }
                                                                            }
                                                                            _ = v135
                                                                        }
                                                                        v120 = (v120 * 366)
                                                                        v163 := v2[566]
                                                                        switch ((("\n\t" + "\n\t") + ("foo bar" + "a")) + ("foo bar" + "foo bar")) {
                                                                        case "":
                                                                            var v164 int32 = int32(v163)
                                                                            _ = v164
                                                                        }
                                                                        _ = v163
                                                                    } else {
                                                                        v2[42] = v3
                                                                    }
                                                                default:
                                                                    if false {
                                                                        v165 := GenS6{F0: -698, F1: 10, F2: (v118 & (v118 ^ v118))}
                                                                        if false {
                                                                            v166 := []int{(fi12 / ((v3) | 1))}
                                                                            _ = v166
                                                                        }
                                                                        v167 := GenS15{F0: 4294967295, F1: 112, F2: v24}
                                                                        v123 = uint64(v118)
                                                                        _ = v165
                                                                        _ = v167
                                                                    } else {
                                                                        if (((v3 != v3) && true) && (((v1 <= ((-2.25 / 1.5) * (v1 * v1))) || (false && false)) || (false || (!false)))) {
                                                                            v168 := GenS8{F0: 1, F1: (v120 >> ((uint(v23)) & 63)), F2: -70}
                                                                            var v169 uint8 = v23
                                                                            v170 := GenS2{F0: v120, F1: v118, F2: genFn0(-755, v24, 113)}
                                                                            _ = v168
                                                                            _ = v169
                                                                            _ = v170
                                                                        }
                                                                        v120 = (v120 - 65535)
                                                                    }
                                                                }
                                                                var v171 uint64 = v123
                                                                fmt.Println(v23, v5)
                                                                var v172 uint = 18446744073709551615
                                                                switch (v3 | (fi13 + fi13)) {
                                                                case -9223372036854775808:
                                                                    v120 = 255
                                                                case 0:
                                                                    v172 -= 200
                                                                case 32768:
                                                                    v173 := GenS16{F0: 18446744073709551615}
                                                                    _ = v173
                                                                default:
                                                                    switch 633 {
                                                                    case 32767:
                                                                        v174 := v2[101]
                                                                        _ = v174
                                                                    case 4294967295:
                                                                        var v175 bool = false
                                                                        _ = v175
                                                                    }
                                                                }
                                                                _ = v171
                                                                _ = v172
                                                            } else {
                                                                v176 := GenS17{F0: (v4 / (((v23 ^ 119)) | 1)), F1: (v123 ^ v123)}
                                                                var v177 uint32 = 0
                                                                _ = v176
                                                                _ = v177
                                                            }
                                                        case 4294967295:
                                                            var v178 string = "\n\t"
                                                            _ = v178
                                                        default:
                                                            var v179 int8 = 0
                                                            _ = v179
                                                        }
                                                        _ = v122
                                                        _ = v123
                                                    }
                                                default:
                                                    delete(v2, uint32(v120))
                                                }
                                                switch v5 {
                                                case 42:
                                                    var v180 uint8 = uint8(v118)
                                                    _ = v180
                                                case 2147483647:
                                                    var v181 uint = (^(v120))
                                                    _ = v181
                                                }
                                                _ = v120
                                            }
                                            fmt.Println(v5, v118, v4)
                                            switch "foo bar" {
                                            case "":
                                                var v182 uint = 439
                                                _ = v182
                                            }
                                            var v183 uint8 = (v23 % ((v23) | 1))
                                            switch (v5 - (v3 & v3)) {
                                            case 1:
                                                switch 2147483647 {
                                                case -2147483648:
                                                    if ((42 != 1) || (true && (false || (("hello" + "\n\t") < "hello")))) {
                                                        v5 = v5
                                                        if true {
                                                            delete(v2, 65535)
                                                            var v184 float32 = ((1.5 - -2.25) - (1.5 - float32(1.5)))
                                                            v24 = (v24 - 2147483647)
                                                            var v185 uint16 = v118
                                                            _ = v184
                                                            _ = v185
                                                        }
                                                        var v186 float64 = 3.14159
                                                        _ = v186
                                                    } else {
                                                        _ = len(v2)
                                                    }
                                                case -9223372036854775808:
                                                    var v187 string = ("hello" + "hello")
                                                    fmt.Println(v4, fi13, v3)
                                                    fmt.Println(v187, v24)
                                                    _ = v187
                                                }
                                            }
                                            _ = v183
                                        default:
                                            v5 += (fi13 << ((0) & 63))
                                        }
                                        _ = v118
                                        _ = v119
                                    }
                                } else {
                                    fmt.Println(v4, v5, fi12)
                                }
                                switch "a" {
                                case "a":
                                    fmt.Println(v3, v1, v3)
                                case "foo bar":
                                    if (float32(0.0) <= v1) {
                                        var v188 int8 = -128
                                        _ = v188
                                    } else {
                                        v23 = (v23 << ((1) & 63))
                                        v1 = (v1 - float32(1.5))
                                        fmt.Println(v23)
                                    }
                                    switch 4294967295 {
                                    case 1:
                                        v189 := GenS18{F0: 1}
                                        _ = v189
                                    default:
                                        v1 = v1
                                    }
                                    switch ("a" + "foo bar") {
                                    case "a":
                                        v4 = (v23 << ((90) & 63))
                                    default:
                                        if false {
                                            var v190 string = "a"
                                            if ((94 == 514) && (!(true || false))) {
                                                v191 := v2[2147483648]
                                                v191 = (v3 ^ (fi12 * (fi12 & (v5 & 1))))
                                                if true {
                                                    var v192 int = 2147483647
                                                    var v193 uint32 = 956
                                                    _ = v192
                                                    _ = v193
                                                }
                                                _ = v191
                                            }
                                            _ = v190
                                        }
                                    }
                                default:
                                    fmt.Println(v3, v3, v1)
                                }
                                _ = v24
                            } else {
                                switch fi13 {
                                case 0:
                                    v2[125] = (v3 * (fi13 ^ 2147483647))
                                }
                                v194 := GenS7{F0: ""}
                                v195 := GenS19{F0: -32768, F1: (9223372036854775807 < 255), F2: 0}
                                if (0 == 32767) {
                                    fmt.Println(v23, v1)
                                } else {
                                    var v196 uint16 = 2
                                    _ = v196
                                }
                                if true {
                                    switch "a" {
                                    case "":
                                        switch "\n\t" {
                                        case "hello":
                                            var v197 float32 = v1
                                            _ = v197
                                        case "a":
                                            if false {
                                                _ = len(v2)
                                            } else {
                                                var v198 uint32 = 2147483648
                                                var v199 int64 = 42
                                                v200 := v2[v198]
                                                v5 = (v5 << ((0) & 63))
                                                switch "foo bar" {
                                                case "":
                                                    var v201 float32 = (v1 - 0.0)
                                                    _ = v201
                                                case "foo bar":
                                                    var v202 bool = false
                                                    _ = v202
                                                case "hello":
                                                    fmt.Println(fi13, v200)
                                                default:
                                                    v203 := v2[1]
                                                    _ = v203
                                                }
                                                v2[(v198 / ((v198) | 1))] = 2147483647
                                                if false {
                                                    if false {
                                                        var v204 int64 = genFn0(-691, 1, (^(v198)))
                                                        var v205 uint8 = v23
                                                        _ = v204
                                                        _ = v205
                                                    } else {
                                                        switch (fi13 % (((fi12 * 32767)) | 1)) {
                                                        case 251:
                                                            if false {
                                                                if true {
                                                                    v206 := GenS15{F0: 4294967295, F1: 42, F2: 693}
                                                                    v4 = 255
                                                                    var v207 int8 = 41
                                                                    _ = v206
                                                                    _ = v207
                                                                } else {
                                                                    var v208 int8 = 100
                                                                    var v209 uint16 = 151
                                                                    v210 := GenS6{F0: (v199 ^ v199), F1: 100, F2: v209}
                                                                    _ = v208
                                                                    _ = v209
                                                                    _ = v210
                                                                }
                                                                if (v23 < 0) {
                                                                    delete(v2, (v198 | v198))
                                                                } else {
                                                                    switch "" {
                                                                    case "\n\t":
                                                                        if (-2.25 != 0.0) {
                                                                            switch int(v198) {
                                                                            case 127:
                                                                                if (!(255 == 1)) {
                                                                                    fmt.Println(v1)
                                                                                    var v211 float64 = (0.0 * 0.0)
                                                                                    if (((((("foo bar" + "a") + "foo bar") + "hello") + "") + ("a" + "hello")) <= "hello") {
                                                                                        v212 := GenS6{F0: (v199 % ((v199) | 1)), F1: 18, F2: 255}
                                                                                        if (!true) {
                                                                                            if (121 == 42) {
                                                                                                switch ("foo bar" + "\n\t") {
                                                                                                case "\n\t":
                                                                                                    if ((!false) || true) {
                                                                                                        v3 = fi13
                                                                                                    }
                                                                                                }
                                                                                                v198 -= (v198 << ((255) & 63))
                                                                                            } else {
                                                                                                var v213 int8 = 1
                                                                                                _ = v213
                                                                                            }
                                                                                            var v214 string = "foo bar"
                                                                                            v4 -= v23
                                                                                            _ = v214
                                                                                        }
                                                                                        _ = v212
                                                                                    }
                                                                                    _ = v211
                                                                                } else {
                                                                                    var v215 bool = true
                                                                                    var v216 int16 = -1
                                                                                    var v217 uint16 = 255
                                                                                    v218 := GenS20{F0: 0, F1: 838}
                                                                                    _ = v215
                                                                                    _ = v216
                                                                                    _ = v217
                                                                                    _ = v218
                                                                                }
                                                                            case -1:
                                                                                v4 -= (v4 << ((0) & 63))
                                                                            case 128:
                                                                                var v219 bool = true
                                                                                _ = v219
                                                                            default:
                                                                                delete(v2, (v198 - 63))
                                                                            }
                                                                            v2[65535] = v3
                                                                            var v220 int16 = 42
                                                                            _ = v220
                                                                        }
                                                                    case "foo bar":
                                                                        v23 = (v23 / ((v23) | 1))
                                                                    default:
                                                                        if true {
                                                                            v221 := []int{(-(fi13)), 127, v200}
                                                                            var v222 uint64 = 163
                                                                            switch (("foo bar" + "hello") + "a") {
                                                                            case "a":
                                                                                var v223 float64 = 1.5
                                                                                _ = v223
                                                                            case "":
                                                                                if len(v221) >= 2 {
                                                                                    v221 = v221[1:len(v221)]
                                                                                }
                                                                            }
                                                                            _ = v221
                                                                            _ = v222
                                                                        } else {
                                                                            v4 = v23
                                                                        }
                                                                    }
                                                                }
                                                                if false {
                                                                    v3 = -1
                                                                    v5 = fi13
                                                                    v5 += 127
                                                                } else {
                                                                    v198 += v198
                                                                    if false {
                                                                        var v224 int = fi13
                                                                        v225 := []int{(^(fi13)), -119, v224, (v200 ^ 505)}
                                                                        var v226 float32 = ((v1 * v1) + (1.5 * v1))
                                                                        if false {
                                                                            v200 = (v224 / (((fi12 >> ((4294967295) & 63))) | 1))
                                                                        }
                                                                        _ = v224
                                                                        _ = v225
                                                                        _ = v226
                                                                    } else {
                                                                        fmt.Println(v199, v1, v4)
                                                                        v199 = 541
                                                                    }
                                                                    var v227 int64 = -561
                                                                    _ = v227
                                                                }
                                                                v4 -= (v23 / ((v4) | 1))
                                                            }
                                                        }
                                                        if false {
                                                            var v228 float32 = v1
                                                            var v229 float32 = (v228 / -2.25)
                                                            if true {
                                                                v230 := GenS21{F0: fi12, F1: 1}
                                                                v5 -= -1
                                                                v231 := GenS5{F0: -89, F1: (v5 * 1), F2: uint32(v23)}
                                                                _ = v230
                                                                _ = v231
                                                            }
                                                            _ = v228
                                                            _ = v229
                                                        } else {
                                                            v232 := []int{v5, (^(v3))}
                                                            var v233 int32 = -50
                                                            _ = v232
                                                            _ = v233
                                                        }
                                                        v199 -= v199
                                                        var v234 int64 = -9223372036854775808
                                                        v235 := GenS22{F0: 0, F1: -2147483648}
                                                        if true {
                                                            var v236 uint16 = 56
                                                            delete(v2, (v198 >> ((4294967295) & 63)))
                                                            v2[v198] = fi12
                                                            _ = v236
                                                        }
                                                        var v237 uint16 = 436
                                                        _ = v234
                                                        _ = v235
                                                        _ = v237
                                                    }
                                                    v238 := []int{(v5 ^ (v3 >> ((18446744073709551615) & 63)))}
                                                    v239 := v2[(v198 * 770)]
                                                    _ = v238
                                                    _ = v239
                                                }
                                                _ = v198
                                                _ = v199
                                                _ = v200
                                            }
                                        case "":
                                            var v240 uint32 = 1
                                            _ = v240
                                        default:
                                            v4 += 2
                                        }
                                    default:
                                        if true {
                                            var v241 uint64 = 597
                                            _ = v241
                                        }
                                    }
                                    delete(v2, 0)
                                    v5 = (fi12 & fi13)
                                    if (-9223372036854775808 == -9223372036854775808) {
                                        switch (("" + "") + ((((("foo bar" + "\n\t") + "a") + "hello") + "foo bar") + "foo bar")) {
                                        case "a":
                                            v1 = ((((1.5 - (((float32(1e10) + (float32(1e10) - ((0.1 + ((v1 / v1) * (v1 / -2.25))) - v1))) / (((v1 - ((v1 + 0.1) - 3.14159)) * 1.5) - float32(1e10))) / v1)) - (v1 * 1e10)) - v1) + (((-2.25 - v1) * (float32(0.0) * (v1 * ((float32((((1e10 - ((3.14159 + float64(3.14159)) / 1e10)) * (3.14159 * (0.1 / -2.25))) / 1.0)) * 1e10) + float32(1.5))))) + v1))
                                        case "hello":
                                            v242 := GenS1{F0: 255, F1: (("foo bar" + "a") + "a")}
                                            switch 710 {
                                            case -9223372036854775808:
                                                var v243 uint16 = 32767
                                                _ = v243
                                            default:
                                                switch (v5 ^ (fi12 - fi12)) {
                                                case 4294967295:
                                                    v244 := v2[493]
                                                    _ = v244
                                                case 32767:
                                                    v4 = (v4 - v23)
                                                    v245 := GenS10{F0: 65535, F1: 0}
                                                    v3 = (v3 & v5)
                                                    _ = len(v2)
                                                    var v246 int32 = 0
                                                    v23 -= (^(v23))
                                                    _ = v245
                                                    _ = v246
                                                }
                                            }
                                            v247 := []int{(-(fi12)), (v3 >> ((18446744073709551615) & 63)), (v3 >> ((855) & 63))}
                                            v3 -= v3
                                            _ = v242
                                            _ = v247
                                        case "foo bar":
                                            fmt.Println(v1, fi13)
                                        default:
                                            _ = len(v2)
                                        }
                                        v248 := []int{128, (fi12 * v3)}
                                        switch (^(v5)) {
                                        case 32768:
                                            v249 := GenS18{F0: 2147483648}
                                            _ = v249
                                        case 127:
                                            v4 += (v4 << ((656) & 63))
                                        case -218:
                                            var v250 int16 = -1
                                            _ = v250
                                        default:
                                            if (65535 <= 1) {
                                                v2[0] = fi13
                                                if len(v248) >= 2 {
                                                    v248 = v248[1:len(v248)]
                                                }
                                                v251 := GenS22{F0: 40, F1: 1}
                                                _ = v251
                                            }
                                        }
                                        v252 := GenS23{F0: v5, F1: 4294967295, F2: int(v23)}
                                        v3 += 4294967295
                                        switch "hello" {
                                        case "a":
                                            var v253 uint32 = 539
                                            _ = v253
                                        case "\n\t":
                                            delete(v2, 13)
                                        case "hello":
                                            v5 = (fi13 ^ (fi12 & 4294967295))
                                        }
                                        if (!(0 >= -1)) {
                                            v254 := v2[2]
                                            _ = v254
                                        }
                                        _ = v248
                                        _ = v252
                                    }
                                    var v255 float64 = float64((float32(float64((0.1 * (0.0 * (((v1 - (3.14159 + 1e10)) - v1) - v1))))) + float32(0.1)))
                                    _ = v255
                                } else {
                                    var v256 float32 = ((1e10 / 1.0) + v1)
                                    if true {
                                        var v257 int = (fi12 / ((1) | 1))
                                        switch "" {
                                        case "a":
                                            fmt.Println(v23, v257)
                                        }
                                        _ = v257
                                    } else {
                                        v258 := []int{}
                                        var v259 uint8 = (v4 % (((v4 ^ (v23 << ((1) & 63)))) | 1))
                                        v260 := GenS24{F0: 42, F1: 1}
                                        v261 := GenS25{F0: 0, F1: 65535}
                                        v262 := GenS26{F0: 793, F1: true, F2: -766}
                                        var v263 string = ""
                                        _ = v258
                                        _ = v259
                                        _ = v260
                                        _ = v261
                                        _ = v262
                                        _ = v263
                                    }
                                    _ = v256
                                }
                                _ = v194
                                _ = v195
                            }
                            _ = v23
                        }
                        _ = fi13
                    }
                    v4 = (^(v4))
                    if ("a" != "\n\t") {
                        for fi264 := 0; fi264 < 2; fi264++ {
                            fmt.Println(fi12, v5, v3)
                            if ((("a" + "\n\t") > "\n\t") || (0.1 <= 0.1)) {
                                var v265 int = -593
                                if (("foo bar" + "\n\t") <= "\n\t") {
                                    if true {
                                        switch ("\n\t" + "") {
                                        case "\n\t":
                                            if false {
                                                v5 += v265
                                                var v266 uint8 = 255
                                                _ = v266
                                            } else {
                                                var v267 int32 = -244
                                                if (!true) {
                                                    var v268 int16 = -260
                                                    v5 = -9223372036854775808
                                                    _ = v268
                                                } else {
                                                    var v269 uint8 = v4
                                                    v5 += (v5 * (^(fi12)))
                                                    fmt.Println(fi264, fi12)
                                                    _ = v269
                                                }
                                                v270 := []int{}
                                                _ = v267
                                                _ = v270
                                            }
                                        }
                                        var v271 float64 = (float64((v1 - (v1 * 1.5))) * -2.25)
                                        v272 := GenS27{F0: true, F1: 18446744073709551615, F2: (^(v4))}
                                        v273 := []int{(v3 / (((v5 >> ((4294967295) & 63))) | 1)), (v5 - v5)}
                                        _ = v271
                                        _ = v272
                                        _ = v273
                                    }
                                    fmt.Println(fi264)
                                    v4 += (v4 & v4)
                                    v274 := []int{(v265 >> ((666) & 63)), fi12}
                                    _ = len(v274)
                                    _ = v274
                                } else {
                                    v275 := GenS1{F0: 18446744073709551615, F1: "hello"}
                                    var v276 int8 = -80
                                    _ = v275
                                    _ = v276
                                }
                                v265 += (-(v3))
                                _ = v265
                            } else {
                                v4 = v4
                            }
                            v4 += v4
                            _ = fi264
                        }
                    } else {
                        v1 += ((v1 / v1) - v1)
                        for fi277 := 0; fi277 < 11; fi277++ {
                            var v278 int8 = 60
                            _ = fi277
                            _ = v278
                        }
                        if true {
                            delete(v2, 2147483647)
                        } else {
                            var v279 int = fi12
                            _ = v279
                        }
                    }
                }
                fmt.Println(fi12, v4)
                v280 := GenS28{F0: -2147483648}
                _ = v280
            } else {
                var v281 int64 = 91
                v281 = (v281 << ((255) & 63))
                fmt.Println(v5, fi12, v281)
                _ = v281
            }
            _ = fi12
        }
    } else {
        v282 := GenS25{F0: 0, F1: 4294967295}
        for fi283 := 0; fi283 < 2; fi283++ {
            v284 := []int{v3}
            for fi285 := 0; fi285 < 2; fi285++ {
                fmt.Println(v4, v1, v5)
                _ = fi285
            }
            v3 += fi283
            _ = fi283
            _ = v284
        }
        _ = v282
    }
    var v286 string = (("a" + "") + "a")
    if false {
        v5 -= 9223372036854775807
        v2[65535] = v3
        fmt.Println(v286, v3, v286)
    } else {
        var v287 uint32 = 141
        _ = v287
    }
    fmt.Println(v286)
    var v288 int = v3
    for fi289 := 0; fi289 < 12; fi289++ {
        v286 = v286
        _ = fi289
    }
    v290 := GenS29{F0: (4294967295 == 189), F1: 255}
    fmt.Println(v5, v1)
    v288 = (v5 & v5)
    var v291 int16 = 771
    switch (v286 + v286) {
    case "foo bar":
        for fi292 := 0; fi292 < 15; fi292++ {
            if false {
                var v293 float64 = 3.14159
                if false {
                    var v294 uint32 = 2
                    v295 := GenS30{F0: 18446744073709551615, F1: ((v291 % ((v291) | 1)) <= v291)}
                    v286 = v286
                    _ = v294
                    _ = v295
                } else {
                    v5 += v5
                }
                _ = v293
            } else {
                v296 := GenS31{F0: 112, F1: 944}
                _ = v296
            }
            _ = fi292
        }
    case "":
        var v297 int8 = 42
        _ = v297
    case "hello":
        var v298 int64 = 0
        _ = v298
    }
    if false {
        var v299 int8 = -1
        var v300 int = 2147483647
        _ = v299
        _ = v300
    } else {
        var v301 uint8 = v4
        var v302 int64 = 9223372036854775807
        _ = v301
        _ = v302
    }
    var v303 int = (-(v3))
    v304 := GenS24{F0: uint16(v4), F1: 4294967295}
    if (!(true || false)) {
        for fi305 := 0; fi305 < 9; fi305++ {
            if ((3.14159 != (0.0 * ((float64(1.5) + 0.0) * 1e10))) && (v288 > v288)) {
                if (true && (0.0 <= (0.0 + (3.14159 * v1)))) {
                    if (((fi305 != v288) || false) || (false && (true || (2147483647 > 2147483647)))) {
                        v306 := GenS32{F0: (!false), F1: -287, F2: 18446744073709551615}
                        _ = v306
                    }
                } else {
                    var v307 int64 = -1
                    switch (v3 >> ((2) & 63)) {
                    case 42:
                        var v308 float64 = -2.25
                        _ = v308
                    case 1:
                        _ = len(v2)
                    case -1:
                        v3 += (^(v5))
                    default:
                        switch "hello" {
                        case "foo bar":
                            v309 := GenS9{F0: int64(v288), F1: 255}
                            _ = v309
                        case "\n\t":
                            fmt.Println(v307, v291)
                        case "hello":
                            v310 := GenS28{F0: 2147483647}
                            _ = v310
                        default:
                            switch v3 {
                            case -193:
                                for fi311 := 0; fi311 < 9; fi311++ {
                                    v312 := GenS25{F0: 42, F1: 42}
                                    v313 := GenS26{F0: 159, F1: (100 == -128), F2: v291}
                                    _ = fi311
                                    _ = v312
                                    _ = v313
                                }
                            }
                        }
                    }
                    _ = v307
                }
                v314 := []int{v288, v303}
                var v315 int64 = -9223372036854775808
                _ = v314
                _ = v315
            } else {
                v316 := GenS33{F0: 0, F1: 2}
                fmt.Println(v291, v1)
                var v317 float32 = 3.14159
                _ = v316
                _ = v317
            }
            v318 := GenS34{F0: (v291 | v291), F1: (v291 % ((v291) | 1))}
            v1 = (v1 * ((v1 - float32(0.1)) / (v1 * ((1.5 + float32(0.1)) + v1))))
            _ = fi305
            _ = v318
        }
        fmt.Println(v288, v3, v291)
    }
    var v319 bool = (!false)
    var v320 uint8 = (v4 >> ((0) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v286:", v286)
    fmt.Println("v288:", v288)
    fmt.Println("v291:", v291)
    fmt.Println("v3:", v3)
    fmt.Println("v303:", v303)
    fmt.Println("v319:", v319)
    fmt.Println("v320:", v320)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v286
    _ = v288
    _ = v290
    _ = v291
    _ = v3
    _ = v303
    _ = v304
    _ = v319
    _ = v320
    _ = v4
    _ = v5
}
