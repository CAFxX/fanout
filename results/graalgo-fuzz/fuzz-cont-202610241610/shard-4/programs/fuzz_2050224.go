package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8; F1 uint64; F2 int8 }
type GenS3 struct { F0 int; F1 string; F2 uint8 }
type GenS4 struct { F0 uint8; F1 string }
type GenS5 struct { F0 uint32; F1 int }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int8; F1 int8; F2 int }
type GenS8 struct { F0 string; F1 string; F2 uint8 }
type GenS9 struct { F0 int32; F1 uint16 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 string; F1 int8; F2 int16 }
type GenS13 struct { F0 string; F1 int8; F2 uint8 }
type GenS14 struct { F0 int8; F1 int64; F2 int16 }
type GenS15 struct { F0 uint; F1 bool; F2 int64 }
type GenS16 struct { F0 uint }

func main() {
    v1 := []int{}
    v2 := make(map[string]string)
    v3 := make(map[string]bool)
    var v4 uint8 = 127
    if false {
        var v5 float32 = ((3.14159 + 0.1) * 1.5)
        for fi6 := 0; fi6 < 7; fi6++ {
            if true {
                var v7 uint16 = 348
                if true {
                    v8 := GenS1{F0: 194}
                    var v9 int = 127
                    _ = v8
                    _ = v9
                } else {
                    for fi10 := 0; fi10 < 16; fi10++ {
                        if false {
                            v4 = (v4 + v4)
                            var v11 int64 = -421
                            v12 := GenS2{F0: (v4 % (((v4 >> ((307) & 63))) | 1)), F1: 18446744073709551615, F2: 100}
                            _ = v11
                            _ = v12
                        }
                        var v13 float64 = 0.0
                        _ = fi10
                        _ = v13
                    }
                    switch fi6 {
                    case 128:
                        v14 := v2["hello"]
                        _ = v14
                    }
                    for fi15 := 0; fi15 < 20; fi15++ {
                        v5 = 3.14159
                        var v16 int = 9223372036854775807
                        v7 = (^(v7))
                        _ = fi15
                        _ = v16
                    }
                }
                if (true || ((("hello" + "\n\t") + "\n\t") != ("hello" + "hello"))) {
                    v17 := GenS1{F0: 18446744073709551615}
                    v4 = v4
                    v18 := v2[""]
                    _ = v17
                    _ = v18
                }
                _ = v7
            }
            var v19 string = ("\n\t" + ("foo bar" + ("a" + "foo bar")))
            var v20 float64 = float64(v5)
            _ = fi6
            _ = v19
            _ = v20
        }
        v5 = v5
        _ = v5
    } else {
        switch -1 {
        case 32768:
            var v21 int16 = 0
            _ = v21
        case -9223372036854775808:
            v22 := GenS2{F0: (v4 * v4), F1: 26, F2: 111}
            _ = v22
        case 471:
            switch "foo bar" {
            case "a":
                var v23 float64 = (((1.5 - 1e10) - (0.0 + 1e10)) - (1.5 / 3.14159))
                _ = v23
            default:
                v24 := GenS3{F0: -1, F1: "hello", F2: v4}
                _ = v24
            }
        default:
            switch "hello" {
            case "":
                var v25 string = "hello"
                _ = v25
            case "hello":
                var v26 int8 = 127
                _ = v26
            }
        }
        fmt.Println(v4)
    }
    _ = len(v3)
    v27 := GenS4{F0: (v4 & 2), F1: ""}
    for fi28 := 0; fi28 < 16; fi28++ {
        fmt.Println(v4, v4)
        _ = fi28
    }
    v3[""] = true
    if true {
        var v29 uint64 = 0
        _ = v29
    } else {
        if (true && true) {
            var v30 int64 = 1
            switch "foo bar" {
            case "a":
                v4 = 127
            }
            _ = v30
        } else {
            if (993 != 42) {
                switch ("\n\t" + "hello") {
                case "hello":
                    if false {
                        switch "hello" {
                        case "":
                            var v31 uint64 = 0
                            _ = v31
                        case "foo bar":
                            for fi32 := 0; fi32 < 2; fi32++ {
                                switch fi32 {
                                case -733:
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                case 32768:
                                    if (false || (!false)) {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        if (2 != 4294967295) {
                                            if (!false) {
                                                v33 := GenS5{F0: 1, F1: 794}
                                                _ = v33
                                            } else {
                                                v34 := GenS1{F0: 18446744073709551615}
                                                v35 := GenS6{F0: 65535}
                                                if (fi32 != fi32) {
                                                    v4 -= 129
                                                    v36 := GenS7{F0: 1, F1: -128, F2: fi32}
                                                    v37 := GenS8{F0: "foo bar", F1: "hello", F2: 226}
                                                    _ = v36
                                                    _ = v37
                                                } else {
                                                    if len(v1) > 0 {
                                                        v1[0] = fi32
                                                    }
                                                }
                                                _ = v34
                                                _ = v35
                                            }
                                        } else {
                                            var v38 int8 = 0
                                            _ = v38
                                        }
                                    }
                                }
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                v39 := GenS8{F0: ("" + "hello"), F1: ("" + ("\n\t" + "a")), F2: (v4 * 54)}
                                _ = fi32
                                _ = v39
                            }
                        case "\n\t":
                            switch "a" {
                            case "hello":
                                v4 -= 100
                            case "":
                                if false {
                                    v40 := GenS9{F0: int32(v4), F1: 0}
                                    for fi41 := 0; fi41 < 4; fi41++ {
                                        v4 = (v4 / ((v4) | 1))
                                        var v42 uint32 = 4294967295
                                        _ = fi41
                                        _ = v42
                                    }
                                    v43 := GenS7{F0: 0, F1: -128, F2: 885}
                                    _ = v40
                                    _ = v43
                                } else {
                                    v44 := GenS9{F0: -623, F1: 65535}
                                    var v45 int16 = 0
                                    _ = v44
                                    _ = v45
                                }
                            }
                        default:
                            if len(v1) > 0 {
                                v1[0] = 1
                            }
                        }
                    }
                }
                v46 := GenS7{F0: 100, F1: 42, F2: 32767}
                var v47 int16 = -32768
                _ = v46
                _ = v47
            } else {
                var v48 int32 = 2147483647
                _ = v48
            }
            v1 = append(v1, -803)
            var v49 int16 = -47
            _ = v49
        }
        if ("\n\t" < "foo bar") {
            switch "\n\t" {
            case "a":
                var v50 uint8 = (v4 - v4)
                _ = v50
            case "\n\t":
                v51 := GenS10{F0: 0}
                _ = v51
            }
            switch -93 {
            case 1:
                v52 := GenS1{F0: 255}
                _ = v52
            case 9223372036854775807:
                fmt.Println(v4)
            case 0:
                switch (((("foo bar" + "a") + "\n\t") + "") + "\n\t") {
                case "\n\t":
                    switch 1 {
                    case 2147483647:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    case 127:
                        v4 = v4
                    case 32767:
                        var v53 uint = 4294967295
                        _ = v53
                    default:
                        var v54 bool = (-9223372036854775808 > -1)
                        _ = v54
                    }
                }
            }
            for fi55 := 0; fi55 < 1; fi55++ {
                var v56 int64 = -1
                v56 = v56
                var v57 uint16 = 42
                _ = fi55
                _ = v56
                _ = v57
            }
        } else {
            for fi58 := 0; fi58 < 2; fi58++ {
                for fi59 := 0; fi59 < 4; fi59++ {
                    v4 -= v4
                    if ((((-2.25 - 0.1) * 3.14159) > (((1e10 + float32(3.14159)) - ((0.0 / (1e10 - 0.0)) * 1.5)) * (float32(0.1) / (1e10 + -2.25)))) || false) {
                        var v60 uint32 = 121
                        var v61 string = "hello"
                        fmt.Println(fi59, v4, v60)
                        switch "hello" {
                        case "\n\t":
                            if (true || false) {
                                var v62 uint8 = (v4 + (v4 / ((v4) | 1)))
                                var v63 int = -257
                                _ = v62
                                _ = v63
                            } else {
                                v60 -= 693
                                var v64 uint = 65535
                                _ = len(v1)
                                var v65 uint32 = (v60 << (((^(v64))) & 63))
                                _ = v64
                                _ = v65
                            }
                        default:
                            if true {
                                v60 -= (^(v60))
                                switch (fi58 >> ((18446744073709551615) & 63)) {
                                case -302:
                                    if true {
                                        v66 := GenS7{F0: 1, F1: -1, F2: fi58}
                                        _ = v66
                                    }
                                case -802:
                                    var v67 uint16 = 2
                                    _ = v67
                                case 67:
                                    switch ("hello" + v61) {
                                    case "foo bar":
                                        switch v61 {
                                        case "hello":
                                            var v68 uint32 = 178
                                            _ = v68
                                        default:
                                            if (-461 >= 2147483647) {
                                                fmt.Println(fi58, fi58, v61)
                                            }
                                        }
                                    default:
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    }
                                default:
                                    if true {
                                        switch fi58 {
                                        case 778:
                                            var v69 float64 = float64(((float32(3.14159) + 0.1) - (1e10 / 1.5)))
                                            _ = v69
                                        }
                                    } else {
                                        v70 := GenS11{F0: -576}
                                        v4 -= v4
                                        _ = v70
                                    }
                                }
                                var v71 int64 = -9223372036854775808
                                _ = v71
                            }
                        }
                        _ = v60
                        _ = v61
                    } else {
                        switch (fi58 / (((fi58 * (^(fi58)))) | 1)) {
                        case 197:
                            v72 := GenS11{F0: -9223372036854775808}
                            _ = v72
                        default:
                            _ = len(v2)
                        }
                    }
                    if true {
                        v73 := GenS12{F0: ("hello" + "a"), F1: 0, F2: -32768}
                        fmt.Println(fi58, fi58)
                        if (!true) {
                            var v74 int64 = 23
                            _ = v74
                        } else {
                            v75 := GenS3{F0: -1, F1: "hello", F2: v4}
                            _ = v75
                        }
                        _ = v73
                    }
                    _ = fi59
                }
                _ = fi58
            }
        }
        switch 32768 {
        case 0:
            v76 := GenS4{F0: v4, F1: ("\n\t" + "hello")}
            _ = v76
        case 32768:
            switch 9223372036854775807 {
            case 987:
                fmt.Println(v4)
            }
        case 31:
            if false {
                if (127 != 9223372036854775807) {
                    v4 = 1
                    v1 = append(v1, 32767)
                    v4 += 127
                } else {
                    if true {
                        for fi77 := 0; fi77 < 13; fi77++ {
                            v78 := GenS4{F0: (v4 & 1), F1: "hello"}
                            _ = fi77
                            _ = v78
                        }
                    } else {
                        v79 := v3["a"]
                        if (false && (v79 || false)) {
                            if false {
                                v79 = (!true)
                                var v80 uint64 = 966
                                v81 := GenS13{F0: ("" + "hello"), F1: 119, F2: (v4 % ((42) | 1))}
                                _ = v80
                                _ = v81
                            }
                            if ((0.1 - ((0.0 * 0.1) + float64(3.14159))) == -2.25) {
                                if (!true) {
                                    if v79 {
                                        var v82 int8 = -1
                                        _ = len(v2)
                                        v4 = v4
                                        _ = v82
                                    }
                                    v4 += v4
                                    v4 += v4
                                }
                                delete(v3, "hello")
                                if v79 {
                                    v83 := GenS2{F0: 128, F1: 18446744073709551615, F2: int8(v4)}
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    v84 := GenS14{F0: 1, F1: -185, F2: int16(v4)}
                                    _ = v83
                                    _ = v84
                                } else {
                                    v85 := GenS15{F0: 354, F1: true, F2: 0}
                                    fmt.Println(v4, v79)
                                    delete(v2, "a")
                                    _ = v85
                                }
                            } else {
                                fmt.Println(v4, v4)
                                for fi86 := 0; fi86 < 11; fi86++ {
                                    v79 = false
                                    var v87 int8 = 1
                                    if len(v1) > 0 {
                                        v1[0] = fi86
                                    }
                                    _ = fi86
                                    _ = v87
                                }
                                var v88 float32 = 0.1
                                _ = v88
                            }
                        }
                        for fi89 := 0; fi89 < 12; fi89++ {
                            v2["\n\t"] = (("hello" + "hello") + "")
                            var v90 int8 = 100
                            _ = fi89
                            _ = v90
                        }
                        _ = v79
                    }
                    v91 := GenS16{F0: 255}
                    var v92 uint16 = 1
                    _ = v91
                    _ = v92
                }
            } else {
                var v93 int8 = 1
                _ = v93
            }
        default:
            v4 = v4
        }
    }
    delete(v2, "\n\t")
    var v94 uint8 = (v4 & v4)
    v95 := GenS12{F0: ("a" + ("foo bar" + ("hello" + ""))), F1: 0, F2: -32768}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v94:", v94)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v94
    _ = v95
}
