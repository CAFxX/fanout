package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64; F1 int64; F2 uint }
type GenS3 struct { F0 int }
type GenS4 struct { F0 bool; F1 int64 }
type GenS5 struct { F0 int32; F1 int16 }
type GenS6 struct { F0 uint32; F1 uint }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int64; F1 int8 }
type GenS9 struct { F0 int16; F1 int16 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 int64; F1 int; F2 int64 }
type GenS13 struct { F0 int8; F1 uint8; F2 int64 }
func genFn0(p0 uint64, p1 bool) int {
    if p1 {
        var v4 int8 = 49
        var v5 uint8 = 1
        var v6 uint16 = 32767
        _ = v4
        _ = v5
        _ = v6
    } else {
        if false {
            var v7 string = "foo bar"
            _ = v7
        }
        switch "" {
        case "\n\t":
            if ((p0 >> ((65535) & 63)) < 0) {
                p0 -= p0
            }
        default:
            fmt.Println(p1)
        }
    }
    var v8 int8 = -1
    _ = p0
    _ = p1
    _ = v8
    return 2147483647
}
func genFn1(p0 int32, p1 uint16) bool {
    switch "hello" {
    case "a":
        for fi9 := 0; fi9 < 2; fi9++ {
            var v10 uint16 = (p1 ^ (p1 << ((255) & 63)))
            var v11 int16 = 0
            _ = fi9
            _ = v10
            _ = v11
        }
    case "hello":
        fmt.Println(p1, p0)
    }
    p0 = p0
    var v12 int16 = 1
    var v13 int8 = 100
    v14 := GenS1{F0: p0}
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    _ = v14
    return ((!false) && ((p1 << ((4294967295) & 63)) == (p1 >> ((173) & 63))))
}

