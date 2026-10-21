package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 uint32 }
type GenS2 struct { F0 bool; F1 int64; F2 bool }
type GenS3 struct { F0 uint16; F1 int32; F2 uint8 }
type GenS4 struct { F0 int; F1 uint8; F2 int32 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 bool; F1 string }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int32; F1 uint64; F2 int8 }
type GenS10 struct { F0 uint16; F1 int16; F2 uint }

func main() {
    v1 := []string{}
    v2 := make(map[uint32]uint32)
    var v3 uint8 = 32
    var v4 int32 = -2147483648
    var v5 bool = true
    v4 -= v4
    if v5 {
        fmt.Println(v5, v5, v3)
    } else {
        switch ("a" + "\n\t") {
        case "a":
            var v6 int16 = 515
            _ = v6
        case "\n\t":
            v3 = (v3 + 128)
        case "hello":
            var v7 uint = 684
            _ = v7
        default:
            for fi8 := 0; fi8 < 5; fi8++ {
                var v9 float32 = float32(1.5)
                v10 := GenS1{F0: 740, F1: 0, F2: 65535}
                _ = fi8
                _ = v10
                _ = v9
            }
        }
        delete(v2, 2147483647)
        if (941 <= 571) {
            switch 9223372036854775807 {
            case 428:
                for fi11 := 0; fi11 < 16; fi11++ {
                    v12 := GenS2{F0: (4294967295 < uint64(v3)), F1: 42, F2: v5}
                    var v13 bool = false
                    var v14 uint16 = 587
                    _ = fi11
                    _ = v12
                    _ = v13
                    _ = v14
                }
            default:
                if (((42 < 1) || v5) && v5) {
                    var v15 int = 32768
                    for fi16 := 0; fi16 < 6; fi16++ {
                        switch (fi16 & 32767) {
                        case 32767:
                            v3 -= (v3 << ((4294967295) & 63))
                        }
                        v17 := GenS2{F0: (false && ((!true) && v5)), F1: -957, F2: v5}
                        _ = fi16
                        _ = v17
                    }
                    _ = v15
                }
            }
            if (-298 == 10) {
                if (((!(1 <= -9223372036854775808)) || ("" != "")) || ((9223372036854775807 <= 1) && v5)) {
                    switch -154 {
                    case 572:
                        v18 := GenS3{F0: 632, F1: (v4 & (v4 - -2147483648)), F2: 128}
                        _ = v18
                    case 367:
                        for fi19 := 0; fi19 < 9; fi19++ {
                            if v5 {
                                var v20 float32 = 0.0
                                v5 = v5
                                _ = len(v1)
                                _ = v20
                            } else {
                                v5 = v5
                                v4 -= -1
                                v21 := GenS2{F0: v5, F1: 588, F2: (0.1 != (float32((0.0 + 1e10)) / 3.14159))}
                                _ = v21
                            }
                            _ = fi19
                        }
                    default:
                        var v22 int64 = -9223372036854775808
                        _ = v22
                    }
                    v3 += (v3 + (v3 << ((42) & 63)))
                    v23 := GenS1{F0: 18446744073709551615, F1: 0, F2: 65535}
                    _ = v23
                }
                if v5 {
                    v3 += v3
                }
            } else {
                for fi24 := 0; fi24 < 9; fi24++ {
                    if (!v5) {
                        switch (("hello" + "\n\t") + ("\n\t" + "hello")) {
                        case "a":
                            fmt.Println(v4, v5)
                        case "hello":
                            var v25 int32 = -258
                            _ = v25
                        case "\n\t":
                            if ((v3 & 0) == v3) {
                                _ = len(v1)
                                for fi26 := 0; fi26 < 2; fi26++ {
                                    v4 = (v4 - (^(v4)))
                                    _ = fi26
                                }
                                for fi27 := 0; fi27 < 14; fi27++ {
                                    v3 += v3
                                    _ = fi27
                                }
                            }
                        }
                        v28 := GenS2{F0: true, F1: -659, F2: (false && v5)}
                        var v29 float32 = (float32((0.0 - (((0.0 - (0.1 + (3.14159 + 1.5))) * 1e10) - (0.0 + 0.0)))) / float32(float64(-2.25)))
                        _ = v28
                        _ = v29
                    }
                    var v30 int16 = -705
                    _ = fi24
                    _ = v30
                }
                for fi31 := 0; fi31 < 20; fi31++ {
                    switch -585 {
                    case -9223372036854775808:
                        v32 := GenS2{F0: (0.1 >= (((((1.5 + 3.14159) / 3.14159) * float32(3.14159)) - 3.14159) / (1.5 - 0.0))), F1: -1, F2: (v5 || (547 <= 976))}
                        _ = v32
                    case -2147483648:
                        var v33 uint8 = v3
                        _ = v33
                    default:
                        if (v5 && ((v4 / (((v4 % ((42) | 1))) | 1)) <= 42)) {
                            var v34 float32 = 1.5
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            for fi35 := 0; fi35 < 19; fi35++ {
                                if (v4 >= v4) {
                                    if v5 {
                                        switch 127 {
                                        case 999:
                                            if (v4 < (v4 << ((0) & 63))) {
                                                if (65535 > 0) {
                                                    v36 := GenS3{F0: 65535, F1: (-(v4)), F2: 29}
                                                    _ = v36
                                                } else {
                                                    var v37 float32 = v34
                                                    v38 := GenS3{F0: 65535, F1: (v4 - (v4 % ((v4) | 1))), F2: (v3 ^ 42)}
                                                    _ = v37
                                                    _ = v38
                                                }
                                                v3 = v3
                                                var v39 float32 = (v34 + (v34 + ((v34 * 3.14159) + (v34 / v34))))
                                                _ = v39
                                            }
                                        case 42:
                                            v3 -= (v3 * 1)
                                        case 0:
                                            v3 = v3
                                            switch (fi35 ^ fi35) {
                                            case 0:
                                                if (!false) {
                                                    var v40 uint16 = 1
                                                    if (true || v5) {
                                                        fmt.Println(fi35, v34)
                                                    }
                                                    switch ("a" + "hello") {
                                                    case "foo bar":
                                                        fmt.Println(v40, v40)
                                                    case "a":
                                                        if (int32(v3) <= 1) {
                                                            fmt.Println(v34, v34)
                                                            var v41 bool = ((v4 | -1) <= (v4 & (^(v4))))
                                                            _ = v41
                                                        }
                                                    case "\n\t":
                                                        var v42 int64 = 1
                                                        v43 := v2[65535]
                                                        var v44 string = (("hello" + "") + ("hello" + "a"))
                                                        switch (fi31 | fi31) {
                                                        case 1:
                                                            var v45 int64 = v42
                                                            _ = v45
                                                        }
                                                        v34 += v34
                                                        _ = v42
                                                        _ = v43
                                                        _ = v44
                                                    default:
                                                        switch fi31 {
                                                        case 42:
                                                            fmt.Println(fi35, fi31, v5)
                                                        case -429:
                                                            v40 = (v40 / ((v40) | 1))
                                                        case 0:
                                                            v3 += (v3 + 127)
                                                        }
                                                    }
                                                    _ = v40
                                                }
                                            }
                                        default:
                                            var v46 uint8 = v3
                                            _ = v46
                                        }
                                    }
                                    v47 := GenS4{F0: 128, F1: (^(v3)), F2: (-(v4))}
                                    _ = v47
                                } else {
                                    if len(v1) > 0 {
                                        v1[0] = ""
                                    }
                                }
                                _ = fi35
                            }
                            _ = v34
                        } else {
                            if ((!true) && (127 >= v3)) {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                v4 += v4
                            }
                        }
                    }
                    switch "hello" {
                    case "a":
                        if (2 < 65535) {
                            var v48 int8 = -128
                            if len(v1) > 0 {
                                v1[0] = "foo bar"
                            }
                            for fi49 := 0; fi49 < 8; fi49++ {
                                delete(v2, 178)
                                switch (fi49 >> ((2) & 63)) {
                                case 0:
                                    if v5 {
                                        var v50 uint = 18446744073709551615
                                        var v51 uint16 = 0
                                        _ = v50
                                        _ = v51
                                    }
                                case 32768:
                                    fmt.Println(v5, fi49)
                                default:
                                    var v52 int8 = (v48 >> ((65535) & 63))
                                    _ = v52
                                }
                                _ = fi49
                            }
                            _ = v48
                        }
                    case "\n\t":
                        var v53 uint8 = v3
                        _ = v53
                    default:
                        v3 = (v3 % ((v3) | 1))
                    }
                    v54 := GenS5{F0: 0}
                    _ = fi31
                    _ = v54
                }
            }
        }
    }
    var v55 string = "\n\t"
    delete(v2, 327)
    fmt.Println(v3)
    var v56 string = ("foo bar" + v55)
    delete(v2, 2)
    if v5 {
        if (13 != -912) {
            fmt.Println(v5, v4, v56)
            var v57 int8 = 74
            var v58 int16 = 286
            _ = v57
            _ = v58
        }
        if (2147483648 >= uint32(v3)) {
            for fi59 := 0; fi59 < 13; fi59++ {
                switch 42 {
                case 128:
                    var v60 uint8 = (v3 * 128)
                    _ = v60
                }
                switch v56 {
                case "foo bar":
                    v3 = v3
                case "a":
                    if false {
                        v4 = v4
                    } else {
                        var v61 uint32 = 146
                        for fi62 := 0; fi62 < 19; fi62++ {
                            var v63 uint8 = uint8(v61)
                            _ = fi62
                            _ = v63
                        }
                        _ = v61
                    }
                case "":
                    v64 := GenS2{F0: v5, F1: -179, F2: false}
                    _ = v64
                default:
                    for fi65 := 0; fi65 < 5; fi65++ {
                        if (((v56 + "foo bar") + v55) < v56) {
                            var v66 int32 = v4
                            v1 = append(v1, "a")
                            v67 := GenS6{F0: v5, F1: v56}
                            var v68 float32 = 3.14159
                            if (int64(v4) <= -9223372036854775808) {
                                delete(v2, 659)
                                _ = len(v2)
                            } else {
                                var v69 float64 = 0.0
                                v70 := GenS7{F0: uint64(v3)}
                                delete(v2, 2147483647)
                                _ = v69
                                _ = v70
                            }
                            _ = v66
                            _ = v67
                            _ = v68
                        } else {
                            switch ("foo bar" + v56) {
                            case "hello":
                                v71 := GenS3{F0: 0, F1: v4, F2: 42}
                                _ = v71
                            case "foo bar":
                                if true {
                                    v72 := GenS8{F0: ("hello" + "a")}
                                    _ = v72
                                }
                            default:
                                v55 = "a"
                            }
                            v3 -= v3
                            v73 := GenS6{F0: v5, F1: (v55 + ("" + v56))}
                            v74 := GenS9{F0: v4, F1: 400, F2: -128}
                            _ = v73
                            _ = v74
                        }
                        _ = fi65
                    }
                }
                switch ((v56 + ((v55 + "foo bar") + (("\n\t" + "hello") + ""))) + v56) {
                case "hello":
                    v75 := GenS10{F0: 65535, F1: -952, F2: 2}
                    _ = v75
                case "foo bar":
                    switch fi59 {
                    case -358:
                        var v76 string = v55
                        _ = v76
                    case -382:
                        for fi77 := 0; fi77 < 9; fi77++ {
                            var v78 float32 = (1e10 + -2.25)
                            var v79 uint8 = v3
                            _ = fi77
                            _ = v78
                            _ = v79
                        }
                    default:
                        for fi80 := 0; fi80 < 13; fi80++ {
                            if false {
                                v1 = append(v1, "foo bar")
                            } else {
                                switch ((v56 + (v55 + v55)) + v56) {
                                case "a":
                                    v4 += v4
                                case "":
                                    v56 = v55
                                }
                                var v81 int16 = 32767
                                _ = v81
                            }
                            fmt.Println(fi59)
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = fi80
                        }
                    }
                default:
                    for fi82 := 0; fi82 < 8; fi82++ {
                        fmt.Println(v3)
                        var v83 uint64 = 2
                        _ = fi82
                        _ = v83
                    }
                }
                _ = fi59
            }
        }
    } else {
        for fi84 := 0; fi84 < 16; fi84++ {
            v2[0] = 42
            v85 := GenS1{F0: 18446744073709551615, F1: 4294967295, F2: 0}
            v2[2] = 375
            _ = fi84
            _ = v85
        }
    }
    var v86 uint32 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v56
    _ = v86
}
