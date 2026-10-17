package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16; F1 int8; F2 int32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 bool; F1 int64 }
type GenS7 struct { F0 uint64; F1 int; F2 uint32 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint8 }

func main() {
    var v1 int32 = 42
    var v2 bool = (-1 <= 709)
    var v3 string = "a"
    v4 := make(map[int]int)
    v5 := make(map[string]int)
    v6 := []int{32767, 224, 2147483647, 9223372036854775807}
    var v7 int = 1
    for fi8 := 0; fi8 < 10; fi8++ {
        for fi9 := 0; fi9 < 18; fi9++ {
            var v10 uint16 = 303
            if (false || (float32(0.1) != 0.0)) {
                if (1 != 42) {
                    v7 = fi9
                    var v11 int8 = 100
                    _ = v11
                }
                if false {
                    v1 += (-(v1))
                    var v12 int8 = 42
                    _ = v12
                } else {
                    if (42 < 1000) {
                        if (v2 || false) {
                            if v2 {
                                switch v3 {
                                case "foo bar":
                                    fmt.Println(v3, v1, v2)
                                }
                                if (false || v2) {
                                    v3 = v3
                                    v13 := GenS1{F0: (v3 + (v3 + "\n\t")), F1: (v2 && v2)}
                                    var v14 float64 = (float64(0.0) * ((float64((0.1 / (float32(1.5) + 0.1))) + float64((3.14159 / (-2.25 + (1e10 + -2.25))))) + 1.5))
                                    _ = v13
                                    _ = v14
                                } else {
                                    if v2 {
                                        v1 = v1
                                        if true {
                                            v5["a"] = (fi8 | (fi8 % (((-(fi8))) | 1)))
                                        } else {
                                            v2 = false
                                        }
                                        var v15 float32 = ((0.1 / float32(1e10)) + (-2.25 * (0.1 * 0.0)))
                                        _ = v15
                                    }
                                }
                                var v16 string = (v3 + "")
                                switch "a" {
                                case "foo bar":
                                    var v17 uint = 177
                                    _ = v17
                                case "\n\t":
                                    var v18 int32 = v1
                                    _ = v18
                                }
                                var v19 uint64 = 0
                                _ = v16
                                _ = v19
                            } else {
                                if v2 {
                                    v20 := GenS1{F0: "foo bar", F1: (!(v10 > (v10 / ((v10) | 1))))}
                                    switch v3 {
                                    case "":
                                        var v21 uint8 = 0
                                        _ = v21
                                    }
                                    switch v3 {
                                    case "hello":
                                        var v22 int = (fi9 * (v7 * (fi8 & -483)))
                                        _ = v22
                                    case "":
                                        v23 := GenS2{F0: 320}
                                        _ = v23
                                    case "a":
                                        _ = len(v6)
                                    default:
                                        if len(v6) > 0 {
                                            v6[0] = (fi9 * (v7 / ((fi9) | 1)))
                                        }
                                    }
                                    _ = v20
                                } else {
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    if v2 {
                                        if true {
                                            var v24 int64 = 1
                                            _ = v24
                                        } else {
                                            fmt.Println(v1, v10, fi9)
                                        }
                                        if (0 <= (v7 | fi8)) {
                                            if v2 {
                                                if ((1000 < -59) || v2) {
                                                    v25 := GenS1{F0: v3, F1: v2}
                                                    if (18 >= 127) {
                                                        if len(v6) > 0 {
                                                            v6[0] = 32768
                                                        }
                                                    }
                                                    var v26 uint16 = (^(v10))
                                                    v26 -= v26
                                                    var v27 int8 = -34
                                                    var v28 int64 = 1
                                                    _ = v25
                                                    _ = v26
                                                    _ = v27
                                                    _ = v28
                                                }
                                                var v29 float32 = (0.0 / float32(1.5))
                                                v30 := GenS2{F0: 642}
                                                _ = v29
                                                _ = v30
                                            }
                                            var v31 bool = ((v3 + (v3 + ((v3 + v3) + "\n\t"))) <= v3)
                                            _ = v31
                                        } else {
                                            v10 = (v10 - v10)
                                            v10 -= 255
                                        }
                                    }
                                    var v32 int64 = -9223372036854775808
                                    var v33 int8 = int8(fi8)
                                    _ = v32
                                    _ = v33
                                }
                                v3 = v3
                                v34 := GenS2{F0: 363}
                                _ = v34
                            }
                            var v35 uint16 = v10
                            _ = v35
                        } else {
                            fmt.Println(fi8, fi8, v2)
                            v7 += fi8
                            var v36 int16 = -1
                            _ = v36
                        }
                        fmt.Println(v2, v7)
                        if len(v6) > 0 {
                            v6[0] = fi8
                        }
                        v37 := GenS2{F0: 0}
                        v38 := v5[v3]
                        switch -1 {
                        case 9223372036854775807:
                            v6 = append(v6, (fi8 & (v7 >> ((2) & 63))))
                        default:
                            var v39 int = int(v1)
                            _ = v39
                        }
                        _ = v37
                        _ = v38
                    } else {
                        if (1 <= 195) {
                            switch 2147483647 {
                            case 1:
                                v7 -= (v7 & 32767)
                            case -391:
                                if (127 <= 2) {
                                    v40 := GenS1{F0: v3, F1: false}
                                    v41 := GenS3{F0: 32767, F1: -128, F2: v1}
                                    _ = v40
                                    _ = v41
                                }
                                var v42 int64 = -1
                                v10 -= (v10 / ((v10) | 1))
                                _ = v42
                            case -2147483648:
                                var v43 int8 = 1
                                _ = v43
                            default:
                                v10 = 2
                            }
                            if false {
                                var v44 int32 = (v1 << ((42) & 63))
                                _ = v44
                            }
                            fmt.Println(fi9)
                            var v45 float64 = (float64((1e10 - -2.25)) - (-2.25 / 1.0))
                            _ = v45
                        }
                        v46 := GenS3{F0: 1, F1: 34, F2: v1}
                        _ = len(v6)
                        _ = v46
                    }
                    if true {
                        v10 -= (v10 << ((1) & 63))
                        switch (v7 | fi8) {
                        case -1:
                            v3 = ""
                        case 32767:
                            v7 -= v7
                        }
                        var v47 int16 = 42
                        _ = v47
                    }
                    fmt.Println(v7, fi9)
                }
            }
            fmt.Println(v2, v1, fi8)
            _ = fi9
            _ = v10
        }
        fmt.Println(v3, fi8)
        _ = fi8
    }
    for fi48 := 0; fi48 < 2; fi48++ {
        for fi49 := 0; fi49 < 8; fi49++ {
            if (v1 <= v1) {
                if (-772 >= -2147483648) {
                    v2 = (255 != 15)
                    v3 = "\n\t"
                    v50 := GenS1{F0: "foo bar", F1: true}
                    _ = v50
                }
                v7 -= v7
                var v51 int16 = -248
                v51 += -1
                var v52 int = 2147483647
                switch ("foo bar" + (("" + ((((("" + "") + "\n\t") + (((v3 + v3) + "\n\t") + "foo bar")) + v3) + v3)) + ((v3 + (v3 + v3)) + "\n\t"))) {
                case "\n\t":
                    v53 := GenS3{F0: 0, F1: 75, F2: (v1 + v1)}
                    _ = v53
                case "foo bar":
                    var v54 bool = true
                    _ = v54
                }
                _ = v51
                _ = v52
            }
            _ = fi49
        }
        if v2 {
            v55 := GenS2{F0: 1}
            _ = len(v6)
            _ = v55
        }
        _ = fi48
    }
    v3 = v3
    v56 := GenS4{F0: -324}
    for fi57 := 0; fi57 < 17; fi57++ {
        for fi58 := 0; fi58 < 13; fi58++ {
            if v2 {
                var v59 uint = 604
                var v60 float64 = 0.1
                switch (-(v7)) {
                case -9223372036854775808:
                    v60 = (v60 + 3.14159)
                case 32768:
                    v3 = (v3 + v3)
                case 2147483647:
                    v61 := v5[(((v3 + v3) + "hello") + ((v3 + v3) + v3))]
                    _ = v61
                default:
                    if len(v6) > 0 {
                        v6[0] = fi57
                    }
                }
                _ = v59
                _ = v60
            } else {
                v62 := GenS5{F0: 127}
                v2 = (!v2)
                _ = v62
            }
            v7 = fi58
            v63 := v4[1]
            _ = fi58
            _ = v63
        }
        if len(v6) > 0 {
            v6[0] = 32767
        }
        _ = fi57
    }
    var v64 uint32 = 341
    v65 := GenS6{F0: false, F1: 42}
    var v66 uint16 = 65535
    if (42 > 42) {
        switch v3 {
        case "a":
            for fi67 := 0; fi67 < 17; fi67++ {
                v68 := GenS7{F0: 0, F1: v7, F2: v64}
                _ = fi67
                _ = v68
            }
        case "":
            var v69 int8 = 1
            _ = v69
        }
        fmt.Println(v64)
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    } else {
        switch (-(v7)) {
        case 0:
            v5[""] = v7
        case -751:
            switch "" {
            case "":
                for fi70 := 0; fi70 < 7; fi70++ {
                    for fi71 := 0; fi71 < 11; fi71++ {
                        if (fi70 <= (fi71 ^ fi71)) {
                            fmt.Println(v7, v7)
                        }
                        switch 42 {
                        case -359:
                            if false {
                                _ = len(v6)
                            }
                        }
                        var v72 int8 = 100
                        _ = fi71
                        _ = v72
                    }
                    v66 = v66
                    v73 := GenS6{F0: false, F1: 1}
                    _ = fi70
                    _ = v73
                }
            case "hello":
                if ((v66 * (v66 << ((1) & 63))) > v66) {
                    if v2 {
                        if ((v1 >> ((255) & 63)) == 42) {
                            var v74 float32 = float32(0.0)
                            var v75 int32 = 0
                            var v76 uint = 42
                            _ = v74
                            _ = v75
                            _ = v76
                        }
                    } else {
                        fmt.Println(v64, v1, v3)
                    }
                    v3 = (("" + v3) + (v3 + v3))
                    fmt.Println(v66, v7)
                } else {
                    var v77 string = v3
                    _ = v77
                }
            }
        case 707:
            for fi78 := 0; fi78 < 17; fi78++ {
                v1 += (v1 << ((893) & 63))
                _ = fi78
            }
        }
        switch v3 {
        case "\n\t":
            switch (v7 ^ v7) {
            case 2147483647:
                if ((-1 < -1) || v2) {
                    v79 := v4[v7]
                    switch ((("" + "hello") + (v3 + (v3 + "a"))) + ("\n\t" + v3)) {
                    case "\n\t":
                        v3 = v3
                    }
                    v80 := GenS8{F0: 1}
                    _ = v79
                    _ = v80
                }
            }
        default:
            v1 = (v1 >> ((65535) & 63))
        }
    }
    var v81 uint8 = uint8(v66)
    var v82 int64 = 0
    var v83 uint8 = (v81 * v81)
    switch v3 {
    case "\n\t":
        v84 := GenS1{F0: v3, F1: (!v2)}
        _ = v84
    case "":
        var v85 string = (v3 + v3)
        _ = v85
    case "foo bar":
        v2 = v2
    }
    for fi86 := 0; fi86 < 9; fi86++ {
        v87 := GenS9{F0: 42}
        _ = fi86
        _ = v87
    }
    fmt.Println(v81)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v64:", v64)
    fmt.Println("v66:", v66)
    fmt.Println("v7:", v7)
    fmt.Println("v81:", v81)
    fmt.Println("v82:", v82)
    fmt.Println("v83:", v83)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v6
    _ = v64
    _ = v65
    _ = v66
    _ = v7
    _ = v81
    _ = v82
    _ = v83
}
