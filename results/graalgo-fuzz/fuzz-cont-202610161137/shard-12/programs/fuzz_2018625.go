package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 bool; F2 uint32 }
type GenS3 struct { F0 int8; F1 bool; F2 uint16 }
type GenS4 struct { F0 string; F1 uint8 }
type GenS5 struct { F0 uint16; F1 uint32; F2 uint32 }
type GenS6 struct { F0 int; F1 int16; F2 uint8 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 int; F1 string; F2 int }
type GenS10 struct { F0 int; F1 int8; F2 int32 }
type GenS11 struct { F0 bool; F1 bool; F2 int8 }
type GenS12 struct { F0 bool }
type GenS13 struct { F0 int64; F1 int; F2 int64 }
func genFn0(p0 int, p1 int64, p2 int64) int {
    var v4 int64 = (p1 - (p2 & p2))
    for fi5 := 0; fi5 < 14; fi5++ {
        v6 := GenS1{F0: 0}
        var v7 uint16 = 2
        switch fi5 {
        case 32767:
            var v8 int8 = 51
            _ = v8
        case 244:
            var v9 int8 = 127
            _ = v9
        case -757:
            if false {
                for fi10 := 0; fi10 < 20; fi10++ {
                    fmt.Println(v4)
                    var v11 int32 = 109
                    _ = fi10
                    _ = v11
                }
                if true {
                    switch "foo bar" {
                    case "\n\t":
                        p2 -= (v4 >> ((2) & 63))
                    }
                    if (42 > 65535) {
                        v12 := GenS2{F0: false, F1: false, F2: 0}
                        _ = v12
                    }
                    for fi13 := 0; fi13 < 15; fi13++ {
                        v14 := []int{fi13}
                        if true {
                            if len(v14) >= 2 {
                                v14 = v14[1:len(v14)]
                            }
                        } else {
                            p1 -= (^(v4))
                            var v15 uint16 = v7
                            v14 = append(v14, fi13)
                            v16 := GenS1{F0: (-(p2))}
                            v4 -= p1
                            _ = v15
                            _ = v16
                        }
                        _ = fi13
                        _ = v14
                    }
                }
                v17 := GenS1{F0: p1}
                _ = v17
            }
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    v18 := []int{(p0 << ((282) & 63)), -258, p0}
    for fi19 := 0; fi19 < 18; fi19++ {
        v20 := GenS1{F0: (p1 - (v4 * 300))}
        for fi21 := 0; fi21 < 18; fi21++ {
            var v22 uint16 = 1
            switch 9223372036854775807 {
            case -359:
                if len(v18) > 0 {
                    v18[0] = (-(p0))
                }
            case 9223372036854775807:
                p1 += -601
            case -1:
                if len(v18) >= 2 {
                    v18 = v18[1:len(v18)]
                }
            default:
                v18 = append(v18, 4294967295)
            }
            _ = fi21
            _ = v22
        }
        _ = fi19
        _ = v20
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v4
    return (p0 | 0)
}

func main() {
    var v1 int16 = 474
    var v2 int8 = 42
    var v3 uint32 = 4294967295
    var v23 int16 = v1
    for fi24 := 0; fi24 < 7; fi24++ {
        if false {
            if (true && false) {
                var v25 float64 = 0.0
                for fi26 := 0; fi26 < 3; fi26++ {
                    if false {
                        switch fi24 {
                        case 362:
                            if true {
                                var v27 int8 = (-(v2))
                                if false {
                                    var v28 uint64 = 1
                                    if (("hello" < "a") || false) {
                                        fmt.Println(v2, fi26, v28)
                                        switch ("foo bar" + "\n\t") {
                                        case "":
                                            fmt.Println(v3, v28, v1)
                                        case "foo bar":
                                            fmt.Println(v25)
                                            v29 := GenS3{F0: (-(v2)), F1: false, F2: uint16(v28)}
                                            v30 := []int{(fi26 * fi24)}
                                            _ = v29
                                            _ = v30
                                        case "\n\t":
                                            var v31 uint8 = 255
                                            _ = v31
                                        }
                                    } else {
                                        v32 := []int{}
                                        switch genFn0(genFn0((fi26 % ((genFn0((^(fi24)), -1, -9223372036854775808)) | 1)), -1, 0), -1, -319) {
                                        case 32768:
                                            v1 += v23
                                        case 0:
                                            switch "a" {
                                            case "":
                                                if false {
                                                    var v33 float64 = 1.5
                                                    v23 = (v1 | 42)
                                                    _ = v33
                                                }
                                                var v34 int32 = 0
                                                var v35 float64 = 3.14159
                                                switch "\n\t" {
                                                case "\n\t":
                                                    if true {
                                                        if (!false) {
                                                            var v36 int32 = v34
                                                            if (!((float64((1e10 + 0.1)) - float64(float32(0.1))) < 1.5)) {
                                                                var v37 uint64 = (v28 >> ((65535) & 63))
                                                                v38 := GenS4{F0: "a", F1: 0}
                                                                _ = v37
                                                                _ = v38
                                                            }
                                                            var v39 string = "foo bar"
                                                            _ = v36
                                                            _ = v39
                                                        } else {
                                                            var v40 uint64 = v28
                                                            _ = v40
                                                        }
                                                    } else {
                                                        fmt.Println(v35, fi26, v27)
                                                        if (((!false) || ((127 == 127) && true)) && true) {
                                                            v41 := GenS2{F0: true, F1: false, F2: (v3 / ((v3) | 1))}
                                                            if (0.0 <= 0.1) {
                                                                var v42 string = ("foo bar" + ("\n\t" + (("a" + "hello") + "a")))
                                                                _ = v42
                                                            }
                                                            v1 = (^(v1))
                                                            _ = v41
                                                        } else {
                                                            if ((^(v3)) < (v3 * v3)) {
                                                                switch "foo bar" {
                                                                case "":
                                                                    v43 := GenS5{F0: uint16(v3), F1: (v3 + 2147483647), F2: (v3 >> ((113) & 63))}
                                                                    _ = v43
                                                                case "a":
                                                                    switch -365 {
                                                                    case 4294967295:
                                                                        switch (("" + "foo bar") + "hello") {
                                                                        case "\n\t":
                                                                            var v44 int16 = (v23 ^ int16(v34))
                                                                            _ = v44
                                                                        case "hello":
                                                                            var v45 uint8 = 127
                                                                            if (true && false) {
                                                                                v3 -= 1
                                                                                if false {
                                                                                    v46 := GenS2{F0: (true || true), F1: true, F2: v3}
                                                                                    if len(v32) > 0 {
                                                                                        v32[0] = 9223372036854775807
                                                                                    }
                                                                                    var v47 int64 = 1
                                                                                    v32 = append(v32, genFn0(634, v47, v47))
                                                                                    _ = v46
                                                                                    _ = v47
                                                                                } else {
                                                                                    v23 += 652
                                                                                    v28 = (v28 | uint64(v45))
                                                                                    var v48 int8 = (^(v27))
                                                                                    _ = v48
                                                                                }
                                                                            } else {
                                                                                var v49 bool = true
                                                                                _ = v49
                                                                            }
                                                                            if false {
                                                                                if (v45 >= v45) {
                                                                                    switch fi24 {
                                                                                    case 0:
                                                                                        v25 -= float64((((0.0 - 1e10) / 1.5) + (((0.1 / (0.1 * (1e10 - (0.0 * (1e10 / (1e10 / 1.5)))))) + (0.0 * 0.1)) + 0.0)))
                                                                                        var v50 uint32 = v3
                                                                                        _ = v50
                                                                                    case -1:
                                                                                        var v51 uint = 0
                                                                                        _ = v51
                                                                                    }
                                                                                    var v52 uint32 = (^(v3))
                                                                                    _ = v52
                                                                                } else {
                                                                                    v53 := GenS6{F0: genFn0(int(v2), 9223372036854775807, -1), F1: -204, F2: (v45 ^ 213)}
                                                                                    _ = v53
                                                                                }
                                                                            }
                                                                            v54 := GenS7{F0: ("a" > "")}
                                                                            v35 -= (v35 - -2.25)
                                                                            _ = v45
                                                                            _ = v54
                                                                        default:
                                                                            fmt.Println(v35, v3)
                                                                        }
                                                                    case 112:
                                                                        fmt.Println(v1, v27)
                                                                    case 128:
                                                                        var v55 uint8 = 42
                                                                        _ = v55
                                                                    default:
                                                                        _ = len(v32)
                                                                        v56 := GenS6{F0: 30, F1: (v23 + v1), F2: 241}
                                                                        _ = v56
                                                                    }
                                                                case "foo bar":
                                                                    if (true || (1 >= 630)) {
                                                                        if (566 != 796) {
                                                                            var v57 int64 = 567
                                                                            _ = v57
                                                                        } else {
                                                                            if (v1 > (v1 + 1000)) {
                                                                                var v58 uint32 = (v3 | v3)
                                                                                _ = v58
                                                                            } else {
                                                                                if true {
                                                                                    if len(v32) >= 2 {
                                                                                        v32 = v32[1:len(v32)]
                                                                                    }
                                                                                    if len(v32) >= 2 {
                                                                                        v32 = v32[1:len(v32)]
                                                                                    }
                                                                                    if ((121 >= 32767) || (fi24 > fi26)) {
                                                                                        v59 := GenS8{F0: 255}
                                                                                        _ = v59
                                                                                    }
                                                                                } else {
                                                                                    v60 := GenS7{F0: true}
                                                                                    fmt.Println(v23, v25)
                                                                                    fmt.Println(v1, v1)
                                                                                    _ = v60
                                                                                }
                                                                            }
                                                                            v61 := GenS9{F0: 32768, F1: "", F2: genFn0((fi26 % ((genFn0(662, 42, int64(v28))) | 1)), 537, 239)}
                                                                            v62 := GenS6{F0: fi26, F1: -228, F2: 80}
                                                                            _ = v61
                                                                            _ = v62
                                                                        }
                                                                        var v63 string = ""
                                                                        v32 = append(v32, 1)
                                                                        v64 := GenS3{F0: 0, F1: (714 == 32767), F2: 42}
                                                                        _ = v63
                                                                        _ = v64
                                                                    }
                                                                }
                                                                fmt.Println(fi24)
                                                            } else {
                                                                fmt.Println(v3)
                                                                if (-182 == (fi24 | -578)) {
                                                                    var v65 int8 = v2
                                                                    switch fi24 {
                                                                    case -2147483648:
                                                                        v2 = v2
                                                                    case 128:
                                                                        _ = len(v32)
                                                                    default:
                                                                        fmt.Println(v2)
                                                                    }
                                                                    _ = v65
                                                                }
                                                                var v66 uint32 = 2147483648
                                                                v67 := GenS10{F0: (fi24 << ((4294967295) & 63)), F1: v27, F2: v34}
                                                                fmt.Println(v27)
                                                                fmt.Println(v34, fi24)
                                                                if false {
                                                                    if false {
                                                                        var v68 float64 = v25
                                                                        _ = v68
                                                                    }
                                                                    if len(v32) >= 2 {
                                                                        v32 = v32[1:len(v32)]
                                                                    }
                                                                } else {
                                                                    var v69 string = ("foo bar" + ("hello" + "hello"))
                                                                    _ = v69
                                                                }
                                                                switch ("foo bar" + (("hello" + "") + ("" + ("" + (("" + "foo bar") + "\n\t"))))) {
                                                                case "foo bar":
                                                                    fmt.Println(v23)
                                                                default:
                                                                    v23 = v23
                                                                }
                                                                var v70 int16 = 1
                                                                _ = v66
                                                                _ = v67
                                                                _ = v70
                                                            }
                                                            if false {
                                                                var v71 float32 = 0.0
                                                                _ = v71
                                                            } else {
                                                                v72 := GenS1{F0: -945}
                                                                _ = v72
                                                            }
                                                            if false {
                                                                if false {
                                                                    v25 = v35
                                                                    if true {
                                                                        fmt.Println(v35, v23)
                                                                    } else {
                                                                        switch "" {
                                                                        case "foo bar":
                                                                            var v73 uint16 = 255
                                                                            _ = v73
                                                                        case "hello":
                                                                            var v74 int64 = 42
                                                                            _ = v74
                                                                        default:
                                                                            var v75 string = "\n\t"
                                                                            _ = v75
                                                                        }
                                                                        fmt.Println(v28, fi24)
                                                                        if false {
                                                                            switch genFn0(fi24, 512, 42) {
                                                                            case -834:
                                                                                _ = len(v32)
                                                                            }
                                                                            var v76 int64 = 0
                                                                            var v77 int64 = (v76 - (v76 % ((-1) | 1)))
                                                                            _ = v76
                                                                            _ = v77
                                                                        } else {
                                                                            v78 := GenS2{F0: true, F1: (!(int8(v34) < v27)), F2: v3}
                                                                            fmt.Println(v2)
                                                                            if ((v2 < (-(v2))) || true) {
                                                                                _ = len(v32)
                                                                                fmt.Println(v2, v23)
                                                                            } else {
                                                                                var v79 uint32 = 42
                                                                                if len(v32) >= 2 {
                                                                                    v32 = v32[1:len(v32)]
                                                                                }
                                                                                _ = v79
                                                                            }
                                                                            var v80 string = ""
                                                                            _ = v78
                                                                            _ = v80
                                                                        }
                                                                    }
                                                                    v34 = v34
                                                                }
                                                                v81 := GenS11{F0: (!false), F1: (40 != 100), F2: (-(v27))}
                                                                var v82 int8 = (v2 * 127)
                                                                _ = v81
                                                                _ = v82
                                                            } else {
                                                                v32 = append(v32, genFn0(1, -509, 9223372036854775807))
                                                                v83 := GenS9{F0: (fi26 - fi24), F1: (("" + "") + ""), F2: fi26}
                                                                fmt.Println(v1)
                                                                _ = v83
                                                            }
                                                        }
                                                    }
                                                case "a":
                                                    v28 = (v28 + (v28 % (((v28 * 42)) | 1)))
                                                }
                                                v34 = (v34 * -916)
                                                _ = v34
                                                _ = v35
                                            case "foo bar":
                                                v32 = append(v32, (fi24 & (fi26 & fi24)))
                                            case "a":
                                                v25 += (v25 + v25)
                                            default:
                                                _ = len(v32)
                                            }
                                        case 464:
                                            var v84 uint8 = 128
                                            _ = v84
                                        }
                                        _ = v32
                                    }
                                    var v85 uint16 = 1
                                    _ = v28
                                    _ = v85
                                } else {
                                    v86 := GenS12{F0: false}
                                    _ = v86
                                }
                                _ = v27
                            }
                        case -182:
                            var v87 uint32 = (^(v3))
                            _ = v87
                        case 42:
                            v88 := GenS7{F0: true}
                            _ = v88
                        default:
                            if (255 >= 18446744073709551615) {
                                v1 += v1
                            }
                            if (!(false && false)) {
                                v1 += 0
                                v3 += v3
                                if true {
                                    fmt.Println(fi24, fi26)
                                    v89 := GenS10{F0: fi26, F1: int8(fi26), F2: 0}
                                    var v90 float32 = (0.0 * 3.14159)
                                    _ = v89
                                    _ = v90
                                }
                                if true {
                                    if false {
                                        var v91 string = ("foo bar" + "foo bar")
                                        v3 += v3
                                        if ((genFn0(-1, 1, 42) <= (-(fi26))) || true) {
                                            var v92 float32 = 0.1
                                            var v93 uint64 = 18446744073709551615
                                            if (false || true) {
                                                var v94 string = ""
                                                v95 := []int{}
                                                switch (fi24 << ((2) & 63)) {
                                                case 128:
                                                    if len(v95) >= 2 {
                                                        v95 = v95[1:len(v95)]
                                                    }
                                                case -9223372036854775808:
                                                    var v96 uint16 = 0
                                                    _ = v96
                                                case 588:
                                                    if len(v95) > 0 {
                                                        v95[0] = fi24
                                                    }
                                                default:
                                                    var v97 uint = 2
                                                    _ = v97
                                                }
                                                _ = v94
                                                _ = v95
                                            } else {
                                                var v98 int64 = 9223372036854775807
                                                _ = v98
                                            }
                                            _ = v92
                                            _ = v93
                                        }
                                        _ = v91
                                    }
                                } else {
                                    v99 := GenS11{F0: (!false), F1: (!(v3 != (v3 % ((2) | 1)))), F2: v2}
                                    v100 := GenS13{F0: int64(v3), F1: (-(fi26)), F2: -670}
                                    switch (-(fi24)) {
                                    case -29:
                                        fmt.Println(fi24)
                                    case 660:
                                        v101 := []int{(fi24 >> ((18446744073709551615) & 63)), fi26}
                                        _ = v101
                                    default:
                                        switch genFn0(128, -1, 0) {
                                        case -2147483648:
                                            v102 := []int{}
                                            _ = v102
                                        case 1:
                                            var v103 uint64 = 0
                                            _ = v103
                                        }
                                    }
                                    _ = v100
                                    _ = v99
                                }
                                v104 := []int{}
                                _ = v104
                            } else {
                                if (fi26 <= fi26) {
                                    v105 := GenS10{F0: int(v23), F1: (v2 & (v2 - (v2 + -101))), F2: 971}
                                    _ = v105
                                } else {
                                    if (v3 > (v3 & v3)) {
                                        v106 := GenS11{F0: ((856 == 4294967295) || false), F1: (!(!true)), F2: (^(v2))}
                                        v107 := []int{fi26}
                                        var v108 float32 = -2.25
                                        var v109 int32 = 495
                                        _ = v106
                                        _ = v107
                                        _ = v108
                                        _ = v109
                                    }
                                    v110 := []int{32767, -1, fi24, genFn0((fi24 << ((42) & 63)), 1, 42)}
                                    _ = v110
                                }
                                var v111 uint = 18446744073709551615
                                _ = v111
                            }
                            fmt.Println(v23)
                        }
                        if true {
                            v25 -= (v25 * (v25 + 3.14159))
                        } else {
                            if false {
                                v25 -= ((((v25 / v25) + v25) - (float64(float32(v25)) * ((0.0 - v25) * 0.0))) * v25)
                            } else {
                                if true {
                                    v25 = v25
                                    v25 -= v25
                                } else {
                                    fmt.Println(v3)
                                    fmt.Println(v3)
                                    var v112 uint32 = (v3 ^ 4294967295)
                                    _ = v112
                                }
                                var v113 uint32 = (v3 % ((2147483648) | 1))
                                _ = v113
                            }
                        }
                    } else {
                        v1 += (v1 - v1)
                        var v114 uint8 = 127
                        switch 2147483647 {
                        case 1:
                            v2 = (v2 >> ((65535) & 63))
                        }
                        _ = v114
                    }
                    _ = fi26
                }
                _ = v25
            }
        } else {
            var v115 uint8 = 121
            _ = v115
        }
        _ = fi24
    }
    for fi116 := 0; fi116 < 18; fi116++ {
        v117 := []int{genFn0(fi116, 1, -1), genFn0(fi116, -9223372036854775808, 9223372036854775807)}
        _ = fi116
        _ = v117
    }
    for fi118 := 0; fi118 < 1; fi118++ {
        switch "a" {
        case "hello":
            v3 -= v3
        default:
            var v119 bool = (2147483647 != 0)
            _ = v119
        }
        for fi120 := 0; fi120 < 4; fi120++ {
            fmt.Println(fi118, fi118, fi118)
            v121 := GenS3{F0: 127, F1: (true || false), F2: 30}
            fmt.Println(fi118, v1)
            var v122 int64 = -42
            _ = fi120
            _ = v121
            _ = v122
        }
        v123 := []int{(fi118 ^ genFn0((fi118 * (fi118 * (fi118 % ((fi118) | 1)))), -1, -9223372036854775808))}
        _ = fi118
        _ = v123
    }
    fmt.Println(v1, v1)
    fmt.Println(v2)
    v124 := []int{}
    v2 = v2
    if len(v124) > 0 {
        v124[0] = -807
    }
    v1 -= (v23 ^ 42)
    _ = len(v124)
    var v125 int32 = -1
    var v126 string = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v124:", len(v124), v124)
    fmt.Println("v125:", v125)
    fmt.Println("v126:", v126)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v124
    _ = v125
    _ = v126
    _ = v2
    _ = v23
    _ = v3
}
