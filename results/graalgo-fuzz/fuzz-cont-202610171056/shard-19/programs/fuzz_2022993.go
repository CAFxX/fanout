package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 bool; F1 uint16; F2 int8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int16; F1 uint16 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 int64; F1 int8; F2 uint32 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 uint8; F1 uint64; F2 uint }
type GenS12 struct { F0 int8 }
func genFn0() int32 {
    for fi4 := 0; fi4 < 15; fi4++ {
        var v5 string = "hello"
        for fi6 := 0; fi6 < 7; fi6++ {
            var v7 uint16 = 0
            if false {
                fmt.Println(v5)
            }
            _ = fi6
            _ = v7
        }
        v8 := GenS1{F0: 2, F1: 404, F2: 854}
        _ = fi4
        _ = v5
        _ = v8
    }
    if ((1 < -1) && false) {
        if ((9223372036854775807 <= 0) || false) {
            v9 := GenS1{F0: 176, F1: 4294967295, F2: 790}
            _ = v9
        } else {
            v10 := GenS1{F0: 666, F1: 741, F2: 2147483648}
            var v11 int16 = 1000
            v12 := GenS1{F0: 0, F1: 42, F2: 766}
            _ = v10
            _ = v11
            _ = v12
        }
        var v13 uint8 = 166
        _ = v13
    }
    return -2147483648
}
func genFn1(p0 string, p1 uint) bool {
    var v14 uint = (p1 % ((p1) | 1))
    var v15 float32 = 3.14159
    v14 += p1
    v16 := GenS2{F0: -1}
    switch 42 {
    case 86:
        p1 = (p1 - (p1 - v14))
    case 1:
        v17 := GenS3{F0: 628}
        _ = v17
    }
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    _ = v16
    return false
}

