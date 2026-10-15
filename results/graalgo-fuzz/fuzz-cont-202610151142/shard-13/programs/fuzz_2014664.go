package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 int64 }
type GenS2 struct { F0 int; F1 uint8; F2 uint64 }
type GenS3 struct { F0 uint8; F1 uint64; F2 uint8 }
type GenS4 struct { F0 uint32; F1 bool; F2 int16 }
type GenS5 struct { F0 uint; F1 int }
type GenS6 struct { F0 uint32; F1 bool }
type GenS7 struct { F0 uint64; F1 int8 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 int8 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]int)
    v3 := []uint32{42, 306}
    var v4 uint8 = 255
    v5 := []int{206, 2147483647, 908}
    var v6 uint64 = 2
    v4 -= (v4 * uint8(v6))
    v4 += (v4 / ((v4) | 1))
    switch "a" {
    case "foo bar":
        v6 -= (v6 - v6)
    case "hello":
        var v7 uint16 = 276
        _ = v7
    default:
        fmt.Println(v6, v6)
    }
    if (1 <= 418) {
        _ = len(v5)
        if (!false) {
            var v8 uint32 = 2
            for fi9 := 0; fi9 < 7; fi9++ {
                var v10 int16 = 1000
                v10 = -1
                _ = fi9
                _ = v10
            }
            _ = v8
        }
        v4 += (v4 << ((65535) & 63))
    } else {
        for fi11 := 0; fi11 < 9; fi11++ {
            v5 = append(v5, (-(fi11)))
            _ = fi11
        }
        if false {
            for fi12 := 0; fi12 < 17; fi12++ {
                v13 := GenS1{F0: -9223372036854775808, F1: 1000, F2: -9223372036854775808}
                v14 := v2[1]
                var v15 int = fi12
                _ = fi12
                _ = v13
                _ = v14
                _ = v15
            }
            var v16 uint32 = 2147483648
            v4 -= 1
            _ = v16
        } else {
            v17 := GenS1{F0: 469, F1: 587, F2: -892}
            var v18 float64 = float64((1.5 - (3.14159 - 3.14159)))
            _ = v17
            _ = v18
        }
        var v19 uint32 = 268
        _ = v19
    }
    if true {
        delete(v2, 65535)
        switch "hello" {
        case "hello":
            v6 -= 42
        default:
            v20 := v1[int(v4)]
            _ = v20
        }
        switch -640 {
        case 650:
            for fi21 := 0; fi21 < 13; fi21++ {
                for fi22 := 0; fi22 < 16; fi22++ {
                    var v23 int32 = 1
                    _ = fi22
                    _ = v23
                }
                _ = fi21
            }
        case 4294967295:
            v24 := v2[42]
            _ = v24
        case 1:
            var v25 int64 = -244
            _ = v25
        default:
            v26 := v2[2147483647]
            _ = v26
        }
    } else {
        switch 127 {
        case 32767:
            v27 := v1[int(v4)]
            _ = v27
        }
        if len(v5) > 0 {
            v5[0] = -2147483648
        }
        var v28 int64 = 698
        _ = v28
    }
    var v29 uint8 = v4
    v29 = (^(v4))
    if len(v5) > 0 {
        v5[0] = 32768
    }
    if true {
        if true {
            v30 := GenS2{F0: 42, F1: (^(v29)), F2: 4294967295}
            switch "a" {
            case "hello":
                fmt.Println(v4, v4)
            case "\n\t":
                if (v29 == (^(v29))) {
                    var v31 bool = false
                    v4 = v29
                    delete(v1, -9223372036854775808)
                    _ = v31
                }
            case "foo bar":
                switch 2147483647 {
                case 42:
                    if (638 != 9223372036854775807) {
                        switch 32767 {
                        case -623:
                            var v32 int8 = 100
                            _ = v32
                        case -104:
                            v4 -= v4
                        case 477:
                            _ = len(v2)
                        default:
                            for fi33 := 0; fi33 < 4; fi33++ {
                                v6 = (v6 * v6)
                                var v34 uint32 = uint32(v6)
                                _ = fi33
                                _ = v34
                            }
                        }
                        var v35 string = ""
                        _ = v35
                    } else {
                        for fi36 := 0; fi36 < 14; fi36++ {
                            switch "hello" {
                            case "hello":
                                for fi37 := 0; fi37 < 1; fi37++ {
                                    var v38 bool = true
                                    v4 = v4
                                    _ = fi37
                                    _ = v38
                                }
                            case "foo bar":
                                if (1 == (v6 | (v6 << ((1) & 63)))) {
                                    fmt.Println(v29, v6)
                                    var v39 int64 = -9223372036854775808
                                    switch "\n\t" {
                                    case "foo bar":
                                        if (!false) {
                                            if true {
                                                v40 := GenS3{F0: (v29 - v29), F1: (^(v6)), F2: v29}
                                                var v41 int64 = v39
                                                for fi42 := 0; fi42 < 9; fi42++ {
                                                    var v43 uint16 = 2
                                                    v44 := GenS3{F0: 127, F1: (v6 / (((v6 ^ (^(v6)))) | 1)), F2: 43}
                                                    switch (fi36 << ((749) & 63)) {
                                                    case 0:
                                                        var v45 uint64 = (^(v6))
                                                        _ = v45
                                                    case 791:
                                                        switch (fi42 >> ((42) & 63)) {
                                                        case 97:
                                                            fmt.Println(v4)
                                                        }
                                                    default:
                                                        var v46 uint8 = uint8(v43)
                                                        _ = v46
                                                    }
                                                    v47 := GenS1{F0: v41, F1: -204, F2: v41}
                                                    v48 := GenS1{F0: (v41 >> ((834) & 63)), F1: 396, F2: v39}
                                                    if true {
                                                        v43 -= 255
                                                    } else {
                                                        var v49 bool = (v6 > (v6 - v6))
                                                        var v50 int8 = -128
                                                        if true {
                                                            var v51 bool = ((fi36 != fi42) && v49)
                                                            if v49 {
                                                                fmt.Println(v6, v39, v6)
                                                            } else {
                                                                if v51 {
                                                                    var v52 string = ("\n\t" + "hello")
                                                                    v53 := GenS4{F0: uint32(v43), F1: false, F2: 1000}
                                                                    var v54 uint64 = v6
                                                                    _ = v52
                                                                    _ = v53
                                                                    _ = v54
                                                                }
                                                                var v55 float32 = -2.25
                                                                if true {
                                                                    switch fi36 {
                                                                    case 42:
                                                                        var v56 uint8 = (v4 << ((65535) & 63))
                                                                        _ = v56
                                                                    case 128:
                                                                        v41 = (^(v41))
                                                                    default:
                                                                        v6 += v6
                                                                    }
                                                                    v57 := GenS3{F0: (v4 ^ (^(v29))), F1: 42, F2: (v4 << ((4294967295) & 63))}
                                                                    var v58 float64 = -2.25
                                                                    _ = v57
                                                                    _ = v58
                                                                } else {
                                                                    var v59 uint64 = (^(v6))
                                                                    _ = v59
                                                                }
                                                                if v49 {
                                                                    v60 := GenS5{F0: 42, F1: 487}
                                                                    var v61 float32 = v55
                                                                    var v62 float32 = v61
                                                                    v51 = false
                                                                    _ = v60
                                                                    _ = v61
                                                                    _ = v62
                                                                }
                                                                _ = v55
                                                            }
                                                            _ = v51
                                                        }
                                                        if (true && (v49 || (((^(v6)) > v6) || (!v49)))) {
                                                            v63 := GenS6{F0: 42, F1: (!false)}
                                                            var v64 uint = 4294967295
                                                            var v65 int16 = 0
                                                            var v66 uint32 = 2
                                                            v67 := GenS2{F0: fi42, F1: v4, F2: 2}
                                                            fmt.Println(v6, v4)
                                                            _ = v63
                                                            _ = v64
                                                            _ = v65
                                                            _ = v66
                                                            _ = v67
                                                        }
                                                        var v68 int8 = v50
                                                        if (false && (v4 < v29)) {
                                                            v4 = 2
                                                        } else {
                                                            if v49 {
                                                                fmt.Println(v49, v43)
                                                                if v49 {
                                                                    var v69 int = 127
                                                                    fmt.Println(v41, v6)
                                                                    v4 = (v29 / ((v4) | 1))
                                                                    var v70 uint64 = (v6 - (^(v6)))
                                                                    _ = v69
                                                                    _ = v70
                                                                } else {
                                                                    var v71 uint8 = 128
                                                                    if v49 {
                                                                        var v72 uint16 = (v43 % ((896) | 1))
                                                                        switch ("foo bar" + "foo bar") {
                                                                        case "a":
                                                                            v73 := GenS4{F0: 1, F1: (v49 || v49), F2: 202}
                                                                            _ = v73
                                                                        case "":
                                                                            switch (fi42 >> ((255) & 63)) {
                                                                            case 32768:
                                                                                if ((-2.25 <= 1e10) || (v49 || (!(-2.25 > (0.1 * (0.0 / 1e10)))))) {
                                                                                    if (v6 > (v6 % ((268) | 1))) {
                                                                                        fmt.Println(v72, fi42)
                                                                                        if len(v5) >= 2 {
                                                                                            v5 = v5[1:len(v5)]
                                                                                        }
                                                                                    }
                                                                                } else {
                                                                                    var v74 float64 = 1e10
                                                                                    switch "" {
                                                                                    case "hello":
                                                                                        v43 += v43
                                                                                    default:
                                                                                        var v75 uint8 = 0
                                                                                        _ = v75
                                                                                    }
                                                                                    var v76 bool = (v49 && false)
                                                                                    _ = v74
                                                                                    _ = v76
                                                                                }
                                                                            case 2147483647:
                                                                                v77 := GenS7{F0: 1, F1: v50}
                                                                                _ = v77
                                                                            }
                                                                        default:
                                                                            var v78 int16 = 1
                                                                            _ = v78
                                                                        }
                                                                        v49 = v49
                                                                        _ = v72
                                                                    }
                                                                    v79 := GenS8{F0: -1}
                                                                    v50 = (v68 & (^(v50)))
                                                                    _ = v71
                                                                    _ = v79
                                                                }
                                                            }
                                                            _ = len(v2)
                                                            switch fi42 {
                                                            case 127:
                                                                fmt.Println(v39)
                                                            }
                                                        }
                                                        var v80 string = ("hello" + "\n\t")
                                                        _ = v49
                                                        _ = v50
                                                        _ = v68
                                                        _ = v80
                                                    }
                                                    v81 := GenS4{F0: 2147483647, F1: true, F2: -32768}
                                                    if false {
                                                        fmt.Println(fi36)
                                                    } else {
                                                        var v82 int64 = (v39 * v41)
                                                        v4 = 1
                                                        switch fi42 {
                                                        case 2147483647:
                                                            v43 += 255
                                                        case 32768:
                                                            v82 += (v41 ^ 42)
                                                        }
                                                        _ = v82
                                                    }
                                                    var v83 int64 = (^(v41))
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                    var v84 int32 = 754
                                                    _ = fi42
                                                    _ = v43
                                                    _ = v44
                                                    _ = v47
                                                    _ = v48
                                                    _ = v81
                                                    _ = v83
                                                    _ = v84
                                                }
                                                _ = v40
                                                _ = v41
                                            }
                                            var v85 int64 = (v39 * v39)
                                            v29 = (^(v29))
                                            _ = v85
                                        }
                                    case "a":
                                        v4 = 0
                                    default:
                                        for fi86 := 0; fi86 < 11; fi86++ {
                                            delete(v2, 501)
                                            var v87 float32 = ((-2.25 + 0.1) + (1e10 * 0.0))
                                            v88 := GenS4{F0: 995, F1: true, F2: 1000}
                                            _ = fi86
                                            _ = v87
                                            _ = v88
                                        }
                                    }
                                    _ = v39
                                }
                            case "a":
                                for fi89 := 0; fi89 < 3; fi89++ {
                                    v90 := GenS9{F0: 127}
                                    switch ("foo bar" + "") {
                                    case "\n\t":
                                        v29 = (v29 << ((42) & 63))
                                    case "":
                                        if ((65535 == 0) && false) {
                                            v6 = v6
                                            v6 -= (^(v6))
                                        }
                                        var v91 int64 = 0
                                        _ = v91
                                    case "foo bar":
                                        if len(v3) > 0 {
                                            v3[0] = 2
                                        }
                                    default:
                                        if (278 < 0) {
                                            v92 := v2[2]
                                            _ = v92
                                        }
                                    }
                                    v1[32767] = 42
                                    _ = fi89
                                    _ = v90
                                }
                            }
                            if (42 > 255) {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                if len(v5) > 0 {
                                    v5[0] = fi36
                                }
                                v4 -= v29
                            }
                            _ = fi36
                        }
                        var v93 int16 = 1
                        switch -1 {
                        case -2147483648:
                            for fi94 := 0; fi94 < 7; fi94++ {
                                v6 = v6
                                _ = fi94
                            }
                        case 634:
                            v93 -= int16(v6)
                        case 0:
                            var v95 int8 = 1
                            _ = v95
                        default:
                            var v96 int = 384
                            _ = v96
                        }
                        _ = v93
                    }
                case 0:
                    v4 = 128
                default:
                    var v97 int = -9223372036854775808
                    _ = v97
                }
            }
            for fi98 := 0; fi98 < 6; fi98++ {
                if true {
                    var v99 float32 = -2.25
                    v99 -= 0.1
                    _ = v99
                }
                var v100 int8 = 42
                _ = fi98
                _ = v100
            }
            _ = v30
        } else {
            var v101 bool = false
            v3 = append(v3, 0)
            _ = v101
        }
    }
    var v102 float32 = 0.0
    var v103 float64 = (0.0 / 1.5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v102:", v102)
    fmt.Println("v103:", v103)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v102
    _ = v103
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
