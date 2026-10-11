package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int; F1 bool; F2 bool }
type GenS5 struct { F0 int; F1 string }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint8; F1 int; F2 uint }
type GenS8 struct { F0 uint16; F1 int16; F2 string }
type GenS9 struct { F0 bool; F1 uint16 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 int64; F1 uint; F2 int32 }
type GenS12 struct { F0 uint8; F1 int8 }
func genFn0(p0 string, p1 int, p2 int64) uint8 {
    var v6 float32 = 1e10
    for fi7 := 0; fi7 < 6; fi7++ {
        v8 := []int{(p1 / ((fi7) | 1))}
        var v9 int32 = 100
        v6 -= v6
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 255
}

func main() {
    var v1 uint32 = 415
    v2 := make(map[int]bool)
    var v3 bool = (!false)
    v4 := make(map[string]string)
    var v5 int = 258
    v3 = (0 < 2)
    delete(v4, "")
    for fi10 := 0; fi10 < 7; fi10++ {
        v11 := []int{fi10}
        switch v5 {
        case 42:
            v12 := GenS1{F0: 42, F1: 51}
            _ = v12
        case 32767:
            switch "\n\t" {
            case "a":
                fmt.Println(v1, v5)
            }
        case -212:
            for fi13 := 0; fi13 < 17; fi13++ {
                if false {
                    v5 = fi10
                } else {
                    v3 = v3
                    var v14 int64 = 9223372036854775807
                    _ = v14
                }
                _ = len(v11)
                var v15 string = ""
                v5 = -1
                _ = fi13
                _ = v15
            }
        default:
            fmt.Println(v5, v1)
        }
        v3 = (127 >= -78)
        _ = fi10
        _ = v11
    }
    if (v5 != v5) {
        v16 := []int{-865, (v5 << ((1) & 63)), v5}
        _ = v16
    } else {
        var v17 uint64 = 4294967295
        var v18 uint8 = 2
        v1 += (v1 << ((65535) & 63))
        _ = v17
        _ = v18
    }
    switch v5 {
    case 392:
        var v19 float32 = (0.0 * (1e10 / (3.14159 / 1e10)))
        _ = v19
    case -816:
        v20 := GenS1{F0: 1, F1: -78}
        _ = v20
    default:
        v21 := GenS1{F0: -220, F1: -13}
        _ = v21
    }
    for fi22 := 0; fi22 < 12; fi22++ {
        for fi23 := 0; fi23 < 17; fi23++ {
            switch "\n\t" {
            case "a":
                if (!(42 != 277)) {
                    if (!v3) {
                        var v24 uint16 = 32767
                        v25 := []int{fi22, (fi23 * fi23), 0}
                        fmt.Println(fi23, v24)
                        _ = v24
                        _ = v25
                    }
                    if (1.5 <= 0.1) {
                        v26 := GenS1{F0: 9223372036854775807, F1: 7}
                        v4["\n\t"] = "\n\t"
                        _ = v26
                    } else {
                        v27 := GenS2{F0: 42}
                        _ = v27
                    }
                } else {
                    var v28 string = "a"
                    _ = v28
                }
            default:
                v29 := GenS1{F0: 392, F1: -32}
                _ = v29
            }
            _ = fi23
        }
        if (-1 > -184) {
            if v3 {
                var v30 bool = v3
                var v31 int8 = 46
                switch "a" {
                case "foo bar":
                    switch ("a" + "\n\t") {
                    case "a":
                        for fi32 := 0; fi32 < 2; fi32++ {
                            v33 := GenS2{F0: 2}
                            fmt.Println(v1)
                            fmt.Println(v5)
                            _ = fi32
                            _ = v33
                        }
                    case "":
                        v34 := []int{}
                        _ = v34
                    }
                case "\n\t":
                    for fi35 := 0; fi35 < 8; fi35++ {
                        if v30 {
                            fmt.Println(v5, v31)
                        }
                        if v3 {
                            fmt.Println(v31, v1)
                            v5 = 32768
                            if v30 {
                                switch "a" {
                                case "hello":
                                    var v36 uint = 683
                                    _ = v36
                                default:
                                    if (255 < 2) {
                                        switch "\n\t" {
                                        case "":
                                            v4["\n\t"] = "a"
                                        case "foo bar":
                                            v37 := v4["a"]
                                            _ = v37
                                        case "\n\t":
                                            v38 := GenS3{F0: 0}
                                            _ = v38
                                        }
                                    } else {
                                        v39 := GenS1{F0: -9223372036854775808, F1: v31}
                                        fmt.Println(v5, fi22, v3)
                                        fmt.Println(v30, fi22, v31)
                                        _ = v39
                                    }
                                }
                                if v3 {
                                    var v40 int8 = (^(v31))
                                    _ = v40
                                } else {
                                    if (342 == 6) {
                                        if v3 {
                                            v5 = fi35
                                            v31 += (v31 | v31)
                                        } else {
                                            var v41 float32 = 3.14159
                                            if (2 <= 2) {
                                                if (true || v30) {
                                                    var v42 float32 = float32(3.14159)
                                                    var v43 float64 = (-2.25 * 0.0)
                                                    _ = v42
                                                    _ = v43
                                                }
                                                v3 = (!v30)
                                            } else {
                                                switch v5 {
                                                case 1:
                                                    var v44 int16 = 1000
                                                    _ = v44
                                                case 32768:
                                                    v41 += 1.5
                                                case -948:
                                                    if true {
                                                        fmt.Println(v30)
                                                        var v45 float32 = 1e10
                                                        if (-9223372036854775808 > -1) {
                                                            v31 = 127
                                                        }
                                                        _ = v45
                                                    } else {
                                                        if (v30 && true) {
                                                            v46 := GenS4{F0: (fi22 << ((65535) & 63)), F1: v30, F2: false}
                                                            delete(v4, "")
                                                            _ = v46
                                                        } else {
                                                            if (v31 < (v31 / (((v31 / ((-56) | 1))) | 1))) {
                                                                var v47 float64 = float64(v41)
                                                                var v48 uint16 = 746
                                                                _ = v47
                                                                _ = v48
                                                            }
                                                        }
                                                    }
                                                }
                                                var v49 float64 = 3.14159
                                                if v3 {
                                                    var v50 uint8 = 117
                                                    switch "" {
                                                    case "":
                                                        fmt.Println(v41, v5)
                                                    default:
                                                        v31 -= v31
                                                    }
                                                    _ = v50
                                                }
                                                var v51 float32 = (v41 + v41)
                                                _ = v49
                                                _ = v51
                                            }
                                            v52 := v2[fi35]
                                            var v53 uint64 = 1
                                            var v54 uint64 = (v53 - 0)
                                            if (4294967295 > 18446744073709551615) {
                                                v3 = v3
                                                v1 = (^(v1))
                                                if (!true) {
                                                    v31 += v31
                                                    v55 := []int{}
                                                    if v52 {
                                                        var v56 bool = (472 > (^(v1)))
                                                        switch (-(fi35)) {
                                                        case 32767:
                                                            if (1.5 > 0.0) {
                                                                if v3 {
                                                                    var v57 uint32 = uint32(v54)
                                                                    _ = v57
                                                                } else {
                                                                    v58 := GenS5{F0: fi22, F1: "a"}
                                                                    _ = v58
                                                                }
                                                                v54 = 42
                                                            }
                                                        case -2147483648:
                                                            v31 = 52
                                                            if true {
                                                                var v59 bool = v56
                                                                var v60 bool = false
                                                                _ = v59
                                                                _ = v60
                                                            } else {
                                                                fmt.Println(v53)
                                                                switch ("hello" + ("hello" + (("" + ("foo bar" + "\n\t")) + "\n\t"))) {
                                                                case "foo bar":
                                                                    if (((0.1 / -2.25) <= 0.1) || ((-984 >= 0) || v52)) {
                                                                        if v56 {
                                                                            var v61 uint16 = 32767
                                                                            var v62 uint8 = 255
                                                                            v63 := GenS2{F0: 353}
                                                                            _ = v61
                                                                            _ = v62
                                                                            _ = v63
                                                                        }
                                                                    }
                                                                case "a":
                                                                    var v64 uint16 = 42
                                                                    _ = v64
                                                                default:
                                                                    v41 = 1e10
                                                                    var v65 string = "a"
                                                                    if (317 < uint(v1)) {
                                                                        if v52 {
                                                                            v66 := GenS6{F0: 65535}
                                                                            var v67 float64 = 0.1
                                                                            v68 := GenS7{F0: 245, F1: -2147483648, F2: 2}
                                                                            _ = v66
                                                                            _ = v67
                                                                            _ = v68
                                                                        }
                                                                        v31 = 0
                                                                        var v69 uint64 = (v54 + 4294967295)
                                                                        _ = v69
                                                                    }
                                                                    _ = v65
                                                                }
                                                                var v70 int8 = v31
                                                                _ = v70
                                                            }
                                                        }
                                                        if (v1 == (v1 - (v1 << ((558) & 63)))) {
                                                            switch int(v53) {
                                                            case 127:
                                                                switch (^(v5)) {
                                                                case 0:
                                                                    v71 := GenS7{F0: 232, F1: fi35, F2: 42}
                                                                    _ = v71
                                                                case 127:
                                                                    v30 = (!v30)
                                                                case 32768:
                                                                    switch (v5 % ((127) | 1)) {
                                                                    case 4294967295:
                                                                        if v52 {
                                                                            delete(v2, fi22)
                                                                        }
                                                                    case 576:
                                                                        var v72 int = int(v53)
                                                                        _ = v72
                                                                    }
                                                                }
                                                            case 32768:
                                                                var v73 int32 = 1
                                                                _ = v73
                                                            case 361:
                                                                var v74 bool = (337 <= 42)
                                                                _ = v74
                                                            default:
                                                                if (!false) {
                                                                    var v75 uint64 = 42
                                                                    var v76 int32 = -1
                                                                    _ = v75
                                                                    _ = v76
                                                                } else {
                                                                    fmt.Println(fi22, v41)
                                                                    var v77 uint16 = 2
                                                                    _ = v77
                                                                }
                                                            }
                                                            fmt.Println(v54, v31, fi22)
                                                            if (genFn0("\n\t", fi35, 953) > 128) {
                                                                var v78 float64 = (3.14159 + 0.1)
                                                                if (18446744073709551615 < 2) {
                                                                    if ((v78 * v78) < v78) {
                                                                        fmt.Println(v30, v52)
                                                                        var v79 uint64 = 2
                                                                        var v80 int = fi35
                                                                        if v30 {
                                                                            fmt.Println(v78)
                                                                            var v81 string = ""
                                                                            _ = v81
                                                                        } else {
                                                                            v52 = (v52 || v3)
                                                                            switch v5 {
                                                                            case 127:
                                                                                _ = len(v4)
                                                                            }
                                                                            fmt.Println(v3)
                                                                            fmt.Println(v30, v30, v41)
                                                                            if (v31 < int8(v53)) {
                                                                                switch "hello" {
                                                                                case "foo bar":
                                                                                    var v82 int64 = int64(v31)
                                                                                    _ = v82
                                                                                case "\n\t":
                                                                                    switch (fi22 | int(v53)) {
                                                                                    case 128:
                                                                                        var v83 int = 890
                                                                                        _ = v83
                                                                                    }
                                                                                default:
                                                                                    if len(v55) > 0 {
                                                                                        v55[0] = fi35
                                                                                    }
                                                                                }
                                                                                v1 = 2147483647
                                                                                if false {
                                                                                    if len(v55) > 0 {
                                                                                        v55[0] = fi22
                                                                                    }
                                                                                    if v30 {
                                                                                        v56 = v30
                                                                                        var v84 string = (("foo bar" + "") + ("" + "foo bar"))
                                                                                        var v85 int8 = v31
                                                                                        _ = v84
                                                                                        _ = v85
                                                                                    } else {
                                                                                        var v86 uint = 255
                                                                                        _ = v86
                                                                                    }
                                                                                    fmt.Println(fi35, v31, v1)
                                                                                }
                                                                            }
                                                                        }
                                                                        v87 := GenS8{F0: 42, F1: 32767, F2: "hello"}
                                                                        delete(v4, "foo bar")
                                                                        switch v80 {
                                                                        case 945:
                                                                            var v88 int16 = -130
                                                                            v89 := GenS8{F0: 255, F1: v88, F2: ((("a" + "") + "a") + "hello")}
                                                                            if (-71 >= 0) {
                                                                                _ = len(v55)
                                                                            } else {
                                                                                v54 += v54
                                                                                fmt.Println(v79, v52)
                                                                            }
                                                                            _ = len(v55)
                                                                            if true {
                                                                                var v90 int8 = v31
                                                                                v91 := GenS6{F0: 65535}
                                                                                _ = v90
                                                                                _ = v91
                                                                            } else {
                                                                                var v92 int8 = v31
                                                                                v93 := GenS1{F0: 1, F1: (v92 << ((2) & 63))}
                                                                                _ = v92
                                                                                _ = v93
                                                                            }
                                                                            _ = v88
                                                                            _ = v89
                                                                        case 1:
                                                                            if v3 {
                                                                                v5 = (fi22 | fi35)
                                                                                if false {
                                                                                    var v94 uint32 = 2147483647
                                                                                    var v95 int32 = -281
                                                                                    _ = v94
                                                                                    _ = v95
                                                                                } else {
                                                                                    var v96 bool = false
                                                                                    _ = v96
                                                                                }
                                                                                var v97 int8 = v31
                                                                                v98 := GenS4{F0: (^(v80)), F1: v30, F2: (false && true)}
                                                                                _ = v97
                                                                                _ = v98
                                                                            }
                                                                        case 32767:
                                                                            if v52 {
                                                                                var v99 uint16 = 0
                                                                                _ = v99
                                                                            }
                                                                        }
                                                                        _ = v79
                                                                        _ = v80
                                                                        _ = v87
                                                                    } else {
                                                                        var v100 bool = true
                                                                        if len(v55) > 0 {
                                                                            v55[0] = 228
                                                                        }
                                                                        v41 += 1e10
                                                                        _ = v100
                                                                    }
                                                                    fmt.Println(fi35, fi22, v52)
                                                                }
                                                                v101 := GenS9{F0: (("a" + ("a" + "hello")) < ""), F1: 42}
                                                                _ = v101
                                                                _ = v78
                                                            }
                                                        } else {
                                                            fmt.Println(v54, v54)
                                                            v31 += (v31 + 0)
                                                            v31 -= v31
                                                        }
                                                        _ = v56
                                                    } else {
                                                        fmt.Println(v3, v52)
                                                    }
                                                    var v102 uint = 18446744073709551615
                                                    if (4294967295 == v53) {
                                                        v53 -= v53
                                                        var v103 uint16 = 645
                                                        v52 = (!(1000 != 42))
                                                        _ = v103
                                                    }
                                                    _ = v102
                                                    _ = v55
                                                } else {
                                                    v31 -= (v31 & v31)
                                                }
                                            } else {
                                                v31 = 87
                                            }
                                            _ = v41
                                            _ = v52
                                            _ = v53
                                            _ = v54
                                        }
                                        v104 := v2[127]
                                        v104 = true
                                        var v105 string = ("foo bar" + "")
                                        var v106 int8 = 0
                                        _ = v104
                                        _ = v105
                                        _ = v106
                                    }
                                    fmt.Println(fi22, v1, v3)
                                    v107 := GenS6{F0: 18446744073709551615}
                                    _ = v107
                                }
                                var v108 uint32 = 1
                                var v109 uint32 = (v108 | (v1 << ((729) & 63)))
                                _ = v108
                                _ = v109
                            } else {
                                v5 = (fi22 & (-(fi35)))
                            }
                        } else {
                            fmt.Println(fi35, fi35, v31)
                        }
                        fmt.Println(v1, v3, v3)
                        v1 = (v1 - (v1 * v1))
                        var v110 uint8 = 127
                        _ = fi35
                        _ = v110
                    }
                }
                _ = v30
                _ = v31
            } else {
                var v111 uint16 = 32767
                _ = v111
            }
            v112 := GenS10{F0: v1}
            _ = v112
        } else {
            var v113 uint8 = 1
            v3 = true
            _ = v113
        }
        _ = fi22
    }
    for fi114 := 0; fi114 < 18; fi114++ {
        switch "a" {
        case "foo bar":
            v5 -= (-(v5))
        case "":
            var v115 uint = 554
            _ = v115
        case "a":
            v1 = 1
        default:
            var v116 int8 = 127
            _ = v116
        }
        v117 := GenS9{F0: v3, F1: 7}
        v5 += fi114
        _ = fi114
        _ = v117
    }
    v118 := GenS11{F0: 9223372036854775807, F1: 2, F2: -1}
    v3 = false
    _ = len(v4)
    var v119 bool = ((!v3) || (71 <= 4294967295))
    v120 := GenS1{F0: 0, F1: -1}
    for fi121 := 0; fi121 < 16; fi121++ {
        v122 := GenS12{F0: 0, F1: -80}
        var v123 float32 = -2.25
        for fi124 := 0; fi124 < 1; fi124++ {
            v125 := []int{fi124, v5, fi124, 2147483647}
            _ = fi124
            _ = v125
        }
        _ = fi121
        _ = v122
        _ = v123
    }
    v126 := GenS12{F0: genFn0("\n\t", v5, -1), F1: 0}
    for fi127 := 0; fi127 < 20; fi127++ {
        for fi128 := 0; fi128 < 12; fi128++ {
            var v129 string = "hello"
            var v130 int = (v5 << ((1) & 63))
            _ = fi128
            _ = v129
            _ = v130
        }
        v5 += 128
        _ = fi127
    }
    v5 -= v5
    if v3 {
        v3 = (-32768 == 0)
    } else {
        var v131 float64 = 1e10
        v131 = -2.25
        _ = v131
    }
    v5 -= (^(v5))
    fmt.Println("v1:", v1)
    fmt.Println("v119:", v119)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v118
    _ = v119
    _ = v120
    _ = v126
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
