package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint32; F1 int8 }
type GenS5 struct { F0 bool; F1 int8 }

func main() {
    var v1 float32 = 1e10
    var v2 float64 = 1.5
    var v3 uint32 = 2147483647
    var v4 int16 = 42
    v5 := make(map[int]string)
    var v6 int32 = 2147483647
    fmt.Println(v6)
    v7 := []int{32768}
    for fi8 := 0; fi8 < 7; fi8++ {
        var v9 uint8 = 65
        _ = fi8
        _ = v9
    }
    if true {
        _ = len(v7)
        if len(v7) > 0 {
            v7[0] = 128
        }
    } else {
        fmt.Println(v1, v6, v6)
    }
    for fi10 := 0; fi10 < 11; fi10++ {
        var v11 int = (fi10 + 1)
        v6 -= 0
        _ = fi10
        _ = v11
    }
    if (!true) {
        for fi12 := 0; fi12 < 14; fi12++ {
            v6 = v6
            v13 := GenS1{F0: (v3 ^ v3)}
            _ = fi12
            _ = v13
        }
        v4 -= 1
        var v14 bool = false
        _ = v14
    } else {
        v6 = 342
        for fi15 := 0; fi15 < 3; fi15++ {
            v1 = 0.0
            if (false && (127 < (fi15 * (fi15 >> ((1) & 63))))) {
                for fi16 := 0; fi16 < 7; fi16++ {
                    var v17 float64 = float64(v1)
                    _ = fi16
                    _ = v17
                }
            } else {
                if true {
                    switch fi15 {
                    case -1:
                        v2 = v2
                    case 2147483647:
                        for fi18 := 0; fi18 < 15; fi18++ {
                            v7 = append(v7, (fi18 * fi18))
                            var v19 uint16 = 1
                            _ = fi18
                            _ = v19
                        }
                    }
                    if (0 >= 1) {
                        v20 := GenS1{F0: (v3 ^ 42)}
                        v21 := GenS1{F0: (v3 << ((255) & 63))}
                        fmt.Println(v6, v3, fi15)
                        _ = v20
                        _ = v21
                    }
                } else {
                    v3 = 2147483648
                    v4 = v4
                    v1 = 3.14159
                }
                if (fi15 >= 9223372036854775807) {
                    for fi22 := 0; fi22 < 17; fi22++ {
                        if ((("a" + "a") + "foo bar") <= "a") {
                            var v23 uint64 = 955
                            fmt.Println(v4, fi15, v6)
                            var v24 uint8 = 1
                            _ = v23
                            _ = v24
                        } else {
                            v6 = int32(v3)
                            v25 := GenS2{F0: 42, F1: 111}
                            _ = v25
                        }
                        _ = fi22
                    }
                    if false {
                        _ = len(v5)
                        var v26 int32 = 42
                        _ = v26
                    } else {
                        v27 := GenS2{F0: v4, F1: 1}
                        if true {
                            if (!(2147483647 > (v6 + -1))) {
                                switch 32767 {
                                case 0:
                                    v28 := GenS1{F0: v3}
                                    _ = v28
                                case 9223372036854775807:
                                    if false {
                                        for fi29 := 0; fi29 < 3; fi29++ {
                                            if (int64(fi29) <= 0) {
                                                switch "\n\t" {
                                                case "a":
                                                    v4 += (v4 ^ 32767)
                                                default:
                                                    delete(v5, (fi15 ^ fi15))
                                                }
                                                var v30 string = ("" + "")
                                                var v31 int8 = -118
                                                _ = v30
                                                _ = v31
                                            } else {
                                                _ = len(v5)
                                            }
                                            if (!true) {
                                                switch (fi15 * 0) {
                                                case 32768:
                                                    var v32 int32 = 378
                                                    _ = v32
                                                }
                                                var v33 bool = true
                                                _ = v33
                                            } else {
                                                fmt.Println(v4, v4, v2)
                                                if true {
                                                    var v34 string = "hello"
                                                    _ = v34
                                                } else {
                                                    var v35 int8 = 1
                                                    v4 += (^(v4))
                                                    switch fi15 {
                                                    case 127:
                                                        v3 += 1
                                                        var v36 uint64 = 2
                                                        _ = v36
                                                    case -72:
                                                        switch "foo bar" {
                                                        case "foo bar":
                                                            delete(v5, fi15)
                                                        case "a":
                                                            if (("hello" + (("foo bar" + ("\n\t" + (("" + "\n\t") + "\n\t"))) + (("" + "a") + "hello"))) != ("" + "foo bar")) {
                                                                var v37 uint16 = 1
                                                                v5[1] = (("foo bar" + ("\n\t" + ((("foo bar" + ((("a" + ("\n\t" + "\n\t")) + ((("\n\t" + "a") + "a") + ("foo bar" + "\n\t"))) + "hello")) + "") + "foo bar"))) + "foo bar")
                                                                var v38 int = (fi15 >> ((649) & 63))
                                                                fmt.Println(fi15)
                                                                _ = v37
                                                                _ = v38
                                                            }
                                                        case "hello":
                                                            v1 = float32(((1.5 - (((((-2.25 / (v2 / 1e10)) / 3.14159) - v2) / float64(v1)) / 1.0)) - v2))
                                                        default:
                                                            var v39 uint32 = (v3 % ((v3) | 1))
                                                            _ = v39
                                                        }
                                                        if (false && (("" + "foo bar") > "\n\t")) {
                                                            if ("foo bar" > ("foo bar" + "foo bar")) {
                                                                delete(v5, fi15)
                                                                fmt.Println(v3, v6)
                                                            }
                                                            if true {
                                                                v1 = 3.14159
                                                            } else {
                                                                if (5 > (v35 + 100)) {
                                                                    v6 = v6
                                                                }
                                                                if len(v7) > 0 {
                                                                    v7[0] = (fi29 << ((1) & 63))
                                                                }
                                                            }
                                                        }
                                                        if len(v7) > 0 {
                                                            v7[0] = (fi29 - 9223372036854775807)
                                                        }
                                                        var v40 float32 = (float32(0.1) / v1)
                                                        _ = v40
                                                    case 28:
                                                        if true {
                                                            fmt.Println(v6)
                                                        } else {
                                                            var v41 uint16 = 724
                                                            _ = v41
                                                        }
                                                        if false {
                                                            switch (^(fi15)) {
                                                            case -798:
                                                                v42 := v5[(fi29 ^ (fi29 - fi15))]
                                                                _ = v42
                                                            case 32768:
                                                                v5[(fi29 / ((127) | 1))] = "a"
                                                            default:
                                                                if (true && false) {
                                                                    if len(v7) > 0 {
                                                                        v7[0] = (-(fi29))
                                                                    }
                                                                    v4 = 42
                                                                    if true {
                                                                        v43 := GenS1{F0: (v3 ^ v3)}
                                                                        v44 := GenS3{F0: v6}
                                                                        _ = v43
                                                                        _ = v44
                                                                    }
                                                                } else {
                                                                    v2 += 1e10
                                                                    switch 80 {
                                                                    case 453:
                                                                        if true {
                                                                            v35 -= v35
                                                                        } else {
                                                                            v1 = (v1 - float32((float64(v1) - v2)))
                                                                        }
                                                                    case 512:
                                                                        switch 42 {
                                                                        case 1:
                                                                            v2 = 0.1
                                                                        case 646:
                                                                            v3 = v3
                                                                        case 128:
                                                                            v5[-9223372036854775808] = ("foo bar" + "\n\t")
                                                                        }
                                                                        var v45 uint32 = (^(v3))
                                                                        v4 = v4
                                                                        _ = v45
                                                                    default:
                                                                        var v46 uint16 = 32767
                                                                        _ = v46
                                                                    }
                                                                    fmt.Println(v6, v6)
                                                                }
                                                            }
                                                            var v47 int16 = v4
                                                            if true {
                                                                v6 = int32(v47)
                                                                v3 += v3
                                                            }
                                                            _ = v47
                                                        } else {
                                                            _ = len(v7)
                                                            v6 += (^(v6))
                                                        }
                                                        var v48 int64 = int64(fi29)
                                                        _ = v48
                                                    default:
                                                        fmt.Println(v6, v6, v35)
                                                    }
                                                    _ = v35
                                                }
                                            }
                                            _ = fi29
                                        }
                                        var v49 int8 = -128
                                        var v50 uint = 0
                                        _ = v49
                                        _ = v50
                                    } else {
                                        if false {
                                            fmt.Println(v6, v3)
                                        } else {
                                            var v51 uint8 = 127
                                            v5[(fi15 - (fi15 ^ (fi15 << ((2) & 63))))] = "hello"
                                            _ = v51
                                        }
                                        if true {
                                            for fi52 := 0; fi52 < 4; fi52++ {
                                                var v53 bool = (true && true)
                                                switch "hello" {
                                                case "":
                                                    var v54 int8 = 42
                                                    _ = v54
                                                case "hello":
                                                    v55 := GenS4{F0: v3, F1: 1}
                                                    _ = v55
                                                case "foo bar":
                                                    fmt.Println(fi15, v53, v53)
                                                }
                                                var v56 int16 = (v4 * v4)
                                                _ = fi52
                                                _ = v53
                                                _ = v56
                                            }
                                        } else {
                                            var v57 uint32 = 2147483647
                                            _ = v57
                                        }
                                    }
                                }
                                v58 := GenS3{F0: v6}
                                for fi59 := 0; fi59 < 1; fi59++ {
                                    var v60 uint = 2
                                    var v61 int32 = v6
                                    if ((!((2 <= 157) || (fi15 <= 2147483647))) && ((1 == 231) && false)) {
                                        fmt.Println(v60, v61, v3)
                                        v4 += 0
                                    } else {
                                        var v62 int8 = 0
                                        var v63 float32 = (v1 + v1)
                                        var v64 int32 = (v6 * (v61 * 0))
                                        _ = v62
                                        _ = v63
                                        _ = v64
                                    }
                                    _ = fi59
                                    _ = v60
                                    _ = v61
                                }
                                _ = v58
                            } else {
                                var v65 uint32 = v3
                                var v66 uint8 = 178
                                switch "foo bar" {
                                case "\n\t":
                                    var v67 uint16 = 0
                                    _ = v67
                                default:
                                    if false {
                                        v2 += (((-2.25 + 1.5) + v2) - 3.14159)
                                        if false {
                                            for fi68 := 0; fi68 < 12; fi68++ {
                                                v69 := GenS3{F0: (v6 * 2147483647)}
                                                if false {
                                                    v70 := GenS5{F0: (!(1e10 < float32(v2))), F1: 0}
                                                    var v71 int = (fi68 & 32768)
                                                    _ = v70
                                                    _ = v71
                                                }
                                                var v72 float32 = v1
                                                var v73 uint64 = 1
                                                _ = fi68
                                                _ = v69
                                                _ = v72
                                                _ = v73
                                            }
                                            v3 -= v65
                                            var v74 int = fi15
                                            _ = v74
                                        }
                                        v2 = 0.0
                                    }
                                }
                                _ = v65
                                _ = v66
                            }
                        } else {
                            v4 = (v4 / ((v4) | 1))
                            var v75 int32 = 2147483647
                            var v76 int16 = (v4 & 1000)
                            _ = v75
                            _ = v76
                        }
                        var v77 int16 = (v4 - (v4 ^ v4))
                        _ = v27
                        _ = v77
                    }
                } else {
                    if false {
                        v4 = 1
                        if (0 >= fi15) {
                            for fi78 := 0; fi78 < 2; fi78++ {
                                var v79 uint64 = 2
                                _ = fi78
                                _ = v79
                            }
                        }
                    }
                    var v80 float64 = -2.25
                    _ = v80
                }
            }
            _ = fi15
        }
    }
    v3 = v3
    var v81 int64 = 502
    var v82 int32 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v81:", v81)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v81
    _ = v82
}
