package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 uint64 }
type GenS2 struct { F0 int8; F1 uint16; F2 string }
type GenS3 struct { F0 int64; F1 bool }
type GenS4 struct { F0 uint8; F1 int32 }
type GenS5 struct { F0 uint16; F1 string }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int32; F1 uint64 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int }

func main() {
    var v1 bool = true
    v2 := []string{}
    var v3 uint = 42
    var v4 int8 = -1
    v5 := make(map[uint32]bool)
    var v6 uint8 = 127
    var v7 int32 = 2147483647
    if len(v2) > 0 {
        v2[0] = "a"
    }
    if v1 {
        v1 = ((v7 ^ 0) != v7)
        if (0 != -600) {
            for fi8 := 0; fi8 < 20; fi8++ {
                v9 := GenS1{F0: (("" + ("a" + "hello")) + ("" + "hello")), F1: v4, F2: 42}
                fmt.Println(v7, v7)
                if (65535 == 32767) {
                    var v10 uint64 = 4294967295
                    _ = v10
                }
                _ = fi8
                _ = v9
            }
            if v1 {
                fmt.Println(v6, v6)
                var v11 uint16 = 470
                fmt.Println(v3, v11, v1)
                _ = v11
            }
            for fi12 := 0; fi12 < 19; fi12++ {
                v13 := GenS1{F0: (("hello" + "a") + ""), F1: (v4 >> ((18446744073709551615) & 63)), F2: 42}
                var v14 float64 = 0.1
                _ = fi12
                _ = v13
                _ = v14
            }
        }
    }
    if (1 > -32768) {
        var v15 uint16 = 964
        switch -2147483648 {
        case 357:
            if (121 < 2147483647) {
                if true {
                    fmt.Println(v4)
                    var v16 uint64 = 1
                    var v17 uint16 = v15
                    _ = v16
                    _ = v17
                }
                var v18 int32 = (v7 + (v7 | (v7 % (((v7 / (((v7 / ((v7) | 1))) | 1))) | 1))))
                var v19 float32 = ((0.0 - 1.5) / (0.1 * 1e10))
                _ = v18
                _ = v19
            } else {
                if v1 {
                    v5[65535] = v1
                }
            }
        case 128:
            v7 = (v7 | int32(v15))
        default:
            if v1 {
                var v20 uint = 31
                switch -2147483648 {
                case 4294967295:
                    var v21 int = 565
                    _ = v21
                }
                switch "\n\t" {
                case "":
                    v4 = v4
                default:
                    var v22 int64 = -9223372036854775808
                    _ = v22
                }
                _ = v20
            } else {
                if (!(0.0 < 3.14159)) {
                    var v23 int32 = v7
                    for fi24 := 0; fi24 < 7; fi24++ {
                        switch "foo bar" {
                        case "":
                            if (!v1) {
                                v25 := GenS2{F0: -128, F1: uint16(v3), F2: "hello"}
                                switch "a" {
                                case "":
                                    for fi26 := 0; fi26 < 4; fi26++ {
                                        var v27 string = "foo bar"
                                        var v28 string = v27
                                        if len(v2) > 0 {
                                            v2[0] = (v28 + ((v28 + "a") + v28))
                                        }
                                        var v29 float32 = 1e10
                                        _ = fi26
                                        _ = v27
                                        _ = v28
                                        _ = v29
                                    }
                                case "hello":
                                    v30 := v5[42]
                                    _ = v30
                                }
                                v31 := GenS1{F0: (("a" + "hello") + ((("\n\t" + "\n\t") + "foo bar") + "hello")), F1: v4, F2: 18446744073709551615}
                                _ = v25
                                _ = v31
                            } else {
                                v32 := GenS1{F0: ("a" + (("" + "foo bar") + "a")), F1: 42, F2: 298}
                                _ = v32
                            }
                        case "hello":
                            fmt.Println(v4, v6)
                        case "foo bar":
                            fmt.Println(v15)
                        default:
                            var v33 int16 = int16(v15)
                            _ = v33
                        }
                        switch "" {
                        case "hello":
                            var v34 uint16 = v15
                            _ = v34
                        case "a":
                            v3 = v3
                        case "":
                            for fi35 := 0; fi35 < 15; fi35++ {
                                v7 = (v7 | int32(v4))
                                v7 = (^(v23))
                                var v36 int64 = int64(v3)
                                _ = fi35
                                _ = v36
                            }
                        default:
                            var v37 int = fi24
                            _ = v37
                        }
                        var v38 int8 = v4
                        _ = fi24
                        _ = v38
                    }
                    _ = v23
                }
            }
        }
        fmt.Println(v6, v6, v6)
        _ = v15
    } else {
        if len(v2) > 0 {
            v2[0] = ""
        }
        switch 42 {
        case 32767:
            if true {
                v39 := v5[65535]
                fmt.Println(v1)
                fmt.Println(v1)
                _ = v39
            }
        case 702:
            v40 := GenS3{F0: 930, F1: ((v1 && (!(1e10 != (1e10 / (1.5 / 1.5))))) || (!("foo bar" != "foo bar")))}
            _ = v40
        default:
            for fi41 := 0; fi41 < 12; fi41++ {
                if ((("\n\t" + "hello") != "foo bar") && (v1 || (v1 && v1))) {
                    for fi42 := 0; fi42 < 16; fi42++ {
                        switch "hello" {
                        case "hello":
                            switch "" {
                            case "a":
                                var v43 uint64 = 1
                                _ = v43
                            case "foo bar":
                                v1 = true
                            case "\n\t":
                                fmt.Println(v3, v7, fi42)
                            default:
                                switch "\n\t" {
                                case "a":
                                    if ((!((312 >= v3) && v1)) || true) {
                                        v5[42] = false
                                        switch fi42 {
                                        case -1:
                                            v3 -= (v3 >> ((769) & 63))
                                        case 127:
                                            var v44 float32 = 0.1
                                            _ = v44
                                        }
                                    } else {
                                        if v1 {
                                            v45 := GenS1{F0: (("\n\t" + "hello") + ("a" + ("" + (("" + "hello") + "a")))), F1: 96, F2: 42}
                                            if v1 {
                                                switch 9223372036854775807 {
                                                case 1:
                                                    var v46 uint16 = 2
                                                    _ = v46
                                                case -1:
                                                    v5[2147483647] = (!v1)
                                                    _ = len(v5)
                                                    _ = len(v2)
                                                    if (v1 || v1) {
                                                        var v47 int32 = (v7 % ((v7) | 1))
                                                        _ = v47
                                                    }
                                                }
                                                v48 := GenS2{F0: v4, F1: 42, F2: ""}
                                                _ = v48
                                            }
                                            if false {
                                                v6 -= (v6 & 127)
                                            }
                                            _ = v45
                                        } else {
                                            v1 = ((v7 ^ (v7 - int32(v6))) >= v7)
                                            v49 := GenS3{F0: 42, F1: true}
                                            _ = v49
                                        }
                                        delete(v5, 886)
                                        fmt.Println(fi42)
                                    }
                                default:
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                }
                            }
                        }
                        _ = fi42
                    }
                    switch -2147483648 {
                    case -2147483648:
                        for fi50 := 0; fi50 < 1; fi50++ {
                            v6 = v6
                            v51 := GenS3{F0: 217, F1: (42 == -1)}
                            _ = fi50
                            _ = v51
                        }
                    case 486:
                        switch (fi41 / ((fi41) | 1)) {
                        case 127:
                            if v1 {
                                v7 += v7
                            }
                        case -398:
                            v52 := GenS3{F0: -583, F1: (3.14159 > float64(0.1))}
                            _ = v52
                        default:
                            if v1 {
                                v53 := v5[65535]
                                if false {
                                    for fi54 := 0; fi54 < 9; fi54++ {
                                        v55 := GenS1{F0: ("" + "hello"), F1: (v4 % (((v4 | v4)) | 1)), F2: 18446744073709551615}
                                        v56 := v5[2147483648]
                                        _ = fi54
                                        _ = v55
                                        _ = v56
                                    }
                                } else {
                                    if v53 {
                                        var v57 int16 = 42
                                        _ = v57
                                    }
                                }
                                var v58 uint16 = 32767
                                _ = v53
                                _ = v58
                            }
                        }
                    case 42:
                        v3 += 0
                    }
                    v6 -= v6
                } else {
                    if ((!v1) || v1) {
                        if v1 {
                            var v59 float32 = 0.0
                            _ = v59
                        }
                        v60 := GenS4{F0: (v6 * v6), F1: v7}
                        _ = v60
                    } else {
                        v4 += v4
                    }
                    v4 -= v4
                    v3 = v3
                }
                _ = len(v2)
                _ = fi41
            }
        }
        switch -9223372036854775808 {
        case -9223372036854775808:
            v4 = 100
        case 2147483647:
            v7 += 58
        case 32768:
            fmt.Println(v7)
        }
    }
    v7 = (v7 + -2147483648)
    v6 += v6
    if (("\n\t" + "hello") != "") {
        var v61 string = "a"
        v62 := GenS5{F0: 208, F1: v61}
        var v63 uint = v3
        _ = v61
        _ = v62
        _ = v63
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v64 string = ((("\n\t" + "hello") + "hello") + "")
        v6 -= v6
        _ = v64
    }
    var v65 uint8 = 1
    v66 := GenS6{F0: v6}
    if true {
        var v67 int8 = 127
        v4 -= v4
        _ = v67
    }
    v68 := GenS7{F0: v7, F1: 0}
    for fi69 := 0; fi69 < 13; fi69++ {
        for fi70 := 0; fi70 < 20; fi70++ {
            if true {
                switch (fi70 % ((-9223372036854775808) | 1)) {
                case -740:
                    var v71 bool = v1
                    var v72 int16 = -750
                    var v73 int64 = 1
                    fmt.Println(v73)
                    fmt.Println(v4, v71, v6)
                    _ = v71
                    _ = v72
                    _ = v73
                case 165:
                    delete(v5, 0)
                case -9223372036854775808:
                    var v74 uint64 = 42
                    _ = v74
                default:
                    v75 := GenS8{F0: (v3 % (((v3 * 4294967295)) | 1))}
                    _ = v75
                }
                v65 = (v65 * 1)
                var v76 int8 = (v4 + 42)
                v77 := GenS9{F0: (fi70 << ((42) & 63))}
                _ = v76
                _ = v77
            }
            v78 := GenS8{F0: (v3 % ((v3) | 1))}
            var v79 uint16 = 65535
            _ = fi70
            _ = v78
            _ = v79
        }
        _ = fi69
    }
    v1 = ((v1 || ("\n\t" != (("foo bar" + "foo bar") + "hello"))) && true)
    fmt.Println(v65, v1, v7)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v65:", v65)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v65
    _ = v66
    _ = v68
    _ = v7
}
