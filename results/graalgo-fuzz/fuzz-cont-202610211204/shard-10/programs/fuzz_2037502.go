package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 uint64 }
type GenS2 struct { F0 string; F1 int16; F2 bool }
type GenS3 struct { F0 int32; F1 int8 }
type GenS4 struct { F0 uint8; F1 bool; F2 uint32 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint; F1 string }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int64; F1 int; F2 uint16 }
type GenS10 struct { F0 uint8; F1 uint64 }
type GenS11 struct { F0 string; F1 int64 }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 bool }
type GenS14 struct { F0 string }
func genFn0() int {
    if true {
        var v4 int32 = -2147483648
        _ = v4
    } else {
        for fi5 := 0; fi5 < 10; fi5++ {
            switch "\n\t" {
            case "a":
                v6 := []int{}
                _ = v6
            case "":
                switch "" {
                case "a":
                    for fi7 := 0; fi7 < 19; fi7++ {
                        switch fi7 {
                        case 0:
                            var v8 float32 = float32(1.5)
                            _ = v8
                        }
                        _ = fi7
                    }
                case "":
                    if (611 != 0) {
                        var v9 bool = (-1 <= 1)
                        if v9 {
                            var v10 string = "\n\t"
                            var v11 int32 = -995
                            var v12 float32 = (float32(((1.5 / float64(-2.25)) + float64(1.5))) - (1.5 * (1e10 / 0.1)))
                            _ = v10
                            _ = v11
                            _ = v12
                        }
                        v13 := []int{780, -1, (fi5 * fi5), fi5}
                        _ = v13
                        _ = v9
                    } else {
                        var v14 int = fi5
                        v14 += v14
                        if (9223372036854775807 == 64) {
                            v15 := []int{226, (v14 & v14)}
                            v14 -= fi5
                            var v16 bool = false
                            _ = v15
                            _ = v16
                        }
                        _ = v14
                    }
                case "hello":
                    for fi17 := 0; fi17 < 17; fi17++ {
                        var v18 uint8 = 179
                        _ = fi17
                        _ = v18
                    }
                default:
                    var v19 uint16 = 1
                    _ = v19
                }
            case "hello":
                if true {
                    var v20 uint64 = 1
                    v21 := GenS1{F0: 2, F1: v20, F2: 4294967295}
                    fmt.Println(v20, v20)
                    _ = v20
                    _ = v21
                } else {
                    v22 := []int{(fi5 % ((-2147483648) | 1)), -1, (-(fi5)), fi5}
                    fmt.Println(fi5)
                    _ = v22
                }
            default:
                for fi23 := 0; fi23 < 4; fi23++ {
                    var v24 uint16 = 32767
                    if ((230 >= -836) || false) {
                        v24 = v24
                    }
                    _ = fi23
                    _ = v24
                }
            }
            _ = fi5
        }
    }
    if true {
        for fi25 := 0; fi25 < 12; fi25++ {
            switch (fi25 * fi25) {
            case -9223372036854775808:
                switch (^(fi25)) {
                case -1:
                    v26 := GenS1{F0: 230, F1: 4294967295, F2: 0}
                    _ = v26
                }
            case 439:
                for fi27 := 0; fi27 < 19; fi27++ {
                    var v28 bool = (2 <= 255)
                    _ = fi27
                    _ = v28
                }
            case 927:
                var v29 uint = 4294967295
                _ = v29
            default:
                for fi30 := 0; fi30 < 20; fi30++ {
                    var v31 int8 = 100
                    var v32 int32 = 1
                    _ = fi30
                    _ = v31
                    _ = v32
                }
            }
            v33 := GenS2{F0: "foo bar", F1: 0, F2: (3.14159 < 1.5)}
            var v34 uint16 = 65535
            _ = fi25
            _ = v33
            _ = v34
        }
        if (!(!(127 >= -69))) {
            switch -889 {
            case -2147483648:
                v35 := []int{-273, 855, 9223372036854775807}
                _ = v35
            case -17:
                v36 := GenS3{F0: -2147483648, F1: -128}
                _ = v36
            case 42:
                var v37 int = 1
                _ = v37
            }
            var v38 uint = 18446744073709551615
            v39 := []int{128, -1}
            _ = v38
            _ = v39
        } else {
            for fi40 := 0; fi40 < 6; fi40++ {
                v41 := GenS2{F0: ("a" + ("" + (("" + "") + "a"))), F1: 370, F2: false}
                var v42 int = fi40
                var v43 int32 = -684
                _ = fi40
                _ = v41
                _ = v42
                _ = v43
            }
        }
    } else {
        var v44 int = -2147483648
        if true {
            var v45 uint = 248
            v46 := GenS4{F0: 128, F1: false, F2: 42}
            var v47 int8 = 127
            _ = v45
            _ = v46
            _ = v47
        }
        _ = v44
    }
    v48 := []int{-2147483648}
    var v49 int64 = 159
    v50 := GenS1{F0: 127, F1: 1, F2: 0}
    _ = v48
    _ = v49
    _ = v50
    return 4294967295
}
func genFn1(p0 int32, p1 int16, p2 uint) int {
    v51 := GenS5{F0: 42}
    fmt.Println(p0, p1, p0)
    for fi52 := 0; fi52 < 6; fi52++ {
        v53 := []int{genFn0(), (-(fi52)), genFn0(), (fi52 + fi52)}
        _ = fi52
        _ = v53
    }
    v54 := GenS4{F0: 168, F1: true, F2: 2147483648}
    v55 := GenS1{F0: 0, F1: 0, F2: uint64(p2)}
    _ = p0
    _ = p1
    _ = p2
    _ = v51
    _ = v54
    _ = v55
    return int(p2)
}

