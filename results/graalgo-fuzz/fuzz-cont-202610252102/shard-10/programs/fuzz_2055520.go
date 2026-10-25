package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 uint; F1 uint8; F2 uint8 }
type GenS3 struct { F0 string; F1 int }
type GenS4 struct { F0 uint; F1 int; F2 int32 }
type GenS5 struct { F0 string; F1 int64; F2 bool }
type GenS6 struct { F0 int16; F1 uint; F2 uint }
type GenS7 struct { F0 int64; F1 int64 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int64; F1 uint32 }
type GenS10 struct { F0 uint; F1 bool }
type GenS11 struct { F0 int; F1 int64; F2 int }
func genFn0(p0 uint32) int32 {
    p0 = 442
    var v7 float64 = 0.0
    var v8 uint64 = 18446744073709551615
    v9 := GenS1{F0: (!true), F1: 1}
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 float64 = float64(1e10)
    var v2 bool = true
    v3 := []uint32{4294967295, 1, 758}
    v4 := make(map[uint32]bool)
    var v5 int64 = -1
    var v6 uint64 = 449
    switch (("\n\t" + ("\n\t" + "hello")) + "a") {
    case "a":
        fmt.Println(v5, v1)
    }
    if (3.14159 != v1) {
        for fi10 := 0; fi10 < 20; fi10++ {
            for fi11 := 0; fi11 < 1; fi11++ {
                var v12 int8 = 85
                if (!v2) {
                    var v13 uint = 2
                    if true {
                        v14 := GenS2{F0: (v13 & v13), F1: 127, F2: 255}
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = v14
                    }
                    var v15 int16 = 453
                    _ = v13
                    _ = v15
                } else {
                    if (42 != 0) {
                        v16 := GenS1{F0: v2, F1: 0}
                        var v17 string = "hello"
                        fmt.Println(fi10, v6)
                        v2 = (1000 == 400)
                        v18 := v4[2147483648]
                        _ = v16
                        _ = v17
                        _ = v18
                    } else {
                        v19 := GenS2{F0: 1, F1: 2, F2: 152}
                        _ = v19
                    }
                    var v20 float32 = (0.0 / float32(3.14159))
                    fmt.Println(v1, v6, fi10)
                    _ = v20
                }
                _ = fi11
                _ = v12
            }
            switch "foo bar" {
            case "a":
                v4[160] = v2
            }
            _ = fi10
        }
        v21 := GenS2{F0: 2, F1: 242, F2: 255}
        if v2 {
            for fi22 := 0; fi22 < 16; fi22++ {
                switch -4 {
                case 128:
                    fmt.Println(fi22, fi22)
                case -111:
                    for fi23 := 0; fi23 < 1; fi23++ {
                        var v24 int = 127
                        fmt.Println(v6)
                        _ = fi23
                        _ = v24
                    }
                case 0:
                    v2 = (!v2)
                default:
                    if len(v3) > 0 {
                        v3[0] = 0
                    }
                }
                if true {
                    fmt.Println(fi22)
                    v6 += 2
                    v6 = v6
                }
                v25 := v4[42]
                _ = fi22
                _ = v25
            }
            var v26 float32 = float32(v1)
            _ = v26
        } else {
            fmt.Println(v1, v6)
            if (65535 > 74) {
                v2 = (-1 > 127)
            }
            var v27 uint64 = (v6 >> ((1) & 63))
            _ = v27
        }
        _ = v21
    } else {
        var v28 float32 = 0.0
        var v29 uint8 = 49
        _ = v28
        _ = v29
    }
    switch ("foo bar" + "\n\t") {
    case "foo bar":
        var v30 uint16 = 2
        _ = v30
    case "hello":
        switch 32768 {
        case -911:
            switch 851 {
            case 0:
                v5 = 1
            default:
                v1 -= (v1 + (float64(1e10) * v1))
            }
        case 883:
            var v31 int32 = genFn0(970)
            _ = v31
        case 32768:
            fmt.Println(v6, v5)
        default:
            if (false && false) {
                if v2 {
                    if (2 != 163) {
                        v32 := GenS2{F0: 2, F1: 251, F2: 0}
                        fmt.Println(v5, v6)
                        v33 := GenS3{F0: "hello", F1: -579}
                        _ = v32
                        _ = v33
                    } else {
                        switch -1 {
                        case 32768:
                            var v34 int8 = -128
                            _ = v34
                        case 502:
                            var v35 string = ((("" + "\n\t") + (("hello" + ("a" + "foo bar")) + ("\n\t" + "\n\t"))) + "foo bar")
                            _ = v35
                        }
                        var v36 int64 = 9223372036854775807
                        _ = v36
                    }
                    var v37 uint = 65535
                    _ = v37
                }
            }
        }
    case "\n\t":
        v38 := GenS4{F0: 1, F1: 410, F2: 2147483647}
        _ = v38
    }
    v1 = (float64(1.5) / ((v1 / 1.5) + float64(1.5)))
    if false {
        v6 = (^(v6))
        v4[2147483647] = v2
        if v2 {
            delete(v4, 42)
        } else {
            v5 = (v5 & (-(v5)))
        }
    }
    for fi39 := 0; fi39 < 4; fi39++ {
        fmt.Println(v1)
        _ = fi39
    }
    var v40 uint = 18446744073709551615
    for fi41 := 0; fi41 < 7; fi41++ {
        if v2 {
            var v42 int8 = 0
            _ = v42
        } else {
            var v43 uint = (v40 / ((uint(v6)) | 1))
            switch "hello" {
            case "\n\t":
                fmt.Println(v43, v43)
            case "":
                v6 -= 390
            case "hello":
                v3 = append(v3, 553)
            }
            _ = v43
        }
        _ = fi41
    }
    var v44 string = ""
    for fi45 := 0; fi45 < 14; fi45++ {
        for fi46 := 0; fi46 < 4; fi46++ {
            if ((-(v5)) <= 0) {
                var v47 int16 = 400
                v47 -= 32767
                _ = v47
            } else {
                fmt.Println(fi46)
            }
            v6 = v6
            _ = fi46
        }
        v1 = (v1 / v1)
        _ = fi45
    }
    var v48 bool = v2
    fmt.Println(v40)
    var v49 int32 = 2147483647
    var v50 float64 = 0.1
    if v48 {
        v51 := GenS1{F0: ((v6 / (((^(v6))) | 1)) == 441), F1: 2}
        _ = v51
    } else {
        v5 -= (^(v5))
        for fi52 := 0; fi52 < 1; fi52++ {
            fmt.Println(v2, v44)
            if (true && (-128 < 9)) {
                if (!v48) {
                    _ = len(v3)
                    if (1 <= 172) {
                        fmt.Println(v5)
                        switch 2147483647 {
                        case 887:
                            v5 -= (v5 << ((42) & 63))
                        case 42:
                            if (int8(v6) > -1) {
                                v50 += ((v1 * float64((float32((((1e10 + float64(((1.5 * (-2.25 - float32(v50))) * 0.1))) * ((1.5 * (v50 / v50)) / v1)) - -2.25)) - 1e10))) - v50)
                            }
                        case 2147483647:
                            var v53 string = "\n\t"
                            _ = v53
                        default:
                            if len(v3) > 0 {
                                v3[0] = 2147483647
                            }
                        }
                        switch fi52 {
                        case 906:
                            for fi54 := 0; fi54 < 19; fi54++ {
                                if (1.5 <= (0.0 - 1e10)) {
                                    var v55 int = fi52
                                    _ = v55
                                } else {
                                    v56 := GenS1{F0: true, F1: 2}
                                    var v57 uint32 = 2147483647
                                    v58 := GenS2{F0: v40, F1: 0, F2: 1}
                                    v59 := GenS5{F0: (v44 + v44), F1: v5, F2: ((float64(1e10) / ((v1 - 0.1) * v50)) >= (float64(((((1.5 + 0.0) + (float32(v1) * -2.25)) / 1e10) / ((3.14159 - 0.0) * float32(((v1 * (3.14159 / v1)) + (float64((float32(1.5) - ((1.5 + 3.14159) - -2.25))) + (3.14159 * v1))))))) - v1))}
                                    switch (fi54 << ((v40) & 63)) {
                                    case 2147483647:
                                        var v60 float64 = v1
                                        v61 := GenS6{F0: 1000, F1: (v40 << ((v40) & 63)), F2: (v40 >> ((v40) & 63))}
                                        if (v44 >= "") {
                                            v57 = v57
                                        } else {
                                            var v62 int = (-(fi52))
                                            _ = v62
                                        }
                                        _ = v60
                                        _ = v61
                                    case 30:
                                        if ((-2.25 == 3.14159) && false) {
                                            switch fi52 {
                                            case 1:
                                                if true {
                                                    v63 := GenS3{F0: "foo bar", F1: -2147483648}
                                                    _ = v63
                                                }
                                            case 4294967295:
                                                delete(v4, uint32(v6))
                                            case -84:
                                                v64 := GenS3{F0: ((("a" + v44) + ("\n\t" + ((("" + v44) + ("foo bar" + v44)) + v44))) + v44), F1: (fi52 >> ((v40) & 63))}
                                                _ = v64
                                            default:
                                                v6 = v6
                                            }
                                            var v65 uint16 = 255
                                            v66 := GenS1{F0: (v57 == v57), F1: 631}
                                            _ = v65
                                            _ = v66
                                        }
                                    case 128:
                                        if len(v3) > 0 {
                                            v3[0] = (v57 >> ((v40) & 63))
                                        }
                                    default:
                                        v67 := GenS2{F0: (^(v40)), F1: 91, F2: 113}
                                        _ = v67
                                    }
                                    _ = v56
                                    _ = v57
                                    _ = v58
                                    _ = v59
                                }
                                var v68 uint64 = (v6 * v6)
                                fmt.Println(fi52)
                                _ = fi54
                                _ = v68
                            }
                        }
                    }
                    fmt.Println(v50, v49)
                } else {
                    if (((v2 && v48) && v48) && ((v50 * (1.5 / 1e10)) <= (((v50 / v1) + 1e10) - 0.0))) {
                        if (false || v2) {
                            v3 = append(v3, 2147483647)
                            for fi69 := 0; fi69 < 2; fi69++ {
                                var v70 int64 = (v5 + -9223372036854775808)
                                if true {
                                    v71 := GenS7{F0: (v5 | 9223372036854775807), F1: v70}
                                    if (-9223372036854775808 != v70) {
                                        v72 := GenS7{F0: v5, F1: (^(v70))}
                                        _ = v72
                                    } else {
                                        v48 = false
                                        if v48 {
                                            if v2 {
                                                if v2 {
                                                    var v73 uint8 = 28
                                                    v74 := v4[65535]
                                                    v75 := GenS8{F0: v40}
                                                    _ = v73
                                                    _ = v74
                                                    _ = v75
                                                }
                                            } else {
                                                v49 += 1
                                                if (!(!false)) {
                                                    if v48 {
                                                        fmt.Println(fi52, fi52)
                                                        v76 := v4[80]
                                                        _ = v76
                                                    }
                                                    if (-72 < 96) {
                                                        switch (fi69 & -1) {
                                                        case -789:
                                                            fmt.Println(v50, v2)
                                                        case 788:
                                                            var v77 uint8 = 182
                                                            if (!v2) {
                                                                if ((!(2147483648 < 65535)) || false) {
                                                                    v70 = -9223372036854775808
                                                                    v48 = (genFn0(2) != v49)
                                                                    fmt.Println(v40, fi52, v77)
                                                                } else {
                                                                    v6 -= 641
                                                                    delete(v4, uint32(v6))
                                                                    if ((float32((((((v50 + (v50 * (float64(1.5) / v50))) / v50) - 0.0) - 0.1) / v50)) + 1.5) <= 3.14159) {
                                                                        fmt.Println(v48, v2, v49)
                                                                        if ((v48 && (!true)) || (!false)) {
                                                                            delete(v4, 42)
                                                                            v78 := GenS7{F0: v70, F1: (v5 + int64(v77))}
                                                                            _ = v78
                                                                        } else {
                                                                            fmt.Println(fi69, v44, v6)
                                                                        }
                                                                    } else {
                                                                        switch (^(fi52)) {
                                                                        case -2147483648:
                                                                            var v79 uint32 = 483
                                                                            _ = v79
                                                                        case 9223372036854775807:
                                                                            _ = len(v3)
                                                                        case 660:
                                                                            v1 = float64((1e10 + (1.5 + 0.0)))
                                                                        default:
                                                                            var v80 int8 = 0
                                                                            _ = v80
                                                                        }
                                                                        v2 = v2
                                                                        v49 -= (v49 ^ genFn0(2147483647))
                                                                        var v81 int = 324
                                                                        if (42 >= -1) {
                                                                            v4[uint32(v6)] = (!(v40 >= v40))
                                                                            if (!(false && false)) {
                                                                                if len(v3) > 0 {
                                                                                    v3[0] = 65535
                                                                                }
                                                                            }
                                                                        }
                                                                        v82 := v4[0]
                                                                        _ = v81
                                                                        _ = v82
                                                                    }
                                                                }
                                                            } else {
                                                                var v83 bool = (v77 < (v77 ^ 0))
                                                                _ = v83
                                                            }
                                                            _ = v77
                                                        case -433:
                                                            if len(v3) > 0 {
                                                                v3[0] = 189
                                                            }
                                                            v70 = 0
                                                        }
                                                        var v84 string = v44
                                                        _ = v84
                                                    } else {
                                                        switch (v44 + (((v44 + ("a" + v44)) + (v44 + v44)) + "hello")) {
                                                        case "\n\t":
                                                            switch 0 {
                                                            case 9223372036854775807:
                                                                v85 := GenS5{F0: (v44 + "foo bar"), F1: v70, F2: (255 > v40)}
                                                                if (-2.25 == 0.1) {
                                                                    if v48 {
                                                                        v86 := GenS8{F0: v40}
                                                                        var v87 uint32 = 0
                                                                        v50 = (0.1 / v1)
                                                                        v88 := v4[4294967295]
                                                                        fmt.Println(fi52)
                                                                        _ = v86
                                                                        _ = v87
                                                                        _ = v88
                                                                    } else {
                                                                        var v89 float32 = 1e10
                                                                        v90 := GenS4{F0: v40, F1: 4294967295, F2: int32(v6)}
                                                                        _ = v89
                                                                        _ = v90
                                                                    }
                                                                    var v91 uint8 = 0
                                                                    v3 = append(v3, 0)
                                                                    _ = v91
                                                                }
                                                                _ = v85
                                                            case -270:
                                                                v6 = (v6 - v6)
                                                            case 32767:
                                                                var v92 uint64 = (^(v6))
                                                                _ = v92
                                                            }
                                                        }
                                                    }
                                                    v5 -= (v70 >> (((v40 * v40)) & 63))
                                                }
                                                if false {
                                                    delete(v4, uint32(v40))
                                                    v93 := GenS2{F0: 42, F1: 255, F2: 127}
                                                    switch fi52 {
                                                    case -1:
                                                        _ = len(v3)
                                                    default:
                                                        var v94 bool = v2
                                                        _ = v94
                                                    }
                                                    _ = v93
                                                }
                                            }
                                            var v95 bool = false
                                            _ = v95
                                        }
                                        var v96 int = fi52
                                        _ = v96
                                    }
                                    _ = v71
                                }
                                var v97 int32 = -2147483648
                                _ = fi69
                                _ = v70
                                _ = v97
                            }
                            switch int(v49) {
                            case -339:
                                if (((v5 & (v5 >> (((^(v40))) & 63))) != -834) && false) {
                                    for fi98 := 0; fi98 < 5; fi98++ {
                                        switch "a" {
                                        case "hello":
                                            var v99 uint16 = 238
                                            _ = v99
                                        case "a":
                                            var v100 bool = (true || false)
                                            _ = v100
                                        }
                                        v49 += v49
                                        _ = fi98
                                    }
                                } else {
                                    v101 := GenS3{F0: ("hello" + v44), F1: 42}
                                    fmt.Println(v1, v49)
                                    fmt.Println(v44)
                                    _ = v101
                                }
                            }
                        }
                        v4[0] = (v2 || (42 == 127))
                    }
                    for fi102 := 0; fi102 < 7; fi102++ {
                        switch fi52 {
                        case 0:
                            var v103 uint64 = v6
                            var v104 int8 = 42
                            v48 = v2
                            switch fi52 {
                            case -129:
                                v105 := GenS9{F0: (v5 | (v5 % ((149) | 1))), F1: 859}
                                _ = v105
                            case -9:
                                switch (v44 + v44) {
                                case "\n\t":
                                    v106 := GenS10{F0: (v40 + (^(v40))), F1: v48}
                                    _ = v106
                                case "hello":
                                    var v107 uint8 = 61
                                    _ = v107
                                default:
                                    if (255 <= 127) {
                                        var v108 uint16 = 65535
                                        var v109 int64 = 9223372036854775807
                                        var v110 int = (fi102 | 32767)
                                        _ = v108
                                        _ = v109
                                        _ = v110
                                    }
                                }
                            case -1:
                                v111 := GenS8{F0: (v40 ^ (^(v40)))}
                                _ = v111
                            }
                            _ = v103
                            _ = v104
                        }
                        _ = fi102
                    }
                    fmt.Println(v49, v5, v49)
                }
                _ = len(v4)
            }
            v112 := GenS11{F0: (fi52 | fi52), F1: (v5 + v5), F2: (fi52 + 1)}
            _ = fi52
            _ = v112
        }
    }
    for fi113 := 0; fi113 < 1; fi113++ {
        for fi114 := 0; fi114 < 8; fi114++ {
            v49 += v49
            _ = fi114
        }
        _ = fi113
    }
    if len(v3) > 0 {
        v3[0] = 2
    }
    switch -2147483648 {
    case 32768:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case 9223372036854775807:
        v115 := v4[0]
        _ = v115
    default:
        v3 = append(v3, 0)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v44:", v44)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v44
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v6
}
