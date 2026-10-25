package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 uint8; F1 int16; F2 uint8 }
type GenS4 struct { F0 uint; F1 uint16; F2 int32 }
type GenS5 struct { F0 int; F1 int; F2 int8 }
type GenS6 struct { F0 int; F1 int32; F2 uint32 }
type GenS7 struct { F0 int32; F1 uint32; F2 int32 }
func genFn0(p0 int, p1 uint16) int64 {
    v7 := GenS1{F0: 42}
    v8 := []int{p0, 42}
    var v9 bool = (-292 <= 42)
    v10 := GenS2{F0: 127, F1: "hello"}
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return -124
}

func main() {
    var v1 int32 = 2147483647
    v2 := make(map[int]uint32)
    var v3 uint = 728
    v4 := make(map[int]string)
    v5 := make(map[int]string)
    v6 := make(map[int]string)
    var v11 int64 = 476
    var v12 uint8 = 0
    if (1.5 != -2.25) {
        v13 := []int{32768, 4294967295, -669}
        v12 = (^(v12))
        _ = v13
    } else {
        v11 -= genFn0(-400, 65535)
        v1 -= (-(v1))
    }
    if false {
        for fi14 := 0; fi14 < 16; fi14++ {
            if false {
                var v15 string = ("hello" + ("a" + ((("\n\t" + "hello") + "hello") + ("hello" + "a"))))
                fmt.Println(v11, v11)
                _ = v15
            }
            if (0 != 0) {
                for fi16 := 0; fi16 < 11; fi16++ {
                    v17 := GenS2{F0: 2, F1: "foo bar"}
                    var v18 uint8 = (^(v12))
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                if true {
                    for fi19 := 0; fi19 < 20; fi19++ {
                        switch "" {
                        case "hello":
                            fmt.Println(fi14, v1, fi19)
                        case "":
                            v12 = (v12 * (^(v12)))
                        case "a":
                            var v20 float32 = float32((3.14159 - (((3.14159 * (1.5 + 1e10)) + ((float64(float32((1e10 + (3.14159 - (-2.25 * 0.0))))) + (1.5 * 0.0)) / -2.25)) * (1.5 - 1.5))))
                            _ = v20
                        default:
                            switch (fi14 * (fi19 << ((v3) & 63))) {
                            case 1:
                                v1 = (v1 + v1)
                            }
                        }
                        _ = len(v4)
                        _ = fi19
                    }
                    switch 4294967295 {
                    case 407:
                        var v21 int = (fi14 / (((fi14 + (fi14 * int(v1)))) | 1))
                        _ = v21
                    case -9223372036854775808:
                        v12 -= 0
                    case 128:
                        var v22 uint64 = 0
                        _ = v22
                    }
                }
            } else {
                v23 := []int{0, (fi14 - -504), int(v11)}
                switch -1 {
                case 555:
                    if ((v1 << (((v3 ^ 255)) & 63)) > (^(v1))) {
                        v12 = v12
                        if false {
                            fmt.Println(v12)
                            var v24 int32 = (v1 >> ((v3) & 63))
                            v24 = 280
                            _ = v24
                        }
                    }
                case -14:
                    var v25 string = ("" + "\n\t")
                    _ = v25
                case 4294967295:
                    switch ("foo bar" + "foo bar") {
                    case "":
                        fmt.Println(fi14)
                    case "hello":
                        v23 = append(v23, (fi14 * fi14))
                    case "\n\t":
                        if true {
                            switch (fi14 >> ((18446744073709551615) & 63)) {
                            case 4294967295:
                                v3 = 255
                            default:
                                switch fi14 {
                                case 9223372036854775807:
                                    for fi26 := 0; fi26 < 9; fi26++ {
                                        switch "a" {
                                        case "\n\t":
                                            delete(v4, (fi26 / ((int(v11)) | 1)))
                                        case "hello":
                                            v27 := GenS2{F0: v12, F1: "foo bar"}
                                            _ = v27
                                        }
                                        v28 := v5[127]
                                        if true {
                                            switch fi14 {
                                            case 529:
                                                v11 = 42
                                                _ = len(v23)
                                                if (v1 >= 2147483647) {
                                                    var v29 uint8 = 255
                                                    _ = v29
                                                }
                                                v30 := GenS1{F0: 4294967295}
                                                var v31 int8 = 100
                                                _ = v30
                                                _ = v31
                                            }
                                            switch "" {
                                            case "a":
                                                v32 := GenS1{F0: uint32(v3)}
                                                _ = v32
                                            default:
                                                if ((32767 == 1) && (!((false || false) || ((v3 | 558) == 65535)))) {
                                                    v33 := GenS2{F0: 127, F1: ("hello" + v28)}
                                                    _ = v33
                                                }
                                            }
                                        } else {
                                            v34 := v4[(^(fi26))]
                                            var v35 int8 = 127
                                            _ = v34
                                            _ = v35
                                        }
                                        var v36 uint16 = uint16(v12)
                                        if len(v23) > 0 {
                                            v23[0] = (fi14 % ((int(v12)) | 1))
                                        }
                                        v37 := GenS2{F0: (v12 + v12), F1: v28}
                                        var v38 int32 = 1
                                        var v39 int8 = 0
                                        _ = fi26
                                        _ = v28
                                        _ = v36
                                        _ = v37
                                        _ = v38
                                        _ = v39
                                    }
                                case 330:
                                    var v40 int16 = 32767
                                    _ = v40
                                default:
                                    v2[fi14] = 0
                                }
                            }
                            v41 := GenS1{F0: 2}
                            _ = v41
                        }
                    default:
                        v5[(fi14 << ((42) & 63))] = "a"
                    }
                default:
                    v11 += (v11 - (v11 ^ v11))
                }
                v1 = (v1 & v1)
                _ = v23
            }
            fmt.Println(v1, v11)
            _ = fi14
        }
        fmt.Println(v12, v1)
        if true {
            if false {
                var v42 int32 = (v1 * v1)
                if (!(100 == -34)) {
                    for fi43 := 0; fi43 < 17; fi43++ {
                        for fi44 := 0; fi44 < 4; fi44++ {
                            v45 := GenS3{F0: (v12 * (v12 >> (((v3 - (v3 << ((v3) & 63)))) & 63))), F1: 1000, F2: (v12 % ((1) | 1))}
                            v12 = (^(v12))
                            if false {
                                v1 += -2147483648
                                v46 := GenS4{F0: 18446744073709551615, F1: 1, F2: v1}
                                _ = v46
                            } else {
                                fmt.Println(v3, v12, v42)
                                switch (-(fi44)) {
                                case 1:
                                    v4[fi44] = "foo bar"
                                default:
                                    v6[fi43] = ""
                                }
                                var v47 int8 = 85
                                _ = v47
                            }
                            _ = fi44
                            _ = v45
                        }
                        if true {
                            var v48 string = "\n\t"
                            _ = v48
                        }
                        _ = fi43
                    }
                    var v49 uint8 = (v12 | v12)
                    _ = v49
                } else {
                    fmt.Println(v3)
                    if false {
                        var v50 float64 = 3.14159
                        v42 = v42
                        v51 := GenS5{F0: 78, F1: -21, F2: -86}
                        _ = v50
                        _ = v51
                    }
                }
                _ = v42
            }
            if false {
                _ = len(v6)
                var v52 uint64 = 42
                _ = v52
            }
        } else {
            switch 2147483647 {
            case 1:
                fmt.Println(v3, v11)
            default:
                switch -209 {
                case 127:
                    var v53 float64 = 0.0
                    _ = v53
                case 42:
                    if (false && ((v12 ^ (^(v12))) > (v12 >> (((v3 + (^(v3)))) & 63)))) {
                        v3 += v3
                        v54 := v2[-756]
                        var v55 int16 = -322
                        _ = v54
                        _ = v55
                    }
                case -85:
                    v56 := v2[42]
                    _ = v56
                }
            }
            if (true || (!(true && (1000 >= 0)))) {
                v57 := []int{4294967295, 32767, -2147483648}
                _ = v57
            }
        }
    } else {
        for fi58 := 0; fi58 < 16; fi58++ {
            var v59 int8 = 42
            for fi60 := 0; fi60 < 5; fi60++ {
                switch (^(fi58)) {
                case 128:
                    var v61 uint64 = 2
                    _ = v61
                case -1:
                    v12 = (v12 * 113)
                }
                _ = len(v5)
                _ = fi60
            }
            _ = fi58
            _ = v59
        }
        fmt.Println(v1, v12)
    }
    for fi62 := 0; fi62 < 16; fi62++ {
        v3 += v3
        v11 += genFn0((fi62 >> ((v3) & 63)), 2)
        _ = fi62
    }
    fmt.Println(v12, v3)
    v1 = (-(v1))
    if false {
        for fi63 := 0; fi63 < 17; fi63++ {
            var v64 uint32 = 4294967295
            v12 = v12
            if (!true) {
                v65 := GenS2{F0: v12, F1: "foo bar"}
                switch "foo bar" {
                case "\n\t":
                    if (false || (!("hello" != ("a" + "hello")))) {
                        if true {
                            switch (fi63 >> (((^(v3))) & 63)) {
                            case -2147483648:
                                if true {
                                    switch -978 {
                                    case -9223372036854775808:
                                        for fi66 := 0; fi66 < 19; fi66++ {
                                            if true {
                                                v11 += v11
                                                var v67 uint16 = 420
                                                v68 := []int{int(v67), int(v11)}
                                                v11 += v11
                                                _ = v67
                                                _ = v68
                                            }
                                            _ = fi66
                                        }
                                    default:
                                        switch (-(fi63)) {
                                        case -9223372036854775808:
                                            var v69 int64 = (v11 - v11)
                                            _ = v69
                                        case 4294967295:
                                            v3 = (v3 * 2)
                                        case 2147483647:
                                            v1 = 42
                                        default:
                                            fmt.Println(v64, v64, v64)
                                        }
                                    }
                                    for fi70 := 0; fi70 < 4; fi70++ {
                                        v2[fi70] = (^(v64))
                                        v64 += v64
                                        _ = fi70
                                    }
                                    switch "\n\t" {
                                    case "hello":
                                        if true {
                                            v71 := GenS4{F0: (v3 / ((v3) | 1)), F1: 65535, F2: (v1 % ((v1) | 1))}
                                            v3 -= v3
                                            _ = v71
                                        } else {
                                            var v72 float32 = (float32(3.14159) + (float32(((0.0 / 1.0) + 3.14159)) - (0.0 - 1e10)))
                                            _ = v72
                                        }
                                    case "a":
                                        v73 := GenS5{F0: -9223372036854775808, F1: 128, F2: -128}
                                        _ = v73
                                    }
                                }
                            case 977:
                                v74 := GenS6{F0: fi63, F1: v1, F2: v64}
                                _ = v74
                            }
                            var v75 float64 = (((0.0 / (((0.0 / 1.5) - 0.1) / 1.5)) / (float64((1.5 - float32(0.1))) * (float64(((1.5 - 3.14159) / ((float32(1e10) / 1.0) * (((0.0 * (((-2.25 - 0.1) / (1.5 / (1.5 / 1.0))) - float32(1.5))) - 0.0) + 1e10)))) - -2.25))) - -2.25)
                            v76 := GenS7{F0: -2147483648, F1: 1, F2: (v1 + 1)}
                            _ = v75
                            _ = v76
                        }
                        if false {
                            var v77 uint16 = 42
                            var v78 uint = (^(v3))
                            _ = v77
                            _ = v78
                        }
                    }
                default:
                    var v79 uint16 = 32767
                    _ = v79
                }
                _ = v65
            } else {
                var v80 uint = v3
                _ = v80
            }
            _ = fi63
            _ = v64
        }
    } else {
        var v81 int64 = v11
        v82 := []int{-1, 105, 32768}
        var v83 float64 = 0.0
        _ = v81
        _ = v82
        _ = v83
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
