package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 string }
type GenS2 struct { F0 uint16; F1 uint; F2 int8 }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 uint64; F1 string }
type GenS5 struct { F0 uint32; F1 string }
type GenS6 struct { F0 uint32; F1 int32; F2 uint8 }
func genFn0(p0 int) string {
    fmt.Println(p0)
    switch "a" {
    case "hello":
        var v7 string = "a"
        _ = v7
    default:
        var v8 uint32 = 65535
        _ = v8
    }
    var v9 uint8 = 0
    p0 = -1
    _ = p0
    _ = v9
    return ("\n\t" + "hello")
}
func genFn1() uint16 {
    v10 := GenS1{F0: 906, F1: 42, F2: (("" + "hello") + "foo bar")}
    var v11 uint = 18446744073709551615
    v11 -= 0
    v11 = 255
    if true {
        switch 128 {
        case -547:
            v12 := []int{32768, 4294967295, 326, 0}
            _ = v12
        }
        fmt.Println(v11)
        var v13 int64 = 0
        _ = v13
    } else {
        fmt.Println(v11)
        for fi14 := 0; fi14 < 12; fi14++ {
            var v15 uint64 = 18446744073709551615
            var v16 bool = false
            var v17 uint8 = 0
            _ = fi14
            _ = v15
            _ = v16
            _ = v17
        }
        for fi18 := 0; fi18 < 2; fi18++ {
            v11 = v11
            v11 -= (^(v11))
            var v19 string = ("" + "\n\t")
            _ = fi18
            _ = v19
        }
    }
    _ = v10
    _ = v11
    return 32767
}

