package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 int8; F1 bool; F2 int8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint; F1 int64 }
type GenS6 struct { F0 bool; F1 int8 }
type GenS7 struct { F0 int16; F1 uint16; F2 uint16 }
type GenS8 struct { F0 int; F1 int64 }
type GenS9 struct { F0 int8; F1 uint }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 int32; F1 uint16 }
type GenS12 struct { F0 uint64; F1 int; F2 uint32 }
type GenS13 struct { F0 string }
type GenS14 struct { F0 uint8; F1 int64; F2 uint16 }
type GenS15 struct { F0 int8; F1 uint32; F2 int32 }
type GenS16 struct { F0 int32; F1 int16 }
type GenS17 struct { F0 uint; F1 uint32 }
type GenS18 struct { F0 int; F1 string }
type GenS19 struct { F0 int8 }
type GenS20 struct { F0 bool; F1 uint16; F2 int32 }
type GenS21 struct { F0 int64; F1 uint8 }
type GenS22 struct { F0 uint64; F1 uint8 }
type GenS23 struct { F0 uint16 }
type GenS24 struct { F0 uint64; F1 uint32; F2 uint8 }
type GenS25 struct { F0 int8; F1 bool; F2 uint8 }
type GenS26 struct { F0 int64; F1 uint16; F2 int32 }
type GenS27 struct { F0 uint64; F1 uint64; F2 uint8 }
type GenS28 struct { F0 uint; F1 uint8 }
type GenS29 struct { F0 uint64; F1 bool; F2 int32 }
type GenS30 struct { F0 int }
type GenS31 struct { F0 int16 }
type GenS32 struct { F0 string; F1 string }
type GenS33 struct { F0 uint64; F1 uint16 }
type GenS34 struct { F0 uint8; F1 uint; F2 int16 }
type GenS35 struct { F0 uint16; F1 string }
type GenS36 struct { F0 uint; F1 uint }
type GenS37 struct { F0 int32 }
type GenS38 struct { F0 bool; F1 string }
type GenS39 struct { F0 uint32; F1 int16 }
type GenS40 struct { F0 uint8 }
func genFn0() uint32 {
    v6 := GenS1{F0: 0, F1: 2}
    var v7 uint32 = 0
    for fi8 := 0; fi8 < 8; fi8++ {
        if true {
            var v9 float32 = (1e10 + 0.1)
            _ = v9
        }
        _ = fi8
    }
    v10 := []int{0, 128, -138}
    _ = v10
    _ = v6
    _ = v7
    return (v7 % ((0) | 1))
}
func genFn1(p0 int32, p1 bool) string {
    fmt.Println(p0)
    p0 = p0
    switch 1 {
    case 127:
        p1 = false
    case -686:
        v11 := []int{-116}
        _ = v11
    case 2147483647:
        for fi12 := 0; fi12 < 2; fi12++ {
            if (-117 < 1) {
                if (255 >= 1) {
                    var v13 bool = (1e10 == 3.14159)
                    _ = v13
                }
                p0 = int32(fi12)
            } else {
                for fi14 := 0; fi14 < 2; fi14++ {
                    v15 := GenS2{F0: 127, F1: (0 > 664), F2: 127}
                    switch ((("" + "hello") + "hello") + "a") {
                    case "hello":
                        if (255 >= 32767) {
                            fmt.Println(p1, fi12)
                            switch "" {
                            case "\n\t":
                                fmt.Println(p0, fi12)
                            case "foo bar":
                                p0 = p0
                                fmt.Println(fi12)
                                var v16 uint8 = 127
                                _ = v16
                            case "":
                                var v17 uint32 = 0
                                _ = v17
                            default:
                                var v18 int64 = -9223372036854775808
                                _ = v18
                            }
                            v19 := GenS1{F0: (fi12 | fi12), F1: 32767}
                            _ = v19
                        }
                    }
                    switch 2147483647 {
                    case -349:
                        if (fi12 >= 0) {
                            v20 := GenS1{F0: (fi12 * fi14), F1: 42}
                            _ = v20
                        }
                        v21 := []int{2, (fi12 & 32768), fi12}
                        p0 = p0
                        _ = v21
                    case 32:
                        var v22 int8 = 0
                        _ = v22
                    case 0:
                        p0 = p0
                    default:
                        fmt.Println(p1, fi12)
                    }
                    _ = fi14
                    _ = v15
                }
                if false {
                    p1 = true
                    for fi23 := 0; fi23 < 1; fi23++ {
                        var v24 float32 = 0.1
                        _ = fi23
                        _ = v24
                    }
                    v25 := GenS1{F0: int(p0), F1: 891}
                    _ = v25
                } else {
                    if p1 {
                        v26 := []int{fi12}
                        p1 = (p1 && p1)
                        v27 := GenS1{F0: (fi12 << ((65535) & 63)), F1: 2}
                        _ = v26
                        _ = v27
                    }
                }
            }
            _ = fi12
        }
    }
    p0 += p0
    p1 = (!((!p1) && p1))
    _ = p0
    _ = p1
    return "a"
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]int)
    var v3 int = -2147483648
    v4 := []string{"hello", "foo bar", (("foo bar" + "\n\t") + ("" + "hello"))}
    v5 := make(map[string]bool)
    var v28 uint = 65535
    var v29 uint = v28
    for fi30 := 0; fi30 < 11; fi30++ {
        v29 = (v29 + 1)
        for fi31 := 0; fi31 < 18; fi31++ {
            if false {
                v32 := GenS3{F0: 219}
                v28 = (v28 ^ v28)
                _ = v32
            } else {
                v4 = append(v4, "a")
                fmt.Println(v3, v3, v3)
            }
            v33 := GenS2{F0: 88, F1: ((-(v3)) >= 32768), F2: -98}
            v34 := GenS1{F0: (fi31 & v3), F1: 2}
            var v35 bool = false
            _ = fi31
            _ = v33
            _ = v34
            _ = v35
        }
        if (0 > 669) {
            if (3.14159 <= 0.1) {
                v29 -= (v29 + v29)
            } else {
                var v36 int = (v3 - (fi30 | fi30))
                fmt.Println(v29, v36)
                _ = v36
            }
            for fi37 := 0; fi37 < 1; fi37++ {
                v38 := v1["foo bar"]
                var v39 uint16 = 957
                _ = fi37
                _ = v38
                _ = v39
            }
            v40 := GenS1{F0: (fi30 / ((v3) | 1)), F1: 65535}
            _ = v40
        } else {
            var v41 uint = 42
            v3 = (v3 & (v3 + (^(v3))))
            _ = v41
        }
        _ = fi30
    }
    v42 := GenS4{F0: (v29 & 94)}
    if (v3 < v3) {
        if ((1 > 114) && (189 != 127)) {
            v3 = (v3 >> ((v28) & 63))
            v43 := GenS1{F0: v3, F1: 255}
            for fi44 := 0; fi44 < 5; fi44++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v45 uint32 = 177
                switch "" {
                case "hello":
                    _ = len(v1)
                default:
                    fmt.Println(v45, fi44)
                }
                _ = fi44
                _ = v45
            }
            _ = v43
        } else {
            if true {
                if (!true) {
                    switch (("hello" + "hello") + "") {
                    case "hello":
                        for fi46 := 0; fi46 < 3; fi46++ {
                            v47 := v2["hello"]
                            _ = fi46
                            _ = v47
                        }
                    case "a":
                        v29 -= (v28 & (^(v29)))
                    case "":
                        var v48 float64 = 0.1
                        _ = v48
                    }
                } else {
                    var v49 uint = (v29 % ((v28) | 1))
                    for fi50 := 0; fi50 < 12; fi50++ {
                        for fi51 := 0; fi51 < 18; fi51++ {
                            fmt.Println(fi51, v28, fi51)
                            var v52 int64 = 42
                            switch (fi50 ^ (fi50 % (((v3 % (((-(fi50))) | 1))) | 1))) {
                            case 4294967295:
                                v49 = 92
                            case -828:
                                v28 = v28
                            case 127:
                                var v53 uint = 0
                                _ = v53
                            default:
                                v54 := GenS4{F0: v29}
                                _ = v54
                            }
                            _ = fi51
                            _ = v52
                        }
                        _ = fi50
                    }
                    _ = v49
                }
                v3 += (-(v3))
            } else {
                v55 := GenS4{F0: (^(v29))}
                for fi56 := 0; fi56 < 9; fi56++ {
                    v57 := GenS1{F0: (-(v3)), F1: 32767}
                    var v58 float64 = 0.1
                    for fi59 := 0; fi59 < 1; fi59++ {
                        switch 42 {
                        case 2147483647:
                            var v60 bool = ((v58 == 1.5) && (-1 >= -1))
                            _ = v60
                        case -168:
                            if (-128 <= 43) {
                                var v61 uint8 = 0
                                fmt.Println(v58)
                                _ = v61
                            } else {
                                var v62 int8 = -3
                                var v63 int64 = -1
                                v3 -= 618
                                fmt.Println(v29)
                                var v64 uint32 = 1
                                _ = v62
                                _ = v63
                                _ = v64
                            }
                        }
                        if true {
                            var v65 float32 = 0.0
                            _ = v65
                        }
                        _ = fi59
                    }
                    _ = fi56
                    _ = v57
                    _ = v58
                }
                v3 -= (v3 + v3)
                _ = v55
            }
            v29 = (v28 | v28)
            for fi66 := 0; fi66 < 17; fi66++ {
                var v67 uint16 = 32767
                _ = fi66
                _ = v67
            }
        }
        var v68 uint8 = 255
        _ = v68
    }
    if (!(255 != 2)) {
        switch (-(v3)) {
        case -714:
            switch v3 {
            case 9223372036854775807:
                v69 := GenS2{F0: 127, F1: true, F2: 127}
                _ = v69
            case 0:
                v3 = (v3 & -115)
            case 127:
                var v70 bool = true
                _ = v70
            }
        default:
            v71 := GenS2{F0: 77, F1: true, F2: -128}
            _ = v71
        }
    } else {
        var v72 bool = (!(!((-608 < 442) || (255 >= 255))))
        switch "hello" {
        case "hello":
            v3 = int(v28)
        default:
            for fi73 := 0; fi73 < 16; fi73++ {
                v74 := GenS1{F0: 880, F1: 751}
                _ = fi73
                _ = v74
            }
        }
        var v75 int = (v3 << ((255) & 63))
        _ = v72
        _ = v75
    }
    for fi76 := 0; fi76 < 12; fi76++ {
        if true {
            var v77 float32 = (0.1 + 0.0)
            if false {
                v78 := v2[("" + (((("foo bar" + ("hello" + "")) + ("" + ("hello" + "a"))) + "\n\t") + ("\n\t" + "hello")))]
                fmt.Println(v78, v29)
                _ = v78
            } else {
                for fi79 := 0; fi79 < 11; fi79++ {
                    var v80 bool = true
                    fmt.Println(fi76)
                    _ = fi79
                    _ = v80
                }
                v81 := v2["a"]
                if true {
                    if false {
                        for fi82 := 0; fi82 < 15; fi82++ {
                            if (false || true) {
                                v29 += v29
                                var v83 int16 = 1000
                                _ = v83
                            } else {
                                var v84 uint64 = 1
                                var v85 float32 = float32((float64(v77) - 1.5))
                                var v86 uint = (v29 + (v28 % (((v28 + v29)) | 1)))
                                if true {
                                    switch (v3 % ((fi76) | 1)) {
                                    case -2147483648:
                                        v87 := GenS1{F0: fi82, F1: 39}
                                        _ = v87
                                    case 32768:
                                        v88 := GenS1{F0: -9223372036854775808, F1: 430}
                                        _ = v88
                                    default:
                                        fmt.Println(v85, v85, fi82)
                                    }
                                } else {
                                    if true {
                                        _ = len(v5)
                                        v77 = (v77 * (v85 / v85))
                                        _ = len(v4)
                                    } else {
                                        v86 += 42
                                        v86 = v86
                                        v89 := GenS5{F0: 343, F1: 759}
                                        _ = v89
                                    }
                                }
                                _ = v84
                                _ = v85
                                _ = v86
                            }
                            var v90 uint16 = 308
                            var v91 int = (fi76 - -56)
                            _ = fi82
                            _ = v90
                            _ = v91
                        }
                        if false {
                            if false {
                                v3 -= v81
                                if false {
                                    var v92 uint32 = genFn0()
                                    if false {
                                        var v93 int8 = -121
                                        _ = v93
                                    } else {
                                        v94 := GenS5{F0: v29, F1: 1}
                                        var v95 string = "a"
                                        var v96 string = v95
                                        _ = v94
                                        _ = v95
                                        _ = v96
                                    }
                                    if true {
                                        v77 = float32(1e10)
                                        switch ("" + "hello") {
                                        case "a":
                                            v97 := v1[""]
                                            _ = v97
                                        case "\n\t":
                                            fmt.Println(v92, v28)
                                        default:
                                            var v98 uint = 255
                                            _ = v98
                                        }
                                    } else {
                                        if false {
                                            fmt.Println(v92, v29, v81)
                                            switch "foo bar" {
                                            case "":
                                                var v99 float32 = v77
                                                _ = v99
                                            default:
                                                fmt.Println(v3)
                                            }
                                        } else {
                                            v77 = v77
                                            v77 -= v77
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        }
                                        fmt.Println(fi76)
                                        for fi100 := 0; fi100 < 10; fi100++ {
                                            if ((v81 >> ((4294967295) & 63)) >= (fi76 ^ fi76)) {
                                                var v101 int32 = 1
                                                v102 := v1["hello"]
                                                if (true && v102) {
                                                    if (true && true) {
                                                        var v103 int32 = v101
                                                        _ = v103
                                                    }
                                                } else {
                                                    fmt.Println(v101, v101, fi100)
                                                    if v102 {
                                                        if true {
                                                            if false {
                                                                v29 -= 1
                                                                v4 = append(v4, "a")
                                                                v77 = (v77 + v77)
                                                                v104 := GenS2{F0: -1, F1: (-600 != 2147483647), F2: 1}
                                                                _ = v104
                                                            } else {
                                                                var v105 uint = v29
                                                                if (!(65535 > 42)) {
                                                                    fmt.Println(v81)
                                                                    if (0.0 != (0.1 + (1.5 + (((3.14159 * (3.14159 - float64(float32(1e10)))) + 0.1) - 0.1)))) {
                                                                        var v106 uint16 = 397
                                                                        switch "a" {
                                                                        case "\n\t":
                                                                            v29 -= 42
                                                                        case "":
                                                                            v101 = v101
                                                                        case "hello":
                                                                            fmt.Println(v101, v102, v3)
                                                                        default:
                                                                            if (v102 && v102) {
                                                                                if true {
                                                                                    var v107 uint8 = 255
                                                                                    v77 = v77
                                                                                    _ = v107
                                                                                } else {
                                                                                    v77 += v77
                                                                                    v1["hello"] = false
                                                                                }
                                                                                v108 := GenS6{F0: v102, F1: 42}
                                                                                if v102 {
                                                                                    v28 += (v105 % ((v28) | 1))
                                                                                    switch "\n\t" {
                                                                                    case "foo bar":
                                                                                        var v109 int64 = 0
                                                                                        _ = v109
                                                                                    case "a":
                                                                                        v5["foo bar"] = false
                                                                                    case "\n\t":
                                                                                        v28 = v29
                                                                                    default:
                                                                                        v110 := GenS2{F0: -54, F1: false, F2: 1}
                                                                                        _ = v110
                                                                                    }
                                                                                    fmt.Println(v77, v106)
                                                                                } else {
                                                                                    if v102 {
                                                                                        var v111 float32 = 3.14159
                                                                                        fmt.Println(v102)
                                                                                        _ = v111
                                                                                    } else {
                                                                                        v112 := GenS7{F0: 706, F1: v106, F2: (v106 - (v106 << (((^(v28))) & 63)))}
                                                                                        _ = v112
                                                                                    }
                                                                                    v106 -= (v106 >> ((65535) & 63))
                                                                                    v81 -= 1
                                                                                    v5[("\n\t" + "")] = (float64(1e10) >= (0.1 * (1.5 * 1e10)))
                                                                                    v77 -= v77
                                                                                }
                                                                                _ = v108
                                                                            }
                                                                        }
                                                                        _ = v106
                                                                    }
                                                                } else {
                                                                    switch "a" {
                                                                    case "hello":
                                                                        var v113 uint64 = 1
                                                                        _ = v113
                                                                    }
                                                                    v2["\n\t"] = (fi76 & fi76)
                                                                }
                                                                v3 = fi76
                                                                fmt.Println(v105, v105)
                                                                if false {
                                                                    if (genFn0() > (^(v92))) {
                                                                        delete(v5, "")
                                                                    } else {
                                                                        if (-1 < 1) {
                                                                            var v114 int32 = 2147483647
                                                                            _ = v114
                                                                        } else {
                                                                            var v115 bool = (1 < 0)
                                                                            v2["a"] = fi100
                                                                            var v116 float64 = -2.25
                                                                            _ = v115
                                                                            _ = v116
                                                                        }
                                                                    }
                                                                    var v117 int64 = 1
                                                                    _ = v117
                                                                }
                                                                v118 := GenS4{F0: 0}
                                                                _ = v105
                                                                _ = v118
                                                            }
                                                            if false {
                                                                v28 -= v29
                                                                fmt.Println(fi100, v29, v92)
                                                                var v119 uint = (^(v28))
                                                                _ = v119
                                                            } else {
                                                                if v102 {
                                                                    var v120 string = (("foo bar" + "foo bar") + ("hello" + "\n\t"))
                                                                    v121 := GenS8{F0: v81, F1: -9223372036854775808}
                                                                    v77 = v77
                                                                    _ = v120
                                                                    _ = v121
                                                                } else {
                                                                    fmt.Println(v81)
                                                                }
                                                                fmt.Println(fi100)
                                                                if len(v4) > 0 {
                                                                    v4[0] = ""
                                                                }
                                                                fmt.Println(v28, v77, v77)
                                                                var v122 int16 = -411
                                                                var v123 int32 = -658
                                                                var v124 string = ""
                                                                _ = v122
                                                                _ = v123
                                                                _ = v124
                                                            }
                                                            v92 -= (v92 + (v92 * (v92 ^ genFn0())))
                                                        }
                                                        v92 = 4294967295
                                                        if v102 {
                                                            var v125 int8 = -111
                                                            _ = v125
                                                        }
                                                    } else {
                                                        if len(v4) > 0 {
                                                            v4[0] = ("hello" + "a")
                                                        }
                                                        var v126 uint32 = v92
                                                        _ = v126
                                                    }
                                                    var v127 float32 = ((0.0 + ((1e10 - ((1e10 - v77) / v77)) / v77)) * ((v77 / float32(float64(v77))) + (0.0 / 1.5)))
                                                    _ = v127
                                                }
                                                _ = v101
                                                _ = v102
                                            } else {
                                                if (!true) {
                                                    fmt.Println(v28, v28)
                                                }
                                            }
                                            if false {
                                                var v128 float64 = ((0.0 + 1.5) / 0.1)
                                                v2[""] = fi76
                                                if false {
                                                    v129 := GenS9{F0: -110, F1: v28}
                                                    fmt.Println(fi76, v77)
                                                    if (2 <= uint64(v92)) {
                                                        _ = len(v5)
                                                        v130 := v5["hello"]
                                                        v81 += 901
                                                        _ = v130
                                                    }
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        var v131 uint16 = 255
                                                        _ = v131
                                                    case "a":
                                                        var v132 bool = false
                                                        _ = v132
                                                    case "hello":
                                                        v4 = append(v4, ("a" + ("foo bar" + "foo bar")))
                                                    default:
                                                        if (660 != 328) {
                                                            var v133 float64 = -2.25
                                                            var v134 int8 = 127
                                                            _ = v133
                                                            _ = v134
                                                        } else {
                                                            v81 = (v3 * 2147483647)
                                                            if (!(!(!true))) {
                                                                var v135 bool = ((fi76 << ((uint(v92)) & 63)) <= (-(v3)))
                                                                var v136 string = "hello"
                                                                _ = v135
                                                                _ = v136
                                                            } else {
                                                                var v137 uint64 = 0
                                                                var v138 float64 = 3.14159
                                                                _ = v137
                                                                _ = v138
                                                            }
                                                        }
                                                        fmt.Println(v81)
                                                        if (false || true) {
                                                            v3 -= fi76
                                                            var v139 uint = (^(v28))
                                                            _ = v139
                                                        }
                                                    }
                                                    if false {
                                                        switch v3 {
                                                        case 32767:
                                                            if (255 > 127) {
                                                                switch ("" + "hello") {
                                                                case "\n\t":
                                                                    v77 = v77
                                                                }
                                                                v140 := v2[""]
                                                                var v141 int16 = 32767
                                                                _ = v140
                                                                _ = v141
                                                            } else {
                                                                var v142 uint32 = (v92 | (v92 * (v92 | (v92 / ((v92) | 1)))))
                                                                _ = v142
                                                            }
                                                        case -830:
                                                            if false {
                                                                var v143 float32 = (v77 * 0.0)
                                                                _ = v143
                                                            } else {
                                                                v81 -= (fi76 & 1)
                                                                var v144 uint32 = (v92 - (^(v92)))
                                                                if (((!false) || true) || false) {
                                                                    v145 := GenS10{F0: -55}
                                                                    _ = v145
                                                                }
                                                                var v146 bool = false
                                                                fmt.Println(v29, v29)
                                                                v147 := GenS11{F0: 2147483647, F1: 234}
                                                                _ = v144
                                                                _ = v146
                                                                _ = v147
                                                            }
                                                        case 1:
                                                            switch fi76 {
                                                            case -9223372036854775808:
                                                                var v148 bool = (false || (!((!((1 >= 1) && false)) || (!true))))
                                                                _ = v148
                                                            default:
                                                                fmt.Println(fi100, v3)
                                                                if (-1 <= 0) {
                                                                    v77 -= (0.0 * ((1e10 - 0.1) * ((-2.25 + v77) * (-2.25 * 1e10))))
                                                                    fmt.Println(v77, fi100)
                                                                    var v149 float32 = float32(float64((0.1 / float32(3.14159))))
                                                                    _ = v149
                                                                }
                                                                v29 += (v29 % (((v28 >> ((v29) & 63))) | 1))
                                                            }
                                                        default:
                                                            switch "hello" {
                                                            case "hello":
                                                                if len(v4) > 0 {
                                                                    v4[0] = ""
                                                                }
                                                            }
                                                        }
                                                    }
                                                    v150 := GenS9{F0: 1, F1: (v28 % (((v29 << (((v29 * (v28 * (v29 >> (((v28 * (v28 % ((255) | 1)))) & 63))))) & 63))) | 1))}
                                                    _ = v129
                                                    _ = v150
                                                }
                                                _ = v128
                                            } else {
                                                var v151 uint64 = 18446744073709551615
                                                fmt.Println(fi76)
                                                if (uint32(v28) > (v92 - (^(v92)))) {
                                                    v92 = (v92 & v92)
                                                    v92 = (v92 % ((1) | 1))
                                                } else {
                                                    if (false && false) {
                                                        v152 := GenS5{F0: v29, F1: 0}
                                                        var v153 int32 = 393
                                                        _ = v152
                                                        _ = v153
                                                    }
                                                    var v154 uint64 = 73
                                                    _ = v154
                                                }
                                                if (!(-671 > -984)) {
                                                    v155 := GenS12{F0: v151, F1: (-(fi76)), F2: (^(v92))}
                                                    _ = v155
                                                }
                                                _ = v151
                                            }
                                            var v156 uint8 = 0
                                            _ = fi100
                                            _ = v156
                                        }
                                    }
                                    _ = v92
                                } else {
                                    _ = len(v5)
                                }
                                var v157 int32 = -2147483648
                                _ = v157
                            } else {
                                for fi158 := 0; fi158 < 14; fi158++ {
                                    v29 += (v28 / (((v28 >> ((v29) & 63))) | 1))
                                    if (794 != 1) {
                                        if (false || false) {
                                            var v159 int16 = -96
                                            _ = len(v4)
                                            _ = v159
                                        }
                                        var v160 uint = (^(v29))
                                        _ = v160
                                    } else {
                                        fmt.Println(v28, v28)
                                        if true {
                                            switch "a" {
                                            case "\n\t":
                                                v3 += (fi76 ^ fi158)
                                            case "hello":
                                                if (!false) {
                                                    fmt.Println(fi76)
                                                    v161 := GenS7{F0: 32767, F1: 255, F2: 228}
                                                    if len(v4) > 0 {
                                                        v4[0] = "\n\t"
                                                    }
                                                    _ = v161
                                                } else {
                                                    var v162 int = 128
                                                    v163 := GenS13{F0: ("a" + (("a" + (((("" + ("foo bar" + ("" + ("foo bar" + ("foo bar" + ("a" + "hello")))))) + "hello") + "a") + ("foo bar" + "\n\t"))) + "foo bar"))}
                                                    fmt.Println(fi158, v29)
                                                    _ = v162
                                                    _ = v163
                                                }
                                                delete(v1, "hello")
                                                switch (("foo bar" + "a") + ((("\n\t" + "foo bar") + "foo bar") + ("foo bar" + ("foo bar" + "hello")))) {
                                                case "hello":
                                                    v164 := GenS14{F0: 2, F1: -9223372036854775808, F2: 32767}
                                                    v3 = 602
                                                    switch v3 {
                                                    case 42:
                                                        fmt.Println(fi76, v81)
                                                    }
                                                    _ = v164
                                                case "":
                                                    var v165 int8 = 0
                                                    _ = v165
                                                case "a":
                                                    v166 := GenS15{F0: -128, F1: genFn0(), F2: 733}
                                                    _ = v166
                                                }
                                            default:
                                                if false {
                                                    var v167 int32 = 1
                                                    v29 = (v29 - (v29 | v28))
                                                    delete(v5, (((("a" + "") + (("a" + ((("" + "hello") + ("hello" + "")) + "\n\t")) + "a")) + "a") + ("\n\t" + "foo bar")))
                                                    _ = v167
                                                } else {
                                                    if false {
                                                        if false {
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                            if true {
                                                                if (float32((0.1 * 0.0)) < 1.5) {
                                                                    _ = len(v2)
                                                                    v168 := GenS16{F0: 296, F1: 0}
                                                                    if true {
                                                                        if (226 < 32767) {
                                                                            v169 := GenS4{F0: v28}
                                                                            if (!false) {
                                                                                if len(v4) > 0 {
                                                                                    v4[0] = ("\n\t" + "foo bar")
                                                                                }
                                                                                if true {
                                                                                    _ = len(v4)
                                                                                }
                                                                                var v170 int64 = -810
                                                                                _ = v170
                                                                            } else {
                                                                                if false {
                                                                                    var v171 uint8 = 2
                                                                                    switch v3 {
                                                                                    case 127:
                                                                                        fmt.Println(v28, v28, v171)
                                                                                    case -490:
                                                                                        v172 := GenS5{F0: v29, F1: 0}
                                                                                        _ = v172
                                                                                    default:
                                                                                        switch fi158 {
                                                                                        case 4294967295:
                                                                                            fmt.Println(v3, v29, v171)
                                                                                        case 359:
                                                                                            var v173 float32 = v77
                                                                                            _ = v173
                                                                                        case 32767:
                                                                                            var v174 int = fi76
                                                                                            _ = v174
                                                                                        default:
                                                                                            fmt.Println(v81, v81, v171)
                                                                                        }
                                                                                    }
                                                                                    v175 := GenS17{F0: (v29 ^ (v29 * 255)), F1: genFn0()}
                                                                                    _ = v171
                                                                                    _ = v175
                                                                                } else {
                                                                                    v81 += v3
                                                                                    v176 := v1[""]
                                                                                    v177 := GenS5{F0: v28, F1: 9223372036854775807}
                                                                                    _ = v176
                                                                                    _ = v177
                                                                                }
                                                                                switch ("\n\t" + (("foo bar" + "") + "a")) {
                                                                                case "foo bar":
                                                                                    fmt.Println(v3, v28, v77)
                                                                                    if (true || false) {
                                                                                        v178 := GenS4{F0: 313}
                                                                                        _ = len(v5)
                                                                                        if false {
                                                                                            v179 := GenS18{F0: (v81 >> ((647) & 63)), F1: ""}
                                                                                            v180 := GenS14{F0: 255, F1: 1, F2: 32767}
                                                                                            var v181 uint8 = 42
                                                                                            _ = v179
                                                                                            _ = v180
                                                                                            _ = v181
                                                                                        }
                                                                                        _ = v178
                                                                                    } else {
                                                                                        v77 -= 1.5
                                                                                    }
                                                                                    switch "hello" {
                                                                                    case "foo bar":
                                                                                        var v182 uint = (v29 ^ v29)
                                                                                        _ = v182
                                                                                    case "\n\t":
                                                                                        v5["a"] = false
                                                                                        fmt.Println(v28, v77)
                                                                                    case "":
                                                                                        delete(v5, ("\n\t" + ""))
                                                                                    }
                                                                                }
                                                                            }
                                                                            var v183 uint8 = 20
                                                                            v183 = 127
                                                                            _ = v169
                                                                            _ = v183
                                                                        } else {
                                                                            var v184 bool = (fi158 > 69)
                                                                            v185 := GenS19{F0: 100}
                                                                            _ = v184
                                                                            _ = v185
                                                                        }
                                                                    }
                                                                    var v186 uint8 = 151
                                                                    _ = v168
                                                                    _ = v186
                                                                }
                                                                var v187 uint16 = 119
                                                                var v188 int = 128
                                                                v29 = v28
                                                                var v189 uint32 = 2
                                                                var v190 bool = true
                                                                switch "\n\t" {
                                                                case "foo bar":
                                                                    v191 := GenS20{F0: ((998 > 855) || v190), F1: v187, F2: 0}
                                                                    _ = v191
                                                                case "":
                                                                    v187 = (v187 * 399)
                                                                    switch v3 {
                                                                    case 964:
                                                                        if (!true) {
                                                                            v188 = fi76
                                                                            var v192 uint8 = 156
                                                                            var v193 uint32 = genFn0()
                                                                            if true {
                                                                                var v194 uint16 = (^(v187))
                                                                                v195 := v1["hello"]
                                                                                v196 := GenS21{F0: -459, F1: v192}
                                                                                _ = v194
                                                                                _ = v195
                                                                                _ = v196
                                                                            }
                                                                            var v197 uint8 = v192
                                                                            _ = v192
                                                                            _ = v193
                                                                            _ = v197
                                                                        }
                                                                    default:
                                                                        if (-561 >= -394) {
                                                                            v198 := GenS18{F0: 0, F1: "foo bar"}
                                                                            _ = v198
                                                                        }
                                                                    }
                                                                }
                                                                switch (fi76 / ((901) | 1)) {
                                                                case 582:
                                                                    v29 = 4294967295
                                                                }
                                                                if (!true) {
                                                                    fmt.Println(v190, v28)
                                                                    switch v81 {
                                                                    case -648:
                                                                        if ((v28 % ((944) | 1)) <= 397) {
                                                                            fmt.Println(v187, v81)
                                                                            if v190 {
                                                                                var v199 uint8 = 127
                                                                                v200 := GenS22{F0: 18446744073709551615, F1: v199}
                                                                                _ = v199
                                                                                _ = v200
                                                                            }
                                                                            v77 = 0.1
                                                                            v5[""] = (v188 != (v3 + 758))
                                                                            if v190 {
                                                                                switch "" {
                                                                                case "\n\t":
                                                                                    v2[(("hello" + ("" + ("\n\t" + "\n\t"))) + ("a" + "a"))] = fi76
                                                                                }
                                                                            } else {
                                                                                if (127 == 100) {
                                                                                    fmt.Println(v29, v190)
                                                                                } else {
                                                                                    var v201 uint64 = 18446744073709551615
                                                                                    if false {
                                                                                        if (false || v190) {
                                                                                            switch "foo bar" {
                                                                                            case "hello":
                                                                                                switch -259 {
                                                                                                case 32767:
                                                                                                    if (!(!((v28 / ((v29) | 1)) > (v29 | (v29 / (((v28 % (((v29 + v29)) | 1))) | 1)))))) {
                                                                                                        v187 -= v187
                                                                                                    } else {
                                                                                                        fmt.Println(fi76, v189, v190)
                                                                                                        if (((^(v187)) >= (v187 / ((v187) | 1))) || v190) {
                                                                                                            if (true || v190) {
                                                                                                                v29 = 104
                                                                                                                v77 = (v77 - float32(1.5))
                                                                                                                var v202 bool = v190
                                                                                                                _ = v202
                                                                                                            } else {
                                                                                                                v203 := GenS3{F0: 169}
                                                                                                                if true {
                                                                                                                    var v204 int32 = 0
                                                                                                                    _ = v204
                                                                                                                } else {
                                                                                                                    v205 := GenS14{F0: 1, F1: 1, F2: (v187 + (^(v187)))}
                                                                                                                    _ = v205
                                                                                                                }
                                                                                                                _ = v203
                                                                                                            }
                                                                                                            if v190 {
                                                                                                                v189 -= v189
                                                                                                            } else {
                                                                                                                if v190 {
                                                                                                                    var v206 float64 = float64(0.0)
                                                                                                                    v188 -= (v3 - 4294967295)
                                                                                                                    fmt.Println(v189, v187)
                                                                                                                    v206 -= v206
                                                                                                                    _ = v206
                                                                                                                } else {
                                                                                                                    if v190 {
                                                                                                                        if true {
                                                                                                                            var v207 bool = v190
                                                                                                                            var v208 uint8 = 163
                                                                                                                            v28 = 739
                                                                                                                            _ = v207
                                                                                                                            _ = v208
                                                                                                                        } else {
                                                                                                                            v81 = 429
                                                                                                                            fmt.Println(v189, v201, v201)
                                                                                                                            v188 = 42
                                                                                                                            v209 := GenS1{F0: (v188 >> ((4294967295) & 63)), F1: 1}
                                                                                                                            _ = v209
                                                                                                                        }
                                                                                                                    }
                                                                                                                    var v210 string = "a"
                                                                                                                    fmt.Println(fi158, fi76)
                                                                                                                    _ = v210
                                                                                                                }
                                                                                                                v201 += 2
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                default:
                                                                                                    v3 += (v3 + v3)
                                                                                                }
                                                                                            case "a":
                                                                                                fmt.Println(v188, v187)
                                                                                            }
                                                                                            fmt.Println(v3)
                                                                                            var v211 bool = true
                                                                                            _ = v211
                                                                                        } else {
                                                                                            v212 := GenS16{F0: int32(v201), F1: 1000}
                                                                                            fmt.Println(fi76)
                                                                                            v213 := GenS23{F0: (v187 / (((^(v187))) | 1))}
                                                                                            _ = v212
                                                                                            _ = v213
                                                                                        }
                                                                                    } else {
                                                                                        if v190 {
                                                                                            if len(v4) > 0 {
                                                                                                v4[0] = (("\n\t" + ("\n\t" + ("foo bar" + "hello"))) + "")
                                                                                            }
                                                                                            var v214 bool = true
                                                                                            var v215 int8 = -1
                                                                                            if len(v4) >= 2 {
                                                                                                v4 = v4[1:len(v4)]
                                                                                            }
                                                                                            _ = v214
                                                                                            _ = v215
                                                                                        } else {
                                                                                            v216 := GenS24{F0: v201, F1: genFn0(), F2: 193}
                                                                                            _ = v216
                                                                                        }
                                                                                        v217 := GenS8{F0: fi158, F1: -340}
                                                                                        var v218 float64 = 1.5
                                                                                        if len(v4) >= 2 {
                                                                                            v4 = v4[1:len(v4)]
                                                                                        }
                                                                                        var v219 string = ("" + "")
                                                                                        v220 := GenS25{F0: -1, F1: ((9223372036854775807 < (v188 % ((9223372036854775807) | 1))) && v190), F2: 100}
                                                                                        _ = v217
                                                                                        _ = v218
                                                                                        _ = v219
                                                                                        _ = v220
                                                                                    }
                                                                                    _ = v201
                                                                                }
                                                                            }
                                                                            var v221 bool = v190
                                                                            _ = v221
                                                                        }
                                                                    case 128:
                                                                        switch (fi76 ^ (^(fi158))) {
                                                                        case -132:
                                                                            if true {
                                                                                if v190 {
                                                                                    v188 -= fi158
                                                                                    _ = len(v4)
                                                                                } else {
                                                                                    if true {
                                                                                        var v222 bool = false
                                                                                        if v190 {
                                                                                            var v223 uint32 = (v189 | v189)
                                                                                            _ = v223
                                                                                        }
                                                                                        if v222 {
                                                                                            v188 += v188
                                                                                        } else {
                                                                                            if len(v4) >= 2 {
                                                                                                v4 = v4[1:len(v4)]
                                                                                            }
                                                                                        }
                                                                                        if ((("foo bar" + "") + "") < "a") {
                                                                                            var v224 uint8 = uint8(v29)
                                                                                            var v225 string = ""
                                                                                            v222 = v190
                                                                                            _ = v224
                                                                                            _ = v225
                                                                                        } else {
                                                                                            v28 += (v29 * (v29 >> ((65535) & 63)))
                                                                                        }
                                                                                        _ = v222
                                                                                    } else {
                                                                                        delete(v5, "a")
                                                                                        fmt.Println(v190)
                                                                                    }
                                                                                    var v226 uint64 = 42
                                                                                    v81 = (v188 + 277)
                                                                                    _ = v226
                                                                                }
                                                                                fmt.Println(v29, v29, fi76)
                                                                            }
                                                                        case 32767:
                                                                            fmt.Println(v187, fi158)
                                                                        default:
                                                                            _ = len(v4)
                                                                            if v190 {
                                                                                fmt.Println(v190, v77, v28)
                                                                                var v227 float32 = v77
                                                                                _ = v227
                                                                            }
                                                                            var v228 uint16 = (v187 + (v187 / (((v187 % ((963) | 1))) | 1)))
                                                                            var v229 uint32 = v189
                                                                            switch fi158 {
                                                                            case 32767:
                                                                                v230 := v5["a"]
                                                                                _ = v230
                                                                            case 138:
                                                                                var v231 float32 = -2.25
                                                                                v232 := GenS15{F0: -51, F1: 42, F2: 1}
                                                                                var v233 uint32 = v189
                                                                                _ = v231
                                                                                _ = v232
                                                                                _ = v233
                                                                            }
                                                                            switch (fi76 ^ v188) {
                                                                            case 2147483647:
                                                                                v188 = fi76
                                                                            case -1:
                                                                                var v234 uint64 = 2
                                                                                _ = v234
                                                                            }
                                                                            if (186 > 255) {
                                                                                fmt.Println(v229)
                                                                            } else {
                                                                                _ = len(v5)
                                                                            }
                                                                            _ = v228
                                                                            _ = v229
                                                                        }
                                                                    }
                                                                    var v235 float32 = v77
                                                                    if v190 {
                                                                        if false {
                                                                            var v236 string = ""
                                                                            _ = v236
                                                                        }
                                                                    }
                                                                    v237 := GenS23{F0: uint16(v189)}
                                                                    v28 = (^(v28))
                                                                    _ = v235
                                                                    _ = v237
                                                                } else {
                                                                    v187 += v187
                                                                    if len(v4) >= 2 {
                                                                        v4 = v4[1:len(v4)]
                                                                    }
                                                                }
                                                                var v238 int32 = 327
                                                                var v239 uint64 = 1
                                                                _ = v187
                                                                _ = v188
                                                                _ = v189
                                                                _ = v190
                                                                _ = v238
                                                                _ = v239
                                                            }
                                                            v81 += 4294967295
                                                        } else {
                                                            v240 := GenS9{F0: 102, F1: 365}
                                                            var v241 uint64 = 1
                                                            v242 := GenS26{F0: 0, F1: 32767, F2: int32(fi76)}
                                                            if (100 >= 14) {
                                                                var v243 uint64 = 18446744073709551615
                                                                switch (v81 + -545) {
                                                                case -212:
                                                                    v77 += v77
                                                                case -522:
                                                                    v28 = 2
                                                                case 128:
                                                                    if true {
                                                                        v244 := GenS27{F0: 4294967295, F1: (v243 - uint64(v29)), F2: 255}
                                                                        v245 := GenS17{F0: v28, F1: genFn0()}
                                                                        switch "foo bar" {
                                                                        case "hello":
                                                                            _ = len(v1)
                                                                        case "\n\t":
                                                                            fmt.Println(v28, fi158, fi158)
                                                                        default:
                                                                            switch "a" {
                                                                            case "hello":
                                                                                if true {
                                                                                    v243 = (^(v241))
                                                                                    fmt.Println(fi158)
                                                                                } else {
                                                                                    if (!false) {
                                                                                        var v246 uint32 = 37
                                                                                        if false {
                                                                                            var v247 uint16 = 0
                                                                                            v28 -= uint(v247)
                                                                                            v246 = genFn0()
                                                                                            _ = v247
                                                                                        } else {
                                                                                            v81 = (v3 + (fi158 | fi158))
                                                                                            delete(v5, "hello")
                                                                                            v81 -= -709
                                                                                            if false {
                                                                                                switch 905 {
                                                                                                case 127:
                                                                                                    _ = len(v1)
                                                                                                case 9223372036854775807:
                                                                                                    if true {
                                                                                                        fmt.Println(v81)
                                                                                                        switch "" {
                                                                                                        case "hello":
                                                                                                            v248 := GenS24{F0: v243, F1: (v246 << ((v28) & 63)), F2: 128}
                                                                                                            _ = v248
                                                                                                        }
                                                                                                    } else {
                                                                                                        if true {
                                                                                                            var v249 int32 = 0
                                                                                                            fmt.Println(v81, v243, v243)
                                                                                                            v250 := GenS28{F0: (v29 % (((^(v28))) | 1)), F1: uint8(v246)}
                                                                                                            _ = v249
                                                                                                            _ = v250
                                                                                                        }
                                                                                                        fmt.Println(v77, v241)
                                                                                                    }
                                                                                                    switch v3 {
                                                                                                    case -980:
                                                                                                        if ("foo bar" > "foo bar") {
                                                                                                            switch ("foo bar" + "\n\t") {
                                                                                                            case "a":
                                                                                                                switch (fi76 ^ -683) {
                                                                                                                case 9223372036854775807:
                                                                                                                    if (("\n\t" + "a") > "foo bar") {
                                                                                                                        v77 += v77
                                                                                                                        switch ("hello" + "") {
                                                                                                                        case "a":
                                                                                                                            v251 := GenS29{F0: v241, F1: (12 > 2), F2: -526}
                                                                                                                            _ = v251
                                                                                                                        case "hello":
                                                                                                                            var v252 uint64 = v243
                                                                                                                            _ = v252
                                                                                                                        case "":
                                                                                                                            if true {
                                                                                                                                var v253 int8 = 100
                                                                                                                                if true {
                                                                                                                                    var v254 uint64 = 2
                                                                                                                                    _ = v254
                                                                                                                                } else {
                                                                                                                                    var v255 float32 = 0.0
                                                                                                                                    _ = len(v5)
                                                                                                                                    _ = v255
                                                                                                                                }
                                                                                                                                v77 += v77
                                                                                                                                _ = v253
                                                                                                                            } else {
                                                                                                                                switch (fi76 / ((fi158) | 1)) {
                                                                                                                                case 32768:
                                                                                                                                    v3 -= v81
                                                                                                                                case 443:
                                                                                                                                    v28 = (v29 * v28)
                                                                                                                                case 4294967295:
                                                                                                                                    var v256 int8 = -1
                                                                                                                                    _ = v256
                                                                                                                                default:
                                                                                                                                    var v257 bool = true
                                                                                                                                    _ = v257
                                                                                                                                }
                                                                                                                                v258 := GenS10{F0: 1}
                                                                                                                                _ = v258
                                                                                                                            }
                                                                                                                        default:
                                                                                                                            v4 = append(v4, ("" + ("a" + ((((("" + "hello") + ("" + ("a" + (("a" + "") + "foo bar")))) + (((("\n\t" + "") + ("" + "hello")) + (("a" + ("hello" + "a")) + "hello")) + "\n\t")) + "foo bar") + "a"))))
                                                                                                                            v29 -= v29
                                                                                                                            _ = len(v2)
                                                                                                                        }
                                                                                                                    } else {
                                                                                                                        v246 = v246
                                                                                                                        if false {
                                                                                                                            v259 := GenS11{F0: -1, F1: 677}
                                                                                                                            _ = v259
                                                                                                                        }
                                                                                                                        v77 = v77
                                                                                                                    }
                                                                                                                }
                                                                                                            case "foo bar":
                                                                                                                var v260 bool = false
                                                                                                                _ = v260
                                                                                                            case "":
                                                                                                                if (v246 > (v246 / ((1) | 1))) {
                                                                                                                    fmt.Println(v241, v243)
                                                                                                                }
                                                                                                            default:
                                                                                                                switch (^(v81)) {
                                                                                                                case 127:
                                                                                                                    v261 := GenS30{F0: (v81 * fi76)}
                                                                                                                    _ = v261
                                                                                                                case -480:
                                                                                                                    var v262 int8 = -112
                                                                                                                    _ = v262
                                                                                                                default:
                                                                                                                    v81 += (v3 / (((fi158 - v3)) | 1))
                                                                                                                }
                                                                                                            }
                                                                                                            v81 = 127
                                                                                                            v29 = v28
                                                                                                            v263 := GenS24{F0: (v243 & v241), F1: (v246 * v246), F2: 127}
                                                                                                            var v264 uint64 = v243
                                                                                                            _ = v263
                                                                                                            _ = v264
                                                                                                        } else {
                                                                                                            v265 := GenS31{F0: -1}
                                                                                                            v77 = v77
                                                                                                            v29 -= (v29 % (((v28 * (v29 | v28))) | 1))
                                                                                                            _ = v265
                                                                                                        }
                                                                                                    case -667:
                                                                                                        if true {
                                                                                                            v29 = v28
                                                                                                        } else {
                                                                                                            var v266 uint8 = 0
                                                                                                            fmt.Println(v266)
                                                                                                            _ = v266
                                                                                                        }
                                                                                                    default:
                                                                                                        switch 596 {
                                                                                                        case 32768:
                                                                                                            var v267 uint16 = 42
                                                                                                            v77 = (v77 * 0.1)
                                                                                                            fmt.Println(v81, v267)
                                                                                                            v268 := GenS32{F0: "a", F1: "hello"}
                                                                                                            var v269 uint64 = v241
                                                                                                            _ = v267
                                                                                                            _ = v268
                                                                                                            _ = v269
                                                                                                        case 128:
                                                                                                            var v270 int32 = 0
                                                                                                            _ = v270
                                                                                                        }
                                                                                                    }
                                                                                                    var v271 uint64 = (v241 / (((^(v241))) | 1))
                                                                                                    _ = v271
                                                                                                default:
                                                                                                    if true {
                                                                                                        switch "" {
                                                                                                        case "foo bar":
                                                                                                            var v272 float32 = v77
                                                                                                            if false {
                                                                                                                if len(v4) > 0 {
                                                                                                                    v4[0] = ("\n\t" + "hello")
                                                                                                                }
                                                                                                                switch 2147483647 {
                                                                                                                case 4294967295:
                                                                                                                    _ = len(v5)
                                                                                                                case -1:
                                                                                                                    v246 = v246
                                                                                                                }
                                                                                                            } else {
                                                                                                                v273 := GenS21{F0: -9223372036854775808, F1: 153}
                                                                                                                var v274 uint8 = 127
                                                                                                                _ = v273
                                                                                                                _ = v274
                                                                                                            }
                                                                                                            v246 += 1
                                                                                                            _ = v272
                                                                                                        case "\n\t":
                                                                                                            fmt.Println(v241, v29, v29)
                                                                                                        default:
                                                                                                            if (1 >= 287) {
                                                                                                                var v275 uint = v28
                                                                                                                v276 := GenS20{F0: true, F1: 1, F2: int32(v243)}
                                                                                                                if (0 <= (v275 + v28)) {
                                                                                                                    if false {
                                                                                                                        fmt.Println(v246, v275, v29)
                                                                                                                    } else {
                                                                                                                        v275 = v28
                                                                                                                    }
                                                                                                                    var v277 uint = v28
                                                                                                                    _ = v277
                                                                                                                } else {
                                                                                                                    v77 = ((v77 - (v77 * v77)) - v77)
                                                                                                                    var v278 float32 = v77
                                                                                                                    _ = v278
                                                                                                                }
                                                                                                                _ = v275
                                                                                                                _ = v276
                                                                                                            }
                                                                                                        }
                                                                                                        v29 = v29
                                                                                                    } else {
                                                                                                        v3 = -960
                                                                                                    }
                                                                                                }
                                                                                                v246 = v246
                                                                                            } else {
                                                                                                var v279 string = "\n\t"
                                                                                                if (1 >= 57) {
                                                                                                    v28 -= uint(v246)
                                                                                                    v280 := GenS33{F0: v241, F1: 1}
                                                                                                    switch fi158 {
                                                                                                    case 42:
                                                                                                        v281 := GenS33{F0: v241, F1: 0}
                                                                                                        _ = v281
                                                                                                    case 701:
                                                                                                        var v282 int8 = -128
                                                                                                        _ = v282
                                                                                                    case 32768:
                                                                                                        v3 -= (v3 + 361)
                                                                                                    }
                                                                                                    _ = v280
                                                                                                }
                                                                                                var v283 uint16 = 2
                                                                                                v77 = (v77 * v77)
                                                                                                _ = v279
                                                                                                _ = v283
                                                                                            }
                                                                                            switch fi158 {
                                                                                            case -9223372036854775808:
                                                                                                v241 = 862
                                                                                            case 128:
                                                                                                if false {
                                                                                                    var v284 float64 = ((-2.25 / 1.5) - -2.25)
                                                                                                    _ = v284
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        if false {
                                                                                            var v285 uint = (^(v29))
                                                                                            var v286 uint16 = 42
                                                                                            _ = v285
                                                                                            _ = v286
                                                                                        } else {
                                                                                            var v287 int32 = -2147483648
                                                                                            var v288 int64 = 560
                                                                                            _ = v287
                                                                                            _ = v288
                                                                                        }
                                                                                        _ = v246
                                                                                    } else {
                                                                                        switch "\n\t" {
                                                                                        case "\n\t":
                                                                                            v289 := GenS34{F0: 128, F1: v29, F2: 0}
                                                                                            _ = v289
                                                                                        }
                                                                                    }
                                                                                }
                                                                            default:
                                                                                var v290 uint64 = (^(v243))
                                                                                _ = v290
                                                                            }
                                                                        }
                                                                        _ = v244
                                                                        _ = v245
                                                                    } else {
                                                                        v241 = v241
                                                                        v241 -= v241
                                                                    }
                                                                }
                                                                var v291 float32 = v77
                                                                var v292 int64 = 80
                                                                _ = v243
                                                                _ = v291
                                                                _ = v292
                                                            }
                                                            _ = v240
                                                            _ = v241
                                                            _ = v242
                                                        }
                                                        var v293 uint8 = 0
                                                        _ = v293
                                                    } else {
                                                        switch "a" {
                                                        case "\n\t":
                                                            if ("foo bar" != "foo bar") {
                                                                v77 = v77
                                                            } else {
                                                                v29 = 716
                                                                var v294 float64 = 3.14159
                                                                v294 -= (v294 + -2.25)
                                                                var v295 int64 = 42
                                                                v29 = v29
                                                                _ = v294
                                                                _ = v295
                                                            }
                                                            if (((1.5 * 3.14159) <= (float64((((v77 * float32(0.1)) + v77) * v77)) - 1e10)) || (!(42 != 2147483647))) {
                                                                switch 32768 {
                                                                case -29:
                                                                    v77 -= 0.0
                                                                default:
                                                                    v296 := GenS32{F0: ("hello" + ("\n\t" + ("foo bar" + "a"))), F1: ("hello" + "\n\t")}
                                                                    _ = v296
                                                                }
                                                                var v297 float64 = 0.1
                                                                _ = v297
                                                            } else {
                                                                if (1000 < 1000) {
                                                                    fmt.Println(v28)
                                                                } else {
                                                                    v5["\n\t"] = true
                                                                    var v298 uint64 = 0
                                                                    var v299 string = ((("" + "\n\t") + ("\n\t" + "")) + "foo bar")
                                                                    _ = v298
                                                                    _ = v299
                                                                }
                                                                v77 = v77
                                                                v77 = (v77 * 0.1)
                                                            }
                                                            v300 := GenS35{F0: 1, F1: "\n\t"}
                                                            _ = v300
                                                        default:
                                                            if true {
                                                                v301 := GenS28{F0: (v28 << ((65535) & 63)), F1: 11}
                                                                fmt.Println(v28, v77)
                                                                if len(v4) >= 2 {
                                                                    v4 = v4[1:len(v4)]
                                                                }
                                                                _ = v301
                                                            } else {
                                                                var v302 float64 = ((float64((v77 - (((v77 + v77) + v77) + v77))) * (1.5 * 3.14159)) - float64((float32(float64(1.5)) - -2.25)))
                                                                if true {
                                                                    v77 -= 1e10
                                                                    var v303 int = (v3 - v3)
                                                                    v304 := GenS19{F0: 16}
                                                                    _ = v303
                                                                    _ = v304
                                                                }
                                                                v77 = 1e10
                                                                v305 := v5["foo bar"]
                                                                _ = v302
                                                                _ = v305
                                                            }
                                                        }
                                                    }
                                                    switch (fi158 + v3) {
                                                    case 1:
                                                        if (2147483648 >= 4294967295) {
                                                            var v306 bool = (585 >= 743)
                                                            v306 = true
                                                            if true {
                                                                var v307 bool = v306
                                                                _ = v307
                                                            }
                                                            v77 -= float32(((1.5 + 3.14159) / 3.14159))
                                                            fmt.Println(v77)
                                                            _ = v306
                                                        } else {
                                                            v308 := GenS36{F0: v28, F1: 2}
                                                            v309 := GenS37{F0: 2147483647}
                                                            _ = v308
                                                            _ = v309
                                                        }
                                                        fmt.Println(v3, fi76, fi158)
                                                    case 127:
                                                        fmt.Println(v81, fi158)
                                                    case 128:
                                                        switch (fi76 % ((-9223372036854775808) | 1)) {
                                                        case -1:
                                                            var v310 float32 = (0.0 + -2.25)
                                                            _ = v310
                                                        case 1:
                                                            fmt.Println(v81, v29)
                                                        case 127:
                                                            if true {
                                                                var v311 uint = v28
                                                                _ = v311
                                                            } else {
                                                                v29 += v28
                                                                v312 := GenS10{F0: -28}
                                                                _ = v312
                                                            }
                                                            v77 -= (((v77 * (v77 / v77)) / ((0.0 - 0.0) + (3.14159 / v77))) - (v77 / (v77 / (v77 / (v77 * v77)))))
                                                            v313 := GenS16{F0: 42, F1: 0}
                                                            _ = v313
                                                        }
                                                    default:
                                                        var v314 int = fi76
                                                        _ = v314
                                                    }
                                                }
                                            }
                                        }
                                    }
                                    v315 := v5["hello"]
                                    _ = fi158
                                    _ = v315
                                }
                            }
                            var v316 string = "hello"
                            _ = v316
                        }
                    }
                    delete(v5, "")
                    if true {
                        if false {
                            if (-2147483648 >= -1) {
                                v317 := GenS3{F0: 0}
                                _ = v317
                            }
                            switch ("hello" + "hello") {
                            case "foo bar":
                                _ = len(v4)
                            case "a":
                                v318 := GenS26{F0: 0, F1: 760, F2: 0}
                                _ = v318
                            default:
                                fmt.Println(v3)
                            }
                        }
                        fmt.Println(v81, fi76, v81)
                    } else {
                        fmt.Println(v29, v29)
                        if false {
                            v319 := v1[("a" + "hello")]
                            _ = v319
                        } else {
                            switch ("foo bar" + "hello") {
                            case "\n\t":
                                switch "a" {
                                case "a":
                                    var v320 int8 = 42
                                    _ = v320
                                case "foo bar":
                                    var v321 int8 = 42
                                    _ = v321
                                default:
                                    fmt.Println(v29, v3)
                                }
                            case "foo bar":
                                var v322 float32 = 1.5
                                _ = v322
                            default:
                                fmt.Println(v3)
                            }
                        }
                        var v323 int = v81
                        _ = v323
                    }
                } else {
                    var v324 uint64 = 2
                    for fi325 := 0; fi325 < 5; fi325++ {
                        if (!((("foo bar" + ((("hello" + "") + (("hello" + "hello") + "foo bar")) + "")) != (((("hello" + "hello") + "hello") + "a") + "a")) && true)) {
                            v81 = int(v324)
                            var v326 float64 = 0.1
                            _ = v326
                        } else {
                            v327 := GenS10{F0: -1}
                            _ = v327
                        }
                        _ = fi325
                    }
                    _ = v324
                }
                _ = v81
            }
            _ = v77
        } else {
            for fi328 := 0; fi328 < 6; fi328++ {
                v29 -= 1
                fmt.Println(v3)
                _ = fi328
            }
            _ = len(v5)
            v3 = fi76
        }
        _ = fi76
    }
    fmt.Println(v28, v3)
    if true {
        var v329 uint16 = 884
        delete(v5, "hello")
        v28 += v28
        _ = v329
    }
    for fi330 := 0; fi330 < 15; fi330++ {
        if false {
            delete(v5, ((("\n\t" + "foo bar") + "a") + "a"))
        }
        if false {
            var v331 uint16 = 1
            v331 = (v331 | (v331 << ((v29) & 63)))
            var v332 int = (fi330 - (^(fi330)))
            _ = v331
            _ = v332
        } else {
            for fi333 := 0; fi333 < 13; fi333++ {
                var v334 bool = false
                v335 := GenS38{F0: (81 == 2147483647), F1: "\n\t"}
                var v336 int32 = 2147483647
                _ = fi333
                _ = v334
                _ = v335
                _ = v336
            }
            for fi337 := 0; fi337 < 17; fi337++ {
                if true {
                    fmt.Println(v28)
                    var v338 uint16 = 255
                    _ = v338
                }
                _ = fi337
            }
            var v339 uint32 = 65535
            _ = v339
        }
        for fi340 := 0; fi340 < 7; fi340++ {
            if false {
                var v341 int8 = 42
                fmt.Println(v341, v3)
                fmt.Println(fi340, fi330, fi340)
                _ = v341
            } else {
                fmt.Println(fi340, v29)
                delete(v1, (("\n\t" + "foo bar") + "hello"))
            }
            _ = fi340
        }
        _ = fi330
    }
    if true {
        v342 := GenS39{F0: 548, F1: 42}
        if false {
            v3 = 2147483647
            _ = len(v1)
            for fi343 := 0; fi343 < 7; fi343++ {
                for fi344 := 0; fi344 < 5; fi344++ {
                    _ = len(v4)
                    var v345 float64 = ((0.1 * (0.1 - 1e10)) - 0.0)
                    if len(v4) > 0 {
                        v4[0] = ("hello" + "")
                    }
                    _ = fi344
                    _ = v345
                }
                if ((3.14159 - ((-2.25 - 0.0) / (1.5 * (1.5 / -2.25)))) > 1.5) {
                    switch ("a" + "\n\t") {
                    case "foo bar":
                        if ("" != "") {
                            v346 := GenS28{F0: 42, F1: 215}
                            for fi347 := 0; fi347 < 6; fi347++ {
                                v29 = (v28 - 4294967295)
                                if ((((("\n\t" + "") + "hello") + "") + "\n\t") <= "hello") {
                                    delete(v1, (((("foo bar" + "") + "\n\t") + "a") + ""))
                                } else {
                                    if true {
                                        fmt.Println(v28)
                                        var v348 int = fi343
                                        _ = v348
                                    }
                                }
                                _ = fi347
                            }
                            v28 = v28
                            _ = v346
                        } else {
                            var v349 int32 = 554
                            var v350 string = ""
                            _ = v349
                            _ = v350
                        }
                    case "a":
                        var v351 int16 = 1
                        _ = v351
                    }
                    v28 -= 317
                } else {
                    fmt.Println(v29, fi343, fi343)
                    var v352 float64 = 0.0
                    if false {
                        v29 += v29
                        var v353 uint64 = 359
                        _ = v353
                    } else {
                        var v354 int8 = -1
                        fmt.Println(v29)
                        var v355 float64 = (v352 - (v352 / v352))
                        _ = v354
                        _ = v355
                    }
                    _ = v352
                }
                var v356 uint32 = 2147483647
                _ = fi343
                _ = v356
            }
        }
        var v357 uint16 = 42
        _ = v342
        _ = v357
    } else {
        v358 := GenS30{F0: 0}
        fmt.Println(v28)
        _ = v358
    }
    v29 = v28
    switch (("foo bar" + "foo bar") + ("foo bar" + ("\n\t" + ("\n\t" + "a")))) {
    case "\n\t":
        v359 := GenS27{F0: 713, F1: 1, F2: 243}
        _ = v359
    case "hello":
        var v360 int = v3
        _ = v360
    default:
        for fi361 := 0; fi361 < 3; fi361++ {
            for fi362 := 0; fi362 < 17; fi362++ {
                v28 += 739
                _ = fi362
            }
            _ = fi361
        }
    }
    v3 = 9223372036854775807
    for fi363 := 0; fi363 < 13; fi363++ {
        var v364 uint = (v28 >> (((v28 & (v29 >> (((v29 << ((v28) & 63))) & 63)))) & 63))
        _ = fi363
        _ = v364
    }
    v365 := GenS17{F0: 18446744073709551615, F1: 4294967295}
    v366 := GenS40{F0: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v365
    _ = v366
    _ = v4
    _ = v42
    _ = v5
}
