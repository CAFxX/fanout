package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 string }
type GenS2 struct { F0 int; F1 int64; F2 uint8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint64; F1 int16 }
type GenS6 struct { F0 uint; F1 bool }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int; F1 int32 }
type GenS9 struct { F0 int32; F1 uint }
type GenS10 struct { F0 uint32; F1 int64 }
func genFn0() uint {
    var v7 uint16 = 65535
    v8 := []int{-623, 42}
    _ = len(v8)
    v9 := GenS1{F0: -408, F1: 2147483647, F2: ((("foo bar" + "hello") + "foo bar") + ("foo bar" + ("hello" + ("" + "hello"))))}
    v7 = (v7 | 725)
    _ = v7
    _ = v8
    _ = v9
    return 65535
}
func genFn1() uint64 {
    if (((float32(((1.5 * 0.1) - (float64((-2.25 * 0.0)) - -2.25))) - (((((1e10 * 0.0) / 3.14159) + (-2.25 / ((0.1 + float32(float64(3.14159))) / 3.14159))) / (((float32(1e10) + (0.1 + 1.5)) - 0.1) * 1e10)) + (3.14159 + 3.14159))) - -2.25) >= 1e10) {
        v10 := GenS1{F0: 32767, F1: 32767, F2: ""}
        if true {
            var v11 uint16 = 580
            v11 += (v11 << ((1) & 63))
            v11 -= 119
            _ = v11
        } else {
            v12 := []int{-2147483648, 32768, 42}
            for fi13 := 0; fi13 < 2; fi13++ {
                fmt.Println(fi13)
                fmt.Println(fi13)
                var v14 bool = false
                _ = fi13
                _ = v14
            }
            v15 := GenS2{F0: -1, F1: 1, F2: 127}
            _ = v12
            _ = v15
        }
        var v16 uint = 42
        _ = v10
        _ = v16
    } else {
        var v17 float32 = (float32(float64(1e10)) - 1e10)
        if true {
            var v18 int = -792
            var v19 int32 = 42
            switch (v18 ^ 32767) {
            case 9223372036854775807:
                v20 := []int{(v18 + (v18 | (v18 >> ((157) & 63)))), 0, v18, 127}
                _ = v20
            default:
                for fi21 := 0; fi21 < 15; fi21++ {
                    var v22 uint64 = 4294967295
                    _ = fi21
                    _ = v22
                }
            }
            _ = v18
            _ = v19
        } else {
            v23 := []int{2147483647, -467, 9223372036854775807}
            _ = v23
        }
        _ = v17
    }
    var v24 float64 = 1e10
    v25 := []int{-981, 0, 32768, -1}
    _ = v24
    _ = v25
    return 4294967295
}

