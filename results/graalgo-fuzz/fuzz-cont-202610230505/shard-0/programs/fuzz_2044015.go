package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 int }
type GenS2 struct { F0 uint32; F1 bool }
type GenS3 struct { F0 uint16; F1 int16 }
type GenS4 struct { F0 uint64; F1 string; F2 int64 }
type GenS5 struct { F0 int8; F1 int8 }
type GenS6 struct { F0 int32 }
func genFn0(p0 uint16, p1 int16) uint16 {
    v7 := GenS1{F0: uint32(p0), F1: 333, F2: 0}
    for fi8 := 0; fi8 < 7; fi8++ {
        v9 := GenS1{F0: 42, F1: 954, F2: fi8}
        _ = fi8
        _ = v9
    }
    var v10 float32 = (1e10 * 1.5)
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    return p0
}
func genFn1() int16 {
    v11 := GenS1{F0: 2147483647, F1: 4294967295, F2: 1}
    var v12 uint64 = 4294967295
    v13 := GenS2{F0: 4294967295, F1: true}
    var v14 int = int(v12)
    for fi15 := 0; fi15 < 18; fi15++ {
        var v16 float64 = (0.0 * -2.25)
        v12 = v12
        var v17 string = ((("\n\t" + ("a" + "foo bar")) + "hello") + ("foo bar" + (("foo bar" + "a") + "\n\t")))
        _ = fi15
        _ = v16
        _ = v17
    }
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    return -672
}

