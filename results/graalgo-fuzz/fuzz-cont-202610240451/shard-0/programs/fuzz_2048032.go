package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 uint32; F1 int8 }
type GenS4 struct { F0 uint; F1 uint16 }
type GenS5 struct { F0 uint8; F1 uint8 }
type GenS6 struct { F0 int32; F1 int8; F2 uint32 }
type GenS7 struct { F0 uint8; F1 uint16 }
type GenS8 struct { F0 uint32; F1 int8 }

func main() {
    v1 := make(map[string]int)
    var v2 uint8 = 122
    v3 := make(map[uint32]bool)
    var v4 int64 = 9223372036854775807
    switch "\n\t" {
    case "foo bar":
        for fi5 := 0; fi5 < 11; fi5++ {
            var v6 int32 = -1
            var v7 int = fi5
            if true {
                for fi8 := 0; fi8 < 1; fi8++ {
                    var v9 int16 = 42
                    _ = fi8
                    _ = v9
                }
                v2 -= v2
                v10 := GenS1{F0: 42, F1: v4}
                _ = v10
            } else {
                v11 := GenS1{F0: int16(v2), F1: int64(v6)}
                for fi12 := 0; fi12 < 1; fi12++ {
                    switch "a" {
                    case "a":
                        v13 := GenS2{F0: (-2.25 <= 1.5), F1: (!((^(v2)) > v2))}
                        _ = v13
                    case "hello":
                        var v14 float64 = 1e10
                        _ = v14
                    default:
                        var v15 uint8 = v2
                        _ = v15
                    }
                    _ = fi12
                }
                _ = v11
            }
            _ = fi5
            _ = v6
            _ = v7
        }
    default:
        if (-109 >= 127) {
            for fi16 := 0; fi16 < 9; fi16++ {
                switch fi16 {
                case 296:
                    v17 := GenS2{F0: (false && ((v2 & v2) >= (v2 | v2))), F1: true}
                    _ = v17
                default:
                    for fi18 := 0; fi18 < 2; fi18++ {
                        v2 = (v2 & v2)
                        var v19 uint8 = (v2 >> ((0) & 63))
                        _ = fi18
                        _ = v19
                    }
                }
                _ = fi16
            }
            fmt.Println(v4)
        }
    }
    switch 0 {
    case 127:
        for fi20 := 0; fi20 < 9; fi20++ {
            for fi21 := 0; fi21 < 12; fi21++ {
                v2 -= v2
                _ = fi21
            }
            var v22 int8 = 42
            switch fi20 {
            case -365:
                v4 = (v4 | (-(v4)))
            case -370:
                if false {
                    var v23 int64 = (v4 % ((v4) | 1))
                    v23 = v4
                    _ = v23
                } else {
                    var v24 float32 = -2.25
                    switch "a" {
                    case "foo bar":
                        var v25 uint64 = 42
                        _ = v25
                    case "\n\t":
                        for fi26 := 0; fi26 < 18; fi26++ {
                            var v27 uint = 42
                            var v28 float64 = 1e10
                            v22 = 127
                            if false {
                                v22 += v22
                                v24 = (v24 * v24)
                                v29 := GenS3{F0: 736, F1: 100}
                                var v30 int32 = 161
                                _ = v29
                                _ = v30
                            } else {
                                if true {
                                    v24 = v24
                                    if (v28 <= v28) {
                                        v4 += v4
                                    }
                                    switch "" {
                                    case "":
                                        v1[("foo bar" + ("" + ("foo bar" + (("hello" + ("a" + "\n\t")) + ("foo bar" + "hello")))))] = (fi26 / ((fi20) | 1))
                                    default:
                                        var v31 uint16 = 874
                                        _ = v31
                                    }
                                } else {
                                    v22 += -108
                                    var v32 string = "hello"
                                    v33 := []int{128, fi20, fi26}
                                    var v34 int32 = -2147483648
                                    _ = v32
                                    _ = v33
                                    _ = v34
                                }
                                v35 := GenS4{F0: (v27 * v27), F1: 255}
                                v2 -= (v2 & v2)
                                v2 = (v2 + (v2 * v2))
                                var v36 int8 = (v22 ^ 1)
                                v37 := []int{fi20}
                                v36 += (^(v36))
                                _ = v35
                                _ = v36
                                _ = v37
                            }
                            _ = fi26
                            _ = v27
                            _ = v28
                        }
                    }
                    _ = v24
                }
            default:
                for fi38 := 0; fi38 < 13; fi38++ {
                    v22 += int8(fi20)
                    v2 -= v2
                    var v39 int64 = (v4 << ((0) & 63))
                    var v40 bool = (369 <= 18446744073709551615)
                    if ((-2.25 - 0.1) == 3.14159) {
                        v41 := []int{fi38}
                        var v42 int16 = -1
                        v43 := GenS3{F0: 0, F1: v22}
                        _ = v41
                        _ = v42
                        _ = v43
                    }
                    _ = fi38
                    _ = v39
                    _ = v40
                }
            }
            _ = fi20
            _ = v22
        }
    case 4294967295:
        var v44 int64 = -9223372036854775808
        _ = v44
    case 0:
        if ((255 != 4294967295) && true) {
            v45 := GenS1{F0: -32768, F1: v4}
            var v46 float64 = ((((((1e10 + 1e10) * ((1.5 / 1.0) / ((0.1 - 0.0) * 1.5))) + float64((-2.25 - 1.5))) * 0.0) + 1e10) + 0.0)
            _ = v45
            _ = v46
        }
    default:
        v2 -= 2
    }
    if true {
        v3[260] = true
        switch "foo bar" {
        case "a":
            if true {
                for fi47 := 0; fi47 < 4; fi47++ {
                    if false {
                        if (-2147483648 <= -2147483648) {
                            for fi48 := 0; fi48 < 7; fi48++ {
                                fmt.Println(v2, v2)
                                v3[238] = (!(false || true))
                                fmt.Println(v2, fi48, fi48)
                                if true {
                                    switch ("foo bar" + "\n\t") {
                                    case "\n\t":
                                        fmt.Println(fi47)
                                    }
                                    v4 = v4
                                    var v49 bool = (42 <= -333)
                                    v49 = true
                                    _ = v49
                                }
                                _ = len(v3)
                                _ = fi48
                            }
                            for fi50 := 0; fi50 < 1; fi50++ {
                                var v51 bool = true
                                _ = fi50
                                _ = v51
                            }
                        } else {
                            v52 := GenS5{F0: 127, F1: v2}
                            _ = v52
                        }
                    } else {
                        v53 := GenS3{F0: 2147483647, F1: 89}
                        for fi54 := 0; fi54 < 2; fi54++ {
                            fmt.Println(fi54)
                            v2 += 128
                            _ = fi54
                        }
                        _ = v53
                    }
                    v55 := GenS1{F0: 964, F1: 0}
                    v56 := GenS6{F0: int32(v2), F1: -101, F2: 65535}
                    _ = fi47
                    _ = v55
                    _ = v56
                }
            }
        case "":
            v2 -= v2
        case "hello":
            v2 = 129
        default:
            for fi57 := 0; fi57 < 2; fi57++ {
                v2 -= 0
                fmt.Println(fi57)
                for fi58 := 0; fi58 < 13; fi58++ {
                    switch (fi57 + -2147483648) {
                    case 140:
                        var v59 int = (fi58 | (fi57 / (((fi58 ^ (fi58 + (fi58 + (fi58 & 4294967295))))) | 1)))
                        _ = v59
                    case 9223372036854775807:
                        v60 := GenS7{F0: v2, F1: 32767}
                        _ = v60
                    case 32768:
                        var v61 int = 42
                        v4 += (v4 & v4)
                        v61 -= v61
                        _ = v61
                    }
                    _ = fi58
                }
                _ = fi57
            }
        }
    }
    for fi62 := 0; fi62 < 18; fi62++ {
        v4 -= (^(v4))
        v2 += (v2 - v2)
        v63 := v1["foo bar"]
        _ = fi62
        _ = v63
    }
    v64 := []int{382, 32768, -856}
    var v65 bool = (((v4 <= v4) && true) && false)
    if (0 != 292) {
        v2 -= (v2 << ((447) & 63))
    } else {
        if (!(float64(float32((1.5 * 3.14159))) <= 1.5)) {
            for fi66 := 0; fi66 < 8; fi66++ {
                fmt.Println(v4, v65)
                for fi67 := 0; fi67 < 20; fi67++ {
                    var v68 int16 = -1
                    _ = fi67
                    _ = v68
                }
                _ = fi66
            }
        }
    }
    var v69 uint64 = 0
    v65 = true
    var v70 int = -883
    switch "foo bar" {
    case "":
        var v71 uint64 = uint64(v2)
        _ = v71
    case "a":
        delete(v3, 42)
    default:
        for fi72 := 0; fi72 < 14; fi72++ {
            var v73 int32 = 0
            _ = fi72
            _ = v73
        }
    }
    v74 := v3[4294967295]
    v75 := GenS3{F0: 43, F1: 1}
    switch (v70 | 0) {
    case 127:
        var v76 bool = true
        _ = v76
    case -9223372036854775808:
        fmt.Println(v69, v65, v70)
    case -1:
        v65 = v65
    default:
        v2 = v2
    }
    if ((^(v69)) >= 2) {
        switch (v70 + 32768) {
        case 9223372036854775807:
            v2 = v2
        default:
            v74 = (((v4 << ((255) & 63)) < v4) || (-128 >= -21))
        }
        var v77 string = ""
        v2 = 42
        _ = v77
    }
    switch 32768 {
    case 367:
        v74 = v74
    case 128:
        v78 := GenS1{F0: 0, F1: int64(v2)}
        _ = v78
    case 0:
        v79 := GenS5{F0: (^(v2)), F1: (v2 >> ((18446744073709551615) & 63))}
        _ = v79
    default:
        var v80 string = "foo bar"
        _ = v80
    }
    v65 = (707 <= 255)
    if (uint16(v2) == 86) {
        v81 := GenS8{F0: 615, F1: 1}
        v82 := GenS4{F0: 763, F1: 417}
        _ = v81
        _ = v82
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v64:", len(v64), v64)
    fmt.Println("v65:", v65)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v64
    _ = v65
    _ = v69
    _ = v70
    _ = v74
    _ = v75
}
