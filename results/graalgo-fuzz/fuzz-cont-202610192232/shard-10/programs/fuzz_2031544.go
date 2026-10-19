package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint8; F1 int64; F2 int64 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint16; F1 int8 }
type GenS8 struct { F0 uint16; F1 int32; F2 uint }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 string; F1 int8 }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 int64; F1 int }
type GenS13 struct { F0 int8; F1 bool }

func main() {
    v1 := []string{}
    var v2 uint8 = 2
    v3 := make(map[uint32]uint32)
    var v4 float64 = (3.14159 - 1.5)
    var v5 uint64 = 0
    if true {
        v2 = 0
        var v6 int32 = -991
        v7 := GenS1{F0: 42, F1: "\n\t"}
        _ = v6
        _ = v7
    }
    _ = len(v3)
    if true {
        switch 42 {
        case 4294967295:
            v2 = v2
        case -2147483648:
            var v8 int8 = 1
            _ = v8
        default:
            var v9 string = "foo bar"
            _ = v9
        }
    } else {
        v10 := GenS2{F0: -9223372036854775808, F1: -32768}
        for fi11 := 0; fi11 < 17; fi11++ {
            if (304 != 2147483647) {
                v2 = v2
                v12 := GenS3{F0: ((v2 ^ (^(v2))) >= v2)}
                var v13 float32 = 3.14159
                _ = v12
                _ = v13
            } else {
                v14 := GenS4{F0: 2147483647}
                for fi15 := 0; fi15 < 16; fi15++ {
                    var v16 string = ("" + "\n\t")
                    v16 = (v16 + ((("\n\t" + v16) + "\n\t") + "\n\t"))
                    v17 := GenS5{F0: (v2 >> ((4294967295) & 63)), F1: -1, F2: 0}
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                var v18 int16 = 32767
                _ = v14
                _ = v18
            }
            switch "\n\t" {
            case "a":
                for fi19 := 0; fi19 < 12; fi19++ {
                    var v20 int8 = 42
                    _ = fi19
                    _ = v20
                }
            case "hello":
                v2 = (v2 + (v2 - v2))
            case "\n\t":
                var v21 uint64 = v5
                _ = v21
            }
            _ = fi11
        }
        _ = v10
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v22 int = 127
    fmt.Println(v22)
    var v23 int = 32768
    for fi24 := 0; fi24 < 12; fi24++ {
        v5 = v5
        var v25 int = (fi24 ^ -305)
        _ = fi24
        _ = v25
    }
    var v26 int8 = 1
    switch "a" {
    case "a":
        for fi27 := 0; fi27 < 20; fi27++ {
            if false {
                switch 128 {
                case 233:
                    v28 := v3[65535]
                    _ = v28
                case -50:
                    if (v5 == (v5 % ((v5) | 1))) {
                        for fi29 := 0; fi29 < 19; fi29++ {
                            var v30 float32 = 0.0
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            v5 -= 18446744073709551615
                            var v31 float64 = -2.25
                            _ = fi29
                            _ = v30
                            _ = v31
                        }
                        var v32 bool = true
                        _ = v32
                    }
                case -420:
                    v23 = (fi27 ^ 32767)
                }
            } else {
                if (v2 == (^(v2))) {
                    v23 += (fi27 - 4294967295)
                    if (65535 == 65535) {
                        var v33 int = v23
                        _ = v33
                    }
                }
                if false {
                    for fi34 := 0; fi34 < 6; fi34++ {
                        var v35 float64 = float64(0.1)
                        v4 = 0.1
                        if (int(v26) != (fi34 | (fi27 | (fi27 ^ v23)))) {
                            if false {
                                fmt.Println(v5, v35)
                                fmt.Println(v26)
                            } else {
                                v23 = v23
                                switch fi34 {
                                case 193:
                                    var v36 int = 32767
                                    _ = v36
                                case 128:
                                    var v37 int64 = 1
                                    _ = v37
                                case 231:
                                    v38 := GenS5{F0: (^(v2)), F1: -494, F2: 42}
                                    _ = v38
                                }
                            }
                        } else {
                            var v39 uint = 881
                            _ = v39
                        }
                        _ = fi34
                        _ = v35
                    }
                    fmt.Println(v26, v22, fi27)
                } else {
                    var v40 bool = (false || ((false || ((false || (377 >= 2147483647)) || false)) && true))
                    _ = v40
                }
                if true {
                    for fi41 := 0; fi41 < 19; fi41++ {
                        var v42 int32 = 2147483647
                        _ = fi41
                        _ = v42
                    }
                    v23 = 2147483647
                }
            }
            switch (v22 >> ((18446744073709551615) & 63)) {
            case 4294967295:
                v23 -= (fi27 >> ((1) & 63))
            case 365:
                if (v4 > v4) {
                    for fi43 := 0; fi43 < 8; fi43++ {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        _ = fi43
                    }
                } else {
                    v4 -= (0.0 + 3.14159)
                    if true {
                        v2 = (v2 - v2)
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        fmt.Println(v4)
                    }
                    v44 := GenS4{F0: 4294967295}
                    _ = v44
                }
            case -9223372036854775808:
                for fi45 := 0; fi45 < 13; fi45++ {
                    var v46 int32 = 0
                    _ = fi45
                    _ = v46
                }
            default:
                for fi47 := 0; fi47 < 2; fi47++ {
                    v48 := GenS6{F0: uint32(v2)}
                    _ = fi47
                    _ = v48
                }
            }
            v3[0] = 576
            _ = fi27
        }
    case "hello":
        var v49 int16 = 673
        _ = v49
    case "":
        var v50 float32 = (float32(v4) + 0.0)
        _ = v50
    default:
        v5 = 42
    }
    fmt.Println(v22, v22)
    if (false && true) {
        v4 += v4
        for fi51 := 0; fi51 < 12; fi51++ {
            for fi52 := 0; fi52 < 18; fi52++ {
                if true {
                    var v53 int32 = int32(v2)
                    switch (v23 / (((fi52 ^ (fi51 >> ((2) & 63)))) | 1)) {
                    case 0:
                        var v54 bool = true
                        _ = v54
                    }
                    v55 := GenS7{F0: 433, F1: -128}
                    _ = v53
                    _ = v55
                }
                var v56 float64 = v4
                if (1e10 >= (float32((3.14159 * 3.14159)) / 0.1)) {
                    switch (fi51 & (^(fi51))) {
                    case 32768:
                        if (947 != 4294967295) {
                            v22 += (-(fi52))
                        } else {
                            v57 := GenS6{F0: 2}
                            _ = v57
                        }
                    default:
                        v58 := v3[0]
                        _ = v58
                    }
                    if ((!true) || ("\n\t" < "\n\t")) {
                        if (65535 < uint32(v2)) {
                            v22 += -430
                        }
                        var v59 uint64 = (v5 - (v5 ^ v5))
                        _ = v59
                    } else {
                        v60 := GenS8{F0: 2, F1: 0, F2: 767}
                        v61 := GenS9{F0: false}
                        var v62 float32 = (1.5 - 0.0)
                        _ = v60
                        _ = v61
                        _ = v62
                    }
                    v63 := GenS5{F0: v2, F1: -1, F2: -813}
                    _ = len(v3)
                    _ = v63
                } else {
                    var v64 uint8 = uint8(v5)
                    v65 := GenS10{F0: "", F1: 1}
                    switch fi52 {
                    case 855:
                        var v66 bool = false
                        _ = v66
                    default:
                        switch v23 {
                        case -9223372036854775808:
                            switch fi51 {
                            case 710:
                                v64 += (v64 | (v2 / (((v2 >> ((1) & 63))) | 1)))
                            default:
                                _ = len(v3)
                                var v67 uint32 = 930
                                switch (-(fi52)) {
                                case 42:
                                    var v68 bool = true
                                    _ = v68
                                default:
                                    fmt.Println(v23)
                                    v23 = int(v26)
                                    if false {
                                        _ = len(v3)
                                        var v69 int = 4294967295
                                        _ = len(v1)
                                        if true {
                                            var v70 uint8 = 42
                                            v2 -= (v2 * (^(v64)))
                                            v4 -= 1e10
                                            _ = v70
                                        }
                                        if (v4 > v56) {
                                            switch fi52 {
                                            case 32767:
                                                var v71 uint64 = v5
                                                _ = v71
                                            case 9223372036854775807:
                                                v72 := GenS7{F0: 0, F1: v26}
                                                _ = v72
                                            case 32768:
                                                v23 -= fi51
                                            default:
                                                v26 = (-(v26))
                                            }
                                        } else {
                                            fmt.Println(v64, v23, v5)
                                            if false {
                                                v73 := GenS1{F0: 255, F1: "a"}
                                                var v74 uint64 = v5
                                                _ = v73
                                                _ = v74
                                            } else {
                                                v2 -= 42
                                                v22 += v22
                                                v75 := GenS1{F0: uint16(v5), F1: ""}
                                                var v76 uint32 = 65535
                                                _ = v75
                                                _ = v76
                                            }
                                        }
                                        _ = v69
                                    }
                                    fmt.Println(v26, v26, v2)
                                }
                                _ = v67
                            }
                        case 42:
                            var v77 uint = 65535
                            _ = v77
                        default:
                            if (1.5 <= (-2.25 + 1.5)) {
                                var v78 float64 = float64(0.1)
                                var v79 int64 = 582
                                _ = v78
                                _ = v79
                            } else {
                                v5 = 2
                            }
                        }
                    }
                    _ = v64
                    _ = v65
                }
                v2 += (v2 | v2)
                var v80 float32 = float32(v56)
                _ = fi52
                _ = v56
                _ = v80
            }
            v4 -= 0.0
            v81 := v3[829]
            _ = fi51
            _ = v81
        }
        v2 = (v2 - (^(v2)))
    } else {
        var v82 int64 = 42
        if false {
            switch "" {
            case "":
                fmt.Println(v23)
            case "foo bar":
                v1 = append(v1, "hello")
            case "hello":
                var v83 int32 = 386
                _ = v83
            }
            for fi84 := 0; fi84 < 19; fi84++ {
                v23 = fi84
                var v85 int64 = v82
                var v86 int32 = 542
                _ = fi84
                _ = v85
                _ = v86
            }
            fmt.Println(v4, v2)
        } else {
            fmt.Println(v4, v2)
        }
        _ = v82
    }
    if (0 > 4294967295) {
        for fi87 := 0; fi87 < 5; fi87++ {
            for fi88 := 0; fi88 < 4; fi88++ {
                if true {
                    var v89 uint64 = v5
                    _ = v89
                } else {
                    var v90 float32 = 1e10
                    v91 := GenS11{F0: 255}
                    _ = v90
                    _ = v91
                }
                _ = fi88
            }
            v92 := GenS6{F0: 158}
            switch ("a" + ("a" + "foo bar")) {
            case "hello":
                var v93 int32 = -962
                _ = v93
            }
            _ = fi87
            _ = v92
        }
    } else {
        if len(v1) > 0 {
            v1[0] = "a"
        }
        if ((!(!(!(!(0.0 <= (0.0 * float32(v4))))))) && (!true)) {
            v23 -= (^(v23))
            for fi94 := 0; fi94 < 16; fi94++ {
                for fi95 := 0; fi95 < 8; fi95++ {
                    v26 -= v26
                    v96 := GenS2{F0: v23, F1: 1000}
                    v26 -= v26
                    _ = fi95
                    _ = v96
                }
                if ((!false) || (true || (!true))) {
                    v97 := GenS7{F0: 414, F1: v26}
                    v98 := GenS12{F0: -1, F1: fi94}
                    v22 = (v22 ^ v22)
                    _ = v97
                    _ = v98
                }
                switch "hello" {
                case "foo bar":
                    var v99 int64 = 710
                    _ = v99
                case "":
                    if true {
                        for fi100 := 0; fi100 < 12; fi100++ {
                            var v101 float64 = (float64((0.0 / float32(-2.25))) * v4)
                            if true {
                                v102 := GenS13{F0: v26, F1: true}
                                _ = v102
                            }
                            v3[1] = 2147483648
                            _ = fi100
                            _ = v101
                        }
                        v3[196] = 2147483648
                        var v103 uint16 = 65535
                        _ = v103
                    } else {
                        for fi104 := 0; fi104 < 17; fi104++ {
                            switch v23 {
                            case -1:
                                _ = len(v1)
                            default:
                                switch "\n\t" {
                                case "a":
                                    var v105 uint8 = 117
                                    _ = v105
                                default:
                                    var v106 int64 = -513
                                    _ = v106
                                }
                            }
                            v2 -= (v2 & v2)
                            _ = fi104
                        }
                        v2 -= (v2 ^ (v2 << ((151) & 63)))
                        var v107 uint32 = 4294967295
                        _ = v107
                    }
                }
                _ = fi94
            }
        }
        v22 = v22
    }
    v23 = (^(v22))
    fmt.Println(v4, v26, v5)
    _ = len(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
