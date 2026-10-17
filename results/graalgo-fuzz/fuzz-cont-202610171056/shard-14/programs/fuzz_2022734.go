package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int8; F1 string; F2 int16 }
type GenS7 struct { F0 int; F1 uint16 }
type GenS8 struct { F0 int; F1 uint16; F2 bool }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 uint32; F1 int }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 int8; F1 uint64 }
type GenS13 struct { F0 bool; F1 uint8; F2 int }
type GenS14 struct { F0 int8 }
func genFn0(p0 int16, p1 int32) int16 {
    v4 := GenS1{F0: 1, F1: 40}
    v5 := []int{-225, 351, 0}
    p1 = -1
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return p0
}
func genFn1(p0 int, p1 bool) uint64 {
    var v6 float32 = float32(1.5)
    if (!true) {
        var v7 float64 = (0.0 - (1e10 * float64(v6)))
        var v8 float32 = v6
        for fi9 := 0; fi9 < 14; fi9++ {
            v10 := []int{(p0 % ((fi9) | 1)), fi9, fi9}
            _ = fi9
            _ = v10
        }
        _ = v7
        _ = v8
    } else {
        for fi11 := 0; fi11 < 5; fi11++ {
            var v12 int8 = 42
            _ = fi11
            _ = v12
        }
        for fi13 := 0; fi13 < 16; fi13++ {
            if (1 < 12) {
                v6 = v6
                v14 := []int{(p0 * 128), 128, -9223372036854775808, (p0 << ((18446744073709551615) & 63))}
                _ = v14
            } else {
                var v15 uint8 = 128
                v16 := GenS1{F0: 1, F1: 1}
                for fi17 := 0; fi17 < 2; fi17++ {
                    var v18 uint64 = 844
                    v18 += 246
                    switch "a" {
                    case "a":
                        v19 := GenS1{F0: 42, F1: (v18 ^ v18)}
                        switch "\n\t" {
                        case "":
                            switch "foo bar" {
                            case "\n\t":
                                var v20 uint8 = 127
                                _ = v20
                            case "a":
                                v21 := GenS1{F0: 255, F1: (^(v18))}
                                _ = v21
                            case "hello":
                                switch fi13 {
                                case -954:
                                    v22 := []int{(fi17 * fi17), p0}
                                    _ = v22
                                default:
                                    if p1 {
                                        v23 := GenS1{F0: 65535, F1: 18446744073709551615}
                                        p0 += 32768
                                        if (!false) {
                                            v24 := []int{}
                                            switch ("foo bar" + "hello") {
                                            case "\n\t":
                                                switch ("a" + "a") {
                                                case "hello":
                                                    fmt.Println(fi17, v15, p0)
                                                case "foo bar":
                                                    v18 = 2
                                                default:
                                                    if len(v24) >= 2 {
                                                        v24 = v24[1:len(v24)]
                                                    }
                                                }
                                            default:
                                                var v25 bool = p1
                                                _ = v25
                                            }
                                            _ = v24
                                        } else {
                                            var v26 uint16 = 953
                                            switch "" {
                                            case "":
                                                v27 := []int{fi13, (fi17 >> ((uint(v18)) & 63)), -1}
                                                var v28 int32 = -747
                                                var v29 int8 = int8(v18)
                                                switch "a" {
                                                case "hello":
                                                    if len(v27) >= 2 {
                                                        v27 = v27[1:len(v27)]
                                                    }
                                                case "\n\t":
                                                    switch "a" {
                                                    case "foo bar":
                                                        v27 = append(v27, fi13)
                                                    case "\n\t":
                                                        v30 := GenS2{F0: -1}
                                                        _ = v30
                                                    }
                                                case "foo bar":
                                                    if len(v27) >= 2 {
                                                        v27 = v27[1:len(v27)]
                                                    }
                                                }
                                                _ = v27
                                                _ = v28
                                                _ = v29
                                            case "a":
                                                v31 := GenS3{F0: p1}
                                                _ = v31
                                            case "\n\t":
                                                p0 += -146
                                            }
                                            if (p1 || ((v26 > (v26 / ((65535) | 1))) && p1)) {
                                                p0 -= 624
                                            } else {
                                                v32 := []int{(p0 | p0), 763}
                                                if (false && true) {
                                                    fmt.Println(fi17)
                                                    v33 := GenS1{F0: (^(v26)), F1: v18}
                                                    fmt.Println(p1, v15)
                                                    _ = v33
                                                }
                                                v6 = -2.25
                                                _ = v32
                                            }
                                            _ = v26
                                        }
                                        _ = v23
                                    } else {
                                        if p1 {
                                            v15 = 250
                                            if (!(p1 && (2 < 42))) {
                                                switch "foo bar" {
                                                case "a":
                                                    if false {
                                                        var v34 string = "\n\t"
                                                        var v35 string = v34
                                                        if (1e10 < (float64(0.0) / 1.5)) {
                                                            var v36 string = (v34 + v35)
                                                            var v37 int16 = 1000
                                                            _ = v36
                                                            _ = v37
                                                        } else {
                                                            switch (v34 + v34) {
                                                            case "foo bar":
                                                                var v38 bool = (v34 <= "")
                                                                _ = v38
                                                            case "hello":
                                                                v34 = "hello"
                                                            case "\n\t":
                                                                var v39 float64 = float64((float32(0.1) + v6))
                                                                _ = v39
                                                            default:
                                                                p0 -= 32768
                                                            }
                                                            v34 = ("\n\t" + (v34 + v34))
                                                            switch fi17 {
                                                            case -9223372036854775808:
                                                                v40 := []int{}
                                                                _ = v40
                                                            default:
                                                                v41 := []int{fi13}
                                                                _ = v41
                                                            }
                                                        }
                                                        _ = v34
                                                        _ = v35
                                                    }
                                                    v18 = v18
                                                case "\n\t":
                                                    v18 += v18
                                                }
                                            }
                                        } else {
                                            var v42 uint = 558
                                            _ = v42
                                        }
                                    }
                                    var v43 float32 = 3.14159
                                    _ = v43
                                }
                            default:
                                v44 := GenS2{F0: -128}
                                _ = v44
                            }
                        case "foo bar":
                            v45 := GenS1{F0: 65535, F1: v18}
                            _ = v45
                        default:
                            p1 = p1
                        }
                        _ = v19
                    default:
                        v18 = 1
                    }
                    fmt.Println(fi13, p0)
                    var v46 uint64 = (v18 * (v18 >> ((1) & 63)))
                    _ = fi17
                    _ = v18
                    _ = v46
                }
                _ = v15
                _ = v16
            }
            if (p1 && (-1 < 1)) {
                for fi47 := 0; fi47 < 12; fi47++ {
                    p1 = p1
                    _ = fi47
                }
                for fi48 := 0; fi48 < 8; fi48++ {
                    var v49 int = (p0 * (fi13 << ((883) & 63)))
                    p1 = true
                    var v50 int8 = 100
                    fmt.Println(p0, p1)
                    _ = fi48
                    _ = v49
                    _ = v50
                }
                fmt.Println(v6, p0)
            }
            for fi51 := 0; fi51 < 4; fi51++ {
                if (!true) {
                    var v52 float32 = v6
                    v53 := GenS1{F0: 997, F1: 4294967295}
                    _ = v52
                    _ = v53
                }
                v54 := []int{(fi51 >> ((18446744073709551615) & 63))}
                p1 = ((!false) && (2 >= 0))
                if (float64(v6) != 0.1) {
                    var v55 int8 = 10
                    _ = v55
                } else {
                    if p1 {
                        var v56 uint = 869
                        _ = v56
                    } else {
                        fmt.Println(fi13)
                        fmt.Println(p0, fi51, fi13)
                    }
                    if len(v54) > 0 {
                        v54[0] = (p0 << ((1) & 63))
                    }
                }
                _ = fi51
                _ = v54
            }
            _ = fi13
        }
        v57 := GenS4{F0: 850}
        _ = v57
    }
    var v58 string = ""
    var v59 uint32 = 1
    if false {
        if (p1 || p1) {
            if p1 {
                v60 := []int{(p0 >> ((230) & 63)), (-(p0))}
                if ((true || (!false)) || (42 >= 1)) {
                    v61 := GenS5{F0: 885}
                    _ = v61
                }
                _ = v60
            }
            if false {
                var v62 string = v58
                v6 = (v6 + v6)
                for fi63 := 0; fi63 < 11; fi63++ {
                    if (!p1) {
                        for fi64 := 0; fi64 < 16; fi64++ {
                            v59 = 2147483647
                            var v65 int32 = 42
                            v66 := []int{p0, fi64}
                            _ = fi64
                            _ = v65
                            _ = v66
                        }
                        var v67 int64 = 9223372036854775807
                        if p1 {
                            v68 := GenS5{F0: 359}
                            _ = v68
                        }
                        _ = v67
                    }
                    p1 = true
                    _ = fi63
                }
                _ = v62
            }
            v69 := GenS6{F0: -128, F1: "", F2: 1}
            _ = v69
        }
        var v70 int32 = -2147483648
        var v71 bool = (65535 != 2147483647)
        _ = v70
        _ = v71
    }
    _ = p0
    _ = p1
    _ = v58
    _ = v59
    _ = v6
    return 42
}

