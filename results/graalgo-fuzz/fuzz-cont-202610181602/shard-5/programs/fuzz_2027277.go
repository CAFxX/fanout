package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 uint8; F1 int16; F2 int }
type GenS4 struct { F0 uint; F1 string }
type GenS5 struct { F0 bool; F1 uint32; F2 int64 }
type GenS6 struct { F0 uint; F1 int32 }
type GenS7 struct { F0 uint; F1 string; F2 int16 }
func genFn0() uint32 {
    for fi5 := 0; fi5 < 11; fi5++ {
        v6 := GenS1{F0: 18446744073709551615, F1: 305}
        _ = fi5
        _ = v6
    }
    var v7 int = 32767
    v8 := []int{v7, (v7 & v7)}
    v8 = append(v8, v7)
    var v9 int = (v7 - (^(v7)))
    _ = v7
    _ = v8
    _ = v9
    return 267
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]string)
    var v3 int64 = -9223372036854775808
    var v4 int64 = (v3 - (v3 ^ (v3 - (v3 * -9223372036854775808))))
    if ((v3 << ((4294967295) & 63)) != (v4 / ((586) | 1))) {
        var v10 uint32 = 2
        var v11 int16 = 69
        _ = v10
        _ = v11
    } else {
        v1[4294967295] = "foo bar"
    }
    v12 := v1[0]
    switch (("\n\t" + "a") + (((v12 + v12) + "a") + "\n\t")) {
    case "hello":
        var v13 int8 = 123
        _ = v13
    case "foo bar":
        v14 := []int{1, 32767}
        _ = v14
    }
    var v15 string = v12
    v16 := []int{32768, 0, 128}
    _ = len(v1)
    for fi17 := 0; fi17 < 3; fi17++ {
        fmt.Println(v3, fi17, v3)
        if true {
            var v18 uint = 32
            delete(v2, 2147483647)
            fmt.Println(v3, fi17, v4)
            _ = v18
        } else {
            switch v12 {
            case "\n\t":
                switch (fi17 / ((fi17) | 1)) {
                case 212:
                    for fi19 := 0; fi19 < 8; fi19++ {
                        if true {
                            v1[821] = "foo bar"
                            if true {
                                var v20 int16 = 1
                                var v21 uint16 = 0
                                var v22 uint8 = 127
                                _ = v20
                                _ = v21
                                _ = v22
                            }
                        }
                        v23 := GenS2{F0: -1, F1: 411}
                        fmt.Println(v15)
                        v12 = (v12 + (v15 + v12))
                        var v24 int8 = 42
                        _ = fi19
                        _ = v23
                        _ = v24
                    }
                case 1:
                    v25 := GenS3{F0: 0, F1: 238, F2: fi17}
                    _ = v25
                default:
                    v26 := GenS3{F0: 1, F1: 1, F2: fi17}
                    _ = v26
                }
            case "foo bar":
                if false {
                    for fi27 := 0; fi27 < 1; fi27++ {
                        v28 := GenS1{F0: 2, F1: 1}
                        _ = fi27
                        _ = v28
                    }
                    v29 := v2[2]
                    _ = v29
                } else {
                    for fi30 := 0; fi30 < 16; fi30++ {
                        v16 = append(v16, (fi17 | fi17))
                        v15 = ("" + ("foo bar" + v15))
                        _ = fi30
                    }
                }
            case "":
                if true {
                    for fi31 := 0; fi31 < 16; fi31++ {
                        v12 = "a"
                        _ = fi31
                    }
                } else {
                    var v32 uint64 = 2
                    switch (fi17 % (((fi17 | 0)) | 1)) {
                    case -2147483648:
                        var v33 bool = true
                        _ = v33
                    case 32767:
                        for fi34 := 0; fi34 < 9; fi34++ {
                            switch (v12 + "a") {
                            case "":
                                if len(v16) >= 2 {
                                    v16 = v16[1:len(v16)]
                                }
                            case "\n\t":
                                fmt.Println(fi17)
                            case "hello":
                                fmt.Println(v32, v15, v3)
                            }
                            delete(v1, 825)
                            _ = fi34
                        }
                    }
                    v4 -= (^(v4))
                    _ = v32
                }
            default:
                var v35 uint8 = 42
                _ = v35
            }
            v16 = append(v16, (fi17 ^ -9223372036854775808))
        }
        if false {
            _ = len(v1)
            var v36 float32 = (1e10 * float32(0.1))
            _ = v36
        }
        _ = fi17
    }
    if true {
        for fi37 := 0; fi37 < 4; fi37++ {
            fmt.Println(v4)
            v15 = "hello"
            _ = fi37
        }
        if (921 >= 2147483647) {
            if true {
                var v38 uint64 = 4294967295
                for fi39 := 0; fi39 < 19; fi39++ {
                    var v40 uint64 = 18446744073709551615
                    v40 = 1
                    v38 = 4294967295
                    _ = fi39
                    _ = v40
                }
                v41 := GenS4{F0: 783, F1: ("" + ("foo bar" + "hello"))}
                _ = v38
                _ = v41
            } else {
                fmt.Println(v4, v3, v12)
                v15 = v12
            }
            for fi42 := 0; fi42 < 18; fi42++ {
                var v43 int8 = -128
                v3 += int64(v43)
                _ = fi42
                _ = v43
            }
            fmt.Println(v12)
        } else {
            for fi44 := 0; fi44 < 9; fi44++ {
                for fi45 := 0; fi45 < 20; fi45++ {
                    if len(v16) > 0 {
                        v16[0] = fi45
                    }
                    _ = fi45
                }
                switch fi44 {
                case 43:
                    if (2147483647 > -397) {
                        for fi46 := 0; fi46 < 8; fi46++ {
                            var v47 int8 = 42
                            v15 = (v12 + v12)
                            if (264 > 42) {
                                if ((v12 + "\n\t") > v15) {
                                    if (v4 > 442) {
                                        var v48 float64 = float64(float32(float64(1e10)))
                                        var v49 float32 = (float32((v48 / v48)) + 0.1)
                                        _ = v48
                                        _ = v49
                                    }
                                    if true {
                                        var v50 uint16 = 1
                                        var v51 uint64 = 42
                                        v51 = v51
                                        _ = v50
                                        _ = v51
                                    }
                                } else {
                                    var v52 int64 = -9223372036854775808
                                    _ = v52
                                }
                                var v53 uint8 = 1
                                if true {
                                    var v54 int8 = (v47 + 0)
                                    var v55 int16 = 42
                                    _ = v54
                                    _ = v55
                                }
                                var v56 float32 = -2.25
                                fmt.Println(v15)
                                _ = v53
                                _ = v56
                            } else {
                                var v57 uint16 = 32767
                                var v58 int8 = 0
                                v59 := GenS4{F0: 65535, F1: "foo bar"}
                                _ = v57
                                _ = v58
                                _ = v59
                            }
                            _ = fi46
                            _ = v47
                        }
                        var v60 float64 = -2.25
                        _ = v60
                    }
                }
                _ = fi44
            }
            if (false || (v12 <= "\n\t")) {
                if len(v16) > 0 {
                    v16[0] = -119
                }
                v4 -= 42
            } else {
                var v61 uint32 = 42
                _ = v61
            }
        }
        v62 := GenS5{F0: (!true), F1: 42, F2: 42}
        _ = v62
    }
    var v63 uint32 = 99
    var v64 uint = uint(v63)
    var v65 uint8 = 127
    switch "foo bar" {
    case "foo bar":
        fmt.Println(v12, v12)
    default:
        if true {
            var v66 int32 = 0
            fmt.Println(v66)
            if len(v16) > 0 {
                v16[0] = -698
            }
            _ = v66
        } else {
            for fi67 := 0; fi67 < 5; fi67++ {
                _ = len(v1)
                switch (fi67 % (((-(fi67))) | 1)) {
                case 9223372036854775807:
                    fmt.Println(fi67)
                case -2147483648:
                    _ = len(v2)
                }
                if (v12 != (v12 + v12)) {
                    v68 := GenS5{F0: false, F1: (v63 & (v63 + genFn0())), F2: (-(v4))}
                    var v69 int8 = 0
                    switch int(v69) {
                    case 42:
                        switch v12 {
                        case "\n\t":
                            var v70 uint64 = 484
                            _ = v70
                        default:
                            v3 -= (^(v4))
                        }
                    case 162:
                        v69 -= v69
                    case -9223372036854775808:
                        v63 = genFn0()
                    }
                    _ = v68
                    _ = v69
                }
                _ = fi67
            }
        }
    }
    delete(v2, (v63 ^ (^(v63))))
    switch v12 {
    case "hello":
        switch 102 {
        case -5:
            switch v15 {
            case "\n\t":
                var v71 bool = ((v15 + v15) < (("a" + v12) + v12))
                _ = v71
            case "hello":
                var v72 float64 = 1.5
                _ = v72
            default:
                var v73 uint32 = (v63 % (((v63 / (((v63 + v63)) | 1))) | 1))
                _ = v73
            }
        case 32768:
            switch "" {
            case "hello":
                var v74 uint = (v64 ^ 966)
                _ = v74
            default:
                v63 = (^(v63))
            }
        case -9223372036854775808:
            for fi75 := 0; fi75 < 13; fi75++ {
                v65 = v65
                _ = fi75
            }
        default:
            switch "a" {
            case "hello":
                v12 = v12
            case "foo bar":
                v12 = "hello"
            case "a":
                if false {
                    for fi76 := 0; fi76 < 19; fi76++ {
                        var v77 uint = v64
                        v78 := GenS4{F0: 4294967295, F1: v15}
                        switch (fi76 & fi76) {
                        case 179:
                            v79 := GenS1{F0: 1, F1: 76}
                            _ = v79
                        default:
                            fmt.Println(v63)
                        }
                        _ = fi76
                        _ = v77
                        _ = v78
                    }
                } else {
                    fmt.Println(v65, v4, v3)
                    for fi80 := 0; fi80 < 5; fi80++ {
                        v12 = (v12 + v12)
                        _ = fi80
                    }
                }
            default:
                switch 625 {
                case 4294967295:
                    var v81 string = v15
                    _ = v81
                case 1:
                    v82 := GenS6{F0: (v64 ^ (v64 / ((v64) | 1))), F1: 1}
                    _ = v82
                case 32768:
                    v83 := v1[105]
                    _ = v83
                default:
                    for fi84 := 0; fi84 < 10; fi84++ {
                        if true {
                            if (665 <= 2) {
                                delete(v2, (v63 ^ (v63 | v63)))
                            } else {
                                v2[genFn0()] = "foo bar"
                                for fi85 := 0; fi85 < 19; fi85++ {
                                    if (true && false) {
                                        v86 := GenS7{F0: v64, F1: v15, F2: -32768}
                                        var v87 float32 = 0.0
                                        v65 = (v65 - (v65 & v65))
                                        var v88 int64 = (v4 % ((v3) | 1))
                                        fmt.Println(fi85, v15, v88)
                                        _ = v86
                                        _ = v87
                                        _ = v88
                                    } else {
                                        v89 := GenS6{F0: v64, F1: 42}
                                        var v90 int64 = v4
                                        _ = v89
                                        _ = v90
                                    }
                                    _ = fi85
                                }
                            }
                            switch v15 {
                            case "\n\t":
                                fmt.Println(v15)
                            case "foo bar":
                                v15 = ""
                            case "a":
                                v91 := GenS3{F0: (v65 & (v65 + v65)), F1: 32767, F2: (fi84 % (((fi84 | 2147483647)) | 1))}
                                _ = v91
                            default:
                                if false {
                                    if false {
                                        v92 := v1[genFn0()]
                                        v93 := GenS6{F0: (v64 - (v64 | (v64 / (((v64 * (v64 * (^(v64))))) | 1)))), F1: -871}
                                        if len(v16) >= 2 {
                                            v16 = v16[1:len(v16)]
                                        }
                                        _ = v92
                                        _ = v93
                                    }
                                } else {
                                    fmt.Println(v12)
                                    v94 := GenS6{F0: (^(v64)), F1: 42}
                                    _ = v94
                                }
                            }
                        }
                        var v95 int32 = 220
                        v64 = (^(v64))
                        _ = fi84
                        _ = v95
                    }
                }
            }
        }
    case "a":
        if true {
            v2[v63] = (v15 + (v15 + "foo bar"))
            switch -790 {
            case 4294967295:
                v12 = (v15 + ("foo bar" + (v12 + v15)))
            default:
                v3 -= 614
            }
        }
    }
    _ = len(v2)
    v65 = (v65 % ((v65) | 1))
    _ = len(v2)
    _ = len(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v63
    _ = v64
    _ = v65
}
