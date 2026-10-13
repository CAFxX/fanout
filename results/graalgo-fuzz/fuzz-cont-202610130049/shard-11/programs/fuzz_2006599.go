package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 uint32; F1 uint64; F2 uint8 }
type GenS4 struct { F0 uint; F1 string; F2 int16 }
type GenS5 struct { F0 int16; F1 int32 }
type GenS6 struct { F0 int64; F1 int }
type GenS7 struct { F0 int8; F1 uint16; F2 uint }
type GenS8 struct { F0 int8; F1 int; F2 uint8 }
func genFn0(p0 uint32, p1 int16) uint16 {
    if (int(p1) <= 1) {
        v7 := []int{42}
        switch 194 {
        case 4294967295:
            p1 = (p1 << ((290) & 63))
        case 923:
            p1 += p1
        case -2147483648:
            var v8 int8 = -1
            _ = v8
        }
        if true {
            var v9 int16 = p1
            _ = v9
        }
        _ = v7
    }
    for fi10 := 0; fi10 < 11; fi10++ {
        v11 := []int{128, 128, fi10, (^(fi10))}
        var v12 uint = 666
        var v13 uint64 = 642
        _ = fi10
        _ = v11
        _ = v12
        _ = v13
    }
    var v14 float32 = (float32(-2.25) * 0.0)
    p1 = p1
    v15 := GenS1{F0: 1}
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    return 892
}

