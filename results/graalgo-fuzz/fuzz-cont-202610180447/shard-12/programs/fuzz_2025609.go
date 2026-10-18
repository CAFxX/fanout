package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 uint64 }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 uint8; F1 bool }
type GenS4 struct { F0 string; F1 uint64; F2 int64 }
type GenS5 struct { F0 int; F1 string; F2 uint }
type GenS6 struct { F0 int8; F1 uint64; F2 uint16 }
type GenS7 struct { F0 uint; F1 int8 }
type GenS8 struct { F0 uint8; F1 int32 }
type GenS9 struct { F0 uint; F1 int8; F2 uint }
type GenS10 struct { F0 int8; F1 uint32; F2 int16 }
type GenS11 struct { F0 int64; F1 int32 }
type GenS12 struct { F0 uint; F1 uint16; F2 int16 }
func genFn0() uint64 {
    var v5 float64 = 3.14159
    for fi6 := 0; fi6 < 6; fi6++ {
        if false {
            v7 := GenS1{F0: 2147483647, F1: 42, F2: 18446744073709551615}
            var v8 float64 = 0.1
            _ = v7
            _ = v8
        }
        _ = fi6
    }
    switch (((("foo bar" + "hello") + ("\n\t" + "a")) + "a") + "") {
    case "foo bar":
        switch 42 {
        case 9223372036854775807:
            v9 := GenS1{F0: 2147483647, F1: 0, F2: 2}
            _ = v9
        case 128:
            fmt.Println(v5)
        case 127:
            v5 = ((3.14159 * 1.5) - 3.14159)
        }
    }
    v10 := []int{-2147483648, -690, 0, 9223372036854775807}
    _ = v10
    _ = v5
    return 42
}
func genFn1() int64 {
    var v11 float32 = ((-2.25 - 0.1) * 3.14159)
    v12 := []int{9223372036854775807, 127, -703}
    v13 := GenS2{F0: 705, F1: 1}
    _ = v11
    _ = v12
    _ = v13
    return 0
}

