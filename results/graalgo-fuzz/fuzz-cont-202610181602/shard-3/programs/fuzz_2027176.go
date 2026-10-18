package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 uint32; F2 int64 }
type GenS3 struct { F0 bool; F1 uint }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int64; F1 int }
type GenS6 struct { F0 int16; F1 int }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 string; F1 uint8; F2 uint16 }
type GenS9 struct { F0 uint; F1 uint8 }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 string }
type GenS14 struct { F0 int16 }
func genFn0(p0 uint, p1 int32, p2 uint) int32 {
    if (true || (!(1000 > -654))) {
        for fi7 := 0; fi7 < 13; fi7++ {
            v8 := GenS1{F0: 601}
            fmt.Println(p1, p1)
            p2 = p0
            _ = fi7
            _ = v8
        }
        if (!false) {
            v9 := []int{32768, -9223372036854775808, 32767}
            v10 := GenS1{F0: -1}
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            _ = v10
            _ = v9
        }
    }
    v11 := []int{32767}
    var v12 float32 = (0.1 / (0.0 + (((3.14159 / -2.25) - (((0.1 - (0.0 + (float32(1.5) / 1.5))) - 0.0) + ((0.0 / float32(1.5)) - 1.5))) / float32(3.14159))))
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return p1
}

func main() {
    var v1 float64 = -2.25
    var v2 string = ("foo bar" + (("\n\t" + "foo bar") + "hello"))
    v3 := make(map[int]uint32)
    var v4 float64 = v1
    var v5 float64 = float64(3.14159)
    v6 := make(map[uint32]int)
    if (1 <= -1) {
        v13 := GenS1{F0: 1}
        var v14 int64 = 246
        switch v2 {
        case "foo bar":
            fmt.Println(v1)
        case "\n\t":
            fmt.Println(v1, v2, v14)
        case "a":
            var v15 int8 = -1
            _ = v15
        default:
            var v16 int32 = -1
            _ = v16
        }
        _ = v13
        _ = v14
    } else {
        for fi17 := 0; fi17 < 17; fi17++ {
            if ((true || true) && true) {
                v5 = 0.0
                v4 = ((v5 * ((0.0 - 3.14159) - v5)) / (v1 + v5))
            } else {
                var v18 float32 = 1.5
                var v19 uint64 = 17
                v20 := []int{232, 4294967295, (fi17 ^ (fi17 - (fi17 * -213)))}
                _ = v18
                _ = v19
                _ = v20
            }
            v1 = 0.0
            for fi21 := 0; fi21 < 6; fi21++ {
                if (127 > 42) {
                    switch (fi17 >> ((255) & 63)) {
                    case -850:
                        fmt.Println(fi21, v1)
                    default:
                        v2 = v2
                    }
                    var v22 bool = true
                    _ = v22
                } else {
                    switch 42 {
                    case 9223372036854775807:
                        v1 = 3.14159
                    case -1:
                        v1 = (0.0 / v5)
                        switch v2 {
                        case "\n\t":
                            switch (^(fi17)) {
                            case 42:
                                v4 -= (v5 * float64((0.1 / 0.1)))
                            }
                        case "a":
                            var v23 int16 = -400
                            _ = v23
                        }
                        v5 = float64(0.1)
                    default:
                        fmt.Println(v4)
                    }
                    if (127 <= -1) {
                        _ = len(v6)
                        v24 := GenS2{F0: (v2 + ""), F1: 42, F2: 186}
                        _ = v24
                    }
                    if true {
                        v4 = v5
                    } else {
                        v5 = ((0.0 / 1.0) * v1)
                        v5 = (v1 - float64(1e10))
                    }
                }
                fmt.Println(v1)
                _ = fi21
            }
            _ = fi17
        }
        var v25 int16 = 0
        v1 += 0.1
        _ = v25
    }
    fmt.Println(v2, v2)
    for fi26 := 0; fi26 < 3; fi26++ {
        v4 = -2.25
        if false {
            for fi27 := 0; fi27 < 3; fi27++ {
                switch (fi27 / ((-459) | 1)) {
                case -865:
                    var v28 uint64 = 1
                    _ = v28
                case 0:
                    v29 := v3[(fi27 | fi27)]
                    _ = v29
                case -298:
                    v5 = 0.1
                }
                var v30 uint32 = 2147483647
                var v31 string = v2
                _ = fi27
                _ = v30
                _ = v31
            }
            var v32 uint8 = 2
            var v33 bool = (false || (!false))
            _ = v32
            _ = v33
        } else {
            if (!false) {
                var v34 int16 = -516
                v5 = 3.14159
                _ = v34
            } else {
                var v35 uint = 0
                var v36 uint32 = 239
                _ = v35
                _ = v36
            }
            for fi37 := 0; fi37 < 11; fi37++ {
                if ((^(fi37)) == (fi37 & fi26)) {
                    _ = len(v3)
                } else {
                    v38 := GenS3{F0: (255 <= 128), F1: 255}
                    v39 := []int{fi37, -538}
                    _ = v38
                    _ = v39
                }
                _ = fi37
            }
            v2 = (v2 + (("a" + v2) + ""))
        }
        v5 -= float64(1.5)
        _ = fi26
    }
    fmt.Println(v2, v5)
    v1 -= (((v4 - v5) + v4) + 1.5)
    var v40 string = "hello"
    if (-32768 > -1) {
        v41 := v3[128]
        _ = v41
    }
    v42 := GenS1{F0: 0}
    fmt.Println(v2, v40)
    if true {
        switch "\n\t" {
        case "\n\t":
            v43 := GenS4{F0: true}
            _ = v43
        case "foo bar":
            var v44 uint64 = 4294967295
            _ = v44
        case "a":
            var v45 string = v40
            _ = v45
        }
        delete(v6, 2147483648)
        for fi46 := 0; fi46 < 3; fi46++ {
            v47 := GenS5{F0: 42, F1: fi46}
            for fi48 := 0; fi48 < 1; fi48++ {
                switch fi46 {
                case 32768:
                    v49 := []int{fi46, fi48, (fi46 % ((fi46) | 1))}
                    delete(v3, (^(fi46)))
                    if false {
                        fmt.Println(v5, fi46)
                    } else {
                        v1 = (v1 - (v5 - ((((v4 + (1.5 / float64(1.5))) * v1) * 0.0) * (v4 - v1))))
                        switch 2147483647 {
                        case 42:
                            var v50 uint32 = 0
                            v51 := GenS6{F0: -118, F1: -280}
                            _ = v50
                            _ = v51
                        case -358:
                            var v52 uint8 = 127
                            _ = v52
                        case 1:
                            if true {
                                fmt.Println(v40, v40)
                                v53 := GenS7{F0: -95}
                                _ = v53
                            }
                        }
                    }
                    _ = v49
                case -2147483648:
                    v54 := []int{-253, (fi48 >> ((255) & 63))}
                    _ = v54
                case 42:
                    if true {
                        v2 = v40
                        var v55 int8 = -1
                        _ = v55
                    } else {
                        v5 += 3.14159
                    }
                }
                _ = fi48
            }
            _ = fi46
            _ = v47
        }
    } else {
        var v56 uint16 = 643
        v3[-107] = 42
        v56 = v56
        _ = v56
    }
    if (true && (0 > 128)) {
        for fi57 := 0; fi57 < 11; fi57++ {
            v58 := GenS7{F0: 834}
            _ = fi57
            _ = v58
        }
        for fi59 := 0; fi59 < 16; fi59++ {
            v4 += v4
            _ = fi59
        }
        if false {
            var v60 float32 = (((3.14159 * 0.1) - 0.1) * ((0.0 - (-2.25 * (-2.25 / 1.0))) + 0.0))
            var v61 float32 = 0.0
            if false {
                v1 += (v1 + ((1.5 / v4) + v1))
                for fi62 := 0; fi62 < 1; fi62++ {
                    v63 := GenS8{F0: "", F1: 255, F2: 0}
                    v64 := GenS3{F0: false, F1: 255}
                    _ = fi62
                    _ = v63
                    _ = v64
                }
            } else {
                fmt.Println(v60, v5)
                v60 = -2.25
                v1 = v5
            }
            _ = v60
            _ = v61
        } else {
            v65 := v3[-2147483648]
            var v66 int32 = 0
            if true {
                _ = len(v3)
            } else {
                switch ((v40 + v40) + "foo bar") {
                case "":
                    delete(v6, (v65 / ((973) | 1)))
                case "foo bar":
                    switch (("\n\t" + "") + v40) {
                    case "":
                        if (v66 <= (v66 / ((-2147483648) | 1))) {
                            if true {
                                v5 = float64(0.1)
                                var v67 int = -1
                                for fi68 := 0; fi68 < 4; fi68++ {
                                    switch ((v2 + (v2 + "foo bar")) + ("a" + ((("hello" + ("" + "foo bar")) + "foo bar") + ((v40 + v40) + ((v40 + (v40 + v2)) + v40))))) {
                                    case "hello":
                                        v1 = (0.0 / (v4 + (v5 * 0.0)))
                                    case "":
                                        if true {
                                            v69 := GenS4{F0: (-128 > -1)}
                                            v70 := GenS6{F0: 32767, F1: fi68}
                                            v71 := GenS5{F0: 9223372036854775807, F1: fi68}
                                            _ = v69
                                            _ = v70
                                            _ = v71
                                        }
                                    case "a":
                                        if false {
                                            if (-2147483648 == v66) {
                                                v2 = v40
                                                v1 += ((v4 * (1e10 / 1.5)) - (v4 + v4))
                                                v40 = "foo bar"
                                            }
                                            delete(v6, (v65 << ((0) & 63)))
                                        }
                                    }
                                    if true {
                                        for fi72 := 0; fi72 < 2; fi72++ {
                                            fmt.Println(v40, v66)
                                            var v73 uint8 = 2
                                            v67 -= fi68
                                            if true {
                                                v74 := []int{fi72, (fi72 << ((2) & 63))}
                                                _ = v74
                                            }
                                            _ = fi72
                                            _ = v73
                                        }
                                        v5 = v1
                                    }
                                    if (0 <= 197) {
                                        for fi75 := 0; fi75 < 3; fi75++ {
                                            v76 := []int{fi68, (fi68 & 670)}
                                            var v77 bool = (v2 > ("" + "foo bar"))
                                            _ = fi75
                                            _ = v76
                                            _ = v77
                                        }
                                        v78 := []int{fi68, (^(v67)), (v67 << ((626) & 63))}
                                        _ = v78
                                    }
                                    _ = fi68
                                }
                                _ = v67
                            } else {
                                for fi79 := 0; fi79 < 5; fi79++ {
                                    var v80 float64 = v1
                                    _ = fi79
                                    _ = v80
                                }
                                if true {
                                    var v81 uint = 255
                                    _ = v81
                                } else {
                                    if true {
                                        v2 = v40
                                        var v82 uint = 18446744073709551615
                                        _ = v82
                                    } else {
                                        var v83 int64 = 0
                                        _ = v83
                                    }
                                    if true {
                                        fmt.Println(v40)
                                        v84 := v3[211]
                                        var v85 int64 = 9223372036854775807
                                        _ = v84
                                        _ = v85
                                    } else {
                                        if false {
                                            switch 2147483647 {
                                            case 32767:
                                                v86 := []int{630}
                                                _ = v86
                                            case 32768:
                                                var v87 int = 127
                                                _ = v87
                                            case 1:
                                                v88 := []int{-310}
                                                _ = v88
                                            }
                                            v89 := []int{220, -2147483648, 1, -2147483648}
                                            for fi90 := 0; fi90 < 15; fi90++ {
                                                _ = len(v89)
                                                fmt.Println(v5, v65)
                                                if (1 >= (fi90 - (fi90 + fi90))) {
                                                    for fi91 := 0; fi91 < 9; fi91++ {
                                                        var v92 int64 = 42
                                                        v93 := GenS9{F0: 1, F1: 128}
                                                        _ = fi91
                                                        _ = v92
                                                        _ = v93
                                                    }
                                                    if len(v89) > 0 {
                                                        v89[0] = (fi90 % (((-(fi90))) | 1))
                                                    }
                                                } else {
                                                    v2 = (("\n\t" + "foo bar") + ((v2 + v2) + (v40 + v2)))
                                                }
                                                _ = fi90
                                            }
                                            _ = v89
                                        } else {
                                            v94 := v6[(v65 ^ (^(v65)))]
                                            if false {
                                                for fi95 := 0; fi95 < 19; fi95++ {
                                                    if false {
                                                        if (!false) {
                                                            var v96 int64 = -1
                                                            _ = v96
                                                        }
                                                        v97 := GenS10{F0: 100}
                                                        _ = v97
                                                    }
                                                    var v98 float32 = 1e10
                                                    _ = fi95
                                                    _ = v98
                                                }
                                            }
                                            var v99 string = ""
                                            _ = v94
                                            _ = v99
                                        }
                                    }
                                    v100 := GenS11{F0: 121}
                                    _ = v100
                                }
                                switch -9223372036854775808 {
                                case 365:
                                    var v101 int64 = 9223372036854775807
                                    _ = v101
                                default:
                                    v40 = ((v40 + v2) + (v2 + v40))
                                }
                            }
                            var v102 string = (v2 + ("a" + v40))
                            _ = v102
                        } else {
                            for fi103 := 0; fi103 < 11; fi103++ {
                                var v104 float32 = 1e10
                                v2 = "hello"
                                if false {
                                    switch 32767 {
                                    case 32767:
                                        fmt.Println(v4, v1)
                                    case 0:
                                        var v105 float64 = (v1 / v4)
                                        _ = v105
                                    case -535:
                                        var v106 int64 = 42
                                        _ = v106
                                    default:
                                        v104 -= v104
                                    }
                                }
                                _ = fi103
                                _ = v104
                            }
                            var v107 int8 = -128
                            var v108 int64 = 42
                            _ = v107
                            _ = v108
                        }
                    case "\n\t":
                        if false {
                            switch -404 {
                            case 128:
                                v1 -= (v4 - v4)
                            case 259:
                                if true {
                                    delete(v3, 127)
                                    fmt.Println(v2)
                                    v2 = v2
                                } else {
                                    var v109 float32 = float32(v4)
                                    v110 := GenS4{F0: (((v4 * float64(v109)) * v4) == (v1 / 1e10))}
                                    _ = v109
                                    _ = v110
                                }
                            case 753:
                                switch ("" + (v2 + "\n\t")) {
                                case "foo bar":
                                    var v111 int32 = v66
                                    _ = v111
                                case "hello":
                                    for fi112 := 0; fi112 < 9; fi112++ {
                                        switch 4294967295 {
                                        case 0:
                                            for fi113 := 0; fi113 < 11; fi113++ {
                                                v114 := []int{(fi112 | (-(fi112))), 42, (fi113 - fi113), 42}
                                                fmt.Println(v65, v2)
                                                _ = fi113
                                                _ = v114
                                            }
                                        case 578:
                                            _ = len(v3)
                                        }
                                        _ = fi112
                                    }
                                case "\n\t":
                                    if false {
                                        if false {
                                            if true {
                                                var v115 float64 = -2.25
                                                var v116 int64 = -1
                                                _ = len(v3)
                                                _ = v115
                                                _ = v116
                                            }
                                            var v117 int = 0
                                            _ = v117
                                        }
                                        var v118 int8 = -1
                                        _ = v118
                                    } else {
                                        var v119 uint32 = (v65 + v65)
                                        delete(v6, v119)
                                        _ = v119
                                    }
                                default:
                                    fmt.Println(v66)
                                }
                            }
                            if false {
                                for fi120 := 0; fi120 < 1; fi120++ {
                                    delete(v6, 0)
                                    _ = fi120
                                }
                                switch "\n\t" {
                                case "\n\t":
                                    v65 += 65535
                                case "a":
                                    for fi121 := 0; fi121 < 19; fi121++ {
                                        var v122 int16 = 32767
                                        _ = fi121
                                        _ = v122
                                    }
                                case "hello":
                                    fmt.Println(v40, v2)
                                default:
                                    var v123 uint64 = 732
                                    _ = v123
                                }
                            } else {
                                if true {
                                    var v124 uint32 = (^(v65))
                                    fmt.Println(v5, v40, v40)
                                    _ = v124
                                }
                                v125 := GenS12{F0: -1}
                                var v126 int16 = 248
                                _ = v125
                                _ = v126
                            }
                        }
                    default:
                        v2 = v40
                    }
                case "hello":
                    fmt.Println(v4, v5, v40)
                default:
                    v127 := GenS13{F0: v2}
                    _ = v127
                }
                fmt.Println(v66)
            }
            _ = v65
            _ = v66
        }
    } else {
        for fi128 := 0; fi128 < 17; fi128++ {
            _ = len(v3)
            var v129 float32 = 1.5
            v130 := []int{fi128}
            _ = fi128
            _ = v129
            _ = v130
        }
    }
    if (!false) {
        var v131 uint8 = 2
        v132 := GenS9{F0: 18446744073709551615, F1: v131}
        _ = v131
        _ = v132
    } else {
        for fi133 := 0; fi133 < 13; fi133++ {
            if true {
                v40 = (v2 + (v2 + ""))
            } else {
                v134 := GenS14{F0: 200}
                v40 = v2
                _ = v134
            }
            _ = fi133
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v42
    _ = v5
    _ = v6
}
