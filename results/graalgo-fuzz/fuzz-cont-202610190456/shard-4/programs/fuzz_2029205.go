package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 int8; F1 uint; F2 int16 }
type GenS3 struct { F0 int; F1 int16; F2 uint8 }
type GenS4 struct { F0 uint16; F1 int8 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int; F1 uint16 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint8; F1 int32; F2 uint32 }
type GenS9 struct { F0 string; F1 int32 }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 uint; F1 int }
type GenS12 struct { F0 string }
type GenS13 struct { F0 int8; F1 int64 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 float32 = 3.14159
    v3 := make(map[string]string)
    var v4 bool = false
    v5 := make(map[uint32]bool)
    for fi6 := 0; fi6 < 4; fi6++ {
        if (!(1000 < 1)) {
            fmt.Println(v4)
            fmt.Println(v4)
        } else {
            _ = len(v3)
            var v7 float64 = ((float64((1e10 / v2)) + 1.5) / 1.0)
            var v8 float64 = v7
            _ = v7
            _ = v8
        }
        if (42 >= -1) {
            v1[0] = (!true)
            var v9 uint16 = 42
            _ = v9
        } else {
            v10 := []int{}
            fmt.Println(fi6, v2, v2)
            v11 := v5[1]
            _ = v10
            _ = v11
        }
        _ = fi6
    }
    var v12 int32 = -901
    fmt.Println(v4, v2, v2)
    switch 127 {
    case 1:
        for fi13 := 0; fi13 < 4; fi13++ {
            v4 = (!((fi13 < -2147483648) && (v12 <= -1)))
            _ = fi13
        }
    case 9223372036854775807:
        fmt.Println(v4, v2)
    case 32768:
        v14 := v3["hello"]
        _ = v14
    }
    v15 := GenS1{F0: 479, F1: v4}
    var v16 int32 = (v12 - (v12 | (^(v12))))
    if (1 >= 2147483647) {
        switch ("a" + "foo bar") {
        case "\n\t":
            v2 = v2
        case "foo bar":
            v17 := []int{}
            _ = v17
        case "":
            var v18 uint8 = 127
            _ = v18
        }
        v19 := GenS2{F0: -1, F1: 290, F2: -598}
        _ = v19
    } else {
        if v4 {
            v12 = -2147483648
        } else {
            var v20 int = -566
            if false {
                v21 := GenS1{F0: v20, F1: v4}
                if v4 {
                    switch v20 {
                    case 740:
                        var v22 uint = 4294967295
                        _ = v22
                    default:
                        if v4 {
                            v23 := GenS2{F0: -128, F1: 0, F2: 1}
                            v20 = 32768
                            v1[2] = v4
                            _ = v23
                        }
                    }
                    for fi24 := 0; fi24 < 4; fi24++ {
                        if (958 > 2) {
                            v16 -= 42
                            var v25 int8 = -128
                            _ = v25
                        }
                        _ = fi24
                    }
                    v26 := []int{(v20 ^ v20)}
                    _ = v26
                } else {
                    if (9223372036854775807 != 466) {
                        v12 = v12
                        if v4 {
                            v12 += 751
                            if true {
                                switch ("hello" + "") {
                                case "foo bar":
                                    var v27 float64 = 3.14159
                                    _ = v27
                                case "":
                                    switch 9223372036854775807 {
                                    case -1:
                                        if v4 {
                                            v28 := []int{}
                                            v29 := GenS3{F0: -287, F1: 32767, F2: 1}
                                            _ = v28
                                            _ = v29
                                        }
                                    case 32767:
                                        v12 += 1
                                    case -170:
                                        if v4 {
                                            switch v20 {
                                            case 32767:
                                                for fi30 := 0; fi30 < 14; fi30++ {
                                                    if ((!true) || (false && (0 == 127))) {
                                                        var v31 uint8 = 35
                                                        v20 -= v20
                                                        _ = v31
                                                    } else {
                                                        v20 = (v20 >> ((942) & 63))
                                                        v32 := GenS4{F0: 255, F1: 0}
                                                        _ = v32
                                                    }
                                                    _ = fi30
                                                }
                                            case 4294967295:
                                                var v33 float64 = -2.25
                                                _ = v33
                                            }
                                            var v34 int64 = 42
                                            _ = v34
                                        } else {
                                            switch (v20 % ((v20) | 1)) {
                                            case 700:
                                                v16 = v12
                                            }
                                            var v35 uint64 = 18446744073709551615
                                            _ = v35
                                        }
                                    }
                                case "hello":
                                    v36 := GenS5{F0: -1}
                                    _ = v36
                                default:
                                    v16 -= (v16 / ((v12) | 1))
                                }
                                v16 += 42
                            }
                            v37 := GenS6{F0: v20, F1: 2}
                            _ = v37
                        }
                        switch v20 {
                        case -1:
                            v38 := GenS7{F0: 1}
                            _ = v38
                        case 127:
                            v4 = (-2147483648 <= v20)
                        case -9223372036854775808:
                            v4 = v4
                        }
                    }
                    if false {
                        switch "" {
                        case "\n\t":
                            if v4 {
                                for fi39 := 0; fi39 < 16; fi39++ {
                                    v16 -= (v16 - 376)
                                    v40 := []int{(v20 >> ((185) & 63)), 128}
                                    var v41 int32 = -921
                                    _ = fi39
                                    _ = v40
                                    _ = v41
                                }
                                if (v4 && (!false)) {
                                    var v42 uint8 = 95
                                    var v43 uint8 = v42
                                    switch "a" {
                                    case "hello":
                                        if (!(!v4)) {
                                            v1[0] = (255 != 32767)
                                        } else {
                                            fmt.Println(v20, v2, v43)
                                        }
                                    case "foo bar":
                                        _ = len(v3)
                                    case "":
                                        var v44 float32 = float32(3.14159)
                                        _ = v44
                                    }
                                    _ = v42
                                    _ = v43
                                }
                                for fi45 := 0; fi45 < 5; fi45++ {
                                    fmt.Println(fi45, v2, v2)
                                    v16 += -693
                                    _ = fi45
                                }
                            } else {
                                var v46 uint64 = 18446744073709551615
                                v47 := GenS1{F0: (v20 | (v20 + 42)), F1: (-292 != int64(v20))}
                                for fi48 := 0; fi48 < 9; fi48++ {
                                    v49 := GenS1{F0: fi48, F1: (v2 != v2)}
                                    _ = fi48
                                    _ = v49
                                }
                                _ = v46
                                _ = v47
                            }
                        }
                        var v50 bool = v4
                        if false {
                            v51 := GenS8{F0: 4, F1: (v16 + v16), F2: 27}
                            _ = v51
                        }
                        _ = v50
                    } else {
                        var v52 bool = v4
                        v5[2147483648] = true
                        v2 = v2
                        _ = v52
                    }
                }
                var v53 uint16 = 2
                _ = v21
                _ = v53
            } else {
                if v4 {
                    v16 -= (v16 + -450)
                    fmt.Println(v2)
                    fmt.Println(v4, v4)
                } else {
                    switch "" {
                    case "foo bar":
                        if (v4 && true) {
                            v4 = (v20 < v20)
                        } else {
                            switch 128 {
                            case 122:
                                switch "hello" {
                                case "foo bar":
                                    v54 := GenS7{F0: 124}
                                    _ = v54
                                case "hello":
                                    v55 := GenS3{F0: v20, F1: -32768, F2: 128}
                                    _ = v55
                                case "\n\t":
                                    for fi56 := 0; fi56 < 17; fi56++ {
                                        if (127 > 1) {
                                            for fi57 := 0; fi57 < 4; fi57++ {
                                                var v58 string = "foo bar"
                                                _ = fi57
                                                _ = v58
                                            }
                                            v16 = v16
                                        } else {
                                            if true {
                                                v59 := GenS9{F0: "a", F1: (v12 + (v16 ^ v12))}
                                                var v60 uint8 = 0
                                                v61 := []int{fi56, v20}
                                                _ = v59
                                                _ = v60
                                                _ = v61
                                            }
                                            v62 := []int{fi56, -1}
                                            v12 = (-(v12))
                                            _ = v62
                                        }
                                        _ = fi56
                                    }
                                }
                            case -647:
                                delete(v5, 42)
                            }
                        }
                    case "a":
                        v16 += (v12 % (((v16 - v12)) | 1))
                    }
                    switch "hello" {
                    case "\n\t":
                        if v4 {
                            v2 = v2
                        } else {
                            if (!(v4 && ((float32((1.5 * (float64(1e10) / 3.14159))) <= v2) || (705 >= 779)))) {
                                v63 := v1[1]
                                v20 = -2147483648
                                var v64 int = (v20 ^ v20)
                                _ = v63
                                _ = v64
                            } else {
                                for fi65 := 0; fi65 < 15; fi65++ {
                                    switch "a" {
                                    case "":
                                        v20 = fi65
                                    case "foo bar":
                                        var v66 int16 = -1
                                        _ = v66
                                    case "a":
                                        for fi67 := 0; fi67 < 12; fi67++ {
                                            v12 += -2147483648
                                            var v68 int32 = (v12 & v16)
                                            v12 = int32(fi67)
                                            _ = fi67
                                            _ = v68
                                        }
                                    default:
                                        var v69 bool = v4
                                        _ = v69
                                    }
                                    var v70 uint64 = 0
                                    switch (v20 ^ 702) {
                                    case 127:
                                        v4 = ((6 < int8(v12)) || (!(!(!v4))))
                                    }
                                    _ = fi65
                                    _ = v70
                                }
                            }
                            if (2 <= 255) {
                                var v71 int8 = 100
                                v72 := []int{(v20 | v20), -523}
                                _ = v71
                                _ = v72
                            }
                            var v73 float64 = 1e10
                            _ = v73
                        }
                    case "":
                        if v4 {
                            for fi74 := 0; fi74 < 6; fi74++ {
                                switch "\n\t" {
                                case "\n\t":
                                    fmt.Println(v16, v20, v4)
                                default:
                                    for fi75 := 0; fi75 < 16; fi75++ {
                                        var v76 float32 = (v2 * 0.1)
                                        v76 += (((0.1 * v2) * 0.1) * 1e10)
                                        _ = fi75
                                        _ = v76
                                    }
                                }
                                for fi77 := 0; fi77 < 2; fi77++ {
                                    if true {
                                        if (!(1e10 == 0.0)) {
                                            var v78 int64 = -516
                                            v79 := GenS2{F0: -1, F1: 2, F2: 0}
                                            _ = v78
                                            _ = v79
                                        } else {
                                            var v80 int16 = -905
                                            v2 = (v2 - (0.0 - v2))
                                            v16 = v16
                                            _ = v80
                                        }
                                        var v81 uint16 = 0
                                        v82 := GenS8{F0: 0, F1: 1, F2: 552}
                                        _ = v81
                                        _ = v82
                                    } else {
                                        var v83 int16 = 1000
                                        _ = v83
                                    }
                                    if ((!v4) && v4) {
                                        v84 := GenS9{F0: "foo bar", F1: v12}
                                        if (9223372036854775807 <= -9223372036854775808) {
                                            var v85 int = (fi77 << ((4294967295) & 63))
                                            switch "\n\t" {
                                            case "\n\t":
                                                var v86 int64 = -1
                                                _ = v86
                                            }
                                            v16 -= v12
                                            _ = v85
                                        }
                                        _ = v84
                                    } else {
                                        var v87 uint64 = 2
                                        if (1 > -9223372036854775808) {
                                            v88 := GenS10{F0: (^(v87))}
                                            v2 += (float32(((-2.25 + float64(0.0)) / 1.5)) / float32(1.5))
                                            _ = v88
                                        } else {
                                            v89 := []int{32768, (-(v20)), (^(fi74)), fi77}
                                            fmt.Println(v16)
                                            if ((!true) || v4) {
                                                var v90 int16 = 898
                                                var v91 string = "\n\t"
                                                _ = v90
                                                _ = v91
                                            } else {
                                                v92 := GenS11{F0: 255, F1: (^(fi74))}
                                                v2 += 0.1
                                                v93 := GenS7{F0: 0}
                                                var v94 float32 = v2
                                                _ = v92
                                                _ = v93
                                                _ = v94
                                            }
                                            _ = v89
                                        }
                                        _ = len(v3)
                                        _ = v87
                                    }
                                    fmt.Println(fi77, fi74, v12)
                                    _ = fi77
                                }
                                _ = fi74
                            }
                        }
                    case "foo bar":
                        v2 = (1e10 / (1e10 + (v2 + float32(float64(((-2.25 * v2) / v2))))))
                    default:
                        switch 2147483647 {
                        case 870:
                            v95 := GenS12{F0: ("hello" + ("" + ("\n\t" + ("hello" + (("" + ("" + "hello")) + "\n\t")))))}
                            _ = v95
                        case 0:
                            for fi96 := 0; fi96 < 7; fi96++ {
                                switch "\n\t" {
                                case "\n\t":
                                    if (!v4) {
                                        v97 := v1[2]
                                        var v98 bool = (!v97)
                                        v99 := []int{(fi96 ^ fi96), (v20 | v20), fi96, (fi96 * v20)}
                                        _ = v97
                                        _ = v98
                                        _ = v99
                                    } else {
                                        if ((-2.25 + 0.0) != 1e10) {
                                            var v100 uint64 = 4294967295
                                            switch fi96 {
                                            case 128:
                                                v101 := []int{(-(v20)), v20, v20}
                                                _ = v101
                                            case 740:
                                                v20 += -9223372036854775808
                                            case 32767:
                                                v12 += (-(v16))
                                            default:
                                                var v102 float32 = (((float32(1.5) / 1.0) + (((1e10 * v2) + -2.25) * v2)) - (-2.25 + v2))
                                                _ = v102
                                            }
                                            switch "\n\t" {
                                            case "":
                                                v16 = (v16 / ((-626) | 1))
                                            default:
                                                v103 := GenS3{F0: -847, F1: -1, F2: 0}
                                                _ = v103
                                            }
                                            _ = v100
                                        }
                                        v104 := GenS13{F0: 127, F1: -9223372036854775808}
                                        _ = v104
                                    }
                                case "a":
                                    v20 = (fi96 | -665)
                                }
                                _ = fi96
                            }
                        case 32768:
                            if (-824 >= -9223372036854775808) {
                                for fi105 := 0; fi105 < 19; fi105++ {
                                    if (841 == 0) {
                                        v106 := GenS13{F0: int8(v20), F1: 0}
                                        _ = v106
                                    }
                                    v107 := []int{v20, v20, fi105, (-(fi105))}
                                    var v108 uint = 65535
                                    _ = fi105
                                    _ = v107
                                    _ = v108
                                }
                            } else {
                                var v109 int32 = v16
                                if v4 {
                                    var v110 uint = 1
                                    var v111 uint64 = 18446744073709551615
                                    _ = v110
                                    _ = v111
                                } else {
                                    for fi112 := 0; fi112 < 4; fi112++ {
                                        var v113 string = ""
                                        v114 := GenS3{F0: fi112, F1: 911, F2: 0}
                                        v115 := GenS6{F0: v20, F1: 0}
                                        _ = fi112
                                        _ = v113
                                        _ = v114
                                        _ = v115
                                    }
                                    for fi116 := 0; fi116 < 4; fi116++ {
                                        v117 := GenS7{F0: -412}
                                        var v118 uint = 255
                                        _ = fi116
                                        _ = v117
                                        _ = v118
                                    }
                                }
                                for fi119 := 0; fi119 < 19; fi119++ {
                                    var v120 int16 = 1000
                                    _ = fi119
                                    _ = v120
                                }
                                _ = v109
                            }
                        }
                    }
                }
                fmt.Println(v2, v20)
                v2 = (1e10 + 3.14159)
            }
            v12 = (v12 | (^(v12)))
            _ = v20
        }
    }
    v121 := v1[228]
    for fi122 := 0; fi122 < 7; fi122++ {
        _ = len(v1)
        var v123 bool = (v4 && ("" >= "hello"))
        _ = fi122
        _ = v123
    }
    fmt.Println(v12, v121)
    var v124 bool = (143 != 4294967295)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v121:", v121)
    fmt.Println("v124:", v124)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v121
    _ = v124
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
