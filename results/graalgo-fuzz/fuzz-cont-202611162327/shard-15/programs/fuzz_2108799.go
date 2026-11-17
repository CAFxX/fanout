package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 bool; F1 uint8; F2 uint }
type GenS3 struct { F0 string; F1 uint16; F2 string }
type GenS4 struct { F0 int; F1 uint8 }
type GenS5 struct { F0 bool; F1 uint8; F2 uint }
type GenS6 struct { F0 uint64; F1 uint16 }
func genFn0() uint16 {
    v5 := GenS1{F0: (true && false), F1: 1}
    var v6 int32 = 897
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 int16 = 32767
        _ = fi7
        _ = v8
    }
    v6 += 1
    var v9 uint16 = 65535
    _ = v5
    _ = v6
    _ = v9
    return 1
}
func genFn1(p0 string, p1 int) int {
    for fi10 := 0; fi10 < 15; fi10++ {
        v11 := []int{}
        for fi12 := 0; fi12 < 7; fi12++ {
            p0 = p0
            _ = fi12
        }
        switch fi10 {
        case 42:
            for fi13 := 0; fi13 < 8; fi13++ {
                if false {
                    p0 = p0
                    p0 = (p0 + p0)
                    var v14 float32 = 3.14159
                    if len(v11) >= 2 {
                        v11 = v11[1:len(v11)]
                    }
                    var v15 uint = 65535
                    _ = v14
                    _ = v15
                } else {
                    var v16 int32 = 167
                    _ = v16
                }
                _ = fi13
            }
        default:
            p1 += 1
        }
        _ = fi10
        _ = v11
    }
    if false {
        p1 -= p1
        for fi17 := 0; fi17 < 5; fi17++ {
            var v18 uint = 2
            var v19 int16 = -853
            _ = fi17
            _ = v18
            _ = v19
        }
    } else {
        v20 := GenS1{F0: true, F1: 0}
        var v21 bool = (255 != 42)
        p1 = (p1 * (p1 + p1))
        _ = v20
        _ = v21
    }
    _ = p0
    _ = p1
    return 9223372036854775807
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]int)
    var v3 bool = (!false)
    var v4 int16 = 1000
    switch -157 {
    case 127:
        fmt.Println(v3, v3)
    }
    switch -2147483648 {
    case 0:
        for fi22 := 0; fi22 < 14; fi22++ {
            if (643 >= 0) {
                v23 := GenS2{F0: v3, F1: 42, F2: 255}
                v24 := GenS3{F0: "hello", F1: 1, F2: "hello"}
                v25 := []int{genFn1(((("foo bar" + "a") + ((("a" + "foo bar") + ("" + (("foo bar" + "a") + ""))) + "a")) + ("a" + ((("hello" + "hello") + "foo bar") + ("\n\t" + "foo bar")))), genFn1(("foo bar" + "foo bar"), int(v4)))}
                _ = v23
                _ = v24
                _ = v25
            } else {
                if true {
                    v26 := GenS3{F0: ((((("a" + ("foo bar" + "hello")) + "\n\t") + "foo bar") + "foo bar") + "\n\t"), F1: 323, F2: ("\n\t" + (("" + "foo bar") + "foo bar"))}
                    _ = v26
                } else {
                    var v27 int = (fi22 / ((fi22) | 1))
                    if (!v3) {
                        var v28 int64 = 517
                        fmt.Println(fi22, v3, fi22)
                        var v29 uint32 = 0
                        _ = v28
                        _ = v29
                    }
                    var v30 int8 = -128
                    _ = v27
                    _ = v30
                }
                v4 = v4
            }
            _ = fi22
        }
    }
    v31 := GenS2{F0: (9223372036854775807 == 984), F1: 161, F2: 65535}
    v32 := GenS2{F0: (628 > 2), F1: 3, F2: 0}
    if true {
        v4 = v4
    } else {
        var v33 int8 = -1
        switch ("hello" + "a") {
        case "":
            fmt.Println(v33)
        case "\n\t":
            var v34 uint = 42
            _ = v34
        }
        fmt.Println(v3)
        _ = v33
    }
    v3 = (int32(v4) != 42)
    if (!v3) {
        switch "\n\t" {
        case "\n\t":
            v35 := []int{}
            _ = v35
        case "a":
            switch 1 {
            case -508:
                if v3 {
                    v3 = v3
                    v4 -= 32767
                } else {
                    if (9223372036854775807 == -929) {
                        var v36 int64 = 42
                        if (3.14159 != 0.1) {
                            var v37 int32 = -1
                            switch 32767 {
                            case 32768:
                                var v38 float64 = -2.25
                                _ = v38
                            case -9223372036854775808:
                                if true {
                                    if ((1e10 > 1.5) || v3) {
                                        _ = len(v1)
                                        switch -696 {
                                        case -1:
                                            if v3 {
                                                v4 += v4
                                                v39 := GenS2{F0: ((false && v3) && (v3 && v3)), F1: 102, F2: 2}
                                                _ = v39
                                            } else {
                                                var v40 string = "a"
                                                _ = v40
                                            }
                                        case -262:
                                            v37 = (v37 | 42)
                                        case -855:
                                            v37 -= int32(v36)
                                        }
                                    } else {
                                        switch 32768 {
                                        case -2147483648:
                                            v2[genFn1("foo bar", 1)] = 1
                                        default:
                                            v41 := []int{genFn1("a", 425)}
                                            _ = v41
                                        }
                                        v42 := v1[581]
                                        fmt.Println(v42, v3)
                                        _ = v42
                                    }
                                } else {
                                    switch ("\n\t" + "hello") {
                                    case "foo bar":
                                        v43 := v1[32768]
                                        _ = v43
                                    case "hello":
                                        switch "hello" {
                                        case "foo bar":
                                            v4 -= (v4 | v4)
                                        case "a":
                                            v44 := GenS2{F0: (42 >= 255), F1: 128, F2: 42}
                                            _ = v44
                                        default:
                                            var v45 uint16 = 32767
                                            _ = v45
                                        }
                                    }
                                }
                            }
                            if v3 {
                                v36 = -9223372036854775808
                            } else {
                                for fi46 := 0; fi46 < 14; fi46++ {
                                    v47 := v1[genFn1("\n\t", (fi46 | (fi46 % ((fi46) | 1))))]
                                    v48 := GenS4{F0: fi46, F1: 127}
                                    v3 = (-579 < (v36 & -477))
                                    _ = fi46
                                    _ = v47
                                    _ = v48
                                }
                                var v49 int = 32767
                                _ = v49
                            }
                            _ = v37
                        }
                        for fi50 := 0; fi50 < 9; fi50++ {
                            if ((292 == 718) && v3) {
                                v4 += (v4 - v4)
                                for fi51 := 0; fi51 < 3; fi51++ {
                                    if v3 {
                                        var v52 int64 = 9223372036854775807
                                        v52 += v36
                                        _ = v52
                                    } else {
                                        v53 := []int{genFn1(("\n\t" + ("foo bar" + "foo bar")), fi50), -9223372036854775808, (fi51 | (fi50 >> ((18446744073709551615) & 63))), fi51}
                                        v36 = -1
                                        _ = len(v53)
                                        _ = v53
                                    }
                                    var v54 uint64 = 42
                                    var v55 int8 = 42
                                    fmt.Println(v54, v54)
                                    v4 += 1000
                                    v4 -= int16(v36)
                                    _ = fi51
                                    _ = v54
                                    _ = v55
                                }
                                for fi56 := 0; fi56 < 15; fi56++ {
                                    if v3 {
                                        switch "" {
                                        case "":
                                            v57 := GenS2{F0: v3, F1: 42, F2: 42}
                                            _ = v57
                                        case "a":
                                            fmt.Println(v3, v3, v4)
                                        case "foo bar":
                                            fmt.Println(fi56, v3)
                                        }
                                        v36 -= (v36 & v36)
                                        v3 = false
                                    }
                                    switch "a" {
                                    case "a":
                                        var v58 uint32 = 4294967295
                                        _ = v58
                                    case "foo bar":
                                        fmt.Println(fi56, v4)
                                    default:
                                        if true {
                                            if false {
                                                v3 = (1e10 <= 0.1)
                                                v4 = v4
                                            }
                                            if v3 {
                                                switch genFn1("foo bar", int(v4)) {
                                                case -951:
                                                    v59 := GenS2{F0: (!(v3 || (false && v3))), F1: 128, F2: 255}
                                                    switch "foo bar" {
                                                    case "a":
                                                        v60 := GenS1{F0: ("a" <= ("" + ("foo bar" + ((("foo bar" + ("a" + "foo bar")) + "hello") + "a")))), F1: 0}
                                                        _ = v60
                                                    case "hello":
                                                        v4 += v4
                                                    case "":
                                                        v3 = ((65535 > 255) || v3)
                                                    }
                                                    v2[(^(fi56))] = fi50
                                                    if (("hello" + "hello") < "hello") {
                                                        var v61 int8 = 91
                                                        _ = v61
                                                    } else {
                                                        v62 := GenS5{F0: true, F1: 255, F2: 1}
                                                        v36 = v36
                                                        v63 := GenS1{F0: v3, F1: -1}
                                                        _ = v62
                                                        _ = v63
                                                    }
                                                    v36 -= (v36 - (v36 >> ((203) & 63)))
                                                    _ = v59
                                                case 2147483647:
                                                    var v64 int8 = -1
                                                    v64 = (^(v64))
                                                    v64 = (v64 / (((v64 & v64)) | 1))
                                                    v36 = 1
                                                    _ = v64
                                                default:
                                                    fmt.Println(v3, v4)
                                                }
                                                if (18446744073709551615 == 4294967295) {
                                                    var v65 int8 = -24
                                                    v4 += v4
                                                    _ = v65
                                                } else {
                                                    v66 := GenS4{F0: (fi56 >> ((65535) & 63)), F1: 128}
                                                    var v67 uint8 = 2
                                                    var v68 uint16 = 65535
                                                    _ = v66
                                                    _ = v67
                                                    _ = v68
                                                }
                                            }
                                        }
                                    }
                                    switch (("\n\t" + "") + "a") {
                                    case "":
                                        var v69 int64 = v36
                                        if v3 {
                                            v4 = 32767
                                        } else {
                                            switch "foo bar" {
                                            case "":
                                                switch "foo bar" {
                                                case "a":
                                                    v3 = false
                                                case "foo bar":
                                                    switch ("" + "hello") {
                                                    case "":
                                                        fmt.Println(v36, v4)
                                                    case "foo bar":
                                                        v69 -= (v36 + -1)
                                                    case "hello":
                                                        v36 = -9223372036854775808
                                                    }
                                                case "hello":
                                                    v36 = v36
                                                }
                                            case "\n\t":
                                                v70 := GenS2{F0: false, F1: 191, F2: 42}
                                                _ = v70
                                            case "foo bar":
                                                v71 := GenS4{F0: fi56, F1: 77}
                                                _ = v71
                                            }
                                        }
                                        _ = v69
                                    case "foo bar":
                                        fmt.Println(fi50, v36, v3)
                                    }
                                    _ = fi56
                                }
                            }
                            switch "hello" {
                            case "foo bar":
                                v72 := []int{(fi50 ^ (-(fi50))), 816, fi50, (^(fi50))}
                                _ = v72
                            default:
                                switch fi50 {
                                case 127:
                                    v73 := []int{fi50}
                                    _ = v73
                                case 4294967295:
                                    if false {
                                        switch (fi50 % (((fi50 >> ((18446744073709551615) & 63))) | 1)) {
                                        case -166:
                                            delete(v1, (fi50 ^ fi50))
                                        }
                                        var v74 uint16 = 32767
                                        for fi75 := 0; fi75 < 12; fi75++ {
                                            if (!v3) {
                                                v4 = (-(v4))
                                            }
                                            v76 := []int{(-(fi75)), 990}
                                            if (!false) {
                                                v77 := GenS3{F0: "a", F1: (v74 + (v74 + v74)), F2: "foo bar"}
                                                v78 := GenS4{F0: (fi50 & -1), F1: 127}
                                                v2[fi75] = 42
                                                _ = v77
                                                _ = v78
                                            } else {
                                                if v3 {
                                                    switch "" {
                                                    case "hello":
                                                        fmt.Println(v3)
                                                    case "":
                                                        v4 += v4
                                                    case "\n\t":
                                                        var v79 float64 = float64((((1.5 - 1.5) + float32((3.14159 + ((((0.1 / (float64(float32(1.5)) / 1.0)) + 0.0) / 1.5) - 1.5)))) + 1.5))
                                                        _ = v79
                                                    default:
                                                        v3 = (!(!(v3 || true)))
                                                    }
                                                    v80 := v2[genFn1(("" + "foo bar"), (fi50 / (((-(fi50))) | 1)))]
                                                    _ = len(v1)
                                                    _ = v80
                                                }
                                                var v81 bool = true
                                                _ = v81
                                            }
                                            var v82 uint8 = 56
                                            _ = fi75
                                            _ = v76
                                            _ = v82
                                        }
                                        _ = v74
                                    } else {
                                        for fi83 := 0; fi83 < 7; fi83++ {
                                            switch 376 {
                                            case 4294967295:
                                                if (!(659 > 351)) {
                                                    fmt.Println(fi83, v36)
                                                    var v84 float64 = 3.14159
                                                    v3 = (2 < 4294967295)
                                                    _ = v84
                                                } else {
                                                    v85 := []int{genFn1("\n\t", (fi83 & (fi50 - (-(fi83))))), 32768, fi83, (fi50 % ((genFn1(("a" + "foo bar"), (^(fi83)))) | 1))}
                                                    v3 = (65535 != 1)
                                                    var v86 int16 = v4
                                                    _ = v85
                                                    _ = v86
                                                }
                                                switch fi83 {
                                                case 0:
                                                    var v87 uint = 18446744073709551615
                                                    _ = v87
                                                default:
                                                    fmt.Println(v4, v3, v36)
                                                }
                                                v88 := GenS4{F0: (fi50 >> ((0) & 63)), F1: 1}
                                                _ = v88
                                            case -9223372036854775808:
                                                var v89 float64 = 3.14159
                                                _ = v89
                                            case -403:
                                                var v90 uint64 = 42
                                                _ = v90
                                            default:
                                                v36 -= (-(v36))
                                            }
                                            _ = fi83
                                        }
                                        for fi91 := 0; fi91 < 8; fi91++ {
                                            fmt.Println(v36, v4, v4)
                                            _ = fi91
                                        }
                                        var v92 string = "\n\t"
                                        _ = v92
                                    }
                                case -950:
                                    v4 -= v4
                                default:
                                    var v93 int16 = 0
                                    _ = v93
                                }
                            }
                            var v94 float32 = float32((((-2.25 * float64(-2.25)) * float64(3.14159)) * 0.0))
                            _ = fi50
                            _ = v94
                        }
                        _ = v36
                    } else {
                        v3 = true
                        v95 := GenS6{F0: 257, F1: genFn0()}
                        _ = v95
                    }
                }
            }
        default:
            var v96 int16 = v4
            _ = v96
        }
        var v97 uint = 4294967295
        _ = v97
    }
    if (v3 && true) {
        fmt.Println(v4, v4)
    }
    var v98 int = 8
    v99 := []int{}
    for fi100 := 0; fi100 < 7; fi100++ {
        v101 := GenS6{F0: 1, F1: 2}
        _ = fi100
        _ = v101
    }
    switch ("\n\t" + "a") {
    case "\n\t":
        delete(v2, genFn1("hello", (v98 >> ((603) & 63))))
    case "foo bar":
        v4 -= (-(v4))
    case "a":
        switch (((("hello" + "") + "\n\t") + "") + "hello") {
        case "a":
            v98 = (v98 ^ 1)
        case "foo bar":
            switch "\n\t" {
            case "a":
                v4 = v4
            case "foo bar":
                if (v3 || (2 < 146)) {
                    var v102 uint = 4294967295
                    switch (v98 | (^(v98))) {
                    case 9223372036854775807:
                        var v103 int8 = -13
                        _ = v103
                    }
                    _ = v102
                } else {
                    v3 = v3
                }
            default:
                var v104 uint = 2
                _ = v104
            }
        default:
            fmt.Println(v98, v3, v4)
        }
    default:
        v98 = (v98 / (((v98 % (((v98 | v98)) | 1))) | 1))
    }
    switch (v98 + v98) {
    case -357:
        v4 = v4
    case 127:
        v99 = append(v99, 628)
    default:
        fmt.Println(v98)
    }
    var v105 bool = v3
    var v106 int16 = v4
    _ = len(v1)
    switch v98 {
    case 923:
        if (v105 || (1.5 <= 0.0)) {
            v3 = v105
        }
    case 766:
        for fi107 := 0; fi107 < 7; fi107++ {
            v108 := v1[(v98 | 32768)]
            var v109 int8 = 1
            _ = fi107
            _ = v108
            _ = v109
        }
    case 2147483647:
        for fi110 := 0; fi110 < 5; fi110++ {
            var v111 uint8 = 229
            v3 = (!true)
            v111 += (v111 % ((v111) | 1))
            _ = fi110
            _ = v111
        }
    }
    var v112 int64 = 9223372036854775807
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v105:", v105)
    fmt.Println("v106:", v106)
    fmt.Println("v112:", v112)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v98:", v98)
    fmt.Println("slicelen v99:", len(v99), v99)
    _ = v1
    _ = v105
    _ = v106
    _ = v112
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v98
    _ = v99
}
