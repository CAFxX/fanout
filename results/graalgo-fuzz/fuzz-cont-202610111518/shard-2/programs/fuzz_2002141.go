package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 uint32; F1 int16; F2 uint16 }
type GenS4 struct { F0 int64; F1 uint8; F2 int32 }
type GenS5 struct { F0 uint; F1 uint64 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 string }
type GenS8 struct { F0 string; F1 int8; F2 int32 }
type GenS9 struct { F0 bool; F1 int16; F2 int64 }
type GenS10 struct { F0 int32; F1 int32 }
type GenS11 struct { F0 int64; F1 uint8; F2 uint64 }
type GenS12 struct { F0 bool; F1 uint16; F2 uint }
type GenS13 struct { F0 string; F1 string; F2 int }
type GenS14 struct { F0 int8; F1 int }
type GenS15 struct { F0 int16; F1 int; F2 bool }
func genFn0(p0 int32) int64 {
    for fi4 := 0; fi4 < 14; fi4++ {
        for fi5 := 0; fi5 < 16; fi5++ {
            var v6 uint16 = 42
            _ = fi5
            _ = v6
        }
        v7 := GenS1{F0: 793, F1: 18446744073709551615}
        _ = fi4
        _ = v7
    }
    v8 := GenS2{F0: -1, F1: 0}
    var v9 int16 = 492
    if false {
        v9 += v9
        if (127 <= -1) {
            v9 = v9
            var v10 uint16 = 255
            for fi11 := 0; fi11 < 2; fi11++ {
                v12 := []int{(-(fi11)), fi11}
                p0 = (p0 / (((p0 >> ((18446744073709551615) & 63))) | 1))
                _ = fi11
                _ = v12
            }
            _ = v10
        }
    } else {
        p0 += (p0 << ((0) & 63))
    }
    var v13 string = (("foo bar" + "") + "hello")
    _ = p0
    _ = v13
    _ = v8
    _ = v9
    return 42
}
func genFn1() string {
    for fi14 := 0; fi14 < 3; fi14++ {
        if true {
            switch "hello" {
            case "a":
                v15 := []int{fi14}
                _ = v15
            case "\n\t":
                switch fi14 {
                case 128:
                    v16 := []int{fi14, fi14, -1, fi14}
                    _ = v16
                case -659:
                    for fi17 := 0; fi17 < 10; fi17++ {
                        if (false || false) {
                            var v18 uint64 = 42
                            v19 := []int{(fi14 & -784), (-(fi17))}
                            v20 := GenS3{F0: 4294967295, F1: 32767, F2: 42}
                            fmt.Println(fi17, fi14)
                            _ = v18
                            _ = v19
                            _ = v20
                        }
                        var v21 float32 = 0.0
                        var v22 int64 = 9223372036854775807
                        _ = fi17
                        _ = v21
                        _ = v22
                    }
                }
            }
        } else {
            var v23 uint64 = 0
            if (0 < v23) {
                v24 := GenS4{F0: 241, F1: 159, F2: int32(v23)}
                v25 := []int{fi14, (fi14 / (((fi14 ^ (fi14 * fi14))) | 1)), -1}
                _ = v24
                _ = v25
            }
            var v26 int16 = 42
            _ = v23
            _ = v26
        }
        v27 := []int{(fi14 - 398), 9223372036854775807, 32768, (fi14 - 973)}
        _ = fi14
        _ = v27
    }
    for fi28 := 0; fi28 < 17; fi28++ {
        for fi29 := 0; fi29 < 11; fi29++ {
            var v30 int32 = -1
            switch "foo bar" {
            case "a":
                switch "" {
                case "":
                    v30 -= v30
                case "\n\t":
                    v31 := GenS5{F0: 255, F1: 42}
                    _ = v31
                default:
                    var v32 bool = false
                    _ = v32
                }
            case "":
                v30 = 42
                switch (^(fi29)) {
                case -1:
                    switch "\n\t" {
                    case "hello":
                        if ((-2.25 - ((0.1 / 3.14159) * ((float32(1.5) / (float32((float64(0.0) - 1.5)) - 3.14159)) * (1.5 - 0.1)))) >= 1e10) {
                            if (-96 >= 0) {
                                v33 := GenS6{F0: ("foo bar" + "")}
                                _ = v33
                            } else {
                                var v34 int32 = v30
                                switch (fi29 ^ 127) {
                                case 32767:
                                    var v35 bool = false
                                    _ = v35
                                case -9223372036854775808:
                                    v34 += (v34 << ((4294967295) & 63))
                                case -585:
                                    var v36 int16 = 1
                                    _ = v36
                                default:
                                    fmt.Println(v34)
                                }
                                switch (fi29 * 32767) {
                                case 0:
                                    v37 := []int{(^(fi28)), (fi29 ^ (^(fi28)))}
                                    _ = v37
                                case 438:
                                    switch fi29 {
                                    case 8:
                                        v38 := []int{}
                                        _ = v38
                                    }
                                case 32768:
                                    v30 = v34
                                default:
                                    var v39 int8 = 127
                                    _ = v39
                                }
                                var v40 int32 = (v34 ^ (v30 << ((663) & 63)))
                                var v41 int64 = genFn0(312)
                                v42 := []int{(fi29 & (fi28 | -2147483648)), 426, (fi28 ^ fi29), 235}
                                var v43 uint = 65535
                                _ = v34
                                _ = v40
                                _ = v41
                                _ = v42
                                _ = v43
                            }
                            v44 := GenS3{F0: 0, F1: -366, F2: 2}
                            var v45 string = "foo bar"
                            if true {
                                v46 := GenS3{F0: 2, F1: -132, F2: 255}
                                _ = v46
                            }
                            _ = v44
                            _ = v45
                        } else {
                            if true {
                                var v47 uint16 = 32767
                                _ = v47
                            } else {
                                v48 := GenS7{F0: "a"}
                                _ = v48
                            }
                            v30 = 2147483647
                        }
                    case "a":
                        if (!true) {
                            if (!true) {
                                fmt.Println(fi28, v30)
                                v49 := GenS8{F0: ("a" + ("hello" + "hello")), F1: 0, F2: (-(v30))}
                                _ = v49
                            } else {
                                if false {
                                    v50 := []int{0, -1, fi28, fi28}
                                    _ = v50
                                } else {
                                    var v51 string = "a"
                                    v52 := []int{-9223372036854775808}
                                    _ = v51
                                    _ = v52
                                }
                                var v53 uint64 = 1
                                v53 += (^(v53))
                                _ = v53
                            }
                            v54 := GenS7{F0: "\n\t"}
                            var v55 uint = 0
                            v55 -= (v55 | v55)
                            _ = v54
                            _ = v55
                        } else {
                            v30 = (v30 + v30)
                            var v56 uint8 = 255
                            _ = v56
                        }
                    }
                case 1:
                    if true {
                        v57 := GenS7{F0: "foo bar"}
                        _ = v57
                    } else {
                        v58 := []int{}
                        _ = v58
                    }
                case 609:
                    var v59 uint8 = 2
                    _ = v59
                default:
                    v60 := GenS6{F0: "\n\t"}
                    _ = v60
                }
                switch "a" {
                case "hello":
                    v30 = (v30 + v30)
                case "":
                    var v61 int16 = -32768
                    _ = v61
                default:
                    var v62 uint16 = 0
                    _ = v62
                }
            case "hello":
                fmt.Println(v30, fi28, fi29)
            }
            _ = fi29
            _ = v30
        }
        var v63 string = ("foo bar" + ("foo bar" + "hello"))
        _ = fi28
        _ = v63
    }
    var v64 int32 = 1
    v65 := []int{-292, 32767, 1}
    switch ("foo bar" + "\n\t") {
    case "\n\t":
        fmt.Println(v64)
    case "a":
        var v66 int64 = 1
        _ = v66
    case "":
        v64 += (v64 >> ((65535) & 63))
    }
    _ = v64
    _ = v65
    return "foo bar"
}

