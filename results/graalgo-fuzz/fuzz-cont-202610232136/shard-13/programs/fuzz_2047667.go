package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string; F1 uint; F2 int }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int; F1 bool }
type GenS5 struct { F0 int64; F1 int64; F2 uint16 }
type GenS6 struct { F0 int; F1 uint16; F2 int64 }
type GenS7 struct { F0 uint16; F1 int16 }
type GenS8 struct { F0 uint }
func genFn0() int {
    var v5 float32 = 0.1
    var v6 uint32 = 42
    if ((3.14159 - 3.14159) >= 0.1) {
        var v7 int8 = 93
        var v8 uint16 = 962
        if (4294967295 != 1) {
            switch ((((("a" + "foo bar") + "a") + "") + "\n\t") + (((("foo bar" + "a") + "a") + ("foo bar" + "a")) + "foo bar")) {
            case "\n\t":
                v9 := []int{int(v7), 2147483647, -9223372036854775808}
                _ = v9
            case "":
                if true {
                    fmt.Println(v6, v6, v7)
                } else {
                    v8 = 255
                    fmt.Println(v8, v8)
                }
            case "a":
                for fi10 := 0; fi10 < 8; fi10++ {
                    if false {
                        fmt.Println(v6, v7, v5)
                        var v11 float64 = 3.14159
                        _ = v11
                    } else {
                        v5 -= float32(0.1)
                    }
                    v8 = (v8 >> ((4294967295) & 63))
                    var v12 int16 = 865
                    _ = fi10
                    _ = v12
                }
            }
        } else {
            v7 += v7
            v13 := []int{}
            v14 := GenS1{F0: (v7 >> ((2) & 63))}
            _ = v13
            _ = v14
        }
        _ = v7
        _ = v8
    } else {
        for fi15 := 0; fi15 < 18; fi15++ {
            for fi16 := 0; fi16 < 7; fi16++ {
                var v17 string = "hello"
                _ = fi16
                _ = v17
            }
            _ = fi15
        }
        v18 := []int{2147483647, 32767}
        v19 := GenS2{F0: "foo bar", F1: 255, F2: 658}
        _ = v18
        _ = v19
    }
    switch 42 {
    case -947:
        if (false && false) {
            v20 := []int{}
            for fi21 := 0; fi21 < 12; fi21++ {
                if len(v20) >= 2 {
                    v20 = v20[1:len(v20)]
                }
                for fi22 := 0; fi22 < 11; fi22++ {
                    switch "" {
                    case "":
                        var v23 float32 = (1e10 + v5)
                        fmt.Println(fi22, fi22, fi22)
                        fmt.Println(fi22, v5)
                        _ = v23
                    case "a":
                        if false {
                            switch ("" + ("" + "")) {
                            case "\n\t":
                                v6 -= 761
                            default:
                                v5 += v5
                            }
                        }
                    }
                    if (49 < -1) {
                        var v24 int8 = -98
                        v5 = (v5 * float32(((((1e10 - (-2.25 - 0.1)) - ((((1.5 * 1e10) * 0.0) - 1e10) / ((((0.0 + 0.1) * (float64((((v5 * v5) + v5) + v5)) + 1.5)) * -2.25) + 3.14159))) / -2.25) / (-2.25 * 3.14159))))
                        _ = v24
                    } else {
                        v5 = v5
                        if (-103 <= -1) {
                            var v25 int8 = -1
                            var v26 int = int(v6)
                            if (0 < 0) {
                                v25 = v25
                                v6 -= (v6 | 249)
                                switch fi22 {
                                case 4294967295:
                                    var v27 string = "a"
                                    _ = v27
                                default:
                                    v5 += (3.14159 - float32((0.1 + float64((((float32(-2.25) * v5) / v5) - v5)))))
                                }
                                var v28 int32 = -2147483648
                                _ = v28
                            } else {
                                var v29 uint = 746
                                v6 = (v6 | 0)
                                var v30 int32 = 2147483647
                                _ = len(v20)
                                _ = v29
                                _ = v30
                            }
                            var v31 int64 = 348
                            _ = v25
                            _ = v26
                            _ = v31
                        } else {
                            var v32 float32 = (0.0 / 1.0)
                            v20 = append(v20, 4294967295)
                            _ = v32
                        }
                    }
                    var v33 int32 = -554
                    _ = fi22
                    _ = v33
                }
                fmt.Println(fi21, v6, v5)
                _ = fi21
            }
            for fi34 := 0; fi34 < 13; fi34++ {
                var v35 uint32 = v6
                switch fi34 {
                case 328:
                    v5 = 1e10
                }
                _ = fi34
                _ = v35
            }
            _ = v20
        } else {
            v5 += v5
        }
    case -968:
        if (52 != 110) {
            v6 += v6
            for fi36 := 0; fi36 < 18; fi36++ {
                if true {
                    var v37 int32 = -1
                    _ = v37
                }
                switch 42 {
                case 9223372036854775807:
                    for fi38 := 0; fi38 < 3; fi38++ {
                        var v39 int16 = 0
                        var v40 uint16 = 845
                        _ = fi38
                        _ = v39
                        _ = v40
                    }
                case 32768:
                    v41 := GenS1{F0: 117}
                    _ = v41
                case -499:
                    fmt.Println(v5)
                default:
                    var v42 uint = 255
                    _ = v42
                }
                _ = fi36
            }
        } else {
            v6 += (^(v6))
        }
    }
    v6 = (v6 | v6)
    _ = v5
    _ = v6
    return 42
}

