package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 bool; F1 uint; F2 int64 }
type GenS3 struct { F0 uint; F1 int }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int8; F1 uint; F2 int64 }
type GenS6 struct { F0 uint8; F1 uint8 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int64; F1 int32 }

func main() {
    var v1 float64 = 3.14159
    var v2 uint8 = 42
    v3 := []string{}
    var v4 float64 = (float64(3.14159) + v1)
    var v5 int64 = 167
    var v6 uint8 = v2
    switch ("foo bar" + "a") {
    case "\n\t":
        for fi7 := 0; fi7 < 20; fi7++ {
            if (32767 <= 993) {
                var v8 uint16 = 936
                for fi9 := 0; fi9 < 2; fi9++ {
                    _ = len(v3)
                    switch (("hello" + ("\n\t" + "\n\t")) + "foo bar") {
                    case "a":
                        v5 += (v5 << ((0) & 63))
                    case "foo bar":
                        v8 += v8
                    default:
                        v2 = v2
                    }
                    fmt.Println(v8)
                    _ = fi9
                }
                v2 -= (v2 * v2)
                _ = v8
            }
            switch (fi7 / ((fi7) | 1)) {
            case 9223372036854775807:
                for fi10 := 0; fi10 < 11; fi10++ {
                    var v11 bool = (0 < (fi7 % (((-(fi10))) | 1)))
                    v12 := GenS1{F0: 0, F1: 665}
                    if len(v3) > 0 {
                        v3[0] = (("hello" + (("hello" + "foo bar") + "foo bar")) + "")
                    }
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            default:
                if len(v3) > 0 {
                    v3[0] = ("a" + ("foo bar" + "foo bar"))
                }
            }
            _ = fi7
        }
    case "hello":
        v13 := GenS2{F0: (!(-32768 < -1)), F1: 1, F2: v5}
        _ = v13
    case "":
        switch ("hello" + (((("" + "foo bar") + "hello") + "foo bar") + "hello")) {
        case "\n\t":
            if (false || true) {
                switch ("hello" + "hello") {
                case "foo bar":
                    if ((-245 == -731) && ((!(!true)) && true)) {
                        fmt.Println(v1, v4, v2)
                        for fi14 := 0; fi14 < 7; fi14++ {
                            v15 := GenS3{F0: 833, F1: fi14}
                            v5 -= (^(v5))
                            _ = fi14
                            _ = v15
                        }
                        if true {
                            for fi16 := 0; fi16 < 7; fi16++ {
                                var v17 string = "hello"
                                switch (fi16 >> ((2) & 63)) {
                                case -9223372036854775808:
                                    v5 = (v5 ^ (v5 * (v5 >> ((255) & 63))))
                                case 1:
                                    v17 = v17
                                case 42:
                                    var v18 int16 = 0
                                    _ = v18
                                default:
                                    v19 := GenS4{F0: -1}
                                    _ = v19
                                }
                                var v20 bool = false
                                _ = fi16
                                _ = v17
                                _ = v20
                            }
                        } else {
                            v3 = append(v3, "\n\t")
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            switch "foo bar" {
                            case "hello":
                                switch "" {
                                case "\n\t":
                                    var v21 float32 = 1e10
                                    _ = v21
                                case "hello":
                                    var v22 uint16 = 1
                                    _ = v22
                                case "":
                                    var v23 int = 871
                                    _ = v23
                                default:
                                    if ("" >= "hello") {
                                        if len(v3) > 0 {
                                            v3[0] = ("\n\t" + "\n\t")
                                        }
                                        var v24 uint16 = 992
                                        _ = v24
                                    }
                                }
                            case "foo bar":
                                v25 := GenS5{F0: 100, F1: uint(v2), F2: v5}
                                _ = v25
                            default:
                                v5 = -562
                            }
                        }
                    } else {
                        v3 = append(v3, "hello")
                        for fi26 := 0; fi26 < 4; fi26++ {
                            v6 += 199
                            v2 = (^(v2))
                            fmt.Println(v6, v2, v4)
                            _ = fi26
                        }
                    }
                case "a":
                    fmt.Println(v6, v4)
                default:
                    fmt.Println(v1)
                }
                if true {
                    for fi27 := 0; fi27 < 15; fi27++ {
                        var v28 float64 = (0.0 + v4)
                        fmt.Println(fi27, v6, v4)
                        v1 = ((1e10 + 0.1) * 3.14159)
                        _ = fi27
                        _ = v28
                    }
                    var v29 int16 = 120
                    _ = v29
                }
            }
        }
    default:
        for fi30 := 0; fi30 < 6; fi30++ {
            var v31 int64 = v5
            switch "\n\t" {
            case "a":
                switch 97 {
                case 59:
                    switch ("\n\t" + "") {
                    case "\n\t":
                        v3 = append(v3, "hello")
                    default:
                        var v32 int64 = (v5 * (v5 >> ((4294967295) & 63)))
                        _ = v32
                    }
                case -9223372036854775808:
                    var v33 int8 = 1
                    _ = v33
                default:
                    v34 := GenS4{F0: fi30}
                    _ = v34
                }
            }
            _ = fi30
            _ = v31
        }
    }
    if false {
        _ = len(v3)
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    var v35 int32 = 2147483647
    if (3.14159 > 1.5) {
        if true {
            if (255 <= 18446744073709551615) {
                fmt.Println(v1)
                var v36 string = "\n\t"
                _ = v36
            } else {
                var v37 float64 = 3.14159
                var v38 uint64 = 449
                _ = v37
                _ = v38
            }
            v1 += 1e10
            for fi39 := 0; fi39 < 18; fi39++ {
                var v40 int32 = v35
                _ = fi39
                _ = v40
            }
        }
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi41 := 0; fi41 < 1; fi41++ {
            v4 += v4
            v42 := GenS6{F0: (v2 << ((uint(v2)) & 63)), F1: v2}
            var v43 uint64 = 18446744073709551615
            _ = fi41
            _ = v42
            _ = v43
        }
        v2 += (v2 >> ((324) & 63))
    }
    if false {
        if false {
            fmt.Println(v6, v5)
            for fi44 := 0; fi44 < 17; fi44++ {
                var v45 uint16 = 219
                v46 := GenS2{F0: false, F1: 214, F2: v5}
                v5 = v5
                _ = fi44
                _ = v45
                _ = v46
            }
        }
    } else {
        v47 := GenS3{F0: 0, F1: 4294967295}
        v3 = append(v3, (("foo bar" + "") + "hello"))
        _ = v47
    }
    var v48 uint = 65535
    for fi49 := 0; fi49 < 6; fi49++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        if (0.1 != 1.5) {
            v50 := GenS5{F0: -128, F1: (v48 * (v48 * (^(v48)))), F2: (-(v5))}
            v2 = v6
            _ = v50
        }
        for fi51 := 0; fi51 < 5; fi51++ {
            v4 = 0.1
            var v52 float32 = 1e10
            v53 := GenS4{F0: 1}
            _ = fi51
            _ = v52
            _ = v53
        }
        _ = fi49
    }
    if true {
        v54 := GenS3{F0: v48, F1: 0}
        _ = v54
    } else {
        var v55 uint = v48
        if ((v2 % ((v6) | 1)) != (^(v6))) {
            for fi56 := 0; fi56 < 1; fi56++ {
                v48 = v55
                _ = fi56
            }
            v57 := GenS7{F0: (v48 | (v48 >> (((^(v48))) & 63)))}
            var v58 uint8 = v2
            _ = v57
            _ = v58
        } else {
            v35 = (v35 / ((v35) | 1))
        }
        var v59 int16 = -32768
        _ = v55
        _ = v59
    }
    var v60 int32 = (v35 % (((v35 + (v35 | int32(v6)))) | 1))
    var v61 int16 = 1000
    for fi62 := 0; fi62 < 13; fi62++ {
        var v63 uint8 = (v6 + (v6 << ((v48) & 63)))
        switch fi62 {
        case 128:
            v4 += v1
        case 9223372036854775807:
            v64 := GenS8{F0: v5, F1: (^(v35))}
            _ = v64
        }
        v65 := GenS7{F0: (v48 ^ v48)}
        _ = fi62
        _ = v63
        _ = v65
    }
    v66 := GenS3{F0: v48, F1: 32768}
    if len(v3) > 0 {
        v3[0] = ("\n\t" + "\n\t")
    }
    _ = len(v3)
    for fi67 := 0; fi67 < 3; fi67++ {
        v4 = ((v1 / -2.25) - 1e10)
        _ = len(v3)
        _ = fi67
    }
    for fi68 := 0; fi68 < 12; fi68++ {
        fmt.Println(v2)
        v4 += 0.0
        var v69 int64 = v5
        _ = fi68
        _ = v69
    }
    for fi70 := 0; fi70 < 3; fi70++ {
        if true {
            if (-1 >= -128) {
                fmt.Println(v5, v4)
                var v71 uint64 = 964
                _ = v71
            }
        }
        var v72 int32 = v60
        for fi73 := 0; fi73 < 10; fi73++ {
            var v74 int64 = -1
            _ = fi73
            _ = v74
        }
        _ = fi70
        _ = v72
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v48
    _ = v5
    _ = v6
    _ = v60
    _ = v61
    _ = v66
}
