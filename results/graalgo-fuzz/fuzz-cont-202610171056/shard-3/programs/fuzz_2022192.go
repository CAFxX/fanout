package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int8; F1 int; F2 string }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int; F1 int64 }
type GenS7 struct { F0 uint16; F1 int32; F2 uint }
type GenS8 struct { F0 uint8; F1 int8 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 uint16 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]int)
    v3 := []int{}
    var v4 float64 = 0.0
    var v5 uint8 = 187
    for fi6 := 0; fi6 < 1; fi6++ {
        var v7 uint32 = 159
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 19; fi8++ {
        _ = len(v1)
        _ = fi8
    }
    for fi9 := 0; fi9 < 6; fi9++ {
        v10 := GenS1{F0: false}
        _ = fi9
        _ = v10
    }
    var v11 bool = true
    if (688 != 1) {
        var v12 string = "hello"
        if (false && v11) {
            if (!v11) {
                switch 128 {
                case 0:
                    v4 += 3.14159
                default:
                    var v13 string = (("hello" + v12) + "hello")
                    _ = v13
                }
            }
            _ = len(v3)
            v5 = v5
        } else {
            switch 42 {
            case 4294967295:
                var v14 int64 = -9223372036854775808
                _ = v14
            case 866:
                for fi15 := 0; fi15 < 2; fi15++ {
                    v16 := GenS2{F0: 1000}
                    v17 := GenS3{F0: 0}
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            default:
                fmt.Println(v4, v11, v4)
            }
            if (!(0 <= 2)) {
                if false {
                    for fi18 := 0; fi18 < 1; fi18++ {
                        for fi19 := 0; fi19 < 13; fi19++ {
                            v12 = (("" + v12) + v12)
                            v4 += v4
                            v20 := GenS4{F0: -95, F1: fi18, F2: v12}
                            _ = fi19
                            _ = v20
                        }
                        if ((916 > 0) || (!(((v5 <= 127) || (!false)) && v11))) {
                            fmt.Println(v5, v12, v5)
                        } else {
                            v21 := GenS5{F0: (v5 - (v5 * 127))}
                            _ = v21
                        }
                        for fi22 := 0; fi22 < 2; fi22++ {
                            v23 := GenS1{F0: true}
                            _ = fi22
                            _ = v23
                        }
                        _ = fi18
                    }
                    switch "foo bar" {
                    case "":
                        v24 := GenS3{F0: -61}
                        _ = v24
                    case "\n\t":
                        switch -2147483648 {
                        case 2147483647:
                            switch 0 {
                            case 32767:
                                var v25 uint16 = 0
                                _ = v25
                            case -663:
                                if v11 {
                                    for fi26 := 0; fi26 < 3; fi26++ {
                                        var v27 int64 = 42
                                        _ = fi26
                                        _ = v27
                                    }
                                } else {
                                    switch (v12 + ("\n\t" + (v12 + ("" + v12)))) {
                                    case "\n\t":
                                        v4 = (v4 - v4)
                                    case "":
                                        v4 += ((v4 + (v4 - v4)) / ((0.0 + 1.5) / v4))
                                    case "hello":
                                        if true {
                                            var v28 float32 = -2.25
                                            fmt.Println(v11, v28, v5)
                                            _ = v28
                                        }
                                    }
                                    fmt.Println(v4, v12)
                                }
                            case 128:
                                if v11 {
                                    var v29 uint16 = 318
                                    _ = v29
                                }
                            }
                        default:
                            for fi30 := 0; fi30 < 3; fi30++ {
                                switch (v12 + (v12 + v12)) {
                                case "foo bar":
                                    var v31 string = v12
                                    _ = v31
                                case "\n\t":
                                    var v32 float32 = 1e10
                                    _ = v32
                                default:
                                    if (!v11) {
                                        if true {
                                            _ = len(v1)
                                            v12 = ((v12 + ("hello" + ("" + v12))) + v12)
                                            v33 := GenS5{F0: (v5 >> ((18446744073709551615) & 63))}
                                            _ = v33
                                        }
                                        v11 = (true && (42 == 410))
                                        for fi34 := 0; fi34 < 1; fi34++ {
                                            if v11 {
                                                v3 = append(v3, (fi34 - fi34))
                                                if (v11 || false) {
                                                    fmt.Println(v4, v12)
                                                    if v11 {
                                                        if v11 {
                                                            v35 := v2[fi30]
                                                            _ = v35
                                                        } else {
                                                            fmt.Println(v11, v4)
                                                        }
                                                        v4 = v4
                                                        v36 := GenS2{F0: 42}
                                                        _ = v36
                                                    }
                                                    v37 := GenS5{F0: v5}
                                                    _ = v37
                                                } else {
                                                    if (-1 == 42) {
                                                        switch 156 {
                                                        case 4294967295:
                                                            fmt.Println(v11, fi34)
                                                        case 127:
                                                            var v38 uint32 = 972
                                                            _ = v38
                                                        case 1:
                                                            v5 = v5
                                                        default:
                                                            v39 := GenS4{F0: -36, F1: fi30, F2: ("\n\t" + ((("foo bar" + v12) + v12) + (v12 + v12)))}
                                                            _ = v39
                                                        }
                                                        if (v4 > 1e10) {
                                                            switch (fi34 * fi34) {
                                                            case 0:
                                                                if (v11 || false) {
                                                                    v40 := v1[fi34]
                                                                    fmt.Println(v40, fi34, fi30)
                                                                    fmt.Println(v40)
                                                                    _ = v40
                                                                } else {
                                                                    v4 = float64((1.5 / -2.25))
                                                                }
                                                                v4 -= v4
                                                            default:
                                                                fmt.Println(fi30, fi34, v12)
                                                            }
                                                        } else {
                                                            if (-1 == 42) {
                                                                v41 := GenS2{F0: -32768}
                                                                switch ("foo bar" + (v12 + "\n\t")) {
                                                                case "\n\t":
                                                                    var v42 string = ("hello" + v12)
                                                                    _ = v42
                                                                case "foo bar":
                                                                    v43 := GenS6{F0: (fi34 + fi30), F1: -9223372036854775808}
                                                                    delete(v1, fi30)
                                                                    v12 = (v12 + v12)
                                                                    _ = v43
                                                                default:
                                                                    var v44 uint = 255
                                                                    _ = v44
                                                                }
                                                                if true {
                                                                    fmt.Println(fi30, v5)
                                                                    var v45 bool = (828 > 18446744073709551615)
                                                                    v5 = 122
                                                                    _ = v45
                                                                }
                                                                _ = v41
                                                            }
                                                        }
                                                        v46 := GenS6{F0: (fi34 - fi30), F1: -235}
                                                        _ = v46
                                                    }
                                                    v4 = v4
                                                    if false {
                                                        v2[fi34] = (fi30 << ((18446744073709551615) & 63))
                                                        v47 := GenS2{F0: 742}
                                                        var v48 bool = (false && true)
                                                        v4 += 1e10
                                                        _ = v47
                                                        _ = v48
                                                    } else {
                                                        if v11 {
                                                            v11 = (!v11)
                                                        }
                                                        var v49 int16 = 1000
                                                        v12 = (v12 + v12)
                                                        _ = v49
                                                    }
                                                }
                                                fmt.Println(fi30, fi34)
                                            }
                                            switch 465 {
                                            case 42:
                                                if false {
                                                    _ = len(v1)
                                                    var v50 uint32 = 2147483647
                                                    _ = v50
                                                } else {
                                                    var v51 int = 32768
                                                    _ = v51
                                                }
                                            }
                                            _ = fi34
                                        }
                                    } else {
                                        v52 := GenS6{F0: (-(fi30)), F1: 0}
                                        _ = v52
                                    }
                                }
                                if (v11 && (2 != 2)) {
                                    var v53 bool = (!false)
                                    switch -9223372036854775808 {
                                    case -577:
                                        var v54 bool = v11
                                        _ = v54
                                    }
                                    v55 := GenS7{F0: 798, F1: 42, F2: 4294967295}
                                    _ = v53
                                    _ = v55
                                } else {
                                    if v11 {
                                        var v56 uint64 = 0
                                        v4 += (v4 + v4)
                                        _ = v56
                                    } else {
                                        fmt.Println(v11)
                                        for fi57 := 0; fi57 < 5; fi57++ {
                                            var v58 uint = 433
                                            fmt.Println(fi57, v4, v11)
                                            var v59 string = v12
                                            _ = fi57
                                            _ = v58
                                            _ = v59
                                        }
                                    }
                                    v12 = "\n\t"
                                    _ = len(v1)
                                }
                                _ = fi30
                            }
                        }
                    case "foo bar":
                        v60 := GenS8{F0: (v5 * v5), F1: -69}
                        _ = v60
                    default:
                        var v61 int = 42
                        _ = v61
                    }
                    var v62 int64 = -9223372036854775808
                    _ = v62
                }
                fmt.Println(v5, v4)
                if ((957 <= 4294967295) || (3.14159 <= (((float64(float32(v4)) / v4) + (1.5 * ((1.5 - (v4 + v4)) + 0.1))) / (0.1 / ((v4 * v4) - (v4 * 3.14159)))))) {
                    if (42 != 985) {
                        var v63 int = -333
                        if v11 {
                            fmt.Println(v12, v12)
                        }
                        v3 = append(v3, v63)
                        _ = v63
                    } else {
                        var v64 bool = v11
                        if (1.5 != float32((v4 - (v4 * (v4 + v4))))) {
                            v5 -= (v5 * 255)
                            fmt.Println(v64, v12, v5)
                            v65 := GenS3{F0: -1}
                            _ = v65
                        } else {
                            v66 := GenS1{F0: true}
                            v4 -= v4
                            fmt.Println(v5, v5, v5)
                            _ = v66
                        }
                        switch 690 {
                        case 128:
                            for fi67 := 0; fi67 < 15; fi67++ {
                                switch fi67 {
                                case 127:
                                    for fi68 := 0; fi68 < 18; fi68++ {
                                        v12 = v12
                                        var v69 int16 = 1
                                        _ = fi68
                                        _ = v69
                                    }
                                default:
                                    for fi70 := 0; fi70 < 4; fi70++ {
                                        var v71 int64 = -1
                                        if len(v3) > 0 {
                                            v3[0] = -282
                                        }
                                        _ = fi70
                                        _ = v71
                                    }
                                }
                                _ = fi67
                            }
                        default:
                            v11 = (false || false)
                        }
                        _ = v64
                    }
                    for fi72 := 0; fi72 < 5; fi72++ {
                        var v73 int8 = 49
                        var v74 bool = (v11 && v11)
                        _ = fi72
                        _ = v73
                        _ = v74
                    }
                    if (!((float64(3.14159) - 3.14159) == 3.14159)) {
                        fmt.Println(v11, v5, v12)
                        for fi75 := 0; fi75 < 18; fi75++ {
                            var v76 float32 = float32(float64(0.0))
                            _ = fi75
                            _ = v76
                        }
                        for fi77 := 0; fi77 < 5; fi77++ {
                            fmt.Println(v4, v4, v4)
                            _ = fi77
                        }
                    }
                } else {
                    v3 = append(v3, -2147483648)
                }
            }
        }
        _ = v12
    } else {
        switch 411 {
        case 4294967295:
            delete(v2, 9223372036854775807)
        case 309:
            switch 4294967295 {
            case -687:
                var v78 uint32 = 323
                _ = v78
            case 42:
                v79 := GenS5{F0: v5}
                _ = v79
            case 1:
                fmt.Println(v5, v4, v4)
            }
        default:
            for fi80 := 0; fi80 < 16; fi80++ {
                switch "" {
                case "":
                    v5 += v5
                default:
                    var v81 float64 = v4
                    _ = v81
                }
                _ = fi80
            }
        }
        v82 := GenS1{F0: (2 != 4294967295)}
        _ = v82
    }
    if true {
        for fi83 := 0; fi83 < 8; fi83++ {
            for fi84 := 0; fi84 < 5; fi84++ {
                if true {
                    var v85 string = "a"
                    var v86 uint8 = (v5 + v5)
                    _ = v85
                    _ = v86
                } else {
                    if v11 {
                        switch ((("foo bar" + ("\n\t" + "foo bar")) + ("\n\t" + "\n\t")) + "foo bar") {
                        case "foo bar":
                            v4 = v4
                        case "\n\t":
                            var v87 uint8 = (v5 * v5)
                            _ = v87
                        default:
                            switch fi84 {
                            case 728:
                                _ = len(v2)
                                var v88 int32 = -1
                                _ = v88
                            case 317:
                                fmt.Println(v11, fi84, v5)
                            }
                        }
                    }
                    switch (fi84 & (fi83 << ((42) & 63))) {
                    case 32768:
                        v89 := v1[fi83]
                        _ = v89
                    case -446:
                        fmt.Println(v11, v4, v4)
                        var v90 float32 = ((1.5 * 1.5) - float32((v4 - v4)))
                        _ = v90
                    case 743:
                        switch -2147483648 {
                        case -57:
                            if (-1 < 724) {
                                v91 := GenS9{F0: 42}
                                v11 = v11
                                _ = v91
                            } else {
                                switch ("foo bar" + ("" + "foo bar")) {
                                case "a":
                                    var v92 int8 = -78
                                    _ = v92
                                case "":
                                    v93 := GenS9{F0: 230}
                                    _ = v93
                                case "foo bar":
                                    if ((-912 > int16(fi84)) && v11) {
                                        fmt.Println(v11, v11)
                                    } else {
                                        var v94 int16 = 733
                                        _ = v94
                                    }
                                }
                                v5 = 127
                            }
                        default:
                            if (0 < 0) {
                                v5 = 0
                                switch "\n\t" {
                                case "foo bar":
                                    v4 = v4
                                }
                                var v95 string = ((("a" + "\n\t") + (("foo bar" + ("" + (("\n\t" + "") + ("\n\t" + "foo bar")))) + "foo bar")) + "")
                                _ = v95
                            }
                        }
                    }
                }
                v2[(^(fi84))] = fi83
                _ = fi84
            }
            _ = fi83
        }
    }
    fmt.Println(v11, v4)
    v1[-350] = v11
    if ((4294967295 == 18446744073709551615) || v11) {
        v11 = true
        fmt.Println(v11, v11)
    } else {
        v5 += v5
        for fi96 := 0; fi96 < 6; fi96++ {
            v4 = float64(1.5)
            v5 = v5
            v11 = (!v11)
            _ = fi96
        }
    }
    v5 += v5
    for fi97 := 0; fi97 < 20; fi97++ {
        var v98 uint64 = 2
        var v99 float64 = 1e10
        var v100 float64 = float64((float32(-2.25) + 0.1))
        _ = fi97
        _ = v100
        _ = v98
        _ = v99
    }
    v11 = (false && v11)
    var v101 uint16 = 0
    v102 := GenS10{F0: (v101 * 512)}
    switch 2147483647 {
    case 163:
        v5 -= (v5 | (^(v5)))
    default:
        if len(v3) > 0 {
            v3[0] = 2147483647
        }
    }
    v4 -= v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v101:", v101)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v101
    _ = v102
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
