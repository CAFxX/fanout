package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 uint; F2 int8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint; F1 uint }
type GenS5 struct { F0 string; F1 bool; F2 uint32 }
type GenS6 struct { F0 int; F1 int32; F2 int64 }
type GenS7 struct { F0 uint8; F1 bool }
type GenS8 struct { F0 string; F1 uint8 }
type GenS9 struct { F0 uint64; F1 int16 }
type GenS10 struct { F0 uint; F1 int32 }
type GenS11 struct { F0 int32; F1 int16 }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 int }
type GenS14 struct { F0 int64 }
func genFn0() int64 {
    if (!(0 != 42)) {
        var v6 int8 = 127
        if (true && false) {
            switch 2147483647 {
            case -76:
                v7 := GenS1{F0: 1000}
                _ = v7
            case -9223372036854775808:
                v6 = (^(v6))
            default:
                var v8 int16 = 42
                _ = v8
            }
        }
        var v9 uint64 = 42
        _ = v6
        _ = v9
    }
    switch -1 {
    case 316:
        if (true && false) {
            v10 := []int{-2147483648, 0, -9223372036854775808}
            _ = v10
        }
    case 0:
        var v11 int = 9223372036854775807
        _ = v11
    case 127:
        for fi12 := 0; fi12 < 5; fi12++ {
            fmt.Println(fi12)
            _ = fi12
        }
    }
    var v13 string = ("" + "a")
    _ = v13
    return -60
}
func genFn1(p0 bool) int64 {
    if false {
        var v14 uint = 43
        _ = v14
    } else {
        v15 := []int{}
        if len(v15) >= 2 {
            v15 = v15[1:len(v15)]
        }
        _ = v15
    }
    p0 = p0
    _ = p0
    return -1
}

