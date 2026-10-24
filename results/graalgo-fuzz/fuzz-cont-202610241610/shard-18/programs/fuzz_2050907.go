package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 uint }
type GenS2 struct { F0 int; F1 int64 }
type GenS3 struct { F0 int8; F1 int }
type GenS4 struct { F0 uint8; F1 int16; F2 int }
type GenS5 struct { F0 int64; F1 int16; F2 uint16 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint64; F1 string }
type GenS8 struct { F0 int8 }

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 bool = (2 < 65535)
    v3 := []string{"a"}
    v4 := []string{"\n\t", "a"}
    switch "a" {
    case "\n\t":
        if len(v3) > 0 {
            v3[0] = "hello"
        }
    case "a":
        for fi5 := 0; fi5 < 2; fi5++ {
            fmt.Println(fi5, v1)
            _ = len(v4)
            _ = fi5
        }
    default:
        if ((((("hello" + ("" + ("foo bar" + ""))) + ("hello" + "\n\t")) + "") >= ("" + "\n\t")) || (18446744073709551615 == 0)) {
            if v2 {
                v2 = (!(3.14159 > -2.25))
            } else {
                v3 = append(v3, ("hello" + ("hello" + "hello")))
            }
            if len(v4) > 0 {
                v4[0] = "foo bar"
            }
            v2 = (-201 <= -9223372036854775808)
        } else {
            v2 = true
            v1 -= 936
        }
    }
    var v6 uint32 = 18
    fmt.Println(v2, v2)
    if len(v3) > 0 {
        v3[0] = ("\n\t" + "foo bar")
    }
    for fi7 := 0; fi7 < 8; fi7++ {
        for fi8 := 0; fi8 < 20; fi8++ {
            var v9 uint64 = (v1 << ((949) & 63))
            if v2 {
                if len(v4) > 0 {
                    v4[0] = ""
                }
            } else {
                fmt.Println(v6, fi7)
                if (!true) {
                    switch fi7 {
                    case 4294967295:
                        v1 -= (v9 & (^(v1)))
                    case -2147483648:
                        v3 = append(v3, ("a" + (("" + "hello") + "hello")))
                    case 1:
                        var v10 uint64 = v9
                        _ = v10
                    default:
                        var v11 int16 = 1
                        _ = v11
                    }
                    if v2 {
                        var v12 int16 = 1
                        var v13 bool = false
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = v12
                        _ = v13
                    } else {
                        if len(v4) > 0 {
                            v4[0] = ""
                        }
                        switch (-(fi8)) {
                        case 32768:
                            v14 := GenS1{F0: "\n\t", F1: 291, F2: 2}
                            _ = v14
                        default:
                            if (!v2) {
                                v2 = (1 <= 1)
                                v9 -= 407
                            } else {
                                v2 = v2
                                if len(v4) > 0 {
                                    v4[0] = "hello"
                                }
                                var v15 int64 = 1
                                _ = v15
                            }
                        }
                    }
                    if (!v2) {
                        if false {
                            v16 := GenS2{F0: fi7, F1: -212}
                            _ = v16
                        } else {
                            var v17 int8 = 0
                            _ = v17
                        }
                        var v18 uint = 252
                        _ = v18
                    } else {
                        if len(v4) > 0 {
                            v4[0] = ""
                        }
                        var v19 int8 = 111
                        _ = v19
                    }
                    var v20 int64 = 530
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = v20
                } else {
                    var v21 int = (fi7 >> ((1) & 63))
                    _ = v21
                }
                var v22 int8 = -118
                v1 -= (v1 - (v9 - v9))
                _ = v22
            }
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    v1 += 963
    var v23 float64 = 0.0
    var v24 int = -2147483648
    for fi25 := 0; fi25 < 2; fi25++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi25
    }
    if false {
        if ("\n\t" > "hello") {
            v26 := GenS1{F0: "foo bar", F1: uint32(v1), F2: 65535}
            _ = v26
        } else {
            var v27 int16 = 0
            for fi28 := 0; fi28 < 5; fi28++ {
                var v29 uint16 = 138
                if len(v3) > 0 {
                    v3[0] = "foo bar"
                }
                _ = fi28
                _ = v29
            }
            v30 := GenS3{F0: 100, F1: 9223372036854775807}
            _ = v27
            _ = v30
        }
    } else {
        fmt.Println(v23)
        if v2 {
            v4 = append(v4, "a")
            v1 -= v1
            v3 = append(v3, ("foo bar" + "hello"))
        } else {
            fmt.Println(v1, v2)
        }
        var v31 uint16 = 65535
        _ = v31
    }
    var v32 uint = 362
    for fi33 := 0; fi33 < 12; fi33++ {
        if true {
            v32 = (v32 * v32)
            if v2 {
                fmt.Println(v24)
            } else {
                var v34 bool = v2
                for fi35 := 0; fi35 < 3; fi35++ {
                    fmt.Println(v23, v1)
                    v24 = -2147483648
                    _ = fi35
                }
                _ = v34
            }
        } else {
            var v36 uint16 = 0
            v37 := GenS4{F0: 127, F1: 344, F2: (fi33 >> (((v32 >> (((v32 | (v32 << (((v32 & 1000)) & 63)))) & 63))) & 63))}
            v6 -= (v6 & v6)
            _ = v36
            _ = v37
        }
        v3 = append(v3, ("hello" + ""))
        _ = fi33
    }
    v38 := GenS2{F0: (v24 + v24), F1: 9223372036854775807}
    switch (("\n\t" + ("foo bar" + "a")) + "") {
    case "foo bar":
        if true {
            v39 := GenS5{F0: int64(v24), F1: -32768, F2: 504}
            _ = v39
        } else {
            for fi40 := 0; fi40 < 5; fi40++ {
                for fi41 := 0; fi41 < 8; fi41++ {
                    switch "foo bar" {
                    case "a":
                        v42 := GenS6{F0: int64(fi41)}
                        _ = v42
                    default:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    v2 = v2
                    v3 = append(v3, "")
                    _ = fi41
                }
                switch ("\n\t" + "hello") {
                case "\n\t":
                    v1 = v1
                case "foo bar":
                    _ = len(v3)
                case "hello":
                    if (false || true) {
                        v43 := GenS6{F0: int64(v1)}
                        for fi44 := 0; fi44 < 6; fi44++ {
                            if (-2.25 >= -2.25) {
                                var v45 uint = 0
                                if ((!(-32768 != -878)) && v2) {
                                    fmt.Println(v6, v24)
                                    switch ("a" + ("a" + "foo bar")) {
                                    case "":
                                        v32 -= (v45 + (v45 ^ v45))
                                    case "a":
                                        if v2 {
                                            fmt.Println(v32, v45, v24)
                                        } else {
                                            var v46 uint = (v32 - (v32 & v45))
                                            var v47 int = (fi44 - 32768)
                                            _ = v46
                                            _ = v47
                                        }
                                    case "hello":
                                        v23 -= 1.5
                                    default:
                                        var v48 uint = (^(v32))
                                        _ = v48
                                    }
                                }
                                _ = v45
                            } else {
                                v6 += (v6 & v6)
                                _ = len(v3)
                                if ((v2 || ((v23 / (v23 - float64(0.1))) < v23)) && v2) {
                                    if (uint32(v32) >= 65535) {
                                        if false {
                                            var v49 int8 = -128
                                            fmt.Println(v32, v32)
                                            switch fi40 {
                                            case 127:
                                                var v50 int32 = 42
                                                _ = v50
                                            case 32767:
                                                var v51 int8 = v49
                                                _ = v51
                                            }
                                            _ = v49
                                        }
                                        v52 := GenS6{F0: 9223372036854775807}
                                        var v53 int = (fi44 + (v24 - fi44))
                                        _ = v52
                                        _ = v53
                                    }
                                }
                            }
                            var v54 float64 = (1.5 / 0.1)
                            _ = fi44
                            _ = v54
                        }
                        _ = v43
                    } else {
                        switch v24 {
                        case 654:
                            v32 -= (v32 << (((v32 / (((v32 & 762)) | 1))) & 63))
                        }
                        for fi55 := 0; fi55 < 7; fi55++ {
                            var v56 uint = (v32 - (v32 / ((832) | 1)))
                            v24 -= v24
                            v1 = v1
                            _ = fi55
                            _ = v56
                        }
                    }
                default:
                    for fi57 := 0; fi57 < 1; fi57++ {
                        var v58 uint8 = 2
                        var v59 int64 = -1
                        _ = fi57
                        _ = v58
                        _ = v59
                    }
                }
                if (65535 <= 153) {
                    v2 = true
                    v2 = v2
                    for fi60 := 0; fi60 < 6; fi60++ {
                        v61 := GenS2{F0: 127, F1: -9223372036854775808}
                        v32 = v32
                        _ = fi60
                        _ = v61
                    }
                } else {
                    for fi62 := 0; fi62 < 6; fi62++ {
                        v63 := GenS3{F0: 112, F1: fi62}
                        var v64 uint16 = 2
                        _ = fi62
                        _ = v63
                        _ = v64
                    }
                }
                _ = fi40
            }
            var v65 uint16 = 0
            _ = v65
        }
    default:
        v66 := GenS4{F0: 69, F1: 1, F2: (v24 * 32768)}
        _ = v66
    }
    v1 += v1
    if ((v24 <= (v24 & v24)) || false) {
        v67 := GenS5{F0: 605, F1: -1, F2: 585}
        switch ("\n\t" + "foo bar") {
        case "hello":
            var v68 int32 = -2147483648
            _ = v68
        default:
            var v69 string = (("a" + ("hello" + (((("hello" + (("" + (("a" + "a") + "a")) + "\n\t")) + "\n\t") + "hello") + ((("foo bar" + "\n\t") + "foo bar") + "")))) + "a")
            _ = v69
        }
        var v70 bool = v2
        _ = v67
        _ = v70
    }
    v1 = (^(v1))
    if ("\n\t" != ((("" + "a") + "") + "a")) {
        _ = len(v4)
    } else {
        var v71 uint32 = (v6 + 485)
        if true {
            v2 = (v71 <= v6)
            v72 := GenS1{F0: ("hello" + ""), F1: v71, F2: (^(v32))}
            _ = v72
        } else {
            var v73 uint = v32
            v74 := GenS6{F0: -9223372036854775808}
            _ = v73
            _ = v74
        }
        if ((^(v6)) == v6) {
            for fi75 := 0; fi75 < 9; fi75++ {
                var v76 uint32 = v6
                for fi77 := 0; fi77 < 13; fi77++ {
                    v78 := GenS7{F0: 18446744073709551615, F1: (((((((("foo bar" + "") + "foo bar") + ("a" + "a")) + "a") + "hello") + "hello") + (("" + ("" + "a")) + "")) + ("foo bar" + "a"))}
                    v6 += v76
                    if (v2 || v2) {
                        fmt.Println(v6, v23)
                    } else {
                        var v79 float32 = 1.5
                        _ = len(v4)
                        switch "\n\t" {
                        case "hello":
                            var v80 int32 = 2147483647
                            _ = v80
                        case "foo bar":
                            if (3.14159 != v79) {
                                v76 -= v71
                                var v81 int64 = 0
                                _ = v81
                            } else {
                                v79 = v79
                            }
                        default:
                            fmt.Println(fi75)
                        }
                        _ = v79
                    }
                    switch "a" {
                    case "foo bar":
                        fmt.Println(v24, v76)
                    default:
                        v82 := GenS8{F0: 42}
                        _ = v82
                    }
                    v71 = (^(v71))
                    _ = fi77
                    _ = v78
                }
                _ = fi75
                _ = v76
            }
            if (float64((1e10 / (0.1 * 3.14159))) != 1e10) {
                var v83 uint8 = 164
                switch v24 {
                case 32768:
                    var v84 uint64 = (v1 << ((v32) & 63))
                    _ = v84
                default:
                    switch 128 {
                    case -2147483648:
                        var v85 uint64 = (^(v1))
                        _ = v85
                    }
                }
                _ = v83
            } else {
                if (((^(v32)) == v32) && (1 != -227)) {
                    v86 := GenS7{F0: v1, F1: "foo bar"}
                    _ = v86
                }
                if (737 <= 255) {
                    v1 = 2
                    for fi87 := 0; fi87 < 10; fi87++ {
                        var v88 int32 = 2147483647
                        v6 += v6
                        for fi89 := 0; fi89 < 9; fi89++ {
                            v71 += 2147483647
                            _ = fi89
                        }
                        _ = fi87
                        _ = v88
                    }
                    for fi90 := 0; fi90 < 15; fi90++ {
                        var v91 uint = (v32 * v32)
                        _ = fi90
                        _ = v91
                    }
                }
            }
        }
        _ = v71
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v32
    _ = v38
    _ = v4
    _ = v6
}
