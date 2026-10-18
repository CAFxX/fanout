package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16; F1 int64; F2 uint }
type GenS4 struct { F0 uint32; F1 int16; F2 string }
type GenS5 struct { F0 uint32; F1 uint64; F2 bool }
type GenS6 struct { F0 int8; F1 int16; F2 int16 }
type GenS7 struct { F0 uint; F1 int }
type GenS8 struct { F0 uint64; F1 uint16; F2 uint16 }
type GenS9 struct { F0 int16; F1 string }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 uint64 }

func main() {
    var v1 uint8 = 2
    var v2 int32 = 1
    v3 := make(map[uint32]uint32)
    var v4 uint8 = 237
    v5 := GenS1{F0: -32768, F1: -468}
    v1 = (v4 & 0)
    if false {
        switch "" {
        case "\n\t":
            v2 = (v2 % (((v2 | -2147483648)) | 1))
        case "hello":
            var v6 uint64 = 1
            _ = v6
        }
        if (719 == 972) {
            var v7 int = 32768
            switch 920 {
            case 127:
                v8 := []int{}
                _ = v8
            }
            v1 -= (v1 % ((2) | 1))
            _ = v7
        } else {
            var v9 int8 = 0
            v1 += (v1 ^ v4)
            var v10 int64 = 693
            _ = v10
            _ = v9
        }
    }
    if (!(!false)) {
        v11 := []int{9223372036854775807, int(v2), -1}
        v1 = 89
        if (42 == 38) {
            v12 := GenS1{F0: int16(v1), F1: 32767}
            _ = v12
        } else {
            if (((471 <= 4294967295) || (!(false || (!true)))) || (!true)) {
                v13 := GenS1{F0: 32767, F1: -168}
                v2 -= (v2 / ((1) | 1))
                var v14 float64 = ((float64(0.1) + (3.14159 - ((-2.25 * 0.1) - (1.5 * 0.0)))) * 0.0)
                _ = v13
                _ = v14
            } else {
                v15 := GenS2{F0: 0}
                _ = v15
            }
            v16 := GenS1{F0: 1000, F1: 773}
            for fi17 := 0; fi17 < 7; fi17++ {
                v2 -= (v2 + v2)
                var v18 int16 = -32768
                _ = fi17
                _ = v18
            }
            _ = v16
        }
        _ = v11
    } else {
        for fi19 := 0; fi19 < 1; fi19++ {
            if false {
                if (985 > 0) {
                    v20 := []int{}
                    for fi21 := 0; fi21 < 6; fi21++ {
                        var v22 string = ""
                        v20 = append(v20, 128)
                        _ = fi21
                        _ = v22
                    }
                    _ = len(v3)
                    _ = v20
                } else {
                    var v23 uint32 = 1
                    v2 -= v2
                    v2 += 0
                    _ = v23
                }
                fmt.Println(v4)
                switch "hello" {
                case "a":
                    if ((((("hello" + ("" + "hello")) + "foo bar") + "") < "\n\t") || (!((-(v2)) >= 2147483647))) {
                        v1 -= (v4 & v4)
                        fmt.Println(fi19, v1, fi19)
                    }
                case "foo bar":
                    v1 -= 237
                }
            }
            v4 = v4
            _ = fi19
        }
    }
    v2 -= 42
    switch "" {
    case "":
        var v24 float32 = float32(0.0)
        _ = v24
    default:
        fmt.Println(v4)
    }
    if false {
        v2 = (v2 * v2)
        if false {
            v4 = (v4 >> ((255) & 63))
            v25 := GenS2{F0: 690}
            fmt.Println(v4, v2)
            _ = v25
        } else {
            var v26 int8 = 100
            _ = v26
        }
        for fi27 := 0; fi27 < 15; fi27++ {
            if false {
                v4 += v4
                v28 := GenS3{F0: 81, F1: 1, F2: 232}
                _ = v28
            } else {
                var v29 int64 = -9223372036854775808
                var v30 uint32 = 1
                _ = v29
                _ = v30
            }
            _ = fi27
        }
    } else {
        v31 := GenS4{F0: 684, F1: 1000, F2: ("hello" + "a")}
        _ = len(v3)
        _ = v31
    }
    var v32 uint32 = 42
    v33 := GenS3{F0: 978, F1: -1, F2: 42}
    v2 += -2147483648
    for fi34 := 0; fi34 < 8; fi34++ {
        for fi35 := 0; fi35 < 11; fi35++ {
            if ((float64(0.1) < 1e10) && (0.1 != 1.5)) {
                switch "\n\t" {
                case "\n\t":
                    v36 := []int{fi34, (fi35 * fi34), 216}
                    _ = v36
                default:
                    v37 := GenS5{F0: v32, F1: 1, F2: (0 > 4294967295)}
                    _ = v37
                }
            } else {
                if (false && (!true)) {
                    v38 := v3[v32]
                    _ = v38
                }
            }
            v2 += v2
            _ = fi35
        }
        if true {
            for fi39 := 0; fi39 < 19; fi39++ {
                v32 = (v32 ^ v32)
                var v40 uint64 = 18446744073709551615
                _ = fi39
                _ = v40
            }
            v41 := []int{}
            switch (^(fi34)) {
            case -9223372036854775808:
                fmt.Println(v4)
            }
            _ = v41
        }
        if (((((0.0 - (1e10 * ((-2.25 + 3.14159) - 3.14159))) - 1.5) <= ((float32(1.5) + 1.5) - ((3.14159 + (0.0 * (3.14159 + 3.14159))) * 1e10))) || false) && false) {
            if (1 > (^(v32))) {
                if false {
                    for fi42 := 0; fi42 < 7; fi42++ {
                        if true {
                            _ = len(v3)
                            if true {
                                v43 := GenS3{F0: 826, F1: -1, F2: 4294967295}
                                _ = v43
                            }
                            switch (("foo bar" + "foo bar") + "foo bar") {
                            case "hello":
                                switch ("hello" + (("" + "hello") + "\n\t")) {
                                case "a":
                                    if (true || false) {
                                        var v44 int8 = -1
                                        if false {
                                            var v45 int16 = 32767
                                            if false {
                                                v46 := GenS5{F0: v32, F1: 42, F2: (1e10 != 3.14159)}
                                                _ = v46
                                            } else {
                                                v47 := []int{}
                                                v4 += 1
                                                _ = v47
                                            }
                                            _ = v45
                                        } else {
                                            v48 := v3[1]
                                            delete(v3, 1)
                                            _ = v48
                                        }
                                        v32 += (^(v32))
                                        var v49 uint64 = 2
                                        var v50 uint = 4294967295
                                        var v51 int32 = v2
                                        if true {
                                            switch (^(fi34)) {
                                            case 4294967295:
                                                v52 := GenS6{F0: v44, F1: 0, F2: 1}
                                                var v53 int64 = -981
                                                _ = v52
                                                _ = v53
                                            case 128:
                                                v51 = int32(v49)
                                            case -684:
                                                var v54 int32 = (v51 << (((^(v50))) & 63))
                                                _ = v54
                                            }
                                        }
                                        fmt.Println(v1, fi34, fi42)
                                        _ = v44
                                        _ = v49
                                        _ = v50
                                        _ = v51
                                    } else {
                                        switch fi42 {
                                        case 42:
                                            var v55 int64 = 474
                                            _ = v55
                                        case -9223372036854775808:
                                            var v56 string = ""
                                            switch v56 {
                                            case "":
                                                v1 -= 2
                                            case "hello":
                                                v3[2147483648] = v32
                                            case "a":
                                                switch (fi34 + 701) {
                                                case 32768:
                                                    v57 := GenS7{F0: 1, F1: fi42}
                                                    v58 := GenS8{F0: 42, F1: 42, F2: 42}
                                                    v32 = v32
                                                    switch v56 {
                                                    case "foo bar":
                                                        if ((1000 > -179) && (1.5 >= 3.14159)) {
                                                            if true {
                                                                if false {
                                                                    v59 := GenS4{F0: (v32 >> ((4294967295) & 63)), F1: 660, F2: "foo bar"}
                                                                    _ = v59
                                                                }
                                                                v2 += v2
                                                                var v60 int8 = -46
                                                                _ = v60
                                                            } else {
                                                                v61 := GenS9{F0: 32767, F1: (v56 + v56)}
                                                                v62 := GenS4{F0: (v32 % (((v32 & 317)) | 1)), F1: 915, F2: v56}
                                                                v63 := GenS8{F0: 463, F1: 32767, F2: 65535}
                                                                _ = v61
                                                                _ = v62
                                                                _ = v63
                                                            }
                                                        }
                                                    case "hello":
                                                        v56 = v56
                                                    }
                                                    fmt.Println(v56, v56)
                                                    _ = v57
                                                    _ = v58
                                                }
                                            default:
                                                if (-1 != -647) {
                                                    v64 := []int{fi34, fi34}
                                                    _ = len(v3)
                                                    if true {
                                                        fmt.Println(v4, v1, v4)
                                                        v65 := GenS10{F0: v2}
                                                        _ = v65
                                                    } else {
                                                        v56 = (v56 + "\n\t")
                                                        if false {
                                                            var v66 int32 = v2
                                                            _ = v66
                                                        }
                                                        switch v56 {
                                                        case "hello":
                                                            delete(v3, (v32 | (^(v32))))
                                                        }
                                                        var v67 uint = 0
                                                        v67 = (v67 / (((v67 + v67)) | 1))
                                                        _ = v67
                                                    }
                                                    _ = v64
                                                }
                                            }
                                            v3[v32] = (^(v32))
                                            v68 := []int{fi34, (^(fi34)), (fi42 / ((fi42) | 1)), fi42}
                                            _ = v56
                                            _ = v68
                                        }
                                        v69 := GenS1{F0: -679, F1: 0}
                                        var v70 int32 = v2
                                        var v71 int64 = 972
                                        v72 := GenS2{F0: 1000}
                                        _ = v69
                                        _ = v70
                                        _ = v71
                                        _ = v72
                                    }
                                default:
                                    fmt.Println(fi34, fi42)
                                }
                            case "foo bar":
                                v1 = v4
                            case "\n\t":
                                v4 = (v4 & (v1 & v4))
                            }
                        } else {
                            delete(v3, v32)
                            v1 = (^(v4))
                            if (!((v2 != (v2 >> ((4294967295) & 63))) || false)) {
                                v4 += (v4 / ((50) | 1))
                                var v73 uint = 0
                                _ = v73
                            }
                            var v74 uint64 = 5
                            var v75 int32 = 2147483647
                            v76 := GenS11{F0: 4294967295}
                            _ = v74
                            _ = v75
                            _ = v76
                        }
                        if (false && (true || false)) {
                            switch 809 {
                            case 0:
                                fmt.Println(v2, fi42)
                            default:
                                var v77 int8 = 127
                                switch "\n\t" {
                                case "foo bar":
                                    v32 = 1
                                case "\n\t":
                                    fmt.Println(v77)
                                case "a":
                                    var v78 int32 = 0
                                    _ = v78
                                default:
                                    v79 := GenS4{F0: 510, F1: -1, F2: (("a" + "") + "")}
                                    _ = v79
                                }
                                v1 = v4
                                _ = v77
                            }
                            var v80 int = fi42
                            _ = v80
                        }
                        v1 = (v4 * 98)
                        var v81 bool = ((v1 % (((v4 - 2)) | 1)) != (^(v4)))
                        v1 -= 255
                        _ = fi42
                        _ = v81
                    }
                    v82 := GenS6{F0: -128, F1: 32767, F2: 1}
                    var v83 int16 = 0
                    _ = v82
                    _ = v83
                } else {
                    var v84 int16 = 32767
                    switch "" {
                    case "hello":
                        var v85 int16 = (v84 ^ v84)
                        _ = v85
                    default:
                        v2 -= v2
                    }
                    _ = v84
                }
                if ((("hello" + "a") <= "foo bar") || (0.1 == 3.14159)) {
                    v86 := []int{(fi34 - (fi34 >> ((3) & 63))), -325, (fi34 & 32768)}
                    for fi87 := 0; fi87 < 15; fi87++ {
                        fmt.Println(v2, v32, fi87)
                        if len(v86) > 0 {
                            v86[0] = (^(fi87))
                        }
                        _ = fi87
                    }
                    _ = v86
                }
                var v88 int = 9223372036854775807
                _ = v88
            }
            fmt.Println(v1)
        } else {
            v89 := []int{-2147483648, fi34, (fi34 / (((-(fi34))) | 1)), (fi34 << ((65535) & 63))}
            if len(v89) >= 2 {
                v89 = v89[1:len(v89)]
            }
            _ = v89
        }
        _ = fi34
    }
    v1 = v1
    if ((^(v32)) <= v32) {
        fmt.Println(v1, v1, v1)
        var v90 float64 = 0.1
        var v91 int8 = -128
        _ = v90
        _ = v91
    }
    var v92 float64 = 1e10
    var v93 int32 = (-(v2))
    v94 := []int{1}
    for fi95 := 0; fi95 < 2; fi95++ {
        v92 += (v92 - v92)
        _ = fi95
    }
    for fi96 := 0; fi96 < 20; fi96++ {
        v2 = -2147483648
        for fi97 := 0; fi97 < 1; fi97++ {
            var v98 int64 = 563
            _ = fi97
            _ = v98
        }
        if (!((-(fi96)) != 836)) {
            switch "" {
            case "":
                var v99 uint64 = 0
                _ = v99
            }
        }
        _ = fi96
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v92:", v92)
    fmt.Println("v93:", v93)
    fmt.Println("slicelen v94:", len(v94), v94)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v92
    _ = v93
    _ = v94
}
