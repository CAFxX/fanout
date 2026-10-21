package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint; F1 uint }
type GenS5 struct { F0 int16; F1 int32; F2 int8 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint32; F1 uint64 }
type GenS8 struct { F0 bool; F1 int32; F2 int32 }
type GenS9 struct { F0 bool; F1 int }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 int32; F1 int; F2 uint32 }
type GenS12 struct { F0 bool; F1 bool }
func genFn0(p0 bool, p1 uint64, p2 int) string {
    var v4 int8 = 0
    v5 := GenS1{F0: (p2 >> ((0) & 63)), F1: 302, F2: v4}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return "a"
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]int)
    v3 := make(map[int]string)
    var v6 string = "\n\t"
    var v7 uint64 = 816
    switch "\n\t" {
    case "\n\t":
        v8 := GenS1{F0: 42, F1: -1, F2: 0}
        _ = v8
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        v7 = (v7 / ((v7) | 1))
        var v10 int = (fi9 / ((fi9) | 1))
        _ = fi9
        _ = v10
    }
    switch v6 {
    case "hello":
        switch "a" {
        case "\n\t":
            for fi11 := 0; fi11 < 11; fi11++ {
                for fi12 := 0; fi12 < 8; fi12++ {
                    v13 := []int{fi12}
                    fmt.Println(v6, v6)
                    _ = fi12
                    _ = v13
                }
                for fi14 := 0; fi14 < 20; fi14++ {
                    var v15 uint16 = 0
                    if ((v15 * v15) <= 0) {
                        if false {
                            var v16 uint8 = 1
                            v17 := []int{fi14, (fi14 % (((^(fi14))) | 1)), 0}
                            if (true && false) {
                                var v18 uint8 = v16
                                if len(v17) >= 2 {
                                    v17 = v17[1:len(v17)]
                                }
                                var v19 uint64 = (v7 >> ((42) & 63))
                                _ = v18
                                _ = v19
                            } else {
                                if true {
                                    v15 += v15
                                    v16 = v16
                                    v7 = v7
                                    v20 := GenS2{F0: -326}
                                    var v21 uint64 = (v7 * (v7 / ((2) | 1)))
                                    switch (fi14 >> ((280) & 63)) {
                                    case -728:
                                        if (v21 <= (v21 / ((v7) | 1))) {
                                            v15 = v15
                                            v6 = v6
                                        }
                                    case 32768:
                                        v22 := GenS1{F0: (-(fi11)), F1: 1, F2: 100}
                                        _ = v22
                                    case 0:
                                        v16 += (v16 / ((v16) | 1))
                                    }
                                    fmt.Println(v16)
                                    _ = v20
                                    _ = v21
                                }
                            }
                            _ = v16
                            _ = v17
                        }
                        switch fi11 {
                        case 42:
                            var v23 uint = 255
                            _ = v23
                        case -1:
                            v24 := v2[(fi11 >> ((2) & 63))]
                            _ = v24
                        }
                        switch (^(fi14)) {
                        case 594:
                            if true {
                                v7 += 1
                            }
                        case -451:
                            if true {
                                if true {
                                    delete(v2, 4294967295)
                                    var v25 int16 = 7
                                    fmt.Println(v6, fi14, v25)
                                    _ = v25
                                }
                                v7 -= (v7 | (v7 % ((401) | 1)))
                                var v26 int64 = 0
                                _ = v26
                            }
                        }
                    }
                    _ = fi14
                    _ = v15
                }
                switch v6 {
                case "":
                    v27 := GenS1{F0: fi11, F1: -506, F2: 100}
                    _ = v27
                case "\n\t":
                    v28 := []int{}
                    _ = v28
                case "hello":
                    v29 := GenS2{F0: 1}
                    _ = v29
                default:
                    var v30 int16 = 42
                    _ = v30
                }
                _ = fi11
            }
        case "":
            for fi31 := 0; fi31 < 6; fi31++ {
                _ = len(v3)
                var v32 uint8 = 128
                _ = fi31
                _ = v32
            }
        default:
            v6 = v6
        }
    case "foo bar":
        switch ("hello" + "a") {
        case "hello":
            v6 = v6
        case "foo bar":
            v6 = (((v6 + v6) + (v6 + v6)) + "\n\t")
        case "a":
            var v33 float32 = (-2.25 - -2.25)
            _ = v33
        }
    case "":
        var v34 int32 = 2147483647
        _ = v34
    }
    v6 = ""
    v7 += v7
    v6 = ""
    v35 := []int{-9223372036854775808, 755, 32768}
    v6 = ("a" + "foo bar")
    fmt.Println(v7, v7)
    for fi36 := 0; fi36 < 15; fi36++ {
        fmt.Println(v6)
        var v37 uint64 = 1
        var v38 float32 = -2.25
        _ = fi36
        _ = v37
        _ = v38
    }
    _ = len(v2)
    for fi39 := 0; fi39 < 1; fi39++ {
        var v40 string = (("foo bar" + "hello") + "foo bar")
        for fi41 := 0; fi41 < 2; fi41++ {
            v2[(fi39 | int(v7))] = fi41
            _ = fi41
        }
        _ = fi39
        _ = v40
    }
    if false {
        for fi42 := 0; fi42 < 14; fi42++ {
            if (true || (!false)) {
                v6 = "foo bar"
            } else {
                for fi43 := 0; fi43 < 16; fi43++ {
                    v44 := GenS3{F0: (v7 & v7)}
                    delete(v2, -87)
                    if (true && (uint16(v7) > 32767)) {
                        v45 := GenS4{F0: 389, F1: 630}
                        if len(v35) > 0 {
                            v35[0] = 32768
                        }
                        _ = v45
                    } else {
                        var v46 int8 = -70
                        var v47 uint16 = 732
                        switch fi42 {
                        case 719:
                            var v48 float64 = (((0.0 + 1.5) + 1.5) - ((0.1 * 3.14159) * float64(1e10)))
                            v49 := GenS5{F0: -712, F1: 380, F2: v46}
                            fmt.Println(v6, v47)
                            _ = v48
                            _ = v49
                        default:
                            switch (v6 + v6) {
                            case "\n\t":
                                fmt.Println(v46, v47)
                            case "foo bar":
                                v50 := GenS6{F0: (v47 << ((603) & 63))}
                                _ = v50
                            case "hello":
                                var v51 float64 = 0.1
                                _ = v51
                            }
                        }
                        v52 := GenS7{F0: 65535, F1: (v7 >> ((4294967295) & 63))}
                        var v53 int = 32768
                        _ = v46
                        _ = v47
                        _ = v52
                        _ = v53
                    }
                    _ = fi43
                    _ = v44
                }
                v6 = ((v6 + "foo bar") + "hello")
                v7 += v7
            }
            var v54 bool = ((1.5 / -2.25) <= 1e10)
            if (9223372036854775807 >= -9223372036854775808) {
                if true {
                    v55 := GenS8{F0: (42 >= -1), F1: -499, F2: 0}
                    _ = v55
                } else {
                    var v56 float64 = 0.1
                    if (!false) {
                        v6 = v6
                    }
                    _ = v56
                }
                var v57 int64 = -809
                v6 = v6
                _ = v57
            } else {
                for fi58 := 0; fi58 < 1; fi58++ {
                    var v59 int16 = 32767
                    if true {
                        var v60 uint32 = 62
                        v61 := GenS9{F0: (false && (16 == 0)), F1: 2147483647}
                        if (true && true) {
                            if (255 == 65535) {
                                var v62 uint16 = 42
                                var v63 float32 = 3.14159
                                _ = v62
                                _ = v63
                            } else {
                                v60 += (v60 | v60)
                            }
                            switch (fi42 & (^(fi58))) {
                            case 735:
                                if (!(false && v54)) {
                                    fmt.Println(v7, fi58, v54)
                                    if (769 != 32767) {
                                        v59 = v59
                                        v64 := GenS2{F0: -186}
                                        if len(v35) > 0 {
                                            v35[0] = fi58
                                        }
                                        _ = v64
                                    } else {
                                        var v65 uint64 = (v7 >> ((18446744073709551615) & 63))
                                        v60 = 4294967295
                                        var v66 bool = (0 != 1)
                                        _ = v65
                                        _ = v66
                                    }
                                } else {
                                    v59 += 42
                                    var v67 string = "a"
                                    v67 = (v6 + (v67 + v67))
                                    _ = v67
                                }
                            case 42:
                                if false {
                                    if v54 {
                                        v7 = (^(v7))
                                        delete(v2, (fi58 / (((fi42 | fi58)) | 1)))
                                    } else {
                                        var v68 uint32 = (v60 & 42)
                                        fmt.Println(fi42)
                                        _ = v68
                                    }
                                    var v69 uint32 = v60
                                    _ = v69
                                } else {
                                    var v70 uint64 = 42
                                    _ = v70
                                }
                            case 1:
                                v59 = (^(v59))
                            default:
                                switch (fi42 ^ -86) {
                                case 277:
                                    if len(v35) > 0 {
                                        v35[0] = 9223372036854775807
                                    }
                                case 42:
                                    switch fi58 {
                                    case 42:
                                        var v71 float64 = float64(((-2.25 * (float32(((1.5 + 1e10) / 1e10)) * ((0.0 / 1.0) / (0.0 - 0.1)))) - 3.14159))
                                        _ = v71
                                    }
                                default:
                                    var v72 float64 = ((1e10 / 1.5) / 1.0)
                                    _ = v72
                                }
                            }
                        }
                        _ = v60
                        _ = v61
                    } else {
                        v73 := v2[fi42]
                        _ = v73
                    }
                    _ = fi58
                    _ = v59
                }
                var v74 float32 = (1e10 + ((3.14159 / 1.0) - 0.1))
                _ = v74
            }
            _ = fi42
            _ = v54
        }
        for fi75 := 0; fi75 < 12; fi75++ {
            var v76 float32 = 3.14159
            _ = fi75
            _ = v76
        }
    } else {
        if (!false) {
            for fi77 := 0; fi77 < 14; fi77++ {
                var v78 string = v6
                v35 = append(v35, (fi77 ^ fi77))
                _ = fi77
                _ = v78
            }
            var v79 uint16 = 2
            _ = v79
        } else {
            v80 := GenS10{F0: 42}
            _ = v80
        }
        v81 := GenS11{F0: 42, F1: -2147483648, F2: 65535}
        v82 := GenS12{F0: (false && (198 <= 128)), F1: (-466 < 1)}
        _ = v81
        _ = v82
    }
    v83 := v2[-649]
    var v84 uint32 = 119
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v83:", v83)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v6
    _ = v7
    _ = v83
    _ = v84
}
