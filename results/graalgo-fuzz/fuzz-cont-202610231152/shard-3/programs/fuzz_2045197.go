package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 int64; F2 int8 }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 int; F1 uint }
type GenS5 struct { F0 uint32; F1 bool; F2 uint64 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint8; F1 uint32; F2 uint8 }

func main() {
    v1 := []uint32{2147483647, 980, 443, 4294967295}
    var v2 bool = false
    var v3 uint8 = 89
    var v4 float64 = 1.5
    var v5 uint16 = 1
    var v6 int = 32768
    for fi7 := 0; fi7 < 4; fi7++ {
        v8 := GenS1{F0: 128}
        var v9 float32 = 3.14159
        _ = fi7
        _ = v8
        _ = v9
    }
    switch "\n\t" {
    case "":
        if v2 {
            if true {
                fmt.Println(v5, v2)
                for fi10 := 0; fi10 < 18; fi10++ {
                    var v11 float32 = ((0.1 * (3.14159 * 0.0)) / (0.1 - ((1e10 - 3.14159) + (3.14159 + (0.1 - (0.1 + (0.1 - 1.5)))))))
                    _ = fi10
                    _ = v11
                }
                var v12 uint16 = v5
                _ = v12
            } else {
                var v13 bool = v2
                if ("foo bar" <= ("hello" + ("a" + ""))) {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v14 := GenS1{F0: v6}
                    if false {
                        v3 += v3
                        v15 := GenS1{F0: (v6 | v6)}
                        _ = v15
                    } else {
                        for fi16 := 0; fi16 < 16; fi16++ {
                            if v2 {
                                _ = len(v1)
                                v6 += fi16
                                var v17 float64 = v4
                                _ = v17
                            } else {
                                switch -432 {
                                case -2147483648:
                                    if (!true) {
                                        var v18 uint16 = v5
                                        fmt.Println(v6)
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        _ = v18
                                    } else {
                                        v19 := GenS1{F0: (v6 + (fi16 & (v6 << ((42) & 63))))}
                                        _ = v19
                                    }
                                default:
                                    v4 = 1.5
                                }
                                _ = len(v1)
                            }
                            var v20 uint16 = (^(v5))
                            _ = fi16
                            _ = v20
                        }
                    }
                    _ = v14
                }
                switch v6 {
                case -9223372036854775808:
                    fmt.Println(v3, v13)
                case 9223372036854775807:
                    switch "hello" {
                    case "hello":
                        v4 += (float64(1.5) - 3.14159)
                    }
                case -644:
                    v21 := GenS1{F0: -1}
                    _ = v21
                }
                _ = v13
            }
            var v22 uint32 = 2
            for fi23 := 0; fi23 < 6; fi23++ {
                if len(v1) > 0 {
                    v1[0] = (v22 ^ (^(v22)))
                }
                _ = fi23
            }
            _ = v22
        }
    case "\n\t":
        v24 := GenS2{F0: ("" + ""), F1: -9223372036854775808, F2: 0}
        _ = v24
    default:
        for fi25 := 0; fi25 < 1; fi25++ {
            if len(v1) > 0 {
                v1[0] = 177
            }
            switch ("foo bar" + "a") {
            case "":
                for fi26 := 0; fi26 < 11; fi26++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v2 = false
                    var v27 uint32 = 2147483648
                    _ = fi26
                    _ = v27
                }
            case "foo bar":
                for fi28 := 0; fi28 < 4; fi28++ {
                    var v29 uint16 = (v5 + 1)
                    _ = len(v1)
                    _ = fi28
                    _ = v29
                }
            }
            _ = fi25
        }
    }
    v30 := GenS3{F0: 65535, F1: 1}
    if v2 {
        for fi31 := 0; fi31 < 15; fi31++ {
            var v32 float64 = (v4 - v4)
            v5 -= (v5 ^ (^(v5)))
            _ = fi31
            _ = v32
        }
        v3 = 127
        var v33 float32 = 1.5
        _ = v33
    }
    switch "foo bar" {
    case "":
        if len(v1) > 0 {
            v1[0] = 0
        }
    default:
        v34 := GenS1{F0: v6}
        _ = v34
    }
    if (v2 && v2) {
        switch v6 {
        case -645:
            if (-663 >= v6) {
                v3 -= v3
            } else {
                v2 = (!v2)
            }
        case 32767:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        default:
            if v2 {
                for fi35 := 0; fi35 < 9; fi35++ {
                    v4 = (((((3.14159 * v4) / float64(float32((1.5 + v4)))) / 1e10) + v4) * (v4 + v4))
                    var v36 int64 = -1
                    for fi37 := 0; fi37 < 14; fi37++ {
                        v2 = v2
                        v38 := GenS2{F0: "", F1: 329, F2: 1}
                        v5 -= 42
                        _ = fi37
                        _ = v38
                    }
                    _ = fi35
                    _ = v36
                }
                switch ("\n\t" + "\n\t") {
                case "":
                    var v39 int = (v6 | v6)
                    _ = v39
                case "hello":
                    v6 = v6
                default:
                    v2 = (42 == int16(v6))
                }
            }
        }
        for fi40 := 0; fi40 < 15; fi40++ {
            if false {
                if (18446744073709551615 < 757) {
                    var v41 uint16 = v5
                    var v42 uint64 = 2
                    _ = v41
                    _ = v42
                } else {
                    var v43 int8 = -1
                    var v44 uint32 = 4294967295
                    fmt.Println(v4, v4)
                    _ = v43
                    _ = v44
                }
            } else {
                v6 = v6
            }
            v45 := GenS2{F0: "", F1: 42, F2: -103}
            var v46 bool = true
            _ = fi40
            _ = v45
            _ = v46
        }
    } else {
        switch "hello" {
        case "a":
            switch v6 {
            case 4294967295:
                if (65535 == 20) {
                    v3 += v3
                    var v47 float32 = 0.1
                    if v2 {
                        var v48 int64 = 1
                        if ((v6 ^ (v6 % ((v6) | 1))) <= v6) {
                            v5 -= v5
                            var v49 string = (("" + ("foo bar" + "a")) + "foo bar")
                            v3 -= (v3 & (^(v3)))
                            _ = v49
                        }
                        _ = v48
                    }
                    _ = v47
                }
            case 180:
                v5 -= (v5 % (((v5 - 255)) | 1))
            case -626:
                if (v3 != 110) {
                    for fi50 := 0; fi50 < 18; fi50++ {
                        v51 := GenS3{F0: 893, F1: 32767}
                        _ = fi50
                        _ = v51
                    }
                    for fi52 := 0; fi52 < 6; fi52++ {
                        if true {
                            switch v6 {
                            case 9223372036854775807:
                                fmt.Println(v2)
                            case 127:
                                for fi53 := 0; fi53 < 9; fi53++ {
                                    v1 = append(v1, 37)
                                    if len(v1) > 0 {
                                        v1[0] = 2
                                    }
                                    _ = fi53
                                }
                            default:
                                v4 = float64(0.0)
                            }
                            v54 := GenS1{F0: v6}
                            _ = v54
                        } else {
                            if len(v1) > 0 {
                                v1[0] = 2
                            }
                            var v55 uint8 = 141
                            var v56 float64 = (v4 * (3.14159 + float64(0.1)))
                            _ = v55
                            _ = v56
                        }
                        switch "foo bar" {
                        case "":
                            v57 := GenS2{F0: "foo bar", F1: -1, F2: 0}
                            _ = v57
                        case "a":
                            v58 := GenS4{F0: fi52, F1: 1}
                            _ = v58
                        }
                        _ = len(v1)
                        _ = fi52
                    }
                    if v2 {
                        v59 := GenS4{F0: (v6 ^ v6), F1: 2}
                        if v2 {
                            switch v6 {
                            case 1:
                                _ = len(v1)
                            case -2147483648:
                                var v60 int64 = 1
                                _ = v60
                            }
                        } else {
                            var v61 float64 = ((v4 * v4) + v4)
                            _ = v61
                        }
                        v62 := GenS5{F0: 2147483647, F1: (2147483647 >= -2147483648), F2: 438}
                        _ = v59
                        _ = v62
                    }
                }
            default:
                var v63 int8 = 42
                _ = v63
            }
        case "":
            v64 := GenS2{F0: "hello", F1: 42, F2: -1}
            _ = v64
        }
    }
    switch "\n\t" {
    case "foo bar":
        v4 -= v4
    case "":
        var v65 int32 = -496
        _ = v65
    case "hello":
        var v66 int64 = 1
        _ = v66
    default:
        if false {
            var v67 int = (v6 & (v6 ^ (v6 * (-(v6)))))
            var v68 float32 = float32(v4)
            switch ("a" + "") {
            case "a":
                var v69 float64 = ((0.0 / v4) - (v4 + 3.14159))
                _ = v69
            case "\n\t":
                v2 = (0 == 729)
            default:
                v68 = 1.5
            }
            _ = v67
            _ = v68
        }
    }
    v70 := GenS6{F0: (v5 | (^(v5)))}
    for fi71 := 0; fi71 < 17; fi71++ {
        var v72 uint64 = 1
        _ = fi71
        _ = v72
    }
    v1 = append(v1, 759)
    switch (("" + ("" + ("" + "a"))) + (((("" + ("hello" + "a")) + ("foo bar" + "a")) + "hello") + "\n\t")) {
    case "hello":
        v4 = 1e10
    case "":
        var v73 bool = true
        _ = v73
    }
    fmt.Println(v3)
    for fi74 := 0; fi74 < 9; fi74++ {
        var v75 uint = 65535
        fmt.Println(v2, v2, v6)
        _ = fi74
        _ = v75
    }
    v2 = v2
    if false {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v76 := GenS2{F0: ("hello" + ("a" + "")), F1: 12, F2: 127}
        _ = v76
    } else {
        v77 := GenS7{F0: v3, F1: 152, F2: 2}
        var v78 int = v6
        var v79 uint8 = v3
        _ = v77
        _ = v78
        _ = v79
    }
    var v80 int16 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v70
    _ = v80
}
