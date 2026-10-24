package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 int16; F1 uint8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint; F1 uint32 }
type GenS6 struct { F0 int64; F1 string; F2 int8 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 uint32; F1 int32; F2 int8 }
type GenS10 struct { F0 uint8 }

func main() {
    v1 := []int{-9223372036854775808}
    v2 := make(map[string]bool)
    v3 := make(map[string]string)
    var v4 uint = 65535
    if true {
        for fi5 := 0; fi5 < 20; fi5++ {
            v4 = 42
            if len(v1) > 0 {
                v1[0] = -9223372036854775808
            }
            var v6 string = ((("\n\t" + (("hello" + "a") + ("a" + (("a" + "hello") + ("a" + "\n\t"))))) + ("\n\t" + "foo bar")) + "a")
            _ = fi5
            _ = v6
        }
    } else {
        fmt.Println(v4)
    }
    v4 -= v4
    fmt.Println(v4)
    v4 = v4
    for fi7 := 0; fi7 < 12; fi7++ {
        if false {
            if ((!(!false)) && false) {
                switch 994 {
                case -9223372036854775808:
                    v8 := GenS1{F0: v4, F1: 128}
                    _ = v8
                default:
                    var v9 uint32 = 774
                    _ = v9
                }
            }
            switch (fi7 & (^(fi7))) {
            case 42:
                fmt.Println(v4, v4)
            default:
                v4 -= 65535
            }
        } else {
            v10 := GenS1{F0: 65535, F1: 1}
            _ = v10
        }
        _ = fi7
    }
    if false {
        for fi11 := 0; fi11 < 6; fi11++ {
            _ = len(v3)
            if len(v1) > 0 {
                v1[0] = fi11
            }
            v4 = (^(v4))
            _ = fi11
        }
        fmt.Println(v4)
        v4 = (v4 ^ 4294967295)
    }
    if false {
        if (true && true) {
            v4 = v4
            v4 -= v4
        } else {
            if (0 < 32767) {
                fmt.Println(v4)
            }
            switch "hello" {
            case "":
                if false {
                    if (0 < -128) {
                        switch "" {
                        case "hello":
                            v12 := GenS2{F0: 4294967295, F1: v4}
                            _ = v12
                        }
                    } else {
                        fmt.Println(v4)
                        fmt.Println(v4)
                        if false {
                            v4 += (v4 + v4)
                        } else {
                            v4 = v4
                        }
                    }
                } else {
                    for fi13 := 0; fi13 < 15; fi13++ {
                        v4 = (^(v4))
                        if len(v1) > 0 {
                            v1[0] = 0
                        }
                        _ = fi13
                    }
                }
            case "hello":
                fmt.Println(v4)
            default:
                switch "hello" {
                case "":
                    v4 += 2
                case "foo bar":
                    v14 := GenS3{F0: -332, F1: 128}
                    _ = v14
                default:
                    fmt.Println(v4)
                }
            }
            var v15 float32 = 1.5
            _ = v15
        }
    } else {
        v16 := GenS2{F0: 2, F1: (^(v4))}
        fmt.Println(v4)
        _ = v16
    }
    var v17 uint16 = 963
    fmt.Println(v4, v17)
    v18 := GenS1{F0: (v4 - v4), F1: 1}
    v19 := GenS4{F0: true}
    if false {
        if false {
            if false {
                var v20 int32 = -1
                if ((v4 * 18446744073709551615) >= 1) {
                    if len(v1) > 0 {
                        v1[0] = -792
                    }
                    var v21 float64 = 0.0
                    if true {
                        v22 := GenS4{F0: (!false)}
                        _ = v22
                    }
                    _ = v21
                } else {
                    if true {
                        for fi23 := 0; fi23 < 7; fi23++ {
                            v24 := v3[""]
                            var v25 string = v24
                            _ = fi23
                            _ = v24
                            _ = v25
                        }
                        if (!(32767 > 814)) {
                            if (v20 == (v20 | v20)) {
                                v17 = v17
                            }
                            var v26 int8 = 127
                            _ = v26
                        }
                        if false {
                            var v27 string = "\n\t"
                            _ = v27
                        } else {
                            if true {
                                var v28 float32 = 1.5
                                if false {
                                    for fi29 := 0; fi29 < 18; fi29++ {
                                        v20 = (^(v20))
                                        if (0 > 42) {
                                            fmt.Println(v20)
                                            switch "a" {
                                            case "foo bar":
                                                v4 += (v4 % ((v4) | 1))
                                            case "\n\t":
                                                var v30 uint8 = 0
                                                _ = v30
                                            }
                                        } else {
                                            var v31 uint16 = v17
                                            _ = v31
                                        }
                                        var v32 string = ("foo bar" + "\n\t")
                                        _ = fi29
                                        _ = v32
                                    }
                                } else {
                                    fmt.Println(v4)
                                    var v33 int32 = v20
                                    switch (("foo bar" + "a") + (("hello" + "hello") + "hello")) {
                                    case "\n\t":
                                        v4 = 42
                                    case "":
                                        var v34 uint = (v4 / ((974) | 1))
                                        _ = v34
                                    case "a":
                                        v35 := GenS5{F0: (v4 - v4), F1: 295}
                                        _ = v35
                                    default:
                                        for fi36 := 0; fi36 < 15; fi36++ {
                                            for fi37 := 0; fi37 < 1; fi37++ {
                                                var v38 float64 = 0.0
                                                var v39 int32 = v33
                                                _ = fi37
                                                _ = v38
                                                _ = v39
                                            }
                                            for fi40 := 0; fi40 < 11; fi40++ {
                                                v4 = (v4 ^ 556)
                                                switch "" {
                                                case "":
                                                    var v41 int8 = 0
                                                    _ = v41
                                                case "hello":
                                                    switch fi36 {
                                                    case 2147483647:
                                                        v4 += (^(v4))
                                                        if (!(80 > 1000)) {
                                                            var v42 float32 = v28
                                                            v43 := GenS6{F0: int64(fi36), F1: "hello", F2: -5}
                                                            v44 := GenS1{F0: uint(v17), F1: 0}
                                                            _ = v42
                                                            _ = v43
                                                            _ = v44
                                                        } else {
                                                            v28 = v28
                                                            _ = len(v1)
                                                        }
                                                        if (v28 < v28) {
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                            var v45 uint8 = 197
                                                            _ = v45
                                                        } else {
                                                            v46 := v2[(((("\n\t" + (((("a" + ("hello" + "\n\t")) + "") + (("" + "a") + "a")) + (("" + "") + "a"))) + "hello") + "hello") + "a")]
                                                            var v47 bool = v46
                                                            v47 = ((((fi36 % ((fi36) | 1)) > -2147483648) || (float32(-2.25) > (v28 / v28))) || v46)
                                                            _ = v46
                                                            _ = v47
                                                        }
                                                        var v48 float32 = ((float32(0.0) / (v28 + 0.1)) + 0.1)
                                                        _ = v48
                                                    case 127:
                                                        v49 := GenS4{F0: (2 == 0)}
                                                        _ = v49
                                                    }
                                                }
                                                var v50 int = (fi40 >> ((uint(v17)) & 63))
                                                _ = fi40
                                                _ = v50
                                            }
                                            switch -9223372036854775808 {
                                            case -9223372036854775808:
                                                var v51 uint8 = 127
                                                _ = v51
                                            case -1:
                                                v52 := GenS6{F0: int64(v20), F1: ("a" + ((("\n\t" + "hello") + "a") + "a")), F2: 42}
                                                _ = v52
                                            case 0:
                                                if (true || false) {
                                                    var v53 float64 = (1e10 * 3.14159)
                                                    _ = v53
                                                }
                                            }
                                            _ = fi36
                                        }
                                    }
                                    _ = v33
                                }
                                delete(v3, "hello")
                                _ = v28
                            } else {
                                for fi54 := 0; fi54 < 13; fi54++ {
                                    fmt.Println(v4, fi54, fi54)
                                    switch "hello" {
                                    case "\n\t":
                                        v4 = v4
                                    case "foo bar":
                                        if (-1 < 42) {
                                            v20 = 0
                                            for fi55 := 0; fi55 < 2; fi55++ {
                                                var v56 uint8 = 1
                                                _ = fi55
                                                _ = v56
                                            }
                                            v4 = 1
                                        } else {
                                            fmt.Println(v4, fi54, v20)
                                            switch ("\n\t" + "a") {
                                            case "\n\t":
                                                var v57 uint16 = v17
                                                _ = v57
                                            case "":
                                                v20 = (v20 % ((v20) | 1))
                                            default:
                                                var v58 int64 = -1
                                                _ = v58
                                            }
                                        }
                                    default:
                                        fmt.Println(v4, v17, v20)
                                    }
                                    _ = fi54
                                }
                                for fi59 := 0; fi59 < 10; fi59++ {
                                    v60 := v3["foo bar"]
                                    _ = fi59
                                    _ = v60
                                }
                                fmt.Println(v4, v20)
                            }
                            v17 += v17
                            if true {
                                v20 += v20
                                for fi61 := 0; fi61 < 6; fi61++ {
                                    for fi62 := 0; fi62 < 7; fi62++ {
                                        if true {
                                            if false {
                                                switch fi61 {
                                                case 0:
                                                    if false {
                                                        if false {
                                                            fmt.Println(v17, v4, v4)
                                                        }
                                                    } else {
                                                        v63 := GenS7{F0: (v20 << (((v4 & v4)) & 63))}
                                                        v20 = 42
                                                        _ = v63
                                                    }
                                                    v2["hello"] = true
                                                    if ((0.0 - 0.1) <= 0.0) {
                                                        var v64 bool = true
                                                        v65 := GenS1{F0: v4, F1: 210}
                                                        switch (fi62 + fi61) {
                                                        case 0:
                                                            var v66 uint64 = 4294967295
                                                            _ = v66
                                                        }
                                                        _ = v64
                                                        _ = v65
                                                    } else {
                                                        v17 = uint16(v4)
                                                        if len(v1) > 0 {
                                                            v1[0] = fi61
                                                        }
                                                    }
                                                case 1:
                                                    var v67 int = (fi62 & fi61)
                                                    _ = v67
                                                }
                                                v17 -= 239
                                                switch ("foo bar" + "hello") {
                                                case "hello":
                                                    if true {
                                                        v4 = v4
                                                        v20 += (v20 << ((433) & 63))
                                                        v4 = v4
                                                        fmt.Println(fi61, v4, fi62)
                                                        v20 += 0
                                                    }
                                                case "a":
                                                    v17 = (^(v17))
                                                case "foo bar":
                                                    var v68 int64 = -9223372036854775808
                                                    _ = v68
                                                }
                                            } else {
                                                if (855 >= 1) {
                                                    _ = len(v2)
                                                    var v69 uint = 42
                                                    _ = v69
                                                } else {
                                                    var v70 uint64 = 314
                                                    _ = v70
                                                }
                                                v20 = (v20 - (v20 % ((v20) | 1)))
                                            }
                                            if false {
                                                switch fi62 {
                                                case 42:
                                                    var v71 float64 = 0.0
                                                    _ = v71
                                                default:
                                                    var v72 int16 = 449
                                                    _ = v72
                                                }
                                                v20 = v20
                                                var v73 int16 = 42
                                                _ = v73
                                            } else {
                                                if false {
                                                    _ = len(v1)
                                                    fmt.Println(v4, v4, v17)
                                                    v1 = append(v1, fi62)
                                                } else {
                                                    var v74 int32 = (v20 + 42)
                                                    if len(v1) > 0 {
                                                        v1[0] = fi61
                                                    }
                                                    _ = v74
                                                }
                                                switch "\n\t" {
                                                case "hello":
                                                    switch "\n\t" {
                                                    case "a":
                                                        var v75 float32 = -2.25
                                                        _ = v75
                                                    case "hello":
                                                        switch (^(fi61)) {
                                                        case 0:
                                                            fmt.Println(v20, fi62, v20)
                                                        case -587:
                                                            if true {
                                                                v4 = 133
                                                            } else {
                                                                if false {
                                                                    switch "\n\t" {
                                                                    case "foo bar":
                                                                        v20 = v20
                                                                    case "\n\t":
                                                                        if (42 < 1) {
                                                                            if false {
                                                                                var v76 float32 = 1.5
                                                                                v4 += v4
                                                                                _ = v76
                                                                            }
                                                                        } else {
                                                                            v77 := GenS4{F0: ((0 != -1) || true)}
                                                                            v78 := GenS8{F0: 65535}
                                                                            _ = v77
                                                                            _ = v78
                                                                        }
                                                                    case "":
                                                                        v20 += v20
                                                                    }
                                                                    var v79 int64 = 9223372036854775807
                                                                    _ = v79
                                                                }
                                                                v80 := GenS8{F0: 0}
                                                                switch -150 {
                                                                case -2147483648:
                                                                    v20 = v20
                                                                case -41:
                                                                    _ = len(v1)
                                                                case 127:
                                                                    v20 = (v20 & 1)
                                                                default:
                                                                    var v81 float32 = 3.14159
                                                                    _ = v81
                                                                }
                                                                _ = v80
                                                            }
                                                        }
                                                    case "foo bar":
                                                        v1 = append(v1, fi62)
                                                    }
                                                case "a":
                                                    v17 = 42
                                                case "foo bar":
                                                    v82 := GenS9{F0: 2147483648, F1: -899, F2: 0}
                                                    _ = v82
                                                }
                                                v17 = (v17 | v17)
                                            }
                                        }
                                        _ = fi62
                                    }
                                    fmt.Println(v4, v4)
                                    _ = fi61
                                }
                                v83 := GenS1{F0: 1, F1: 255}
                                _ = v83
                            } else {
                                var v84 int16 = 32767
                                if false {
                                    var v85 int8 = 42
                                    fmt.Println(v17, v85)
                                    _ = v85
                                } else {
                                    v86 := GenS6{F0: -720, F1: (("foo bar" + "hello") + ""), F2: -128}
                                    _ = v86
                                }
                                _ = v84
                            }
                        }
                    }
                }
                v87 := GenS1{F0: v4, F1: 143}
                _ = v20
                _ = v87
            }
            v88 := GenS10{F0: 255}
            _ = v88
        } else {
            v17 += (^(v17))
            if (!true) {
                v2["hello"] = (1 >= -1)
                v89 := GenS3{F0: 1000, F1: 174}
                _ = v89
            }
        }
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        for fi90 := 0; fi90 < 16; fi90++ {
            var v91 string = "hello"
            _ = fi90
            _ = v91
        }
    }
    switch -180 {
    case -719:
        v4 += v4
    case -425:
        var v92 float32 = 0.0
        _ = v92
    }
    var v93 uint32 = 0
    for fi94 := 0; fi94 < 18; fi94++ {
        var v95 uint16 = (v17 % ((487) | 1))
        var v96 bool = (!(v93 <= 591))
        _ = fi94
        _ = v95
        _ = v96
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v93
}
