package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int; F1 int; F2 int16 }
type GenS5 struct { F0 int16; F1 uint }
type GenS6 struct { F0 uint; F1 uint8; F2 int8 }
type GenS7 struct { F0 bool; F1 bool }
type GenS8 struct { F0 bool }
func genFn0(p0 int32, p1 uint, p2 int64) string {
    if ((p0 | (-(p0))) <= p0) {
        if true {
            v7 := []int{1, 42}
            v8 := GenS1{F0: 42, F1: 18446744073709551615}
            for fi9 := 0; fi9 < 17; fi9++ {
                if (0 < 0) {
                    p2 = 42
                    var v10 uint64 = uint64(p1)
                    _ = len(v7)
                    _ = v10
                } else {
                    fmt.Println(fi9)
                    fmt.Println(p0, p2, p0)
                    if false {
                        p0 = p0
                        p0 -= p0
                    } else {
                        fmt.Println(p2)
                        var v11 int = fi9
                        for fi12 := 0; fi12 < 14; fi12++ {
                            p0 = p0
                            v13 := GenS2{F0: 42}
                            p0 += -1
                            _ = fi12
                            _ = v13
                        }
                        _ = v11
                    }
                }
                v14 := GenS1{F0: 42, F1: 255}
                for fi15 := 0; fi15 < 4; fi15++ {
                    v7 = append(v7, (fi15 / ((fi15) | 1)))
                    _ = fi15
                }
                _ = fi9
                _ = v14
            }
            _ = v7
            _ = v8
        }
        v16 := GenS3{F0: -7}
        _ = v16
    }
    var v17 int32 = (-(p0))
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return ""
}
func genFn1() uint16 {
    switch 1 {
    case -800:
        var v18 int16 = 32767
        _ = v18
    default:
        switch (("foo bar" + (((("a" + "a") + "a") + "hello") + "\n\t")) + "a") {
        case "foo bar":
            if false {
                switch 843 {
                case 128:
                    v19 := GenS3{F0: 123}
                    _ = v19
                case -9223372036854775808:
                    switch 32767 {
                    case -2147483648:
                        switch 2147483647 {
                        case -463:
                            for fi20 := 0; fi20 < 17; fi20++ {
                                var v21 uint64 = 71
                                if false {
                                    v21 -= 0
                                } else {
                                    v21 = (v21 % ((539) | 1))
                                    var v22 int8 = 1
                                    v21 += v21
                                    _ = v22
                                }
                                _ = fi20
                                _ = v21
                            }
                        case -39:
                            v23 := GenS1{F0: 65535, F1: 4294967295}
                            _ = v23
                        default:
                            var v24 int32 = -2147483648
                            _ = v24
                        }
                    }
                default:
                    v25 := []int{32768}
                    _ = v25
                }
                var v26 int = -2147483648
                var v27 float32 = (3.14159 / (float32((3.14159 - (1e10 / 1.5))) - 1.5))
                _ = v26
                _ = v27
            } else {
                var v28 uint32 = 254
                for fi29 := 0; fi29 < 5; fi29++ {
                    var v30 int16 = 0
                    var v31 float64 = 3.14159
                    v32 := GenS1{F0: v28, F1: 785}
                    _ = fi29
                    _ = v30
                    _ = v31
                    _ = v32
                }
                _ = v28
            }
        default:
            var v33 bool = true
            _ = v33
        }
    }
    var v34 string = ("hello" + "\n\t")
    v34 = "\n\t"
    switch "foo bar" {
    case "hello":
        if (false && (1e10 >= (float64(3.14159) * 1.5))) {
            for fi35 := 0; fi35 < 3; fi35++ {
                v34 = v34
                fmt.Println(v34, v34)
                var v36 uint = 216
                _ = fi35
                _ = v36
            }
        }
    case "a":
        if false {
            var v37 int64 = 1
            _ = v37
        }
    default:
        if true {
            fmt.Println(v34)
        } else {
            switch "foo bar" {
            case "\n\t":
                var v38 int32 = 0
                _ = v38
            default:
                var v39 uint = 1
                _ = v39
            }
        }
    }
    var v40 string = ("\n\t" + v34)
    _ = v34
    _ = v40
    return 0
}

