package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 int64; F1 string }
type GenS4 struct { F0 int8; F1 bool }
type GenS5 struct { F0 int8; F1 uint64; F2 int64 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint8; F1 uint8; F2 int }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint = 4294967295
    var v3 bool = false
    v4 := []uint32{2147483647}
    var v5 int64 = 1
    var v6 int32 = -2147483648
    var v7 float64 = 0.0
    v2 += v2
    v5 = v5
    v8 := GenS1{F0: 1, F1: 2147483647}
    v5 = v5
    for fi9 := 0; fi9 < 19; fi9++ {
        var v10 uint16 = 42
        if v3 {
            var v11 int8 = 0
            if v3 {
                var v12 bool = false
                switch fi9 {
                case -9223372036854775808:
                    fmt.Println(v11, v7)
                case 843:
                    for fi13 := 0; fi13 < 20; fi13++ {
                        var v14 float32 = 3.14159
                        _ = fi13
                        _ = v14
                    }
                default:
                    v3 = (65535 != 42)
                }
                var v15 uint = (v2 | v2)
                _ = v12
                _ = v15
            } else {
                fmt.Println(v5, v10, v6)
            }
            if (v3 && ((65535 <= 758) || (!(32767 != v10)))) {
                var v16 uint = (v2 / ((v2) | 1))
                var v17 uint32 = 909
                _ = v16
                _ = v17
            } else {
                switch "\n\t" {
                case "\n\t":
                    v18 := GenS2{F0: (fi9 / (((^(fi9))) | 1)), F1: fi9}
                    _ = v18
                case "a":
                    switch (fi9 + (fi9 << ((0) & 63))) {
                    case 128:
                        _ = len(v4)
                    case 42:
                        if len(v4) > 0 {
                            v4[0] = 2147483647
                        }
                    }
                case "foo bar":
                    v5 -= v5
                }
                var v19 int64 = (v5 | v5)
                _ = v19
            }
            _ = v11
        }
        _ = fi9
        _ = v10
    }
    v20 := GenS1{F0: 1, F1: 548}
    for fi21 := 0; fi21 < 11; fi21++ {
        fmt.Println(v6)
        _ = fi21
    }
    for fi22 := 0; fi22 < 3; fi22++ {
        if (1 > 929) {
            switch fi22 {
            case 42:
                v23 := GenS3{F0: v5, F1: "hello"}
                _ = v23
            case 9223372036854775807:
                switch (("" + "foo bar") + ("" + "a")) {
                case "":
                    var v24 uint32 = 1
                    _ = v24
                case "hello":
                    v3 = (v3 || (int8(v6) > 42))
                default:
                    v3 = (2 != 907)
                }
            case 2147483647:
                var v25 int8 = -32
                _ = v25
            default:
                var v26 bool = (v3 && (500 != 298))
                _ = v26
            }
            v2 = v2
        } else {
            var v27 uint64 = 18446744073709551615
            fmt.Println(fi22, v3)
            for fi28 := 0; fi28 < 19; fi28++ {
                if v3 {
                    v29 := GenS2{F0: fi22, F1: (fi22 % ((fi28) | 1))}
                    fmt.Println(fi22, v3)
                    v27 += v27
                    _ = v29
                }
                v30 := v1[2]
                v31 := GenS4{F0: 111, F1: v3}
                _ = fi28
                _ = v30
                _ = v31
            }
            _ = v27
        }
        _ = fi22
    }
    for fi32 := 0; fi32 < 20; fi32++ {
        switch (fi32 - (fi32 % ((fi32) | 1))) {
        case -532:
            fmt.Println(v6, v3, v2)
        default:
            for fi33 := 0; fi33 < 6; fi33++ {
                v34 := GenS5{F0: -99, F1: 650, F2: -1}
                var v35 float64 = ((float64(1.5) / 1e10) - v7)
                _ = fi33
                _ = v34
                _ = v35
            }
        }
        var v36 int32 = v6
        v36 = (v6 << ((v2) & 63))
        _ = fi32
        _ = v36
    }
    if ((!((1e10 * 1.5) > v7)) && false) {
        fmt.Println(v6, v5, v2)
        switch -2147483648 {
        case -274:
            v5 += -1
        }
        v37 := GenS4{F0: -120, F1: v3}
        _ = v37
    } else {
        v3 = (-1 < v6)
        v38 := GenS4{F0: 66, F1: false}
        var v39 uint16 = 523
        _ = v38
        _ = v39
    }
    var v40 int32 = v6
    if v3 {
        if (-952 <= -2147483648) {
            for fi41 := 0; fi41 < 7; fi41++ {
                if (!(42 == 1)) {
                    for fi42 := 0; fi42 < 2; fi42++ {
                        v43 := v1[4294967295]
                        switch (("hello" + "foo bar") + "hello") {
                        case "":
                            var v44 int32 = (v6 % ((v6) | 1))
                            var v45 string = (("" + "a") + (("" + "hello") + "\n\t"))
                            v4 = append(v4, 65535)
                            switch (^(fi42)) {
                            case -240:
                                if v3 {
                                    v4 = append(v4, 881)
                                    if false {
                                        v46 := GenS5{F0: -128, F1: 660, F2: v5}
                                        v45 = ("" + v45)
                                        _ = v46
                                    }
                                } else {
                                    var v47 int16 = 32767
                                    _ = v47
                                }
                            default:
                                var v48 bool = (v5 < (-(v5)))
                                _ = v48
                            }
                            var v49 string = (v45 + "")
                            v44 -= (v44 + 2147483647)
                            _ = v44
                            _ = v45
                            _ = v49
                        case "foo bar":
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            switch v43 {
                            case 128:
                                v40 = int32(fi41)
                            default:
                                var v50 uint = v2
                                _ = v50
                            }
                        case "\n\t":
                            var v51 int64 = (v5 >> (((v2 / ((255) | 1))) & 63))
                            _ = v51
                        default:
                            v2 = (^(v2))
                        }
                        v52 := GenS1{F0: 204, F1: (fi42 << (((v2 << (((v2 - v2)) & 63))) & 63))}
                        _ = fi42
                        _ = v43
                        _ = v52
                    }
                    var v53 int32 = v40
                    _ = v53
                } else {
                    switch (fi41 >> (((^(v2))) & 63)) {
                    case 4294967295:
                        switch ("hello" + ("\n\t" + ("\n\t" + ""))) {
                        case "a":
                            v7 -= v7
                        default:
                            delete(v1, 42)
                        }
                    case -1:
                        v5 -= -420
                    case 128:
                        switch "hello" {
                        case "\n\t":
                            switch (fi41 + -9223372036854775808) {
                            case 382:
                                v54 := GenS1{F0: 32767, F1: (-(fi41))}
                                _ = v54
                            case 4294967295:
                                v3 = v3
                            case 128:
                                if (851 > 2) {
                                    v7 += v7
                                } else {
                                    if (-1 > -128) {
                                        var v55 float32 = (float32(v7) - ((3.14159 * (float32(float64(-2.25)) - 3.14159)) / 1.0))
                                        switch "" {
                                        case "a":
                                            v56 := GenS2{F0: fi41, F1: (fi41 ^ fi41)}
                                            _ = v56
                                        case "hello":
                                            v55 -= (float32(v7) - (float32(((float64(((3.14159 - v55) - 0.1)) + float64(v55)) * 0.0)) + (1.5 * v55)))
                                        case "\n\t":
                                            v40 += 1
                                        }
                                        _ = v55
                                    } else {
                                        v5 -= -1
                                    }
                                }
                            default:
                                v57 := GenS3{F0: (v5 << (((v2 & (v2 << (((v2 & v2)) & 63)))) & 63)), F1: "hello"}
                                _ = v57
                            }
                        case "hello":
                            switch "" {
                            case "hello":
                                v6 = (v40 - v6)
                            case "":
                                if (v3 && ((!(v3 || (295 > 1))) || v3)) {
                                    var v58 uint8 = 1
                                    _ = v58
                                }
                            case "\n\t":
                                v59 := v1[4294967295]
                                _ = v59
                            default:
                                var v60 uint32 = 42
                                _ = v60
                            }
                        case "":
                            if (1 < 0) {
                                v61 := GenS5{F0: 0, F1: 70, F2: (v5 << (((v2 >> ((4294967295) & 63))) & 63))}
                                _ = len(v1)
                                v62 := GenS5{F0: 1, F1: 306, F2: (v5 << ((v2) & 63))}
                                _ = v61
                                _ = v62
                            } else {
                                if (("a" <= "a") || true) {
                                    if true {
                                        v63 := v1[1]
                                        _ = v63
                                    }
                                    delete(v1, 2147483647)
                                    fmt.Println(v2, v6)
                                } else {
                                    v64 := v1[580]
                                    if (v3 || v3) {
                                        if v3 {
                                            switch fi41 {
                                            case 1:
                                                var v65 uint16 = 65535
                                                _ = v65
                                            case 42:
                                                for fi66 := 0; fi66 < 19; fi66++ {
                                                    v67 := v1[737]
                                                    _ = fi66
                                                    _ = v67
                                                }
                                            default:
                                                if true {
                                                    for fi68 := 0; fi68 < 1; fi68++ {
                                                        v40 += (v6 ^ v6)
                                                        _ = fi68
                                                    }
                                                }
                                            }
                                            v69 := GenS6{F0: 2}
                                            var v70 float64 = ((v7 + v7) - float64(float32((v7 / (v7 + float64((3.14159 - 0.1)))))))
                                            _ = v69
                                            _ = v70
                                        } else {
                                            var v71 string = (("" + "a") + ("" + "a"))
                                            _ = v71
                                        }
                                        var v72 int32 = v6
                                        _ = v72
                                    }
                                    fmt.Println(v3)
                                    _ = v64
                                }
                            }
                        }
                    }
                    switch (fi41 ^ fi41) {
                    case 0:
                        v4 = append(v4, 154)
                    }
                    var v73 uint64 = 0
                    _ = v73
                }
                _ = fi41
            }
        } else {
            v6 += v6
            if ((!v3) || (!(0 != 42))) {
                v74 := GenS7{F0: 213, F1: 127, F2: 4294967295}
                _ = v74
            } else {
                for fi75 := 0; fi75 < 10; fi75++ {
                    for fi76 := 0; fi76 < 11; fi76++ {
                        var v77 int64 = 9223372036854775807
                        _ = fi76
                        _ = v77
                    }
                    var v78 int16 = 906
                    var v79 uint32 = 2147483648
                    _ = fi75
                    _ = v78
                    _ = v79
                }
                switch -71 {
                case -169:
                    fmt.Println(v40, v3, v40)
                case -986:
                    v80 := GenS7{F0: 144, F1: 0, F2: 662}
                    _ = v80
                }
                var v81 int = -495
                _ = v81
            }
            for fi82 := 0; fi82 < 4; fi82++ {
                if v3 {
                    if (1 <= 815) {
                        _ = len(v4)
                    }
                    var v83 uint8 = 93
                    if len(v4) > 0 {
                        v4[0] = 2147483647
                    }
                    _ = v83
                }
                v5 = v5
                _ = fi82
            }
        }
        v2 = v2
    }
    var v84 uint32 = 4294967295
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v84
}
