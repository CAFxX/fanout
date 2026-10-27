package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 int64; F1 int8; F2 uint32 }
type GenS3 struct { F0 int; F1 int8; F2 uint32 }
type GenS4 struct { F0 uint64; F1 uint8 }
type GenS5 struct { F0 string; F1 string; F2 uint64 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 string; F1 uint }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 int8; F1 int16 }
type GenS11 struct { F0 uint16; F1 uint32 }
type GenS12 struct { F0 uint64; F1 int64 }
type GenS13 struct { F0 bool; F1 uint16 }
type GenS14 struct { F0 int8 }

func main() {
    var v1 int = 1
    var v2 int32 = 1
    v3 := make(map[int]int)
    var v4 bool = true
    for fi5 := 0; fi5 < 12; fi5++ {
        for fi6 := 0; fi6 < 1; fi6++ {
            fmt.Println(v2)
            if v4 {
                fmt.Println(v2, fi6, v4)
                v7 := []int{(v1 & fi5), (^(fi6)), v1}
                _ = v7
            } else {
                if (v4 || (4294967295 > 531)) {
                    if (v4 || v4) {
                        v8 := v3[(v1 >> ((2) & 63))]
                        fmt.Println(v8, v1)
                        _ = v8
                    } else {
                        if v4 {
                            var v9 float64 = float64(0.0)
                            _ = v9
                        } else {
                            switch (fi6 & 42) {
                            case -2147483648:
                                var v10 uint8 = 197
                                _ = v10
                            case 207:
                                switch "" {
                                case "foo bar":
                                    v11 := GenS1{F0: 127, F1: 602}
                                    _ = v11
                                case "a":
                                    if (v4 && (v2 <= v2)) {
                                        switch (^(fi6)) {
                                        case -9223372036854775808:
                                            v12 := GenS1{F0: 42, F1: -32768}
                                            _ = v12
                                        }
                                        switch (v1 - fi5) {
                                        case 42:
                                            v13 := GenS1{F0: 42, F1: -656}
                                            _ = v13
                                        default:
                                            v14 := []int{(v1 ^ v1), v1, 4294967295, (fi6 | (fi5 ^ (v1 / ((127) | 1))))}
                                            _ = v14
                                        }
                                        v15 := GenS1{F0: 2, F1: 42}
                                        var v16 float64 = 1.5
                                        v4 = v4
                                        _ = v15
                                        _ = v16
                                    } else {
                                        switch (v1 & -251) {
                                        case -536:
                                            if ((1000 <= 42) || v4) {
                                                v4 = ((4294967295 >= 860) && false)
                                                var v17 int32 = v2
                                                _ = v17
                                            } else {
                                                var v18 uint8 = 42
                                                v19 := GenS1{F0: (v18 ^ (^(v18))), F1: -38}
                                                v18 -= v18
                                                v4 = v4
                                                v1 -= fi6
                                                _ = v18
                                                _ = v19
                                            }
                                        case 109:
                                            v2 = 0
                                        case 127:
                                            if (3.14159 <= 1e10) {
                                                v2 += (v2 * (v2 << ((42) & 63)))
                                                if v4 {
                                                    v20 := []int{(fi5 * 42), fi5, (fi5 / (((fi6 - -2147483648)) | 1))}
                                                    v21 := GenS2{F0: int64(fi5), F1: 0, F2: 2147483647}
                                                    var v22 int8 = int8(fi5)
                                                    _ = v20
                                                    _ = v21
                                                    _ = v22
                                                }
                                                v4 = true
                                            } else {
                                                if (35 >= 547) {
                                                    fmt.Println(v2, fi5)
                                                    fmt.Println(fi6)
                                                    var v23 bool = v4
                                                    fmt.Println(fi5)
                                                    v4 = false
                                                    _ = v23
                                                } else {
                                                    if v4 {
                                                        if v4 {
                                                            v24 := GenS3{F0: v1, F1: 121, F2: 65535}
                                                            _ = v24
                                                        }
                                                    } else {
                                                        v25 := GenS1{F0: 1, F1: 32767}
                                                        var v26 bool = ("\n\t" >= "a")
                                                        if v4 {
                                                            fmt.Println(v4, v26)
                                                            v2 = -662
                                                        } else {
                                                            if (fi6 > fi5) {
                                                                if ((v4 || v4) && ((1e10 * 1e10) <= 3.14159)) {
                                                                    delete(v3, (fi5 / ((32768) | 1)))
                                                                    v2 += (v2 >> ((42) & 63))
                                                                    v2 += (-(v2))
                                                                }
                                                                v27 := []int{(fi6 >> ((361) & 63))}
                                                                if len(v27) >= 2 {
                                                                    v27 = v27[1:len(v27)]
                                                                }
                                                                _ = v27
                                                            }
                                                        }
                                                        _ = v25
                                                        _ = v26
                                                    }
                                                }
                                                v4 = ((!((("\n\t" + ("foo bar" + ("\n\t" + ""))) != "") && v4)) || (v4 || true))
                                                v28 := GenS3{F0: (fi5 + (fi5 >> ((163) & 63))), F1: -1, F2: 2147483647}
                                                _ = v28
                                            }
                                        }
                                    }
                                case "\n\t":
                                    var v29 uint64 = 863
                                    _ = v29
                                default:
                                    var v30 uint64 = 451
                                    v31 := []int{fi6}
                                    if (-32768 == 0) {
                                        var v32 uint16 = 65535
                                        _ = v32
                                    } else {
                                        var v33 bool = (v4 || true)
                                        switch "\n\t" {
                                        case "foo bar":
                                            var v34 uint = 2
                                            _ = v34
                                        case "":
                                            if true {
                                                v35 := GenS1{F0: 22, F1: 1000}
                                                v31 = append(v31, (-(fi5)))
                                                var v36 string = ("a" + "foo bar")
                                                _ = v35
                                                _ = v36
                                            } else {
                                                v31 = append(v31, (fi6 << ((4294967295) & 63)))
                                                v37 := GenS2{F0: 9223372036854775807, F1: 29, F2: uint32(v30)}
                                                v30 -= 0
                                                _ = v37
                                            }
                                        case "hello":
                                            var v38 string = (("\n\t" + ("\n\t" + "")) + "")
                                            _ = v38
                                        }
                                        if (!true) {
                                            var v39 uint8 = 42
                                            fmt.Println(v1, v4)
                                            _ = v39
                                        }
                                        _ = v33
                                    }
                                    _ = v30
                                    _ = v31
                                }
                            case 2147483647:
                                switch "foo bar" {
                                case "\n\t":
                                    v2 -= (-(v2))
                                }
                            }
                            v40 := GenS2{F0: 0, F1: 0, F2: 2}
                            switch 655 {
                            case 656:
                                var v41 int8 = 42
                                _ = v41
                            case 4294967295:
                                if (v2 == v2) {
                                    fmt.Println(v4, fi5, fi6)
                                }
                            default:
                                v1 = (fi6 + v1)
                            }
                            _ = v40
                        }
                    }
                    var v42 float32 = (1e10 + ((((-2.25 * 0.0) - 1e10) * (3.14159 - 3.14159)) * 0.0))
                    v43 := []int{}
                    _ = v42
                    _ = v43
                } else {
                    if true {
                        if v4 {
                            var v44 uint = 255
                            switch int(v2) {
                            case 42:
                                v1 = fi6
                            case -192:
                                delete(v3, fi5)
                            default:
                                var v45 float32 = -2.25
                                _ = v45
                            }
                            v2 += v2
                            _ = v44
                        } else {
                            v4 = true
                            if (!(false && (826 == 568))) {
                                v46 := GenS3{F0: fi6, F1: -128, F2: 2147483648}
                                fmt.Println(fi5)
                                v47 := GenS4{F0: 1, F1: 194}
                                _ = v46
                                _ = v47
                            }
                        }
                        switch "\n\t" {
                        case "\n\t":
                            v48 := GenS5{F0: "hello", F1: "foo bar", F2: 256}
                            _ = v48
                        case "a":
                            v4 = ((-1 <= -1) || (28 <= 42))
                        }
                        var v49 int = (fi6 + fi5)
                        _ = v49
                    } else {
                        var v50 int = v1
                        var v51 uint = 419
                        _ = v50
                        _ = v51
                    }
                }
                fmt.Println(v2, fi5)
                var v52 uint = 2
                v4 = ((0.0 * -2.25) == 0.0)
                v2 = v2
                v1 += 325
                _ = v52
            }
            var v53 bool = ("" > ("\n\t" + ("a" + "foo bar")))
            if (2147483648 > 320) {
                if v53 {
                    switch ("\n\t" + ("foo bar" + "")) {
                    case "":
                        v2 = -603
                    case "foo bar":
                        switch 32768 {
                        case 42:
                            switch v1 {
                            case -1:
                                v4 = (v53 || false)
                            default:
                                if false {
                                    v4 = v53
                                    switch "hello" {
                                    case "hello":
                                        fmt.Println(v1, v4, v1)
                                        var v54 string = ((("hello" + "hello") + "a") + "")
                                        _ = v54
                                    }
                                    switch "hello" {
                                    case "\n\t":
                                        switch ("" + "hello") {
                                        case "foo bar":
                                            v2 = (v2 % (((v2 + v2)) | 1))
                                        }
                                    }
                                    fmt.Println(v4, fi5)
                                    v1 = fi5
                                    v55 := v3[127]
                                    _ = v55
                                }
                            }
                        default:
                            if v4 {
                                var v56 uint = 711
                                _ = v56
                            } else {
                                v57 := GenS6{F0: (2147483647 > 809)}
                                if v53 {
                                    var v58 int8 = -120
                                    _ = v58
                                }
                                _ = v57
                            }
                        }
                        v59 := GenS3{F0: fi6, F1: 117, F2: 403}
                        if v4 {
                            v60 := GenS7{F0: 1}
                            if (v53 && v53) {
                                if (true || v53) {
                                    if v4 {
                                        switch "hello" {
                                        case "":
                                            var v61 uint = 708
                                            _ = v61
                                        case "a":
                                            v53 = ((127 < 128) && ((v2 / ((v2) | 1)) < v2))
                                        case "\n\t":
                                            var v62 uint64 = 914
                                            _ = v62
                                        }
                                    } else {
                                        var v63 int32 = 42
                                        if v4 {
                                            v64 := []int{}
                                            if false {
                                                var v65 int16 = int16(v63)
                                                if ((926 > v65) || v53) {
                                                    var v66 bool = v4
                                                    v2 += (v63 % ((0) | 1))
                                                    if len(v64) > 0 {
                                                        v64[0] = (fi6 >> ((89) & 63))
                                                    }
                                                    _ = v66
                                                } else {
                                                    v4 = true
                                                }
                                                _ = v65
                                            } else {
                                                v67 := GenS6{F0: ((((((-2.25 / (0.1 + float64(0.1))) / 0.1) + (float64(1.5) / 1.5)) + 3.14159) * ((0.0 - ((0.1 / float64(((float32(0.1) / float32(0.1)) * 1e10))) - 1.5)) * (0.1 + 3.14159))) <= (1e10 / 1.5))}
                                                fmt.Println(v2, v53, v1)
                                                if len(v64) >= 2 {
                                                    v64 = v64[1:len(v64)]
                                                }
                                                v3[fi6] = 32768
                                                _ = v67
                                            }
                                            v68 := GenS4{F0: 0, F1: 1}
                                            _ = v64
                                            _ = v68
                                        }
                                        v53 = (!(v4 || true))
                                        _ = v63
                                    }
                                    if (!v4) {
                                        v69 := []int{(v1 / ((v1) | 1)), fi6, (fi6 - (-(fi6)))}
                                        var v70 int64 = 9223372036854775807
                                        v4 = (182 > 4294967295)
                                        _ = v69
                                        _ = v70
                                    }
                                    fmt.Println(v4, v4, v2)
                                }
                                v71 := GenS7{F0: -32768}
                                if v4 {
                                    v72 := []int{}
                                    v1 -= -9223372036854775808
                                    var v73 bool = true
                                    _ = v72
                                    _ = v73
                                } else {
                                    var v74 int8 = -106
                                    _ = v74
                                }
                                if (-443 >= 513) {
                                    if (0.1 <= float64(((((float32(0.1) * (float32(0.1) + 0.0)) * 1.5) * (3.14159 * float32((1.5 / 1e10)))) * -2.25))) {
                                        if v4 {
                                            fmt.Println(v53, v1, v53)
                                            var v75 float32 = 1e10
                                            _ = v75
                                        }
                                        v2 += (v2 * -663)
                                        v76 := []int{(^(fi6)), (fi5 / ((-1) | 1)), (fi6 % (((v1 + (fi6 | fi5))) | 1)), fi5}
                                        _ = v76
                                    }
                                }
                                if v4 {
                                    fmt.Println(v53, v1)
                                    v3[fi5] = 4294967295
                                } else {
                                    var v77 int64 = 42
                                    _ = v77
                                }
                                v78 := v3[-9223372036854775808]
                                v78 += v78
                                _ = v71
                                _ = v78
                            }
                            _ = v60
                        } else {
                            v79 := []int{(fi6 + v1), -909, -1, 128}
                            _ = v79
                        }
                        _ = v59
                    case "\n\t":
                        switch -1 {
                        case 128:
                            v4 = v4
                        case 4294967295:
                            if false {
                                v80 := GenS7{F0: 1000}
                                var v81 uint32 = 65535
                                v53 = (("a" + "foo bar") != "hello")
                                switch (v1 - fi5) {
                                case 32768:
                                    v1 -= 127
                                case -558:
                                    v81 = (v81 ^ (v81 | v81))
                                case 273:
                                    var v82 int = -9223372036854775808
                                    _ = v82
                                }
                                v4 = false
                                v83 := GenS1{F0: 93, F1: 1000}
                                _ = v80
                                _ = v81
                                _ = v83
                            } else {
                                v1 = 4294967295
                            }
                        }
                        v1 -= 32768
                        v53 = (42 == -219)
                    default:
                        var v84 int8 = -128
                        _ = v84
                    }
                    switch (fi6 - (fi5 >> ((797) & 63))) {
                    case 982:
                        fmt.Println(v2)
                    default:
                        if ((-2.25 / -2.25) == -2.25) {
                            fmt.Println(v1, v4, v53)
                            v2 -= v2
                        } else {
                            delete(v3, v1)
                            if (false && v53) {
                                fmt.Println(fi6)
                                var v85 uint8 = 1
                                _ = v85
                            } else {
                                fmt.Println(v1, v1, v53)
                            }
                        }
                    }
                    if false {
                        var v86 int32 = -337
                        if v4 {
                            v1 = (v1 & v1)
                            var v87 uint8 = 6
                            _ = v87
                        } else {
                            fmt.Println(v86, fi6, fi5)
                            if v4 {
                                fmt.Println(v1)
                            }
                            v88 := GenS8{F0: ("" + (("a" + ((("hello" + "") + "a") + "\n\t")) + "a")), F1: 878}
                            _ = v88
                        }
                        _ = v86
                    } else {
                        var v89 float64 = 0.1
                        fmt.Println(v53)
                        v3[(fi5 >> ((1) & 63))] = (^(fi5))
                        _ = v89
                    }
                }
                v1 += 128
            }
            _ = fi6
            _ = v53
        }
        v90 := GenS5{F0: ("a" + ("hello" + "foo bar")), F1: "a", F2: 105}
        _ = fi5
        _ = v90
    }
    if (v4 && true) {
        fmt.Println(v2, v4, v1)
        var v91 uint32 = 522
        _ = v91
    } else {
        switch -1 {
        case 3:
            for fi92 := 0; fi92 < 8; fi92++ {
                for fi93 := 0; fi93 < 4; fi93++ {
                    if (3 < 42) {
                        v94 := GenS5{F0: "\n\t", F1: "hello", F2: 0}
                        _ = v94
                    } else {
                        switch "foo bar" {
                        case "\n\t":
                            var v95 uint = 0
                            _ = v95
                        default:
                            var v96 int16 = 430
                            if (-1 >= 100) {
                                switch "\n\t" {
                                case "\n\t":
                                    v97 := GenS9{F0: 677}
                                    _ = v97
                                case "":
                                    var v98 int64 = 42
                                    _ = v98
                                default:
                                    var v99 uint8 = 0
                                    _ = v99
                                }
                                v96 = (v96 << ((1) & 63))
                                switch 9223372036854775807 {
                                case 318:
                                    var v100 int32 = v2
                                    _ = v100
                                case 9223372036854775807:
                                    v96 -= v96
                                case -687:
                                    v3[(v1 - 2147483647)] = (^(fi92))
                                default:
                                    var v101 uint8 = 255
                                    _ = v101
                                }
                                var v102 uint64 = 1
                                _ = v102
                            } else {
                                if false {
                                    if v4 {
                                        v2 -= (^(v2))
                                    } else {
                                        v2 = 178
                                        v4 = (true || v4)
                                        if false {
                                            v103 := GenS10{F0: 57, F1: v96}
                                            var v104 uint = 761
                                            if v4 {
                                                var v105 uint = v104
                                                fmt.Println(v2, v4, v1)
                                                v4 = (!v4)
                                                _ = v105
                                            } else {
                                                if (v4 || (1 == v2)) {
                                                    var v106 uint32 = 2147483647
                                                    _ = v106
                                                }
                                                v1 = (fi93 - fi92)
                                                if (1 > 280) {
                                                    v107 := []int{}
                                                    _ = v107
                                                }
                                            }
                                            _ = v103
                                            _ = v104
                                        }
                                    }
                                } else {
                                    v108 := GenS11{F0: 757, F1: 2}
                                    var v109 uint = 2
                                    _ = v108
                                    _ = v109
                                }
                            }
                            _ = v96
                        }
                    }
                    _ = fi93
                }
                _ = fi92
            }
        default:
            v110 := GenS12{F0: 42, F1: 1}
            _ = v110
        }
        var v111 int32 = (v2 - 0)
        _ = v111
    }
    for fi112 := 0; fi112 < 2; fi112++ {
        v113 := GenS13{F0: true, F1: 2}
        v114 := GenS14{F0: 127}
        if false {
            v1 = fi112
            v2 += (^(v2))
        } else {
            var v115 float64 = 1.5
            v4 = (false || (32 == 42))
            fmt.Println(v115)
            _ = v115
        }
        _ = fi112
        _ = v113
        _ = v114
    }
    v4 = ((!false) || (v2 > int32(v1)))
    v1 = 863
    var v116 float32 = 1e10
    switch "" {
    case "\n\t":
        v2 = v2
    case "foo bar":
        var v117 int8 = 0
        _ = v117
    case "hello":
        var v118 int8 = 44
        _ = v118
    }
    var v119 uint = 4294967295
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v116:", v116)
    fmt.Println("v119:", v119)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v116
    _ = v119
    _ = v2
    _ = v3
    _ = v4
}
