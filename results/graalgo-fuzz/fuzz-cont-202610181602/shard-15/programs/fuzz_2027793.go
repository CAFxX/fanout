package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool; F1 uint32; F2 uint8 }
type GenS4 struct { F0 uint32; F1 bool }
type GenS5 struct { F0 int16; F1 int32 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint32 }

func main() {
    v1 := make(map[uint32]string)
    v2 := []string{("a" + "foo bar")}
    var v3 uint32 = 42
    var v4 uint16 = 793
    var v5 uint = 90
    v4 -= (v4 << ((v5) & 63))
    fmt.Println(v3, v4, v3)
    v6 := GenS1{F0: 358, F1: 2, F2: false}
    v3 = (v3 / ((v3) | 1))
    if true {
        switch 42 {
        case 0:
            var v7 uint16 = (v4 % ((v4) | 1))
            _ = v7
        default:
            var v8 uint64 = 18446744073709551615
            _ = v8
        }
    } else {
        var v9 int16 = 1
        v3 += (v3 - v3)
        if false {
            v1[(v3 >> ((v5) & 63))] = "a"
            v4 = v4
        }
        _ = v9
    }
    var v10 float32 = (3.14159 / 1e10)
    v11 := GenS1{F0: 0, F1: v4, F2: (!false)}
    switch 32767 {
    case -9223372036854775808:
        var v12 uint32 = 41
        _ = v12
    case 127:
        v13 := GenS1{F0: 42, F1: (v4 % ((523) | 1)), F2: (0 != 128)}
        _ = v13
    case 864:
        v14 := GenS2{F0: 151}
        _ = v14
    default:
        for fi15 := 0; fi15 < 20; fi15++ {
            for fi16 := 0; fi16 < 11; fi16++ {
                if (1 > 883) {
                    fmt.Println(v3, fi15, fi15)
                    switch "hello" {
                    case "":
                        if true {
                            var v17 uint16 = (v4 ^ v4)
                            fmt.Println(v4, fi15, fi16)
                            v18 := GenS2{F0: 4294967295}
                            _ = v17
                            _ = v18
                        } else {
                            if (!false) {
                                if (0.0 <= ((3.14159 + 0.1) + (float64((float32(1.5) * float32(((((3.14159 + (1e10 / 1.5)) - float64(v10)) + float64(1.5)) * 1.5)))) + 0.1))) {
                                    var v19 bool = true
                                    var v20 uint16 = 2
                                    v4 = (v20 | uint16(v5))
                                    _ = v19
                                    _ = v20
                                } else {
                                    var v21 uint64 = 42
                                    var v22 int32 = 127
                                    _ = v21
                                    _ = v22
                                }
                                switch "" {
                                case "foo bar":
                                    var v23 int32 = -1
                                    _ = v23
                                case "\n\t":
                                    var v24 float32 = v10
                                    _ = v24
                                default:
                                    v10 -= v10
                                }
                            }
                            if (!true) {
                                var v25 int8 = 127
                                if false {
                                    var v26 int16 = -272
                                    v27 := GenS2{F0: 18446744073709551615}
                                    var v28 uint8 = 253
                                    _ = v26
                                    _ = v27
                                    _ = v28
                                }
                                var v29 float32 = v10
                                _ = v25
                                _ = v29
                            }
                        }
                    case "foo bar":
                        var v30 int64 = 42
                        _ = v30
                    default:
                        fmt.Println(v4, v5)
                    }
                    var v31 uint = (v5 >> ((v5) & 63))
                    _ = v31
                } else {
                    fmt.Println(v3, v3, fi15)
                    fmt.Println(fi16)
                    if ((v10 >= ((((v10 - v10) / float32(3.14159)) / v10) / (3.14159 / v10))) && (true && true)) {
                        if true {
                            v32 := GenS3{F0: (true || (false || false)), F1: v3, F2: 128}
                            var v33 string = "\n\t"
                            _ = v32
                            _ = v33
                        }
                        switch (fi15 % ((-508) | 1)) {
                        case 0:
                            v34 := GenS3{F0: true, F1: v3, F2: 1}
                            _ = v34
                        case -760:
                            var v35 string = ("" + "\n\t")
                            _ = v35
                        case -9223372036854775808:
                            if true {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                v36 := GenS1{F0: 737, F1: (v4 - (v4 & 383)), F2: true}
                                _ = v36
                            }
                        default:
                            if (!false) {
                                var v37 int32 = 42
                                var v38 float32 = (((v10 * 1e10) + 0.1) * float32(-2.25))
                                _ = v37
                                _ = v38
                            } else {
                                v39 := GenS4{F0: uint32(v5), F1: (912 >= 357)}
                                _ = len(v1)
                                v40 := GenS1{F0: 32767, F1: v4, F2: (!((v5 - uint(v4)) != (v5 | v5)))}
                                _ = v39
                                _ = v40
                            }
                        }
                        var v41 int32 = 2147483647
                        v5 = (v5 << ((v5) & 63))
                        switch 0 {
                        case 32768:
                            v3 += (v3 >> (((v5 ^ 18446744073709551615)) & 63))
                        case -787:
                            _ = len(v2)
                        }
                        fmt.Println(v41, fi16, v10)
                        _ = v41
                    }
                    if true {
                        var v42 uint = v5
                        _ = v42
                    }
                    switch "foo bar" {
                    case "\n\t":
                        if true {
                            fmt.Println(v5, v10)
                            fmt.Println(v10)
                        } else {
                            delete(v1, uint32(v4))
                            delete(v1, v3)
                        }
                        fmt.Println(v10)
                    case "hello":
                        v4 += 42
                    default:
                        v5 = 65535
                    }
                    var v43 uint16 = v4
                    switch 128 {
                    case -2147483648:
                        if false {
                            var v44 uint8 = 42
                            _ = v44
                        }
                        var v45 float32 = v10
                        if true {
                            var v46 uint32 = (v3 / (((v3 >> ((v5) & 63))) | 1))
                            _ = v46
                        } else {
                            v45 = 3.14159
                        }
                        fmt.Println(fi15, v4)
                        if ((!false) && ((v3 & (v3 / (((^(v3))) | 1))) >= (^(v3)))) {
                            switch (^(fi16)) {
                            case 513:
                                v10 = float32(-2.25)
                            case 4294967295:
                                var v47 float64 = 1e10
                                _ = v47
                            default:
                                if true {
                                    v48 := GenS5{F0: 0, F1: -2147483648}
                                    v4 = (v4 >> ((v5) & 63))
                                    v45 = v45
                                    _ = v48
                                }
                            }
                            fmt.Println(fi15, fi16)
                        }
                        _ = v45
                    }
                    var v49 int16 = -417
                    _ = v43
                    _ = v49
                }
                if ((false || (true && false)) && (-1 <= -1)) {
                    switch fi16 {
                    case -1:
                        v50 := GenS6{F0: -112}
                        fmt.Println(v10)
                        _ = v50
                    case -671:
                        if true {
                            var v51 uint64 = 4294967295
                            fmt.Println(v5)
                            _ = v51
                        } else {
                            if (!(true || false)) {
                                v52 := GenS3{F0: ((50 > 65535) && true), F1: (^(v3)), F2: 1}
                                var v53 float64 = 3.14159
                                _ = v52
                                _ = v53
                            } else {
                                var v54 uint32 = 2147483647
                                v55 := GenS3{F0: false, F1: 4294967295, F2: 42}
                                if true {
                                    v10 += v10
                                    if false {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    } else {
                                        v56 := GenS2{F0: 2}
                                        v57 := GenS2{F0: 0}
                                        v58 := GenS4{F0: 0, F1: false}
                                        _ = v56
                                        _ = v57
                                        _ = v58
                                    }
                                    v59 := GenS5{F0: -32768, F1: int32(v4)}
                                    _ = v59
                                }
                                _ = v54
                                _ = v55
                            }
                            v4 -= v4
                            var v60 int16 = -698
                            _ = v60
                        }
                    default:
                        var v61 uint32 = (^(v3))
                        _ = v61
                    }
                    var v62 float64 = ((1.5 + 1e10) - 0.0)
                    _ = v62
                } else {
                    var v63 int64 = 9223372036854775807
                    if len(v2) > 0 {
                        v2[0] = (("" + ("" + ("" + "\n\t"))) + "\n\t")
                    }
                    _ = v63
                }
                _ = fi16
            }
            switch fi15 {
            case -284:
                switch fi15 {
                case 127:
                    var v64 int8 = -1
                    _ = v64
                case -801:
                    v10 += v10
                case 32767:
                    switch 128 {
                    case -762:
                        v65 := GenS4{F0: v3, F1: false}
                        _ = v65
                    case 42:
                        var v66 uint = 1
                        _ = v66
                    case -792:
                        fmt.Println(v5)
                    default:
                        var v67 int64 = 0
                        _ = v67
                    }
                }
            default:
                v68 := GenS1{F0: 32767, F1: (^(v4)), F2: ((1 < 837) && (!(!false)))}
                _ = v68
            }
            _ = fi15
        }
    }
    var v69 int32 = -492
    for fi70 := 0; fi70 < 10; fi70++ {
        if (((!((!false) || false)) || (false || ("hello" >= "foo bar"))) || false) {
            for fi71 := 0; fi71 < 9; fi71++ {
                delete(v1, 1)
                v72 := GenS5{F0: 1000, F1: (-(v69))}
                _ = fi71
                _ = v72
            }
            switch fi70 {
            case 0:
                var v73 string = "\n\t"
                _ = v73
            default:
                if true {
                    v69 -= int32(v5)
                } else {
                    fmt.Println(v3)
                    for fi74 := 0; fi74 < 4; fi74++ {
                        switch (fi74 ^ -1) {
                        case -9223372036854775808:
                            switch -385 {
                            case 42:
                                var v75 float32 = 3.14159
                                _ = v75
                            case -9223372036854775808:
                                switch "" {
                                case "":
                                    if true {
                                        v76 := GenS4{F0: 487, F1: true}
                                        var v77 int8 = 0
                                        v78 := GenS7{F0: 447}
                                        _ = v76
                                        _ = v77
                                        _ = v78
                                    }
                                }
                            case 9223372036854775807:
                                v5 = (^(v5))
                            }
                            var v79 uint8 = uint8(v4)
                            _ = len(v2)
                            _ = v79
                        }
                        _ = fi74
                    }
                    v5 = v5
                }
            }
            if true {
                var v80 bool = (((1.5 * -2.25) * (3.14159 * (((0.1 / 1e10) * (3.14159 - 0.1)) / 0.1))) < 0.1)
                _ = v80
            }
        }
        v4 -= v4
        _ = fi70
    }
    fmt.Println(v5, v69)
    if (2 < 116) {
        v81 := v1[(^(v3))]
        v81 = (v81 + v81)
        var v82 bool = false
        _ = v81
        _ = v82
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v69
}
