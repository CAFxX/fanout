package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int; F1 int32; F2 uint32 }
type GenS4 struct { F0 uint64; F1 int; F2 uint32 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint64 }

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 int = -826
    var v3 uint32 = 792
    var v4 float32 = (float32(0.0) * float32(1.5))
    switch v2 {
    case 127:
        fmt.Println(v2)
    case -9223372036854775808:
        v5 := []int{v2, 599, v2}
        _ = v5
    }
    var v6 uint = 4294967295
    var v7 uint16 = 290
    for fi8 := 0; fi8 < 19; fi8++ {
        v2 = v2
        var v9 uint16 = v7
        var v10 int32 = -499
        _ = fi8
        _ = v10
        _ = v9
    }
    switch ("a" + "") {
    case "foo bar":
        var v11 int32 = -546
        _ = v11
    case "":
        var v12 float64 = 3.14159
        _ = v12
    case "a":
        fmt.Println(v4)
    default:
        if (!true) {
            switch ("hello" + "hello") {
            case "a":
                switch v2 {
                case -75:
                    v13 := GenS1{F0: (2 != (v1 - v1))}
                    _ = v13
                }
            }
        } else {
            var v14 int64 = -1
            v15 := []int{}
            _ = v14
            _ = v15
        }
    }
    v6 += (^(v6))
    var v16 uint64 = 0
    for fi17 := 0; fi17 < 9; fi17++ {
        var v18 int64 = -9223372036854775808
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 16; fi19++ {
        if (!false) {
            fmt.Println(v16, v7)
            for fi20 := 0; fi20 < 17; fi20++ {
                fmt.Println(fi19, fi20)
                var v21 int8 = 127
                v22 := GenS1{F0: false}
                _ = fi20
                _ = v21
                _ = v22
            }
            for fi23 := 0; fi23 < 19; fi23++ {
                var v24 uint64 = 541
                v25 := GenS1{F0: (128 < 221)}
                var v26 float32 = v4
                v27 := GenS1{F0: ((^(v3)) > v3)}
                v7 += 255
                v28 := []int{(fi19 & -1), v2, (fi19 | fi19), 127}
                _ = fi23
                _ = v24
                _ = v25
                _ = v26
                _ = v27
                _ = v28
            }
        } else {
            if false {
                for fi29 := 0; fi29 < 3; fi29++ {
                    v4 = -2.25
                    fmt.Println(v3)
                    v30 := []int{}
                    _ = fi29
                    _ = v30
                }
            }
            var v31 int16 = -1
            _ = v31
        }
        switch -1 {
        case 114:
            for fi32 := 0; fi32 < 7; fi32++ {
                var v33 int8 = 0
                var v34 int16 = -708
                v35 := []int{v2, 42, int(v7), (^(v2))}
                v4 = 3.14159
                _ = fi32
                _ = v33
                _ = v34
                _ = v35
            }
        }
        var v36 string = "a"
        _ = fi19
        _ = v36
    }
    v2 += v2
    switch (^(v2)) {
    case -849:
        if (v3 >= v3) {
            if true {
                var v37 int8 = 112
                _ = v37
            } else {
                if false {
                    v4 += (v4 + 1e10)
                } else {
                    v7 = (v7 * (v7 % ((42) | 1)))
                    fmt.Println(v7, v3, v2)
                    v38 := []int{0, (v2 << ((18446744073709551615) & 63)), (v2 << ((2) & 63)), 42}
                    _ = v38
                }
                v3 = (v3 * (v3 - 65535))
                v2 = (v2 >> (((v6 / ((uint(v7)) | 1))) & 63))
            }
            if (!true) {
                var v39 uint32 = v3
                v40 := GenS1{F0: (false || true)}
                _ = v39
                _ = v40
            } else {
                if true {
                    if true {
                        var v41 uint8 = 81
                        v3 = (v3 << ((304) & 63))
                        for fi42 := 0; fi42 < 15; fi42++ {
                            if false {
                                v43 := GenS1{F0: false}
                                _ = v43
                            }
                            v16 -= (v1 | v16)
                            if (v3 < (v3 % ((v3) | 1))) {
                                var v44 int64 = -176
                                v45 := []int{fi42, fi42}
                                _ = v44
                                _ = v45
                            } else {
                                switch "\n\t" {
                                case "a":
                                    var v46 int16 = -1
                                    _ = v46
                                case "\n\t":
                                    var v47 uint16 = (v7 << (((v6 * (v6 | v6))) & 63))
                                    _ = v47
                                case "":
                                    v7 = (^(v7))
                                }
                                switch (-(v2)) {
                                case -9223372036854775808:
                                    if ((^(v3)) > v3) {
                                        v2 += 127
                                    }
                                case 4294967295:
                                    v48 := []int{fi42, v2}
                                    _ = v48
                                default:
                                    if false {
                                        var v49 uint8 = (v41 << ((v6) & 63))
                                        var v50 int64 = -1
                                        _ = v49
                                        _ = v50
                                    } else {
                                        for fi51 := 0; fi51 < 12; fi51++ {
                                            v52 := GenS2{F0: 140}
                                            _ = fi51
                                            _ = v52
                                        }
                                    }
                                }
                            }
                            _ = fi42
                        }
                        _ = v41
                    }
                } else {
                    for fi53 := 0; fi53 < 9; fi53++ {
                        if ((1e10 * (0.0 * (0.0 * float64(-2.25)))) == 1.5) {
                            switch ("a" + "a") {
                            case "hello":
                                v3 += (^(v3))
                            case "foo bar":
                                if true {
                                    var v54 string = "\n\t"
                                    _ = v54
                                } else {
                                    if true {
                                        v16 = v1
                                    }
                                    if true {
                                        v55 := []int{v2, 32768, int(v3)}
                                        var v56 uint8 = 247
                                        switch "foo bar" {
                                        case "foo bar":
                                            v56 = v56
                                        case "a":
                                            v3 = uint32(v1)
                                        }
                                        _ = v55
                                        _ = v56
                                    }
                                    fmt.Println(v2, v16, v16)
                                }
                            default:
                                for fi57 := 0; fi57 < 14; fi57++ {
                                    var v58 int = (fi57 % ((v2) | 1))
                                    var v59 int8 = 1
                                    if (((^(v7)) > 32767) && (-66 <= (v59 + -95))) {
                                        v58 -= fi57
                                    }
                                    _ = fi57
                                    _ = v58
                                    _ = v59
                                }
                            }
                            v60 := GenS3{F0: v2, F1: 787, F2: (v3 + 2)}
                            _ = v60
                        } else {
                            switch (v2 ^ (^(v2))) {
                            case 1:
                                switch fi53 {
                                case 32767:
                                    for fi61 := 0; fi61 < 5; fi61++ {
                                        v4 -= v4
                                        var v62 float32 = 1.5
                                        _ = fi61
                                        _ = v62
                                    }
                                default:
                                    fmt.Println(v2, v6, v16)
                                }
                            }
                            if true {
                                for fi63 := 0; fi63 < 8; fi63++ {
                                    if (true || true) {
                                        if (219 >= (^(v3))) {
                                            v64 := []int{}
                                            _ = v64
                                        }
                                        v4 = v4
                                        if (((("hello" + (("hello" + "foo bar") + "a")) + "\n\t") + "foo bar") == "\n\t") {
                                            v65 := GenS2{F0: (^(v3))}
                                            if false {
                                                v66 := GenS2{F0: v3}
                                                v3 -= (v3 - (v3 / ((v3) | 1)))
                                                _ = v66
                                            }
                                            v67 := GenS1{F0: true}
                                            _ = v65
                                            _ = v67
                                        }
                                    } else {
                                        v6 = v6
                                    }
                                    var v68 float32 = float32((0.1 - (1e10 / float64(v4))))
                                    v69 := []int{(v2 << ((0) & 63)), (v2 | (v2 << ((v6) & 63)))}
                                    _ = fi63
                                    _ = v68
                                    _ = v69
                                }
                            }
                        }
                        var v70 int16 = 703
                        _ = fi53
                        _ = v70
                    }
                }
                v7 += v7
                if (v6 > (^(v6))) {
                    v71 := GenS4{F0: (v16 - v1), F1: (v2 & (v2 & v2)), F2: (^(v3))}
                    fmt.Println(v6, v1, v4)
                    var v72 uint32 = v3
                    _ = v71
                    _ = v72
                } else {
                    var v73 int16 = -702
                    _ = v73
                }
            }
        } else {
            switch 693 {
            case 32767:
                switch ("\n\t" + ("hello" + "foo bar")) {
                case "a":
                    var v74 float32 = v4
                    _ = v74
                default:
                    switch 127 {
                    case -1:
                        for fi75 := 0; fi75 < 12; fi75++ {
                            for fi76 := 0; fi76 < 11; fi76++ {
                                v4 -= v4
                                switch "\n\t" {
                                case "hello":
                                    v77 := GenS4{F0: (v1 << ((4294967295) & 63)), F1: fi76, F2: v3}
                                    _ = v77
                                }
                                var v78 bool = true
                                _ = fi76
                                _ = v78
                            }
                            _ = fi75
                        }
                    default:
                        var v79 int32 = -1
                        _ = v79
                    }
                }
            }
        }
    case 9223372036854775807:
        var v80 uint8 = 128
        _ = v80
    case 906:
        v3 = (^(v3))
    default:
        if (!(v6 >= 108)) {
            var v81 int32 = 42
            fmt.Println(v1)
            _ = v81
        } else {
            for fi82 := 0; fi82 < 5; fi82++ {
                var v83 float64 = 1e10
                var v84 uint16 = uint16(v1)
                var v85 int16 = 1000
                _ = fi82
                _ = v83
                _ = v84
                _ = v85
            }
        }
    }
    for fi86 := 0; fi86 < 17; fi86++ {
        v1 = v1
        var v87 uint = 2
        var v88 float64 = 0.1
        _ = fi86
        _ = v87
        _ = v88
    }
    v89 := []int{(v2 / (((v2 ^ (v2 >> ((v6) & 63)))) | 1))}
    for fi90 := 0; fi90 < 9; fi90++ {
        if false {
            v91 := GenS3{F0: (v2 ^ fi90), F1: 2147483647, F2: v3}
            _ = v91
        } else {
            if (42 < 819) {
                switch (v2 & (v2 % (((-(fi90))) | 1))) {
                case -162:
                    var v92 float32 = v4
                    _ = v92
                default:
                    if ("a" == "a") {
                        fmt.Println(v3)
                        v89 = append(v89, 312)
                        if true {
                            fmt.Println(v1, v1, v7)
                        } else {
                            v4 += 1e10
                            for fi93 := 0; fi93 < 20; fi93++ {
                                if (false || (v3 <= v3)) {
                                    v94 := GenS3{F0: fi93, F1: -475, F2: v3}
                                    _ = v94
                                }
                                switch (fi90 % ((v2) | 1)) {
                                case 2147483647:
                                    v7 = (v7 + 0)
                                case -1:
                                    var v95 int16 = -32768
                                    _ = v95
                                default:
                                    fmt.Println(v7, v16, v6)
                                    if len(v89) > 0 {
                                        v89[0] = 677
                                    }
                                    var v96 uint8 = 0
                                    var v97 int8 = -1
                                    _ = v96
                                    _ = v97
                                }
                                v89 = append(v89, (v2 % (((v2 ^ (-(fi90)))) | 1)))
                                _ = fi93
                            }
                        }
                    }
                }
                var v98 uint16 = (v7 & v7)
                var v99 int16 = -102
                _ = v98
                _ = v99
            }
            if ((!(!(!(("foo bar" != "\n\t") || true)))) && (-32768 > 0)) {
                v1 = v16
            } else {
                v100 := GenS5{F0: (v1 * v1)}
                _ = len(v89)
                _ = v100
            }
            var v101 uint32 = v3
            _ = v101
        }
        _ = fi90
    }
    v102 := GenS6{F0: 429}
    fmt.Println(v2, v16, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v89:", len(v89), v89)
    _ = v1
    _ = v102
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v89
}
