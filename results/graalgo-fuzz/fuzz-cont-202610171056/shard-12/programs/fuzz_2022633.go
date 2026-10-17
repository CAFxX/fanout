package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 uint32; F1 int16; F2 int64 }
type GenS3 struct { F0 string; F1 uint32; F2 bool }
type GenS4 struct { F0 bool; F1 int }
type GenS5 struct { F0 int; F1 int8; F2 uint64 }
type GenS6 struct { F0 uint64; F1 bool; F2 uint64 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 string; F1 uint8 }
type GenS9 struct { F0 int16; F1 int; F2 uint64 }

func main() {
    var v1 int = -20
    v2 := []string{}
    var v3 uint16 = 101
    v3 += v3
    for fi4 := 0; fi4 < 15; fi4++ {
        if len(v2) > 0 {
            v2[0] = "hello"
        }
        fmt.Println(v3, v3)
        _ = fi4
    }
    _ = len(v2)
    var v5 uint64 = 748
    if true {
        if false {
            fmt.Println(v5, v3, v5)
            for fi6 := 0; fi6 < 15; fi6++ {
                if (-15 != 115) {
                    switch v1 {
                    case -846:
                        for fi7 := 0; fi7 < 11; fi7++ {
                            var v8 float32 = float32(0.1)
                            if (15 != 42) {
                                v3 = 32767
                                var v9 uint8 = 42
                                v10 := GenS1{F0: (fi6 / (((fi7 >> ((uint(v5)) & 63))) | 1)), F1: 252}
                                _ = v10
                                _ = v9
                            }
                            _ = fi7
                            _ = v8
                        }
                    case -2147483648:
                        v11 := GenS2{F0: 307, F1: 0, F2: -25}
                        _ = v11
                    case 899:
                        v12 := GenS1{F0: fi6, F1: 128}
                        _ = v12
                    }
                    var v13 uint = 4294967295
                    if false {
                        v13 -= (v13 % ((v13) | 1))
                    }
                    _ = v13
                } else {
                    if (fi6 < v1) {
                        if len(v2) > 0 {
                            v2[0] = "foo bar"
                        }
                    }
                }
                v14 := GenS3{F0: "a", F1: 4294967295, F2: false}
                _ = fi6
                _ = v14
            }
            var v15 uint32 = 65535
            _ = v15
        }
        var v16 int = (^(v1))
        _ = v16
    }
    v17 := GenS2{F0: 2, F1: 1000, F2: int64(v1)}
    v18 := GenS4{F0: true, F1: 38}
    v1 += -1
    v5 -= v5
    var v19 uint8 = 129
    var v20 int64 = int64(v5)
    var v21 int64 = v20
    if false {
        v3 = (v3 - (v3 % (((v3 / ((uint16(v19)) | 1))) | 1)))
        for fi22 := 0; fi22 < 19; fi22++ {
            v23 := GenS2{F0: 4294967295, F1: int16(v19), F2: 960}
            switch v1 {
            case -2147483648:
                if false {
                    if true {
                        if false {
                            v24 := GenS1{F0: v1, F1: (v19 & uint8(v3))}
                            var v25 uint8 = v19
                            for fi26 := 0; fi26 < 2; fi26++ {
                                if (true || (42 <= v1)) {
                                    v2 = append(v2, "hello")
                                    if ("hello" < "hello") {
                                        switch "foo bar" {
                                        case "":
                                            if len(v2) > 0 {
                                                v2[0] = "hello"
                                            }
                                        case "hello":
                                            v27 := GenS3{F0: "", F1: 42, F2: (1.5 >= 0.1)}
                                            _ = v27
                                        case "foo bar":
                                            if len(v2) > 0 {
                                                v2[0] = "\n\t"
                                            }
                                        default:
                                            var v28 float64 = 0.1
                                            if true {
                                                v29 := GenS4{F0: true, F1: (fi26 >> ((18446744073709551615) & 63))}
                                                var v30 uint16 = (^(v3))
                                                _ = v29
                                                _ = v30
                                            } else {
                                                v25 -= v19
                                                switch fi22 {
                                                case 42:
                                                    if false {
                                                        var v31 uint8 = v19
                                                        _ = v31
                                                    }
                                                case 457:
                                                    var v32 float64 = v28
                                                    _ = v32
                                                case 32767:
                                                    var v33 uint8 = 42
                                                    _ = v33
                                                }
                                            }
                                            _ = v28
                                        }
                                        var v34 string = ("hello" + "foo bar")
                                        var v35 string = v34
                                        _ = v34
                                        _ = v35
                                    }
                                    v2 = append(v2, ("\n\t" + ("hello" + "a")))
                                } else {
                                    var v36 int8 = -1
                                    v20 += (^(v20))
                                    if false {
                                        if ((!(-2147483648 <= 0)) || true) {
                                            v25 += 0
                                            _ = len(v2)
                                            if ("hello" > "\n\t") {
                                                if (!false) {
                                                    if (!true) {
                                                        var v37 uint64 = (v5 / (((v5 % ((v5) | 1))) | 1))
                                                        var v38 uint = 0
                                                        _ = v37
                                                        _ = v38
                                                    } else {
                                                        if (0 <= v5) {
                                                            v21 -= (v21 - (v21 & (^(v20))))
                                                            if len(v2) > 0 {
                                                                v2[0] = "foo bar"
                                                            }
                                                            v39 := GenS5{F0: v1, F1: v36, F2: (v5 ^ v5)}
                                                            _ = v39
                                                        } else {
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                            _ = len(v2)
                                                            var v40 int32 = int32(fi22)
                                                            v2 = append(v2, "hello")
                                                            _ = v40
                                                        }
                                                        var v41 int64 = (v20 | int64(v36))
                                                        v21 += 614
                                                        _ = v41
                                                    }
                                                    v20 += 42
                                                    if (42 != v36) {
                                                        v25 = v19
                                                        fmt.Println(v21, v3, v20)
                                                    } else {
                                                        if (v19 > (v19 * 217)) {
                                                            _ = len(v2)
                                                        }
                                                        var v42 uint32 = uint32(v3)
                                                        var v43 bool = false
                                                        _ = v42
                                                        _ = v43
                                                    }
                                                } else {
                                                    fmt.Println(v1, fi22)
                                                }
                                                v19 += (v19 / (((v19 >> ((295) & 63))) | 1))
                                                v25 += 0
                                            } else {
                                                var v44 int8 = (^(v36))
                                                _ = v44
                                            }
                                        }
                                        v20 = 0
                                        var v45 int64 = (v20 / (((v21 * (v20 << ((255) & 63)))) | 1))
                                        _ = v45
                                    } else {
                                        var v46 float64 = 3.14159
                                        var v47 int64 = (v20 >> ((837) & 63))
                                        v2 = append(v2, "\n\t")
                                        _ = v46
                                        _ = v47
                                    }
                                    v48 := GenS5{F0: (v1 / (((v1 * fi22)) | 1)), F1: v36, F2: v5}
                                    _ = v36
                                    _ = v48
                                }
                                if ((-(fi26)) < v1) {
                                    v49 := GenS4{F0: false, F1: 1}
                                    fmt.Println(v20, v3)
                                    fmt.Println(v1, v19)
                                    _ = v49
                                }
                                _ = fi26
                            }
                            _ = v24
                            _ = v25
                        } else {
                            v21 -= v21
                            var v50 float64 = 0.0
                            _ = v50
                        }
                        var v51 float64 = 1.5
                        v5 += (v5 * 0)
                        _ = v51
                    } else {
                        v52 := GenS6{F0: 374, F1: false, F2: 1}
                        _ = v52
                    }
                } else {
                    v53 := GenS7{F0: 9223372036854775807}
                    _ = v53
                }
            case 1:
                v54 := GenS3{F0: "\n\t", F1: 42, F2: true}
                _ = v54
            case 490:
                if (!(!true)) {
                    var v55 uint64 = v5
                    _ = v55
                }
            default:
                switch "foo bar" {
                case "\n\t":
                    var v56 int8 = 42
                    _ = v56
                }
            }
            _ = fi22
            _ = v23
        }
        switch (("\n\t" + (("\n\t" + "") + "hello")) + "hello") {
        case "":
            if false {
                if (true || ((4294967295 == 4294967295) && (475 == 1))) {
                    _ = len(v2)
                } else {
                    if (!false) {
                        fmt.Println(v5, v3)
                        var v57 string = ("" + ("\n\t" + "foo bar"))
                        var v58 uint32 = 42
                        _ = v57
                        _ = v58
                    }
                    var v59 bool = (v19 <= v19)
                    _ = v59
                }
                v60 := GenS8{F0: "", F1: v19}
                _ = v60
            }
        default:
            v61 := GenS6{F0: v5, F1: true, F2: (v5 | 18446744073709551615)}
            _ = v61
        }
    }
    fmt.Println(v1, v5, v1)
    for fi62 := 0; fi62 < 10; fi62++ {
        v21 = v21
        _ = len(v2)
        _ = fi62
    }
    for fi63 := 0; fi63 < 6; fi63++ {
        fmt.Println(v1, v3, v1)
        _ = fi63
    }
    for fi64 := 0; fi64 < 15; fi64++ {
        for fi65 := 0; fi65 < 18; fi65++ {
            v66 := GenS1{F0: 32768, F1: (v19 >> ((uint(v5)) & 63))}
            switch "a" {
            case "hello":
                v20 += (v21 / (((v21 / ((v20) | 1))) | 1))
            case "a":
                v1 = fi65
                var v67 float64 = float64(3.14159)
                if false {
                    v68 := GenS9{F0: -32768, F1: (v1 | (^(fi65))), F2: v5}
                    var v69 int16 = 1
                    _ = v68
                    _ = v69
                } else {
                    v2 = append(v2, "a")
                }
                _ = v67
            case "\n\t":
                var v70 int64 = 42
                _ = v70
            default:
                v3 += v3
            }
            _ = fi65
            _ = v66
        }
        var v71 int64 = 1
        switch "hello" {
        case "":
            switch (v1 | (fi64 % ((9223372036854775807) | 1))) {
            case 9223372036854775807:
                switch v1 {
                case 393:
                    var v72 string = ("\n\t" + "\n\t")
                    _ = v72
                case 35:
                    fmt.Println(v1, v19, v5)
                case -355:
                    for fi73 := 0; fi73 < 14; fi73++ {
                        fmt.Println(v3, v5, fi73)
                        if false {
                            var v74 int16 = -32768
                            _ = len(v2)
                            _ = v74
                        } else {
                            if false {
                                var v75 int32 = 1
                                var v76 int32 = v75
                                _ = v75
                                _ = v76
                            } else {
                                if (4294967295 > 2) {
                                    fmt.Println(fi64)
                                    var v77 uint = 18446744073709551615
                                    _ = v77
                                } else {
                                    var v78 bool = true
                                    fmt.Println(v20)
                                    if (!(v78 && ((v19 << ((65535) & 63)) > uint8(v5)))) {
                                        v2 = append(v2, "foo bar")
                                        v2 = append(v2, "\n\t")
                                    }
                                    if true {
                                        fmt.Println(v78)
                                    }
                                    _ = v78
                                }
                            }
                        }
                        switch "foo bar" {
                        case "\n\t":
                            if true {
                                v1 += 2147483647
                                v5 += v5
                                v1 = (fi73 + 1)
                            }
                        default:
                            var v79 uint8 = uint8(v3)
                            _ = v79
                        }
                        _ = fi73
                    }
                }
            default:
                if (0 < 2) {
                    for fi80 := 0; fi80 < 11; fi80++ {
                        _ = len(v2)
                        if false {
                            switch "foo bar" {
                            case "\n\t":
                                var v81 string = ""
                                _ = v81
                            case "foo bar":
                                var v82 uint32 = 2147483647
                                _ = v82
                            }
                            fmt.Println(v20)
                        }
                        _ = fi80
                    }
                    for fi83 := 0; fi83 < 8; fi83++ {
                        v20 += (^(v71))
                        _ = fi83
                    }
                    var v84 int32 = 771
                    _ = v84
                }
            }
        default:
            switch "hello" {
            case "":
                fmt.Println(v71, v71, fi64)
            case "\n\t":
                v1 = v1
            }
        }
        _ = fi64
        _ = v71
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v5
}
