package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16; F1 uint64 }
type GenS5 struct { F0 uint64; F1 bool; F2 int }
type GenS6 struct { F0 int16; F1 int8 }
type GenS7 struct { F0 int8; F1 uint32; F2 int16 }
type GenS8 struct { F0 bool; F1 bool }
func genFn0(p0 int64, p1 bool) int32 {
    v4 := []int{-107}
    var v5 uint = 2
    p1 = true
    v5 = v5
    p0 = p0
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return -576
}

func main() {
    var v1 string = "a"
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]uint32)
    for fi6 := 0; fi6 < 6; fi6++ {
        v1 = "foo bar"
        switch 704 {
        case 32767:
            var v7 uint16 = 32767
            _ = v7
        case 0:
            if (!(!(false && true))) {
                v1 = "hello"
                v8 := GenS1{F0: 113}
                _ = v8
            } else {
                for fi9 := 0; fi9 < 1; fi9++ {
                    if false {
                        if ((42 <= 172) || true) {
                            v10 := GenS2{F0: 1, F1: false}
                            _ = v10
                        }
                        switch v1 {
                        case "hello":
                            v11 := GenS3{F0: 63}
                            _ = v11
                        case "":
                            if ((2 > 2) || (!true)) {
                                v12 := GenS3{F0: 213}
                                if (!(1000 >= 1000)) {
                                    if (-1 == 1) {
                                        v1 = (((v1 + ("a" + v1)) + ("" + ((v1 + ((v1 + "hello") + v1)) + ((v1 + ("\n\t" + v1)) + ("foo bar" + (v1 + v1)))))) + ((v1 + "") + "a"))
                                        v1 = "foo bar"
                                        v13 := v2[v1]
                                        var v14 uint8 = 0
                                        v1 = "foo bar"
                                        if false {
                                            var v15 string = v1
                                            _ = v15
                                        } else {
                                            fmt.Println(fi6, v13)
                                        }
                                        _ = v13
                                        _ = v14
                                    } else {
                                        v2[""] = 2147483648
                                        var v16 uint8 = 1
                                        v16 += 255
                                        _ = v16
                                    }
                                } else {
                                    switch v1 {
                                    case "\n\t":
                                        fmt.Println(fi6, v1)
                                    case "hello":
                                        var v17 int8 = 0
                                        _ = v17
                                    case "":
                                        var v18 int = 127
                                        _ = v18
                                    }
                                }
                                switch "" {
                                case "hello":
                                    v19 := GenS3{F0: 128}
                                    _ = v19
                                case "":
                                    var v20 uint64 = 1
                                    _ = v20
                                default:
                                    fmt.Println(fi6, v1, fi9)
                                }
                                _ = v12
                            } else {
                                v21 := GenS4{F0: 2, F1: 1}
                                _ = v21
                            }
                        case "foo bar":
                            v22 := GenS2{F0: 384, F1: (0 != 892)}
                            _ = v22
                        default:
                            v23 := GenS1{F0: 42}
                            _ = v23
                        }
                        v1 = v1
                        v24 := []int{(fi9 - (fi9 / (((fi6 * fi9)) | 1))), fi6, (fi9 & fi6)}
                        _ = v24
                    }
                    _ = fi9
                }
            }
        case -9223372036854775808:
            v25 := GenS5{F0: 637, F1: (0 < 32767), F2: fi6}
            _ = v25
        default:
            v26 := []int{-2147483648, (fi6 % ((fi6) | 1)), fi6, (fi6 & fi6)}
            _ = v26
        }
        _ = fi6
    }
    for fi27 := 0; fi27 < 12; fi27++ {
        v28 := GenS6{F0: 42, F1: -1}
        _ = fi27
        _ = v28
    }
    for fi29 := 0; fi29 < 11; fi29++ {
        v30 := []int{fi29, fi29}
        _ = fi29
        _ = v30
    }
    if (42 != 42) {
        v1 = v1
        var v31 float32 = 1.5
        for fi32 := 0; fi32 < 20; fi32++ {
            var v33 float64 = 0.0
            _ = fi32
            _ = v33
        }
        _ = v31
    }
    v2[(v1 + (v1 + v1))] = 1
    for fi34 := 0; fi34 < 19; fi34++ {
        var v35 uint32 = 2
        _ = fi34
        _ = v35
    }
    if false {
        v1 = v1
        if (144 < 255) {
            var v36 uint8 = 2
            for fi37 := 0; fi37 < 12; fi37++ {
                if true {
                    switch (fi37 | (fi37 >> ((4294967295) & 63))) {
                    case -856:
                        for fi38 := 0; fi38 < 16; fi38++ {
                            switch v1 {
                            case "a":
                                if (!false) {
                                    v1 = v1
                                    v36 += (v36 % (((^(v36))) | 1))
                                }
                            case "hello":
                                if false {
                                    if (507 >= 0) {
                                        v1 = "foo bar"
                                        fmt.Println(v36)
                                        var v39 uint16 = 2
                                        _ = v39
                                    }
                                }
                            default:
                                switch (-(fi38)) {
                                case -66:
                                    if false {
                                        v1 = ("hello" + v1)
                                        var v40 int64 = -9223372036854775808
                                        _ = v40
                                    }
                                case 127:
                                    v1 = (v1 + v1)
                                default:
                                    v36 += (v36 >> ((874) & 63))
                                }
                                var v41 int32 = -574
                                _ = v41
                            }
                            fmt.Println(fi38, v36, fi37)
                            _ = fi38
                        }
                    case 128:
                        _ = len(v3)
                    case -2147483648:
                        v42 := GenS4{F0: 490, F1: 42}
                        _ = v42
                    default:
                        if (("" + v1) != v1) {
                            fmt.Println(v36, v1, v36)
                            var v43 int8 = 2
                            _ = v43
                        }
                    }
                    fmt.Println(v1, fi37)
                    var v44 uint8 = (v36 / ((v36) | 1))
                    _ = v44
                } else {
                    v1 = v1
                }
                v36 = 128
                var v45 uint8 = (^(v36))
                _ = fi37
                _ = v45
            }
            _ = v36
        } else {
            var v46 float64 = float64(0.0)
            if false {
                switch v1 {
                case "\n\t":
                    if (2147483647 <= 42) {
                        switch -2147483648 {
                        case 32767:
                            if false {
                                for fi47 := 0; fi47 < 15; fi47++ {
                                    v1 = ""
                                    v1 = ((v1 + v1) + ("foo bar" + v1))
                                    _ = fi47
                                }
                                for fi48 := 0; fi48 < 16; fi48++ {
                                    var v49 int16 = -237
                                    switch v1 {
                                    case "a":
                                        var v50 bool = true
                                        _ = v50
                                    }
                                    fmt.Println(v49, v49)
                                    _ = fi48
                                    _ = v49
                                }
                                if ((419 < 65535) || ((float64(3.14159) + (v46 / (v46 - (((v46 - v46) + float64(1e10)) / 1.0)))) <= ((v46 * 0.1) - v46))) {
                                    fmt.Println(v46, v46)
                                    fmt.Println(v46)
                                    v51 := GenS7{F0: -29, F1: 2147483648, F2: -32768}
                                    _ = v51
                                }
                            } else {
                                v52 := GenS3{F0: 127}
                                v46 += v46
                                _ = v52
                            }
                        case 9223372036854775807:
                            for fi53 := 0; fi53 < 14; fi53++ {
                                for fi54 := 0; fi54 < 10; fi54++ {
                                    v46 -= v46
                                    v46 += 0.0
                                    var v55 int = fi53
                                    if (!((true && (!(128 != 1))) || false)) {
                                        var v56 uint32 = 2
                                        _ = v56
                                    } else {
                                        v57 := GenS2{F0: 0, F1: (!(false || (690 != 42)))}
                                        _ = v57
                                    }
                                    _ = fi54
                                    _ = v55
                                }
                                _ = fi53
                            }
                        case 32768:
                            fmt.Println(v1, v46)
                        }
                        var v58 uint32 = 4294967295
                        _ = v58
                    } else {
                        for fi59 := 0; fi59 < 8; fi59++ {
                            var v60 uint32 = 4294967295
                            _ = fi59
                            _ = v60
                        }
                    }
                }
                v46 = float64((float32((v46 / -2.25)) - 0.1))
            } else {
                v1 = ((v1 + ("foo bar" + ("a" + "foo bar"))) + "foo bar")
            }
            _ = v46
        }
        switch v1 {
        case "a":
            fmt.Println(v1)
        case "hello":
            for fi61 := 0; fi61 < 20; fi61++ {
                switch (fi61 % (((fi61 | fi61)) | 1)) {
                case 1:
                    for fi62 := 0; fi62 < 16; fi62++ {
                        var v63 uint32 = 1
                        v63 = v63
                        v63 += 4294967295
                        switch fi62 {
                        case 32768:
                            delete(v3, 976)
                        }
                        v1 = (v1 + "a")
                        _ = fi62
                        _ = v63
                    }
                }
                _ = fi61
            }
        case "\n\t":
            v64 := GenS3{F0: 243}
            _ = v64
        default:
            for fi65 := 0; fi65 < 12; fi65++ {
                var v66 uint64 = 4294967295
                v1 = "foo bar"
                for fi67 := 0; fi67 < 10; fi67++ {
                    var v68 uint16 = 65535
                    var v69 int64 = -1
                    _ = fi67
                    _ = v68
                    _ = v69
                }
                _ = fi65
                _ = v66
            }
        }
    }
    switch 2147483647 {
    case 965:
        for fi70 := 0; fi70 < 16; fi70++ {
            switch v1 {
            case "\n\t":
                v71 := []int{(fi70 & fi70)}
                _ = v71
            default:
                v72 := GenS8{F0: true, F1: false}
                _ = v72
            }
            v1 = ((("foo bar" + "a") + "foo bar") + v1)
            _ = fi70
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v3
}
