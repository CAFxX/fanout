package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int; F1 uint32; F2 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int64; F1 int; F2 string }
type GenS6 struct { F0 uint64; F1 int64 }
type GenS7 struct { F0 uint8; F1 int32; F2 uint16 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 int8; F1 int8; F2 string }
type GenS10 struct { F0 int8; F1 int32; F2 uint8 }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 uint32 }

func main() {
    var v1 int32 = 364
    v2 := []string{"foo bar", "a", "a", "a"}
    var v3 uint32 = 2147483648
    var v4 int16 = -1
    var v5 string = ("" + ("a" + ("foo bar" + "hello")))
    var v6 float64 = 0.1
    if true {
        var v7 float32 = 1e10
        _ = v7
    }
    fmt.Println(v3)
    for fi8 := 0; fi8 < 14; fi8++ {
        v3 += v3
        var v9 bool = ((v4 >> ((41) & 63)) >= v4)
        v10 := GenS1{F0: (v9 || (-1 == 1))}
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := GenS1{F0: ((-(v4)) <= (^(v4)))}
    if true {
        for fi12 := 0; fi12 < 10; fi12++ {
            v4 = v4
            _ = fi12
        }
    } else {
        if true {
            fmt.Println(v6, v6, v3)
            _ = len(v2)
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v13 int = 32767
        _ = v13
    }
    v4 = v4
    switch "\n\t" {
    case "":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v14 := GenS2{F0: -1, F1: v3, F2: ((!((1.5 > ((1e10 - float32((0.1 + float64(3.14159)))) / 3.14159)) || ((!(2 >= 418)) && ((-(v4)) == 1000)))) || (32767 > (v4 ^ 116)))}
    for fi15 := 0; fi15 < 16; fi15++ {
        var v16 uint8 = 128
        if true {
            var v17 string = "hello"
            _ = v17
        }
        _ = fi15
        _ = v16
    }
    var v18 uint64 = 4294967295
    var v19 int32 = (v1 << ((65535) & 63))
    fmt.Println(v18, v4, v4)
    switch v5 {
    case "a":
        switch (v5 + (("" + v5) + v5)) {
        case "a":
            fmt.Println(v4, v3)
        default:
            switch -740 {
            case 4294967295:
                var v20 uint = 0
                _ = v20
            case 143:
                switch 128 {
                case -1:
                    fmt.Println(v18, v19)
                case -718:
                    var v21 uint64 = v18
                    _ = v21
                default:
                    if (0 > -207) {
                        v22 := GenS3{F0: (v18 - v18)}
                        if (42 <= -611) {
                            if (v5 >= "a") {
                                var v23 int16 = v4
                                _ = v23
                            } else {
                                for fi24 := 0; fi24 < 12; fi24++ {
                                    if (false || (127 <= 0)) {
                                        v5 = v5
                                        v25 := GenS4{F0: v4}
                                        _ = v25
                                    } else {
                                        if true {
                                            v1 -= 42
                                            var v26 float32 = 1e10
                                            if true {
                                                var v27 float64 = v6
                                                _ = len(v2)
                                                var v28 uint64 = (v18 % (((v18 % ((v18) | 1))) | 1))
                                                _ = v27
                                                _ = v28
                                            } else {
                                                if true {
                                                    var v29 int32 = 0
                                                    v30 := GenS3{F0: (v18 / (((v18 - v18)) | 1))}
                                                    v26 = v26
                                                    _ = v29
                                                    _ = v30
                                                } else {
                                                    v18 += 1
                                                }
                                            }
                                            _ = v26
                                        } else {
                                            fmt.Println(v18)
                                        }
                                        v5 = "\n\t"
                                        v31 := GenS2{F0: 4294967295, F1: (v3 ^ v3), F2: (false && (!false))}
                                        _ = v31
                                    }
                                    for fi32 := 0; fi32 < 9; fi32++ {
                                        if true {
                                            fmt.Println(v6)
                                            var v33 string = v5
                                            fmt.Println(fi32)
                                            _ = v33
                                        } else {
                                            var v34 int8 = 127
                                            _ = v34
                                        }
                                        v5 = v5
                                        v19 = 0
                                        _ = fi32
                                    }
                                    _ = fi24
                                }
                                switch v5 {
                                case "hello":
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                }
                                for fi35 := 0; fi35 < 6; fi35++ {
                                    if ((v5 == ("a" + "")) && ((fi35 != (-(fi35))) && false)) {
                                        if true {
                                            v36 := GenS5{F0: -909, F1: fi35, F2: v5}
                                            _ = v36
                                        } else {
                                            switch v5 {
                                            case "":
                                                var v37 float32 = (((0.0 - 3.14159) + 0.1) + (((0.1 / -2.25) + 0.0) + (1.5 * 1e10)))
                                                _ = v37
                                            case "a":
                                                for fi38 := 0; fi38 < 19; fi38++ {
                                                    v4 += (^(v4))
                                                    var v39 int32 = (v19 % ((v1) | 1))
                                                    _ = fi38
                                                    _ = v39
                                                }
                                            case "hello":
                                                switch (v5 + ("" + v5)) {
                                                case "foo bar":
                                                    if (((v6 * v6) + 3.14159) < (3.14159 + 0.1)) {
                                                        v40 := GenS5{F0: -1, F1: -743, F2: v5}
                                                        if true {
                                                            for fi41 := 0; fi41 < 15; fi41++ {
                                                                fmt.Println(v4, v6, v6)
                                                                v1 += 0
                                                                _ = fi41
                                                            }
                                                        }
                                                        var v42 int32 = v19
                                                        _ = v40
                                                        _ = v42
                                                    } else {
                                                        _ = len(v2)
                                                        v43 := GenS6{F0: (v18 ^ 4294967295), F1: 42}
                                                        _ = v43
                                                    }
                                                case "hello":
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                }
                                            }
                                            if true {
                                                for fi44 := 0; fi44 < 5; fi44++ {
                                                    var v45 float64 = v6
                                                    var v46 float64 = (v6 - (v6 * -2.25))
                                                    _ = fi44
                                                    _ = v45
                                                    _ = v46
                                                }
                                            }
                                            v47 := GenS2{F0: fi35, F1: 558, F2: (!(!false))}
                                            _ = v47
                                        }
                                        if (-1 < (v4 >> ((379) & 63))) {
                                            v4 = 0
                                            if (-9223372036854775808 >= 1) {
                                                v6 = v6
                                                v18 = (v18 ^ (v18 << ((982) & 63)))
                                            } else {
                                                v4 = (v4 >> ((65535) & 63))
                                                var v48 int8 = 1
                                                v1 += v1
                                                _ = v48
                                            }
                                            if false {
                                                v18 += uint64(v3)
                                                var v49 int16 = (v4 * v4)
                                                var v50 uint8 = 42
                                                _ = v49
                                                _ = v50
                                            }
                                        }
                                    } else {
                                        v51 := GenS7{F0: 1, F1: (v1 << ((0) & 63)), F2: 65535}
                                        v52 := GenS6{F0: 2, F1: 0}
                                        switch 4294967295 {
                                        case -871:
                                            if len(v2) > 0 {
                                                v2[0] = (v5 + "\n\t")
                                            }
                                        case -1:
                                            v53 := GenS4{F0: v4}
                                            _ = v53
                                        case 339:
                                            v54 := GenS8{F0: 187}
                                            _ = v54
                                        default:
                                            _ = len(v2)
                                        }
                                        _ = v51
                                        _ = v52
                                    }
                                    _ = fi35
                                }
                            }
                            var v55 uint64 = v18
                            var v56 int64 = -9223372036854775808
                            _ = v55
                            _ = v56
                        }
                        fmt.Println(v4, v1)
                        _ = v22
                    }
                }
            case -2147483648:
                v57 := GenS5{F0: -9223372036854775808, F1: 2147483647, F2: "a"}
                _ = v57
            default:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
        }
    default:
        if true {
            switch 32768 {
            case 4294967295:
                if (42 == 127) {
                    switch int(v3) {
                    case -767:
                        v6 -= (((v6 * v6) - ((v6 / float64(-2.25)) / v6)) / (v6 + v6))
                    }
                    for fi58 := 0; fi58 < 14; fi58++ {
                        for fi59 := 0; fi59 < 14; fi59++ {
                            var v60 float32 = 3.14159
                            _ = fi59
                            _ = v60
                        }
                        v61 := GenS9{F0: 100, F1: 100, F2: "\n\t"}
                        v4 = (v4 << ((0) & 63))
                        _ = fi58
                        _ = v61
                    }
                    switch -240 {
                    case 424:
                        v1 = (v19 & v19)
                    }
                } else {
                    for fi62 := 0; fi62 < 1; fi62++ {
                        for fi63 := 0; fi63 < 10; fi63++ {
                            v6 = 1e10
                            var v64 int64 = 319
                            v2 = append(v2, v5)
                            switch int(v3) {
                            case -1:
                                if false {
                                    switch ((v5 + v5) + (v5 + v5)) {
                                    case "hello":
                                        if len(v2) > 0 {
                                            v2[0] = "a"
                                        }
                                    case "":
                                        var v65 float32 = 3.14159
                                        _ = v65
                                    }
                                    if (32767 <= 903) {
                                        v1 += (^(v19))
                                    } else {
                                        _ = len(v2)
                                        v66 := GenS10{F0: -128, F1: (v19 | v1), F2: 247}
                                        _ = v66
                                    }
                                    v67 := GenS8{F0: 2}
                                    _ = v67
                                } else {
                                    var v68 int = 4294967295
                                    v6 = ((float64(0.1) + float64(1.5)) * (v6 / -2.25))
                                    _ = v68
                                }
                            default:
                                v5 = v5
                            }
                            _ = fi63
                            _ = v64
                        }
                        _ = fi62
                    }
                }
            default:
                switch "hello" {
                case "":
                    var v69 uint32 = v3
                    _ = v69
                case "\n\t":
                    v70 := GenS7{F0: 128, F1: v1, F2: 409}
                    _ = v70
                case "hello":
                    if true {
                        for fi71 := 0; fi71 < 3; fi71++ {
                            var v72 int64 = -9223372036854775808
                            switch fi71 {
                            case 525:
                                v1 -= (v1 | (-(v1)))
                            default:
                                var v73 uint = 255
                                _ = v73
                            }
                            switch v5 {
                            case "a":
                                switch ((v5 + "hello") + v5) {
                                case "foo bar":
                                    v74 := GenS8{F0: (^(v18))}
                                    _ = v74
                                }
                            default:
                                v75 := GenS11{F0: 424}
                                _ = v75
                            }
                            _ = fi71
                            _ = v72
                        }
                    } else {
                        if false {
                            if ((18446744073709551615 <= (v18 % (((v18 + 18446744073709551615)) | 1))) || true) {
                                for fi76 := 0; fi76 < 3; fi76++ {
                                    v77 := GenS9{F0: -40, F1: 100, F2: (("" + (v5 + v5)) + v5)}
                                    v78 := GenS4{F0: -32768}
                                    _ = fi76
                                    _ = v77
                                    _ = v78
                                }
                                for fi79 := 0; fi79 < 18; fi79++ {
                                    for fi80 := 0; fi80 < 12; fi80++ {
                                        var v81 int64 = -352
                                        if (65535 != 805) {
                                            var v82 int = (fi80 ^ fi79)
                                            _ = v82
                                        } else {
                                            fmt.Println(v5, v5, v81)
                                            var v83 uint32 = uint32(v18)
                                            _ = v83
                                        }
                                        var v84 float32 = -2.25
                                        _ = fi80
                                        _ = v81
                                        _ = v84
                                    }
                                    _ = fi79
                                }
                                v85 := GenS12{F0: (v3 * (v3 >> ((537) & 63)))}
                                _ = v85
                            }
                            var v86 uint32 = v3
                            _ = v86
                        } else {
                            var v87 int = 9223372036854775807
                            _ = len(v2)
                            if len(v2) > 0 {
                                v2[0] = v5
                            }
                            _ = v87
                        }
                        v5 = (v5 + (v5 + "a"))
                    }
                default:
                    var v88 int64 = 430
                    _ = v88
                }
            }
        } else {
            var v89 uint64 = 42
            v90 := GenS7{F0: 2, F1: 0, F2: 1}
            _ = v89
            _ = v90
        }
    }
    if (127 != -1) {
        var v91 uint64 = v18
        for fi92 := 0; fi92 < 9; fi92++ {
            v91 = uint64(v3)
            _ = fi92
        }
        v93 := GenS4{F0: 969}
        _ = v91
        _ = v93
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
