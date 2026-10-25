package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int16; F1 uint }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 bool; F1 int; F2 uint32 }
type GenS10 struct { F0 int64; F1 int64 }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 int8; F1 bool; F2 uint8 }
type GenS14 struct { F0 uint16; F1 uint32 }
func genFn0() int {
    for fi4 := 0; fi4 < 11; fi4++ {
        if true {
            var v5 int8 = 127
            _ = v5
        }
        fmt.Println(fi4)
        var v6 int32 = 2147483647
        _ = fi4
        _ = v6
    }
    if false {
        var v7 uint8 = 255
        _ = v7
    } else {
        v8 := []int{128}
        var v9 string = "foo bar"
        _ = v8
        _ = v9
    }
    return -395
}
func genFn1(p0 bool) uint32 {
    var v10 bool = (127 != -33)
    fmt.Println(v10, v10)
    switch "" {
    case "foo bar":
        if true {
            var v11 uint32 = 943
            var v12 bool = p0
            if (686 < 1) {
                for fi13 := 0; fi13 < 13; fi13++ {
                    var v14 string = ""
                    v10 = (127 != 174)
                    p0 = p0
                    _ = fi13
                    _ = v14
                }
            } else {
                v15 := GenS1{F0: 255}
                if v10 {
                    v12 = v10
                    v16 := []int{0, 42, 1}
                    _ = v16
                } else {
                    if v12 {
                        if v12 {
                            for fi17 := 0; fi17 < 14; fi17++ {
                                for fi18 := 0; fi18 < 20; fi18++ {
                                    switch -2147483648 {
                                    case 182:
                                        p0 = (-362 == -9223372036854775808)
                                    case 0:
                                        if (!(4294967295 < 799)) {
                                            v19 := GenS2{F0: 2147483647}
                                            _ = v19
                                        } else {
                                            switch "hello" {
                                            case "":
                                                var v20 uint = 847
                                                _ = v20
                                            case "a":
                                                v11 += (v11 - v11)
                                            case "foo bar":
                                                switch genFn0() {
                                                case 513:
                                                    fmt.Println(fi17, v10, p0)
                                                default:
                                                    var v21 uint16 = 2
                                                    _ = v21
                                                }
                                            }
                                        }
                                    }
                                    v12 = (127 >= 0)
                                    _ = fi18
                                }
                                _ = fi17
                            }
                            if true {
                                v12 = (-1 == 929)
                                v22 := []int{-1, 694}
                                _ = v22
                            } else {
                                for fi23 := 0; fi23 < 9; fi23++ {
                                    v24 := GenS2{F0: 2147483647}
                                    var v25 int32 = -99
                                    _ = fi23
                                    _ = v24
                                    _ = v25
                                }
                                var v26 float64 = (float64((1e10 * 1e10)) * (1.5 * (0.1 * ((float64(1e10) * (((1.5 * ((1e10 * 1e10) * 0.0)) / (0.1 * 0.1)) + 3.14159)) + 0.0))))
                                _ = v26
                            }
                        }
                        for fi27 := 0; fi27 < 1; fi27++ {
                            v12 = ((487 == 956) && v12)
                            v28 := GenS1{F0: 42}
                            v29 := GenS3{F0: (1.5 == 1e10)}
                            _ = fi27
                            _ = v28
                            _ = v29
                        }
                    }
                }
                _ = v15
            }
            _ = v11
            _ = v12
        } else {
            if (42 != 946) {
                for fi30 := 0; fi30 < 1; fi30++ {
                    p0 = (0.0 < 1e10)
                    _ = fi30
                }
                if false {
                    p0 = v10
                }
                v31 := GenS2{F0: -2147483648}
                _ = v31
            }
            v32 := GenS3{F0: ((0.1 > ((3.14159 * -2.25) + 3.14159)) && (65535 >= 32767))}
            _ = v32
        }
    case "hello":
        v10 = v10
    case "\n\t":
        for fi33 := 0; fi33 < 3; fi33++ {
            var v34 int16 = -32768
            fmt.Println(fi33)
            _ = fi33
            _ = v34
        }
    default:
        if p0 {
            v10 = (42 > -1)
            if p0 {
                v35 := GenS1{F0: 255}
                v36 := []int{2147483647, 127}
                _ = v35
                _ = v36
            } else {
                var v37 int8 = 100
                _ = v37
            }
            p0 = (!(995 < 1))
        }
    }
    _ = p0
    _ = v10
    return 320
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint64 = 927
    v3 := []string{"foo bar", "hello", "\n\t", "a"}
    var v38 string = "a"
    for fi39 := 0; fi39 < 19; fi39++ {
        v2 += (v2 ^ 42)
        v38 = "foo bar"
        switch fi39 {
        case 4294967295:
            for fi40 := 0; fi40 < 5; fi40++ {
                fmt.Println(v2, v38, fi39)
                var v41 uint16 = 42
                _ = fi40
                _ = v41
            }
        case 127:
            if true {
                v38 = ("foo bar" + v38)
                var v42 int16 = 32767
                _ = v42
            }
        }
        _ = fi39
    }
    for fi43 := 0; fi43 < 10; fi43++ {
        for fi44 := 0; fi44 < 5; fi44++ {
            switch genFn0() {
            case 0:
                v45 := GenS2{F0: -16}
                _ = v45
            }
            _ = fi44
        }
        if true {
            switch v38 {
            case "foo bar":
                if (42 < 2147483647) {
                    switch v38 {
                    case "\n\t":
                        fmt.Println(fi43)
                    default:
                        if false {
                            for fi46 := 0; fi46 < 15; fi46++ {
                                v47 := GenS4{F0: 42}
                                var v48 uint32 = 693
                                var v49 string = v38
                                _ = fi46
                                _ = v47
                                _ = v48
                                _ = v49
                            }
                        } else {
                            var v50 uint32 = genFn1(true)
                            _ = v50
                        }
                    }
                    switch v38 {
                    case "foo bar":
                        delete(v1, "\n\t")
                    case "":
                        v38 = v38
                    }
                    var v51 float32 = 0.0
                    _ = v51
                }
            case "\n\t":
                if (907 != -1) {
                    v3 = append(v3, "a")
                    var v52 uint = 18446744073709551615
                    var v53 int = 156
                    _ = v52
                    _ = v53
                } else {
                    for fi54 := 0; fi54 < 20; fi54++ {
                        var v55 float64 = (3.14159 * 1e10)
                        var v56 string = v38
                        _ = fi54
                        _ = v55
                        _ = v56
                    }
                    fmt.Println(v2)
                }
            case "a":
                v57 := GenS5{F0: 558}
                _ = v57
            default:
                switch ("hello" + v38) {
                case "\n\t":
                    for fi58 := 0; fi58 < 16; fi58++ {
                        if len(v3) > 0 {
                            v3[0] = ""
                        }
                        var v59 uint8 = 81
                        _ = fi58
                        _ = v59
                    }
                }
            }
        } else {
            v3 = append(v3, v38)
        }
        _ = fi43
    }
    for fi60 := 0; fi60 < 7; fi60++ {
        for fi61 := 0; fi61 < 5; fi61++ {
            switch "a" {
            case "":
                v62 := GenS2{F0: -1}
                _ = v62
            default:
                fmt.Println(fi60)
            }
            var v63 int8 = 1
            switch v38 {
            case "a":
                fmt.Println(fi61)
            case "foo bar":
                switch v38 {
                case "a":
                    if (!true) {
                        delete(v1, "a")
                        v63 = v63
                        v63 -= 1
                    }
                default:
                    var v64 bool = (0 < -2147483648)
                    _ = v64
                }
                if (!(!((v63 | v63) > v63))) {
                    var v65 string = "a"
                    if false {
                        var v66 int32 = -2147483648
                        v67 := GenS6{F0: 42}
                        _ = v66
                        _ = v67
                    } else {
                        v2 = (v2 + v2)
                        v63 = -115
                        var v68 float32 = float32(0.0)
                        if (42 <= (v63 % ((-62) | 1))) {
                            v69 := GenS1{F0: 127}
                            _ = v69
                        }
                        switch (fi60 * genFn0()) {
                        case 78:
                            var v70 string = ((v38 + "hello") + "foo bar")
                            _ = v70
                        case 32768:
                            var v71 float32 = ((((v68 * (float32(-2.25) * ((((0.0 / 1.5) / v68) / v68) - (-2.25 + v68)))) * v68) * v68) / 1.5)
                            _ = v71
                        }
                        if (true && true) {
                            var v72 uint64 = (v2 * 18446744073709551615)
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            fmt.Println(v65)
                            v73 := GenS5{F0: (v72 >> ((4294967295) & 63))}
                            v72 -= 18446744073709551615
                            v1[v65] = 932
                            if ((fi60 + (fi61 >> ((18446744073709551615) & 63))) > fi60) {
                                v74 := v1["a"]
                                if len(v3) > 0 {
                                    v3[0] = v38
                                }
                                switch (fi61 | genFn0()) {
                                case 1:
                                    v75 := GenS7{F0: -464, F1: 65535}
                                    _ = v75
                                default:
                                    switch genFn0() {
                                    case -310:
                                        switch fi61 {
                                        case 0:
                                            v76 := v1["foo bar"]
                                            _ = v76
                                        case -494:
                                            v77 := GenS8{F0: (^(v74))}
                                            _ = v77
                                        case 9223372036854775807:
                                            var v78 uint8 = 2
                                            _ = v78
                                        default:
                                            v2 += v72
                                        }
                                    }
                                }
                                _ = v74
                            } else {
                                if ((true && (false || (v63 == v63))) || (-9223372036854775808 < 0)) {
                                    var v79 int16 = -76
                                    _ = v79
                                }
                                if (true && true) {
                                    var v80 int64 = -9223372036854775808
                                    _ = v80
                                } else {
                                    var v81 int32 = 0
                                    switch (fi60 >> ((647) & 63)) {
                                    case -13:
                                        if false {
                                            if false {
                                                if false {
                                                    v82 := GenS9{F0: (!(32768 <= fi61)), F1: fi60, F2: 2}
                                                    if false {
                                                        switch fi61 {
                                                        case 1:
                                                            var v83 int8 = v63
                                                            _ = v83
                                                        case 681:
                                                            v2 = v72
                                                        case 32768:
                                                            v84 := GenS6{F0: 66}
                                                            _ = v84
                                                        }
                                                        v38 = v65
                                                        _ = len(v1)
                                                    }
                                                    _ = v82
                                                }
                                                fmt.Println(v81)
                                            }
                                            v85 := GenS3{F0: false}
                                            fmt.Println(v68, v68, v2)
                                            _ = v85
                                        } else {
                                            fmt.Println(v81, fi60, v38)
                                            var v86 uint16 = 32767
                                            v87 := v1["\n\t"]
                                            _ = v86
                                            _ = v87
                                        }
                                        fmt.Println(fi61, v63)
                                        v72 = (v72 & v72)
                                    case 32767:
                                        if (!(32767 > 42)) {
                                            if (false && true) {
                                                if false {
                                                    v63 = (v63 * int8(fi61))
                                                } else {
                                                    var v88 uint = 2
                                                    v88 += v88
                                                    if (-394 > 741) {
                                                        v3 = append(v3, v38)
                                                    } else {
                                                        var v89 int32 = (v81 + v81)
                                                        if (true && false) {
                                                            v89 = (v89 | (v81 | (v81 & v89)))
                                                            v81 = v81
                                                            fmt.Println(v81, v89)
                                                            var v90 int64 = 1
                                                            v90 = (v90 << ((4294967295) & 63))
                                                            _ = v90
                                                        } else {
                                                            var v91 int16 = -1
                                                            v92 := GenS10{F0: 631, F1: 9223372036854775807}
                                                            v81 -= 0
                                                            _ = v91
                                                            _ = v92
                                                        }
                                                        var v93 int64 = -444
                                                        _ = v89
                                                        _ = v93
                                                    }
                                                    _ = v88
                                                }
                                            } else {
                                                var v94 int = fi61
                                                var v95 int32 = v81
                                                _ = v94
                                                _ = v95
                                            }
                                            v96 := GenS8{F0: 4294967295}
                                            _ = len(v1)
                                            switch v65 {
                                            case "":
                                                if true {
                                                    if false {
                                                        v63 += (v63 - 0)
                                                        v97 := GenS11{F0: 210}
                                                        v98 := GenS12{F0: 2}
                                                        _ = v97
                                                        _ = v98
                                                    }
                                                    v81 = v81
                                                } else {
                                                    v1[v65] = 2
                                                    if false {
                                                        var v99 uint64 = (v72 | (^(v72)))
                                                        fmt.Println(fi60, v2, v63)
                                                        _ = v99
                                                    }
                                                }
                                            case "foo bar":
                                                v72 += v72
                                            default:
                                                var v100 float64 = 0.1
                                                _ = v100
                                            }
                                            _ = v96
                                        }
                                    case -1:
                                        v101 := GenS13{F0: (v63 % ((v63) | 1)), F1: (255 < 1), F2: 0}
                                        _ = v101
                                    default:
                                        var v102 uint16 = 255
                                        _ = v102
                                    }
                                    switch -1 {
                                    case 128:
                                        if false {
                                            v68 += v68
                                        }
                                    default:
                                        var v103 uint64 = 4294967295
                                        _ = v103
                                    }
                                    _ = v81
                                }
                            }
                            _ = v72
                            _ = v73
                        } else {
                            switch ("foo bar" + v65) {
                            case "\n\t":
                                if true {
                                    v2 = (v2 * 18446744073709551615)
                                    v104 := GenS14{F0: 65535, F1: 272}
                                    _ = v104
                                }
                            case "":
                                v105 := v1[("\n\t" + v65)]
                                _ = v105
                            }
                            v38 = "hello"
                        }
                        _ = v68
                    }
                    _ = v65
                } else {
                    var v106 uint8 = 127
                    v106 = 1
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    if true {
                        var v107 float64 = ((0.0 * (float64(1.5) * (-2.25 + 1.5))) / 0.1)
                        fmt.Println(fi60)
                        _ = v107
                    } else {
                        fmt.Println(v2, v106)
                        var v108 float64 = ((3.14159 * 3.14159) * ((-2.25 - 0.0) - (1.5 / (float64(1.5) - 0.1))))
                        v38 = "hello"
                        _ = v108
                    }
                    var v109 float32 = float32(3.14159)
                    _ = v106
                    _ = v109
                }
                if true {
                    v2 = 0
                    v38 = ((v38 + ((v38 + (v38 + v38)) + "a")) + ("hello" + "foo bar"))
                    v2 = v2
                    v1[v38] = 2147483648
                }
            case "hello":
                switch genFn0() {
                case 0:
                    v3 = append(v3, v38)
                case -9223372036854775808:
                    fmt.Println(fi61)
                case 42:
                    fmt.Println(v38, fi61, fi60)
                }
            default:
                v63 = v63
            }
            v63 = v63
            _ = fi61
            _ = v63
        }
        for fi110 := 0; fi110 < 17; fi110++ {
            var v111 string = v38
            fmt.Println(fi60)
            v112 := GenS6{F0: 4294967295}
            _ = fi110
            _ = v111
            _ = v112
        }
        _ = fi60
    }
    v113 := GenS1{F0: 42}
    if (!((!true) && true)) {
        v2 = 748
        v2 += (^(v2))
    }
    fmt.Println(v38)
    if true {
        v2 += (^(v2))
    } else {
        var v114 int16 = -746
        for fi115 := 0; fi115 < 16; fi115++ {
            switch (v38 + ((v38 + (v38 + v38)) + (((((v38 + v38) + (("a" + "hello") + "a")) + ((v38 + v38) + v38)) + ("foo bar" + (("a" + v38) + v38))) + v38))) {
            case "hello":
                var v116 uint32 = 0
                _ = v116
            case "foo bar":
                v117 := GenS12{F0: 2147483647}
                _ = v117
            case "a":
                var v118 uint = 2
                _ = v118
            }
            _ = fi115
        }
        for fi119 := 0; fi119 < 19; fi119++ {
            v120 := GenS11{F0: 2}
            v121 := GenS6{F0: 555}
            v122 := GenS3{F0: (false && (-26 > 0))}
            _ = fi119
            _ = v120
            _ = v121
            _ = v122
        }
        _ = v114
    }
    switch 1 {
    case 1:
        for fi123 := 0; fi123 < 2; fi123++ {
            for fi124 := 0; fi124 < 17; fi124++ {
                delete(v1, ((v38 + v38) + "hello"))
                _ = fi124
            }
            _ = fi123
        }
    case 252:
        v38 = (v38 + (("" + v38) + "\n\t"))
    default:
        switch "foo bar" {
        case "foo bar":
            fmt.Println(v38)
        case "a":
            for fi125 := 0; fi125 < 20; fi125++ {
                delete(v1, "foo bar")
                _ = fi125
            }
        case "hello":
            v38 = v38
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v113
    _ = v2
    _ = v3
    _ = v38
}
