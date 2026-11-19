package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 int32; F1 bool; F2 int64 }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 uint64; F1 int16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint; F1 int; F2 string }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint16; F1 uint }
type GenS9 struct { F0 int32; F1 uint }
type GenS10 struct { F0 int8; F1 int; F2 int16 }
type GenS11 struct { F0 int64; F1 uint32 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 bool = (!true)
    var v3 bool = (((-2.25 * float64(3.14159)) + 1e10) > 0.0)
    v4 := []int{42, 879}
    v3 = (996 < 2)
    for fi5 := 0; fi5 < 10; fi5++ {
        switch "hello" {
        case "hello":
            if false {
                v6 := GenS1{F0: 32767, F1: false}
                v2 = false
                v7 := GenS2{F0: -835, F1: false, F2: -9223372036854775808}
                _ = v6
                _ = v7
            } else {
                switch 135 {
                case 630:
                    var v8 string = "foo bar"
                    _ = v8
                case 491:
                    v2 = (681 <= 806)
                }
                for fi9 := 0; fi9 < 2; fi9++ {
                    var v10 bool = v3
                    _ = len(v4)
                    fmt.Println(fi9, fi5, v2)
                    _ = fi9
                    _ = v10
                }
            }
        default:
            v3 = true
        }
        _ = fi5
    }
    if (-9223372036854775808 > 952) {
        switch ((("hello" + "foo bar") + "\n\t") + ("hello" + ((((("foo bar" + (("" + ("foo bar" + "hello")) + "foo bar")) + "foo bar") + "\n\t") + ("hello" + "a")) + ("foo bar" + ("" + "a"))))) {
        case "foo bar":
            v3 = v2
        case "hello":
            if (float32(-2.25) <= ((((-2.25 / 1.0) + (0.1 * ((0.1 - (float32((0.1 * ((float64(0.0) * 3.14159) * -2.25))) * 0.0)) + -2.25))) / 1e10) / (0.1 / 3.14159))) {
                for fi11 := 0; fi11 < 1; fi11++ {
                    for fi12 := 0; fi12 < 12; fi12++ {
                        v2 = (!((float32((3.14159 + -2.25)) * 3.14159) == 0.0))
                        var v13 int32 = 418
                        _ = fi12
                        _ = v13
                    }
                    v4 = append(v4, fi11)
                    _ = fi11
                }
            }
        }
        v14 := GenS1{F0: 0, F1: (18446744073709551615 > 18446744073709551615)}
        var v15 float32 = 0.0
        _ = v14
        _ = v15
    } else {
        delete(v1, 4294967295)
        var v16 int8 = 42
        _ = len(v1)
        _ = v16
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v17 := GenS2{F0: 1, F1: false, F2: 1}
    v2 = (197 <= 42)
    v2 = (true || v3)
    v2 = (!(415 > 285))
    v3 = v3
    if v2 {
        v18 := GenS3{F0: -1, F1: 0}
        v2 = (0.1 < 3.14159)
        if v2 {
            v19 := GenS4{F0: 18446744073709551615, F1: -1}
            if (4294967295 > 1) {
                switch ("" + ("a" + "a")) {
                case "foo bar":
                    var v20 int16 = 784
                    _ = v20
                case "\n\t":
                    v3 = (!(true && true))
                default:
                    var v21 bool = false
                    _ = v21
                }
                if (!(0 > 0)) {
                    for fi22 := 0; fi22 < 12; fi22++ {
                        var v23 uint8 = 255
                        delete(v1, 233)
                        _ = fi22
                        _ = v23
                    }
                    switch "" {
                    case "a":
                        for fi24 := 0; fi24 < 17; fi24++ {
                            if len(v4) > 0 {
                                v4[0] = (fi24 | fi24)
                            }
                            var v25 uint8 = 35
                            v26 := GenS5{F0: -324}
                            _ = fi24
                            _ = v25
                            _ = v26
                        }
                    case "":
                        v27 := GenS6{F0: 1, F1: 128, F2: "a"}
                        _ = v27
                    case "foo bar":
                        if v3 {
                            v28 := GenS6{F0: 667, F1: 1, F2: ""}
                            _ = v28
                        }
                    default:
                        var v29 int32 = 2147483647
                        _ = v29
                    }
                    v30 := GenS3{F0: 1, F1: 18446744073709551615}
                    _ = v30
                } else {
                    for fi31 := 0; fi31 < 16; fi31++ {
                        var v32 uint64 = 1
                        if (2147483648 == 1) {
                            v2 = v2
                            switch "hello" {
                            case "a":
                                var v33 float64 = 1e10
                                _ = v33
                            case "hello":
                                if true {
                                    for fi34 := 0; fi34 < 10; fi34++ {
                                        if v3 {
                                            fmt.Println(v32)
                                            v32 = (v32 - (v32 << ((2) & 63)))
                                            switch "\n\t" {
                                            case "hello":
                                                if (((v32 <= (v32 % (((v32 >> ((18446744073709551615) & 63))) | 1))) && (32767 != 124)) || (-1 == 0)) {
                                                    v35 := GenS3{F0: -102, F1: 4294967295}
                                                    var v36 string = (("" + "") + "\n\t")
                                                    v37 := GenS1{F0: -32768, F1: (!false)}
                                                    _ = v35
                                                    _ = v36
                                                    _ = v37
                                                }
                                                switch "hello" {
                                                case "hello":
                                                    switch fi34 {
                                                    case -2147483648:
                                                        var v38 int = -9223372036854775808
                                                        _ = v38
                                                    case -9223372036854775808:
                                                        v1[42] = 688
                                                    case 1:
                                                        switch (fi31 % (((fi31 & (fi31 / ((fi34) | 1)))) | 1)) {
                                                        case 695:
                                                            v39 := GenS7{F0: 255}
                                                            _ = v39
                                                        case 736:
                                                            var v40 string = "\n\t"
                                                            _ = v40
                                                        default:
                                                            var v41 uint64 = (v32 << ((21) & 63))
                                                            _ = v41
                                                        }
                                                    }
                                                case "\n\t":
                                                    switch "a" {
                                                    case "":
                                                        v42 := GenS2{F0: 42, F1: (342 != 642), F2: 192}
                                                        var v43 int16 = 0
                                                        var v44 uint8 = 1
                                                        _ = v42
                                                        _ = v43
                                                        _ = v44
                                                    case "\n\t":
                                                        v45 := v1[483]
                                                        _ = v45
                                                    default:
                                                        switch "hello" {
                                                        case "a":
                                                            var v46 float64 = 0.0
                                                            _ = v46
                                                        case "foo bar":
                                                            v47 := GenS8{F0: 653, F1: 2}
                                                            _ = v47
                                                        default:
                                                            if v3 {
                                                                v48 := GenS9{F0: -2147483648, F1: 0}
                                                                var v49 uint = 65535
                                                                var v50 float64 = 1.5
                                                                switch "hello" {
                                                                case "a":
                                                                    fmt.Println(v32, v32, v49)
                                                                case "foo bar":
                                                                    if len(v4) > 0 {
                                                                        v4[0] = (fi34 * 773)
                                                                    }
                                                                case "":
                                                                    v1[1] = 65535
                                                                default:
                                                                    if v3 {
                                                                        var v51 string = "foo bar"
                                                                        switch (fi34 ^ (fi31 - (fi34 / ((fi34) | 1)))) {
                                                                        case 0:
                                                                            fmt.Println(v2)
                                                                        case 32768:
                                                                            var v52 float64 = -2.25
                                                                            _ = v52
                                                                        case -916:
                                                                            if (v2 || false) {
                                                                                var v53 int64 = int64(fi34)
                                                                                _ = v53
                                                                            } else {
                                                                                v2 = ((42 != 2147483647) || false)
                                                                                v54 := GenS8{F0: 65535, F1: (v49 * (v49 << (((v49 << ((v49) & 63))) & 63)))}
                                                                                var v55 uint64 = uint64(v49)
                                                                                _ = v54
                                                                                _ = v55
                                                                            }
                                                                        }
                                                                        _ = v51
                                                                    } else {
                                                                        fmt.Println(v3)
                                                                        v32 = (^(v32))
                                                                        var v56 string = "\n\t"
                                                                        v49 += 4294967295
                                                                        _ = v56
                                                                    }
                                                                }
                                                                _ = v48
                                                                _ = v49
                                                                _ = v50
                                                            }
                                                            var v57 uint16 = 589
                                                            _ = v57
                                                        }
                                                    }
                                                case "foo bar":
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                }
                                            case "foo bar":
                                                fmt.Println(v32, fi34, fi34)
                                            }
                                        } else {
                                            if (true && (v2 && (42 <= 0))) {
                                                var v58 uint = 255
                                                var v59 int64 = 9223372036854775807
                                                _ = v58
                                                _ = v59
                                            }
                                        }
                                        var v60 int = fi34
                                        v61 := GenS10{F0: 127, F1: fi31, F2: -538}
                                        v2 = (v2 && (false || (!false)))
                                        v32 -= 4294967295
                                        _ = fi34
                                        _ = v60
                                        _ = v61
                                    }
                                    var v62 uint = 925
                                    _ = v62
                                } else {
                                    v32 -= (v32 & 1)
                                }
                            case "foo bar":
                                v2 = v3
                            }
                            for fi63 := 0; fi63 < 4; fi63++ {
                                if v2 {
                                    v2 = ((!(217 != 43)) && v3)
                                } else {
                                    if ((v32 | (^(v32))) <= v32) {
                                        var v64 uint16 = 311
                                        _ = v64
                                    }
                                    v32 = 905
                                }
                                v65 := v1[2147483647]
                                if false {
                                    v65 = (v65 * 0)
                                }
                                _ = fi63
                                _ = v65
                            }
                        }
                        _ = fi31
                        _ = v32
                    }
                    v3 = (4294967295 <= 0)
                    v66 := GenS11{F0: -1, F1: 0}
                    _ = v66
                }
                if v3 {
                    var v67 int32 = -2147483648
                    _ = v67
                }
            }
            _ = v19
        } else {
            var v68 float64 = float64(((((3.14159 / 1.0) + 0.1) + float32(3.14159)) / (-2.25 - (0.0 - 1e10))))
            if v2 {
                for fi69 := 0; fi69 < 11; fi69++ {
                    v70 := GenS6{F0: 0, F1: -2147483648, F2: "hello"}
                    if (42 > 2) {
                        fmt.Println(v2, v3, v2)
                        switch fi69 {
                        case -2147483648:
                            fmt.Println(v68, v68, fi69)
                        case 0:
                            v71 := GenS7{F0: 255}
                            _ = v71
                        }
                        v4 = append(v4, fi69)
                    }
                    _ = fi69
                    _ = v70
                }
            } else {
                var v72 int = 1
                for fi73 := 0; fi73 < 17; fi73++ {
                    var v74 uint32 = 2
                    v75 := GenS7{F0: 621}
                    _ = fi73
                    _ = v74
                    _ = v75
                }
                var v76 string = (("\n\t" + "a") + "\n\t")
                _ = v72
                _ = v76
            }
            _ = v68
        }
        _ = v18
    }
    var v77 int8 = -38
    var v78 int16 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v77
    _ = v78
}