func main() {
    var v1 uint64 = 268
    var v2 uint16 = 2
    var v3 int8 = -124
    for fi56 := 0; fi56 < 2; fi56++ {
        var v57 float32 = 0.1
        for fi58 := 0; fi58 < 18; fi58++ {
            v59 := GenS5{F0: 4294967295}
            v60 := []int{}
            _ = fi58
            _ = v59
            _ = v60
        }
        _ = fi56
        _ = v57
    }
    if (118 < -380) {
        var v61 bool = true
        for fi62 := 0; fi62 < 10; fi62++ {
            var v63 uint16 = 65535
            switch (fi62 << ((65535) & 63)) {
            case 32768:
                v64 := GenS6{F0: 0}
                _ = v64
            case -503:
                v65 := GenS7{F0: 42, F1: ((("\n\t" + "foo bar") + "hello") + ("\n\t" + ""))}
                _ = v65
            default:
                var v66 uint16 = 65535
                _ = v66
            }
            _ = fi62
            _ = v63
        }
        v67 := GenS8{F0: "a"}
        _ = v61
        _ = v67
    } else {
        fmt.Println(v3, v2, v3)
    }
    v68 := GenS3{F0: -481, F1: v3}
    if (410 < 18446744073709551615) {
        v69 := []int{9223372036854775807, 166}
        _ = v69
    } else {
        var v70 int32 = -1
        _ = v70
    }
    for fi71 := 0; fi71 < 10; fi71++ {
        var v72 uint = 0
        var v73 bool = ((false || (true && (128 > 255))) || true)
        _ = fi71
        _ = v72
        _ = v73
    }
    fmt.Println(v2, v2)
    if true {
        var v74 uint16 = 0
        _ = v74
    } else {
        var v75 bool = true
        fmt.Println(v2, v3, v75)
        for fi76 := 0; fi76 < 18; fi76++ {
            v77 := GenS8{F0: ("hello" + "hello")}
            _ = fi76
            _ = v77
        }
        _ = v75
    }
    switch "hello" {
    case "":
        if true {
            if (1e10 == 1e10) {
                v1 -= (v1 + 733)
                v3 += (v3 ^ -1)
            }
            if true {
                switch -950 {
                case -9223372036854775808:
                    v78 := GenS5{F0: 0}
                    _ = v78
                case 2147483647:
                    if (false && (-1 >= -1)) {
                        fmt.Println(v2)
                        fmt.Println(v1, v1, v2)
                        var v79 uint = 1
                        _ = v79
                    }
                case 4294967295:
                    switch "\n\t" {
                    case "\n\t":
                        if (-2.25 < ((((float64(1e10) * float64((((3.14159 + (-2.25 - 0.0)) + 1.5) / ((3.14159 + 1e10) * (1e10 - 0.0))))) + (1.5 - 3.14159)) * (3.14159 - 1.5)) / 1.5)) {
                            if false {
                                var v80 int16 = -32768
                                for fi81 := 0; fi81 < 8; fi81++ {
                                    if (!((fi81 == (fi81 << ((594) & 63))) && ((!(false && ((false && true) || true))) || true))) {
                                        v1 = 220
                                    } else {
                                        if (128 <= 127) {
                                            var v82 int64 = -1
                                            _ = v82
                                        } else {
                                            var v83 float32 = float32(0.1)
                                            if true {
                                                var v84 float64 = 3.14159
                                                for fi85 := 0; fi85 < 15; fi85++ {
                                                    v3 -= (v3 * 127)
                                                    if true {
                                                        v84 -= v84
                                                        var v86 int16 = (^(v80))
                                                        if (true || (-2147483648 > 0)) {
                                                            if (true || false) {
                                                                fmt.Println(v80)
                                                                if (false && (65535 > v2)) {
                                                                    switch (fi85 & fi85) {
                                                                    case -2147483648:
                                                                        if ((v3 | v3) != (v3 << ((18446744073709551615) & 63))) {
                                                                            var v87 bool = false
                                                                            v1 = (v1 << ((2) & 63))
                                                                            _ = v87
                                                                        }
                                                                    case 128:
                                                                        var v88 int16 = 32767
                                                                        _ = v88
                                                                    case -9223372036854775808:
                                                                        if (!false) {
                                                                            if (!(755 >= 1)) {
                                                                                switch fi85 {
                                                                                case 127:
                                                                                    switch 32768 {
                                                                                    case -9223372036854775808:
                                                                                        v89 := []int{(fi85 << ((787) & 63)), fi81, 0, (fi85 << ((765) & 63))}
                                                                                        _ = v89
                                                                                    case 0:
                                                                                        var v90 string = "hello"
                                                                                        _ = v90
                                                                                    case 128:
                                                                                        if true {
                                                                                            v91 := []int{(fi85 ^ (fi81 / ((fi85) | 1)))}
                                                                                            fmt.Println(v86, v86)
                                                                                            var v92 float64 = v84
                                                                                            _ = v91
                                                                                            _ = v92
                                                                                        }
                                                                                    }
                                                                                default:
                                                                                    var v93 int32 = -930
                                                                                    _ = v93
                                                                                }
                                                                                v94 := []int{fi81}
                                                                                _ = v94
                                                                            }
                                                                            v95 := GenS9{F0: -9223372036854775808, F1: fi85, F2: (^(v2))}
                                                                            _ = v95
                                                                        }
                                                                    default:
                                                                        fmt.Println(v1, v83)
                                                                    }
                                                                    var v96 string = ("\n\t" + ("" + "hello"))
                                                                    v2 -= v2
                                                                    _ = v96
                                                                }
                                                                var v97 bool = false
                                                                v3 -= (v3 + v3)
                                                                _ = v97
                                                            }
                                                        } else {
                                                            var v98 uint16 = v2
                                                            _ = v98
                                                        }
                                                        switch (("\n\t" + "") + "\n\t") {
                                                        case "hello":
                                                            v99 := GenS2{F0: "", F1: 128, F2: true}
                                                            _ = v99
                                                        default:
                                                            if (v83 > v83) {
                                                                if ((v2 - v2) > 42) {
                                                                    v100 := []int{genFn1(599, -1, 4294967295), (fi85 & (fi85 / (((fi85 >> ((1) & 63))) | 1))), (fi85 / (((-(fi81))) | 1)), (fi85 << ((18446744073709551615) & 63))}
                                                                    if (true || (1 == 2147483647)) {
                                                                        v101 := GenS10{F0: 128, F1: (^(v1))}
                                                                        v1 = v1
                                                                        _ = v101
                                                                    } else {
                                                                        if (!false) {
                                                                            v83 += (1e10 + ((v83 + float32(1.5)) - 1.5))
                                                                            var v102 float64 = 1e10
                                                                            _ = v102
                                                                        }
                                                                    }
                                                                    if false {
                                                                        var v103 uint16 = v2
                                                                        switch "foo bar" {
                                                                        case "a":
                                                                            if (true && (1 != 1)) {
                                                                                var v104 uint64 = 466
                                                                                var v105 uint8 = 2
                                                                                if false {
                                                                                    if len(v100) >= 2 {
                                                                                        v100 = v100[1:len(v100)]
                                                                                    }
                                                                                    v106 := GenS1{F0: v105, F1: v104, F2: v1}
                                                                                    v2 = v2
                                                                                    _ = v106
                                                                                } else {
                                                                                    v80 -= 32767
                                                                                    v107 := GenS11{F0: "hello", F1: int64(v104)}
                                                                                    _ = v107
                                                                                }
                                                                                _ = v104
                                                                                _ = v105
                                                                            } else {
                                                                                var v108 bool = (false && ("" >= "foo bar"))
                                                                                _ = v108
                                                                            }
                                                                        default:
                                                                            v109 := GenS12{F0: 65535}
                                                                            _ = v109
                                                                        }
                                                                        _ = v103
                                                                    }
                                                                    _ = v100
                                                                } else {
                                                                    v110 := []int{}
                                                                    switch fi81 {
                                                                    case -192:
                                                                        v111 := GenS5{F0: 1}
                                                                        _ = v111
                                                                    }
                                                                    _ = v110
                                                                }
                                                                if (false || (false || false)) {
                                                                    var v112 string = ("" + "\n\t")
                                                                    _ = v112
                                                                } else {
                                                                    if (!false) {
                                                                        v113 := []int{fi81, fi81, (^(fi81)), (fi81 - (fi85 >> ((120) & 63)))}
                                                                        if (!(((0.0 - v83) <= (((v83 - float32((v84 / 1.0))) - 0.1) + float32(v84))) && ("hello" == "a"))) {
                                                                            v113 = append(v113, genFn0())
                                                                            if false {
                                                                                v114 := GenS6{F0: 42}
                                                                                _ = v114
                                                                            } else {
                                                                                var v115 int32 = -609
                                                                                _ = v115
                                                                            }
                                                                        }
                                                                        v86 = (v86 + (v80 % (((v86 * 1)) | 1)))
                                                                        _ = v113
                                                                    }
                                                                    switch "hello" {
                                                                    case "":
                                                                        v116 := []int{(^(fi85)), genFn0(), -1, genFn0()}
                                                                        _ = v116
                                                                    case "a":
                                                                        v2 = v2
                                                                    case "hello":
                                                                        fmt.Println(v1, v3, fi81)
                                                                    default:
                                                                        v117 := GenS2{F0: "", F1: -86, F2: true}
                                                                        _ = v117
                                                                    }
                                                                }
                                                                v118 := GenS2{F0: ("" + ""), F1: v86, F2: (3.14159 != v84)}
                                                                _ = v118
                                                            }
                                                        }
                                                        fmt.Println(fi81)
                                                        _ = v86
                                                    }
                                                    _ = fi85
                                                }
                                                for fi119 := 0; fi119 < 17; fi119++ {
                                                    v83 = v83
                                                    var v120 uint = 629
                                                    if true {
                                                        var v121 uint32 = 2147483648
                                                        var v122 float32 = v83
                                                        v123 := GenS13{F0: true}
                                                        fmt.Println(v121, v84, v121)
                                                        v124 := []int{(fi81 / (((fi119 % (((fi81 | (fi119 - fi81))) | 1))) | 1)), (fi81 * (-(fi119)))}
                                                        _ = v121
                                                        _ = v122
                                                        _ = v123
                                                        _ = v124
                                                    }
                                                    _ = fi119
                                                    _ = v120
                                                }
                                                _ = v84
                                            } else {
                                                v125 := GenS14{F0: "\n\t"}
                                                _ = v125
                                            }
                                            v126 := GenS4{F0: 1, F1: (255 == 128), F2: 2}
                                            _ = v126
                                            _ = v83
                                        }
                                        fmt.Println(v2)
                                    }
                                    _ = fi81
                                }
                                _ = v80
                            }
                        } else {
                            v2 = (v2 * v2)
                            var v127 bool = ("hello" > "a")
                            _ = v127
                        }
                    case "hello":
                        fmt.Println(v2, v2)
                    }
                default:
                    switch "a" {
                    case "\n\t":
                        if true {
                            for fi128 := 0; fi128 < 16; fi128++ {
                                fmt.Println(v1, v2, v1)
                                for fi129 := 0; fi129 < 12; fi129++ {
                                    if false {
                                        v2 -= v2
                                    } else {
                                        var v130 int64 = 1
                                        fmt.Println(v2)
                                        v131 := GenS1{F0: 127, F1: 395, F2: (v1 >> ((18446744073709551615) & 63))}
                                        _ = v130
                                        _ = v131
                                    }
                                    v132 := GenS8{F0: ""}
                                    _ = fi129
                                    _ = v132
                                }
                                v133 := []int{fi128, (-(fi128)), (fi128 - -2147483648), 1}
                                _ = fi128
                                _ = v133
                            }
                            var v134 int16 = 256
                            _ = v134
                        }
                    }
                }
                v135 := []int{}
                _ = v135
            } else {
                var v136 int8 = v3
                _ = v136
            }
        }
    }
    var v137 float64 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("v137:", v137)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v137
    _ = v2
    _ = v3
    _ = v68
}
