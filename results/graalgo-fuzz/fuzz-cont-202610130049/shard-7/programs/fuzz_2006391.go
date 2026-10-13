package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16; F1 uint8 }
type GenS4 struct { F0 uint32; F1 bool; F2 int64 }
type GenS5 struct { F0 uint; F1 uint; F2 uint64 }
type GenS6 struct { F0 string; F1 int16 }
type GenS7 struct { F0 string; F1 int32 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 string; F1 uint8 }
type GenS10 struct { F0 int64; F1 uint32; F2 int }
type GenS11 struct { F0 uint8; F1 uint; F2 bool }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 int8 }
func genFn0(p0 uint, p1 bool) int16 {
    v7 := GenS1{F0: (("foo bar" + ("\n\t" + "foo bar")) + "\n\t"), F1: -2147483648, F2: 32767}
    for fi8 := 0; fi8 < 13; fi8++ {
        switch fi8 {
        case -10:
            v9 := []int{32768, fi8, (-(fi8)), (fi8 - fi8)}
            _ = v9
        default:
            for fi10 := 0; fi10 < 10; fi10++ {
                var v11 bool = (0.0 == 0.0)
                if v11 {
                    switch "hello" {
                    case "a":
                        p1 = true
                    case "foo bar":
                        var v12 string = "foo bar"
                        _ = v12
                    case "hello":
                        var v13 uint16 = uint16(p0)
                        _ = v13
                    default:
                        p1 = ((!p1) || true)
                    }
                } else {
                    var v14 int8 = -128
                    p0 = (^(p0))
                    var v15 string = ("foo bar" + "foo bar")
                    _ = v14
                    _ = v15
                }
                if (int32(fi10) <= 2147483647) {
                    var v16 bool = (2 > 127)
                    v17 := []int{fi8, 4294967295}
                    _ = v16
                    _ = v17
                } else {
                    v18 := GenS1{F0: (("a" + ("" + "")) + ""), F1: -750, F2: 32767}
                    v19 := GenS1{F0: "foo bar", F1: -705, F2: 274}
                    if v11 {
                        if v11 {
                            var v20 float32 = (0.1 / 1.0)
                            if p1 {
                                var v21 bool = ((float64(v20) != (1e10 / -2.25)) || false)
                                if v11 {
                                    v20 = float32(0.1)
                                } else {
                                    p0 -= p0
                                    var v22 bool = (p0 != (^(p0)))
                                    p1 = v11
                                    v23 := GenS2{F0: 4294967295}
                                    _ = v22
                                    _ = v23
                                }
                                _ = v21
                            }
                            if v11 {
                                switch "" {
                                case "\n\t":
                                    if false {
                                        p1 = v11
                                        switch "foo bar" {
                                        case "":
                                            fmt.Println(fi10, v11, p0)
                                        case "a":
                                            v24 := GenS3{F0: 42, F1: 146}
                                            _ = v24
                                        case "\n\t":
                                            switch 957 {
                                            case 1:
                                                if (-1 < 42) {
                                                    var v25 bool = (!("" < ("" + "\n\t")))
                                                    if (false || v11) {
                                                        var v26 int16 = 1
                                                        _ = v26
                                                    }
                                                    switch (fi8 + fi10) {
                                                    case 566:
                                                        var v27 float64 = -2.25
                                                        _ = v27
                                                    case 1:
                                                        if ("a" < "\n\t") {
                                                            v20 = float32((float64((v20 * (float32(-2.25) - v20))) * 0.1))
                                                            var v28 uint16 = 550
                                                            if p1 {
                                                                p0 = 588
                                                                v29 := []int{fi8, (fi10 % (((fi8 / ((fi8) | 1))) | 1))}
                                                                _ = v29
                                                            }
                                                            _ = v28
                                                        }
                                                    }
                                                    _ = v25
                                                } else {
                                                    if (false || (!false)) {
                                                        var v30 uint32 = 675
                                                        switch (-(fi10)) {
                                                        case 482:
                                                            v31 := []int{-1, 182}
                                                            _ = v31
                                                        case 2147483647:
                                                            if ((-(fi10)) < fi8) {
                                                                v32 := GenS4{F0: (v30 % (((v30 & v30)) | 1)), F1: (!v11), F2: -9223372036854775808}
                                                                _ = v32
                                                            }
                                                        default:
                                                            fmt.Println(v11, fi8, p0)
                                                        }
                                                        v20 += v20
                                                        _ = v30
                                                    }
                                                }
                                            case 32767:
                                                switch (fi8 % (((fi8 ^ 2147483647)) | 1)) {
                                                case 2147483647:
                                                    v11 = (255 <= 255)
                                                default:
                                                    p0 = p0
                                                }
                                            default:
                                                var v33 int8 = -1
                                                _ = v33
                                            }
                                        default:
                                            fmt.Println(p1, fi10, p1)
                                            if true {
                                                var v34 bool = false
                                                switch (fi10 >> ((229) & 63)) {
                                                case -904:
                                                    var v35 int16 = 42
                                                    _ = v35
                                                case 32768:
                                                    fmt.Println(v20, fi10, v11)
                                                default:
                                                    switch ("a" + "a") {
                                                    case "":
                                                        if (v34 || (true || true)) {
                                                            switch ("" + (((("a" + (("a" + "foo bar") + ("foo bar" + ("foo bar" + ("foo bar" + (("" + "hello") + "foo bar")))))) + "") + "") + "")) {
                                                            case "foo bar":
                                                                if p1 {
                                                                    v20 = 0.1
                                                                    if (((fi10 ^ 128) == fi8) && (v11 || ((v11 || (28 != 206)) && v11))) {
                                                                        if v11 {
                                                                            v36 := []int{fi10, -644, fi10, (fi10 % ((fi10) | 1))}
                                                                            v20 = 0.1
                                                                            if v11 {
                                                                                if v11 {
                                                                                    p0 += (p0 + p0)
                                                                                    var v37 uint16 = 65535
                                                                                    _ = v37
                                                                                }
                                                                            }
                                                                            _ = v36
                                                                        } else {
                                                                            v11 = v11
                                                                            v38 := []int{}
                                                                            p0 = (^(p0))
                                                                            _ = v38
                                                                        }
                                                                        var v39 uint8 = 0
                                                                        var v40 string = ("hello" + "foo bar")
                                                                        _ = v39
                                                                        _ = v40
                                                                    } else {
                                                                        if ((0.0 * -2.25) != 3.14159) {
                                                                            if p1 {
                                                                                var v41 string = "foo bar"
                                                                                if p1 {
                                                                                    if (92 != 0) {
                                                                                        var v42 int8 = 100
                                                                                        if false {
                                                                                            var v43 uint16 = 255
                                                                                            v44 := GenS2{F0: 0}
                                                                                            var v45 int = (fi10 >> ((701) & 63))
                                                                                            _ = v43
                                                                                            _ = v44
                                                                                            _ = v45
                                                                                        } else {
                                                                                            switch (v41 + v41) {
                                                                                            case "\n\t":
                                                                                                switch (fi10 % (((fi10 + (fi8 ^ fi8))) | 1)) {
                                                                                                case 128:
                                                                                                    var v46 int16 = 1000
                                                                                                    _ = v46
                                                                                                default:
                                                                                                    v47 := GenS5{F0: 255, F1: p0, F2: uint64(p0)}
                                                                                                    _ = v47
                                                                                                }
                                                                                            case "a":
                                                                                                var v48 uint32 = 1
                                                                                                _ = v48
                                                                                            }
                                                                                        }
                                                                                        _ = v42
                                                                                    } else {
                                                                                        fmt.Println(v34)
                                                                                    }
                                                                                    var v49 bool = (!((1 == 127) || (!(!(78 == 255)))))
                                                                                    v50 := GenS5{F0: (p0 - 1), F1: p0, F2: 42}
                                                                                    v49 = false
                                                                                    v34 = v11
                                                                                    _ = v49
                                                                                    _ = v50
                                                                                } else {
                                                                                    if (true || p1) {
                                                                                        if v11 {
                                                                                            v11 = (0.1 != ((3.14159 * 3.14159) * -2.25))
                                                                                            var v51 int = (fi10 << (((p0 & 42)) & 63))
                                                                                            v52 := GenS4{F0: uint32(p0), F1: false, F2: -1}
                                                                                            _ = v51
                                                                                            _ = v52
                                                                                        }
                                                                                        if v11 {
                                                                                            var v53 uint32 = 505
                                                                                            _ = v53
                                                                                        } else {
                                                                                            v54 := []int{fi8}
                                                                                            _ = v54
                                                                                        }
                                                                                    }
                                                                                    var v55 string = ((((v41 + "hello") + (v41 + "")) + (v41 + v41)) + v41)
                                                                                    _ = v55
                                                                                }
                                                                                _ = v41
                                                                            }
                                                                            var v56 uint16 = 32767
                                                                            _ = v56
                                                                        }
                                                                        p1 = v11
                                                                    }
                                                                    fmt.Println(fi8, v20)
                                                                } else {
                                                                    v57 := []int{}
                                                                    _ = v57
                                                                }
                                                            }
                                                            v58 := []int{(fi10 % ((fi10) | 1))}
                                                            _ = v58
                                                        } else {
                                                            v59 := []int{fi10, fi8}
                                                            if false {
                                                                v34 = true
                                                                fmt.Println(v34, fi10, v34)
                                                            } else {
                                                                v34 = ((-407 == -32768) && v34)
                                                                if (true || false) {
                                                                    var v60 uint = 1
                                                                    _ = v60
                                                                } else {
                                                                    v61 := GenS6{F0: "", F1: 0}
                                                                    _ = v61
                                                                }
                                                                p1 = true
                                                                switch (fi8 / (((fi8 ^ (-(fi10)))) | 1)) {
                                                                case 4294967295:
                                                                    v62 := GenS7{F0: "foo bar", F1: 42}
                                                                    _ = v62
                                                                }
                                                            }
                                                            _ = v59
                                                        }
                                                    case "\n\t":
                                                        v63 := []int{}
                                                        _ = v63
                                                    }
                                                }
                                                if ("a" < (("foo bar" + ((("foo bar" + ((("hello" + "\n\t") + ("a" + "\n\t")) + "\n\t")) + "a") + "a")) + "\n\t")) {
                                                    fmt.Println(p1, v20)
                                                    var v64 float64 = 1e10
                                                    var v65 uint8 = 42
                                                    _ = v64
                                                    _ = v65
                                                }
                                                if p1 {
                                                    var v66 uint = p0
                                                    _ = v66
                                                } else {
                                                    if (v11 || (!true)) {
                                                        if v11 {
                                                            if p1 {
                                                                fmt.Println(p1)
                                                                switch "a" {
                                                                case "":
                                                                    p0 += 42
                                                                }
                                                                v67 := GenS5{F0: (p0 * 191), F1: (p0 + (p0 & p0)), F2: 18446744073709551615}
                                                                _ = v67
                                                            }
                                                        } else {
                                                            v20 += 1e10
                                                        }
                                                        if v34 {
                                                            v11 = false
                                                            switch 9223372036854775807 {
                                                            case 32767:
                                                                v11 = (p1 || v11)
                                                            default:
                                                                v68 := []int{-475}
                                                                v69 := GenS8{F0: (v11 && (-312 >= 1))}
                                                                p0 += (p0 / ((730) | 1))
                                                                v70 := GenS4{F0: 747, F1: true, F2: -1}
                                                                _ = v68
                                                                _ = v69
                                                                _ = v70
                                                            }
                                                            p0 = 2
                                                        }
                                                    }
                                                    var v71 uint32 = 1
                                                    var v72 int8 = 127
                                                    _ = v71
                                                    _ = v72
                                                }
                                                v73 := GenS2{F0: 949}
                                                var v74 float64 = float64(((v20 - (0.1 + v20)) + ((3.14159 * v20) * v20)))
                                                _ = v34
                                                _ = v73
                                                _ = v74
                                            }
                                            var v75 int = (fi10 % ((0) | 1))
                                            v76 := GenS9{F0: (("foo bar" + "foo bar") + "hello"), F1: 0}
                                            switch "" {
                                            case "\n\t":
                                                switch "hello" {
                                                case "":
                                                    v77 := []int{-25, 2147483647, (fi10 % (((fi10 / ((fi8) | 1))) | 1))}
                                                    _ = v77
                                                case "a":
                                                    v20 -= -2.25
                                                default:
                                                    v78 := GenS10{F0: 0, F1: 2, F2: fi8}
                                                    _ = v78
                                                }
                                            case "":
                                                if (!p1) {
                                                    var v79 string = "a"
                                                    v75 -= fi10
                                                    var v80 int16 = -120
                                                    var v81 int32 = 0
                                                    v80 = v80
                                                    _ = v79
                                                    _ = v80
                                                    _ = v81
                                                } else {
                                                    v82 := []int{(v75 & (fi10 % ((127) | 1))), (v75 | (fi10 << (((p0 * p0)) & 63))), fi8}
                                                    _ = v82
                                                }
                                            }
                                            v11 = (v11 || ((42 == -128) || (0 >= 4294967295)))
                                            var v83 uint = p0
                                            _ = v75
                                            _ = v76
                                            _ = v83
                                        }
                                    } else {
                                        p0 = (p0 & 4294967295)
                                        fmt.Println(fi10)
                                        fmt.Println(fi10, fi8)
                                    }
                                case "foo bar":
                                    v84 := GenS5{F0: 480, F1: 42, F2: 341}
                                    _ = v84
                                case "":
                                    v20 -= (v20 + v20)
                                }
                                if ((731 > 42) || ((42 >= -53) && (!p1))) {
                                    var v85 int = fi10
                                    p1 = v11
                                    _ = v85
                                } else {
                                    switch "foo bar" {
                                    case "a":
                                        v86 := []int{42, (fi8 / (((^(fi8))) | 1)), 913, fi8}
                                        _ = v86
                                    case "foo bar":
                                        v11 = v11
                                    case "\n\t":
                                        var v87 int = (fi10 % (((fi10 + -792)) | 1))
                                        _ = v87
                                    default:
                                        switch "foo bar" {
                                        case "foo bar":
                                            if p1 {
                                                v88 := GenS2{F0: 65535}
                                                var v89 uint16 = 0
                                                _ = v88
                                                _ = v89
                                            }
                                        case "hello":
                                            p1 = ((((((-2.25 / (3.14159 + 3.14159)) * 0.0) + 1e10) / 1.5) / (float64((v20 / 1.0)) * (1.5 * 1e10))) != 0.0)
                                        case "":
                                            v90 := []int{(fi8 | (^(fi8))), 4294967295, (^(fi8))}
                                            _ = v90
                                        default:
                                            v91 := []int{(fi10 - fi10), fi8}
                                            _ = v91
                                        }
                                    }
                                    v11 = true
                                    v20 = 1.5
                                }
                                v92 := []int{32767, fi8, (fi8 | 128), (-(fi10))}
                                _ = v92
                            }
                            _ = v20
                        } else {
                            v93 := []int{-2147483648}
                            fmt.Println(p0, p0, fi10)
                            v93 = append(v93, (fi8 * fi8))
                            _ = v93
                        }
                    } else {
                        var v94 uint32 = 52
                        if (!p1) {
                            if v11 {
                                fmt.Println(fi10)
                                v95 := []int{fi8, (fi10 % (((fi10 ^ fi10)) | 1))}
                                _ = v95
                            }
                            v96 := GenS9{F0: (("a" + "") + ("foo bar" + "")), F1: 1}
                            v97 := GenS4{F0: 2147483647, F1: (-1 != 111), F2: 42}
                            _ = v96
                            _ = v97
                        }
                        var v98 int16 = 0
                        _ = v94
                        _ = v98
                    }
                    _ = v18
                    _ = v19
                }
                v11 = p1
                _ = fi10
                _ = v11
            }
        }
        _ = fi8
    }
    var v99 uint8 = 12
    for fi100 := 0; fi100 < 8; fi100++ {
        var v101 uint16 = 1
        var v102 string = (("" + "\n\t") + ("" + "a"))
        switch "\n\t" {
        case "foo bar":
            var v103 float32 = 1.5
            _ = v103
        }
        _ = fi100
        _ = v101
        _ = v102
    }
    var v104 string = ("\n\t" + "foo bar")
    _ = p0
    _ = p1
    _ = v104
    _ = v7
    _ = v99
    return 0
}
func genFn1(p0 uint, p1 int16, p2 uint) uint64 {
    if true {
        v105 := []int{32768}
        fmt.Println(p2)
        var v106 float32 = 1.5
        _ = v105
        _ = v106
    }
    v107 := GenS11{F0: 54, F1: (p2 / ((880) | 1)), F2: false}
    v108 := []int{}
    v109 := GenS4{F0: 939, F1: ((1e10 != 1.5) && (!true)), F2: 9223372036854775807}
    v110 := GenS11{F0: 54, F1: (^(p2)), F2: true}
    _ = p0
    _ = p1
    _ = p2
    _ = v107
    _ = v108
    _ = v109
    _ = v110
    return 1
}

func main() {
    var v1 uint8 = 0
    v2 := make(map[string]bool)
    var v3 int = 1
    var v4 int32 = -1
    var v5 string = ("foo bar" + "hello")
    v6 := make(map[string]int)
    v5 = v5
    var v111 uint8 = (v1 % ((205) | 1))
    var v112 int32 = (v4 ^ v4)
    var v113 float32 = 0.1
    v114 := GenS12{F0: 27}
    for fi115 := 0; fi115 < 13; fi115++ {
        v116 := GenS13{F0: -1}
        _ = fi115
        _ = v116
    }
    v3 += (v3 - v3)
    v117 := GenS4{F0: 4294967295, F1: true, F2: -297}
    fmt.Println("v1:", v1)
    fmt.Println("v111:", v111)
    fmt.Println("v112:", v112)
    fmt.Println("v113:", v113)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v111
    _ = v112
    _ = v113
    _ = v114
    _ = v117
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
