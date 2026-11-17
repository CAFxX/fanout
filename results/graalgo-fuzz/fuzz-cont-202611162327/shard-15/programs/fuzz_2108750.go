package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint; F1 string; F2 int }
type GenS3 struct { F0 bool; F1 uint }
type GenS4 struct { F0 int16; F1 int64 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int16; F1 uint8; F2 uint32 }
type GenS8 struct { F0 string; F1 string; F2 uint32 }
type GenS9 struct { F0 int8; F1 int; F2 int16 }
type GenS10 struct { F0 uint64 }
func genFn0(p0 uint, p1 uint32) uint8 {
    var v7 uint64 = 93
    if true {
        switch "hello" {
        case "a":
            for fi8 := 0; fi8 < 12; fi8++ {
                var v9 uint64 = 4294967295
                _ = fi8
                _ = v9
            }
        case "foo bar":
            for fi10 := 0; fi10 < 20; fi10++ {
                fmt.Println(fi10, p0)
                p0 -= (^(p0))
                _ = fi10
            }
        default:
            v11 := []int{int(v7), 0, -80}
            _ = v11
        }
        var v12 float64 = 1.5
        _ = v12
    } else {
        for fi13 := 0; fi13 < 8; fi13++ {
            for fi14 := 0; fi14 < 18; fi14++ {
                var v15 uint32 = (^(p1))
                v16 := GenS1{F0: 633}
                _ = fi14
                _ = v15
                _ = v16
            }
            v17 := []int{(-(fi13))}
            _ = fi13
            _ = v17
        }
    }
    if (true || ((!false) && false)) {
        if (-9223372036854775808 != 696) {
            if (3.14159 > ((0.1 + -2.25) / (1e10 - (1e10 * (((0.1 - 1e10) - (float64(1e10) + (0.0 - (1e10 / 3.14159)))) + 1.5))))) {
                if true {
                    if false {
                        var v18 uint = p0
                        var v19 uint = v18
                        for fi20 := 0; fi20 < 20; fi20++ {
                            v21 := GenS1{F0: v7}
                            v7 += v7
                            _ = fi20
                            _ = v21
                        }
                        _ = v18
                        _ = v19
                    } else {
                        switch 32767 {
                        case 128:
                            var v22 uint16 = 0
                            _ = v22
                        }
                        p0 += 18446744073709551615
                    }
                } else {
                    var v23 uint32 = (p1 - 2)
                    v23 = (p1 ^ (p1 / ((931) | 1)))
                    _ = v23
                }
                switch -343 {
                case 9223372036854775807:
                    p1 -= 1
                default:
                    for fi24 := 0; fi24 < 1; fi24++ {
                        v25 := []int{}
                        _ = fi24
                        _ = v25
                    }
                }
                if ("foo bar" == "\n\t") {
                    for fi26 := 0; fi26 < 17; fi26++ {
                        if false {
                            var v27 int = -9223372036854775808
                            v7 = v7
                            _ = v27
                        } else {
                            fmt.Println(fi26, fi26, v7)
                            v28 := GenS1{F0: (^(v7))}
                            _ = v28
                        }
                        for fi29 := 0; fi29 < 1; fi29++ {
                            v30 := []int{fi26, (fi29 ^ int(p1)), 1, (fi29 << ((p0) & 63))}
                            p1 -= 285
                            var v31 uint64 = v7
                            _ = fi29
                            _ = v30
                            _ = v31
                        }
                        var v32 uint16 = 2
                        _ = fi26
                        _ = v32
                    }
                }
            } else {
                p1 = uint32(p0)
                for fi33 := 0; fi33 < 3; fi33++ {
                    var v34 float64 = 0.1
                    var v35 int8 = 34
                    _ = fi33
                    _ = v34
                    _ = v35
                }
            }
            var v36 int16 = 139
            var v37 int16 = v36
            _ = v36
            _ = v37
        } else {
            p1 = 2147483647
            p0 += p0
            var v38 string = "foo bar"
            _ = v38
        }
    }
    _ = p0
    _ = p1
    _ = v7
    return 2
}
func genFn1() int32 {
    switch -451 {
    case 0:
        switch "foo bar" {
        case "foo bar":
            if (1 > 2147483648) {
                var v39 int8 = 1
                if true {
                    v39 += 100
                    var v40 uint16 = 32767
                    var v41 int8 = (v39 << ((2) & 63))
                    _ = v40
                    _ = v41
                } else {
                    v42 := []int{}
                    _ = v42
                }
                for fi43 := 0; fi43 < 14; fi43++ {
                    var v44 float32 = float32(0.1)
                    for fi45 := 0; fi45 < 10; fi45++ {
                        var v46 uint32 = 2147483648
                        _ = fi45
                        _ = v46
                    }
                    _ = fi43
                    _ = v44
                }
                _ = v39
            }
        case "\n\t":
            for fi47 := 0; fi47 < 6; fi47++ {
                v48 := []int{(-(fi47)), fi47, 32768, fi47}
                _ = fi47
                _ = v48
            }
        default:
            v49 := []int{}
            _ = v49
        }
    case 128:
        var v50 uint32 = 776
        _ = v50
    case -9223372036854775808:
        v51 := []int{1, 42}
        _ = v51
    default:
        v52 := GenS2{F0: 65535, F1: ("foo bar" + "hello"), F2: 9223372036854775807}
        _ = v52
    }
    var v53 uint = 826
    v54 := []int{127, -515, -315}
    var v55 float64 = 3.14159
    var v56 bool = ("foo bar" == "\n\t")
    _ = v53
    _ = v54
    _ = v55
    _ = v56
    return -2147483648
}

