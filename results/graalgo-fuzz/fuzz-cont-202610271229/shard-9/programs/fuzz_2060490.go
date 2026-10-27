package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint; F1 uint }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 string; F1 uint64; F2 uint64 }
type GenS8 struct { F0 int16 }
func genFn0() int64 {
    var v4 uint8 = 134
    v5 := []int{-1}
    switch "" {
    case "a":
        v6 := GenS1{F0: 2147483647}
        _ = v6
    case "":
        v7 := GenS2{F0: 287}
        _ = v7
    }
    _ = v4
    _ = v5
    return -1
}
func genFn1(p0 int8) uint {
    if true {
        p0 = p0
        switch 42 {
        case 42:
            v8 := GenS2{F0: 65535}
            _ = v8
        default:
            fmt.Println(p0)
        }
        fmt.Println(p0)
    }
    v9 := GenS2{F0: 850}
    switch "" {
    case "a":
        v10 := GenS2{F0: 4294967295}
        _ = v10
    case "\n\t":
        p0 += (-(p0))
    case "foo bar":
        var v11 int32 = -1
        _ = v11
    default:
        var v12 uint16 = 1
        _ = v12
    }
    var v13 float32 = ((1.5 - 1e10) - 1e10)
    switch -952 {
    case 742:
        p0 = (p0 % ((p0) | 1))
    default:
        if true {
            switch (("a" + ("" + "foo bar")) + "") {
            case "a":
                var v14 float32 = ((1e10 * (v13 / 1e10)) * (0.0 / v13))
                _ = v14
            case "foo bar":
                var v15 int16 = 42
                _ = v15
            case "hello":
                if (127 <= p0) {
                    var v16 uint64 = 526
                    _ = v16
                } else {
                    if false {
                        if false {
                            v17 := []int{127, 32767, 32768}
                            var v18 string = ("" + ("a" + "\n\t"))
                            _ = v17
                            _ = v18
                        }
                    } else {
                        if false {
                            p0 -= 0
                            if (!true) {
                                v19 := []int{32767, 128, 32767, 4294967295}
                                for fi20 := 0; fi20 < 13; fi20++ {
                                    p0 += p0
                                    _ = fi20
                                }
                                var v21 int = -764
                                _ = v19
                                _ = v21
                            }
                        }
                        for fi22 := 0; fi22 < 13; fi22++ {
                            for fi23 := 0; fi23 < 6; fi23++ {
                                if (true && false) {
                                    v24 := GenS2{F0: 42}
                                    _ = v24
                                } else {
                                    v25 := []int{}
                                    v13 -= v13
                                    v26 := GenS1{F0: 1}
                                    _ = v25
                                    _ = v26
                                }
                                if (!((((3.14159 != float32(0.0)) || (true || (v13 < 1e10))) || (65535 >= 1)) || (!((3.14159 + 1e10) < 0.0)))) {
                                    v27 := []int{int(p0), fi22, fi23}
                                    _ = v27
                                } else {
                                    fmt.Println(fi22, fi22)
                                }
                                v28 := GenS2{F0: 60}
                                fmt.Println(fi22)
                                v13 -= (1e10 / v13)
                                _ = fi23
                                _ = v28
                            }
                            for fi29 := 0; fi29 < 16; fi29++ {
                                var v30 uint32 = 2
                                var v31 string = ("a" + "hello")
                                v32 := GenS2{F0: 413}
                                fmt.Println(p0)
                                _ = fi29
                                _ = v30
                                _ = v31
                                _ = v32
                            }
                            var v33 float32 = v13
                            _ = fi22
                            _ = v33
                        }
                    }
                    v34 := []int{4294967295, 127}
                    _ = v34
                }
            }
            v13 += v13
        }
    }
    _ = p0
    _ = v13
    _ = v9
    return 0
}

