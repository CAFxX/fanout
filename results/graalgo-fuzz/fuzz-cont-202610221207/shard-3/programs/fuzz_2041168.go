package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 int16; F2 int8 }
type GenS3 struct { F0 int64; F1 int32; F2 int }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint8; F1 bool }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int16; F1 uint; F2 bool }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 uint16; F1 uint16; F2 uint16 }
type GenS11 struct { F0 bool; F1 int32; F2 uint32 }
func genFn0() int {
    v5 := GenS1{F0: 452}
    var v6 int8 = 0
    if true {
        v6 += v6
        if true {
            for fi7 := 0; fi7 < 12; fi7++ {
                switch fi7 {
                case -711:
                    v8 := GenS2{F0: 175, F1: -1, F2: v6}
                    _ = v8
                }
                _ = fi7
            }
        } else {
            v9 := GenS3{F0: 9223372036854775807, F1: -786, F2: 546}
            v10 := []int{-1, -433}
            switch "hello" {
            case "\n\t":
                fmt.Println(v6)
            }
            _ = v10
            _ = v9
        }
    } else {
        for fi11 := 0; fi11 < 11; fi11++ {
            v6 += (v6 * v6)
            v12 := GenS4{F0: 2}
            _ = fi11
            _ = v12
        }
        switch 127 {
        case 127:
            v13 := GenS1{F0: 168}
            _ = v13
        case -636:
            for fi14 := 0; fi14 < 8; fi14++ {
                fmt.Println(v6, v6)
                var v15 string = ("a" + ("" + ("" + "")))
                v16 := GenS4{F0: 1}
                _ = fi14
                _ = v15
                _ = v16
            }
        case -9223372036854775808:
            var v17 int = 937
            _ = v17
        default:
            v6 += (v6 << ((174) & 63))
        }
        v18 := []int{32767, -450, -321}
        _ = v18
    }
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 int16, p1 int16) int64 {
    for fi19 := 0; fi19 < 15; fi19++ {
        for fi20 := 0; fi20 < 2; fi20++ {
            p1 = p1
            p0 -= p1
            _ = fi20
        }
        _ = fi19
    }
    if false {
        v21 := GenS4{F0: 32767}
        _ = v21
    }
    p1 += p0
    _ = p0
    _ = p1
    return -1
}

