package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8; F1 string; F2 uint8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int; F1 int16 }
type GenS7 struct { F0 int64 }
func genFn0() bool {
    for fi5 := 0; fi5 < 8; fi5++ {
        switch ("a" + "foo bar") {
        case "a":
            v6 := GenS1{F0: -1}
            _ = v6
        case "foo bar":
            v7 := GenS1{F0: 9223372036854775807}
            _ = v7
        case "hello":
            for fi8 := 0; fi8 < 4; fi8++ {
                v9 := GenS1{F0: -1}
                if (("a" + "hello") != "\n\t") {
                    if false {
                        v10 := GenS2{F0: 1}
                        _ = v10
                    }
                    v11 := GenS1{F0: -476}
                    _ = v11
                } else {
                    var v12 bool = true
                    _ = v12
                }
                v13 := []int{fi8, 32767, (-(fi5)), (fi8 / (((fi8 / ((fi8) | 1))) | 1))}
                _ = fi8
                _ = v13
                _ = v9
            }
        default:
            var v14 uint = 0
            _ = v14
        }
        v15 := []int{(fi5 + fi5), (fi5 % (((fi5 / ((fi5) | 1))) | 1))}
        if false {
            if len(v15) > 0 {
                v15[0] = 676
            }
            v16 := GenS3{F0: 1, F1: "hello", F2: 124}
            if (!((32767 <= 1000) || false)) {
                if ((1e10 * (0.0 + 1e10)) > (1.5 - 1.5)) {
                    fmt.Println(fi5)
                    var v17 uint64 = 4294967295
                    _ = v17
                }
            }
            _ = v16
        }
        _ = fi5
        _ = v15
    }
    var v18 uint64 = 1
    if ((!false) && (0.0 >= 0.0)) {
        if (("foo bar" + "\n\t") <= "foo bar") {
            fmt.Println(v18)
        }
        v19 := []int{-2147483648, -37}
        var v20 uint64 = v18
        _ = v19
        _ = v20
    } else {
        v18 += (^(v18))
        for fi21 := 0; fi21 < 19; fi21++ {
            v18 += v18
            v18 -= (v18 * v18)
            _ = fi21
        }
    }
    var v22 float32 = 1e10
    v23 := GenS2{F0: uint(v18)}
    _ = v18
    _ = v22
    _ = v23
    return true
}

