package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 uint32; F2 int64 }
type GenS3 struct { F0 bool; F1 int32; F2 uint }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int32; F1 int32 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint64; F1 uint8; F2 int64 }
type GenS9 struct { F0 uint; F1 uint16 }
type GenS10 struct { F0 int64; F1 uint8; F2 uint32 }
type GenS11 struct { F0 bool; F1 uint32; F2 uint8 }
type GenS12 struct { F0 uint; F1 uint32; F2 string }
func genFn0(p0 uint8) uint16 {
    p0 -= (p0 >> ((18446744073709551615) & 63))
    v6 := GenS1{F0: 32767}
    for fi7 := 0; fi7 < 8; fi7++ {
        p0 += p0
        v8 := GenS1{F0: (fi7 >> ((255) & 63))}
        _ = fi7
        _ = v8
    }
    var v9 int32 = 42
    if false {
        if false {
            fmt.Println(p0)
            for fi10 := 0; fi10 < 20; fi10++ {
                for fi11 := 0; fi11 < 6; fi11++ {
                    v12 := GenS1{F0: (fi11 + (fi10 * 4294967295))}
                    _ = fi11
                    _ = v12
                }
                _ = fi10
            }
            switch "foo bar" {
            case "foo bar":
                fmt.Println(v9, v9)
            case "\n\t":
                var v13 int32 = (^(v9))
                _ = v13
            }
        } else {
            var v14 float32 = 0.0
            for fi15 := 0; fi15 < 1; fi15++ {
                v9 += (v9 - (v9 << ((2) & 63)))
                v16 := GenS2{F0: 255, F1: 42, F2: 0}
                _ = fi15
                _ = v16
            }
            switch -9223372036854775808 {
            case 2147483647:
                for fi17 := 0; fi17 < 8; fi17++ {
                    p0 = (p0 + (p0 * (p0 % ((p0) | 1))))
                    v18 := []int{0, (fi17 / ((fi17) | 1)), fi17}
                    _ = fi17
                    _ = v18
                }
            }
            _ = v14
        }
        switch "" {
        case "foo bar":
            for fi19 := 0; fi19 < 5; fi19++ {
                switch "\n\t" {
                case "a":
                    var v20 int64 = -646
                    _ = v20
                case "":
                    for fi21 := 0; fi21 < 13; fi21++ {
                        var v22 int64 = 9223372036854775807
                        _ = fi21
                        _ = v22
                    }
                }
                for fi23 := 0; fi23 < 19; fi23++ {
                    fmt.Println(p0)
                    fmt.Println(p0, p0, fi23)
                    _ = fi23
                }
                _ = fi19
            }
        default:
            v24 := []int{32768}
            _ = v24
        }
        v9 = (v9 * 385)
    } else {
        var v25 bool = (2 != 0)
        fmt.Println(p0, v25)
        var v26 int32 = v9
        _ = v25
        _ = v26
    }
    _ = p0
    _ = v6
    _ = v9
    return 971
}
func genFn1(p0 uint16, p1 int32) int32 {
    var v27 float64 = 0.0
    v28 := []int{41, -400, 32768, -9223372036854775808}
    _ = p0
    _ = p1
    _ = v27
    _ = v28
    return p1
}

