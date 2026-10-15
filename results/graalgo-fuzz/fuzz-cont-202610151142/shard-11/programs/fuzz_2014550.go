package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int32; F1 int64; F2 int }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int; F1 int; F2 int }
type GenS8 struct { F0 bool; F1 int8; F2 bool }
type GenS9 struct { F0 bool; F1 uint64 }
type GenS10 struct { F0 uint16; F1 uint16 }
type GenS11 struct { F0 int }

func main() {
    v1 := make(map[string]bool)
    var v2 uint64 = 4294967295
    var v3 uint = 42
    v4 := []uint32{2147483648}
    var v5 float64 = -2.25
    switch "\n\t" {
    case "\n\t":
        _ = len(v1)
    default:
        if true {
            if false {
                for fi6 := 0; fi6 < 2; fi6++ {
                    if true {
                        for fi7 := 0; fi7 < 13; fi7++ {
                            var v8 int16 = 32767
                            switch ("hello" + "a") {
                            case "hello":
                                fmt.Println(fi6, v2, fi6)
                            case "foo bar":
                                var v9 uint32 = 759
                                _ = v9
                            }
                            v8 = 216
                            _ = fi7
                            _ = v8
                        }
                        v5 += (v5 / v5)
                        v3 -= (v3 ^ (v3 / (((v3 >> ((1) & 63))) | 1)))
                    } else {
                        v10 := GenS1{F0: 1}
                        for fi11 := 0; fi11 < 17; fi11++ {
                            var v12 int8 = 1
                            if true {
                                var v13 int64 = -528
                                v14 := v1[""]
                                if ("hello" == ("" + (("a" + "hello") + "\n\t"))) {
                                    var v15 string = ""
                                    var v16 uint8 = 101
                                    v4 = append(v4, 65535)
                                    v2 = (v2 << (((v3 % (((v3 / (((^(v3))) | 1))) | 1))) & 63))
                                    _ = v15
                                    _ = v16
                                } else {
                                    v17 := GenS1{F0: (-(v12))}
                                    if false {
                                        switch (fi6 << ((18446744073709551615) & 63)) {
                                        case 2147483647:
                                            if (-431 <= 32767) {
                                                fmt.Println(v14, fi11, fi6)
                                            }
                                        default:
                                            v2 = (v2 | v2)
                                        }
                                        v18 := GenS2{F0: 1}
                                        if v14 {
                                            v2 += v2
                                            v19 := v1["\n\t"]
                                            if (v19 && (-777 > 1000)) {
                                                v14 = (-32768 < 1)
                                                v3 -= v3
                                                if false {
                                                    if v19 {
                                                        if v14 {
                                                            v20 := GenS3{F0: 255}
                                                            _ = v20
                                                        } else {
                                                            switch ("a" + "hello") {
                                                            case "foo bar":
                                                                if false {
                                                                    v13 -= 9223372036854775807
                                                                    v3 -= (^(v3))
                                                                    v13 = (v13 >> ((v3) & 63))
                                                                } else {
                                                                    v13 -= v13
                                                                }
                                                            case "a":
                                                                fmt.Println(v3, v5)
                                                            case "hello":
                                                                v3 -= v3
                                                            }
                                                            var v21 int64 = v13
                                                            var v22 string = "a"
                                                            _ = v21
                                                            _ = v22
                                                        }
                                                        fmt.Println(v19, fi11)
                                                    }
                                                } else {
                                                    v2 = (^(v2))
                                                }
                                                if true {
                                                    v4 = append(v4, 971)
                                                    var v23 int8 = 42
                                                    _ = v23
                                                } else {
                                                    v2 = 18446744073709551615
                                                }
                                                fmt.Println(v2, v12)
                                                delete(v1, "foo bar")
                                            } else {
                                                v2 += v2
                                                switch fi6 {
                                                case -445:
                                                    _ = len(v4)
                                                default:
                                                    if (!false) {
                                                        if (!((v12 & v12) == v12)) {
                                                            _ = len(v1)
                                                        }
                                                        if (!(true || v19)) {
                                                            v5 -= v5
                                                            if false {
                                                                v12 += (v12 / ((v12) | 1))
                                                                v24 := GenS2{F0: -32768}
                                                                v25 := GenS3{F0: 500}
                                                                v13 = v13
                                                                v14 = v14
                                                                _ = v24
                                                                _ = v25
                                                            }
                                                            var v26 int8 = (^(v12))
                                                            if (v3 >= 255) {
                                                                var v27 uint8 = 2
                                                                if v14 {
                                                                    var v28 uint32 = 4294967295
                                                                    var v29 int = (fi11 + fi11)
                                                                    v2 = 1
                                                                    _ = v28
                                                                    _ = v29
                                                                } else {
                                                                    if false {
                                                                        fmt.Println(v2, v3)
                                                                    }
                                                                }
                                                                var v30 int64 = v13
                                                                _ = v27
                                                                _ = v30
                                                            }
                                                            _ = v26
                                                        }
                                                        var v31 uint8 = 141
                                                        _ = v31
                                                    }
                                                }
                                                v32 := GenS4{F0: v3}
                                                var v33 int = (fi11 << ((uint(v2)) & 63))
                                                v34 := GenS1{F0: (v12 + 127)}
                                                _ = v32
                                                _ = v33
                                                _ = v34
                                            }
                                            v1[""] = (v3 > (v3 << ((42) & 63)))
                                            _ = v19
                                        } else {
                                            v35 := GenS2{F0: int16(v13)}
                                            fmt.Println(fi6, fi6)
                                            if v14 {
                                                v2 = (v2 << ((v3) & 63))
                                            }
                                            _ = v35
                                        }
                                        if false {
                                            v36 := GenS1{F0: 42}
                                            v37 := GenS1{F0: v12}
                                            _ = v36
                                            _ = v37
                                        } else {
                                            var v38 bool = (v14 && ((v3 + 1) > v3))
                                            fmt.Println(v14, v3, v14)
                                            if len(v4) > 0 {
                                                v4[0] = 2147483648
                                            }
                                            _ = v38
                                        }
                                        fmt.Println(fi6, fi6, v2)
                                        _ = v18
                                    }
                                    _ = v17
                                }
                                _ = v13
                                _ = v14
                            }
                            v5 = v5
                            _ = fi11
                            _ = v12
                        }
                        v4 = append(v4, uint32(v2))
                        _ = v10
                    }
                    var v39 float64 = (1e10 + v5)
                    _ = fi6
                    _ = v39
                }
                v40 := GenS5{F0: 1, F1: 1, F2: 760}
                switch 525 {
                case 162:
                    delete(v1, ("a" + ""))
                default:
                    var v41 int64 = 9223372036854775807
                    _ = v41
                }
                _ = v40
            }
            for fi42 := 0; fi42 < 3; fi42++ {
                var v43 int = fi42
                switch fi42 {
                case -434:
                    v43 = int(v2)
                case 32767:
                    v5 -= 0.1
                case 865:
                    v44 := GenS6{F0: true}
                    _ = v44
                default:
                    switch ("\n\t" + "foo bar") {
                    case "a":
                        v3 += uint(v2)
                    default:
                        for fi45 := 0; fi45 < 12; fi45++ {
                            v2 = (v2 + 944)
                            switch "" {
                            case "\n\t":
                                var v46 string = (("a" + "foo bar") + ("a" + ("" + "hello")))
                                _ = v46
                            }
                            v43 = (v43 & (fi42 >> (((^(v3))) & 63)))
                            _ = fi45
                        }
                    }
                }
                v47 := GenS2{F0: -1}
                _ = fi42
                _ = v43
                _ = v47
            }
        } else {
            if (v2 >= (v2 ^ v2)) {
                if ("hello" > (("\n\t" + "foo bar") + ("foo bar" + "a"))) {
                    if (true && (int32(v2) <= 713)) {
                        var v48 uint8 = 0
                        _ = v48
                    }
                    var v49 uint16 = 0
                    v49 -= v49
                    _ = v49
                } else {
                    for fi50 := 0; fi50 < 18; fi50++ {
                        if (-1 < 32767) {
                            v51 := v1["a"]
                            _ = v51
                        } else {
                            if false {
                                if (0 == 255) {
                                    v3 = (v3 % ((827) | 1))
                                    if (2 != 0) {
                                        var v52 uint64 = 4294967295
                                        for fi53 := 0; fi53 < 11; fi53++ {
                                            fmt.Println(v3, v3, v5)
                                            var v54 int32 = 0
                                            var v55 int16 = -1
                                            _ = fi53
                                            _ = v54
                                            _ = v55
                                        }
                                        v3 -= 65535
                                        _ = v52
                                    }
                                    if false {
                                        v56 := GenS7{F0: -2147483648, F1: (-(fi50)), F2: fi50}
                                        for fi57 := 0; fi57 < 14; fi57++ {
                                            v58 := GenS4{F0: 2}
                                            _ = fi57
                                            _ = v58
                                        }
                                        _ = v56
                                    } else {
                                        if ((v2 >> (((^(v3))) & 63)) >= (v2 & 18446744073709551615)) {
                                            for fi59 := 0; fi59 < 5; fi59++ {
                                                var v60 int16 = 1000
                                                _ = fi59
                                                _ = v60
                                            }
                                            v5 = v5
                                            var v61 int32 = 277
                                            _ = v61
                                        } else {
                                            v5 = v5
                                            for fi62 := 0; fi62 < 12; fi62++ {
                                                var v63 uint16 = 32767
                                                v64 := GenS8{F0: (42 < 0), F1: -128, F2: (!(2147483648 < 336))}
                                                _ = fi62
                                                _ = v63
                                                _ = v64
                                            }
                                        }
                                        var v65 float64 = 3.14159
                                        v2 = 1
                                        _ = v65
                                    }
                                }
                                var v66 int16 = -32768
                                _ = v66
                            }
                        }
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        fmt.Println(v5)
                        _ = fi50
                    }
                    v2 = (v2 | (v2 + 18446744073709551615))
                    v2 += v2
                }
                if (true && ((3.14159 >= -2.25) && (uint(v2) < 255))) {
                    for fi67 := 0; fi67 < 5; fi67++ {
                        v5 -= ((float64(-2.25) * (float64(3.14159) + float64(3.14159))) - (0.1 + 3.14159))
                        _ = fi67
                    }
                } else {
                    var v68 int8 = -119
                    switch ((((("a" + "foo bar") + (("hello" + ("a" + ("foo bar" + ("\n\t" + "a")))) + "a")) + "foo bar") + ((("foo bar" + "foo bar") + "foo bar") + "\n\t")) + "") {
                    case "hello":
                        var v69 float32 = (float32(v5) * (3.14159 * (((float32(1.5) / -2.25) / 0.1) * ((-2.25 - 0.0) * ((0.1 / 3.14159) / 3.14159)))))
                        _ = v69
                    }
                    _ = v68
                }
            }
            var v70 uint32 = 359
            _ = v70
        }
    }
    _ = len(v4)
    for fi71 := 0; fi71 < 20; fi71++ {
        var v72 int32 = 2147483647
        v72 += v72
        _ = fi71
        _ = v72
    }
    var v73 uint64 = v2
    fmt.Println(v2, v73)
    if false {
        switch 502 {
        case -9223372036854775808:
            var v74 uint32 = 1
            _ = v74
        case -786:
            fmt.Println(v3, v73)
        }
        var v75 float64 = v5
        if (!false) {
            v2 = (v73 | 4294967295)
            for fi76 := 0; fi76 < 10; fi76++ {
                for fi77 := 0; fi77 < 20; fi77++ {
                    if (true || (!(!true))) {
                        v3 -= v3
                        v73 = v2
                    } else {
                        v78 := GenS9{F0: (100 == 1), F1: (v73 & (^(v73)))}
                        v1["hello"] = (1 == 65535)
                        _ = v78
                    }
                    v3 += 351
                    _ = fi77
                }
                for fi79 := 0; fi79 < 8; fi79++ {
                    var v80 uint64 = v73
                    var v81 uint32 = 2147483647
                    var v82 int64 = 0
                    if (false && false) {
                        if len(v4) > 0 {
                            v4[0] = v81
                        }
                    }
                    delete(v1, "a")
                    _ = fi79
                    _ = v80
                    _ = v81
                    _ = v82
                }
                _ = fi76
            }
            for fi83 := 0; fi83 < 10; fi83++ {
                switch 137 {
                case -9223372036854775808:
                    for fi84 := 0; fi84 < 7; fi84++ {
                        if (true && (-1 == 111)) {
                            var v85 uint64 = (v73 << ((v3) & 63))
                            var v86 uint16 = 289
                            _ = v85
                            _ = v86
                        }
                        var v87 int64 = 9223372036854775807
                        v88 := GenS10{F0: 42, F1: 42}
                        var v89 string = "hello"
                        _ = fi84
                        _ = v87
                        _ = v88
                        _ = v89
                    }
                case 273:
                    var v90 int32 = 42
                    _ = v90
                case -1:
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                default:
                    switch "\n\t" {
                    case "":
                        v3 -= (v3 + (v3 & 2))
                    case "hello":
                        v4 = append(v4, 2147483647)
                    case "a":
                        for fi91 := 0; fi91 < 10; fi91++ {
                            v1[(("" + "") + "foo bar")] = true
                            var v92 uint8 = 128
                            _ = fi91
                            _ = v92
                        }
                    }
                }
                _ = fi83
            }
        }
        _ = v75
    } else {
        v93 := GenS11{F0: 128}
        _ = v93
    }
    var v94 uint16 = 255
    if true {
        v5 += (v5 - v5)
        var v95 int64 = 9223372036854775807
        v96 := GenS3{F0: (^(v3))}
        _ = v95
        _ = v96
    } else {
        fmt.Println(v2, v2)
        switch 128 {
        case 128:
            fmt.Println(v94, v5, v5)
        case -1:
            v94 += 0
        }
    }
    if (-1 > 2147483647) {
        var v97 int = -139
        for fi98 := 0; fi98 < 17; fi98++ {
            v97 = -9223372036854775808
            switch "foo bar" {
            case "a":
                var v99 uint = (v3 >> ((v3) & 63))
                _ = v99
            case "":
                v100 := GenS7{F0: 42, F1: (v97 * (v97 >> (((^(v3))) & 63))), F2: (fi98 << (((v3 / (((^(v3))) | 1))) & 63))}
                _ = v100
            case "\n\t":
                v5 -= v5
            }
            fmt.Println(v2, v3)
            _ = fi98
        }
        _ = v97
    }
    for fi101 := 0; fi101 < 19; fi101++ {
        v5 = v5
        _ = fi101
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v73:", v73)
    fmt.Println("v94:", v94)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v73
    _ = v94
}
