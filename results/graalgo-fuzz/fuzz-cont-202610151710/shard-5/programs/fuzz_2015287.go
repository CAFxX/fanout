package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint; F1 uint64; F2 int64 }
type GenS5 struct { F0 uint16; F1 uint32; F2 int16 }
type GenS6 struct { F0 uint32; F1 uint; F2 uint8 }
type GenS7 struct { F0 uint32; F1 int; F2 string }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 int64; F1 uint8 }
type GenS10 struct { F0 int32; F1 uint; F2 int64 }
type GenS11 struct { F0 string; F1 uint16 }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 string; F1 string }
type GenS14 struct { F0 int64; F1 uint16 }
type GenS15 struct { F0 uint16; F1 uint8 }
type GenS16 struct { F0 int16; F1 int32; F2 bool }
func genFn0(p0 int32) uint8 {
    v6 := GenS1{F0: (-(p0)), F1: -1}
    p0 = 2147483647
    var v7 uint8 = 128
    _ = p0
    _ = v6
    _ = v7
    return v7
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]bool)
    v3 := make(map[string]uint32)
    var v4 uint32 = 42
    v5 := make(map[string]int)
    v8 := []int{0, -1, 32768, 127}
    if (403 > 42) {
        var v9 uint32 = (^(v4))
        _ = v9
    } else {
        for fi10 := 0; fi10 < 16; fi10++ {
            v4 -= 2
            _ = fi10
        }
    }
    v4 = v4
    if (127 < 9223372036854775807) {
        fmt.Println(v4)
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        fmt.Println(v4)
    } else {
        v4 -= v4
        v4 += v4
    }
    v11 := GenS2{F0: 32767, F1: (((-2147483648 > int32(v4)) || (v4 == (^(v4)))) || true)}
    for fi12 := 0; fi12 < 11; fi12++ {
        switch "foo bar" {
        case "a":
            var v13 int16 = -728
            _ = v13
        case "hello":
            if (40 >= 133) {
                fmt.Println(fi12, v4)
                v14 := GenS1{F0: -2147483648, F1: 1}
                _ = v14
            } else {
                if ((^(fi12)) > (-(fi12))) {
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                    switch "hello" {
                    case "hello":
                        fmt.Println(fi12, v4)
                    case "foo bar":
                        if (((!(((v4 + v4) < (v4 | (v4 >> ((65535) & 63)))) && false)) || true) && false) {
                            v4 = (^(v4))
                            if (1e10 > (1e10 * -2.25)) {
                                v8 = append(v8, fi12)
                                for fi15 := 0; fi15 < 20; fi15++ {
                                    fmt.Println(fi12, fi15, fi15)
                                    v16 := GenS3{F0: 0}
                                    _ = fi15
                                    _ = v16
                                }
                            }
                            if false {
                                if (-1 >= -128) {
                                    v4 = v4
                                    var v17 int32 = -696
                                    _ = v17
                                } else {
                                    var v18 bool = (327 <= 45)
                                    v19 := GenS2{F0: 65535, F1: false}
                                    _ = v18
                                    _ = v19
                                }
                                v4 = v4
                                if (1.5 != -2.25) {
                                    for fi20 := 0; fi20 < 17; fi20++ {
                                        if false {
                                            var v21 float64 = ((0.0 + float64(3.14159)) + -2.25)
                                            v4 = 0
                                            if (-9223372036854775808 == (fi12 * (^(fi12)))) {
                                                v21 -= (1.5 / 1.0)
                                                var v22 int8 = -92
                                                v23 := v1[(v4 >> ((582) & 63))]
                                                _ = v22
                                                _ = v23
                                            }
                                            _ = v21
                                        }
                                        if false {
                                            _ = len(v3)
                                            v24 := GenS4{F0: 420, F1: 638, F2: -942}
                                            v4 = (v4 * v4)
                                            _ = v24
                                        } else {
                                            var v25 float32 = float32(3.14159)
                                            _ = v25
                                        }
                                        v4 -= 0
                                        _ = fi20
                                    }
                                    if (true && false) {
                                        v26 := GenS5{F0: 32767, F1: v4, F2: int16(v4)}
                                        delete(v2, v4)
                                        v4 = 42
                                        _ = v26
                                    } else {
                                        if (1 != 441) {
                                            v4 -= 1
                                        }
                                        if (42 >= 0) {
                                            var v27 float64 = 1.5
                                            fmt.Println(v4, v4)
                                            _ = v27
                                        } else {
                                            var v28 int = fi12
                                            if len(v8) >= 2 {
                                                v8 = v8[1:len(v8)]
                                            }
                                            if (66 == 4294967295) {
                                                if false {
                                                    for fi29 := 0; fi29 < 3; fi29++ {
                                                        var v30 int64 = 1
                                                        _ = fi29
                                                        _ = v30
                                                    }
                                                    fmt.Println(fi12, v28, fi12)
                                                    v4 = 903
                                                } else {
                                                    v28 = 2147483647
                                                    var v31 float64 = -2.25
                                                    v32 := GenS3{F0: 1}
                                                    _ = v31
                                                    _ = v32
                                                }
                                                fmt.Println(v4, v4)
                                            }
                                            _ = v28
                                        }
                                        v8 = append(v8, (fi12 ^ (-(fi12))))
                                    }
                                }
                            }
                        }
                    default:
                        if (false || (!(1 < -32768))) {
                            v4 -= v4
                        } else {
                            v4 = (v4 & v4)
                            if ((948 > 924) && false) {
                                if true {
                                    for fi33 := 0; fi33 < 3; fi33++ {
                                        switch 670 {
                                        case 128:
                                            var v34 string = "\n\t"
                                            _ = v34
                                        case 9223372036854775807:
                                            switch fi12 {
                                            case 32768:
                                                switch "\n\t" {
                                                case "foo bar":
                                                    delete(v1, v4)
                                                case "a":
                                                    v35 := v3["foo bar"]
                                                    _ = v35
                                                default:
                                                    v36 := GenS1{F0: 1, F1: -987}
                                                    _ = v36
                                                }
                                            case 543:
                                                var v37 bool = false
                                                _ = v37
                                            }
                                        case 2147483647:
                                            var v38 uint64 = 560
                                            _ = v38
                                        }
                                        v4 = (v4 / ((65535) | 1))
                                        _ = fi33
                                    }
                                    var v39 uint = 65535
                                    v40 := GenS6{F0: 4294967295, F1: uint(v4), F2: 2}
                                    _ = v39
                                    _ = v40
                                } else {
                                    _ = len(v2)
                                }
                                switch ((("a" + ("" + "")) + "a") + "foo bar") {
                                case "":
                                    if len(v8) >= 2 {
                                        v8 = v8[1:len(v8)]
                                    }
                                case "hello":
                                    var v41 uint64 = 2
                                    _ = v41
                                case "a":
                                    if ((2 != 379) || false) {
                                        var v42 uint = 510
                                        _ = v42
                                    } else {
                                        var v43 int8 = 1
                                        v44 := GenS7{F0: (^(v4)), F1: 128, F2: "a"}
                                        _ = v43
                                        _ = v44
                                    }
                                default:
                                    for fi45 := 0; fi45 < 16; fi45++ {
                                        switch (fi12 / ((-2147483648) | 1)) {
                                        case 1:
                                            if (true || (true && (32767 <= 1))) {
                                                v4 = (v4 / (((v4 ^ (^(v4)))) | 1))
                                            } else {
                                                v46 := GenS4{F0: 586, F1: 42, F2: 512}
                                                _ = v46
                                            }
                                        }
                                        if (0.1 < 0.1) {
                                            v47 := GenS8{F0: 127}
                                            v48 := GenS2{F0: 1, F1: (4294967295 == 0)}
                                            v4 += v4
                                            _ = v47
                                            _ = v48
                                        } else {
                                            if len(v8) >= 2 {
                                                v8 = v8[1:len(v8)]
                                            }
                                        }
                                        _ = fi45
                                    }
                                }
                                switch (fi12 + fi12) {
                                case 638:
                                    var v49 int64 = 9223372036854775807
                                    _ = v49
                                default:
                                    var v50 int = (^(fi12))
                                    _ = v50
                                }
                            }
                            var v51 uint64 = 108
                            _ = v51
                        }
                    }
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                } else {
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                    var v52 int = (fi12 * (-(fi12)))
                    _ = v52
                }
                v53 := GenS7{F0: v4, F1: -9223372036854775808, F2: "a"}
                _ = v53
            }
        case "":
            var v54 int16 = 1
            _ = v54
        }
        for fi55 := 0; fi55 < 8; fi55++ {
            v56 := GenS6{F0: v4, F1: 0, F2: 225}
            if ((0 == -941) || false) {
                fmt.Println(fi12, fi12)
            }
            _ = fi55
            _ = v56
        }
        _ = fi12
    }
    v57 := GenS9{F0: 9223372036854775807, F1: genFn0(-962)}
    for fi58 := 0; fi58 < 12; fi58++ {
        v59 := GenS10{F0: -876, F1: 2, F2: 9223372036854775807}
        _ = fi58
        _ = v59
    }
    var v60 uint32 = 998
    var v61 string = (("hello" + "hello") + "foo bar")
    switch v61 {
    case "hello":
        for fi62 := 0; fi62 < 6; fi62++ {
            if (!false) {
                if true {
                    v61 = v61
                } else {
                    if (-2.25 == float32(0.1)) {
                        var v63 uint16 = 2
                        for fi64 := 0; fi64 < 14; fi64++ {
                            v8 = append(v8, fi64)
                            v61 = "\n\t"
                            if ((18446744073709551615 < 1) && true) {
                                var v65 bool = false
                                v60 += v4
                                _ = v65
                            }
                            v4 = 365
                            _ = fi64
                        }
                        _ = v63
                    }
                    v60 -= 1
                }
                switch (fi62 * -342) {
                case 1:
                    for fi66 := 0; fi66 < 10; fi66++ {
                        switch v61 {
                        case "hello":
                            v61 = (((v61 + v61) + v61) + v61)
                            var v67 string = v61
                            var v68 int = fi66
                            _ = v67
                            _ = v68
                        case "foo bar":
                            v60 = v60
                            var v69 uint8 = 42
                            var v70 uint = 18446744073709551615
                            switch "" {
                            case "foo bar":
                                v71 := v1[(v60 - 2147483647)]
                                _ = v71
                            case "":
                                v72 := GenS3{F0: 42}
                                _ = v72
                            default:
                                if true {
                                    v61 = v61
                                }
                                var v73 uint = v70
                                if false {
                                    v74 := GenS6{F0: v60, F1: v73, F2: 31}
                                    if (false && ((v60 & 544) == 65535)) {
                                        if (v61 == "foo bar") {
                                            if len(v8) > 0 {
                                                v8[0] = 4294967295
                                            }
                                            fmt.Println(v73, v4, fi62)
                                        } else {
                                            fmt.Println(v73, v60)
                                            if false {
                                                v75 := GenS11{F0: v61, F1: 222}
                                                v76 := GenS12{F0: 2147483647}
                                                _ = v75
                                                _ = v76
                                            }
                                            if (42 > -710) {
                                                var v77 uint64 = 2
                                                v73 -= (^(v70))
                                                fmt.Println(fi66, v69)
                                                _ = v77
                                            }
                                        }
                                        if false {
                                            var v78 int = 32767
                                            if ((((1.5 / 3.14159) - 0.0) / 1e10) <= ((1.5 + (((0.1 + -2.25) + 3.14159) * 3.14159)) * (3.14159 / 1e10))) {
                                                var v79 uint = (^(v70))
                                                fmt.Println(fi62, v69, v60)
                                                if false {
                                                    v80 := GenS5{F0: 965, F1: v60, F2: -851}
                                                    if false {
                                                        v81 := GenS13{F0: (("" + v61) + v61), F1: ("a" + v61)}
                                                        _ = v81
                                                    } else {
                                                        switch 2147483647 {
                                                        case 59:
                                                            var v82 uint8 = uint8(v70)
                                                            _ = v82
                                                        case -630:
                                                            var v83 uint = v79
                                                            _ = v83
                                                        case -837:
                                                            if (!false) {
                                                                if true {
                                                                    if (65535 <= (v60 ^ 2)) {
                                                                        var v84 int32 = 267
                                                                        var v85 uint8 = (^(v69))
                                                                        _ = len(v8)
                                                                        _ = v84
                                                                        _ = v85
                                                                    }
                                                                } else {
                                                                    v79 += v70
                                                                    fmt.Println(v60)
                                                                }
                                                                v69 += genFn0(0)
                                                                fmt.Println(v79)
                                                            }
                                                        default:
                                                            v4 += v4
                                                        }
                                                    }
                                                    delete(v3, (("" + ("" + v61)) + "\n\t"))
                                                    _ = v80
                                                }
                                                switch -9223372036854775808 {
                                                case 32767:
                                                    fmt.Println(v61, v69, fi62)
                                                case -882:
                                                    if len(v8) >= 2 {
                                                        v8 = v8[1:len(v8)]
                                                    }
                                                default:
                                                    v86 := v3["hello"]
                                                    _ = v86
                                                }
                                                _ = v79
                                            } else {
                                                var v87 int64 = 538
                                                v70 = 4294967295
                                                _ = v87
                                            }
                                            var v88 uint8 = uint8(v70)
                                            _ = v78
                                            _ = v88
                                        }
                                        if true {
                                            v89 := GenS14{F0: 1, F1: 307}
                                            _ = v89
                                        } else {
                                            delete(v1, v4)
                                        }
                                    }
                                    if (-2.25 != -2.25) {
                                        v90 := GenS15{F0: 0, F1: (^(v69))}
                                        _ = len(v1)
                                        switch v61 {
                                        case "hello":
                                            var v91 float64 = 0.1
                                            _ = v91
                                        }
                                        _ = v90
                                    } else {
                                        var v92 int8 = 100
                                        _ = v92
                                    }
                                    _ = v74
                                } else {
                                    v93 := v5[(v61 + (("hello" + v61) + "hello"))]
                                    var v94 int16 = int16(v73)
                                    if false {
                                        v73 = v70
                                        switch -1 {
                                        case -2147483648:
                                            v69 = genFn0(42)
                                        case 127:
                                            v95 := v1[(v4 >> ((583) & 63))]
                                            _ = v95
                                        default:
                                            switch 4294967295 {
                                            case -1:
                                                var v96 int8 = -31
                                                _ = v96
                                            default:
                                                delete(v3, (v61 + "\n\t"))
                                            }
                                        }
                                    }
                                    _ = v93
                                    _ = v94
                                }
                                _ = v73
                            }
                            _ = v69
                            _ = v70
                        default:
                            fmt.Println(v60, v61)
                        }
                        var v97 int32 = 0
                        _ = fi66
                        _ = v97
                    }
                case 0:
                    switch "hello" {
                    case "a":
                        var v98 int16 = 1000
                        _ = v98
                    default:
                        for fi99 := 0; fi99 < 2; fi99++ {
                            var v100 bool = (1 >= 733)
                            _ = fi99
                            _ = v100
                        }
                    }
                }
                var v101 int32 = -206
                _ = v101
            } else {
                if (false && (fi62 >= fi62)) {
                    delete(v5, ((v61 + "\n\t") + (v61 + ((v61 + v61) + v61))))
                    v8 = append(v8, fi62)
                } else {
                    delete(v2, 4294967295)
                }
                fmt.Println(fi62, v60, v61)
            }
            v61 = ("" + "")
            _ = fi62
        }
    case "a":
        v61 = v61
    }
    v102 := GenS4{F0: 758, F1: 1, F2: -1}
    v61 = (v61 + (((v61 + v61) + v61) + ("a" + v61)))
    for fi103 := 0; fi103 < 7; fi103++ {
        var v104 bool = true
        var v105 uint = 0
        _ = fi103
        _ = v104
        _ = v105
    }
    switch 42 {
    case -855:
        v60 = v4
    case 4294967295:
        var v106 uint64 = 42
        _ = v106
    case -803:
        switch ("" + v61) {
        case "foo bar":
            var v107 bool = (false || (!true))
            _ = v107
        }
    }
    v61 = (((v61 + (v61 + ("" + "\n\t"))) + "") + "a")
    v108 := GenS16{F0: -955, F1: -2147483648, F2: true}
    v109 := GenS14{F0: 215, F1: 32767}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v102
    _ = v108
    _ = v109
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v60
    _ = v61
    _ = v8
}
