package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64; F1 uint8; F2 int64 }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 bool; F1 string }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint; F1 uint32 }
type GenS9 struct { F0 int64; F1 uint64; F2 uint64 }
type GenS10 struct { F0 int32; F1 int64; F2 bool }
type GenS11 struct { F0 int; F1 int; F2 uint }
func genFn0(p0 uint, p1 string, p2 uint32) uint64 {
    if (255 != 114) {
        var v7 float64 = 1.5
        v8 := []int{-9223372036854775808, -207, -589, -2147483648}
        _ = v7
        _ = v8
    }
    switch -561 {
    case 42:
        v9 := []int{}
        _ = v9
    case -365:
        var v10 uint8 = 42
        _ = v10
    case -1:
        if (float32(3.14159) <= (1.5 * 0.0)) {
            fmt.Println(p0)
            for fi11 := 0; fi11 < 15; fi11++ {
                for fi12 := 0; fi12 < 5; fi12++ {
                    if false {
                        v13 := GenS1{F0: -9223372036854775808}
                        _ = v13
                    }
                    if ((true && false) && true) {
                        v14 := []int{-9223372036854775808, fi11}
                        _ = v14
                    }
                    var v15 int8 = -12
                    _ = fi12
                    _ = v15
                }
                fmt.Println(fi11, fi11, p0)
                v16 := GenS1{F0: 42}
                _ = fi11
                _ = v16
            }
            var v17 uint16 = 42
            _ = v17
        }
    }
    p2 = 4294967295
    for fi18 := 0; fi18 < 3; fi18++ {
        v19 := GenS1{F0: 1}
        _ = fi18
        _ = v19
    }
    v20 := []int{4294967295, -9223372036854775808, 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    return 746
}
func genFn1(p0 int64, p1 int64) int {
    var v21 float64 = 0.0
    var v22 int32 = -2147483648
    var v23 uint64 = 762
    var v24 int32 = (^(v22))
    _ = p0
    _ = p1
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    return 127
}

func main() {
    var v1 uint64 = 937
    v2 := make(map[string]int)
    var v3 int64 = 1
    var v4 uint32 = 601
    var v5 string = ""
    var v6 string = ""
    switch -686 {
    case 2147483647:
        delete(v2, ("hello" + (v5 + "foo bar")))
    case -422:
        switch "a" {
        case "foo bar":
            v25 := []int{9223372036854775807, -9223372036854775808, -1, 845}
            _ = v25
        }
    }
    switch 128 {
    case 2147483647:
        v26 := GenS1{F0: (v3 | 65)}
        _ = v26
    case -9223372036854775808:
        if true {
            switch v5 {
            case "hello":
                for fi27 := 0; fi27 < 6; fi27++ {
                    if false {
                        switch v5 {
                        case "a":
                            switch (fi27 + (fi27 >> ((42) & 63))) {
                            case -346:
                                fmt.Println(v4, v5)
                            }
                        default:
                            var v28 bool = true
                            _ = v28
                        }
                    } else {
                        v29 := GenS2{F0: v3, F1: 211, F2: v3}
                        v5 = ""
                        _ = v29
                    }
                    for fi30 := 0; fi30 < 5; fi30++ {
                        v3 = (^(v3))
                        v31 := GenS1{F0: (v3 + (v3 ^ int64(fi27)))}
                        var v32 string = "a"
                        var v33 uint = 674
                        _ = fi30
                        _ = v31
                        _ = v32
                        _ = v33
                    }
                    var v34 int64 = -496
                    _ = fi27
                    _ = v34
                }
            case "":
                var v35 int64 = v3
                _ = v35
            case "foo bar":
                var v36 int = -9223372036854775808
                _ = v36
            }
            v6 = "\n\t"
        } else {
            v37 := GenS3{F0: 0, F1: genFn0(42, "", v4)}
            var v38 uint64 = v1
            _ = v37
            _ = v38
        }
    case 9223372036854775807:
        v39 := []int{279, -2147483648, genFn1((-(v3)), 0)}
        _ = v39
    default:
        switch 128 {
        case 42:
            v40 := GenS3{F0: 0, F1: (^(v1))}
            _ = v40
        case 127:
            v6 = v6
        }
    }
    if (2147483647 != 0) {
        v6 = v5
        _ = len(v2)
    }
    for fi41 := 0; fi41 < 8; fi41++ {
        v42 := GenS4{F0: (633 != -5), F1: "hello"}
        _ = fi41
        _ = v42
    }
    fmt.Println(v4, v6)
    v43 := GenS1{F0: v3}
    for fi44 := 0; fi44 < 1; fi44++ {
        var v45 int = genFn1(v3, 1)
        switch 1 {
        case -256:
            if (!false) {
                delete(v2, "\n\t")
                v3 -= 42
                for fi46 := 0; fi46 < 18; fi46++ {
                    switch (^(v45)) {
                    case 4294967295:
                        v47 := GenS2{F0: (-(v3)), F1: 121, F2: (-(v3))}
                        _ = v47
                    case 89:
                        v48 := GenS5{F0: (v3 << ((2) & 63))}
                        _ = v48
                    default:
                        if (65535 >= 82) {
                            var v49 uint8 = 128
                            if (!true) {
                                var v50 uint64 = (v1 & 95)
                                v3 += -646
                                fmt.Println(v49)
                                _ = v50
                            } else {
                                v45 -= (fi44 + (^(fi46)))
                                fmt.Println(v6)
                                if false {
                                    if true {
                                        v51 := GenS6{F0: 762}
                                        _ = v51
                                    } else {
                                        var v52 int8 = 42
                                        var v53 int16 = -735
                                        switch 128 {
                                        case 0:
                                            if (896 != -1) {
                                                v45 = v45
                                            } else {
                                                v54 := []int{genFn1((v3 ^ 388), v3), genFn1(1, v3)}
                                                _ = v54
                                            }
                                            fmt.Println(v45)
                                            fmt.Println(fi44, v53)
                                        case 94:
                                            v3 -= 42
                                        default:
                                            var v55 int = (fi44 - 42)
                                            _ = v55
                                        }
                                        _ = v52
                                        _ = v53
                                    }
                                    var v56 uint8 = (v49 - (v49 - 127))
                                    v57 := GenS7{F0: genFn0(4294967295, "foo bar", 65535)}
                                    _ = v56
                                    _ = v57
                                }
                                v2["\n\t"] = fi44
                                v49 -= (v49 << ((18446744073709551615) & 63))
                                v49 -= v49
                            }
                            if ((v1 % ((v1) | 1)) > (v1 << ((1) & 63))) {
                                v58 := GenS7{F0: 660}
                                var v59 int = v45
                                _ = v58
                                _ = v59
                            }
                            if (!((v49 << ((4294967295) & 63)) != v49)) {
                                if (!(-86 < 1)) {
                                    if (0.1 >= 0.1) {
                                        v60 := []int{}
                                        _ = v60
                                    }
                                    v61 := []int{(fi44 >> ((15) & 63))}
                                    _ = v61
                                }
                            } else {
                                var v62 uint = uint(v1)
                                v63 := []int{fi44}
                                switch "hello" {
                                case "":
                                    switch "foo bar" {
                                    case "foo bar":
                                        fmt.Println(v4)
                                    }
                                }
                                v45 -= (fi44 / ((fi44) | 1))
                                _ = v62
                                _ = v63
                            }
                            v4 = 632
                            _ = v49
                        } else {
                            var v64 bool = true
                            var v65 int64 = 1
                            var v66 float32 = (0.1 - 3.14159)
                            _ = v64
                            _ = v65
                            _ = v66
                        }
                    }
                    _ = len(v2)
                    _ = fi46
                }
            } else {
                if (((v6 + "foo bar") + "foo bar") != v6) {
                    v67 := GenS8{F0: 2, F1: (v4 * 321)}
                    _ = v67
                } else {
                    var v68 string = "\n\t"
                    _ = v68
                }
                switch v6 {
                case "hello":
                    switch fi44 {
                    case -526:
                        fmt.Println(fi44)
                    case -355:
                        var v69 uint32 = 0
                        _ = v69
                    case -767:
                        var v70 string = (v5 + "a")
                        _ = v70
                    }
                case "\n\t":
                    for fi71 := 0; fi71 < 8; fi71++ {
                        v6 = "a"
                        v72 := []int{fi71, genFn1(39, v3)}
                        _ = fi71
                        _ = v72
                    }
                case "":
                    if false {
                        for fi73 := 0; fi73 < 10; fi73++ {
                            fmt.Println(v45, fi44)
                            if true {
                                v74 := GenS1{F0: v3}
                                v5 = "foo bar"
                                _ = v74
                            } else {
                                if (42 <= 42) {
                                    var v75 int32 = 42
                                    v76 := GenS3{F0: (v45 & genFn1((v3 >> ((2) & 63)), v3)), F1: genFn0(42, v5, (v4 & v4))}
                                    if true {
                                        v5 = "\n\t"
                                        v77 := GenS5{F0: -754}
                                        switch "foo bar" {
                                        case "":
                                            v3 += (v3 * v3)
                                        default:
                                            fmt.Println(v4, v3)
                                        }
                                        _ = v77
                                    }
                                    v75 = -37
                                    _ = v75
                                    _ = v76
                                }
                                if (!false) {
                                    fmt.Println(fi73)
                                }
                                var v78 uint16 = 985
                                _ = v78
                            }
                            v79 := []int{4294967295, (v45 % ((fi44) | 1))}
                            fmt.Println(fi44)
                            v1 += (v1 + v1)
                            _ = fi73
                            _ = v79
                        }
                    }
                default:
                    v6 = v5
                }
            }
        default:
            switch 0 {
            case -603:
                v45 = -1
            default:
                for fi80 := 0; fi80 < 1; fi80++ {
                    if ((v5 + ("" + "")) > v6) {
                        var v81 string = v6
                        if ("a" > v81) {
                            fmt.Println(v81)
                            _ = len(v2)
                        }
                        if false {
                            v82 := GenS9{F0: -1, F1: 2, F2: (v1 / ((genFn0(uint(v1), (("a" + ("" + ((v5 + v81) + "hello"))) + v81), (^(v4)))) | 1))}
                            _ = v82
                        }
                        _ = v81
                    } else {
                        v83 := GenS8{F0: 759, F1: v4}
                        if ((v4 | v4) > (v4 % ((v4) | 1))) {
                            v1 -= 447
                            delete(v2, v5)
                        }
                        _ = v83
                    }
                    if false {
                        var v84 uint64 = 674
                        _ = v84
                    }
                    _ = fi80
                }
            }
        }
        if (42 > 2) {
            var v85 float32 = (3.14159 / 0.1)
            _ = v85
        } else {
            fmt.Println(v5, v1, v45)
        }
        _ = fi44
        _ = v45
    }
    v6 = "a"
    if (1e10 > (1.5 + 0.0)) {
        v86 := GenS3{F0: -705, F1: v1}
        for fi87 := 0; fi87 < 20; fi87++ {
            var v88 float32 = (1.5 / 1.5)
            var v89 uint16 = 65535
            _ = fi87
            _ = v88
            _ = v89
        }
        _ = v86
    } else {
        v90 := GenS4{F0: (1 >= v3), F1: v5}
        _ = v90
    }
    for fi91 := 0; fi91 < 16; fi91++ {
        var v92 uint64 = v1
        v93 := []int{(fi91 + (fi91 | (fi91 | (^(fi91))))), genFn1(v3, (-(v3))), fi91}
        _ = fi91
        _ = v92
        _ = v93
    }
    var v94 int = genFn1(v3, (-(v3)))
    switch (v94 / ((v94) | 1)) {
    case 32768:
        if (!(((("\n\t" + v6) + "a") != (("\n\t" + v5) + ("hello" + v6))) || (!(1 != 100)))) {
            switch "hello" {
            case "\n\t":
                if false {
                    for fi95 := 0; fi95 < 17; fi95++ {
                        fmt.Println(fi95, v6)
                        _ = fi95
                    }
                    var v96 uint = 0
                    if (119 > 56) {
                        v2[(v5 + v5)] = v94
                    } else {
                        fmt.Println(v3, v1)
                        switch "foo bar" {
                        case "a":
                            var v97 string = v5
                            _ = v97
                        case "hello":
                            var v98 uint8 = 42
                            _ = v98
                        }
                    }
                    _ = v96
                } else {
                    var v99 int32 = 42
                    if (397 == 2) {
                        if true {
                            for fi100 := 0; fi100 < 8; fi100++ {
                                for fi101 := 0; fi101 < 10; fi101++ {
                                    if false {
                                        if false {
                                            v102 := v2[v5]
                                            var v103 bool = false
                                            _ = v102
                                            _ = v103
                                        }
                                    }
                                    var v104 string = (((v5 + v5) + v6) + (v5 + v5))
                                    switch (((("foo bar" + v104) + v104) + "foo bar") + ("\n\t" + (v5 + v6))) {
                                    case "hello":
                                        var v105 string = "foo bar"
                                        _ = v105
                                    case "foo bar":
                                        v106 := []int{(^(fi100)), v94, fi100, fi101}
                                        _ = v106
                                    case "a":
                                        fmt.Println(v1)
                                    }
                                    _ = fi101
                                    _ = v104
                                }
                                _ = fi100
                            }
                        }
                        v107 := GenS3{F0: int(v4), F1: genFn0(4294967295, v5, v4)}
                        var v108 uint = 885
                        _ = v107
                        _ = v108
                    }
                    for fi109 := 0; fi109 < 19; fi109++ {
                        var v110 string = "a"
                        _ = fi109
                        _ = v110
                    }
                    _ = v99
                }
            }
            var v111 uint8 = 2
            _ = v111
        }
    case 1:
        var v112 uint8 = 255
        _ = v112
    case 194:
        if (2 == 573) {
            if ((v6 + v6) <= v6) {
                v6 = ((("" + v6) + (v6 + v5)) + v5)
                v3 += v3
                for fi113 := 0; fi113 < 8; fi113++ {
                    fmt.Println(v3, v94, v4)
                    _ = fi113
                }
            }
            if false {
                var v114 uint = uint(v4)
                for fi115 := 0; fi115 < 3; fi115++ {
                    v6 = ""
                    v2[(v5 + (v6 + "hello"))] = 601
                    _ = fi115
                }
                v3 -= v3
                _ = v114
            }
            for fi116 := 0; fi116 < 11; fi116++ {
                switch v6 {
                case "\n\t":
                    var v117 uint32 = (v4 + v4)
                    _ = v117
                case "hello":
                    v2[v5] = (fi116 / (((fi116 ^ fi116)) | 1))
                case "":
                    var v118 int32 = 2147483647
                    _ = v118
                default:
                    if true {
                        switch (v94 ^ (-(v94))) {
                        case -2147483648:
                            v2[v5] = (^(fi116))
                        case 1:
                            var v119 int16 = -32768
                            _ = v119
                        default:
                            switch 2147483647 {
                            case 9223372036854775807:
                                v120 := []int{788, genFn1((v3 & v3), (v3 % ((-1) | 1)))}
                                _ = v120
                            case -402:
                                v1 += (v1 >> ((4294967295) & 63))
                            default:
                                for fi121 := 0; fi121 < 13; fi121++ {
                                    var v122 int16 = 32767
                                    var v123 uint8 = 255
                                    v5 = "foo bar"
                                    _ = fi121
                                    _ = v122
                                    _ = v123
                                }
                            }
                        }
                    }
                }
                _ = fi116
            }
        }
    }
    for fi124 := 0; fi124 < 3; fi124++ {
        v125 := GenS10{F0: -450, F1: (v3 | v3), F2: true}
        if true {
            v126 := GenS9{F0: (v3 / ((-237) | 1)), F1: v1, F2: (^(v1))}
            _ = v126
        }
        _ = fi124
        _ = v125
    }
    v94 += v94
    v127 := GenS1{F0: v3}
    for fi128 := 0; fi128 < 1; fi128++ {
        v129 := GenS11{F0: (^(fi128)), F1: (v94 >> ((4294967295) & 63)), F2: 42}
        v6 = "hello"
        switch 127 {
        case 128:
            v4 = v4
        }
        _ = fi128
        _ = v129
    }
    var v130 uint16 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v130:", v130)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v94:", v94)
    _ = v1
    _ = v127
    _ = v130
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
    _ = v94
}