func main() {
    var v1 uint32 = 65535
    var v2 float32 = 0.1
    var v3 int8 = 74
    var v4 bool = (!(255 != 536))
    var v5 uint32 = v1
    var v6 int16 = -1
    v16 := []int{341, -244, 371, -9223372036854775808}
    var v17 string = "hello"
    if v4 {
        var v18 int64 = 42
        if ((((1e10 + 0.0) * ((3.14159 / 1.0) + ((3.14159 + 1.5) - -2.25))) / (-2.25 - 0.0)) != 0.1) {
            v16 = append(v16, 4294967295)
            if len(v16) > 0 {
                v16[0] = 32768
            }
            for fi19 := 0; fi19 < 1; fi19++ {
                if (v4 || ((!false) || (42 >= 179))) {
                    var v20 float64 = 1.5
                    for fi21 := 0; fi21 < 8; fi21++ {
                        var v22 uint = 0
                        v20 -= (v20 - ((float64(v2) - float64((3.14159 * ((0.0 / v2) - v2)))) / 1.0))
                        fmt.Println(v1, v5)
                        _ = fi21
                        _ = v22
                    }
                    _ = v20
                }
                _ = fi19
            }
        } else {
            if true {
                switch -216 {
                case 237:
                    var v23 uint8 = 42
                    _ = v23
                }
                v5 += (^(v1))
                if len(v16) > 0 {
                    v16[0] = 42
                }
            }
            var v24 float64 = 3.14159
            if v4 {
                for fi25 := 0; fi25 < 5; fi25++ {
                    var v26 uint32 = v1
                    _ = fi25
                    _ = v26
                }
            } else {
                for fi27 := 0; fi27 < 13; fi27++ {
                    if (2147483647 == 1) {
                        for fi28 := 0; fi28 < 20; fi28++ {
                            switch (fi28 + fi28) {
                            case 1:
                                if len(v16) >= 2 {
                                    v16 = v16[1:len(v16)]
                                }
                                var v29 bool = false
                                _ = v29
                            default:
                                var v30 bool = true
                                _ = v30
                            }
                            _ = len(v16)
                            switch fi28 {
                            case -931:
                                v1 = (v5 * v5)
                                fmt.Println(v3, v1)
                                var v31 bool = (!(!v4))
                                _ = v31
                            case -144:
                                switch fi27 {
                                case 2147483647:
                                    fmt.Println(v4)
                                case 0:
                                    var v32 uint64 = 2
                                    _ = v32
                                default:
                                    var v33 int64 = v18
                                    _ = v33
                                }
                            case -683:
                                v17 = ("a" + v17)
                            default:
                                fmt.Println(v6, fi28, v6)
                            }
                            _ = fi28
                        }
                        fmt.Println(v17, v24, v24)
                        switch "\n\t" {
                        case "foo bar":
                            if true {
                                v3 += (v3 ^ int8(v5))
                            }
                        case "a":
                            switch v17 {
                            case "foo bar":
                                v24 = v24
                            default:
                                _ = len(v16)
                            }
                        }
                    } else {
                        if (!v4) {
                            var v34 uint16 = 0
                            for fi35 := 0; fi35 < 14; fi35++ {
                                v16 = append(v16, fi27)
                                v36 := GenS1{F0: v3}
                                _ = fi35
                                _ = v36
                            }
                            _ = v34
                        }
                    }
                    _ = fi27
                }
                var v37 int = 32768
                fmt.Println(v18, v4)
                _ = v37
            }
            _ = v24
        }
        fmt.Println(v2, v4, v2)
        _ = v18
    }
    v2 += 3.14159
    for fi38 := 0; fi38 < 8; fi38++ {
        if (127 >= 220) {
            for fi39 := 0; fi39 < 5; fi39++ {
                v17 = (v17 + v17)
                v40 := GenS1{F0: (v3 | (^(v3)))}
                switch ((v17 + ((v17 + v17) + "hello")) + ("a" + (v17 + ((((v17 + v17) + (v17 + v17)) + (v17 + v17)) + "hello")))) {
                case "foo bar":
                    v17 = v17
                default:
                    switch (fi39 | 82) {
                    case -292:
                        fmt.Println(v5, v6, fi39)
                    case -2147483648:
                        if ("\n\t" < "a") {
                            if (0 < 2) {
                                v16 = append(v16, fi39)
                                v17 = v17
                                if (847 == -892) {
                                    v4 = (!(42 > 1))
                                } else {
                                    if (!v4) {
                                        v4 = (!(v5 <= (^(v1))))
                                        var v41 int16 = (^(v6))
                                        _ = v41
                                    } else {
                                        v6 += v6
                                        var v42 uint64 = 47
                                        if len(v16) > 0 {
                                            v16[0] = (fi38 / (((fi38 + (fi39 | 1))) | 1))
                                        }
                                        _ = v42
                                    }
                                    if (v4 && v4) {
                                        v43 := GenS2{F0: 251, F1: "a"}
                                        if (1000 >= (-(v6))) {
                                            var v44 string = ((v17 + "") + "")
                                            _ = v44
                                        }
                                        _ = v43
                                    } else {
                                        fmt.Println(v17)
                                    }
                                    if v4 {
                                        v1 -= 4294967295
                                    }
                                    var v45 float32 = v2
                                    if v4 {
                                        v17 = "\n\t"
                                        switch fi38 {
                                        case 738:
                                            fmt.Println(v2)
                                            fmt.Println(v17)
                                            var v46 uint64 = 4294967295
                                            _ = v46
                                        }
                                        var v47 int64 = int64(v1)
                                        _ = v47
                                    }
                                    _ = v45
                                }
                            } else {
                                _ = len(v16)
                                v5 = 175
                            }
                        }
                    case -10:
                        if len(v16) > 0 {
                            v16[0] = 0
                        }
                    }
                }
                _ = fi39
                _ = v40
            }
            v2 = 1e10
        } else {
            for fi48 := 0; fi48 < 1; fi48++ {
                v49 := GenS2{F0: 1, F1: v17}
                v50 := GenS3{F0: (v1 - (v1 | 4294967295)), F1: 4294967295, F2: 47}
                var v51 int8 = (v3 | (v3 << ((159) & 63)))
                _ = fi48
                _ = v49
                _ = v50
                _ = v51
            }
            var v52 float64 = 1e10
            v53 := GenS1{F0: (v3 << ((846) & 63))}
            _ = v52
            _ = v53
        }
        if (!(311 >= -709)) {
            for fi54 := 0; fi54 < 10; fi54++ {
                var v55 uint8 = 39
                if (0.1 > -2.25) {
                    if (!v4) {
                        v16 = append(v16, -2147483648)
                        var v56 int32 = -1
                        fmt.Println(v56)
                        _ = v56
                    } else {
                        var v57 int = fi54
                        var v58 int = 4294967295
                        switch (fi54 << ((65535) & 63)) {
                        case -408:
                            if (v5 < 1) {
                                var v59 uint8 = 0
                                if v4 {
                                    var v60 uint8 = v59
                                    switch "foo bar" {
                                    case "hello":
                                        var v61 string = v17
                                        switch (fi54 / (((fi38 - (fi54 & v57))) | 1)) {
                                        case 127:
                                            v59 = v59
                                        case 42:
                                            v62 := GenS3{F0: (v1 + 65535), F1: 0, F2: v60}
                                            _ = v62
                                        }
                                        var v63 uint = 42
                                        v6 = -308
                                        _ = v61
                                        _ = v63
                                    case "a":
                                        var v64 int8 = v3
                                        _ = v64
                                    case "\n\t":
                                        v65 := GenS4{F0: 4294967295, F1: ((v17 + v17) + (("hello" + v17) + v17)), F2: 1000}
                                        _ = v65
                                    }
                                    _ = v60
                                }
                                var v66 uint32 = v1
                                _ = v59
                                _ = v66
                            } else {
                                var v67 uint16 = 2
                                _ = v67
                            }
                        case 4294967295:
                            v4 = (65535 >= v1)
                        }
                        _ = v57
                        _ = v58
                    }
                    if (v17 >= (v17 + ((v17 + "hello") + ("a" + ((((v17 + (v17 + v17)) + "\n\t") + "\n\t") + (("foo bar" + v17) + v17)))))) {
                        var v68 int16 = v6
                        switch ((v17 + "") + ("foo bar" + "a")) {
                        case "\n\t":
                            if len(v16) >= 2 {
                                v16 = v16[1:len(v16)]
                            }
                        }
                        var v69 float64 = 1.5
                        _ = v68
                        _ = v69
                    } else {
                        v1 += 2
                    }
                    var v70 float32 = v2
                    _ = v70
                }
                var v71 int64 = -276
                _ = fi54
                _ = v55
                _ = v71
            }
        } else {
            if len(v16) >= 2 {
                v16 = v16[1:len(v16)]
            }
        }
        _ = fi38
    }
    for fi72 := 0; fi72 < 11; fi72++ {
        if ((v3 << ((116) & 63)) >= v3) {
            if false {
                var v73 uint32 = (^(v5))
                _ = v73
            } else {
                if ((v17 + "hello") >= ("hello" + ((("" + ((v17 + v17) + "a")) + v17) + v17))) {
                    var v74 int64 = 0
                    _ = v74
                } else {
                    for fi75 := 0; fi75 < 3; fi75++ {
                        v3 = v3
                        switch v17 {
                        case "":
                            v3 += v3
                        }
                        v17 = v17
                        switch "hello" {
                        case "foo bar":
                            var v76 int8 = -128
                            _ = v76
                        case "\n\t":
                            var v77 int64 = 11
                            _ = v77
                        }
                        v2 -= ((v2 - float32(3.14159)) * (v2 - 1e10))
                        _ = fi75
                    }
                }
            }
        } else {
            fmt.Println(fi72)
            if false {
                switch v17 {
                case "":
                    var v78 int16 = 251
                    _ = v78
                }
                v16 = append(v16, fi72)
                if false {
                    if (!v4) {
                        var v79 string = "a"
                        v80 := GenS1{F0: v3}
                        if v4 {
                            var v81 uint64 = 2
                            v82 := GenS2{F0: 569, F1: ("" + ((v79 + (v79 + v17)) + v79))}
                            _ = v81
                            _ = v82
                        }
                        _ = v79
                        _ = v80
                    }
                    var v83 string = v17
                    _ = v83
                } else {
                    if len(v16) > 0 {
                        v16[0] = 32767
                    }
                }
            } else {
                switch (fi72 | -730) {
                case 0:
                    for fi84 := 0; fi84 < 2; fi84++ {
                        if (!(true && v4)) {
                            if (255 < 638) {
                                v85 := GenS3{F0: (v1 - (v1 - (v5 & 2))), F1: 42, F2: 190}
                                _ = v85
                            } else {
                                v4 = v4
                            }
                            v2 = v2
                            v86 := GenS3{F0: (^(v5)), F1: 3, F2: 1}
                            _ = v86
                        } else {
                            if (true && v4) {
                                switch fi84 {
                                case 438:
                                    var v87 int8 = (v3 >> ((65535) & 63))
                                    _ = v87
                                case 42:
                                    v88 := GenS5{F0: v6, F1: 0}
                                    _ = v88
                                default:
                                    if true {
                                        _ = len(v16)
                                        v3 = (v3 % (((v3 ^ (v3 * (v3 >> ((1) & 63))))) | 1))
                                        fmt.Println(v6)
                                    }
                                }
                            } else {
                                fmt.Println(v2)
                                if ((1e10 <= 0.0) && v4) {
                                    var v89 int64 = 0
                                    if (1 <= v89) {
                                        var v90 uint16 = 255
                                        _ = v90
                                    }
                                    var v91 int32 = 1
                                    _ = v89
                                    _ = v91
                                } else {
                                    if ((128 < (fi72 << ((18446744073709551615) & 63))) && v4) {
                                        if true {
                                            var v92 int64 = 9223372036854775807
                                            if len(v16) >= 2 {
                                                v16 = v16[1:len(v16)]
                                            }
                                            switch v17 {
                                            case "a":
                                                var v93 uint64 = 4294967295
                                                _ = v93
                                            }
                                            _ = v92
                                        } else {
                                            if (v4 || v4) {
                                                fmt.Println(fi84, v1, v1)
                                                fmt.Println(fi84, v1, v6)
                                            } else {
                                                var v94 int32 = -1
                                                if (!v4) {
                                                    switch (fi84 | fi72) {
                                                    case -2147483648:
                                                        var v95 string = (v17 + "")
                                                        _ = v95
                                                    default:
                                                        if len(v16) >= 2 {
                                                            v16 = v16[1:len(v16)]
                                                        }
                                                    }
                                                    _ = len(v16)
                                                    fmt.Println(v3)
                                                }
                                                v16 = append(v16, fi84)
                                                _ = v94
                                            }
                                            v4 = true
                                            v96 := GenS6{F0: -985, F1: (fi84 - (fi72 * fi84))}
                                            _ = v96
                                        }
                                        if (-181 >= -1) {
                                            v97 := GenS4{F0: 914, F1: (v17 + v17), F2: (v6 << ((2) & 63))}
                                            if false {
                                                v1 -= (v1 >> ((65535) & 63))
                                                var v98 int32 = 1
                                                _ = v98
                                            }
                                            _ = v97
                                        }
                                    } else {
                                        v99 := GenS2{F0: 255, F1: v17}
                                        switch fi84 {
                                        case 1:
                                            if (v4 || (168 > 2)) {
                                                v1 += (v1 + (^(v1)))
                                                v4 = ((v2 - float32(1e10)) != 1.5)
                                                fmt.Println(v17)
                                                v6 = (v6 ^ -753)
                                            }
                                        case 128:
                                            v17 = "hello"
                                        default:
                                            var v100 int64 = -9223372036854775808
                                            _ = v100
                                        }
                                        v3 += v3
                                        v5 = (v1 + (^(v1)))
                                        _ = v99
                                    }
                                    if (true && true) {
                                        switch ((v17 + "\n\t") + "a") {
                                        case "foo bar":
                                            if len(v16) >= 2 {
                                                v16 = v16[1:len(v16)]
                                            }
                                        default:
                                            _ = len(v16)
                                        }
                                        var v101 uint = 1
                                        var v102 int8 = int8(fi72)
                                        _ = v101
                                        _ = v102
                                    } else {
                                        var v103 bool = v4
                                        _ = v103
                                    }
                                }
                                if v4 {
                                    if true {
                                        if (2147483647 >= 42) {
                                            var v104 int = fi84
                                            _ = v104
                                        } else {
                                            var v105 float32 = (((v2 / v2) + v2) + 1.5)
                                            v106 := GenS2{F0: 415, F1: (v17 + ((v17 + (v17 + "foo bar")) + "a"))}
                                            var v107 int32 = 0
                                            _ = v105
                                            _ = v106
                                            _ = v107
                                        }
                                    } else {
                                        if ("hello" < v17) {
                                            v108 := GenS7{F0: v3, F1: 2, F2: 2}
                                            _ = v108
                                        }
                                    }
                                    if len(v16) > 0 {
                                        v16[0] = (fi72 ^ (fi84 << ((42) & 63)))
                                    }
                                    var v109 float32 = v2
                                    _ = v109
                                }
                                v17 = (v17 + v17)
                                var v110 float64 = 0.1
                                var v111 uint16 = 385
                                if v4 {
                                    switch v17 {
                                    case "\n\t":
                                        var v112 float32 = (float32((0.1 + v110)) / 1.5)
                                        _ = v112
                                    case "a":
                                        v113 := GenS8{F0: -128, F1: (fi72 ^ fi84), F2: 0}
                                        v1 = v5
                                        _ = v113
                                    case "":
                                        var v114 int64 = 420
                                        _ = v114
                                    }
                                } else {
                                    v115 := GenS1{F0: int8(v111)}
                                    _ = v115
                                }
                                var v116 int64 = -9223372036854775808
                                var v117 uint32 = v5
                                if len(v16) > 0 {
                                    v16[0] = fi72
                                }
                                _ = v110
                                _ = v111
                                _ = v116
                                _ = v117
                            }
                            var v118 uint = 42
                            fmt.Println(fi72, v5)
                            _ = v118
                        }
                        v5 = v1
                        _ = fi84
                    }
                case 839:
                    switch v17 {
                    case "\n\t":
                        switch "a" {
                        case "hello":
                            var v119 int16 = v6
                            _ = v119
                        case "":
                            _ = len(v16)
                        default:
                            var v120 int64 = -596
                            _ = v120
                        }
                    }
                case -9223372036854775808:
                    v121 := GenS2{F0: 65535, F1: ((v17 + "") + (("hello" + "\n\t") + (v17 + (("" + ("a" + (v17 + "\n\t"))) + "\n\t"))))}
                    _ = v121
                }
                fmt.Println(fi72, v4)
            }
        }
        v5 -= 2
        var v122 float32 = v2
        _ = fi72
        _ = v122
    }
    var v123 int32 = 42
    v124 := GenS5{F0: 42, F1: -339}
    fmt.Println(v6, v5)
    var v125 int64 = int64(v123)
    v126 := GenS8{F0: v3, F1: 4294967295, F2: 1}
    switch 911 {
    case -155:
        v127 := GenS4{F0: 188, F1: ((v17 + (v17 + "")) + "hello"), F2: (-(v6))}
        _ = v127
    default:
        var v128 string = "\n\t"
        _ = v128
    }
    fmt.Println(v3)
    if len(v16) > 0 {
        v16[0] = int(v1)
    }
    v123 = 1
    if (0 > -9223372036854775808) {
        var v129 uint8 = 185
        v16 = append(v16, 32767)
        _ = v129
    }
    fmt.Println(v123, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v123:", v123)
    fmt.Println("v125:", v125)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v123
    _ = v124
    _ = v125
    _ = v126
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