func main() {
    v1 := make(map[string]string)
    var v2 float64 = ((((-2.25 + 0.1) * float64(((float32(float64((((3.14159 / 3.14159) * (1e10 * (float32(1.5) + 0.0))) + -2.25))) * -2.25) - 0.1))) / 1.5) - (float64((1.5 * 0.0)) / 1.0))
    v3 := make(map[uint32]int)
    var v4 uint = 255
    v5 := make(map[uint32]uint32)
    v6 := make(map[uint32]string)
    v1["foo bar"] = "a"
    v4 = 4294967295
    v57 := GenS2{F0: 2, F1: "", F2: -9223372036854775808}
    v58 := []int{}
    v2 += v2
    var v59 int64 = 42
    var v60 int64 = v59
    var v61 uint8 = 16
    switch 42 {
    case 42:
        var v62 bool = true
        _ = v62
    default:
        v63 := GenS1{F0: 18446744073709551615}
        _ = v63
    }
    v2 += v2
    switch 0 {
    case 127:
        var v64 float32 = 3.14159
        _ = v64
    case -2147483648:
        v61 += 199
    default:
        var v65 float64 = (((v2 - v2) * v2) * v2)
        _ = v65
    }
    v61 = (v61 >> (((^(v4))) & 63))
    for fi66 := 0; fi66 < 3; fi66++ {
        switch "\n\t" {
        case "a":
            v2 -= v2
        }
        _ = fi66
    }
    for fi67 := 0; fi67 < 11; fi67++ {
        var v68 float64 = 0.0
        if (false && (32767 == 1000)) {
            v69 := GenS3{F0: ("\n\t" != ("foo bar" + "foo bar")), F1: (v4 * 2)}
            _ = v69
        } else {
            for fi70 := 0; fi70 < 10; fi70++ {
                var v71 int32 = -1
                _ = fi70
                _ = v71
            }
            for fi72 := 0; fi72 < 15; fi72++ {
                v59 += v59
                switch "hello" {
                case "foo bar":
                    delete(v3, 4294967295)
                case "":
                    switch fi72 {
                    case 127:
                        switch (fi67 >> (((v4 % ((v4) | 1))) & 63)) {
                        case 509:
                            switch (^(fi72)) {
                            case 4294967295:
                                v2 = (float64(float32(v68)) + ((v68 / v2) / v2))
                            default:
                                fmt.Println(v59)
                            }
                            v2 -= (float64((3.14159 * (1e10 + (-2.25 - 0.1)))) * (0.1 - ((1e10 + (v2 + 1.5)) - v2)))
                            var v73 uint32 = 1
                            var v74 uint = v4
                            _ = v73
                            _ = v74
                        case -244:
                            if false {
                                var v75 float32 = 0.0
                                _ = v75
                            }
                        case 128:
                            v60 -= v59
                        default:
                            v68 = 0.1
                        }
                    default:
                        if (!false) {
                            if false {
                                fmt.Println(fi67)
                            } else {
                                fmt.Println(v60, fi72)
                            }
                            v60 -= (v60 * -749)
                        } else {
                            switch ("a" + "a") {
                            case "hello":
                                fmt.Println(fi67)
                            case "a":
                                var v76 uint8 = v61
                                _ = v76
                            default:
                                if true {
                                    v59 += 0
                                    var v77 float64 = 3.14159
                                    _ = v77
                                } else {
                                    v2 += v2
                                    v2 = 1.5
                                    var v78 string = "\n\t"
                                    _ = v78
                                }
                            }
                            if len(v58) >= 2 {
                                v58 = v58[1:len(v58)]
                            }
                            if true {
                                if (!(((v2 + ((float64(0.0) + v2) + 3.14159)) + float64(((((0.1 / 1.5) * (0.1 - float32(((float64(1.5) - 0.0) + v68)))) - 1e10) / 3.14159))) <= (1.5 * float64(((1.5 / 3.14159) - 3.14159))))) {
                                    if (v61 >= 38) {
                                        delete(v5, 2147483648)
                                        _ = len(v3)
                                    } else {
                                        var v79 uint16 = 854
                                        _ = v79
                                    }
                                    var v80 uint64 = 0
                                    v81 := GenS4{F0: 142, F1: v59}
                                    _ = v80
                                    _ = v81
                                }
                            } else {
                                v61 = v61
                                var v82 float32 = float32((v68 + v68))
                                var v83 uint16 = 255
                                _ = v82
                                _ = v83
                            }
                            if (29 == 127) {
                                var v84 bool = ((v60 % ((v60) | 1)) <= 0)
                                v4 -= v4
                                switch "" {
                                case "a":
                                    switch 42 {
                                    case -725:
                                        delete(v3, 2)
                                    case -480:
                                        var v85 int16 = -32768
                                        _ = v85
                                    }
                                case "hello":
                                    if (!(1000 != 32767)) {
                                        v86 := GenS4{F0: 1, F1: v59}
                                        if (2147483648 >= 65535) {
                                            if (18 < -128) {
                                                if ((v68 >= (0.1 + -2.25)) && true) {
                                                    var v87 int = fi72
                                                    v60 += (v59 % (((v59 * v59)) | 1))
                                                    _ = v87
                                                }
                                                if true {
                                                    v2 += v2
                                                } else {
                                                    v2 += v68
                                                }
                                                fmt.Println(v2, v61, v4)
                                            } else {
                                                var v88 bool = true
                                                var v89 uint8 = v61
                                                _ = v88
                                                _ = v89
                                            }
                                            v4 = 65535
                                            if (v84 && ((v4 - v4) > 255)) {
                                                _ = len(v58)
                                                var v90 int32 = -314
                                                _ = v90
                                            }
                                        } else {
                                            var v91 uint64 = 2
                                            if ((58 >= -1) && v84) {
                                                var v92 uint = v4
                                                if (!v84) {
                                                    var v93 uint32 = 1
                                                    v93 += (v93 * (^(v93)))
                                                    var v94 int = fi67
                                                    _ = v93
                                                    _ = v94
                                                } else {
                                                    switch ("foo bar" + ("" + "")) {
                                                    case "a":
                                                        if true {
                                                            v91 -= (v91 / ((uint64(v61)) | 1))
                                                        } else {
                                                            var v95 string = ("" + "a")
                                                            v96 := GenS1{F0: 42}
                                                            v97 := GenS5{F0: v95}
                                                            _ = v95
                                                            _ = v96
                                                            _ = v97
                                                        }
                                                    case "":
                                                        _ = len(v58)
                                                    }
                                                    fmt.Println(fi72, v84)
                                                }
                                                v98 := GenS6{F0: (v61 + (^(v61)))}
                                                switch ("a" + "foo bar") {
                                                case "\n\t":
                                                    v99 := GenS6{F0: 255}
                                                    _ = v99
                                                case "":
                                                    var v100 float32 = (0.1 + -2.25)
                                                    _ = v100
                                                default:
                                                    fmt.Println(v91)
                                                }
                                                var v101 int32 = 42
                                                _ = v101
                                                _ = v92
                                                _ = v98
                                            }
                                            _ = v91
                                        }
                                        v102 := GenS7{F0: 32767, F1: (v61 + 0), F2: 750}
                                        _ = v102
                                        _ = v86
                                    } else {
                                        v103 := GenS3{F0: v84, F1: v4}
                                        if (3.14159 >= (-2.25 / 1.0)) {
                                            v104 := GenS3{F0: ((fi72 / ((23) | 1)) >= fi72), F1: 1}
                                            fmt.Println(v4, v2, v2)
                                            switch int(v4) {
                                            case -2147483648:
                                                v105 := GenS8{F0: ("foo bar" + ("hello" + "a")), F1: "foo bar", F2: 2147483647}
                                                _ = v105
                                            }
                                            if v84 {
                                                fmt.Println(v84, v68, v2)
                                                v4 = v4
                                            } else {
                                                switch "hello" {
                                                case "":
                                                    v6[359] = "foo bar"
                                                default:
                                                    v58 = append(v58, fi67)
                                                }
                                            }
                                            var v106 bool = (-571 <= 0)
                                            var v107 uint64 = 238
                                            _ = v104
                                            _ = v106
                                            _ = v107
                                        }
                                        if (!(1 != (^(v59)))) {
                                            if len(v58) >= 2 {
                                                v58 = v58[1:len(v58)]
                                            }
                                            fmt.Println(fi72)
                                            v58 = append(v58, (fi67 - fi67))
                                            v4 = (^(v4))
                                            delete(v1, "")
                                            if (genFn0((^(v4)), 0) > (v61 / (((v61 ^ genFn0(4294967295, 468))) | 1))) {
                                                fmt.Println(v61, v4, fi72)
                                            } else {
                                                if v84 {
                                                    switch ("" + "a") {
                                                    case "hello":
                                                        v61 = 42
                                                    default:
                                                        v108 := v6[uint32(v61)]
                                                        _ = v108
                                                    }
                                                } else {
                                                    fmt.Println(v84)
                                                    switch fi67 {
                                                    case 1:
                                                        switch 32768 {
                                                        case 966:
                                                            v109 := GenS8{F0: "", F1: "hello", F2: 1}
                                                            _ = v109
                                                        default:
                                                            v110 := GenS9{F0: 68, F1: fi67, F2: -32768}
                                                            fmt.Println(v84, v4)
                                                            var v111 int = 128
                                                            _ = v110
                                                            _ = v111
                                                        }
                                                    }
                                                    v61 -= (v61 * 127)
                                                }
                                            }
                                            v68 -= v68
                                        }
                                        _ = v103
                                    }
                                }
                                switch fi72 {
                                case -133:
                                    delete(v3, uint32(v61))
                                case -910:
                                    var v112 uint8 = v61
                                    _ = v112
                                case -31:
                                    if v84 {
                                        v3[230] = (fi67 ^ fi67)
                                        fmt.Println(v61, v59, v60)
                                        if (18446744073709551615 == 741) {
                                            v113 := GenS5{F0: ""}
                                            v114 := GenS10{F0: 830}
                                            _ = v113
                                            _ = v114
                                        }
                                    } else {
                                        v3[4294967295] = (fi72 ^ (fi72 ^ -1))
                                    }
                                default:
                                    var v115 bool = (true || v84)
                                    _ = v115
                                }
                                if true {
                                    v2 -= (float64(float32((1.5 * (1e10 / (((v2 + v68) - -2.25) - (1e10 / (float64(-2.25) + 1.5))))))) + v2)
                                    switch (fi72 - fi72) {
                                    case 455:
                                        v68 = 0.0
                                    case -2147483648:
                                        v61 += (v61 / (((^(v61))) | 1))
                                    default:
                                        v59 -= v60
                                    }
                                }
                                _ = v84
                            } else {
                                var v116 int8 = -78
                                delete(v5, 2147483648)
                                if false {
                                    v59 += int64(v61)
                                    v117 := v1[(("\n\t" + ("a" + ((("a" + "\n\t") + "\n\t") + "hello"))) + "hello")]
                                    var v118 int = (fi67 % ((2147483647) | 1))
                                    _ = v117
                                    _ = v118
                                }
                                _ = v116
                            }
                            fmt.Println(v2)
                        }
                    }
                }
                _ = fi72
            }
        }
        _ = fi67
        _ = v68
    }
    for fi119 := 0; fi119 < 19; fi119++ {
        v120 := GenS8{F0: "hello", F1: (("foo bar" + "hello") + ("hello" + "hello")), F2: 276}
        _ = fi119
        _ = v120
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v58:", len(v58), v58)
    fmt.Println("v59:", v59)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v58
    _ = v59
    _ = v6
    _ = v60
    _ = v61
}
