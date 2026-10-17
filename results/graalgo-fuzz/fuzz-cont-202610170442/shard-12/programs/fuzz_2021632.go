package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 int32; F1 int64; F2 int }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int8; F1 int64; F2 int16 }
type GenS5 struct { F0 uint32; F1 uint32; F2 bool }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int16; F1 uint64; F2 int64 }
type GenS9 struct { F0 int32; F1 uint64; F2 int }
func genFn0() uint16 {
    var v7 string = ("hello" + ((((("hello" + "\n\t") + "\n\t") + "\n\t") + "hello") + ""))
    for fi8 := 0; fi8 < 9; fi8++ {
        v7 = v7
        _ = fi8
    }
    var v9 uint16 = 32767
    var v10 uint32 = 4294967295
    _ = v10
    _ = v7
    _ = v9
    return v9
}

func main() {
    v1 := make(map[int]bool)
    v2 := []uint32{}
    var v3 uint32 = 2147483647
    v4 := make(map[uint32]string)
    var v5 bool = (v3 >= v3)
    var v6 uint32 = v3
    switch "a" {
    case "":
        for fi11 := 0; fi11 < 6; fi11++ {
            for fi12 := 0; fi12 < 11; fi12++ {
                switch "hello" {
                case "hello":
                    switch (fi11 % ((fi12) | 1)) {
                    case 128:
                        v2 = append(v2, v6)
                    case 127:
                        v3 = (v6 >> ((4294967295) & 63))
                    case 0:
                        v5 = v5
                    default:
                        var v13 uint32 = (v6 ^ 694)
                        _ = v13
                    }
                default:
                    switch ("foo bar" + "\n\t") {
                    case "\n\t":
                        if (!(1 == 677)) {
                            var v14 int32 = 715
                            _ = v14
                        }
                    }
                }
                var v15 float64 = 0.0
                var v16 int16 = -92
                v16 -= v16
                _ = fi12
                _ = v15
                _ = v16
            }
            _ = fi11
        }
    case "a":
        fmt.Println(v6, v3, v6)
    case "\n\t":
        v17 := v1[32768]
        _ = v17
    }
    v18 := v4[v3]
    switch v18 {
    case "\n\t":
        v18 = (v18 + "foo bar")
    case "a":
        if (false && ((!(42 <= 775)) || true)) {
            v5 = (255 > 2)
            for fi19 := 0; fi19 < 13; fi19++ {
                switch v18 {
                case "":
                    if true {
                        if v5 {
                            var v20 uint32 = (v6 * 744)
                            var v21 uint8 = 9
                            _ = v20
                            _ = v21
                        }
                        for fi22 := 0; fi22 < 10; fi22++ {
                            v5 = (3.14159 >= float64(1.5))
                            if len(v2) > 0 {
                                v2[0] = v6
                            }
                            _ = fi22
                        }
                        v6 -= 498
                    } else {
                        var v23 uint16 = 1
                        var v24 uint = 1
                        _ = v23
                        _ = v24
                    }
                case "\n\t":
                    for fi25 := 0; fi25 < 15; fi25++ {
                        v18 = v18
                        if v5 {
                            switch 1 {
                            case -9223372036854775808:
                                if (281 >= v6) {
                                    var v26 uint64 = 42
                                    v3 -= (v6 | 947)
                                    _ = v26
                                } else {
                                    v27 := GenS1{F0: v18, F1: 489}
                                    switch v18 {
                                    case "foo bar":
                                        if (!true) {
                                            if v5 {
                                                switch 1 {
                                                case 491:
                                                    var v28 int8 = 62
                                                    _ = v28
                                                }
                                                if (32767 <= 0) {
                                                    v3 = 2147483648
                                                    v29 := GenS1{F0: "", F1: 18446744073709551615}
                                                    _ = v29
                                                } else {
                                                    fmt.Println(fi19, v18)
                                                }
                                                if v5 {
                                                    if (255 >= 37) {
                                                        v3 = v6
                                                    } else {
                                                        v2 = append(v2, (v3 - v3))
                                                        switch fi25 {
                                                        case 790:
                                                            var v30 bool = (float32((3.14159 + 3.14159)) > (1e10 * 0.0))
                                                            _ = v30
                                                        case -985:
                                                            if (45 != (v6 - (v3 << ((4294967295) & 63)))) {
                                                                v3 -= v6
                                                            } else {
                                                                v6 = v6
                                                                v3 = 65535
                                                                var v31 uint32 = 2
                                                                var v32 float32 = 1.5
                                                                var v33 string = v18
                                                                var v34 int = 32768
                                                                _ = v31
                                                                _ = v32
                                                                _ = v33
                                                                _ = v34
                                                            }
                                                        }
                                                    }
                                                    fmt.Println(v5, fi19, fi25)
                                                } else {
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    v35 := GenS1{F0: v18, F1: 4294967295}
                                                    _ = v35
                                                }
                                            }
                                            fmt.Println(v6, v3, v5)
                                            switch v18 {
                                            case "":
                                                switch (-(fi19)) {
                                                case 702:
                                                    if v5 {
                                                        v36 := GenS2{F0: -2147483648, F1: 1, F2: fi25}
                                                        _ = v36
                                                    } else {
                                                        if (!v5) {
                                                            v6 += (^(v3))
                                                            delete(v1, (fi25 | (fi19 % ((fi19) | 1))))
                                                            v37 := GenS2{F0: 349, F1: -1, F2: (^(fi19))}
                                                            _ = v37
                                                        }
                                                        v38 := GenS3{F0: v18}
                                                        _ = v38
                                                    }
                                                default:
                                                    v39 := GenS4{F0: -1, F1: 0, F2: 32767}
                                                    _ = v39
                                                }
                                            case "foo bar":
                                                v40 := GenS5{F0: 42, F1: (v3 % ((v6) | 1)), F2: false}
                                                _ = v40
                                            case "hello":
                                                var v41 string = v18
                                                _ = v41
                                            default:
                                                switch v18 {
                                                case "a":
                                                    var v42 int32 = -37
                                                    _ = v42
                                                case "foo bar":
                                                    if (32767 >= -32768) {
                                                        switch (fi25 - fi25) {
                                                        case 660:
                                                            switch "foo bar" {
                                                            case "hello":
                                                                v4[4294967295] = v18
                                                            }
                                                        case -1:
                                                            var v43 uint = 0
                                                            _ = v43
                                                        default:
                                                            switch ("hello" + v18) {
                                                            case "hello":
                                                                var v44 int8 = 21
                                                                _ = v44
                                                            default:
                                                                var v45 uint32 = 0
                                                                _ = v45
                                                            }
                                                        }
                                                        v46 := GenS2{F0: 42, F1: -1, F2: fi19}
                                                        v18 = "hello"
                                                        _ = v46
                                                    } else {
                                                        v3 -= (v6 * v6)
                                                        v47 := GenS3{F0: ("foo bar" + ("a" + (v18 + v18)))}
                                                        switch "a" {
                                                        case "a":
                                                            v48 := GenS6{F0: 2}
                                                            _ = v48
                                                        default:
                                                            var v49 float64 = (-2.25 / -2.25)
                                                            _ = v49
                                                        }
                                                        _ = v47
                                                    }
                                                case "\n\t":
                                                    var v50 int32 = 42
                                                    _ = v50
                                                default:
                                                    var v51 int = fi25
                                                    _ = v51
                                                }
                                            }
                                        } else {
                                            v6 += (v3 | 4294967295)
                                        }
                                    default:
                                        v2 = append(v2, v6)
                                    }
                                    _ = v27
                                }
                            case 9223372036854775807:
                                switch ("hello" + "a") {
                                case "\n\t":
                                    v52 := v4[(v6 >> ((645) & 63))]
                                    _ = v52
                                case "foo bar":
                                    if v5 {
                                        var v53 string = "\n\t"
                                        fmt.Println(v3, v5, v6)
                                        _ = v53
                                    } else {
                                        v3 = v3
                                        if v5 {
                                            _ = len(v2)
                                            fmt.Println(v18)
                                        }
                                        if len(v2) > 0 {
                                            v2[0] = v6
                                        }
                                    }
                                default:
                                    var v54 int8 = -1
                                    _ = v54
                                }
                            case 0:
                                v55 := GenS5{F0: v3, F1: (v3 / (((v3 ^ 65535)) | 1)), F2: false}
                                _ = v55
                            }
                            v6 -= v3
                        } else {
                            v6 += 42
                        }
                        _ = fi25
                    }
                }
                var v56 int64 = 0
                _ = fi19
                _ = v56
            }
        }
    case "foo bar":
        for fi57 := 0; fi57 < 14; fi57++ {
            _ = len(v1)
            switch (("foo bar" + "\n\t") + "a") {
            case "\n\t":
                fmt.Println(v5, fi57, fi57)
            }
            var v58 float64 = 0.0
            _ = fi57
            _ = v58
        }
    default:
        var v59 int64 = 42
        _ = v59
    }
    v60 := GenS7{F0: 1}
    v18 = (v18 + "")
    _ = len(v2)
    for fi61 := 0; fi61 < 6; fi61++ {
        if len(v2) > 0 {
            v2[0] = v3
        }
        for fi62 := 0; fi62 < 11; fi62++ {
            if v5 {
                switch "hello" {
                case "\n\t":
                    if (!v5) {
                        v3 = v6
                    }
                case "hello":
                    v2 = append(v2, (v6 * 0))
                case "":
                    switch v18 {
                    case "":
                        v63 := GenS8{F0: 985, F1: 1, F2: -1}
                        _ = v63
                    case "a":
                        switch v18 {
                        case "foo bar":
                            v64 := GenS9{F0: -533, F1: 2, F2: fi61}
                            _ = v64
                        case "a":
                            v65 := v4[(v6 | (^(v3)))]
                            _ = v65
                        }
                        _ = len(v1)
                    case "foo bar":
                        if v5 {
                            v2 = append(v2, 2147483647)
                            v66 := v1[(fi61 >> ((255) & 63))]
                            var v67 uint = 328
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            if (!v5) {
                                v18 = ((v18 + v18) + v18)
                                v3 -= (v6 ^ 1)
                                v66 = false
                            }
                            var v68 uint16 = 1
                            v67 = (v67 ^ (v67 << ((v67) & 63)))
                            v5 = ((!false) && false)
                            var v69 bool = false
                            var v70 int = (fi62 - 1)
                            _ = v66
                            _ = v67
                            _ = v68
                            _ = v69
                            _ = v70
                        } else {
                            switch "a" {
                            case "\n\t":
                                var v71 int8 = 42
                                _ = v71
                            case "":
                                v72 := GenS5{F0: (v6 / (((v3 << ((0) & 63))) | 1)), F1: (v6 * 464), F2: ((v6 * 4294967295) <= v6)}
                                _ = v72
                            }
                        }
                    default:
                        v73 := GenS5{F0: (v6 ^ v3), F1: (v3 ^ 2147483647), F2: (v5 && (!false))}
                        _ = v73
                    }
                default:
                    var v74 uint32 = v3
                    _ = v74
                }
                v2 = append(v2, (v3 / ((v3) | 1)))
            }
            _ = fi62
        }
        v75 := GenS1{F0: (v18 + v18), F1: 0}
        _ = fi61
        _ = v75
    }
    v3 += 4294967295
    v5 = (!(("hello" == "a") || v5))
    delete(v1, 32767)
    v76 := GenS6{F0: 1}
    v6 = 4294967295
    for fi77 := 0; fi77 < 13; fi77++ {
        var v78 string = (("" + v18) + v18)
        _ = fi77
        _ = v78
    }
    if (127 == int(v6)) {
        v79 := GenS7{F0: -153}
        if (!(!(false && (((!(262 > 18446744073709551615)) && v5) || true)))) {
            var v80 float64 = 1e10
            _ = v80
        } else {
            v81 := GenS6{F0: 167}
            if (!(!v5)) {
                fmt.Println(v18, v3)
            } else {
                for fi82 := 0; fi82 < 14; fi82++ {
                    var v83 float64 = 1e10
                    _ = fi82
                    _ = v83
                }
                if (-1 >= 0) {
                    fmt.Println(v5, v6, v6)
                }
                fmt.Println(v6)
            }
            _ = v81
        }
        _ = v79
    } else {
        v4[4294967295] = ("\n\t" + v18)
        for fi84 := 0; fi84 < 14; fi84++ {
            v85 := GenS4{F0: -128, F1: -232, F2: 0}
            fmt.Println(fi84, v18, v3)
            fmt.Println(v3, v5)
            _ = fi84
            _ = v85
        }
        _ = len(v2)
    }
    v86 := GenS7{F0: 42}
    v87 := GenS8{F0: -1, F1: 482, F2: -423}
    for fi88 := 0; fi88 < 9; fi88++ {
        var v89 int16 = -32768
        _ = fi88
        _ = v89
    }
    var v90 uint64 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v60
    _ = v76
    _ = v86
    _ = v87
    _ = v90
}