func main() {
    var v1 uint = 111
    var v2 bool = (194 < 1)
    v3 := make(map[int]uint32)
    delete(v3, 339)
    if v2 {
        _ = len(v3)
    } else {
        if v2 {
            _ = len(v3)
            v35 := []int{127, 127}
            _ = v35
        } else {
            switch ("\n\t" + ("hello" + ("hello" + ("\n\t" + "")))) {
            case "":
                v2 = (-705 <= 32767)
            default:
                if (-2147483648 < 2147483647) {
                    switch -1 {
                    case 4294967295:
                        v36 := []int{1, 42, 9223372036854775807}
                        _ = v36
                    case 127:
                        v37 := GenS3{F0: v1, F1: v1}
                        _ = v37
                    }
                } else {
                    v38 := []int{32767, 607}
                    if v2 {
                        var v39 bool = (-9223372036854775808 == 127)
                        _ = v39
                    }
                    _ = v38
                }
            }
        }
        var v40 uint8 = 128
        _ = v40
    }
    switch "\n\t" {
    case "\n\t":
        v2 = v2
    case "hello":
        v41 := GenS3{F0: (v1 * (v1 * (v1 & v1))), F1: (v1 << ((genFn1(100)) & 63))}
        _ = v41
    default:
        fmt.Println(v2)
    }
    v2 = (true && (genFn0() == 365))
    v42 := GenS4{F0: 100}
    if true {
        for fi43 := 0; fi43 < 3; fi43++ {
            v2 = true
            switch "a" {
            case "a":
                v2 = v2
            case "hello":
                if (v2 && (!false)) {
                    for fi44 := 0; fi44 < 16; fi44++ {
                        var v45 uint = (v1 & 838)
                        if false {
                            v2 = v2
                            fmt.Println(v45)
                            fmt.Println(fi43)
                        } else {
                            v46 := []int{-535, (fi44 | -709), (fi43 | (fi44 % (((fi43 >> ((genFn1(42)) & 63))) | 1)))}
                            if (0 > 1) {
                                if false {
                                    switch "hello" {
                                    case "":
                                        v2 = true
                                    case "foo bar":
                                        var v47 int16 = 936
                                        _ = v47
                                    }
                                } else {
                                    var v48 int32 = -422
                                    v45 += 0
                                    v49 := GenS4{F0: 42}
                                    v50 := v3[fi43]
                                    v45 += v1
                                    _ = v48
                                    _ = v49
                                    _ = v50
                                }
                                var v51 float32 = ((1e10 * 0.0) * 3.14159)
                                _ = v51
                            } else {
                                var v52 int8 = 0
                                v45 += genFn1(-128)
                                var v53 int32 = -1
                                _ = v52
                                _ = v53
                            }
                            _ = v46
                        }
                        _ = fi44
                        _ = v45
                    }
                    var v54 float32 = -2.25
                    if (1 < 2) {
                        v54 = float32((0.0 / 1e10))
                        for fi55 := 0; fi55 < 3; fi55++ {
                            v56 := GenS1{F0: 1}
                            var v57 int32 = 262
                            v3[fi55] = 65535
                            if (!(!true)) {
                                v1 -= v1
                                if (449 == uint64(v1)) {
                                    v57 -= (v57 | (v57 ^ v57))
                                    v58 := GenS5{F0: 607}
                                    if (!(v2 || v2)) {
                                        var v59 uint16 = 32767
                                        _ = v59
                                    } else {
                                        v2 = (1 > 74)
                                        if v2 {
                                            v54 = 1.5
                                            var v60 int32 = v57
                                            var v61 uint32 = 133
                                            _ = v60
                                            _ = v61
                                        }
                                        v62 := GenS2{F0: 4294967295}
                                        _ = v62
                                    }
                                    if v2 {
                                        v1 += 0
                                        if v2 {
                                            fmt.Println(v57, v54)
                                        }
                                        fmt.Println(v54, v57, v2)
                                    } else {
                                        v54 = 1.5
                                        fmt.Println(fi55)
                                    }
                                    _ = v58
                                }
                                var v63 uint64 = 2
                                _ = v63
                            } else {
                                v64 := GenS5{F0: -32768}
                                var v65 int32 = 649
                                v65 = v57
                                _ = v64
                                _ = v65
                            }
                            _ = fi55
                            _ = v56
                            _ = v57
                        }
                    } else {
                        if (((0.0 - (1e10 - 1e10)) >= 1e10) || v2) {
                            if false {
                                for fi66 := 0; fi66 < 16; fi66++ {
                                    var v67 float32 = (((float32(float64(1e10)) - float32(((-2.25 - (1e10 * 3.14159)) - 1.5))) * 0.1) * float32(1.5))
                                    v2 = v2
                                    v68 := GenS3{F0: genFn1(-102), F1: (v1 + genFn1(32))}
                                    v54 = 1e10
                                    _ = fi66
                                    _ = v67
                                    _ = v68
                                }
                                switch fi43 {
                                case -2147483648:
                                    v2 = v2
                                case -185:
                                    var v69 uint32 = 1
                                    _ = v69
                                case 2147483647:
                                    for fi70 := 0; fi70 < 15; fi70++ {
                                        v1 += v1
                                        if false {
                                            v71 := GenS5{F0: 42}
                                            switch fi70 {
                                            case 1:
                                                v2 = true
                                            case 2147483647:
                                                v1 = (v1 & genFn1(-64))
                                            case 127:
                                                switch fi43 {
                                                case 62:
                                                    if (!(!v2)) {
                                                        if (32767 <= 945) {
                                                            if (42 < 0) {
                                                                var v72 int = (^(fi70))
                                                                _ = v72
                                                            } else {
                                                                var v73 int16 = 1
                                                                var v74 uint16 = 813
                                                                v2 = (!v2)
                                                                _ = v73
                                                                _ = v74
                                                            }
                                                            v2 = (!(false || v2))
                                                            v75 := v3[-106]
                                                            _ = v75
                                                        } else {
                                                            v1 = 0
                                                            if false {
                                                                var v76 uint8 = 42
                                                                _ = v76
                                                            } else {
                                                                if true {
                                                                    var v77 int = int(v1)
                                                                    switch fi70 {
                                                                    case 42:
                                                                        v78 := v3[(^(fi70))]
                                                                        _ = v78
                                                                    case 646:
                                                                        _ = len(v3)
                                                                    }
                                                                    if true {
                                                                        v1 = (v1 + v1)
                                                                        var v79 uint8 = 208
                                                                        if v2 {
                                                                            if (true || v2) {
                                                                                fmt.Println(v1, fi70)
                                                                                var v80 uint8 = v79
                                                                                v80 = (v80 & v79)
                                                                                _ = v80
                                                                            }
                                                                            v77 += fi70
                                                                            v81 := []int{(fi43 % (((^(fi43))) | 1)), fi70}
                                                                            _ = v81
                                                                        } else {
                                                                            v1 = v1
                                                                            var v82 int = 933
                                                                            v83 := []int{fi43, v77, 32768, 42}
                                                                            v54 += (v54 / (-2.25 / v54))
                                                                            _ = v82
                                                                            _ = v83
                                                                        }
                                                                        _ = v79
                                                                    } else {
                                                                        if v2 {
                                                                            v2 = (1 <= -780)
                                                                            v84 := GenS6{F0: 42}
                                                                            _ = v84
                                                                        }
                                                                        v1 += v1
                                                                        fmt.Println(v77)
                                                                    }
                                                                    var v85 int16 = -874
                                                                    _ = v77
                                                                    _ = v85
                                                                }
                                                            }
                                                        }
                                                        var v86 float32 = v54
                                                        if (!(v54 <= ((v54 + v86) * (v54 + -2.25)))) {
                                                            var v87 uint8 = 42
                                                            _ = v87
                                                        }
                                                        _ = v86
                                                    } else {
                                                        v88 := GenS3{F0: (^(v1)), F1: v1}
                                                        v89 := []int{fi70, -1, (fi70 << ((255) & 63))}
                                                        var v90 string = "hello"
                                                        _ = v88
                                                        _ = v89
                                                        _ = v90
                                                    }
                                                    fmt.Println(v54, v2)
                                                    _ = len(v3)
                                                case 32768:
                                                    var v91 float64 = float64(v54)
                                                    _ = v91
                                                }
                                                if (-1 >= 2147483647) {
                                                    var v92 string = ""
                                                    v1 += (v1 | (v1 | v1))
                                                    v54 = (v54 - v54)
                                                    _ = v92
                                                } else {
                                                    v1 = (^(v1))
                                                }
                                            }
                                            v54 += (-2.25 - v54)
                                            _ = v71
                                        }
                                        v93 := GenS7{F0: ("" + "hello"), F1: 962, F2: 2}
                                        var v94 uint16 = 65535
                                        _ = fi70
                                        _ = v93
                                        _ = v94
                                    }
                                }
                            } else {
                                var v95 int8 = -1
                                v96 := GenS7{F0: "hello", F1: 836, F2: 42}
                                _ = v95
                                _ = v96
                            }
                            v97 := GenS7{F0: "", F1: 18446744073709551615, F2: 0}
                            switch "a" {
                            case "\n\t":
                                for fi98 := 0; fi98 < 14; fi98++ {
                                    if (977 >= 32767) {
                                        var v99 uint64 = 183
                                        v100 := GenS6{F0: (v99 % ((uint64(v1)) | 1))}
                                        var v101 int8 = -128
                                        _ = v100
                                        _ = v101
                                        _ = v99
                                    } else {
                                        var v102 string = (("\n\t" + ("" + "foo bar")) + "\n\t")
                                        switch fi43 {
                                        case 42:
                                            delete(v3, fi98)
                                        case 594:
                                            fmt.Println(v2, fi43, v102)
                                        default:
                                            v1 = (v1 & 18446744073709551615)
                                        }
                                        v1 -= (^(v1))
                                        _ = v102
                                    }
                                    v54 = (float32((1e10 * (1e10 - (-2.25 - 3.14159)))) + v54)
                                    _ = fi98
                                }
                            case "foo bar":
                                var v103 int8 = 104
                                _ = v103
                            case "":
                                var v104 int = fi43
                                _ = v104
                            }
                            _ = v97
                        }
                        if true {
                            switch "\n\t" {
                            case "a":
                                fmt.Println(v1)
                            case "":
                                v3[fi43] = 0
                            }
                        }
                        switch (^(fi43)) {
                        case -9:
                            delete(v3, fi43)
                        case 128:
                            fmt.Println(v1, fi43)
                        case 42:
                            var v105 int8 = 41
                            _ = v105
                        }
                    }
                    _ = v54
                } else {
                    v106 := GenS4{F0: 42}
                    v2 = true
                    _ = v106
                }
            }
            _ = fi43
        }
        v2 = v2
        var v107 uint32 = 2
        _ = v107
    }
    var v108 uint = v1
    v2 = false
    fmt.Println(v108)
    v109 := []int{32767, -661, 896, 42}
    var v110 uint32 = 2147483647
    for fi111 := 0; fi111 < 3; fi111++ {
        if len(v109) >= 2 {
            v109 = v109[1:len(v109)]
        }
        v110 = (v110 * (v110 + v110))
        fmt.Println(v108, v108, v110)
        _ = fi111
    }
    if v2 {
        fmt.Println(v108, v1)
        var v112 int8 = 100
        if (0.1 != 3.14159) {
            var v113 int = 590
            _ = v113
        }
        _ = v112
    } else {
        fmt.Println(v2, v108, v1)
        if v2 {
            var v114 uint64 = 0
            v115 := GenS2{F0: v1}
            _ = v114
            _ = v115
        }
        v1 += v108
    }
    v116 := GenS7{F0: "", F1: uint64(v110), F2: 955}
    if ((v110 * (v110 - (v110 - v110))) != 2147483647) {
        if len(v109) > 0 {
            v109[0] = 2147483647
        }
    }
    switch "\n\t" {
    case "\n\t":
        v117 := GenS8{F0: -32768}
        _ = v117
    default:
        if len(v109) >= 2 {
            v109 = v109[1:len(v109)]
        }
    }
    if (127 < -119) {
        if true {
            v118 := GenS7{F0: "a", F1: 2, F2: 432}
            _ = v118
        }
    }
    switch 210 {
    case 2147483647:
        switch 42 {
        case 128:
            v119 := GenS1{F0: 0}
            _ = v119
        case -909:
            v109 = append(v109, 42)
        default:
            fmt.Println(v2, v1)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v108:", v108)
    fmt.Println("slicelen v109:", len(v109), v109)
    fmt.Println("v110:", v110)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v108
    _ = v109
    _ = v110
    _ = v116
    _ = v2
    _ = v3
    _ = v42
}