func main() {
    var v1 uint64 = 486
    var v2 uint8 = 246
    v3 := make(map[uint32]bool)
    v4 := make(map[int]string)
    for fi43 := 0; fi43 < 20; fi43++ {
        fmt.Println(v2, v1, v2)
        v44 := GenS1{F0: 88}
        var v45 string = "\n\t"
        _ = fi43
        _ = v44
        _ = v45
    }
    fmt.Println(v1, v1)
    for fi46 := 0; fi46 < 4; fi46++ {
        var v47 int32 = -2147483648
        _ = fi46
        _ = v47
    }
    v1 = v1
    v48 := []int{593, 2147483647}
    for fi49 := 0; fi49 < 7; fi49++ {
        if (v2 >= 2) {
            if (-1 < -57) {
                var v50 uint64 = 42
                _ = v50
            } else {
                fmt.Println(fi49, fi49)
                for fi51 := 0; fi51 < 17; fi51++ {
                    fmt.Println(fi51, fi49, v1)
                    _ = fi51
                }
            }
        }
        v52 := GenS3{F0: 0}
        var v53 float64 = 1.5
        _ = fi49
        _ = v52
        _ = v53
    }
    v54 := GenS1{F0: 42}
    for fi55 := 0; fi55 < 6; fi55++ {
        _ = len(v3)
        if true {
            switch ("\n\t" + "foo bar") {
            case "\n\t":
                switch "" {
                case "\n\t":
                    if true {
                        v56 := GenS4{F0: -546, F1: true}
                        if true {
                            v2 += v2
                        } else {
                            v1 -= v1
                        }
                        _ = v56
                    } else {
                        fmt.Println(v2, fi55)
                        for fi57 := 0; fi57 < 12; fi57++ {
                            var v58 uint = 18446744073709551615
                            var v59 int16 = 213
                            fmt.Println(v58, v59)
                            _ = fi57
                            _ = v58
                            _ = v59
                        }
                    }
                }
            case "foo bar":
                for fi60 := 0; fi60 < 11; fi60++ {
                    v61 := GenS1{F0: 66}
                    v1 += v1
                    _ = fi60
                    _ = v61
                }
            case "":
                for fi62 := 0; fi62 < 16; fi62++ {
                    var v63 bool = (!(false || (412 < 42)))
                    _ = fi62
                    _ = v63
                }
            default:
                switch -665 {
                case -130:
                    for fi64 := 0; fi64 < 6; fi64++ {
                        switch 127 {
                        case 1:
                            v65 := v4[(-(fi64))]
                            if true {
                                var v66 int8 = -27
                                if ((v1 + (^(v1))) == (v1 - v1)) {
                                    switch (^(fi64)) {
                                    case 42:
                                        v67 := GenS5{F0: 78, F1: -703, F2: 255}
                                        _ = v67
                                    case -943:
                                        var v68 int16 = 1
                                        _ = v68
                                    }
                                    if true {
                                        v1 -= 2
                                    }
                                }
                                v1 = v1
                                _ = v66
                            } else {
                                v69 := GenS6{F0: genFn0(), F1: 255, F2: 9223372036854775807}
                                fmt.Println(fi64, fi64, fi55)
                                _ = v69
                            }
                            v1 += 4294967295
                            _ = v65
                        case -335:
                            if (true || (true || (1.5 < 0.1))) {
                                v1 = 0
                                if true {
                                    v70 := GenS3{F0: -200}
                                    fmt.Println(fi64, fi55)
                                    _ = v70
                                } else {
                                    v71 := GenS7{F0: 0, F1: 756}
                                    _ = v71
                                }
                            }
                        case 887:
                            var v72 float32 = 0.0
                            _ = v72
                        default:
                            if (true || true) {
                                _ = len(v3)
                                switch 32768 {
                                case 42:
                                    var v73 uint64 = uint64(v2)
                                    _ = v73
                                case 472:
                                    if (705 >= 65535) {
                                        if (false && (!true)) {
                                            v2 += (v2 >> ((293) & 63))
                                        } else {
                                            var v74 int64 = -608
                                            var v75 int64 = v74
                                            var v76 int = (fi55 << ((42) & 63))
                                            _ = v74
                                            _ = v75
                                            _ = v76
                                        }
                                        v1 = (v1 - (v1 / (((^(v1))) | 1)))
                                    }
                                case 9223372036854775807:
                                    switch (("hello" + ("\n\t" + "a")) + (("foo bar" + "\n\t") + "")) {
                                    case "hello":
                                        fmt.Println(v2, fi64)
                                    case "":
                                        v1 += (^(v1))
                                    case "\n\t":
                                        var v77 int8 = -62
                                        _ = v77
                                    }
                                default:
                                    v48 = append(v48, 2147483647)
                                }
                                v1 = 1
                            }
                        }
                        _ = fi64
                    }
                case 32767:
                    switch (-(fi55)) {
                    case -917:
                        if false {
                            var v78 int64 = 9223372036854775807
                            v79 := GenS8{F0: 4294967295}
                            _ = len(v48)
                            _ = v78
                            _ = v79
                        }
                    default:
                        var v80 int16 = 992
                        _ = v80
                    }
                }
            }
        }
        if (42 == 0) {
            if false {
                if len(v48) >= 2 {
                    v48 = v48[1:len(v48)]
                }
                switch int(v2) {
                case 0:
                    switch "a" {
                    case "":
                        var v81 uint64 = (^(v1))
                        _ = v81
                    case "a":
                        delete(v4, genFn0())
                    case "foo bar":
                        for fi82 := 0; fi82 < 17; fi82++ {
                            if ((^(v1)) < v1) {
                                var v83 int = (^(fi55))
                                fmt.Println(fi82, fi55, v2)
                                switch (fi82 / ((-2147483648) | 1)) {
                                case -9223372036854775808:
                                    v84 := v4[fi82]
                                    _ = v84
                                case 32767:
                                    if true {
                                        fmt.Println(fi82, fi55)
                                        switch "foo bar" {
                                        case "a":
                                            v2 -= (v2 + (v2 + v2))
                                        case "\n\t":
                                            var v85 string = ""
                                            _ = v85
                                        }
                                        v1 = v1
                                    }
                                case 0:
                                    v86 := v4[int(v1)]
                                    _ = v86
                                default:
                                    var v87 uint64 = v1
                                    _ = v87
                                }
                                var v88 int32 = -267
                                fmt.Println(v2)
                                _ = v83
                                _ = v88
                            }
                            var v89 int8 = -1
                            _ = fi82
                            _ = v89
                        }
                    }
                case -929:
                    fmt.Println(v1, fi55, fi55)
                case -9223372036854775808:
                    v90 := GenS7{F0: 2, F1: -599}
                    _ = v90
                default:
                    var v91 uint16 = 0
                    _ = v91
                }
            } else {
                delete(v4, 1)
                _ = len(v48)
                fmt.Println(v1, fi55)
            }
        }
        _ = fi55
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v48:", len(v48), v48)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v54
}