func main() {
    var v1 int = 907
    var v2 int = (v1 >> ((930) & 63))
    var v3 uint32 = 42
    v2 -= (v2 * v1)
    if false {
        if (!(1 != 9223372036854775807)) {
            var v15 int = (v2 / ((v2) | 1))
            v3 = v3
            v16 := []int{(^(v2)), v1, genFn0(598, (-667 >= 328))}
            _ = v15
            _ = v16
        }
        for fi17 := 0; fi17 < 12; fi17++ {
            var v18 float64 = 0.1
            _ = fi17
            _ = v18
        }
        for fi19 := 0; fi19 < 15; fi19++ {
            v20 := []int{(v2 >> ((786) & 63)), 32767, (-(fi19))}
            fmt.Println(v3, v1, v1)
            switch v1 {
            case -35:
                _ = len(v20)
            case 32768:
                for fi21 := 0; fi21 < 9; fi21++ {
                    v1 = fi19
                    _ = fi21
                }
            default:
                if (2 > 2) {
                    for fi22 := 0; fi22 < 10; fi22++ {
                        v23 := GenS2{F0: 42, F1: -9223372036854775808, F2: 959}
                        fmt.Println(v3, v1, v1)
                        _ = fi22
                        _ = v23
                    }
                } else {
                    for fi24 := 0; fi24 < 10; fi24++ {
                        if len(v20) > 0 {
                            v20[0] = fi19
                        }
                        _ = fi24
                    }
                }
            }
            _ = fi19
            _ = v20
        }
    }
    v3 += v3
    if (((3.14159 * 1.5) + (3.14159 - (0.0 - 1e10))) >= 3.14159) {
        var v25 uint16 = 606
        v26 := []int{}
        _ = v25
        _ = v26
    } else {
        switch v1 {
        case -473:
            switch (-(v1)) {
            case 127:
                fmt.Println(v1, v2, v2)
            case -9223372036854775808:
                v3 += v3
            }
        case 2147483647:
            if false {
                fmt.Println(v3, v1, v1)
                for fi27 := 0; fi27 < 4; fi27++ {
                    if (4294967295 <= 204) {
                        var v28 float64 = (1e10 + (3.14159 + (3.14159 * (float64(1.5) / (((3.14159 - 0.1) - (1e10 / 0.1)) + -2.25)))))
                        for fi29 := 0; fi29 < 19; fi29++ {
                            v30 := GenS3{F0: (fi29 - 14)}
                            v28 += 3.14159
                            switch "hello" {
                            case "hello":
                                v1 += v1
                            case "":
                                v31 := GenS4{F0: (539 < 0), F1: 9223372036854775807}
                                _ = v31
                            case "\n\t":
                                v32 := []int{(v1 - (fi29 % (((fi29 >> ((978) & 63))) | 1))), v1}
                                _ = v32
                            default:
                                var v33 uint32 = (v3 << ((65535) & 63))
                                _ = v33
                            }
                            _ = fi29
                            _ = v30
                        }
                        var v34 string = ("" + "hello")
                        _ = v28
                        _ = v34
                    } else {
                        if false {
                            var v35 int8 = 42
                            _ = v35
                        } else {
                            v1 = v2
                            for fi36 := 0; fi36 < 18; fi36++ {
                                var v37 bool = (true && (!((fi27 / ((fi27) | 1)) <= fi36)))
                                fmt.Println(v37)
                                if (!v37) {
                                    var v38 uint8 = 207
                                    switch 32768 {
                                    case -222:
                                        var v39 bool = v37
                                        _ = v39
                                    default:
                                        fmt.Println(v1, fi27, v2)
                                    }
                                    v37 = v37
                                    _ = v38
                                } else {
                                    if ((fi36 >= (v1 ^ (v1 | (v2 << ((601) & 63))))) && v37) {
                                        switch "hello" {
                                        case "\n\t":
                                            if v37 {
                                                switch v2 {
                                                case -739:
                                                    fmt.Println(fi27, fi36)
                                                case 324:
                                                    v3 -= v3
                                                }
                                                v37 = (0 != 1)
                                            }
                                        case "hello":
                                            var v40 bool = v37
                                            _ = v40
                                        default:
                                            var v41 bool = v37
                                            switch (v1 * fi27) {
                                            case -977:
                                                switch "hello" {
                                                case "a":
                                                    var v42 uint64 = 0
                                                    _ = v42
                                                case "hello":
                                                    fmt.Println(fi27, v37, fi27)
                                                }
                                                if v37 {
                                                    var v43 uint8 = 2
                                                    v44 := GenS3{F0: (v2 << ((515) & 63))}
                                                    var v45 bool = v41
                                                    _ = v43
                                                    _ = v44
                                                    _ = v45
                                                }
                                                var v46 int = genFn0(550, (!true))
                                                _ = v46
                                            case 2147483647:
                                                v37 = (!false)
                                                var v47 uint = 42
                                                v48 := GenS3{F0: fi27}
                                                _ = v47
                                                _ = v48
                                            case 391:
                                                v2 += -618
                                            default:
                                                fmt.Println(v2)
                                            }
                                            _ = v41
                                        }
                                        v1 -= (^(v1))
                                        v49 := []int{42, v1}
                                        _ = v49
                                    }
                                    v50 := GenS4{F0: false, F1: 0}
                                    _ = v50
                                }
                                if (v37 && false) {
                                    if v37 {
                                        var v51 uint64 = 0
                                        _ = v51
                                    }
                                    var v52 uint64 = 82
                                    v53 := []int{fi27, fi36, (fi27 & v2), genFn0((^(v52)), v37)}
                                    _ = v52
                                    _ = v53
                                } else {
                                    if (-1 == 237) {
                                        if true {
                                            v54 := []int{(^(fi36)), v2}
                                            v54 = append(v54, -9223372036854775808)
                                            _ = v54
                                        }
                                    } else {
                                        var v55 uint64 = 42
                                        _ = v55
                                    }
                                    switch -966 {
                                    case -1:
                                        if v37 {
                                            var v56 uint8 = 47
                                            v57 := GenS3{F0: genFn0(1, v37)}
                                            switch "hello" {
                                            case "\n\t":
                                                v3 = v3
                                            case "hello":
                                                if (("" + ("\n\t" + (("" + (("" + ("a" + ((((("\n\t" + "foo bar") + ("\n\t" + ("foo bar" + ""))) + ("\n\t" + "\n\t")) + "\n\t") + "\n\t"))) + ("" + "foo bar"))) + ("" + "foo bar")))) >= ("foo bar" + ("" + "\n\t"))) {
                                                    switch ("hello" + "") {
                                                    case "\n\t":
                                                        switch v1 {
                                                        case -828:
                                                            v58 := GenS5{F0: 2147483647, F1: 735}
                                                            _ = v58
                                                        case 608:
                                                            var v59 uint8 = v56
                                                            _ = v59
                                                        case 127:
                                                            v2 = (v1 + (v2 - (v2 & 4294967295)))
                                                        }
                                                    case "":
                                                        switch 50 {
                                                        case 32767:
                                                            var v60 uint16 = 57
                                                            _ = v60
                                                        case 994:
                                                            v61 := GenS5{F0: -874, F1: -620}
                                                            _ = v61
                                                        case 128:
                                                            var v62 float64 = (1.5 - 3.14159)
                                                            _ = v62
                                                        }
                                                    case "foo bar":
                                                        if true {
                                                            switch "" {
                                                            case "":
                                                                v63 := GenS5{F0: 42, F1: -634}
                                                                _ = v63
                                                            case "\n\t":
                                                                v1 = (fi27 % (((-(v2))) | 1))
                                                            case "hello":
                                                                if v37 {
                                                                    fmt.Println(v56)
                                                                    v64 := []int{v1, 128}
                                                                    _ = v64
                                                                } else {
                                                                    v1 = (v2 % ((651) | 1))
                                                                    if (9223372036854775807 < 373) {
                                                                        if (((false || (v37 || v37)) && true) && (42 <= 42)) {
                                                                            if v37 {
                                                                                v65 := GenS6{F0: (v3 % (((v3 << ((1) & 63))) | 1)), F1: 42}
                                                                                _ = v65
                                                                            }
                                                                            v66 := []int{128, int(v3)}
                                                                            _ = v66
                                                                        } else {
                                                                            if v37 {
                                                                                var v67 uint16 = 204
                                                                                _ = v67
                                                                            }
                                                                        }
                                                                        var v68 uint64 = 1
                                                                        v68 = v68
                                                                        _ = v68
                                                                    } else {
                                                                        var v69 bool = (!true)
                                                                        var v70 uint8 = (^(v56))
                                                                        _ = v69
                                                                        _ = v70
                                                                    }
                                                                }
                                                            }
                                                            fmt.Println(v37, v1, v1)
                                                            fmt.Println(v2, v3)
                                                        } else {
                                                            v1 += (fi36 | (-(fi36)))
                                                            if (18446744073709551615 > 0) {
                                                                v71 := []int{(fi36 | -2147483648), v1, genFn0(2, (false || v37)), v2}
                                                                switch (fi36 * genFn0(1, (1 != -32768))) {
                                                                case 128:
                                                                    v37 = v37
                                                                }
                                                                v1 -= genFn0(18446744073709551615, (v37 && v37))
                                                                var v72 uint = 4294967295
                                                                _ = v71
                                                                _ = v72
                                                            } else {
                                                                v73 := []int{30, genFn0(18446744073709551615, (v37 || true))}
                                                                switch genFn0(1, false) {
                                                                case -2147483648:
                                                                    var v74 uint = 65535
                                                                    _ = v74
                                                                default:
                                                                    _ = len(v73)
                                                                }
                                                                _ = v73
                                                            }
                                                            var v75 int8 = 0
                                                            _ = v75
                                                        }
                                                    }
                                                }
                                            default:
                                                fmt.Println(v1, fi27, v56)
                                                if true {
                                                    var v76 uint64 = 389
                                                    switch fi27 {
                                                    case -651:
                                                        v77 := GenS6{F0: v3, F1: 65535}
                                                        _ = v77
                                                    }
                                                    if v37 {
                                                        if (!v37) {
                                                            switch ((("foo bar" + "foo bar") + "hello") + "hello") {
                                                            case "hello":
                                                                v78 := GenS7{F0: 42}
                                                                _ = v78
                                                            case "foo bar":
                                                                fmt.Println(v1, v2)
                                                            }
                                                            var v79 uint = 255
                                                            _ = v79
                                                        }
                                                        var v80 float64 = (3.14159 - 1e10)
                                                        _ = v80
                                                    } else {
                                                        if (917 != -120) {
                                                            if ((v76 ^ (v76 + (v76 >> ((4294967295) & 63)))) <= (v76 ^ 18446744073709551615)) {
                                                                if v37 {
                                                                    var v81 int8 = 127
                                                                    _ = v81
                                                                } else {
                                                                    v37 = ((1e10 * (0.1 + ((3.14159 + ((float32((0.1 * 1e10)) / float32(1e10)) + -2.25)) / 1e10))) >= 1.5)
                                                                }
                                                                var v82 uint32 = v3
                                                                if true {
                                                                    if (!true) {
                                                                        switch "a" {
                                                                        case "a":
                                                                            if (true || (!v37)) {
                                                                                var v83 string = "\n\t"
                                                                                switch 4294967295 {
                                                                                case 9223372036854775807:
                                                                                    v84 := []int{(v1 * fi27), fi27, (v2 * (^(fi27)))}
                                                                                    _ = v84
                                                                                }
                                                                                v85 := []int{genFn0(200, false)}
                                                                                _ = v83
                                                                                _ = v85
                                                                            }
                                                                        case "hello":
                                                                            v86 := GenS8{F0: 537, F1: 100}
                                                                            _ = v86
                                                                        case "":
                                                                            var v87 uint16 = 1
                                                                            _ = v87
                                                                        }
                                                                        switch fi36 {
                                                                        case 2147483647:
                                                                            var v88 int16 = 42
                                                                            _ = v88
                                                                        default:
                                                                            fmt.Println(v56)
                                                                        }
                                                                    }
                                                                    var v89 int32 = 2147483647
                                                                    var v90 uint64 = 102
                                                                    _ = v89
                                                                    _ = v90
                                                                } else {
                                                                    switch v2 {
                                                                    case 269:
                                                                        if (fi36 == (-(fi27))) {
                                                                            var v91 int64 = -455
                                                                            _ = v91
                                                                        }
                                                                    default:
                                                                        var v92 int = (v2 << ((255) & 63))
                                                                        _ = v92
                                                                    }
                                                                }
                                                                _ = v82
                                                            } else {
                                                                v56 = (v56 << ((0) & 63))
                                                                v56 -= 1
                                                                v76 += v76
                                                            }
                                                            v3 = 65535
                                                            v93 := []int{fi27}
                                                            switch (v2 & v1) {
                                                            case -9223372036854775808:
                                                                if (!v37) {
                                                                    var v94 int32 = 1
                                                                    var v95 string = "a"
                                                                    v93 = append(v93, genFn0(uint64(v56), (v37 || (true || false))))
                                                                    _ = v94
                                                                    _ = v95
                                                                } else {
                                                                    var v96 uint16 = 255
                                                                    v56 -= (v56 | 2)
                                                                    var v97 int64 = -9223372036854775808
                                                                    _ = v96
                                                                    _ = v97
                                                                }
                                                            case 128:
                                                                switch "" {
                                                                case "hello":
                                                                    if len(v93) > 0 {
                                                                        v93[0] = fi27
                                                                    }
                                                                default:
                                                                    if len(v93) >= 2 {
                                                                        v93 = v93[1:len(v93)]
                                                                    }
                                                                }
                                                            case 1:
                                                                var v98 int64 = -1
                                                                _ = v98
                                                            default:
                                                                if true {
                                                                    if (!true) {
                                                                        if v37 {
                                                                            v99 := GenS7{F0: 0}
                                                                            _ = v99
                                                                        } else {
                                                                            v37 = (!(((v3 > (^(v3))) && v37) || false))
                                                                            switch genFn0(4294967295, (int32(v76) >= 1)) {
                                                                            case 42:
                                                                                if (fi36 >= fi27) {
                                                                                    fmt.Println(v76)
                                                                                    if len(v93) > 0 {
                                                                                        v93[0] = 526
                                                                                    }
                                                                                    if (v3 > (^(v3))) {
                                                                                        var v100 bool = (-9223372036854775808 < 509)
                                                                                        var v101 float32 = float32(float64(1.5))
                                                                                        _ = v100
                                                                                        _ = v101
                                                                                    } else {
                                                                                        var v102 string = ("\n\t" + "foo bar")
                                                                                        var v103 uint8 = v56
                                                                                        _ = v102
                                                                                        _ = v103
                                                                                    }
                                                                                    if v37 {
                                                                                        var v104 string = "hello"
                                                                                        _ = v104
                                                                                    } else {
                                                                                        if (32767 > -1) {
                                                                                            if len(v93) > 0 {
                                                                                                v93[0] = v2
                                                                                            }
                                                                                            var v105 int8 = 65
                                                                                            v105 -= (v105 % ((v105) | 1))
                                                                                            _ = v105
                                                                                        } else {
                                                                                            v1 -= v1
                                                                                        }
                                                                                        var v106 float64 = 1e10
                                                                                        v37 = (v37 || v37)
                                                                                        _ = v106
                                                                                    }
                                                                                    fmt.Println(v56)
                                                                                    var v107 string = ""
                                                                                    _ = v107
                                                                                }
                                                                            case 400:
                                                                                v37 = (!true)
                                                                            case -2147483648:
                                                                                v1 = (v1 - fi36)
                                                                            }
                                                                            v108 := GenS5{F0: 2147483647, F1: 0}
                                                                            if len(v93) > 0 {
                                                                                v93[0] = (v2 % ((fi27) | 1))
                                                                            }
                                                                            _ = v108
                                                                        }
                                                                    } else {
                                                                        if ((v37 || v37) || ((v76 & (v76 | v76)) != v76)) {
                                                                            fmt.Println(v1, v76)
                                                                            v2 -= fi36
                                                                        } else {
                                                                            v109 := GenS2{F0: -119, F1: 545, F2: 255}
                                                                            _ = v109
                                                                        }
                                                                        var v110 int = v2
                                                                        v2 = (fi36 << ((535) & 63))
                                                                        _ = v110
                                                                    }
                                                                } else {
                                                                    v37 = (uint8(v3) != uint8(v3))
                                                                    if (((0.0 > float32((((3.14159 * 0.0) + ((1e10 / 1.0) / -2.25)) - 1.5))) && true) && false) {
                                                                        v93 = append(v93, genFn0((^(v76)), (-2147483648 == -1)))
                                                                        var v111 uint32 = (v3 + uint32(v76))
                                                                        _ = v111
                                                                    }
                                                                    if len(v93) > 0 {
                                                                        v93[0] = 702
                                                                    }
                                                                    v76 = (v76 | 595)
                                                                }
                                                                if false {
                                                                    if v37 {
                                                                        var v112 int = 42
                                                                        var v113 int32 = 1
                                                                        _ = v112
                                                                        _ = v113
                                                                    }
                                                                    v114 := GenS9{F0: 1, F1: -409}
                                                                    _ = v114
                                                                }
                                                            }
                                                            var v115 uint32 = v3
                                                            if (3.14159 != 0.0) {
                                                                fmt.Println(fi36)
                                                                if v37 {
                                                                    var v116 int32 = int32(v56)
                                                                    v117 := GenS7{F0: 255}
                                                                    _ = v116
                                                                    _ = v117
                                                                } else {
                                                                    var v118 float64 = -2.25
                                                                    var v119 bool = v37
                                                                    if v119 {
                                                                        switch "" {
                                                                        case "\n\t":
                                                                            v37 = false
                                                                        }
                                                                        v1 += v1
                                                                        v93 = append(v93, -2147483648)
                                                                    } else {
                                                                        v76 -= 1
                                                                        if (genFn0(312, true) >= v2) {
                                                                            var v120 uint = 65535
                                                                            fmt.Println(v2, fi36, v115)
                                                                            _ = v120
                                                                        } else {
                                                                            fmt.Println(fi36)
                                                                        }
                                                                        v118 = v118
                                                                    }
                                                                    v76 = v76
                                                                    _ = v118
                                                                    _ = v119
                                                                }
                                                            } else {
                                                                fmt.Println(v3, v115, v2)
                                                                v121 := GenS9{F0: 42, F1: 1000}
                                                                if v37 {
                                                                    switch fi36 {
                                                                    case 9223372036854775807:
                                                                        var v122 uint8 = v56
                                                                        _ = v122
                                                                    }
                                                                    v56 = (v56 / (((v56 + v56)) | 1))
                                                                } else {
                                                                    fmt.Println(fi27, v1)
                                                                    if ((v1 << ((312) & 63)) == (fi27 % ((fi36) | 1))) {
                                                                        v123 := GenS5{F0: 2147483647, F1: -32768}
                                                                        _ = v123
                                                                    }
                                                                    v1 += v1
                                                                }
                                                                _ = v121
                                                            }
                                                            _ = v115
                                                            _ = v93
                                                        } else {
                                                            v37 = v37
                                                        }
                                                    }
                                                    var v124 uint32 = v3
                                                    v3 = (v124 / (((v3 & v124)) | 1))
                                                    if ((fi36 * (^(v1))) <= (v2 / ((genFn0(42, v37)) | 1))) {
                                                        if (42 < int64(v124)) {
                                                            if (255 >= 32767) {
                                                                var v125 uint8 = (^(v56))
                                                                _ = v125
                                                            }
                                                        }
                                                    } else {
                                                        v76 = v76
                                                        if v37 {
                                                            var v126 uint32 = 65535
                                                            v127 := GenS10{F0: 1}
                                                            _ = v126
                                                            _ = v127
                                                        } else {
                                                            var v128 uint16 = 255
                                                            _ = v128
                                                        }
                                                    }
                                                    var v129 int16 = 32767
                                                    _ = v124
                                                    _ = v129
                                                    _ = v76
                                                } else {
                                                    var v130 int = (^(fi27))
                                                    var v131 bool = false
                                                    if v131 {
                                                        var v132 int32 = -217
                                                        v133 := []int{v2, genFn0(4294967295, true), fi36, (v2 >> ((1) & 63))}
                                                        var v134 int8 = 0
                                                        v132 = v132
                                                        _ = v132
                                                        _ = v133
                                                        _ = v134
                                                    } else {
                                                        v3 -= (v3 << ((4294967295) & 63))
                                                        if v37 {
                                                            switch fi36 {
                                                            case 1:
                                                                v1 = genFn0(529, v131)
                                                            default:
                                                                switch -9223372036854775808 {
                                                                case 32767:
                                                                    var v135 int8 = 1
                                                                    _ = v135
                                                                }
                                                            }
                                                            var v136 float32 = 0.0
                                                            _ = v136
                                                        }
                                                        fmt.Println(fi27)
                                                        fmt.Println(v2)
                                                    }
                                                    _ = v130
                                                    _ = v131
                                                }
                                                v137 := GenS3{F0: fi36}
                                                v1 = 742
                                                _ = v137
                                            }
                                            _ = v56
                                            _ = v57
                                        }
                                    }
                                    v3 -= 836
                                }
                                _ = fi36
                                _ = v37
                            }
                        }
                        var v138 int16 = 713
                        _ = v138
                    }
                    var v139 uint = 18446744073709551615
                    _ = fi27
                    _ = v139
                }
                v3 = 2147483648
            }
        case 265:
            var v140 uint8 = 243
            _ = v140
        }
    }
    v141 := []int{(v1 - (v1 & v2)), genFn0(2, true), (-(v1)), (^(v1))}
    v142 := GenS4{F0: (42 >= 2), F1: -9223372036854775808}
    if (!true) {
        v143 := GenS11{F0: 1}
        v141 = append(v141, int(v3))
        _ = v143
    }
    v2 = genFn0(905, false)
    var v144 int = v2
    var v145 string = "hello"
    switch "hello" {
    case "hello":
        v146 := GenS2{F0: 0, F1: -9223372036854775808, F2: uint(v3)}
        _ = v146
    case "a":
        v144 = genFn0(42, true)
    case "\n\t":
        v147 := GenS12{F0: 42, F1: 635, F2: 9223372036854775807}
        _ = v147
    default:
        if true {
            v1 = 2147483647
            for fi148 := 0; fi148 < 5; fi148++ {
                for fi149 := 0; fi149 < 5; fi149++ {
                    if len(v141) > 0 {
                        v141[0] = fi149
                    }
                    v150 := GenS13{F0: 42, F1: 127, F2: 0}
                    var v151 uint = 643
                    _ = fi149
                    _ = v150
                    _ = v151
                }
                switch (^(fi148)) {
                case -1:
                    v144 = (fi148 >> ((0) & 63))
                default:
                    v1 = 42
                }
                _ = fi148
            }
        } else {
            _ = len(v141)
            if false {
                v3 += 2147483647
            }
            var v152 uint16 = 255
            _ = v152
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v141:", len(v141), v141)
    fmt.Println("v144:", v144)
    fmt.Println("v145:", v145)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v141
    _ = v142
    _ = v144
    _ = v145
    _ = v2
    _ = v3
}
