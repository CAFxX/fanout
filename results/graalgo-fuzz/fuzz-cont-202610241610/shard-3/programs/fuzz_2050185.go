package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int; F1 int; F2 int16 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 bool; F1 int64 }
type GenS6 struct { F0 uint32; F1 uint16; F2 uint8 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int16; F1 uint8; F2 int8 }
type GenS9 struct { F0 string; F1 int; F2 int16 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 string; F1 uint32 }
type GenS12 struct { F0 uint16; F1 uint32; F2 string }

func main() {
    var v1 int = 724
    var v2 int32 = 1
    var v3 int64 = -332
    var v4 uint8 = 255
    for fi5 := 0; fi5 < 11; fi5++ {
        v6 := GenS1{F0: 559, F1: 2, F2: v2}
        var v7 uint8 = 0
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 uint32 = 708
    v9 := GenS2{F0: 2}
    if ((42 == 42) || (true || (!(("\n\t" != ("\n\t" + "\n\t")) || (!false))))) {
        v3 = (^(v3))
        for fi10 := 0; fi10 < 7; fi10++ {
            for fi11 := 0; fi11 < 10; fi11++ {
                if true {
                    v12 := []int{(-(fi11))}
                    _ = v12
                }
                fmt.Println(fi10, v4, v3)
                switch "foo bar" {
                case "\n\t":
                    switch fi11 {
                    case 1:
                        switch fi11 {
                        case 0:
                            v3 -= -9223372036854775808
                        }
                    case 2147483647:
                        v8 = v8
                    }
                case "":
                    v13 := GenS2{F0: 255}
                    _ = v13
                case "a":
                    v14 := GenS3{F0: fi10, F1: 32768, F2: -32768}
                    _ = v14
                default:
                    v1 = -9223372036854775808
                }
                if false {
                    if true {
                        if true {
                            if true {
                                if ((v3 | 1) == v3) {
                                    if (4294967295 > 42) {
                                        var v15 float32 = 1.5
                                        if (!false) {
                                            v3 += (v3 >> ((18446744073709551615) & 63))
                                            fmt.Println(v1, v15, v2)
                                            var v16 float32 = 0.1
                                            _ = v16
                                        } else {
                                            v17 := []int{}
                                            switch 32768 {
                                            case -670:
                                                v18 := GenS3{F0: v1, F1: (^(v1)), F2: 32767}
                                                _ = v18
                                            case 9223372036854775807:
                                                var v19 string = "\n\t"
                                                _ = v19
                                            case 51:
                                                switch 32767 {
                                                case 127:
                                                    switch fi11 {
                                                    case 2147483647:
                                                        v2 -= v2
                                                        var v20 int8 = 127
                                                        _ = v20
                                                    case -638:
                                                        var v21 string = "\n\t"
                                                        _ = v21
                                                    }
                                                case 684:
                                                    var v22 uint = uint(v4)
                                                    v22 = 65535
                                                    v2 = 2147483647
                                                    v2 -= 0
                                                    var v23 uint = (v22 >> ((v22) & 63))
                                                    _ = v22
                                                    _ = v23
                                                }
                                            default:
                                                if (false || false) {
                                                    v24 := GenS4{F0: false}
                                                    v25 := GenS1{F0: 42, F1: 18446744073709551615, F2: v2}
                                                    v2 = -86
                                                    _ = v24
                                                    _ = v25
                                                }
                                            }
                                            _ = v17
                                        }
                                        switch "foo bar" {
                                        case "\n\t":
                                            fmt.Println(v3, v15)
                                        case "hello":
                                            v26 := []int{0}
                                            _ = v26
                                        case "foo bar":
                                            v8 = (v8 >> ((2) & 63))
                                            v27 := GenS3{F0: (fi10 - (^(fi10))), F1: fi10, F2: 1000}
                                            var v28 uint64 = 481
                                            var v29 float32 = ((v15 * ((3.14159 / 1.0) * ((v15 - (v15 + (v15 / ((float32((-2.25 / -2.25)) - (v15 / float32(3.14159))) + v15)))) * 0.1))) / v15)
                                            _ = v27
                                            _ = v28
                                            _ = v29
                                        default:
                                            v30 := GenS5{F0: true, F1: v3}
                                            _ = v30
                                        }
                                        _ = v15
                                    } else {
                                        var v31 uint = 42
                                        var v32 uint = 255
                                        _ = v31
                                        _ = v32
                                    }
                                }
                                v33 := GenS6{F0: 390, F1: 66, F2: (v4 - v4)}
                                _ = v33
                            } else {
                                var v34 int64 = v3
                                var v35 int8 = 8
                                v36 := GenS5{F0: true, F1: int64(v35)}
                                _ = v34
                                _ = v35
                                _ = v36
                            }
                            v37 := []int{-904, v1, 0}
                            v4 -= (^(v4))
                            _ = v37
                        }
                        if (((0.1 > 1.5) || (v3 == v3)) && (((757 != 42) || false) || ("\n\t" < "hello"))) {
                            var v38 uint8 = (v4 * 255)
                            if false {
                                v39 := GenS7{F0: 42}
                                var v40 float32 = float32((((((((float64((-2.25 - float32((0.1 - 0.0)))) * 3.14159) / (0.1 + 3.14159)) * ((((((float64(3.14159) - 3.14159) + 0.1) - 1e10) - 1e10) / 1e10) * 0.0)) + 1.5) + 0.0) + ((3.14159 * (0.1 * (1e10 - 0.0))) - float64(-2.25))) + -2.25))
                                v4 = (v4 * v4)
                                _ = v39
                                _ = v40
                            }
                            _ = v38
                        }
                    } else {
                        v41 := []int{fi11, fi11}
                        var v42 int8 = 100
                        fmt.Println(fi10, fi11)
                        _ = v41
                        _ = v42
                    }
                    var v43 int8 = 42
                    _ = v43
                } else {
                    switch v1 {
                    case -1:
                        var v44 int8 = 42
                        _ = v44
                    case -364:
                        v45 := []int{}
                        _ = v45
                    }
                    v46 := GenS2{F0: 586}
                    v47 := []int{}
                    _ = v46
                    _ = v47
                }
                var v48 uint8 = v4
                var v49 uint8 = v48
                switch "" {
                case "a":
                    fmt.Println(v1, v1, v1)
                default:
                    if (!true) {
                        fmt.Println(v48)
                    }
                }
                v50 := GenS6{F0: 470, F1: 215, F2: v49}
                _ = fi11
                _ = v48
                _ = v49
                _ = v50
            }
            _ = fi10
        }
        switch -422 {
        case 0:
            var v51 float32 = (float32(-2.25) + float32(((3.14159 / (float64((((0.1 / 1e10) - 0.1) + (0.0 / 1e10))) / 1.0)) * 1.5)))
            _ = v51
        }
    }
    v52 := GenS6{F0: v8, F1: 427, F2: v4}
    v8 = (^(v8))
    v53 := GenS6{F0: 0, F1: 2, F2: (v4 * 255)}
    v54 := GenS8{F0: 0, F1: (v4 + (v4 >> ((1) & 63))), F2: 42}
    var v55 uint8 = v4
    if ((255 != 1) || true) {
        for fi56 := 0; fi56 < 18; fi56++ {
            if false {
                fmt.Println(v1)
            } else {
                if false {
                    v2 = -181
                    if (2147483648 >= 215) {
                        if (42 <= 0) {
                            var v57 int32 = (-(v2))
                            for fi58 := 0; fi58 < 19; fi58++ {
                                v59 := []int{v1}
                                if ((^(v55)) >= (^(v4))) {
                                    var v60 float32 = 3.14159
                                    var v61 uint32 = (v8 << ((65535) & 63))
                                    _ = v60
                                    _ = v61
                                } else {
                                    v2 -= v57
                                    if (18446744073709551615 >= 360) {
                                        v8 += v8
                                        if (1 >= fi56) {
                                            fmt.Println(fi56)
                                            v62 := GenS9{F0: "hello", F1: 32767, F2: -1}
                                            var v63 int8 = 127
                                            _ = v62
                                            _ = v63
                                        }
                                    }
                                    fmt.Println(v1)
                                    if len(v59) > 0 {
                                        v59[0] = fi56
                                    }
                                }
                                _ = fi58
                                _ = v59
                            }
                            for fi64 := 0; fi64 < 7; fi64++ {
                                v2 -= v2
                                if true {
                                    var v65 int32 = (^(v57))
                                    v66 := GenS1{F0: 65535, F1: 1, F2: (v65 ^ v57)}
                                    _ = v65
                                    _ = v66
                                }
                                _ = fi64
                            }
                            _ = v57
                        } else {
                            v67 := []int{}
                            _ = v67
                        }
                    } else {
                        v68 := GenS3{F0: -30, F1: int(v4), F2: -32768}
                        var v69 uint32 = (v8 + (v8 | v8))
                        _ = v68
                        _ = v69
                    }
                } else {
                    if false {
                        var v70 int64 = -1
                        fmt.Println(v4, v2, v4)
                        fmt.Println(v1, v8)
                        _ = v70
                    } else {
                        var v71 bool = (!(116 != 103))
                        var v72 int64 = (v3 - v3)
                        v73 := []int{}
                        _ = v71
                        _ = v72
                        _ = v73
                    }
                    v3 = v3
                    if (int64(v4) == (v3 - 79)) {
                        v74 := GenS8{F0: -165, F1: 2, F2: 127}
                        _ = v74
                    }
                }
            }
            for fi75 := 0; fi75 < 1; fi75++ {
                v76 := GenS5{F0: true, F1: (v3 | 42)}
                var v77 string = "hello"
                v78 := []int{fi56, fi75, 127}
                _ = fi75
                _ = v76
                _ = v77
                _ = v78
            }
            switch "foo bar" {
            case "a":
                if ((((3.14159 * 0.0) + 0.0) - 1e10) <= (3.14159 / 1.5)) {
                    v79 := []int{v1, (fi56 << ((4294967295) & 63)), v1}
                    v79 = append(v79, (fi56 % ((fi56) | 1)))
                    _ = v79
                }
            }
            _ = fi56
        }
    } else {
        var v80 string = "hello"
        for fi81 := 0; fi81 < 19; fi81++ {
            v82 := []int{fi81, (fi81 >> ((2) & 63)), v1}
            if ((1.5 - (0.0 + 1.5)) > ((0.1 * (((-2.25 / 1e10) + (1e10 - 0.1)) + ((3.14159 - (0.0 / 1.5)) + 0.0))) + (0.1 + float64((1e10 + 0.1))))) {
                for fi83 := 0; fi83 < 19; fi83++ {
                    switch "\n\t" {
                    case "":
                        var v84 uint16 = 378
                        _ = v84
                    case "hello":
                        var v85 uint32 = v8
                        _ = v85
                    default:
                        if (true || true) {
                            switch fi81 {
                            case 949:
                                var v86 uint = 1
                                if ((v2 << (((^(v86))) & 63)) < (v2 | -30)) {
                                    if false {
                                        v1 -= int(v4)
                                        v55 = 1
                                    } else {
                                        _ = len(v82)
                                        var v87 bool = false
                                        fmt.Println(fi81, v86, v80)
                                        _ = v87
                                    }
                                    if (fi81 >= 32767) {
                                        if false {
                                            var v88 int16 = -587
                                            _ = v88
                                        } else {
                                            v86 += 65535
                                            v2 = (-(v2))
                                            v89 := GenS8{F0: 292, F1: v55, F2: -3}
                                            _ = v89
                                        }
                                        if false {
                                            var v90 uint32 = v8
                                            fmt.Println(v1, v86, v1)
                                            _ = v90
                                        }
                                    }
                                } else {
                                    var v91 uint16 = 358
                                    v91 = v91
                                    v3 = -528
                                    fmt.Println(fi81, v8, fi81)
                                    _ = len(v82)
                                    _ = v91
                                }
                                _ = v86
                            case -641:
                                v80 = ""
                            case 0:
                                if len(v82) >= 2 {
                                    v82 = v82[1:len(v82)]
                                }
                            }
                            if (!true) {
                                v92 := GenS5{F0: false, F1: (v3 - (v3 + v3))}
                                switch 4294967295 {
                                case -128:
                                    v82 = append(v82, (^(fi81)))
                                case 4294967295:
                                    v3 = v3
                                default:
                                    var v93 uint32 = v8
                                    _ = v93
                                }
                                if (true && (!(1 < 682))) {
                                    if true {
                                        v2 -= 0
                                        v82 = append(v82, 2147483647)
                                    } else {
                                        if (true || true) {
                                            v4 -= v55
                                        } else {
                                            if false {
                                                v80 = (v80 + "hello")
                                                if len(v82) > 0 {
                                                    v82[0] = 968
                                                }
                                                v1 -= fi81
                                            } else {
                                                var v94 uint16 = 1
                                                var v95 int32 = 0
                                                var v96 uint = 0
                                                _ = v94
                                                _ = v95
                                                _ = v96
                                            }
                                        }
                                        switch (v80 + ("foo bar" + "hello")) {
                                        case "hello":
                                            v97 := GenS10{F0: (^(v55))}
                                            _ = v97
                                        }
                                    }
                                } else {
                                    var v98 string = v80
                                    _ = v98
                                }
                                _ = v92
                            } else {
                                v99 := GenS6{F0: 562, F1: uint16(v8), F2: 255}
                                fmt.Println(fi81, v80)
                                _ = v99
                            }
                            switch v80 {
                            case "foo bar":
                                var v100 string = "a"
                                _ = v100
                            }
                        } else {
                            switch (v1 / (((fi81 >> ((18446744073709551615) & 63))) | 1)) {
                            case 157:
                                if false {
                                    v8 = (^(v8))
                                    v101 := GenS3{F0: fi81, F1: v1, F2: 1000}
                                    _ = v101
                                }
                                if len(v82) > 0 {
                                    v82[0] = 2147483647
                                }
                                switch "a" {
                                case "foo bar":
                                    _ = len(v82)
                                case "\n\t":
                                    if true {
                                        fmt.Println(v1, fi81)
                                        var v102 uint32 = v8
                                        _ = v102
                                    }
                                case "hello":
                                    if len(v82) > 0 {
                                        v82[0] = fi83
                                    }
                                default:
                                    var v103 uint64 = 18446744073709551615
                                    _ = v103
                                }
                            case 32767:
                                v104 := GenS11{F0: "", F1: v8}
                                _ = v104
                            case 127:
                                fmt.Println(v3)
                                if ((0.1 < 0.1) || false) {
                                    v105 := GenS10{F0: 128}
                                    _ = v105
                                }
                                var v106 float64 = 3.14159
                                _ = v106
                            }
                            v82 = append(v82, fi81)
                        }
                    }
                    var v107 float32 = ((((((0.1 - ((-2.25 + 0.0) * 1.5)) - 3.14159) / 3.14159) - (3.14159 * 0.0)) / 1.0) * 0.0)
                    v8 += (v8 | 352)
                    _ = fi83
                    _ = v107
                }
                v4 -= (v55 & v55)
                if false {
                    fmt.Println(v3, v4)
                }
            } else {
                if true {
                    v108 := GenS4{F0: (4294967295 < 0)}
                    var v109 int16 = 221
                    _ = v108
                    _ = v109
                }
                var v110 uint8 = 128
                _ = v110
            }
            fmt.Println(v4)
            _ = fi81
            _ = v82
        }
        _ = v80
    }
    for fi111 := 0; fi111 < 10; fi111++ {
        v112 := GenS12{F0: 255, F1: 41, F2: "foo bar"}
        _ = fi111
        _ = v112
    }
    for fi113 := 0; fi113 < 10; fi113++ {
        fmt.Println(v8, v3)
        v114 := []int{}
        var v115 float64 = 1.5
        _ = fi113
        _ = v114
        _ = v115
    }
    switch "\n\t" {
    case "hello":
        var v116 uint64 = 42
        _ = v116
    case "a":
        if (2 != 4294967295) {
            v117 := []int{}
            _ = v117
        } else {
            switch ("hello" + "") {
            case "foo bar":
                var v118 int64 = v3
                _ = v118
            case "\n\t":
                v4 += v4
            default:
                var v119 int = v1
                _ = v119
            }
            var v120 float32 = -2.25
            _ = v120
        }
    }
    var v121 uint32 = 0
    v121 = v121
    v122 := []int{4294967295, (v1 << ((4294967295) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("v121:", v121)
    fmt.Println("slicelen v122:", len(v122), v122)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v55:", v55)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v121
    _ = v122
    _ = v2
    _ = v3
    _ = v4
    _ = v52
    _ = v53
    _ = v54
    _ = v55
    _ = v8
    _ = v9
}
