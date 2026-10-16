package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint16; F1 int32; F2 int }
type GenS5 struct { F0 uint16; F1 int }
type GenS6 struct { F0 uint16; F1 bool }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int8; F1 uint8 }
type GenS9 struct { F0 int16; F1 uint64 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 int64; F1 int64; F2 uint }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 bool; F1 uint32; F2 int }
type GenS14 struct { F0 int; F1 int; F2 int32 }
type GenS15 struct { F0 string; F1 string; F2 uint16 }
type GenS16 struct { F0 uint16; F1 uint; F2 uint8 }
type GenS17 struct { F0 uint8 }
type GenS18 struct { F0 int64; F1 int32 }
type GenS19 struct { F0 int32 }
type GenS20 struct { F0 uint64; F1 int; F2 int16 }
type GenS21 struct { F0 int16 }
type GenS22 struct { F0 uint16; F1 uint16; F2 uint8 }
type GenS23 struct { F0 bool; F1 string }
type GenS24 struct { F0 int32; F1 int32; F2 string }
type GenS25 struct { F0 int8; F1 uint32; F2 uint8 }
type GenS26 struct { F0 string }
type GenS27 struct { F0 int8; F1 string }
type GenS28 struct { F0 int64 }
type GenS29 struct { F0 uint32 }
type GenS30 struct { F0 string }
type GenS31 struct { F0 uint; F1 bool; F2 bool }
type GenS32 struct { F0 uint16; F1 uint64; F2 int64 }
type GenS33 struct { F0 int16; F1 int64 }
type GenS34 struct { F0 int8 }
type GenS35 struct { F0 int8; F1 uint64; F2 uint8 }
type GenS36 struct { F0 uint8; F1 string }
type GenS37 struct { F0 uint8 }
type GenS38 struct { F0 uint8; F1 uint }
type GenS39 struct { F0 uint16; F1 int8; F2 int16 }
type GenS40 struct { F0 int32 }
type GenS41 struct { F0 uint8; F1 uint32; F2 int }
type GenS42 struct { F0 int64 }
func genFn0() bool {
    if (0 >= 215) {
        if (float64(0.1) != ((3.14159 * 0.1) + 3.14159)) {
            var v7 bool = (42 > 1)
            v7 = v7
            v8 := GenS1{F0: 729}
            _ = v7
            _ = v8
        } else {
            var v9 string = "hello"
            fmt.Println(v9)
            _ = v9
        }
        var v10 uint32 = 1
        _ = v10
    }
    v11 := GenS1{F0: 18446744073709551615}
    for fi12 := 0; fi12 < 5; fi12++ {
        if (true || false) {
            if false {
                if (!true) {
                    if (("\n\t" + "foo bar") == (((("" + "a") + "foo bar") + "\n\t") + "hello")) {
                        for fi13 := 0; fi13 < 6; fi13++ {
                            var v14 uint32 = 42
                            v15 := GenS1{F0: 2}
                            _ = fi13
                            _ = v14
                            _ = v15
                        }
                    } else {
                        var v16 uint64 = 2
                        var v17 int8 = -37
                        _ = v16
                        _ = v17
                    }
                    var v18 int = fi12
                    _ = v18
                }
            }
        } else {
            var v19 uint8 = 42
            v20 := []int{}
            if (1 > 767) {
                for fi21 := 0; fi21 < 19; fi21++ {
                    v22 := GenS2{F0: 1000, F1: "foo bar"}
                    _ = fi21
                    _ = v22
                }
                v23 := GenS1{F0: 42}
                var v24 int32 = -1
                _ = v23
                _ = v24
            } else {
                if true {
                    fmt.Println(fi12, v19)
                    switch fi12 {
                    case -2147483648:
                        var v25 int = (fi12 >> ((4294967295) & 63))
                        _ = v25
                    case 760:
                        for fi26 := 0; fi26 < 13; fi26++ {
                            v19 -= (v19 & v19)
                            switch ("\n\t" + "\n\t") {
                            case "":
                                var v27 float64 = 1.5
                                _ = v27
                            default:
                                var v28 string = ((((("\n\t" + "foo bar") + "") + ("foo bar" + ((("\n\t" + ("\n\t" + "")) + "foo bar") + "\n\t"))) + "a") + "\n\t")
                                _ = v28
                            }
                            v29 := GenS1{F0: uint64(v19)}
                            v19 -= 2
                            switch "hello" {
                            case "hello":
                                switch (^(fi12)) {
                                case 248:
                                    v30 := GenS2{F0: -1, F1: ("foo bar" + "a")}
                                    _ = v30
                                default:
                                    switch "\n\t" {
                                    case "a":
                                        var v31 string = "a"
                                        _ = v31
                                    case "foo bar":
                                        var v32 uint8 = v19
                                        _ = v32
                                    }
                                }
                            }
                            _ = fi26
                            _ = v29
                        }
                    }
                    switch ("a" + "hello") {
                    case "foo bar":
                        if ((v19 / ((v19) | 1)) == (^(v19))) {
                            if len(v20) > 0 {
                                v20[0] = -268
                            }
                            v19 -= v19
                            v33 := GenS2{F0: -917, F1: ("\n\t" + "a")}
                            _ = v33
                        } else {
                            for fi34 := 0; fi34 < 14; fi34++ {
                                v19 += (v19 + (v19 << ((65535) & 63)))
                                v19 -= (v19 << ((352) & 63))
                                _ = fi34
                            }
                            if len(v20) >= 2 {
                                v20 = v20[1:len(v20)]
                            }
                            v35 := GenS1{F0: 42}
                            _ = v35
                        }
                    case "hello":
                        var v36 int = (fi12 * fi12)
                        _ = v36
                    default:
                        if len(v20) >= 2 {
                            v20 = v20[1:len(v20)]
                        }
                    }
                } else {
                    switch (fi12 / ((fi12) | 1)) {
                    case -9223372036854775808:
                        v19 = v19
                    case -1:
                        _ = len(v20)
                    }
                    fmt.Println(v19, v19)
                    if false {
                        if false {
                            v19 += v19
                            if false {
                                for fi37 := 0; fi37 < 15; fi37++ {
                                    var v38 uint8 = v19
                                    if true {
                                        v38 -= (v19 << ((255) & 63))
                                        if len(v20) >= 2 {
                                            v20 = v20[1:len(v20)]
                                        }
                                        var v39 int64 = 0
                                        _ = v39
                                    } else {
                                        var v40 int32 = 1
                                        _ = v40
                                    }
                                    _ = fi37
                                    _ = v38
                                }
                                var v41 uint = 1
                                for fi42 := 0; fi42 < 13; fi42++ {
                                    if false {
                                        var v43 uint = (^(v41))
                                        if false {
                                            if (int64(v43) <= -591) {
                                                v44 := GenS3{F0: 2}
                                                if ("" >= "\n\t") {
                                                    v43 -= (^(v41))
                                                    var v45 int = (fi42 << (((v41 >> (((v41 << ((18446744073709551615) & 63))) & 63))) & 63))
                                                    _ = v45
                                                }
                                                fmt.Println(v19, v19, v41)
                                                v19 += 154
                                                _ = v44
                                            } else {
                                                if (86 == 127) {
                                                    v43 = v41
                                                }
                                            }
                                            if (36 < 56) {
                                                fmt.Println(v19)
                                            }
                                            if false {
                                                if (!true) {
                                                    if true {
                                                        v19 = (v19 >> (((v43 / ((v43) | 1))) & 63))
                                                        v41 -= uint(v19)
                                                    }
                                                    if len(v20) > 0 {
                                                        v20[0] = fi42
                                                    }
                                                }
                                            } else {
                                                switch (fi42 * (fi42 >> ((v43) & 63))) {
                                                case 442:
                                                    v41 = (^(v43))
                                                case -296:
                                                    v41 = 925
                                                case -1:
                                                    v46 := GenS3{F0: 1}
                                                    _ = v46
                                                }
                                                v41 = 806
                                                if len(v20) >= 2 {
                                                    v20 = v20[1:len(v20)]
                                                }
                                            }
                                        }
                                        if len(v20) > 0 {
                                            v20[0] = (fi12 ^ (fi12 + 42))
                                        }
                                        _ = v43
                                    } else {
                                        fmt.Println(v19)
                                        fmt.Println(v41)
                                    }
                                    var v47 int = (fi12 % ((32767) | 1))
                                    v20 = append(v20, (fi12 & 0))
                                    if (2147483647 == 881) {
                                        v48 := GenS3{F0: 255}
                                        if ((1 > 127) || ((!(float64(1e10) == 0.1)) && (false && (2147483648 > 954)))) {
                                            var v49 int8 = 1
                                            var v50 int32 = -2147483648
                                            _ = v49
                                            _ = v50
                                        } else {
                                            var v51 bool = true
                                            if false {
                                                v41 = (v41 >> (((^(v41))) & 63))
                                                if (v51 || (!(-9223372036854775808 < 42))) {
                                                    var v52 int32 = -26
                                                    _ = v52
                                                } else {
                                                    if (v51 || v51) {
                                                        var v53 uint64 = 1
                                                        if true {
                                                            fmt.Println(v47, v51, v53)
                                                            v54 := GenS4{F0: 1, F1: -172, F2: (fi42 >> ((42) & 63))}
                                                            v19 += (v19 & (v19 ^ (v19 & v19)))
                                                            _ = v54
                                                        } else {
                                                            var v55 int64 = -441
                                                            v56 := GenS4{F0: 564, F1: 42, F2: 519}
                                                            fmt.Println(fi42)
                                                            _ = v55
                                                            _ = v56
                                                        }
                                                        v53 = 18446744073709551615
                                                        _ = v53
                                                    } else {
                                                        v47 += 128
                                                        v19 -= v19
                                                    }
                                                    var v57 int8 = -128
                                                    fmt.Println(v57)
                                                    _ = v57
                                                }
                                                if v51 {
                                                    if (1 != 9223372036854775807) {
                                                        var v58 string = "hello"
                                                        _ = v58
                                                    } else {
                                                        var v59 int = (v47 & 32768)
                                                        _ = v59
                                                    }
                                                } else {
                                                    v60 := GenS2{F0: 42, F1: ("\n\t" + "a")}
                                                    if v51 {
                                                        if v51 {
                                                            if len(v20) >= 2 {
                                                                v20 = v20[1:len(v20)]
                                                            }
                                                        } else {
                                                            v61 := GenS5{F0: 2, F1: (-(v47))}
                                                            v20 = append(v20, v47)
                                                            _ = v61
                                                        }
                                                        if ((!false) || v51) {
                                                            var v62 int8 = 107
                                                            v63 := GenS6{F0: 1, F1: v51}
                                                            v41 = v41
                                                            _ = v62
                                                            _ = v63
                                                        }
                                                    } else {
                                                        v64 := GenS3{F0: 1}
                                                        _ = v64
                                                    }
                                                    _ = len(v20)
                                                    _ = v60
                                                }
                                            } else {
                                                var v65 uint = (v41 & (v41 + (v41 >> ((349) & 63))))
                                                var v66 string = "foo bar"
                                                v19 = v19
                                                _ = v65
                                                _ = v66
                                            }
                                            fmt.Println(v41, fi12, v19)
                                            _ = v51
                                        }
                                        switch fi42 {
                                        case -945:
                                            if true {
                                                v67 := GenS2{F0: 1, F1: "a"}
                                                _ = v67
                                            } else {
                                                v68 := GenS6{F0: 32767, F1: (-73 > 127)}
                                                if true {
                                                    v20 = append(v20, v47)
                                                } else {
                                                    var v69 bool = false
                                                    if (0 >= 32767) {
                                                        v47 -= 42
                                                    }
                                                    fmt.Println(fi42, v69, fi42)
                                                    var v70 uint8 = 198
                                                    v71 := GenS1{F0: 407}
                                                    _ = v69
                                                    _ = v70
                                                    _ = v71
                                                }
                                                var v72 float64 = 3.14159
                                                if (!((1 <= -1) || (false || (!(!(true && false)))))) {
                                                    if false {
                                                        v73 := GenS7{F0: -9223372036854775808}
                                                        _ = v73
                                                    }
                                                    fmt.Println(v41, v41, v41)
                                                    fmt.Println(v72)
                                                }
                                                _ = v68
                                                _ = v72
                                            }
                                        case -1:
                                            var v74 int16 = 32767
                                            _ = v74
                                        case 756:
                                            fmt.Println(fi42)
                                        default:
                                            var v75 int = fi12
                                            _ = v75
                                        }
                                        v47 = 900
                                        _ = v48
                                    }
                                    fmt.Println(v47, v41, v41)
                                    _ = fi42
                                    _ = v47
                                }
                                _ = v41
                            }
                            v19 = v19
                        } else {
                            fmt.Println(fi12, v19)
                            var v76 string = ""
                            _ = v76
                        }
                    }
                }
                v19 = v19
            }
            _ = v19
            _ = v20
        }
        var v77 uint32 = 425
        fmt.Println(fi12, v77)
        _ = fi12
        _ = v77
    }
    _ = v11
    return (129 != 42)
}
func genFn1() int {
    var v78 int16 = 0
    fmt.Println(v78)
    if true {
        for fi79 := 0; fi79 < 20; fi79++ {
            v80 := GenS8{F0: 42, F1: 36}
            fmt.Println(v78)
            _ = fi79
            _ = v80
        }
        for fi81 := 0; fi81 < 6; fi81++ {
            var v82 float64 = 0.0
            _ = fi81
            _ = v82
        }
    } else {
        var v83 int64 = -9223372036854775808
        v78 = (-(v78))
        _ = v83
    }
    _ = v78
    return 32767
}

