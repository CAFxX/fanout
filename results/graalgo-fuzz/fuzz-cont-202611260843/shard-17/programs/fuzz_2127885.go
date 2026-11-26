package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 uint; F1 uint64; F2 string }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint32; F1 int; F2 int }
type GenS5 struct { F0 uint64; F1 int64 }
type GenS6 struct { F0 uint; F1 uint8 }
type GenS7 struct { F0 int16; F1 uint16; F2 uint8 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint64; F1 bool }
func genFn0() uint {
    for fi6 := 0; fi6 < 18; fi6++ {
        var v7 int32 = 188
        v8 := []int{(fi6 | (^(fi6))), fi6}
        if (true && ((351 > 1) && (923 <= 0))) {
            v7 = v7
        } else {
            var v9 uint16 = 42
            switch fi6 {
            case 874:
                for fi10 := 0; fi10 < 13; fi10++ {
                    fmt.Println(v7, fi6)
                    _ = fi10
                }
            default:
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
            }
            _ = v9
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    switch "foo bar" {
    case "hello":
        var v11 uint = 974
        _ = v11
    case "\n\t":
        var v12 int = -793
        _ = v12
    case "a":
        if ((!((1 > 18446744073709551615) || (0 != 2147483647))) && false) {
            if true {
                var v13 uint64 = 2
                _ = v13
            }
            var v14 bool = (1 > 127)
            _ = v14
        }
    default:
        var v15 uint = 4294967295
        _ = v15
    }
    switch 32768 {
    case 32767:
        var v16 uint16 = 2
        _ = v16
    }
    return 4294967295
}
func genFn1(p0 int64, p1 int64, p2 int8) int32 {
    for fi17 := 0; fi17 < 17; fi17++ {
        if false {
            var v18 string = "hello"
            v19 := []int{}
            p0 += -9223372036854775808
            _ = v18
            _ = v19
        } else {
            switch fi17 {
            case 0:
                if (!(genFn0() >= 889)) {
                    switch ("a" + "hello") {
                    case "hello":
                        for fi20 := 0; fi20 < 1; fi20++ {
                            var v21 int16 = 0
                            _ = fi20
                            _ = v21
                        }
                    default:
                        var v22 uint32 = 4294967295
                        _ = v22
                    }
                }
            case -477:
                if true {
                    fmt.Println(fi17, fi17, fi17)
                    v23 := []int{775}
                    if len(v23) >= 2 {
                        v23 = v23[1:len(v23)]
                    }
                    _ = v23
                }
            default:
                var v24 int = (fi17 + 32767)
                _ = v24
            }
        }
        var v25 string = "foo bar"
        p0 = (-(p1))
        _ = fi17
        _ = v25
    }
    switch "hello" {
    case "hello":
        if ((-1 != 834) || (-450 >= -243)) {
            for fi26 := 0; fi26 < 11; fi26++ {
                v27 := GenS1{F0: 2147483648, F1: 34}
                switch 127 {
                case -185:
                    p2 = (^(p2))
                }
                v28 := GenS2{F0: 2, F1: 42, F2: (("a" + "foo bar") + ((("" + "foo bar") + (("\n\t" + ("\n\t" + "")) + "foo bar")) + "a"))}
                _ = fi26
                _ = v27
                _ = v28
            }
            switch "hello" {
            case "\n\t":
                for fi29 := 0; fi29 < 11; fi29++ {
                    v30 := []int{fi29, (fi29 / ((fi29) | 1)), 4294967295, 30}
                    v31 := GenS3{F0: -1}
                    _ = fi29
                    _ = v30
                    _ = v31
                }
            case "":
                var v32 int64 = (^(p0))
                _ = v32
            case "foo bar":
                v33 := GenS4{F0: 2147483648, F1: 152, F2: -618}
                _ = v33
            default:
                if true {
                    for fi34 := 0; fi34 < 11; fi34++ {
                        if (!(685 <= 42)) {
                            if ((-(p2)) <= (p2 | p2)) {
                                v35 := GenS2{F0: 348, F1: 42, F2: "foo bar"}
                                var v36 float64 = 3.14159
                                _ = v35
                                _ = v36
                            }
                        } else {
                            var v37 float32 = 0.1
                            p2 = (p2 >> ((65535) & 63))
                            _ = v37
                        }
                        var v38 int16 = 647
                        v39 := []int{32767, int(v38)}
                        _ = fi34
                        _ = v38
                        _ = v39
                    }
                }
            }
            switch -1 {
            case -1:
                p2 = p2
            case -227:
                var v40 int32 = 1
                _ = v40
            }
        } else {
            v41 := GenS2{F0: 1, F1: 2, F2: ""}
            _ = v41
        }
    default:
        v42 := GenS2{F0: 590, F1: 830, F2: "a"}
        _ = v42
    }
    if true {
        if ((!(2147483648 == 855)) || false) {
            v43 := GenS1{F0: 42, F1: 18446744073709551615}
            p0 = p0
            v44 := []int{}
            _ = v43
            _ = v44
        }
        v45 := GenS5{F0: 0, F1: (p0 | p1)}
        p2 -= -128
        _ = v45
    }
    p0 += (p0 % (((p1 - p0)) | 1))
    _ = p0
    _ = p1
    _ = p2
    return -951
}

func main() {
    var v1 uint64 = 1
    v2 := make(map[uint32]uint32)
    var v3 int = -512
    var v4 uint32 = 2
    v5 := []int{}
    for fi46 := 0; fi46 < 8; fi46++ {
        if true {
            fmt.Println(v1, v3, v3)
            switch (("hello" + "") + (("\n\t" + "hello") + "")) {
            case "foo bar":
                if ("a" != "") {
                    for fi47 := 0; fi47 < 8; fi47++ {
                        switch ("a" + "\n\t") {
                        case "foo bar":
                            v48 := v2[(^(v4))]
                            _ = v48
                        case "\n\t":
                            v1 -= v1
                        }
                        _ = fi47
                    }
                    v1 += v1
                } else {
                    var v49 uint8 = uint8(v4)
                    v50 := GenS5{F0: v1, F1: 42}
                    _ = v49
                    _ = v50
                }
            case "a":
                var v51 int16 = 1
                _ = v51
            }
            v2[220] = v4
        } else {
            var v52 int64 = -151
            var v53 uint32 = (v4 - v4)
            if len(v5) > 0 {
                v5[0] = v3
            }
            _ = v52
            _ = v53
        }
        if false {
            for fi54 := 0; fi54 < 14; fi54++ {
                var v55 uint16 = 65535
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                var v56 uint32 = (v4 ^ (v4 + (v4 | 65535)))
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                var v57 string = "a"
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = fi54
                _ = v55
                _ = v56
                _ = v57
            }
        } else {
            switch (v3 * v3) {
            case 4294967295:
                fmt.Println(v4, v1)
            }
            var v58 int64 = 845
            _ = v58
        }
        _ = fi46
    }
    var v59 int16 = 1
    if false {
        for fi60 := 0; fi60 < 1; fi60++ {
            var v61 int64 = 189
            for fi62 := 0; fi62 < 5; fi62++ {
                v63 := GenS4{F0: 1, F1: fi62, F2: -316}
                if (false || (((^(v61)) > v61) && ((true || (!(false && false))) && true))) {
                    v64 := GenS3{F0: v61}
                    v59 += v59
                    v1 += (v1 * v1)
                    var v65 float32 = 0.1
                    var v66 int32 = 0
                    _ = v64
                    _ = v65
                    _ = v66
                } else {
                    fmt.Println(v61, v61, v59)
                    switch (fi60 / (((fi62 & (fi60 * 128))) | 1)) {
                    case -1:
                        var v67 bool = false
                        _ = v67
                    case 2147483647:
                        v59 = v59
                    case 833:
                        v4 = v4
                    default:
                        v68 := GenS2{F0: 255, F1: 4294967295, F2: ""}
                        _ = v68
                    }
                }
                _ = fi62
                _ = v63
            }
            switch "foo bar" {
            case "foo bar":
                var v69 uint32 = (v4 - 0)
                _ = v69
            }
            _ = fi60
            _ = v61
        }
        var v70 int16 = (v59 / ((-32768) | 1))
        _ = v70
    }
    if (false || false) {
        v71 := v2[42]
        for fi72 := 0; fi72 < 19; fi72++ {
            var v73 uint32 = (v4 ^ v71)
            if false {
                fmt.Println(v3)
                for fi74 := 0; fi74 < 6; fi74++ {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    fmt.Println(fi74, fi72)
                    _ = fi74
                }
            } else {
                if true {
                    for fi75 := 0; fi75 < 8; fi75++ {
                        if ((-(fi75)) <= fi75) {
                            if (v3 != int(v1)) {
                                v76 := GenS6{F0: genFn0(), F1: 128}
                                v77 := GenS3{F0: 42}
                                _ = v76
                                _ = v77
                            } else {
                                var v78 uint = 255
                                switch "\n\t" {
                                case "foo bar":
                                    v79 := GenS7{F0: v59, F1: 1, F2: 42}
                                    _ = v79
                                }
                                v5 = append(v5, fi72)
                                _ = v78
                            }
                        }
                        var v80 uint64 = v1
                        _ = fi75
                        _ = v80
                    }
                    v73 = 0
                    if (18446744073709551615 <= genFn0()) {
                        v4 += (^(v71))
                        for fi81 := 0; fi81 < 4; fi81++ {
                            if true {
                                fmt.Println(v71, fi81)
                                if (0 > genFn0()) {
                                    var v82 uint64 = 0
                                    _ = v82
                                } else {
                                    if (9223372036854775807 == -10) {
                                        v73 += v73
                                        var v83 int8 = 42
                                        v84 := GenS3{F0: -9223372036854775808}
                                        _ = v83
                                        _ = v84
                                    }
                                    if false {
                                        if ((v3 != 4294967295) && false) {
                                            v71 = (^(v73))
                                            switch "foo bar" {
                                            case "a":
                                                v1 -= v1
                                            case "hello":
                                                switch "\n\t" {
                                                case "foo bar":
                                                    switch 32767 {
                                                    case 0:
                                                        var v85 int = (v3 * (fi72 + (fi81 | fi72)))
                                                        var v86 int64 = 499
                                                        _ = v85
                                                        _ = v86
                                                    default:
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                    }
                                                case "a":
                                                    if (v3 > fi72) {
                                                        if (false || true) {
                                                            var v87 int32 = 91
                                                            var v88 int = (v3 & v3)
                                                            if false {
                                                                fmt.Println(v3)
                                                                v89 := GenS8{F0: "a"}
                                                                _ = v89
                                                            }
                                                            _ = v87
                                                            _ = v88
                                                        }
                                                        v59 = 1000
                                                        if len(v5) > 0 {
                                                            v5[0] = 54
                                                        }
                                                    }
                                                case "":
                                                    var v90 int32 = 2147483647
                                                    _ = v90
                                                default:
                                                    var v91 uint32 = 2147483648
                                                    _ = v91
                                                }
                                            case "foo bar":
                                                if false {
                                                    v71 += (v4 >> ((genFn0()) & 63))
                                                    if len(v5) > 0 {
                                                        v5[0] = v3
                                                    }
                                                    v73 = (v71 + (^(v4)))
                                                }
                                            default:
                                                if (int64(v1) >= -1) {
                                                    v71 += v73
                                                    v1 += v1
                                                    fmt.Println(v73, v59, fi72)
                                                }
                                                v92 := GenS9{F0: v1, F1: false}
                                                v93 := GenS5{F0: 0, F1: 42}
                                                _ = v92
                                                _ = v93
                                            }
                                            fmt.Println(fi81)
                                        }
                                    } else {
                                        v71 = (v4 % (((v71 % ((2147483648) | 1))) | 1))
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    }
                                }
                                _ = len(v2)
                            } else {
                                v94 := GenS2{F0: 876, F1: (v1 % ((47) | 1)), F2: "hello"}
                                v3 = v3
                                _ = v94
                            }
                            _ = fi81
                        }
                        v71 -= (v71 ^ (^(v71)))
                    }
                }
            }
            fmt.Println(v1, v71, v3)
            _ = fi72
            _ = v73
        }
        _ = v71
    } else {
        if false {
            v4 += (^(v4))
        }
    }
    if (v1 >= (^(v1))) {
        for fi95 := 0; fi95 < 1; fi95++ {
            var v96 float32 = 1.5
            var v97 int8 = 100
            var v98 uint64 = (v1 + (v1 + v1))
            _ = fi95
            _ = v96
            _ = v97
            _ = v98
        }
    } else {
        if (float64(float32(float64(1.5))) != ((1e10 * (0.0 / (1e10 - (1.5 + 0.1)))) * float64((1.5 / 3.14159)))) {
            v3 += (^(v3))
        } else {
            for fi99 := 0; fi99 < 20; fi99++ {
                v4 += v4
                _ = fi99
            }
            v59 = (-(v59))
        }
        v3 = (-(v3))
    }
    var v100 bool = (!((v3 | (v3 % (((v3 % ((v3) | 1))) | 1))) >= (v3 << ((683) & 63))))
    fmt.Println(v1, v4, v4)
    switch "" {
    case "\n\t":
        fmt.Println(v1, v3)
    case "foo bar":
        var v101 int64 = -9223372036854775808
        _ = v101
    }
    v102 := GenS6{F0: 4294967295, F1: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v100:", v100)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v100
    _ = v102
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v59
}
