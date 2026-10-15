package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 uint32 }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 bool; F1 int8 }
type GenS4 struct { F0 int32; F1 int; F2 uint8 }
type GenS5 struct { F0 string; F1 int64 }
type GenS6 struct { F0 int16; F1 int64 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint32; F1 int16; F2 int8 }
type GenS9 struct { F0 uint64; F1 uint; F2 string }

func main() {
    v1 := []string{"hello", "a", "a", ("foo bar" + ((("" + "\n\t") + "foo bar") + "\n\t"))}
    v2 := make(map[int]int)
    v3 := []int{141, 128, -151}
    v4 := []int{9223372036854775807}
    v5 := []int{0, 263}
    v6 := make(map[uint32]uint32)
    var v7 int64 = 0
    fmt.Println(v7)
    switch 2147483647 {
    case 0:
        fmt.Println(v7)
    case 9223372036854775807:
        if (0.1 <= 0.0) {
            var v8 uint8 = 153
            var v9 uint64 = 18446744073709551615
            _ = v8
            _ = v9
        } else {
            for fi10 := 0; fi10 < 15; fi10++ {
                var v11 string = ""
                _ = fi10
                _ = v11
            }
        }
    default:
        fmt.Println(v7)
    }
    switch 32767 {
    case 474:
        if false {
            for fi12 := 0; fi12 < 13; fi12++ {
                _ = len(v5)
                if (!(0.0 <= (1e10 + (1.5 * -2.25)))) {
                    v13 := GenS1{F0: 579, F1: (fi12 << ((246) & 63)), F2: 2147483647}
                    v14 := GenS1{F0: -671, F1: (fi12 + (fi12 / ((fi12) | 1))), F2: 2}
                    for fi15 := 0; fi15 < 9; fi15++ {
                        var v16 bool = ((("hello" + "\n\t") + "foo bar") >= "foo bar")
                        if (1e10 != 3.14159) {
                            var v17 int64 = 1
                            if (!v16) {
                                v7 -= (^(v17))
                                var v18 int32 = -723
                                var v19 int64 = -9223372036854775808
                                _ = v18
                                _ = v19
                            } else {
                                fmt.Println(fi12)
                            }
                            _ = v17
                        } else {
                            switch 0 {
                            case 549:
                                v7 = 9223372036854775807
                            case 9223372036854775807:
                                var v20 bool = (!v16)
                                _ = v20
                            default:
                                fmt.Println(v7, v16, v7)
                            }
                            v6[65535] = 1
                            fmt.Println(v16, v16, fi15)
                        }
                        if v16 {
                            if ((!false) && (585 >= 4294967295)) {
                                var v21 int8 = 1
                                _ = v21
                            }
                            switch ("" + (("a" + ("" + ("" + "hello"))) + ("" + ("a" + ("\n\t" + "foo bar"))))) {
                            case "":
                                if (!(false || (-1 == 2147483647))) {
                                    var v22 bool = v16
                                    _ = len(v2)
                                    v7 += (-(v7))
                                    _ = v22
                                } else {
                                    var v23 int16 = 1
                                    if len(v1) > 0 {
                                        v1[0] = "foo bar"
                                    }
                                    var v24 int64 = v7
                                    _ = len(v3)
                                    var v25 uint = 4294967295
                                    _ = v23
                                    _ = v24
                                    _ = v25
                                }
                            case "a":
                                v26 := GenS1{F0: 2147483647, F1: (fi12 >> ((1) & 63)), F2: 65535}
                                _ = v26
                            case "\n\t":
                                v27 := GenS1{F0: 0, F1: fi12, F2: 42}
                                _ = v27
                            default:
                                if (65 != 1) {
                                    v28 := v2[(fi12 ^ fi15)]
                                    v16 = true
                                    if (4294967295 >= 377) {
                                        _ = len(v6)
                                    }
                                    _ = v28
                                }
                            }
                        } else {
                            v29 := GenS2{F0: -689, F1: v7}
                            if (2 <= 25) {
                                var v30 uint32 = 2147483647
                                _ = v30
                            } else {
                                fmt.Println(fi12)
                                v31 := GenS2{F0: 2147483647, F1: v7}
                                if (0 >= 172) {
                                    switch "\n\t" {
                                    case "a":
                                        var v32 float32 = 1e10
                                        _ = v32
                                    case "\n\t":
                                        v16 = (128 >= 42)
                                    default:
                                        var v33 int64 = -1
                                        _ = v33
                                    }
                                    v16 = v16
                                    v34 := GenS2{F0: 42, F1: (v7 * (v7 ^ v7))}
                                    _ = v34
                                } else {
                                    v35 := GenS3{F0: v16, F1: -1}
                                    _ = v35
                                }
                                _ = v31
                            }
                            var v36 int8 = 19
                            _ = v29
                            _ = v36
                        }
                        fmt.Println(fi12, v7, v7)
                        _ = fi15
                        _ = v16
                    }
                    _ = v13
                    _ = v14
                }
                _ = fi12
            }
            var v37 uint64 = 4294967295
            var v38 uint64 = (v37 >> ((433) & 63))
            _ = v37
            _ = v38
        } else {
            v39 := GenS3{F0: false, F1: -128}
            v7 = v7
            var v40 uint8 = 7
            _ = v39
            _ = v40
        }
    case -767:
        var v41 uint16 = 363
        _ = v41
    case -1:
        switch "a" {
        case "hello":
            v42 := GenS4{F0: 1, F1: 1, F2: 42}
            _ = v42
        }
    default:
        var v43 float64 = 0.0
        _ = v43
    }
    v7 = v7
    if (!true) {
        v7 = v7
    } else {
        var v44 uint = 0
        switch "\n\t" {
        case "foo bar":
            if (!(0 == 2147483647)) {
                v7 += v7
                if ((v7 >> (((v44 ^ v44)) & 63)) > v7) {
                    v45 := GenS2{F0: -2147483648, F1: (v7 + (-(v7)))}
                    _ = v45
                }
                v44 = (v44 | v44)
            }
        case "\n\t":
            var v46 string = "foo bar"
            _ = v46
        default:
            for fi47 := 0; fi47 < 20; fi47++ {
                if (127 >= 0) {
                    v44 -= 2
                } else {
                    for fi48 := 0; fi48 < 18; fi48++ {
                        var v49 float32 = -2.25
                        switch "hello" {
                        case "a":
                            var v50 uint16 = 65535
                            _ = v50
                        case "\n\t":
                            v51 := GenS5{F0: "hello", F1: 9223372036854775807}
                            _ = v51
                        case "":
                            var v52 int = (fi48 >> ((2) & 63))
                            if false {
                                v49 = v49
                                switch fi48 {
                                case 128:
                                    v44 = (^(v44))
                                case -654:
                                    var v53 uint16 = 1
                                    _ = v53
                                case -2147483648:
                                    switch v52 {
                                    case 4294967295:
                                        v44 = v44
                                    default:
                                        var v54 uint16 = 1
                                        _ = v54
                                    }
                                default:
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                            }
                            v55 := GenS5{F0: "\n\t", F1: (v7 | v7)}
                            v56 := GenS6{F0: 1, F1: 9223372036854775807}
                            _ = v52
                            _ = v55
                            _ = v56
                        }
                        if (!false) {
                            v44 = v44
                            v57 := GenS3{F0: true, F1: -48}
                            _ = v57
                        }
                        _ = fi48
                        _ = v49
                    }
                    if false {
                        fmt.Println(v7)
                    }
                }
                switch fi47 {
                case 2147483647:
                    v7 -= (v7 ^ v7)
                case -9223372036854775808:
                    for fi58 := 0; fi58 < 10; fi58++ {
                        var v59 float64 = float64((0.1 - float32(3.14159)))
                        _ = fi58
                        _ = v59
                    }
                }
                _ = fi47
            }
        }
        switch 580 {
        case -47:
            var v60 int64 = v7
            _ = v60
        }
        _ = v44
    }
    v7 = (v7 * (v7 ^ v7))
    v61 := GenS7{F0: 42}
    var v62 uint64 = 2
    for fi63 := 0; fi63 < 16; fi63++ {
        v62 -= (v62 - v62)
        v62 = v62
        _ = fi63
    }
    var v64 int16 = -32768
    var v65 uint64 = (^(v62))
    for fi66 := 0; fi66 < 16; fi66++ {
        if true {
            var v67 int64 = (v7 % ((v7) | 1))
            _ = v67
        } else {
            for fi68 := 0; fi68 < 2; fi68++ {
                if true {
                    v65 = 1
                    var v69 uint32 = 463
                    v70 := GenS6{F0: (v64 << ((865) & 63)), F1: 594}
                    _ = v69
                    _ = v70
                } else {
                    v2[-736] = fi66
                    fmt.Println(v64, v7)
                }
                delete(v2, int(v65))
                v71 := GenS8{F0: 4294967295, F1: int16(fi68), F2: 0}
                _ = fi68
                _ = v71
            }
        }
        v7 = 1
        fmt.Println(fi66, fi66, v7)
        _ = fi66
    }
    v62 += v62
    if true {
        delete(v6, 681)
        for fi72 := 0; fi72 < 3; fi72++ {
            if (0.1 <= float64((-2.25 - (1.5 * 3.14159)))) {
                if false {
                    var v73 int32 = -2147483648
                    for fi74 := 0; fi74 < 9; fi74++ {
                        v73 = (v73 / ((-807) | 1))
                        v75 := v2[32767]
                        if true {
                            fmt.Println(v65, v75, v65)
                            if true {
                                fmt.Println(v75, v75, v65)
                                v73 = v73
                                if len(v5) > 0 {
                                    v5[0] = fi74
                                }
                            }
                            fmt.Println(fi72)
                        }
                        _ = fi74
                        _ = v75
                    }
                    v6[683] = 1
                    _ = v73
                }
            }
            v76 := GenS9{F0: (^(v62)), F1: 93, F2: "hello"}
            _ = fi72
            _ = v76
        }
    }
    var v77 bool = false
    var v78 int16 = v64
    fmt.Println(v7)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v62:", v62)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    fmt.Println("v7:", v7)
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v61
    _ = v62
    _ = v64
    _ = v65
    _ = v7
    _ = v77
    _ = v78
}
