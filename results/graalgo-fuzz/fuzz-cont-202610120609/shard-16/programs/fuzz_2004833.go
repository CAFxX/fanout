package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 uint32; F2 bool }
type GenS3 struct { F0 uint16; F1 int }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int32; F1 bool }
type GenS7 struct { F0 uint8; F1 int64; F2 uint64 }
type GenS8 struct { F0 uint16; F1 uint32; F2 uint32 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 int16; F1 int64; F2 int64 }
type GenS11 struct { F0 int64; F1 int32 }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 uint16; F1 int32; F2 int }
type GenS15 struct { F0 uint16; F1 bool }
type GenS16 struct { F0 uint16; F1 int }
type GenS17 struct { F0 int8; F1 uint16 }
type GenS18 struct { F0 uint16; F1 bool }
type GenS19 struct { F0 uint8; F1 uint8; F2 int }

func main() {
    var v1 int16 = -32768
    var v2 int64 = -1
    var v3 int = 1
    v4 := make(map[int]bool)
    v5 := make(map[uint32]string)
    v6 := make(map[uint32]uint32)
    if false {
        for fi7 := 0; fi7 < 7; fi7++ {
            var v8 string = "foo bar"
            _ = fi7
            _ = v8
        }
    }
    v9 := GenS1{F0: -1}
    if true {
        v3 += (v3 & (^(v3)))
    }
    if true {
        var v10 int16 = v1
        switch 4294967295 {
        case -1:
            var v11 int32 = -1
            _ = v11
        case 4294967295:
            if true {
                for fi12 := 0; fi12 < 12; fi12++ {
                    switch fi12 {
                    case -1:
                        switch 9223372036854775807 {
                        case 1:
                            delete(v6, 360)
                        }
                    case 0:
                        for fi13 := 0; fi13 < 15; fi13++ {
                            switch "" {
                            case "\n\t":
                                if (4294967295 == 2) {
                                    v3 = 586
                                    v2 = (v2 << ((988) & 63))
                                }
                            case "a":
                                switch "" {
                                case "foo bar":
                                    v14 := GenS2{F0: 1, F1: 65535, F2: true}
                                    _ = v14
                                case "":
                                    v15 := GenS3{F0: 449, F1: (v3 - 127)}
                                    _ = v15
                                }
                            }
                            fmt.Println(fi12, v2, v1)
                            _ = fi13
                        }
                    default:
                        fmt.Println(v1, v1)
                    }
                    _ = fi12
                }
                switch -1 {
                case 364:
                    fmt.Println(v10, v10, v3)
                case 42:
                    v3 += (v3 << ((869) & 63))
                }
            } else {
                v16 := []int{(-(v3))}
                if (!true) {
                    v3 = v3
                    fmt.Println(v2, v3, v3)
                }
                fmt.Println(v2, v10)
                _ = v16
            }
        case 0:
            switch (v3 & (v3 + v3)) {
            case 799:
                for fi17 := 0; fi17 < 6; fi17++ {
                    for fi18 := 0; fi18 < 8; fi18++ {
                        if false {
                            fmt.Println(fi18)
                        }
                        var v19 uint16 = 909
                        _ = fi18
                        _ = v19
                    }
                    _ = fi17
                }
            default:
                if ((v2 < (^(v2))) || (592 != 32767)) {
                    if (!true) {
                        for fi20 := 0; fi20 < 4; fi20++ {
                            if true {
                                v21 := GenS2{F0: 128, F1: 216, F2: (!(75 < 255))}
                                var v22 float32 = (0.1 / float32((3.14159 * (1.5 * (-2.25 / (float64(1.5) + float64(0.1)))))))
                                _ = v21
                                _ = v22
                            }
                            v23 := GenS1{F0: v3}
                            _ = fi20
                            _ = v23
                        }
                        if (!(3.14159 == 1.5)) {
                            var v24 string = "hello"
                            _ = v24
                        }
                        for fi25 := 0; fi25 < 1; fi25++ {
                            v10 -= v1
                            _ = fi25
                        }
                    }
                    delete(v6, 1)
                }
            }
        }
        for fi26 := 0; fi26 < 4; fi26++ {
            for fi27 := 0; fi27 < 11; fi27++ {
                v28 := GenS4{F0: "foo bar"}
                var v29 bool = true
                _ = fi27
                _ = v28
                _ = v29
            }
            _ = fi26
        }
        _ = v10
    } else {
        v2 -= v2
    }
    v4[32768] = (!(!((v1 >> ((0) & 63)) == (^(v1)))))
    v1 = 0
    switch "" {
    case "\n\t":
        v30 := GenS5{F0: (v2 % (((v2 | (v2 - v2))) | 1))}
        _ = v30
    case "a":
        var v31 uint8 = 114
        _ = v31
    default:
        for fi32 := 0; fi32 < 8; fi32++ {
            v33 := GenS6{F0: -1, F1: (true && (!("\n\t" < "")))}
            var v34 uint32 = 2147483648
            fmt.Println(v3, fi32)
            _ = fi32
            _ = v33
            _ = v34
        }
    }
    var v35 float64 = 1e10
    for fi36 := 0; fi36 < 4; fi36++ {
        v37 := GenS7{F0: 128, F1: 363, F2: 1}
        _ = fi36
        _ = v37
    }
    var v38 uint64 = 4294967295
    switch v3 {
    case -2147483648:
        if (!("a" < "hello")) {
            v39 := []int{0, v3, (v3 ^ (v3 << ((18446744073709551615) & 63)))}
            for fi40 := 0; fi40 < 4; fi40++ {
                if (!true) {
                    v41 := GenS8{F0: 42, F1: 1, F2: 291}
                    _ = v41
                } else {
                    for fi42 := 0; fi42 < 11; fi42++ {
                        if false {
                            var v43 uint64 = v38
                            v35 = v35
                            _ = v43
                        }
                        v44 := GenS9{F0: -128}
                        var v45 uint8 = 128
                        _ = fi42
                        _ = v44
                        _ = v45
                    }
                    for fi46 := 0; fi46 < 14; fi46++ {
                        var v47 int = (fi40 + (fi40 & fi46))
                        var v48 bool = true
                        if v48 {
                            if (925 == 32767) {
                                var v49 int = 83
                                v2 += v2
                                if v48 {
                                    if (true && v48) {
                                        v2 -= (v2 + -1)
                                        v50 := GenS10{F0: v1, F1: (^(v2)), F2: -1}
                                        _ = v50
                                    }
                                    switch "" {
                                    case "":
                                        v3 = fi40
                                    case "a":
                                        v1 += (v1 & (v1 << ((2) & 63)))
                                    case "foo bar":
                                        var v51 uint32 = 4294967295
                                        _ = v51
                                    default:
                                        if true {
                                            var v52 uint16 = 0
                                            _ = v52
                                        } else {
                                            if ((!(255 != 26)) || false) {
                                                var v53 string = ("hello" + "hello")
                                                v54 := GenS2{F0: 128, F1: 87, F2: true}
                                                v48 = v48
                                                var v55 uint64 = v38
                                                _ = v53
                                                _ = v54
                                                _ = v55
                                            }
                                        }
                                    }
                                } else {
                                    var v56 int32 = -770
                                    v3 = v47
                                    switch (^(fi40)) {
                                    case 2147483647:
                                        var v57 int32 = (v56 << ((1) & 63))
                                        _ = v57
                                    case 32767:
                                        if v48 {
                                            v3 += 127
                                            if false {
                                                var v58 float32 = (-2.25 / 1e10)
                                                _ = v58
                                            } else {
                                                if v48 {
                                                    v49 = (v3 - (v3 + v3))
                                                    var v59 uint64 = v38
                                                    _ = v59
                                                }
                                            }
                                            v38 = v38
                                        }
                                        if v48 {
                                            var v60 uint32 = 2147483648
                                            switch 127 {
                                            case 32768:
                                                v61 := GenS11{F0: 9223372036854775807, F1: v56}
                                                _ = v61
                                            case -1:
                                                var v62 string = ("foo bar" + "\n\t")
                                                _ = v62
                                            }
                                            _ = v60
                                        } else {
                                            v49 = v3
                                            v63 := GenS3{F0: 65535, F1: v47}
                                            if (!(v49 == (v49 | (fi46 + 128)))) {
                                                if v48 {
                                                    if (65535 != 4294967295) {
                                                        var v64 int64 = int64(fi46)
                                                        var v65 int16 = -32768
                                                        _ = v64
                                                        _ = v65
                                                    }
                                                    _ = len(v5)
                                                }
                                                switch fi46 {
                                                case 233:
                                                    if v48 {
                                                        v35 = ((((v35 / 1e10) - v35) + v35) + v35)
                                                        var v66 uint = 18446744073709551615
                                                        _ = v66
                                                    } else {
                                                        var v67 uint64 = (v38 * v38)
                                                        switch "foo bar" {
                                                        case "":
                                                            var v68 float64 = 0.0
                                                            _ = v68
                                                        case "\n\t":
                                                            if v48 {
                                                                if ((!false) && ((((v47 >> ((18446744073709551615) & 63)) > -1) && false) && false)) {
                                                                    v69 := v4[(fi40 | v3)]
                                                                    v70 := GenS8{F0: 0, F1: 112, F2: 2147483647}
                                                                    v71 := GenS2{F0: 2, F1: 1, F2: true}
                                                                    _ = v69
                                                                    _ = v70
                                                                    _ = v71
                                                                }
                                                                v72 := GenS12{F0: -71}
                                                                v73 := v6[2147483648]
                                                                fmt.Println(v38)
                                                                if v48 {
                                                                    if (((-115 > v2) && false) || v48) {
                                                                        v74 := GenS6{F0: (^(v56)), F1: false}
                                                                        fmt.Println(v35, v67, v1)
                                                                        fmt.Println(fi40, fi40, fi40)
                                                                        _ = v74
                                                                    }
                                                                    v73 = (v73 * v73)
                                                                    var v75 string = "hello"
                                                                    _ = v75
                                                                } else {
                                                                    if (v48 || true) {
                                                                        v35 = -2.25
                                                                        var v76 string = ("hello" + ("\n\t" + ("foo bar" + "hello")))
                                                                        _ = v76
                                                                    }
                                                                }
                                                                _ = v72
                                                                _ = v73
                                                            }
                                                        case "a":
                                                            v47 = fi40
                                                        }
                                                        fmt.Println(v1, v3)
                                                        _ = v67
                                                    }
                                                case 32767:
                                                    v77 := GenS11{F0: 42, F1: -1}
                                                    _ = v77
                                                case -2147483648:
                                                    fmt.Println(v56, fi46, v1)
                                                }
                                                if false {
                                                    if (0 <= 137) {
                                                        var v78 float64 = -2.25
                                                        var v79 float64 = v78
                                                        _ = len(v39)
                                                        _ = v78
                                                        _ = v79
                                                    } else {
                                                        var v80 float64 = (-2.25 + (1.5 + float64((0.0 / 1e10))))
                                                        var v81 uint = 18446744073709551615
                                                        if v48 {
                                                            delete(v6, 2)
                                                            switch "" {
                                                            case "\n\t":
                                                                if ("a" <= ("\n\t" + "\n\t")) {
                                                                    if v48 {
                                                                        var v82 int = v3
                                                                        _ = v82
                                                                    } else {
                                                                        fmt.Println(v35)
                                                                        v83 := GenS9{F0: 53}
                                                                        _ = v83
                                                                    }
                                                                    switch (fi46 >> ((v81) & 63)) {
                                                                    case 127:
                                                                        fmt.Println(v35)
                                                                    }
                                                                } else {
                                                                    var v84 float32 = 0.0
                                                                    v39 = append(v39, fi46)
                                                                    v48 = false
                                                                    _ = v84
                                                                }
                                                            default:
                                                                if (v48 && false) {
                                                                    if (1e10 >= float32(float64(0.1))) {
                                                                        if v48 {
                                                                            var v85 string = "foo bar"
                                                                            v86 := GenS9{F0: 100}
                                                                            _ = v85
                                                                            _ = v86
                                                                        }
                                                                    }
                                                                    switch "\n\t" {
                                                                    case "":
                                                                        v39 = append(v39, 9223372036854775807)
                                                                    case "foo bar":
                                                                        var v87 int = v47
                                                                        _ = v87
                                                                    }
                                                                    v88 := GenS12{F0: 28}
                                                                    v89 := GenS13{F0: (^(v56))}
                                                                    _ = v88
                                                                    _ = v89
                                                                } else {
                                                                    fmt.Println(v80, v47, v47)
                                                                }
                                                            }
                                                        } else {
                                                            v90 := GenS5{F0: v2}
                                                            _ = v90
                                                        }
                                                        var v91 int8 = -128
                                                        _ = v80
                                                        _ = v81
                                                        _ = v91
                                                    }
                                                    var v92 uint8 = 128
                                                    _ = v92
                                                } else {
                                                    var v93 uint16 = 0
                                                    if (!((v2 == 1) || ((fi46 % ((v3) | 1)) < fi40))) {
                                                        _ = len(v5)
                                                        if true {
                                                            if false {
                                                                v4[(fi46 % (((v47 << ((1) & 63))) | 1))] = v48
                                                                v94 := GenS5{F0: v2}
                                                                v95 := GenS14{F0: 255, F1: 262, F2: (v3 & 4294967295)}
                                                                switch "" {
                                                                case "hello":
                                                                    v96 := v6[2147483647]
                                                                    _ = v96
                                                                case "\n\t":
                                                                    v47 -= int(v38)
                                                                    _ = len(v39)
                                                                    v47 = fi46
                                                                default:
                                                                    var v97 uint64 = 685
                                                                    _ = v97
                                                                }
                                                                _ = v94
                                                                _ = v95
                                                            } else {
                                                                v98 := GenS15{F0: (^(v93)), F1: (253 == (fi40 >> ((2) & 63)))}
                                                                v48 = v48
                                                                var v99 int64 = 42
                                                                _ = v98
                                                                _ = v99
                                                            }
                                                        } else {
                                                            var v100 int64 = (v2 >> ((1) & 63))
                                                            v101 := GenS14{F0: 42, F1: v56, F2: (fi46 / ((v3) | 1))}
                                                            _ = len(v5)
                                                            _ = v100
                                                            _ = v101
                                                        }
                                                        var v102 int64 = (v2 % ((v2) | 1))
                                                        _ = v102
                                                    }
                                                    _ = v93
                                                }
                                            }
                                            fmt.Println(v48, v1)
                                            var v103 uint64 = 794
                                            switch v3 {
                                            case 43:
                                                if len(v39) > 0 {
                                                    v39[0] = (fi40 + -483)
                                                }
                                            case 2147483647:
                                                var v104 float32 = -2.25
                                                _ = v104
                                            case 9223372036854775807:
                                                v105 := GenS14{F0: 255, F1: (v56 | (v56 - (v56 << ((1) & 63)))), F2: fi46}
                                                _ = v105
                                            }
                                            if len(v39) >= 2 {
                                                v39 = v39[1:len(v39)]
                                            }
                                            _ = v103
                                            _ = v63
                                        }
                                        delete(v4, 4294967295)
                                    case 127:
                                        v49 -= 4294967295
                                    }
                                    v56 -= v56
                                    _ = v56
                                }
                                _ = v49
                            } else {
                                var v106 int32 = int32(v38)
                                v2 -= -9223372036854775808
                                var v107 uint64 = v38
                                _ = v106
                                _ = v107
                            }
                            v108 := GenS16{F0: 32767, F1: (^(fi40))}
                            var v109 float32 = (-2.25 - (3.14159 + (-2.25 + 1.5)))
                            _ = v108
                            _ = v109
                        }
                        _ = fi46
                        _ = v47
                        _ = v48
                    }
                    fmt.Println(v1, v1)
                }
                v1 = v1
                v110 := GenS17{F0: 127, F1: 743}
                _ = fi40
                _ = v110
            }
            _ = v39
        } else {
            if true {
                switch (-(v3)) {
                case 127:
                    for fi111 := 0; fi111 < 13; fi111++ {
                        switch ("" + "") {
                        case "hello":
                            v5[297] = (("hello" + "a") + ("a" + (("" + "\n\t") + ("" + "foo bar"))))
                        case "":
                            for fi112 := 0; fi112 < 13; fi112++ {
                                v1 = -32768
                                var v113 int16 = (^(v1))
                                v1 += 32767
                                if false {
                                    v114 := []int{(fi111 * -383), (fi112 - fi111), fi112}
                                    var v115 float64 = (v35 + ((1e10 - (v35 + (v35 - v35))) / v35))
                                    _ = v114
                                    _ = v115
                                }
                                _ = fi112
                                _ = v113
                            }
                        default:
                            var v116 uint64 = v38
                            _ = v116
                        }
                        v117 := GenS6{F0: 2147483647, F1: true}
                        var v118 string = ""
                        _ = fi111
                        _ = v117
                        _ = v118
                    }
                case 995:
                    fmt.Println(v2)
                }
                v1 = v1
                var v119 uint8 = 18
                _ = v119
            } else {
                var v120 uint8 = 127
                _ = v120
            }
        }
    case 128:
        v121 := GenS14{F0: 1, F1: -2147483648, F2: v3}
        _ = v121
    }
    v38 = (v38 & 1)
    v122 := GenS18{F0: 0, F1: true}
    var v123 uint64 = (v38 * v38)
    var v124 uint = 1
    v125 := GenS19{F0: 188, F1: 253, F2: (-(v3))}
    v126 := GenS16{F0: 32767, F1: -2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("v123:", v123)
    fmt.Println("v124:", v124)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v122
    _ = v123
    _ = v124
    _ = v125
    _ = v126
    _ = v2
    _ = v3
    _ = v35
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
