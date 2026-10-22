package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 uint16 }
type GenS2 struct { F0 uint16; F1 uint; F2 int8 }
type GenS3 struct { F0 int; F1 bool }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 bool; F1 int32; F2 bool }
type GenS6 struct { F0 uint64 }
func genFn0() uint {
    if true {
        if false {
            v5 := []int{-261, -447}
            var v6 float32 = 0.0
            var v7 float64 = 0.0
            _ = v5
            _ = v6
            _ = v7
        }
        var v8 float64 = ((3.14159 + 1e10) / 3.14159)
        var v9 int16 = 1
        _ = v8
        _ = v9
    } else {
        var v10 string = ""
        for fi11 := 0; fi11 < 3; fi11++ {
            fmt.Println(v10)
            for fi12 := 0; fi12 < 18; fi12++ {
                v13 := []int{fi12, 127, -870}
                fmt.Println(fi11, fi11, v10)
                _ = fi12
                _ = v13
            }
            v14 := []int{fi11, fi11, fi11, (fi11 ^ fi11)}
            _ = fi11
            _ = v14
        }
        var v15 uint64 = 0
        _ = v10
        _ = v15
    }
    var v16 int64 = 9223372036854775807
    for fi17 := 0; fi17 < 19; fi17++ {
        if ((!false) && (2147483648 <= 0)) {
            var v18 uint8 = 127
            _ = v18
        }
        _ = fi17
    }
    _ = v16
    return 930
}
func genFn1(p0 uint16, p1 int8, p2 int8) uint32 {
    v19 := GenS1{F0: ("hello" + "hello"), F1: 0, F2: (p0 ^ 610)}
    var v20 float32 = -2.25
    switch -9223372036854775808 {
    case 405:
        var v21 uint8 = 2
        _ = v21
    default:
        v22 := []int{2147483647, -734, 127, 683}
        _ = v22
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    _ = v20
    return 65535
}

func main() {
    var v1 bool = (-776 <= 0)
    var v2 uint32 = 379
    v3 := make(map[int]string)
    var v4 string = ("\n\t" + ("" + "a"))
    if ((^(v2)) < v2) {
        v2 -= v2
        v23 := GenS1{F0: v4, F1: -9223372036854775808, F2: 911}
        _ = v23
    }
    switch 2147483647 {
    case 2147483647:
        fmt.Println(v2)
    }
    for fi24 := 0; fi24 < 14; fi24++ {
        if (128 < (-(fi24))) {
            var v25 bool = (2147483647 <= 42)
            _ = v25
        } else {
            var v26 uint = 2
            if v1 {
                if true {
                    fmt.Println(v4, v4)
                    v27 := GenS1{F0: ("hello" + v4), F1: 42, F2: 32767}
                    delete(v3, fi24)
                    _ = v27
                } else {
                    if false {
                        if (850 >= v26) {
                            var v28 uint8 = 42
                            for fi29 := 0; fi29 < 5; fi29++ {
                                if v1 {
                                    var v30 uint8 = (v28 >> (((v26 % ((genFn0()) | 1))) & 63))
                                    v31 := GenS1{F0: v4, F1: 42, F2: 32767}
                                    v2 = (v2 / ((genFn1(uint16(v26), 127, -84)) | 1))
                                    _ = v30
                                    _ = v31
                                }
                                fmt.Println(v26)
                                v32 := GenS2{F0: 32767, F1: (v26 * (^(v26))), F2: -99}
                                _ = fi29
                                _ = v32
                            }
                            if (1e10 != (0.1 * -2.25)) {
                                var v33 float64 = 0.1
                                var v34 int = fi24
                                _ = v33
                                _ = v34
                            } else {
                                if false {
                                    for fi35 := 0; fi35 < 2; fi35++ {
                                        fmt.Println(v26, v28)
                                        _ = fi35
                                    }
                                    v36 := []int{}
                                    v4 = v4
                                    _ = v36
                                } else {
                                    v37 := GenS3{F0: 128, F1: false}
                                    switch v4 {
                                    case "\n\t":
                                        v38 := []int{(^(fi24)), fi24, fi24}
                                        _ = v38
                                    case "a":
                                        if v1 {
                                            if false {
                                                fmt.Println(fi24)
                                                _ = len(v3)
                                            } else {
                                                v1 = v1
                                            }
                                            _ = len(v3)
                                        }
                                    case "":
                                        v28 = (v28 * v28)
                                    default:
                                        v26 += 598
                                    }
                                    if ((!(!(!false))) || v1) {
                                        v39 := []int{}
                                        if len(v39) > 0 {
                                            v39[0] = fi24
                                        }
                                        v40 := GenS1{F0: (("\n\t" + "") + (v4 + (v4 + "foo bar"))), F1: 1, F2: 682}
                                        _ = v39
                                        _ = v40
                                    } else {
                                        fmt.Println(v2, fi24, v1)
                                    }
                                    _ = v37
                                }
                                if ((1.5 + 0.1) == float32(1.5)) {
                                    v26 = v26
                                    for fi41 := 0; fi41 < 10; fi41++ {
                                        var v42 float32 = 1.5
                                        v1 = v1
                                        _ = fi41
                                        _ = v42
                                    }
                                    for fi43 := 0; fi43 < 12; fi43++ {
                                        var v44 uint8 = (v28 >> ((genFn0()) & 63))
                                        v45 := GenS1{F0: v4, F1: -26, F2: 32767}
                                        var v46 float32 = (float32(0.0) + float32(((0.1 * 1e10) - 0.1)))
                                        v47 := []int{fi24, (^(fi43))}
                                        _ = fi43
                                        _ = v44
                                        _ = v45
                                        _ = v46
                                        _ = v47
                                    }
                                }
                            }
                            _ = v28
                        } else {
                            fmt.Println(v2, v26)
                            for fi48 := 0; fi48 < 14; fi48++ {
                                v26 = genFn0()
                                v49 := GenS1{F0: ("hello" + ""), F1: 42, F2: 65535}
                                var v50 int16 = int16(v26)
                                _ = fi48
                                _ = v49
                                _ = v50
                            }
                            v4 = ""
                        }
                        if v1 {
                            var v51 uint = v26
                            for fi52 := 0; fi52 < 4; fi52++ {
                                var v53 uint = v26
                                v1 = (1e10 != 1e10)
                                _ = fi52
                                _ = v53
                            }
                            _ = v51
                        } else {
                            v4 = ("foo bar" + "")
                            switch fi24 {
                            case 32767:
                                for fi54 := 0; fi54 < 2; fi54++ {
                                    var v55 bool = (!(1.5 == ((-2.25 + 0.0) / 3.14159)))
                                    _ = fi54
                                    _ = v55
                                }
                            }
                            for fi56 := 0; fi56 < 20; fi56++ {
                                switch (fi56 * (fi56 >> ((v26) & 63))) {
                                case 32767:
                                    v2 -= 2147483647
                                }
                                _ = fi56
                            }
                        }
                        v26 -= (v26 & 42)
                    }
                    switch v4 {
                    case "\n\t":
                        var v57 int32 = 810
                        _ = v57
                    case "hello":
                        v2 = v2
                    case "a":
                        if ((2147483647 < 42) && true) {
                            if v1 {
                                var v58 int8 = 0
                                var v59 int = fi24
                                _ = v58
                                _ = v59
                            }
                        } else {
                            v60 := []int{fi24, (fi24 ^ 42), fi24}
                            if (fi24 < 32768) {
                                if (!v1) {
                                    v2 = (^(v2))
                                    var v61 uint = (v26 + (v26 * genFn0()))
                                    v62 := GenS2{F0: 866, F1: (v26 / (((v26 * (^(v61)))) | 1)), F2: int8(v61)}
                                    _ = v61
                                    _ = v62
                                } else {
                                    for fi63 := 0; fi63 < 10; fi63++ {
                                        if len(v60) > 0 {
                                            v60[0] = fi63
                                        }
                                        v2 += v2
                                        switch (v4 + v4) {
                                        case "\n\t":
                                            if (4294967295 < (fi63 | fi24)) {
                                                v1 = (2 == 1)
                                                var v64 float32 = (0.1 + (0.1 / 0.1))
                                                var v65 uint16 = 255
                                                v66 := GenS2{F0: v65, F1: 18446744073709551615, F2: -128}
                                                _ = v64
                                                _ = v65
                                                _ = v66
                                            }
                                        case "a":
                                            v1 = (!(!true))
                                        case "foo bar":
                                            if (v1 && true) {
                                                var v67 uint8 = 56
                                                if v1 {
                                                    v3[fi24] = (("hello" + "\n\t") + ("" + ("\n\t" + "a")))
                                                    v68 := GenS3{F0: 4294967295, F1: (!v1)}
                                                    _ = v68
                                                }
                                                _ = v67
                                            } else {
                                                if v1 {
                                                    v3[fi24] = (v4 + "hello")
                                                    v1 = (153 < 0)
                                                } else {
                                                    fmt.Println(v2, fi24)
                                                    var v69 bool = (-9223372036854775808 == 42)
                                                    v1 = v69
                                                    _ = v69
                                                }
                                            }
                                        default:
                                            v2 = v2
                                        }
                                        _ = fi63
                                    }
                                    if v1 {
                                        if (156 == 42) {
                                            if (fi24 >= int(v2)) {
                                                var v70 int64 = 0
                                                var v71 uint = v26
                                                v72 := v3[127]
                                                _ = v70
                                                _ = v71
                                                _ = v72
                                            } else {
                                                if false {
                                                    v2 -= v2
                                                }
                                            }
                                        } else {
                                            var v73 uint64 = 0
                                            v2 = (v2 % ((genFn1(956, 127, 60)) | 1))
                                            _ = v73
                                        }
                                    }
                                }
                            }
                            _ = v60
                        }
                    default:
                        switch fi24 {
                        case -137:
                            if (1e10 == 1e10) {
                                for fi74 := 0; fi74 < 19; fi74++ {
                                    var v75 float64 = float64(-2.25)
                                    fmt.Println(fi74)
                                    var v76 int64 = -717
                                    _ = fi74
                                    _ = v75
                                    _ = v76
                                }
                            }
                        case 128:
                            if (0.1 >= 0.1) {
                                v26 -= 53
                                v77 := v3[2147483647]
                                var v78 int16 = 1
                                _ = v77
                                _ = v78
                            }
                        }
                    }
                }
            } else {
                for fi79 := 0; fi79 < 6; fi79++ {
                    v26 -= genFn0()
                    _ = fi79
                }
                switch "hello" {
                case "foo bar":
                    v26 = genFn0()
                case "a":
                    var v80 float32 = -2.25
                    _ = v80
                case "hello":
                    var v81 float64 = (3.14159 - 1.5)
                    _ = v81
                }
            }
            _ = v26
        }
        if ((144 < -990) && v1) {
            if v1 {
                switch (fi24 >> ((genFn0()) & 63)) {
                case 4294967295:
                    var v82 float64 = (3.14159 + 1.5)
                    _ = v82
                case 32768:
                    var v83 bool = (v1 && v1)
                    _ = v83
                case 1:
                    _ = len(v3)
                }
            } else {
                v84 := v3[(-(fi24))]
                _ = v84
            }
            v4 = v4
        }
        if v1 {
            for fi85 := 0; fi85 < 20; fi85++ {
                v4 = v4
                var v86 uint8 = 232
                fmt.Println(v86)
                _ = fi85
                _ = v86
            }
            switch v4 {
            case "a":
                for fi87 := 0; fi87 < 9; fi87++ {
                    var v88 uint16 = 65535
                    v4 = v4
                    var v89 int64 = -1
                    _ = fi87
                    _ = v88
                    _ = v89
                }
            case "\n\t":
                fmt.Println(fi24)
            }
        }
        _ = fi24
    }
    if v1 {
        if v1 {
            var v90 int16 = 42
            _ = v90
        } else {
            fmt.Println(v2, v2, v1)
            for fi91 := 0; fi91 < 19; fi91++ {
                var v92 int16 = 1000
                _ = fi91
                _ = v92
            }
        }
        for fi93 := 0; fi93 < 10; fi93++ {
            v1 = v1
            _ = fi93
        }
        v94 := GenS4{F0: v1}
        _ = v94
    }
    if (false || (-2.25 < float32(((1e10 - 1.5) + 1e10)))) {
        var v95 uint16 = 65535
        if (!(((!true) || v1) && v1)) {
            switch "a" {
            case "\n\t":
                v96 := GenS4{F0: (v95 != 42)}
                _ = v96
            case "a":
                if ("hello" <= (v4 + "\n\t")) {
                    v97 := []int{}
                    v1 = ((-2.25 - float32((1.5 - 1.5))) > 0.0)
                    var v98 uint16 = v95
                    _ = v97
                    _ = v98
                }
            default:
                if (v2 != (v2 ^ (v2 / ((638) | 1)))) {
                    for fi99 := 0; fi99 < 20; fi99++ {
                        v95 = (v95 << ((569) & 63))
                        v95 -= 158
                        _ = fi99
                    }
                    v100 := []int{685, 781, 32768}
                    _ = v100
                }
            }
            var v101 int16 = 1000
            if (-19 > 0) {
                switch "\n\t" {
                case "\n\t":
                    var v102 uint64 = 211
                    _ = v102
                case "foo bar":
                    if (42 <= 546) {
                        v103 := v3[516]
                        v104 := []int{}
                        _ = len(v3)
                        _ = v103
                        _ = v104
                    }
                case "hello":
                    var v105 uint32 = (v2 ^ 550)
                    _ = v105
                }
            } else {
                var v106 bool = true
                if false {
                    v101 += (v101 / (((v101 ^ (v101 << ((65535) & 63)))) | 1))
                    v101 = v101
                    for fi107 := 0; fi107 < 10; fi107++ {
                        if false {
                            v108 := GenS2{F0: 385, F1: 869, F2: -80}
                            _ = v108
                        } else {
                            var v109 string = "hello"
                            v110 := GenS2{F0: (^(v95)), F1: 255, F2: 127}
                            v4 = ((((((v109 + (v4 + v4)) + ("foo bar" + "a")) + v109) + (("\n\t" + v109) + "foo bar")) + ((v109 + (v4 + v4)) + "foo bar")) + ("a" + ((((v109 + ("" + (((("foo bar" + v4) + (v109 + ("" + "hello"))) + (("\n\t" + "") + ("a" + ""))) + v4))) + "") + "hello") + (v109 + ((("" + ((("\n\t" + v4) + v109) + "")) + (v109 + ((v4 + v109) + "foo bar"))) + "")))))
                            _ = v109
                            _ = v110
                        }
                        var v111 int64 = -9223372036854775808
                        for fi112 := 0; fi112 < 8; fi112++ {
                            var v113 uint8 = 42
                            _ = fi112
                            _ = v113
                        }
                        _ = fi107
                        _ = v111
                    }
                }
                _ = v106
            }
            _ = v101
        } else {
            v114 := GenS5{F0: (128 != -2147483648), F1: 1, F2: (!(!(-2147483648 < 0)))}
            _ = len(v3)
            for fi115 := 0; fi115 < 1; fi115++ {
                var v116 int16 = -1
                _ = fi115
                _ = v116
            }
            _ = v114
        }
        _ = v95
    }
    fmt.Println(v2, v1)
    v117 := []int{42, 1, 32767}
    for fi118 := 0; fi118 < 10; fi118++ {
        switch fi118 {
        case -1:
            for fi119 := 0; fi119 < 9; fi119++ {
                v3[(fi119 << ((746) & 63))] = "\n\t"
                _ = fi119
            }
        default:
            if (1 != 42) {
                v117 = append(v117, -9223372036854775808)
                v120 := GenS1{F0: v4, F1: 1, F2: 42}
                fmt.Println(fi118, v4)
                _ = v120
            }
        }
        if len(v117) > 0 {
            v117[0] = fi118
        }
        if false {
            if (42 < 65535) {
                fmt.Println(v4)
                if (!true) {
                    for fi121 := 0; fi121 < 6; fi121++ {
                        v3[(fi121 & (-(fi121)))] = (v4 + (v4 + ("foo bar" + "\n\t")))
                        _ = fi121
                    }
                    v122 := GenS6{F0: 824}
                    var v123 float64 = (float64(1e10) / (1e10 - 0.0))
                    _ = v122
                    _ = v123
                } else {
                    v117 = append(v117, fi118)
                    v124 := v3[(^(fi118))]
                    _ = v124
                }
                var v125 float64 = 0.0
                _ = v125
            }
            v126 := GenS6{F0: uint64(v2)}
            _ = v126
        }
        _ = fi118
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v117:", len(v117), v117)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v117
    _ = v2
    _ = v3
    _ = v4
}
