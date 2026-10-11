package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 int16; F2 uint }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 bool; F1 bool }
type GenS7 struct { F0 bool; F1 uint64; F2 uint }
type GenS8 struct { F0 uint8 }
func genFn0(p0 bool, p1 uint8, p2 uint16) uint8 {
    v5 := GenS1{F0: p2}
    fmt.Println(p1)
    switch 9223372036854775807 {
    case 658:
        if p0 {
            for fi6 := 0; fi6 < 14; fi6++ {
                p0 = p0
                _ = fi6
            }
            v7 := []int{32768}
            if len(v7) > 0 {
                v7[0] = -74
            }
            _ = v7
        } else {
            for fi8 := 0; fi8 < 8; fi8++ {
                var v9 uint32 = 2147483647
                var v10 float64 = -2.25
                _ = fi8
                _ = v10
                _ = v9
            }
            if p0 {
                var v11 int = -2147483648
                _ = v11
            } else {
                for fi12 := 0; fi12 < 18; fi12++ {
                    var v13 int16 = -357
                    for fi14 := 0; fi14 < 15; fi14++ {
                        if ((1e10 / 1.0) > (1.5 + ((1e10 / ((1.5 - 0.1) * 1e10)) / 1e10))) {
                            v15 := []int{fi12, fi14, (fi14 << ((0) & 63))}
                            p0 = false
                            _ = v15
                        }
                        v16 := GenS1{F0: (p2 << ((1) & 63))}
                        v17 := GenS1{F0: (p2 ^ 438)}
                        _ = fi14
                        _ = v16
                        _ = v17
                    }
                    v18 := []int{(fi12 ^ fi12), (fi12 / ((fi12) | 1)), fi12}
                    _ = fi12
                    _ = v13
                    _ = v18
                }
                var v19 float64 = 0.1
                switch 290 {
                case -1:
                    for fi20 := 0; fi20 < 7; fi20++ {
                        v21 := []int{fi20, (fi20 / ((32767) | 1)), fi20, fi20}
                        var v22 int8 = 42
                        for fi23 := 0; fi23 < 9; fi23++ {
                            v22 = (-(v22))
                            _ = fi23
                        }
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                case 9223372036854775807:
                    var v24 int32 = -1
                    _ = v24
                }
                _ = v19
            }
            v25 := []int{1, 451, 128, -2147483648}
            _ = v25
        }
    case 106:
        switch "a" {
        case "foo bar":
            switch "hello" {
            case "a":
                p0 = (42 >= 688)
            case "\n\t":
                var v26 float32 = 0.1
                _ = v26
            case "":
                var v27 int64 = 1
                _ = v27
            }
        case "\n\t":
            p2 -= (p2 * (^(p2)))
        case "a":
            if (!(2147483647 == 32768)) {
                for fi28 := 0; fi28 < 7; fi28++ {
                    fmt.Println(p1)
                    _ = fi28
                }
                v29 := GenS2{F0: -1, F1: 42, F2: 109}
                for fi30 := 0; fi30 < 1; fi30++ {
                    if false {
                        for fi31 := 0; fi31 < 3; fi31++ {
                            var v32 int8 = -128
                            _ = fi31
                            _ = v32
                        }
                        for fi33 := 0; fi33 < 11; fi33++ {
                            fmt.Println(fi33, p0, p1)
                            var v34 uint32 = 0
                            p0 = true
                            var v35 int64 = 9223372036854775807
                            _ = fi33
                            _ = v34
                            _ = v35
                        }
                        var v36 float32 = 3.14159
                        _ = v36
                    }
                    _ = fi30
                }
                _ = v29
            } else {
                var v37 int16 = 721
                switch -198 {
                case 42:
                    v38 := GenS1{F0: 69}
                    _ = v38
                case -9223372036854775808:
                    var v39 float64 = -2.25
                    _ = v39
                default:
                    v40 := GenS1{F0: p2}
                    _ = v40
                }
                _ = v37
            }
        }
    case 127:
        p2 += 288
    default:
        switch "hello" {
        case "":
            var v41 float32 = 3.14159
            _ = v41
        case "\n\t":
            p1 = 1
        }
    }
    switch "a" {
    case "foo bar":
        fmt.Println(p1, p1, p2)
    case "hello":
        if (42 < 162) {
            for fi42 := 0; fi42 < 19; fi42++ {
                var v43 int16 = int16(fi42)
                fmt.Println(p1, fi42)
                p1 = (p1 % ((p1) | 1))
                _ = fi42
                _ = v43
            }
            for fi44 := 0; fi44 < 5; fi44++ {
                fmt.Println(fi44, p2)
                p0 = (1 >= 42)
                _ = fi44
            }
        } else {
            var v45 uint64 = 18446744073709551615
            _ = v45
        }
    case "\n\t":
        v46 := GenS1{F0: (p2 & (^(p2)))}
        _ = v46
    }
    v47 := []int{32768, 32768}
    _ = p0
    _ = p1
    _ = p2
    _ = v47
    _ = v5
    return p1
}

func main() {
    v1 := make(map[string]int)
    var v2 int16 = 1000
    var v3 string = ("foo bar" + (("foo bar" + "\n\t") + ("foo bar" + "foo bar")))
    var v4 uint32 = 387
    v1[v3] = 764
    v4 -= (^(v4))
    _ = len(v1)
    for fi48 := 0; fi48 < 8; fi48++ {
        if false {
            if (1 != (v2 & 1)) {
                v49 := GenS2{F0: 9223372036854775807, F1: v2, F2: 42}
                v4 = v4
                var v50 bool = false
                _ = v49
                _ = v50
            } else {
                var v51 int64 = 1
                for fi52 := 0; fi52 < 12; fi52++ {
                    v3 = (v3 + v3)
                    var v53 uint64 = 534
                    var v54 int = fi48
                    _ = fi52
                    _ = v53
                    _ = v54
                }
                _ = v51
            }
            for fi55 := 0; fi55 < 4; fi55++ {
                if true {
                    var v56 float64 = (((-2.25 + (0.1 * ((((0.1 * 1.5) - (0.1 * -2.25)) / 0.1) / 1.5))) / float64((float32(1.5) / ((3.14159 + (1e10 / 1.0)) * 1e10)))) / 3.14159)
                    _ = v56
                }
                v57 := []int{(fi55 / ((32767) | 1)), fi48, fi55}
                _ = fi55
                _ = v57
            }
        } else {
            switch "" {
            case "hello":
                v3 = (v3 + v3)
            case "a":
                var v58 int64 = 545
                _ = v58
            case "foo bar":
                fmt.Println(v2, v4, v2)
            default:
                v3 = ""
            }
            fmt.Println(fi48)
        }
        fmt.Println(fi48, v3)
        v59 := GenS1{F0: 0}
        _ = fi48
        _ = v59
    }
    if (1 <= 127) {
        fmt.Println(v3, v3, v2)
    }
    v3 = v3
    v2 = (v2 + 42)
    if true {
        for fi60 := 0; fi60 < 13; fi60++ {
            v3 = "\n\t"
            _ = fi60
        }
        _ = len(v1)
        v4 = (v4 << ((1) & 63))
    } else {
        if true {
            v4 = (v4 + v4)
            if (false && (32768 > 9223372036854775807)) {
                if (255 > 255) {
                    v2 -= (-(v2))
                    var v61 uint32 = 2147483647
                    _ = v61
                } else {
                    var v62 uint16 = 756
                    for fi63 := 0; fi63 < 18; fi63++ {
                        v64 := GenS3{F0: true}
                        _ = len(v1)
                        if true {
                            v62 += v62
                            if (178 < (v4 / ((41) | 1))) {
                                v65 := GenS1{F0: 0}
                                v2 = (v2 >> ((18446744073709551615) & 63))
                                _ = v65
                            }
                        } else {
                            switch v3 {
                            case "":
                                var v66 int16 = (v2 & -32768)
                                _ = v66
                            case "hello":
                                v67 := GenS4{F0: 42}
                                _ = v67
                            default:
                                v68 := GenS5{F0: 896}
                                _ = v68
                            }
                            fmt.Println(v3, v2)
                        }
                        _ = fi63
                        _ = v64
                    }
                    var v69 float64 = 3.14159
                    _ = v62
                    _ = v69
                }
                fmt.Println(v4, v2)
            }
        } else {
            if true {
                for fi70 := 0; fi70 < 9; fi70++ {
                    v3 = "\n\t"
                    var v71 int = 9223372036854775807
                    _ = fi70
                    _ = v71
                }
            } else {
                if true {
                    var v72 bool = (65535 == 65535)
                    _ = v72
                }
                v73 := GenS4{F0: 0}
                _ = v73
            }
            switch 42 {
            case 480:
                v74 := GenS5{F0: 4294967295}
                _ = v74
            case 914:
                if (true || (42 >= 0)) {
                    switch 9223372036854775807 {
                    case -2147483648:
                        for fi75 := 0; fi75 < 1; fi75++ {
                            if false {
                                var v76 bool = (true && (42 == uint(v4)))
                                for fi77 := 0; fi77 < 8; fi77++ {
                                    v78 := GenS5{F0: 0}
                                    var v79 int32 = 2147483647
                                    _ = fi77
                                    _ = v78
                                    _ = v79
                                }
                                v80 := v1[v3]
                                _ = v76
                                _ = v80
                            }
                            _ = fi75
                        }
                    case 127:
                        for fi81 := 0; fi81 < 19; fi81++ {
                            v82 := GenS4{F0: 565}
                            _ = fi81
                            _ = v82
                        }
                    case 114:
                        fmt.Println(v4, v3, v2)
                    default:
                        for fi83 := 0; fi83 < 18; fi83++ {
                            var v84 float32 = -2.25
                            if false {
                                v84 -= v84
                                v2 -= (v2 | (v2 % ((v2) | 1)))
                            } else {
                                var v85 float32 = ((1e10 + 0.1) * float32(1.5))
                                v4 += v4
                                for fi86 := 0; fi86 < 8; fi86++ {
                                    if true {
                                        var v87 uint8 = 123
                                        v4 = (v4 - 2147483647)
                                        fmt.Println(v4, v2, v3)
                                        _ = v87
                                    } else {
                                        v88 := []int{}
                                        v4 += (v4 << ((42) & 63))
                                        _ = v88
                                    }
                                    _ = fi86
                                }
                                _ = v85
                            }
                            v1[(v3 + v3)] = (fi83 / ((127) | 1))
                            _ = fi83
                            _ = v84
                        }
                    }
                    v4 -= 2147483647
                    if ((0 != 42) && true) {
                        for fi89 := 0; fi89 < 20; fi89++ {
                            v90 := GenS5{F0: 1}
                            _ = fi89
                            _ = v90
                        }
                        fmt.Println(v3, v2, v4)
                        for fi91 := 0; fi91 < 2; fi91++ {
                            v3 = "hello"
                            var v92 uint64 = 18446744073709551615
                            for fi93 := 0; fi93 < 13; fi93++ {
                                v3 = ""
                                switch ("" + v3) {
                                case "a":
                                    var v94 bool = (!false)
                                    _ = v94
                                case "":
                                    var v95 int8 = 100
                                    _ = v95
                                }
                                _ = fi93
                            }
                            _ = fi91
                            _ = v92
                        }
                    } else {
                        v96 := []int{127, 2147483647}
                        if (int(v4) != 4294967295) {
                            var v97 float64 = (-2.25 * -2.25)
                            v97 = (((float64(float32(v97)) * (v97 * (v97 * v97))) - 1.5) + (v97 * (float64(1e10) + -2.25)))
                            v2 = v2
                            _ = v97
                        }
                        if true {
                            v98 := GenS6{F0: (245 > 255), F1: false}
                            _ = v98
                        } else {
                            fmt.Println(v3)
                        }
                        _ = v96
                    }
                }
            case 194:
                v99 := GenS4{F0: 42}
                _ = v99
            default:
                if (249 != uint(v4)) {
                    var v100 uint16 = 715
                    if false {
                        fmt.Println(v2, v100)
                        v101 := GenS1{F0: v100}
                        _ = v101
                    }
                    v4 = v4
                    _ = v100
                } else {
                    var v102 uint64 = 726
                    if ((!((594 <= 2) || false)) && (!(0 >= 255))) {
                        for fi103 := 0; fi103 < 6; fi103++ {
                            if true {
                                var v104 int16 = (v2 + (v2 % ((190) | 1)))
                                var v105 uint = 458
                                _ = v104
                                _ = v105
                            }
                            _ = fi103
                        }
                    }
                    _ = v102
                }
            }
            v106 := []int{0, -296}
            _ = v106
        }
    }
    switch int(v2) {
    case 4294967295:
        for fi107 := 0; fi107 < 11; fi107++ {
            v108 := []int{fi107, fi107, fi107, fi107}
            if true {
                v109 := v1[(v3 + v3)]
                _ = v109
            }
            var v110 bool = (!(true && false))
            _ = fi107
            _ = v108
            _ = v110
        }
    case 0:
        switch v3 {
        case "foo bar":
            fmt.Println(v4, v2)
        case "hello":
            v111 := GenS7{F0: ((1.5 * 1.5) > -2.25), F1: 172, F2: 255}
            _ = v111
        default:
            fmt.Println(v2, v3, v4)
        }
    case 127:
        v1[v3] = -1
    }
    var v112 int32 = 0
    switch 860 {
    case 1:
        fmt.Println(v2, v4)
    case 0:
        if false {
            v113 := GenS4{F0: -9223372036854775808}
            for fi114 := 0; fi114 < 3; fi114++ {
                var v115 bool = false
                _ = fi114
                _ = v115
            }
            _ = v113
        } else {
            for fi116 := 0; fi116 < 11; fi116++ {
                v117 := []int{(fi116 / ((fi116) | 1)), fi116, (fi116 >> ((18446744073709551615) & 63))}
                fmt.Println(v112, fi116)
                _ = fi116
                _ = v117
            }
            fmt.Println(v4)
        }
    case -2147483648:
        var v118 uint32 = v4
        _ = v118
    default:
        var v119 uint8 = 1
        _ = v119
    }
    for fi120 := 0; fi120 < 15; fi120++ {
        var v121 float32 = -2.25
        fmt.Println(v2)
        fmt.Println(v121, v121)
        _ = fi120
        _ = v121
    }
    delete(v1, "hello")
    var v122 uint = 2
    v123 := GenS8{F0: 255}
    var v124 string = v3
    fmt.Println(v112, v4)
    v125 := GenS1{F0: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v112:", v112)
    fmt.Println("v122:", v122)
    fmt.Println("v124:", v124)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v112
    _ = v122
    _ = v123
    _ = v124
    _ = v125
    _ = v2
    _ = v3
    _ = v4
}
