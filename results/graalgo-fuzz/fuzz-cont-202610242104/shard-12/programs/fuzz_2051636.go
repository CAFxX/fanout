package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8; F1 int32; F2 uint }
type GenS4 struct { F0 int32; F1 uint }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 string; F1 bool }
type GenS8 struct { F0 bool; F1 bool; F2 int8 }
type GenS9 struct { F0 uint; F1 uint16; F2 uint64 }
type GenS10 struct { F0 string }
func genFn0(p0 uint32, p1 string) uint16 {
    p0 = 157
    v6 := GenS1{F0: 900}
    p0 = p0
    _ = p0
    _ = p1
    _ = v6
    return 32767
}
func genFn1() int32 {
    var v7 float32 = 0.1
    v8 := []int{32768, 32768, 4294967295}
    _ = v7
    _ = v8
    return 900
}

func main() {
    var v1 float64 = 1e10
    var v2 uint64 = 42
    var v3 uint32 = 717
    var v4 int32 = -2147483648
    var v5 uint8 = 128
    switch 39 {
    case -1:
        v9 := GenS1{F0: 540}
        _ = v9
    }
    var v10 string = "hello"
    if false {
        var v11 int16 = 32767
        _ = v11
    } else {
        if (true && (-128 <= 1)) {
            for fi12 := 0; fi12 < 3; fi12++ {
                var v13 uint16 = 2
                v14 := GenS2{F0: (!true)}
                _ = fi12
                _ = v13
                _ = v14
            }
            v15 := []int{-9223372036854775808, 32768}
            fmt.Println(v4)
            _ = v15
        }
        var v16 int32 = genFn1()
        var v17 string = ((v10 + v10) + "foo bar")
        _ = v16
        _ = v17
    }
    v1 = v1
    var v18 int16 = -32768
    v3 = v3
    if (v4 > (v4 << ((42) & 63))) {
        var v19 float64 = float64(0.1)
        v3 = (v3 * (v3 ^ (v3 - (v3 << ((30) & 63)))))
        v1 += v19
        _ = v19
    } else {
        fmt.Println(v10, v10, v3)
    }
    if false {
        v20 := []int{}
        _ = v20
    } else {
        v21 := GenS3{F0: v5, F1: 1, F2: 597}
        _ = v21
    }
    v5 = v5
    if (v4 <= genFn1()) {
        var v22 string = v10
        for fi23 := 0; fi23 < 15; fi23++ {
            v18 += -788
            v22 = v22
            _ = fi23
        }
        _ = v22
    } else {
        for fi24 := 0; fi24 < 15; fi24++ {
            for fi25 := 0; fi25 < 12; fi25++ {
                v2 = v2
                switch (fi25 & (fi24 / ((4294967295) | 1))) {
                case 32767:
                    switch fi24 {
                    case 2147483647:
                        v5 = (v5 << ((421) & 63))
                    case -2147483648:
                        if true {
                            if true {
                                v26 := []int{(fi25 ^ 327)}
                                var v27 int64 = -668
                                _ = v26
                                _ = v27
                            } else {
                                v28 := []int{(fi24 / ((fi24) | 1)), 9223372036854775807, fi25, (fi25 << ((397) & 63))}
                                v29 := GenS3{F0: (^(v5)), F1: genFn1(), F2: 421}
                                v10 = v10
                                var v30 uint64 = (v2 / ((v2) | 1))
                                _ = v28
                                _ = v29
                                _ = v30
                            }
                            var v31 float32 = (0.1 - (0.0 * float32(float64(((((3.14159 * 1e10) + 1e10) - 1.5) + float32(float64(0.1)))))))
                            var v32 uint64 = (v2 - (v2 - 18446744073709551615))
                            _ = v31
                            _ = v32
                        }
                    case 127:
                        var v33 int8 = -128
                        _ = v33
                    default:
                        switch fi24 {
                        case -195:
                            v34 := []int{}
                            _ = v34
                        case -2147483648:
                            var v35 int64 = -1
                            _ = v35
                        case 73:
                            var v36 int64 = 9223372036854775807
                            _ = v36
                        default:
                            fmt.Println(v10, v3, fi24)
                        }
                    }
                }
                _ = fi25
            }
            v10 = (v10 + "foo bar")
            v37 := []int{fi24, (fi24 - (fi24 / (((fi24 & -9223372036854775808)) | 1))), 127}
            _ = fi24
            _ = v37
        }
        v5 = 0
        switch v10 {
        case "":
            v5 = 127
        case "hello":
            if true {
                switch v10 {
                case "foo bar":
                    v4 -= (v4 ^ v4)
                case "a":
                    if (!(-1 < -128)) {
                        for fi38 := 0; fi38 < 11; fi38++ {
                            v39 := GenS3{F0: v5, F1: (v4 * -2147483648), F2: 0}
                            v3 += 65535
                            var v40 int16 = v18
                            _ = fi38
                            _ = v39
                            _ = v40
                        }
                        var v41 string = v10
                        v42 := []int{9223372036854775807}
                        _ = v41
                        _ = v42
                    } else {
                        v43 := []int{127, 4294967295, 678, 31}
                        _ = v43
                    }
                default:
                    var v44 int32 = -888
                    _ = v44
                }
                v45 := []int{-2147483648, -2147483648, 32768, 0}
                _ = v45
            } else {
                if ((v1 <= -2.25) || (!true)) {
                    var v46 float32 = (1e10 * (0.0 / 1e10))
                    var v47 uint64 = (^(v2))
                    if (!false) {
                        v4 = (v4 % ((v4) | 1))
                    }
                    _ = v46
                    _ = v47
                } else {
                    v48 := GenS4{F0: (v4 | (-(v4))), F1: 42}
                    var v49 uint8 = (v5 + (v5 ^ v5))
                    _ = v48
                    _ = v49
                }
                v1 += (v1 / (v1 - 3.14159))
            }
        case "foo bar":
            for fi50 := 0; fi50 < 16; fi50++ {
                v51 := []int{fi50, fi50, fi50}
                _ = fi50
                _ = v51
            }
        }
    }
    var v52 uint64 = 0
    v10 = v10
    if false {
        v2 = v2
        for fi53 := 0; fi53 < 6; fi53++ {
            switch 0 {
            case -722:
                for fi54 := 0; fi54 < 16; fi54++ {
                    v55 := []int{9223372036854775807}
                    _ = fi54
                    _ = v55
                }
            }
            switch fi53 {
            case 2147483647:
                if (v3 != v3) {
                    if false {
                        v56 := []int{fi53, (-(fi53)), fi53, (fi53 - fi53)}
                        if false {
                            for fi57 := 0; fi57 < 1; fi57++ {
                                v10 = (v10 + v10)
                                var v58 int = (fi57 ^ (fi57 ^ (fi57 ^ 459)))
                                _ = fi57
                                _ = v58
                            }
                            var v59 int32 = genFn1()
                            if (!(true && ((v59 & (-(v59))) <= (v59 ^ (v59 >> ((321) & 63)))))) {
                                v18 -= v18
                                if (917 == genFn0(v3, "\n\t")) {
                                    for fi60 := 0; fi60 < 9; fi60++ {
                                        var v61 uint = 65535
                                        if (!false) {
                                            if true {
                                                v62 := GenS1{F0: 255}
                                                var v63 uint8 = 255
                                                var v64 float32 = 0.0
                                                v10 = "a"
                                                v65 := GenS2{F0: (false && true)}
                                                if false {
                                                    v63 += (v5 | v5)
                                                    fmt.Println(v63, fi60)
                                                    if ((fi60 & (fi60 * fi60)) <= -94) {
                                                        switch v10 {
                                                        case "a":
                                                            if ((v59 / ((-2147483648) | 1)) == (v4 + v59)) {
                                                                fmt.Println(v18)
                                                                v66 := GenS5{F0: uint16(v61)}
                                                                var v67 uint16 = 255
                                                                var v68 int = 127
                                                                _ = v66
                                                                _ = v67
                                                                _ = v68
                                                            } else {
                                                                if (!true) {
                                                                    if (v2 < (v2 - 4294967295)) {
                                                                        var v69 uint32 = 2
                                                                        _ = v69
                                                                    } else {
                                                                        var v70 float64 = 3.14159
                                                                        _ = v70
                                                                    }
                                                                    var v71 int16 = v18
                                                                    _ = v71
                                                                } else {
                                                                    if (0 <= (fi53 & fi60)) {
                                                                        var v72 uint32 = v3
                                                                        _ = v72
                                                                    }
                                                                    var v73 uint8 = (v63 >> (((v61 ^ (v61 & (v61 % ((200) | 1))))) & 63))
                                                                    fmt.Println(v4)
                                                                    v74 := GenS3{F0: 195, F1: 2147483647, F2: 983}
                                                                    v75 := GenS3{F0: (v63 * v73), F1: (v4 & (-(v59))), F2: (^(v61))}
                                                                    if (645 <= genFn0((v3 | (v3 / ((uint32(v2)) | 1))), "foo bar")) {
                                                                        var v76 int64 = 0
                                                                        v77 := GenS6{F0: v59}
                                                                        fmt.Println(v63, v61)
                                                                        _ = v76
                                                                        _ = v77
                                                                    } else {
                                                                        fmt.Println(v64, v5)
                                                                        var v78 float32 = 0.0
                                                                        _ = v78
                                                                    }
                                                                    var v79 uint32 = (v3 & (v3 % (((v3 & v3)) | 1)))
                                                                    fmt.Println(v64, fi53)
                                                                    var v80 string = v10
                                                                    _ = v73
                                                                    _ = v74
                                                                    _ = v75
                                                                    _ = v79
                                                                    _ = v80
                                                                }
                                                            }
                                                        case "\n\t":
                                                            var v81 int = fi53
                                                            _ = v81
                                                        case "":
                                                            var v82 string = v10
                                                            _ = v82
                                                        default:
                                                            fmt.Println(v64, v10)
                                                        }
                                                        fmt.Println(v59, fi60)
                                                    }
                                                }
                                                _ = v62
                                                _ = v63
                                                _ = v64
                                                _ = v65
                                            } else {
                                                if len(v56) > 0 {
                                                    v56[0] = fi53
                                                }
                                                _ = len(v56)
                                            }
                                        } else {
                                            v4 = (v59 >> (((v61 / (((^(v61))) | 1))) & 63))
                                            v83 := GenS7{F0: v10, F1: false}
                                            v84 := GenS2{F0: ((^(v61)) >= 27)}
                                            _ = v83
                                            _ = v84
                                        }
                                        v59 -= (v59 << (((^(v61))) & 63))
                                        _ = fi60
                                        _ = v61
                                    }
                                    v1 += v1
                                    for fi85 := 0; fi85 < 19; fi85++ {
                                        fmt.Println(fi53, v18)
                                        v59 += (v59 | v4)
                                        v3 += (^(v3))
                                        _ = fi85
                                    }
                                } else {
                                    v86 := GenS8{F0: (0.0 == ((-2.25 * -2.25) / ((-2.25 * 1.5) * 0.1))), F1: false, F2: 0}
                                    if true {
                                        var v87 float32 = (float32(1.5) - 1e10)
                                        var v88 int8 = 0
                                        _ = v87
                                        _ = v88
                                    }
                                    v89 := GenS7{F0: v10, F1: ((false || (0 < 0)) && true)}
                                    _ = v86
                                    _ = v89
                                }
                                v59 = -242
                            }
                            _ = v59
                        }
                        _ = v56
                    }
                    for fi90 := 0; fi90 < 10; fi90++ {
                        if true {
                            if (false && false) {
                                var v91 bool = false
                                var v92 uint32 = 2147483647
                                var v93 uint = 32
                                _ = v91
                                _ = v92
                                _ = v93
                            } else {
                                if ((((v10 + (v10 + v10)) + ("foo bar" + (v10 + ("foo bar" + (v10 + v10))))) + "a") == (v10 + ("\n\t" + v10))) {
                                    if true {
                                        v1 -= v1
                                    } else {
                                        var v94 uint = uint(v3)
                                        _ = v94
                                    }
                                    var v95 string = "foo bar"
                                    _ = v95
                                }
                                if false {
                                    var v96 int64 = 16
                                    _ = v96
                                }
                                v97 := []int{(^(fi90)), (fi90 / ((fi90) | 1)), 127}
                                if (false || false) {
                                    fmt.Println(v52, fi53, v4)
                                    v98 := GenS8{F0: (0 >= 1), F1: false, F2: 123}
                                    var v99 int16 = v18
                                    _ = v98
                                    _ = v99
                                } else {
                                    var v100 uint32 = v3
                                    _ = len(v97)
                                    if true {
                                        if true {
                                            if (true && false) {
                                                switch "a" {
                                                case "\n\t":
                                                    switch 2147483647 {
                                                    case 4294967295:
                                                        var v101 uint16 = 714
                                                        _ = v101
                                                    case -893:
                                                        v2 = (v52 << ((65535) & 63))
                                                    }
                                                case "hello":
                                                    var v102 uint8 = v5
                                                    _ = v102
                                                case "foo bar":
                                                    fmt.Println(v52)
                                                }
                                                v103 := GenS8{F0: true, F1: (false || false), F2: -128}
                                                _ = v103
                                            }
                                        } else {
                                            if (false && (false || false)) {
                                                fmt.Println(v18, fi90, v5)
                                                if len(v97) >= 2 {
                                                    v97 = v97[1:len(v97)]
                                                }
                                                v18 = (v18 % ((v18) | 1))
                                            }
                                            v2 += (^(v52))
                                        }
                                        v104 := GenS5{F0: 32767}
                                        v4 = v4
                                        v105 := GenS9{F0: 1, F1: 65535, F2: v2}
                                        _ = v104
                                        _ = v105
                                    }
                                    switch fi90 {
                                    case -2147483648:
                                        if true {
                                            var v106 int64 = 4
                                            v107 := GenS4{F0: (v4 + genFn1()), F1: 1}
                                            _ = v106
                                            _ = v107
                                        } else {
                                            if (!false) {
                                                v108 := GenS2{F0: (365 >= 823)}
                                                _ = v108
                                            }
                                        }
                                    case 42:
                                        var v109 int = 29
                                        _ = v109
                                    case 2147483647:
                                        switch v10 {
                                        case "hello":
                                            v97 = append(v97, (fi53 >> ((196) & 63)))
                                        default:
                                            fmt.Println(v18, v4)
                                        }
                                    default:
                                        v5 += v5
                                    }
                                    v2 -= (^(v52))
                                    _ = v100
                                }
                                _ = v97
                            }
                        } else {
                            v2 += 985
                            if false {
                                v110 := GenS10{F0: v10}
                                var v111 int = fi90
                                _ = v110
                                _ = v111
                            }
                            fmt.Println(fi53, v18)
                        }
                        _ = fi90
                    }
                } else {
                    var v112 float64 = (3.14159 + float64(((1e10 - -2.25) / float32((-2.25 / (v1 + (v1 / float64(((3.14159 - ((((-2.25 - 1e10) * -2.25) - (((3.14159 - -2.25) - (1e10 - 1.5)) - 1.5)) / (3.14159 + 0.1))) * 3.14159)))))))))
                    _ = v112
                }
            default:
                var v113 int16 = (^(v18))
                _ = v113
            }
            _ = fi53
        }
    } else {
        v114 := []int{}
        v2 = v52
        _ = v114
    }
    var v115 bool = (uint64(v3) >= (v2 | v2))
    switch int(v5) {
    case 42:
        var v116 string = v10
        _ = v116
    default:
        v117 := GenS4{F0: genFn1(), F1: 665}
        _ = v117
    }
    v115 = (!((v1 != (v1 * 3.14159)) && v115))
    v118 := []int{-929, int(v4), -895, 177}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v115:", v115)
    fmt.Println("slicelen v118:", len(v118), v118)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v10
    _ = v115
    _ = v118
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
}
