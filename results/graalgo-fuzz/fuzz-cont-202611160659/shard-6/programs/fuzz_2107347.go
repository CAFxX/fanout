package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 string; F1 int64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int32; F1 uint32 }
type GenS7 struct { F0 bool; F1 uint8; F2 int32 }

func main() {
    v1 := []int{9223372036854775807}
    var v2 int32 = -139
    v3 := []uint32{832, 2147483648, 4294967295, 42}
    var v4 int8 = -10
    v5 := GenS1{F0: ("foo bar" + "")}
    v6 := GenS1{F0: "\n\t"}
    v2 += (v2 + v2)
    v4 -= -1
    if (955 < 42) {
        var v7 int64 = 0
        switch 32768 {
        case 1:
            if (-117 >= (v4 & v4)) {
                if ((true || (0.1 <= (-2.25 * float32(1e10)))) || false) {
                    _ = len(v1)
                    v7 += 1
                } else {
                    v7 -= v7
                    switch 0 {
                    case -9223372036854775808:
                        for fi8 := 0; fi8 < 17; fi8++ {
                            var v9 uint = 42
                            v10 := GenS2{F0: 2, F1: v4}
                            var v11 int16 = 0
                            _ = fi8
                            _ = v10
                            _ = v11
                            _ = v9
                        }
                    case 628:
                        if false {
                            for fi12 := 0; fi12 < 5; fi12++ {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                var v13 uint64 = 2
                                _ = fi12
                                _ = v13
                            }
                        }
                    }
                    v2 = (v2 / ((v2) | 1))
                }
                for fi14 := 0; fi14 < 17; fi14++ {
                    for fi15 := 0; fi15 < 9; fi15++ {
                        var v16 string = ("hello" + "a")
                        var v17 int32 = -799
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                    for fi18 := 0; fi18 < 8; fi18++ {
                        if true {
                            var v19 int32 = (v2 ^ v2)
                            switch fi18 {
                            case 32768:
                                v19 = v19
                            default:
                                var v20 bool = true
                                if (!v20) {
                                    var v21 string = "foo bar"
                                    var v22 int = (fi14 % ((fi14) | 1))
                                    v23 := GenS3{F0: "a", F1: -9223372036854775808}
                                    v4 += v4
                                    var v24 bool = (42 < 260)
                                    var v25 uint8 = 66
                                    _ = v21
                                    _ = v22
                                    _ = v23
                                    _ = v24
                                    _ = v25
                                } else {
                                    v7 += v7
                                    if (32 > 0) {
                                        _ = len(v1)
                                        var v26 int8 = v4
                                        _ = v26
                                    }
                                }
                                _ = v20
                            }
                            fmt.Println(fi18)
                            _ = v19
                        } else {
                            _ = len(v1)
                        }
                        v2 = -1
                        if (2 > 2147483648) {
                            v27 := GenS1{F0: ""}
                            var v28 uint32 = 2147483647
                            v7 += (-(v7))
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = v27
                            _ = v28
                        }
                        _ = fi18
                    }
                    var v29 int32 = (v2 + 371)
                    _ = fi14
                    _ = v29
                }
                var v30 float64 = 0.1
                _ = v30
            } else {
                for fi31 := 0; fi31 < 4; fi31++ {
                    for fi32 := 0; fi32 < 9; fi32++ {
                        var v33 uint64 = 563
                        _ = fi32
                        _ = v33
                    }
                    var v34 int = (fi31 & fi31)
                    if false {
                        v35 := GenS1{F0: "a"}
                        _ = v35
                    }
                    _ = fi31
                    _ = v34
                }
                var v36 bool = true
                _ = v36
            }
        default:
            v4 += (v4 % ((100) | 1))
        }
        if false {
            for fi37 := 0; fi37 < 12; fi37++ {
                var v38 bool = false
                var v39 uint8 = 149
                _ = fi37
                _ = v38
                _ = v39
            }
        } else {
            switch "a" {
            case "a":
                var v40 int = 1
                _ = v40
            case "foo bar":
                fmt.Println(v7, v4)
            case "hello":
                v7 = (v7 - (v7 | (v7 + v7)))
            }
            fmt.Println(v4, v4)
        }
        _ = v7
    }
    v41 := GenS3{F0: (("hello" + "") + (("foo bar" + "a") + "a")), F1: -1}
    if len(v3) > 0 {
        v3[0] = 189
    }
    var v42 float64 = 0.1
    v43 := GenS3{F0: ("hello" + "hello"), F1: -1}
    var v44 string = "foo bar"
    v45 := GenS4{F0: true}
    var v46 float64 = ((v42 - (v42 / 1.0)) + 3.14159)
    switch 4294967295 {
    case -1:
        switch v44 {
        case "\n\t":
            v42 = v46
        case "":
            v44 = (v44 + (v44 + "\n\t"))
        }
    case -642:
        v47 := GenS3{F0: v44, F1: int64(v4)}
        _ = v47
    case 9223372036854775807:
        var v48 int32 = (v2 - 2147483647)
        _ = v48
    }
    v3 = append(v3, 183)
    if (v4 < -103) {
        if true {
            if (2 < 0) {
                v46 = v46
                for fi49 := 0; fi49 < 20; fi49++ {
                    v50 := GenS2{F0: 228, F1: v4}
                    switch int(v4) {
                    case 0:
                        v46 -= 0.0
                    case -159:
                        v51 := GenS5{F0: 347}
                        _ = v51
                    case 127:
                        var v52 int32 = v2
                        _ = v52
                    default:
                        var v53 bool = true
                        _ = v53
                    }
                    if len(v3) > 0 {
                        v3[0] = 696
                    }
                    _ = fi49
                    _ = v50
                }
            }
            switch (v44 + (v44 + v44)) {
            case "foo bar":
                v54 := GenS2{F0: 42, F1: 0}
                _ = v54
            case "":
                var v55 uint16 = 778
                _ = v55
            case "\n\t":
                switch v44 {
                case "hello":
                    v3 = append(v3, 4294967295)
                case "foo bar":
                    if false {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                default:
                    if false {
                        var v56 int64 = -9223372036854775808
                        switch v44 {
                        case "":
                            for fi57 := 0; fi57 < 7; fi57++ {
                                switch v44 {
                                case "":
                                    fmt.Println(v2, v4, v44)
                                case "a":
                                    v58 := GenS6{F0: v2, F1: 265}
                                    _ = v58
                                case "foo bar":
                                    switch fi57 {
                                    case 4294967295:
                                        var v59 int = (^(fi57))
                                        _ = v59
                                    case -100:
                                        for fi60 := 0; fi60 < 18; fi60++ {
                                            v46 = v42
                                            v42 -= v46
                                            _ = fi60
                                        }
                                    case 32767:
                                        v61 := GenS6{F0: (v2 - (v2 - v2)), F1: 2}
                                        _ = v61
                                    }
                                }
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                fmt.Println(v46)
                                _ = fi57
                            }
                        case "foo bar":
                            if false {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            } else {
                                fmt.Println(v44)
                                var v62 int64 = 9223372036854775807
                                _ = v62
                            }
                        }
                        _ = v56
                    } else {
                        switch "a" {
                        case "a":
                            var v63 int = int(v2)
                            _ = v63
                        case "":
                            var v64 uint32 = 2
                            _ = v64
                        }
                        var v65 uint8 = 1
                        for fi66 := 0; fi66 < 1; fi66++ {
                            v67 := GenS7{F0: false, F1: v65, F2: -2147483648}
                            if len(v1) > 0 {
                                v1[0] = (fi66 * (fi66 + 1))
                            }
                            _ = fi66
                            _ = v67
                        }
                        _ = v65
                    }
                }
            default:
                if len(v1) > 0 {
                    v1[0] = 1
                }
            }
            v3 = append(v3, 920)
        }
        var v68 uint = 473
        fmt.Println(v68)
        _ = v68
    }
    if len(v1) > 0 {
        v1[0] = 42
    }
    var v69 uint8 = 42
    if true {
        v2 = v2
    } else {
        for fi70 := 0; fi70 < 20; fi70++ {
            v44 = v44
            _ = fi70
        }
        var v71 int64 = -1
        var v72 uint16 = 320
        _ = v71
        _ = v72
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    fmt.Println("v46:", v46)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v5
    _ = v6
    _ = v69
}
