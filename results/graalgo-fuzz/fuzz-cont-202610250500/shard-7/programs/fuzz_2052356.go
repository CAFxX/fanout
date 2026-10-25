package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 bool }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64; F1 uint32 }
type GenS4 struct { F0 uint; F1 int64 }
type GenS5 struct { F0 bool; F1 uint64; F2 int8 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 bool; F1 int8; F2 int }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int64; F1 int }
func genFn0() int16 {
    var v6 uint = 253
    var v7 uint8 = 128
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 int = (fi8 % ((fi8) | 1))
        for fi10 := 0; fi10 < 18; fi10++ {
            if (!(!true)) {
                v7 = 2
                v9 = 128
                var v11 uint = v6
                _ = v11
            }
            var v12 int32 = -1
            _ = fi10
            _ = v12
        }
        var v13 uint32 = 685
        _ = fi8
        _ = v13
        _ = v9
    }
    _ = v6
    _ = v7
    return -605
}
func genFn1() int32 {
    switch ((("hello" + "") + ("a" + "")) + "\n\t") {
    case "":
        if (0 < 1000) {
            v14 := GenS1{F0: 225, F1: 64, F2: (!(210 == 1))}
            if true {
                var v15 float32 = 3.14159
                fmt.Println(v15)
                var v16 uint64 = 42
                _ = v15
                _ = v16
            } else {
                for fi17 := 0; fi17 < 9; fi17++ {
                    var v18 uint32 = 2147483647
                    v19 := []int{(fi17 + (fi17 + 977))}
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                var v20 uint16 = 65535
                fmt.Println(v20)
                _ = v20
            }
            _ = v14
        }
    case "\n\t":
        for fi21 := 0; fi21 < 8; fi21++ {
            var v22 bool = false
            for fi23 := 0; fi23 < 19; fi23++ {
                var v24 int = (^(fi23))
                _ = fi23
                _ = v24
            }
            v25 := []int{(fi21 ^ fi21), fi21, fi21}
            _ = fi21
            _ = v22
            _ = v25
        }
    case "a":
        for fi26 := 0; fi26 < 2; fi26++ {
            switch "" {
            case "\n\t":
                var v27 float32 = (-2.25 * 0.1)
                _ = v27
            case "a":
                var v28 float64 = -2.25
                _ = v28
            case "":
                if (255 < 18446744073709551615) {
                    if (("" <= ("a" + ("\n\t" + (("\n\t" + "foo bar") + "hello")))) && (true || true)) {
                        var v29 uint16 = 42
                        v29 = v29
                        _ = v29
                    } else {
                        for fi30 := 0; fi30 < 4; fi30++ {
                            var v31 bool = (!(true || (457 == 2)))
                            _ = fi30
                            _ = v31
                        }
                    }
                    if true {
                        var v32 int = 42
                        for fi33 := 0; fi33 < 9; fi33++ {
                            v32 -= fi26
                            var v34 uint = 65535
                            fmt.Println(fi26, v32, v32)
                            _ = fi33
                            _ = v34
                        }
                        _ = v32
                    } else {
                        v35 := GenS2{F0: 65535}
                        var v36 uint32 = 42
                        _ = v35
                        _ = v36
                    }
                    if false {
                        v37 := []int{fi26}
                        for fi38 := 0; fi38 < 5; fi38++ {
                            v39 := GenS1{F0: 436, F1: 42, F2: (4294967295 > 18446744073709551615)}
                            _ = fi38
                            _ = v39
                        }
                        for fi40 := 0; fi40 < 1; fi40++ {
                            var v41 uint64 = 18446744073709551615
                            var v42 uint64 = v41
                            var v43 string = ""
                            _ = fi40
                            _ = v41
                            _ = v42
                            _ = v43
                        }
                        _ = v37
                    }
                } else {
                    if (!(42 < 0)) {
                        fmt.Println(fi26)
                    } else {
                        var v44 uint = 65535
                        v44 -= v44
                        if false {
                            v44 += (v44 & v44)
                            if false {
                                switch "foo bar" {
                                case "":
                                    fmt.Println(fi26, fi26)
                                case "a":
                                    v44 = 1
                                case "hello":
                                    v45 := GenS3{F0: 518, F1: 65535}
                                    _ = v45
                                default:
                                    v44 = (^(v44))
                                }
                                switch "" {
                                case "\n\t":
                                    v46 := []int{(^(fi26)), (fi26 % ((0) | 1)), (fi26 << (((^(v44))) & 63))}
                                    _ = v46
                                case "a":
                                    if (1 < 1) {
                                        fmt.Println(fi26, fi26)
                                        var v47 uint = v44
                                        _ = v47
                                    }
                                case "foo bar":
                                    v48 := GenS4{F0: 4294967295, F1: 9223372036854775807}
                                    _ = v48
                                }
                                for fi49 := 0; fi49 < 6; fi49++ {
                                    fmt.Println(fi26, fi49)
                                    v50 := GenS5{F0: (!false), F1: 18446744073709551615, F2: -1}
                                    v51 := []int{(fi49 << (((^(v44))) & 63))}
                                    fmt.Println(v44, fi49)
                                    _ = fi49
                                    _ = v50
                                    _ = v51
                                }
                            }
                        } else {
                            fmt.Println(fi26, v44)
                            for fi52 := 0; fi52 < 14; fi52++ {
                                fmt.Println(fi26, fi52)
                                v44 = (v44 >> ((2) & 63))
                                if (genFn0() != 32767) {
                                    var v53 uint16 = 42
                                    _ = v53
                                } else {
                                    v44 += v44
                                }
                                var v54 string = "\n\t"
                                v44 += (v44 * 2)
                                _ = fi52
                                _ = v54
                            }
                            fmt.Println(v44)
                        }
                        _ = v44
                    }
                    if true {
                        var v55 uint16 = 65535
                        var v56 int16 = -32768
                        if false {
                            fmt.Println(v56, fi26, fi26)
                            v57 := []int{fi26, (fi26 % ((fi26) | 1))}
                            _ = v57
                        }
                        _ = v55
                        _ = v56
                    } else {
                        var v58 float32 = 0.1
                        v58 = 0.1
                        for fi59 := 0; fi59 < 17; fi59++ {
                            v60 := GenS4{F0: 1, F1: 9223372036854775807}
                            v58 += v58
                            var v61 int32 = -2147483648
                            _ = fi59
                            _ = v60
                            _ = v61
                        }
                        _ = v58
                    }
                    if false {
                        v62 := []int{fi26, (fi26 / ((fi26) | 1))}
                        fmt.Println(fi26)
                        _ = v62
                    } else {
                        for fi63 := 0; fi63 < 13; fi63++ {
                            var v64 int32 = 1
                            _ = fi63
                            _ = v64
                        }
                    }
                }
            default:
                var v65 float64 = 3.14159
                _ = v65
            }
            switch 1 {
            case -1:
                if (-448 <= 1) {
                    switch "" {
                    case "a":
                        v66 := []int{-2147483648, 1}
                        _ = v66
                    default:
                        if ((("\n\t" + "a") + "foo bar") > "a") {
                            v67 := GenS3{F0: 685, F1: 65535}
                            _ = v67
                        } else {
                            var v68 uint8 = 0
                            var v69 float32 = 0.0
                            for fi70 := 0; fi70 < 18; fi70++ {
                                var v71 string = "a"
                                v71 = v71
                                v72 := GenS1{F0: 2147483647, F1: 1, F2: true}
                                var v73 int8 = 127
                                _ = fi70
                                _ = v71
                                _ = v72
                                _ = v73
                            }
                            _ = v68
                            _ = v69
                        }
                    }
                    var v74 uint = 0
                    _ = v74
                } else {
                    var v75 uint64 = 0
                    v76 := GenS3{F0: v75, F1: 0}
                    _ = v75
                    _ = v76
                }
            case -386:
                var v77 string = ""
                _ = v77
            }
            for fi78 := 0; fi78 < 9; fi78++ {
                v79 := GenS4{F0: 93, F1: -10}
                fmt.Println(fi78, fi26)
                _ = fi78
                _ = v79
            }
            _ = fi26
        }
    default:
        v80 := GenS1{F0: 690, F1: 65535, F2: false}
        _ = v80
    }
    v81 := GenS2{F0: 358}
    if (745 <= 352) {
        v82 := GenS1{F0: 398, F1: 255, F2: (1 >= 127)}
        v83 := []int{}
        var v84 int8 = -128
        _ = v82
        _ = v83
        _ = v84
    } else {
        for fi85 := 0; fi85 < 12; fi85++ {
            for fi86 := 0; fi86 < 9; fi86++ {
                var v87 int8 = 3
                var v88 float64 = ((1.5 * -2.25) - (0.0 - 0.0))
                v88 -= (v88 - (v88 - v88))
                var v89 float32 = (((-2.25 + (-2.25 * 3.14159)) - 1.5) * 1.5)
                _ = fi86
                _ = v87
                _ = v88
                _ = v89
            }
            if false {
                var v90 string = ("\n\t" + ("hello" + ("\n\t" + "a")))
                _ = v90
            } else {
                v91 := []int{}
                _ = v91
            }
            var v92 uint16 = 42
            _ = fi85
            _ = v92
        }
        var v93 uint64 = 964
        var v94 int16 = -1
        _ = v93
        _ = v94
    }
    v95 := GenS5{F0: (!(42 >= 2147483648)), F1: 2, F2: 1}
    _ = v81
    _ = v95
    return -94
}

func main() {
    var v1 uint32 = 1
    var v2 float32 = (0.0 / (((float32((0.0 / ((1.5 / 1.0) + -2.25))) - (0.1 - -2.25)) / 1.0) / 1.5))
    var v3 int16 = 1000
    var v4 int64 = 42
    var v5 int64 = (v4 & (-(v4)))
    for fi96 := 0; fi96 < 3; fi96++ {
        for fi97 := 0; fi97 < 2; fi97++ {
            v2 -= 1.5
            var v98 float64 = ((1e10 / (float64(((((1e10 - ((((v2 + 1e10) / v2) - v2) - (float32(0.1) / float32(-2.25)))) + 0.1) - float32(float64(v2))) + (((1e10 + (1e10 / v2)) / float32(3.14159)) + ((v2 - (v2 - v2)) + v2)))) - 0.0)) / (3.14159 - 1.5))
            v99 := []int{1, 32768}
            if (16 == -1) {
                var v100 int = 128
                _ = v100
            } else {
                if true {
                    v101 := GenS4{F0: 371, F1: 1}
                    _ = v101
                }
                v1 = 42
            }
            if false {
                v2 += 1.5
            } else {
                var v102 int64 = -241
                _ = v102
            }
            v4 += v4
            _ = fi97
            _ = v98
            _ = v99
        }
        _ = fi96
    }
    if true {
        switch 127 {
        case -785:
            for fi103 := 0; fi103 < 18; fi103++ {
                switch (fi103 & fi103) {
                case 32767:
                    v104 := GenS3{F0: 994, F1: (^(v1))}
                    _ = v104
                default:
                    switch "a" {
                    case "\n\t":
                        v105 := GenS6{F0: 42}
                        _ = v105
                    default:
                        for fi106 := 0; fi106 < 13; fi106++ {
                            var v107 uint = 255
                            fmt.Println(v3, v2)
                            _ = fi106
                            _ = v107
                        }
                    }
                }
                v108 := GenS5{F0: ((^(v3)) != v3), F1: 2, F2: 27}
                v109 := GenS3{F0: 4294967295, F1: v1}
                _ = fi103
                _ = v108
                _ = v109
            }
        }
        v110 := []int{32768}
        v3 += v3
        _ = v110
    } else {
        v111 := []int{128}
        v112 := GenS7{F0: true, F1: 0, F2: 9223372036854775807}
        _ = v111
        _ = v112
    }
    var v113 int16 = v3
    var v114 float64 = (0.0 / 3.14159)
    if true {
        var v115 uint16 = 255
        _ = v115
    }
    v116 := []int{42}
    if true {
        var v117 uint16 = 65535
        if ((^(v3)) != (v3 & genFn0())) {
            v118 := GenS8{F0: 128}
            _ = v118
        } else {
            fmt.Println(v4, v2)
            if len(v116) >= 2 {
                v116 = v116[1:len(v116)]
            }
            if true {
                var v119 float64 = 0.0
                _ = v119
            } else {
                v2 = ((v2 + (((v2 - 3.14159) + v2) - (v2 + v2))) - v2)
            }
        }
        switch -2147483648 {
        case -1:
            v114 += ((v114 / (v114 * v114)) / (((v114 - v114) * 0.1) + (float64(0.0) * ((1.5 * v114) - v114))))
        case -617:
            v117 = (v117 / ((v117) | 1))
        default:
            v1 = 0
        }
        _ = v117
    } else {
        v120 := GenS7{F0: false, F1: 0, F2: 4294967295}
        v121 := GenS9{F0: v4, F1: 1}
        v122 := GenS7{F0: ((-2.25 - float64(v2)) != 0.0), F1: 42, F2: -2147483648}
        _ = v120
        _ = v121
        _ = v122
    }
    fmt.Println(v1, v5, v1)
    v113 = genFn0()
    fmt.Println("v1:", v1)
    fmt.Println("v113:", v113)
    fmt.Println("v114:", v114)
    fmt.Println("slicelen v116:", len(v116), v116)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v113
    _ = v114
    _ = v116
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