func main() {
    var v1 float32 = 0.1
    v2 := []int{680, 42}
    var v3 int8 = -128
    v1 += 0.0
    var v18 string = "\n\t"
    if true {
        switch "\n\t" {
        case "hello":
            v1 += 0.1
        default:
            var v19 int64 = 42
            _ = v19
        }
    }
    if (42 == 582) {
        for fi20 := 0; fi20 < 17; fi20++ {
            v21 := GenS4{F0: ((v3 << ((369) & 63)) <= 1), F1: 42, F2: v3}
            _ = fi20
            _ = v21
        }
        fmt.Println(v3, v18)
    } else {
        var v22 uint32 = 0
        _ = v22
    }
    v3 = (v3 & v3)
    v2 = append(v2, -1)
    switch ("a" + "foo bar") {
    case "hello":
        for fi23 := 0; fi23 < 14; fi23++ {
            v1 = 1e10
            if false {
                v18 = (v18 + v18)
                if (0.0 != (0.0 / 1.5)) {
                    switch "\n\t" {
                    case "foo bar":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    default:
                        v3 -= v3
                    }
                    if true {
                        var v24 float32 = v1
                        v1 -= v24
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = v24
                    } else {
                        v2 = append(v2, (fi23 | fi23))
                        for fi25 := 0; fi25 < 5; fi25++ {
                            fmt.Println(v18, v18, fi23)
                            if (2147483648 <= 1) {
                                if false {
                                    fmt.Println(v3, fi23)
                                    if true {
                                        switch "foo bar" {
                                        case "hello":
                                            if (!(!true)) {
                                                v26 := GenS5{F0: 20}
                                                var v27 int64 = 508
                                                _ = v26
                                                _ = v27
                                            }
                                        }
                                        if false {
                                            v28 := GenS3{F0: 519}
                                            switch v18 {
                                            case "a":
                                                var v29 int32 = genFn0()
                                                _ = v29
                                            case "\n\t":
                                                if false {
                                                    switch 4294967295 {
                                                    case -9223372036854775808:
                                                        if ((!false) || true) {
                                                            v30 := GenS3{F0: 0}
                                                            v2 = append(v2, (fi25 - -21))
                                                            _ = v30
                                                        } else {
                                                            var v31 int16 = 0
                                                            v31 = (v31 | 379)
                                                            fmt.Println(fi23, v3, v3)
                                                            _ = v31
                                                        }
                                                    case 4294967295:
                                                        v32 := GenS2{F0: -128}
                                                        _ = v32
                                                    case 166:
                                                        v3 = -1
                                                    }
                                                    v18 = ""
                                                    fmt.Println(fi25, fi25)
                                                } else {
                                                    var v33 int32 = 2147483647
                                                    v3 -= 100
                                                    v34 := GenS6{F0: (v3 << ((65535) & 63))}
                                                    _ = v33
                                                    _ = v34
                                                }
                                            case "foo bar":
                                                var v35 int8 = 29
                                                _ = v35
                                            default:
                                                v36 := GenS3{F0: 2147483647}
                                                var v37 uint = 18446744073709551615
                                                var v38 float64 = 1e10
                                                _ = v36
                                                _ = v37
                                                _ = v38
                                            }
                                            var v39 bool = false
                                            _ = len(v2)
                                            var v40 int32 = int32(fi23)
                                            var v41 int32 = 42
                                            if v39 {
                                                var v42 int16 = -1
                                                _ = v42
                                            } else {
                                                if true {
                                                    var v43 int16 = 32767
                                                    if false {
                                                        _ = len(v2)
                                                        fmt.Println(v3)
                                                        switch "\n\t" {
                                                        case "\n\t":
                                                            var v44 uint64 = 4294967295
                                                            var v45 int8 = -1
                                                            _ = v44
                                                            _ = v45
                                                        default:
                                                            var v46 int8 = (v3 ^ (v3 % ((v3) | 1)))
                                                            _ = v46
                                                        }
                                                    }
                                                    _ = v43
                                                } else {
                                                    _ = len(v2)
                                                }
                                                var v47 string = v18
                                                var v48 string = v18
                                                _ = v47
                                                _ = v48
                                            }
                                            _ = v28
                                            _ = v39
                                            _ = v40
                                            _ = v41
                                        }
                                    }
                                    if false {
                                        if (!false) {
                                            if false {
                                                if true {
                                                    v49 := GenS1{F0: 2, F1: 4294967295, F2: 147}
                                                    _ = len(v2)
                                                    if (42 > 169) {
                                                        if (true || false) {
                                                            if false {
                                                                if (9223372036854775807 == 294) {
                                                                    v2 = append(v2, 42)
                                                                    v50 := GenS1{F0: 4294967295, F1: 282, F2: 2}
                                                                    switch (fi25 % ((fi23) | 1)) {
                                                                    case 1:
                                                                        if (true && false) {
                                                                            switch (fi23 + 9223372036854775807) {
                                                                            case 127:
                                                                                switch "\n\t" {
                                                                                case "":
                                                                                    v51 := GenS4{F0: ((42 != 592) && (0 >= 9223372036854775807)), F1: 401, F2: (v3 + 6)}
                                                                                    _ = v51
                                                                                case "a":
                                                                                    fmt.Println(v18, fi23)
                                                                                    v52 := GenS4{F0: (65535 != 255), F1: 764, F2: 1}
                                                                                    if len(v2) >= 2 {
                                                                                        v2 = v2[1:len(v2)]
                                                                                    }
                                                                                    v2 = append(v2, (fi25 >> ((65535) & 63)))
                                                                                    _ = v52
                                                                                case "foo bar":
                                                                                    v2 = append(v2, fi25)
                                                                                }
                                                                            }
                                                                            var v53 int = fi23
                                                                            _ = v53
                                                                        }
                                                                        var v54 int = (-(fi25))
                                                                        _ = v54
                                                                    case 9223372036854775807:
                                                                        if (v18 >= v18) {
                                                                            v1 -= 1e10
                                                                            switch fi23 {
                                                                            case -693:
                                                                                v1 = float32(-2.25)
                                                                            case 4294967295:
                                                                                v55 := GenS1{F0: 700, F1: 2, F2: 0}
                                                                                _ = v55
                                                                            }
                                                                        } else {
                                                                            v18 = (v18 + "\n\t")
                                                                            v56 := GenS3{F0: 42}
                                                                            _ = v56
                                                                        }
                                                                    case -2147483648:
                                                                        var v57 string = "foo bar"
                                                                        _ = v57
                                                                    }
                                                                    _ = v50
                                                                } else {
                                                                    fmt.Println(v1)
                                                                    v1 += -2.25
                                                                }
                                                            }
                                                        }
                                                    } else {
                                                        v2 = append(v2, fi23)
                                                        fmt.Println(fi23, v18)
                                                        var v58 float32 = v1
                                                        if (2147483647 == (fi23 % ((9223372036854775807) | 1))) {
                                                            if false {
                                                                v59 := GenS7{F0: 0, F1: 7}
                                                                _ = v59
                                                            }
                                                            v60 := GenS8{F0: 95}
                                                            v18 = v18
                                                            v61 := GenS4{F0: true, F1: 763, F2: v3}
                                                            var v62 int16 = 32767
                                                            fmt.Println(v3, v1, fi25)
                                                            _ = v60
                                                            _ = v61
                                                            _ = v62
                                                        } else {
                                                            var v63 uint64 = 42
                                                            _ = v63
                                                        }
                                                        fmt.Println(v18, fi23)
                                                        var v64 uint8 = 153
                                                        v1 = v1
                                                        v64 += 102
                                                        switch (v18 + v18) {
                                                        case "foo bar":
                                                            var v65 bool = true
                                                            _ = v65
                                                        }
                                                        var v66 uint8 = (v64 / ((v64) | 1))
                                                        fmt.Println(fi25, v3)
                                                        _ = v58
                                                        _ = v64
                                                        _ = v66
                                                    }
                                                    v3 = -86
                                                    _ = v49
                                                } else {
                                                    switch "foo bar" {
                                                    case "":
                                                        switch fi23 {
                                                        case -9223372036854775808:
                                                            var v67 int32 = -1
                                                            _ = v67
                                                        }
                                                    case "a":
                                                        var v68 int32 = genFn0()
                                                        _ = v68
                                                    case "hello":
                                                        switch 127 {
                                                        case -393:
                                                            v69 := GenS9{F0: 1, F1: v3, F2: 2147483647}
                                                            _ = v69
                                                        default:
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        }
                                                    }
                                                    var v70 float64 = 1.5
                                                    _ = v70
                                                }
                                            } else {
                                                _ = len(v2)
                                            }
                                            v18 = v18
                                            if true {
                                                var v71 float64 = 0.1
                                                v2 = append(v2, fi25)
                                                _ = v71
                                            }
                                        } else {
                                            var v72 int32 = 0
                                            var v73 float64 = 1.5
                                            v74 := GenS8{F0: 42}
                                            if len(v2) > 0 {
                                                v2[0] = int(v3)
                                            }
                                            _ = v72
                                            _ = v73
                                            _ = v74
                                        }
                                        v18 = "foo bar"
                                        var v75 uint64 = 2
                                        _ = v75
                                    } else {
                                        switch (("\n\t" + (v18 + v18)) + ("\n\t" + ("a" + ("" + v18)))) {
                                        case "a":
                                            var v76 float32 = (-2.25 / 1e10)
                                            v76 -= -2.25
                                            var v77 uint8 = 2
                                            _ = v76
                                            _ = v77
                                        case "\n\t":
                                            if true {
                                                fmt.Println(fi23)
                                                switch v18 {
                                                case "\n\t":
                                                    switch "a" {
                                                    case "foo bar":
                                                        var v78 int8 = v3
                                                        var v79 string = v18
                                                        _ = v78
                                                        _ = v79
                                                    }
                                                default:
                                                    var v80 int8 = (v3 << ((1) & 63))
                                                    _ = v80
                                                }
                                                switch -1 {
                                                case 1:
                                                    var v81 uint = 42
                                                    _ = v81
                                                default:
                                                    var v82 uint64 = 2
                                                    _ = v82
                                                }
                                            }
                                        }
                                        switch ("" + v18) {
                                        case "a":
                                            fmt.Println(v3, fi25, v18)
                                        case "foo bar":
                                            v3 += 0
                                        case "":
                                            v1 -= 3.14159
                                        }
                                        v83 := GenS9{F0: -728, F1: (v3 + v3), F2: 420}
                                        _ = v83
                                    }
                                    var v84 string = (((v18 + (v18 + "\n\t")) + (((v18 + v18) + v18) + v18)) + (v18 + v18))
                                    var v85 int64 = -408
                                    if true {
                                        _ = len(v2)
                                    }
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    _ = v84
                                    _ = v85
                                }
                                fmt.Println(v18)
                                var v86 uint8 = 2
                                _ = v86
                            }
                            _ = fi25
                        }
                    }
                    var v87 uint8 = 152
                    _ = v87
                } else {
                    v18 = "\n\t"
                    var v88 int = fi23
                    if false {
                        switch (fi23 * (fi23 << ((937) & 63))) {
                        case 32767:
                            switch "foo bar" {
                            case "\n\t":
                                var v89 int8 = v3
                                _ = v89
                            case "":
                                var v90 uint32 = 966
                                _ = v90
                            default:
                                var v91 float64 = ((1.5 - ((0.1 - ((float64(0.1) + -2.25) * float64(3.14159))) + 0.0)) + 0.0)
                                _ = v91
                            }
                        }
                    } else {
                        v1 = float32((0.0 / -2.25))
                        if ((fi23 >> ((18446744073709551615) & 63)) <= (fi23 | v88)) {
                            for fi92 := 0; fi92 < 13; fi92++ {
                                var v93 uint32 = 2
                                v3 -= 100
                                v94 := GenS9{F0: 1, F1: int8(v93), F2: 0}
                                v95 := GenS4{F0: (false || true), F1: 812, F2: (^(v3))}
                                _ = fi92
                                _ = v93
                                _ = v94
                                _ = v95
                            }
                            var v96 int16 = 1
                            _ = v96
                        }
                        var v97 uint64 = 312
                        _ = v97
                    }
                    _ = v88
                }
                if true {
                    v2 = append(v2, fi23)
                    v18 = ((v18 + "a") + (((v18 + v18) + v18) + v18))
                }
            }
            fmt.Println(v18)
            _ = fi23
        }
    }
    switch -9223372036854775808 {
    case 32768:
        v3 = 1
    default:
        if false {
            if (true && false) {
                if (!((!false) && (65535 != 1))) {
                    v98 := GenS10{F0: true}
                    switch "" {
                    case "hello":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    default:
                        fmt.Println(v3)
                    }
                    _ = v98
                }
                v1 += 1.5
            } else {
                v99 := GenS2{F0: -128}
                switch v18 {
                case "a":
                    var v100 uint8 = 2
                    _ = v100
                case "\n\t":
                    for fi101 := 0; fi101 < 1; fi101++ {
                        if true {
                            v3 = v3
                            v1 += float32(3.14159)
                            v18 = v18
                        }
                        if (13 <= 255) {
                            if false {
                                v18 = v18
                                v3 = v3
                                if false {
                                    v102 := GenS11{F0: 2, F1: 4294967295, F2: 651}
                                    _ = v102
                                }
                            } else {
                                switch "a" {
                                case "foo bar":
                                    var v103 float32 = v1
                                    _ = v103
                                case "":
                                    _ = len(v2)
                                case "a":
                                    switch (((v18 + v18) + "a") + v18) {
                                    case "":
                                        if true {
                                            fmt.Println(v3, v3)
                                        }
                                    case "a":
                                        v3 = 127
                                    }
                                }
                            }
                        }
                        _ = fi101
                    }
                case "foo bar":
                    v18 = v18
                }
                for fi104 := 0; fi104 < 17; fi104++ {
                    v105 := GenS3{F0: 0}
                    v106 := GenS9{F0: 23, F1: 42, F2: 2}
                    v3 = -92
                    _ = fi104
                    _ = v105
                    _ = v106
                }
                _ = v99
            }
            if (65535 == 2) {
                fmt.Println(v18, v18)
            }
            var v107 float32 = (v1 * v1)
            _ = v107
        } else {
            v108 := GenS9{F0: 1, F1: v3, F2: 466}
            _ = v108
        }
    }
    v1 = v1
    for fi109 := 0; fi109 < 10; fi109++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi109
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if false {
        _ = len(v2)
        if (true && false) {
            var v110 int = -624
            _ = v110
        } else {
            v2 = append(v2, 914)
        }
    } else {
        if true {
            var v111 string = v18
            v112 := GenS12{F0: (^(v3))}
            _ = v111
            _ = v112
        } else {
            v113 := GenS1{F0: 42, F1: 0, F2: 2147483648}
            v18 = (v18 + v18)
            var v114 bool = (1 <= -746)
            _ = v113
            _ = v114
        }
    }
    var v115 int16 = 0
    v18 = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("v115:", v115)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v115
    _ = v18
    _ = v2
    _ = v3
}
