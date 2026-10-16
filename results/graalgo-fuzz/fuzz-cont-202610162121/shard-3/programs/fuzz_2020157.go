package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 int16; F1 uint16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint; F1 int16 }
type GenS6 struct { F0 bool; F1 uint64; F2 string }
type GenS7 struct { F0 int32; F1 uint32 }
type GenS8 struct { F0 uint64; F1 bool }
type GenS9 struct { F0 bool; F1 int64 }
type GenS10 struct { F0 uint8; F1 uint32; F2 int8 }

func main() {
    var v1 bool = (!false)
    v2 := make(map[int]uint32)
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]int)
    var v5 uint16 = 296
    for fi6 := 0; fi6 < 19; fi6++ {
        v5 = (v5 >> ((630) & 63))
        _ = fi6
    }
    v5 += v5
    v7 := []int{0, 127, 2147483647, 42}
    if (255 > 18446744073709551615) {
        v8 := GenS1{F0: (v5 >> ((42) & 63)), F1: 58}
        for fi9 := 0; fi9 < 14; fi9++ {
            var v10 uint = 716
            v5 -= v5
            _ = fi9
            _ = v10
        }
        _ = v8
    } else {
        if v1 {
            if (!v1) {
                switch 127 {
                case 9223372036854775807:
                    v11 := GenS2{F0: (v5 << ((168) & 63)), F1: -96}
                    _ = v11
                case 127:
                    if (222 > 423) {
                        v12 := GenS3{F0: -1, F1: v5}
                        _ = v12
                    }
                case 310:
                    var v13 uint = 255
                    _ = v13
                default:
                    v5 = v5
                }
                if v1 {
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                    v4[0] = -458
                    switch ((("" + "a") + "\n\t") + "") {
                    case "foo bar":
                        v1 = true
                    default:
                        if v1 {
                            if false {
                                v14 := GenS2{F0: (^(v5)), F1: 100}
                                for fi15 := 0; fi15 < 1; fi15++ {
                                    v16 := GenS1{F0: v5, F1: 65535}
                                    v1 = (1 < 1)
                                    _ = fi15
                                    _ = v16
                                }
                                v17 := GenS4{F0: 255}
                                _ = v14
                                _ = v17
                            }
                        } else {
                            for fi18 := 0; fi18 < 6; fi18++ {
                                var v19 float64 = 1.5
                                var v20 int32 = 1
                                fmt.Println(fi18, v1)
                                _ = fi18
                                _ = v19
                                _ = v20
                            }
                            var v21 int8 = 123
                            _ = v21
                        }
                    }
                }
            } else {
                var v22 uint64 = 0
                _ = v22
            }
            v1 = (1 < 1)
        } else {
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            var v23 float64 = 3.14159
            switch "a" {
            case "hello":
                if true {
                    if len(v7) > 0 {
                        v7[0] = -52
                    }
                    switch 617 {
                    case 127:
                        if true {
                            v23 = v23
                            fmt.Println(v1, v5, v5)
                        }
                    case 128:
                        v23 = (((v23 * -2.25) / (v23 - v23)) * v23)
                    case -9223372036854775808:
                        v24 := v3[32767]
                        _ = v24
                    }
                } else {
                    switch "hello" {
                    case "hello":
                        v25 := GenS5{F0: 65535, F1: -32768}
                        _ = v25
                    case "":
                        switch "a" {
                        case "foo bar":
                            if len(v7) > 0 {
                                v7[0] = -2147483648
                            }
                        case "hello":
                            for fi26 := 0; fi26 < 13; fi26++ {
                                switch (^(fi26)) {
                                case -1:
                                    if v1 {
                                        for fi27 := 0; fi27 < 7; fi27++ {
                                            v5 = v5
                                            fmt.Println(fi27, v1, v23)
                                            if false {
                                                if (1 > -2147483648) {
                                                    v28 := GenS2{F0: 32767, F1: -58}
                                                    var v29 uint16 = v5
                                                    var v30 float32 = ((-2.25 * float32(((v23 * (v23 * 0.0)) * 1.5))) - 1e10)
                                                    _ = v28
                                                    _ = v29
                                                    _ = v30
                                                } else {
                                                    v1 = v1
                                                    fmt.Println(fi27, fi26)
                                                    if (0 > -2147483648) {
                                                        var v31 int32 = -1
                                                        if len(v7) >= 2 {
                                                            v7 = v7[1:len(v7)]
                                                        }
                                                        var v32 int32 = v31
                                                        _ = v31
                                                        _ = v32
                                                    }
                                                }
                                                if len(v7) > 0 {
                                                    v7[0] = fi26
                                                }
                                                v33 := GenS5{F0: 255, F1: -1}
                                                _ = v33
                                            }
                                            v5 = v5
                                            _ = fi27
                                        }
                                    }
                                case 950:
                                    fmt.Println(v23, v23)
                                case 1:
                                    var v34 string = "a"
                                    _ = v34
                                }
                                switch 128 {
                                case 539:
                                    v35 := GenS5{F0: 255, F1: -775}
                                    _ = v35
                                case 32767:
                                    v5 = (v5 ^ (v5 << ((85) & 63)))
                                }
                                _ = fi26
                            }
                        case "a":
                            if true {
                                v1 = (2 <= 2)
                                fmt.Println(v1, v1, v5)
                                v36 := GenS6{F0: (-2.25 >= 0.0), F1: 2, F2: ""}
                                _ = v36
                            } else {
                                var v37 string = "\n\t"
                                if v1 {
                                    for fi38 := 0; fi38 < 15; fi38++ {
                                        if (2 < 117) {
                                            fmt.Println(fi38, v1, fi38)
                                            switch v37 {
                                            case "foo bar":
                                                v23 -= v23
                                            }
                                            switch v37 {
                                            case "a":
                                                _ = len(v7)
                                            case "":
                                                for fi39 := 0; fi39 < 6; fi39++ {
                                                    switch -9223372036854775808 {
                                                    case 42:
                                                        v23 = 0.0
                                                    case -70:
                                                        v40 := GenS7{F0: -351, F1: 168}
                                                        _ = v40
                                                    default:
                                                        v7 = append(v7, 1)
                                                    }
                                                    _ = fi39
                                                }
                                            case "\n\t":
                                                var v41 uint64 = 42
                                                _ = v41
                                            }
                                        }
                                        if v1 {
                                            v7 = append(v7, (fi38 * fi38))
                                        }
                                        _ = fi38
                                    }
                                } else {
                                    switch ("hello" + v37) {
                                    case "":
                                        if v1 {
                                            for fi42 := 0; fi42 < 16; fi42++ {
                                                v43 := GenS1{F0: v5, F1: 180}
                                                _ = fi42
                                                _ = v43
                                            }
                                            if v1 {
                                                var v44 int64 = 1
                                                v5 = 813
                                                _ = v44
                                            }
                                            if len(v7) > 0 {
                                                v7[0] = -2147483648
                                            }
                                        } else {
                                            var v45 float32 = (((0.1 / ((1e10 / 1e10) - float32(v23))) + 3.14159) + (float32((((1.5 * ((((v23 * ((v23 / ((v23 * float64(-2.25)) * float64(0.0))) * v23)) - (v23 + 0.0)) / 1.5) * 1.5)) / ((0.1 + 1e10) / (1.5 * (v23 + (1e10 - float64(0.1)))))) / v23)) + (1e10 / ((1.5 + 3.14159) * 1.5))))
                                            for fi46 := 0; fi46 < 19; fi46++ {
                                                v4[224] = -1
                                                v45 -= (float32((v23 - v23)) - (float32(v23) / v45))
                                                _ = fi46
                                            }
                                            var v47 int8 = 127
                                            _ = v45
                                            _ = v47
                                        }
                                    case "foo bar":
                                        var v48 string = "a"
                                        _ = v48
                                    case "\n\t":
                                        for fi49 := 0; fi49 < 20; fi49++ {
                                            v23 = ((v23 / (v23 * v23)) * -2.25)
                                            _ = fi49
                                        }
                                    default:
                                        v5 = (v5 >> ((255) & 63))
                                    }
                                    if (v5 < (v5 | v5)) {
                                        v4[65535] = 9223372036854775807
                                        var v50 int8 = 0
                                        _ = v50
                                    }
                                }
                                _ = v37
                            }
                        }
                    case "a":
                        var v51 uint32 = 322
                        _ = v51
                    default:
                        v7 = append(v7, -723)
                    }
                }
            }
            _ = v23
        }
        var v52 bool = (-32768 != 32767)
        _ = v52
    }
    delete(v4, 65535)
    fmt.Println(v1, v1)
    if ((!false) || v1) {
        v5 -= v5
        v53 := GenS8{F0: 0, F1: (false && (!false))}
        _ = v53
    } else {
        switch -107 {
        case 445:
            var v54 int = -281
            _ = v54
        }
        var v55 bool = false
        _ = v55
    }
    v5 += v5
    v1 = v1
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    v56 := GenS3{F0: 1000, F1: (v5 / (((v5 | (v5 & 181))) | 1))}
    v5 -= 1
    if false {
        if false {
            switch "hello" {
            case "hello":
                v5 += (v5 % ((v5) | 1))
            case "foo bar":
                if (false || true) {
                    for fi57 := 0; fi57 < 20; fi57++ {
                        v58 := GenS3{F0: -32768, F1: (v5 - v5)}
                        _ = fi57
                        _ = v58
                    }
                }
            }
        }
    }
    fmt.Println(v1, v5)
    switch "a" {
    case "foo bar":
        if v1 {
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            v59 := GenS4{F0: 603}
            if (2147483648 != 4294967295) {
                switch 128 {
                case -166:
                    switch 1 {
                    case 970:
                        if v1 {
                            v60 := GenS4{F0: 18446744073709551615}
                            var v61 bool = (4294967295 == -241)
                            if ((4294967295 == 98) || (1 > -1)) {
                                switch ("a" + "") {
                                case "":
                                    if len(v7) >= 2 {
                                        v7 = v7[1:len(v7)]
                                    }
                                case "hello":
                                    var v62 float64 = float64((float32(-2.25) / float32((((-2.25 / 1.0) / (3.14159 / 1.5)) * -2.25))))
                                    _ = v62
                                case "foo bar":
                                    var v63 float64 = 1.5
                                    _ = v63
                                default:
                                    var v64 string = "foo bar"
                                    _ = v64
                                }
                                v65 := GenS2{F0: (v5 * v5), F1: 1}
                                switch 128 {
                                case 2147483647:
                                    var v66 uint32 = 843
                                    _ = v66
                                case -77:
                                    v67 := v4[2147483647]
                                    _ = v67
                                }
                                _ = v65
                            } else {
                                v1 = true
                                v61 = (!(0 < 11))
                            }
                            _ = v60
                            _ = v61
                        } else {
                            switch "" {
                            case "":
                                for fi68 := 0; fi68 < 17; fi68++ {
                                    var v69 string = "a"
                                    v69 = "\n\t"
                                    _ = fi68
                                    _ = v69
                                }
                            case "foo bar":
                                v70 := GenS9{F0: ((((-2.25 * (1e10 / 1.5)) - 0.0) / float32((-2.25 - 1.5))) <= 3.14159), F1: -9223372036854775808}
                                _ = v70
                            case "a":
                                var v71 float32 = 3.14159
                                _ = v71
                            }
                            if ((!(838 == 2147483648)) && v1) {
                                v1 = v1
                            } else {
                                var v72 float64 = (-2.25 + ((-2.25 + 3.14159) + float64(0.0)))
                                fmt.Println(v5, v5, v1)
                                v73 := v3[4294967295]
                                _ = v72
                                _ = v73
                            }
                        }
                    default:
                        v5 = v5
                    }
                case 42:
                    v1 = (-1 > -2147483648)
                default:
                    v74 := GenS4{F0: 4294967295}
                    _ = v74
                }
            } else {
                for fi75 := 0; fi75 < 10; fi75++ {
                    if (-2147483648 == -2147483648) {
                        for fi76 := 0; fi76 < 1; fi76++ {
                            if (((true || v1) || (0 == 110)) && true) {
                                var v77 int16 = 42
                                v77 += (v77 - v77)
                                _ = v77
                            }
                            if v1 {
                                switch fi75 {
                                case 32768:
                                    v1 = (!v1)
                                case 2147483647:
                                    var v78 float64 = float64(float32(0.1))
                                    _ = v78
                                case 127:
                                    v79 := v3[fi75]
                                    _ = v79
                                default:
                                    fmt.Println(v5, fi76, fi75)
                                }
                                v5 -= 337
                            } else {
                                if len(v7) >= 2 {
                                    v7 = v7[1:len(v7)]
                                }
                                var v80 bool = true
                                _ = v80
                            }
                            var v81 int32 = -1
                            _ = fi76
                            _ = v81
                        }
                        v5 -= 42
                        switch fi75 {
                        case 127:
                            if v1 {
                                _ = len(v3)
                                v1 = (!(!(!v1)))
                            } else {
                                v1 = false
                            }
                        case 42:
                            v82 := GenS10{F0: 255, F1: 1, F2: -1}
                            _ = v82
                        }
                    }
                    _ = fi75
                }
            }
            _ = v59
        } else {
            fmt.Println(v1, v5)
            v83 := v3[-872]
            _ = v83
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v7
}