func main() {
    v1 := []int{976, 21}
    var v2 bool = false
    var v3 int16 = 1
    var v72 uint8 = 255
    v3 = v3
    if (v2 || (genFn0((v3 / ((936) | 1)), -2147483648) != v3)) {
        v72 = v72
        if false {
            v73 := GenS7{F0: 127, F1: 1}
            _ = v73
        } else {
            var v74 int16 = v3
            if ((1e10 != -2.25) && v2) {
                v1 = append(v1, 848)
                var v75 string = "foo bar"
                _ = v75
            } else {
                var v76 uint16 = 1
                _ = v76
            }
            _ = v74
        }
        var v77 int64 = 42
        _ = v77
    }
    v72 = 42
    if len(v1) > 0 {
        v1[0] = 908
    }
    v3 = int16(v72)
    v78 := GenS2{F0: -66}
    for fi79 := 0; fi79 < 2; fi79++ {
        v80 := GenS8{F0: (fi79 % (((fi79 & (fi79 * 643))) | 1)), F1: 146, F2: v2}
        fmt.Println(fi79)
        _ = fi79
        _ = v80
    }
    fmt.Println(v3, v2, v3)
    switch "hello" {
    case "a":
        switch 0 {
        case 32767:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case 42:
            var v81 int64 = 452
            _ = v81
        case 941:
            for fi82 := 0; fi82 < 19; fi82++ {
                var v83 bool = v2
                if v2 {
                    switch fi82 {
                    case 127:
                        var v84 int8 = 100
                        _ = v84
                    case 128:
                        if v2 {
                            v85 := GenS9{F0: 1}
                            _ = v85
                        } else {
                            switch (fi82 % ((32767) | 1)) {
                            case 358:
                                if v2 {
                                    v1 = append(v1, fi82)
                                    v86 := GenS6{F0: 70, F1: "\n\t", F2: (v3 + -1)}
                                    switch "foo bar" {
                                    case "a":
                                        var v87 int32 = -2147483648
                                        _ = v87
                                    case "\n\t":
                                        if v2 {
                                            for fi88 := 0; fi88 < 7; fi88++ {
                                                var v89 uint16 = 255
                                                _ = fi88
                                                _ = v89
                                            }
                                            var v90 uint64 = 2
                                            if false {
                                                v1 = append(v1, fi82)
                                                v91 := GenS4{F0: genFn1((-(fi82)), v83)}
                                                var v92 float32 = 1e10
                                                _ = v91
                                                _ = v92
                                            }
                                            _ = v90
                                        }
                                    }
                                    _ = v86
                                }
                            case 42:
                                if false {
                                    var v93 uint8 = 42
                                    _ = len(v1)
                                    _ = v93
                                } else {
                                    if v2 {
                                        var v94 uint8 = (v72 * 42)
                                        _ = v94
                                    } else {
                                        var v95 uint32 = 94
                                        _ = v95
                                    }
                                }
                            case -768:
                                fmt.Println(v2, v3)
                            default:
                                if v2 {
                                    if (!v83) {
                                        if false {
                                            switch "" {
                                            case "foo bar":
                                                var v96 int32 = 2147483647
                                                _ = v96
                                            case "a":
                                                v3 = 1
                                            }
                                        }
                                        var v97 bool = v2
                                        _ = v97
                                    }
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    var v98 int = 920
                                    _ = v98
                                } else {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                            }
                        }
                    case -271:
                        var v99 uint = 356
                        _ = v99
                    default:
                        if v2 {
                            var v100 uint16 = 32767
                            v101 := GenS10{F0: 1, F1: (fi82 / ((fi82) | 1))}
                            if ((v100 / ((v100) | 1)) != 32767) {
                                v3 = -32768
                                var v102 uint32 = 65535
                                _ = v102
                            }
                            _ = v100
                            _ = v101
                        } else {
                            if v83 {
                                var v103 float32 = -2.25
                                _ = v103
                            }
                            v3 += 32767
                        }
                    }
                    v83 = true
                } else {
                    _ = len(v1)
                    if (42 < 42) {
                        fmt.Println(v2, v83, v3)
                    } else {
                        var v104 float64 = ((0.1 * 3.14159) + 3.14159)
                        _ = v104
                    }
                }
                _ = fi82
                _ = v83
            }
        }
    default:
        v2 = v2
    }
    v3 = genFn0((v3 & genFn0(genFn0((v3 * (v3 - v3)), 2147483647), 13)), 42)
    switch 133 {
    case -9223372036854775808:
        for fi105 := 0; fi105 < 10; fi105++ {
            v106 := GenS11{F0: -872}
            if len(v1) > 0 {
                v1[0] = (^(fi105))
            }
            _ = fi105
            _ = v106
        }
    case -503:
        var v107 int = -1
        _ = v107
    case -2147483648:
        if (v2 || v2) {
            v72 += 42
            for fi108 := 0; fi108 < 11; fi108++ {
                fmt.Println(v72, v3, v2)
                _ = fi108
            }
            var v109 uint8 = (v72 >> ((1) & 63))
            _ = v109
        }
    }
    v2 = (v2 && (3.14159 <= (0.0 + -2.25)))
    fmt.Println(v72, v72)
    if v2 {
        v110 := GenS12{F0: -128, F1: 2}
        fmt.Println(v72, v72, v72)
        _ = v110
    } else {
        v111 := GenS13{F0: v2, F1: v72, F2: 12}
        v112 := GenS14{F0: 42}
        switch -102 {
        case 60:
            var v113 bool = (!v2)
            _ = v113
        case 9223372036854775807:
            v3 += 32767
        default:
            var v114 uint32 = 0
            _ = v114
        }
        _ = v111
        _ = v112
    }
    v115 := GenS4{F0: genFn1(127, v2)}
    switch "hello" {
    case "\n\t":
        v2 = false
    case "a":
        v1 = append(v1, 42)
    case "foo bar":
        var v116 string = "\n\t"
        _ = v116
    default:
        v2 = v2
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v115
    _ = v2
    _ = v3
    _ = v72
    _ = v78
}
