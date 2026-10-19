package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 int32 }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 int32; F1 uint8 }
type GenS4 struct { F0 int64; F1 bool; F2 int }
type GenS5 struct { F0 uint64; F1 uint16; F2 int32 }
type GenS6 struct { F0 int64; F1 int64; F2 int16 }
func genFn0() uint64 {
    var v6 string = "foo bar"
    fmt.Println(v6)
    v7 := GenS1{F0: 0, F1: false, F2: 0}
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 uint8, p1 uint16) string {
    fmt.Println(p0, p1)
    var v8 uint32 = 2147483647
    v9 := []int{4294967295, 2147483647, 32767}
    if false {
        switch 127 {
        case -1:
            p0 += 1
        case 1:
            var v10 int64 = 42
            _ = v10
        case 32767:
            var v11 int8 = 100
            _ = v11
        default:
            p1 = p1
        }
        p0 += p0
    } else {
        for fi12 := 0; fi12 < 18; fi12++ {
            if ((!(450 == 267)) && (!(false || true))) {
                switch "foo bar" {
                case "a":
                    for fi13 := 0; fi13 < 8; fi13++ {
                        if (-1 == 1000) {
                            v14 := GenS1{F0: 1000, F1: (uint(p0) > uint(p1)), F2: -1}
                            if (v8 < v8) {
                                if false {
                                    v8 = (v8 * (v8 % ((v8) | 1)))
                                    p1 -= p1
                                    var v15 uint = 4294967295
                                    var v16 bool = ((p1 << ((v15) & 63)) == (p1 - uint16(v15)))
                                    if v16 {
                                        if (9223372036854775807 >= 58) {
                                            if v16 {
                                                var v17 string = ""
                                                if (320 <= (p1 * (p1 / ((p1) | 1)))) {
                                                    v18 := GenS2{F0: v16, F1: fi13}
                                                    var v19 string = (v17 + v17)
                                                    if len(v9) > 0 {
                                                        v9[0] = fi12
                                                    }
                                                    if ((fi12 | -2147483648) > 9223372036854775807) {
                                                        v17 = v17
                                                        if len(v9) >= 2 {
                                                            v9 = v9[1:len(v9)]
                                                        }
                                                    }
                                                    _ = v18
                                                    _ = v19
                                                } else {
                                                    var v20 uint32 = 308
                                                    switch (((v17 + "\n\t") + (v17 + (v17 + v17))) + (("" + "") + "")) {
                                                    case "a":
                                                        v21 := GenS1{F0: 414, F1: ("" > v17), F2: -2147483648}
                                                        _ = v21
                                                    case "foo bar":
                                                        if (v15 > 39) {
                                                            v9 = append(v9, fi12)
                                                        } else {
                                                            p1 -= p1
                                                            v20 += v8
                                                        }
                                                    default:
                                                        _ = len(v9)
                                                    }
                                                    v22 := GenS1{F0: 898, F1: v16, F2: 2147483647}
                                                    if v16 {
                                                        v16 = true
                                                        v8 = 0
                                                    } else {
                                                        p1 = p1
                                                        var v23 uint16 = uint16(v15)
                                                        _ = v23
                                                    }
                                                    v20 = v8
                                                    if len(v9) > 0 {
                                                        v9[0] = (fi12 & fi13)
                                                    }
                                                    var v24 uint64 = genFn0()
                                                    _ = v20
                                                    _ = v22
                                                    _ = v24
                                                }
                                                var v25 uint16 = uint16(p0)
                                                _ = v17
                                                _ = v25
                                            }
                                        } else {
                                            v15 = v15
                                        }
                                    }
                                    if (176 != p0) {
                                        if ((v16 || v16) || v16) {
                                            if ((p0 & p0) >= 2) {
                                                var v26 int = fi13
                                                v27 := GenS2{F0: false, F1: fi13}
                                                p0 += (p0 & uint8(p1))
                                                _ = v26
                                                _ = v27
                                            } else {
                                                v16 = true
                                                switch fi12 {
                                                case 797:
                                                    fmt.Println(p0, p1, fi13)
                                                case 127:
                                                    v8 -= v8
                                                case -55:
                                                    fmt.Println(v16, p0)
                                                default:
                                                    v8 = v8
                                                }
                                                if ((!((^(v8)) >= 672)) && v16) {
                                                    if len(v9) >= 2 {
                                                        v9 = v9[1:len(v9)]
                                                    }
                                                    var v28 uint = (v15 + v15)
                                                    p0 -= (p0 * p0)
                                                    _ = v28
                                                } else {
                                                    v8 += (v8 >> (((v15 << (((v15 / ((v15) | 1))) & 63))) & 63))
                                                }
                                            }
                                            fmt.Println(p0)
                                        }
                                        v8 = v8
                                        if v16 {
                                            v16 = (4 <= 42)
                                        } else {
                                            var v29 int = (-(fi13))
                                            v9 = append(v9, (fi13 % ((-1) | 1)))
                                            _ = v29
                                        }
                                    } else {
                                        var v30 uint32 = (^(v8))
                                        v30 -= (v8 - v8)
                                        if v16 {
                                            p0 = (p0 >> ((v15) & 63))
                                            v9 = append(v9, 9223372036854775807)
                                        }
                                        if (v16 && ((("" >= "hello") || (!v16)) && true)) {
                                            switch "a" {
                                            case "foo bar":
                                                v15 += v15
                                            case "":
                                                switch "foo bar" {
                                                case "foo bar":
                                                    if ((float64((float32(float64(((3.14159 + (0.1 * -2.25)) + 1.5))) - (((1.5 / 3.14159) - 0.1) + (1e10 * 3.14159)))) * ((float64(-2.25) / 1.0) * -2.25)) > 0.0) {
                                                        p0 = (p0 + (p0 / (((p0 | (p0 & (p0 >> ((uint(p1)) & 63))))) | 1)))
                                                        v8 = v8
                                                    }
                                                case "\n\t":
                                                    switch "hello" {
                                                    case "\n\t":
                                                        var v31 int64 = -554
                                                        if len(v9) >= 2 {
                                                            v9 = v9[1:len(v9)]
                                                        }
                                                        _ = v31
                                                    }
                                                case "a":
                                                    p0 = (^(p0))
                                                default:
                                                    p0 = (p0 + p0)
                                                }
                                            case "a":
                                                if v16 {
                                                    var v32 float64 = 1e10
                                                    _ = v32
                                                } else {
                                                    var v33 uint8 = 93
                                                    _ = v33
                                                }
                                            }
                                        }
                                        var v34 int64 = -119
                                        v35 := GenS1{F0: 1, F1: true, F2: 1}
                                        _ = v30
                                        _ = v34
                                        _ = v35
                                    }
                                    _ = len(v9)
                                    _ = v15
                                    _ = v16
                                }
                            } else {
                                switch "foo bar" {
                                case "a":
                                    if len(v9) > 0 {
                                        v9[0] = fi12
                                    }
                                default:
                                    var v36 int16 = 1
                                    _ = v36
                                }
                                v37 := GenS2{F0: (127 > 15), F1: fi12}
                                var v38 float32 = float32(-2.25)
                                v39 := GenS1{F0: -610, F1: (v38 >= (v38 + v38)), F2: 42}
                                var v40 float64 = -2.25
                                _ = v37
                                _ = v38
                                _ = v39
                                _ = v40
                            }
                            _ = v14
                        } else {
                            if (p1 <= (p1 % ((0) | 1))) {
                                fmt.Println(p0)
                                var v41 uint32 = v8
                                switch "hello" {
                                case "\n\t":
                                    if true {
                                        switch fi12 {
                                        case 4294967295:
                                            p1 += p1
                                        case 127:
                                            var v42 bool = (true && true)
                                            _ = v42
                                        }
                                        v8 += 864
                                        p1 -= (p1 ^ p1)
                                        var v43 int16 = -32768
                                        if ((0.0 + float32(3.14159)) < 3.14159) {
                                            var v44 int64 = 42
                                            _ = v44
                                        }
                                        fmt.Println(fi12, p0, v43)
                                        v41 = 4294967295
                                        if len(v9) > 0 {
                                            v9[0] = fi12
                                        }
                                        _ = v43
                                    } else {
                                        switch (fi12 - 2147483647) {
                                        case 0:
                                            if true {
                                                if ((fi12 % ((127) | 1)) > (fi12 + (fi12 >> ((uint(p1)) & 63)))) {
                                                    var v45 uint8 = 138
                                                    _ = v45
                                                } else {
                                                    var v46 int16 = 0
                                                    v47 := GenS3{F0: -1, F1: 127}
                                                    fmt.Println(v46, v41)
                                                    _ = v46
                                                    _ = v47
                                                }
                                                var v48 uint32 = v8
                                                _ = v48
                                            }
                                        default:
                                            fmt.Println(v41, v8)
                                            var v49 uint64 = 2
                                            var v50 uint = 0
                                            _ = v49
                                            _ = v50
                                        }
                                    }
                                    fmt.Println(p1, fi12)
                                    fmt.Println(v8)
                                    p0 = (p0 * (^(p0)))
                                case "":
                                    v8 = (^(v8))
                                default:
                                    p0 = 0
                                }
                                _ = v41
                            } else {
                                if ((v8 | v8) < (v8 & (v8 * (v8 >> ((255) & 63))))) {
                                    if len(v9) > 0 {
                                        v9[0] = (-(fi13))
                                    }
                                    v8 = (v8 + 0)
                                    switch ("\n\t" + "hello") {
                                    case "\n\t":
                                        if len(v9) >= 2 {
                                            v9 = v9[1:len(v9)]
                                        }
                                    case "foo bar":
                                        v8 = (v8 % ((v8) | 1))
                                    default:
                                        fmt.Println(v8)
                                    }
                                }
                            }
                            fmt.Println(p1, fi13, p1)
                        }
                        _ = fi13
                    }
                case "\n\t":
                    var v51 int16 = -32768
                    _ = v51
                case "hello":
                    p1 += p1
                default:
                    if true {
                        switch fi12 {
                        case 112:
                            v8 -= v8
                        }
                    } else {
                        v8 = v8
                        v8 = 65535
                    }
                }
            } else {
                v52 := GenS2{F0: ((false && (false && ((-40 < 100) && (!(!true))))) && false), F1: (fi12 - (fi12 << ((510) & 63)))}
                switch "a" {
                case "foo bar":
                    switch 128 {
                    case 127:
                        var v53 int64 = 1
                        _ = v53
                    case 32768:
                        if (!(!false)) {
                            switch fi12 {
                            case -713:
                                if len(v9) >= 2 {
                                    v9 = v9[1:len(v9)]
                                }
                            case 9223372036854775807:
                                var v54 uint8 = 1
                                _ = v54
                            }
                        } else {
                            var v55 uint32 = (v8 / (((v8 % (((v8 % (((v8 & uint32(p1))) | 1))) | 1))) | 1))
                            if true {
                                var v56 uint = 2
                                _ = v56
                            }
                            fmt.Println(v8, p0)
                            _ = v55
                        }
                    }
                case "":
                    v57 := GenS4{F0: -1, F1: false, F2: -9223372036854775808}
                    _ = v57
                default:
                    p1 = (p1 | 65535)
                }
                switch 462 {
                case -9223372036854775808:
                    var v58 float32 = 0.0
                    _ = v58
                }
                _ = v52
            }
            _ = fi12
        }
        for fi59 := 0; fi59 < 11; fi59++ {
            if ((p1 << ((627) & 63)) < (p1 | p1)) {
                v60 := GenS3{F0: -859, F1: p0}
                var v61 string = "\n\t"
                _ = v60
                _ = v61
            } else {
                v62 := GenS5{F0: 18446744073709551615, F1: uint16(p0), F2: int32(p1)}
                _ = v62
            }
            _ = fi59
        }
    }
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return "hello"
}

