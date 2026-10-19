package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string }
type GenS2 struct { F0 uint8; F1 uint16; F2 uint8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int64; F1 int64 }
type GenS5 struct { F0 string; F1 string; F2 int16 }
type GenS6 struct { F0 bool; F1 uint16; F2 uint8 }
type GenS7 struct { F0 int16; F1 uint8; F2 uint32 }
type GenS8 struct { F0 int16; F1 uint64; F2 uint16 }
type GenS9 struct { F0 uint8; F1 int64; F2 uint32 }
type GenS10 struct { F0 int; F1 string; F2 uint64 }
type GenS11 struct { F0 uint32 }
type GenS12 struct { F0 int8 }
type GenS13 struct { F0 int16 }
func genFn0() uint64 {
    v5 := []int{-2147483648, 9223372036854775807, 32768, 679}
    if len(v5) > 0 {
        v5[0] = 32767
    }
    _ = v5
    return 1
}

func main() {
    v1 := make(map[int]string)
    var v2 int8 = -114
    v3 := make(map[uint32]bool)
    var v4 int16 = 464
    if (!true) {
        if ((true || (0 > 529)) && false) {
            v6 := []int{127, -2147483648, 4294967295}
            _ = v6
        } else {
            var v7 uint32 = 2147483648
            _ = v7
        }
        for fi8 := 0; fi8 < 19; fi8++ {
            if true {
                var v9 float32 = 0.1
                if true {
                    v10 := GenS1{F0: 0, F1: ""}
                    for fi11 := 0; fi11 < 6; fi11++ {
                        var v12 uint8 = 128
                        var v13 int64 = 0
                        fmt.Println(v13)
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                    fmt.Println(v4, fi8)
                    _ = v10
                }
                _ = v9
            } else {
                if true {
                    v14 := []int{(fi8 - fi8), (fi8 + fi8)}
                    var v15 uint8 = 1
                    switch "" {
                    case "foo bar":
                        if (42 <= 711) {
                            v4 = (v4 << ((4294967295) & 63))
                        }
                    default:
                        v16 := GenS2{F0: (v15 << ((255) & 63)), F1: 0, F2: 128}
                        _ = v16
                    }
                    _ = v14
                    _ = v15
                } else {
                    switch "hello" {
                    case "\n\t":
                        if true {
                            fmt.Println(fi8, v4)
                        } else {
                            if (2147483647 >= -1) {
                                var v17 int64 = 42
                                v4 -= (v4 - 171)
                                _ = v17
                            } else {
                                switch (fi8 * (fi8 << ((18446744073709551615) & 63))) {
                                case 32767:
                                    var v18 int16 = (^(v4))
                                    _ = v18
                                case 32768:
                                    switch -1 {
                                    case 42:
                                        var v19 uint64 = genFn0()
                                        _ = v19
                                    case 9223372036854775807:
                                        v20 := GenS2{F0: 186, F1: 0, F2: 1}
                                        _ = v20
                                    case 128:
                                        var v21 uint = 65535
                                        _ = v21
                                    }
                                case -390:
                                    v4 = (v4 | (v4 + v4))
                                }
                                v22 := GenS1{F0: 171, F1: ("" + "hello")}
                                _ = v22
                            }
                            v4 = v4
                        }
                    }
                    for fi23 := 0; fi23 < 18; fi23++ {
                        v24 := GenS2{F0: 118, F1: 610, F2: 1}
                        v2 = (v2 >> ((255) & 63))
                        v25 := v3[611]
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                }
                fmt.Println(fi8, fi8)
            }
            for fi26 := 0; fi26 < 8; fi26++ {
                if (v4 <= v4) {
                    fmt.Println(fi8, fi26)
                }
                fmt.Println(fi8, fi26, fi8)
                switch fi26 {
                case 289:
                    if true {
                        if false {
                            var v27 uint64 = 673
                            v28 := []int{(-(fi8)), (fi26 & 128), -2147483648}
                            v29 := GenS3{F0: 2147483647}
                            _ = v27
                            _ = v28
                            _ = v29
                        } else {
                            var v30 int32 = 0
                            var v31 float64 = 0.0
                            _ = v30
                            _ = v31
                        }
                    } else {
                        v32 := GenS2{F0: 1, F1: 404, F2: 185}
                        v33 := GenS4{F0: -1, F1: 718}
                        switch "hello" {
                        case "a":
                            if false {
                                v34 := []int{}
                                v2 = (v2 | (v2 << ((0) & 63)))
                                v35 := GenS3{F0: 0}
                                v4 -= v4
                                _ = v34
                                _ = v35
                            }
                        case "\n\t":
                            v4 = (v4 / ((v4) | 1))
                        default:
                            v36 := GenS3{F0: 0}
                            _ = v36
                        }
                        _ = v32
                        _ = v33
                    }
                case 127:
                    v37 := v1[508]
                    _ = v37
                case 128:
                    fmt.Println(v2, fi8, v2)
                default:
                    var v38 uint8 = 140
                    fmt.Println(fi8)
                    _ = v38
                }
                var v39 int16 = v4
                _ = len(v1)
                v40 := GenS5{F0: ("\n\t" + "hello"), F1: (("foo bar" + ("hello" + "hello")) + ""), F2: (v39 + (v4 ^ v39))}
                v41 := GenS6{F0: false, F1: 274, F2: 0}
                if true {
                    v39 = v39
                    v42 := v1[fi8]
                    _ = v42
                } else {
                    var v43 int32 = 0
                    fmt.Println(v2, v39)
                    _ = v43
                }
                _ = fi26
                _ = v39
                _ = v40
                _ = v41
            }
            _ = fi8
        }
    }
    var v44 int16 = v4
    v45 := []int{-76, -657, 0}
    for fi46 := 0; fi46 < 4; fi46++ {
        switch ("hello" + "a") {
        case "\n\t":
            for fi47 := 0; fi47 < 4; fi47++ {
                if (2147483647 <= 1) {
                    if true {
                        if (2 < 2) {
                            switch (fi47 & fi46) {
                            case 42:
                                fmt.Println(v4)
                            }
                        }
                        v48 := v1[(fi47 + (fi47 >> ((65535) & 63)))]
                        _ = v48
                    } else {
                        var v49 uint8 = 1
                        _ = v49
                    }
                } else {
                    v50 := GenS7{F0: 1000, F1: 127, F2: 0}
                    v3[65535] = true
                    var v51 string = "foo bar"
                    _ = v50
                    _ = v51
                }
                v52 := GenS8{F0: v4, F1: 42, F2: 131}
                _ = fi47
                _ = v52
            }
        case "a":
            if true {
                var v53 uint = 8
                _ = v53
            }
        }
        v44 = v4
        _ = fi46
    }
    if false {
        v2 += v2
        v4 -= v44
    } else {
        if ((0.0 != 1e10) && (false || (-968 != 128))) {
            var v54 uint16 = 555
            for fi55 := 0; fi55 < 16; fi55++ {
                fmt.Println(v54, v4, fi55)
                switch "" {
                case "hello":
                    v2 = v2
                case "":
                    var v56 int64 = 0
                    _ = v56
                default:
                    var v57 bool = (503 < (v54 % ((v54) | 1)))
                    _ = v57
                }
                _ = fi55
            }
            _ = v54
        } else {
            var v58 int64 = -453
            var v59 int64 = (v58 | v58)
            _ = v58
            _ = v59
        }
        for fi60 := 0; fi60 < 15; fi60++ {
            var v61 int = fi60
            if (563 == 2147483648) {
                var v62 uint = 1
                _ = v62
            } else {
                var v63 int32 = 383
                v44 = int16(v63)
                if ((100 > v2) || false) {
                    switch 466 {
                    case -2147483648:
                        v44 -= (v4 >> ((827) & 63))
                    }
                    v64 := v1[v61]
                    _ = v64
                }
                _ = v63
            }
            _ = fi60
            _ = v61
        }
        v1[42] = "hello"
    }
    if false {
        v65 := GenS6{F0: (!(!true)), F1: 2, F2: 127}
        _ = v65
    } else {
        if (false && true) {
            v44 += (v4 * -1)
            v3[0] = (1.5 == (0.0 / 1.0))
            switch ("a" + "hello") {
            case "\n\t":
                v66 := GenS9{F0: 194, F1: 1, F2: 958}
                _ = v66
            case "a":
                switch 75 {
                case 0:
                    fmt.Println(v44, v44, v44)
                case -751:
                    switch 128 {
                    case -2147483648:
                        v4 = -1
                    case 1:
                        _ = len(v45)
                    case 455:
                        fmt.Println(v44, v4)
                    default:
                        if true {
                            if (false || (false || false)) {
                                var v67 uint16 = 743
                                _ = v67
                            }
                        } else {
                            if false {
                                fmt.Println(v4, v2)
                                if false {
                                    if (!(42 < 32767)) {
                                        v2 += 0
                                        v2 += 100
                                    } else {
                                        fmt.Println(v44, v2)
                                        if (3.14159 < float64(((float32(1.5) + -2.25) / 1e10))) {
                                            if true {
                                                var v68 int64 = -913
                                                _ = v68
                                            }
                                        }
                                    }
                                } else {
                                    v69 := GenS6{F0: (float64(3.14159) != 0.0), F1: 312, F2: 0}
                                    v45 = append(v45, 128)
                                    v44 -= v44
                                    _ = v69
                                }
                            } else {
                                v4 = 1000
                                v44 -= (v4 | int16(v2))
                                v3[4294967295] = false
                            }
                            fmt.Println(v4)
                            v4 = (-(v44))
                        }
                    }
                }
            default:
                if false {
                    if ((1.5 * 1.5) >= (float64(0.1) - float64(1.5))) {
                        v2 = 1
                    } else {
                        fmt.Println(v4)
                    }
                    for fi70 := 0; fi70 < 10; fi70++ {
                        var v71 float64 = (-2.25 / 1.0)
                        _ = fi70
                        _ = v71
                    }
                }
            }
        } else {
            switch 0 {
            case 127:
                fmt.Println(v2, v44)
            case 2147483647:
                if true {
                    v72 := GenS10{F0: 1, F1: "foo bar", F2: genFn0()}
                    _ = v72
                } else {
                    if false {
                        switch 496 {
                        case -484:
                            v1[127] = ("" + ((("" + ("" + ("" + "a"))) + ("foo bar" + ("hello" + ("foo bar" + "foo bar")))) + "foo bar"))
                        case 127:
                            if true {
                                v44 = 1
                                if len(v45) >= 2 {
                                    v45 = v45[1:len(v45)]
                                }
                                var v73 uint16 = 512
                                _ = v73
                            } else {
                                v74 := GenS4{F0: 0, F1: 1}
                                if (141 < 42) {
                                    v2 -= v2
                                    _ = len(v45)
                                } else {
                                    v44 = (v44 ^ int16(v2))
                                    v75 := GenS11{F0: 1}
                                    _ = v75
                                }
                                _ = v74
                            }
                        default:
                            fmt.Println(v2)
                        }
                        fmt.Println(v4, v44, v4)
                        switch 957 {
                        case 128:
                            for fi76 := 0; fi76 < 19; fi76++ {
                                if (42 >= 127) {
                                    switch "hello" {
                                    case "foo bar":
                                        var v77 float32 = (1e10 + (0.1 / 1e10))
                                        _ = v77
                                    case "\n\t":
                                        v78 := GenS12{F0: -77}
                                        _ = v78
                                    }
                                    switch fi76 {
                                    case 128:
                                        for fi79 := 0; fi79 < 2; fi79++ {
                                            switch 128 {
                                            case 32767:
                                                var v80 string = "hello"
                                                _ = v80
                                            case 127:
                                                var v81 int8 = v2
                                                _ = v81
                                            case -9223372036854775808:
                                                fmt.Println(fi76)
                                            }
                                            _ = fi79
                                        }
                                    case -92:
                                        var v82 uint32 = 4294967295
                                        _ = v82
                                    }
                                    v3[2147483648] = false
                                } else {
                                    _ = len(v45)
                                }
                                v83 := GenS13{F0: 1000}
                                for fi84 := 0; fi84 < 1; fi84++ {
                                    v4 -= v44
                                    _ = fi84
                                }
                                _ = fi76
                                _ = v83
                            }
                        case 509:
                            var v85 uint64 = 474
                            _ = v85
                        case 1:
                            var v86 uint8 = 109
                            _ = v86
                        default:
                            switch -704 {
                            case 128:
                                v44 = (-(v44))
                            case 4294967295:
                                v3[213] = true
                            case -2147483648:
                                v44 += -830
                            }
                        }
                    } else {
                        var v87 uint16 = 1
                        delete(v3, 2147483648)
                        var v88 int8 = (v2 / ((-119) | 1))
                        _ = v87
                        _ = v88
                    }
                    switch 32768 {
                    case 127:
                        fmt.Println(v2, v4, v2)
                    case 803:
                        var v89 float64 = 3.14159
                        _ = v89
                    case -1:
                        delete(v3, 2)
                    }
                    fmt.Println(v2)
                }
            case 0:
                if (false || (v4 < v44)) {
                    var v90 uint = 124
                    if true {
                        v91 := v1[9223372036854775807]
                        var v92 int16 = v44
                        if len(v45) > 0 {
                            v45[0] = 9223372036854775807
                        }
                        _ = v91
                        _ = v92
                    }
                    _ = v90
                }
            }
            v4 = (v4 >> ((995) & 63))
            var v93 string = "\n\t"
            _ = v93
        }
        var v94 int16 = 1000
        _ = v94
    }
    fmt.Println(v2, v2, v44)
    for fi95 := 0; fi95 < 20; fi95++ {
        var v96 int = fi95
        if (42 > 0) {
            v97 := GenS4{F0: 0, F1: -818}
            _ = v97
        } else {
            var v98 uint32 = 945
            if ("" != ("" + "\n\t")) {
                if (false || (!true)) {
                    var v99 uint = 65535
                    fmt.Println(v99, v98, fi95)
                    _ = v99
                }
                var v100 float32 = float32(0.0)
                var v101 uint64 = 42
                _ = v100
                _ = v101
            } else {
                delete(v3, v98)
            }
            v44 -= v44
            _ = v98
        }
        _ = fi95
        _ = v96
    }
    switch -2147483648 {
    case 42:
        if len(v45) > 0 {
            v45[0] = -974
        }
    }
    v102 := GenS13{F0: (v4 | 0)}
    if true {
        for fi103 := 0; fi103 < 19; fi103++ {
            delete(v1, int(v2))
            v2 = v2
            v2 = (v2 % ((v2) | 1))
            _ = fi103
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v45:", len(v45), v45)
    _ = v1
    _ = v102
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
}
