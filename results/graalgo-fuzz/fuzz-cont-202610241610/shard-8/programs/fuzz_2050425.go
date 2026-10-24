package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 int16 }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int16; F1 uint8; F2 int64 }
type GenS5 struct { F0 bool; F1 uint8 }
type GenS6 struct { F0 int16; F1 uint; F2 int8 }
type GenS7 struct { F0 uint8; F1 int64 }
func genFn0() int8 {
    if true {
        v7 := []int{}
        _ = v7
    } else {
        var v8 uint32 = 379
        v8 = v8
        var v9 uint = 42
        _ = v8
        _ = v9
    }
    v10 := []int{0, 32768, 32767}
    var v11 int16 = 992
    var v12 int = 9223372036854775807
    _ = v10
    _ = v11
    _ = v12
    return 11
}
func genFn1() int64 {
    var v13 float64 = ((1e10 * -2.25) / (-2.25 / 1.0))
    v13 -= (v13 + v13)
    v14 := []int{-1, 4294967295, 648, 42}
    v14 = append(v14, -2147483648)
    fmt.Println(v13)
    _ = v13
    _ = v14
    return 9223372036854775807
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]string)
    var v3 uint32 = 728
    v4 := make(map[int]string)
    v5 := []int{}
    v6 := make(map[int]bool)
    v3 = (v3 >> ((273) & 63))
    for fi15 := 0; fi15 < 16; fi15++ {
        if false {
            var v16 uint32 = 2
            v17 := GenS1{F0: 128, F1: 116, F2: 348}
            if (-128 > genFn0()) {
                for fi18 := 0; fi18 < 14; fi18++ {
                    v19 := GenS1{F0: 165, F1: 2, F2: -357}
                    v20 := GenS1{F0: 127, F1: 255, F2: 993}
                    if (((fi18 - (^(fi18))) >= fi15) && false) {
                        v5 = append(v5, (fi18 / (((fi18 + (fi18 - -9223372036854775808))) | 1)))
                    }
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                var v21 uint32 = v16
                _ = v21
            } else {
                v16 = v16
            }
            _ = v16
            _ = v17
        } else {
            if (121 < 102) {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                v22 := v4[(-(fi15))]
                for fi23 := 0; fi23 < 6; fi23++ {
                    _ = len(v5)
                    v24 := GenS2{F0: 820, F1: true}
                    _ = fi23
                    _ = v24
                }
                _ = v22
            } else {
                for fi25 := 0; fi25 < 10; fi25++ {
                    v3 = (v3 & (v3 >> ((65535) & 63)))
                    v3 = 2147483648
                    v2[(-(fi15))] = "foo bar"
                    _ = fi25
                }
                switch fi15 {
                case -409:
                    switch ("foo bar" + "hello") {
                    case "foo bar":
                        for fi26 := 0; fi26 < 13; fi26++ {
                            var v27 int = -9223372036854775808
                            var v28 int16 = 332
                            var v29 int32 = 0
                            _ = fi26
                            _ = v27
                            _ = v28
                            _ = v29
                        }
                    case "hello":
                        v30 := GenS3{F0: true}
                        _ = v30
                    default:
                        var v31 uint32 = 1
                        _ = v31
                    }
                case 0:
                    var v32 int8 = 127
                    _ = v32
                }
                var v33 int = fi15
                _ = v33
            }
            for fi34 := 0; fi34 < 13; fi34++ {
                if true {
                    var v35 int64 = -1
                    switch "a" {
                    case "hello":
                        switch "foo bar" {
                        case "a":
                            var v36 uint8 = 120
                            _ = v36
                        case "hello":
                            if (333 > 769) {
                                var v37 int64 = (v35 ^ v35)
                                var v38 string = "foo bar"
                                switch (fi15 * (fi15 % ((2147483647) | 1))) {
                                case 42:
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                case 582:
                                    if false {
                                        v35 = v37
                                        if (-2.25 >= 3.14159) {
                                            v5 = append(v5, (fi34 * 4294967295))
                                        }
                                    } else {
                                        v39 := GenS3{F0: false}
                                        if ((127 == 255) && true) {
                                            switch (fi15 | 4294967295) {
                                            case -9223372036854775808:
                                                if ((true && ("foo bar" == "a")) || true) {
                                                    v40 := GenS1{F0: 2, F1: 127, F2: 1000}
                                                    var v41 float64 = float64(1.5)
                                                    switch (fi15 % ((fi15) | 1)) {
                                                    case -2147483648:
                                                        var v42 int16 = 1
                                                        _ = v42
                                                    }
                                                    _ = v40
                                                    _ = v41
                                                }
                                            case -1:
                                                var v43 uint = 42
                                                _ = v43
                                            case 128:
                                                v37 = -1
                                            }
                                        }
                                        if (42 == 884) {
                                            fmt.Println(v35)
                                            v2[fi15] = ((v38 + v38) + "foo bar")
                                        }
                                        _ = v39
                                    }
                                case 32768:
                                    _ = len(v4)
                                default:
                                    if ((255 == 2) && false) {
                                        if false {
                                            v37 -= genFn1()
                                            v44 := GenS2{F0: 42, F1: false}
                                            _ = v44
                                        } else {
                                            var v45 int32 = -347
                                            v46 := GenS2{F0: -74, F1: false}
                                            var v47 uint8 = 42
                                            _ = v45
                                            _ = v46
                                            _ = v47
                                        }
                                        if (827 > 1) {
                                            if ((^(v35)) >= genFn1()) {
                                                v48 := GenS1{F0: 128, F1: 106, F2: 1000}
                                                if true {
                                                    var v49 uint = 0
                                                    var v50 uint8 = 255
                                                    _ = v49
                                                    _ = v50
                                                }
                                                v51 := GenS4{F0: 32767, F1: 235, F2: 9223372036854775807}
                                                _ = v48
                                                _ = v51
                                            }
                                            if false {
                                                _ = len(v2)
                                            } else {
                                                v52 := GenS2{F0: -278, F1: (false || true)}
                                                v37 += v37
                                                v3 = (v3 * 4294967295)
                                                _ = v52
                                            }
                                        }
                                        fmt.Println(v38)
                                    }
                                }
                                _ = v37
                                _ = v38
                            } else {
                                var v53 uint32 = (^(v3))
                                var v54 uint16 = 65535
                                _ = v53
                                _ = v54
                            }
                        }
                    case "":
                        switch (fi34 & fi34) {
                        case 32768:
                            var v55 int16 = 0
                            _ = v55
                        case 853:
                            v56 := v4[(fi15 % ((-568) | 1))]
                            _ = v56
                        case 678:
                            fmt.Println(v3, v3)
                        }
                    case "\n\t":
                        v3 = (v3 + v3)
                    }
                    var v57 int8 = 0
                    _ = v35
                    _ = v57
                }
                _ = fi34
            }
            for fi58 := 0; fi58 < 1; fi58++ {
                v3 = v3
                fmt.Println(fi15, fi58, fi58)
                v59 := GenS5{F0: true, F1: 2}
                _ = fi58
                _ = v59
            }
        }
        for fi60 := 0; fi60 < 3; fi60++ {
            v3 += 65535
            if ((963 < 255) || true) {
                var v61 int8 = -1
                if (42 > 2) {
                    var v62 int16 = 32767
                    _ = v62
                } else {
                    v61 -= genFn0()
                }
                delete(v4, fi60)
                _ = v61
            } else {
                v3 = (v3 >> ((308) & 63))
                if true {
                    if len(v5) > 0 {
                        v5[0] = (-(fi60))
                    }
                    v63 := GenS1{F0: 121, F1: uint8(v3), F2: -13}
                    v3 = v3
                    _ = v63
                } else {
                    if ((-2.25 / 1.0) > 1.5) {
                        v64 := v4[fi60]
                        var v65 bool = false
                        _ = v64
                        _ = v65
                    } else {
                        switch "a" {
                        case "foo bar":
                            v3 += v3
                        case "\n\t":
                            if (!(439 == 659)) {
                                var v66 int = (^(fi15))
                                v3 = 2147483648
                                switch "a" {
                                case "foo bar":
                                    switch v66 {
                                    case 42:
                                        var v67 uint16 = 140
                                        _ = v67
                                    default:
                                        delete(v1, v3)
                                    }
                                case "\n\t":
                                    var v68 int64 = -245
                                    v68 += v68
                                    _ = v68
                                default:
                                    v3 += v3
                                    v3 = (v3 & (v3 | 4294967295))
                                    switch v66 {
                                    case -122:
                                        v66 -= 769
                                    case -263:
                                        fmt.Println(fi60, fi60)
                                        v3 = (v3 >> ((4294967295) & 63))
                                        var v69 int = (v66 | v66)
                                        v3 -= (v3 & 4294967295)
                                        var v70 int = (fi60 + v66)
                                        _ = v69
                                        _ = v70
                                    default:
                                        fmt.Println(v66)
                                    }
                                }
                                v3 = (v3 >> ((4294967295) & 63))
                                var v71 uint64 = 18446744073709551615
                                _ = v66
                                _ = v71
                            }
                        default:
                            var v72 int64 = 25
                            _ = v72
                        }
                    }
                    v3 = v3
                    if (1 <= -930) {
                        v73 := GenS6{F0: 1, F1: 2, F2: 99}
                        v74 := GenS7{F0: 128, F1: 9223372036854775807}
                        var v75 int32 = 376
                        _ = v73
                        _ = v74
                        _ = v75
                    } else {
                        switch 32768 {
                        case 128:
                            var v76 uint32 = 1
                            _ = v76
                        }
                        v3 = v3
                        var v77 uint32 = v3
                        _ = v77
                    }
                }
            }
            var v78 float64 = 1e10
            _ = fi60
            _ = v78
        }
        for fi79 := 0; fi79 < 17; fi79++ {
            var v80 uint16 = 2
            _ = fi79
            _ = v80
        }
        _ = fi15
    }
    if (!false) {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    } else {
        switch "" {
        case "\n\t":
            v3 = (v3 << ((42) & 63))
        }
        v3 = (v3 >> ((18446744073709551615) & 63))
    }
    v3 = (^(v3))
    for fi81 := 0; fi81 < 8; fi81++ {
        switch "foo bar" {
        case "foo bar":
            if ((v3 - v3) > 4294967295) {
                var v82 int8 = int8(fi81)
                for fi83 := 0; fi83 < 15; fi83++ {
                    switch fi81 {
                    case -873:
                        var v84 float32 = (float32((1.5 / 1.5)) + 0.0)
                        _ = v84
                    case -9223372036854775808:
                        if (((v82 <= 6) || (!true)) && (76 > 1000)) {
                            v82 -= (-(v82))
                            fmt.Println(fi83, v3)
                        } else {
                            _ = len(v2)
                            v82 -= v82
                            switch (fi83 << ((274) & 63)) {
                            case 128:
                                fmt.Println(v82, v82, v82)
                            case -9223372036854775808:
                                var v85 string = "hello"
                                _ = v85
                            default:
                                if len(v5) > 0 {
                                    v5[0] = 9223372036854775807
                                }
                            }
                            var v86 string = "hello"
                            v3 += v3
                            v3 = (v3 ^ (v3 & 221))
                            _ = v86
                        }
                    default:
                        _ = len(v1)
                    }
                    _ = fi83
                }
                v3 -= v3
                _ = v82
            }
        case "hello":
            var v87 uint64 = 4294967295
            _ = v87
        default:
            fmt.Println(v3, fi81)
        }
        var v88 int8 = -1
        var v89 string = ("\n\t" + (("a" + "hello") + ""))
        _ = fi81
        _ = v88
        _ = v89
    }
    var v90 int16 = 52
    fmt.Println(v3, v3)
    fmt.Println(v90)
    var v91 int8 = genFn0()
    v92 := GenS2{F0: 0, F1: (2 != 42)}
    v3 = (v3 << ((18446744073709551615) & 63))
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if len(v5) > 0 {
        v5[0] = -9223372036854775808
    }
    var v93 uint16 = 835
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v90:", v90)
    fmt.Println("v91:", v91)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v90
    _ = v91
    _ = v92
    _ = v93
}
