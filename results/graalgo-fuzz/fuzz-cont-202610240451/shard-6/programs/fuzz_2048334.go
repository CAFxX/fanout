package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 int; F1 int16; F2 int32 }
type GenS3 struct { F0 int16; F1 int16 }
type GenS4 struct { F0 uint; F1 bool; F2 uint64 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 string; F1 uint32; F2 int32 }
type GenS8 struct { F0 string; F1 uint32 }
type GenS9 struct { F0 int64; F1 uint16 }
type GenS10 struct { F0 string; F1 string }
type GenS11 struct { F0 int16; F1 uint; F2 uint16 }

func main() {
    v1 := []uint32{}
    var v2 uint = 65535
    v3 := make(map[uint32]bool)
    var v4 int16 = 42
    var v5 int = -720
    var v6 float64 = (3.14159 - (1e10 * 0.1))
    var v7 uint64 = 1
    switch "" {
    case "\n\t":
        v8 := GenS1{F0: 42, F1: 1}
        _ = v8
    case "hello":
        var v9 uint64 = (v7 ^ 0)
        _ = v9
    default:
        if (!((128 > 128) || (126 > v4))) {
            v2 = v2
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            for fi10 := 0; fi10 < 1; fi10++ {
                v11 := GenS1{F0: 849, F1: 863}
                if (955 <= -9223372036854775808) {
                    switch "hello" {
                    case "":
                        fmt.Println(fi10, v2)
                    case "hello":
                        v5 += (v5 - fi10)
                    case "\n\t":
                        switch "" {
                        case "hello":
                            v7 += 42
                        case "a":
                            var v12 uint64 = 0
                            _ = v12
                        default:
                            for fi13 := 0; fi13 < 9; fi13++ {
                                if false {
                                    v6 -= v6
                                    v14 := GenS2{F0: 127, F1: int16(v7), F2: -1}
                                    v15 := GenS3{F0: 644, F1: (-(v4))}
                                    v16 := GenS4{F0: (^(v2)), F1: true, F2: (v7 << (((^(v2))) & 63))}
                                    _ = v14
                                    _ = v15
                                    _ = v16
                                }
                                v5 = fi10
                                var v17 string = "\n\t"
                                _ = fi13
                                _ = v17
                            }
                        }
                    }
                    _ = len(v3)
                    var v18 int16 = v4
                    _ = v18
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi10
                _ = v11
            }
        }
    }
    for fi19 := 0; fi19 < 11; fi19++ {
        var v20 float32 = (3.14159 * 1e10)
        for fi21 := 0; fi21 < 16; fi21++ {
            v7 -= (v7 | (v7 - 706))
            _ = fi21
        }
        if true {
            if true {
                v22 := GenS5{F0: 255}
                _ = v22
            } else {
                switch -347 {
                case -450:
                    for fi23 := 0; fi23 < 2; fi23++ {
                        switch v5 {
                        case 127:
                            v5 -= v5
                        }
                        _ = fi23
                    }
                case 32767:
                    var v24 int32 = -2147483648
                    _ = v24
                case -9223372036854775808:
                    v7 = (v7 + 18446744073709551615)
                default:
                    v5 = (fi19 & fi19)
                }
                if true {
                    v25 := GenS3{F0: (v4 << ((18446744073709551615) & 63)), F1: 464}
                    v26 := GenS6{F0: (fi19 | (v5 - (v5 | v5)))}
                    switch "foo bar" {
                    case "a":
                        if true {
                            if true {
                                v5 += (v5 << (((v2 - v2)) & 63))
                            }
                            v6 = float64((((1.5 * 1.5) * v20) + -2.25))
                        }
                    default:
                        v5 -= (-(fi19))
                    }
                    _ = v25
                    _ = v26
                } else {
                    for fi27 := 0; fi27 < 16; fi27++ {
                        v7 = v7
                        _ = fi27
                    }
                    v4 -= v4
                    v5 += (fi19 - (fi19 & v5))
                }
                if (3 >= 91) {
                    var v28 int32 = -735
                    v29 := GenS1{F0: 42, F1: 342}
                    _ = v28
                    _ = v29
                } else {
                    var v30 uint = v2
                    for fi31 := 0; fi31 < 4; fi31++ {
                        v5 = v5
                        v1 = append(v1, 198)
                        _ = fi31
                    }
                    _ = v30
                }
            }
            if true {
                var v32 float32 = v20
                if false {
                    for fi33 := 0; fi33 < 12; fi33++ {
                        var v34 float64 = 3.14159
                        v7 -= v7
                        _ = fi33
                        _ = v34
                    }
                }
                if (!false) {
                    v35 := GenS1{F0: 1, F1: 2147483648}
                    _ = v35
                }
                _ = v32
            }
            v4 += 32767
        }
        _ = fi19
        _ = v20
    }
    var v36 int16 = (v4 + v4)
    v37 := GenS6{F0: v5}
    var v38 int8 = -1
    v6 = (0.0 * v6)
    switch v5 {
    case -99:
        var v39 uint64 = (^(v7))
        _ = v39
    default:
        v6 += (v6 * float64(0.1))
    }
    for fi40 := 0; fi40 < 12; fi40++ {
        if true {
            if (!("hello" < "")) {
                var v41 uint64 = v7
                v4 = (v36 << ((42) & 63))
                _ = v41
            } else {
                var v42 float32 = 1.5
                v36 = (v4 % (((v36 & (v36 & (^(v36))))) | 1))
                switch ("hello" + "a") {
                case "foo bar":
                    if (!false) {
                        fmt.Println(v5, v6, v4)
                        v43 := GenS3{F0: 581, F1: v4}
                        _ = v43
                    } else {
                        var v44 int = 9223372036854775807
                        if (42 == 78) {
                            fmt.Println(v7, v6)
                            for fi45 := 0; fi45 < 19; fi45++ {
                                if (247 != 127) {
                                    v42 = v42
                                    if false {
                                        var v46 float32 = v42
                                        v1 = append(v1, 42)
                                        _ = v46
                                    }
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    if (0 <= 0) {
                                        if (!false) {
                                            delete(v3, 2147483648)
                                            v47 := GenS7{F0: "\n\t", F1: 4294967295, F2: -1}
                                            if true {
                                                v36 = v36
                                                switch "" {
                                                case "hello":
                                                    var v48 float64 = (1e10 / ((0.0 - (v6 - -2.25)) + v6))
                                                    _ = v48
                                                case "foo bar":
                                                    v7 = (v7 + 2)
                                                case "a":
                                                    v49 := GenS6{F0: fi45}
                                                    _ = v49
                                                default:
                                                    switch v5 {
                                                    case 32768:
                                                        fmt.Println(v7, v6, fi45)
                                                    case 665:
                                                        v50 := GenS3{F0: 34, F1: (v4 << (((v2 ^ v2)) & 63))}
                                                        _ = v50
                                                    case 2147483647:
                                                        var v51 uint8 = 91
                                                        _ = v51
                                                    }
                                                }
                                                if true {
                                                    var v52 uint16 = 42
                                                    if (2 >= (v2 | v2)) {
                                                        v53 := GenS2{F0: (fi40 << ((v2) & 63)), F1: v36, F2: 1}
                                                        switch (("a" + "\n\t") + (("hello" + "foo bar") + "foo bar")) {
                                                        case "":
                                                            var v54 uint16 = (v52 | v52)
                                                            _ = v54
                                                        case "foo bar":
                                                            fmt.Println(v2, v52)
                                                        }
                                                        _ = v53
                                                    } else {
                                                        var v55 int32 = 538
                                                        _ = v55
                                                    }
                                                    v56 := GenS1{F0: 2, F1: 65535}
                                                    _ = v52
                                                    _ = v56
                                                } else {
                                                    var v57 uint16 = uint16(v2)
                                                    _ = v57
                                                }
                                            } else {
                                                var v58 uint16 = 521
                                                var v59 uint16 = (v58 - uint16(v7))
                                                _ = v58
                                                _ = v59
                                            }
                                            _ = v47
                                        } else {
                                            v60 := GenS8{F0: "\n\t", F1: 42}
                                            var v61 float32 = (1.5 / float32(v6))
                                            _ = v60
                                            _ = v61
                                        }
                                        var v62 int8 = (v38 >> ((529) & 63))
                                        _ = v62
                                    } else {
                                        if (!(("hello" + ("foo bar" + "a")) > ((("foo bar" + ("a" + "a")) + "a") + "hello"))) {
                                            v5 += fi40
                                            switch (-(v44)) {
                                            case 128:
                                                v44 = 9223372036854775807
                                                var v63 int = v5
                                                v5 -= 685
                                                _ = v63
                                            case 47:
                                                var v64 int16 = (-(v36))
                                                _ = v64
                                            }
                                            var v65 uint8 = 127
                                            _ = v65
                                        }
                                    }
                                    var v66 int = (v44 / (((v5 + 32768)) | 1))
                                    _ = v66
                                }
                                fmt.Println(v36)
                                v3[2] = false
                                _ = fi45
                            }
                        }
                        v6 -= v6
                        _ = v44
                    }
                }
                _ = v42
            }
            v3[42] = true
        } else {
            if (v7 < (v7 - (v7 + v7))) {
                var v67 uint16 = 255
                _ = len(v1)
                var v68 uint = v2
                _ = v67
                _ = v68
            } else {
                switch ("\n\t" + ("hello" + "")) {
                case "a":
                    v6 -= (((0.0 / float64((0.1 * (3.14159 / 3.14159)))) * (float64(0.1) / 1e10)) * v6)
                case "":
                    var v69 string = "\n\t"
                    _ = v69
                default:
                    var v70 int = fi40
                    _ = v70
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            var v71 int8 = (v38 + (v38 << ((v2) & 63)))
            _ = v71
        }
        _ = fi40
    }
    if (-925 < 42) {
        if false {
            v36 -= v36
            fmt.Println(v4, v5)
            v4 = v36
        } else {
            fmt.Println(v2, v7, v2)
        }
        for fi72 := 0; fi72 < 18; fi72++ {
            switch -533 {
            case 399:
                for fi73 := 0; fi73 < 9; fi73++ {
                    fmt.Println(v7, v4)
                    var v74 int = (fi73 * fi72)
                    _ = fi73
                    _ = v74
                }
            default:
                v75 := GenS9{F0: 9223372036854775807, F1: 42}
                _ = v75
            }
            _ = fi72
        }
        v38 -= 127
    } else {
        var v76 int16 = v4
        fmt.Println(v5, v38)
        switch "\n\t" {
        case "hello":
            for fi77 := 0; fi77 < 14; fi77++ {
                var v78 string = ("\n\t" + "")
                _ = fi77
                _ = v78
            }
        case "foo bar":
            v79 := GenS2{F0: v5, F1: 1000, F2: -140}
            _ = v79
        case "\n\t":
            v80 := GenS10{F0: "a", F1: "foo bar"}
            _ = v80
        }
        _ = v76
    }
    v81 := GenS7{F0: "", F1: 671, F2: 0}
    v82 := GenS11{F0: v36, F1: v2, F2: 1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v81
    _ = v82
}
