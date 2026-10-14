package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint; F1 uint; F2 string }
type GenS7 struct { F0 int; F1 uint64; F2 uint }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 int16; F1 int16; F2 int32 }
type GenS10 struct { F0 int64; F1 uint32; F2 int8 }

func main() {
    v1 := []int{32767, -735, -2147483648}
    v2 := make(map[int]int)
    var v3 uint8 = 0
    var v4 uint = 4294967295
    v5 := []int{127, 128}
    if false {
        var v6 uint64 = 18446744073709551615
        v7 := GenS1{F0: 442, F1: 1}
        var v8 uint32 = 42
        _ = v6
        _ = v7
        _ = v8
    } else {
        var v9 string = (("hello" + "\n\t") + "\n\t")
        v10 := GenS2{F0: ("" + (v9 + v9))}
        for fi11 := 0; fi11 < 13; fi11++ {
            var v12 int = (fi11 / (((^(fi11))) | 1))
            _ = fi11
            _ = v12
        }
        _ = v10
        _ = v9
    }
    fmt.Println(v3, v4)
    v3 += v3
    v13 := GenS2{F0: (("\n\t" + "a") + "\n\t")}
    if (0.1 < 1.5) {
        v3 -= (v3 | v3)
    }
    v4 += 0
    var v14 float32 = 0.0
    var v15 int8 = -23
    v4 += (^(v4))
    var v16 uint64 = 627
    v4 += 308
    if true {
        if true {
            for fi17 := 0; fi17 < 11; fi17++ {
                if (65535 <= uint16(v4)) {
                    v16 = 479
                } else {
                    fmt.Println(fi17)
                }
                _ = fi17
            }
            var v18 uint64 = v16
            switch "\n\t" {
            case "foo bar":
                v15 = v15
            }
            _ = v18
        }
    }
    if (false && true) {
        var v19 uint32 = 4294967295
        v20 := v2[9223372036854775807]
        switch "foo bar" {
        case "a":
            if (!(-10 > -432)) {
                for fi21 := 0; fi21 < 6; fi21++ {
                    switch (fi21 >> ((v4) & 63)) {
                    case 599:
                        v1 = append(v1, 786)
                    case 2147483647:
                        v4 = v4
                    case 128:
                        for fi22 := 0; fi22 < 3; fi22++ {
                            _ = len(v2)
                            v19 -= uint32(v3)
                            if false {
                                v19 -= 2147483648
                                v23 := GenS3{F0: (v3 + (v3 / ((v3) | 1)))}
                                if (-2.25 == (0.1 + 0.0)) {
                                    if true {
                                        var v24 string = (("\n\t" + "hello") + ("foo bar" + ("" + "")))
                                        var v25 bool = true
                                        _ = v24
                                        _ = v25
                                    }
                                    v4 += v4
                                }
                                _ = v23
                            } else {
                                _ = len(v2)
                                v4 += 255
                            }
                            _ = fi22
                        }
                    default:
                        var v26 uint32 = 65535
                        _ = v26
                    }
                    fmt.Println(v14, v3)
                    _ = fi21
                }
                _ = len(v2)
            }
        case "hello":
            var v27 float64 = -2.25
            _ = v27
        }
        _ = v19
        _ = v20
    } else {
        v5 = append(v5, 4294967295)
        v15 -= v15
        _ = len(v5)
    }
    switch "" {
    case "a":
        switch "hello" {
        case "\n\t":
            v3 = 1
        case "":
            for fi28 := 0; fi28 < 4; fi28++ {
                v3 += 240
                _ = fi28
            }
        default:
            for fi29 := 0; fi29 < 15; fi29++ {
                v30 := GenS3{F0: v3}
                v31 := GenS3{F0: (v3 * v3)}
                _ = fi29
                _ = v30
                _ = v31
            }
        }
    case "":
        if true {
            for fi32 := 0; fi32 < 6; fi32++ {
                v16 -= (v16 % ((v16) | 1))
                v4 += v4
                _ = fi32
            }
        } else {
            v33 := GenS4{F0: (v4 | 42)}
            switch 0 {
            case -594:
                v34 := GenS5{F0: (^(v3))}
                _ = v34
            default:
                var v35 uint8 = v3
                _ = v35
            }
            _ = v33
        }
    }
    v3 -= v3
    if true {
        switch -2147483648 {
        case 0:
            for fi36 := 0; fi36 < 3; fi36++ {
                for fi37 := 0; fi37 < 10; fi37++ {
                    v4 = 765
                    switch (-(fi36)) {
                    case -9223372036854775808:
                        var v38 uint32 = 1
                        var v39 int8 = v15
                        var v40 bool = true
                        _ = v38
                        _ = v39
                        _ = v40
                    }
                    if false {
                        var v41 int8 = int8(v4)
                        _ = v41
                    } else {
                        var v42 int16 = int16(v15)
                        _ = v42
                    }
                    _ = fi37
                }
                if false {
                    var v43 int16 = -1
                    if false {
                        for fi44 := 0; fi44 < 9; fi44++ {
                            v45 := GenS6{F0: 0, F1: v4, F2: "foo bar"}
                            v46 := GenS2{F0: ((("" + "\n\t") + (("foo bar" + "a") + ("" + "\n\t"))) + "\n\t")}
                            var v47 uint16 = 2
                            _ = fi44
                            _ = v45
                            _ = v46
                            _ = v47
                        }
                        v16 = 18446744073709551615
                    }
                    if (!(!(fi36 != fi36))) {
                        for fi48 := 0; fi48 < 9; fi48++ {
                            if (-2147483648 > 1) {
                                v49 := GenS2{F0: "foo bar"}
                                v3 = (^(v3))
                                if ((v14 <= (1e10 - (v14 * 0.1))) || (v4 < v4)) {
                                    v14 -= v14
                                    v50 := GenS6{F0: uint(v3), F1: v4, F2: ("hello" + (("hello" + "") + "a"))}
                                    _ = v50
                                } else {
                                    switch "hello" {
                                    case "a":
                                        if (!true) {
                                            v43 = 1000
                                        }
                                    case "hello":
                                        v51 := GenS4{F0: v4}
                                        _ = v51
                                    case "foo bar":
                                        v52 := GenS7{F0: (fi48 / ((128) | 1)), F1: v16, F2: 1}
                                        _ = v52
                                    }
                                    var v53 int32 = 42
                                    v53 += (v53 << (((^(v4))) & 63))
                                    var v54 int16 = (v43 >> (((v4 * (v4 << ((v4) & 63)))) & 63))
                                    _ = v53
                                    _ = v54
                                }
                                _ = v49
                            } else {
                                var v55 uint8 = v3
                                if (!(false && true)) {
                                    if true {
                                        v2[(fi48 ^ (fi36 + fi48))] = fi48
                                    } else {
                                        if false {
                                            var v56 uint64 = v16
                                            _ = v56
                                        } else {
                                            if true {
                                                var v57 int16 = (v43 >> ((v4) & 63))
                                                var v58 string = ((((("\n\t" + "a") + "") + "") + "a") + "a")
                                                _ = v57
                                                _ = v58
                                            } else {
                                                var v59 uint16 = 2
                                                var v60 uint64 = (v16 >> ((18446744073709551615) & 63))
                                                _ = v59
                                                _ = v60
                                            }
                                            if true {
                                                var v61 bool = (false && (v4 >= v4))
                                                if v61 {
                                                    v43 += v43
                                                } else {
                                                    var v62 uint = (v4 | v4)
                                                    _ = v62
                                                }
                                                _ = v61
                                            }
                                            v15 = (v15 * 1)
                                        }
                                    }
                                    switch ("hello" + "a") {
                                    case "foo bar":
                                        var v63 string = "a"
                                        _ = v63
                                    case "\n\t":
                                        v55 = (v55 * (v3 * 128))
                                    default:
                                        switch (fi36 - int(v4)) {
                                        case 4294967295:
                                            v64 := v2[(fi48 & 157)]
                                            _ = v64
                                        }
                                    }
                                } else {
                                    if len(v1) > 0 {
                                        v1[0] = int(v43)
                                    }
                                    v16 += v16
                                }
                                _ = v55
                            }
                            switch 127 {
                            case 489:
                                fmt.Println(v15, fi48)
                            case -9223372036854775808:
                                v4 = v4
                            case 42:
                                switch fi48 {
                                case 32767:
                                    v65 := GenS3{F0: (^(v3))}
                                    _ = v65
                                case -584:
                                    v66 := GenS8{F0: 508}
                                    _ = v66
                                case -9223372036854775808:
                                    var v67 string = ""
                                    _ = v67
                                default:
                                    fmt.Println(fi36, v3, v16)
                                }
                            }
                            _ = fi48
                        }
                        var v68 uint = 2
                        _ = v68
                    }
                    _ = v43
                }
                _ = fi36
            }
        }
    } else {
        var v69 int = 2147483647
        var v70 uint8 = v3
        _ = v69
        _ = v70
    }
    switch -1 {
    case -901:
        var v71 bool = true
        _ = v71
    }
    switch ("\n\t" + "foo bar") {
    case "foo bar":
        if (284 < 9223372036854775807) {
            for fi72 := 0; fi72 < 19; fi72++ {
                switch "foo bar" {
                case "hello":
                    var v73 float64 = 1e10
                    _ = v73
                case "\n\t":
                    v4 -= (v4 | v4)
                case "foo bar":
                    if false {
                        if false {
                            switch -1 {
                            case 42:
                                if false {
                                    fmt.Println(v4, v16)
                                    fmt.Println(v14, fi72)
                                    v74 := GenS9{F0: -32768, F1: -67, F2: -832}
                                    _ = v74
                                }
                            default:
                                _ = len(v2)
                            }
                            for fi75 := 0; fi75 < 11; fi75++ {
                                switch (fi72 | (-(fi72))) {
                                case 42:
                                    v76 := GenS7{F0: fi75, F1: (v16 ^ v16), F2: (v4 & v4)}
                                    _ = v76
                                case 4294967295:
                                    var v77 uint = (v4 | (v4 * 26))
                                    _ = v77
                                }
                                _ = fi75
                            }
                            v14 = (0.1 + ((((v14 - v14) / (((1.5 + 1.5) - v14) / ((v14 / v14) * v14))) * (((v14 - (0.1 - v14)) * v14) / ((v14 * (0.0 - v14)) / (v14 - v14)))) / v14))
                        } else {
                            v14 -= v14
                            v4 += (v4 * v4)
                            v3 = (v3 + v3)
                        }
                        delete(v2, fi72)
                    }
                default:
                    if len(v5) > 0 {
                        v5[0] = (fi72 % (((fi72 % (((-(fi72))) | 1))) | 1))
                    }
                }
                if (false && (!true)) {
                    var v78 uint16 = 0
                    for fi79 := 0; fi79 < 15; fi79++ {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        if false {
                            v78 -= (v78 * v78)
                            if ((^(v4)) >= (v4 / (((v4 / ((v4) | 1))) | 1))) {
                                if false {
                                    if true {
                                        var v80 int = 32768
                                        switch (v80 << ((v4) & 63)) {
                                        case -1:
                                            var v81 int32 = 42
                                            if (42 != -1) {
                                                v4 = 4294967295
                                                v16 = v16
                                                v82 := GenS9{F0: -1, F1: 1000, F2: -511}
                                                if ((true || false) && false) {
                                                    if true {
                                                        var v83 int16 = 42
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                        _ = v83
                                                    } else {
                                                        v16 -= v16
                                                    }
                                                }
                                                _ = v82
                                            }
                                            v81 = (-(v81))
                                            _ = v81
                                        case 128:
                                            v84 := v2[v80]
                                            _ = v84
                                        case 32768:
                                            if true {
                                                if (0 >= 81) {
                                                    var v85 string = ("a" + (("" + "foo bar") + ("\n\t" + ("" + "foo bar"))))
                                                    var v86 uint = (v4 | v4)
                                                    _ = v85
                                                    _ = v86
                                                } else {
                                                    v3 += v3
                                                    var v87 int32 = -2147483648
                                                    _ = v87
                                                }
                                                v4 += v4
                                                v80 -= 32767
                                                if len(v5) > 0 {
                                                    v5[0] = 967
                                                }
                                            }
                                        default:
                                            if ((v16 | (v16 % ((uint64(v3)) | 1))) >= (^(v16))) {
                                                v88 := GenS10{F0: -1, F1: 2147483648, F2: (v15 % ((-128) | 1))}
                                                v4 = v4
                                                var v89 float64 = 3.14159
                                                _ = v88
                                                _ = v89
                                            }
                                        }
                                        _ = v80
                                    } else {
                                        var v90 int32 = 1
                                        _ = v90
                                    }
                                    var v91 uint64 = (v16 & 160)
                                    _ = v91
                                }
                                fmt.Println(v4, fi79)
                                var v92 float64 = ((1.5 - (((1e10 + ((((0.0 + (1e10 - float64((v14 * v14)))) - 0.0) / float64(((v14 + v14) * v14))) / float64(v14))) * (float64(v14) - 1e10)) - (float64(v14) / 1e10))) / (1.5 / float64(float32((3.14159 * float64(v14))))))
                                _ = v92
                            }
                            v78 = (v78 + v78)
                            fmt.Println(v3, v78)
                        } else {
                            fmt.Println(v14, v15)
                            var v93 uint8 = v3
                            _ = v93
                        }
                        _ = fi79
                    }
                    v94 := GenS10{F0: -9223372036854775808, F1: 1, F2: (^(v15))}
                    _ = v78
                    _ = v94
                }
                var v95 float32 = v14
                _ = fi72
                _ = v95
            }
            _ = len(v5)
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
