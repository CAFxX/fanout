package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 uint }
type GenS2 struct { F0 int64; F1 bool; F2 int16 }
type GenS3 struct { F0 int32; F1 uint8; F2 string }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint16; F1 int32 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint; F1 bool }

func main() {
    v1 := make(map[int]bool)
    v2 := []uint32{}
    var v3 int32 = -2147483648
    var v4 uint8 = 200
    var v5 int32 = (v3 / (((v3 * v3)) | 1))
    var v6 uint32 = 2147483648
    v7 := v1[128]
    var v8 int64 = 39
    var v9 int64 = v8
    if v7 {
        v4 = v4
        var v10 int8 = 0
        for fi11 := 0; fi11 < 2; fi11++ {
            switch "a" {
            case "foo bar":
                for fi12 := 0; fi12 < 9; fi12++ {
                    var v13 float32 = 0.0
                    if true {
                        var v14 int = fi12
                        _ = v14
                    }
                    v5 = v3
                    _ = fi12
                    _ = v13
                }
            default:
                v15 := GenS1{F0: 912, F1: fi11, F2: 255}
                _ = v15
            }
            v8 += -1
            _ = fi11
        }
        _ = v10
    }
    v5 += v3
    v2 = append(v2, v6)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v16 int8 = -1
    v17 := GenS1{F0: (v6 % (((v6 >> ((277) & 63))) | 1)), F1: -1, F2: 4294967295}
    var v18 uint16 = 1
    fmt.Println(v8)
    if v7 {
        for fi19 := 0; fi19 < 16; fi19++ {
            for fi20 := 0; fi20 < 9; fi20++ {
                v21 := GenS2{F0: (v8 << ((32) & 63)), F1: (3.14159 <= 1.5), F2: -1}
                if false {
                    fmt.Println(v4, v3)
                    v16 = v16
                    if len(v2) > 0 {
                        v2[0] = (v6 ^ v6)
                    }
                } else {
                    switch "" {
                    case "\n\t":
                        var v22 int = (fi19 + 32768)
                        if (0 <= v5) {
                            v23 := GenS1{F0: v6, F1: v22, F2: 505}
                            v3 = (v5 << ((980) & 63))
                            if false {
                                _ = len(v1)
                            }
                            _ = v23
                        } else {
                            v24 := GenS1{F0: (v6 - (v6 & 47)), F1: v22, F2: 65535}
                            var v25 uint16 = (^(v18))
                            v2 = append(v2, v6)
                            _ = v24
                            _ = v25
                        }
                        _ = v22
                    case "":
                        var v26 uint8 = (v4 >> ((116) & 63))
                        _ = v26
                    }
                }
                v2 = append(v2, v6)
                _ = fi20
                _ = v21
            }
            var v27 int16 = 0
            _ = fi19
            _ = v27
        }
        v28 := GenS2{F0: (v9 + v9), F1: (!(!((!true) && v7))), F2: -500}
        _ = v28
    }
    if (127 < -2147483648) {
        v8 += (v8 * (v8 >> ((0) & 63)))
        var v29 uint64 = 18446744073709551615
        _ = v29
    }
    var v30 int64 = v9
    var v31 uint16 = uint16(v4)
    fmt.Println(v5, v9)
    for fi32 := 0; fi32 < 17; fi32++ {
        if v7 {
            var v33 uint16 = (v18 | (v18 ^ (v18 + (v31 - v18))))
            v34 := GenS2{F0: v30, F1: (!v7), F2: -32768}
            if v7 {
                for fi35 := 0; fi35 < 5; fi35++ {
                    if v7 {
                        v36 := GenS3{F0: 42, F1: 1, F2: "foo bar"}
                        _ = v36
                    }
                    if v7 {
                        v31 = 273
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        v5 += (v5 << ((1) & 63))
                    } else {
                        var v37 bool = true
                        var v38 int32 = v5
                        v31 = (v33 | (v31 / ((0) | 1)))
                        _ = v37
                        _ = v38
                    }
                    var v39 uint16 = 1
                    _ = fi35
                    _ = v39
                }
                for fi40 := 0; fi40 < 19; fi40++ {
                    switch "hello" {
                    case "a":
                        if len(v2) > 0 {
                            v2[0] = (^(v6))
                        }
                    }
                    var v41 bool = ((fi40 / ((32767) | 1)) <= 127)
                    var v42 uint8 = v4
                    _ = fi40
                    _ = v41
                    _ = v42
                }
            }
            _ = v33
            _ = v34
        } else {
            _ = len(v1)
            if ((v18 - 32767) != (v18 * (v31 << ((389) & 63)))) {
                if (!(v7 && (0 >= 4294967295))) {
                    if (!v7) {
                        for fi43 := 0; fi43 < 3; fi43++ {
                            var v44 int64 = -1
                            _ = fi43
                            _ = v44
                        }
                        for fi45 := 0; fi45 < 19; fi45++ {
                            var v46 uint64 = 149
                            fmt.Println(v30, v3, v4)
                            var v47 int16 = 0
                            var v48 uint64 = (v46 | 18446744073709551615)
                            _ = fi45
                            _ = v46
                            _ = v47
                            _ = v48
                        }
                        if (!(v7 || false)) {
                            var v49 uint16 = v31
                            if false {
                                switch (^(fi32)) {
                                case -460:
                                    if false {
                                        v50 := GenS3{F0: 42, F1: (v4 * v4), F2: "foo bar"}
                                        _ = v50
                                    }
                                case 128:
                                    switch ("" + "") {
                                    case "\n\t":
                                        for fi51 := 0; fi51 < 17; fi51++ {
                                            if v7 {
                                                v4 = (v4 * 206)
                                                v52 := GenS4{F0: v49}
                                                if (!((v4 * v4) > v4)) {
                                                    v53 := v1[(fi51 >> ((2) & 63))]
                                                    fmt.Println(v18, v3, v16)
                                                    v4 = uint8(v6)
                                                    _ = v53
                                                }
                                                _ = v52
                                            }
                                            var v54 uint8 = 255
                                            v30 -= -9223372036854775808
                                            var v55 uint64 = 2
                                            delete(v1, (fi32 % (((fi51 % (((fi32 ^ fi51)) | 1))) | 1)))
                                            _ = fi51
                                            _ = v54
                                            _ = v55
                                        }
                                    case "hello":
                                        v56 := GenS2{F0: (v30 >> ((18446744073709551615) & 63)), F1: (0.0 == (float32(0.1) * (0.1 - 0.0))), F2: 1}
                                        _ = v56
                                    case "foo bar":
                                        v5 = v3
                                    }
                                case -1:
                                    v31 = (v49 ^ v49)
                                }
                                var v57 int8 = (v16 / ((-3) | 1))
                                if (!v7) {
                                    var v58 int64 = (v9 * v8)
                                    v5 = v3
                                    _ = v58
                                } else {
                                    v59 := GenS5{F0: 2, F1: v5}
                                    switch ("" + "") {
                                    case "hello":
                                        v2 = append(v2, (v6 - (^(v6))))
                                    case "":
                                        _ = len(v1)
                                    case "\n\t":
                                        var v60 string = "a"
                                        _ = v60
                                    default:
                                        var v61 uint16 = v49
                                        _ = v61
                                    }
                                    fmt.Println(v18, v6)
                                    _ = v59
                                }
                                _ = v57
                            } else {
                                for fi62 := 0; fi62 < 17; fi62++ {
                                    v18 = 884
                                    _ = fi62
                                }
                                switch (fi32 & 127) {
                                case -2147483648:
                                    switch "hello" {
                                    case "a":
                                        var v63 uint32 = 1
                                        _ = v63
                                    }
                                case -248:
                                    switch "a" {
                                    case "":
                                        if v7 {
                                            var v64 int32 = -2147483648
                                            if false {
                                                for fi65 := 0; fi65 < 9; fi65++ {
                                                    var v66 uint8 = uint8(v6)
                                                    if false {
                                                        if true {
                                                            var v67 int64 = v30
                                                            _ = v67
                                                        }
                                                        var v68 uint16 = (^(v18))
                                                        _ = v68
                                                    } else {
                                                        var v69 int32 = -786
                                                        delete(v1, 330)
                                                        if true {
                                                            if (!(v7 || v7)) {
                                                                switch 32767 {
                                                                case -1:
                                                                    v4 = v66
                                                                case -757:
                                                                    switch ((("foo bar" + "") + "") + "") {
                                                                    case "a":
                                                                        switch ("foo bar" + "") {
                                                                        case "\n\t":
                                                                            v66 -= v4
                                                                        }
                                                                    default:
                                                                        v70 := GenS1{F0: v6, F1: 128, F2: 1}
                                                                        _ = v70
                                                                    }
                                                                default:
                                                                    fmt.Println(v4, v49)
                                                                }
                                                                var v71 int64 = int64(v16)
                                                                if v7 {
                                                                    var v72 uint16 = (v49 >> ((2) & 63))
                                                                    if false {
                                                                        var v73 string = "hello"
                                                                        var v74 int8 = (^(v16))
                                                                        v8 += -323
                                                                        var v75 uint32 = 65535
                                                                        if ((!false) && v7) {
                                                                            var v76 uint8 = v4
                                                                            v2 = append(v2, (^(v75)))
                                                                            v77 := GenS2{F0: v8, F1: v7, F2: int16(v8)}
                                                                            _ = v76
                                                                            _ = v77
                                                                        } else {
                                                                            v78 := GenS4{F0: 388}
                                                                            if false {
                                                                                v2 = append(v2, (^(v75)))
                                                                                var v79 int8 = (v16 >> ((916) & 63))
                                                                                v16 = (v74 & (v79 - v79))
                                                                                if (!v7) {
                                                                                    switch int(v64) {
                                                                                    case -2147483648:
                                                                                        if (v7 || false) {
                                                                                            v66 -= (v4 * v66)
                                                                                        }
                                                                                    case 127:
                                                                                        if len(v2) >= 2 {
                                                                                            v2 = v2[1:len(v2)]
                                                                                        }
                                                                                    default:
                                                                                        fmt.Println(v7)
                                                                                    }
                                                                                }
                                                                                _ = v79
                                                                            }
                                                                            if len(v2) >= 2 {
                                                                                v2 = v2[1:len(v2)]
                                                                            }
                                                                            v64 = v3
                                                                            if v7 {
                                                                                if v7 {
                                                                                    if false {
                                                                                        v1[fi65] = v7
                                                                                    }
                                                                                } else {
                                                                                    v80 := v1[(fi65 ^ (^(fi32)))]
                                                                                    fmt.Println(v66, v71)
                                                                                    _ = v80
                                                                                }
                                                                                var v81 uint32 = uint32(v49)
                                                                                if len(v2) >= 2 {
                                                                                    v2 = v2[1:len(v2)]
                                                                                }
                                                                                _ = v81
                                                                            } else {
                                                                                if (!v7) {
                                                                                    var v82 int64 = (v8 * v71)
                                                                                    v83 := GenS4{F0: (v49 / ((v18) | 1))}
                                                                                    var v84 int = 9223372036854775807
                                                                                    _ = v82
                                                                                    _ = v83
                                                                                    _ = v84
                                                                                }
                                                                            }
                                                                            _ = v78
                                                                        }
                                                                        v85 := GenS2{F0: (v71 * 42), F1: true, F2: 1}
                                                                        v86 := GenS6{F0: (v69 % ((42) | 1))}
                                                                        _ = v73
                                                                        _ = v74
                                                                        _ = v75
                                                                        _ = v85
                                                                        _ = v86
                                                                    } else {
                                                                        v64 = (v64 >> ((691) & 63))
                                                                    }
                                                                    v87 := v1[(fi65 >> ((0) & 63))]
                                                                    _ = v72
                                                                    _ = v87
                                                                } else {
                                                                    var v88 int32 = 1
                                                                    _ = v88
                                                                }
                                                                _ = v71
                                                            }
                                                        } else {
                                                            if len(v2) > 0 {
                                                                v2[0] = v6
                                                            }
                                                            v16 -= v16
                                                        }
                                                        if (!false) {
                                                            fmt.Println(v69)
                                                            if ((v49 + 32767) != v49) {
                                                                v2 = append(v2, v6)
                                                                var v89 int8 = v16
                                                                v90 := GenS4{F0: 32767}
                                                                _ = v89
                                                                _ = v90
                                                            } else {
                                                                var v91 uint32 = 65535
                                                                v6 = 4294967295
                                                                var v92 int32 = (v5 % (((v64 + (v5 * 42))) | 1))
                                                                var v93 uint16 = v18
                                                                var v94 string = ""
                                                                _ = v91
                                                                _ = v92
                                                                _ = v93
                                                                _ = v94
                                                            }
                                                            var v95 float64 = (((((-2.25 - 3.14159) / -2.25) - (float64(1e10) * (((0.1 - ((float64(((1e10 + (-2.25 + 3.14159)) + 1.5)) - (float64(3.14159) + -2.25)) / (float64(1e10) * 1.5))) - -2.25) - float64(0.0)))) / 1.0) + 0.1)
                                                            _ = v95
                                                        }
                                                        var v96 uint16 = 0
                                                        if (!(("" + "foo bar") <= (("\n\t" + "foo bar") + "hello"))) {
                                                            if len(v2) > 0 {
                                                                v2[0] = v6
                                                            }
                                                            fmt.Println(v30, v49, v69)
                                                            var v97 uint16 = 255
                                                            _ = v97
                                                        } else {
                                                            var v98 int64 = (v9 >> ((65535) & 63))
                                                            v49 = v96
                                                            var v99 uint16 = (v96 - (v31 / ((263) | 1)))
                                                            _ = v98
                                                            _ = v99
                                                        }
                                                        _ = v69
                                                        _ = v96
                                                    }
                                                    v8 = (v30 % ((-611) | 1))
                                                    if v7 {
                                                        var v100 float64 = ((((1.5 + 1.5) * 1.5) + 1.5) * (1.5 - ((3.14159 - 0.0) / 1.5)))
                                                        var v101 int32 = int32(v18)
                                                        v102 := GenS7{F0: 266}
                                                        _ = v100
                                                        _ = v101
                                                        _ = v102
                                                    }
                                                    fmt.Println(v7, v66, v8)
                                                    _ = fi65
                                                    _ = v66
                                                }
                                                v103 := GenS5{F0: 32767, F1: 1}
                                                var v104 bool = (255 != 18446744073709551615)
                                                _ = v103
                                                _ = v104
                                            }
                                            var v105 uint16 = v31
                                            _ = v105
                                            _ = v64
                                        } else {
                                            if (!(v31 >= v31)) {
                                                var v106 uint = 18446744073709551615
                                                if v7 {
                                                    for fi107 := 0; fi107 < 18; fi107++ {
                                                        v108 := GenS8{F0: 2, F1: (("a" + (((("\n\t" + "hello") + "a") + "\n\t") + "foo bar")) < (("hello" + ("" + ("foo bar" + ""))) + ("\n\t" + ("a" + "foo bar"))))}
                                                        var v109 int32 = -219
                                                        switch "\n\t" {
                                                        case "":
                                                            var v110 uint16 = (^(v18))
                                                            _ = v110
                                                        }
                                                        _ = fi107
                                                        _ = v108
                                                        _ = v109
                                                    }
                                                } else {
                                                    for fi111 := 0; fi111 < 4; fi111++ {
                                                        fmt.Println(v49, fi111, v18)
                                                        _ = fi111
                                                    }
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    if v7 {
                                                        fmt.Println(v8, v49, v49)
                                                        var v112 int = (fi32 << (((v106 << ((uint(v49)) & 63))) & 63))
                                                        _ = v112
                                                    }
                                                }
                                                switch (("a" + ("" + "a")) + "foo bar") {
                                                case "foo bar":
                                                    for fi113 := 0; fi113 < 3; fi113++ {
                                                        v114 := GenS1{F0: 4294967295, F1: -1, F2: v106}
                                                        _ = fi113
                                                        _ = v114
                                                    }
                                                case "\n\t":
                                                    delete(v1, fi32)
                                                default:
                                                    v8 += -1
                                                }
                                                _ = v106
                                            }
                                        }
                                    }
                                default:
                                    var v115 float32 = 0.0
                                    _ = v115
                                }
                                fmt.Println(v16)
                            }
                            var v116 bool = true
                            _ = v116
                            _ = v49
                        }
                    }
                    var v117 uint32 = uint32(v18)
                    _ = v117
                } else {
                    _ = len(v1)
                }
                v30 = (v8 * int64(fi32))
                for fi118 := 0; fi118 < 8; fi118++ {
                    if true {
                        var v119 uint16 = v18
                        _ = v119
                    } else {
                        var v120 int32 = (v5 + int32(v6))
                        var v121 int8 = (v16 * -1)
                        var v122 uint16 = (v18 - (v18 / ((795) | 1)))
                        _ = v120
                        _ = v121
                        _ = v122
                    }
                    _ = fi118
                }
            }
            switch fi32 {
            case -2147483648:
                var v123 float64 = 0.0
                _ = v123
            }
        }
        _ = fi32
    }
    v2 = append(v2, (v6 + v6))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
