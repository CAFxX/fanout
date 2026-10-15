package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int8; F1 uint16 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint; F1 bool }
type GenS7 struct { F0 bool; F1 uint }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 bool; F1 int; F2 int }
type GenS10 struct { F0 uint16; F1 int }
type GenS11 struct { F0 int8; F1 string }
type GenS12 struct { F0 int8 }

func main() {
    v1 := make(map[string]int)
    var v2 uint16 = 32767
    v3 := make(map[int]uint32)
    v4 := []string{"", (("\n\t" + (("" + "\n\t") + ("\n\t" + "hello"))) + "a"), "", "\n\t"}
    var v5 uint8 = 0
    v6 := GenS1{F0: 666, F1: -128}
    for fi7 := 0; fi7 < 2; fi7++ {
        for fi8 := 0; fi8 < 7; fi8++ {
            v5 = (v5 % ((v5) | 1))
            v9 := GenS1{F0: 2, F1: 127}
            v4 = append(v4, "hello")
            v2 = 65535
            v5 = (v5 / ((v5) | 1))
            _ = fi8
            _ = v9
        }
        v10 := GenS2{F0: v2}
        if true {
            switch (("foo bar" + "a") + "") {
            case "hello":
                var v11 int = 192
                _ = v11
            case "foo bar":
                switch "a" {
                case "\n\t":
                    var v12 int64 = -1
                    _ = v12
                case "hello":
                    v5 += (^(v5))
                case "":
                    v4 = append(v4, "foo bar")
                }
            default:
                fmt.Println(v2)
            }
            fmt.Println(fi7, v2, fi7)
            v13 := GenS2{F0: v2}
            _ = v13
        } else {
            if false {
                for fi14 := 0; fi14 < 14; fi14++ {
                    v15 := GenS3{F0: 75, F1: v2}
                    switch fi7 {
                    case -9223372036854775808:
                        var v16 string = ""
                        _ = v16
                    case -444:
                        var v17 string = ""
                        _ = v17
                    case -2147483648:
                        v18 := GenS1{F0: 52, F1: 127}
                        _ = v18
                    }
                    v19 := GenS4{F0: 2147483647}
                    _ = fi14
                    _ = v15
                    _ = v19
                }
                for fi20 := 0; fi20 < 19; fi20++ {
                    if ("" >= ("hello" + "hello")) {
                        var v21 uint32 = 4
                        if (!(0 != 255)) {
                            if (uint8(v21) > v5) {
                                v4 = append(v4, "a")
                                v21 = (v21 / (((v21 & (^(v21)))) | 1))
                                if (true || false) {
                                    v22 := GenS5{F0: 58}
                                    v2 -= v2
                                    switch "" {
                                    case "a":
                                        fmt.Println(v2, fi7)
                                    case "hello":
                                        if false {
                                            v23 := GenS1{F0: 0, F1: 100}
                                            delete(v1, "foo bar")
                                            if false {
                                                v5 = 188
                                                if (-533 != 42) {
                                                    var v24 uint16 = (v2 % ((65535) | 1))
                                                    if true {
                                                        v25 := GenS4{F0: 1}
                                                        _ = v25
                                                    } else {
                                                        var v26 int16 = -1
                                                        delete(v1, "hello")
                                                        fmt.Println(v5, v26, fi20)
                                                        _ = v26
                                                    }
                                                    v27 := GenS6{F0: 65535, F1: true}
                                                    _ = v24
                                                    _ = v27
                                                } else {
                                                    v28 := GenS6{F0: 4294967295, F1: false}
                                                    _ = v28
                                                }
                                                v21 = v21
                                                var v29 bool = (((("hello" >= "hello") || true) && ((v5 <= 42) && false)) && false)
                                                _ = v29
                                            } else {
                                                v30 := v1["a"]
                                                var v31 bool = ((false || (1000 <= -837)) && true)
                                                _ = v30
                                                _ = v31
                                            }
                                            v21 = v21
                                            switch fi20 {
                                            case 0:
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                            case 42:
                                                fmt.Println(fi20)
                                            default:
                                                v3[fi7] = v21
                                            }
                                            _ = v23
                                        } else {
                                            delete(v1, ("" + ("" + (("foo bar" + "\n\t") + "\n\t"))))
                                        }
                                    default:
                                        switch (fi7 | fi7) {
                                        case 32768:
                                            switch 32767 {
                                            case -2147483648:
                                                v32 := GenS7{F0: true, F1: 4294967295}
                                                _ = v32
                                            }
                                        case 503:
                                            switch ("hello" + (("" + "") + "")) {
                                            case "\n\t":
                                                if true {
                                                    var v33 uint8 = (^(v5))
                                                    switch fi7 {
                                                    case 41:
                                                        v2 -= v2
                                                    }
                                                    _ = v33
                                                } else {
                                                    v21 -= v21
                                                    if false {
                                                        fmt.Println(v21)
                                                    }
                                                }
                                            }
                                        }
                                    }
                                    _ = v22
                                }
                            } else {
                                fmt.Println(v2, fi20)
                                v21 -= (^(v21))
                                var v34 uint = 605
                                _ = v34
                            }
                            v5 = v5
                        }
                        _ = v21
                    }
                    _ = fi20
                }
            }
        }
        _ = fi7
        _ = v10
    }
    var v35 uint8 = (v5 / (((v5 >> ((0) & 63))) | 1))
    if false {
        if true {
            var v36 uint = 65535
            for fi37 := 0; fi37 < 19; fi37++ {
                var v38 int16 = 32767
                var v39 uint = 65535
                _ = fi37
                _ = v38
                _ = v39
            }
            for fi40 := 0; fi40 < 20; fi40++ {
                v41 := GenS4{F0: -2147483648}
                _ = fi40
                _ = v41
            }
            _ = v36
        } else {
            var v42 float64 = 0.1
            fmt.Println(v5, v35)
            _ = v42
        }
        for fi43 := 0; fi43 < 4; fi43++ {
            var v44 float32 = 3.14159
            _ = fi43
            _ = v44
        }
        for fi45 := 0; fi45 < 16; fi45++ {
            v46 := GenS8{F0: 1}
            _ = fi45
            _ = v46
        }
    } else {
        var v47 uint = 18446744073709551615
        v47 = v47
        var v48 uint32 = 2147483648
        _ = v47
        _ = v48
    }
    if (823 == 0) {
        for fi49 := 0; fi49 < 5; fi49++ {
            v50 := GenS9{F0: ("foo bar" < "\n\t"), F1: -2147483648, F2: (-(fi49))}
            v35 = (v35 % ((v5) | 1))
            v51 := GenS7{F0: (!false), F1: 65535}
            _ = fi49
            _ = v50
            _ = v51
        }
    }
    switch "" {
    case "a":
        var v52 float32 = 1.5
        _ = v52
    default:
        switch 127 {
        case 2147483647:
            var v53 uint16 = uint16(v35)
            _ = v53
        case 984:
            var v54 uint8 = 72
            _ = v54
        case 4294967295:
            for fi55 := 0; fi55 < 8; fi55++ {
                v56 := GenS10{F0: 255, F1: -9223372036854775808}
                var v57 int16 = 1
                var v58 int8 = 100
                _ = fi55
                _ = v56
                _ = v57
                _ = v58
            }
        default:
            switch "" {
            case "\n\t":
                var v59 int16 = -32768
                _ = v59
            default:
                fmt.Println(v5, v5, v35)
            }
        }
    }
    switch -9223372036854775808 {
    case 1:
        v5 = (v35 * (v5 >> ((65535) & 63)))
    case -9223372036854775808:
        v2 = (^(v2))
    case 32767:
        fmt.Println(v35, v35, v2)
    default:
        if (0.1 > ((((3.14159 / 1.5) / 1e10) * 0.1) * 1e10)) {
            for fi60 := 0; fi60 < 16; fi60++ {
                v35 += (v35 % (((^(v5))) | 1))
                for fi61 := 0; fi61 < 9; fi61++ {
                    if (fi60 != fi60) {
                        v2 += (v2 >> ((18446744073709551615) & 63))
                        switch fi60 {
                        case -9223372036854775808:
                            v62 := GenS11{F0: 0, F1: "\n\t"}
                            _ = v62
                        case 9223372036854775807:
                            var v63 int64 = -9223372036854775808
                            _ = v63
                        case -979:
                            var v64 bool = (false || (9223372036854775807 >= 0))
                            _ = v64
                        }
                        if false {
                            fmt.Println(fi60)
                            if len(v4) > 0 {
                                v4[0] = ("foo bar" + "")
                            }
                        }
                        v65 := GenS6{F0: 65535, F1: false}
                        if (int32(v2) > 665) {
                            switch ("\n\t" + "a") {
                            case "foo bar":
                                var v66 int64 = -405
                                _ = v66
                            }
                        } else {
                            if true {
                                switch "" {
                                case "hello":
                                    fmt.Println(v5)
                                case "foo bar":
                                    switch (fi61 - -235) {
                                    case 42:
                                        delete(v3, fi60)
                                    case 9223372036854775807:
                                        if len(v4) > 0 {
                                            v4[0] = "\n\t"
                                        }
                                    }
                                case "\n\t":
                                    v67 := GenS3{F0: -102, F1: (v2 % ((v2) | 1))}
                                    _ = v67
                                }
                                v68 := GenS7{F0: (1000 <= -32768), F1: 878}
                                v69 := GenS12{F0: 42}
                                _ = v68
                                _ = v69
                            }
                            v35 -= (v35 ^ v5)
                        }
                        _ = v65
                    }
                    var v70 uint16 = v2
                    if (false && true) {
                        fmt.Println(fi61)
                        var v71 bool = true
                        if (-91 < 0) {
                            fmt.Println(v70)
                            fmt.Println(v2)
                        } else {
                            v35 += v5
                        }
                        _ = v71
                    } else {
                        v2 += v2
                        switch fi60 {
                        case 128:
                            v1[("a" + "hello")] = (fi60 / (((fi61 & int(v35))) | 1))
                        }
                        v4 = append(v4, "hello")
                    }
                    _ = fi61
                    _ = v70
                }
                _ = fi60
            }
        } else {
            var v72 int = 128
            if (2 >= 2) {
                if true {
                    v5 = v35
                    for fi73 := 0; fi73 < 14; fi73++ {
                        for fi74 := 0; fi74 < 11; fi74++ {
                            if (false && (!true)) {
                                var v75 uint8 = (v35 - (v35 - v5))
                                if (100 >= 17) {
                                    v72 = v72
                                    var v76 float64 = -2.25
                                    v72 = 127
                                    _ = v76
                                }
                                if (-2.25 >= -2.25) {
                                    delete(v3, fi74)
                                    if (true && (true && false)) {
                                        v35 += 42
                                        delete(v1, "hello")
                                    } else {
                                        v5 += 0
                                        v77 := GenS2{F0: 207}
                                        _ = v77
                                    }
                                    fmt.Println(fi73, fi73)
                                }
                                _ = v75
                            } else {
                                fmt.Println(v72, fi73, fi74)
                            }
                            _ = fi74
                        }
                        _ = fi73
                    }
                }
                var v78 float64 = -2.25
                _ = v78
            }
            fmt.Println(v2, v72)
            _ = v72
        }
    }
    v5 = 128
    v5 -= (v5 | v35)
    fmt.Println(v2, v2)
    if (1 >= 254) {
        v5 = v35
        _ = len(v3)
        for fi79 := 0; fi79 < 19; fi79++ {
            v4 = append(v4, "a")
            _ = fi79
        }
    } else {
        for fi80 := 0; fi80 < 14; fi80++ {
            v81 := v1[("foo bar" + ("\n\t" + "hello"))]
            var v82 bool = true
            _ = fi80
            _ = v81
            _ = v82
        }
        v35 = 13
        v5 += v35
    }
    delete(v3, 558)
    v35 = v5
    if ((!(915 == 18446744073709551615)) && (572 <= -1)) {
        fmt.Println(v35, v2, v5)
    } else {
        if (111 > (v5 % (((v5 / ((52) | 1))) | 1))) {
            if (!(!false)) {
                v3[800] = 1
            } else {
                v83 := GenS6{F0: 0, F1: false}
                _ = v83
            }
        }
        v5 = 215
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v84 bool = (255 != 2)
    var v85 int64 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v84:", v84)
    fmt.Println("v85:", v85)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v84
    _ = v85
}
