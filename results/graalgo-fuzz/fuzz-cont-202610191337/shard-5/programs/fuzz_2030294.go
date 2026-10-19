package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 uint }
type GenS2 struct { F0 int; F1 int32; F2 uint64 }
type GenS3 struct { F0 string; F1 int64 }
type GenS4 struct { F0 uint; F1 int32 }
type GenS5 struct { F0 int16; F1 int16; F2 int16 }

func main() {
    v1 := []uint32{}
    v2 := []string{}
    v3 := make(map[int]string)
    var v4 uint32 = 2147483647
    v5 := GenS1{F0: ("a" + "a"), F1: 42, F2: 2}
    if (!false) {
        var v6 uint64 = 4294967295
        _ = v6
    }
    switch 128 {
    case -9223372036854775808:
        var v7 uint = 65535
        _ = v7
    }
    v8 := GenS1{F0: "\n\t", F1: -423, F2: 0}
    v9 := GenS1{F0: "hello", F1: 2147483647, F2: 1}
    for fi10 := 0; fi10 < 19; fi10++ {
        v4 += 751
        for fi11 := 0; fi11 < 10; fi11++ {
            v4 -= (v4 & v4)
            _ = fi11
        }
        var v12 int16 = int16(v4)
        _ = fi10
        _ = v12
    }
    if (0 <= 877) {
        if len(v2) > 0 {
            v2[0] = "\n\t"
        }
        switch ("\n\t" + ((("foo bar" + (("" + "a") + "hello")) + "\n\t") + "\n\t")) {
        case "\n\t":
            v3[9223372036854775807] = ("hello" + "hello")
        case "foo bar":
            if (32767 > 127) {
                for fi13 := 0; fi13 < 18; fi13++ {
                    v14 := v3[(^(fi13))]
                    switch 0 {
                    case 32767:
                        v14 = ((("" + v14) + ((v14 + "") + v14)) + v14)
                    case 0:
                        if (7 <= 27) {
                            for fi15 := 0; fi15 < 11; fi15++ {
                                var v16 int8 = 42
                                var v17 uint16 = 0
                                _ = fi15
                                _ = v16
                                _ = v17
                            }
                            if len(v2) > 0 {
                                v2[0] = (v14 + "\n\t")
                            }
                            if (-2.25 >= 0.1) {
                                if false {
                                    var v18 uint8 = 127
                                    _ = v18
                                }
                                var v19 uint8 = 127
                                _ = v19
                            } else {
                                v4 += (v4 / (((v4 / ((v4) | 1))) | 1))
                            }
                        } else {
                            for fi20 := 0; fi20 < 4; fi20++ {
                                var v21 bool = (false && true)
                                v1 = append(v1, (^(v4)))
                                _ = fi20
                                _ = v21
                            }
                        }
                    case -9223372036854775808:
                        var v22 int = -942
                        _ = v22
                    default:
                        var v23 uint = 345
                        _ = v23
                    }
                    _ = fi13
                    _ = v14
                }
            } else {
                fmt.Println(v4)
                v24 := GenS2{F0: 128, F1: -564, F2: 42}
                v4 += v4
                _ = v24
            }
        case "":
            if (true || false) {
                v3[-1] = "a"
                for fi25 := 0; fi25 < 6; fi25++ {
                    switch ("foo bar" + "") {
                    case "a":
                        if false {
                            switch "\n\t" {
                            case "\n\t":
                                v4 += v4
                            }
                            if (!((0 < 1) && (42 >= 1))) {
                                var v26 uint = 634
                                if (int64(v4) > -9223372036854775808) {
                                    var v27 string = (("foo bar" + "") + "foo bar")
                                    _ = v27
                                }
                                v3[121] = "\n\t"
                                _ = v26
                            }
                            v3[(fi25 << ((65535) & 63))] = ("a" + ("\n\t" + "a"))
                        } else {
                            switch "foo bar" {
                            case "":
                                v28 := GenS3{F0: ((((("\n\t" + "a") + ((("hello" + ("a" + ("\n\t" + ""))) + ("" + "\n\t")) + (("" + "") + "a"))) + "foo bar") + "\n\t") + ""), F1: 1}
                                _ = v28
                            case "\n\t":
                                v29 := GenS3{F0: "foo bar", F1: -273}
                                _ = v29
                            case "foo bar":
                                if false {
                                    fmt.Println(v4)
                                    var v30 float64 = float64(((((0.1 * -2.25) - ((float32((1.5 - 3.14159)) / ((float32(-2.25) / (0.1 * 3.14159)) * 1e10)) - 1.5)) - float32((1e10 * (-2.25 + ((1.5 * 0.1) / (-2.25 - float64(((0.1 - (-2.25 + 1e10)) + (0.1 * (float32((1e10 * (1e10 - 1.5))) + (((0.0 / (0.1 - ((((1e10 - 3.14159) + float32(1e10)) - (float32(((0.0 * float64((1.5 - 0.1))) * float64(1.5))) - 0.1)) * (1e10 + -2.25)))) - 0.0) - float32(3.14159)))))))))))) * 1e10))
                                    _ = v30
                                } else {
                                    var v31 int64 = 111
                                    _ = v31
                                }
                            default:
                                var v32 uint32 = (v4 / ((315) | 1))
                                _ = v32
                            }
                            switch "\n\t" {
                            case "\n\t":
                                var v33 uint64 = 1
                                _ = v33
                            case "a":
                                for fi34 := 0; fi34 < 19; fi34++ {
                                    var v35 int8 = 1
                                    _ = fi34
                                    _ = v35
                                }
                            default:
                                v4 = (^(v4))
                            }
                        }
                    case "\n\t":
                        v3[fi25] = ""
                    case "":
                        if false {
                            fmt.Println(v4, fi25)
                            v4 = v4
                            for fi36 := 0; fi36 < 18; fi36++ {
                                var v37 string = ""
                                switch 127 {
                                case -2147483648:
                                    var v38 float64 = (1e10 - 0.0)
                                    v38 -= v38
                                    _ = v38
                                case 32768:
                                    if true {
                                        v39 := GenS1{F0: "", F1: 682, F2: 396}
                                        if false {
                                            var v40 uint64 = 4294967295
                                            v41 := GenS2{F0: fi36, F1: 1, F2: (v40 >> ((0) & 63))}
                                            _ = v40
                                            _ = v41
                                        } else {
                                            _ = len(v1)
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            switch v37 {
                                            case "":
                                                fmt.Println(fi25, fi25)
                                            case "a":
                                                _ = len(v3)
                                            default:
                                                if (0.0 < (-2.25 / 3.14159)) {
                                                    v3[-2147483648] = (v37 + "foo bar")
                                                    switch ((v37 + v37) + ("hello" + ((v37 + v37) + (((v37 + v37) + ("\n\t" + "\n\t")) + "hello")))) {
                                                    case "hello":
                                                        v42 := GenS2{F0: 1, F1: 2147483647, F2: 42}
                                                        fmt.Println(fi36, v4, v37)
                                                        if (((4294967295 > 115) && (v4 != (v4 + v4))) && ((^(v4)) > 0)) {
                                                            var v43 float64 = (0.1 - 0.0)
                                                            switch fi25 {
                                                            case 776:
                                                                if ("hello" != v37) {
                                                                    fmt.Println(fi25)
                                                                    fmt.Println(v43, v37, v43)
                                                                    if false {
                                                                        v37 = "a"
                                                                        if false {
                                                                            fmt.Println(v43)
                                                                            var v44 uint64 = 661
                                                                            v3[2147483647] = (v37 + "hello")
                                                                            _ = v44
                                                                        }
                                                                        if true {
                                                                            v37 = ""
                                                                            var v45 uint64 = 913
                                                                            _ = v45
                                                                        } else {
                                                                            switch v37 {
                                                                            case "":
                                                                                v37 = v37
                                                                            default:
                                                                                v4 = v4
                                                                            }
                                                                        }
                                                                        v46 := GenS2{F0: 42, F1: -473, F2: uint64(v4)}
                                                                        v47 := GenS1{F0: "hello", F1: -1, F2: 18446744073709551615}
                                                                        _ = v46
                                                                        _ = v47
                                                                    }
                                                                }
                                                            }
                                                            v2 = append(v2, "a")
                                                            _ = v43
                                                        }
                                                        _ = v42
                                                    case "foo bar":
                                                        v37 = v37
                                                    case "a":
                                                        var v48 int64 = -150
                                                        _ = v48
                                                    default:
                                                        var v49 float32 = ((-2.25 * 3.14159) - -2.25)
                                                        _ = v49
                                                    }
                                                }
                                            }
                                        }
                                        _ = v39
                                    } else {
                                        v37 = (v37 + ("foo bar" + v37))
                                        var v50 int32 = -156
                                        var v51 int = (^(fi36))
                                        _ = v50
                                        _ = v51
                                    }
                                default:
                                    v37 = "a"
                                }
                                v4 = 374
                                _ = fi36
                                _ = v37
                            }
                        }
                    }
                    var v52 uint16 = 2
                    for fi53 := 0; fi53 < 19; fi53++ {
                        v54 := v3[fi25]
                        fmt.Println(v52, fi53)
                        if false {
                            v54 = v54
                            var v55 uint64 = 0
                            _ = v55
                        }
                        _ = fi53
                        _ = v54
                    }
                    _ = fi25
                    _ = v52
                }
            }
        }
    } else {
        for fi56 := 0; fi56 < 19; fi56++ {
            for fi57 := 0; fi57 < 20; fi57++ {
                v4 = v4
                v4 = v4
                v4 -= 36
                v58 := GenS2{F0: -1, F1: 2147483647, F2: 0}
                v59 := GenS4{F0: 0, F1: 1}
                _ = fi57
                _ = v58
                _ = v59
            }
            _ = fi56
        }
    }
    var v60 uint16 = 1
    if true {
        var v61 int32 = -2147483648
        if false {
            v4 = v4
            if false {
                for fi62 := 0; fi62 < 8; fi62++ {
                    for fi63 := 0; fi63 < 4; fi63++ {
                        switch ("hello" + "foo bar") {
                        case "":
                            v61 = (v61 + v61)
                        default:
                            var v64 float64 = -2.25
                            fmt.Println(v4, v61)
                            _ = v64
                        }
                        v60 = v60
                        var v65 int32 = v61
                        _ = fi63
                        _ = v65
                    }
                    _ = fi62
                }
                if (false && ((!(false || (-31 == 0))) || true)) {
                    var v66 uint16 = (v60 + v60)
                    var v67 float32 = ((float32(0.0) + (1e10 / 1.0)) * 3.14159)
                    v68 := v3[-9223372036854775808]
                    _ = v66
                    _ = v67
                    _ = v68
                } else {
                    if (-1 > 127) {
                        v61 += v61
                        v60 -= v60
                    } else {
                        v60 = v60
                    }
                    v60 = v60
                }
            }
            if false {
                fmt.Println(v4, v61)
                v69 := GenS2{F0: 127, F1: -2147483648, F2: 0}
                var v70 uint64 = 0
                _ = v69
                _ = v70
            } else {
                _ = len(v3)
                var v71 int16 = 1000
                v71 += (v71 & v71)
                _ = v71
            }
        }
        _ = v61
    } else {
        for fi72 := 0; fi72 < 20; fi72++ {
            var v73 uint16 = v60
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v3[fi72] = ("foo bar" + "foo bar")
            _ = fi72
            _ = v73
        }
        v60 = v60
        v4 -= (v4 % ((2) | 1))
    }
    v60 += v60
    if true {
        v60 += (v60 * v60)
        var v74 uint64 = 1
        v75 := GenS5{F0: 509, F1: 42, F2: 42}
        _ = v74
        _ = v75
    }
    v4 = v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v60
    _ = v8
    _ = v9
}