func main() {
    v1 := []int{}
    var v2 uint16 = 535
    var v3 bool = (127 == 42)
    v4 := make(map[uint32]uint32)
    v5 := make(map[int]uint32)
    var v63 int64 = 839
    v2 = (v2 ^ (^(v2)))
    if v3 {
        _ = len(v1)
    } else {
        var v64 int16 = -1
        v65 := v5[1]
        _ = v64
        _ = v65
    }
    switch 492 {
    case 32767:
        var v66 float32 = (1e10 / ((-2.25 - 1.5) - -2.25))
        _ = v66
    case 504:
        var v67 float64 = 1e10
        _ = v67
    case 612:
        for fi68 := 0; fi68 < 8; fi68++ {
            v63 = v63
            _ = fi68
        }
    }
    v2 += 32767
    if (!v3) {
        fmt.Println(v3, v63)
        switch 2147483647 {
        case -1:
            v2 += (^(v2))
        default:
            for fi69 := 0; fi69 < 9; fi69++ {
                v2 = v2
                _ = fi69
            }
        }
    }
    switch "" {
    case "hello":
        v70 := GenS6{F0: (v63 * (v63 - (-(v63)))), F1: v63, F2: -1}
        _ = v70
    default:
        v2 = (v2 + 0)
    }
    fmt.Println(v2, v2)
    v3 = v3
    if v3 {
        if (42 >= 9223372036854775807) {
            v3 = (4294967295 == 8)
            fmt.Println(v3)
        }
    }
    var v71 string = "\n\t"
    delete(v4, 2)
    delete(v5, -2147483648)
    for fi72 := 0; fi72 < 10; fi72++ {
        fmt.Println(fi72)
        _ = fi72
    }
    if false {
        if v3 {
            var v73 uint = 5
            _ = v73
        } else {
            v74 := GenS4{F0: v63, F1: (v3 && true), F2: 9223372036854775807}
            var v75 uint8 = 127
            _ = v74
            _ = v75
        }
        if true {
            if (v3 && (v3 && (false && (!v3)))) {
                if v3 {
                    fmt.Println(v63, v71)
                    for fi76 := 0; fi76 < 3; fi76++ {
                        v71 = v71
                        v5[(fi76 | (fi76 >> ((908) & 63)))] = 595
                        _ = fi76
                    }
                    v63 -= v63
                } else {
                    var v77 string = v71
                    _ = v77
                }
                for fi78 := 0; fi78 < 19; fi78++ {
                    switch (^(fi78)) {
                    case -1:
                        var v79 float32 = 0.1
                        _ = v79
                    }
                    v80 := GenS3{F0: 1, F1: 128}
                    _ = fi78
                    _ = v80
                }
            }
            fmt.Println(v3, v71, v63)
        } else {
            for fi81 := 0; fi81 < 14; fi81++ {
                v63 = (v63 / (((v63 & 9223372036854775807)) | 1))
                _ = fi81
            }
            fmt.Println(v71)
        }
    } else {
        fmt.Println(v71, v2)
        var v82 int64 = (v63 % (((v63 >> ((65535) & 63))) | 1))
        var v83 uint32 = 4294967295
        _ = v82
        _ = v83
    }
    v84 := GenS1{F0: -308, F1: (float32(float64(1.5)) >= float32((float64(((0.1 / 0.1) + (0.0 / ((float32(1.5) / 3.14159) - 0.0)))) / 1e10))), F2: 0}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v63:", v63)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v63
    _ = v71
    _ = v84
}
