package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int; F1 int16; F2 int64 }
type GenS4 struct { F0 int64; F1 bool; F2 uint }
type GenS5 struct { F0 int64; F1 int64 }
type GenS6 struct { F0 bool; F1 int64 }
type GenS7 struct { F0 int; F1 int64 }
type GenS8 struct { F0 uint8; F1 uint32; F2 uint64 }
type GenS9 struct { F0 int16; F1 int64 }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 uint16 }

func main() {
    var v1 int32 = 42
    var v2 uint = 18446744073709551615
    var v3 uint32 = 291
    v4 := []int{}
    v5 := []int{}
    switch 127 {
    case 127:
        if false {
            var v6 uint64 = 820
            switch 32768 {
            case 2147483647:
                for fi7 := 0; fi7 < 4; fi7++ {
                    var v8 int32 = v1
                    v4 = append(v4, fi7)
                    for fi9 := 0; fi9 < 16; fi9++ {
                        v8 += (v1 >> (((v2 / (((v2 ^ 2)) | 1))) & 63))
                        _ = fi9
                    }
                    _ = fi7
                    _ = v8
                }
            case 0:
                v1 = v1
            case 128:
                var v10 int16 = 32767
                _ = v10
            default:
                v6 = v6
            }
            for fi11 := 0; fi11 < 5; fi11++ {
                var v12 uint32 = (^(v3))
                for fi13 := 0; fi13 < 1; fi13++ {
                    _ = len(v4)
                    _ = fi13
                }
                var v14 int = 9223372036854775807
                _ = fi11
                _ = v12
                _ = v14
            }
            _ = v6
        } else {
            var v15 uint8 = 164
            _ = v15
        }
    }
    if (!("a" < "foo bar")) {
        v16 := GenS1{F0: 42, F1: true}
        _ = v16
    }
    v2 += (v2 << ((42) & 63))
    for fi17 := 0; fi17 < 15; fi17++ {
        for fi18 := 0; fi18 < 11; fi18++ {
            v2 -= 255
            _ = len(v5)
            _ = fi18
        }
        if (!true) {
            if (!(fi17 != 9223372036854775807)) {
                for fi19 := 0; fi19 < 12; fi19++ {
                    v1 = v1
                    _ = fi19
                }
                v3 -= v3
            } else {
                var v20 bool = true
                _ = v20
            }
            var v21 int = (^(fi17))
            _ = v21
        } else {
            v22 := GenS1{F0: 127, F1: false}
            _ = v22
        }
        if true {
            v3 = (^(v3))
            var v23 int64 = 1
            if false {
                for fi24 := 0; fi24 < 13; fi24++ {
                    var v25 uint32 = (v3 & (v3 / (((v3 + v3)) | 1)))
                    v26 := GenS2{F0: v1}
                    var v27 uint16 = 255
                    _ = len(v4)
                    switch "" {
                    case "foo bar":
                        fmt.Println(v1, v3)
                    default:
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        v28 := GenS3{F0: 9223372036854775807, F1: -750, F2: v23}
                        v4 = append(v4, (fi17 << ((0) & 63)))
                        _ = v28
                    }
                    _ = fi24
                    _ = v25
                    _ = v26
                    _ = v27
                }
                v1 = v1
                fmt.Println(v3)
            } else {
                for fi29 := 0; fi29 < 5; fi29++ {
                    fmt.Println(fi17, v1, v1)
                    _ = fi29
                }
            }
            _ = v23
        } else {
            var v30 uint16 = 1
            if true {
                var v31 uint8 = 255
                var v32 uint16 = (v30 & (v30 - (v30 % ((v30) | 1))))
                v1 -= 570
                _ = v31
                _ = v32
            }
            _ = v30
        }
        _ = fi17
    }
    for fi33 := 0; fi33 < 6; fi33++ {
        v1 += (v1 / ((v1) | 1))
        switch 42 {
        case 32767:
            for fi34 := 0; fi34 < 12; fi34++ {
                switch fi34 {
                case 128:
                    if false {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    } else {
                        if true {
                            v1 = v1
                            fmt.Println(fi33, v2, v2)
                        }
                        if false {
                            if (!true) {
                                if false {
                                    var v35 int = fi33
                                    fmt.Println(v2)
                                    _ = v35
                                } else {
                                    v5 = append(v5, (-(fi33)))
                                }
                                v3 = v3
                                fmt.Println(v2)
                            }
                        } else {
                            if ((0 >= -1) || true) {
                                if len(v4) > 0 {
                                    v4[0] = (fi33 / (((fi33 | (fi34 + fi33))) | 1))
                                }
                                v2 = v2
                                var v36 int64 = -9223372036854775808
                                _ = v36
                            } else {
                                if false {
                                    if false {
                                        if false {
                                            v2 -= v2
                                            if (true || (32767 == 42)) {
                                                v37 := GenS4{F0: 9223372036854775807, F1: true, F2: (v2 % ((v2) | 1))}
                                                _ = v37
                                            }
                                        }
                                        var v38 uint16 = 42
                                        _ = v38
                                    }
                                } else {
                                    switch "\n\t" {
                                    case "\n\t":
                                        var v39 int16 = -32768
                                        _ = v39
                                    default:
                                        fmt.Println(fi33)
                                    }
                                    switch "hello" {
                                    case "\n\t":
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    case "":
                                        var v40 int32 = (-(v1))
                                        _ = v40
                                    case "a":
                                        v41 := GenS3{F0: fi33, F1: -32768, F2: 467}
                                        _ = v41
                                    default:
                                        if true {
                                            switch 0 {
                                            case 2147483647:
                                                switch (("foo bar" + ("foo bar" + "\n\t")) + "a") {
                                                case "a":
                                                    v42 := GenS5{F0: 42, F1: -9223372036854775808}
                                                    _ = v42
                                                case "hello":
                                                    if (-1 > 461) {
                                                        if false {
                                                            v43 := GenS3{F0: (fi33 - fi33), F1: 1000, F2: 9223372036854775807}
                                                            fmt.Println(fi34, fi33)
                                                            _ = len(v5)
                                                            fmt.Println(fi34, v3, fi34)
                                                            _ = v43
                                                        } else {
                                                            switch "a" {
                                                            case "a":
                                                                fmt.Println(v3)
                                                            case "foo bar":
                                                                fmt.Println(fi34, v2, fi34)
                                                            }
                                                        }
                                                    } else {
                                                        v1 = v1
                                                    }
                                                default:
                                                    fmt.Println(fi33)
                                                }
                                                var v44 uint64 = 4294967295
                                                _ = v44
                                            case 317:
                                                var v45 uint16 = 2
                                                _ = v45
                                            case -203:
                                                v1 += (v1 >> ((688) & 63))
                                            }
                                            v3 -= v3
                                            if ((false || false) || (((0.0 * (1.5 - -2.25)) + 1.5) < float32(((3.14159 * (3.14159 + 1e10)) * (3.14159 - ((((float64(1e10) + float64(3.14159)) * 1.5) / (float64(1.5) + float64(0.0))) + (1.5 * ((3.14159 + 1.5) * (-2.25 + (0.0 / float64(1.5))))))))))) {
                                                if true {
                                                    var v46 int = (-(fi33))
                                                    v2 += v2
                                                    v46 = -15
                                                    _ = v46
                                                } else {
                                                    if len(v5) > 0 {
                                                        v5[0] = -2147483648
                                                    }
                                                    _ = len(v4)
                                                    var v47 int = fi34
                                                    _ = v47
                                                }
                                            } else {
                                                fmt.Println(v3, fi33, v3)
                                            }
                                        } else {
                                            v3 = 2147483648
                                            var v48 int32 = v1
                                            _ = v48
                                        }
                                    }
                                    var v49 float64 = -2.25
                                    _ = v49
                                }
                                var v50 float32 = float32((3.14159 + float64(0.1)))
                                var v51 float32 = v50
                                _ = v50
                                _ = v51
                            }
                            fmt.Println(fi33)
                            v2 = 0
                        }
                        v1 = v1
                    }
                }
                var v52 uint16 = 819
                v2 = v2
                v4 = append(v4, (^(fi34)))
                switch "hello" {
                case "\n\t":
                    _ = len(v5)
                case "hello":
                    var v53 int8 = 42
                    _ = v53
                case "foo bar":
                    v2 += v2
                default:
                    _ = len(v4)
                }
                _ = fi34
                _ = v52
            }
        }
        _ = fi33
    }
    v2 = (v2 + (v2 - v2))
    if false {
        if (v2 < (v2 / (((v2 + v2)) | 1))) {
            v54 := GenS6{F0: false, F1: 1}
            for fi55 := 0; fi55 < 20; fi55++ {
                _ = len(v4)
                for fi56 := 0; fi56 < 10; fi56++ {
                    if len(v5) > 0 {
                        v5[0] = (-(fi55))
                    }
                    _ = fi56
                }
                _ = fi55
            }
            _ = v54
        }
        var v57 int = -2147483648
        if (!false) {
            v58 := GenS7{F0: v57, F1: 9223372036854775807}
            v3 -= v3
            switch v57 {
            case 1:
                if true {
                    var v59 float32 = (1e10 / (-2.25 - (float32((3.14159 / 1.0)) + 1.5)))
                    var v60 int64 = -508
                    var v61 uint64 = 0
                    _ = v59
                    _ = v60
                    _ = v61
                } else {
                    var v62 int64 = int64(v57)
                    var v63 uint16 = 255
                    fmt.Println(v3, v62)
                    _ = v62
                    _ = v63
                }
            }
            _ = v58
        }
        _ = v57
    } else {
        for fi64 := 0; fi64 < 2; fi64++ {
            fmt.Println(v3, fi64, v3)
            var v65 string = ("a" + "\n\t")
            for fi66 := 0; fi66 < 13; fi66++ {
                if (!(-9223372036854775808 != -344)) {
                    if (false && true) {
                        switch (v65 + ("a" + (v65 + v65))) {
                        case "":
                            if false {
                                fmt.Println(v3, fi64)
                            } else {
                                if true {
                                    v1 -= (v1 + int32(v3))
                                    v2 += (v2 * (v2 % (((v2 * 0)) | 1)))
                                    v3 = v3
                                    switch fi64 {
                                    case 4294967295:
                                        var v67 float64 = 3.14159
                                        _ = v67
                                    }
                                    if (0 >= 42) {
                                        var v68 bool = (false && (18446744073709551615 > 42))
                                        _ = v68
                                    } else {
                                        var v69 int32 = 42
                                        var v70 int32 = 593
                                        if (1e10 == 3.14159) {
                                            if false {
                                                var v71 string = "\n\t"
                                                v71 = ("" + "foo bar")
                                                _ = v71
                                            }
                                        } else {
                                            if (!(4294967295 != (-(fi66)))) {
                                                v65 = v65
                                                v4 = append(v4, fi66)
                                                if (fi66 >= fi64) {
                                                    fmt.Println(v70, v2)
                                                }
                                                fmt.Println(fi66, v65, fi66)
                                                fmt.Println(fi64)
                                            } else {
                                                v72 := GenS8{F0: 255, F1: 2147483647, F2: 2}
                                                _ = len(v5)
                                                _ = v72
                                            }
                                            var v73 uint32 = (v3 | v3)
                                            _ = v73
                                        }
                                        if true {
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            v1 = v70
                                            if len(v4) > 0 {
                                                v4[0] = fi66
                                            }
                                            if ((v1 - 2147483647) != v69) {
                                                v69 += (v69 ^ (v1 & (v1 - v1)))
                                                v74 := GenS3{F0: -691, F1: -883, F2: -9223372036854775808}
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                _ = v74
                                            } else {
                                                if (((-69 >= 127) && false) && false) {
                                                    var v75 int8 = 127
                                                    v70 -= v69
                                                    var v76 bool = (fi64 > (fi66 / (((fi64 | (fi64 - (fi64 & -113)))) | 1)))
                                                    v69 = (v1 + (v1 + v70))
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                    _ = v75
                                                    _ = v76
                                                }
                                            }
                                            var v77 uint16 = 65535
                                            _ = v77
                                        } else {
                                            if true {
                                                v3 -= 0
                                                v1 += -2147483648
                                                v78 := GenS3{F0: fi66, F1: -336, F2: -1}
                                                _ = v78
                                            } else {
                                                fmt.Println(v65, fi64, v69)
                                                v79 := GenS9{F0: -167, F1: -1}
                                                _ = v79
                                            }
                                            v65 = v65
                                            switch fi64 {
                                            case 32768:
                                                v80 := GenS1{F0: 0, F1: (uint16(v2) >= 390)}
                                                _ = v80
                                            case -767:
                                                v4 = append(v4, -2147483648)
                                                if (!true) {
                                                    var v81 bool = (1 > -1)
                                                    _ = len(v4)
                                                    _ = v81
                                                }
                                            default:
                                                var v82 bool = true
                                                _ = v82
                                            }
                                        }
                                        _ = v69
                                        _ = v70
                                    }
                                    v1 += v1
                                } else {
                                    var v83 float64 = 1e10
                                    fmt.Println(v1, fi64)
                                    _ = v83
                                }
                                v84 := GenS10{F0: 2}
                                v85 := GenS4{F0: 718, F1: (!((v3 >> ((v2) & 63)) <= v3)), F2: (v2 + (v2 / ((66) | 1)))}
                                _ = v84
                                _ = v85
                            }
                        case "foo bar":
                            v86 := GenS11{F0: 329}
                            _ = v86
                        default:
                            if len(v4) > 0 {
                                v4[0] = fi66
                            }
                        }
                        switch "a" {
                        case "":
                            if false {
                                if (!(true || true)) {
                                    v87 := GenS3{F0: (^(fi66)), F1: 1000, F2: 1}
                                    v3 += 65535
                                    _ = v87
                                }
                                var v88 float64 = -2.25
                                if (true && false) {
                                    v89 := GenS3{F0: fi66, F1: 1, F2: 42}
                                    var v90 int = (fi66 / ((-2147483648) | 1))
                                    _ = v89
                                    _ = v90
                                }
                                _ = v88
                            }
                        case "a":
                            var v91 float32 = 1e10
                            _ = v91
                        case "foo bar":
                            v1 -= v1
                            switch (fi64 >> ((v2) & 63)) {
                            case 1:
                                switch fi64 {
                                case 4294967295:
                                    v2 -= 0
                                case 1:
                                    var v92 uint8 = 175
                                    _ = v92
                                case 32767:
                                    var v93 int = (^(fi66))
                                    _ = v93
                                }
                            case 127:
                                v3 -= 118
                            case -2147483648:
                                if true {
                                    v65 = v65
                                } else {
                                    v2 += v2
                                    var v94 uint = v2
                                    _ = v94
                                }
                            }
                            switch (^(fi64)) {
                            case 568:
                                _ = len(v4)
                            }
                            if true {
                                v95 := GenS4{F0: 42, F1: false, F2: 172}
                                var v96 uint32 = v3
                                v65 = "\n\t"
                                v96 += (v96 + v96)
                                _ = v95
                                _ = v96
                            } else {
                                var v97 uint16 = 399
                                var v98 int32 = (v1 >> (((v2 / ((v2) | 1))) & 63))
                                if (1e10 >= 0.0) {
                                    v99 := GenS11{F0: (v97 * v97)}
                                    v100 := GenS11{F0: (v97 - v97)}
                                    v65 = v65
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    fmt.Println(v3, v97, v97)
                                    _ = v100
                                    _ = v99
                                } else {
                                    var v101 uint32 = (v3 % ((v3) | 1))
                                    v101 -= v3
                                    _ = v101
                                }
                                _ = v97
                                _ = v98
                            }
                            v3 += v3
                        }
                    } else {
                        var v102 uint64 = 453
                        _ = v102
                    }
                    var v103 int = 42
                    _ = v103
                } else {
                    switch v65 {
                    case "hello":
                        _ = len(v4)
                    case "a":
                        var v104 int32 = (v1 / ((42) | 1))
                        _ = v104
                    default:
                        fmt.Println(v65, v65, fi66)
                    }
                    fmt.Println(v3)
                }
                _ = fi66
            }
            _ = fi64
            _ = v65
        }
    }
    var v105 uint = (v2 >> (((v2 & v2)) & 63))
    for fi106 := 0; fi106 < 16; fi106++ {
        v4 = append(v4, fi106)
        _ = fi106
    }
    var v107 float32 = (0.1 - (0.1 - (0.1 - 1.5)))
    var v108 uint32 = 651
    if (-32768 > 0) {
        fmt.Println(v107)
    }
    v109 := GenS7{F0: 32768, F1: -1}
    if false {
        for fi110 := 0; fi110 < 14; fi110++ {
            _ = len(v4)
            _ = fi110
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    v105 = (v105 << ((v105) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v105:", v105)
    fmt.Println("v107:", v107)
    fmt.Println("v108:", v108)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v105
    _ = v107
    _ = v108
    _ = v109
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
