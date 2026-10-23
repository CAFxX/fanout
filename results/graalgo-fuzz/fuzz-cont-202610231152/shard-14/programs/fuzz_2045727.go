package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16; F1 int16; F2 bool }
type GenS3 struct { F0 int; F1 uint32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int }

func main() {
    var v1 uint64 = 4294967295
    v2 := make(map[string]bool)
    var v3 bool = false
    v4 := make(map[uint32]int)
    v5 := make(map[int]string)
    v6 := []uint32{403, 0, 65535}
    for fi7 := 0; fi7 < 18; fi7++ {
        for fi8 := 0; fi8 < 9; fi8++ {
            v9 := GenS1{F0: ""}
            if v3 {
                var v10 string = ("a" + "hello")
                _ = v10
            } else {
                v3 = v3
            }
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    var v11 float32 = 0.0
    delete(v4, 65535)
    switch ("hello" + "foo bar") {
    case "a":
        for fi12 := 0; fi12 < 11; fi12++ {
            v1 += v1
            v4[4294967295] = fi12
            v3 = (!(!(!v3)))
            _ = fi12
        }
    case "":
        delete(v5, 0)
    default:
        switch ("a" + "foo bar") {
        case "":
            v11 += (v11 * v11)
        case "foo bar":
            var v13 uint64 = v1
            _ = v13
        default:
            if len(v6) > 0 {
                v6[0] = 2147483648
            }
        }
    }
    if false {
        switch 664 {
        case -1:
            switch 194 {
            case -1:
                v4[65535] = -291
            case 1:
                var v14 float32 = v11
                _ = v14
            default:
                switch ("" + "\n\t") {
                case "hello":
                    if len(v6) > 0 {
                        v6[0] = uint32(v1)
                    }
                default:
                    for fi15 := 0; fi15 < 13; fi15++ {
                        v1 = v1
                        _ = fi15
                    }
                }
            }
        case -122:
            if false {
                if (-1 == 12) {
                    if false {
                        fmt.Println(v3)
                        fmt.Println(v1)
                        v3 = (v3 || v3)
                    }
                    v3 = v3
                    fmt.Println(v11)
                }
                for fi16 := 0; fi16 < 14; fi16++ {
                    v2["a"] = (1.5 <= (float64(float32(((-2.25 + (1.5 + 0.1)) + 0.0))) / 0.1))
                    switch ("foo bar" + ((("a" + (("hello" + "\n\t") + "a")) + "foo bar") + "hello")) {
                    case "\n\t":
                        fmt.Println(v3, v11)
                    case "a":
                        switch (("\n\t" + ("" + "hello")) + "foo bar") {
                        case "a":
                            for fi17 := 0; fi17 < 4; fi17++ {
                                v18 := GenS2{F0: 2, F1: -1, F2: (2 < 229)}
                                var v19 uint16 = 1
                                _ = fi17
                                _ = v18
                                _ = v19
                            }
                        default:
                            if (true && (!v3)) {
                                var v20 uint16 = 898
                                fmt.Println(fi16, v1, v3)
                                v21 := GenS2{F0: (v20 | (v20 ^ 42)), F1: 32767, F2: (65535 != 42)}
                                _ = v20
                                _ = v21
                            } else {
                                var v22 uint = 803
                                if (((100 == 98) || v3) && (!v3)) {
                                    if (0 != 1000) {
                                        v23 := GenS1{F0: ("" + ("foo bar" + ""))}
                                        if v3 {
                                            if (0 < 2147483647) {
                                                for fi24 := 0; fi24 < 9; fi24++ {
                                                    v11 -= ((1.5 * ((v11 / ((0.0 - ((v11 / (float32(3.14159) * v11)) / (float32(float64(v11)) / v11))) + v11)) + v11)) / 0.1)
                                                    v1 -= 1
                                                    _ = fi24
                                                }
                                                if v3 {
                                                    var v25 uint32 = 2147483648
                                                    v26 := GenS3{F0: 1, F1: 2147483647}
                                                    for fi27 := 0; fi27 < 18; fi27++ {
                                                        v22 -= 255
                                                        _ = fi27
                                                    }
                                                    _ = v25
                                                    _ = v26
                                                }
                                            }
                                            var v28 int8 = 1
                                            fmt.Println(v11)
                                            _ = v28
                                        }
                                        delete(v5, fi16)
                                        _ = v23
                                    } else {
                                        var v29 int16 = 32767
                                        var v30 bool = ((72 > 2147483647) && v3)
                                        _ = v29
                                        _ = v30
                                    }
                                } else {
                                    for fi31 := 0; fi31 < 4; fi31++ {
                                        v32 := GenS4{F0: (v22 & v22)}
                                        _ = fi31
                                        _ = v32
                                    }
                                }
                                _ = v22
                            }
                        }
                    case "hello":
                        var v33 uint64 = 2
                        _ = v33
                    }
                    _ = fi16
                }
            }
        default:
            v3 = ((v11 - ((v11 / v11) + (float32(0.0) + v11))) != v11)
        }
        fmt.Println(v11)
        fmt.Println(v1)
    } else {
        for fi34 := 0; fi34 < 9; fi34++ {
            v1 -= (v1 | v1)
            _ = len(v5)
            _ = fi34
        }
        delete(v4, 791)
    }
    var v35 uint32 = 4294967295
    v36 := GenS4{F0: 1}
    v37 := GenS5{F0: -100}
    if false {
        v38 := GenS2{F0: 2, F1: -767, F2: (!(9223372036854775807 > 9223372036854775807))}
        for fi39 := 0; fi39 < 3; fi39++ {
            v40 := GenS3{F0: (fi39 - (^(fi39))), F1: uint32(v1)}
            if ((v1 ^ v1) <= v1) {
                v11 += v11
                switch "foo bar" {
                case "a":
                    if v3 {
                        for fi41 := 0; fi41 < 6; fi41++ {
                            if v3 {
                                fmt.Println(fi41, v1, v1)
                                if true {
                                    var v42 float64 = (-2.25 + 0.0)
                                    if (65535 == v35) {
                                        _ = len(v5)
                                    }
                                    if v3 {
                                        v6 = append(v6, (^(v35)))
                                        if (!v3) {
                                            fmt.Println(v1, v1, v35)
                                            v43 := GenS3{F0: (fi39 >> ((478) & 63)), F1: v35}
                                            fmt.Println(fi41, v1)
                                            _ = v43
                                        }
                                        v42 -= 1.5
                                    }
                                    v11 = v11
                                    var v44 int16 = 0
                                    _ = v42
                                    _ = v44
                                }
                                if (1.5 >= -2.25) {
                                    fmt.Println(v3, v11, v1)
                                    _ = len(v5)
                                }
                                var v45 string = ("a" + "a")
                                _ = v45
                            } else {
                                var v46 uint = 423
                                if (!v3) {
                                    if v3 {
                                        var v47 uint64 = v1
                                        if (("foo bar" + "hello") < "hello") {
                                            fmt.Println(fi39, v35)
                                            v48 := GenS3{F0: (fi39 * int(v46)), F1: 65535}
                                            _ = v48
                                        }
                                        v11 -= 1e10
                                        _ = v47
                                    }
                                    v49 := GenS3{F0: -2147483648, F1: 65535}
                                    v50 := GenS3{F0: (fi39 * fi41), F1: v35}
                                    _ = v49
                                    _ = v50
                                } else {
                                    var v51 string = "foo bar"
                                    v35 = (v35 ^ (v35 & v35))
                                    _ = v51
                                }
                                switch fi39 {
                                case 127:
                                    var v52 int32 = 42
                                    v6 = append(v6, (v35 << ((v46) & 63)))
                                    var v53 string = "foo bar"
                                    _ = v52
                                    _ = v53
                                case 0:
                                    var v54 int16 = -1
                                    _ = v54
                                default:
                                    var v55 uint64 = v1
                                    _ = v55
                                }
                                _ = v46
                            }
                            _ = fi41
                        }
                        if (((v11 * 1e10) * float32(float64(1e10))) <= v11) {
                            v56 := GenS6{F0: -596}
                            if (42 == -753) {
                                var v57 int8 = 0
                                if v3 {
                                    var v58 int32 = int32(v35)
                                    for fi59 := 0; fi59 < 16; fi59++ {
                                        v60 := GenS1{F0: ""}
                                        _ = fi59
                                        _ = v60
                                    }
                                    var v61 uint = 18446744073709551615
                                    _ = v58
                                    _ = v61
                                } else {
                                    var v62 int16 = 42
                                    v1 += 727
                                    switch "\n\t" {
                                    case "a":
                                        if len(v6) > 0 {
                                            v6[0] = 2
                                        }
                                    case "hello":
                                        var v63 uint8 = 42
                                        _ = v63
                                    case "":
                                        if ((1.5 / (float64(v11) + 0.0)) < 0.0) {
                                            v1 += v1
                                            switch "a" {
                                            case "foo bar":
                                                _ = len(v6)
                                            }
                                        } else {
                                            fmt.Println(v3, v3, v35)
                                            var v64 float32 = v11
                                            _ = v64
                                        }
                                    }
                                    _ = v62
                                }
                                for fi65 := 0; fi65 < 8; fi65++ {
                                    v6 = append(v6, (v35 * v35))
                                    _ = fi65
                                }
                                _ = v57
                            }
                            _ = v56
                        }
                        v35 += (v35 | (^(v35)))
                    } else {
                        var v66 int = (fi39 + (^(fi39)))
                        v67 := GenS7{F0: v66}
                        v3 = (4294967295 >= (v1 ^ (v1 * v1)))
                        _ = v66
                        _ = v67
                    }
                case "\n\t":
                    v3 = false
                }
            } else {
                fmt.Println(v1, v11, v1)
                v35 = v35
            }
            _ = fi39
            _ = v40
        }
        fmt.Println(v3)
        _ = v38
    } else {
        v68 := GenS4{F0: 947}
        if (((v1 << ((0) & 63)) < 1) || false) {
            v3 = (0 > 52)
            v69 := GenS3{F0: 32767, F1: (^(v35))}
            if ("hello" > "\n\t") {
                v35 = v35
                var v70 uint = 18446744073709551615
                _ = v70
            } else {
                for fi71 := 0; fi71 < 3; fi71++ {
                    for fi72 := 0; fi72 < 7; fi72++ {
                        var v73 int32 = 2147483647
                        v73 = (v73 / ((1) | 1))
                        _ = fi72
                        _ = v73
                    }
                    _ = fi71
                }
                fmt.Println(v1, v11)
            }
            _ = v69
        }
        _ = v68
    }
    v35 = (v35 >> ((1) & 63))
    v74 := GenS3{F0: 683, F1: (^(v35))}
    for fi75 := 0; fi75 < 18; fi75++ {
        var v76 uint64 = 0
        v76 = v1
        fmt.Println(v1)
        _ = fi75
        _ = v76
    }
    var v77 float32 = v11
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v74
    _ = v77
}