func main() {
    var v1 int = 0
    v2 := make(map[uint32]bool)
    var v3 bool = (0 >= 644)
    v67 := []int{(v1 & -320), v1}
    var v68 int64 = 42
    v69 := v2[862]
    fmt.Println(v69, v1, v68)
    for fi70 := 0; fi70 < 15; fi70++ {
        for fi71 := 0; fi71 < 16; fi71++ {
            fmt.Println(v69, fi71, v3)
            _ = fi71
        }
        if (((1 > 66) && v3) || (0 < 0)) {
            v72 := GenS1{F0: v68, F1: 42}
            var v73 int32 = 453
            switch "foo bar" {
            case "hello":
                var v74 string = ("\n\t" + "\n\t")
                _ = v74
            case "\n\t":
                var v75 float32 = 3.14159
                _ = v75
            case "":
                for fi76 := 0; fi76 < 14; fi76++ {
                    v69 = (!(((0.0 * (((3.14159 - (0.1 - 1.5)) - (3.14159 + (1.5 * -2.25))) / 1.0)) < 0.0) && (v69 || v3)))
                    if len(v67) > 0 {
                        v67[0] = (fi70 ^ (fi70 ^ fi70))
                    }
                    _ = fi76
                }
            default:
                if v69 {
                    v77 := GenS9{F0: v3, F1: 0, F2: (v68 + 0)}
                    v3 = (2147483647 < 510)
                    v73 -= v73
                    _ = v77
                } else {
                    for fi78 := 0; fi78 < 18; fi78++ {
                        if true {
                            v79 := GenS10{F0: 1, F1: v73}
                            switch ("" + (("hello" + ((("" + ("foo bar" + "hello")) + "\n\t") + "a")) + "hello")) {
                            case "foo bar":
                                if v3 {
                                    v80 := GenS11{F0: v68, F1: 255, F2: 18446744073709551615}
                                    _ = v80
                                } else {
                                    var v81 float32 = -2.25
                                    var v82 string = (("\n\t" + "foo bar") + "hello")
                                    if ((v81 <= (3.14159 / float32(float64(((float32(1e10) - ((float32(0.1) * ((v81 / 1.0) + float32(((0.1 + float64(3.14159)) / (1e10 * ((float64(v81) + 0.1) / (0.0 - 1e10))))))) + v81)) - v81))))) && (42 < 1)) {
                                        _ = len(v67)
                                        var v83 int16 = -422
                                        switch (v1 - 428) {
                                        case -9223372036854775808:
                                            delete(v2, 2147483648)
                                        default:
                                            var v84 int8 = 100
                                            _ = v84
                                        }
                                        _ = v83
                                    }
                                    _ = v81
                                    _ = v82
                                }
                            case "hello":
                                if v69 {
                                    fmt.Println(v73)
                                    if ((3.14159 != 1.5) && (0 > 21)) {
                                        v73 -= -215
                                        _ = len(v2)
                                        v1 = (fi70 >> ((752) & 63))
                                    } else {
                                        var v85 float64 = ((-2.25 * -2.25) * (1e10 / (1e10 / 1e10)))
                                        var v86 float32 = (0.1 - float32(v85))
                                        v87 := GenS12{F0: (true && ((747 > 0) && (0 <= 135))), F1: 65535, F2: 65535}
                                        var v88 uint64 = 0
                                        var v89 string = "hello"
                                        _ = v85
                                        _ = v86
                                        _ = v87
                                        _ = v88
                                        _ = v89
                                    }
                                }
                            default:
                                v90 := GenS8{F0: ((("\n\t" + "hello") + "\n\t") + "foo bar"), F1: 42, F2: 2147483647}
                                _ = v90
                            }
                            _ = v79
                        } else {
                            var v91 uint64 = 18446744073709551615
                            _ = v91
                        }
                        fmt.Println(v69, v1)
                        if (-771 <= -1) {
                            v92 := GenS13{F0: "\n\t", F1: "foo bar", F2: v1}
                            _ = len(v2)
                            if (!v69) {
                                switch (("\n\t" + ((((("a" + "\n\t") + "\n\t") + "\n\t") + ("" + ("a" + "\n\t"))) + "foo bar")) + ("a" + "a")) {
                                case "a":
                                    v2[1] = (1.5 > 3.14159)
                                default:
                                    if false {
                                        if v3 {
                                            fmt.Println(v1, v69)
                                        }
                                        var v93 bool = v3
                                        _ = v93
                                    } else {
                                        var v94 int32 = v73
                                        if (!false) {
                                            v67 = append(v67, fi78)
                                            fmt.Println(v1)
                                            var v95 uint8 = 120
                                            _ = v95
                                        }
                                        _ = v94
                                    }
                                }
                                switch "" {
                                case "":
                                    var v96 int64 = (v68 ^ int64(fi78))
                                    _ = v96
                                case "hello":
                                    switch fi78 {
                                    case 345:
                                        if (v1 < fi78) {
                                            v73 -= (^(v73))
                                            if ((!v69) || v3) {
                                                v1 = 0
                                            }
                                            v3 = v3
                                        }
                                    }
                                case "foo bar":
                                    var v97 uint = 859
                                    _ = v97
                                default:
                                    var v98 bool = (v73 != v73)
                                    _ = v98
                                }
                            }
                            _ = v92
                        } else {
                            var v99 float64 = (3.14159 / (-2.25 - 0.0))
                            _ = v99
                        }
                        _ = fi78
                    }
                }
            }
            _ = v72
            _ = v73
        }
        _ = fi70
    }
    var v100 int64 = genFn0(-725)
    var v101 string = "a"
    v68 = (-(v68))
    var v102 int16 = 0
    if (v3 || ((0.1 < 0.0) && v3)) {
        if false {
            var v103 int = v1
            var v104 int32 = int32(v68)
            _ = v103
            _ = v104
        }
        v102 = 975
    }
    _ = len(v2)
    var v105 uint = 652
    v106 := GenS5{F0: v105, F1: 0}
    var v107 int = 4294967295
    if len(v67) >= 2 {
        v67 = v67[1:len(v67)]
    }
    var v108 int32 = -2147483648
    for fi109 := 0; fi109 < 9; fi109++ {
        var v110 uint = 2
        v68 -= 1
        switch "a" {
        case "hello":
            var v111 int64 = (v100 - -9223372036854775808)
            _ = v111
        case "foo bar":
            v112 := GenS2{F0: v102, F1: 0}
            _ = v112
        default:
            switch ((((v101 + (((((v101 + ("foo bar" + v101)) + (v101 + "a")) + v101) + (v101 + v101)) + ("" + v101))) + ((v101 + "\n\t") + v101)) + v101) + ("hello" + "foo bar")) {
            case "hello":
                v113 := GenS14{F0: 127, F1: (-(v107))}
                _ = v113
            default:
                for fi114 := 0; fi114 < 2; fi114++ {
                    v115 := GenS15{F0: v102, F1: v107, F2: v69}
                    var v116 uint64 = 0
                    v68 -= 224
                    if false {
                        v100 = v68
                    }
                    var v117 bool = (2147483647 == 906)
                    switch 79 {
                    case 42:
                        v100 -= -1
                    case -9223372036854775808:
                        _ = len(v2)
                    }
                    var v118 bool = (v69 && (v3 || ((v101 < v101) && true)))
                    _ = fi114
                    _ = v115
                    _ = v116
                    _ = v117
                    _ = v118
                }
            }
        }
        _ = fi109
        _ = v110
    }
    fmt.Println(v69, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v100:", v100)
    fmt.Println("v101:", v101)
    fmt.Println("v102:", v102)
    fmt.Println("v105:", v105)
    fmt.Println("v107:", v107)
    fmt.Println("v108:", v108)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v67:", len(v67), v67)
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v100
    _ = v101
    _ = v102
    _ = v105
    _ = v106
    _ = v107
    _ = v108
    _ = v2
    _ = v3
    _ = v67
    _ = v68
    _ = v69
}
