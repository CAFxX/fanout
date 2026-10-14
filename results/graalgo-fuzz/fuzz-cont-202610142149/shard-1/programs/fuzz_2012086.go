package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 uint64; F1 uint64; F2 uint8 }
type GenS3 struct { F0 int64; F1 int32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 bool; F1 string }
type GenS6 struct { F0 uint8; F1 uint8 }
type GenS7 struct { F0 bool; F1 uint16 }
type GenS8 struct { F0 int64; F1 int64; F2 int64 }
type GenS9 struct { F0 uint8; F1 uint16 }

func main() {
    v1 := []uint32{}
    var v2 float64 = 0.1
    var v3 int32 = 42
    var v4 bool = true
    v3 += v3
    v3 += 0
    for fi5 := 0; fi5 < 4; fi5++ {
        var v6 float64 = (v2 / v2)
        _ = fi5
        _ = v6
    }
    var v7 int16 = -32768
    var v8 int64 = -1
    switch ((((("hello" + (("foo bar" + "hello") + "")) + ("a" + "")) + "") + ("\n\t" + "hello")) + "a") {
    case "\n\t":
        for fi9 := 0; fi9 < 9; fi9++ {
            var v10 int = fi9
            v3 = 0
            _ = fi9
            _ = v10
        }
    case "hello":
        if true {
            v11 := GenS1{F0: 42, F1: (116 > 57)}
            _ = v11
        } else {
            v12 := GenS1{F0: 42, F1: v4}
            var v13 float64 = (v2 + 0.0)
            var v14 float64 = 1e10
            _ = v12
            _ = v13
            _ = v14
        }
    default:
        if (!true) {
            v15 := GenS2{F0: 560, F1: 42, F2: 0}
            v16 := GenS3{F0: (v8 >> ((2) & 63)), F1: (-(v3))}
            _ = v15
            _ = v16
        } else {
            var v17 uint16 = 936
            var v18 float64 = v2
            if (("" + "") < "") {
                if (v4 && (740 <= 4294967295)) {
                    var v19 uint16 = v17
                    v7 -= v7
                    _ = v19
                }
                for fi20 := 0; fi20 < 9; fi20++ {
                    var v21 string = ((((("foo bar" + ("\n\t" + "foo bar")) + ("foo bar" + ("hello" + "hello"))) + "\n\t") + "\n\t") + "foo bar")
                    for fi22 := 0; fi22 < 20; fi22++ {
                        v2 = float64(-2.25)
                        var v23 uint = 435
                        _ = fi22
                        _ = v23
                    }
                    var v24 uint64 = 1
                    _ = fi20
                    _ = v21
                    _ = v24
                }
                switch ("\n\t" + "hello") {
                case "hello":
                    var v25 uint32 = 42
                    _ = v25
                }
            } else {
                v1 = append(v1, 185)
                v26 := GenS4{F0: 2}
                _ = v26
            }
            _ = v17
            _ = v18
        }
    }
    if (false || (2 >= 18446744073709551615)) {
        if ((-(v8)) == (v8 * v8)) {
            if false {
                var v27 int16 = int16(v8)
                _ = v27
            }
            var v28 uint16 = 32767
            v8 += v8
            _ = v28
        } else {
            v2 += (3.14159 + float64((0.1 * ((0.1 * 1.5) - (1e10 - -2.25)))))
            switch "hello" {
            case "":
                if len(v1) > 0 {
                    v1[0] = 1
                }
            case "a":
                var v29 uint32 = 2147483648
                _ = v29
            case "hello":
                if (2147483647 < 32767) {
                    switch -661 {
                    case -2147483648:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    case 475:
                        if (!(!v4)) {
                            var v30 int32 = 0
                            switch "hello" {
                            case "":
                                for fi31 := 0; fi31 < 7; fi31++ {
                                    v4 = (!v4)
                                    v32 := GenS1{F0: 167, F1: v4}
                                    if v4 {
                                        v3 = -1
                                        for fi33 := 0; fi33 < 6; fi33++ {
                                            v34 := GenS4{F0: 0}
                                            var v35 float64 = (v2 * v2)
                                            var v36 int32 = (v30 | 2147483647)
                                            _ = fi33
                                            _ = v34
                                            _ = v35
                                            _ = v36
                                        }
                                        var v37 int64 = v8
                                        _ = v37
                                    }
                                    _ = fi31
                                    _ = v32
                                }
                            default:
                                if (128 > 1) {
                                    v4 = v4
                                    if len(v1) > 0 {
                                        v1[0] = 42
                                    }
                                    if (v8 <= (v8 & (v8 | int64(v7)))) {
                                        _ = len(v1)
                                    }
                                } else {
                                    v38 := GenS3{F0: v8, F1: v3}
                                    v39 := GenS5{F0: v4, F1: "a"}
                                    v40 := GenS4{F0: 4294967295}
                                    _ = v38
                                    _ = v39
                                    _ = v40
                                }
                            }
                            _ = v30
                        } else {
                            var v41 uint16 = 624
                            _ = v41
                        }
                    case -1:
                        if v4 {
                            var v42 uint16 = 65535
                            fmt.Println(v8, v3)
                            _ = v42
                        } else {
                            for fi43 := 0; fi43 < 8; fi43++ {
                                if (1.5 == 3.14159) {
                                    var v44 uint16 = 1
                                    var v45 string = ("foo bar" + "a")
                                    _ = v44
                                    _ = v45
                                }
                                v2 += float64(3.14159)
                                var v46 int16 = v7
                                _ = fi43
                                _ = v46
                            }
                            fmt.Println(v8, v7, v3)
                            if false {
                                if true {
                                    if ((!true) && v4) {
                                        switch ("hello" + ("hello" + ("foo bar" + "\n\t"))) {
                                        case "hello":
                                            var v47 int = 0
                                            _ = v47
                                        case "":
                                            v8 = v8
                                        }
                                        v48 := GenS2{F0: 419, F1: 508, F2: 255}
                                        _ = v48
                                    }
                                    for fi49 := 0; fi49 < 20; fi49++ {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        _ = len(v1)
                                        v1 = append(v1, 73)
                                        _ = fi49
                                    }
                                } else {
                                    if (65535 == 0) {
                                        for fi50 := 0; fi50 < 11; fi50++ {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            v8 = v8
                                            if true {
                                                v51 := GenS6{F0: 65, F1: 255}
                                                v2 -= 1e10
                                                _ = v51
                                            } else {
                                                v52 := GenS2{F0: 18446744073709551615, F1: 53, F2: 255}
                                                v4 = (!v4)
                                                if len(v1) > 0 {
                                                    v1[0] = 42
                                                }
                                                _ = v52
                                            }
                                            _ = fi50
                                        }
                                    } else {
                                        switch "foo bar" {
                                        case "a":
                                            fmt.Println(v2, v7)
                                        default:
                                            fmt.Println(v7, v2)
                                        }
                                    }
                                    var v53 float32 = ((3.14159 / float32(0.1)) - float32(float64(((1.5 * (0.1 * ((0.1 + 0.1) - (0.0 * 1e10)))) * 3.14159))))
                                    _ = v53
                                }
                                for fi54 := 0; fi54 < 16; fi54++ {
                                    v55 := GenS7{F0: (((2 == 128) || (v7 != 998)) && v4), F1: 65535}
                                    v4 = ("a" <= "")
                                    if (v4 || (0 > 42)) {
                                        fmt.Println(v7, v7)
                                    }
                                    _ = fi54
                                    _ = v55
                                }
                            } else {
                                v56 := GenS4{F0: 2147483647}
                                v1 = append(v1, 42)
                                var v57 uint8 = 2
                                _ = v56
                                _ = v57
                            }
                        }
                    }
                } else {
                    var v58 string = "a"
                    _ = v58
                }
            }
        }
        if (!v4) {
            v8 = (v8 | v8)
            for fi59 := 0; fi59 < 3; fi59++ {
                if v4 {
                    var v60 float64 = v2
                    _ = v60
                } else {
                    var v61 int64 = (v8 / (((v8 << ((321) & 63))) | 1))
                    _ = v61
                }
                if ((!true) && v4) {
                    fmt.Println(fi59)
                    switch "foo bar" {
                    case "\n\t":
                        switch (fi59 | (^(fi59))) {
                        case 32767:
                            v7 += v7
                        }
                    case "":
                        switch "foo bar" {
                        case "foo bar":
                            var v62 uint32 = 0
                            _ = v62
                        case "a":
                            if len(v1) > 0 {
                                v1[0] = 627
                            }
                        case "":
                            fmt.Println(v2)
                        }
                    case "a":
                        v63 := GenS8{F0: (v8 << ((65535) & 63)), F1: (v8 ^ (v8 % ((9223372036854775807) | 1))), F2: v8}
                        _ = v63
                    }
                    if (v4 || ((!(1 != 928)) && v4)) {
                        if len(v1) > 0 {
                            v1[0] = 2147483647
                        }
                        var v64 uint32 = 42
                        if true {
                            v4 = v4
                            v1 = append(v1, 220)
                        }
                        _ = v64
                    } else {
                        var v65 uint8 = 255
                        switch int(v3) {
                        case 32768:
                            if ("a" != "\n\t") {
                                var v66 uint32 = 247
                                v4 = true
                                _ = v66
                            } else {
                                var v67 uint = 784
                                if len(v1) > 0 {
                                    v1[0] = 691
                                }
                                _ = v67
                            }
                        case 0:
                            fmt.Println(fi59, v2)
                        default:
                            var v68 uint8 = 127
                            _ = v68
                        }
                        var v69 bool = (127 != 127)
                        _ = v65
                        _ = v69
                    }
                } else {
                    if len(v1) > 0 {
                        v1[0] = 2147483648
                    }
                }
                for fi70 := 0; fi70 < 12; fi70++ {
                    v3 = (v3 << ((4294967295) & 63))
                    v71 := GenS9{F0: 0, F1: 297}
                    v2 = -2.25
                    _ = fi70
                    _ = v71
                }
                _ = fi59
            }
            var v72 uint = 971
            _ = v72
        }
    } else {
        for fi73 := 0; fi73 < 20; fi73++ {
            v7 = (v7 & (v7 - (v7 % (((v7 | v7)) | 1))))
            _ = fi73
        }
        v4 = (!false)
    }
    v2 = -2.25
    if true {
        var v74 int16 = 42
        if v4 {
            v8 += (v8 | v8)
            for fi75 := 0; fi75 < 9; fi75++ {
                var v76 string = "foo bar"
                var v77 string = "\n\t"
                _ = fi75
                _ = v76
                _ = v77
            }
        }
        _ = v74
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
