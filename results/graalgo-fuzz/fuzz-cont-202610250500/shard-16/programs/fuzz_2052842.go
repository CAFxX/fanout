package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 int64 }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 uint8; F1 uint8; F2 uint16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int64; F1 int32; F2 string }
type GenS7 struct { F0 string; F1 int32 }
type GenS8 struct { F0 string; F1 uint8; F2 bool }
type GenS9 struct { F0 int32; F1 bool }
type GenS10 struct { F0 uint8; F1 uint }
type GenS11 struct { F0 uint; F1 uint32 }
type GenS12 struct { F0 int; F1 int; F2 int64 }
type GenS13 struct { F0 int; F1 uint32; F2 uint16 }
type GenS14 struct { F0 string; F1 int; F2 uint64 }
type GenS15 struct { F0 int32; F1 string; F2 uint16 }
type GenS16 struct { F0 uint; F1 int8 }
type GenS17 struct { F0 int; F1 string; F2 uint32 }

func main() {
    v1 := []uint32{}
    var v2 uint32 = 0
    v3 := []int{128}
    v4 := GenS1{F0: 375, F1: 32767, F2: 9223372036854775807}
    fmt.Println(v2)
    var v5 uint8 = 42
    for fi6 := 0; fi6 < 6; fi6++ {
        var v7 int64 = -9223372036854775808
        switch fi6 {
        case -9223372036854775808:
            if (!("a" == "a")) {
                var v8 uint8 = (v5 << ((2) & 63))
                fmt.Println(fi6)
                if ((v2 % ((v2) | 1)) != v2) {
                    v9 := GenS1{F0: 255, F1: -260, F2: v7}
                    var v10 string = ("hello" + "a")
                    switch v10 {
                    case "\n\t":
                        var v11 float64 = float64(1.5)
                        _ = v11
                    case "foo bar":
                        var v12 int8 = -128
                        _ = v12
                    }
                    _ = v10
                    _ = v9
                }
                _ = v8
            }
        default:
            var v13 int = (^(fi6))
            _ = v13
        }
        _ = fi6
        _ = v7
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v5 += v5
    switch ("hello" + "foo bar") {
    case "a":
        var v14 uint = 255
        _ = v14
    case "hello":
        v15 := GenS1{F0: 552, F1: -32768, F2: 0}
        _ = v15
    default:
        switch ((("foo bar" + "a") + "") + "hello") {
        case "a":
            v16 := GenS1{F0: 257, F1: -384, F2: -1}
            _ = v16
        default:
            v5 = 158
        }
    }
    if ((65535 < 1) && (false || (127 >= 1))) {
        switch "foo bar" {
        case "\n\t":
            v17 := GenS1{F0: 18446744073709551615, F1: int16(v5), F2: 660}
            _ = v17
        case "a":
            for fi18 := 0; fi18 < 1; fi18++ {
                fmt.Println(v5)
                for fi19 := 0; fi19 < 16; fi19++ {
                    if ((v5 ^ (v5 + 0)) < 255) {
                        switch "foo bar" {
                        case "":
                            switch fi18 {
                            case -541:
                                switch "hello" {
                                case "a":
                                    if (693 <= 4294967295) {
                                        v20 := GenS1{F0: 4294967295, F1: -1, F2: 42}
                                        var v21 uint32 = v2
                                        _ = v20
                                        _ = v21
                                    } else {
                                        fmt.Println(fi19)
                                        v5 += (^(v5))
                                        v2 += 1
                                    }
                                default:
                                    if (9223372036854775807 >= 9223372036854775807) {
                                        if false {
                                            var v22 float64 = float64(float32(1.5))
                                            v5 = (v5 & 255)
                                            _ = v22
                                        } else {
                                            if true {
                                                var v23 bool = (-425 != -1)
                                                if (!false) {
                                                    var v24 int32 = -1
                                                    v25 := GenS2{F0: 472, F1: v2}
                                                    _ = v24
                                                    _ = v25
                                                } else {
                                                    var v26 int32 = 0
                                                    var v27 bool = v23
                                                    _ = v26
                                                    _ = v27
                                                }
                                                if v23 {
                                                    v2 = (v2 ^ v2)
                                                }
                                                fmt.Println(fi19, v2, v2)
                                                if v23 {
                                                    v5 = v5
                                                }
                                                switch 42 {
                                                case -860:
                                                    v5 = 1
                                                case 4294967295:
                                                    var v28 uint8 = (v5 << ((248) & 63))
                                                    _ = v28
                                                }
                                                if (254 != (v2 >> ((18446744073709551615) & 63))) {
                                                    v29 := GenS2{F0: 0, F1: (v2 & v2)}
                                                    v5 = v5
                                                    _ = v29
                                                } else {
                                                    var v30 int = (-(fi18))
                                                    var v31 int64 = 273
                                                    var v32 uint = 1
                                                    v32 = (v32 >> ((v32) & 63))
                                                    var v33 int = 128
                                                    _ = v30
                                                    _ = v31
                                                    _ = v32
                                                    _ = v33
                                                }
                                                _ = v23
                                            } else {
                                                v2 = (v2 << ((18446744073709551615) & 63))
                                                var v34 int16 = 876
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                _ = v34
                                            }
                                            var v35 int8 = 127
                                            _ = v35
                                        }
                                        if (-1 < 991) {
                                            switch "a" {
                                            case "hello":
                                                _ = len(v1)
                                                fmt.Println(fi18)
                                                var v36 int8 = 127
                                                v37 := GenS3{F0: 128, F1: v5, F2: 65535}
                                                switch (fi18 ^ -13) {
                                                case 4294967295:
                                                    v38 := GenS1{F0: 126, F1: 32767, F2: 187}
                                                    _ = v38
                                                default:
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                }
                                                v39 := GenS3{F0: v5, F1: 127, F2: 2}
                                                var v40 float32 = 1e10
                                                _ = v36
                                                _ = v37
                                                _ = v39
                                                _ = v40
                                            default:
                                                if len(v1) > 0 {
                                                    v1[0] = 74
                                                }
                                            }
                                            switch ("a" + ("\n\t" + ("\n\t" + (("a" + ("foo bar" + (("\n\t" + "foo bar") + ("hello" + "\n\t")))) + "hello")))) {
                                            case "foo bar":
                                                fmt.Println(fi19, v2)
                                                var v41 int64 = 1
                                                v42 := GenS4{F0: 42}
                                                _ = v41
                                                _ = v42
                                            default:
                                                if true {
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                }
                                            }
                                        } else {
                                            v5 -= v5
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        }
                                    } else {
                                        var v43 bool = true
                                        switch "foo bar" {
                                        case "\n\t":
                                            fmt.Println(v2)
                                        case "foo bar":
                                            v43 = v43
                                        case "":
                                            v5 -= (v5 ^ (v5 >> ((20) & 63)))
                                        }
                                        _ = v43
                                    }
                                }
                                var v44 uint64 = 42
                                _ = v44
                            case 32768:
                                v45 := GenS5{F0: 556}
                                _ = v45
                            case 235:
                                switch (("a" + (("a" + ("hello" + (("hello" + ("foo bar" + "a")) + (("\n\t" + "a") + ("\n\t" + ("" + (("foo bar" + ("\n\t" + "foo bar")) + ((((("a" + "foo bar") + "foo bar") + ((("a" + "") + "a") + "")) + "a") + (("hello" + (("foo bar" + "a") + ("foo bar" + ""))) + ("foo bar" + "a")))))))))) + "\n\t")) + ("a" + (("\n\t" + ("a" + (("\n\t" + "a") + "\n\t"))) + "foo bar"))) {
                                case "foo bar":
                                    var v46 float64 = 3.14159
                                    _ = v46
                                case "":
                                    if (!false) {
                                        v47 := GenS6{F0: 849, F1: 2147483647, F2: "a"}
                                        v48 := GenS7{F0: ("a" + "foo bar"), F1: -228}
                                        var v49 uint64 = 2
                                        _ = v47
                                        _ = v48
                                        _ = v49
                                    }
                                case "\n\t":
                                    if true {
                                        var v50 uint32 = (^(v2))
                                        if (0 != 255) {
                                            fmt.Println(fi19, v2)
                                            switch fi19 {
                                            case 4294967295:
                                                var v51 int32 = -538
                                                _ = v51
                                            case 352:
                                                v5 = v5
                                            case 878:
                                                var v52 int16 = -506
                                                var v53 bool = (int8(v2) == 42)
                                                v54 := GenS2{F0: 647, F1: (^(v2))}
                                                var v55 uint8 = 165
                                                v53 = (false && v53)
                                                _ = v52
                                                _ = v53
                                                _ = v54
                                                _ = v55
                                            default:
                                                fmt.Println(fi19)
                                            }
                                            if false {
                                                switch -214 {
                                                case 364:
                                                    v2 += 365
                                                case 128:
                                                    if true {
                                                        v56 := GenS6{F0: 42, F1: -1, F2: "a"}
                                                        _ = v56
                                                    }
                                                case 727:
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                default:
                                                    fmt.Println(v2, fi19)
                                                }
                                            } else {
                                                v2 -= (v50 - (v2 + 0))
                                            }
                                        }
                                        v57 := GenS2{F0: -789, F1: 0}
                                        _ = v50
                                        _ = v57
                                    } else {
                                        if (9223372036854775807 <= 9223372036854775807) {
                                            var v58 int = fi19
                                            _ = v58
                                        } else {
                                            v5 -= (v5 - v5)
                                            v5 = (v5 / (((v5 * (v5 / (((v5 >> ((950) & 63))) | 1)))) | 1))
                                        }
                                        v2 = uint32(v5)
                                        var v59 uint16 = 65535
                                        if (-32768 != int16(v5)) {
                                            var v60 int8 = 0
                                            _ = v60
                                        } else {
                                            if false {
                                                switch 1 {
                                                case -52:
                                                    v2 += (v2 >> ((uint(v5)) & 63))
                                                }
                                            } else {
                                                if true {
                                                    v61 := GenS8{F0: ("hello" + ((("" + "hello") + ("\n\t" + "hello")) + ("\n\t" + "\n\t"))), F1: uint8(v59), F2: false}
                                                    _ = v61
                                                }
                                            }
                                            var v62 uint64 = 42
                                            v59 = 65535
                                            _ = v62
                                        }
                                        fmt.Println(fi19, v5)
                                        _ = v59
                                    }
                                }
                            default:
                                fmt.Println(fi19, v5, fi19)
                            }
                        case "a":
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        case "\n\t":
                            v63 := GenS6{F0: -9223372036854775808, F1: -2147483648, F2: "foo bar"}
                            _ = v63
                        default:
                            if false {
                                if (576 >= -32768) {
                                    switch fi19 {
                                    case 2147483647:
                                        if true {
                                            v64 := GenS4{F0: int64(fi18)}
                                            _ = v64
                                        } else {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            v65 := GenS2{F0: 1, F1: v2}
                                            var v66 int = (fi19 % ((fi18) | 1))
                                            switch "" {
                                            case "hello":
                                                fmt.Println(v2, v66)
                                            }
                                            _ = v65
                                            _ = v66
                                        }
                                    case 574:
                                        var v67 int = int(v2)
                                        _ = v67
                                    }
                                    v2 += (^(v2))
                                } else {
                                    v3 = append(v3, (fi18 - fi18))
                                }
                                switch "hello" {
                                case "\n\t":
                                    var v68 float64 = 3.14159
                                    _ = v68
                                case "hello":
                                    v69 := GenS9{F0: -1, F1: true}
                                    v5 += (v5 % ((v5) | 1))
                                    if true {
                                        v3 = append(v3, (^(fi18)))
                                        v2 -= (v2 + 2)
                                        if false {
                                            v2 -= (^(v2))
                                            if false {
                                                fmt.Println(fi18, v5, v2)
                                            } else {
                                                switch ("a" + ("" + "hello")) {
                                                case "foo bar":
                                                    switch 565 {
                                                    case -2147483648:
                                                        var v70 bool = ((("\n\t" + "hello") + ((("\n\t" + "a") + ("foo bar" + "\n\t")) + "a")) <= "foo bar")
                                                        _ = v70
                                                    case 992:
                                                        v2 -= v2
                                                    case 341:
                                                        fmt.Println(v5, fi19, fi19)
                                                    }
                                                    v2 = v2
                                                }
                                                fmt.Println(fi19, v2)
                                            }
                                            v5 = v5
                                        } else {
                                            if (18446744073709551615 > 0) {
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            } else {
                                                switch "foo bar" {
                                                case "\n\t":
                                                    if true {
                                                        v71 := GenS9{F0: -468, F1: false}
                                                        _ = v71
                                                    } else {
                                                        v2 = (v2 + v2)
                                                        var v72 uint = 436
                                                        _ = v72
                                                    }
                                                default:
                                                    var v73 float64 = 0.0
                                                    _ = v73
                                                }
                                                var v74 uint16 = 32767
                                                if (!true) {
                                                    v2 -= v2
                                                    fmt.Println(v74, v2)
                                                    var v75 int16 = 0
                                                    _ = v75
                                                } else {
                                                    var v76 float64 = ((-2.25 * ((((-2.25 / 3.14159) - -2.25) + -2.25) / 0.1)) + ((-2.25 - 1e10) + (0.0 * float64(1.5))))
                                                    _ = v76
                                                }
                                                if (100 <= -1) {
                                                    if true {
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        if false {
                                                            fmt.Println(fi19, v2, v5)
                                                        }
                                                        _ = len(v3)
                                                    } else {
                                                        var v77 int64 = -9223372036854775808
                                                        fmt.Println(fi18, v74, v77)
                                                        _ = v77
                                                    }
                                                    v2 -= v2
                                                    v5 -= (v5 - 255)
                                                } else {
                                                    v78 := GenS10{F0: (v5 % ((v5) | 1)), F1: 65535}
                                                    var v79 uint64 = uint64(v74)
                                                    _ = v78
                                                    _ = v79
                                                }
                                                _ = v74
                                            }
                                        }
                                    } else {
                                        if false {
                                            var v80 uint32 = 200
                                            var v81 float32 = 0.1
                                            if ((!(true && (v80 <= v80))) && (true || (!(uint(v5) >= 2)))) {
                                                v81 -= v81
                                                v80 = (v80 & 4294967295)
                                            } else {
                                                fmt.Println(v81)
                                                v82 := GenS1{F0: 2, F1: 332, F2: -9223372036854775808}
                                                _ = v82
                                            }
                                            _ = v80
                                            _ = v81
                                        } else {
                                            fmt.Println(fi18, v5, fi18)
                                            switch "a" {
                                            case "a":
                                                var v83 int16 = -986
                                                _ = v83
                                            case "\n\t":
                                                v84 := GenS11{F0: 65535, F1: 594}
                                                _ = v84
                                            case "hello":
                                                if false {
                                                    v3 = append(v3, 32767)
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    var v85 int32 = -1
                                                    _ = v85
                                                }
                                            default:
                                                if (true || false) {
                                                    v2 += v2
                                                    var v86 float64 = (-2.25 - 1e10)
                                                    v2 += (v2 + v2)
                                                    _ = v86
                                                } else {
                                                    var v87 bool = true
                                                    v88 := GenS12{F0: -2147483648, F1: fi19, F2: 515}
                                                    var v89 float64 = 1e10
                                                    fmt.Println(fi18, v87)
                                                    _ = v87
                                                    _ = v88
                                                    _ = v89
                                                }
                                            }
                                            switch -961 {
                                            case 42:
                                                fmt.Println(v5, v5, fi18)
                                            case 330:
                                                var v90 int = 127
                                                _ = v90
                                            default:
                                                v91 := GenS13{F0: fi18, F1: (v2 ^ 772), F2: 255}
                                                _ = v91
                                            }
                                            v2 += (v2 + v2)
                                        }
                                        v5 = (v5 - (v5 * (v5 & v5)))
                                        v2 -= 2
                                    }
                                    _ = v69
                                default:
                                    var v92 float64 = ((0.0 / 3.14159) + ((1e10 * (0.1 * ((3.14159 * float64(3.14159)) - -2.25))) / (0.0 + 3.14159)))
                                    _ = v92
                                }
                                fmt.Println(v2)
                                if (403 > 0) {
                                    if true {
                                        var v93 float64 = -2.25
                                        var v94 int64 = -1
                                        _ = v93
                                        _ = v94
                                    }
                                    switch "a" {
                                    case "":
                                        switch "\n\t" {
                                        case "":
                                            v5 -= 127
                                        default:
                                            fmt.Println(fi18, v5)
                                        }
                                    }
                                } else {
                                    if false {
                                        var v95 int8 = 127
                                        _ = v95
                                    } else {
                                        var v96 uint8 = 114
                                        _ = len(v1)
                                        _ = v96
                                    }
                                    fmt.Println(fi19)
                                    v97 := GenS14{F0: ("foo bar" + ("hello" + "a")), F1: (fi19 << ((65535) & 63)), F2: 4294967295}
                                    _ = v97
                                }
                            } else {
                                var v98 float64 = (0.0 + -2.25)
                                if false {
                                    _ = len(v3)
                                }
                                var v99 uint64 = 42
                                _ = v98
                                _ = v99
                            }
                        }
                        var v100 uint16 = 487
                        _ = v100
                    } else {
                        v5 = (v5 * 1)
                        switch "\n\t" {
                        case "\n\t":
                            v5 = 0
                            var v101 uint8 = v5
                            _ = v101
                        }
                        if false {
                            fmt.Println(v5, fi19)
                            fmt.Println(fi18, fi19, v5)
                        } else {
                            v2 -= v2
                        }
                        var v102 int = (fi19 << ((736) & 63))
                        if false {
                            v102 = fi18
                            fmt.Println(v2, v5)
                        }
                        _ = v102
                    }
                    var v103 int = 9223372036854775807
                    v104 := GenS15{F0: 385, F1: "foo bar", F2: 855}
                    _ = fi19
                    _ = v103
                    _ = v104
                }
                if true {
                    if ((float32((1e10 + -2.25)) - float32((1.5 - 0.0))) >= float32(1e10)) {
                        switch fi18 {
                        case 871:
                            _ = len(v3)
                        default:
                            v5 = v5
                        }
                    }
                }
                _ = fi18
            }
        default:
            for fi105 := 0; fi105 < 10; fi105++ {
                if true {
                    if true {
                        v106 := GenS16{F0: 255, F1: 100}
                        for fi107 := 0; fi107 < 12; fi107++ {
                            var v108 bool = ((fi105 % ((fi105) | 1)) <= -9223372036854775808)
                            var v109 bool = v108
                            var v110 int64 = 42
                            _ = fi107
                            _ = v108
                            _ = v109
                            _ = v110
                        }
                        _ = v106
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    var v111 int32 = 42
                    _ = v111
                } else {
                    for fi112 := 0; fi112 < 18; fi112++ {
                        var v113 bool = true
                        _ = fi112
                        _ = v113
                    }
                    v3 = append(v3, (fi105 * fi105))
                    for fi114 := 0; fi114 < 20; fi114++ {
                        v115 := GenS17{F0: fi105, F1: "\n\t", F2: v2}
                        _ = fi114
                        _ = v115
                    }
                }
                fmt.Println(v5)
                v2 = (v2 ^ v2)
                _ = fi105
            }
        }
        var v116 int32 = -1
        _ = v116
    } else {
        _ = len(v3)
    }
    v2 = (^(v2))
    var v117 int32 = -851
    var v118 uint64 = 220
    var v119 uint64 = (^(v118))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v117:", v117)
    fmt.Println("v118:", v118)
    fmt.Println("v119:", v119)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v117
    _ = v118
    _ = v119
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
