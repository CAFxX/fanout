package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 string; F2 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint64; F1 uint32; F2 uint }
type GenS5 struct { F0 uint64; F1 int8 }
type GenS6 struct { F0 uint16; F1 uint16; F2 uint8 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint32; F1 uint16 }
type GenS9 struct { F0 string; F1 uint16; F2 uint }

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{("foo bar" + ""), ("" + "")}
    v3 := make(map[string]string)
    v4 := make(map[int]string)
    v5 := make(map[uint32]bool)
    var v6 int16 = -1
    delete(v1, 0)
    if len(v2) > 0 {
        v2[0] = "hello"
    }
    var v7 uint8 = 128
    if (float64(1.5) < (3.14159 / 1.0)) {
        for fi8 := 0; fi8 < 19; fi8++ {
            v6 -= v6
            for fi9 := 0; fi9 < 9; fi9++ {
                var v10 float64 = -2.25
                var v11 uint8 = v7
                var v12 uint16 = 309
                _ = fi9
                _ = v10
                _ = v11
                _ = v12
            }
            v13 := v4[fi8]
            _ = fi8
            _ = v13
        }
        for fi14 := 0; fi14 < 18; fi14++ {
            v15 := GenS1{F0: 880}
            _ = fi14
            _ = v15
        }
    } else {
        for fi16 := 0; fi16 < 15; fi16++ {
            v6 -= 1
            v6 -= (v6 & -32768)
            _ = fi16
        }
        for fi17 := 0; fi17 < 9; fi17++ {
            v18 := GenS2{F0: -201, F1: "a", F2: 255}
            _ = fi17
            _ = v18
        }
        fmt.Println(v7, v7)
    }
    switch -9223372036854775808 {
    case -930:
        var v19 uint16 = 37
        _ = v19
    case 2147483647:
        for fi20 := 0; fi20 < 3; fi20++ {
            var v21 uint16 = 1
            _ = fi20
            _ = v21
        }
    case 9223372036854775807:
        var v22 uint16 = 663
        _ = v22
    }
    if (false || ((false && (-9223372036854775808 < 486)) || true)) {
        v23 := GenS1{F0: 782}
        v6 = (v6 + 42)
        _ = v23
    } else {
        v24 := GenS1{F0: 730}
        var v25 int32 = 225
        _ = v24
        _ = v25
    }
    delete(v5, 42)
    v7 -= v7
    v7 = 34
    switch 32767 {
    case 32767:
        var v26 int16 = 207
        _ = v26
    case 32768:
        if true {
            var v27 float32 = float32(((float64(0.1) + 0.1) * 0.0))
            _ = v27
        } else {
            v28 := GenS1{F0: 65535}
            v6 -= (v6 + v6)
            var v29 int8 = 0
            _ = v28
            _ = v29
        }
    case 128:
        for fi30 := 0; fi30 < 17; fi30++ {
            v7 = v7
            _ = fi30
        }
    }
    switch 2147483647 {
    case 128:
        for fi31 := 0; fi31 < 12; fi31++ {
            if false {
                var v32 uint64 = 2
                switch "a" {
                case "\n\t":
                    var v33 uint32 = 2147483648
                    _ = v33
                case "a":
                    for fi34 := 0; fi34 < 3; fi34++ {
                        switch fi34 {
                        case 32768:
                            switch "\n\t" {
                            case "\n\t":
                                var v35 uint = 2
                                v32 += 42
                                _ = len(v3)
                                _ = v35
                            }
                        case 0:
                            fmt.Println(v7, fi31, fi34)
                        case 724:
                            if (false || (!false)) {
                                if true {
                                    if ((((0.1 - 3.14159) - (-2.25 - (0.0 - 0.1))) / (1.5 * 3.14159)) <= float32(float64(1e10))) {
                                        v36 := GenS2{F0: v6, F1: "a", F2: 255}
                                        if (true || true) {
                                            if ((617 >= 231) && false) {
                                                var v37 int64 = -1
                                                if true {
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    _ = len(v3)
                                                    var v38 int16 = 32767
                                                    _ = v38
                                                }
                                                var v39 int = (fi31 >> ((323) & 63))
                                                _ = v37
                                                _ = v39
                                            } else {
                                                switch ("foo bar" + ("a" + (("" + "a") + "a"))) {
                                                case "\n\t":
                                                    v2 = append(v2, "hello")
                                                    v40 := GenS3{F0: v6}
                                                    var v41 int32 = -2147483648
                                                    _ = v40
                                                    _ = v41
                                                case "":
                                                    v7 = 128
                                                case "hello":
                                                    var v42 uint = 42
                                                    _ = v42
                                                }
                                            }
                                            v43 := GenS3{F0: v6}
                                            var v44 int = fi34
                                            if true {
                                                v45 := v1[2147483648]
                                                var v46 float32 = (3.14159 - 3.14159)
                                                var v47 uint16 = 1
                                                _ = v45
                                                _ = v46
                                                _ = v47
                                            }
                                            var v48 uint32 = 4294967295
                                            switch "hello" {
                                            case "a":
                                                fmt.Println(v6)
                                            case "":
                                                v49 := GenS3{F0: v6}
                                                _ = v49
                                            default:
                                                switch "\n\t" {
                                                case "":
                                                    v7 = v7
                                                case "\n\t":
                                                    delete(v4, (^(fi31)))
                                                default:
                                                    fmt.Println(fi34)
                                                }
                                                switch (fi34 ^ (fi31 - 32767)) {
                                                case -2147483648:
                                                    v32 = v32
                                                case 32768:
                                                    var v50 uint = 0
                                                    _ = v50
                                                }
                                                if true {
                                                    v51 := GenS4{F0: 0, F1: (v48 * (v48 & (^(v48)))), F2: 1}
                                                    _ = v51
                                                } else {
                                                    v44 += 614
                                                    v52 := GenS5{F0: 2, F1: -52}
                                                    v48 -= v48
                                                    _ = v52
                                                }
                                            }
                                            _ = v43
                                            _ = v44
                                            _ = v48
                                        }
                                        var v53 uint32 = 797
                                        if true {
                                            fmt.Println(fi31)
                                        }
                                        var v54 uint32 = v53
                                        switch "foo bar" {
                                        case "foo bar":
                                            v53 = 1
                                        case "a":
                                            _ = len(v2)
                                        default:
                                            fmt.Println(v32)
                                        }
                                        var v55 float64 = 0.0
                                        _ = v36
                                        _ = v53
                                        _ = v54
                                        _ = v55
                                    }
                                    v56 := GenS4{F0: v32, F1: 42, F2: 822}
                                    _ = v56
                                } else {
                                    switch "hello" {
                                    case "foo bar":
                                        var v57 uint32 = 2147483648
                                        v32 += (v32 << ((42) & 63))
                                        v32 += (^(v32))
                                        _ = v57
                                    }
                                    var v58 bool = false
                                    v7 = 2
                                    _ = v58
                                }
                                var v59 float64 = float64(float32(0.0))
                                _ = v59
                            } else {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                var v60 float32 = (((1.5 + ((float32((1e10 - 3.14159)) * 1e10) / -2.25)) * 1.5) / (1e10 + float32(1.5)))
                                v61 := GenS2{F0: v6, F1: (("a" + ("" + "foo bar")) + (("\n\t" + "hello") + ("" + "a"))), F2: 255}
                                _ = v60
                                _ = v61
                            }
                        }
                        var v62 uint32 = 27
                        v6 -= (v6 | (v6 % (((v6 >> ((42) & 63))) | 1)))
                        _ = fi34
                        _ = v62
                    }
                default:
                    if false {
                        for fi63 := 0; fi63 < 16; fi63++ {
                            var v64 int32 = 2147483647
                            _ = fi63
                            _ = v64
                        }
                    }
                }
                if true {
                    for fi65 := 0; fi65 < 14; fi65++ {
                        v7 = (v7 ^ 127)
                        v7 += v7
                        fmt.Println(v6, fi65, v32)
                        v66 := GenS6{F0: 932, F1: 161, F2: v7}
                        switch (fi65 - (fi65 + (fi65 ^ (fi65 >> ((255) & 63))))) {
                        case 4294967295:
                            if true {
                                var v67 uint = 2
                                var v68 float64 = 3.14159
                                v69 := v4[fi65]
                                _ = v67
                                _ = v68
                                _ = v69
                            } else {
                                if (42 <= 2147483647) {
                                    v32 += 18446744073709551615
                                    v70 := GenS7{F0: (4294967295 < 2)}
                                    if (true || false) {
                                        _ = len(v2)
                                        v6 -= 1
                                        if (((((0.0 - 1e10) - 1.5) + (0.0 + float32(0.1))) / 1e10) >= (3.14159 / ((0.0 * ((3.14159 + 3.14159) * -2.25)) + float32(1e10)))) {
                                            v71 := GenS2{F0: (v6 & v6), F1: "", F2: 2}
                                            v72 := GenS8{F0: 65535, F1: 42}
                                            _ = v71
                                            _ = v72
                                        } else {
                                            var v73 uint32 = 834
                                            _ = v73
                                        }
                                    } else {
                                        v6 -= -616
                                        v74 := GenS8{F0: uint32(v32), F1: 32767}
                                        _ = v74
                                    }
                                    _ = v70
                                }
                                _ = len(v2)
                            }
                        default:
                            v7 += (v7 + v7)
                        }
                        v75 := GenS9{F0: "a", F1: 255, F2: 255}
                        _ = fi65
                        _ = v66
                        _ = v75
                    }
                    v32 = v32
                }
                _ = v32
            }
            var v76 int64 = 498
            _ = fi31
            _ = v76
        }
    default:
        v7 = 136
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
