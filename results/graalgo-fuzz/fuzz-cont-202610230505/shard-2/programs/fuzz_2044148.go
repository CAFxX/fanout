package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64; F1 uint32; F2 uint32 }
type GenS4 struct { F0 int64; F1 int16; F2 int8 }
type GenS5 struct { F0 int64; F1 bool }
type GenS6 struct { F0 int; F1 string }
func genFn0(p0 string, p1 int32) int16 {
    v7 := GenS1{F0: 768}
    v8 := GenS1{F0: 42}
    for fi9 := 0; fi9 < 3; fi9++ {
        v10 := []int{(^(fi9))}
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 1000
}
func genFn1(p0 string, p1 uint8, p2 int32) uint {
    v11 := []int{}
    p2 = p2
    var v12 int64 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return 65535
}

func main() {
    var v1 int = 128
    var v2 int8 = 1
    v3 := []uint32{}
    var v4 int8 = -1
    v5 := make(map[uint32]string)
    var v6 float32 = 1.5
    if true {
        var v13 uint16 = 42
        v14 := v5[2147483647]
        _ = v13
        _ = v14
    }
    if true {
        var v15 uint32 = 776
        _ = v15
    } else {
        v4 = (v2 | 46)
        delete(v5, 65535)
        switch ("a" + ((("hello" + ("foo bar" + ("foo bar" + "\n\t"))) + ("\n\t" + "foo bar")) + ("" + ("foo bar" + "\n\t")))) {
        case "foo bar":
            if (-737 <= -1) {
                switch "" {
                case "foo bar":
                    switch ("\n\t" + "\n\t") {
                    case "\n\t":
                        for fi16 := 0; fi16 < 8; fi16++ {
                            var v17 uint32 = 1
                            v4 = (v2 << ((2) & 63))
                            _ = fi16
                            _ = v17
                        }
                    case "a":
                        var v18 int32 = 1
                        _ = v18
                    }
                case "\n\t":
                    switch 9223372036854775807 {
                    case -9223372036854775808:
                        fmt.Println(v4, v4)
                    }
                default:
                    for fi19 := 0; fi19 < 5; fi19++ {
                        for fi20 := 0; fi20 < 8; fi20++ {
                            var v21 uint8 = 42
                            fmt.Println(v21, fi20, v4)
                            fmt.Println(v6, fi20, v1)
                            _ = fi20
                            _ = v21
                        }
                        if (fi19 == 2147483647) {
                            for fi22 := 0; fi22 < 12; fi22++ {
                                v23 := GenS2{F0: (^(v1))}
                                if (2 < 4294967295) {
                                    if ("foo bar" == "hello") {
                                        v4 = (v4 & (v2 >> ((1) & 63)))
                                        fmt.Println(v1, v1, v2)
                                        v1 = (fi19 / ((-2147483648) | 1))
                                    }
                                    var v24 uint = 255
                                    _ = v24
                                } else {
                                    if true {
                                        if false {
                                            var v25 float64 = (((1e10 * 3.14159) + (0.0 - 3.14159)) + -2.25)
                                            var v26 uint64 = 4294967295
                                            switch fi19 {
                                            case 128:
                                                fmt.Println(v26, v4, v26)
                                            case -2147483648:
                                                v4 -= (v2 >> ((650) & 63))
                                            default:
                                                if ((v4 * 33) < -1) {
                                                    v5[65535] = ("hello" + "hello")
                                                } else {
                                                    switch -1 {
                                                    case -2147483648:
                                                        if (2147483647 != 580) {
                                                            fmt.Println(v1, v26)
                                                            if len(v3) > 0 {
                                                                v3[0] = 65535
                                                            }
                                                        } else {
                                                            v27 := GenS1{F0: 494}
                                                            var v28 int8 = 0
                                                            _ = v27
                                                            _ = v28
                                                        }
                                                    }
                                                    if len(v3) > 0 {
                                                        v3[0] = 149
                                                    }
                                                }
                                            }
                                            v1 -= (-(fi19))
                                            _ = v25
                                            _ = v26
                                        }
                                        var v29 int16 = 1000
                                        var v30 uint64 = 645
                                        if false {
                                            if (true || true) {
                                                v4 -= (v4 - 1)
                                                v31 := GenS1{F0: 262}
                                                _ = v31
                                            } else {
                                                v30 = (v30 % (((v30 * 42)) | 1))
                                                if true {
                                                    v2 = (v2 / ((v4) | 1))
                                                    switch fi19 {
                                                    case 130:
                                                        v32 := v5[2]
                                                        _ = v32
                                                    case 42:
                                                        v29 -= (-(v29))
                                                    default:
                                                        var v33 int64 = -1
                                                        _ = v33
                                                    }
                                                    var v34 int32 = 42
                                                    _ = v34
                                                }
                                            }
                                        }
                                        v6 = v6
                                        var v35 uint8 = 0
                                        if true {
                                            v29 -= v29
                                            var v36 int = fi19
                                            if false {
                                                v37 := GenS1{F0: 2}
                                                v1 = v1
                                                var v38 float64 = -2.25
                                                switch "hello" {
                                                case "foo bar":
                                                    if (0 <= 114) {
                                                        v29 = 42
                                                        v36 += v36
                                                    }
                                                case "":
                                                    _ = len(v3)
                                                case "\n\t":
                                                    if len(v3) > 0 {
                                                        v3[0] = 574
                                                    }
                                                }
                                                _ = v37
                                                _ = v38
                                            }
                                            var v39 int = 32768
                                            _ = v36
                                            _ = v39
                                        }
                                        _ = v29
                                        _ = v30
                                        _ = v35
                                    } else {
                                        var v40 int64 = -72
                                        v40 += (v40 * (^(v40)))
                                        _ = v40
                                    }
                                    if (2 <= 24) {
                                        v41 := GenS1{F0: 649}
                                        var v42 uint = 1
                                        _ = v41
                                        _ = v42
                                    } else {
                                        if ((v2 % ((v4) | 1)) < v2) {
                                            v43 := GenS1{F0: 460}
                                            switch (fi19 / (((^(v1))) | 1)) {
                                            case 0:
                                                v2 = v4
                                                _ = len(v5)
                                                if (v2 >= -1) {
                                                    var v44 uint8 = 2
                                                    switch 163 {
                                                    case -2147483648:
                                                        v45 := GenS3{F0: 1, F1: 348, F2: 2147483647}
                                                        if false {
                                                            var v46 uint = 4294967295
                                                            _ = v46
                                                        } else {
                                                            v47 := GenS4{F0: -377, F1: 42, F2: (v2 ^ (^(v4)))}
                                                            if ((-(fi19)) <= v1) {
                                                                fmt.Println(v2, v1, fi19)
                                                            }
                                                            v6 = (v6 / (v6 - (v6 * v6)))
                                                            _ = v47
                                                        }
                                                        _ = v45
                                                    case -9223372036854775808:
                                                        if (((1e10 + (1e10 * v6)) == (v6 - (v6 + (3.14159 - v6)))) && false) {
                                                            v44 -= v44
                                                        }
                                                    case -24:
                                                        var v48 float32 = v6
                                                        _ = v48
                                                    default:
                                                        v1 = -1
                                                        fmt.Println(fi22, v2, fi19)
                                                    }
                                                    switch v1 {
                                                    case 9223372036854775807:
                                                        fmt.Println(fi22, fi22)
                                                    case 127:
                                                        var v49 uint16 = 774
                                                        _ = v49
                                                    case 2147483647:
                                                        switch "a" {
                                                        case "foo bar":
                                                            switch v1 {
                                                            case 42:
                                                                switch "" {
                                                                case "foo bar":
                                                                    v1 += fi19
                                                                case "":
                                                                    v6 = ((1.5 / (((1.5 + (v6 + v6)) - v6) + (v6 * v6))) - v6)
                                                                default:
                                                                    fmt.Println(v1)
                                                                }
                                                            default:
                                                                var v50 uint32 = 1
                                                                _ = v50
                                                            }
                                                        case "a":
                                                            var v51 bool = (0.1 == (1.5 / 1e10))
                                                            _ = v51
                                                        case "\n\t":
                                                            switch fi19 {
                                                            case 1:
                                                                v52 := v5[65535]
                                                                _ = v52
                                                            case -9223372036854775808:
                                                                var v53 string = "hello"
                                                                _ = v53
                                                            }
                                                        default:
                                                            v2 += (v4 * v2)
                                                        }
                                                    }
                                                    _ = v44
                                                }
                                            case 9223372036854775807:
                                                fmt.Println(v1, v2)
                                            case -1:
                                                v54 := GenS5{F0: 1, F1: (4294967295 > 0)}
                                                _ = v54
                                            }
                                            if ((true || ((!(!false)) || false)) && false) {
                                                if false {
                                                    fmt.Println(v1, fi19, v2)
                                                }
                                            } else {
                                                fmt.Println(v6, v4, v6)
                                                v6 = (0.1 * (1e10 / (float32(1e10) - v6)))
                                                v2 += (v2 / ((-128) | 1))
                                                v5[42] = "\n\t"
                                                switch ("hello" + "foo bar") {
                                                case "hello":
                                                    switch "\n\t" {
                                                    case "a":
                                                        var v55 int8 = 127
                                                        _ = v55
                                                    case "hello":
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                    case "foo bar":
                                                        v1 = fi22
                                                    default:
                                                        switch ((("" + "foo bar") + ("foo bar" + "")) + "\n\t") {
                                                        case "hello":
                                                            if (true && true) {
                                                                var v56 string = "a"
                                                                v1 += v1
                                                                v1 += v1
                                                                _ = v56
                                                            }
                                                        case "":
                                                            var v57 bool = (((0.1 + (float32(float64(v6)) * v6)) - v6) < 3.14159)
                                                            _ = v57
                                                        }
                                                        v2 += (-(v2))
                                                    }
                                                default:
                                                    v6 -= (v6 * (0.0 + v6))
                                                }
                                            }
                                            _ = v43
                                        } else {
                                            switch ("" + ("" + "")) {
                                            case "":
                                                v58 := GenS4{F0: -9223372036854775808, F1: 32767, F2: -128}
                                                _ = v58
                                            }
                                        }
                                        v59 := GenS2{F0: (fi22 + (-(v1)))}
                                        _ = v59
                                    }
                                    if len(v3) > 0 {
                                        v3[0] = 314
                                    }
                                    if (!true) {
                                        switch fi19 {
                                        case 0:
                                            v1 -= -2147483648
                                        default:
                                            v5[2147483648] = ""
                                        }
                                    } else {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        switch v1 {
                                        case 770:
                                            fmt.Println(v2, v1, fi19)
                                        case 0:
                                            var v60 int8 = 42
                                            v60 = 52
                                            v6 += float32(1e10)
                                            _ = v60
                                        case 9223372036854775807:
                                            v1 -= fi22
                                        default:
                                            var v61 uint = 408
                                            _ = v61
                                        }
                                    }
                                    switch "a" {
                                    case "a":
                                        fmt.Println(v2, fi19, fi19)
                                    case "":
                                        var v62 int = (fi22 * fi19)
                                        _ = v62
                                    case "foo bar":
                                        _ = len(v5)
                                    }
                                }
                                fmt.Println(v2)
                                fmt.Println(v4)
                                var v63 uint8 = 0
                                _ = fi22
                                _ = v23
                                _ = v63
                            }
                            var v64 string = "hello"
                            _ = v64
                        } else {
                            v1 = (v1 << ((756) & 63))
                        }
                        if false {
                            v2 = 127
                            for fi65 := 0; fi65 < 9; fi65++ {
                                switch "foo bar" {
                                case "":
                                    v4 = int8(fi19)
                                default:
                                    if false {
                                        var v66 uint8 = 8
                                        v67 := GenS2{F0: 565}
                                        _ = v66
                                        _ = v67
                                    }
                                }
                                if len(v3) > 0 {
                                    v3[0] = 49
                                }
                                if false {
                                    if len(v3) > 0 {
                                        v3[0] = 2147483647
                                    }
                                    if false {
                                        _ = len(v3)
                                    } else {
                                        var v68 uint32 = 670
                                        var v69 int16 = genFn0(("a" + "hello"), -2147483648)
                                        _ = len(v3)
                                        _ = v68
                                        _ = v69
                                    }
                                    _ = len(v5)
                                } else {
                                    if len(v3) > 0 {
                                        v3[0] = 0
                                    }
                                    v4 -= (v2 | (v4 ^ -46))
                                }
                                _ = fi65
                            }
                        } else {
                            v4 -= v2
                        }
                        _ = fi19
                    }
                }
            } else {
                var v70 float64 = 1e10
                fmt.Println(v4, v70, v2)
                _ = v70
            }
        case "":
            if ((false && (0 > 80)) && (((v6 * ((1e10 / (0.0 * v6)) / -2.25)) * ((v6 * 0.1) + v6)) != v6)) {
                v6 -= v6
                v2 = v2
                var v71 int32 = 1
                _ = v71
            }
        default:
            v72 := GenS6{F0: 32767, F1: ("foo bar" + "\n\t")}
            _ = v72
        }
    }
    fmt.Println(v1)
    _ = len(v5)
    v2 += v4
    _ = len(v5)
    fmt.Println(v4)
    fmt.Println(v2)
    v4 = 32
    v73 := GenS5{F0: -572, F1: true}
    switch v1 {
    case -2147483648:
        var v74 uint64 = 42
        _ = v74
    case -9223372036854775808:
        if false {
            switch (v1 * (v1 + v1)) {
            case 127:
                var v75 float32 = 0.0
                _ = v75
            case -9223372036854775808:
                for fi76 := 0; fi76 < 16; fi76++ {
                    delete(v5, 4294967295)
                    _ = fi76
                }
            }
            if false {
                var v77 float32 = 0.1
                if len(v3) > 0 {
                    v3[0] = 2147483647
                }
                _ = v77
            }
            for fi78 := 0; fi78 < 7; fi78++ {
                var v79 bool = true
                _ = fi78
                _ = v79
            }
        } else {
            for fi80 := 0; fi80 < 4; fi80++ {
                v6 = (1.5 * 1e10)
                _ = fi80
            }
            if (true && false) {
                var v81 uint8 = 0
                _ = v81
            }
        }
    default:
        if (545 != -2147483648) {
            v4 -= -1
        } else {
            v4 -= (-(v2))
        }
    }
    switch -1 {
    case 32768:
        for fi82 := 0; fi82 < 15; fi82++ {
            v83 := v5[809]
            switch (v83 + ("foo bar" + "a")) {
            case "a":
                v4 = (-(v4))
            case "foo bar":
                fmt.Println(v83)
            }
            var v84 float32 = 0.1
            _ = fi82
            _ = v83
            _ = v84
        }
    default:
        v85 := v5[2]
        _ = v85
    }
    var v86 string = ""
    var v87 bool = false
    v88 := GenS1{F0: 32767}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v86:", v86)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v73
    _ = v86
    _ = v87
    _ = v88
}
