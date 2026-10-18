package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 int64 }
type GenS2 struct { F0 uint64; F1 int }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int; F1 int32; F2 int64 }
type GenS6 struct { F0 bool; F1 uint32; F2 uint32 }
type GenS7 struct { F0 uint8; F1 int16; F2 string }
type GenS8 struct { F0 bool; F1 int16; F2 int64 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 string }
type GenS12 struct { F0 int16 }
type GenS13 struct { F0 uint32; F1 int64 }
type GenS14 struct { F0 uint64; F1 uint }
type GenS15 struct { F0 string; F1 uint32; F2 bool }
type GenS16 struct { F0 int32; F1 bool; F2 int8 }

func main() {
    var v1 int = 0
    v2 := []int{v1, -647, v1}
    var v3 uint32 = 671
    v4 := []uint32{(v3 - v3), 65535, (v3 % (((^(v3))) | 1)), (v3 + 4294967295)}
    v5 := make(map[uint32]string)
    var v6 int16 = 153
    switch v1 {
    case -1:
        for fi7 := 0; fi7 < 6; fi7++ {
            for fi8 := 0; fi8 < 13; fi8++ {
                var v9 int16 = (v6 << ((832) & 63))
                switch (fi7 << ((255) & 63)) {
                case 128:
                    var v10 uint16 = 459
                    var v11 float64 = -2.25
                    v1 = fi8
                    _ = v10
                    _ = v11
                case -7:
                    v12 := v5[(v3 % (((^(v3))) | 1))]
                    _ = v12
                default:
                    v13 := GenS1{F0: 128, F1: (v3 & v3), F2: 474}
                    _ = v13
                }
                var v14 int64 = 1
                _ = fi8
                _ = v14
                _ = v9
            }
            v15 := GenS2{F0: 101, F1: v1}
            _ = fi7
            _ = v15
        }
    case 42:
        if len(v4) > 0 {
            v4[0] = (^(v3))
        }
    }
    if (1e10 != 0.1) {
        var v16 float32 = 1e10
        _ = v16
    } else {
        if true {
            var v17 int32 = 1
            _ = v17
        }
        var v18 bool = (42 != 866)
        var v19 string = "hello"
        _ = v18
        _ = v19
    }
    var v20 int8 = 127
    if false {
        if (float32(0.0) == 1.5) {
            var v21 float32 = -2.25
            var v22 uint32 = 824
            _ = v21
            _ = v22
        } else {
            fmt.Println(v20, v20)
        }
        switch "foo bar" {
        case "foo bar":
            v1 += v1
        case "\n\t":
            v23 := GenS3{F0: (v20 & (v20 & (v20 ^ v20)))}
            _ = v23
        }
        fmt.Println(v3, v1, v20)
    }
    var v24 int64 = -1
    if (true && false) {
        for fi25 := 0; fi25 < 2; fi25++ {
            fmt.Println(v20)
            for fi26 := 0; fi26 < 16; fi26++ {
                if true {
                    v27 := GenS4{F0: 461}
                    if len(v2) > 0 {
                        v2[0] = (v1 >> ((18446744073709551615) & 63))
                    }
                    var v28 int8 = v20
                    v1 = 1
                    v24 = v24
                    _ = v27
                    _ = v28
                }
                v29 := GenS5{F0: (fi25 % ((-1) | 1)), F1: 0, F2: (v24 / ((v24) | 1))}
                fmt.Println(v24)
                _ = fi26
                _ = v29
            }
            v3 -= v3
            _ = fi25
        }
        v30 := GenS6{F0: (v3 != (v3 + (^(v3)))), F1: (v3 - v3), F2: (v3 >> ((255) & 63))}
        if false {
            v6 = -1
            v31 := GenS7{F0: 2, F1: v6, F2: "\n\t"}
            _ = v31
        }
        _ = v30
    } else {
        if true {
            switch (v1 / ((-414) | 1)) {
            case 4294967295:
                var v32 int = (v1 - int(v3))
                _ = v32
            }
        } else {
            fmt.Println(v6, v3)
            switch int(v24) {
            case 667:
                v24 -= v24
            case 127:
                for fi33 := 0; fi33 < 8; fi33++ {
                    for fi34 := 0; fi34 < 4; fi34++ {
                        v35 := GenS8{F0: true, F1: (v6 | (v6 >> ((1) & 63))), F2: 42}
                        var v36 string = "\n\t"
                        _ = fi34
                        _ = v35
                        _ = v36
                    }
                    var v37 float32 = (((1e10 - (float32((-2.25 / (0.1 - -2.25))) - 0.1)) * (float32(-2.25) - 1.5)) * -2.25)
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi33
                    _ = v37
                }
            case -9223372036854775808:
                if false {
                    for fi38 := 0; fi38 < 16; fi38++ {
                        fmt.Println(v20, v3, v1)
                        _ = fi38
                    }
                    v39 := GenS3{F0: v20}
                    v20 = 0
                    _ = v39
                }
            default:
                v24 = v24
            }
            v40 := GenS9{F0: "\n\t"}
            _ = v40
        }
    }
    switch "a" {
    case "hello":
        var v41 uint64 = 1
        _ = v41
    case "":
        fmt.Println(v1, v6, v1)
    case "foo bar":
        if true {
            v42 := GenS3{F0: (v20 >> ((3) & 63))}
            for fi43 := 0; fi43 < 18; fi43++ {
                _ = len(v5)
                switch fi43 {
                case 4294967295:
                    fmt.Println(v6, v1)
                default:
                    if false {
                        switch fi43 {
                        case -484:
                            switch (((("" + ("" + "a")) + "") + ("\n\t" + "hello")) + "foo bar") {
                            case "a":
                                if true {
                                    v44 := v5[v3]
                                    for fi45 := 0; fi45 < 15; fi45++ {
                                        fmt.Println(v1)
                                        _ = fi45
                                    }
                                    if (!(v44 != v44)) {
                                        for fi46 := 0; fi46 < 16; fi46++ {
                                            var v47 int = (fi46 % ((v1) | 1))
                                            _ = fi46
                                            _ = v47
                                        }
                                        var v48 int32 = 2147483647
                                        _ = v48
                                    } else {
                                        v49 := GenS10{F0: 1}
                                        _ = v49
                                    }
                                    _ = v44
                                }
                            default:
                                var v50 string = ("\n\t" + "a")
                                _ = v50
                            }
                        case 128:
                            var v51 uint = 4294967295
                            _ = v51
                        }
                        if (!false) {
                            delete(v5, v3)
                            if len(v4) > 0 {
                                v4[0] = (v3 * v3)
                            }
                            if (-1 >= (v24 + (v24 << ((13) & 63)))) {
                                if (!true) {
                                    v20 = v20
                                    var v52 uint16 = 255
                                    var v53 int64 = (v24 * v24)
                                    _ = v52
                                    _ = v53
                                } else {
                                    _ = len(v4)
                                }
                            } else {
                                if false {
                                    fmt.Println(v24, v3, v24)
                                    var v54 float64 = ((-2.25 - 1.5) * (((0.0 - -2.25) * 0.0) + float64(0.0)))
                                    _ = v54
                                } else {
                                    switch fi43 {
                                    case -9223372036854775808:
                                        var v55 uint8 = 255
                                        _ = v55
                                    case -483:
                                        v56 := GenS11{F0: "\n\t"}
                                        _ = v56
                                    }
                                    fmt.Println(v6, v1, v6)
                                }
                                var v57 uint = 18446744073709551615
                                _ = v57
                            }
                        } else {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            if (42 == v3) {
                                fmt.Println(v6)
                                v58 := v5[65535]
                                for fi59 := 0; fi59 < 5; fi59++ {
                                    if (1 > 4294967295) {
                                        var v60 float64 = 0.0
                                        var v61 int32 = 1
                                        _ = v60
                                        _ = v61
                                    } else {
                                        if (false || false) {
                                            v24 += (v24 & (v24 ^ v24))
                                            fmt.Println(fi43)
                                            v62 := GenS6{F0: ((v3 < v3) && (237 == 18446744073709551615)), F1: v3, F2: 65535}
                                            _ = v62
                                        }
                                        v63 := v5[42]
                                        _ = v63
                                    }
                                    v20 = (v20 + -128)
                                    if false {
                                        v20 -= (v20 % ((v20) | 1))
                                    }
                                    _ = fi59
                                }
                                _ = v58
                            } else {
                                if (2 != 2) {
                                    var v64 uint32 = (^(v3))
                                    v65 := v5[v3]
                                    v24 -= v24
                                    _ = v64
                                    _ = v65
                                }
                                for fi66 := 0; fi66 < 19; fi66++ {
                                    v6 += v6
                                    var v67 uint16 = 2
                                    _ = fi66
                                    _ = v67
                                }
                                if true {
                                    v68 := GenS10{F0: 540}
                                    if (!(3.14159 == 0.1)) {
                                        fmt.Println(v24)
                                        var v69 uint8 = 255
                                        _ = v69
                                    }
                                    _ = v68
                                }
                            }
                        }
                        fmt.Println(v20, v1, v24)
                    } else {
                        for fi70 := 0; fi70 < 3; fi70++ {
                            v71 := GenS12{F0: v6}
                            if true {
                                if (!(0 >= 0)) {
                                    v72 := v5[(v3 & v3)]
                                    _ = v72
                                } else {
                                    if true {
                                        if true {
                                            v73 := GenS13{F0: (v3 >> ((65535) & 63)), F1: int64(v3)}
                                            _ = v73
                                        } else {
                                            var v74 uint16 = 32767
                                            _ = v74
                                        }
                                        var v75 uint = 760
                                        if false {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        } else {
                                            _ = len(v2)
                                        }
                                        _ = v75
                                    } else {
                                        var v76 int32 = -2147483648
                                        v2 = append(v2, 32767)
                                        v20 = int8(v76)
                                        var v77 uint32 = (v3 - 42)
                                        v6 -= (-(v6))
                                        var v78 bool = (0 == 583)
                                        _ = v76
                                        _ = v77
                                        _ = v78
                                    }
                                }
                            }
                            if ((fi70 % ((int(v3)) | 1)) >= v1) {
                                v24 += v24
                                delete(v5, v3)
                            } else {
                                v6 = (v6 / ((v6) | 1))
                                v79 := GenS9{F0: ("foo bar" + "\n\t")}
                                var v80 int32 = 863
                                _ = v79
                                _ = v80
                            }
                            _ = fi70
                            _ = v71
                        }
                    }
                }
                v5[v3] = ""
                _ = fi43
            }
            if true {
                v6 = v6
            } else {
                fmt.Println(v6, v24, v20)
            }
            _ = v42
        } else {
            if (!true) {
                var v81 uint32 = (v3 * 25)
                _ = v81
            }
            if false {
                var v82 uint64 = 2
                _ = v82
            } else {
                fmt.Println(v1, v6)
                var v83 float32 = 0.1
                v1 = v1
                _ = v83
            }
            fmt.Println(v1)
        }
    }
    var v84 string = "a"
    if ((v1 >> ((65535) & 63)) == v1) {
        var v85 bool = (18446744073709551615 <= 4294967295)
        _ = v85
    } else {
        v6 += (v6 << ((0) & 63))
    }
    if false {
        for fi86 := 0; fi86 < 15; fi86++ {
            switch 1 {
            case 32768:
                for fi87 := 0; fi87 < 17; fi87++ {
                    v3 -= v3
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = fi87
                }
            case 2147483647:
                for fi88 := 0; fi88 < 5; fi88++ {
                    v84 = (v84 + "hello")
                    v5[v3] = v84
                    _ = fi88
                }
            case -137:
                for fi89 := 0; fi89 < 11; fi89++ {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    fmt.Println(v1)
                    var v90 uint = 347
                    _ = fi89
                    _ = v90
                }
            default:
                v3 = (^(v3))
            }
            var v91 float64 = (0.1 - 0.0)
            _ = fi86
            _ = v91
        }
        v24 += -1
        v92 := GenS14{F0: 4294967295, F1: 1}
        _ = v92
    } else {
        v24 -= v24
        switch ((v84 + "a") + "hello") {
        case "":
            v93 := GenS12{F0: 572}
            _ = v93
        case "hello":
            v1 -= (v1 * v1)
        case "\n\t":
            v94 := GenS5{F0: v1, F1: -1, F2: (v24 ^ int64(v3))}
            _ = v94
        default:
            switch "" {
            case "":
                v6 -= -257
            default:
                fmt.Println(v24)
            }
        }
    }
    fmt.Println(v84, v3)
    if len(v4) > 0 {
        v4[0] = v3
    }
    var v95 float64 = float64(1e10)
    if (0 <= 541) {
        v1 = v1
        v96 := GenS10{F0: 0}
        v97 := GenS15{F0: "a", F1: v3, F2: ("\n\t" == v84)}
        _ = v96
        _ = v97
    } else {
        if false {
            var v98 string = v84
            _ = v98
        }
        switch v84 {
        case "foo bar":
            v99 := GenS13{F0: 4294967295, F1: v24}
            _ = v99
        case "hello":
            switch v84 {
            case "hello":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            default:
                var v100 float64 = (v95 / 0.1)
                _ = v100
            }
        default:
            v101 := GenS16{F0: -437, F1: (!true), F2: v20}
            _ = v101
        }
    }
    fmt.Println(v6, v6)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v84:", v84)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v84
    _ = v95
}
