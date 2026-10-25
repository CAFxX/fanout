package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 int32; F1 int; F2 int64 }
type GenS3 struct { F0 int32; F1 uint8 }
type GenS4 struct { F0 int32; F1 int8 }
type GenS5 struct { F0 int64; F1 uint }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int16; F1 bool; F2 uint64 }
type GenS9 struct { F0 int64 }
func genFn0(p0 uint8, p1 int16) int16 {
    v4 := GenS1{F0: (false || (0 == 4294967295)), F1: 32767}
    p1 -= -32768
    _ = p0
    _ = p1
    _ = v4
    return p1
}
func genFn1(p0 uint16, p1 uint8, p2 uint32) int32 {
    p0 -= (p0 * p0)
    v5 := GenS1{F0: (!(-2.25 >= (3.14159 + (3.14159 - 0.0)))), F1: 444}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 268
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]int)
    var v3 int64 = 9223372036854775807
    if (((!false) && ((3.14159 + 0.0) < float64(3.14159))) && (9223372036854775807 != 471)) {
        _ = len(v2)
        v1[1] = (42 <= 0)
    }
    switch ("" + ("" + "hello")) {
    case "\n\t":
        v6 := GenS2{F0: -2147483648, F1: 0, F2: (v3 * v3)}
        _ = v6
    case "foo bar":
        for fi7 := 0; fi7 < 4; fi7++ {
            v3 += v3
            _ = fi7
        }
    case "":
        v8 := []int{128, -1, 9223372036854775807, 32767}
        _ = v8
    }
    if ((v3 + v3) >= v3) {
        v3 = -812
        fmt.Println(v3)
        for fi9 := 0; fi9 < 4; fi9++ {
            v10 := []int{fi9, (fi9 / (((-(fi9))) | 1)), 4294967295}
            var v11 bool = true
            _ = fi9
            _ = v10
            _ = v11
        }
    } else {
        v3 += v3
        for fi12 := 0; fi12 < 6; fi12++ {
            v3 += (v3 / ((v3) | 1))
            fmt.Println(v3, fi12)
            v13 := GenS3{F0: genFn1(65535, 42, 0), F1: 42}
            _ = fi12
            _ = v13
        }
    }
    var v14 uint8 = 127
    for fi15 := 0; fi15 < 3; fi15++ {
        v14 -= (v14 << ((0) & 63))
        _ = fi15
    }
    var v16 int64 = v3
    v14 -= v14
    v17 := []int{268, -984, 520}
    if (v16 != v3) {
        var v18 uint = 0
        switch ((("hello" + (("\n\t" + "a") + "a")) + ("a" + "hello")) + "") {
        case "a":
            if (true && true) {
                v18 -= (v18 & v18)
                if (!(!false)) {
                    if false {
                        v16 -= v16
                    } else {
                        fmt.Println(v18, v16, v16)
                        var v19 int = 2147483647
                        _ = v19
                    }
                } else {
                    if true {
                        var v20 uint16 = 847
                        _ = v20
                    }
                    var v21 uint8 = uint8(v18)
                    var v22 int8 = 1
                    _ = v21
                    _ = v22
                }
            }
        case "hello":
            v23 := v2[65535]
            _ = v23
        }
        _ = v18
    }
    fmt.Println(v16, v16, v14)
    v24 := v1[2]
    if (v24 || v24) {
        fmt.Println(v3, v3, v16)
        if false {
            for fi25 := 0; fi25 < 1; fi25++ {
                var v26 int8 = -128
                for fi27 := 0; fi27 < 3; fi27++ {
                    v28 := GenS4{F0: 2147483647, F1: -2}
                    if ((!(v14 == 127)) || false) {
                        if (2 > 32767) {
                            var v29 float32 = 1.5
                            v26 += v26
                            v29 = v29
                            _ = v29
                        } else {
                            delete(v2, 0)
                        }
                    } else {
                        if false {
                            if ((v3 - 0) >= 42) {
                                var v30 int32 = 0
                                _ = v30
                            }
                            v31 := GenS5{F0: v16, F1: 18446744073709551615}
                            v26 += (v26 - (v26 * (v26 ^ v26)))
                            v32 := GenS4{F0: -1, F1: (v26 + v26)}
                            _ = v31
                            _ = v32
                        } else {
                            var v33 uint8 = 255
                            if len(v17) > 0 {
                                v17[0] = 32767
                            }
                            _ = v33
                        }
                        var v34 float32 = 3.14159
                        var v35 int16 = -223
                        fmt.Println(fi25, v3, v26)
                        _ = v34
                        _ = v35
                    }
                    _ = fi27
                    _ = v28
                }
                _ = fi25
                _ = v26
            }
        } else {
            v16 += v3
            delete(v1, 336)
        }
        switch -150 {
        case 9223372036854775807:
            for fi36 := 0; fi36 < 19; fi36++ {
                var v37 uint16 = 1
                if len(v17) >= 2 {
                    v17 = v17[1:len(v17)]
                }
                _ = fi36
                _ = v37
            }
        case 658:
            for fi38 := 0; fi38 < 4; fi38++ {
                if (65535 != 2) {
                    if v24 {
                        for fi39 := 0; fi39 < 7; fi39++ {
                            var v40 int64 = int64(fi39)
                            _ = fi39
                            _ = v40
                        }
                        v14 = (v14 - 38)
                        for fi41 := 0; fi41 < 20; fi41++ {
                            v3 = v3
                            switch "a" {
                            case "\n\t":
                                var v42 int32 = 181
                                _ = v42
                            case "a":
                                var v43 int = 128
                                _ = v43
                            case "hello":
                                v44 := GenS6{F0: 2}
                                _ = v44
                            }
                            _ = fi41
                        }
                    } else {
                        for fi45 := 0; fi45 < 14; fi45++ {
                            var v46 uint8 = 2
                            _ = fi45
                            _ = v46
                        }
                        var v47 int = (^(fi38))
                        var v48 uint64 = 996
                        _ = v47
                        _ = v48
                    }
                    for fi49 := 0; fi49 < 20; fi49++ {
                        switch fi49 {
                        case 32768:
                            var v50 int = fi38
                            _ = v50
                        default:
                            var v51 bool = (!false)
                            _ = v51
                        }
                        var v52 int64 = (v3 & -9223372036854775808)
                        if (27 == 2) {
                            v3 -= v3
                            switch (^(fi38)) {
                            case -774:
                                v24 = v24
                            }
                            var v53 uint32 = 42
                            fmt.Println(v16, v53)
                            v53 = 0
                            _ = v53
                        } else {
                            v54 := v1[2147483648]
                            _ = v54
                        }
                        _ = fi49
                        _ = v52
                    }
                }
                _ = fi38
            }
        case -585:
            v24 = (0.0 < 3.14159)
        }
    } else {
        _ = len(v1)
        if true {
            var v55 uint16 = 32767
            _ = v55
        } else {
            v14 = v14
            if (!((888 < 18446744073709551615) || v24)) {
                switch 32767 {
                case -539:
                    v3 -= 1
                case 4294967295:
                    var v56 uint64 = 266
                    _ = v56
                case 540:
                    for fi57 := 0; fi57 < 13; fi57++ {
                        var v58 bool = v24
                        _ = fi57
                        _ = v58
                    }
                }
            } else {
                if len(v17) > 0 {
                    v17[0] = 42
                }
                for fi59 := 0; fi59 < 12; fi59++ {
                    for fi60 := 0; fi60 < 13; fi60++ {
                        var v61 uint16 = 42
                        _ = fi60
                        _ = v61
                    }
                    var v62 bool = (false || ((v14 == 246) && (2147483647 == 1)))
                    v14 = v14
                    _ = fi59
                    _ = v62
                }
            }
            var v63 uint = 900
            _ = v63
        }
    }
    fmt.Println(v16, v16, v24)
    var v64 int8 = 42
    v64 = v64
    if len(v17) > 0 {
        v17[0] = int(v16)
    }
    if (32768 >= -1) {
        if false {
            var v65 int16 = -1
            var v66 int64 = (-(v3))
            _ = v65
            _ = v66
        }
        switch ("\n\t" + "\n\t") {
        case "hello":
            for fi67 := 0; fi67 < 4; fi67++ {
                v16 += v3
                _ = fi67
            }
        case "a":
            if v24 {
                switch "hello" {
                case "a":
                    switch 127 {
                    case -767:
                        if (-1 >= 1000) {
                            v3 = v3
                            var v68 int8 = (^(v64))
                            _ = v68
                        } else {
                            v64 = (-(v64))
                        }
                    }
                case "\n\t":
                    if len(v17) >= 2 {
                        v17 = v17[1:len(v17)]
                    }
                default:
                    switch "a" {
                    case "\n\t":
                        if len(v17) >= 2 {
                            v17 = v17[1:len(v17)]
                        }
                    }
                }
                if len(v17) > 0 {
                    v17[0] = 32768
                }
            } else {
                var v69 int8 = (v64 * v64)
                v70 := GenS7{F0: -376}
                _ = v69
                _ = v70
            }
        case "foo bar":
            if v24 {
                switch -244 {
                case 32767:
                    switch "foo bar" {
                    case "":
                        if (1 == 42) {
                            var v71 uint = 570
                            switch -225 {
                            case -1:
                                _ = len(v17)
                            case 127:
                                v72 := GenS1{F0: v24, F1: -222}
                                _ = v72
                            }
                            delete(v1, 0)
                            _ = v71
                        }
                    case "a":
                        if ("\n\t" != (("" + "") + "foo bar")) {
                            var v73 string = "foo bar"
                            _ = v73
                        } else {
                            var v74 bool = (v24 && v24)
                            fmt.Println(v3)
                            v24 = v24
                            _ = v74
                        }
                    default:
                        v75 := GenS8{F0: 732, F1: (!(uint32(v14) == 2)), F2: 18446744073709551615}
                        _ = v75
                    }
                case 42:
                    v76 := GenS9{F0: 9223372036854775807}
                    _ = v76
                case 9223372036854775807:
                    if len(v17) >= 2 {
                        v17 = v17[1:len(v17)]
                    }
                default:
                    for fi77 := 0; fi77 < 4; fi77++ {
                        fmt.Println(v14)
                        for fi78 := 0; fi78 < 16; fi78++ {
                            v17 = append(v17, (fi78 ^ (fi78 & (-(fi77)))))
                            _ = fi78
                        }
                        _ = fi77
                    }
                }
            } else {
                v16 = -9223372036854775808
                if v24 {
                    var v79 int8 = v64
                    v16 += int64(v14)
                    fmt.Println(v79, v3, v64)
                    _ = v79
                }
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v14
    _ = v16
    _ = v17
    _ = v2
    _ = v24
    _ = v3
    _ = v64
}
