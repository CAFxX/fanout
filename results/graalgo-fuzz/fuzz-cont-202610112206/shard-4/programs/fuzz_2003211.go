package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 int8; F1 int; F2 uint }
type GenS3 struct { F0 string; F1 uint64 }
type GenS4 struct { F0 uint8; F1 int32; F2 int16 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int8; F1 uint64; F2 string }
type GenS9 struct { F0 int; F1 int32; F2 int64 }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 int16; F1 int }
func genFn0(p0 uint64, p1 uint32) uint64 {
    var v5 int8 = -128
    v6 := GenS1{F0: 0, F1: 1}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return p0
}

func main() {
    var v1 uint16 = 897
    var v2 uint64 = 314
    var v3 float32 = 3.14159
    v4 := make(map[string]uint32)
    var v7 uint = 2
    if (0.1 != (1.5 * (((1.5 * float64(1e10)) + 0.0) + (1e10 * 3.14159)))) {
        fmt.Println(v1)
        var v8 int16 = -768
        switch "hello" {
        case "\n\t":
            switch 42 {
            case -355:
                if true {
                    switch (("hello" + "a") + ("a" + "hello")) {
                    case "":
                        for fi9 := 0; fi9 < 4; fi9++ {
                            var v10 int16 = v8
                            var v11 bool = (!(0 >= 2147483647))
                            var v12 int16 = 32767
                            _ = fi9
                            _ = v10
                            _ = v11
                            _ = v12
                        }
                    default:
                        if (!true) {
                            var v13 int64 = -1
                            v14 := GenS2{F0: 1, F1: 127, F2: 888}
                            _ = v13
                            _ = v14
                        } else {
                            v2 -= (v2 << ((v7) & 63))
                            v2 -= genFn0((v2 & (v2 - genFn0(genFn0(v2, 530), 2147483648))), 2)
                        }
                    }
                    v2 -= (v2 - genFn0(v2, 2147483647))
                    v2 -= (v2 * genFn0(uint64(v1), uint32(v2)))
                }
            default:
                if false {
                    if (("" + "foo bar") > (("\n\t" + "\n\t") + "")) {
                        var v15 float64 = (float64(1.5) * 1e10)
                        _ = v15
                    } else {
                        v16 := []int{}
                        switch "foo bar" {
                        case "hello":
                            for fi17 := 0; fi17 < 20; fi17++ {
                                var v18 int32 = -2147483648
                                switch "foo bar" {
                                case "foo bar":
                                    switch fi17 {
                                    case 2147483647:
                                        var v19 string = ""
                                        _ = v19
                                    case -1:
                                        v20 := GenS2{F0: 100, F1: fi17, F2: (^(v7))}
                                        _ = v20
                                    case -737:
                                        var v21 uint64 = v2
                                        _ = v21
                                    default:
                                        v22 := GenS3{F0: "foo bar", F1: v2}
                                        _ = v22
                                    }
                                case "":
                                    var v23 uint32 = 4294967295
                                    _ = v23
                                case "hello":
                                    for fi24 := 0; fi24 < 12; fi24++ {
                                        _ = len(v16)
                                        var v25 uint8 = 46
                                        var v26 float32 = (v3 - v3)
                                        _ = fi24
                                        _ = v25
                                        _ = v26
                                    }
                                default:
                                    var v27 float64 = (1.5 / (0.1 / float64((v3 - ((float32(3.14159) - 1.5) * -2.25)))))
                                    _ = v27
                                }
                                _ = fi17
                                _ = v18
                            }
                        }
                        v28 := GenS2{F0: 100, F1: 2147483647, F2: 2}
                        _ = v16
                        _ = v28
                    }
                    var v29 uint32 = uint32(v1)
                    _ = v29
                } else {
                    v30 := GenS4{F0: 127, F1: 0, F2: v8}
                    _ = v30
                }
            }
        case "a":
            var v31 uint16 = 42
            _ = v31
        default:
            var v32 int32 = 0
            _ = v32
        }
        _ = v8
    } else {
        var v33 float32 = float32((3.14159 / 1.5))
        _ = v33
    }
    var v34 string = ("foo bar" + "\n\t")
    v35 := GenS5{F0: 42}
    fmt.Println(v1, v3, v7)
    v7 = v7
    var v36 uint16 = 0
    for fi37 := 0; fi37 < 17; fi37++ {
        v38 := GenS6{F0: 107}
        fmt.Println(v7)
        _ = fi37
        _ = v38
    }
    for fi39 := 0; fi39 < 12; fi39++ {
        v3 += v3
        _ = fi39
    }
    var v40 int64 = -9223372036854775808
    fmt.Println(v7)
    _ = len(v4)
    v2 += v2
    switch (v34 + (v34 + v34)) {
    case "hello":
        for fi41 := 0; fi41 < 3; fi41++ {
            for fi42 := 0; fi42 < 15; fi42++ {
                v43 := []int{fi42}
                _ = len(v43)
                _ = fi42
                _ = v43
            }
            v34 = ("a" + (v34 + v34))
            v4[v34] = 0
            _ = fi41
        }
    case "":
        delete(v4, ((v34 + v34) + (v34 + ("\n\t" + "foo bar"))))
    case "a":
        v2 -= (^(v2))
    default:
        if false {
            v44 := []int{}
            for fi45 := 0; fi45 < 16; fi45++ {
                var v46 bool = (fi45 < fi45)
                _ = fi45
                _ = v46
            }
            if true {
                if true {
                    switch -960 {
                    case 1:
                        var v47 string = "hello"
                        _ = v47
                    case -2147483648:
                        for fi48 := 0; fi48 < 14; fi48++ {
                            var v49 uint8 = 255
                            v2 = genFn0((v2 & (v2 | v2)), 244)
                            _ = fi48
                            _ = v49
                        }
                    case 9223372036854775807:
                        if len(v44) > 0 {
                            v44[0] = -781
                        }
                    default:
                        if (1 == 255) {
                            var v50 uint = v7
                            _ = v50
                        }
                    }
                } else {
                    if (1 <= 0) {
                        delete(v4, v34)
                        var v51 uint32 = 343
                        _ = v51
                    }
                }
                if (("" >= "a") || (!((v40 << (((v7 ^ v7)) & 63)) <= 42))) {
                    for fi52 := 0; fi52 < 14; fi52++ {
                        v34 = v34
                        _ = len(v44)
                        _ = fi52
                    }
                    if true {
                        for fi53 := 0; fi53 < 12; fi53++ {
                            switch v34 {
                            case "":
                                var v54 int16 = 1000
                                _ = v54
                            case "\n\t":
                                _ = len(v4)
                            case "hello":
                                var v55 float64 = float64((v3 + v3))
                                _ = v55
                            default:
                                for fi56 := 0; fi56 < 16; fi56++ {
                                    fmt.Println(v7, v34)
                                    v57 := GenS4{F0: uint8(v2), F1: 1, F2: 42}
                                    v58 := GenS7{F0: 18446744073709551615}
                                    _ = fi56
                                    _ = v57
                                    _ = v58
                                }
                            }
                            _ = fi53
                        }
                    } else {
                        v44 = append(v44, 32768)
                    }
                    switch -9223372036854775808 {
                    case -1:
                        if len(v44) > 0 {
                            v44[0] = -9223372036854775808
                        }
                    }
                } else {
                    v59 := GenS8{F0: -117, F1: 1, F2: (v34 + v34)}
                    v7 = v7
                    _ = len(v44)
                    _ = v59
                }
            }
            _ = v44
        } else {
            var v60 string = ("" + ("" + ""))
            var v61 uint16 = (^(v36))
            _ = v60
            _ = v61
        }
    }
    v3 = v3
    switch -51 {
    case -1:
        for fi62 := 0; fi62 < 11; fi62++ {
            switch fi62 {
            case 32768:
                for fi63 := 0; fi63 < 11; fi63++ {
                    var v64 string = ((v34 + v34) + v34)
                    fmt.Println(v34, v64)
                    switch "foo bar" {
                    case "foo bar":
                        if true {
                            if false {
                                v65 := GenS7{F0: v7}
                                v34 = (v64 + "\n\t")
                                _ = v65
                            } else {
                                v66 := GenS9{F0: (fi62 + (fi63 >> ((v7) & 63))), F1: -1, F2: v40}
                                v67 := GenS9{F0: fi63, F1: 507, F2: -1}
                                _ = len(v4)
                                v40 = v40
                                fmt.Println(v3, v36)
                                var v68 uint8 = 0
                                _ = v66
                                _ = v67
                                _ = v68
                            }
                            var v69 float64 = 1e10
                            if true {
                                v70 := []int{(fi62 | (fi62 << ((v7) & 63))), (fi62 % (((fi63 - 128)) | 1)), (fi63 + (^(fi63))), fi62}
                                _ = v70
                            } else {
                                var v71 int32 = -2147483648
                                var v72 float32 = (-2.25 - v3)
                                if ((-(fi62)) >= 2147483647) {
                                    v64 = v34
                                } else {
                                    var v73 int32 = (v71 / ((v71) | 1))
                                    _ = v73
                                }
                                _ = v71
                                _ = v72
                            }
                            _ = v69
                        } else {
                            v3 += float32(0.1)
                            if ((v40 >> ((v7) & 63)) <= (v40 - v40)) {
                                if (((uint64(v1) < uint64(v7)) && (v7 > (v7 / ((v7) | 1)))) && (true || (true || false))) {
                                    v7 -= 346
                                    v4[v34] = 42
                                }
                            }
                        }
                        var v74 uint = 42
                        _ = v74
                    case "\n\t":
                        if (!(4294967295 <= 4294967295)) {
                            v75 := GenS10{F0: v1}
                            switch ((v64 + v64) + "hello") {
                            case "foo bar":
                                fmt.Println(v36)
                            case "":
                                var v76 string = (v64 + (v34 + v64))
                                _ = v76
                            case "\n\t":
                                if true {
                                    v77 := []int{}
                                    switch "a" {
                                    case "a":
                                        v77 = append(v77, -416)
                                    }
                                    _ = v77
                                } else {
                                    v3 = float32((3.14159 / ((float64(((0.1 * float32(1.5)) + v3)) - float64(-2.25)) * 1.5)))
                                }
                            default:
                                var v78 uint = (^(v7))
                                _ = v78
                            }
                            _ = v75
                        } else {
                            v79 := []int{fi63}
                            var v80 uint16 = v36
                            var v81 uint8 = 127
                            var v82 int = -686
                            _ = v79
                            _ = v80
                            _ = v81
                            _ = v82
                        }
                        switch int(v40) {
                        case -2147483648:
                            v64 = v64
                        case 32767:
                            if true {
                                v34 = "foo bar"
                                if true {
                                    if (0.0 >= 0.1) {
                                        v83 := []int{(fi63 % ((-194) | 1))}
                                        v84 := GenS6{F0: 255}
                                        v83 = append(v83, (-(fi63)))
                                        _ = v83
                                        _ = v84
                                    }
                                    var v85 float32 = 3.14159
                                    _ = v85
                                } else {
                                    var v86 uint16 = v1
                                    if (v40 <= (v40 | int64(v7))) {
                                        if true {
                                            v87 := GenS9{F0: fi62, F1: -2147483648, F2: (-(v40))}
                                            delete(v4, (v34 + v34))
                                            _ = v87
                                        } else {
                                            v88 := GenS11{F0: int16(fi63), F1: (fi63 >> ((uint(v36)) & 63))}
                                            var v89 uint32 = 0
                                            _ = v88
                                            _ = v89
                                        }
                                    }
                                    _ = len(v4)
                                    _ = v86
                                }
                                v4[v64] = 65535
                            }
                        default:
                            fmt.Println(v36, v3, v3)
                        }
                        v90 := []int{(fi62 - -393), (fi63 >> ((527) & 63))}
                        _ = v90
                    case "a":
                        if true {
                            var v91 int = -182
                            fmt.Println(v64)
                            _ = v91
                        } else {
                            v92 := GenS6{F0: 10}
                            switch fi63 {
                            case 9223372036854775807:
                                fmt.Println(v36, fi63)
                            case -69:
                                fmt.Println(v2, v64)
                            default:
                                v93 := GenS2{F0: 1, F1: (fi62 ^ 0), F2: v7}
                                _ = v93
                            }
                            _ = v92
                        }
                    }
                    _ = fi63
                    _ = v64
                }
            default:
                var v94 uint32 = 2
                _ = v94
            }
            v95 := []int{fi62, fi62, -942}
            _ = fi62
            _ = v95
        }
    }
    var v96 float32 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v7:", v7)
    fmt.Println("v96:", v96)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v40
    _ = v7
    _ = v96
}
