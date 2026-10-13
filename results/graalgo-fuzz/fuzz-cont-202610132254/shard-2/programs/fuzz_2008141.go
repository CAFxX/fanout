package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 int64; F1 int8; F2 uint }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint16; F1 uint8; F2 int16 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int16; F1 string; F2 uint64 }
type GenS8 struct { F0 int32; F1 int8 }
type GenS9 struct { F0 uint64 }
func genFn0() uint16 {
    if true {
        var v4 int32 = -117
        _ = v4
    } else {
        var v5 int16 = -1
        _ = v5
    }
    for fi6 := 0; fi6 < 9; fi6++ {
        fmt.Println(fi6)
        if false {
            var v7 float64 = 0.1
            _ = v7
        } else {
            for fi8 := 0; fi8 < 8; fi8++ {
                var v9 uint64 = 1
                var v10 uint8 = 42
                _ = fi8
                _ = v10
                _ = v9
            }
        }
        switch (fi6 * fi6) {
        case 1:
            var v11 bool = ((fi6 - 128) > (fi6 >> ((806) & 63)))
            _ = v11
        case 128:
            var v12 int16 = -32768
            _ = v12
        case -9223372036854775808:
            for fi13 := 0; fi13 < 7; fi13++ {
                v14 := GenS1{F0: 19, F1: -701}
                _ = fi13
                _ = v14
            }
        default:
            if false {
                var v15 bool = (0 != 209)
                fmt.Println(v15, v15)
                _ = v15
            }
        }
        _ = fi6
    }
    v16 := GenS1{F0: 42, F1: -420}
    var v17 int32 = -2147483648
    if false {
        if false {
            v17 = v17
            fmt.Println(v17)
            v17 -= -2147483648
        }
        fmt.Println(v17)
    } else {
        if ((-128 == -1) || (!true)) {
            for fi18 := 0; fi18 < 8; fi18++ {
                if (!false) {
                    if true {
                        fmt.Println(v17, v17)
                        switch "hello" {
                        case "a":
                            v17 = v17
                        default:
                            v19 := []int{}
                            _ = v19
                        }
                    } else {
                        var v20 uint16 = 0
                        v17 = -1
                        switch "" {
                        case "\n\t":
                            v20 -= (v20 - v20)
                        case "foo bar":
                            fmt.Println(v17, fi18)
                        case "hello":
                            v21 := []int{}
                            _ = v21
                        default:
                            v22 := GenS1{F0: -862, F1: 32767}
                            _ = v22
                        }
                        _ = v20
                    }
                    v17 = -183
                }
                if true {
                    v17 = v17
                    if true {
                        switch (fi18 | (fi18 ^ fi18)) {
                        case -158:
                            if (((fi18 << ((18446744073709551615) & 63)) > (fi18 + 909)) || false) {
                                v17 += (v17 - v17)
                                for fi23 := 0; fi23 < 17; fi23++ {
                                    if (0 == 128) {
                                        if ((2147483648 > 300) && false) {
                                            v17 -= v17
                                            var v24 uint16 = 255
                                            fmt.Println(fi23)
                                            _ = v24
                                        }
                                        v25 := []int{363}
                                        switch "a" {
                                        case "":
                                            v26 := GenS2{F0: 9223372036854775807, F1: 0, F2: 42}
                                            _ = v26
                                        case "foo bar":
                                            if false {
                                                var v27 float32 = 1.5
                                                var v28 float64 = 0.1
                                                v25 = append(v25, fi23)
                                                _ = v27
                                                _ = v28
                                            }
                                        }
                                        _ = v25
                                    }
                                    if ("hello" != "a") {
                                        if (!true) {
                                            v29 := []int{(^(fi23)), (^(fi23)), fi23, fi23}
                                            if true {
                                                var v30 int64 = 42
                                                _ = v30
                                            }
                                            var v31 int32 = (v17 % (((v17 % ((v17) | 1))) | 1))
                                            _ = v29
                                            _ = v31
                                        } else {
                                            v17 += (v17 + (v17 << ((771) & 63)))
                                        }
                                        var v32 string = "a"
                                        var v33 uint64 = 1
                                        v34 := GenS3{F0: (((v32 + (("\n\t" + "hello") + v32)) + ("hello" + "a")) + v32)}
                                        var v35 int64 = -1
                                        _ = v32
                                        _ = v33
                                        _ = v34
                                        _ = v35
                                    }
                                    _ = fi23
                                }
                            } else {
                                fmt.Println(v17, fi18)
                                var v36 float32 = ((0.0 * 1.5) / 1.0)
                                _ = v36
                            }
                        case 832:
                            for fi37 := 0; fi37 < 9; fi37++ {
                                v17 = (v17 | (v17 % (((v17 << ((137) & 63))) | 1)))
                                switch ("\n\t" + "a") {
                                case "foo bar":
                                    var v38 int32 = v17
                                    v39 := []int{fi18, fi18}
                                    v17 = (v17 | v38)
                                    _ = v38
                                    _ = v39
                                case "hello":
                                    v17 += -2147483648
                                default:
                                    v17 += -2147483648
                                }
                                _ = fi37
                            }
                        case -634:
                            var v40 float64 = (-2.25 + float64(0.1))
                            _ = v40
                        }
                    } else {
                        v41 := []int{fi18, (^(fi18)), fi18, fi18}
                        _ = v41
                    }
                }
                var v42 int16 = -1
                _ = fi18
                _ = v42
            }
        }
    }
    _ = v16
    _ = v17
    return 65535
}
func genFn1(p0 int8, p1 uint8, p2 int) uint {
    v43 := GenS1{F0: -874, F1: 0}
    p2 = (p2 << ((875) & 63))
    p1 = 127
    _ = p0
    _ = p1
    _ = p2
    _ = v43
    return 609
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 int8 = 1
    var v3 int8 = (v2 >> ((42) & 63))
    var v44 uint16 = 2
    if ((true || ((1.5 != float64(-2.25)) || true)) || (!true)) {
        for fi45 := 0; fi45 < 14; fi45++ {
            if (false || true) {
                if (495 != 65535) {
                    var v46 uint16 = v44
                    _ = v46
                } else {
                    var v47 string = "hello"
                    v1 -= (v1 << ((18446744073709551615) & 63))
                    v48 := GenS1{F0: 1, F1: 1000}
                    _ = v47
                    _ = v48
                }
            } else {
                v49 := GenS2{F0: 0, F1: v3, F2: 145}
                v50 := GenS1{F0: -224, F1: 321}
                v51 := GenS2{F0: (v1 - -9223372036854775808), F1: 124, F2: 65535}
                _ = v49
                _ = v50
                _ = v51
            }
            for fi52 := 0; fi52 < 11; fi52++ {
                if (true && false) {
                    if true {
                        v53 := GenS1{F0: 42, F1: -562}
                        var v54 float32 = 1e10
                        fmt.Println(fi52, v1)
                        _ = v53
                        _ = v54
                    } else {
                        var v55 uint64 = 303
                        var v56 bool = true
                        v56 = (v56 || (((v2 - 42) != v3) || ((v44 - 909) < (v44 << ((0) & 63)))))
                        if v56 {
                            if (!(v56 && (((float64(-2.25) + (0.1 + -2.25)) - 1.5) != 1e10))) {
                                v57 := []int{100, (^(fi45)), (fi52 + (fi52 % (((^(fi52))) | 1)))}
                                v56 = (v56 || v56)
                                v56 = v56
                                if ((v2 % ((v3) | 1)) < (v2 & (v2 * v3))) {
                                    switch fi45 {
                                    case 42:
                                        var v58 uint8 = 128
                                        _ = v58
                                    case 0:
                                        v56 = (v56 || v56)
                                    default:
                                        if len(v57) >= 2 {
                                            v57 = v57[1:len(v57)]
                                        }
                                    }
                                    var v59 uint16 = genFn0()
                                    _ = v59
                                } else {
                                    v60 := GenS1{F0: 42, F1: -32768}
                                    if v56 {
                                        v3 = (v3 << ((0) & 63))
                                        v61 := GenS3{F0: ("a" + (("\n\t" + "a") + "\n\t"))}
                                        _ = v61
                                    }
                                    _ = v60
                                }
                                if (412 >= 336) {
                                    fmt.Println(v44, fi52)
                                    if len(v57) >= 2 {
                                        v57 = v57[1:len(v57)]
                                    }
                                    if len(v57) > 0 {
                                        v57[0] = fi45
                                    }
                                }
                                var v62 uint64 = v55
                                _ = len(v57)
                                _ = v57
                                _ = v62
                            }
                        }
                        fmt.Println(v55)
                        _ = v55
                        _ = v56
                    }
                    var v63 uint64 = 1
                    switch 42 {
                    case 323:
                        if false {
                            v64 := []int{int(v63)}
                            var v65 uint64 = v63
                            _ = v64
                            _ = v65
                        } else {
                            v63 -= 18446744073709551615
                        }
                    case 9223372036854775807:
                        v66 := GenS1{F0: -32768, F1: int16(v63)}
                        _ = v66
                    default:
                        v63 += (v63 ^ 18446744073709551615)
                    }
                    _ = v63
                }
                switch fi45 {
                case 929:
                    v67 := []int{}
                    _ = v67
                case -15:
                    var v68 int = -2147483648
                    _ = v68
                default:
                    v3 += v3
                }
                v69 := GenS1{F0: -32768, F1: int16(fi52)}
                _ = fi52
                _ = v69
            }
            switch "" {
            case "":
                v3 = v2
            case "\n\t":
                switch ("foo bar" + (("foo bar" + ("" + "\n\t")) + ("" + ("a" + ((("\n\t" + ("a" + "hello")) + "") + "foo bar"))))) {
                case "hello":
                    switch "hello" {
                    case "a":
                        for fi70 := 0; fi70 < 19; fi70++ {
                            if true {
                                var v71 uint = 42
                                switch fi70 {
                                case -2147483648:
                                    v72 := []int{}
                                    _ = v72
                                case -1:
                                    v3 = (v2 / (((-(v2))) | 1))
                                case 9223372036854775807:
                                    v73 := GenS4{F0: 207}
                                    _ = v73
                                }
                                v74 := GenS3{F0: ""}
                                _ = v71
                                _ = v74
                            }
                            fmt.Println(fi70, v2)
                            v1 = v1
                            var v75 int16 = 32767
                            if true {
                                var v76 int = (fi70 * fi45)
                                var v77 int32 = 377
                                if false {
                                    if true {
                                        v77 += 2147483647
                                        v77 = v77
                                        if (!true) {
                                            var v78 int = (v76 << ((42) & 63))
                                            switch v76 {
                                            case 9223372036854775807:
                                                v79 := []int{fi45, (v76 * (fi70 >> ((591) & 63)))}
                                                var v80 int = (-(v76))
                                                _ = v79
                                                _ = v80
                                            case 190:
                                                v77 += -2147483648
                                                v81 := GenS5{F0: v44, F1: 127, F2: v75}
                                                _ = v81
                                            default:
                                                var v82 int = int(v44)
                                                _ = v82
                                            }
                                            if false {
                                                v83 := GenS6{F0: 868}
                                                v77 -= v77
                                                v84 := []int{42, (fi70 >> ((470) & 63)), (v76 | fi45), 98}
                                                if len(v84) > 0 {
                                                    v84[0] = fi45
                                                }
                                                fmt.Println(fi45)
                                                fmt.Println(v76, v77)
                                                _ = v83
                                                _ = v84
                                            } else {
                                                var v85 int8 = v2
                                                v76 += (-(v78))
                                                _ = v85
                                            }
                                            _ = v78
                                        }
                                    }
                                    if true {
                                        var v86 string = "\n\t"
                                        _ = v86
                                    } else {
                                        var v87 int16 = (v75 | v75)
                                        var v88 uint8 = uint8(v44)
                                        if true {
                                            var v89 string = ("" + "foo bar")
                                            v90 := GenS4{F0: (v88 << ((0) & 63))}
                                            if true {
                                                if true {
                                                    var v91 int32 = (^(v77))
                                                    _ = v91
                                                }
                                                v92 := []int{fi70, (fi70 << ((0) & 63)), 127, -1}
                                                _ = v92
                                            }
                                            _ = v89
                                            _ = v90
                                        } else {
                                            v44 -= v44
                                        }
                                        var v93 bool = (false || (((0.0 * (1.5 / -2.25)) + ((3.14159 / (3.14159 / 3.14159)) * float32(0.1))) >= (((1e10 + -2.25) - 3.14159) / 0.1)))
                                        v94 := []int{(v76 % ((32768) | 1))}
                                        _ = v87
                                        _ = v88
                                        _ = v93
                                        _ = v94
                                    }
                                    v1 = v1
                                }
                                _ = v76
                                _ = v77
                            }
                            v95 := GenS7{F0: v75, F1: "a", F2: 42}
                            _ = fi70
                            _ = v75
                            _ = v95
                        }
                    default:
                        var v96 int64 = (v1 - 127)
                        _ = v96
                    }
                case "a":
                    fmt.Println(v44)
                default:
                    for fi97 := 0; fi97 < 2; fi97++ {
                        var v98 uint = genFn1((v2 ^ v2), 255, (^(fi45)))
                        fmt.Println(v44, v1, fi45)
                        _ = fi97
                        _ = v98
                    }
                }
            case "foo bar":
                v2 += (v3 / ((v3) | 1))
            default:
                switch fi45 {
                case 42:
                    v2 = v3
                case -852:
                    v99 := []int{(fi45 ^ 42), (fi45 & int(v1))}
                    _ = v99
                case 127:
                    for fi100 := 0; fi100 < 6; fi100++ {
                        v101 := []int{(fi45 / (((fi45 / ((589) | 1))) | 1)), 2147483647, 4294967295, fi45}
                        v102 := GenS8{F0: 42, F1: (^(v2))}
                        var v103 string = "hello"
                        if (32767 >= 974) {
                            v44 += v44
                        } else {
                            if true {
                                v101 = append(v101, fi100)
                                if (-657 <= 32767) {
                                    var v104 string = v103
                                    _ = v104
                                }
                                var v105 float64 = float64(0.0)
                                _ = v105
                            } else {
                                var v106 string = ""
                                _ = len(v101)
                                v3 = 125
                                _ = v106
                            }
                        }
                        if false {
                            fmt.Println(v2)
                            fmt.Println(v103, v3)
                            v103 = ("foo bar" + ((("a" + v103) + (v103 + v103)) + (v103 + ("" + (((((v103 + v103) + "foo bar") + ("foo bar" + v103)) + ((v103 + v103) + v103)) + v103)))))
                        } else {
                            switch int(v2) {
                            case 128:
                                v2 += (v3 * (-(v3)))
                            case 9223372036854775807:
                                fmt.Println(v3, v2, fi45)
                            default:
                                if (0 < 176) {
                                    if ((v1 - (v1 >> ((0) & 63))) > 1) {
                                        v3 = 127
                                        var v107 uint8 = 42
                                        var v108 uint8 = (v107 | 255)
                                        _ = v107
                                        _ = v108
                                    } else {
                                        v2 = v2
                                        fmt.Println(v3, v2)
                                    }
                                    var v109 bool = true
                                    _ = v109
                                }
                            }
                        }
                        _ = fi100
                        _ = v101
                        _ = v102
                        _ = v103
                    }
                default:
                    v110 := GenS3{F0: "a"}
                    _ = v110
                }
            }
            _ = fi45
        }
        var v111 float64 = 1.5
        _ = v111
    } else {
        v44 = v44
        v44 += (v44 * v44)
    }
    v112 := GenS9{F0: 906}
    v113 := []int{128, -2147483648}
    for fi114 := 0; fi114 < 8; fi114++ {
        var v115 int64 = v1
        _ = fi114
        _ = v115
    }
    if len(v113) > 0 {
        v113[0] = 4294967295
    }
    var v116 uint16 = v44
    v3 += v3
    v117 := GenS8{F0: 2147483647, F1: (v2 >> ((42) & 63))}
    v116 = (v44 + 1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v113:", len(v113), v113)
    fmt.Println("v116:", v116)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v112
    _ = v113
    _ = v116
    _ = v117
    _ = v2
    _ = v3
    _ = v44
}
