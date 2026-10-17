package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32; F1 bool; F2 uint8 }
type GenS5 struct { F0 uint8; F1 int8; F2 bool }
type GenS6 struct { F0 uint8; F1 int8 }
type GenS7 struct { F0 int16; F1 int64 }
type GenS8 struct { F0 uint16; F1 int64 }
type GenS9 struct { F0 int32; F1 int16 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 int32 }

func main() {
    var v1 int8 = 5
    var v2 int16 = -32768
    v3 := make(map[uint32]string)
    var v4 int16 = v2
    var v5 string = "foo bar"
    fmt.Println(v1, v5, v2)
    switch v5 {
    case "":
        var v6 float64 = -2.25
        _ = v6
    case "a":
        v7 := []int{-1, -2147483648, 32767, 940}
        _ = v7
    default:
        switch "a" {
        case "a":
            var v8 string = (("hello" + v5) + (v5 + v5))
            _ = v8
        case "hello":
            var v9 int16 = -1
            _ = v9
        }
    }
    var v10 uint16 = 1
    for fi11 := 0; fi11 < 8; fi11++ {
        v5 = v5
        _ = fi11
    }
    var v12 string = ((v5 + "foo bar") + ("hello" + ("\n\t" + "foo bar")))
    fmt.Println(v12)
    for fi13 := 0; fi13 < 1; fi13++ {
        if true {
            var v14 uint8 = 127
            _ = v14
        } else {
            var v15 uint16 = v10
            _ = v15
        }
        _ = fi13
    }
    for fi16 := 0; fi16 < 11; fi16++ {
        for fi17 := 0; fi17 < 3; fi17++ {
            var v18 string = v12
            v5 = "a"
            _ = fi17
            _ = v18
        }
        delete(v3, 65535)
        if (388 != 0) {
            fmt.Println(v10, v12, v5)
            for fi19 := 0; fi19 < 20; fi19++ {
                v10 += (v10 << ((2) & 63))
                delete(v3, 717)
                var v20 uint8 = 95
                _ = fi19
                _ = v20
            }
            for fi21 := 0; fi21 < 10; fi21++ {
                v5 = (("hello" + "hello") + "")
                _ = fi21
            }
        } else {
            for fi22 := 0; fi22 < 17; fi22++ {
                var v23 uint32 = 349
                _ = fi22
                _ = v23
            }
        }
        _ = fi16
    }
    switch v12 {
    case "a":
        _ = len(v3)
    case "foo bar":
        v4 = 1
    }
    switch "hello" {
    case "\n\t":
        for fi24 := 0; fi24 < 3; fi24++ {
            switch fi24 {
            case 406:
                fmt.Println(v1, v5, v1)
            }
            switch fi24 {
            case 456:
                for fi25 := 0; fi25 < 4; fi25++ {
                    if (-29 != 1) {
                        if (708 < 726) {
                            v26 := GenS1{F0: 1}
                            v5 = ("" + (v5 + v5))
                            var v27 float32 = ((1e10 / 1.0) - 3.14159)
                            if false {
                                v28 := v3[170]
                                switch (fi24 / (((^(fi24))) | 1)) {
                                case 1:
                                    fmt.Println(v10)
                                case 152:
                                    fmt.Println(v12, fi24, v12)
                                case 2147483647:
                                    if (65535 != (^(v10))) {
                                        var v29 uint32 = 148
                                        var v30 uint32 = (v29 - 2)
                                        if true {
                                            v31 := GenS1{F0: 489}
                                            var v32 uint8 = 43
                                            _ = v31
                                            _ = v32
                                        }
                                        _ = v29
                                        _ = v30
                                    } else {
                                        var v33 float64 = -2.25
                                        switch fi24 {
                                        case 754:
                                            v34 := GenS1{F0: 4294967295}
                                            _ = v34
                                        case 9223372036854775807:
                                            switch (v28 + v28) {
                                            case "a":
                                                var v35 uint = 42
                                                _ = v35
                                            case "":
                                                if (((!false) || (!(true || false))) && false) {
                                                    v33 -= (v33 + v33)
                                                    if (1 != 9223372036854775807) {
                                                        v36 := []int{}
                                                        if (v27 == v27) {
                                                            v37 := GenS1{F0: 4294967295}
                                                            v38 := GenS2{F0: "", F1: v2}
                                                            v39 := GenS3{F0: 1}
                                                            _ = v37
                                                            _ = v38
                                                            _ = v39
                                                        } else {
                                                            switch (v5 + v5) {
                                                            case "a":
                                                                if true {
                                                                    v28 = (v12 + (((v5 + "\n\t") + v28) + v5))
                                                                } else {
                                                                    switch v12 {
                                                                    case "foo bar":
                                                                        if false {
                                                                            v33 = v33
                                                                            if true {
                                                                                var v40 int8 = v1
                                                                                switch ((v12 + ((((v28 + v12) + (("" + "a") + v5)) + "a") + v28)) + v5) {
                                                                                case "":
                                                                                    var v41 uint64 = 336
                                                                                    _ = v41
                                                                                }
                                                                                _ = v40
                                                                            }
                                                                            v4 = v2
                                                                            v42 := GenS3{F0: v1}
                                                                            v36 = append(v36, fi25)
                                                                            _ = v42
                                                                        }
                                                                    default:
                                                                        var v43 int16 = (v2 >> ((144) & 63))
                                                                        _ = v43
                                                                    }
                                                                }
                                                            case "":
                                                                if false {
                                                                    if false {
                                                                        v1 += v1
                                                                    }
                                                                }
                                                            default:
                                                                if len(v36) >= 2 {
                                                                    v36 = v36[1:len(v36)]
                                                                }
                                                                fmt.Println(fi25, v5, v2)
                                                                delete(v3, 42)
                                                            }
                                                        }
                                                        _ = v36
                                                    } else {
                                                        fmt.Println(v33, v10)
                                                        v44 := GenS4{F0: 4294967295, F1: (497 != 201), F2: 233}
                                                        _ = v44
                                                    }
                                                    v5 = v12
                                                    switch 2147483647 {
                                                    case 0:
                                                        v45 := GenS5{F0: 2, F1: (v1 * (v1 + (v1 + (v1 & (v1 * int8(fi24)))))), F2: true}
                                                        _ = v45
                                                    case 32767:
                                                        var v46 uint64 = 42
                                                        _ = v46
                                                    default:
                                                        if false {
                                                            delete(v3, 2147483648)
                                                        }
                                                    }
                                                    if (0.0 < -2.25) {
                                                        v47 := GenS1{F0: 185}
                                                        v48 := []int{42, (fi25 % ((-1) | 1)), fi24}
                                                        v49 := GenS6{F0: 2, F1: v1}
                                                        _ = v47
                                                        _ = v48
                                                        _ = v49
                                                    } else {
                                                        var v50 bool = true
                                                        v4 -= (v2 | v4)
                                                        _ = v50
                                                    }
                                                    if (v4 >= (-(v4))) {
                                                        var v51 uint32 = 525
                                                        v1 += (v1 >> ((0) & 63))
                                                        v3[(v51 * v51)] = v12
                                                        var v52 int64 = -9223372036854775808
                                                        _ = v51
                                                        _ = v52
                                                    }
                                                    v5 = ("a" + v5)
                                                } else {
                                                    switch fi24 {
                                                    case -9223372036854775808:
                                                        var v53 bool = true
                                                        _ = v53
                                                    case 1:
                                                        v4 = (v2 - v2)
                                                    case 127:
                                                        v27 = v27
                                                        var v54 int32 = -745
                                                        if true {
                                                            switch fi25 {
                                                            case -339:
                                                                v55 := v3[65535]
                                                                _ = v55
                                                            case 681:
                                                                switch (-(fi24)) {
                                                                case 127:
                                                                    switch fi25 {
                                                                    case 32768:
                                                                        fmt.Println(v4)
                                                                    case 780:
                                                                        v5 = (("" + v12) + v5)
                                                                    }
                                                                    _ = len(v3)
                                                                case 352:
                                                                    var v56 uint8 = 255
                                                                    _ = v56
                                                                default:
                                                                    if (!true) {
                                                                        fmt.Println(v4, v28, v33)
                                                                        v57 := GenS1{F0: 0}
                                                                        _ = v57
                                                                    } else {
                                                                        var v58 int8 = 72
                                                                        switch "foo bar" {
                                                                        case "hello":
                                                                            v10 += v10
                                                                        }
                                                                        v59 := []int{fi24, (fi25 ^ 32767)}
                                                                        _ = v58
                                                                        _ = v59
                                                                    }
                                                                }
                                                            case 32767:
                                                                fmt.Println(v12, v1)
                                                            }
                                                        }
                                                        _ = v54
                                                    }
                                                }
                                            case "hello":
                                                v10 += (v10 / ((v10) | 1))
                                            }
                                        case -460:
                                            v60 := GenS7{F0: (v2 >> ((70) & 63)), F1: 9223372036854775807}
                                            _ = v60
                                        }
                                        var v61 int64 = 0
                                        _ = v33
                                        _ = v61
                                    }
                                default:
                                    v4 = (v4 | v4)
                                }
                                _ = v28
                            } else {
                                v5 = ("\n\t" + "foo bar")
                                v62 := GenS8{F0: v10, F1: 9223372036854775807}
                                _ = v62
                            }
                            if (42 == 9223372036854775807) {
                                v12 = ""
                                v63 := GenS5{F0: 0, F1: v1, F2: (-9223372036854775808 != 0)}
                                _ = v63
                            } else {
                                if ((v10 ^ (v10 ^ (^(v10)))) >= v10) {
                                    if (128 >= 1) {
                                        var v64 int64 = 42
                                        fmt.Println(v1, fi25)
                                        v65 := GenS1{F0: 4294967295}
                                        var v66 uint64 = 42
                                        var v67 float32 = float32((float64(float32(-2.25)) / 1e10))
                                        _ = v64
                                        _ = v65
                                        _ = v66
                                        _ = v67
                                    } else {
                                        v68 := []int{}
                                        _ = v68
                                    }
                                    v1 += 89
                                    switch (^(fi24)) {
                                    case -9223372036854775808:
                                        v2 = v2
                                    case 1:
                                        var v69 int = int(v2)
                                        _ = v69
                                    case 91:
                                        var v70 int16 = (^(v2))
                                        _ = v70
                                    default:
                                        if false {
                                            var v71 int8 = (v1 ^ (v1 * v1))
                                            var v72 uint32 = uint32(v10)
                                            _ = v71
                                            _ = v72
                                        } else {
                                            fmt.Println(v10, v5)
                                            var v73 string = "\n\t"
                                            v1 = 1
                                            _ = v73
                                        }
                                    }
                                    v2 = v4
                                    delete(v3, 0)
                                } else {
                                    v4 += (v2 ^ 42)
                                    v74 := GenS3{F0: v1}
                                    var v75 uint64 = 2
                                    _ = v74
                                    _ = v75
                                }
                            }
                            v4 += v2
                            _ = v26
                            _ = v27
                        }
                        var v76 uint64 = 0
                        _ = v76
                    } else {
                        _ = len(v3)
                    }
                    _ = fi25
                }
            case 32767:
                var v77 uint32 = 1
                _ = v77
            default:
                v78 := GenS6{F0: 2, F1: (v1 | (v1 | (v1 & (v1 / ((v1) | 1)))))}
                _ = v78
            }
            v79 := []int{(fi24 & fi24), 2147483647, fi24, fi24}
            _ = fi24
            _ = v79
        }
    default:
        if (v10 > v10) {
            fmt.Println(v5, v5)
            var v80 int = -2147483648
            var v81 int8 = (v1 / (((v1 ^ v1)) | 1))
            _ = v80
            _ = v81
        }
    }
    v4 -= (v4 / ((32767) | 1))
    v10 -= v10
    v2 = v4
    var v82 float32 = 0.0
    var v83 uint8 = 9
    v12 = v12
    for fi84 := 0; fi84 < 18; fi84++ {
        var v85 string = v5
        v82 += float32(0.1)
        switch fi84 {
        case 271:
            switch (fi84 + (fi84 + (fi84 ^ -213))) {
            case -1:
                fmt.Println(v5)
            case 4294967295:
                v86 := GenS9{F0: 933, F1: v4}
                _ = v86
            default:
                for fi87 := 0; fi87 < 17; fi87++ {
                    if false {
                        switch "hello" {
                        case "":
                            fmt.Println(v5)
                            v88 := GenS1{F0: 4294967295}
                            v5 = (v5 + (v12 + "\n\t"))
                            _ = v88
                        case "\n\t":
                            v89 := GenS10{F0: 42}
                            _ = v89
                        case "a":
                            v90 := []int{fi84, 119, fi87}
                            _ = v90
                        default:
                            v12 = v5
                        }
                        v91 := GenS5{F0: 127, F1: (v1 ^ -4), F2: false}
                        var v92 float32 = v82
                        _ = v91
                        _ = v92
                    } else {
                        fmt.Println(fi84)
                        var v93 int = (^(fi84))
                        v94 := GenS6{F0: (v83 % ((v83) | 1)), F1: -128}
                        _ = v93
                        _ = v94
                    }
                    v95 := GenS11{F0: -1}
                    var v96 bool = (42 != -218)
                    switch fi84 {
                    case 32767:
                        switch fi87 {
                        case -1:
                            fmt.Println(fi87, v12)
                        case -2147483648:
                            v97 := GenS5{F0: v83, F1: v1, F2: v96}
                            _ = v97
                        default:
                            var v98 int = -2147483648
                            _ = v98
                        }
                    case -1:
                        fmt.Println(v12, v85, v10)
                    default:
                        delete(v3, 0)
                    }
                    _ = fi87
                    _ = v95
                    _ = v96
                }
            }
        case 980:
            var v99 int = fi84
            _ = v99
        default:
            v100 := v3[4294967295]
            _ = v100
        }
        _ = fi84
        _ = v85
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v82:", v82)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v82
    _ = v83
}
