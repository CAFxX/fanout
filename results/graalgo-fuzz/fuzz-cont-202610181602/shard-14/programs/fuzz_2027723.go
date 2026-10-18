package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 uint8 }
type GenS2 struct { F0 uint64; F1 int16; F2 uint8 }
type GenS3 struct { F0 bool; F1 uint; F2 uint32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 bool; F1 string }
type GenS6 struct { F0 int64; F1 int64 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 string; F1 int64; F2 uint64 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 uint64; F1 uint }
type GenS11 struct { F0 int64; F1 uint8; F2 uint8 }
type GenS12 struct { F0 int64; F1 int16 }
func genFn0(p0 uint8, p1 uint64, p2 bool) string {
    v7 := GenS1{F0: (p0 % (((p0 >> ((2) & 63))) | 1)), F1: 201, F2: (p0 | p0)}
    if (1 < uint(p0)) {
        if false {
            if p2 {
                p1 += p1
            } else {
                fmt.Println(p0)
            }
            p2 = (2 < (p0 ^ (p0 * p0)))
            var v8 bool = (p1 >= p1)
            _ = v8
        }
        var v9 int8 = 127
        _ = v9
    } else {
        v10 := []int{32767, -9223372036854775808, 128}
        v10 = append(v10, -9223372036854775808)
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (("a" + "foo bar") + ("\n\t" + ("foo bar" + "a")))
}

func main() {
    var v1 uint64 = 2
    var v2 string = "hello"
    v3 := []uint32{}
    var v4 int = 692
    var v5 int16 = 32767
    v6 := make(map[int]int)
    for fi11 := 0; fi11 < 9; fi11++ {
        if true {
            var v12 int8 = 100
            for fi13 := 0; fi13 < 2; fi13++ {
                switch (fi13 % ((127) | 1)) {
                case 32767:
                    v6[4294967295] = fi11
                case 0:
                    var v14 uint64 = v1
                    _ = v14
                default:
                    _ = len(v6)
                }
                var v15 float32 = 3.14159
                _ = fi13
                _ = v15
            }
            _ = v12
        } else {
            v2 = ("foo bar" + v2)
            var v16 uint16 = 255
            _ = v16
        }
        fmt.Println(v2, fi11, v5)
        _ = fi11
    }
    v4 -= v4
    switch 2147483647 {
    case 32768:
        switch v2 {
        case "hello":
            for fi17 := 0; fi17 < 10; fi17++ {
                var v18 uint = 42
                for fi19 := 0; fi19 < 6; fi19++ {
                    var v20 int32 = 42
                    var v21 int64 = 1
                    _ = fi19
                    _ = v20
                    _ = v21
                }
                fmt.Println(v1)
                _ = fi17
                _ = v18
            }
        }
    case 861:
        switch v4 {
        case 42:
            v22 := GenS1{F0: 93, F1: 2147483647, F2: 1}
            _ = v22
        case 128:
            v23 := GenS1{F0: 1, F1: 65535, F2: 0}
            _ = v23
        case 32768:
            v24 := GenS2{F0: (v1 ^ 4294967295), F1: -1, F2: 128}
            _ = v24
        }
    }
    var v25 bool = (int16(v1) > (v5 - (v5 << ((0) & 63))))
    if v25 {
        v26 := v6[(v4 ^ v4)]
        v27 := GenS1{F0: 255, F1: 261, F2: 255}
        _ = v26
        _ = v27
    } else {
        _ = len(v6)
        for fi28 := 0; fi28 < 10; fi28++ {
            v4 -= (v4 / (((v4 - v4)) | 1))
            _ = fi28
        }
        if v25 {
            if false {
                if v25 {
                    if len(v3) > 0 {
                        v3[0] = 2147483647
                    }
                    v29 := GenS3{F0: (!((-(v5)) > v5)), F1: 2, F2: 42}
                    switch v4 {
                    case 2147483647:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    case -732:
                        if (!v25) {
                            var v30 int64 = 42
                            v1 += (v1 / ((939) | 1))
                            v1 += 0
                            _ = v30
                        }
                    case -9223372036854775808:
                        fmt.Println(v4, v4)
                    default:
                        if v25 {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            if v25 {
                                var v31 bool = (false || v25)
                                fmt.Println(v5, v31)
                                var v32 float32 = (0.0 - (3.14159 - (-2.25 * 0.0)))
                                _ = v31
                                _ = v32
                            }
                            if true {
                                v25 = (((v2 + v2) > "a") || v25)
                                var v33 uint32 = 2147483648
                                _ = v33
                            }
                        } else {
                            for fi34 := 0; fi34 < 1; fi34++ {
                                v5 += v5
                                if false {
                                    fmt.Println(v4, v5)
                                }
                                _ = fi34
                            }
                            if v25 {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                if (("hello" + "\n\t") > "a") {
                                    for fi35 := 0; fi35 < 14; fi35++ {
                                        v36 := GenS1{F0: 127, F1: 0, F2: 42}
                                        if v25 {
                                            if v25 {
                                                fmt.Println(v4, v4)
                                            }
                                            if v25 {
                                                v2 = v2
                                                var v37 int64 = -9223372036854775808
                                                _ = v37
                                            } else {
                                                for fi38 := 0; fi38 < 17; fi38++ {
                                                    delete(v6, v4)
                                                    v1 = (^(v1))
                                                    v39 := GenS4{F0: (v25 || (2147483648 >= 149))}
                                                    _ = fi38
                                                    _ = v39
                                                }
                                            }
                                        } else {
                                            var v40 string = ((((v2 + (v2 + v2)) + "a") + ("hello" + v2)) + v2)
                                            fmt.Println(v5, v40)
                                            for fi41 := 0; fi41 < 18; fi41++ {
                                                var v42 float32 = 1.5
                                                delete(v6, v4)
                                                if (!((-781 >= 42) || (!v25))) {
                                                    v25 = (!v25)
                                                    var v43 uint8 = 42
                                                    v2 = "hello"
                                                    _ = v43
                                                } else {
                                                    var v44 uint8 = 128
                                                    v6[int(v44)] = 0
                                                    v25 = v25
                                                    _ = v44
                                                }
                                                if true {
                                                    v45 := GenS5{F0: (v25 && false), F1: (((v40 + ("" + v40)) + "foo bar") + ("\n\t" + (v40 + ("hello" + ""))))}
                                                    var v46 float32 = 0.1
                                                    v6[fi41] = v4
                                                    _ = v45
                                                    _ = v46
                                                } else {
                                                    if v25 {
                                                        var v47 uint8 = 128
                                                        switch (v2 + ("foo bar" + v40)) {
                                                        case "a":
                                                            v48 := GenS5{F0: ((v47 - 128) != v47), F1: v40}
                                                            _ = v48
                                                        case "foo bar":
                                                            var v49 uint = 1
                                                            _ = v49
                                                        default:
                                                            v3 = append(v3, 42)
                                                        }
                                                        if (!true) {
                                                            v50 := GenS2{F0: (v1 ^ 18446744073709551615), F1: (v5 << ((65535) & 63)), F2: (v47 >> ((1) & 63))}
                                                            v51 := v6[fi41]
                                                            if (!true) {
                                                                if true {
                                                                    var v52 bool = (!(v2 > "foo bar"))
                                                                    _ = v52
                                                                } else {
                                                                    if (66 >= 0) {
                                                                        var v53 float32 = 3.14159
                                                                        v25 = v25
                                                                        var v54 bool = true
                                                                        _ = v53
                                                                        _ = v54
                                                                    }
                                                                }
                                                                var v55 uint32 = 1
                                                                _ = len(v3)
                                                                _ = v55
                                                            } else {
                                                                v47 = 42
                                                                var v56 uint64 = v1
                                                                var v57 float64 = -2.25
                                                                var v58 int = (fi35 - fi41)
                                                                if len(v3) > 0 {
                                                                    v3[0] = 42
                                                                }
                                                                _ = v56
                                                                _ = v57
                                                                _ = v58
                                                            }
                                                            _ = v50
                                                            _ = v51
                                                        }
                                                        _ = v47
                                                    } else {
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        if false {
                                                            switch (fi35 % ((v4) | 1)) {
                                                            case -9223372036854775808:
                                                                if ((!false) && (0.0 > v42)) {
                                                                    var v59 float64 = 0.0
                                                                    _ = v59
                                                                } else {
                                                                    switch ((v2 + v40) + v40) {
                                                                    case "a":
                                                                        var v60 uint16 = uint16(v1)
                                                                        _ = v60
                                                                    case "foo bar":
                                                                        var v61 int = v4
                                                                        _ = v61
                                                                    default:
                                                                        var v62 uint16 = 0
                                                                        _ = v62
                                                                    }
                                                                    v42 -= -2.25
                                                                    v40 = (((v40 + (v2 + (v2 + (("hello" + ((v2 + (v2 + "")) + "a")) + "\n\t")))) + (v40 + v2)) + ("" + ("" + ((v2 + (v40 + v40)) + ((((v2 + "\n\t") + ((v40 + v2) + ("" + v2))) + (v2 + (v40 + ""))) + ("\n\t" + v2))))))
                                                                    fmt.Println(fi35, v25)
                                                                }
                                                            case 32768:
                                                                fmt.Println(v42, v40, v42)
                                                            case 127:
                                                                if v25 {
                                                                    var v63 int = (v4 << ((18446744073709551615) & 63))
                                                                    v2 = ""
                                                                    var v64 float64 = 1.5
                                                                    _ = v63
                                                                    _ = v64
                                                                } else {
                                                                    var v65 int32 = 42
                                                                    v66 := GenS1{F0: 153, F1: 447, F2: 127}
                                                                    v4 += (v4 + 32767)
                                                                    _ = v65
                                                                    _ = v66
                                                                }
                                                            }
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                        }
                                                        if (v25 || v25) {
                                                            switch (v4 + fi41) {
                                                            case 32767:
                                                                var v67 float64 = 1.5
                                                                _ = v67
                                                            case -9223372036854775808:
                                                                var v68 int16 = 538
                                                                _ = v68
                                                            default:
                                                                v1 = v1
                                                            }
                                                            v69 := GenS5{F0: (407 == fi35), F1: v2}
                                                            v70 := GenS6{F0: -766, F1: -949}
                                                            _ = v69
                                                            _ = v70
                                                        }
                                                    }
                                                    v71 := GenS4{F0: false}
                                                    _ = v71
                                                }
                                                v25 = v25
                                                _ = fi41
                                                _ = v42
                                            }
                                            _ = v40
                                        }
                                        v72 := GenS7{F0: 0}
                                        _ = fi35
                                        _ = v36
                                        _ = v72
                                    }
                                    var v73 int32 = 2147483647
                                    if (v73 >= v73) {
                                        var v74 int64 = 928
                                        _ = len(v3)
                                        var v75 float64 = 0.0
                                        _ = v74
                                        _ = v75
                                    } else {
                                        var v76 uint8 = 0
                                        _ = v76
                                    }
                                    _ = v73
                                }
                            }
                        }
                    }
                    _ = v29
                }
                var v77 uint64 = (v1 >> ((4294967295) & 63))
                _ = v77
            }
        }
    }
    if v25 {
        switch v2 {
        case "":
            v1 = (v1 << ((255) & 63))
        default:
            switch ("\n\t" + v2) {
            case "foo bar":
                for fi78 := 0; fi78 < 10; fi78++ {
                    var v79 int = v4
                    v80 := GenS8{F0: v2, F1: -9223372036854775808, F2: v1}
                    fmt.Println(v79, v4, fi78)
                    _ = fi78
                    _ = v79
                    _ = v80
                }
            }
        }
        var v81 bool = v25
        switch 159 {
        case 32768:
            switch ((v2 + v2) + (v2 + "hello")) {
            case "a":
                for fi82 := 0; fi82 < 15; fi82++ {
                    switch v2 {
                    case "":
                        var v83 float64 = 0.0
                        _ = v83
                    case "a":
                        if v81 {
                            for fi84 := 0; fi84 < 15; fi84++ {
                                v85 := GenS6{F0: -9223372036854775808, F1: 9223372036854775807}
                                if (65535 != 739) {
                                    var v86 int8 = -128
                                    if (v25 || v81) {
                                        fmt.Println(v81, fi84, v4)
                                        v25 = (-117 <= v86)
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        v5 = v5
                                        v25 = v25
                                    } else {
                                        fmt.Println(fi84)
                                        v2 = "hello"
                                        switch (fi82 >> ((335) & 63)) {
                                        case 2147483647:
                                            if ((v4 == (fi84 | v4)) || ((v1 ^ (v1 >> ((255) & 63))) >= v1)) {
                                                v87 := GenS5{F0: v81, F1: ""}
                                                _ = v87
                                            } else {
                                                v3 = append(v3, 2)
                                                var v88 uint32 = 669
                                                v89 := GenS9{F0: int16(v1)}
                                                _ = v88
                                                _ = v89
                                            }
                                            var v90 string = ("foo bar" + "hello")
                                            _ = v90
                                        case 128:
                                            var v91 float32 = ((3.14159 * -2.25) + float32(float64((3.14159 / 1.5))))
                                            _ = v91
                                        }
                                        if v81 {
                                            _ = len(v3)
                                        }
                                    }
                                    fmt.Println(v81, v1, fi82)
                                    _ = v86
                                } else {
                                    if (-1 > 9223372036854775807) {
                                        var v92 int = v4
                                        switch (v4 | 9223372036854775807) {
                                        case -631:
                                            v81 = (1 != int64(v1))
                                        default:
                                            var v93 float64 = 0.0
                                            _ = v93
                                        }
                                        _ = v92
                                    } else {
                                        if v81 {
                                            var v94 uint64 = (v1 ^ (v1 >> ((1) & 63)))
                                            if (false || false) {
                                                if (!v25) {
                                                    v2 = ((((v2 + v2) + "foo bar") + "a") + (("a" + "") + (v2 + v2)))
                                                    switch "" {
                                                    case "foo bar":
                                                        var v95 int8 = 4
                                                        _ = v95
                                                    case "hello":
                                                        v1 = v94
                                                    default:
                                                        var v96 bool = false
                                                        _ = v96
                                                    }
                                                }
                                                var v97 string = v2
                                                var v98 uint64 = (v94 & (v1 & (v1 % ((4294967295) | 1))))
                                                _ = v97
                                                _ = v98
                                            }
                                            var v99 int16 = (v5 % (((v5 * v5)) | 1))
                                            _ = v94
                                            _ = v99
                                        } else {
                                            var v100 uint16 = 65535
                                            fmt.Println(v81, fi84)
                                            v25 = true
                                            v5 += (v5 % ((567) | 1))
                                            _ = v100
                                        }
                                        v101 := GenS7{F0: 94}
                                        v5 = v5
                                        _ = v101
                                    }
                                    v25 = true
                                    if v81 {
                                        var v102 uint = 255
                                        _ = v102
                                    } else {
                                        var v103 uint = 4294967295
                                        v104 := GenS10{F0: v1, F1: v103}
                                        _ = v103
                                        _ = v104
                                    }
                                }
                                _ = fi84
                                _ = v85
                            }
                            v5 += v5
                            if false {
                                for fi105 := 0; fi105 < 8; fi105++ {
                                    var v106 uint16 = 0
                                    v3 = append(v3, 2147483647)
                                    _ = fi105
                                    _ = v106
                                }
                                for fi107 := 0; fi107 < 10; fi107++ {
                                    switch (v2 + "hello") {
                                    case "\n\t":
                                        v108 := GenS4{F0: (!(3.14159 != ((0.1 / ((0.0 - -2.25) - 0.0)) + 1.5)))}
                                        _ = v108
                                    case "":
                                        v25 = (1 == 127)
                                        v109 := GenS9{F0: (v5 << ((uint(v1)) & 63))}
                                        _ = v109
                                    case "hello":
                                        switch 1 {
                                        case -395:
                                            v4 -= (fi107 / ((32768) | 1))
                                        case 127:
                                            switch ("\n\t" + "a") {
                                            case "":
                                                v110 := GenS11{F0: 42, F1: 127, F2: 0}
                                                var v111 uint = 2
                                                v111 = v111
                                                v2 = v2
                                                _ = len(v6)
                                                _ = v110
                                                _ = v111
                                            }
                                        }
                                    default:
                                        v25 = (!v81)
                                    }
                                    _ = fi107
                                }
                            } else {
                                var v112 float32 = 0.1
                                v1 -= (v1 + v1)
                                _ = v112
                            }
                        }
                    case "\n\t":
                        fmt.Println(fi82)
                    }
                    for fi113 := 0; fi113 < 6; fi113++ {
                        var v114 float32 = 3.14159
                        v5 += 1
                        fmt.Println(fi113, v2, v2)
                        _ = fi113
                        _ = v114
                    }
                    fmt.Println(v25, v4)
                    _ = fi82
                }
            case "hello":
                var v115 float32 = float32(-2.25)
                _ = v115
            case "":
                v4 += 0
            }
        case 4294967295:
            for fi116 := 0; fi116 < 2; fi116++ {
                for fi117 := 0; fi117 < 15; fi117++ {
                    if v25 {
                        var v118 uint8 = 0
                        v4 = -4
                        _ = v118
                    }
                    _ = fi117
                }
                switch fi116 {
                case 127:
                    if (42 == 1) {
                        v6[-313] = (v4 % (((v4 & -9223372036854775808)) | 1))
                    }
                default:
                    var v119 int64 = 0
                    _ = v119
                }
                _ = fi116
            }
        case 2147483647:
            for fi120 := 0; fi120 < 20; fi120++ {
                var v121 string = ("a" + v2)
                _ = fi120
                _ = v121
            }
        default:
            v2 = "\n\t"
        }
        _ = v81
    } else {
        if v25 {
            var v122 bool = false
            fmt.Println(v25)
            if ((("foo bar" + (("\n\t" + "\n\t") + "a")) + "") < (v2 + (v2 + v2))) {
                if (0 != 97) {
                    v123 := GenS12{F0: -1, F1: 1}
                    _ = v123
                } else {
                    var v124 uint16 = 2
                    v1 -= (v1 * (v1 + (^(v1))))
                    switch v2 {
                    case "":
                        v6[(v4 & (v4 << ((2) & 63)))] = -44
                    case "foo bar":
                        var v125 uint8 = 255
                        _ = v125
                    case "\n\t":
                        fmt.Println(v122, v25)
                    default:
                        var v126 float64 = (1.5 * float64(0.0))
                        _ = v126
                    }
                    _ = v124
                }
            } else {
                v25 = (v1 <= v1)
                v5 -= v5
                var v127 int8 = 0
                _ = v127
            }
            _ = v122
        }
        if v25 {
            if false {
                if len(v3) > 0 {
                    v3[0] = 396
                }
            }
            switch v2 {
            case "hello":
                v5 = v5
            default:
                var v128 float32 = ((-2.25 * ((0.0 * (-2.25 - 1.5)) + -2.25)) * 0.0)
                _ = v128
            }
            delete(v6, (v4 << ((397) & 63)))
        }
    }
    fmt.Println(v2, v2, v25)
    for fi129 := 0; fi129 < 2; fi129++ {
        v25 = (609 < 42)
        var v130 int = -826
        _ = fi129
        _ = v130
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