func main() {
    var v1 int32 = -1
    var v2 bool = true
    v3 := []int{}
    v4 := []int{32767, 230, int(v1)}
    var v5 uint8 = 0
    var v6 int32 = (v1 / (((v1 & v1)) | 1))
    if (!false) {
        fmt.Println(v2)
        switch "hello" {
        case "":
            if false {
                fmt.Println(v5, v1)
                switch 110 {
                case 127:
                    for fi26 := 0; fi26 < 16; fi26++ {
                        for fi27 := 0; fi27 < 17; fi27++ {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            _ = fi27
                        }
                        for fi28 := 0; fi28 < 19; fi28++ {
                            if true {
                                v5 += 127
                                if (255 > (v5 & (^(v5)))) {
                                    switch (fi26 | (-(fi26))) {
                                    case -1:
                                        switch "\n\t" {
                                        case "hello":
                                            v2 = v2
                                        case "foo bar":
                                            v5 -= v5
                                        }
                                    case 127:
                                        v29 := GenS2{F0: (fi28 + 42), F1: 9223372036854775807, F2: v5}
                                        _ = v29
                                    default:
                                        v30 := GenS1{F0: fi26, F1: fi28, F2: "hello"}
                                        _ = v30
                                    }
                                    _ = len(v3)
                                } else {
                                    var v31 uint8 = v5
                                    v5 = 42
                                    var v32 float32 = 0.0
                                    _ = v31
                                    _ = v32
                                }
                                _ = len(v4)
                                if true {
                                    if len(v4) > 0 {
                                        v4[0] = 42
                                    }
                                    v5 = v5
                                    v6 = v6
                                }
                                switch -435 {
                                case 32768:
                                    var v33 uint16 = 255
                                    _ = v33
                                case 32767:
                                    v2 = false
                                default:
                                    var v34 uint = 4294967295
                                    _ = v34
                                }
                            } else {
                                if false {
                                    if len(v4) > 0 {
                                        v4[0] = 128
                                    }
                                    if (genFn1() == 0) {
                                        if (1 <= uint64(v5)) {
                                            var v35 uint = genFn0()
                                            var v36 int8 = 127
                                            var v37 uint = (^(v35))
                                            v38 := GenS1{F0: (^(fi28)), F1: (fi26 + fi28), F2: (("foo bar" + "foo bar") + "hello")}
                                            _ = v35
                                            _ = v36
                                            _ = v37
                                            _ = v38
                                        }
                                        _ = len(v3)
                                        var v39 float32 = (1.5 / ((((0.0 * ((0.0 - (float32(-2.25) * ((((float32((3.14159 / 0.1)) - 3.14159) * 3.14159) / 0.1) - 1.5))) - 0.1)) + 0.0) * 1.5) + (float32(float64(0.0)) - (-2.25 / float32(float64((3.14159 + 3.14159)))))))
                                        _ = v39
                                    }
                                    var v40 uint16 = 0
                                    _ = v40
                                } else {
                                    v2 = v2
                                }
                                _ = len(v4)
                                fmt.Println(v5, fi28, fi26)
                            }
                            v41 := GenS3{F0: v5}
                            v1 = v1
                            var v42 int = fi28
                            _ = fi28
                            _ = v41
                            _ = v42
                        }
                        var v43 int = fi26
                        _ = fi26
                        _ = v43
                    }
                default:
                    v5 += (v5 + (^(v5)))
                }
            } else {
                if ("\n\t" == "a") {
                    fmt.Println(v6, v6, v6)
                    var v44 string = "\n\t"
                    _ = v44
                } else {
                    v2 = true
                    fmt.Println(v2)
                    var v45 int64 = 0
                    _ = v45
                }
                var v46 uint64 = 2
                _ = v46
            }
        case "hello":
            _ = len(v4)
        default:
            var v47 int8 = -113
            _ = v47
        }
    } else {
        if (!v2) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            for fi48 := 0; fi48 < 17; fi48++ {
                fmt.Println(v1, v1)
                if len(v3) > 0 {
                    v3[0] = 32768
                }
                var v49 float64 = (-2.25 + ((float64(-2.25) * 1.5) / (1e10 / -2.25)))
                _ = fi48
                _ = v49
            }
            v5 = (^(v5))
        } else {
            var v50 int16 = 878
            _ = v50
        }
        if false {
            v51 := GenS3{F0: (v5 ^ (v5 | (v5 / ((v5) | 1))))}
            _ = v51
        } else {
            var v52 int8 = -128
            var v53 uint16 = 1
            _ = v52
            _ = v53
        }
    }
    if v2 {
        var v54 uint16 = 65535
        if v2 {
            if (("" + "a") == ("\n\t" + "")) {
                for fi55 := 0; fi55 < 20; fi55++ {
                    for fi56 := 0; fi56 < 3; fi56++ {
                        if (-84 < -819) {
                            if (!v2) {
                                v57 := GenS4{F0: v5}
                                var v58 uint16 = v54
                                var v59 int16 = 42
                                v60 := GenS5{F0: 2, F1: (-(v59))}
                                v54 = v54
                                if v2 {
                                    fmt.Println(v5, fi55)
                                    switch -519 {
                                    case -872:
                                        v59 -= (^(v59))
                                        v61 := GenS6{F0: 729, F1: (((^(v6)) > (v6 & (v1 * (v6 | v1)))) && (!v2))}
                                        var v62 uint = 2
                                        _ = v61
                                        _ = v62
                                    case 2147483647:
                                        switch "foo bar" {
                                        case "foo bar":
                                            var v63 float32 = 1.5
                                            _ = v63
                                        }
                                    }
                                } else {
                                    var v64 int8 = 42
                                    v4 = append(v4, 127)
                                    _ = v64
                                }
                                _ = v57
                                _ = v58
                                _ = v59
                                _ = v60
                            } else {
                                _ = len(v4)
                                switch (fi56 / ((-997) | 1)) {
                                case 245:
                                    var v65 uint64 = 4294967295
                                    _ = v65
                                }
                            }
                        }
                        v1 = v1
                        _ = fi56
                    }
                    v6 -= 2147483647
                    var v66 int64 = 1
                    _ = fi55
                    _ = v66
                }
            } else {
                if (4294967295 == genFn1()) {
                    v54 += (v54 - 32767)
                    v5 = 231
                } else {
                    var v67 uint32 = 683
                    v3 = append(v3, -832)
                    _ = v67
                }
                if true {
                    var v68 uint64 = 18446744073709551615
                    if false {
                        v69 := GenS7{F0: (v5 & (^(v5)))}
                        var v70 bool = v2
                        switch "hello" {
                        case "foo bar":
                            v2 = (-2147483648 >= -9223372036854775808)
                        }
                        _ = v69
                        _ = v70
                    } else {
                        for fi71 := 0; fi71 < 4; fi71++ {
                            for fi72 := 0; fi72 < 4; fi72++ {
                                v73 := GenS4{F0: (^(v5))}
                                switch int(v54) {
                                case 172:
                                    v1 -= -2147483648
                                }
                                var v74 int = fi72
                                if false {
                                    switch (fi72 * 770) {
                                    case 128:
                                        if (("foo bar" + "hello") < ("foo bar" + ("hello" + "hello"))) {
                                            v74 += (fi71 & fi71)
                                        }
                                    case 32767:
                                        if len(v3) > 0 {
                                            v3[0] = (fi71 + int(v1))
                                        }
                                    default:
                                        if v2 {
                                            var v75 float64 = 0.1
                                            var v76 uint16 = 255
                                            v1 += 42
                                            v1 -= v6
                                            _ = v75
                                            _ = v76
                                        } else {
                                            if (uint64(v5) <= v68) {
                                                var v77 float64 = (-2.25 - 0.0)
                                                v54 += 0
                                                _ = v77
                                            }
                                            var v78 uint32 = 2147483647
                                            _ = v78
                                        }
                                        var v79 uint64 = (^(v68))
                                        switch (^(fi71)) {
                                        case 0:
                                            v80 := GenS4{F0: (v5 * (v5 / (((v5 / ((0) | 1))) | 1)))}
                                            _ = v80
                                        case 128:
                                            var v81 string = "a"
                                            v5 += (v5 / ((v5) | 1))
                                            switch (fi71 / ((fi72) | 1)) {
                                            case 674:
                                                switch ("a" + "foo bar") {
                                                case "hello":
                                                    var v82 int8 = 42
                                                    _ = v82
                                                case "foo bar":
                                                    if v2 {
                                                        v2 = (59 <= 100)
                                                        if (fi71 > 196) {
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                        }
                                                    }
                                                case "a":
                                                    if true {
                                                        var v83 int8 = 48
                                                        _ = v83
                                                    }
                                                default:
                                                    if len(v4) > 0 {
                                                        v4[0] = (v74 - fi71)
                                                    }
                                                }
                                                var v84 uint8 = v5
                                                _ = v84
                                            }
                                            _ = v81
                                        case -138:
                                            if (-1 > (fi72 >> ((42) & 63))) {
                                                if len(v3) > 0 {
                                                    v3[0] = int(v1)
                                                }
                                            } else {
                                                v68 = uint64(v5)
                                                if v2 {
                                                    var v85 uint64 = 719
                                                    v1 = -313
                                                    fmt.Println(v5, v79)
                                                    v85 += v68
                                                    _ = v85
                                                }
                                            }
                                        }
                                        _ = v79
                                    }
                                }
                                _ = fi72
                                _ = v73
                                _ = v74
                            }
                            var v86 uint64 = (v68 / (((^(v68))) | 1))
                            _ = fi71
                            _ = v86
                        }
                        if (9223372036854775807 < 0) {
                            v87 := GenS4{F0: 128}
                            switch int(v5) {
                            case -9223372036854775808:
                                var v88 string = "\n\t"
                                _ = v88
                            case -243:
                                v68 = 2
                            default:
                                fmt.Println(v6, v2, v1)
                            }
                            _ = v87
                        } else {
                            var v89 int32 = (v6 / ((v6) | 1))
                            _ = v89
                        }
                    }
                    if v2 {
                        if v2 {
                            for fi90 := 0; fi90 < 11; fi90++ {
                                for fi91 := 0; fi91 < 1; fi91++ {
                                    if v2 {
                                        fmt.Println(v2, v1, fi91)
                                        _ = len(v3)
                                        var v92 int32 = v1
                                        _ = v92
                                    }
                                    var v93 string = ""
                                    _ = fi91
                                    _ = v93
                                }
                                switch (("a" + "hello") + (("\n\t" + "hello") + "")) {
                                case "":
                                    var v94 uint32 = 2
                                    _ = v94
                                case "\n\t":
                                    v95 := GenS5{F0: (v68 + v68), F1: -740}
                                    _ = v95
                                case "foo bar":
                                    v1 += v1
                                }
                                v96 := GenS5{F0: (v68 | v68), F1: 389}
                                _ = fi90
                                _ = v96
                            }
                        }
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    _ = v68
                }
            }
            var v97 int8 = -54
            _ = v97
        } else {
            var v98 float64 = 1.5
            _ = v98
        }
        _ = v54
    }
    for fi99 := 0; fi99 < 20; fi99++ {
        if v2 {
            var v100 int = fi99
            for fi101 := 0; fi101 < 5; fi101++ {
                v6 += (-(v6))
                switch 127 {
                case 2147483647:
                    fmt.Println(v1, v2, v5)
                }
                _ = fi101
            }
            _ = v100
        }
        _ = fi99
    }
    for fi102 := 0; fi102 < 12; fi102++ {
        var v103 float32 = -2.25
        for fi104 := 0; fi104 < 16; fi104++ {
            var v105 int32 = (v6 % ((v6) | 1))
            if (724 < 1) {
                switch 0 {
                case -884:
                    fmt.Println(fi104, v103, v103)
                case 234:
                    v106 := GenS8{F0: fi102, F1: v6}
                    var v107 float32 = v103
                    _ = v106
                    _ = v107
                default:
                    v103 -= (3.14159 / v103)
                }
                if v2 {
                    if true {
                        fmt.Println(fi104, fi104)
                        switch (fi102 * (fi104 - (fi102 >> ((genFn0()) & 63)))) {
                        case 32768:
                            var v108 string = "\n\t"
                            _ = v108
                        default:
                            if (v2 || (((0.1 * (float64(v103) * -2.25)) * (0.0 * 0.1)) >= (0.1 + 0.1))) {
                                v2 = (float32((3.14159 * 1e10)) < (-2.25 * v103))
                                v109 := GenS9{F0: v6, F1: uint(v5)}
                                _ = v109
                            }
                        }
                    } else {
                        v110 := GenS5{F0: 1, F1: -579}
                        var v111 float32 = (((v103 / (1e10 - v103)) + 3.14159) * 1.5)
                        _ = len(v3)
                        _ = v110
                        _ = v111
                    }
                    v2 = ("a" <= "\n\t")
                }
                var v112 uint16 = 42
                _ = v112
            }
            if v2 {
                v113 := GenS10{F0: 37, F1: 0}
                _ = v113
            } else {
                if len(v4) > 0 {
                    v4[0] = (fi102 % ((2147483647) | 1))
                }
            }
            _ = fi104
            _ = v105
        }
        _ = fi102
        _ = v103
    }
    v114 := GenS4{F0: 98}
    switch 42 {
    case -492:
        for fi115 := 0; fi115 < 8; fi115++ {
            for fi116 := 0; fi116 < 18; fi116++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi116
            }
            _ = fi115
        }
    }
    for fi117 := 0; fi117 < 16; fi117++ {
        v118 := GenS5{F0: 4294967295, F1: -1}
        _ = fi117
        _ = v118
    }
    if (v2 && ((v2 || (!v2)) || false)) {
        fmt.Println(v2, v6)
        v5 = 255
    } else {
        v2 = (1 <= 32767)
        if ((550 == 1) || v2) {
            var v119 int = -9223372036854775808
            _ = v119
        }
        v1 += v6
    }
    var v120 int = 32767
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi121 := 0; fi121 < 13; fi121++ {
        fmt.Println(v5, v2)
        _ = fi121
    }
    v1 = (v1 ^ 884)
    var v122 uint = 255
    if true {
        for fi123 := 0; fi123 < 5; fi123++ {
            v120 -= (v120 >> (((^(v122))) & 63))
            switch ("a" + "a") {
            case "hello":
                for fi124 := 0; fi124 < 8; fi124++ {
                    var v125 uint = (v122 >> ((0) & 63))
                    var v126 int64 = -9223372036854775808
                    if false {
                        fmt.Println(v1, v122)
                    } else {
                        switch (fi123 ^ (fi124 >> (((^(v125))) & 63))) {
                        case 32767:
                            var v127 string = "foo bar"
                            _ = v127
                        case -2147483648:
                            var v128 int32 = v6
                            _ = v128
                        default:
                            var v129 uint = v125
                            _ = v129
                        }
                        v126 = v126
                        v122 -= (^(v125))
                    }
                    v1 = v6
                    _ = fi124
                    _ = v125
                    _ = v126
                }
            case "foo bar":
                var v130 bool = v2
                _ = v130
            case "":
                var v131 float32 = 0.1
                _ = v131
            default:
                v5 = v5
            }
            var v132 uint8 = v5
            _ = fi123
            _ = v132
        }
        var v133 int8 = 1
        _ = v133
    }
    for fi134 := 0; fi134 < 20; fi134++ {
        if len(v4) > 0 {
            v4[0] = v120
        }
        _ = fi134
    }
    _ = len(v3)
    switch "foo bar" {
    case "\n\t":
        v135 := GenS6{F0: genFn0(), F1: (-2.25 != float64(3.14159))}
        _ = v135
    default:
        v3 = append(v3, v120)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v120:", v120)
    fmt.Println("v122:", v122)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v114
    _ = v120
    _ = v122
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
