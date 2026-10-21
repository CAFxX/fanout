package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 bool; F1 int32; F2 uint8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int64; F1 uint16 }
type GenS6 struct { F0 int16; F1 uint8 }
type GenS7 struct { F0 string; F1 int; F2 uint }
func genFn0() uint32 {
    var v6 int8 = -86
    if false {
        for fi7 := 0; fi7 < 1; fi7++ {
            var v8 int32 = 0
            v9 := GenS1{F0: (fi7 * (fi7 | fi7)), F1: -9223372036854775808}
            _ = fi7
            _ = v8
            _ = v9
        }
        if false {
            fmt.Println(v6)
            v6 += (v6 * v6)
            v10 := []int{-2147483648, 746, 42, 2147483647}
            _ = v10
        } else {
            var v11 string = "a"
            _ = v11
        }
        v6 = v6
    } else {
        for fi12 := 0; fi12 < 10; fi12++ {
            v13 := []int{}
            if (true && (0 == 65535)) {
                switch fi12 {
                case 650:
                    var v14 int = -2147483648
                    _ = v14
                case -1:
                    var v15 float64 = 3.14159
                    _ = v15
                }
            } else {
                var v16 uint32 = 782
                _ = v16
            }
            for fi17 := 0; fi17 < 10; fi17++ {
                var v18 float32 = (0.1 - 3.14159)
                _ = fi17
                _ = v18
            }
            _ = fi12
            _ = v13
        }
        v19 := GenS2{F0: 42}
        _ = v19
    }
    var v20 uint16 = 2
    _ = v20
    _ = v6
    return 4294967295
}
func genFn1(p0 uint32, p1 uint64) string {
    for fi21 := 0; fi21 < 10; fi21++ {
        var v22 int8 = -1
        p1 -= p1
        _ = fi21
        _ = v22
    }
    p1 += 261
    var v23 uint = 720
    p0 = (p0 * (p0 | 4294967295))
    v24 := []int{}
    _ = p0
    _ = p1
    _ = v23
    _ = v24
    return "hello"
}

