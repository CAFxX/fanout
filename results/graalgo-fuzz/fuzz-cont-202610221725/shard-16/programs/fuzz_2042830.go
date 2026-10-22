package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 uint8 }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 string; F1 uint; F2 uint32 }
type GenS5 struct { F0 int16; F1 uint64 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 int; F1 uint16; F2 uint }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 int16 }
type GenS13 struct { F0 uint16; F1 uint16 }
func genFn0() uint {
    for fi7 := 0; fi7 < 13; fi7++ {
        fmt.Println(fi7)
        _ = fi7
    }
    if true {
        var v8 string = ("hello" + "foo bar")
        var v9 uint16 = 32767
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 65535, F1: 0, F2: 2}
    var v11 int16 = -835
    v12 := GenS1{F0: 42, F1: 0, F2: 58}
    _ = v10
    _ = v11
    _ = v12
    return 0
}

func main() {
    v1 := []uint32{}
    v2 := []string{"hello"}
    var v3 int16 = -32768
    var v4 int = 32767
    v5 := make(map[int]bool)
    v6 := []int{(v4 % (((v4 / ((9223372036854775807) | 1))) | 1)), 818, 1, (v4 % ((int(v3)) | 1))}
    fmt.Println(v3)
    if true {
        for fi13 := 0; fi13 < 5; fi13++ {
            if (-9223372036854775808 != -864) {
                for fi14 := 0; fi14 < 19; fi14++ {
                    v15 := GenS1{F0: 2147483648, F1: 23, F2: 199}
                    _ = fi14
                    _ = v15
                }
            } else {
                if (0 == 1) {
                    fmt.Println(fi13, v3, v3)
                } else {
                    var v16 int16 = 42
                    var v17 int8 = 95
                    v4 -= 32767
                    _ = v16
                    _ = v17
                }
            }
            v4 = (-(v4))
            _ = fi13
        }
        switch ("a" + ("" + "\n\t")) {
        case "":
            var v18 uint64 = 4294967295
            _ = v18
        default:
            v19 := GenS1{F0: 1, F1: 914, F2: 2}
            _ = v19
        }
        if false {
            if true {
                var v20 bool = (255 >= 2)
                v3 -= v3
                v21 := GenS1{F0: 4294967295, F1: 1, F2: 255}
                _ = v20
                _ = v21
            } else {
                if len(v2) > 0 {
                    v2[0] = ("\n\t" + "a")
                }
                if (1 <= -1) {
                    var v22 uint = 706
                    _ = v22
                }
            }
        }
    } else {
        fmt.Println(v3)
        var v23 bool = (90 < -9223372036854775808)
        for fi24 := 0; fi24 < 20; fi24++ {
            var v25 int16 = (v3 * 1)
            if (-2.25 > 1e10) {
                var v26 int8 = -128
                switch ((("a" + ("\n\t" + "a")) + "a") + "\n\t") {
                case "\n\t":
                    if false {
                        if (2147483647 >= -332) {
                            v25 = (-(v3))
                        } else {
                            fmt.Println(v26)
                            for fi27 := 0; fi27 < 8; fi27++ {
                                v26 = v26
                                _ = fi27
                            }
                            v28 := GenS2{F0: 2, F1: 0}
                            _ = v28
                        }
                        fmt.Println(v25, v26, v4)
                    }
                default:
                    var v29 int32 = -216
                    _ = v29
                }
                _ = v26
            } else {
                switch (v4 % (((v4 >> ((65535) & 63))) | 1)) {
                case 785:
                    for fi30 := 0; fi30 < 10; fi30++ {
                        var v31 float32 = 0.0
                        if (2147483647 == 42) {
                            v4 = 907
                            v3 = (v25 | v25)
                            v25 = v25
                        }
                        v31 += v31
                        var v32 uint16 = 1
                        var v33 float32 = v31
                        _ = fi30
                        _ = v31
                        _ = v32
                        _ = v33
                    }
                case -1:
                    v34 := GenS2{F0: 981, F1: 0}
                    _ = v34
                case 128:
                    v3 += (v25 % ((-1) | 1))
                }
                v35 := GenS3{F0: fi24}
                _ = v35
            }
            _ = fi24
            _ = v25
        }
        _ = v23
    }
    v36 := GenS3{F0: v4}
    v4 = (v4 >> ((2) & 63))
    switch "hello" {
    case "\n\t":
        v3 = (v3 - -932)
    case "":
        v37 := GenS2{F0: 32767, F1: 102}
        _ = v37
    case "foo bar":
        switch 1 {
        case -9223372036854775808:
            v38 := GenS2{F0: 1, F1: 128}
            _ = v38
        case 9223372036854775807:
            for fi39 := 0; fi39 < 9; fi39++ {
                var v40 int64 = 114
                _ = fi39
                _ = v40
            }
        case 477:
            v41 := GenS2{F0: 0, F1: 1}
            _ = v41
        default:
            v42 := GenS4{F0: "\n\t", F1: 42, F2: 177}
            _ = v42
        }
    default:
        v3 = (v3 + v3)
    }
    _ = len(v2)
    v43 := GenS1{F0: 1, F1: 2, F2: 247}
    v44 := GenS5{F0: v3, F1: 42}
    fmt.Println(v3)
    v45 := v5[(v4 ^ (v4 * 1))]
    switch ("foo bar" + ((("hello" + (("\n\t" + (("hello" + "hello") + ("hello" + ("" + "foo bar")))) + ("\n\t" + ("\n\t" + "a")))) + "\n\t") + "\n\t")) {
    case "":
        switch ("a" + "\n\t") {
        case "hello":
            v45 = ((-2147483648 == 2147483647) && v45)
        }
    default:
        var v46 float64 = 1.5
        _ = v46
    }
    var v47 float64 = 0.1
    v48 := GenS6{F0: (v4 ^ (v4 >> ((255) & 63)))}
    if (!v45) {
        switch (v4 >> ((925) & 63)) {
        case 9223372036854775807:
            var v49 int64 = -637
            _ = v49
        case 897:
            var v50 int32 = int32(v3)
            _ = v50
        case -153:
            if v45 {
                var v51 uint64 = 4294967295
                v45 = false
                _ = v51
            }
        }
        v47 = (v47 - 0.0)
    } else {
        for fi52 := 0; fi52 < 20; fi52++ {
            if v45 {
                for fi53 := 0; fi53 < 18; fi53++ {
                    delete(v5, (-(v4)))
                    switch 2147483647 {
                    case -9223372036854775808:
                        v4 = fi53
                    case -2147483648:
                        if (2147483648 == 2147483648) {
                            fmt.Println(v45, v45)
                            v54 := GenS5{F0: (v3 << ((0) & 63)), F1: 18446744073709551615}
                            _ = v54
                        }
                    case 764:
                        if true {
                            switch fi53 {
                            case 32767:
                                var v55 uint16 = 255
                                _ = v55
                            case 700:
                                if ("foo bar" <= "") {
                                    v56 := GenS7{F0: 2}
                                    _ = v56
                                }
                            case 32768:
                                var v57 int = 127
                                _ = v57
                            default:
                                if v45 {
                                    switch ("" + "\n\t") {
                                    case "":
                                        v3 = v3
                                        if v45 {
                                            var v58 int = (fi53 * -98)
                                            fmt.Println(v3, fi52, v4)
                                            var v59 uint = genFn0()
                                            _ = v58
                                            _ = v59
                                        }
                                        if v45 {
                                            v60 := GenS5{F0: (v3 | -1), F1: 4294967295}
                                            v1 = append(v1, 2)
                                            switch ("hello" + "hello") {
                                            case "a":
                                                switch fi52 {
                                                case 42:
                                                    if (!true) {
                                                        v45 = v45
                                                        if v45 {
                                                            v5[(fi53 >> ((0) & 63))] = v45
                                                            v6 = append(v6, 4294967295)
                                                            if len(v6) >= 2 {
                                                                v6 = v6[1:len(v6)]
                                                            }
                                                        } else {
                                                            v4 = fi52
                                                            v45 = (!false)
                                                            var v61 string = "a"
                                                            _ = v61
                                                        }
                                                    }
                                                case 128:
                                                    v3 += -280
                                                default:
                                                    if (0 == 2) {
                                                        v47 = (v47 / 1.5)
                                                        if v45 {
                                                            v62 := GenS6{F0: (v4 & v4)}
                                                            _ = v62
                                                        } else {
                                                            fmt.Println(v3, v3)
                                                            var v63 uint8 = 42
                                                            _ = v63
                                                        }
                                                        fmt.Println(v4)
                                                        v64 := GenS7{F0: 220}
                                                        var v65 uint32 = 42
                                                        _ = v64
                                                        _ = v65
                                                    } else {
                                                        delete(v5, (^(fi53)))
                                                        v2 = append(v2, ("hello" + "foo bar"))
                                                    }
                                                }
                                            case "\n\t":
                                                fmt.Println(v45, v3, fi52)
                                            }
                                            _ = v60
                                        }
                                    case "a":
                                        if true {
                                            var v66 float64 = 0.0
                                            _ = v66
                                        } else {
                                            v47 += float64(-2.25)
                                            v1 = append(v1, 503)
                                            var v67 uint64 = 314
                                            _ = v67
                                        }
                                    case "hello":
                                        switch (fi52 ^ fi52) {
                                        case 1:
                                            v68 := GenS8{F0: v4, F1: 55, F2: 42}
                                            _ = v68
                                        case -739:
                                            if true {
                                                _ = len(v5)
                                                v69 := GenS9{F0: 62}
                                                if (18446744073709551615 > 44) {
                                                    var v70 uint8 = 128
                                                    var v71 int32 = -724
                                                    var v72 int16 = -135
                                                    v47 = v47
                                                    v45 = false
                                                    _ = v70
                                                    _ = v71
                                                    _ = v72
                                                } else {
                                                    v47 += v47
                                                    _ = len(v1)
                                                    if v45 {
                                                        fmt.Println(v3)
                                                        if v45 {
                                                            if (!v45) {
                                                                v73 := GenS10{F0: -739}
                                                                v74 := GenS11{F0: v45}
                                                                _ = v73
                                                                _ = v74
                                                            } else {
                                                                if (true || (true && (v45 && v45))) {
                                                                    v75 := GenS1{F0: 2147483648, F1: 1, F2: 153}
                                                                    v45 = (-133 >= -1)
                                                                    if v45 {
                                                                        v47 = v47
                                                                        if v45 {
                                                                            var v76 bool = ((185 < 771) || false)
                                                                            var v77 uint8 = 2
                                                                            v77 += v77
                                                                            _ = v76
                                                                            _ = v77
                                                                        } else {
                                                                            v4 += (fi53 / ((fi52) | 1))
                                                                            switch 241 {
                                                                            case -9223372036854775808:
                                                                                v5[(v4 | v4)] = false
                                                                            case 0:
                                                                                v78 := GenS2{F0: 0, F1: 64}
                                                                                _ = v78
                                                                            }
                                                                        }
                                                                    }
                                                                    _ = v75
                                                                }
                                                            }
                                                            delete(v5, (fi53 - (^(v4))))
                                                        } else {
                                                            if (1 > 1) {
                                                                v45 = (!(2 > 559))
                                                                delete(v5, v4)
                                                                if true {
                                                                    fmt.Println(fi52, v4)
                                                                }
                                                            } else {
                                                                v79 := GenS8{F0: fi53, F1: 435, F2: 424}
                                                                var v80 uint8 = 255
                                                                var v81 int64 = 1
                                                                _ = v79
                                                                _ = v80
                                                                _ = v81
                                                            }
                                                            v3 = (v3 + v3)
                                                            v4 = (^(v4))
                                                            var v82 string = "\n\t"
                                                            var v83 uint = 65535
                                                            v83 -= genFn0()
                                                            _ = v82
                                                            _ = v83
                                                        }
                                                        var v84 int16 = (v3 << ((65535) & 63))
                                                        _ = v84
                                                    } else {
                                                        fmt.Println(fi53, v3)
                                                        v45 = (42 >= 0)
                                                        _ = len(v5)
                                                    }
                                                }
                                                if (!false) {
                                                    v47 += 0.0
                                                    v3 += (v3 ^ v3)
                                                    if len(v2) > 0 {
                                                        v2[0] = "hello"
                                                    }
                                                }
                                                switch 9223372036854775807 {
                                                case 127:
                                                    switch "hello" {
                                                    case "\n\t":
                                                        switch (fi52 - (fi52 - 0)) {
                                                        case 980:
                                                            fmt.Println(fi52)
                                                        case 918:
                                                            var v85 int32 = 738
                                                            _ = v85
                                                        case 4294967295:
                                                            fmt.Println(fi52, fi53, v4)
                                                        }
                                                    case "foo bar":
                                                        if v45 {
                                                            if ((((((3.14159 + ((0.0 - (-2.25 - float32(v47))) - float32(((v47 * float64(0.1)) / v47)))) * 1.5) > 0.0) && false) || ((fi53 << ((65535) & 63)) > -21)) && (false && (v47 > float64(1.5)))) {
                                                                v3 -= 908
                                                            } else {
                                                                fmt.Println(fi53, fi53, v45)
                                                                var v86 bool = (!v45)
                                                                _ = v86
                                                            }
                                                        }
                                                    case "hello":
                                                        fmt.Println(fi53, v3, v4)
                                                    default:
                                                        v45 = (((255 > 2) || v45) && false)
                                                    }
                                                default:
                                                    var v87 bool = true
                                                    _ = v87
                                                }
                                                v4 = v4
                                                _ = v69
                                            }
                                        }
                                    }
                                    var v88 float64 = 3.14159
                                    fmt.Println(v45, v47, fi53)
                                    _ = v88
                                } else {
                                    var v89 int32 = -2147483648
                                    if true {
                                        var v90 bool = (((v3 + (v3 ^ (v3 + -192))) != v3) && v45)
                                        _ = len(v5)
                                        _ = v90
                                    }
                                    fmt.Println(v47, v47, v45)
                                    _ = v89
                                }
                            }
                            v91 := GenS5{F0: -920, F1: 527}
                            _ = v91
                        } else {
                            v47 = v47
                            if v45 {
                                if true {
                                    var v92 uint = 18446744073709551615
                                    v3 = v3
                                    _ = v92
                                } else {
                                    var v93 float32 = (((float32((v47 + (v47 + v47))) - float32(v47)) - 0.0) / float32(1e10))
                                    _ = v93
                                }
                                var v94 uint16 = 32767
                                if true {
                                    v47 = v47
                                } else {
                                    v95 := GenS12{F0: (v3 & (v3 + int16(v4)))}
                                    v96 := GenS12{F0: v3}
                                    _ = len(v5)
                                    _ = v95
                                    _ = v96
                                }
                                _ = v94
                            } else {
                                _ = len(v6)
                                v97 := GenS4{F0: ("foo bar" + ""), F1: 0, F2: 911}
                                v47 += float64(0.0)
                                _ = v97
                            }
                        }
                    default:
                        var v98 uint8 = 128
                        _ = v98
                    }
                    v47 += (1e10 * v47)
                    v99 := v5[42]
                    var v100 uint16 = 0
                    fmt.Println(fi53, v99)
                    _ = fi53
                    _ = v100
                    _ = v99
                }
                v101 := GenS13{F0: 42, F1: 1}
                v102 := GenS7{F0: 0}
                _ = v101
                _ = v102
            }
            fmt.Println(v3, v47, v3)
            _ = fi52
        }
        var v103 int64 = -1
        var v104 uint8 = 255
        _ = v103
        _ = v104
    }
    var v105 float32 = (0.0 - 3.14159)
    var v106 int16 = 245
    for fi107 := 0; fi107 < 3; fi107++ {
        switch "foo bar" {
        case "hello":
            fmt.Println(v106)
        case "a":
            var v108 uint64 = 845
            _ = v108
        case "\n\t":
            if (!(255 > 1)) {
                var v109 uint = 794
                v109 = genFn0()
                _ = v109
            }
        default:
            v47 = 0.1
        }
        var v110 float32 = ((3.14159 - v105) / (1.5 - v105))
        _ = fi107
        _ = v110
    }
    for fi111 := 0; fi111 < 17; fi111++ {
        v2 = append(v2, ((("" + "hello") + "hello") + ("a" + (("a" + "") + ("" + (("\n\t" + "foo bar") + ((("hello" + "\n\t") + "foo bar") + "")))))))
        _ = fi111
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v105:", v105)
    fmt.Println("v106:", v106)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v105
    _ = v106
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v47
    _ = v48
    _ = v5
    _ = v6
}
