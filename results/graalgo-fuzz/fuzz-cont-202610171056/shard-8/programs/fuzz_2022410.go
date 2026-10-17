package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 uint8 }
type GenS2 struct { F0 uint64; F1 uint16; F2 int16 }
type GenS3 struct { F0 uint32; F1 int64; F2 int8 }
type GenS4 struct { F0 string; F1 uint8 }
type GenS5 struct { F0 int32; F1 bool }
type GenS6 struct { F0 uint16; F1 int }
type GenS7 struct { F0 int16; F1 uint16 }
type GenS8 struct { F0 int16; F1 uint32; F2 uint32 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 int64; F1 int64; F2 uint }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 int16; F1 string }
type GenS13 struct { F0 int }
type GenS14 struct { F0 int8 }
func genFn0(p0 uint, p1 string, p2 uint16) bool {
    v6 := []int{-2147483648}
    var v7 float64 = 0.0
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return (32767 == 32767)
}
func genFn1(p0 uint64, p1 string, p2 int64) uint32 {
    var v8 int16 = 759
    if false {
        v9 := []int{}
        _ = v9
    } else {
        if true {
            var v10 int64 = 1
            if (3.14159 == (0.1 + 1.5)) {
                v11 := GenS1{F0: (v8 >> ((42) & 63)), F1: 0, F2: 255}
                var v12 int16 = (v8 - (v8 + v8))
                _ = v11
                _ = v12
            } else {
                fmt.Println(v8, v10, v10)
            }
            v8 = v8
            _ = v10
        } else {
            p2 = int64(p0)
            v13 := GenS1{F0: (v8 >> ((182) & 63)), F1: 0, F2: 3}
            _ = v13
        }
        for fi14 := 0; fi14 < 9; fi14++ {
            v15 := []int{(-(fi14)), 32768}
            _ = fi14
            _ = v15
        }
        var v16 int32 = 1
        _ = v16
    }
    p2 -= 0
    if (false || (750 <= 0)) {
        fmt.Println(p1, p1, p1)
        fmt.Println(v8)
    } else {
        for fi17 := 0; fi17 < 20; fi17++ {
            switch 128 {
            case -155:
                v18 := []int{fi17, int(p0), fi17}
                _ = v18
            case 128:
                if false {
                    if true {
                        for fi19 := 0; fi19 < 13; fi19++ {
                            if (!false) {
                                v20 := []int{(fi17 * 32768)}
                                var v21 int = 1
                                var v22 float64 = 0.0
                                _ = v20
                                _ = v21
                                _ = v22
                            }
                            v8 = -216
                            v8 += -32768
                            _ = fi19
                        }
                        var v23 int = int(p2)
                        var v24 int = (fi17 - fi17)
                        _ = v23
                        _ = v24
                    }
                    var v25 int16 = v8
                    _ = v25
                } else {
                    v26 := GenS2{F0: p0, F1: 32767, F2: (v8 & (-(v8)))}
                    v27 := []int{fi17, fi17, 32767}
                    _ = v26
                    _ = v27
                }
            case -359:
                v28 := GenS2{F0: 1, F1: 0, F2: (v8 + v8)}
                _ = v28
            }
            if (!(!false)) {
                v8 -= 963
                fmt.Println(v8, v8)
                v29 := GenS1{F0: v8, F1: -382, F2: 123}
                _ = v29
            } else {
                fmt.Println(p2, fi17, p0)
                v30 := GenS3{F0: 0, F1: int64(p0), F2: 100}
                _ = v30
            }
            p1 = p1
            _ = fi17
        }
        for fi31 := 0; fi31 < 15; fi31++ {
            var v32 uint32 = uint32(p0)
            var v33 uint32 = (v32 / (((^(v32))) | 1))
            var v34 float32 = 0.1
            _ = fi31
            _ = v32
            _ = v33
            _ = v34
        }
    }
    if true {
        if true {
            v35 := GenS3{F0: 2, F1: p2, F2: -128}
            p1 = (p1 + p1)
            var v36 int64 = (^(p2))
            _ = v35
            _ = v36
        } else {
            v37 := []int{}
            _ = v37
        }
    } else {
        v38 := []int{9223372036854775807}
        _ = v38
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 0
}

func main() {
    var v1 int32 = -2147483648
    var v2 int16 = 32767
    var v3 int32 = (v1 & -852)
    v4 := []uint32{2}
    var v5 string = "hello"
    v39 := GenS4{F0: v5, F1: 1}
    for fi40 := 0; fi40 < 5; fi40++ {
        v41 := GenS5{F0: v3, F1: (v5 <= "foo bar")}
        v5 = v5
        _ = fi40
        _ = v41
    }
    fmt.Println(v1)
    v42 := GenS1{F0: 1, F1: (v3 % ((v1) | 1)), F2: 128}
    if false {
        v43 := GenS6{F0: 2, F1: 566}
        _ = v43
    } else {
        var v44 int16 = -777
        v5 = v5
        var v45 int8 = 1
        _ = v44
        _ = v45
    }
    fmt.Println(v1, v2)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v46 string = ("foo bar" + ((v5 + "hello") + ("a" + "a")))
    for fi47 := 0; fi47 < 18; fi47++ {
        if true {
            var v48 uint8 = 42
            _ = v48
        } else {
            v3 -= v1
        }
        v4 = append(v4, 1)
        _ = fi47
    }
    v5 = v46
    var v49 float32 = 0.1
    if false {
        switch v5 {
        case "hello":
            switch 813 {
            case 9223372036854775807:
                var v50 int8 = 1
                _ = v50
            case 561:
                for fi51 := 0; fi51 < 13; fi51++ {
                    fmt.Println(v5, v5)
                    v2 -= v2
                    _ = fi51
                }
            case -1:
                var v52 uint16 = 2
                _ = v52
            }
        case "a":
            v53 := GenS7{F0: 0, F1: 1}
            _ = v53
        case "foo bar":
            var v54 uint64 = 304
            _ = v54
        default:
            if true {
                v55 := GenS7{F0: (v2 - v2), F1: 42}
                if true {
                    var v56 int16 = (-(v2))
                    var v57 int32 = v3
                    var v58 uint = 4294967295
                    _ = v56
                    _ = v57
                    _ = v58
                } else {
                    for fi59 := 0; fi59 < 10; fi59++ {
                        switch 144 {
                        case 405:
                            if (2 >= 18446744073709551615) {
                                for fi60 := 0; fi60 < 16; fi60++ {
                                    var v61 string = ("" + ("\n\t" + ""))
                                    var v62 int8 = -126
                                    _ = fi60
                                    _ = v61
                                    _ = v62
                                }
                                for fi63 := 0; fi63 < 7; fi63++ {
                                    if false {
                                        v64 := GenS1{F0: (v2 - (v2 << ((0) & 63))), F1: (v1 + -1), F2: 58}
                                        v3 = (^(v3))
                                        _ = v64
                                    } else {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        v65 := GenS6{F0: 255, F1: (fi59 - fi59)}
                                        if false {
                                            fmt.Println(fi59)
                                            v49 += v49
                                            var v66 string = v5
                                            _ = v66
                                        } else {
                                            fmt.Println(v2, v46)
                                            switch 128 {
                                            case -9223372036854775808:
                                                _ = len(v4)
                                                if false {
                                                    if len(v4) > 0 {
                                                        v4[0] = genFn1(855, ((v46 + ("foo bar" + "a")) + (("" + "foo bar") + v46)), int64(fi59))
                                                    }
                                                } else {
                                                    v2 = -559
                                                }
                                                v46 = (v46 + ((((v5 + "a") + v46) + "foo bar") + "hello"))
                                            case 0:
                                                switch -861 {
                                                case 32767:
                                                    var v67 uint8 = 2
                                                    _ = v67
                                                }
                                            case 42:
                                                v1 = -1
                                            default:
                                                fmt.Println(v3, v3)
                                                v2 = (v2 ^ v2)
                                                switch v5 {
                                                case "\n\t":
                                                    v46 = v5
                                                }
                                            }
                                        }
                                        v68 := GenS4{F0: "foo bar", F1: 7}
                                        var v69 float64 = (0.0 / 1e10)
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        _ = v65
                                        _ = v68
                                        _ = v69
                                    }
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    _ = fi63
                                }
                                v70 := GenS5{F0: (^(v1)), F1: true}
                                _ = v70
                            } else {
                                v2 -= 32767
                                fmt.Println(v5, fi59, v1)
                            }
                        case -9223372036854775808:
                            var v71 string = v46
                            _ = v71
                        case 0:
                            if ((false || true) && (((1 < 981) || false) || true)) {
                                var v72 int32 = (v3 << ((913) & 63))
                                switch fi59 {
                                case 207:
                                    var v73 float32 = 1e10
                                    _ = v73
                                case 180:
                                    v3 += v1
                                default:
                                    if true {
                                        var v74 int = 4294967295
                                        _ = v74
                                    } else {
                                        v75 := GenS5{F0: v3, F1: false}
                                        _ = v75
                                    }
                                }
                                _ = v72
                            } else {
                                if len(v4) > 0 {
                                    v4[0] = 65535
                                }
                                v76 := GenS4{F0: (v5 + (("hello" + v46) + ("foo bar" + ""))), F1: 1}
                                fmt.Println(v1)
                                _ = v76
                            }
                        }
                        var v77 int16 = (^(v2))
                        var v78 int64 = 1
                        _ = fi59
                        _ = v77
                        _ = v78
                    }
                    for fi79 := 0; fi79 < 1; fi79++ {
                        if (255 >= 2) {
                            if (1 > 1) {
                                if false {
                                    if (42 == 42) {
                                        v5 = v5
                                        var v80 float32 = 3.14159
                                        var v81 int16 = (^(v2))
                                        _ = v80
                                        _ = v81
                                    }
                                    v49 += v49
                                    fmt.Println(fi79)
                                }
                                for fi82 := 0; fi82 < 17; fi82++ {
                                    v46 = "hello"
                                    var v83 uint64 = 2
                                    var v84 int32 = v1
                                    _ = fi82
                                    _ = v83
                                    _ = v84
                                }
                            } else {
                                for fi85 := 0; fi85 < 14; fi85++ {
                                    v86 := GenS8{F0: int16(fi85), F1: 554, F2: 1}
                                    fmt.Println(v1, v2, fi79)
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    switch (-(fi85)) {
                                    case -2147483648:
                                        var v87 float64 = (float64(1.5) / (3.14159 - 1.5))
                                        _ = v87
                                    default:
                                        var v88 uint64 = 638
                                        _ = v88
                                    }
                                    _ = fi85
                                    _ = v86
                                }
                                v3 = v1
                            }
                            v89 := GenS2{F0: 787, F1: 65535, F2: -1}
                            var v90 string = ("" + v5)
                            _ = v89
                            _ = v90
                        }
                        for fi91 := 0; fi91 < 19; fi91++ {
                            switch (fi79 - (fi79 / (((-(fi79))) | 1))) {
                            case -2147483648:
                                var v92 string = (v5 + "\n\t")
                                _ = v92
                            case 1:
                                v1 = (v3 / ((v3) | 1))
                            case -918:
                                v1 = (v1 / ((42) | 1))
                            default:
                                var v93 string = ""
                                _ = v93
                            }
                            if (!((!true) || (!false))) {
                                if false {
                                    v94 := GenS9{F0: -9223372036854775808}
                                    v95 := GenS4{F0: "", F1: 1}
                                    _ = v94
                                    _ = v95
                                } else {
                                    v96 := GenS10{F0: 1, F1: 42, F2: 65535}
                                    _ = v96
                                }
                                if true {
                                    var v97 int64 = -192
                                    v5 = ("" + (v46 + "\n\t"))
                                    var v98 uint8 = 2
                                    _ = v97
                                    _ = v98
                                } else {
                                    if (!true) {
                                        var v99 int16 = (v2 + -712)
                                        _ = v99
                                    }
                                    var v100 uint16 = 2
                                    var v101 float32 = 3.14159
                                    _ = v100
                                    _ = v101
                                }
                                if (!true) {
                                    v46 = v5
                                }
                            }
                            _ = fi91
                        }
                        _ = fi79
                    }
                }
                _ = len(v4)
                _ = v55
            }
        }
    } else {
        for fi102 := 0; fi102 < 6; fi102++ {
            if false {
                var v103 uint8 = 0
                v104 := GenS11{F0: 18446744073709551615}
                v46 = (("\n\t" + (("foo bar" + (v5 + "a")) + ((v46 + ((("a" + "hello") + v46) + "hello")) + "\n\t"))) + (("a" + "hello") + v46))
                _ = v103
                _ = v104
            }
            var v105 uint = 4294967295
            v106 := GenS1{F0: 32767, F1: v3, F2: 45}
            _ = fi102
            _ = v105
            _ = v106
        }
        var v107 int64 = 364
        _ = v107
    }
    switch v5 {
    case "a":
        for fi108 := 0; fi108 < 16; fi108++ {
            v1 = (^(v1))
            if (false && (!false)) {
                _ = len(v4)
                for fi109 := 0; fi109 < 15; fi109++ {
                    var v110 string = v5
                    switch (v46 + "hello") {
                    case "hello":
                        if (889 < 65535) {
                            switch fi109 {
                            case 127:
                                var v111 uint64 = 1
                                _ = v111
                            case 445:
                                if false {
                                    var v112 int = (fi108 >> ((905) & 63))
                                    var v113 uint16 = 255
                                    fmt.Println(fi108, fi109, v46)
                                    _ = v112
                                    _ = v113
                                } else {
                                    v1 = (^(v1))
                                    var v114 int8 = 1
                                    var v115 bool = ((v3 + (v1 | v3)) >= (v1 % ((int32(fi108)) | 1)))
                                    fmt.Println(v2, v2, fi109)
                                    if len(v4) > 0 {
                                        v4[0] = 1
                                    }
                                    v116 := GenS5{F0: 0, F1: v115}
                                    if (505 >= v3) {
                                        var v117 int8 = v114
                                        v2 = v2
                                        _ = v117
                                    }
                                    var v118 float64 = 1.5
                                    _ = v114
                                    _ = v115
                                    _ = v116
                                    _ = v118
                                }
                            default:
                                fmt.Println(v1, v49)
                                v119 := GenS12{F0: -1, F1: v110}
                                _ = v119
                            }
                            _ = len(v4)
                        } else {
                            _ = len(v4)
                            if false {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                            var v120 float64 = (1e10 * -2.25)
                            _ = v120
                        }
                    case "\n\t":
                        v110 = ("foo bar" + "foo bar")
                    case "foo bar":
                        if true {
                            v5 = v5
                            fmt.Println(v49, fi108)
                        } else {
                            switch "foo bar" {
                            case "a":
                                if false {
                                    if len(v4) > 0 {
                                        v4[0] = 65535
                                    }
                                } else {
                                    v121 := GenS4{F0: v5, F1: 223}
                                    _ = v121
                                }
                            case "foo bar":
                                switch "foo bar" {
                                case "hello":
                                    v122 := GenS13{F0: (fi109 + -613)}
                                    _ = v122
                                }
                            }
                            _ = len(v4)
                            var v123 int8 = -128
                            _ = v123
                        }
                    default:
                        v5 = ("foo bar" + (v46 + v46))
                    }
                    _ = fi109
                    _ = v110
                }
                var v124 int64 = -600
                _ = v124
            }
            v125 := GenS14{F0: 100}
            _ = fi108
            _ = v125
        }
    default:
        if (!false) {
            v126 := GenS11{F0: 0}
            fmt.Println(v3)
            _ = v126
        } else {
            var v127 uint64 = 877
            _ = v127
        }
    }
    if false {
        if (1 == -3) {
            fmt.Println(v1)
        } else {
            var v128 float64 = ((3.14159 / float64(v49)) - 1e10)
            _ = v128
        }
    }
    v3 += -1
    for fi129 := 0; fi129 < 19; fi129++ {
        for fi130 := 0; fi130 < 3; fi130++ {
            if (!(false || true)) {
                var v131 string = v5
                _ = v131
            }
            v46 = v46
            var v132 int8 = int8(v2)
            _ = fi130
            _ = v132
        }
        _ = fi129
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v46:", v46)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v42
    _ = v46
    _ = v49
    _ = v5
}