func main() {
    v1 := make(map[int]bool)
    var v2 float32 = (0.0 / (0.0 + 3.14159))
    v3 := make(map[string]int)
    var v4 int32 = 970
    if false {
        v14 := GenS2{F0: 977, F1: (v4 / ((1) | 1))}
        _ = v14
    }
    v15 := GenS2{F0: 42, F1: (-(v4))}
    v1[403] = true
    v16 := GenS3{F0: 127, F1: true}
    v17 := GenS4{F0: "\n\t", F1: 690, F2: genFn1()}
    for fi18 := 0; fi18 < 19; fi18++ {
        fmt.Println(fi18, v4)
        v2 = float32(0.0)
        for fi19 := 0; fi19 < 15; fi19++ {
            switch ("a" + (("" + "\n\t") + "")) {
            case "\n\t":
                v4 += (v4 - -2147483648)
            case "a":
                v4 += (v4 / ((2147483647) | 1))
            default:
                v2 = (v2 / (v2 - (float32(3.14159) - (v2 / -2.25))))
            }
            _ = fi19
        }
        _ = fi18
    }
    for fi20 := 0; fi20 < 3; fi20++ {
        v21 := []int{}
        _ = fi20
        _ = v21
    }
    var v22 float64 = 0.1
    var v23 uint32 = 4294967295
    switch 9223372036854775807 {
    case 127:
        switch "hello" {
        case "\n\t":
            if false {
                if (((3.14159 / float32(float64(((0.1 + (v2 + 3.14159)) + v2)))) * (v2 / (1.5 / v2))) != v2) {
                    v24 := GenS2{F0: v23, F1: v4}
                    for fi25 := 0; fi25 < 11; fi25++ {
                        if (127 != 1) {
                            var v26 int64 = -283
                            _ = v26
                        }
                        if (float64(float32(v22)) == (-2.25 * -2.25)) {
                            if (3.14159 >= 3.14159) {
                                if (4294967295 >= (^(v23))) {
                                    if (!(!(false || (466 != 255)))) {
                                        if (32767 == 352) {
                                            v27 := []int{fi25}
                                            if true {
                                                if ("\n\t" <= "") {
                                                    var v28 uint32 = (v23 / (((v23 & v23)) | 1))
                                                    if false {
                                                        for fi29 := 0; fi29 < 1; fi29++ {
                                                            if (!(v22 > (0.1 / v22))) {
                                                                if false {
                                                                    v30 := v1[(-(fi25))]
                                                                    _ = v30
                                                                } else {
                                                                    var v31 uint64 = 2
                                                                    v4 += 612
                                                                    _ = len(v27)
                                                                    _ = v31
                                                                }
                                                                var v32 int64 = 0
                                                                if false {
                                                                    _ = len(v1)
                                                                    v33 := GenS5{F0: (fi25 + (fi29 + 529)), F1: "a", F2: 312}
                                                                    _ = v33
                                                                }
                                                                _ = v32
                                                            } else {
                                                                v34 := GenS4{F0: (("" + "a") + "foo bar"), F1: 239, F2: 1}
                                                                _ = v34
                                                            }
                                                            _ = fi29
                                                        }
                                                        for fi35 := 0; fi35 < 2; fi35++ {
                                                            v36 := GenS4{F0: ((("\n\t" + "hello") + "foo bar") + ((("foo bar" + "foo bar") + ("a" + "hello")) + "hello")), F1: 255, F2: genFn1()}
                                                            v37 := v3["a"]
                                                            _ = fi35
                                                            _ = v36
                                                            _ = v37
                                                        }
                                                        fmt.Println(v28, v23)
                                                    }
                                                    v23 = v23
                                                    _ = v28
                                                }
                                            } else {
                                                if ((2 < 89) || (1 >= 42)) {
                                                    v4 -= v4
                                                    var v38 uint32 = (v23 + (v23 * 2))
                                                    _ = v38
                                                } else {
                                                    delete(v3, "\n\t")
                                                    var v39 string = ("foo bar" + "hello")
                                                    _ = v39
                                                }
                                            }
                                            fmt.Println(v2, v2, v4)
                                            _ = v27
                                        } else {
                                            if false {
                                                var v40 uint16 = 152
                                                _ = v40
                                            } else {
                                                if true {
                                                    v41 := v1[fi25]
                                                    _ = v41
                                                } else {
                                                    v23 -= 4294967295
                                                }
                                                _ = len(v1)
                                                v42 := GenS2{F0: (v23 * (v23 * v23)), F1: (v4 | v4)}
                                                _ = v42
                                            }
                                            if (v23 >= (v23 % ((2) | 1))) {
                                                fmt.Println(v2)
                                            }
                                            for fi43 := 0; fi43 < 9; fi43++ {
                                                if (255 < 127) {
                                                    v22 -= float64(float32(0.0))
                                                    if true {
                                                        v23 -= 42
                                                        fmt.Println(v22, v4)
                                                    }
                                                    var v44 uint = 4294967295
                                                    _ = v44
                                                } else {
                                                    var v45 int64 = 365
                                                    var v46 string = "foo bar"
                                                    _ = v45
                                                    _ = v46
                                                }
                                                switch ("" + "hello") {
                                                case "\n\t":
                                                    var v47 int = -570
                                                    _ = v47
                                                case "foo bar":
                                                    switch "foo bar" {
                                                    case "":
                                                        delete(v3, "hello")
                                                    case "a":
                                                        v48 := GenS6{F0: 100, F1: 264, F2: 1}
                                                        _ = v48
                                                    case "hello":
                                                        var v49 int16 = -1
                                                        _ = v49
                                                    }
                                                }
                                                v2 += (1.5 / 1.5)
                                                switch 9223372036854775807 {
                                                case 2147483647:
                                                    _ = len(v1)
                                                case -1:
                                                    v22 += 1e10
                                                case -2147483648:
                                                    v50 := []int{}
                                                    _ = v50
                                                }
                                                _ = fi43
                                            }
                                        }
                                        fmt.Println(v2, v2, v4)
                                    }
                                } else {
                                    var v51 int16 = 1000
                                    for fi52 := 0; fi52 < 12; fi52++ {
                                        v53 := GenS5{F0: (fi52 >> ((4294967295) & 63)), F1: ("foo bar" + ("" + "a")), F2: 255}
                                        v54 := []int{int(v4), 127, (fi25 ^ fi52)}
                                        v54 = append(v54, (fi25 - fi25))
                                        if (fi25 <= fi52) {
                                            switch "\n\t" {
                                            case "foo bar":
                                                switch (("a" + "foo bar") + "foo bar") {
                                                case "\n\t":
                                                    var v55 float32 = ((v2 - v2) / (v2 * float32(1e10)))
                                                    _ = v55
                                                default:
                                                    v2 -= v2
                                                }
                                            case "a":
                                                v2 -= (v2 + 0.1)
                                            case "":
                                                var v56 bool = true
                                                _ = v56
                                            }
                                            v57 := GenS2{F0: 108, F1: v4}
                                            _ = v57
                                        }
                                        v23 -= (^(v23))
                                        var v58 float64 = (v22 / (((v22 + 1.5) * (-2.25 + float64(((v2 + v2) - ((1.5 / v2) + float32(v22)))))) + float64((v2 + v2))))
                                        var v59 float64 = (3.14159 * (((float64(v2) - 0.1) - v58) * 3.14159))
                                        _ = fi52
                                        _ = v53
                                        _ = v54
                                        _ = v58
                                        _ = v59
                                    }
                                    _ = v51
                                }
                            } else {
                                switch ("\n\t" + "") {
                                case "hello":
                                    v2 += -2.25
                                case "":
                                    if (!false) {
                                        if true {
                                            var v60 int16 = -308
                                            v22 = v22
                                            v60 = v60
                                            _ = v60
                                        } else {
                                            if (1 < 103) {
                                                v61 := GenS7{F0: 65535, F1: -128}
                                                if false {
                                                    v62 := GenS7{F0: 0, F1: -20}
                                                    _ = v62
                                                } else {
                                                    for fi63 := 0; fi63 < 7; fi63++ {
                                                        var v64 int = (fi25 % ((fi25) | 1))
                                                        _ = fi63
                                                        _ = v64
                                                    }
                                                    var v65 uint32 = v23
                                                    _ = v65
                                                }
                                                _ = v61
                                            } else {
                                                v2 = v2
                                                v66 := GenS8{F0: 239, F1: 42}
                                                fmt.Println(v22, fi25)
                                                _ = v66
                                            }
                                            v67 := v3["foo bar"]
                                            _ = v67
                                        }
                                        fmt.Println(v23, v4)
                                        var v68 uint = 4294967295
                                        _ = v68
                                    }
                                }
                            }
                        }
                        var v69 uint64 = 18446744073709551615
                        _ = fi25
                        _ = v69
                    }
                    var v70 uint16 = 65535
                    _ = v24
                    _ = v70
                } else {
                    for fi71 := 0; fi71 < 17; fi71++ {
                        switch int(v4) {
                        case 32768:
                            v23 = 467
                        }
                        if (!false) {
                            v2 = 0.1
                        } else {
                            v1[fi71] = true
                            v72 := GenS5{F0: (fi71 >> ((2) & 63)), F1: "foo bar", F2: 0}
                            _ = v72
                        }
                        _ = fi71
                    }
                    fmt.Println(v4, v2)
                }
                v73 := []int{2147483647, 9223372036854775807, -1, 4294967295}
                _ = v73
            }
        case "a":
            v22 = (v22 + v22)
        }
    case -500:
        switch 4294967295 {
        case 1:
            if (false && ((!false) || (!false))) {
                v74 := GenS9{F0: 65535, F1: 100, F2: 2}
                _ = v74
            } else {
                if true {
                    for fi75 := 0; fi75 < 15; fi75++ {
                        v2 = (1.5 / v2)
                        switch (fi75 << ((2) & 63)) {
                        case 0:
                            var v76 int8 = 127
                            _ = v76
                        case 9223372036854775807:
                            if false {
                                var v77 float64 = 0.0
                                var v78 bool = false
                                var v79 bool = (genFn0() > 42)
                                _ = v77
                                _ = v78
                                _ = v79
                            } else {
                                if false {
                                    if true {
                                        v80 := GenS10{F0: 1, F1: v23, F2: 42}
                                        for fi81 := 0; fi81 < 10; fi81++ {
                                            v2 += v2
                                            var v82 uint32 = (v23 + (v23 * v23))
                                            _ = fi81
                                            _ = v82
                                        }
                                        _ = v80
                                    }
                                } else {
                                    v23 = v23
                                }
                                if (1 > 85) {
                                    v1[fi75] = true
                                } else {
                                    v83 := GenS4{F0: "a", F1: 2, F2: -1}
                                    _ = v83
                                }
                            }
                        default:
                            switch 32768 {
                            case -115:
                                for fi84 := 0; fi84 < 11; fi84++ {
                                    var v85 float32 = 1.5
                                    fmt.Println(v4)
                                    _ = fi84
                                    _ = v85
                                }
                            case 9223372036854775807:
                                v2 = 0.0
                            case 1:
                                if (false || (!false)) {
                                    fmt.Println(v4)
                                    for fi86 := 0; fi86 < 2; fi86++ {
                                        v4 = 2147483647
                                        v87 := GenS3{F0: 255, F1: (false && (!(false || (287 != 205))))}
                                        _ = fi86
                                        _ = v87
                                    }
                                } else {
                                    v2 = v2
                                    fmt.Println(v4)
                                    v88 := []int{(fi75 - (fi75 ^ fi75)), (-(fi75)), fi75, (fi75 << ((0) & 63))}
                                    _ = v88
                                }
                            default:
                                if false {
                                    v23 = 0
                                }
                            }
                        }
                        _ = fi75
                    }
                    v89 := GenS5{F0: 128, F1: "a", F2: 18446744073709551615}
                    for fi90 := 0; fi90 < 15; fi90++ {
                        switch (-(fi90)) {
                        case 128:
                            v22 += v22
                        case -9223372036854775808:
                            v91 := GenS1{F0: v4, F1: 1, F2: 4294967295}
                            _ = v91
                        default:
                            switch ("" + "a") {
                            case "a":
                                for fi92 := 0; fi92 < 3; fi92++ {
                                    v93 := v1[(^(fi90))]
                                    v2 = v2
                                    if (!v93) {
                                        v94 := GenS9{F0: 18446744073709551615, F1: -1, F2: 940}
                                        _ = v94
                                    }
                                    var v95 float32 = float32(float64(v2))
                                    _ = fi92
                                    _ = v93
                                    _ = v95
                                }
                            case "hello":
                                if (90 == 0) {
                                    if false {
                                        v1[fi90] = false
                                        v23 += v23
                                    }
                                    var v96 uint64 = 2
                                    _ = v96
                                }
                            case "":
                                fmt.Println(v22, v23, v22)
                            }
                        }
                        _ = fi90
                    }
                    _ = v89
                } else {
                    var v97 uint64 = 42
                    _ = v97
                }
            }
        case -648:
            v4 = v4
        case 0:
            fmt.Println(v22)
        }
    case 42:
        var v98 bool = true
        _ = v98
    default:
        v2 -= 3.14159
    }
    switch -205 {
    case 128:
        for fi99 := 0; fi99 < 12; fi99++ {
            v23 = 366
            v100 := GenS11{F0: genFn1(), F1: v4}
            _ = fi99
            _ = v100
        }
    case 127:
        v101 := []int{}
        _ = v101
    case -144:
        var v102 float32 = v2
        _ = v102
    default:
        if false {
            v103 := []int{32768}
            v104 := GenS12{F0: 0, F1: 65535, F2: 652}
            _ = v103
            _ = v104
        }
    }
    fmt.Println(v4, v4, v4)
    v22 += float64(((v2 / (0.0 + v2)) / (3.14159 / v2)))
    v23 = v23
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