func main() {
    var v1 uint32 = 56
    v2 := make(map[int]string)
    var v3 uint = 4294967295
    var v4 int8 = 5
    v5 := make(map[uint32]uint32)
    var v29 float64 = -2.25
    var v30 uint64 = 4294967295
    switch "a" {
    case "":
        switch 32768 {
        case 520:
            switch 32768 {
            case 127:
                if false {
                    switch "hello" {
                    case "hello":
                        if false {
                            v31 := []int{}
                            for fi32 := 0; fi32 < 5; fi32++ {
                                for fi33 := 0; fi33 < 17; fi33++ {
                                    switch (("foo bar" + "a") + "foo bar") {
                                    case "foo bar":
                                        switch (fi33 | fi32) {
                                        case -1:
                                            if (v29 <= ((1e10 - v29) + float64(1e10))) {
                                                switch (-(fi32)) {
                                                case 4294967295:
                                                    v34 := GenS2{F0: 42, F1: (v1 & (v1 << (((v3 >> ((2) & 63))) & 63))), F2: -9223372036854775808}
                                                    _ = v34
                                                case 532:
                                                    switch fi33 {
                                                    case 32768:
                                                        v35 := GenS2{F0: uint16(v30), F1: 952, F2: 1}
                                                        _ = v35
                                                    default:
                                                        fmt.Println(v29)
                                                    }
                                                case 1:
                                                    var v36 uint32 = v1
                                                    _ = v36
                                                }
                                                v37 := GenS1{F0: 891}
                                                _ = v37
                                            } else {
                                                if true {
                                                    fmt.Println(v3)
                                                    if (128 < (fi32 * -2147483648)) {
                                                        if false {
                                                            if true {
                                                                if (("foo bar" > ("a" + (((("hello" + "a") + "") + "a") + "foo bar"))) && true) {
                                                                    switch (fi32 % ((32767) | 1)) {
                                                                    case -2147483648:
                                                                        switch "a" {
                                                                        case "foo bar":
                                                                            switch "a" {
                                                                            case "a":
                                                                                v38 := GenS3{F0: ((v3 >> ((1) & 63)) >= v3), F1: -2147483648, F2: 18446744073709551615}
                                                                                v29 = 1.5
                                                                                v1 += (v1 % (((v1 << ((910) & 63))) | 1))
                                                                                _ = v38
                                                                            case "hello":
                                                                                v39 := GenS2{F0: 1, F1: v1, F2: 42}
                                                                                _ = v39
                                                                            case "foo bar":
                                                                                var v40 uint = v3
                                                                                _ = v40
                                                                            default:
                                                                                delete(v2, 32767)
                                                                            }
                                                                        default:
                                                                            var v41 float64 = v29
                                                                            _ = v41
                                                                        }
                                                                        if true {
                                                                            _ = len(v2)
                                                                            fmt.Println(v3)
                                                                            var v42 int16 = -581
                                                                            _ = v42
                                                                        }
                                                                    default:
                                                                        v29 = v29
                                                                    }
                                                                    v29 = v29
                                                                    if len(v31) >= 2 {
                                                                        v31 = v31[1:len(v31)]
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        var v43 uint16 = 255
                                                        _ = v43
                                                    } else {
                                                        if true {
                                                            v44 := GenS4{F0: 255}
                                                            var v45 int64 = -304
                                                            _ = v44
                                                            _ = v45
                                                        }
                                                        v3 = (v3 - uint(v30))
                                                    }
                                                }
                                                var v46 uint16 = genFn0(255)
                                                _ = v46
                                            }
                                        case 4294967295:
                                            if false {
                                                fmt.Println(fi33, fi32)
                                            } else {
                                                var v47 uint = (v3 - (v3 & (v3 << ((v3) & 63))))
                                                _ = v47
                                            }
                                        default:
                                            var v48 float64 = float64(float32(v29))
                                            _ = v48
                                        }
                                    case "a":
                                        v5[v1] = (v1 >> (((v3 * (^(v3)))) & 63))
                                    case "":
                                        fmt.Println(fi32, v29)
                                    default:
                                        v49 := GenS5{F0: false}
                                        _ = v49
                                    }
                                    if (false && true) {
                                        v30 += v30
                                    }
                                    _ = fi33
                                }
                                _ = fi32
                            }
                            _ = v31
                        } else {
                            for fi50 := 0; fi50 < 4; fi50++ {
                                v51 := GenS6{F0: 803, F1: 42}
                                _ = fi50
                                _ = v51
                            }
                        }
                    case "foo bar":
                        switch 366 {
                        case 305:
                            v52 := GenS1{F0: 127}
                            _ = v52
                        default:
                            for fi53 := 0; fi53 < 15; fi53++ {
                                var v54 float64 = (v29 - v29)
                                _ = fi53
                                _ = v54
                            }
                        }
                    default:
                        switch "foo bar" {
                        case "foo bar":
                            for fi55 := 0; fi55 < 6; fi55++ {
                                var v56 uint = v3
                                _ = fi55
                                _ = v56
                            }
                        case "hello":
                            v57 := GenS2{F0: genFn0(255), F1: v1, F2: 9223372036854775807}
                            _ = v57
                        default:
                            fmt.Println(v3)
                        }
                    }
                    v58 := []int{42, -307}
                    v1 -= (^(v1))
                    _ = v58
                } else {
                    var v59 int16 = 1
                    if (!(1.5 <= 1e10)) {
                        delete(v2, -252)
                        v30 += (^(v30))
                    } else {
                        for fi60 := 0; fi60 < 7; fi60++ {
                            var v61 uint16 = 166
                            _ = fi60
                            _ = v61
                        }
                        var v62 uint = v3
                        switch "\n\t" {
                        case "":
                            for fi63 := 0; fi63 < 8; fi63++ {
                                var v64 uint32 = 65535
                                var v65 int64 = -9223372036854775808
                                for fi66 := 0; fi66 < 6; fi66++ {
                                    v67 := []int{(fi63 * (fi66 ^ fi63))}
                                    _ = fi66
                                    _ = v67
                                }
                                _ = fi63
                                _ = v64
                                _ = v65
                            }
                        default:
                            v68 := GenS7{F0: (v1 ^ v1)}
                            _ = v68
                        }
                        _ = v62
                    }
                    _ = v59
                }
            default:
                v29 = v29
            }
        case 2147483647:
            v69 := GenS8{F0: 2, F1: 42, F2: 9223372036854775807}
            _ = v69
        }
    }
    if false {
        for fi70 := 0; fi70 < 10; fi70++ {
            fmt.Println(v3, v29, v1)
            if true {
                for fi71 := 0; fi71 < 7; fi71++ {
                    var v72 int16 = -1
                    var v73 uint = (v3 & (^(v3)))
                    v74 := v5[(^(v1))]
                    _ = fi71
                    _ = v72
                    _ = v73
                    _ = v74
                }
                _ = len(v5)
                v3 -= v3
            }
            _ = fi70
        }
    }
    switch "a" {
    case "\n\t":
        if true {
            v1 = v1
            delete(v2, 9223372036854775807)
            if true {
                var v75 uint16 = genFn0(128)
                if ((^(v75)) < genFn0(52)) {
                    var v76 uint32 = v1
                    var v77 uint64 = uint64(v3)
                    v76 -= v76
                    _ = v76
                    _ = v77
                }
                if true {
                    v78 := GenS2{F0: genFn0(1), F1: v1, F2: -570}
                    v4 = v4
                    if ((v29 <= v29) && false) {
                        _ = len(v5)
                        var v79 uint64 = (v30 / (((^(v30))) | 1))
                        _ = v79
                    } else {
                        var v80 int = 460
                        if (!false) {
                            v81 := GenS9{F0: 65535, F1: uint16(v30)}
                            for fi82 := 0; fi82 < 4; fi82++ {
                                v4 = v4
                                v2[(fi82 / (((v80 >> (((v3 * 1)) & 63))) | 1))] = "foo bar"
                                for fi83 := 0; fi83 < 16; fi83++ {
                                    v84 := v5[v1]
                                    v29 -= v29
                                    _ = fi83
                                    _ = v84
                                }
                                _ = fi82
                            }
                            var v85 bool = false
                            _ = v81
                            _ = v85
                        }
                        for fi86 := 0; fi86 < 13; fi86++ {
                            v3 = v3
                            var v87 float64 = v29
                            _ = fi86
                            _ = v87
                        }
                        _ = v80
                    }
                    _ = v78
                } else {
                    v88 := GenS10{F0: -224, F1: 138, F2: (v1 | (v1 ^ (^(v1))))}
                    for fi89 := 0; fi89 < 8; fi89++ {
                        fmt.Println(v3)
                        _ = fi89
                    }
                    if (true && (false || true)) {
                        v90 := []int{128}
                        v3 = (v3 / ((515) | 1))
                        _ = v90
                    } else {
                        switch "a" {
                        case "\n\t":
                            for fi91 := 0; fi91 < 8; fi91++ {
                                v92 := []int{fi91, (-(fi91))}
                                var v93 uint = (v3 | (v3 * 18446744073709551615))
                                if false {
                                    v29 += (0.1 - v29)
                                }
                                _ = fi91
                                _ = v92
                                _ = v93
                            }
                        }
                        if (true || ((-(v4)) < -128)) {
                            v94 := GenS11{F0: (42 != -226), F1: (v1 & uint32(v3)), F2: 2}
                            switch -347 {
                            case 1:
                                for fi95 := 0; fi95 < 17; fi95++ {
                                    v3 = v3
                                    v1 += 2147483647
                                    _ = fi95
                                }
                            case 32768:
                                var v96 int = -194
                                _ = v96
                            case -808:
                                var v97 uint32 = (v1 | v1)
                                _ = v97
                            }
                            var v98 uint32 = v1
                            _ = v94
                            _ = v98
                        }
                    }
                    _ = v88
                }
                _ = v75
            } else {
                v99 := GenS12{F0: (v3 - v3), F1: 42, F2: "foo bar"}
                var v100 int16 = -32768
                fmt.Println(v3, v100, v3)
                _ = v100
                _ = v99
            }
        }
    case "hello":
        v29 += ((((1e10 * (v29 * v29)) + (1e10 * -2.25)) * (1e10 + float64(1.5))) / v29)
    case "a":
        switch -2147483648 {
        case -997:
            v101 := []int{}
            _ = v101
        case -534:
            if true {
                v30 += v30
                v3 = v3
                if false {
                    delete(v2, 32768)
                    for fi102 := 0; fi102 < 14; fi102++ {
                        for fi103 := 0; fi103 < 11; fi103++ {
                            fmt.Println(v1)
                            var v104 uint64 = v30
                            var v105 uint32 = 759
                            _ = fi103
                            _ = v104
                            _ = v105
                        }
                        if true {
                            v4 = v4
                            v3 -= v3
                        } else {
                            var v106 int32 = -2147483648
                            v107 := []int{128, fi102}
                            v30 = (v30 & v30)
                            _ = v106
                            _ = v107
                        }
                        _ = fi102
                    }
                    for fi108 := 0; fi108 < 3; fi108++ {
                        var v109 int8 = 1
                        _ = fi108
                        _ = v109
                    }
                }
            } else {
                v110 := GenS4{F0: (v3 + v3)}
                fmt.Println(v4, v3, v1)
                _ = v110
            }
        }
    }
    v111 := GenS10{F0: -673, F1: 221, F2: (v1 & 2147483648)}
    v3 += v3
    v4 = v4
    v1 = v1
    var v112 int16 = -32768
    v113 := []int{493, int(v112), 9223372036854775807, -2147483648}
    var v114 uint64 = v30
    if (9223372036854775807 < -9223372036854775808) {
        if len(v113) >= 2 {
            v113 = v113[1:len(v113)]
        }
        var v115 uint16 = uint16(v1)
        _ = v115
    }
    var v116 float64 = v29
    var v117 int64 = int64(v112)
    fmt.Println("v1:", v1)
    fmt.Println("v112:", v112)
    fmt.Println("slicelen v113:", len(v113), v113)
    fmt.Println("v114:", v114)
    fmt.Println("v116:", v116)
    fmt.Println("v117:", v117)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v111
    _ = v112
    _ = v113
    _ = v114
    _ = v116
    _ = v117
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