func main() {
    var v1 int16 = 42
    v2 := []string{("" + ("a" + "hello")), "foo bar", "", "foo bar"}
    v3 := []int{32768, 0, 32768, 127}
    v4 := []string{("\n\t" + "hello")}
    v5 := make(map[uint32]string)
    v6 := make(map[uint32]bool)
    var v18 uint64 = 0
    var v19 bool = false
    var v20 uint8 = 127
    var v21 uint32 = 4294967295
    var v22 int16 = (v1 & v1)
    fmt.Println(v1, v22, v19)
    for fi23 := 0; fi23 < 7; fi23++ {
        var v24 uint64 = (v18 + (v18 << ((uint(v21)) & 63)))
        _ = fi23
        _ = v24
    }
    var v25 uint64 = (v18 - uint64(v21))
    if v19 {
        if (true || v19) {
            if v19 {
                var v26 int16 = genFn1()
                var v27 int16 = (v26 - (v1 - (v1 / ((genFn1()) | 1))))
                v28 := GenS2{F0: v21, F1: v19}
                _ = v26
                _ = v27
                _ = v28
            }
            switch 128 {
            case 128:
                _ = len(v5)
            case 2147483647:
                v29 := GenS3{F0: 2, F1: (-(v22))}
                _ = v29
            default:
                switch 4294967295 {
                case 94:
                    var v30 uint64 = 0
                    _ = v30
                default:
                    if (0.0 != 0.1) {
                        var v31 int8 = 100
                        if (-119 < -2147483648) {
                            var v32 uint16 = 879
                            v33 := GenS1{F0: (v21 << ((4294967295) & 63)), F1: v25, F2: 1}
                            _ = v32
                            _ = v33
                        } else {
                            for fi34 := 0; fi34 < 1; fi34++ {
                                var v35 int16 = 1000
                                _ = fi34
                                _ = v35
                            }
                        }
                        _ = v31
                    } else {
                        for fi36 := 0; fi36 < 14; fi36++ {
                            for fi37 := 0; fi37 < 3; fi37++ {
                                fmt.Println(v18, fi36)
                                var v38 float32 = ((1e10 / ((1.5 / 1.0) * 3.14159)) / 0.1)
                                _ = fi37
                                _ = v38
                            }
                            v5[42] = ("\n\t" + "\n\t")
                            if ((0.0 / -2.25) > 0.1) {
                                switch fi36 {
                                case 32767:
                                    if false {
                                        switch 954 {
                                        case 4294967295:
                                            for fi39 := 0; fi39 < 10; fi39++ {
                                                var v40 int = (fi39 << ((469) & 63))
                                                _ = fi39
                                                _ = v40
                                            }
                                        case -9223372036854775808:
                                            v18 -= (v25 - v25)
                                        default:
                                            if (42 == 22) {
                                                v19 = (((530 == 4294967295) || v19) || (!(true || v19)))
                                                switch "a" {
                                                case "foo bar":
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                default:
                                                    var v41 uint16 = 65535
                                                    _ = v41
                                                }
                                                v42 := GenS1{F0: 696, F1: 1, F2: (^(fi36))}
                                                _ = v42
                                            }
                                        }
                                    } else {
                                        switch -2147483648 {
                                        case 749:
                                            if (63 > 285) {
                                                v19 = (("\n\t" >= "") && v19)
                                                v19 = (-9223372036854775808 <= 42)
                                                var v43 int8 = 100
                                                _ = v43
                                            } else {
                                                v18 = uint64(v20)
                                                v44 := v6[v21]
                                                var v45 uint = 42
                                                _ = v44
                                                _ = v45
                                            }
                                        case -2147483648:
                                            switch 0 {
                                            case 42:
                                                var v46 float64 = ((1.5 + float64(((3.14159 / 1e10) * 3.14159))) + float64(0.0))
                                                _ = v46
                                            case -9223372036854775808:
                                                v47 := v6[uint32(v18)]
                                                _ = v47
                                            default:
                                                fmt.Println(v21, v21, v19)
                                            }
                                        case 128:
                                            fmt.Println(fi36, v19)
                                        }
                                        var v48 string = "\n\t"
                                        _ = v48
                                    }
                                case -9223372036854775808:
                                    if (!(100 > 0)) {
                                        v21 += v21
                                        if (!v19) {
                                            v49 := GenS4{F0: (^(v18)), F1: "a", F2: int64(v1)}
                                            if (false || (true && v19)) {
                                                v19 = (v19 && v19)
                                            } else {
                                                var v50 uint16 = 5
                                                if (2 < 1) {
                                                    if (1 == (v22 - genFn1())) {
                                                        switch -2147483648 {
                                                        case 32767:
                                                            v19 = v19
                                                        default:
                                                            v51 := GenS1{F0: v21, F1: 2, F2: 4294967295}
                                                            _ = v51
                                                        }
                                                        var v52 uint = 2
                                                        _ = v52
                                                    } else {
                                                        if ((^(v21)) >= v21) {
                                                            v20 -= v20
                                                        } else {
                                                            var v53 int8 = 38
                                                            var v54 uint8 = v20
                                                            _ = v53
                                                            _ = v54
                                                        }
                                                        v20 = v20
                                                        var v55 string = "hello"
                                                        _ = v55
                                                    }
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                }
                                                _ = v50
                                            }
                                            _ = len(v5)
                                            _ = v49
                                        } else {
                                            var v56 float32 = (float32(1e10) - -2.25)
                                            var v57 uint64 = v18
                                            _ = v56
                                            _ = v57
                                        }
                                        for fi58 := 0; fi58 < 8; fi58++ {
                                            v1 = genFn1()
                                            _ = fi58
                                        }
                                    } else {
                                        v1 = 830
                                    }
                                default:
                                    var v59 uint32 = (^(v21))
                                    _ = v59
                                }
                                _ = len(v6)
                            }
                            _ = fi36
                        }
                    }
                }
            }
        }
        var v60 uint = 42
        v2 = append(v2, "hello")
        _ = v60
    } else {
        var v61 uint32 = 42
        for fi62 := 0; fi62 < 6; fi62++ {
            var v63 int64 = 9223372036854775807
            for fi64 := 0; fi64 < 17; fi64++ {
                delete(v6, v21)
                _ = fi64
            }
            _ = fi62
            _ = v63
        }
        _ = v61
    }
    v25 += 2
    if true {
        var v65 uint = 65535
        fmt.Println(v18, v22, v25)
        for fi66 := 0; fi66 < 10; fi66++ {
            v67 := GenS2{F0: v21, F1: false}
            switch "" {
            case "hello":
                fmt.Println(fi66)
            default:
                if v19 {
                    var v68 uint16 = 0
                    switch (fi66 % ((fi66) | 1)) {
                    case 4294967295:
                        v21 -= 1
                    case -9223372036854775808:
                        switch fi66 {
                        case 582:
                            v1 -= 42
                        case 0:
                            var v69 bool = ((((!false) && v19) || false) || (genFn0((v68 & genFn0(42, int16(fi66))), -32768) != (^(v68))))
                            _ = v69
                        case 32767:
                            v70 := GenS5{F0: int8(v20), F1: 42}
                            _ = v70
                        }
                    }
                    _ = v68
                } else {
                    for fi71 := 0; fi71 < 12; fi71++ {
                        if (!(1 >= -1)) {
                            v72 := v6[v21]
                            _ = len(v2)
                            _ = v72
                        }
                        var v73 uint = v65
                        _ = fi71
                        _ = v73
                    }
                    fmt.Println(v25, v25)
                    var v74 int64 = -726
                    _ = v74
                }
            }
            if v19 {
                for fi75 := 0; fi75 < 15; fi75++ {
                    var v76 string = "\n\t"
                    switch v76 {
                    case "a":
                        v76 = v76
                    case "\n\t":
                        v22 = (v1 % ((genFn1()) | 1))
                    }
                    v77 := GenS5{F0: -128, F1: -128}
                    var v78 uint64 = (v18 | (v18 | (v25 & v25)))
                    if v19 {
                        switch (v76 + v76) {
                        case "hello":
                            var v79 uint16 = 32767
                            _ = v79
                        default:
                            var v80 int32 = 42
                            _ = v80
                        }
                    } else {
                        fmt.Println(v22, v19)
                    }
                    var v81 int64 = 9223372036854775807
                    var v82 float32 = (0.1 + (((((((0.0 + 3.14159) * (0.1 - -2.25)) / -2.25) * ((3.14159 / 1e10) * 1.5)) / ((3.14159 - 1e10) + (float32(1.5) / -2.25))) * 3.14159) / 1e10))
                    _ = fi75
                    _ = v76
                    _ = v77
                    _ = v78
                    _ = v81
                    _ = v82
                }
                fmt.Println(v18)
            }
            _ = fi66
            _ = v67
        }
        _ = v65
    } else {
        v21 = 65535
        if len(v4) > 0 {
            v4[0] = ("" + "foo bar")
        }
    }
    v22 = -32768
    switch "" {
    case "":
        v21 = uint32(v20)
    case "a":
        v83 := GenS3{F0: 1, F1: (v22 * (v22 ^ v1))}
        _ = v83
    case "foo bar":
        if ("foo bar" != "") {
            switch "" {
            case "\n\t":
                var v84 int = -517
                _ = v84
            case "":
                for fi85 := 0; fi85 < 1; fi85++ {
                    v22 = (v22 % (((v1 | (-(v22)))) | 1))
                    _ = fi85
                }
            case "foo bar":
                var v86 uint8 = v20
                _ = v86
            }
        }
    }
    for fi87 := 0; fi87 < 6; fi87++ {
        v18 -= (v25 + (v18 >> ((701) & 63)))
        _ = fi87
    }
    if (!v19) {
        if v19 {
            var v88 int32 = -1
            v89 := GenS5{F0: 1, F1: -92}
            _ = v88
            _ = v89
        }
        fmt.Println(v1, v18)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    } else {
        fmt.Println(v20, v19, v22)
        switch "a" {
        case "":
            var v90 int16 = (v22 % (((v22 | (v22 | (v22 ^ -493)))) | 1))
            _ = v90
        case "foo bar":
            var v91 uint = 65535
            _ = v91
        }
        if ((("" == "a") || (v20 != uint8(v21))) || (!(255 > 709))) {
            for fi92 := 0; fi92 < 12; fi92++ {
                _ = len(v5)
                v93 := GenS6{F0: 2147483647}
                _ = fi92
                _ = v93
            }
            fmt.Println(v25, v25, v22)
            var v94 uint32 = v21
            _ = v94
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