func main() {
    var v1 int = 32767
    v2 := make(map[string]string)
    var v3 int32 = 737
    v4 := []int{(v1 ^ (v1 >> ((255) & 63)))}
    var v5 uint32 = 42
    switch "foo bar" {
    case "hello":
        var v25 uint = 18446744073709551615
        _ = v25
    case "foo bar":
        for fi26 := 0; fi26 < 5; fi26++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            for fi27 := 0; fi27 < 6; fi27++ {
                if (69 == 65535) {
                    v28 := v2[(((("hello" + "hello") + ("hello" + "hello")) + ("hello" + "foo bar")) + "foo bar")]
                    _ = v28
                } else {
                    var v29 int64 = 42
                    if true {
                        var v30 uint64 = 384
                        _ = v30
                    }
                    _ = v29
                }
                _ = fi27
            }
            var v31 uint64 = 447
            _ = fi26
            _ = v31
        }
    case "":
        if (-60 != -128) {
            v1 = v1
            switch v1 {
            case 32767:
                if true {
                    fmt.Println(v3, v1, v5)
                    fmt.Println(v3)
                } else {
                    var v32 int32 = v3
                    _ = v32
                }
            default:
                var v33 uint32 = (v5 & (v5 + 2147483648))
                _ = v33
            }
            for fi34 := 0; fi34 < 1; fi34++ {
                switch "foo bar" {
                case "":
                    v5 -= (v5 << ((255) & 63))
                case "hello":
                    if true {
                        for fi35 := 0; fi35 < 1; fi35++ {
                            if false {
                                v2["a"] = "foo bar"
                                v5 = (v5 ^ v5)
                                if (("hello" + "") > "hello") {
                                    v3 += v3
                                    var v36 int8 = 100
                                    _ = v36
                                }
                            }
                            _ = fi35
                        }
                        v2[((((((("a" + "") + ("hello" + "a")) + "hello") + "foo bar") + "\n\t") + ("a" + "foo bar")) + ("\n\t" + "a"))] = ""
                    }
                default:
                    var v37 int8 = 0
                    _ = v37
                }
                fmt.Println(v5, v5)
                if len(v4) > 0 {
                    v4[0] = -201
                }
                _ = fi34
            }
        } else {
            if (false || false) {
                v38 := v2[(("" + "\n\t") + "hello")]
                _ = v38
            } else {
                v39 := v2[""]
                for fi40 := 0; fi40 < 13; fi40++ {
                    var v41 int32 = -108
                    fmt.Println(v3, v39, v41)
                    _ = fi40
                    _ = v41
                }
                _ = v39
            }
            switch "foo bar" {
            case "foo bar":
                var v42 uint16 = 0
                _ = v42
            }
            fmt.Println(v1)
        }
    }
    var v43 uint8 = 1
    fmt.Println(v5)
    v44 := GenS1{F0: (-(v1)), F1: -9223372036854775808}
    if true {
        switch (v1 - v1) {
        case -2147483648:
            if true {
                v45 := GenS1{F0: v1, F1: -1}
                switch v1 {
                case 604:
                    var v46 uint = 42
                    _ = v46
                default:
                    v47 := GenS3{F0: false, F1: -1, F2: v43}
                    _ = v47
                }
                _ = v45
            }
        case 431:
            if false {
                v2[("" + "\n\t")] = ""
            } else {
                v5 = (v5 ^ genFn0())
                var v48 uint = 4294967295
                for fi49 := 0; fi49 < 3; fi49++ {
                    v50 := GenS2{F0: v5}
                    _ = fi49
                    _ = v50
                }
                _ = v48
            }
        default:
            if false {
                v43 = (^(v43))
                if true {
                    fmt.Println(v43, v43)
                    v5 -= genFn0()
                    if true {
                        var v51 float32 = (0.0 - (1e10 + float32(1.5)))
                        switch "hello" {
                        case "a":
                            var v52 bool = (0.0 == (v51 / 1e10))
                            _ = v52
                        case "hello":
                            var v53 int = v1
                            _ = v53
                        case "":
                            var v54 int64 = 9223372036854775807
                            _ = v54
                        default:
                            for fi55 := 0; fi55 < 8; fi55++ {
                                if true {
                                    var v56 int64 = 1
                                    v57 := GenS3{F0: (!false), F1: -2147483648, F2: v43}
                                    v58 := GenS4{F0: 42}
                                    _ = v56
                                    _ = v57
                                    _ = v58
                                } else {
                                    _ = len(v4)
                                    if (!false) {
                                        switch v1 {
                                        case 1:
                                            for fi59 := 0; fi59 < 20; fi59++ {
                                                var v60 uint = uint(v43)
                                                if (fi59 >= (fi59 * (fi55 ^ int(v60)))) {
                                                    var v61 uint16 = 65535
                                                    var v62 uint64 = 1
                                                    v5 = (v5 + (v5 + (v5 << ((v60) & 63))))
                                                    _ = v61
                                                    _ = v62
                                                } else {
                                                    if (v3 <= (v3 & -667)) {
                                                        v3 = v3
                                                        if len(v4) > 0 {
                                                            v4[0] = fi55
                                                        }
                                                        v63 := GenS5{F0: 42, F1: 756}
                                                        _ = v63
                                                    } else {
                                                        var v64 float64 = ((1.5 * (float64(v51) - float64(v51))) * 0.0)
                                                        var v65 int32 = v3
                                                        switch (v1 >> ((0) & 63)) {
                                                        case 696:
                                                            if ((-1 > -1) || true) {
                                                                if (-1 == fi55) {
                                                                    var v66 uint64 = 42
                                                                    var v67 int64 = 9223372036854775807
                                                                    v64 = 0.1
                                                                    _ = v66
                                                                    _ = v67
                                                                }
                                                                var v68 int8 = 42
                                                                _ = v68
                                                            } else {
                                                                switch 32767 {
                                                                case -706:
                                                                    if (uint64(v60) > 249) {
                                                                        _ = len(v4)
                                                                        if false {
                                                                            fmt.Println(v1, fi59)
                                                                        } else {
                                                                            if true {
                                                                                var v69 int16 = int16(v5)
                                                                                v60 += 390
                                                                                if (false || (v3 <= (v65 * (-(v65))))) {
                                                                                    v1 = (v1 >> (((v60 + (^(v60)))) & 63))
                                                                                    if len(v4) >= 2 {
                                                                                        v4 = v4[1:len(v4)]
                                                                                    }
                                                                                } else {
                                                                                    var v70 int32 = v65
                                                                                    _ = v70
                                                                                }
                                                                                _ = v69
                                                                            }
                                                                        }
                                                                        switch "foo bar" {
                                                                        case "":
                                                                            var v71 bool = true
                                                                            _ = v71
                                                                        case "foo bar":
                                                                            switch (fi55 << (((v60 & uint(v5))) & 63)) {
                                                                            case 32767:
                                                                                v43 += (v43 + v43)
                                                                            }
                                                                        case "hello":
                                                                            v64 = float64(((float32((v64 - v64)) + (1e10 * (v51 * (-2.25 + v51)))) - v51))
                                                                        default:
                                                                            var v72 uint64 = 317
                                                                            _ = v72
                                                                        }
                                                                        var v73 uint8 = 42
                                                                        fmt.Println(fi59)
                                                                        if (v51 == v51) {
                                                                            var v74 int = 0
                                                                            _ = v74
                                                                        } else {
                                                                            v3 = v3
                                                                            var v75 float64 = ((v64 - (3.14159 - float64(v51))) * v64)
                                                                            _ = v75
                                                                        }
                                                                        v2[(((("\n\t" + "a") + ("foo bar" + "foo bar")) + ("\n\t" + "\n\t")) + "")] = ("hello" + "")
                                                                        _ = v73
                                                                    } else {
                                                                        if (false && true) {
                                                                            var v76 uint8 = (^(v43))
                                                                            var v77 uint16 = 963
                                                                            var v78 uint64 = 686
                                                                            _ = v76
                                                                            _ = v77
                                                                            _ = v78
                                                                        }
                                                                        v65 += -1
                                                                        var v79 string = ((((("" + ("" + ("hello" + (("a" + ("\n\t" + ("hello" + ("foo bar" + ("foo bar" + ("foo bar" + "")))))) + ("foo bar" + "foo bar"))))) + ("a" + "foo bar")) + "") + "hello") + "\n\t")
                                                                        _ = v79
                                                                    }
                                                                default:
                                                                    fmt.Println(v60, v3)
                                                                }
                                                                switch "foo bar" {
                                                                case "foo bar":
                                                                    v51 = ((v51 + float32(v64)) - v51)
                                                                default:
                                                                    v80 := v2["a"]
                                                                    _ = v80
                                                                }
                                                            }
                                                        case 4294967295:
                                                            v81 := GenS6{F0: 200, F1: uint8(v5)}
                                                            delete(v2, "\n\t")
                                                            if (((true && ("" == "\n\t")) || true) && false) {
                                                                v3 += -1
                                                                if (42 > -71) {
                                                                    var v82 float32 = (v51 / 0.1)
                                                                    _ = v82
                                                                } else {
                                                                    v83 := GenS2{F0: (v5 % ((v5) | 1))}
                                                                    _ = v83
                                                                }
                                                            } else {
                                                                v51 = float32(v64)
                                                                _ = len(v4)
                                                                var v84 int8 = 1
                                                                _ = v84
                                                            }
                                                            _ = v81
                                                        case -2147483648:
                                                            if (false && (!(!(81 == 1)))) {
                                                                if (((3.14159 - (-2.25 + v51)) - (-2.25 - (v51 / (0.1 - (1.5 - 0.0))))) == (0.1 / (v51 * v51))) {
                                                                    fmt.Println(v43)
                                                                    fmt.Println(v5)
                                                                    v3 = (v65 ^ (-(v65)))
                                                                }
                                                            } else {
                                                                v5 = (v5 / (((v5 << ((18446744073709551615) & 63))) | 1))
                                                            }
                                                        default:
                                                            var v85 float64 = 1.5
                                                            _ = v85
                                                        }
                                                        _ = v64
                                                        _ = v65
                                                    }
                                                    var v86 float64 = 1e10
                                                    _ = v86
                                                }
                                                _ = fi59
                                                _ = v60
                                            }
                                        }
                                        var v87 uint32 = (v5 | v5)
                                        _ = v87
                                    } else {
                                        v88 := GenS7{F0: "hello", F1: 9223372036854775807, F2: 42}
                                        v43 -= (^(v43))
                                        for fi89 := 0; fi89 < 16; fi89++ {
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            _ = fi89
                                        }
                                        _ = v88
                                    }
                                    fmt.Println(v3, v43)
                                }
                                if true {
                                    var v90 int64 = 42
                                    _ = v90
                                } else {
                                    for fi91 := 0; fi91 < 1; fi91++ {
                                        v51 += -2.25
                                        _ = fi91
                                    }
                                    var v92 int = (-(v1))
                                    if ("hello" <= "\n\t") {
                                        v51 = float32(-2.25)
                                    }
                                    _ = v92
                                }
                                if false {
                                    _ = len(v4)
                                } else {
                                    fmt.Println(v5, v5, v3)
                                    v43 -= v43
                                    var v93 uint = 534
                                    _ = v93
                                }
                                _ = fi55
                            }
                        }
                        _ = v51
                    } else {
                        if true {
                            fmt.Println(v3)
                            var v94 float64 = -2.25
                            _ = v94
                        } else {
                            var v95 float64 = 1e10
                            v96 := GenS6{F0: 41, F1: v43}
                            v5 = v5
                            _ = v95
                            _ = v96
                        }
                    }
                }
            }
        }
    } else {
        var v97 bool = false
        var v98 int8 = -115
        _ = v97
        _ = v98
    }
    fmt.Println(v3)
    v3 += 0
    v43 = v43
    if false {
        switch "foo bar" {
        case "a":
            var v99 int64 = -1
            _ = v99
        }
    }
    v100 := GenS7{F0: "", F1: (-(v1)), F2: 1}
    v101 := GenS4{F0: v3}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v100
    _ = v101
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v5
}