func main() {
    var v1 int64 = 1
    var v2 float32 = 0.0
    v3 := make(map[int]uint32)
    v4 := []int{}
    var v24 uint = 1
    v2 = 0.0
    switch "a" {
    case "\n\t":
        for fi25 := 0; fi25 < 6; fi25++ {
            switch -435 {
            case 194:
                var v26 uint16 = 1
                _ = v26
            case 32768:
                v2 -= 0.1
            case 127:
                if false {
                    fmt.Println(v2)
                }
            }
            switch fi25 {
            case 2147483647:
                v1 += 1
            case -496:
                if (1 > 32767) {
                    if false {
                        v1 += v1
                        var v27 uint64 = 18446744073709551615
                        _ = v27
                    } else {
                        for fi28 := 0; fi28 < 16; fi28++ {
                            v1 += -992
                            v29 := GenS4{F0: v24}
                            if false {
                                var v30 string = ("foo bar" + "hello")
                                if (((-(fi28)) > -262) && (true || (!false))) {
                                    switch 42 {
                                    case 32768:
                                        if true {
                                            v30 = v30
                                        } else {
                                            v24 += (v24 % (((v24 * (v24 / ((v24) | 1)))) | 1))
                                            _ = len(v3)
                                            var v31 bool = true
                                            var v32 uint16 = 65535
                                            switch int(v32) {
                                            case -123:
                                                v24 += uint(v32)
                                            case 525:
                                                v33 := v3[128]
                                                _ = v33
                                            default:
                                                v1 += -248
                                            }
                                            _ = v31
                                            _ = v32
                                        }
                                    case -396:
                                        v34 := GenS4{F0: 1}
                                        _ = v34
                                    case 803:
                                        switch v30 {
                                        case "\n\t":
                                            v30 = "\n\t"
                                        default:
                                            var v35 uint32 = 299
                                            _ = v35
                                        }
                                    default:
                                        if (65535 <= 4294967295) {
                                            if (24 <= 209) {
                                                v2 = v2
                                                switch "hello" {
                                                case "":
                                                    v2 -= (((((float32(0.0) - 1.5) + (0.0 * ((v2 - 0.1) * float32((-2.25 - 3.14159))))) / v2) + ((v2 / ((((1.5 * v2) * 1e10) / 0.1) - (v2 - ((3.14159 / (v2 * 0.0)) * 0.1)))) * v2)) + v2)
                                                    v36 := GenS4{F0: 42}
                                                    _ = v36
                                                case "\n\t":
                                                    var v37 bool = true
                                                    _ = v37
                                                }
                                                switch 1 {
                                                case 32767:
                                                    fmt.Println(v30, v30, v1)
                                                case 32768:
                                                    v38 := GenS5{F0: 0}
                                                    _ = v38
                                                case -2147483648:
                                                    var v39 uint32 = 1
                                                    _ = v39
                                                default:
                                                    var v40 string = "foo bar"
                                                    _ = v40
                                                }
                                                v30 = v30
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                            }
                                            v2 -= (((v2 / v2) - 3.14159) + (v2 - 1e10))
                                        } else {
                                            fmt.Println(fi25, v24)
                                            fmt.Println(v24)
                                        }
                                    }
                                } else {
                                    fmt.Println(v24)
                                }
                                var v41 uint32 = 376
                                _ = v30
                                _ = v41
                            } else {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                            _ = fi28
                            _ = v29
                        }
                        for fi42 := 0; fi42 < 10; fi42++ {
                            if ((fi25 % ((44) | 1)) < fi42) {
                                if true {
                                    var v43 int16 = 1000
                                    v44 := GenS4{F0: 946}
                                    _ = v43
                                    _ = v44
                                } else {
                                    var v45 float32 = (v2 + (3.14159 + 1.5))
                                    v2 += float32(0.1)
                                    _ = v45
                                }
                                if ((!true) || (65535 != 42)) {
                                    var v46 int = (fi25 ^ 9223372036854775807)
                                    _ = v46
                                }
                            } else {
                                var v47 string = "foo bar"
                                fmt.Println(v1, v47, fi25)
                                var v48 int = fi42
                                _ = v47
                                _ = v48
                            }
                            v1 -= (v1 >> ((v24) & 63))
                            fmt.Println(v1, fi25, fi42)
                            if false {
                                fmt.Println(v2)
                                var v49 uint32 = 1
                                switch (fi42 << (((v24 & v24)) & 63)) {
                                case 32767:
                                    if (2 > 42) {
                                        v50 := GenS1{F0: v1}
                                        if true {
                                            v51 := GenS5{F0: (^(v1))}
                                            _ = v51
                                        }
                                        if len(v4) > 0 {
                                            v4[0] = 358
                                        }
                                        switch "foo bar" {
                                        case "hello":
                                            v52 := v3[-300]
                                            _ = v52
                                        case "":
                                            fmt.Println(fi42, v24)
                                        case "foo bar":
                                            v53 := v3[(fi25 / ((fi42) | 1))]
                                            _ = v53
                                        default:
                                            fmt.Println(fi42)
                                        }
                                        var v54 uint64 = 0
                                        if true {
                                            var v55 bool = true
                                            switch "" {
                                            case "hello":
                                                var v56 int8 = -1
                                                _ = v56
                                            case "\n\t":
                                                v57 := GenS6{F0: (fi42 % (((fi25 + (fi42 | (fi42 / ((597) | 1))))) | 1)), F1: -32768}
                                                _ = v57
                                            case "foo bar":
                                                var v58 bool = (true || v55)
                                                _ = v58
                                            default:
                                                switch "" {
                                                case "a":
                                                    v59 := GenS6{F0: (^(fi25)), F1: 1}
                                                    v3[9223372036854775807] = v49
                                                    if v55 {
                                                        var v60 float32 = ((float32(1.5) - 1.5) / ((0.1 / v2) - v2))
                                                        _ = v60
                                                    }
                                                    _ = v59
                                                }
                                                v4 = append(v4, -1)
                                                var v61 bool = v55
                                                _ = v61
                                            }
                                            _ = v55
                                        }
                                        var v62 float32 = float32(1.5)
                                        _ = v50
                                        _ = v54
                                        _ = v62
                                    } else {
                                        if false {
                                            var v63 bool = true
                                            var v64 float64 = (1.5 / float64(v2))
                                            _ = v63
                                            _ = v64
                                        } else {
                                            var v65 uint32 = v49
                                            v2 += (v2 * v2)
                                            _ = v65
                                        }
                                        v66 := GenS4{F0: (^(v24))}
                                        _ = v66
                                    }
                                }
                                v49 = 2147483647
                                _ = v49
                            }
                            switch fi25 {
                            case -652:
                                v3[0] = 559
                                fmt.Println(fi25, fi42, fi25)
                                v67 := GenS6{F0: 2147483647, F1: 32767}
                                _ = v67
                            default:
                                v1 = 42
                            }
                            _ = fi42
                        }
                        v68 := GenS6{F0: fi25, F1: 261}
                        _ = v68
                    }
                }
            case 42:
                for fi69 := 0; fi69 < 13; fi69++ {
                    switch ("\n\t" + ("" + "a")) {
                    case "foo bar":
                        fmt.Println(v1, v1, v24)
                    case "hello":
                        if len(v4) > 0 {
                            v4[0] = fi25
                        }
                    case "\n\t":
                        var v70 uint32 = 653
                        _ = v70
                    }
                    _ = fi69
                }
            }
            _ = fi25
        }
    default:
        fmt.Println(v1, v2, v1)
    }
    v2 = 3.14159
    v1 += v1
    for fi71 := 0; fi71 < 11; fi71++ {
        _ = len(v3)
        _ = fi71
    }
    if true {
        var v72 uint8 = 180
        _ = v72
    } else {
        var v73 int8 = -1
        v2 -= v2
        _ = v73
    }
    var v74 uint64 = 338
    v75 := v3[42]
    fmt.Println(v2, v24)
    v76 := GenS7{F0: v1}
    v77 := GenS2{F0: v24}
    var v78 int = 32767
    v78 += v78
    fmt.Println(v74)
    for fi79 := 0; fi79 < 1; fi79++ {
        if ((fi79 - (^(v78))) > v78) {
            var v80 uint32 = (v75 % ((v75) | 1))
            _ = len(v3)
            fmt.Println(v2, v78)
            _ = v80
        } else {
            v81 := GenS7{F0: -9223372036854775808}
            _ = v81
        }
        _ = len(v4)
        var v82 int64 = v1
        _ = fi79
        _ = v82
    }
    fmt.Println(v1, v75)
    fmt.Println(v75, v2, v74)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v74:", v74)
    fmt.Println("v75:", v75)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v74
    _ = v75
    _ = v76
    _ = v77
    _ = v78
}
