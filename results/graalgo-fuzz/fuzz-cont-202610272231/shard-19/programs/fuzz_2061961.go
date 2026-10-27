package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 int16; F1 bool }
type GenS4 struct { F0 int; F1 uint8 }
type GenS5 struct { F0 uint; F1 int32 }
type GenS6 struct { F0 uint; F1 uint64; F2 uint64 }
type GenS7 struct { F0 uint16; F1 int8; F2 int8 }
type GenS8 struct { F0 uint; F1 string; F2 uint16 }

func main() {
    var v1 uint64 = 169
    v2 := make(map[uint32]uint32)
    var v3 int64 = 778
    var v4 int16 = 103
    var v5 uint = 0
    v6 := []int{553}
    if ((127 > 0) && false) {
        var v7 uint64 = (^(v1))
        v8 := GenS1{F0: (v3 ^ v3), F1: 0}
        _ = v7
        _ = v8
    } else {
        var v9 uint = 4294967295
        _ = v9
    }
    for fi10 := 0; fi10 < 10; fi10++ {
        if false {
            var v11 uint32 = 36
            _ = v11
        }
        v3 -= (v3 | (^(v3)))
        v4 += v4
        _ = fi10
    }
    delete(v2, 4294967295)
    if (137 < 0) {
        if true {
            var v12 uint8 = 127
            fmt.Println(v4, v3, v1)
            switch "\n\t" {
            case "hello":
                if false {
                    fmt.Println(v3, v5, v4)
                    if ((v1 * (^(v1))) >= (v1 % (((v1 ^ (v1 << (((v5 - 2)) & 63)))) | 1))) {
                        v1 = 868
                        var v13 int = -9223372036854775808
                        _ = v13
                    } else {
                        var v14 uint64 = (v1 - 1)
                        _ = v14
                    }
                } else {
                    fmt.Println(v3, v4, v1)
                }
            case "\n\t":
                if (112 > 1) {
                    v1 -= v1
                    v15 := GenS1{F0: (^(v3)), F1: 2147483648}
                    for fi16 := 0; fi16 < 9; fi16++ {
                        fmt.Println(v4, v12)
                        if len(v6) > 0 {
                            v6[0] = fi16
                        }
                        _ = fi16
                    }
                    _ = v15
                } else {
                    var v17 int8 = -1
                    switch "\n\t" {
                    case "a":
                        v3 += (v3 - int64(v1))
                    case "foo bar":
                        for fi18 := 0; fi18 < 3; fi18++ {
                            var v19 float64 = 0.0
                            v5 += v5
                            _ = fi18
                            _ = v19
                        }
                    case "hello":
                        if (true && ((^(v17)) >= (v17 / ((-53) | 1)))) {
                            var v20 float32 = ((-2.25 - 0.1) / 1.0)
                            fmt.Println(v3)
                            var v21 uint16 = 65535
                            _ = v20
                            _ = v21
                        }
                    default:
                        for fi22 := 0; fi22 < 1; fi22++ {
                            var v23 int64 = (v3 >> (((v5 & (v5 ^ (v5 % ((uint(v12)) | 1))))) & 63))
                            var v24 uint16 = 255
                            v1 = v1
                            _ = fi22
                            _ = v23
                            _ = v24
                        }
                    }
                    var v25 int32 = int32(v12)
                    _ = v17
                    _ = v25
                }
            case "foo bar":
                v26 := GenS1{F0: v3, F1: 2147483647}
                _ = v26
            }
            _ = v12
        }
        var v27 int64 = (-(v3))
        if (true || (false || (128 <= 127))) {
            if true {
                for fi28 := 0; fi28 < 13; fi28++ {
                    v29 := GenS2{F0: 42, F1: 80}
                    var v30 float64 = ((float64(((0.1 * 1.5) * -2.25)) + 3.14159) - -2.25)
                    for fi31 := 0; fi31 < 5; fi31++ {
                        switch "" {
                        case "":
                            v32 := GenS2{F0: -1, F1: 32767}
                            _ = v32
                        case "\n\t":
                            var v33 int32 = 190
                            _ = v33
                        case "foo bar":
                            fmt.Println(v5, v3, v27)
                        }
                        var v34 string = "foo bar"
                        v3 = (v3 - v3)
                        _ = fi31
                        _ = v34
                    }
                    _ = fi28
                    _ = v29
                    _ = v30
                }
                switch 32767 {
                case 2147483647:
                    var v35 bool = (false || (("a" + ("foo bar" + "hello")) <= "\n\t"))
                    _ = v35
                case 377:
                    var v36 uint8 = 127
                    _ = v36
                case -375:
                    fmt.Println(v1, v3, v3)
                default:
                    _ = len(v6)
                }
                v3 = v3
            } else {
                for fi37 := 0; fi37 < 2; fi37++ {
                    if ((v27 + v27) <= v3) {
                        for fi38 := 0; fi38 < 8; fi38++ {
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            v3 = (v3 << (((v5 * v5)) & 63))
                            _ = fi38
                        }
                    } else {
                        v4 = -61
                    }
                    v1 = 1
                    _ = fi37
                }
                if false {
                    var v39 uint32 = 619
                    v40 := GenS3{F0: 32767, F1: false}
                    _ = v39
                    _ = v40
                } else {
                    if true {
                        v6 = append(v6, 9223372036854775807)
                    }
                }
                switch 32767 {
                case 745:
                    v41 := GenS2{F0: -554, F1: 134}
                    _ = v41
                }
            }
            var v42 int16 = 0
            _ = v42
        }
        _ = v27
    }
    switch 4294967295 {
    case 9223372036854775807:
        if ("foo bar" < "") {
            v4 -= (v4 * v4)
            if (-28 != 100) {
                switch 4294967295 {
                case -1:
                    fmt.Println(v1, v3)
                case 9223372036854775807:
                    if true {
                        switch -760 {
                        case 771:
                            var v43 uint16 = 32767
                            _ = v43
                        case 42:
                            v44 := GenS4{F0: 1, F1: 42}
                            _ = v44
                        default:
                            switch "\n\t" {
                            case "\n\t":
                                v45 := GenS1{F0: v3, F1: 120}
                                _ = v45
                            case "foo bar":
                                v46 := GenS5{F0: (v5 + v5), F1: -953}
                                _ = v46
                            }
                        }
                    } else {
                        if (v1 == 28) {
                            v3 = v3
                            v4 = 0
                        }
                        v5 = v5
                    }
                default:
                    var v47 uint16 = 42
                    _ = v47
                }
            }
        } else {
            for fi48 := 0; fi48 < 7; fi48++ {
                for fi49 := 0; fi49 < 8; fi49++ {
                    if ((-74 >= -1) && true) {
                        v50 := GenS5{F0: 335, F1: -579}
                        _ = v50
                    }
                    v3 = v3
                    switch "hello" {
                    case "hello":
                        fmt.Println(v1)
                    case "\n\t":
                        var v51 uint16 = 65535
                        _ = v51
                    case "foo bar":
                        var v52 int32 = 2147483647
                        _ = v52
                    }
                    _ = fi49
                }
                if false {
                    if (!(!(579 <= 313))) {
                        fmt.Println(v1, fi48)
                        for fi53 := 0; fi53 < 4; fi53++ {
                            switch "foo bar" {
                            case "\n\t":
                                v54 := GenS6{F0: v5, F1: (v1 ^ v1), F2: 410}
                                _ = v54
                            case "hello":
                                fmt.Println(v5, fi53, v1)
                            case "a":
                                v6 = append(v6, -9223372036854775808)
                            }
                            v55 := GenS3{F0: v4, F1: true}
                            _ = fi53
                            _ = v55
                        }
                    } else {
                        switch (fi48 + (fi48 << ((v5) & 63))) {
                        case 0:
                            var v56 bool = (false || true)
                            _ = v56
                        case 176:
                            var v57 bool = false
                            _ = v57
                        default:
                            var v58 uint16 = 42
                            _ = v58
                        }
                        for fi59 := 0; fi59 < 1; fi59++ {
                            switch ("" + "foo bar") {
                            case "a":
                                var v60 int = fi48
                                _ = v60
                            case "\n\t":
                                var v61 uint8 = 128
                                _ = v61
                            case "foo bar":
                                fmt.Println(v3, v3, v5)
                            }
                            _ = fi59
                        }
                    }
                    _ = len(v2)
                } else {
                    var v62 uint64 = v1
                    _ = v62
                }
                _ = fi48
            }
            v4 = v4
        }
    case 127:
        if (true && (v4 != (v4 / ((v4) | 1)))) {
            v5 = v5
            v5 = (v5 | v5)
            for fi63 := 0; fi63 < 16; fi63++ {
                v1 -= (v1 | v1)
                _ = fi63
            }
        }
    default:
        fmt.Println(v5, v4, v1)
    }
    v1 = uint64(v5)
    v64 := GenS1{F0: (^(v3)), F1: 4294967295}
    fmt.Println(v1, v5, v1)
    v3 = v3
    if true {
        var v65 string = "\n\t"
        _ = v65
    } else {
        fmt.Println(v1, v3)
    }
    if ("a" > ("" + ("a" + (((("a" + "hello") + "foo bar") + (("" + ((("\n\t" + "hello") + "a") + ("a" + "hello"))) + ("foo bar" + (("" + "hello") + ("foo bar" + "a"))))) + ("foo bar" + "hello"))))) {
        v66 := v2[606]
        _ = v66
    } else {
        v4 -= v4
        v5 = (v5 % (((v5 - (v5 ^ v5))) | 1))
        if false {
            if false {
                v67 := GenS1{F0: v3, F1: 4294967295}
                _ = v67
            } else {
                switch "foo bar" {
                case "a":
                    switch ("foo bar" + "\n\t") {
                    case "hello":
                        var v68 uint64 = v1
                        _ = v68
                    case "\n\t":
                        for fi69 := 0; fi69 < 13; fi69++ {
                            v5 = (^(v5))
                            if true {
                                var v70 int32 = 1
                                _ = v70
                            }
                            var v71 uint = 873
                            _ = fi69
                            _ = v71
                        }
                    default:
                        if true {
                            var v72 float32 = 0.1
                            _ = v72
                        }
                    }
                default:
                    v73 := GenS7{F0: uint16(v5), F1: int8(v3), F2: 18}
                    _ = v73
                }
                v74 := GenS8{F0: (v5 ^ uint(v1)), F1: "foo bar", F2: 498}
                _ = v74
            }
            v1 += (v1 | (v1 ^ (v1 << ((599) & 63))))
            for fi75 := 0; fi75 < 11; fi75++ {
                var v76 uint64 = v1
                if ((^(v4)) >= (v4 * v4)) {
                    var v77 int8 = 1
                    _ = v77
                } else {
                    var v78 int32 = 42
                    _ = v78
                }
                var v79 bool = ((fi75 | (fi75 * fi75)) == fi75)
                _ = fi75
                _ = v76
                _ = v79
            }
        } else {
            if true {
                var v80 float32 = ((1e10 / float32((float64(1.5) - ((3.14159 + (float64((-2.25 - -2.25)) / 1.5)) / ((float64(1.5) / float64(float32(1.5))) + -2.25))))) + 0.1)
                v81 := GenS5{F0: (v5 / (((v5 >> ((v5) & 63))) | 1)), F1: -1}
                _ = v80
                _ = v81
            }
            var v82 uint16 = 32767
            v83 := GenS6{F0: v5, F1: 42, F2: 0}
            _ = v82
            _ = v83
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v64
}