func main() {
    v1 := []int{127, 127, -1}
    var v2 uint8 = 2
    v3 := []string{("a" + "\n\t"), "a"}
    v4 := make(map[uint32]uint32)
    var v5 uint16 = 42
    var v6 uint64 = 1
    _ = len(v4)
    fmt.Println(v5, v6)
    v6 += (v6 * (v6 % ((0) | 1)))
    for fi84 := 0; fi84 < 15; fi84++ {
        switch (fi84 % ((fi84) | 1)) {
        case 32767:
            for fi85 := 0; fi85 < 18; fi85++ {
                v5 = v5
                switch 2147483647 {
                case 577:
                    _ = len(v4)
                default:
                    var v86 int8 = 127
                    _ = v86
                }
                _ = fi85
            }
        case -49:
            var v87 uint8 = (v2 * v2)
            _ = v87
        case -9223372036854775808:
            for fi88 := 0; fi88 < 12; fi88++ {
                var v89 int64 = 1
                if (true && (-1 > 0)) {
                    _ = len(v4)
                    _ = len(v4)
                    if true {
                        switch "" {
                        case "":
                            v89 = 757
                        }
                        v90 := GenS2{F0: -32768, F1: ""}
                        _ = v90
                    } else {
                        var v91 int8 = -128
                        if false {
                            if ((-2.25 * 1e10) != float64((0.1 + -2.25))) {
                                var v92 float64 = 3.14159
                                if true {
                                    switch (fi88 << ((4294967295) & 63)) {
                                    case 0:
                                        delete(v4, 0)
                                    case -904:
                                        switch genFn1() {
                                        case 2147483647:
                                            v93 := GenS9{F0: -373, F1: (v6 * (v6 >> ((45) & 63)))}
                                            _ = v93
                                        case 0:
                                            if false {
                                                var v94 float64 = (float64(float32(3.14159)) / float64((0.0 - 1e10)))
                                                v95 := GenS8{F0: (v91 % ((v91) | 1)), F1: (v2 ^ v2)}
                                                _ = v94
                                                _ = v95
                                            }
                                        case 923:
                                            var v96 bool = (false && false)
                                            _ = v96
                                        default:
                                            v89 -= v89
                                        }
                                        if true {
                                            v92 += float64(3.14159)
                                            if true {
                                                v6 += v6
                                                var v97 string = "a"
                                                v98 := GenS6{F0: (v5 | 0), F1: false}
                                                _ = v97
                                                _ = v98
                                            } else {
                                                fmt.Println(v91)
                                            }
                                            switch (("hello" + "\n\t") + "foo bar") {
                                            case "\n\t":
                                                if (!(!((false || false) && true))) {
                                                    switch fi88 {
                                                    case 42:
                                                        v5 -= (^(v5))
                                                    default:
                                                        _ = len(v4)
                                                    }
                                                    v92 += ((v92 * (((-2.25 / v92) * (-2.25 - -2.25)) - float64(1e10))) - (0.0 - (((v92 / float64(((1e10 + 1e10) * 1.5))) / 1.0) / v92)))
                                                    if true {
                                                        fmt.Println(v91, v92)
                                                        v99 := GenS10{F0: -978}
                                                        v100 := v4[286]
                                                        _ = v100
                                                        _ = v99
                                                    } else {
                                                        v92 -= v92
                                                        if true {
                                                            var v101 float64 = 3.14159
                                                            if ((fi88 * 32768) == (fi84 * fi84)) {
                                                                if len(v3) > 0 {
                                                                    v3[0] = "hello"
                                                                }
                                                                fmt.Println(v101, v89)
                                                                var v102 int64 = (-(v89))
                                                                var v103 int16 = 1000
                                                                switch (("" + "") + "") {
                                                                case "foo bar":
                                                                    v1 = append(v1, (fi88 | (fi88 - (^(fi88)))))
                                                                default:
                                                                    v4[42] = 65535
                                                                }
                                                                _ = v102
                                                                _ = v103
                                                            }
                                                            _ = v101
                                                        } else {
                                                            if false {
                                                                v104 := GenS11{F0: int64(v5), F1: v89, F2: uint(v5)}
                                                                switch fi84 {
                                                                case 107:
                                                                    if false {
                                                                        v89 = 223
                                                                    }
                                                                case 2147483647:
                                                                    if false {
                                                                        var v105 int = fi84
                                                                        if false {
                                                                            if (1 <= 1) {
                                                                                v5 = 0
                                                                            }
                                                                        } else {
                                                                            v4[2] = 2147483648
                                                                            v106 := GenS12{F0: 65535}
                                                                            _ = v106
                                                                        }
                                                                        _ = v105
                                                                    } else {
                                                                        if false {
                                                                            if true {
                                                                                var v107 uint32 = 0
                                                                                if (!(631 >= v89)) {
                                                                                    var v108 int8 = v91
                                                                                    v2 = (v2 >> ((460) & 63))
                                                                                    var v109 uint = 255
                                                                                    _ = v108
                                                                                    _ = v109
                                                                                }
                                                                                if (0 != 1) {
                                                                                    if true {
                                                                                        v110 := GenS7{F0: v89}
                                                                                        fmt.Println(v5, v6, v91)
                                                                                        var v111 int = fi88
                                                                                        _ = v110
                                                                                        _ = v111
                                                                                    } else {
                                                                                        v89 -= int64(v91)
                                                                                        var v112 int = (fi88 ^ fi88)
                                                                                        var v113 uint8 = 1
                                                                                        _ = v112
                                                                                        _ = v113
                                                                                    }
                                                                                    if ("foo bar" > "") {
                                                                                        fmt.Println(v5, v89)
                                                                                    } else {
                                                                                        v107 = (v107 * v107)
                                                                                        v6 += (v6 / ((4294967295) | 1))
                                                                                        var v114 int64 = (^(v89))
                                                                                        var v115 float32 = (0.0 / 0.1)
                                                                                        _ = v114
                                                                                        _ = v115
                                                                                    }
                                                                                    v116 := v4[(v107 * (^(v107)))]
                                                                                    _ = v116
                                                                                }
                                                                                _ = v107
                                                                            }
                                                                            if (!(!false)) {
                                                                                switch ("\n\t" + ("a" + "")) {
                                                                                case "a":
                                                                                    fmt.Println(v91, v5)
                                                                                }
                                                                                v92 = 0.0
                                                                                if (944 < -987) {
                                                                                    if false {
                                                                                        fmt.Println(v92, v89)
                                                                                        if ((-(v91)) < (v91 & 1)) {
                                                                                            if len(v3) >= 2 {
                                                                                                v3 = v3[1:len(v3)]
                                                                                            }
                                                                                            v89 = (v89 & v89)
                                                                                            var v117 uint32 = 0
                                                                                            _ = v117
                                                                                        }
                                                                                    } else {
                                                                                        var v118 string = (("foo bar" + "hello") + "a")
                                                                                        var v119 uint64 = 1
                                                                                        _ = v118
                                                                                        _ = v119
                                                                                    }
                                                                                    var v120 int16 = -32768
                                                                                    _ = v120
                                                                                } else {
                                                                                    v5 -= uint16(v6)
                                                                                }
                                                                            }
                                                                        }
                                                                        v121 := GenS3{F0: 343}
                                                                        _ = v121
                                                                    }
                                                                case -2147483648:
                                                                    if true {
                                                                        v92 += (v92 - (((0.1 * (v92 * (0.0 / -2.25))) + -2.25) / v92))
                                                                    } else {
                                                                        fmt.Println(v2)
                                                                        var v122 float64 = (v92 * (v92 - v92))
                                                                        _ = v122
                                                                    }
                                                                }
                                                                _ = v104
                                                            } else {
                                                                switch (fi88 / (((fi88 / ((fi88) | 1))) | 1)) {
                                                                case -2147483648:
                                                                    _ = len(v1)
                                                                case -905:
                                                                    v123 := GenS3{F0: (^(v5))}
                                                                    _ = v123
                                                                case 610:
                                                                    var v124 int32 = 1
                                                                    _ = v124
                                                                default:
                                                                    v4[uint32(v2)] = 213
                                                                }
                                                                v5 -= (^(v5))
                                                            }
                                                        }
                                                        var v125 uint16 = 217
                                                        _ = v125
                                                    }
                                                } else {
                                                    switch fi88 {
                                                    case 143:
                                                        if true {
                                                            v6 -= (^(v6))
                                                            v92 = v92
                                                            var v126 uint32 = 2147483647
                                                            var v127 uint = 42
                                                            delete(v4, uint32(v6))
                                                            _ = v126
                                                            _ = v127
                                                        } else {
                                                            v89 = 941
                                                            if true {
                                                                var v128 bool = ((fi84 * (fi88 % ((int(v2)) | 1))) <= (fi84 & int(v91)))
                                                                v129 := GenS13{F0: (36 >= v5), F1: 751, F2: (fi84 - fi84)}
                                                                v5 = v5
                                                                _ = v128
                                                                _ = v129
                                                            } else {
                                                                if (377 == 42) {
                                                                    var v130 uint16 = (v5 + v5)
                                                                    _ = len(v3)
                                                                    var v131 bool = true
                                                                    _ = v130
                                                                    _ = v131
                                                                } else {
                                                                    var v132 int64 = 1
                                                                    if (!((true && (v91 < (v91 ^ (v91 + (v91 << ((4294967295) & 63)))))) || (true && true))) {
                                                                        if true {
                                                                            switch "" {
                                                                            case "":
                                                                                switch "foo bar" {
                                                                                case "hello":
                                                                                    fmt.Println(v2, v89, fi84)
                                                                                }
                                                                            case "hello":
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = -1
                                                                                }
                                                                                var v133 bool = true
                                                                                _ = v133
                                                                            case "a":
                                                                                v134 := GenS2{F0: 1000, F1: ("a" + "\n\t")}
                                                                                _ = v134
                                                                            }
                                                                            var v135 float32 = 0.1
                                                                            v136 := v4[2147483648]
                                                                            var v137 int64 = v89
                                                                            _ = v135
                                                                            _ = v136
                                                                            _ = v137
                                                                        }
                                                                        if (!false) {
                                                                            var v138 int16 = -904
                                                                            var v139 float32 = 0.1
                                                                            fmt.Println(fi84)
                                                                            _ = v138
                                                                            _ = v139
                                                                        }
                                                                        if true {
                                                                            switch fi84 {
                                                                            case -564:
                                                                                v91 = v91
                                                                            }
                                                                            var v140 int16 = 42
                                                                            if true {
                                                                                var v141 uint = uint(v6)
                                                                                _ = v141
                                                                            }
                                                                            _ = v140
                                                                        }
                                                                        v132 = (v132 + v89)
                                                                        if (!false) {
                                                                            _ = len(v4)
                                                                        } else {
                                                                            if true {
                                                                                v4[858] = 2147483648
                                                                            } else {
                                                                                switch genFn1() {
                                                                                case -964:
                                                                                    _ = len(v4)
                                                                                }
                                                                            }
                                                                            v91 += (-(v91))
                                                                        }
                                                                    }
                                                                    _ = v132
                                                                }
                                                                v142 := GenS11{F0: (-(v89)), F1: v89, F2: 1}
                                                                if true {
                                                                    v91 = 110
                                                                } else {
                                                                    var v143 int64 = v89
                                                                    v143 = v89
                                                                    if (v92 <= (float64(((0.1 - ((-2.25 / (float32(v92) * (((3.14159 / (3.14159 * 1.5)) / ((0.1 - (((-2.25 * 1e10) + ((3.14159 * 0.0) / 0.1)) - ((-2.25 - float32(v92)) - 1e10))) / float32((v92 * 1e10)))) * 0.0))) * 0.1)) * -2.25)) + 0.1)) {
                                                                        fmt.Println(v6, v92)
                                                                        if (-2.25 >= (((0.0 / 1.5) / 0.1) / -2.25)) {
                                                                            v144 := GenS7{F0: (^(v89))}
                                                                            _ = v144
                                                                        }
                                                                        v145 := GenS11{F0: v143, F1: (v89 - 0), F2: 255}
                                                                        v146 := GenS6{F0: 317, F1: (!(1e10 == float32(0.1)))}
                                                                        switch (^(fi88)) {
                                                                        case 32767:
                                                                            fmt.Println(fi84, v6, v6)
                                                                        case 4294967295:
                                                                            fmt.Println(fi88, fi88)
                                                                        case -400:
                                                                            if (147 != 675) {
                                                                                var v147 uint8 = v2
                                                                                _ = v147
                                                                            }
                                                                        default:
                                                                            _ = len(v4)
                                                                        }
                                                                        _ = v145
                                                                        _ = v146
                                                                    } else {
                                                                        var v148 float32 = ((3.14159 + 1e10) / 1.0)
                                                                        v1 = append(v1, genFn1())
                                                                        v3 = append(v3, "hello")
                                                                        var v149 uint32 = 0
                                                                        _ = v148
                                                                        _ = v149
                                                                    }
                                                                    _ = v143
                                                                }
                                                                v150 := GenS13{F0: ((false && false) || (0 <= 0)), F1: 2147483648, F2: int(v5)}
                                                                _ = v142
                                                                _ = v150
                                                            }
                                                            v4[2] = 0
                                                        }
                                                    default:
                                                        switch (^(fi84)) {
                                                        case -973:
                                                            v92 += (0.1 * v92)
                                                            var v151 uint32 = 4294967295
                                                            _ = v151
                                                        case -9223372036854775808:
                                                            var v152 int = (fi84 >> ((836) & 63))
                                                            fmt.Println(fi84, v91, v89)
                                                            _ = v152
                                                        default:
                                                            v89 -= v89
                                                        }
                                                    }
                                                    v89 = v89
                                                    switch (^(fi88)) {
                                                    case -2147483648:
                                                        switch (fi88 ^ 32768) {
                                                        case 793:
                                                            v153 := GenS14{F0: (fi84 + (fi88 * fi84)), F1: fi84, F2: -1}
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                            _ = v153
                                                        case 42:
                                                            var v154 int16 = 1
                                                            _ = v154
                                                        case -170:
                                                            switch "" {
                                                            case "a":
                                                                v2 += v2
                                                            case "foo bar":
                                                                var v155 int32 = 0
                                                                _ = v155
                                                            case "":
                                                                switch "" {
                                                                case "foo bar":
                                                                    v156 := GenS15{F0: "\n\t", F1: ((("hello" + "foo bar") + "\n\t") + ""), F2: (v5 ^ v5)}
                                                                    _ = v156
                                                                case "\n\t":
                                                                    if true {
                                                                        _ = len(v1)
                                                                        fmt.Println(fi88, v89)
                                                                    } else {
                                                                        var v157 uint = 538
                                                                        v158 := GenS16{F0: (^(v5)), F1: v157, F2: v2}
                                                                        _ = v157
                                                                        _ = v158
                                                                    }
                                                                case "":
                                                                    _ = len(v4)
                                                                    v159 := GenS6{F0: v5, F1: true}
                                                                    _ = v159
                                                                default:
                                                                    v5 += 1
                                                                }
                                                            default:
                                                                fmt.Println(v91, v6, v91)
                                                            }
                                                        }
                                                    case -1:
                                                        v6 = 42
                                                    case -244:
                                                        v4[uint32(v2)] = 42
                                                        switch "foo bar" {
                                                        case "":
                                                            fmt.Println(fi84, v6)
                                                        case "\n\t":
                                                            if (18446744073709551615 >= 4294967295) {
                                                                var v160 int = (^(fi88))
                                                                var v161 int16 = -651
                                                                var v162 uint16 = 592
                                                                v89 = v89
                                                                _ = v160
                                                                _ = v161
                                                                _ = v162
                                                            } else {
                                                                var v163 uint64 = (v6 & (v6 & (v6 * v6)))
                                                                v164 := GenS17{F0: (v2 ^ 81)}
                                                                var v165 int8 = (v91 | -46)
                                                                var v166 int8 = (v91 / ((v91) | 1))
                                                                v1 = append(v1, fi84)
                                                                v167 := GenS18{F0: v89, F1: -126}
                                                                if false {
                                                                    v89 += -392
                                                                    v168 := GenS2{F0: 642, F1: ("hello" + "")}
                                                                    _ = v168
                                                                }
                                                                _ = v163
                                                                _ = v164
                                                                _ = v165
                                                                _ = v166
                                                                _ = v167
                                                            }
                                                        case "a":
                                                            if (!(!true)) {
                                                                var v169 int = fi84
                                                                _ = v169
                                                            }
                                                        default:
                                                            v6 -= (v6 & uint64(v5))
                                                        }
                                                        if false {
                                                            v170 := GenS19{F0: 42}
                                                            _ = v170
                                                        } else {
                                                            var v171 bool = (-2.25 > (1e10 / 1.5))
                                                            var v172 int32 = -2147483648
                                                            _ = v171
                                                            _ = v172
                                                        }
                                                    default:
                                                        var v173 float64 = v92
                                                        _ = v173
                                                    }
                                                }
                                                v92 = v92
                                            case "":
                                                v174 := GenS6{F0: (v5 + (v5 & v5)), F1: (!(!false))}
                                                _ = v174
                                            }
                                        }
                                    }
                                } else {
                                    v175 := GenS15{F0: ("foo bar" + ("" + "a")), F1: "foo bar", F2: (v5 << ((148) & 63))}
                                    if true {
                                        var v176 string = "\n\t"
                                        _ = v176
                                    }
                                    if true {
                                        if ((^(fi88)) > (fi84 / ((fi88) | 1))) {
                                            var v177 int64 = v89
                                            var v178 float64 = ((v92 - v92) * 0.0)
                                            switch "" {
                                            case "hello":
                                                v5 -= v5
                                            default:
                                                switch fi84 {
                                                case -1:
                                                    if false {
                                                        switch (("\n\t" + "hello") + (("foo bar" + ("hello" + (("hello" + ("\n\t" + (("hello" + "") + "a"))) + ((("foo bar" + (("a" + "hello") + ("" + "foo bar"))) + "\n\t") + "hello")))) + "hello")) {
                                                        case "hello":
                                                            v89 = (v177 >> ((42) & 63))
                                                        case "":
                                                            delete(v4, 2)
                                                        default:
                                                            v179 := GenS15{F0: "\n\t", F1: "a", F2: (v5 ^ (v5 ^ (v5 | v5)))}
                                                            _ = v179
                                                        }
                                                        v180 := GenS17{F0: 2}
                                                        switch (fi84 * -2147483648) {
                                                        case -623:
                                                            var v181 float64 = v178
                                                            _ = v181
                                                        case 42:
                                                            v6 = (v6 | v6)
                                                        }
                                                        v89 = (v89 * v89)
                                                        _ = v180
                                                    }
                                                case 489:
                                                    _ = len(v3)
                                                case -539:
                                                    v182 := GenS13{F0: ((("\n\t" + "") > "a") || (!true)), F1: uint32(v6), F2: fi88}
                                                    _ = v182
                                                default:
                                                    fmt.Println(v89, v6)
                                                }
                                            }
                                            _ = v177
                                            _ = v178
                                        }
                                        fmt.Println(fi84, v2)
                                        if (float64(0.0) < float64(0.1)) {
                                            if false {
                                                if true {
                                                    fmt.Println(v92, v91)
                                                    if len(v1) > 0 {
                                                        v1[0] = (fi88 << ((42) & 63))
                                                    }
                                                }
                                                v183 := GenS20{F0: v6, F1: 410, F2: -906}
                                                switch genFn1() {
                                                case -9223372036854775808:
                                                    v92 = v92
                                                default:
                                                    var v184 uint64 = 42
                                                    _ = v184
                                                }
                                                _ = v183
                                            } else {
                                                if (4294967295 != 455) {
                                                    var v185 uint64 = (v6 * (v6 & (v6 ^ (v6 - (v6 - (v6 + (v6 & (v6 | uint64(v2)))))))))
                                                    _ = v185
                                                }
                                            }
                                            v186 := GenS21{F0: 1000}
                                            _ = v186
                                        } else {
                                            switch fi88 {
                                            case 42:
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            case 32768:
                                                if false {
                                                    _ = len(v4)
                                                    switch fi88 {
                                                    case -9223372036854775808:
                                                        if (fi84 < genFn1()) {
                                                            v187 := GenS18{F0: v89, F1: 789}
                                                            fmt.Println(v89, fi88)
                                                            _ = v187
                                                        }
                                                    default:
                                                        fmt.Println(v6, v6)
                                                    }
                                                    var v188 int64 = v89
                                                    _ = v188
                                                } else {
                                                    v2 += 255
                                                    delete(v4, 4294967295)
                                                }
                                            }
                                        }
                                        fmt.Println(v91, fi84, v92)
                                    }
                                    _ = v175
                                }
                                _ = v92
                            }
                            v189 := GenS22{F0: v5, F1: (v5 ^ v5), F2: v2}
                            _ = v189
                        } else {
                            var v190 uint8 = (^(v2))
                            switch "foo bar" {
                            case "hello":
                                v91 += -117
                            default:
                                if false {
                                    _ = len(v1)
                                    v191 := GenS15{F0: "\n\t", F1: "hello", F2: v5}
                                    _ = v191
                                }
                            }
                            var v192 int32 = 1
                            _ = v190
                            _ = v192
                        }
                        v91 = v91
                        if true {
                            v193 := v4[65535]
                            v194 := v4[v193]
                            _ = v193
                            _ = v194
                        } else {
                            if (v5 <= v5) {
                                v91 = (v91 / ((1) | 1))
                                v195 := GenS11{F0: (v89 ^ v89), F1: (v89 % ((v89) | 1)), F2: 18446744073709551615}
                                _ = v195
                            } else {
                                v196 := GenS23{F0: (!(1 == -492)), F1: "hello"}
                                var v197 float32 = (1.5 - -2.25)
                                var v198 uint8 = v2
                                _ = v196
                                _ = v197
                                _ = v198
                            }
                            if false {
                                switch (fi84 ^ 127) {
                                case 0:
                                    fmt.Println(fi88, v91, fi88)
                                case 128:
                                    v199 := GenS24{F0: -552, F1: 1, F2: (("hello" + ("a" + (("" + "") + ("" + (("" + (("" + "foo bar") + "a")) + "foo bar"))))) + "hello")}
                                    switch (fi84 + (fi84 ^ genFn1())) {
                                    case -906:
                                        switch "foo bar" {
                                        case "":
                                            if false {
                                                v200 := v4[267]
                                                switch 32767 {
                                                case 32767:
                                                    switch (fi84 / ((int(v2)) | 1)) {
                                                    case -1:
                                                        v4[(v200 << ((42) & 63))] = v200
                                                    }
                                                case -560:
                                                    if (!(!(!true))) {
                                                        fmt.Println(v6, v6)
                                                    }
                                                }
                                                _ = v200
                                            } else {
                                                switch (fi84 >> ((42) & 63)) {
                                                case 0:
                                                    var v201 uint = 0
                                                    _ = v201
                                                }
                                            }
                                        default:
                                            v4[723] = 2147483648
                                        }
                                    default:
                                        fmt.Println(fi88, fi88, v91)
                                    }
                                    var v202 uint8 = (v2 & 1)
                                    _ = v199
                                    _ = v202
                                case 2147483647:
                                    var v203 int32 = 634
                                    _ = v203
                                default:
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    fmt.Println(v2, v5)
                                }
                            } else {
                                if (true || (!(42 <= -1))) {
                                    v5 -= v5
                                } else {
                                    switch "foo bar" {
                                    case "hello":
                                        switch "" {
                                        case "a":
                                            var v204 int8 = (v91 - -76)
                                            _ = v204
                                        default:
                                            if ((false && true) && (1 >= (v2 * v2))) {
                                                if false {
                                                    var v205 int32 = -819
                                                    if (v91 < int8(v2)) {
                                                        v91 -= (-(v91))
                                                    } else {
                                                        var v206 int8 = v91
                                                        _ = v206
                                                    }
                                                    var v207 uint8 = v2
                                                    _ = v205
                                                    _ = v207
                                                }
                                            } else {
                                                v5 += uint16(v2)
                                                if (!(((("" + "") + "foo bar") + "\n\t") > ("a" + "hello"))) {
                                                    var v208 string = (((((("a" + "a") + "a") + "") + "foo bar") + (("\n\t" + "hello") + "\n\t")) + (("hello" + "a") + "hello"))
                                                    if ((v91 < v91) || true) {
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        v209 := v4[uint32(v6)]
                                                        if true {
                                                            v210 := GenS3{F0: v5}
                                                            _ = v210
                                                        } else {
                                                            v4[v209] = (v209 + 2147483647)
                                                        }
                                                        v211 := GenS22{F0: v5, F1: (v5 * v5), F2: 2}
                                                        v212 := GenS25{F0: v91, F1: v209, F2: (v2 & (v2 << ((20) & 63)))}
                                                        _ = v209
                                                        _ = v211
                                                        _ = v212
                                                    } else {
                                                        v91 += v91
                                                        fmt.Println(v6)
                                                        var v213 uint64 = v6
                                                        _ = v213
                                                    }
                                                    _ = v208
                                                } else {
                                                    v5 = (^(v5))
                                                    if ((882 <= genFn1()) || false) {
                                                        var v214 uint64 = (v6 / ((uint64(v5)) | 1))
                                                        _ = v214
                                                    } else {
                                                        switch (^(fi84)) {
                                                        case 42:
                                                            v2 -= (v2 | v2)
                                                        case 9:
                                                            var v215 uint64 = (v6 & v6)
                                                            _ = v215
                                                        }
                                                    }
                                                    var v216 uint32 = 2
                                                    switch 50 {
                                                    case 32768:
                                                        if (!((v89 % (((v89 * v89)) | 1)) > (v89 - int64(v6)))) {
                                                            var v217 int32 = 42
                                                            v6 -= (v6 * 482)
                                                            fmt.Println(v91, v5)
                                                            _ = v217
                                                        } else {
                                                            if (false || true) {
                                                                fmt.Println(v6)
                                                                fmt.Println(v89)
                                                            }
                                                        }
                                                    case 0:
                                                        if len(v1) > 0 {
                                                            v1[0] = fi84
                                                        }
                                                    default:
                                                        if false {
                                                            switch "a" {
                                                            case "\n\t":
                                                                _ = len(v3)
                                                            case "foo bar":
                                                                fmt.Println(v91)
                                                            case "a":
                                                                var v218 uint16 = (v5 ^ (v5 >> ((4294967295) & 63)))
                                                                _ = v218
                                                            default:
                                                                if true {
                                                                    switch (("\n\t" + ("\n\t" + ("\n\t" + "hello"))) + "hello") {
                                                                    case "a":
                                                                        fmt.Println(v216)
                                                                    case "hello":
                                                                        v89 += (v89 * v89)
                                                                    }
                                                                    v219 := GenS26{F0: ((("foo bar" + "foo bar") + "") + ("a" + ""))}
                                                                    _ = v219
                                                                }
                                                            }
                                                        } else {
                                                            var v220 string = "\n\t"
                                                            _ = v220
                                                        }
                                                        v221 := GenS20{F0: v6, F1: fi84, F2: -567}
                                                        _ = v221
                                                    }
                                                    _ = v216
                                                }
                                                if true {
                                                    v2 -= v2
                                                }
                                                var v222 int = fi88
                                                _ = v222
                                            }
                                        }
                                    case "\n\t":
                                        switch (fi88 / ((int(v6)) | 1)) {
                                        case 128:
                                            v4[54] = 2
                                        case 42:
                                            v223 := GenS19{F0: -364}
                                            _ = v223
                                        }
                                    default:
                                        v2 = (^(v2))
                                    }
                                    fmt.Println(v89, v2, fi84)
                                    var v224 int16 = 32767
                                    _ = v224
                                }
                                if (false && true) {
                                    _ = len(v4)
                                    delete(v4, 212)
                                    if false {
                                        var v225 uint64 = (v6 ^ (v6 % ((2) | 1)))
                                        _ = v225
                                    }
                                    var v226 int = fi88
                                    var v227 string = ("" + (("\n\t" + "") + "foo bar"))
                                    _ = v226
                                    _ = v227
                                } else {
                                    v91 = v91
                                    fmt.Println(v89, v5)
                                }
                                var v228 int64 = v89
                                _ = v228
                            }
                        }
                        if true {
                            var v229 int8 = v91
                            v230 := v4[2147483648]
                            _ = v229
                            _ = v230
                        } else {
                            fmt.Println(v5, v2, v2)
                            switch int(v5) {
                            case 474:
                                if (v91 <= -128) {
                                    v231 := GenS23{F0: (0.1 == 3.14159), F1: "foo bar"}
                                    switch 294 {
                                    case 0:
                                        if false {
                                            if (42 == (fi84 | genFn1())) {
                                                var v232 uint = 543
                                                _ = v232
                                            }
                                            fmt.Println(v89, fi88, fi88)
                                        } else {
                                            if true {
                                                _ = len(v4)
                                            } else {
                                                v6 += (v6 / ((v6) | 1))
                                                v1 = append(v1, -707)
                                                var v233 uint8 = 5
                                                _ = v233
                                            }
                                            var v234 float32 = -2.25
                                            delete(v4, 42)
                                            _ = v234
                                        }
                                    case 4294967295:
                                        v89 += 0
                                        v235 := GenS5{F0: v5, F1: fi88}
                                        if true {
                                            v236 := GenS12{F0: 42}
                                            _ = v236
                                        }
                                        _ = v235
                                    }
                                    _ = v231
                                } else {
                                    var v237 float64 = 0.1
                                    var v238 string = "a"
                                    switch (("" + v238) + v238) {
                                    case "":
                                        v239 := GenS17{F0: (^(v2))}
                                        _ = v239
                                    case "a":
                                        switch (-(fi84)) {
                                        case -970:
                                            fmt.Println(v91)
                                        case 579:
                                            v238 = "foo bar"
                                        case 127:
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        }
                                    default:
                                        switch v238 {
                                        case "hello":
                                            fmt.Println(v91, v6, fi88)
                                        case "\n\t":
                                            if (2 > (v5 | v5)) {
                                                v240 := GenS11{F0: (v89 * (v89 / (((v89 >> ((4294967295) & 63))) | 1))), F1: (^(v89)), F2: 2}
                                                _ = v240
                                            }
                                        case "a":
                                            fmt.Println(v238, v6)
                                        default:
                                            var v241 float32 = (-2.25 + -2.25)
                                            _ = v241
                                        }
                                    }
                                    var v242 uint16 = 42
                                    var v243 int8 = v91
                                    var v244 uint64 = (v6 % ((v6) | 1))
                                    if true {
                                        var v245 int16 = 1
                                        if false {
                                            if false {
                                                var v246 int64 = (^(v89))
                                                _ = v246
                                            } else {
                                                var v247 float64 = v237
                                                v4[624] = 1
                                                _ = v247
                                            }
                                            if true {
                                                if true {
                                                    v248 := GenS11{F0: v89, F1: int64(v6), F2: 255}
                                                    _ = v248
                                                } else {
                                                    if (((1.5 < (0.1 * v237)) && true) && true) {
                                                        if false {
                                                            var v249 int64 = (v89 ^ v89)
                                                            v250 := GenS2{F0: v245, F1: v238}
                                                            if ((v245 / ((-32768) | 1)) >= v245) {
                                                                v237 = v237
                                                            } else {
                                                                v4[42] = 4294967295
                                                            }
                                                            _ = v249
                                                            _ = v250
                                                        } else {
                                                            delete(v4, 42)
                                                            if false {
                                                                if (114 < (v5 - (v5 & (v5 | 2)))) {
                                                                    var v251 uint64 = (v244 * v244)
                                                                    switch "" {
                                                                    case "a":
                                                                        v252 := GenS12{F0: 4294967295}
                                                                        _ = v252
                                                                    case "foo bar":
                                                                        var v253 uint = 219
                                                                        _ = v253
                                                                    }
                                                                    fmt.Println(v245, v91, v6)
                                                                    v254 := GenS27{F0: (v243 & (^(v243))), F1: "a"}
                                                                    _ = v251
                                                                    _ = v254
                                                                }
                                                            }
                                                            var v255 uint16 = v242
                                                            switch (fi88 << ((759) & 63)) {
                                                            case -2147483648:
                                                                var v256 int64 = (v89 + -415)
                                                                _ = v256
                                                            }
                                                            fmt.Println(v91)
                                                            v257 := GenS23{F0: (2 <= (^(v5))), F1: v238}
                                                            v242 -= (v5 << ((4294967295) & 63))
                                                            _ = v255
                                                            _ = v257
                                                        }
                                                        var v258 uint = 971
                                                        _ = v258
                                                    } else {
                                                        var v259 int64 = -993
                                                        _ = len(v4)
                                                        v245 += (^(v245))
                                                        _ = v259
                                                    }
                                                }
                                                if false {
                                                    fmt.Println(v242)
                                                    if true {
                                                        v260 := GenS13{F0: true, F1: 550, F2: genFn1()}
                                                        var v261 int16 = v245
                                                        if ((v91 > (v243 - v243)) && (((1.5 + (float64(0.0) / v237)) > 0.0) || (v237 <= 1.5))) {
                                                            switch v238 {
                                                            case "foo bar":
                                                                var v262 uint8 = (v2 + (v2 >> ((65535) & 63)))
                                                                _ = v262
                                                            case "\n\t":
                                                                var v263 int32 = -295
                                                                fmt.Println(v244)
                                                                var v264 int32 = (-(v263))
                                                                _ = v263
                                                                _ = v264
                                                            default:
                                                                v237 -= ((v237 + (0.1 * (1.5 - v237))) / float64(1e10))
                                                            }
                                                        } else {
                                                            v245 = (v261 % ((329) | 1))
                                                            v265 := GenS18{F0: 9223372036854775807, F1: 42}
                                                            v4[uint32(v5)] = 438
                                                            _ = v265
                                                        }
                                                        switch genFn1() {
                                                        case 128:
                                                            v2 += 42
                                                        case 968:
                                                            fmt.Println(v245)
                                                        case 774:
                                                            if true {
                                                                v6 -= v6
                                                                if false {
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                }
                                                            }
                                                        default:
                                                            fmt.Println(v237)
                                                            v1 = append(v1, genFn1())
                                                            v242 -= v5
                                                        }
                                                        var v266 int8 = v243
                                                        _ = v260
                                                        _ = v261
                                                        _ = v266
                                                    } else {
                                                        switch (-(fi88)) {
                                                        case 42:
                                                            var v267 string = "foo bar"
                                                            _ = v267
                                                        case 199:
                                                            v268 := GenS28{F0: (v89 << ((4294967295) & 63))}
                                                            _ = v268
                                                        case -1:
                                                            var v269 float32 = (0.0 + 1e10)
                                                            _ = v269
                                                        }
                                                        v243 = v243
                                                        if true {
                                                            fmt.Println(v243)
                                                        }
                                                    }
                                                    var v270 uint = 2
                                                    switch fi84 {
                                                    case 128:
                                                        v243 = 127
                                                    case 397:
                                                        v91 -= v243
                                                    }
                                                    v271 := GenS12{F0: 966}
                                                    _ = v270
                                                    _ = v271
                                                } else {
                                                    if true {
                                                        var v272 float32 = 1.5
                                                        _ = v272
                                                    }
                                                }
                                                var v273 uint16 = v5
                                                fmt.Println(v245, v91)
                                                var v274 int8 = v91
                                                fmt.Println(v237, v242)
                                                _ = v273
                                                _ = v274
                                            } else {
                                                v89 -= v89
                                                v243 = 100
                                            }
                                        } else {
                                            v237 -= (0.1 * v237)
                                            switch -9223372036854775808 {
                                            case 42:
                                                if false {
                                                    switch (^(fi88)) {
                                                    case 178:
                                                        v275 := GenS29{F0: 2147483648}
                                                        _ = v275
                                                    }
                                                    var v276 int = fi84
                                                    _ = len(v3)
                                                    v237 -= (v237 - v237)
                                                    _ = v276
                                                } else {
                                                    switch ((v238 + v238) + v238) {
                                                    case "foo bar":
                                                        v91 = -128
                                                    }
                                                    v5 += 65535
                                                    var v277 uint = 18446744073709551615
                                                    _ = v277
                                                }
                                            case 806:
                                                fmt.Println(v243)
                                            default:
                                                var v278 string = v238
                                                v279 := GenS30{F0: v278}
                                                v280 := GenS31{F0: 740, F1: false, F2: true}
                                                _ = v278
                                                _ = v279
                                                _ = v280
                                            }
                                        }
                                        fmt.Println(v238, v91, v244)
                                        var v281 string = "a"
                                        if false {
                                            v243 += (^(v243))
                                            if false {
                                                switch v238 {
                                                case "":
                                                    var v282 int = fi84
                                                    _ = v282
                                                }
                                                switch v238 {
                                                case "\n\t":
                                                    fmt.Println(v5, v237)
                                                default:
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                }
                                            }
                                        } else {
                                            if ((0.1 / (3.14159 + (0.1 + 3.14159))) > float32((float64(0.1) - ((v237 + ((v237 / 1.0) + v237)) * 1e10)))) {
                                                if true {
                                                    switch fi84 {
                                                    case -2147483648:
                                                        var v283 uint = 526
                                                        _ = v283
                                                    default:
                                                        v284 := GenS5{F0: v5, F1: (fi84 & (fi84 << ((40) & 63)))}
                                                        _ = v284
                                                    }
                                                    var v285 float32 = 1e10
                                                    _ = v285
                                                }
                                                var v286 uint16 = 255
                                                _ = v286
                                            }
                                            v287 := GenS28{F0: 9223372036854775807}
                                            if false {
                                                fmt.Println(fi84, v91, v91)
                                                _ = len(v1)
                                            }
                                            switch v281 {
                                            case "\n\t":
                                                v245 -= (v245 * (v245 % ((v245) | 1)))
                                            case "a":
                                                var v288 int8 = v91
                                                _ = v288
                                            case "":
                                                v89 = (v89 >> ((uint(v242)) & 63))
                                            default:
                                                v2 = (v2 / ((v2) | 1))
                                            }
                                            _ = v287
                                        }
                                        _ = v245
                                        _ = v281
                                    }
                                    var v289 bool = true
                                    var v290 int16 = 1
                                    _ = v237
                                    _ = v238
                                    _ = v242
                                    _ = v243
                                    _ = v244
                                    _ = v289
                                    _ = v290
                                }
                            case 1:
                                v291 := GenS32{F0: 32767, F1: 42, F2: v89}
                                _ = v291
                            case 2147483647:
                                v1 = append(v1, fi88)
                            }
                            switch fi88 {
                            case 42:
                                v5 = (^(v5))
                            case -9223372036854775808:
                                if (3.14159 <= ((0.1 / (0.1 + -2.25)) * 1e10)) {
                                    switch "" {
                                    case "hello":
                                        delete(v4, 356)
                                    }
                                }
                            default:
                                var v292 int16 = 42
                                v293 := GenS33{F0: (v292 * 32767), F1: 42}
                                delete(v4, 721)
                                _ = v292
                                _ = v293
                            }
                            switch (fi88 / ((128) | 1)) {
                            case -9223372036854775808:
                                switch ("a" + "\n\t") {
                                case "\n\t":
                                    v294 := GenS15{F0: "a", F1: "hello", F2: 32767}
                                    _ = v294
                                case "":
                                    v295 := GenS34{F0: v91}
                                    _ = v295
                                default:
                                    var v296 int32 = 2147483647
                                    _ = v296
                                }
                            case -1:
                                if (0.0 == ((float64(1.5) + (1e10 * 0.0)) + 3.14159)) {
                                    v2 += (^(v2))
                                } else {
                                    var v297 float64 = 0.0
                                    _ = v297
                                }
                            default:
                                var v298 int8 = v91
                                _ = v298
                            }
                            v6 = (v6 ^ (^(v6)))
                            if len(v3) > 0 {
                                v3[0] = "foo bar"
                            }
                            if true {
                                v91 -= (v91 + (^(v91)))
                                v91 = v91
                                if (3.14159 > float32(3.14159)) {
                                    fmt.Println(v2)
                                    if false {
                                        v299 := GenS17{F0: v2}
                                        _ = len(v4)
                                        v6 = (v6 ^ v6)
                                        v300 := GenS23{F0: false, F1: "hello"}
                                        _ = v299
                                        _ = v300
                                    } else {
                                        if len(v3) > 0 {
                                            v3[0] = "hello"
                                        }
                                    }
                                    if ((v2 % (((v2 & (v2 + v2))) | 1)) < (v2 - v2)) {
                                        var v301 float64 = float64(1e10)
                                        var v302 uint = 124
                                        var v303 uint = uint(v2)
                                        _ = v301
                                        _ = v302
                                        _ = v303
                                    }
                                    var v304 string = ("a" + "\n\t")
                                    _ = v304
                                }
                            } else {
                                switch -261 {
                                case -1:
                                    v6 = v6
                                    v305 := GenS35{F0: (v91 - (-(v91))), F1: (^(v6)), F2: uint8(v5)}
                                    var v306 int8 = (v91 / (((v91 & (v91 >> ((0) & 63)))) | 1))
                                    v2 = (v2 + 188)
                                    var v307 int8 = (^(v91))
                                    _ = v305
                                    _ = v306
                                    _ = v307
                                case -2147483648:
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                case 2147483647:
                                    v91 -= (-(v91))
                                }
                                v308 := GenS36{F0: (^(v2)), F1: "\n\t"}
                                _ = v308
                            }
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        _ = v91
                    }
                }
                _ = fi88
                _ = v89
            }
        default:
            var v309 int16 = -836
            _ = v309
        }
        if (42 <= -1) {
            for fi310 := 0; fi310 < 1; fi310++ {
                switch -651 {
                case 32768:
                    var v311 int64 = 0
                    _ = v311
                case -32:
                    var v312 int8 = -37
                    _ = v312
                }
                v6 -= 216
                _ = fi310
            }
            switch 32768 {
            case 127:
                for fi313 := 0; fi313 < 11; fi313++ {
                    if true {
                        v314 := GenS27{F0: 42, F1: ("" + "")}
                        if true {
                            v315 := GenS37{F0: (v2 / ((v2) | 1))}
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            switch "a" {
                            case "a":
                                _ = len(v1)
                                if true {
                                    fmt.Println(v5)
                                    _ = len(v3)
                                    if (-837 == 1) {
                                        v6 = v6
                                    }
                                    v316 := GenS38{F0: 2, F1: 18446744073709551615}
                                    var v317 int = 42
                                    _ = v316
                                    _ = v317
                                }
                            case "\n\t":
                                if true {
                                    fmt.Println(v5)
                                }
                            default:
                                v5 = v5
                            }
                            _ = v315
                        } else {
                            var v318 uint8 = uint8(v5)
                            _ = v318
                        }
                        var v319 uint = 255
                        switch "\n\t" {
                        case "a":
                            v3 = append(v3, ("hello" + "a"))
                        case "foo bar":
                            v320 := GenS33{F0: 42, F1: -829}
                            _ = v320
                        default:
                            var v321 int8 = 0
                            _ = v321
                        }
                        _ = v314
                        _ = v319
                    } else {
                        v5 = (^(v5))
                    }
                    var v322 uint64 = v6
                    _ = fi313
                    _ = v322
                }
            case 32768:
                fmt.Println(v5)
            }
        } else {
            v6 += 0
            for fi323 := 0; fi323 < 20; fi323++ {
                var v324 uint64 = (v6 | (v6 - (v6 & 573)))
                var v325 uint8 = v2
                _ = fi323
                _ = v324
                _ = v325
            }
            for fi326 := 0; fi326 < 4; fi326++ {
                if true {
                    v4[0] = 844
                    v5 = v5
                } else {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    if ((^(v5)) < (v5 / ((uint16(v6)) | 1))) {
                        v327 := GenS11{F0: -1, F1: 0, F2: 0}
                        var v328 int32 = 1
                        _ = v327
                        _ = v328
                    } else {
                        switch 127 {
                        case 9223372036854775807:
                            v6 += (v6 >> ((42) & 63))
                        case 1:
                            var v329 bool = false
                            _ = v329
                        }
                    }
                    fmt.Println(fi84)
                }
                v6 = 2
                v330 := GenS39{F0: (^(v5)), F1: 1, F2: 687}
                v6 -= v6
                if len(v3) > 0 {
                    v3[0] = "foo bar"
                }
                var v331 float32 = 3.14159
                var v332 uint64 = 0
                _ = fi326
                _ = v330
                _ = v331
                _ = v332
            }
        }
        _ = fi84
    }
    v6 += (^(v6))
    v4[0] = 42
    switch 36 {
    case 32768:
        v2 = (v2 - v2)
    default:
        v333 := GenS40{F0: -686}
        _ = v333
    }
    delete(v4, 2147483647)
    v2 = 164
    v6 -= (^(v6))
    for fi334 := 0; fi334 < 10; fi334++ {
        if (true || (false && (!(false && true)))) {
            var v335 uint = 4294967295
            for fi336 := 0; fi336 < 4; fi336++ {
                var v337 uint32 = uint32(v5)
                _ = fi336
                _ = v337
            }
            var v338 string = (((((("" + "\n\t") + (("a" + "hello") + "")) + "\n\t") + "foo bar") + "") + "hello")
            _ = v335
            _ = v338
        } else {
            switch "\n\t" {
            case "\n\t":
                for fi339 := 0; fi339 < 2; fi339++ {
                    v2 = (^(v2))
                    _ = fi339
                }
            }
            var v340 uint32 = 398
            _ = v340
        }
        var v341 uint16 = v5
        _ = fi334
        _ = v341
    }
    switch -200 {
    case 164:
        var v342 float64 = 3.14159
        _ = v342
    case 42:
        switch ((("a" + "foo bar") + "\n\t") + "foo bar") {
        case "":
            var v343 uint8 = 255
            _ = v343
        }
    case -9223372036854775808:
        switch "a" {
        case "\n\t":
            var v344 uint = 4294967295
            _ = v344
        case "foo bar":
            v2 += 255
        }
    }
    if (true && (true && true)) {
        _ = len(v3)
        if true {
            v6 = (v6 & (v6 / ((233) | 1)))
            v5 += 32767
            v345 := v4[566]
            _ = v345
        }
    } else {
        switch "" {
        case "a":
            var v346 uint = 0
            _ = v346
        case "\n\t":
            var v347 float64 = (1.5 / 1e10)
            _ = v347
        }
        v348 := GenS41{F0: v2, F1: 4294967295, F2: genFn1()}
        var v349 uint8 = 128
        _ = v348
        _ = v349
    }
    v5 += (v5 | v5)
    v350 := GenS42{F0: -1}
    var v351 uint8 = 0
    v352 := v4[699]
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v351:", v351)
    fmt.Println("v352:", v352)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v350
    _ = v351
    _ = v352
    _ = v4
    _ = v5
    _ = v6
}
