package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint8; F1 int32 }
type GenS5 struct { F0 int64; F1 uint64 }
type GenS6 struct { F0 int32; F1 int16 }
type GenS7 struct { F0 uint64; F1 int16 }
type GenS8 struct { F0 uint64; F1 uint; F2 int32 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 uint64; F1 int64 }
type GenS12 struct { F0 int32; F1 string; F2 int32 }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 uint32; F1 int16; F2 int64 }
type GenS15 struct { F0 int }
type GenS16 struct { F0 int32; F1 uint64; F2 int }
func genFn0(p0 uint32, p1 uint32, p2 string) uint {
    var v5 int = 2147483647
    p1 -= p0
    v6 := GenS1{F0: ((1 < 32767) || (!(0 > -1))), F1: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 255
}

func main() {
    var v1 float64 = 0.1
    var v2 int16 = 0
    var v3 uint32 = 65535
    v4 := make(map[int]int)
    if (((0.0 / 1.0) <= (1e10 - (3.14159 - (1.5 * 3.14159)))) || (v3 > (v3 ^ (v3 - (v3 - (v3 ^ (v3 << ((18446744073709551615) & 63)))))))) {
        for fi7 := 0; fi7 < 16; fi7++ {
            switch "a" {
            case "foo bar":
                var v8 string = "hello"
                _ = v8
            case "\n\t":
                if false {
                    fmt.Println(v2)
                    fmt.Println(fi7, v2)
                } else {
                    v4[728] = (fi7 / ((fi7) | 1))
                }
            case "a":
                if (true && false) {
                    v2 += v2
                    if false {
                        _ = len(v4)
                    } else {
                        var v9 string = ""
                        var v10 bool = (fi7 == 2147483647)
                        _ = v10
                        _ = v9
                    }
                }
            default:
                v2 -= 32767
            }
            if (v3 >= (v3 >> ((65535) & 63))) {
                for fi11 := 0; fi11 < 9; fi11++ {
                    v3 -= v3
                    var v12 bool = (true || false)
                    _ = fi11
                    _ = v12
                }
            } else {
                v13 := GenS2{F0: (("" + "hello") + ""), F1: 255}
                v14 := GenS1{F0: false, F1: 204}
                _ = v13
                _ = v14
            }
            _ = fi7
        }
        delete(v4, 127)
    }
    if (42 == 2) {
        v15 := GenS2{F0: "\n\t", F1: 32767}
        var v16 uint64 = 0
        if (631 > genFn0((v3 & (v3 | (v3 & v3))), 918, ("" + ""))) {
            if (false && (!(true || (false || false)))) {
                v2 -= int16(v3)
                switch 127 {
                case 128:
                    fmt.Println(v1, v2)
                }
                var v17 int16 = 1000
                _ = v17
            } else {
                var v18 uint = genFn0(v3, v3, "\n\t")
                var v19 string = ("hello" + "")
                if (0.0 >= float64(1e10)) {
                    var v20 string = v19
                    var v21 uint16 = 623
                    for fi22 := 0; fi22 < 4; fi22++ {
                        var v23 int = fi22
                        _ = fi22
                        _ = v23
                    }
                    _ = v20
                    _ = v21
                }
                _ = v18
                _ = v19
            }
            v2 -= v2
            if true {
                for fi24 := 0; fi24 < 20; fi24++ {
                    if false {
                        v3 += v3
                        v25 := GenS3{F0: "a"}
                        _ = v25
                    } else {
                        v26 := GenS4{F0: 7, F1: 42}
                        v2 += (v2 << ((42) & 63))
                        _ = v26
                    }
                    if true {
                        var v27 uint64 = (v16 ^ (^(v16)))
                        _ = v27
                    } else {
                        var v28 uint32 = v3
                        _ = v28
                    }
                    for fi29 := 0; fi29 < 15; fi29++ {
                        v16 -= (^(v16))
                        _ = fi29
                    }
                    _ = fi24
                }
                fmt.Println(v1, v16)
            } else {
                fmt.Println(v3, v16, v3)
                v16 = (^(v16))
                v30 := GenS5{F0: 0, F1: 113}
                _ = v30
            }
        } else {
            v3 += (v3 * (v3 + (v3 << ((562) & 63))))
            var v31 float32 = 0.0
            _ = v31
        }
        _ = v15
        _ = v16
    } else {
        v32 := GenS4{F0: 127, F1: 0}
        var v33 int8 = 1
        v3 = 427
        _ = v32
        _ = v33
    }
    if (320 <= 0) {
        switch "hello" {
        case "a":
            v3 = 2147483647
        case "":
            var v34 string = ""
            _ = v34
        default:
            v35 := GenS3{F0: "a"}
            _ = v35
        }
        fmt.Println(v1, v3, v3)
        fmt.Println(v1, v2)
    } else {
        fmt.Println(v1, v1, v2)
        switch "hello" {
        case "a":
            v3 = 0
        }
    }
    v36 := GenS6{F0: -443, F1: (v2 + (v2 ^ v2))}
    v37 := GenS7{F0: 4294967295, F1: v2}
    for fi38 := 0; fi38 < 12; fi38++ {
        if false {
            if (!((v3 / (((v3 ^ (^(v3)))) | 1)) >= (v3 & (v3 + v3)))) {
                v2 -= v2
            }
        }
        v39 := GenS3{F0: ("hello" + "foo bar")}
        v40 := GenS5{F0: -939, F1: 1}
        _ = fi38
        _ = v39
        _ = v40
    }
    var v41 int32 = 2147483647
    switch -390 {
    case 2147483647:
        v42 := []int{4294967295, 127, 1}
        _ = v42
    case 32767:
        if (v2 == v2) {
            for fi43 := 0; fi43 < 11; fi43++ {
                v44 := []int{32767, (fi43 << ((255) & 63)), fi43}
                _ = fi43
                _ = v44
            }
            for fi45 := 0; fi45 < 6; fi45++ {
                var v46 uint8 = 42
                v47 := GenS8{F0: 2, F1: 4294967295, F2: -1}
                fmt.Println(fi45, v41)
                _ = fi45
                _ = v46
                _ = v47
            }
        } else {
            v48 := []int{}
            switch 0 {
            case -9223372036854775808:
                v1 += float64(3.14159)
            case 32768:
                v2 += (-(v2))
            case 575:
                switch ("hello" + "") {
                case "\n\t":
                    v1 = v1
                case "foo bar":
                    v41 = 0
                }
            default:
                var v49 string = ""
                _ = v49
            }
            _ = v48
        }
    case 4294967295:
        if true {
            var v50 uint = 65535
            _ = v50
        } else {
            for fi51 := 0; fi51 < 12; fi51++ {
                var v52 uint8 = 128
                v53 := GenS9{F0: 42}
                for fi54 := 0; fi54 < 16; fi54++ {
                    if true {
                        v55 := GenS10{F0: 32767}
                        v56 := GenS8{F0: 2, F1: 1, F2: 868}
                        _ = v55
                        _ = v56
                    } else {
                        var v57 float64 = (v1 * 1.5)
                        if false {
                            fmt.Println(v52, v52, v41)
                            v58 := GenS9{F0: 808}
                            _ = v58
                        }
                        if (!((true || (((1e10 / (0.0 + (float32((v57 / float64(1.5))) - float32(3.14159)))) + float32((v1 + v1))) == ((1.5 - 0.1) + (0.0 / 0.1)))) || (true && (int8(v3) <= 96)))) {
                            switch "a" {
                            case "a":
                                v41 = -1
                            case "\n\t":
                                switch "a" {
                                case "hello":
                                    v59 := GenS6{F0: v41, F1: (v2 - v2)}
                                    _ = v59
                                case "foo bar":
                                    v60 := []int{fi51, (^(fi51)), 886}
                                    _ = v60
                                }
                            case "":
                                v61 := GenS3{F0: ((("\n\t" + "\n\t") + ("a" + "")) + "")}
                                if (v3 != (v3 % (((v3 / ((v3) | 1))) | 1))) {
                                    var v62 uint8 = 255
                                    _ = v62
                                }
                                _ = v61
                            }
                        } else {
                            switch ("hello" + "hello") {
                            case "hello":
                                fmt.Println(fi54)
                            default:
                                if false {
                                    v63 := []int{fi51}
                                    var v64 string = ""
                                    _ = v63
                                    _ = v64
                                } else {
                                    switch fi51 {
                                    case 9223372036854775807:
                                        var v65 int16 = (v2 + 451)
                                        _ = v65
                                    case 32767:
                                        _ = len(v4)
                                    }
                                }
                            }
                            var v66 uint = 255
                            var v67 uint = v66
                            _ = v66
                            _ = v67
                        }
                        _ = v57
                    }
                    v2 -= (v2 << ((116) & 63))
                    v41 = v41
                    _ = fi54
                }
                _ = fi51
                _ = v52
                _ = v53
            }
            var v68 bool = false
            _ = v68
        }
    default:
        v69 := GenS3{F0: "foo bar"}
        _ = v69
    }
    if true {
        if true {
            v1 -= 0.0
            var v70 int32 = (v41 >> ((genFn0(v3, v3, "")) & 63))
            v41 = int32(v3)
            _ = v70
        }
        switch ((("a" + "foo bar") + "foo bar") + "foo bar") {
        case "hello":
            for fi71 := 0; fi71 < 15; fi71++ {
                for fi72 := 0; fi72 < 14; fi72++ {
                    var v73 int8 = 127
                    if false {
                        switch fi72 {
                        case -369:
                            switch (fi71 | fi71) {
                            case -2147483648:
                                v3 = (v3 + (v3 - (v3 ^ 325)))
                            case -868:
                                v41 -= v41
                            }
                        case 32768:
                            v3 = 42
                        }
                        var v74 uint16 = 0
                        v75 := []int{(fi72 / (((fi72 >> ((88) & 63))) | 1))}
                        if (-1 == 0) {
                            var v76 uint64 = 2
                            var v77 float64 = float64(0.0)
                            v74 -= v74
                            _ = v76
                            _ = v77
                        } else {
                            v73 = v73
                            if true {
                                if len(v75) >= 2 {
                                    v75 = v75[1:len(v75)]
                                }
                            } else {
                                v41 += (v41 | (v41 ^ 867))
                            }
                            v73 += v73
                            var v78 int64 = 9223372036854775807
                            if (!false) {
                                var v79 string = ("\n\t" + (("\n\t" + ("hello" + "a")) + "hello"))
                                var v80 uint32 = (v3 * (v3 / (((v3 * (v3 << ((18446744073709551615) & 63)))) | 1)))
                                switch v79 {
                                case "foo bar":
                                    if (true || true) {
                                        if len(v75) >= 2 {
                                            v75 = v75[1:len(v75)]
                                        }
                                        v81 := GenS10{F0: v2}
                                        if (!(fi72 != fi72)) {
                                            if false {
                                                var v82 bool = ((float64(3.14159) + v1) < v1)
                                                _ = v82
                                            }
                                        }
                                        _ = v81
                                    }
                                case "a":
                                    if true {
                                        var v83 int16 = (v2 - v2)
                                        var v84 uint = 319
                                        _ = len(v75)
                                        _ = v83
                                        _ = v84
                                    } else {
                                        var v85 bool = true
                                        v41 += (v41 - (v41 % ((v41) | 1)))
                                        if true {
                                            if (255 < 255) {
                                                v75 = append(v75, -937)
                                                v80 = 2147483648
                                            }
                                            if false {
                                                v85 = (255 > 2)
                                            } else {
                                                v86 := GenS5{F0: (v78 / ((42) | 1)), F1: 5}
                                                _ = v86
                                            }
                                            v87 := GenS11{F0: 2, F1: v78}
                                            v1 = -2.25
                                            _ = v87
                                        } else {
                                            if true {
                                                var v88 int = -2147483648
                                                if (v85 || v85) {
                                                    v80 = v3
                                                }
                                                _ = v88
                                            } else {
                                                if ((v74 / ((v74) | 1)) <= v74) {
                                                    v89 := GenS12{F0: 38, F1: (v79 + (v79 + v79)), F2: (v41 + v41)}
                                                    var v90 bool = v85
                                                    _ = v89
                                                    _ = v90
                                                }
                                                v1 = 0.1
                                            }
                                            v91 := GenS8{F0: 1, F1: 1, F2: v41}
                                            _ = v91
                                        }
                                        var v92 uint8 = 1
                                        v93 := GenS13{F0: -29}
                                        _ = v85
                                        _ = v92
                                        _ = v93
                                    }
                                default:
                                    var v94 uint16 = (v74 << ((2) & 63))
                                    _ = v94
                                }
                                _ = v79
                                _ = v80
                            } else {
                                v4[9223372036854775807] = fi71
                                if (v1 <= 1e10) {
                                    v95 := GenS2{F0: "a", F1: uint16(v3)}
                                    _ = v95
                                } else {
                                    fmt.Println(v2, v78, v73)
                                    delete(v4, 0)
                                }
                            }
                            _ = v78
                        }
                        if (2 != 42) {
                            v96 := GenS4{F0: 42, F1: (v41 % ((-1) | 1))}
                            switch "hello" {
                            case "hello":
                                var v97 int64 = 949
                                _ = v97
                            case "a":
                                switch "\n\t" {
                                case "\n\t":
                                    var v98 float64 = float64(0.0)
                                    _ = v98
                                case "foo bar":
                                    fmt.Println(v3)
                                }
                            }
                            if len(v75) > 0 {
                                v75[0] = (fi71 * fi72)
                            }
                            var v99 string = ("a" + "hello")
                            _ = v96
                            _ = v99
                        }
                        if true {
                            v4[fi71] = (-(fi71))
                            fmt.Println(v2)
                            v100 := GenS14{F0: (v3 | (v3 >> ((18446744073709551615) & 63))), F1: (v2 * v2), F2: -9223372036854775808}
                            _ = v100
                        } else {
                            v75 = append(v75, 9223372036854775807)
                            var v101 uint = 1
                            fmt.Println(v101, v1, v74)
                            _ = len(v75)
                            _ = v101
                        }
                        switch fi72 {
                        case 42:
                            var v102 float64 = (0.1 / v1)
                            _ = v102
                        case 32768:
                            var v103 float64 = float64(3.14159)
                            _ = v103
                        default:
                            v104 := GenS13{F0: (v41 * (^(v41)))}
                            _ = v104
                        }
                        _ = v74
                        _ = v75
                    } else {
                        v3 = 2147483647
                        var v105 int64 = 42
                        var v106 string = ""
                        v107 := GenS15{F0: 32768}
                        _ = v105
                        _ = v106
                        _ = v107
                    }
                    var v108 bool = (2 < 127)
                    v73 = (v73 & (v73 * int8(fi71)))
                    var v109 int32 = (v41 ^ (v41 >> ((0) & 63)))
                    _ = fi72
                    _ = v108
                    _ = v109
                    _ = v73
                }
                _ = fi71
            }
        default:
            if (!false) {
                var v110 uint32 = 42
                v111 := GenS14{F0: v110, F1: v2, F2: 42}
                var v112 uint = 4294967295
                _ = v110
                _ = v111
                _ = v112
            } else {
                fmt.Println(v3, v3, v1)
                v2 += v2
            }
        }
        switch "hello" {
        case "\n\t":
            for fi113 := 0; fi113 < 10; fi113++ {
                v114 := GenS14{F0: v3, F1: 1000, F2: 42}
                _ = fi113
                _ = v114
            }
        case "a":
            fmt.Println(v41)
        case "foo bar":
            v41 += v41
        }
    } else {
        var v115 int32 = (v41 >> ((2) & 63))
        fmt.Println(v1, v1, v1)
        v4[496] = -722
        _ = v115
    }
    for fi116 := 0; fi116 < 1; fi116++ {
        v117 := GenS16{F0: int32(v3), F1: 225, F2: fi116}
        _ = fi116
        _ = v117
    }
    if false {
        v3 = 138
    } else {
        var v118 uint = 42
        v1 -= v1
        var v119 int8 = -13
        _ = v118
        _ = v119
    }
    var v120 uint32 = (v3 & v3)
    v121 := GenS1{F0: (487 >= 2), F1: 1}
    for fi122 := 0; fi122 < 10; fi122++ {
        fmt.Println(v3)
        var v123 float64 = float64(float32(v1))
        _ = fi122
        _ = v123
    }
    var v124 bool = (42 > -128)
    v2 = v2
    if true {
        var v125 uint16 = 65535
        v124 = false
        _ = v125
    }
    fmt.Println("v1:", v1)
    fmt.Println("v120:", v120)
    fmt.Println("v124:", v124)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    _ = v1
    _ = v120
    _ = v121
    _ = v124
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v41
}
