package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 uint16; F1 int32 }
type GenS4 struct { F0 uint; F1 bool; F2 int32 }
type GenS5 struct { F0 uint32; F1 int8; F2 uint8 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint; F1 uint16 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 int; F1 bool }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 int8; F1 int8 }
type GenS13 struct { F0 uint8 }

func main() {
    v1 := []uint32{65535, 2147483647}
    v2 := make(map[uint32]bool)
    var v3 uint32 = 42
    var v4 float64 = -2.25
    v5 := make(map[string]string)
    var v6 string = ("\n\t" + "")
    fmt.Println(v4, v3, v4)
    var v7 float32 = -2.25
    v6 = (v6 + v6)
    for fi8 := 0; fi8 < 11; fi8++ {
        if true {
            v7 = float32(float64(v7))
        } else {
            v9 := GenS1{F0: "foo bar"}
            if (11 < 18446744073709551615) {
                var v10 int16 = 0
                var v11 int16 = (v10 << ((974) & 63))
                var v12 uint16 = 32767
                _ = v10
                _ = v11
                _ = v12
            } else {
                v13 := v2[(v3 << ((897) & 63))]
                for fi14 := 0; fi14 < 10; fi14++ {
                    v15 := GenS2{F0: 0, F1: 32767}
                    switch fi14 {
                    case 42:
                        v16 := GenS3{F0: 65535, F1: 0}
                        _ = v16
                    case 370:
                        var v17 int16 = 1
                        _ = v17
                    case 871:
                        v18 := GenS3{F0: 32767, F1: -2147483648}
                        _ = v18
                    }
                    var v19 int32 = -655
                    _ = fi14
                    _ = v15
                    _ = v19
                }
                v13 = v13
                _ = v13
            }
            _ = v9
        }
        var v20 int32 = -1
        _ = fi8
        _ = v20
    }
    if (273 != 65535) {
        v21 := GenS4{F0: 4294967295, F1: false, F2: 1}
        for fi22 := 0; fi22 < 9; fi22++ {
            for fi23 := 0; fi23 < 6; fi23++ {
                v4 = v4
                var v24 uint = 0
                _ = fi23
                _ = v24
            }
            v25 := GenS2{F0: 42, F1: 451}
            _ = fi22
            _ = v25
        }
        v6 = "\n\t"
        _ = v21
    } else {
        v1 = append(v1, (^(v3)))
        _ = len(v2)
    }
    fmt.Println(v7, v7)
    if len(v1) > 0 {
        v1[0] = v3
    }
    var v26 string = v6
    v6 = ((v6 + "\n\t") + "foo bar")
    _ = len(v1)
    var v27 uint16 = 42
    fmt.Println(v26)
    var v28 int = 189
    switch (v28 * (v28 / ((v28) | 1))) {
    case -847:
        if (42 != 111) {
            if (!true) {
                var v29 int16 = 42
                var v30 string = v26
                var v31 int64 = -1
                _ = v29
                _ = v30
                _ = v31
            } else {
                for fi32 := 0; fi32 < 13; fi32++ {
                    v33 := GenS5{F0: v3, F1: 1, F2: 255}
                    var v34 bool = (42 < 65535)
                    for fi35 := 0; fi35 < 13; fi35++ {
                        if (-522 >= 0) {
                            v28 = 32768
                            v1 = append(v1, 4294967295)
                        }
                        if (651 <= fi32) {
                            var v36 string = v6
                            switch v36 {
                            case "":
                                var v37 uint8 = 128
                                _ = v37
                            default:
                                v2[0] = (((fi35 ^ (fi35 & (fi35 - fi32))) > (^(fi35))) || true)
                            }
                            if (v3 != v3) {
                                if v34 {
                                    switch (v36 + v36) {
                                    case "\n\t":
                                        v38 := GenS4{F0: 255, F1: v34, F2: -787}
                                        _ = v38
                                    case "":
                                        v36 = v26
                                    case "foo bar":
                                        var v39 int16 = int16(v27)
                                        _ = v39
                                    default:
                                        if (v7 != (v7 * float32(((0.0 * v4) + v4)))) {
                                            fmt.Println(v34)
                                            v40 := GenS6{F0: v26}
                                            _ = v40
                                        } else {
                                            switch (v6 + v26) {
                                            case "hello":
                                                v41 := GenS7{F0: (v27 & 85)}
                                                _ = v41
                                            case "foo bar":
                                                switch ("a" + ((v36 + ("" + v6)) + v26)) {
                                                case "foo bar":
                                                    v27 += v27
                                                case "":
                                                    if v34 {
                                                        v42 := GenS8{F0: 92, F1: 65535}
                                                        var v43 uint32 = v3
                                                        _ = v42
                                                        _ = v43
                                                    }
                                                case "a":
                                                    var v44 int = (fi32 % ((fi32) | 1))
                                                    _ = v44
                                                }
                                            case "\n\t":
                                                v45 := GenS8{F0: 20, F1: (v27 * v27)}
                                                _ = v45
                                            }
                                            if v34 {
                                                v46 := GenS7{F0: 1}
                                                _ = v46
                                            } else {
                                                var v47 int = (fi32 - (fi35 & v28))
                                                if false {
                                                    v47 = (v28 / ((v47) | 1))
                                                    v48 := GenS1{F0: v26}
                                                    var v49 uint8 = 1
                                                    _ = v48
                                                    _ = v49
                                                }
                                                v3 = (v3 - (v3 | v3))
                                                _ = v47
                                            }
                                            v26 = "hello"
                                            v4 = 0.0
                                            switch fi32 {
                                            case 32767:
                                                var v50 int8 = 42
                                                _ = v50
                                            }
                                        }
                                    }
                                    if v34 {
                                        v4 = v4
                                        v34 = (532 != 255)
                                    } else {
                                        v6 = (v6 + ((((("" + v6) + "") + v26) + v26) + v36))
                                        switch (fi35 ^ (v28 + 1)) {
                                        case 4294967295:
                                            var v51 uint16 = v27
                                            _ = v51
                                        case 42:
                                            if (816 <= -727) {
                                                fmt.Println(fi35, v36)
                                            }
                                            if ((((v7 * 0.0) / ((v7 * v7) - (0.0 * float32(v4)))) * 3.14159) >= ((-2.25 / v7) + v7)) {
                                                fmt.Println(v34, v34, v3)
                                                var v52 uint16 = v27
                                                _ = v52
                                            }
                                            var v53 uint32 = (v3 & v3)
                                            v36 = v6
                                            _ = v53
                                        default:
                                            v3 = (^(v3))
                                        }
                                    }
                                }
                                switch v26 {
                                case "\n\t":
                                    v54 := GenS8{F0: 1, F1: v27}
                                    _ = v54
                                case "hello":
                                    v55 := GenS6{F0: (v36 + ((v36 + v26) + ("a" + (v6 + "foo bar"))))}
                                    _ = v55
                                default:
                                    var v56 uint64 = 4294967295
                                    _ = v56
                                }
                            } else {
                                if (v36 != ((v26 + (v26 + ("" + "\n\t"))) + "hello")) {
                                    _ = len(v1)
                                    fmt.Println(v27)
                                }
                                switch "a" {
                                case "\n\t":
                                    fmt.Println(v7, v36)
                                case "":
                                    if len(v1) > 0 {
                                        v1[0] = v3
                                    }
                                case "foo bar":
                                    var v57 int8 = 127
                                    _ = v57
                                default:
                                    if true {
                                        v58 := GenS9{F0: (v26 + v6)}
                                        fmt.Println(v6, fi32)
                                        var v59 uint8 = 255
                                        _ = v58
                                        _ = v59
                                    }
                                }
                                v5[(v6 + (v26 + ((v36 + ((v36 + v36) + ("" + v36))) + ("a" + "hello"))))] = (v36 + v36)
                                var v60 uint = 4294967295
                                v34 = (195 < 2)
                                switch -624 {
                                case -1:
                                    fmt.Println(v28, v27, v6)
                                }
                                v28 = (fi32 * 9223372036854775807)
                                _ = v60
                            }
                            switch "\n\t" {
                            case "hello":
                                v3 -= 153
                            case "foo bar":
                                v61 := GenS10{F0: fi32, F1: (v27 < v27)}
                                _ = v61
                            default:
                                v34 = (9223372036854775807 != 42)
                            }
                            v36 = v36
                            _ = v36
                        } else {
                            v62 := GenS11{F0: 42}
                            if (65535 < 0) {
                                var v63 int8 = int8(v27)
                                var v64 bool = (!true)
                                v6 = "foo bar"
                                _ = v63
                                _ = v64
                            } else {
                                var v65 int = (v28 * 32767)
                                v28 = (fi32 * v28)
                                v66 := GenS4{F0: 4294967295, F1: false, F2: 1}
                                _ = v65
                                _ = v66
                            }
                            var v67 uint = 35
                            _ = v62
                            _ = v67
                        }
                        v26 = v6
                        _ = fi35
                    }
                    _ = fi32
                    _ = v33
                    _ = v34
                }
                var v68 float32 = (((1e10 / float32(1.5)) - (v7 * v7)) + float32((v4 / 1.5)))
                v27 = v27
                _ = v68
            }
        }
    default:
        var v69 string = ((v6 + v6) + "hello")
        _ = v69
    }
    for fi70 := 0; fi70 < 6; fi70++ {
        for fi71 := 0; fi71 < 13; fi71++ {
            switch -847 {
            case 42:
                v6 = ("" + (v26 + ((v6 + v6) + "")))
            case -332:
                if len(v1) > 0 {
                    v1[0] = (v3 - (v3 / (((v3 | v3)) | 1)))
                }
            case 373:
                var v72 uint16 = 32767
                _ = v72
            default:
                var v73 uint = 0
                _ = v73
            }
            var v74 int32 = 1
            if true {
                if ((((v4 / v4) + ((0.0 / ((v4 - (3.14159 * 0.0)) / (-2.25 - ((v4 - (v4 / v4)) * v4)))) * v4)) <= ((0.0 * v4) * v4)) && (v26 > v26)) {
                    var v75 int8 = -1
                    if (v75 > -1) {
                        v76 := GenS7{F0: (v27 ^ (v27 << ((4294967295) & 63)))}
                        v75 += v75
                        if (v26 < v26) {
                            delete(v2, 2147483648)
                            switch "a" {
                            case "a":
                                v74 += (-(v74))
                            case "\n\t":
                                delete(v2, 30)
                            default:
                                var v77 float32 = v7
                                _ = v77
                            }
                            if true {
                                delete(v5, (v26 + (v26 + "")))
                                v5[v6] = "a"
                            } else {
                                if false {
                                    v4 = -2.25
                                    switch ((v6 + v6) + "\n\t") {
                                    case "hello":
                                        if ((0.0 + (v7 / (0.1 / v7))) > 0.1) {
                                            if false {
                                                v78 := GenS1{F0: v6}
                                                var v79 uint8 = 128
                                                _ = v78
                                                _ = v79
                                            }
                                            _ = len(v5)
                                        } else {
                                            v5["\n\t"] = v26
                                            v80 := v5["a"]
                                            var v81 int16 = 384
                                            _ = v80
                                            _ = v81
                                        }
                                    case "foo bar":
                                        fmt.Println(v6, fi70, v26)
                                    }
                                    if false {
                                        v27 += 42
                                        switch fi70 {
                                        case 32768:
                                            if (false && true) {
                                                var v82 string = "foo bar"
                                                var v83 uint8 = 128
                                                v84 := v2[v3]
                                                _ = v82
                                                _ = v83
                                                _ = v84
                                            }
                                        case -964:
                                            fmt.Println(v27, v6)
                                        }
                                        v27 += (^(v27))
                                    } else {
                                        switch v6 {
                                        case "":
                                            var v85 uint32 = 417
                                            _ = v85
                                        case "foo bar":
                                            v86 := GenS5{F0: (v3 | 1), F1: v75, F2: 42}
                                            _ = v86
                                        case "\n\t":
                                            v5[""] = "hello"
                                        }
                                    }
                                } else {
                                    fmt.Println(v6, v4, v4)
                                }
                                if true {
                                    delete(v5, v6)
                                }
                                switch v28 {
                                case -94:
                                    var v87 uint64 = 49
                                    _ = v87
                                case -183:
                                    v74 -= v74
                                }
                                var v88 int16 = 42
                                var v89 int = v28
                                _ = len(v1)
                                v28 += (fi71 - v28)
                                fmt.Println(v26, v4, v88)
                                _ = v88
                                _ = v89
                            }
                        }
                        _ = v76
                    }
                    if (!false) {
                        v90 := GenS12{F0: v75, F1: (v75 << ((0) & 63))}
                        v91 := GenS1{F0: v6}
                        if false {
                            v3 = (v3 * 2147483647)
                            switch (v26 + ("" + (((v6 + v6) + v6) + ("" + v6)))) {
                            case "":
                                v92 := GenS13{F0: 128}
                                _ = v92
                            case "\n\t":
                                v28 = -874
                            }
                            var v93 uint16 = v27
                            _ = v93
                        } else {
                            var v94 int32 = (^(v74))
                            var v95 int16 = 32767
                            _ = v94
                            _ = v95
                        }
                        var v96 int64 = 595
                        _ = v90
                        _ = v91
                        _ = v96
                    } else {
                        _ = len(v1)
                        if (v74 <= 2147483647) {
                            if true {
                                v6 = v26
                            }
                            v97 := GenS1{F0: ((v26 + (((v6 + v6) + v6) + (((v26 + "foo bar") + "hello") + (((v6 + v6) + "") + ("hello" + ((v6 + ("hello" + v26)) + v6)))))) + v26)}
                            _ = v97
                        }
                    }
                    fmt.Println(v26)
                    var v98 uint32 = (v3 >> ((1) & 63))
                    v99 := v5[v26]
                    v98 += v3
                    if ((!(true && false)) && false) {
                        switch (-(fi71)) {
                        case 32768:
                            var v100 uint32 = 1
                            _ = v100
                        case 32767:
                            v27 = v27
                        }
                        v4 = (v4 + ((3.14159 - v4) - float64((1.5 * v7))))
                    }
                    _ = v75
                    _ = v98
                    _ = v99
                }
                switch "\n\t" {
                case "":
                    delete(v2, v3)
                case "\n\t":
                    v1 = append(v1, v3)
                case "a":
                    v7 += (1e10 + 1.5)
                }
                v1 = append(v1, 0)
            }
            _ = fi71
            _ = v74
        }
        var v101 uint8 = 0
        _ = fi70
        _ = v101
    }
    var v102 uint32 = (v3 + (v3 << ((18446744073709551615) & 63)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v102:", v102)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v102
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