func main() {
    var v1 int64 = 1
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]uint32)
    v4 := []int{42, 9223372036854775807, 127}
    v5 := make(map[uint32]uint32)
    var v6 uint16 = 1
    var v20 uint64 = 4294967295
    for fi21 := 0; fi21 < 9; fi21++ {
        if true {
            var v22 int8 = 41
            _ = v22
        } else {
            var v23 uint64 = v20
            var v24 string = ("" + "a")
            fmt.Println(v1, v20, v20)
            _ = v23
            _ = v24
        }
        fmt.Println(fi21)
        _ = fi21
    }
    v25 := v3[-732]
    v20 = 2
    if ((1.5 != 1e10) || (65535 >= 0)) {
        if (!(-1 <= 0)) {
            for fi26 := 0; fi26 < 10; fi26++ {
                for fi27 := 0; fi27 < 1; fi27++ {
                    var v28 string = ("\n\t" + "")
                    _ = fi27
                    _ = v28
                }
                if false {
                    for fi29 := 0; fi29 < 6; fi29++ {
                        delete(v3, 0)
                        v30 := GenS2{F0: genFn1(), F1: 576, F2: 40}
                        var v31 float64 = 1.5
                        _ = fi29
                        _ = v30
                        _ = v31
                    }
                    if ((v20 > (v20 ^ v20)) || true) {
                        for fi32 := 0; fi32 < 2; fi32++ {
                            var v33 int = (fi26 ^ (fi26 % ((int(v25)) | 1)))
                            if (false && ((404 < v6) && true)) {
                                v20 += (v20 ^ 42)
                                v20 = (v20 / (((v20 % ((1) | 1))) | 1))
                            } else {
                                v34 := GenS2{F0: (v6 * uint16(v25)), F1: 18446744073709551615, F2: -128}
                                _ = v34
                            }
                            _ = fi32
                            _ = v33
                        }
                        var v35 bool = true
                        if len(v4) > 0 {
                            v4[0] = 42
                        }
                        _ = v35
                    }
                }
                for fi36 := 0; fi36 < 7; fi36++ {
                    var v37 float32 = ((0.0 + (3.14159 + (3.14159 + 1.5))) - float32((0.1 + 1e10)))
                    var v38 int8 = 100
                    v25 += (v25 & v25)
                    _ = fi36
                    _ = v37
                    _ = v38
                }
                _ = fi26
            }
        } else {
            switch ("\n\t" + "hello") {
            case "":
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            case "foo bar":
                v6 += (v6 & uint16(v25))
            default:
                v39 := GenS2{F0: 42, F1: 1, F2: -128}
                _ = v39
            }
        }
        if false {
            fmt.Println(v1, v6)
        } else {
            v40 := GenS2{F0: (v6 & (^(v6))), F1: 255, F2: 127}
            _ = v40
        }
    }
    for fi41 := 0; fi41 < 4; fi41++ {
        for fi42 := 0; fi42 < 5; fi42++ {
            var v43 int8 = 100
            _ = fi42
            _ = v43
        }
        var v44 int16 = 42
        fmt.Println(fi41, fi41, v1)
        _ = fi41
        _ = v44
    }
    for fi45 := 0; fi45 < 11; fi45++ {
        var v46 uint32 = v25
        fmt.Println(fi45)
        if false {
            if (false || false) {
                if (true || true) {
                    for fi47 := 0; fi47 < 17; fi47++ {
                        v46 += 42
                        fmt.Println(fi47, fi47)
                        _ = fi47
                    }
                    if ((v6 << ((1) & 63)) == v6) {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        v1 -= v1
                        var v48 float32 = (float32(float64(float32(0.0))) - ((((((-2.25 - (1e10 / float32(-2.25))) + 0.1) * (-2.25 - 3.14159)) / (float32((3.14159 - (1.5 + (1e10 * 3.14159)))) + (1.5 * 3.14159))) * 3.14159) + (1e10 * float32(-2.25))))
                        _ = v48
                    } else {
                        if (!false) {
                            var v49 int = -1
                            _ = v49
                        }
                    }
                    var v50 int8 = -128
                    _ = v50
                } else {
                    for fi51 := 0; fi51 < 4; fi51++ {
                        v25 = (v46 - (v46 ^ (v25 % (((v25 + v25)) | 1))))
                        var v52 uint8 = 2
                        _ = fi51
                        _ = v52
                    }
                    var v53 float64 = 3.14159
                    _ = v53
                }
                v4 = append(v4, 32768)
            }
            fmt.Println(v20)
        }
        _ = fi45
        _ = v46
    }
    v4 = append(v4, 32768)
    var v54 float32 = 1.5
    switch "" {
    case "hello":
        v20 = (v20 & v20)
    default:
        fmt.Println(v25, v6)
    }
    for fi55 := 0; fi55 < 2; fi55++ {
        if (!(-371 == 1000)) {
            v54 = (((0.0 - v54) - v54) + float32(((-2.25 + 0.1) / (1e10 * 1e10))))
        }
        _ = fi55
    }
    for fi56 := 0; fi56 < 14; fi56++ {
        v4 = append(v4, fi56)
        var v57 int16 = -32768
        _ = fi56
        _ = v57
    }
    if true {
        v1 -= v1
        switch 841 {
        case 0:
            if false {
                var v58 int = 0
                v59 := GenS3{F0: 490, F1: 100}
                _ = v58
                _ = v59
            } else {
                var v60 uint64 = v20
                for fi61 := 0; fi61 < 17; fi61++ {
                    for fi62 := 0; fi62 < 16; fi62++ {
                        v20 = v20
                        v2[(v25 + (^(v25)))] = v25
                        _ = fi62
                    }
                    var v63 int32 = 1
                    v64 := GenS3{F0: -9223372036854775808, F1: -128}
                    _ = fi61
                    _ = v63
                    _ = v64
                }
                _ = v60
            }
        case -802:
            switch 0 {
            case 2147483647:
                var v65 uint = uint(v25)
                _ = v65
            case -78:
                if (288 <= 369) {
                    v66 := GenS2{F0: (v6 >> ((18446744073709551615) & 63)), F1: 1, F2: -98}
                    var v67 uint64 = 42
                    v68 := GenS2{F0: (v6 | genFn1()), F1: 0, F2: 75}
                    _ = v66
                    _ = v67
                    _ = v68
                }
            }
        case 465:
            for fi69 := 0; fi69 < 20; fi69++ {
                v20 -= v20
                _ = fi69
            }
        default:
            if (658 > 2) {
                switch ("foo bar" + "hello") {
                case "foo bar":
                    if (v54 >= v54) {
                        v70 := GenS2{F0: genFn1(), F1: 1, F2: -66}
                        v71 := GenS4{F0: 4294967295, F1: ((("foo bar" + "a") + ("hello" + ("" + (("hello" + ("hello" + "")) + "\n\t")))) + "\n\t")}
                        if true {
                            v72 := GenS5{F0: (^(v25)), F1: ("a" + (("hello" + "") + "foo bar"))}
                            _ = v72
                        } else {
                            v6 = v6
                            if (false && (true && ((v54 - ((0.1 * (float32(0.0) * v54)) * v54)) < (v54 - 1.5)))) {
                                fmt.Println(v25, v6, v25)
                            } else {
                                _ = len(v4)
                                fmt.Println(v1)
                            }
                        }
                        _ = v70
                        _ = v71
                    }
                case "hello":
                    for fi73 := 0; fi73 < 18; fi73++ {
                        v25 = v25
                        v74 := GenS3{F0: (fi73 & (fi73 ^ (fi73 << ((uint(v6)) & 63)))), F1: 105}
                        _ = fi73
                        _ = v74
                    }
                case "":
                    fmt.Println(v20, v20, v20)
                default:
                    v20 = v20
                }
                v1 += -9223372036854775808
                v54 -= (v54 * -2.25)
            } else {
                fmt.Println(v20)
                var v75 string = "\n\t"
                v76 := GenS6{F0: v25, F1: -1, F2: 255}
                _ = v75
                _ = v76
            }
        }
    }
    if (!false) {
        var v77 int8 = 0
        _ = v77
    }
    _ = len(v4)
    var v78 string = "\n\t"
    var v79 uint8 = 2
    if (3.14159 >= (0.1 + 1.5)) {
        _ = len(v4)
        if (!false) {
            fmt.Println(v78)
            v80 := v3[32767]
            if false {
                v81 := v2[(v80 << ((65535) & 63))]
                for fi82 := 0; fi82 < 11; fi82++ {
                    v5[(^(v81))] = v25
                    var v83 bool = true
                    _ = fi82
                    _ = v83
                }
                _ = v81
            }
            _ = v80
        }
    } else {
        var v84 uint32 = v25
        _ = v84
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v54:", v54)
    fmt.Println("v6:", v6)
    fmt.Println("v78:", v78)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v6
    _ = v78
    _ = v79
}