func main() {
    v1 := []uint32{2147483647, 2147483647, 65535}
    v2 := make(map[int]int)
    v3 := make(map[int]uint32)
    var v4 bool = true
    var v5 int16 = 0
    var v16 bool = true
    v16 = v16
    fmt.Println(v5)
    fmt.Println(v4)
    _ = len(v1)
    for fi17 := 0; fi17 < 11; fi17++ {
        var v18 uint64 = 558
        v4 = true
        _ = fi17
        _ = v18
    }
    fmt.Println(v16, v5)
    delete(v2, 32768)
    var v19 bool = false
    v16 = (1 == 42)
    for fi20 := 0; fi20 < 3; fi20++ {
        for fi21 := 0; fi21 < 15; fi21++ {
            v22 := GenS1{F0: 0}
            fmt.Println(fi21)
            _ = len(v2)
            _ = fi21
            _ = v22
        }
        if false {
            v16 = (((1e10 * 0.0) + 3.14159) < float32((3.14159 - float64((1e10 - 0.1)))))
        } else {
            var v23 int8 = 0
            v23 = (v23 ^ int8(fi20))
            if (false && v16) {
                var v24 uint = 290
                switch 127 {
                case -147:
                    if (2 < 42) {
                        switch (fi20 + 128) {
                        case 2147483647:
                            if (65535 >= 2) {
                                if false {
                                    if (v19 && (-2147483648 < -2147483648)) {
                                        switch (fi20 >> ((v24) & 63)) {
                                        case 9223372036854775807:
                                            _ = len(v2)
                                        }
                                        for fi25 := 0; fi25 < 6; fi25++ {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            _ = fi25
                                        }
                                    } else {
                                        for fi26 := 0; fi26 < 9; fi26++ {
                                            if (524 > 2) {
                                                delete(v2, fi26)
                                                v27 := GenS2{F0: 42, F1: 255, F2: 1}
                                                _ = v27
                                            } else {
                                                v5 = int16(v24)
                                                var v28 int8 = v23
                                                var v29 float32 = (((3.14159 / 1.5) - (1.5 * float32(0.1))) * 0.0)
                                                _ = v28
                                                _ = v29
                                            }
                                            v24 = (v24 | (v24 & v24))
                                            var v30 float32 = 1e10
                                            var v31 float32 = v30
                                            switch "" {
                                            case "hello":
                                                if v4 {
                                                    v2[-724] = -1
                                                    if false {
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        v30 += v31
                                                        if (32768 == 2147483647) {
                                                            fmt.Println(v16, fi20)
                                                            switch "\n\t" {
                                                            case "foo bar":
                                                                var v32 int8 = v23
                                                                _ = v32
                                                            case "a":
                                                                v33 := GenS1{F0: (v5 - int16(v23))}
                                                                _ = v33
                                                            case "hello":
                                                                v30 = v30
                                                            default:
                                                                if v19 {
                                                                    switch (((("a" + "foo bar") + "") + "") + "hello") {
                                                                    case "\n\t":
                                                                        v16 = true
                                                                    case "foo bar":
                                                                        v34 := GenS3{F0: 2147483647}
                                                                        _ = v34
                                                                    case "hello":
                                                                        v3[32768] = 532
                                                                    default:
                                                                        if v4 {
                                                                            var v35 uint32 = 558
                                                                            if len(v1) > 0 {
                                                                                v1[0] = (v35 * 42)
                                                                            }
                                                                            if v19 {
                                                                                v30 = v31
                                                                                var v36 uint64 = 188
                                                                                _ = v36
                                                                            }
                                                                            _ = v35
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                        }
                                                    } else {
                                                        _ = len(v1)
                                                        v5 -= (v5 - 1000)
                                                    }
                                                    if (v30 >= (((v31 / v31) / v30) + (v30 + 0.1))) {
                                                        v4 = true
                                                        v37 := GenS4{F0: (v24 / (((v24 / ((0) | 1))) | 1)), F1: v24}
                                                        _ = v37
                                                    }
                                                }
                                            }
                                            _ = fi26
                                            _ = v30
                                            _ = v31
                                        }
                                        switch fi20 {
                                        case 32768:
                                            if v19 {
                                                switch (^(fi20)) {
                                                case 0:
                                                    var v38 int32 = -1
                                                    _ = v38
                                                case 2147483647:
                                                    if false {
                                                        v39 := GenS1{F0: 0}
                                                        if (v19 && (0.0 > (1.5 - 1.5))) {
                                                            v40 := GenS2{F0: 0, F1: 42, F2: -9}
                                                            fmt.Println(v4)
                                                            _ = v40
                                                        } else {
                                                            v16 = true
                                                        }
                                                        _ = v39
                                                    } else {
                                                        if v4 {
                                                            if v4 {
                                                                var v41 int32 = 2147483647
                                                                if v16 {
                                                                    for fi42 := 0; fi42 < 20; fi42++ {
                                                                        v43 := v2[(fi42 << (((^(v24))) & 63))]
                                                                        _ = fi42
                                                                        _ = v43
                                                                    }
                                                                    for fi44 := 0; fi44 < 18; fi44++ {
                                                                        _ = len(v2)
                                                                        v4 = false
                                                                        _ = fi44
                                                                    }
                                                                    var v45 int = fi20
                                                                    _ = v45
                                                                } else {
                                                                    v46 := GenS1{F0: (v5 * -32768)}
                                                                    _ = v46
                                                                }
                                                                var v47 bool = v16
                                                                _ = v41
                                                                _ = v47
                                                            } else {
                                                                switch "" {
                                                                case "a":
                                                                    if v19 {
                                                                        v48 := GenS5{F0: "hello", F1: true, F2: 65535}
                                                                        v19 = v16
                                                                        _ = v48
                                                                    }
                                                                case "":
                                                                    fmt.Println(fi20, v4)
                                                                default:
                                                                    if true {
                                                                        v49 := GenS6{F0: (fi20 - fi20), F1: 2147483647, F2: -65}
                                                                        if (255 <= 0) {
                                                                            for fi50 := 0; fi50 < 15; fi50++ {
                                                                                if false {
                                                                                    v51 := GenS3{F0: 137}
                                                                                    var v52 int64 = -280
                                                                                    _ = v51
                                                                                    _ = v52
                                                                                } else {
                                                                                    if false {
                                                                                        v53 := GenS1{F0: 42}
                                                                                        v54 := GenS6{F0: fi50, F1: 0, F2: 9223372036854775807}
                                                                                        _ = v53
                                                                                        _ = v54
                                                                                    }
                                                                                    var v55 uint64 = 610
                                                                                    _ = v55
                                                                                }
                                                                                _ = fi50
                                                                            }
                                                                            v56 := GenS3{F0: 2147483647}
                                                                            _ = v56
                                                                        } else {
                                                                            for fi57 := 0; fi57 < 10; fi57++ {
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = 2
                                                                                }
                                                                                v4 = false
                                                                                _ = fi57
                                                                            }
                                                                        }
                                                                        v58 := GenS7{F0: 255, F1: v19}
                                                                        _ = v49
                                                                        _ = v58
                                                                    }
                                                                }
                                                                var v59 int16 = (v5 * v5)
                                                                if (!v16) {
                                                                    fmt.Println(fi20, v19, v19)
                                                                } else {
                                                                    v4 = false
                                                                    if (v19 || (255 > (v24 + (v24 & v24)))) {
                                                                        fmt.Println(v16)
                                                                    } else {
                                                                        var v60 uint8 = 1
                                                                        switch (("a" + "\n\t") + "") {
                                                                        case "hello":
                                                                            v61 := GenS8{F0: "", F1: (^(v60))}
                                                                            _ = v61
                                                                        default:
                                                                            v62 := GenS2{F0: v60, F1: (v24 / ((1) | 1)), F2: 0}
                                                                            _ = v62
                                                                        }
                                                                        _ = v60
                                                                    }
                                                                }
                                                                _ = v59
                                                            }
                                                            for fi63 := 0; fi63 < 6; fi63++ {
                                                                switch "hello" {
                                                                case "foo bar":
                                                                    switch -110 {
                                                                    case 32767:
                                                                        fmt.Println(v24)
                                                                    default:
                                                                        if v19 {
                                                                            v23 -= (v23 - (v23 / ((1) | 1)))
                                                                            v64 := v2[fi63]
                                                                            v65 := GenS7{F0: 127, F1: ((1 != 42) && (!(901 != 516)))}
                                                                            _ = v64
                                                                            _ = v65
                                                                        } else {
                                                                            if (((1e10 < 0.0) || (901 == 0)) && (0 >= 1)) {
                                                                                if true {
                                                                                    var v66 int64 = 29
                                                                                    switch (("" + (("hello" + "a") + "foo bar")) + "foo bar") {
                                                                                    case "foo bar":
                                                                                        var v67 uint = (v24 >> ((761) & 63))
                                                                                        _ = v67
                                                                                    default:
                                                                                        var v68 uint8 = 127
                                                                                        _ = v68
                                                                                    }
                                                                                    var v69 int8 = v23
                                                                                    _ = v66
                                                                                    _ = v69
                                                                                } else {
                                                                                    if (genFn1((128 > fi20)) > 1) {
                                                                                        v24 = v24
                                                                                    }
                                                                                    switch (fi63 << ((v24) & 63)) {
                                                                                    case 4294967295:
                                                                                        if v16 {
                                                                                            var v70 int32 = 2147483647
                                                                                            v70 = 1
                                                                                            v71 := GenS9{F0: 0, F1: (^(v5))}
                                                                                            _ = v70
                                                                                            _ = v71
                                                                                        } else {
                                                                                            if len(v1) > 0 {
                                                                                                v1[0] = 991
                                                                                            }
                                                                                            v72 := v2[128]
                                                                                            v72 = 32768
                                                                                            switch "" {
                                                                                            case "foo bar":
                                                                                                v19 = v4
                                                                                            case "":
                                                                                                v73 := v2[(fi20 << ((65535) & 63))]
                                                                                                _ = v73
                                                                                            case "\n\t":
                                                                                                var v74 float32 = 1.5
                                                                                                _ = v74
                                                                                            default:
                                                                                                if false {
                                                                                                    v24 = 571
                                                                                                    v75 := GenS10{F0: v24, F1: -948}
                                                                                                    v76 := GenS3{F0: 42}
                                                                                                    _ = v75
                                                                                                    _ = v76
                                                                                                }
                                                                                            }
                                                                                            v1 = append(v1, 589)
                                                                                            if v19 {
                                                                                                v77 := v3[-1]
                                                                                                v19 = (v19 || false)
                                                                                                var v78 uint64 = 74
                                                                                                _ = v77
                                                                                                _ = v78
                                                                                            } else {
                                                                                                if len(v1) > 0 {
                                                                                                    v1[0] = 759
                                                                                                }
                                                                                                if (!v16) {
                                                                                                    v19 = v4
                                                                                                    if len(v1) > 0 {
                                                                                                        v1[0] = 42
                                                                                                    }
                                                                                                    fmt.Println(v72, v5, v24)
                                                                                                }
                                                                                            }
                                                                                            _ = v72
                                                                                        }
                                                                                    default:
                                                                                        v19 = v16
                                                                                    }
                                                                                    var v79 int64 = 293
                                                                                    _ = v79
                                                                                }
                                                                            } else {
                                                                                v23 = v23
                                                                                var v80 uint = 18446744073709551615
                                                                                v80 -= v24
                                                                                _ = v80
                                                                            }
                                                                            switch (fi63 + -15) {
                                                                            case 0:
                                                                                v81 := GenS1{F0: (-(v5))}
                                                                                _ = v81
                                                                            case -9223372036854775808:
                                                                                v16 = (42 > -1)
                                                                                if false {
                                                                                    var v82 string = ("hello" + "hello")
                                                                                    _ = v82
                                                                                }
                                                                            case 709:
                                                                                v5 += v5
                                                                            }
                                                                        }
                                                                    }
                                                                case "\n\t":
                                                                    if v16 {
                                                                        v23 += (v23 ^ v23)
                                                                        var v83 uint16 = 42
                                                                        var v84 int16 = (^(v5))
                                                                        _ = len(v1)
                                                                        _ = v83
                                                                        _ = v84
                                                                    } else {
                                                                        fmt.Println(v24)
                                                                    }
                                                                case "hello":
                                                                    switch "foo bar" {
                                                                    case "":
                                                                        var v85 int8 = (-(v23))
                                                                        _ = v85
                                                                    case "a":
                                                                        var v86 uint64 = 1
                                                                        _ = v86
                                                                    case "\n\t":
                                                                        fmt.Println(v24, v19)
                                                                    default:
                                                                        v24 = v24
                                                                    }
                                                                }
                                                                v3[fi20] = 283
                                                                fmt.Println(v4)
                                                                _ = fi63
                                                            }
                                                            for fi87 := 0; fi87 < 16; fi87++ {
                                                                v88 := GenS11{F0: 42, F1: int16(v23)}
                                                                var v89 uint64 = 805
                                                                if v19 {
                                                                    if false {
                                                                        if (42 >= 125) {
                                                                            if ((-2.25 + 0.0) > 0.0) {
                                                                                if v16 {
                                                                                    if (!(2147483647 >= 1)) {
                                                                                        if len(v1) > 0 {
                                                                                            v1[0] = 0
                                                                                        }
                                                                                        var v90 float64 = 1.5
                                                                                        _ = v90
                                                                                    } else {
                                                                                        var v91 int16 = (-(v5))
                                                                                        var v92 string = ""
                                                                                        v93 := GenS12{F0: 2147483648}
                                                                                        _ = v91
                                                                                        _ = v92
                                                                                        _ = v93
                                                                                    }
                                                                                } else {
                                                                                    var v94 uint = 272
                                                                                    if (!((v23 >= (v23 | v23)) && ((fi20 << ((v94) & 63)) != int(v94)))) {
                                                                                        if len(v1) >= 2 {
                                                                                            v1 = v1[1:len(v1)]
                                                                                        }
                                                                                        v95 := GenS1{F0: 1000}
                                                                                        v4 = true
                                                                                        _ = v95
                                                                                    }
                                                                                    _ = v94
                                                                                }
                                                                                v4 = (249 > 110)
                                                                                var v96 float64 = (-2.25 * 1e10)
                                                                                _ = v96
                                                                            }
                                                                            v97 := GenS1{F0: (v5 * (^(v5)))}
                                                                            _ = v97
                                                                        } else {
                                                                            fmt.Println(v24, v89, v23)
                                                                            fmt.Println(v89, v24)
                                                                            v23 += (v23 + v23)
                                                                        }
                                                                        _ = len(v2)
                                                                    } else {
                                                                        v24 = (v24 + 65535)
                                                                        v5 += (v5 ^ 600)
                                                                        var v98 uint8 = 66
                                                                        switch fi87 {
                                                                        case -253:
                                                                            if true {
                                                                                var v99 int8 = v23
                                                                                _ = v99
                                                                            }
                                                                        case -9223372036854775808:
                                                                            v5 = v5
                                                                        case 32768:
                                                                            if v19 {
                                                                                switch "hello" {
                                                                                case "":
                                                                                    var v100 int = (fi87 * 4294967295)
                                                                                    _ = v100
                                                                                case "hello":
                                                                                    v4 = true
                                                                                case "\n\t":
                                                                                    v101 := GenS3{F0: 42}
                                                                                    _ = v101
                                                                                default:
                                                                                    if ((!(int(v5) != fi87)) || (!(-32768 < v5))) {
                                                                                        var v102 uint64 = v89
                                                                                        v4 = v19
                                                                                        v103 := GenS6{F0: (fi20 >> (((v24 % ((v24) | 1))) & 63)), F1: 42, F2: -84}
                                                                                        _ = v102
                                                                                        _ = v103
                                                                                    } else {
                                                                                        if (!(0 < 2147483647)) {
                                                                                            v104 := GenS13{F0: (fi87 % ((-2147483648) | 1))}
                                                                                            _ = v104
                                                                                        } else {
                                                                                            v23 += (v23 * 100)
                                                                                        }
                                                                                        var v105 int8 = (^(v23))
                                                                                        _ = v105
                                                                                    }
                                                                                }
                                                                                var v106 float64 = 1e10
                                                                                if false {
                                                                                    switch fi87 {
                                                                                    case 42:
                                                                                        v16 = v4
                                                                                    }
                                                                                    fmt.Println(v24)
                                                                                } else {
                                                                                    var v107 uint8 = 128
                                                                                    var v108 string = "\n\t"
                                                                                    fmt.Println(v5)
                                                                                    _ = v107
                                                                                    _ = v108
                                                                                }
                                                                                _ = v106
                                                                            } else {
                                                                                var v109 int16 = -86
                                                                                _ = v109
                                                                            }
                                                                        }
                                                                        var v110 uint64 = (v89 + v89)
                                                                        _ = v110
                                                                        _ = v98
                                                                    }
                                                                    v111 := GenS2{F0: 2, F1: (v24 << ((65535) & 63)), F2: (v23 / ((v23) | 1))}
                                                                    _ = v111
                                                                }
                                                                _ = fi87
                                                                _ = v88
                                                                _ = v89
                                                            }
                                                        }
                                                    }
                                                case 32767:
                                                    fmt.Println(fi20, v23)
                                                }
                                                v1 = append(v1, 65535)
                                                for fi112 := 0; fi112 < 3; fi112++ {
                                                    v5 += v5
                                                    v113 := GenS13{F0: 0}
                                                    if len(v1) > 0 {
                                                        v1[0] = 0
                                                    }
                                                    _ = fi112
                                                    _ = v113
                                                }
                                            }
                                        }
                                    }
                                    var v114 int64 = genFn0()
                                    _ = len(v1)
                                    _ = v114
                                } else {
                                    var v115 int16 = v5
                                    _ = v115
                                }
                                v116 := GenS14{F0: 325}
                                _ = v116
                            }
                        case 631:
                            var v117 uint = (v24 / (((v24 << ((v24) & 63))) | 1))
                            _ = v117
                        default:
                            if (v19 && (!v16)) {
                                v23 -= v23
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                v118 := GenS8{F0: "", F1: 255}
                                _ = v118
                            } else {
                                var v119 int16 = 42
                                _ = v119
                            }
                        }
                        var v120 int16 = v5
                        _ = v120
                    } else {
                        v4 = false
                    }
                case 299:
                    v19 = ((1 <= 0) && (1 < -9223372036854775808))
                case 0:
                    v16 = (!v19)
                }
                _ = v24
            } else {
                var v121 uint8 = 0
                for fi122 := 0; fi122 < 4; fi122++ {
                    var v123 bool = ((v121 & v121) != v121)
                    v16 = v16
                    var v124 float64 = 1.5
                    _ = fi122
                    _ = v123
                    _ = v124
                }
                _ = v121
            }
            _ = v23
        }
        _ = fi20
    }
    if (!v19) {
        v125 := GenS3{F0: 0}
        _ = v125
    }
    _ = len(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
