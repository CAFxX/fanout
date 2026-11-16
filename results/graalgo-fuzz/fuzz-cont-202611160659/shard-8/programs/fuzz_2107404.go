package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16; F1 int64; F2 int32 }
type GenS4 struct { F0 uint16; F1 bool }
type GenS5 struct { F0 int32; F1 bool }
type GenS6 struct { F0 int32; F1 int8; F2 int16 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint; F1 bool; F2 uint16 }
type GenS9 struct { F0 int64; F1 uint32; F2 uint }
type GenS10 struct { F0 bool; F1 int; F2 int16 }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 int8; F1 int8; F2 int16 }
func genFn0() uint {
    var v6 int16 = 802
    var v7 int16 = 1
    var v8 uint = 18446744073709551615
    _ = v6
    _ = v7
    _ = v8
    return v8
}

func main() {
    var v1 uint16 = 255
    var v2 float32 = 0.1
    v3 := make(map[int]uint32)
    var v4 string = "foo bar"
    var v5 float32 = (-2.25 * v2)
    v9 := GenS1{F0: 1}
    var v10 int32 = 283
    var v11 int16 = 83
    _ = len(v3)
    for fi12 := 0; fi12 < 19; fi12++ {
        var v13 float32 = 1.5
        _ = fi12
        _ = v13
    }
    switch ("a" + v4) {
    case "foo bar":
        for fi14 := 0; fi14 < 8; fi14++ {
            var v15 uint16 = 2
            var v16 uint8 = 255
            _ = fi14
            _ = v15
            _ = v16
        }
    case "hello":
        switch v4 {
        case "foo bar":
            switch (v4 + "") {
            case "a":
                v5 = v2
            case "":
                var v17 bool = false
                _ = v17
            }
        case "\n\t":
            for fi18 := 0; fi18 < 17; fi18++ {
                v10 += v10
                v10 += (v10 & (v10 ^ (v10 * (v10 | v10))))
                if true {
                    for fi19 := 0; fi19 < 5; fi19++ {
                        v20 := GenS2{F0: fi19}
                        _ = fi19
                        _ = v20
                    }
                    var v21 float32 = (v2 * v5)
                    if ((!(42 < 0)) || (!(1 > -1))) {
                        switch "hello" {
                        case "\n\t":
                            v22 := GenS1{F0: 1}
                            _ = v22
                        case "":
                            v23 := []int{fi18, (fi18 ^ fi18), fi18}
                            _ = v23
                        case "hello":
                            v24 := []int{fi18, -9223372036854775808}
                            _ = v24
                        }
                        if true {
                            v11 = v11
                            switch -694 {
                            case -1:
                                var v25 int32 = (v10 | -1)
                                _ = v25
                            case 981:
                                if ((v11 + v11) < (v11 / ((v11) | 1))) {
                                    v10 -= v10
                                }
                            case 15:
                                if false {
                                    v21 -= ((v21 * v2) + -2.25)
                                    fmt.Println(fi18)
                                    v26 := GenS3{F0: v1, F1: 461, F2: (v10 % ((v10) | 1))}
                                    _ = v26
                                } else {
                                    v27 := []int{fi18, (fi18 / ((fi18) | 1)), 71, (fi18 * fi18)}
                                    v28 := GenS2{F0: (fi18 + 2147483647)}
                                    _ = v27
                                    _ = v28
                                }
                            }
                        } else {
                            v29 := GenS2{F0: (fi18 / (((fi18 * 1)) | 1))}
                            switch v4 {
                            case "foo bar":
                                var v30 float32 = -2.25
                                _ = v30
                            }
                            _ = v29
                        }
                    } else {
                        for fi31 := 0; fi31 < 15; fi31++ {
                            v1 -= (^(v1))
                            if (true && ((0.0 + (float32(-2.25) / (1.5 * float32(-2.25)))) == v21)) {
                                if (!true) {
                                    var v32 int32 = -2147483648
                                    var v33 int = fi18
                                    _ = v32
                                    _ = v33
                                } else {
                                    var v34 int32 = v10
                                    _ = v34
                                }
                                var v35 uint = 65535
                                _ = v35
                            }
                            _ = fi31
                        }
                        if false {
                            for fi36 := 0; fi36 < 11; fi36++ {
                                var v37 uint32 = 2
                                fmt.Println(v4, fi18)
                                v4 = v4
                                _ = fi36
                                _ = v37
                            }
                            for fi38 := 0; fi38 < 2; fi38++ {
                                v2 = v2
                                _ = fi38
                            }
                            for fi39 := 0; fi39 < 5; fi39++ {
                                if false {
                                    var v40 int16 = (-(v11))
                                    switch fi18 {
                                    case 2147483647:
                                        var v41 int64 = -762
                                        _ = v41
                                    case 1:
                                        if (!(127 > 128)) {
                                            var v42 int = (^(fi39))
                                            _ = v42
                                        } else {
                                            if (fi39 == (fi18 & (fi39 >> ((4294967295) & 63)))) {
                                                fmt.Println(fi18, v10, v5)
                                                var v43 float64 = 0.1
                                                v44 := GenS4{F0: (v1 >> ((42) & 63)), F1: (!true)}
                                                switch v4 {
                                                case "foo bar":
                                                    if ((fi39 | fi18) > fi39) {
                                                        v1 = (v1 << ((42) & 63))
                                                    } else {
                                                        v21 = (v2 * 0.1)
                                                        v43 -= (3.14159 - (v43 / (float64(v21) * (v43 / 3.14159))))
                                                        v45 := GenS2{F0: fi39}
                                                        if (v43 != v43) {
                                                            var v46 int64 = 9223372036854775807
                                                            _ = len(v3)
                                                            var v47 int = int(v11)
                                                            var v48 uint32 = 1
                                                            _ = v46
                                                            _ = v47
                                                            _ = v48
                                                        } else {
                                                            var v49 int32 = 2
                                                            fmt.Println(v11, fi39, v5)
                                                            _ = v49
                                                        }
                                                        _ = v45
                                                    }
                                                default:
                                                    var v50 int8 = -128
                                                    _ = v50
                                                }
                                                switch v4 {
                                                case "":
                                                    _ = len(v3)
                                                case "\n\t":
                                                    switch (fi18 - fi18) {
                                                    case 319:
                                                        var v51 uint32 = 881
                                                        _ = v51
                                                    }
                                                case "hello":
                                                    if false {
                                                        v43 += float64(-2.25)
                                                        if false {
                                                            switch fi18 {
                                                            case 1:
                                                                v4 = v4
                                                            case 9223372036854775807:
                                                                fmt.Println(fi18, v11, v11)
                                                            case 32768:
                                                                v52 := GenS1{F0: 65535}
                                                                if (!(((-2.25 >= v43) && ((!false) && (uint32(v1) <= 1))) || true)) {
                                                                    if false {
                                                                        var v53 float64 = v43
                                                                        v1 -= 0
                                                                        _ = v53
                                                                    } else {
                                                                        v54 := GenS5{F0: (v10 | -1), F1: (false && (4294967295 > 4294967295))}
                                                                        _ = v54
                                                                    }
                                                                }
                                                                _ = v52
                                                            }
                                                        } else {
                                                            v10 = (v10 - 251)
                                                            if true {
                                                                v10 = (v10 >> ((genFn0()) & 63))
                                                            }
                                                        }
                                                    }
                                                default:
                                                    var v55 uint32 = 466
                                                    _ = v55
                                                }
                                                var v56 bool = false
                                                if ((552 < 18446744073709551615) && v56) {
                                                    fmt.Println(v10, v4, v5)
                                                    var v57 uint16 = 255
                                                    _ = v57
                                                } else {
                                                    if (((!v56) && true) || false) {
                                                        var v58 int64 = 1
                                                        _ = v58
                                                    }
                                                    var v59 int = fi39
                                                    fmt.Println(v2, v59, v11)
                                                    v60 := GenS6{F0: 125, F1: 100, F2: -32768}
                                                    fmt.Println(v40)
                                                    v61 := []int{-1, (fi39 / (((v59 | (v59 & -591))) | 1)), fi39, (v59 - v59)}
                                                    _ = v59
                                                    _ = v60
                                                    _ = v61
                                                }
                                                v11 = v40
                                                _ = v43
                                                _ = v44
                                                _ = v56
                                            }
                                            v40 = (v40 % (((v11 / (((v11 << ((0) & 63))) | 1))) | 1))
                                        }
                                    }
                                    v62 := GenS7{F0: v1}
                                    _ = v40
                                    _ = v62
                                } else {
                                    if (v1 > v1) {
                                        v1 -= v1
                                        v63 := GenS6{F0: -2147483648, F1: -1, F2: 1000}
                                        _ = v63
                                    }
                                }
                                _ = fi39
                            }
                        }
                    }
                    _ = v21
                }
                _ = fi18
            }
        default:
            switch ("hello" + v4) {
            case "\n\t":
                v5 = v2
            case "":
                for fi64 := 0; fi64 < 19; fi64++ {
                    var v65 uint64 = 337
                    _ = fi64
                    _ = v65
                }
            case "a":
                switch ("a" + v4) {
                case "":
                    v2 = v2
                case "hello":
                    for fi66 := 0; fi66 < 9; fi66++ {
                        if (2147483647 < 0) {
                            v67 := GenS8{F0: 2, F1: ((-(v10)) != v10), F2: v1}
                            fmt.Println(v2)
                            var v68 bool = (185 == 4294967295)
                            _ = v67
                            _ = v68
                        } else {
                            var v69 int32 = (v10 * (v10 | v10))
                            fmt.Println(v69, v1)
                            _ = v69
                        }
                        _ = fi66
                    }
                }
            default:
                if ((v2 / (3.14159 / v5)) > v2) {
                    for fi70 := 0; fi70 < 13; fi70++ {
                        if true {
                            v4 = "foo bar"
                            var v71 int8 = 127
                            _ = v71
                        } else {
                            for fi72 := 0; fi72 < 18; fi72++ {
                                v1 = v1
                                switch (fi72 * (fi70 >> ((genFn0()) & 63))) {
                                case 123:
                                    var v73 uint16 = v1
                                    _ = v73
                                }
                                _ = fi72
                            }
                        }
                        if (0 >= 2) {
                            fmt.Println(v10, v5, v1)
                            for fi74 := 0; fi74 < 17; fi74++ {
                                var v75 int8 = 100
                                v4 = v4
                                _ = fi74
                                _ = v75
                            }
                        } else {
                            v2 -= 1.5
                            switch "foo bar" {
                            case "hello":
                                v76 := v3[-2147483648]
                                _ = v76
                            default:
                                switch ("hello" + v4) {
                                case "a":
                                    var v77 uint64 = 1
                                    _ = v77
                                }
                            }
                            v78 := []int{2147483647, -1, (fi70 & fi70)}
                            _ = v78
                        }
                        _ = fi70
                    }
                    v79 := GenS5{F0: 794, F1: true}
                    _ = v79
                }
            }
        }
    default:
        if true {
            v5 = 0.0
        }
    }
    if ((v11 << ((genFn0()) & 63)) <= v11) {
        switch "foo bar" {
        case "hello":
            if false {
                if true {
                    if (!(42 > 135)) {
                        for fi80 := 0; fi80 < 4; fi80++ {
                            v11 = v11
                            fmt.Println(v2)
                            v5 -= v2
                            _ = fi80
                        }
                    }
                } else {
                    v1 += v1
                    v81 := []int{976, -546, 4294967295}
                    switch "" {
                    case "foo bar":
                        var v82 int16 = (v11 & -987)
                        _ = v82
                    case "a":
                        v83 := v3[-179]
                        _ = v83
                    }
                    _ = v81
                }
            } else {
                v84 := GenS9{F0: -1, F1: 1, F2: 908}
                var v85 int16 = v11
                for fi86 := 0; fi86 < 10; fi86++ {
                    for fi87 := 0; fi87 < 4; fi87++ {
                        if (true && false) {
                            v2 = v2
                            var v88 bool = (genFn0() >= 782)
                            _ = v88
                        }
                        v89 := []int{4294967295}
                        if ((("\n\t" < "a") && false) && false) {
                            if false {
                                fmt.Println(fi86, v10, fi86)
                            } else {
                                if (((3.14159 * 1e10) / v2) > v2) {
                                    v90 := GenS9{F0: 333, F1: 2147483647, F2: 65535}
                                    fmt.Println(v2, v2, v10)
                                    if true {
                                        fmt.Println(v1, fi87)
                                        v91 := GenS10{F0: true, F1: (fi87 >> ((65535) & 63)), F2: v11}
                                        _ = v91
                                    }
                                    _ = v90
                                } else {
                                    v92 := v3[(-(fi86))]
                                    var v93 int32 = 764
                                    v93 = v93
                                    _ = v92
                                    _ = v93
                                }
                                var v94 uint = 647
                                if (v11 < 42) {
                                    switch fi86 {
                                    case -9223372036854775808:
                                        delete(v3, (fi86 << ((42) & 63)))
                                    }
                                    var v95 uint8 = 223
                                    _ = v95
                                } else {
                                    v96 := GenS10{F0: false, F1: 1, F2: (-(v11))}
                                    v11 = 32767
                                    if false {
                                        v97 := GenS11{F0: 1}
                                        _ = v97
                                    }
                                    _ = v96
                                }
                                _ = v94
                            }
                            if true {
                                switch "hello" {
                                case "":
                                    var v98 string = "foo bar"
                                    _ = v98
                                case "a":
                                    var v99 uint32 = 2
                                    _ = v99
                                }
                            } else {
                                if len(v89) >= 2 {
                                    v89 = v89[1:len(v89)]
                                }
                                fmt.Println(v85, v5, v10)
                            }
                            if false {
                                var v100 int64 = -147
                                _ = v100
                            }
                        } else {
                            v10 -= 273
                        }
                        _ = fi87
                        _ = v89
                    }
                    _ = fi86
                }
                _ = v84
                _ = v85
            }
        case "foo bar":
            if (42 == 127) {
                if (1 <= 0) {
                    v101 := GenS2{F0: -431}
                    _ = v101
                } else {
                    for fi102 := 0; fi102 < 20; fi102++ {
                        v11 += (^(v11))
                        var v103 float32 = ((v2 / (0.0 + v2)) - 1e10)
                        v104 := []int{(^(fi102)), -20, 127}
                        _ = fi102
                        _ = v103
                        _ = v104
                    }
                    v11 = v11
                }
            }
        }
        switch v4 {
        case "a":
            for fi105 := 0; fi105 < 8; fi105++ {
                v106 := GenS6{F0: -562, F1: 1, F2: (-(v11))}
                switch fi105 {
                case -1:
                    fmt.Println(fi105, v5)
                case 32767:
                    v107 := GenS9{F0: 0, F1: 2147483648, F2: genFn0()}
                    _ = v107
                default:
                    var v108 bool = false
                    _ = v108
                }
                var v109 int32 = v10
                _ = fi105
                _ = v106
                _ = v109
            }
        case "":
            v110 := []int{-9223372036854775808, 42}
            _ = v110
        }
    }
    var v111 bool = true
    fmt.Println(v5)
    v11 += v11
    v2 += (((v2 / float32((((1.5 + (float64(float32((0.1 / ((3.14159 * 1e10) - 0.1)))) + 1e10)) / ((1e10 - 1.5) + (-2.25 / 0.1))) * (0.0 + ((1.5 + 3.14159) + float64((v5 / v2))))))) - v2) * 3.14159)
    v4 = ""
    v112 := GenS7{F0: v1}
    v4 = ((v4 + "") + "foo bar")
    for fi113 := 0; fi113 < 7; fi113++ {
        if v111 {
            v4 = v4
            v10 = v10
        }
        _ = fi113
    }
    for fi114 := 0; fi114 < 18; fi114++ {
        for fi115 := 0; fi115 < 5; fi115++ {
            v116 := []int{(fi114 * (fi115 >> ((0) & 63))), fi115}
            _ = fi115
            _ = v116
        }
        _ = fi114
    }
    if false {
        for fi117 := 0; fi117 < 2; fi117++ {
            v1 = 32767
            _ = fi117
        }
    }
    v118 := GenS12{F0: 100, F1: 0, F2: v11}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v111:", v111)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v111
    _ = v112
    _ = v118
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
