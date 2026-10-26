package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 bool; F1 int8; F2 uint64 }
type GenS4 struct { F0 bool; F1 int32; F2 int64 }
type GenS5 struct { F0 int; F1 uint; F2 uint16 }
type GenS6 struct { F0 int16; F1 uint }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 uint64; F1 int32; F2 string }
type GenS10 struct { F0 uint32; F1 uint }
type GenS11 struct { F0 uint32; F1 string; F2 uint16 }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 bool }
type GenS14 struct { F0 int32; F1 int8; F2 int32 }
type GenS15 struct { F0 string }
func genFn0() int16 {
    if true {
        if ((((0 <= 42) && false) || true) || false) {
            v5 := []int{}
            var v6 string = "\n\t"
            _ = v5
            _ = v6
        }
        for fi7 := 0; fi7 < 5; fi7++ {
            fmt.Println(fi7)
            for fi8 := 0; fi8 < 1; fi8++ {
                var v9 uint8 = 110
                v10 := GenS1{F0: 42, F1: v9}
                var v11 int32 = 2147483647
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
            v12 := GenS2{F0: 42, F1: 1}
            _ = fi7
            _ = v12
        }
        if (42 >= 32767) {
            if (9223372036854775807 >= 32768) {
                var v13 int64 = -1
                v13 = v13
                if false {
                    v13 = -9223372036854775808
                } else {
                    for fi14 := 0; fi14 < 5; fi14++ {
                        var v15 string = ""
                        _ = fi14
                        _ = v15
                    }
                    if (0 == -2147483648) {
                        v13 = 42
                        v16 := GenS3{F0: true, F1: 0, F2: 828}
                        _ = v16
                    }
                    switch -264 {
                    case 0:
                        if false {
                            for fi17 := 0; fi17 < 17; fi17++ {
                                if (!(2 == 255)) {
                                    if true {
                                        v13 = v13
                                        fmt.Println(v13, fi17)
                                        v18 := GenS4{F0: ((255 > 160) || false), F1: 852, F2: v13}
                                        _ = v18
                                    }
                                }
                                _ = fi17
                            }
                        } else {
                            v19 := GenS5{F0: 128, F1: 18, F2: 582}
                            for fi20 := 0; fi20 < 1; fi20++ {
                                var v21 string = (("\n\t" + "") + "a")
                                _ = fi20
                                _ = v21
                            }
                            for fi22 := 0; fi22 < 4; fi22++ {
                                v23 := []int{(-(fi22)), (fi22 % ((fi22) | 1)), (fi22 >> ((973) & 63)), fi22}
                                _ = fi22
                                _ = v23
                            }
                            _ = v19
                        }
                    case -9223372036854775808:
                        v13 = (-(v13))
                    }
                }
                _ = v13
            } else {
                if (1 > 9223372036854775807) {
                    var v24 uint64 = 531
                    fmt.Println(v24)
                    _ = v24
                } else {
                    for fi25 := 0; fi25 < 9; fi25++ {
                        v26 := GenS6{F0: 0, F1: 42}
                        for fi27 := 0; fi27 < 5; fi27++ {
                            if false {
                                v28 := GenS1{F0: 1, F1: 2}
                                var v29 int16 = 0
                                var v30 uint16 = 2
                                v31 := []int{fi25, fi27, fi25}
                                _ = v28
                                _ = v29
                                _ = v30
                                _ = v31
                            } else {
                                if true {
                                    if (!false) {
                                        v32 := []int{328}
                                        switch (fi27 / (((-(fi27))) | 1)) {
                                        case 32767:
                                            fmt.Println(fi25)
                                        default:
                                            if false {
                                                if false {
                                                    if (true && (833 == 2147483648)) {
                                                        var v33 string = "hello"
                                                        v33 = (((v33 + "\n\t") + ("a" + v33)) + v33)
                                                        _ = v33
                                                    } else {
                                                        if (!true) {
                                                            var v34 uint32 = 42
                                                            v34 -= v34
                                                            _ = v34
                                                        } else {
                                                            switch 0 {
                                                            case 1:
                                                                var v35 uint64 = 228
                                                                _ = v35
                                                            case 21:
                                                                if ((fi25 >> ((255) & 63)) < (fi27 | fi25)) {
                                                                    fmt.Println(fi25, fi27)
                                                                    if false {
                                                                        if (42 < (fi25 << ((18446744073709551615) & 63))) {
                                                                            var v36 float64 = 3.14159
                                                                            var v37 string = "foo bar"
                                                                            var v38 uint32 = 2
                                                                            if true {
                                                                                _ = len(v32)
                                                                            }
                                                                            _ = v36
                                                                            _ = v37
                                                                            _ = v38
                                                                        }
                                                                        v39 := GenS7{F0: 617}
                                                                        _ = v39
                                                                    } else {
                                                                        if true {
                                                                            v40 := GenS8{F0: -1}
                                                                            v41 := GenS2{F0: 255, F1: 2147483647}
                                                                            _ = v40
                                                                            _ = v41
                                                                        } else {
                                                                            if len(v32) >= 2 {
                                                                                v32 = v32[1:len(v32)]
                                                                            }
                                                                            fmt.Println(fi25, fi25)
                                                                            var v42 string = ((((((("\n\t" + "foo bar") + (("\n\t" + (((("" + "foo bar") + "a") + ("" + ("" + "foo bar"))) + "")) + (("\n\t" + "\n\t") + "hello"))) + (("hello" + ("a" + "")) + "\n\t")) + "\n\t") + "\n\t") + "\n\t") + "hello")
                                                                            _ = v42
                                                                        }
                                                                        switch "\n\t" {
                                                                        case "\n\t":
                                                                            if true {
                                                                                v43 := GenS9{F0: 2, F1: -1, F2: ("hello" + ("hello" + "foo bar"))}
                                                                                var v44 int = fi25
                                                                                _ = v43
                                                                                _ = v44
                                                                            } else {
                                                                                var v45 int32 = 1
                                                                                var v46 int32 = (v45 & (v45 - 888))
                                                                                v45 = (v45 << ((65535) & 63))
                                                                                _ = v45
                                                                                _ = v46
                                                                            }
                                                                        default:
                                                                            var v47 uint64 = 348
                                                                            _ = v47
                                                                        }
                                                                        var v48 float32 = (0.0 + (1e10 * 1e10))
                                                                        _ = len(v32)
                                                                        _ = v48
                                                                    }
                                                                    v32 = append(v32, (fi25 & fi25))
                                                                }
                                                            case 375:
                                                                var v49 int8 = 100
                                                                _ = v49
                                                            default:
                                                                if len(v32) > 0 {
                                                                    v32[0] = 42
                                                                }
                                                            }
                                                        }
                                                        switch fi25 {
                                                        case 32768:
                                                            switch "\n\t" {
                                                            case "hello":
                                                                v50 := GenS9{F0: 4294967295, F1: -1, F2: "foo bar"}
                                                                _ = v50
                                                            case "":
                                                                var v51 float32 = (-2.25 - 1e10)
                                                                _ = v51
                                                            case "foo bar":
                                                                var v52 int16 = 1
                                                                _ = v52
                                                            }
                                                        case -2147483648:
                                                            var v53 float64 = float64((0.0 - 0.1))
                                                            _ = v53
                                                        case -176:
                                                            if true {
                                                                fmt.Println(fi25, fi27)
                                                                var v54 int16 = 1000
                                                                v55 := GenS2{F0: 32767, F1: int32(fi25)}
                                                                _ = v54
                                                                _ = v55
                                                            }
                                                        }
                                                    }
                                                    v56 := GenS10{F0: 2147483647, F1: 255}
                                                    if (-1 != 2147483647) {
                                                        var v57 int = 0
                                                        switch 32767 {
                                                        case -2147483648:
                                                            var v58 int16 = 0
                                                            _ = v58
                                                        }
                                                        var v59 int8 = -128
                                                        _ = v57
                                                        _ = v59
                                                    }
                                                    v32 = append(v32, fi25)
                                                    if len(v32) > 0 {
                                                        v32[0] = fi27
                                                    }
                                                    _ = v56
                                                } else {
                                                    var v60 int16 = 1
                                                    v61 := GenS11{F0: 1, F1: ("hello" + "hello"), F2: 42}
                                                    _ = v60
                                                    _ = v61
                                                }
                                                var v62 int8 = 1
                                                _ = v62
                                            } else {
                                                switch "" {
                                                case "\n\t":
                                                    var v63 uint8 = 128
                                                    _ = v63
                                                case "foo bar":
                                                    var v64 int32 = -689
                                                    _ = v64
                                                }
                                            }
                                        }
                                        if (false || true) {
                                            var v65 bool = true
                                            var v66 int8 = -1
                                            _ = v65
                                            _ = v66
                                        } else {
                                            fmt.Println(fi27, fi25)
                                            var v67 int16 = 1000
                                            v67 = (v67 >> ((65535) & 63))
                                            _ = v67
                                        }
                                        _ = v32
                                    } else {
                                        v68 := GenS3{F0: (100 == 1), F1: 100, F2: 515}
                                        var v69 int64 = -9223372036854775808
                                        var v70 float64 = (0.1 / (-2.25 + 1.5))
                                        _ = v68
                                        _ = v69
                                        _ = v70
                                    }
                                }
                                var v71 float32 = float32(((-2.25 + 0.1) + ((1e10 * 0.0) + ((0.0 * (1e10 - (1.5 * ((3.14159 / float64(1.5)) / 0.1)))) / 3.14159))))
                                fmt.Println(v71, v71, fi25)
                                switch ("\n\t" + "\n\t") {
                                case "\n\t":
                                    v72 := GenS5{F0: (fi27 >> ((7) & 63)), F1: 235, F2: 556}
                                    _ = v72
                                case "foo bar":
                                    fmt.Println(v71, v71)
                                case "hello":
                                    switch (fi27 % (((fi25 ^ 42)) | 1)) {
                                    case 1:
                                        if true {
                                            var v73 float64 = 0.0
                                            _ = v73
                                        }
                                    case 128:
                                        if (!false) {
                                            v71 -= v71
                                            v74 := []int{32767}
                                            v75 := GenS9{F0: 189, F1: -1, F2: "\n\t"}
                                            _ = v74
                                            _ = v75
                                        } else {
                                            v76 := []int{4294967295, fi27, (fi25 | fi27)}
                                            _ = v76
                                        }
                                    case 506:
                                        v77 := GenS9{F0: 168, F1: -1, F2: ""}
                                        _ = v77
                                    default:
                                        v71 -= float32(((0.0 - -2.25) / float64((1e10 - v71))))
                                    }
                                default:
                                    v71 -= -2.25
                                }
                                _ = v71
                            }
                            var v78 uint32 = 1
                            var v79 bool = (false || true)
                            _ = fi27
                            _ = v78
                            _ = v79
                        }
                        v80 := GenS12{F0: 0}
                        _ = fi25
                        _ = v26
                        _ = v80
                    }
                }
            }
            var v81 uint = 490
            _ = v81
        }
    }
    var v82 float32 = 0.1
    if true {
        var v83 uint32 = 42
        _ = v83
    }
    if (!((42 <= 565) || true)) {
        v84 := []int{32767, 128, -792}
        v82 += float32(-2.25)
        v84 = append(v84, 9223372036854775807)
        _ = v84
    } else {
        switch (("" + "") + ("\n\t" + ("" + "foo bar"))) {
        case "foo bar":
            if (-249 <= 32768) {
                if false {
                    v85 := GenS13{F0: true}
                    _ = v85
                }
                v82 += (v82 - 0.1)
                if (float64((v82 * 1e10)) == 1.5) {
                    var v86 float32 = (float32(0.1) - (v82 / v82))
                    v87 := GenS14{F0: -118, F1: 127, F2: -2147483648}
                    _ = v86
                    _ = v87
                } else {
                    v88 := GenS5{F0: 32768, F1: 56, F2: 1}
                    v89 := GenS6{F0: 266, F1: 42}
                    _ = v88
                    _ = v89
                }
            }
        case "a":
            switch 127 {
            case -128:
                var v90 uint16 = 848
                _ = v90
            default:
                v91 := []int{}
                _ = v91
            }
        case "":
            switch "hello" {
            case "hello":
                if true {
                    v92 := []int{2147483647}
                    var v93 uint8 = 0
                    if len(v92) > 0 {
                        v92[0] = -727
                    }
                    _ = v92
                    _ = v93
                }
            case "a":
                for fi94 := 0; fi94 < 5; fi94++ {
                    fmt.Println(v82)
                    _ = fi94
                }
            case "":
                v82 -= v82
            default:
                v82 = float32(float64(3.14159))
            }
        default:
            switch ("hello" + "\n\t") {
            case "a":
                fmt.Println(v82)
            case "\n\t":
                if (548 >= 2) {
                    switch 2147483647 {
                    case 33:
                        v82 = v82
                    case 438:
                        v95 := GenS8{F0: 443}
                        _ = v95
                    }
                    v82 += v82
                }
            case "":
                if true {
                    for fi96 := 0; fi96 < 7; fi96++ {
                        var v97 int32 = -1
                        _ = fi96
                        _ = v97
                    }
                    v98 := []int{}
                    _ = v98
                }
            default:
                v82 += 0.0
            }
        }
        var v99 uint32 = 2147483648
        _ = v99
    }
    var v100 string = "foo bar"
    _ = v100
    _ = v82
    return 1000
}
func genFn1(p0 int32, p1 uint8, p2 bool) uint8 {
    p1 = 0
    var v101 uint8 = 127
    fmt.Println(p2, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v101
    return v101
}

func main() {
    var v1 uint16 = 255
    var v2 int8 = 90
    v3 := make(map[string]uint32)
    var v4 bool = (true && true)
    v102 := GenS6{F0: 411, F1: 940}
    var v103 uint32 = 1
    v4 = v4
    if (!((3.14159 * (float32((1e10 * 0.0)) - ((3.14159 - (3.14159 * (1.5 - float32(((0.1 - 0.1) * 1.5))))) + (0.1 + 3.14159)))) < 1.5)) {
        var v104 uint64 = 0
        v105 := GenS15{F0: "foo bar"}
        _ = v104
        _ = v105
    }
    switch (((("\n\t" + "foo bar") + "a") + "foo bar") + "a") {
    case "a":
        v4 = (-1 >= 1000)
    default:
        if v4 {
            var v106 uint32 = v103
            _ = v106
        }
    }
    if false {
        var v107 bool = (470 < 4294967295)
        _ = v107
    }
    if v4 {
        var v108 uint8 = 241
        _ = v108
    } else {
        switch -2147483648 {
        case 2147483647:
            var v109 int64 = 9223372036854775807
            _ = v109
        case -655:
            v110 := GenS9{F0: 2, F1: -671, F2: "hello"}
            _ = v110
        case 584:
            for fi111 := 0; fi111 < 20; fi111++ {
                v112 := []int{}
                _ = fi111
                _ = v112
            }
        }
    }
    var v113 uint = 42
    fmt.Println("v1:", v1)
    fmt.Println("v103:", v103)
    fmt.Println("v113:", v113)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v102
    _ = v103
    _ = v113
    _ = v2
    _ = v3
    _ = v4
}
