package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 bool }
type GenS2 struct { F0 uint16; F1 uint32; F2 uint32 }
type GenS3 struct { F0 int8; F1 string; F2 string }
type GenS4 struct { F0 uint32; F1 bool; F2 int }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint8; F1 uint }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 bool; F1 int16 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 int8; F1 int16; F2 uint32 }
type GenS11 struct { F0 uint32; F1 int8; F2 uint }
type GenS12 struct { F0 bool; F1 uint16 }
type GenS13 struct { F0 uint8; F1 int8; F2 int32 }
func genFn0(p0 bool, p1 int32) string {
    var v7 float32 = 1.5
    if (2147483647 != 42) {
        v7 += v7
    } else {
        p1 -= p1
        v8 := GenS1{F0: 0, F1: 18446744073709551615, F2: p0}
        v7 -= v7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v7
    return "foo bar"
}
func genFn1(p0 uint8, p1 uint16, p2 uint) int {
    var v9 uint32 = 2
    p1 -= p1
    p0 = (p0 & p0)
    p1 += p1
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return -388
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]int)
    var v3 uint16 = 2
    var v4 int16 = 32767
    var v5 int16 = (^(v4))
    var v6 uint64 = 4294967295
    for fi10 := 0; fi10 < 17; fi10++ {
        for fi11 := 0; fi11 < 11; fi11++ {
            switch ("a" + "hello") {
            case "\n\t":
                v12 := []int{(fi11 ^ fi11), (fi11 % ((genFn1(42, v3, 42)) | 1))}
                _ = v12
            default:
                v13 := []int{(fi10 >> ((1) & 63)), (fi10 - (fi10 << ((2) & 63))), (fi10 & fi11), 127}
                _ = v13
            }
            if false {
                var v14 string = ("hello" + "\n\t")
                if false {
                    v3 = (v3 ^ v3)
                    var v15 uint = 0
                    _ = v15
                } else {
                    switch (fi10 - fi11) {
                    case -140:
                        var v16 float32 = 0.0
                        _ = v16
                    case 32768:
                        switch (fi10 & fi11) {
                        case -507:
                            var v17 uint64 = v6
                            _ = v17
                        case 4294967295:
                            fmt.Println(fi11, v5)
                        }
                    }
                    delete(v1, v14)
                    var v18 float64 = ((1e10 + 1e10) / -2.25)
                    v6 = 958
                    var v19 int16 = v5
                    _ = v18
                    _ = v19
                }
                var v20 uint32 = 898
                _ = v14
                _ = v20
            } else {
                v6 = v6
            }
            _ = fi11
        }
        v21 := []int{(fi10 + 4294967295), (fi10 % ((genFn1(127, v3, 603)) | 1))}
        for fi22 := 0; fi22 < 4; fi22++ {
            v6 -= v6
            switch fi10 {
            case -263:
                switch ("\n\t" + ((("hello" + "a") + "\n\t") + ("foo bar" + "\n\t"))) {
                case "":
                    fmt.Println(v5, v4, v6)
                case "\n\t":
                    v4 -= (v5 & (v4 / ((v4) | 1)))
                case "foo bar":
                    v5 = (v5 / (((v4 & v4)) | 1))
                }
            case -1:
                switch ("hello" + ("a" + "\n\t")) {
                case "\n\t":
                    var v23 int8 = 7
                    _ = v23
                case "":
                    v21 = append(v21, genFn1(42, 2, 752))
                case "foo bar":
                    v21 = append(v21, 4294967295)
                }
            case -2147483648:
                switch ("foo bar" + "") {
                case "":
                    var v24 int32 = -2147483648
                    if true {
                        fmt.Println(fi22, v5, v4)
                    }
                    _ = v24
                case "\n\t":
                    v4 += int16(v3)
                case "a":
                    v1["a"] = 0
                default:
                    switch fi22 {
                    case 452:
                        var v25 uint = 18446744073709551615
                        _ = v25
                    case -1:
                        var v26 uint32 = 4294967295
                        _ = v26
                    default:
                        v27 := v2[-9223372036854775808]
                        if true {
                            var v28 int = fi22
                            v4 += (v4 >> ((65535) & 63))
                            var v29 uint8 = 128
                            _ = v28
                            _ = v29
                        } else {
                            v30 := GenS1{F0: 191, F1: (v6 / (((v6 * (v6 | (v6 & v6)))) | 1)), F2: false}
                            if false {
                                v5 = (v4 % ((v4) | 1))
                                if (-1 >= -32768) {
                                    var v31 int8 = 0
                                    switch (v27 % ((fi22) | 1)) {
                                    case 42:
                                        if (!(65535 > 1)) {
                                            var v32 int8 = (v31 + (^(v31)))
                                            _ = v32
                                        } else {
                                            if false {
                                                v31 += v31
                                                if false {
                                                    if false {
                                                        var v33 int32 = -274
                                                        v6 += (v6 << ((71) & 63))
                                                        v4 += (v5 + 42)
                                                        _ = v33
                                                    } else {
                                                        v34 := GenS1{F0: v6, F1: v6, F2: ((!false) && (-1 > -410))}
                                                        _ = v34
                                                    }
                                                    if true {
                                                        switch "a" {
                                                        case "a":
                                                            v6 -= v6
                                                        default:
                                                            var v35 bool = ((v31 % ((v31) | 1)) == v31)
                                                            _ = v35
                                                        }
                                                    } else {
                                                        if true {
                                                            v21 = append(v21, v27)
                                                        }
                                                    }
                                                }
                                                var v36 uint = 330
                                                _ = v36
                                            }
                                            switch (fi10 - 42) {
                                            case 127:
                                                if ((v3 & (v3 & (v3 & (v3 * v3)))) > v3) {
                                                    fmt.Println(fi10, v4)
                                                    var v37 uint32 = 671
                                                    switch v27 {
                                                    case 4294967295:
                                                        v38 := GenS1{F0: (v6 | (v6 - 2)), F1: v6, F2: ((v3 & 42) <= (v3 / ((v3) | 1)))}
                                                        _ = v38
                                                    case 0:
                                                        v27 = fi22
                                                    default:
                                                        var v39 bool = (42 == 9223372036854775807)
                                                        _ = v39
                                                    }
                                                    _ = v37
                                                }
                                            }
                                            if true {
                                                switch -1 {
                                                case -1:
                                                    v3 -= (v3 ^ v3)
                                                case 772:
                                                    if (float32(0.1) > (0.1 + 0.0)) {
                                                        v2[489] = genFn1(225, v3, 696)
                                                        v31 = (v31 - (-(v31)))
                                                        v40 := v1["\n\t"]
                                                        var v41 string = ("a" + "hello")
                                                        var v42 uint16 = (v3 / ((v3) | 1))
                                                        v43 := v1[(("foo bar" + v41) + ((("hello" + v41) + v41) + v41))]
                                                        _ = v40
                                                        _ = v41
                                                        _ = v42
                                                        _ = v43
                                                    }
                                                default:
                                                    v2[(fi22 % ((v27) | 1))] = fi10
                                                }
                                            }
                                        }
                                    case -9223372036854775808:
                                        if false {
                                            switch ((("" + "\n\t") + "\n\t") + "\n\t") {
                                            case "":
                                                if true {
                                                    if true {
                                                        v31 = v31
                                                    }
                                                    var v44 int64 = -643
                                                    _ = v44
                                                } else {
                                                    fmt.Println(fi22, v27)
                                                    fmt.Println(v31, v4, v4)
                                                    fmt.Println(v27, v6)
                                                }
                                                var v45 uint = 1
                                                if true {
                                                    v46 := GenS2{F0: (^(v3)), F1: 1, F2: 2147483648}
                                                    fmt.Println(v31, fi22)
                                                    v45 += 18446744073709551615
                                                    fmt.Println(fi10, fi22)
                                                    _ = v46
                                                } else {
                                                    v31 -= (v31 / ((1) | 1))
                                                    if (false && (-2.25 >= 1.5)) {
                                                        if true {
                                                            switch "\n\t" {
                                                            case "":
                                                                switch 32768 {
                                                                case 0:
                                                                    fmt.Println(v31, v27)
                                                                case -767:
                                                                    v3 -= (v3 & v3)
                                                                default:
                                                                    switch (("foo bar" + "") + "") {
                                                                    case "a":
                                                                        v47 := GenS3{F0: (^(v31)), F1: "", F2: "\n\t"}
                                                                        _ = v47
                                                                    }
                                                                }
                                                                v1["\n\t"] = 2
                                                                fmt.Println(fi10, v4, fi10)
                                                            default:
                                                                var v48 uint = (v45 >> ((v45) & 63))
                                                                _ = v48
                                                            }
                                                            fmt.Println(fi22)
                                                            v49 := GenS4{F0: 2, F1: (true && (false || true)), F2: (fi10 ^ genFn1(92, v3, 1))}
                                                            var v50 bool = false
                                                            v51 := GenS2{F0: v3, F1: 65535, F2: 579}
                                                            if len(v21) >= 2 {
                                                                v21 = v21[1:len(v21)]
                                                            }
                                                            v52 := GenS5{F0: (-(v31))}
                                                            _ = v49
                                                            _ = v50
                                                            _ = v51
                                                            _ = v52
                                                        }
                                                        fmt.Println(v5, v6, fi10)
                                                    } else {
                                                        if len(v21) >= 2 {
                                                            v21 = v21[1:len(v21)]
                                                        }
                                                        v53 := GenS3{F0: (^(v31)), F1: "", F2: "a"}
                                                        _ = v53
                                                    }
                                                    if true {
                                                        v54 := GenS6{F0: 255, F1: v45}
                                                        var v55 string = "hello"
                                                        fmt.Println(fi22, fi22)
                                                        _ = v54
                                                        _ = v55
                                                    } else {
                                                        v56 := GenS7{F0: 255}
                                                        v4 -= v5
                                                        var v57 int32 = 895
                                                        var v58 float32 = 3.14159
                                                        v59 := GenS2{F0: (^(v3)), F1: 669, F2: 781}
                                                        var v60 int32 = (v57 & 1)
                                                        _ = v56
                                                        _ = v57
                                                        _ = v58
                                                        _ = v59
                                                        _ = v60
                                                    }
                                                }
                                                _ = v45
                                            }
                                            var v61 int8 = v31
                                            _ = v61
                                        } else {
                                            v62 := GenS8{F0: false, F1: v4}
                                            if len(v21) > 0 {
                                                v21[0] = 32768
                                            }
                                            if (!(-2.25 > (1.5 + (0.0 + 0.1)))) {
                                                var v63 uint16 = (^(v3))
                                                v3 = v63
                                                switch "hello" {
                                                case "hello":
                                                    var v64 uint64 = 18446744073709551615
                                                    _ = v64
                                                case "a":
                                                    fmt.Println(v5, v63)
                                                case "foo bar":
                                                    v27 -= v27
                                                default:
                                                    v63 = (^(v63))
                                                }
                                                fmt.Println(fi10, v27, fi10)
                                                if false {
                                                    if false {
                                                        var v65 uint64 = 1
                                                        _ = v65
                                                    } else {
                                                        v66 := GenS9{F0: 0}
                                                        _ = v66
                                                    }
                                                    var v67 string = ("a" + "foo bar")
                                                    _ = v67
                                                } else {
                                                    _ = len(v1)
                                                    v31 = 100
                                                }
                                                _ = v63
                                            }
                                            _ = v62
                                        }
                                    default:
                                        v68 := GenS6{F0: 127, F1: 2}
                                        _ = v68
                                    }
                                    v1["hello"] = 2147483647
                                    _ = v31
                                }
                                switch genFn1(65, v3, 865) {
                                case 32768:
                                    v69 := GenS4{F0: 2147483648, F1: (v6 > 42), F2: (v27 | fi22)}
                                    if (0.1 < 1e10) {
                                        var v70 int32 = 42
                                        _ = v70
                                    } else {
                                        if ((((1.5 - 1e10) / 1.0) / 1.5) >= 0.1) {
                                            v71 := GenS2{F0: v3, F1: 2, F2: 2147483648}
                                            if false {
                                                var v72 uint64 = (v6 & (v6 << ((375) & 63)))
                                                _ = v72
                                            } else {
                                                if (0.0 < ((-2.25 / 0.1) * float32(-2.25))) {
                                                    if false {
                                                        var v73 float32 = ((1.5 / ((0.0 * float32((3.14159 * 0.0))) - float32(3.14159))) / 3.14159)
                                                        _ = v73
                                                    } else {
                                                        switch v27 {
                                                        case 803:
                                                            if true {
                                                                _ = len(v21)
                                                                if false {
                                                                    var v74 uint8 = 0
                                                                    _ = v74
                                                                }
                                                                var v75 uint = 0
                                                                _ = v75
                                                            }
                                                        case 2147483647:
                                                            delete(v1, "a")
                                                        case -9223372036854775808:
                                                            fmt.Println(fi10)
                                                        }
                                                        var v76 float64 = float64((float32(((1.5 + 1.5) - 1.5)) * (float32(-2.25) / 1.0)))
                                                        _ = v76
                                                    }
                                                    var v77 int = v27
                                                    v3 = (v3 & 2)
                                                    _ = v77
                                                } else {
                                                    v78 := GenS2{F0: (v3 * (^(v3))), F1: 4294967295, F2: 617}
                                                    _ = v78
                                                }
                                                var v79 uint32 = 738
                                                _ = v79
                                            }
                                            var v80 int64 = 0
                                            _ = v71
                                            _ = v80
                                        }
                                        var v81 int64 = 0
                                        _ = v81
                                    }
                                    _ = v69
                                case 184:
                                    if true {
                                        var v82 int8 = 44
                                        _ = v82
                                    } else {
                                        v6 = v6
                                        var v83 int64 = -605
                                        switch genFn1(93, 255, 496) {
                                        case -1:
                                            if len(v21) > 0 {
                                                v21[0] = fi10
                                            }
                                        case 184:
                                            switch (^(fi22)) {
                                            case 32767:
                                                switch ("a" + "a") {
                                                case "foo bar":
                                                    var v84 uint16 = (v3 - 255)
                                                    _ = v84
                                                case "hello":
                                                    fmt.Println(v5)
                                                default:
                                                    v5 = (v5 / ((v4) | 1))
                                                }
                                            case -405:
                                                v4 = (^(v5))
                                            }
                                        case 127:
                                            v5 = v5
                                        }
                                        _ = v83
                                    }
                                case -2147483648:
                                    var v85 uint64 = 0
                                    _ = v85
                                default:
                                    var v86 uint32 = 148
                                    _ = v86
                                }
                            }
                            _ = v30
                        }
                        var v87 int32 = 2147483647
                        _ = v27
                        _ = v87
                    }
                    v4 -= -1
                }
            default:
                var v88 int64 = 9223372036854775807
                fmt.Println(v3)
                if (!(false || false)) {
                    var v89 int = 0
                    v3 = v3
                    var v90 float64 = (0.0 + 1.5)
                    var v91 float32 = (3.14159 - 1.5)
                    switch "a" {
                    case "":
                        fmt.Println(v90)
                    case "foo bar":
                        switch "foo bar" {
                        case "hello":
                            v92 := GenS8{F0: true, F1: -1}
                            _ = v92
                        case "":
                            v91 = v91
                        case "a":
                            v93 := GenS10{F0: 127, F1: v4, F2: 2}
                            v94 := GenS7{F0: 230}
                            _ = v93
                            _ = v94
                        default:
                            v3 = (v3 / ((1) | 1))
                        }
                    case "\n\t":
                        var v95 int64 = v88
                        _ = v95
                    }
                    var v96 int64 = v88
                    v97 := GenS4{F0: 1, F1: true, F2: 9223372036854775807}
                    _ = v89
                    _ = v90
                    _ = v91
                    _ = v96
                    _ = v97
                }
                if true {
                    fmt.Println(fi22, v88, v4)
                    v4 -= v4
                    v88 += -1
                }
                v5 += (-(v4))
                _ = len(v1)
                v1["a"] = 2147483648
                _ = v88
            }
            _ = fi22
        }
        _ = fi10
        _ = v21
    }
    v98 := GenS8{F0: (false || (true && (!false))), F1: v5}
    if (((v3 ^ v3) == 629) || false) {
        for fi99 := 0; fi99 < 17; fi99++ {
            var v100 string = ""
            _ = fi99
            _ = v100
        }
        var v101 int = 0
        _ = v101
    }
    for fi102 := 0; fi102 < 18; fi102++ {
        v103 := GenS11{F0: 2147483648, F1: -1, F2: 1}
        _ = fi102
        _ = v103
    }
    var v104 uint8 = 29
    var v105 uint8 = 0
    var v106 float32 = 0.1
    v107 := []int{-9223372036854775808, -2147483648, 149}
    for fi108 := 0; fi108 < 5; fi108++ {
        _ = len(v107)
        _ = len(v107)
        switch "hello" {
        case "\n\t":
            v109 := v1["hello"]
            _ = v109
        case "":
            v107 = append(v107, (fi108 | genFn1(v105, (v3 % ((v3) | 1)), 607)))
        }
        _ = fi108
    }
    v104 = (v104 / ((127) | 1))
    for fi110 := 0; fi110 < 5; fi110++ {
        v111 := GenS12{F0: (!(2147483648 == 678)), F1: 65535}
        _ = fi110
        _ = v111
    }
    if false {
        var v112 string = "foo bar"
        _ = v112
    } else {
        fmt.Println(v6)
        fmt.Println(v105, v106, v4)
        var v113 string = ("a" + ("" + (("hello" + "a") + "foo bar")))
        _ = v113
    }
    switch "\n\t" {
    case "\n\t":
        var v114 uint32 = 4294967295
        _ = v114
    case "foo bar":
        var v115 uint16 = 577
        _ = v115
    case "":
        v5 = (v4 / (((v4 / ((v4) | 1))) | 1))
    }
    for fi116 := 0; fi116 < 16; fi116++ {
        switch "" {
        case "\n\t":
            var v117 uint8 = (^(v105))
            _ = v117
        case "a":
            v118 := GenS13{F0: 227, F1: 0, F2: -2147483648}
            _ = v118
        case "foo bar":
            v2[(fi116 >> ((782) & 63))] = fi116
        }
        _ = fi116
    }
    if true {
        delete(v2, 9223372036854775807)
        var v119 int64 = 42
        _ = v119
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v104:", v104)
    fmt.Println("v105:", v105)
    fmt.Println("v106:", v106)
    fmt.Println("slicelen v107:", len(v107), v107)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v104
    _ = v105
    _ = v106
    _ = v107
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v98
}
