package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int64; F1 uint16; F2 int64 }
type GenS5 struct { F0 int16; F1 string; F2 int8 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 bool; F1 bool }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 int32; F1 uint16; F2 string }
type GenS11 struct { F0 int16; F1 int8; F2 int64 }
type GenS12 struct { F0 int; F1 uint8 }
type GenS13 struct { F0 bool }

func main() {
    var v1 int8 = 0
    var v2 float64 = (0.1 * 1.5)
    v3 := make(map[int]int)
    v1 -= v1
    var v4 bool = false
    for fi5 := 0; fi5 < 9; fi5++ {
        var v6 uint32 = 65535
        for fi7 := 0; fi7 < 19; fi7++ {
            v8 := GenS1{F0: 1, F1: 154}
            v9 := GenS2{F0: 753}
            _ = fi7
            _ = v8
            _ = v9
        }
        v2 -= float64(((((1e10 + -2.25) - (((-2.25 - (-2.25 + (((-2.25 / 3.14159) * 1e10) / 0.1))) + 3.14159) + 3.14159)) - float32(0.1)) - 3.14159))
        _ = fi5
        _ = v6
    }
    var v10 uint64 = 735
    var v11 uint = 1
    for fi12 := 0; fi12 < 11; fi12++ {
        v13 := GenS2{F0: uint32(v10)}
        v11 = (v11 ^ (v11 >> (((v11 / (((v11 ^ 0)) | 1))) & 63)))
        _ = fi12
        _ = v13
    }
    var v14 int = -277
    v15 := []int{}
    switch ("a" + ("\n\t" + ("a" + "hello"))) {
    case "\n\t":
        if true {
            fmt.Println(v4, v4, v11)
        } else {
            v1 += -1
        }
    default:
        v2 = 0.1
    }
    var v16 uint16 = 858
    if ((!true) || ((!(((v2 < v2) || v4) && true)) || (!v4))) {
        if len(v15) > 0 {
            v15[0] = v14
        }
    } else {
        switch "foo bar" {
        case "\n\t":
            if v4 {
                fmt.Println(v1, v16)
                v17 := GenS2{F0: 2}
                _ = v17
            } else {
                switch v14 {
                case -1:
                    v11 += v11
                case 32767:
                    for fi18 := 0; fi18 < 10; fi18++ {
                        v19 := GenS3{F0: 859}
                        switch ("\n\t" + "a") {
                        case "a":
                            if (!v4) {
                                var v20 uint64 = v10
                                v21 := GenS2{F0: 643}
                                v16 += v16
                                _ = v20
                                _ = v21
                            }
                        default:
                            if (-50 > (v1 | (v1 * 127))) {
                                v22 := GenS4{F0: 980, F1: v16, F2: 654}
                                if v4 {
                                    fmt.Println(v14)
                                } else {
                                    switch "\n\t" {
                                    case "hello":
                                        v23 := GenS4{F0: -610, F1: v16, F2: -1}
                                        _ = v23
                                    case "foo bar":
                                        if (v10 <= v10) {
                                            v24 := GenS5{F0: -526, F1: ("foo bar" + ""), F2: v1}
                                            _ = v24
                                        }
                                    default:
                                        switch (("\n\t" + "foo bar") + "") {
                                        case "hello":
                                            if (0.1 != (1e10 + float32(v2))) {
                                                v16 = 1
                                            } else {
                                                v4 = true
                                                switch -627 {
                                                case 0:
                                                    fmt.Println(v10, v2, v11)
                                                case 4294967295:
                                                    if true {
                                                        v1 += (v1 ^ (v1 | (^(v1))))
                                                        var v25 int16 = -1
                                                        _ = v25
                                                    } else {
                                                        v15 = append(v15, -1)
                                                    }
                                                case -2147483648:
                                                    for fi26 := 0; fi26 < 6; fi26++ {
                                                        v2 += v2
                                                        _ = fi26
                                                    }
                                                default:
                                                    if v4 {
                                                        if (237 >= 255) {
                                                            var v27 string = "a"
                                                            if (v14 == fi18) {
                                                                v28 := GenS6{F0: 127}
                                                                var v29 uint8 = 255
                                                                _ = v28
                                                                _ = v29
                                                            } else {
                                                                for fi30 := 0; fi30 < 12; fi30++ {
                                                                    switch 550 {
                                                                    case 32767:
                                                                        if true {
                                                                            fmt.Println(v2)
                                                                        }
                                                                    case 1:
                                                                        var v31 float32 = 1e10
                                                                        v16 += v16
                                                                        _ = len(v3)
                                                                        v10 -= (v10 | (v10 % (((^(v10))) | 1)))
                                                                        var v32 uint = (v11 * v11)
                                                                        switch (((v27 + (("a" + "") + ((("hello" + v27) + "") + ("\n\t" + (v27 + (v27 + (("a" + (v27 + ((((v27 + (v27 + v27)) + ("\n\t" + "a")) + v27) + v27))) + ("hello" + "foo bar")))))))) + v27) + ("" + v27)) {
                                                                        case "foo bar":
                                                                            v16 -= v16
                                                                        case "hello":
                                                                            var v33 string = v27
                                                                            _ = v33
                                                                        }
                                                                        var v34 string = "foo bar"
                                                                        if v4 {
                                                                            fmt.Println(v14, v34)
                                                                        } else {
                                                                            v27 = "a"
                                                                        }
                                                                        v35 := GenS3{F0: 65535}
                                                                        _ = v31
                                                                        _ = v32
                                                                        _ = v34
                                                                        _ = v35
                                                                    case 32768:
                                                                        switch (fi30 % (((v14 / ((310) | 1))) | 1)) {
                                                                        case -9223372036854775808:
                                                                            if (!(!(v1 != v1))) {
                                                                                if len(v15) > 0 {
                                                                                    v15[0] = v14
                                                                                }
                                                                                v14 = (v14 ^ -324)
                                                                            }
                                                                        case -327:
                                                                            if (v10 != v10) {
                                                                                var v36 int8 = (v1 + v1)
                                                                                switch v27 {
                                                                                case "\n\t":
                                                                                    var v37 int = (fi18 | (^(fi18)))
                                                                                    _ = v37
                                                                                case "a":
                                                                                    if v4 {
                                                                                        v4 = v4
                                                                                        var v38 bool = v4
                                                                                        var v39 float64 = v2
                                                                                        fmt.Println(v1)
                                                                                        _ = v38
                                                                                        _ = v39
                                                                                    } else {
                                                                                        if true {
                                                                                            fmt.Println(v14, v2, v36)
                                                                                            if (v16 < (v16 / ((1) | 1))) {
                                                                                                v40 := GenS7{F0: 7}
                                                                                                if (v11 < (v11 - 348)) {
                                                                                                    var v41 string = (v27 + "hello")
                                                                                                    _ = v41
                                                                                                }
                                                                                                var v42 uint8 = 255
                                                                                                _ = v40
                                                                                                _ = v42
                                                                                            } else {
                                                                                                if (v4 && (v2 != 0.0)) {
                                                                                                    v16 = 32767
                                                                                                    var v43 uint32 = 4294967295
                                                                                                    var v44 bool = ((^(v36)) < (v36 + (^(v1))))
                                                                                                    switch v27 {
                                                                                                    case "\n\t":
                                                                                                        var v45 int64 = 0
                                                                                                        _ = v45
                                                                                                    case "foo bar":
                                                                                                        var v46 uint8 = 52
                                                                                                        _ = v46
                                                                                                    default:
                                                                                                        switch fi18 {
                                                                                                        case -442:
                                                                                                            var v47 string = v27
                                                                                                            _ = v47
                                                                                                        case -1:
                                                                                                            if (127 < 128) {
                                                                                                                if (64 <= 1) {
                                                                                                                    v48 := GenS8{F0: v4, F1: v44}
                                                                                                                    fmt.Println(v27)
                                                                                                                    _ = v48
                                                                                                                } else {
                                                                                                                    if (!(!false)) {
                                                                                                                        fmt.Println(v43, v14, v14)
                                                                                                                        fmt.Println(v4, v4, v14)
                                                                                                                    } else {
                                                                                                                        if ((v27 + "") == v27) {
                                                                                                                            switch (fi18 & fi30) {
                                                                                                                            case -2147483648:
                                                                                                                                var v49 int32 = -1
                                                                                                                                _ = v49
                                                                                                                            case -122:
                                                                                                                                v1 -= v36
                                                                                                                            case 128:
                                                                                                                                v3[1] = -1
                                                                                                                            }
                                                                                                                            switch (fi30 * fi18) {
                                                                                                                            case 42:
                                                                                                                                var v50 uint = 18446744073709551615
                                                                                                                                _ = v50
                                                                                                                            default:
                                                                                                                                if v4 {
                                                                                                                                    var v51 uint64 = 899
                                                                                                                                    fmt.Println(fi30)
                                                                                                                                    _ = v51
                                                                                                                                } else {
                                                                                                                                    switch -9223372036854775808 {
                                                                                                                                    case -572:
                                                                                                                                        _ = len(v3)
                                                                                                                                    case 4294967295:
                                                                                                                                        v52 := GenS5{F0: 0, F1: "", F2: int8(v11)}
                                                                                                                                        _ = v52
                                                                                                                                    }
                                                                                                                                }
                                                                                                                            }
                                                                                                                            var v53 uint64 = (v10 << ((uint(v43)) & 63))
                                                                                                                            _ = v53
                                                                                                                        } else {
                                                                                                                            switch (fi18 * fi18) {
                                                                                                                            case 408:
                                                                                                                                var v54 uint8 = 224
                                                                                                                                _ = v54
                                                                                                                            case 37:
                                                                                                                                v55 := GenS9{F0: 2}
                                                                                                                                _ = v55
                                                                                                                            default:
                                                                                                                                v11 += v11
                                                                                                                            }
                                                                                                                            switch -1 {
                                                                                                                            case 2147483647:
                                                                                                                                v16 -= (v16 & v16)
                                                                                                                            default:
                                                                                                                                v56 := GenS9{F0: 152}
                                                                                                                                _ = v56
                                                                                                                            }
                                                                                                                        }
                                                                                                                    }
                                                                                                                    switch "foo bar" {
                                                                                                                    case "a":
                                                                                                                        if v4 {
                                                                                                                            _ = len(v3)
                                                                                                                            switch (fi30 << (((v11 + v11)) & 63)) {
                                                                                                                            case 1:
                                                                                                                                var v57 float32 = ((1e10 * (float32(float64((((-2.25 / float32(v2)) + (-2.25 / 1.0)) / 1.0))) + float32(1.5))) / 1e10)
                                                                                                                                _ = v57
                                                                                                                            case 32767:
                                                                                                                                var v58 float32 = 1e10
                                                                                                                                _ = v58
                                                                                                                            }
                                                                                                                            fmt.Println(v11)
                                                                                                                        } else {
                                                                                                                            delete(v3, (fi30 % ((127) | 1)))
                                                                                                                        }
                                                                                                                    case "":
                                                                                                                        v1 -= (v36 - -128)
                                                                                                                    case "hello":
                                                                                                                        var v59 int64 = -9223372036854775808
                                                                                                                        _ = v59
                                                                                                                    }
                                                                                                                }
                                                                                                                v60 := GenS10{F0: -978, F1: 255, F2: (v27 + v27)}
                                                                                                                _ = v60
                                                                                                            }
                                                                                                        default:
                                                                                                            switch "" {
                                                                                                            case "":
                                                                                                                if (-1 <= 1000) {
                                                                                                                    fmt.Println(fi30, v14, v36)
                                                                                                                }
                                                                                                            case "a":
                                                                                                                var v61 float64 = ((v2 - v2) - (v2 + (v2 + (v2 - v2))))
                                                                                                                _ = v61
                                                                                                            default:
                                                                                                                var v62 float32 = 1.5
                                                                                                                _ = v62
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                    _ = v43
                                                                                                    _ = v44
                                                                                                } else {
                                                                                                    var v63 bool = (67 < (v36 >> (((v11 >> ((v11) & 63))) & 63)))
                                                                                                    var v64 int64 = 9223372036854775807
                                                                                                    if v4 {
                                                                                                        v14 = v14
                                                                                                        if v4 {
                                                                                                            v14 = fi30
                                                                                                            if ((v27 <= "foo bar") || true) {
                                                                                                                switch (v14 >> ((179) & 63)) {
                                                                                                                case 0:
                                                                                                                    var v65 uint = v11
                                                                                                                    _ = v65
                                                                                                                default:
                                                                                                                    fmt.Println(v27, v2)
                                                                                                                }
                                                                                                                v3[(fi30 / ((fi30) | 1))] = (fi18 | fi18)
                                                                                                            } else {
                                                                                                                switch (v27 + (("hello" + "") + "")) {
                                                                                                                case "":
                                                                                                                    v10 += v10
                                                                                                                case "\n\t":
                                                                                                                    var v66 int32 = 809
                                                                                                                    _ = v66
                                                                                                                default:
                                                                                                                    v3[(fi18 & 32767)] = (fi30 - -2147483648)
                                                                                                                }
                                                                                                                fmt.Println(v1, fi30, v27)
                                                                                                                var v67 uint8 = 19
                                                                                                                _ = v67
                                                                                                            }
                                                                                                            v3[-1] = v14
                                                                                                        }
                                                                                                        var v68 bool = v63
                                                                                                        _ = v68
                                                                                                    }
                                                                                                    _ = v63
                                                                                                    _ = v64
                                                                                                }
                                                                                            }
                                                                                            v69 := GenS11{F0: 0, F1: v1, F2: -925}
                                                                                            _ = v69
                                                                                        }
                                                                                    }
                                                                                default:
                                                                                    var v70 bool = v4
                                                                                    _ = v70
                                                                                }
                                                                                _ = v36
                                                                            }
                                                                        case 127:
                                                                            v14 = fi18
                                                                        }
                                                                    default:
                                                                        var v71 float32 = (-2.25 + 3.14159)
                                                                        var v72 float32 = (1e10 * v71)
                                                                        _ = v71
                                                                        _ = v72
                                                                    }
                                                                    var v73 int32 = -2147483648
                                                                    _ = fi30
                                                                    _ = v73
                                                                }
                                                                if (false && v4) {
                                                                    _ = len(v15)
                                                                }
                                                                if v4 {
                                                                    v11 -= v11
                                                                    v14 = 679
                                                                }
                                                            }
                                                            _ = v27
                                                        }
                                                        if (((v11 & (v11 | 2)) < v11) && v4) {
                                                            for fi74 := 0; fi74 < 3; fi74++ {
                                                                v75 := GenS8{F0: false, F1: (2147483648 >= 2147483647)}
                                                                _ = fi74
                                                                _ = v75
                                                            }
                                                            for fi76 := 0; fi76 < 11; fi76++ {
                                                                v3[v14] = (fi76 - (v14 >> ((888) & 63)))
                                                                var v77 uint64 = (v10 % (((^(v10))) | 1))
                                                                fmt.Println(v11)
                                                                _ = fi76
                                                                _ = v77
                                                            }
                                                            switch "\n\t" {
                                                            case "hello":
                                                                v78 := GenS9{F0: 128}
                                                                _ = v78
                                                            default:
                                                                fmt.Println(v1)
                                                            }
                                                        } else {
                                                            var v79 float32 = (3.14159 * 0.0)
                                                            _ = v79
                                                        }
                                                        if (!(!(!(2147483647 == 1)))) {
                                                            v80 := GenS10{F0: -2147483648, F1: (v16 >> (((v11 / ((v11) | 1))) & 63)), F2: ("\n\t" + ("hello" + ""))}
                                                            v81 := GenS11{F0: 388, F1: 1, F2: -1}
                                                            _ = v80
                                                            _ = v81
                                                        }
                                                    } else {
                                                        v11 = v11
                                                        fmt.Println(v10, v16, fi18)
                                                    }
                                                }
                                                v82 := GenS4{F0: -972, F1: (v16 | v16), F2: 0}
                                                _ = v82
                                            }
                                        case "foo bar":
                                            v15 = append(v15, fi18)
                                        case "a":
                                            fmt.Println(v11)
                                        }
                                    }
                                }
                                v10 += 18446744073709551615
                                _ = v22
                            }
                        }
                        v83 := v3[42]
                        _ = fi18
                        _ = v19
                        _ = v83
                    }
                case 128:
                    for fi84 := 0; fi84 < 19; fi84++ {
                        v4 = (32767 < (v16 << ((2) & 63)))
                        for fi85 := 0; fi85 < 5; fi85++ {
                            v86 := GenS12{F0: fi84, F1: 1}
                            if (!(1 > -1)) {
                                v4 = false
                            } else {
                                v1 = v1
                            }
                            switch "\n\t" {
                            case "hello":
                                fmt.Println(fi84)
                            case "":
                                if true {
                                    if v4 {
                                        v87 := GenS13{F0: true}
                                        _ = v87
                                    } else {
                                        var v88 int8 = (v1 ^ (v1 & v1))
                                        v89 := GenS2{F0: 4294967295}
                                        _ = v88
                                        _ = v89
                                    }
                                    fmt.Println(v1)
                                }
                            }
                            switch fi85 {
                            case 32767:
                                fmt.Println(v14, v14)
                            case -120:
                                v1 = (v1 & (v1 >> (((v11 * v11)) & 63)))
                            case 936:
                                var v90 uint32 = uint32(v16)
                                _ = v90
                            default:
                                var v91 uint16 = 532
                                _ = v91
                            }
                            _ = fi85
                            _ = v86
                        }
                        _ = fi84
                    }
                }
                var v92 int16 = -32768
                _ = v92
            }
        default:
            v93 := GenS8{F0: true, F1: (128 >= 1)}
            _ = v93
        }
        if v4 {
            for fi94 := 0; fi94 < 7; fi94++ {
                var v95 float64 = 1.5
                v96 := GenS2{F0: 742}
                var v97 int64 = -9223372036854775808
                _ = fi94
                _ = v95
                _ = v96
                _ = v97
            }
            fmt.Println(v10)
            delete(v3, (v14 - -2147483648))
        } else {
            var v98 int16 = 1000
            _ = v98
        }
        for fi99 := 0; fi99 < 11; fi99++ {
            v100 := GenS6{F0: 0}
            v101 := GenS2{F0: 158}
            _ = fi99
            _ = v100
            _ = v101
        }
    }
    v11 = 65535
    var v102 string = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v102:", v102)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v102
    _ = v11
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