func main() {
    var v1 string = ""
    var v2 int8 = 0
    var v3 float64 = ((float64(((1e10 - 3.14159) - 1e10)) - 1e10) * 1.5)
    v4 := make(map[string]string)
    var v22 int64 = 1
    v22 = 345
    var v23 float64 = (float64((0.0 + 1e10)) * ((v3 - v3) - ((v3 * float64(0.1)) - v3)))
    switch 127 {
    case 4294967295:
        if (!false) {
            var v24 bool = false
            delete(v4, (((v1 + v1) + ((v1 + ("\n\t" + v1)) + v1)) + v1))
            if (256 != 32767) {
                v25 := []int{828, -354}
                if len(v25) >= 2 {
                    v25 = v25[1:len(v25)]
                }
                _ = v25
            } else {
                if false {
                    v26 := []int{128, 127}
                    if (v24 || v24) {
                        if (65535 != 796) {
                            v1 = (v1 + (v1 + (v1 + v1)))
                            v27 := GenS2{F0: 0, F1: 1, F2: v2}
                            fmt.Println(v1)
                            _ = v27
                        }
                        v28 := v4["hello"]
                        var v29 float64 = (v3 + float64(((0.0 * 0.0) + (1e10 / 3.14159))))
                        _ = v28
                        _ = v29
                    } else {
                        v30 := GenS2{F0: 42, F1: 430, F2: 0}
                        _ = v30
                    }
                    v22 = v22
                    _ = v26
                } else {
                    v31 := GenS3{F0: v22, F1: 2147483647, F2: 128}
                    v24 = v24
                    _ = v31
                }
                for fi32 := 0; fi32 < 19; fi32++ {
                    for fi33 := 0; fi33 < 8; fi33++ {
                        v34 := GenS1{F0: 255}
                        switch v1 {
                        case "\n\t":
                            v35 := GenS3{F0: genFn1(-487, -1), F1: 0, F2: (fi32 / (((fi33 ^ fi32)) | 1))}
                            _ = v35
                        case "a":
                            v36 := GenS1{F0: 2}
                            _ = v36
                        }
                        _ = fi33
                        _ = v34
                    }
                    v37 := GenS5{F0: 81, F1: v24}
                    _ = fi32
                    _ = v37
                }
                var v38 uint16 = 2
                _ = v38
            }
            _ = v24
        }
    default:
        for fi39 := 0; fi39 < 17; fi39++ {
            v23 -= v3
            v40 := []int{}
            for fi41 := 0; fi41 < 1; fi41++ {
                if true {
                    v42 := GenS6{F0: (^(fi39))}
                    if true {
                        v43 := GenS4{F0: 32767}
                        _ = v43
                    }
                    var v44 int = fi41
                    _ = v42
                    _ = v44
                }
                delete(v4, ("a" + "\n\t"))
                var v45 uint32 = 1000
                var v46 int = (fi39 ^ 1)
                switch (v1 + ("a" + v1)) {
                case "foo bar":
                    v1 = "foo bar"
                    if true {
                        var v47 uint32 = v45
                        v1 = v1
                        v48 := GenS1{F0: 2}
                        _ = v47
                        _ = v48
                    } else {
                        v46 -= (fi39 + (fi39 / ((32767) | 1)))
                        fmt.Println(v3, v3, v2)
                        fmt.Println(fi41)
                    }
                    v2 = (^(v2))
                case "\n\t":
                    if (true || true) {
                        var v49 int = (fi39 | 128)
                        if false {
                            if len(v40) >= 2 {
                                v40 = v40[1:len(v40)]
                            }
                            v50 := GenS6{F0: fi39}
                            _ = v50
                        } else {
                            switch v49 {
                            case -9223372036854775808:
                                var v51 string = v1
                                _ = v51
                            case 128:
                                var v52 float32 = float32(((v3 * v23) - float64(-2.25)))
                                _ = v52
                            }
                        }
                        var v53 int8 = v2
                        if false {
                            if (!(false && false)) {
                                v22 += genFn1(-989, 0)
                                v46 = genFn0()
                                var v54 int = (v49 >> ((671) & 63))
                                _ = v54
                            }
                            switch v1 {
                            case "":
                                var v55 int32 = 0
                                _ = v55
                            case "\n\t":
                                v40 = append(v40, (v46 * (^(v49))))
                            case "a":
                                var v56 int64 = -1
                                _ = v56
                            }
                            v45 += v45
                        } else {
                            v2 = v53
                            v45 += (v45 ^ 2147483647)
                            var v57 int8 = v2
                            _ = v57
                        }
                        _ = v49
                        _ = v53
                    }
                case "hello":
                    v2 = -128
                    if ((v1 + ("" + ("\n\t" + (v1 + (v1 + "\n\t"))))) >= "foo bar") {
                        v23 = (-2.25 * float64(1e10))
                    }
                    if len(v40) >= 2 {
                        v40 = v40[1:len(v40)]
                    }
                default:
                    switch v1 {
                    case "":
                        switch v46 {
                        case -1:
                            fmt.Println(v46, v3, v23)
                        case 2147483647:
                            if (1 >= 1) {
                                var v58 int64 = v22
                                var v59 uint64 = 42
                                var v60 uint64 = (v59 | 208)
                                _ = v58
                                _ = v59
                                _ = v60
                            } else {
                                switch 128 {
                                case 663:
                                    var v61 string = ("" + "foo bar")
                                    _ = v61
                                case -2147483648:
                                    var v62 int = -2147483648
                                    _ = v62
                                case 32767:
                                    switch 226 {
                                    case 127:
                                        var v63 uint8 = 255
                                        if false {
                                            var v64 uint = 18446744073709551615
                                            var v65 bool = (42 <= 42)
                                            _ = v64
                                            _ = v65
                                        } else {
                                            v3 -= 0.1
                                        }
                                        _ = len(v40)
                                        _ = v63
                                    }
                                default:
                                    v66 := GenS1{F0: 247}
                                    fmt.Println(v22, v1, v23)
                                    _ = v66
                                }
                                var v67 int8 = 127
                                _ = v67
                            }
                        }
                    }
                    v68 := GenS5{F0: 1, F1: (2 == 319)}
                    _ = v68
                }
                _ = fi41
                _ = v45
                _ = v46
            }
            _ = fi39
            _ = v40
        }
    }
    switch ("" + (("\n\t" + ((((v1 + "foo bar") + v1) + v1) + ("\n\t" + "hello"))) + v1)) {
    case "":
        for fi69 := 0; fi69 < 20; fi69++ {
            fmt.Println(v1, v2)
            for fi70 := 0; fi70 < 18; fi70++ {
                v22 = (v22 | -1)
                v71 := GenS7{F0: (v22 >> ((255) & 63))}
                if (!(genFn1(-32768, 0) == (v22 ^ genFn1(1000, -448)))) {
                    v22 = -1
                } else {
                    if false {
                        v22 += v22
                        var v72 float32 = (1.5 / ((((0.1 + ((3.14159 - ((1e10 - (-2.25 + (1e10 + (1.5 * 1e10)))) + 1.5)) + 1.5)) + 1.5) / -2.25) * float32(0.1)))
                        v73 := GenS6{F0: (fi69 ^ genFn0())}
                        _ = v72
                        _ = v73
                    } else {
                        if (834 <= 1) {
                            if true {
                                switch fi70 {
                                case -9223372036854775808:
                                    v74 := []int{(fi69 | fi69), fi69, -9223372036854775808, fi70}
                                    _ = v74
                                case 807:
                                    v23 = 3.14159
                                case 2147483647:
                                    fmt.Println(v2, v23)
                                default:
                                    switch genFn0() {
                                    case -884:
                                        if (!(!true)) {
                                            v23 += ((v3 - ((v3 + (v23 / v23)) + (v23 / float64(((((-2.25 * 3.14159) * (-2.25 * 1.5)) / 1e10) / 0.1))))) + (((float64(1.5) / (v3 * (v3 / v23))) + ((v3 + (float64((-2.25 + float32(v3))) - float64((0.1 + 1.5)))) * -2.25)) * float64(1.5)))
                                        } else {
                                            v4[(v1 + "foo bar")] = "a"
                                            var v75 bool = (0.0 != 1e10)
                                            _ = v75
                                        }
                                    case 1:
                                        v76 := []int{genFn0()}
                                        _ = v76
                                    case 32767:
                                        v77 := GenS1{F0: 444}
                                        _ = v77
                                    }
                                }
                                if false {
                                    switch (v1 + v1) {
                                    case "foo bar":
                                        var v78 uint16 = 137
                                        _ = v78
                                    case "a":
                                        _ = len(v4)
                                    case "hello":
                                        if true {
                                            if false {
                                                fmt.Println(v22, v22)
                                            } else {
                                                if true {
                                                    switch "hello" {
                                                    case "hello":
                                                        if false {
                                                            var v79 uint16 = 140
                                                            if (!((fi69 & (fi70 * (fi69 - (^(fi70))))) < genFn0())) {
                                                                var v80 uint64 = 655
                                                                if ((v22 - genFn1(-621, 32767)) > (v22 >> ((2) & 63))) {
                                                                    var v81 bool = (!((v2 | v2) <= 100))
                                                                    v4[v1] = "foo bar"
                                                                    v82 := GenS7{F0: int64(fi69)}
                                                                    delete(v4, "a")
                                                                    _ = v81
                                                                    _ = v82
                                                                } else {
                                                                    var v83 int8 = v2
                                                                    var v84 int8 = v2
                                                                    _ = v83
                                                                    _ = v84
                                                                }
                                                                _ = v80
                                                            } else {
                                                                v79 = v79
                                                                switch (fi70 + 42) {
                                                                case -2147483648:
                                                                    if false {
                                                                        v85 := GenS7{F0: genFn1(-1, int16(fi69))}
                                                                        var v86 bool = (genFn1(462, 849) > genFn1(876, -32768))
                                                                        _ = len(v4)
                                                                        var v87 float32 = -2.25
                                                                        v23 = (v3 * -2.25)
                                                                        switch genFn0() {
                                                                        case 1:
                                                                            var v88 float32 = (v87 * (float32(v3) * -2.25))
                                                                            _ = v88
                                                                        case 127:
                                                                            if (v86 || (v86 || (false || (v86 && (((v22 >= (v22 - genFn1(int16(v79), -32768))) || (true || ((!(-46 >= (v2 >> ((65535) & 63)))) || true))) || (!v86)))))) {
                                                                                fmt.Println(v2)
                                                                            } else {
                                                                                var v89 float32 = v87
                                                                                if v86 {
                                                                                    v90 := GenS6{F0: 666}
                                                                                    _ = v90
                                                                                }
                                                                                fmt.Println(v1, fi69, v3)
                                                                                var v91 int32 = -1
                                                                                var v92 string = (v1 + (v1 + "hello"))
                                                                                _ = len(v4)
                                                                                _ = v89
                                                                                _ = v91
                                                                                _ = v92
                                                                            }
                                                                        case 599:
                                                                            fmt.Println(v22)
                                                                        }
                                                                        v2 = v2
                                                                        v93 := GenS8{F0: 242, F1: 18446744073709551615, F2: v86}
                                                                        _ = v85
                                                                        _ = v86
                                                                        _ = v87
                                                                        _ = v93
                                                                    }
                                                                    if (42 != 810) {
                                                                        v3 -= v23
                                                                        if true {
                                                                            v94 := GenS9{F0: 2147483647}
                                                                            _ = v94
                                                                        } else {
                                                                            switch v1 {
                                                                            case "hello":
                                                                                var v95 float64 = 0.0
                                                                                _ = v95
                                                                            default:
                                                                                fmt.Println(fi69)
                                                                            }
                                                                            if (1.5 <= (3.14159 + v23)) {
                                                                                var v96 uint8 = 0
                                                                                var v97 string = ((v1 + ("a" + (v1 + (v1 + v1)))) + v1)
                                                                                _ = v96
                                                                                _ = v97
                                                                            } else {
                                                                                if (false && true) {
                                                                                    var v98 bool = (0 <= 1)
                                                                                    v23 = float64((3.14159 * 1.5))
                                                                                    _ = len(v4)
                                                                                    _ = v98
                                                                                }
                                                                                switch ("" + v1) {
                                                                                case "foo bar":
                                                                                    if true {
                                                                                        if (42 > genFn1(32767, 1)) {
                                                                                            var v99 float32 = ((1.5 - (-2.25 / ((1.5 + 1.5) - float32((0.1 + -2.25))))) + (-2.25 * -2.25))
                                                                                            var v100 uint32 = 42
                                                                                            _ = v100
                                                                                            _ = v99
                                                                                        }
                                                                                    }
                                                                                case "":
                                                                                    v4[v1] = v1
                                                                                case "a":
                                                                                    fmt.Println(v79, v22, v2)
                                                                                }
                                                                            }
                                                                        }
                                                                        if false {
                                                                            fmt.Println(fi70, v23, v23)
                                                                            v101 := []int{fi70, (-(fi69)), (fi70 & (fi70 - 772)), (fi69 + 42)}
                                                                            _ = v101
                                                                        } else {
                                                                            v4[(v1 + v1)] = ("foo bar" + "\n\t")
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                            var v102 uint16 = (^(v79))
                                                            _ = v102
                                                            _ = v79
                                                        } else {
                                                            _ = len(v4)
                                                            var v103 int32 = 42
                                                            _ = v103
                                                        }
                                                    default:
                                                        v104 := v4["foo bar"]
                                                        _ = v104
                                                    }
                                                } else {
                                                    v22 += (v22 + int64(fi69))
                                                    v105 := []int{genFn0(), (fi69 & 1), fi69, genFn0()}
                                                    if (65535 > 65535) {
                                                        var v106 uint32 = 2147483647
                                                        var v107 int64 = (v22 / (((-(v22))) | 1))
                                                        fmt.Println(v106, v107)
                                                        _ = v106
                                                        _ = v107
                                                    } else {
                                                        v3 -= ((v3 + (v23 * v23)) - v23)
                                                        var v108 int8 = v2
                                                        _ = v108
                                                    }
                                                    _ = v105
                                                }
                                            }
                                            if (!(2 > 42)) {
                                                v109 := []int{fi70, genFn0(), fi69}
                                                _ = v109
                                            }
                                        } else {
                                            v1 = "a"
                                            if (255 < 2) {
                                                v110 := GenS6{F0: fi70}
                                                v2 = 1
                                                if false {
                                                    switch v1 {
                                                    case "hello":
                                                        v111 := GenS6{F0: 886}
                                                        _ = v111
                                                    case "foo bar":
                                                        var v112 uint8 = 2
                                                        _ = v112
                                                    default:
                                                        if false {
                                                            var v113 bool = (("hello" + "\n\t") > (v1 + v1))
                                                            _ = v113
                                                        }
                                                    }
                                                }
                                                _ = v110
                                            }
                                        }
                                    default:
                                        fmt.Println(v2, v22, v23)
                                        fmt.Println(v2, v2, v23)
                                    }
                                    _ = len(v4)
                                    _ = len(v4)
                                } else {
                                    v114 := v4[(("a" + ("hello" + ((v1 + v1) + v1))) + v1)]
                                    var v115 uint = 1
                                    fmt.Println(v114, v22)
                                    _ = v114
                                    _ = v115
                                }
                                v116 := GenS5{F0: 127, F1: (42 == 75)}
                                _ = v116
                            }
                            v117 := GenS10{F0: 2, F1: 530, F2: 583}
                            var v118 int32 = 973
                            _ = v117
                            _ = v118
                        }
                    }
                    v1 = v1
                    var v119 uint32 = 596
                    _ = v119
                }
                _ = fi70
                _ = v71
            }
            _ = fi69
        }
    case "hello":
        var v120 int8 = (v2 ^ v2)
        _ = v120
    }
    switch 128 {
    case 61:
        switch "a" {
        case "hello":
            v121 := GenS11{F0: false, F1: 1, F2: 2147483647}
            _ = v121
        default:
            var v122 float32 = float32(1.5)
            _ = v122
        }
    }
    delete(v4, "")
    v23 = v23
    var v123 uint16 = 42
    if true {
        v123 = v123
        fmt.Println(v2, v1, v2)
        fmt.Println(v2, v123)
    }
    v123 = (v123 % (((v123 | 255)) | 1))
    v2 += v2
    fmt.Println("v1:", v1)
    fmt.Println("v123:", v123)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v123
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
