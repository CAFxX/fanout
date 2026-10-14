package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint16; F1 uint32; F2 string }
type GenS5 struct { F0 uint16; F1 bool }
func genFn0() uint8 {
    var v6 int = 42
    v6 += 2147483647
    fmt.Println(v6)
    _ = v6
    return 42
}
func genFn1() uint8 {
    if (1 == 9223372036854775807) {
        var v7 uint64 = 410
        v7 -= (v7 ^ v7)
        var v8 uint8 = genFn0()
        _ = v7
        _ = v8
    } else {
        for fi9 := 0; fi9 < 20; fi9++ {
            var v10 uint32 = 0
            for fi11 := 0; fi11 < 10; fi11++ {
                v10 += v10
                v12 := GenS1{F0: 42, F1: 9223372036854775807}
                v10 -= 4294967295
                _ = fi11
                _ = v12
            }
            _ = fi9
            _ = v10
        }
        v13 := []int{32767}
        var v14 uint8 = 40
        _ = v13
        _ = v14
    }
    v15 := GenS2{F0: 42, F1: 2}
    var v16 uint = 438
    v16 = (^(v16))
    _ = v15
    _ = v16
    return 0
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]int)
    v3 := make(map[int]string)
    var v4 uint = 18446744073709551615
    var v5 int8 = -128
    v4 = 18446744073709551615
    fmt.Println(v4, v5)
    if true {
        if (-9223372036854775808 < int64(v5)) {
            switch "\n\t" {
            case "foo bar":
                for fi17 := 0; fi17 < 3; fi17++ {
                    for fi18 := 0; fi18 < 12; fi18++ {
                        v19 := GenS2{F0: 974, F1: 2}
                        fmt.Println(v5)
                        fmt.Println(fi18, v5, fi17)
                        _ = fi18
                        _ = v19
                    }
                    var v20 uint16 = 349
                    _ = fi17
                    _ = v20
                }
            case "a":
                v5 -= -118
            case "hello":
                if (!(true || true)) {
                    v5 -= (v5 >> ((v4) & 63))
                    v21 := []int{0, 2147483647, 997, 40}
                    _ = v21
                }
            default:
                for fi22 := 0; fi22 < 17; fi22++ {
                    for fi23 := 0; fi23 < 5; fi23++ {
                        fmt.Println(v5, fi23)
                        var v24 int32 = 42
                        fmt.Println(fi22, fi23, v5)
                        _ = fi23
                        _ = v24
                    }
                    _ = fi22
                }
            }
        }
        for fi25 := 0; fi25 < 15; fi25++ {
            if true {
                switch -425 {
                case 2147483647:
                    v4 -= 951
                case 1:
                    v26 := GenS1{F0: v5, F1: fi25}
                    _ = v26
                case -9223372036854775808:
                    for fi27 := 0; fi27 < 14; fi27++ {
                        if false {
                            if (906 == 0) {
                                v28 := []int{}
                                fmt.Println(fi25, fi25, fi25)
                                _ = v28
                            } else {
                                var v29 int8 = v5
                                if ((0.0 + 0.1) > (3.14159 - ((-2.25 + float64(1e10)) + (((0.0 * 1e10) / (1.5 / (1.5 / 0.1))) / ((0.0 - ((3.14159 + 1e10) + (0.0 + -2.25))) + 0.0))))) {
                                    var v30 uint = (v4 | (v4 ^ 658))
                                    _ = v30
                                } else {
                                    v5 -= (v29 ^ 100)
                                    v5 = 127
                                    var v31 uint32 = 0
                                    switch fi25 {
                                    case -994:
                                        var v32 int32 = 220
                                        _ = v32
                                    case 0:
                                        fmt.Println(v4)
                                    default:
                                        fmt.Println(fi27, v5, v5)
                                    }
                                    var v33 int8 = 42
                                    _ = v31
                                    _ = v33
                                }
                                if (-1 >= -421) {
                                    v34 := GenS1{F0: v29, F1: fi25}
                                    if false {
                                        fmt.Println(v4)
                                    } else {
                                        fmt.Println(fi25, fi27)
                                        if false {
                                            fmt.Println(fi27, v5, fi25)
                                        }
                                        var v35 float32 = float32((((0.1 / 1e10) / (float64(1e10) - 0.0)) / 1.0))
                                        var v36 int8 = -128
                                        v4 += v4
                                        v37 := GenS3{F0: (!true)}
                                        v35 += 0.1
                                        _ = v35
                                        _ = v36
                                        _ = v37
                                    }
                                    _ = v34
                                } else {
                                    var v38 int32 = -1
                                    v29 = (^(v5))
                                    _ = v38
                                }
                                _ = v29
                            }
                            if (1e10 == -2.25) {
                                if (127 != 127) {
                                    var v39 uint = 740
                                    _ = v39
                                }
                                var v40 int8 = 100
                                switch "foo bar" {
                                case "foo bar":
                                    v5 = (v5 * (v5 >> ((2) & 63)))
                                default:
                                    switch (fi27 >> ((v4) & 63)) {
                                    case -9223372036854775808:
                                        fmt.Println(fi25)
                                    }
                                }
                                v2[1] = -51
                                fmt.Println(fi27, v5, fi25)
                                var v41 string = ""
                                v41 = v41
                                _ = v40
                                _ = v41
                            } else {
                                v42 := GenS2{F0: 42, F1: 543}
                                if (((float64(float32((1e10 - 1.5))) / (((0.0 * (1e10 + 0.1)) / 1.5) - -2.25)) / 1.0) == -2.25) {
                                    var v43 uint8 = genFn0()
                                    _ = len(v3)
                                    _ = v43
                                }
                                switch (-(fi27)) {
                                case 42:
                                    var v44 int16 = -1
                                    _ = v44
                                case 2147483647:
                                    var v45 float64 = 3.14159
                                    _ = v45
                                }
                                _ = v42
                            }
                            v4 = (v4 + (v4 - v4))
                        }
                        switch ("" + "a") {
                        case "a":
                            fmt.Println(fi25, fi27)
                        case "\n\t":
                            v4 += (v4 - (^(v4)))
                        case "foo bar":
                            v46 := []int{(fi25 >> (((v4 - (v4 << (((^(v4))) & 63)))) & 63)), fi25}
                            _ = v46
                        default:
                            fmt.Println(v4, fi27)
                        }
                        var v47 string = ("" + (("\n\t" + ("hello" + "hello")) + ("a" + "foo bar")))
                        v4 -= v4
                        v48 := []int{(fi27 ^ -694), fi25, fi25, 9223372036854775807}
                        _ = fi27
                        _ = v47
                        _ = v48
                    }
                }
                if false {
                    var v49 int32 = -1
                    if (893 > 1) {
                        var v50 int32 = (v49 & v49)
                        var v51 float64 = 1e10
                        _ = v50
                        _ = v51
                    }
                    fmt.Println(v4)
                    _ = v49
                }
            } else {
                v52 := GenS3{F0: (!true)}
                v53 := GenS4{F0: 1, F1: 168, F2: ("hello" + "foo bar")}
                var v54 bool = false
                _ = v52
                _ = v53
                _ = v54
            }
            _ = fi25
        }
        for fi55 := 0; fi55 < 14; fi55++ {
            for fi56 := 0; fi56 < 6; fi56++ {
                var v57 int32 = 0
                _ = fi56
                _ = v57
            }
            v2[42] = fi55
            _ = fi55
        }
    } else {
        v5 = v5
        for fi58 := 0; fi58 < 17; fi58++ {
            delete(v3, 127)
            v59 := v3[fi58]
            _ = fi58
            _ = v59
        }
        v4 -= v4
    }
    if false {
        v60 := []int{2147483647, 494}
        delete(v3, 474)
        _ = v60
    }
    fmt.Println(v5, v4)
    var v61 uint16 = 32767
    for fi62 := 0; fi62 < 4; fi62++ {
        var v63 bool = (!(203 >= 38))
        switch fi62 {
        case 829:
            fmt.Println(v63, v4)
        case -1:
            v64 := GenS1{F0: (v5 + (v5 - v5)), F1: 502}
            _ = v64
        default:
            v63 = (((0.1 + float32(0.0)) * float32(1.5)) != ((1.5 * 0.1) - -2.25))
        }
        if (((v5 % (((v5 * (v5 << (((v4 | (^(v4)))) & 63)))) | 1)) < (-(v5))) && (v4 >= (v4 >> ((676) & 63)))) {
            v65 := GenS3{F0: v63}
            _ = v65
        }
        _ = fi62
        _ = v63
    }
    v61 -= v61
    var v66 uint64 = 1
    var v67 int8 = (v5 | -1)
    var v68 uint8 = 111
    v1[-2147483648] = -262
    for fi69 := 0; fi69 < 20; fi69++ {
        v4 += (v4 % (((^(v4))) | 1))
        _ = fi69
    }
    for fi70 := 0; fi70 < 16; fi70++ {
        if true {
            v68 += 167
        }
        if (false || true) {
            v71 := GenS1{F0: 0, F1: (-(fi70))}
            _ = v71
        }
        v72 := []int{fi70, -2147483648, (fi70 % (((fi70 + 0)) | 1))}
        _ = fi70
        _ = v72
    }
    v66 = uint64(v68)
    if true {
        if (1000 >= int16(v66)) {
            v73 := GenS5{F0: (v61 / (((v61 % (((v61 << ((v4) & 63))) | 1))) | 1)), F1: false}
            v66 = (v66 & (v66 * v66))
            _ = v73
        }
        v74 := []int{int(v66), 4294967295, 9223372036854775807}
        var v75 float64 = 3.14159
        _ = v74
        _ = v75
    } else {
        v61 -= uint16(v66)
    }
    switch 9223372036854775807 {
    case 362:
        v61 += v61
    case -821:
        v61 = uint16(v4)
    case -366:
        if (1 > -1) {
            for fi76 := 0; fi76 < 19; fi76++ {
                v61 = uint16(v66)
                v1[fi76] = -1
                if ("hello" != "\n\t") {
                    v66 = (^(v66))
                    var v77 int = fi76
                    v4 -= v4
                    _ = v77
                } else {
                    v4 = (v4 << (((v4 >> ((uint(v61)) & 63))) & 63))
                    fmt.Println(v5, v66)
                    var v78 uint32 = 964
                    _ = v78
                }
                _ = fi76
            }
            var v79 int32 = -2147483648
            _ = v79
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v61:", v61)
    fmt.Println("v66:", v66)
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v61
    _ = v66
    _ = v67
    _ = v68
}
