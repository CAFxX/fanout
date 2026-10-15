package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 int32 }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint64; F1 uint16; F2 uint64 }
type GenS5 struct { F0 uint32 }
func genFn0(p0 bool, p1 string, p2 int16) uint8 {
    p1 = p1
    for fi4 := 0; fi4 < 20; fi4++ {
        p0 = p0
        var v5 uint16 = 534
        _ = fi4
        _ = v5
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}
func genFn1(p0 bool, p1 uint64, p2 int64) int16 {
    var v6 uint = 0
    var v7 int16 = 723
    for fi8 := 0; fi8 < 7; fi8++ {
        if (-2.25 == 1.5) {
            p0 = (v7 == (v7 / (((v7 + -913)) | 1)))
        } else {
            var v9 uint64 = p1
            v6 = v6
            _ = v9
        }
        var v10 uint16 = 0
        for fi11 := 0; fi11 < 1; fi11++ {
            v12 := GenS1{F0: p1, F1: ((p1 & 287) >= 18446744073709551615), F2: -2147483648}
            switch (^(fi8)) {
            case -995:
                if (false && true) {
                    if (fi8 > (fi8 * fi8)) {
                        v13 := []int{}
                        var v14 uint = (^(v6))
                        _ = v13
                        _ = v14
                    }
                    v15 := []int{fi8, -1, (^(fi11))}
                    p0 = p0
                    p1 = p1
                    _ = v15
                } else {
                    v16 := GenS1{F0: p1, F1: (!true), F2: 42}
                    _ = v16
                }
            case 1:
                if p0 {
                    var v17 uint8 = 185
                    p2 = p2
                    switch 1 {
                    case 4294967295:
                        v7 = (v7 / (((v7 << ((0) & 63))) | 1))
                    case 32767:
                        p1 = (p1 & (^(p1)))
                    case -9223372036854775808:
                        v17 += v17
                    default:
                        v18 := []int{}
                        _ = v18
                    }
                    _ = v17
                } else {
                    var v19 bool = (p0 && (true || ((!(((((float32((0.0 - 0.1)) - 1e10) - 3.14159) * (1.5 * (1e10 + float32((float64(float32(1e10)) + (1.5 - 1e10)))))) - 0.1) <= ((-2.25 / 0.1) * 0.1))) && (!p0))))
                    _ = v19
                }
            }
            v10 = uint16(p1)
            _ = fi11
            _ = v12
        }
        _ = fi8
        _ = v10
    }
    v6 += (v6 ^ 110)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return (v7 * (v7 * (^(v7))))
}

func main() {
    var v1 uint = 475
    v2 := make(map[int]uint32)
    var v3 bool = (9223372036854775807 == 1)
    switch "foo bar" {
    case "hello":
        v20 := GenS2{F0: 18446744073709551615, F1: 0}
        _ = v20
    case "\n\t":
        v21 := GenS2{F0: 1, F1: 87}
        _ = v21
    case "foo bar":
        for fi22 := 0; fi22 < 12; fi22++ {
            switch fi22 {
            case -146:
                if ((!(v3 || v3)) && false) {
                    var v23 uint = v1
                    v23 -= (v23 - 65535)
                    _ = v23
                }
            case -365:
                for fi24 := 0; fi24 < 6; fi24++ {
                    fmt.Println(fi22)
                    if (!(!(!true))) {
                        fmt.Println(fi22, v1)
                    } else {
                        if false {
                            fmt.Println(fi22, v3, fi24)
                            _ = len(v2)
                        }
                        v25 := GenS2{F0: 0, F1: genFn0(false, (("foo bar" + "") + "a"), 0)}
                        v3 = true
                        _ = v25
                    }
                    var v26 uint64 = 42
                    var v27 int8 = 1
                    v2[fi24] = 333
                    v1 = (v1 / ((989) | 1))
                    var v28 bool = v3
                    switch int(v1) {
                    case -2147483648:
                        if ((v27 & (v27 % ((1) | 1))) >= (v27 << (((v1 + (v1 - 255))) & 63))) {
                            var v29 int = -1
                            fmt.Println(fi22, v28, v27)
                            switch ("\n\t" + ("" + ("hello" + (("\n\t" + "hello") + ("hello" + "\n\t"))))) {
                            case "":
                                v30 := []int{fi24, 127, 438, (fi22 / ((213) | 1))}
                                _ = v30
                            case "hello":
                                v31 := GenS3{F0: (v3 || v3)}
                                _ = v31
                            case "\n\t":
                                if ((1e10 * 1e10) != 0.0) {
                                    v28 = (0.1 != 0.1)
                                    switch "\n\t" {
                                    case "a":
                                        var v32 float64 = 0.1
                                        _ = v32
                                    case "":
                                        var v33 float64 = (1e10 * 0.1)
                                        _ = v33
                                    default:
                                        v34 := GenS1{F0: (v26 | (v26 << ((v1) & 63))), F1: v28, F2: 42}
                                        _ = v34
                                    }
                                } else {
                                    var v35 float32 = 0.0
                                    v1 = (v1 | (v1 * 4294967295))
                                    _ = v35
                                }
                            default:
                                fmt.Println(v29, fi24)
                            }
                            _ = v29
                        }
                    }
                    if ((v26 / ((2) | 1)) > uint64(v1)) {
                        v27 = v27
                    }
                    _ = fi24
                    _ = v26
                    _ = v27
                    _ = v28
                }
            default:
                v3 = (145 > 127)
            }
            v1 -= (v1 / ((v1) | 1))
            _ = fi22
        }
    default:
        v1 += v1
    }
    v1 -= v1
    var v36 uint8 = 90
    if (false && v3) {
        switch "" {
        case "\n\t":
            switch ("a" + ("a" + "hello")) {
            case "":
                var v37 int64 = 42
                _ = v37
            case "a":
                var v38 int16 = -32768
                _ = v38
            case "\n\t":
                _ = len(v2)
            default:
                v36 = v36
            }
        case "hello":
            for fi39 := 0; fi39 < 15; fi39++ {
                v36 -= 89
                v40 := GenS4{F0: 386, F1: 65535, F2: 350}
                _ = fi39
                _ = v40
            }
        case "":
            v41 := GenS1{F0: 1, F1: (0 != (v36 ^ 127)), F2: 1}
            _ = v41
        default:
            var v42 float32 = 1.5
            _ = v42
        }
        var v43 int64 = 835
        v44 := GenS4{F0: 403, F1: uint16(v1), F2: 4294967295}
        _ = v43
        _ = v44
    }
    if v3 {
        v45 := GenS2{F0: 1, F1: (v36 / (((v36 + (v36 % (((^(v36))) | 1)))) | 1))}
        _ = v45
    }
    fmt.Println(v36, v1, v1)
    var v46 int32 = -221
    for fi47 := 0; fi47 < 17; fi47++ {
        v1 -= 925
        _ = fi47
    }
    var v48 uint32 = 2147483648
    var v49 int = -9223372036854775808
    v50 := GenS2{F0: 42, F1: v36}
    switch (v49 * (v49 << (((v1 - v1)) & 63))) {
    case -326:
        v51 := GenS2{F0: 2, F1: v36}
        _ = v51
    case 209:
        var v52 string = "\n\t"
        _ = v52
    case 2147483647:
        var v53 int16 = 32767
        _ = v53
    }
    switch ("hello" + "hello") {
    case "":
        for fi54 := 0; fi54 < 8; fi54++ {
            v3 = (1 != 934)
            _ = fi54
        }
    case "\n\t":
        for fi55 := 0; fi55 < 14; fi55++ {
            for fi56 := 0; fi56 < 10; fi56++ {
                v3 = v3
                var v57 float32 = 0.1
                _ = fi56
                _ = v57
            }
            fmt.Println(v46)
            _ = fi55
        }
    case "a":
        switch (v49 >> ((235) & 63)) {
        case 880:
            v1 -= uint(v48)
        case 127:
            fmt.Println(v48, v36, v48)
        }
    }
    v49 = 127
    for fi58 := 0; fi58 < 2; fi58++ {
        v46 = v46
        switch (fi58 + (fi58 ^ 127)) {
        case 42:
            fmt.Println(fi58, v46)
        case -789:
            switch (v49 + (-(v49))) {
            case -281:
                if false {
                    switch (v49 + (-(v49))) {
                    case 8:
                        v59 := GenS5{F0: uint32(v1)}
                        _ = v59
                    case -216:
                        var v60 bool = true
                        _ = v60
                    case 127:
                        v61 := GenS1{F0: 1, F1: false, F2: (v46 / ((int32(v48)) | 1))}
                        _ = v61
                    }
                    switch ("foo bar" + "") {
                    case "\n\t":
                        for fi62 := 0; fi62 < 7; fi62++ {
                            var v63 uint8 = v36
                            v63 = (v36 << (((v1 >> ((v1) & 63))) & 63))
                            _ = fi62
                            _ = v63
                        }
                    case "hello":
                        for fi64 := 0; fi64 < 14; fi64++ {
                            if v3 {
                                if v3 {
                                    v1 -= v1
                                }
                                v65 := []int{(fi64 & (v49 + (fi64 | fi58))), 0}
                                _ = v65
                            }
                            _ = fi64
                        }
                    case "":
                        fmt.Println(v1, v36, v1)
                    default:
                        v46 = (v46 / ((-2147483648) | 1))
                    }
                    var v66 float32 = 1e10
                    _ = v66
                } else {
                    fmt.Println(fi58)
                    for fi67 := 0; fi67 < 18; fi67++ {
                        v68 := []int{(^(fi67)), -866, 128}
                        var v69 uint16 = 1
                        _ = fi67
                        _ = v68
                        _ = v69
                    }
                }
            default:
                var v70 float32 = 3.14159
                _ = v70
            }
        case -2147483648:
            for fi71 := 0; fi71 < 12; fi71++ {
                fmt.Println(v46, v48, fi71)
                _ = fi71
            }
        default:
            if v3 {
                switch -203 {
                case -496:
                    switch "foo bar" {
                    case "hello":
                        var v72 uint8 = (^(v36))
                        _ = v72
                    case "foo bar":
                        var v73 uint = v1
                        _ = v73
                    default:
                        var v74 int64 = 1
                        _ = v74
                    }
                case -400:
                    v36 = 76
                case 1:
                    _ = len(v2)
                }
            }
        }
        _ = fi58
    }
    v75 := GenS3{F0: v3}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v46:", v46)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v46
    _ = v48
    _ = v49
    _ = v50
    _ = v75
}
