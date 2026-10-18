package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 uint8 }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 int; F1 uint8; F2 uint64 }
type GenS4 struct { F0 uint; F1 bool }
type GenS5 struct { F0 bool; F1 int16 }
type GenS6 struct { F0 int32; F1 uint }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint32; F1 uint32 }
func genFn0() uint {
    for fi5 := 0; fi5 < 7; fi5++ {
        if (!(0 < 42)) {
            if true {
                var v6 int16 = 1000
                for fi7 := 0; fi7 < 12; fi7++ {
                    v8 := GenS1{F0: 128, F1: false, F2: 0}
                    _ = fi7
                    _ = v8
                }
                v9 := []int{fi5, (fi5 - fi5), fi5}
                _ = v6
                _ = v9
            } else {
                v10 := GenS2{F0: 654, F1: 255}
                _ = v10
            }
            if false {
                v11 := GenS2{F0: 572, F1: 231}
                for fi12 := 0; fi12 < 19; fi12++ {
                    if (-225 == 0) {
                        if false {
                            if (!false) {
                                var v13 uint32 = 706
                                fmt.Println(fi5)
                                var v14 int8 = 127
                                _ = v13
                                _ = v14
                            } else {
                                if (-9223372036854775808 == 42) {
                                    if false {
                                        var v15 uint32 = 1
                                        v15 = 0
                                        _ = v15
                                    } else {
                                        var v16 int = (fi12 >> ((451) & 63))
                                        if (!true) {
                                            var v17 string = ("" + "")
                                            _ = v17
                                        }
                                        _ = v16
                                    }
                                } else {
                                    var v18 uint32 = 1
                                    fmt.Println(fi12, v18)
                                    v19 := GenS1{F0: fi5, F1: false, F2: 1}
                                    v20 := []int{}
                                    _ = v18
                                    _ = v19
                                    _ = v20
                                }
                                var v21 float32 = (0.0 - 1.5)
                                v22 := GenS1{F0: fi5, F1: ((false && true) || true), F2: 207}
                                _ = v21
                                _ = v22
                            }
                            v23 := []int{}
                            var v24 uint32 = 403
                            _ = v23
                            _ = v24
                        }
                        var v25 uint16 = 2
                        if false {
                            var v26 uint8 = 49
                            var v27 int16 = 32767
                            _ = v26
                            _ = v27
                        } else {
                            v25 += 1
                            if false {
                                v25 -= (v25 / ((255) | 1))
                                fmt.Println(fi5, v25, v25)
                            }
                        }
                        switch ("\n\t" + ("" + "foo bar")) {
                        case "foo bar":
                            var v28 uint8 = 189
                            _ = v28
                        case "hello":
                            v29 := []int{fi12, (fi12 | (fi5 & (fi5 & fi12))), 948}
                            _ = v29
                        default:
                            v30 := GenS2{F0: 0, F1: 128}
                            _ = v30
                        }
                        _ = v25
                    } else {
                        if (!(103 >= 124)) {
                            var v31 uint32 = 65535
                            _ = v31
                        }
                        v32 := GenS3{F0: -429, F1: 67, F2: 328}
                        var v33 float64 = (0.1 / (1.5 * (0.0 - (3.14159 - ((1.5 * 1e10) + (3.14159 - 1e10))))))
                        _ = v32
                        _ = v33
                    }
                    _ = fi12
                }
                if ((-2.25 / 1e10) == 0.0) {
                    v34 := GenS2{F0: 832, F1: 255}
                    fmt.Println(fi5)
                    _ = v34
                } else {
                    v35 := GenS2{F0: 2147483647, F1: 2}
                    _ = v35
                }
                _ = v11
            } else {
                var v36 uint8 = 128
                _ = v36
            }
            if (true || false) {
                var v37 uint16 = 890
                _ = v37
            } else {
                v38 := []int{}
                if false {
                    var v39 int16 = -26
                    if ((0.1 < 1.5) && (-1 <= -14)) {
                        v38 = append(v38, (fi5 - 2147483647))
                        if false {
                            var v40 bool = (87 == -128)
                            var v41 uint32 = 12
                            _ = v40
                            _ = v41
                        } else {
                            switch ("hello" + "a") {
                            case "a":
                                _ = len(v38)
                            case "":
                                for fi42 := 0; fi42 < 18; fi42++ {
                                    if len(v38) >= 2 {
                                        v38 = v38[1:len(v38)]
                                    }
                                    fmt.Println(fi5, fi5, fi5)
                                    _ = fi42
                                }
                            }
                        }
                        v39 = v39
                    } else {
                        _ = len(v38)
                        v43 := GenS3{F0: fi5, F1: 0, F2: 0}
                        var v44 int16 = v39
                        _ = v43
                        _ = v44
                    }
                    _ = v39
                }
                var v45 int64 = 9223372036854775807
                _ = v38
                _ = v45
            }
        }
        var v46 int64 = 42
        fmt.Println(v46, fi5)
        _ = fi5
        _ = v46
    }
    var v47 string = ((("foo bar" + "foo bar") + "a") + ("a" + ("\n\t" + ("\n\t" + "\n\t"))))
    v48 := GenS2{F0: 0, F1: 0}
    var v49 bool = ((-80 == -1) || (!(42 < 0)))
    _ = v47
    _ = v48
    _ = v49
    return 55
}
func genFn1(p0 uint32, p1 int32, p2 int32) uint32 {
    switch 42 {
    case 826:
        v50 := []int{}
        _ = v50
    case -49:
        v51 := GenS3{F0: 9223372036854775807, F1: 0, F2: 597}
        _ = v51
    }
    v52 := []int{-20, -9223372036854775808}
    var v53 int16 = 736
    fmt.Println(p0)
    if len(v52) > 0 {
        v52[0] = 2147483647
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v52
    _ = v53
    return 1
}

func main() {
    var v1 uint64 = 253
    var v2 uint = 133
    v3 := make(map[uint32]uint32)
    v4 := make(map[uint32]bool)
    switch "hello" {
    case "":
        if true {
            for fi54 := 0; fi54 < 20; fi54++ {
                v55 := []int{fi54, -1, 32767, fi54}
                _ = fi54
                _ = v55
            }
            v4[genFn1(2147483648, -1, 0)] = false
            if true {
                for fi56 := 0; fi56 < 7; fi56++ {
                    if false {
                        fmt.Println(fi56)
                        if (4294967295 != 2147483648) {
                            switch "" {
                            case "\n\t":
                                v57 := []int{32768, fi56, fi56, 737}
                                _ = v57
                            case "foo bar":
                                v58 := []int{-9223372036854775808, (fi56 ^ int(v2)), (-(fi56)), (fi56 ^ fi56)}
                                _ = v58
                            case "":
                                var v59 int32 = 1
                                _ = v59
                            }
                            v60 := []int{1, (fi56 & fi56), (-(fi56)), fi56}
                            _ = v60
                        } else {
                            if true {
                                v1 = 4294967295
                                v2 -= genFn0()
                            } else {
                                v61 := []int{fi56}
                                _ = v61
                            }
                            switch (fi56 / ((fi56) | 1)) {
                            case 2147483647:
                                if true {
                                    if false {
                                        var v62 float64 = float64(3.14159)
                                        for fi63 := 0; fi63 < 12; fi63++ {
                                            switch "\n\t" {
                                            case "hello":
                                                if (0 < int8(fi56)) {
                                                    v1 = v1
                                                } else {
                                                    v64 := v4[865]
                                                    v2 = 255
                                                    _ = v64
                                                }
                                            case "\n\t":
                                                var v65 bool = (!false)
                                                _ = v65
                                            case "foo bar":
                                                fmt.Println(v2, v2, v62)
                                            }
                                            _ = fi63
                                        }
                                        _ = len(v4)
                                        _ = v62
                                    } else {
                                        v1 = (^(v1))
                                        v66 := GenS1{F0: (^(fi56)), F1: (42 == -1), F2: 127}
                                        switch (fi56 % (((fi56 << ((genFn0()) & 63))) | 1)) {
                                        case -9223372036854775808:
                                            v67 := GenS2{F0: 2147483648, F1: 128}
                                            _ = v67
                                        case -2147483648:
                                            var v68 bool = (128 <= 97)
                                            _ = v68
                                        default:
                                            var v69 string = "\n\t"
                                            _ = v69
                                        }
                                        _ = v66
                                    }
                                    for fi70 := 0; fi70 < 9; fi70++ {
                                        var v71 int = -884
                                        _ = fi70
                                        _ = v71
                                    }
                                    for fi72 := 0; fi72 < 5; fi72++ {
                                        if (false || ((661 <= -940) && (-29 != 0))) {
                                            var v73 int32 = -434
                                            v74 := GenS4{F0: 65535, F1: ((^(v2)) >= (^(v2)))}
                                            _ = v73
                                            _ = v74
                                        } else {
                                            if false {
                                                var v75 int64 = 259
                                                _ = v75
                                            }
                                            switch "" {
                                            case "a":
                                                if ((((3.14159 / 3.14159) - (0.0 * 3.14159)) - -2.25) != 3.14159) {
                                                    if (fi72 > (fi72 / ((81) | 1))) {
                                                        v1 = (v1 % ((v1) | 1))
                                                    } else {
                                                        var v76 uint = v2
                                                        var v77 bool = (((1.5 - 0.1) * ((((0.0 / 1e10) / (float64(1e10) + (0.0 - float64(0.1)))) + (1e10 * (0.0 + (float64(0.1) * 1.5)))) * ((1e10 * 0.0) - 1.5))) == ((0.0 + 1e10) / 1e10))
                                                        _ = v76
                                                        _ = v77
                                                    }
                                                    var v78 bool = (false && true)
                                                    _ = v78
                                                }
                                            default:
                                                v2 = 161
                                                var v79 int8 = -1
                                                var v80 uint32 = genFn1(uint32(v1), -2147483648, 1)
                                                v81 := GenS4{F0: (v2 ^ (v2 / ((uint(v1)) | 1))), F1: false}
                                                _ = v79
                                                _ = v80
                                                _ = v81
                                            }
                                        }
                                        _ = fi72
                                    }
                                }
                            }
                            var v82 int32 = 0
                            _ = v82
                        }
                    } else {
                        v2 -= genFn0()
                        v2 -= genFn0()
                    }
                    _ = fi56
                }
                if ("a" == "") {
                    v2 -= 524
                    for fi83 := 0; fi83 < 14; fi83++ {
                        for fi84 := 0; fi84 < 15; fi84++ {
                            var v85 int = 128
                            _ = fi84
                            _ = v85
                        }
                        _ = fi83
                    }
                }
            } else {
                var v86 int = 607
                _ = v86
            }
        }
    case "hello":
        var v87 uint = v2
        _ = v87
    default:
        var v88 int64 = int64(v2)
        _ = v88
    }
    if (2147483647 > -2147483648) {
        var v89 int = 32767
        fmt.Println(v1, v89)
        var v90 uint = (v2 / (((v2 | uint(v1))) | 1))
        _ = v89
        _ = v90
    }
    if true {
        v2 -= v2
        for fi91 := 0; fi91 < 2; fi91++ {
            for fi92 := 0; fi92 < 12; fi92++ {
                if (float32((0.1 / 1.0)) >= (((1.5 * -2.25) - (1e10 / float32(1.5))) / 1.5)) {
                    var v93 uint8 = 169
                    _ = v93
                }
                _ = fi92
            }
            v1 = (^(v1))
            if (genFn1(1, -2147483648, 1) >= 1) {
                var v94 int32 = 0
                var v95 uint8 = uint8(v2)
                v96 := v4[42]
                _ = v94
                _ = v95
                _ = v96
            }
            _ = fi91
        }
        fmt.Println(v2)
    }
    v2 += genFn0()
    var v97 int8 = 0
    v98 := GenS5{F0: ("\n\t" >= ("hello" + ("hello" + "a"))), F1: 1}
    switch 127 {
    case 32767:
        v99 := GenS6{F0: -1, F1: (v2 - (v2 << ((v2) & 63)))}
        _ = v99
    case 42:
        for fi100 := 0; fi100 < 12; fi100++ {
            fmt.Println(v97)
            v101 := []int{4294967295, fi100, fi100, (fi100 | 891)}
            var v102 uint8 = 127
            _ = fi100
            _ = v101
            _ = v102
        }
    case -1:
        v103 := []int{418, -2147483648, 128}
        _ = v103
    default:
        var v104 int8 = int8(v1)
        _ = v104
    }
    v1 = v1
    for fi105 := 0; fi105 < 14; fi105++ {
        var v106 uint32 = 42
        switch fi105 {
        case 48:
            v107 := GenS1{F0: fi105, F1: false, F2: 2}
            _ = v107
        case 9223372036854775807:
            var v108 uint32 = (v106 % ((uint32(v2)) | 1))
            _ = v108
        case 835:
            v1 += v1
        }
        _ = fi105
        _ = v106
    }
    v109 := []int{538, -351, -143}
    for fi110 := 0; fi110 < 18; fi110++ {
        v1 = uint64(v2)
        var v111 int64 = 356
        _ = fi110
        _ = v111
    }
    switch ("" + ("a" + "")) {
    case "\n\t":
        v1 = 42
    case "":
        fmt.Println(v1)
    default:
        v112 := v4[2147483647]
        _ = v112
    }
    v2 = genFn0()
    switch -9223372036854775808 {
    case -9223372036854775808:
        fmt.Println(v97, v2, v1)
    case -521:
        var v113 uint16 = 2
        _ = v113
    case 4294967295:
        var v114 float32 = (0.1 - 3.14159)
        _ = v114
    }
    if true {
        v115 := GenS4{F0: (v2 / (((^(v2))) | 1)), F1: (v97 != v97)}
        var v116 float64 = (-2.25 / 1.5)
        _ = v115
        _ = v116
    } else {
        v117 := GenS7{F0: 42}
        _ = v117
    }
    v3[2] = 483
    if false {
        switch 1 {
        case 9223372036854775807:
            v97 = v97
        case 1:
            v118 := GenS8{F0: uint32(v2), F1: genFn1(1, -1, 42)}
            _ = v118
        case 32767:
            v2 += v2
        default:
            var v119 int64 = 1
            _ = v119
        }
        var v120 float64 = 3.14159
        _ = v120
    } else {
        var v121 uint64 = (v1 - (v1 * 30))
        switch ("\n\t" + "foo bar") {
        case "foo bar":
            v122 := GenS1{F0: 42, F1: (2147483647 <= -2147483648), F2: 1}
            _ = v122
        case "hello":
            v1 = 0
        }
        _ = v121
    }
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v109:", len(v109), v109)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v97:", v97)
    _ = v1
    _ = v109
    _ = v2
    _ = v3
    _ = v4
    _ = v97
    _ = v98
}