func main() {
    var v1 int16 = 32767
    var v2 bool = true
    v3 := []uint32{2, 255, 384, 2}
    v4 := make(map[int]string)
    var v5 bool = ((v1 / ((808) | 1)) >= 1)
    v6 := make(map[uint32]string)
    if (0.1 < 1e10) {
        for fi41 := 0; fi41 < 9; fi41++ {
            var v42 float64 = (float64(0.0) * (1.5 - (-2.25 * 1.5)))
            _ = fi41
            _ = v42
        }
    }
    fmt.Println(v1, v2)
    switch 127 {
    case 42:
        v2 = false
    case -1:
        var v43 int8 = 42
        _ = v43
    case 127:
        v44 := GenS4{F0: -9223372036854775808, F1: 128, F2: 384}
        _ = v44
    }
    if (-2147483648 < 128) {
        var v45 uint8 = 35
        _ = v45
    }
    v46 := GenS4{F0: 4294967295, F1: 42, F2: (v1 % ((0) | 1))}
    fmt.Println(v2)
    if ((0 <= 0) || (32768 == -761)) {
        v5 = true
        var v47 int = 653
        _ = v47
    } else {
        switch "\n\t" {
        case "":
            v5 = false
        default:
            v5 = v5
        }
    }
    switch 9223372036854775807 {
    case 839:
        switch 128 {
        case -9223372036854775808:
            for fi48 := 0; fi48 < 17; fi48++ {
                for fi49 := 0; fi49 < 18; fi49++ {
                    if false {
                        fmt.Println(v2, v5)
                        v5 = (false || true)
                    } else {
                        var v50 int64 = 1
                        fmt.Println(fi49, v5, v50)
                        var v51 int32 = 42
                        _ = v50
                        _ = v51
                    }
                    var v52 int64 = -792
                    if true {
                        v52 = v52
                        v1 += v1
                        switch fi49 {
                        case 32767:
                            v2 = (2 >= 2147483648)
                        }
                    }
                    _ = fi49
                    _ = v52
                }
                for fi53 := 0; fi53 < 17; fi53++ {
                    fmt.Println(v2, v2, v1)
                    switch "a" {
                    case "a":
                        v5 = (0 >= -9223372036854775808)
                    case "\n\t":
                        fmt.Println(fi53, v2)
                    case "hello":
                        switch ("foo bar" + ("" + "hello")) {
                        case "":
                            var v54 uint64 = 2
                            _ = v54
                        case "\n\t":
                            fmt.Println(fi53, fi53)
                        case "foo bar":
                            if (true && v2) {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                        default:
                            v2 = (v5 || (0.1 >= ((0.0 * float32((1e10 - (1.5 / 3.14159)))) + (float32((1.5 + 3.14159)) + 0.1))))
                        }
                    default:
                        switch "hello" {
                        case "foo bar":
                            v3 = append(v3, 1)
                        case "a":
                            if false {
                                if false {
                                    switch "\n\t" {
                                    case "":
                                        var v55 int16 = 1000
                                        _ = v55
                                    case "foo bar":
                                        v56 := GenS4{F0: (fi48 ^ fi48), F1: fi53, F2: 0}
                                        _ = v56
                                    case "a":
                                        var v57 string = ""
                                        _ = v57
                                    }
                                    var v58 int8 = 33
                                    var v59 uint8 = 1
                                    _ = v58
                                    _ = v59
                                }
                                switch "\n\t" {
                                case "\n\t":
                                    v2 = v2
                                case "":
                                    var v60 uint8 = 1
                                    _ = v60
                                }
                            }
                        default:
                            var v61 uint = 4294967295
                            _ = v61
                        }
                        v62 := GenS2{F0: 372}
                        _ = v62
                    }
                    _ = fi53
                }
                _ = fi48
            }
        case 32768:
            var v63 uint = 42
            _ = v63
        case 32767:
            fmt.Println(v5)
        default:
            switch "foo bar" {
            case "\n\t":
                v1 -= (^(v1))
            case "a":
                var v64 uint = 1
                _ = v64
            case "hello":
                var v65 uint64 = 18446744073709551615
                _ = v65
            default:
                switch -17 {
                case -843:
                    if (!(346 != 2)) {
                        var v66 bool = (!(!(71 > 100)))
                        delete(v4, -9223372036854775808)
                        _ = v66
                    } else {
                        _ = len(v3)
                        if v2 {
                            fmt.Println(v2, v2)
                            if v5 {
                                _ = len(v3)
                                v67 := v4[128]
                                if (false && v2) {
                                    fmt.Println(v5)
                                    v68 := GenS1{F0: 2147483648, F1: 2}
                                    v69 := GenS5{F0: v1, F1: 2}
                                    _ = v68
                                    _ = v69
                                }
                                _ = v67
                            } else {
                                v70 := GenS2{F0: -821}
                                for fi71 := 0; fi71 < 16; fi71++ {
                                    var v72 int64 = 9223372036854775807
                                    for fi73 := 0; fi73 < 15; fi73++ {
                                        v72 = v72
                                        switch (^(fi73)) {
                                        case -9223372036854775808:
                                            switch "\n\t" {
                                            case "a":
                                                v74 := GenS2{F0: v72}
                                                _ = v74
                                            }
                                        case 0:
                                            var v75 uint = 65535
                                            _ = v75
                                        }
                                        _ = fi73
                                    }
                                    if v5 {
                                        v76 := GenS4{F0: (^(fi71)), F1: fi71, F2: v1}
                                        v77 := GenS3{F0: -1}
                                        _ = v76
                                        _ = v77
                                    }
                                    _ = fi71
                                    _ = v72
                                }
                                _ = v70
                            }
                            v4[350] = "a"
                        } else {
                            v1 = -32768
                        }
                        for fi78 := 0; fi78 < 20; fi78++ {
                            v79 := v6[4294967295]
                            _ = fi78
                            _ = v79
                        }
                    }
                case 42:
                    for fi80 := 0; fi80 < 4; fi80++ {
                        for fi81 := 0; fi81 < 14; fi81++ {
                            v82 := GenS3{F0: 0}
                            v1 = (v1 << ((1) & 63))
                            v2 = true
                            _ = fi81
                            _ = v82
                        }
                        _ = fi80
                    }
                case 127:
                    v2 = (0.0 > 1.5)
                }
            }
        }
    case 4294967295:
        v83 := GenS6{F0: 0, F1: 130, F2: 72}
        _ = v83
    }
    v84 := GenS6{F0: 255, F1: 42, F2: 86}
    switch "\n\t" {
    case "\n\t":
        var v85 int64 = 42
        _ = v85
    }
    v86 := GenS5{F0: (-(v1)), F1: 18446744073709551615}
    if (("" != ("hello" + ((("foo bar" + "hello") + "hello") + ("\n\t" + ("hello" + ("a" + ("a" + "hello"))))))) && (!(42 == 42))) {
        v87 := GenS6{F0: 2, F1: 42, F2: 100}
        var v88 int64 = -1
        _ = v87
        _ = v88
    } else {
        v2 = true
    }
    if (!((-554 == 1) && (1e10 > (1.5 * (1e10 / (-2.25 * 0.1)))))) {
        var v89 uint8 = 14
        for fi90 := 0; fi90 < 18; fi90++ {
            var v91 uint32 = 2
            for fi92 := 0; fi92 < 6; fi92++ {
                var v93 string = "foo bar"
                fmt.Println(fi90, v1, v93)
                var v94 uint = uint(v91)
                _ = fi92
                _ = v93
                _ = v94
            }
            v95 := GenS7{F0: true, F1: v2}
            _ = fi90
            _ = v91
            _ = v95
        }
        var v96 uint = 42
        _ = v89
        _ = v96
    } else {
        var v97 bool = v2
        _ = v97
    }
    v2 = ((977 == 0) && ((!(0.0 <= 0.1)) && v2))
    var v98 uint64 = 2
    v99 := GenS7{F0: (v5 || (255 < 2)), F1: false}
    var v100 int32 = 0
    if (v2 && (-9223372036854775808 == 226)) {
        v101 := GenS1{F0: 1, F1: (v98 / ((957) | 1))}
        var v102 string = ""
        var v103 uint32 = 402
        _ = v101
        _ = v102
        _ = v103
    } else {
        fmt.Println(v2, v5, v1)
        var v104 float32 = float32((-2.25 / 1.0))
        v105 := GenS8{F0: false}
        _ = v104
        _ = v105
    }
    fmt.Println("v1:", v1)
    fmt.Println("v100:", v100)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v98:", v98)
    _ = v1
    _ = v100
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v6
    _ = v84
    _ = v86
    _ = v98
    _ = v99
}
