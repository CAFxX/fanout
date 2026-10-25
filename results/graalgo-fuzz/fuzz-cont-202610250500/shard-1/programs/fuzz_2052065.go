package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 string; F1 int8 }
type GenS4 struct { F0 string; F1 int8; F2 uint }
type GenS5 struct { F0 int64; F1 int16; F2 bool }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 string; F1 int16 }
type GenS9 struct { F0 uint16; F1 int }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 uint8 }
type GenS12 struct { F0 int64; F1 uint }
type GenS13 struct { F0 int16 }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint8 = 127
    v3 := make(map[uint32]string)
    v4 := []uint32{495, 398}
    v5 := make(map[uint32]bool)
    if (!false) {
        v2 = (v2 ^ v2)
        fmt.Println(v2)
    } else {
        v6 := GenS1{F0: (^(v2)), F1: 713}
        v2 = 242
        switch -9223372036854775808 {
        case 4294967295:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        default:
            for fi7 := 0; fi7 < 12; fi7++ {
                var v8 uint16 = 0
                v2 += v2
                _ = fi7
                _ = v8
            }
        }
        _ = v6
    }
    switch -878 {
    case 127:
        v2 = v2
    case 301:
        if false {
            v9 := GenS1{F0: 20, F1: 760}
            _ = v9
        } else {
            v2 = v2
            v10 := GenS2{F0: false}
            _ = v10
        }
    }
    fmt.Println(v2)
    for fi11 := 0; fi11 < 15; fi11++ {
        switch "" {
        case "a":
            if (1 < 42) {
                v2 -= (v2 & 76)
                if true {
                    var v12 float32 = (3.14159 + 3.14159)
                    _ = v12
                } else {
                    var v13 uint8 = (v2 | (v2 / ((v2) | 1)))
                    _ = v13
                }
            } else {
                var v14 int = 42
                fmt.Println(v2, v2)
                fmt.Println(v14, fi11)
                _ = v14
            }
        case "\n\t":
            if (0.1 != 0.0) {
                v3[65535] = ((((((("a" + ("" + "")) + ("" + ("" + ("\n\t" + "a")))) + "\n\t") + "\n\t") + "foo bar") + "") + "hello")
                var v15 uint8 = v2
                _ = v15
            } else {
                if ((v2 ^ (v2 >> ((4294967295) & 63))) != 93) {
                    if true {
                        if false {
                            if len(v4) > 0 {
                                v4[0] = 1
                            }
                            if (32767 < 30) {
                                var v16 float32 = 1e10
                                fmt.Println(v16, v16)
                                _ = v16
                            }
                            fmt.Println(v2, v2)
                        } else {
                            v2 += v2
                            fmt.Println(fi11, fi11)
                            var v17 uint64 = 985
                            _ = v17
                        }
                        for fi18 := 0; fi18 < 13; fi18++ {
                            v19 := GenS2{F0: (213 < 1000)}
                            v2 = v2
                            _ = fi18
                            _ = v19
                        }
                    } else {
                        for fi20 := 0; fi20 < 5; fi20++ {
                            if (100 < 127) {
                                v2 += 42
                            }
                            _ = fi20
                        }
                        switch (fi11 - fi11) {
                        case 850:
                            if false {
                                var v21 int8 = 1
                                var v22 uint32 = 1
                                var v23 uint32 = 179
                                _ = v21
                                _ = v22
                                _ = v23
                            }
                        case -2147483648:
                            v2 -= 138
                        }
                        for fi24 := 0; fi24 < 16; fi24++ {
                            v4 = append(v4, 802)
                            _ = fi24
                        }
                    }
                } else {
                    if true {
                        var v25 bool = true
                        fmt.Println(v2)
                        _ = v25
                    }
                    fmt.Println(v2, v2)
                }
                for fi26 := 0; fi26 < 6; fi26++ {
                    v2 -= 69
                    var v27 int32 = 0
                    var v28 bool = true
                    _ = fi26
                    _ = v27
                    _ = v28
                }
            }
        case "":
            switch "foo bar" {
            case "\n\t":
                fmt.Println(fi11, v2)
            case "hello":
                switch fi11 {
                case 4294967295:
                    var v29 int32 = 89
                    _ = v29
                default:
                    v30 := GenS2{F0: true}
                    _ = v30
                }
            }
        default:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        for fi31 := 0; fi31 < 16; fi31++ {
            if ("\n\t" != "foo bar") {
                if false {
                    v32 := GenS2{F0: (351 <= 2)}
                    _ = v32
                }
            }
            if (516 < 4294967295) {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v2 = (v2 % ((v2) | 1))
                v33 := GenS3{F0: "foo bar", F1: 42}
                v2 = (v2 << ((636) & 63))
                v34 := GenS2{F0: (uint16(v2) > 42)}
                _ = v33
                _ = v34
            }
            var v35 bool = true
            _ = fi31
            _ = v35
        }
        var v36 float32 = 0.1
        _ = fi11
        _ = v36
    }
    v2 += 128
    for fi37 := 0; fi37 < 10; fi37++ {
        if ((-2.25 != (0.0 + 0.1)) || ((65535 < 65535) || ((v2 / ((v2) | 1)) >= v2))) {
            v5[913] = false
        }
        _ = fi37
    }
    v2 += (v2 / ((v2) | 1))
    switch 128 {
    case 42:
        v2 = (v2 >> ((2) & 63))
    case 506:
        v2 += 2
    }
    v38 := GenS4{F0: "hello", F1: -128, F2: 843}
    var v39 uint32 = 0
    var v40 int32 = -279
    v2 = v2
    for fi41 := 0; fi41 < 9; fi41++ {
        v40 = (v40 % (((v40 - (^(v40)))) | 1))
        v42 := GenS1{F0: 255, F1: 689}
        if (!true) {
            var v43 uint8 = (^(v2))
            _ = v43
        }
        _ = fi41
        _ = v42
    }
    if (0 != 1) {
        if (((true && true) && true) || ((v40 - 2147483647) <= 1)) {
            v44 := GenS5{F0: 42, F1: -861, F2: true}
            var v45 uint = 65535
            _ = len(v1)
            _ = v44
            _ = v45
        }
        for fi46 := 0; fi46 < 20; fi46++ {
            if false {
                var v47 int64 = 873
                if (0 != 0) {
                    v48 := GenS6{F0: 1}
                    _ = v48
                } else {
                    v49 := GenS7{F0: 196}
                    fmt.Println(v40, v2, v2)
                    _ = v49
                }
                v39 = (v39 * v39)
                _ = v47
            } else {
                v40 -= v40
                switch "a" {
                case "hello":
                    var v50 string = ""
                    _ = v50
                case "a":
                    v4 = append(v4, (^(v39)))
                case "":
                    v2 -= v2
                }
                for fi51 := 0; fi51 < 1; fi51++ {
                    switch 42 {
                    case 0:
                        v52 := GenS8{F0: (("foo bar" + ("hello" + "a")) + "foo bar"), F1: 42}
                        _ = v52
                    case 127:
                        v2 += 2
                    case 32767:
                        switch fi46 {
                        case -31:
                            var v53 uint = 4294967295
                            _ = v53
                        }
                        var v54 int16 = int16(v40)
                        var v55 uint = 2
                        _ = v54
                        _ = v55
                    default:
                        v2 -= 53
                    }
                    v2 += (v2 + (v2 >> ((65535) & 63)))
                    _ = fi51
                }
            }
            for fi56 := 0; fi56 < 6; fi56++ {
                switch "\n\t" {
                case "a":
                    fmt.Println(fi46, v39, v40)
                default:
                    switch ("\n\t" + ("foo bar" + "")) {
                    case "":
                        fmt.Println(v39, v40, fi46)
                    default:
                        v39 = 4294967295
                    }
                }
                _ = fi56
            }
            _ = fi46
        }
    } else {
        var v57 string = "foo bar"
        _ = v57
    }
    _ = len(v4)
    for fi58 := 0; fi58 < 15; fi58++ {
        for fi59 := 0; fi59 < 9; fi59++ {
            var v60 int64 = 1
            _ = fi59
            _ = v60
        }
        _ = fi58
    }
    v61 := GenS2{F0: true}
    for fi62 := 0; fi62 < 15; fi62++ {
        for fi63 := 0; fi63 < 2; fi63++ {
            var v64 string = "a"
            if (42 > 0) {
                v65 := GenS1{F0: v2, F1: 365}
                if (false && (false || true)) {
                    switch (fi63 >> ((18446744073709551615) & 63)) {
                    case -9223372036854775808:
                        fmt.Println(fi63)
                    case -1:
                        fmt.Println(v40)
                    case -466:
                        if true {
                            var v66 int16 = 1
                            switch v64 {
                            case "a":
                                var v67 int = (fi62 >> ((711) & 63))
                                _ = v67
                            case "hello":
                                fmt.Println(fi63)
                            case "foo bar":
                                v68 := GenS2{F0: (766 != 843)}
                                _ = v68
                            }
                            _ = v66
                        } else {
                            var v69 int = 9223372036854775807
                            switch fi63 {
                            case 42:
                                v64 = v64
                            case 1:
                                switch ("\n\t" + v64) {
                                case "a":
                                    v64 = (v64 + ("a" + (((((v64 + ("hello" + v64)) + "") + v64) + "") + (v64 + v64))))
                                }
                            default:
                                if ((!(true || true)) || ((!(v64 == v64)) && (42 > 18446744073709551615))) {
                                    if len(v4) > 0 {
                                        v4[0] = v39
                                    }
                                }
                            }
                            if ((84 >= 2) && (1000 == -1)) {
                                v70 := GenS9{F0: 1, F1: (fi62 & int(v40))}
                                if (93 < v40) {
                                    var v71 float32 = 3.14159
                                    _ = len(v1)
                                    fmt.Println(v71)
                                    _ = v71
                                }
                                v39 += 272
                                _ = v70
                            } else {
                                switch v69 {
                                case 4294967295:
                                    switch (fi63 * fi63) {
                                    case 594:
                                        v2 = v2
                                    case 2147483647:
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    case 1:
                                        v4 = append(v4, (v39 % ((v39) | 1)))
                                    default:
                                        v2 = (v2 & (^(v2)))
                                        var v72 int = (^(fi63))
                                        _ = len(v5)
                                        _ = v72
                                    }
                                default:
                                    if (((1e10 * ((-2.25 * (0.1 * 1.5)) - (1.5 / 1.5))) + float32(1.5)) == 0.0) {
                                        if ((0.0 * 0.1) < ((((0.0 * 1e10) / 0.1) * (-2.25 - 0.0)) / 1.0)) {
                                            var v73 int = v69
                                            var v74 int16 = 1
                                            _ = v73
                                            _ = v74
                                        }
                                        if (false || false) {
                                            v69 = (-(fi62))
                                            if ((1e10 <= 1e10) && (9223372036854775807 > 1)) {
                                                v75 := GenS10{F0: 2}
                                                _ = v75
                                            } else {
                                                _ = len(v5)
                                            }
                                        }
                                        v76 := GenS5{F0: 1, F1: 1, F2: false}
                                        _ = v76
                                    } else {
                                        var v77 uint64 = 18446744073709551615
                                        var v78 float32 = float32(1.5)
                                        v79 := GenS11{F0: v2}
                                        var v80 uint16 = 2
                                        v81 := GenS12{F0: -9223372036854775808, F1: 255}
                                        v82 := GenS12{F0: -711, F1: 2}
                                        _ = v77
                                        _ = v78
                                        _ = v79
                                        _ = v80
                                        _ = v81
                                        _ = v82
                                    }
                                }
                                v83 := GenS13{F0: 32767}
                                _ = len(v3)
                                _ = v83
                            }
                            _ = v69
                        }
                    }
                    if true {
                        switch (v64 + "") {
                        case "foo bar":
                            switch v64 {
                            case "a":
                                var v84 uint64 = 42
                                _ = v84
                            case "\n\t":
                                v85 := GenS6{F0: 658}
                                _ = v85
                            default:
                                var v86 float32 = float32(0.1)
                                _ = v86
                            }
                        case "\n\t":
                            delete(v1, v64)
                        default:
                            fmt.Println(v40, fi63, fi63)
                        }
                        fmt.Println(v39, v2, v39)
                    }
                    v64 = v64
                    var v87 string = ((v64 + v64) + v64)
                    _ = v87
                } else {
                    var v88 uint64 = 0
                    _ = v88
                }
                var v89 string = ("" + ("\n\t" + (v64 + ("foo bar" + "foo bar"))))
                var v90 uint = 1
                var v91 bool = (false && false)
                v92 := GenS2{F0: true}
                _ = v65
                _ = v89
                _ = v90
                _ = v91
                _ = v92
            } else {
                fmt.Println(v39, v64)
                if ((fi62 <= (fi62 + fi62)) || (false && true)) {
                    switch fi62 {
                    case -1:
                        v40 = (-(v40))
                    case 525:
                        v93 := GenS4{F0: v64, F1: -128, F2: 976}
                        _ = v93
                    case -235:
                        fmt.Println(v64, fi62, v64)
                    default:
                        v4 = append(v4, (v39 - (v39 << ((65535) & 63))))
                    }
                } else {
                    v39 -= 1
                }
            }
            if true {
                var v94 uint16 = 255
                _ = len(v1)
                _ = v94
            }
            _ = fi63
            _ = v64
        }
        var v95 uint32 = v39
        _ = fi62
        _ = v95
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v61
}
