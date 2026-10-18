package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32; F1 int; F2 bool }
type GenS3 struct { F0 int16; F1 int32; F2 uint64 }
type GenS4 struct { F0 uint16; F1 int32; F2 uint64 }
type GenS5 struct { F0 uint16; F1 string }

func main() {
    var v1 uint8 = 127
    var v2 uint8 = 166
    var v3 int = 109
    var v4 int = (v3 ^ v3)
    var v5 bool = false
    v6 := make(map[uint32]string)
    if v5 {
        v2 -= v1
        _ = len(v6)
    } else {
        var v7 float64 = (3.14159 + 0.0)
        _ = v7
    }
    v8 := []int{(v4 & 42), 9223372036854775807, (v3 - (v3 + v4))}
    v9 := GenS1{F0: 222}
    fmt.Println(v5)
    v2 -= v2
    fmt.Println(v3, v3, v5)
    var v10 int64 = -512
    _ = len(v6)
    switch "hello" {
    case "hello":
        if true {
            if (!(65535 == 1)) {
                v2 = v2
            }
            var v11 int16 = 464
            _ = v11
        } else {
            _ = len(v6)
            if v5 {
                switch "" {
                case "":
                    if v5 {
                        _ = len(v6)
                        var v12 uint16 = 67
                        var v13 uint32 = 348
                        _ = v12
                        _ = v13
                    } else {
                        var v14 float32 = 0.0
                        var v15 uint = 65535
                        _ = v14
                        _ = v15
                    }
                case "foo bar":
                    switch "hello" {
                    case "foo bar":
                        _ = len(v6)
                    }
                default:
                    if v5 {
                        v16 := GenS2{F0: 2147483648, F1: v3, F2: (!(4294967295 > 1))}
                        var v17 float64 = 0.0
                        _ = v16
                        _ = v17
                    } else {
                        var v18 uint8 = (v1 | v1)
                        switch "" {
                        case "":
                            v19 := GenS1{F0: 4}
                            _ = v19
                        }
                        v20 := GenS1{F0: 0}
                        _ = v18
                        _ = v20
                    }
                }
                for fi21 := 0; fi21 < 10; fi21++ {
                    v3 -= (v4 ^ 32768)
                    _ = fi21
                }
            }
            if v5 {
                v22 := GenS3{F0: -1, F1: -496, F2: 42}
                _ = v22
            } else {
                v5 = ("\n\t" < "\n\t")
            }
        }
    default:
        var v23 bool = false
        _ = v23
    }
    var v24 uint16 = 0
    var v25 uint64 = 670
    var v26 uint64 = 740
    var v27 bool = (v5 && (false && false))
    for fi28 := 0; fi28 < 15; fi28++ {
        if (!v27) {
            v29 := GenS4{F0: v24, F1: 42, F2: v26}
            _ = len(v8)
            if false {
                _ = len(v8)
                if v27 {
                    v5 = false
                    v5 = true
                    fmt.Println(fi28)
                } else {
                    var v30 uint = 2
                    _ = v30
                }
            }
            _ = v29
        }
        var v31 uint32 = 2147483648
        switch ("" + ("\n\t" + "\n\t")) {
        case "foo bar":
            _ = len(v8)
        case "\n\t":
            for fi32 := 0; fi32 < 4; fi32++ {
                switch fi32 {
                case -9223372036854775808:
                    v27 = (!(!(((1e10 / (0.0 + -2.25)) <= 3.14159) && true)))
                case 4294967295:
                    switch v3 {
                    case -838:
                        v25 = (v25 % (((v25 - uint64(v31))) | 1))
                    case 1:
                        v1 = uint8(v31)
                    default:
                        switch fi32 {
                        case 1:
                            var v33 bool = v5
                            _ = v33
                        case 127:
                            var v34 int16 = 1
                            _ = v34
                        case 42:
                            v35 := v6[(^(v31))]
                            _ = v35
                        }
                    }
                case 2147483647:
                    var v36 uint8 = (v2 ^ (v1 | v2))
                    _ = v36
                }
                fmt.Println(v1, v24)
                v8 = append(v8, 128)
                v6[(v31 - 2)] = "hello"
                var v37 uint32 = (v31 & v31)
                _ = fi32
                _ = v37
            }
        default:
            v38 := v6[v31]
            _ = v38
        }
        _ = fi28
        _ = v31
    }
    if (!(("a" >= "foo bar") || v5)) {
        var v39 int16 = 0
        _ = v39
    } else {
        for fi40 := 0; fi40 < 7; fi40++ {
            if false {
                v27 = false
                if v27 {
                    var v41 uint8 = 255
                    _ = v41
                }
                _ = len(v6)
            } else {
                if ((v24 - v24) != uint16(v1)) {
                    for fi42 := 0; fi42 < 20; fi42++ {
                        v4 = 1
                        v43 := GenS3{F0: 32767, F1: -2147483648, F2: (v25 * (v25 + (v25 * v25)))}
                        fmt.Println(v2)
                        if (42 >= v4) {
                            v44 := GenS4{F0: 32767, F1: 2147483647, F2: v26}
                            var v45 int = (fi42 & (v3 % (((fi42 % ((v4) | 1))) | 1)))
                            var v46 bool = v27
                            _ = v44
                            _ = v45
                            _ = v46
                        } else {
                            if v5 {
                                var v47 int8 = 3
                                var v48 string = ""
                                _ = v47
                                _ = v48
                            }
                            v8 = append(v8, fi40)
                        }
                        v8 = append(v8, 4294967295)
                        _ = fi42
                        _ = v43
                    }
                    switch 323 {
                    case 2147483647:
                        if ((3.14159 <= 1e10) || v27) {
                            for fi49 := 0; fi49 < 4; fi49++ {
                                v50 := GenS5{F0: uint16(v1), F1: ((("" + ("\n\t" + "a")) + "a") + "hello")}
                                v6[0] = (((("a" + ("a" + "a")) + "") + (("a" + "\n\t") + "hello")) + ("" + (("hello" + "a") + "")))
                                _ = fi49
                                _ = v50
                            }
                            var v51 int16 = 637
                            _ = v51
                        }
                    case -2147483648:
                        for fi52 := 0; fi52 < 1; fi52++ {
                            var v53 uint64 = 104
                            var v54 int64 = (v10 % (((v10 - (v10 ^ v10))) | 1))
                            _ = fi52
                            _ = v53
                            _ = v54
                        }
                    case 42:
                        if true {
                            var v55 bool = v5
                            var v56 int64 = v10
                            var v57 int8 = -128
                            _ = v55
                            _ = v56
                            _ = v57
                        }
                    default:
                        var v58 uint32 = 4294967295
                        _ = v58
                    }
                } else {
                    for fi59 := 0; fi59 < 15; fi59++ {
                        v26 = v26
                        if (!v5) {
                            v4 += fi59
                            v27 = v27
                            if (!(702 == (v24 % ((v24) | 1)))) {
                                v6[42] = "hello"
                                var v60 float64 = 0.0
                                _ = v60
                            } else {
                                v1 -= v1
                                fmt.Println(v25)
                            }
                        } else {
                            v1 = (v1 % ((v1) | 1))
                            var v61 uint32 = 0
                            _ = v61
                        }
                        _ = len(v8)
                        _ = fi59
                    }
                    if v27 {
                        var v62 float32 = -2.25
                        v8 = append(v8, (v4 << ((uint(v24)) & 63)))
                        _ = v62
                    }
                }
                v24 = 2
                for fi63 := 0; fi63 < 18; fi63++ {
                    v64 := v6[955]
                    _ = fi63
                    _ = v64
                }
            }
            for fi65 := 0; fi65 < 13; fi65++ {
                delete(v6, 483)
                _ = fi65
            }
            v6[4294967295] = ("foo bar" + "a")
            _ = fi40
        }
    }
    var v66 int64 = v10
    if false {
        var v67 int16 = -32768
        v68 := GenS1{F0: -2147483648}
        v67 -= (v67 & (v67 + int16(v26)))
        _ = v67
        _ = v68
    } else {
        switch 1 {
        case 305:
            var v69 bool = ((v3 - 632) >= 32767)
            _ = v69
        default:
            v70 := GenS3{F0: 1000, F1: -293, F2: 4294967295}
            _ = v70
        }
        v3 = v4
        fmt.Println(v10, v26, v27)
    }
    fmt.Println(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v66:", v66)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v66
    _ = v8
    _ = v9
}
