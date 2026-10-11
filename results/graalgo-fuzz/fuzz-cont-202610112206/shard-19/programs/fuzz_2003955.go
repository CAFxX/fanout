package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 int; F1 string; F2 uint32 }
type GenS3 struct { F0 string; F1 bool }
type GenS4 struct { F0 bool; F1 uint8; F2 int64 }
type GenS5 struct { F0 uint64; F1 uint8 }
type GenS6 struct { F0 uint64; F1 uint; F2 int16 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint8; F1 int32; F2 bool }

func main() {
    v1 := make(map[string]string)
    var v2 int = 9223372036854775807
    v3 := []uint32{2}
    var v4 bool = (!(-2147483648 != 42))
    v5 := make(map[uint32]bool)
    for fi6 := 0; fi6 < 8; fi6++ {
        for fi7 := 0; fi7 < 7; fi7++ {
            fmt.Println(fi6, v4)
            _ = fi7
        }
        if v4 {
            v8 := GenS1{F0: -283, F1: 1}
            v9 := GenS1{F0: -814, F1: int16(fi6)}
            v10 := GenS2{F0: -483, F1: "", F2: 42}
            _ = v10
            _ = v8
            _ = v9
        } else {
            if ((42 > 42) || false) {
                if (v4 || v4) {
                    if v4 {
                        if v4 {
                            if v4 {
                                for fi11 := 0; fi11 < 3; fi11++ {
                                    var v12 bool = (true || (!v4))
                                    switch "foo bar" {
                                    case "foo bar":
                                        v13 := GenS1{F0: -32768, F1: 32767}
                                        _ = v13
                                    case "":
                                        var v14 uint = 577
                                        _ = v14
                                    }
                                    var v15 string = ("" + "foo bar")
                                    _ = fi11
                                    _ = v12
                                    _ = v15
                                }
                                v4 = (3.14159 < 1.5)
                                v2 = 2147483647
                            }
                            var v16 uint16 = 32767
                            var v17 int64 = 954
                            _ = v16
                            _ = v17
                        }
                        _ = len(v3)
                        if v4 {
                            fmt.Println(v4, v4, v4)
                        }
                    }
                    var v18 int = v2
                    _ = v18
                } else {
                    var v19 int = v2
                    switch ("foo bar" + "") {
                    case "foo bar":
                        var v20 uint32 = 1
                        _ = v20
                    case "hello":
                        var v21 uint32 = 42
                        _ = v21
                    default:
                        if v4 {
                            if (734 == 0) {
                                if v4 {
                                    if (100 == 42) {
                                        fmt.Println(v4, v19, fi6)
                                    } else {
                                        for fi22 := 0; fi22 < 5; fi22++ {
                                            var v23 float64 = ((0.1 + 0.0) * (((0.0 * 3.14159) * 1e10) * 1.5))
                                            var v24 uint8 = 255
                                            v19 = (-(fi22))
                                            v19 = (fi22 % ((fi6) | 1))
                                            var v25 uint8 = (v24 / ((v24) | 1))
                                            _ = fi22
                                            _ = v23
                                            _ = v24
                                            _ = v25
                                        }
                                        for fi26 := 0; fi26 < 15; fi26++ {
                                            v4 = false
                                            var v27 string = (("" + ((("" + (("hello" + "hello") + "\n\t")) + "foo bar") + "\n\t")) + ("hello" + "a"))
                                            _ = fi26
                                            _ = v27
                                        }
                                    }
                                }
                                fmt.Println(v4, v2)
                            }
                            var v28 string = "a"
                            _ = v28
                        } else {
                            var v29 uint16 = 42
                            if (v4 && (!v4)) {
                                var v30 uint64 = 0
                                _ = len(v5)
                                v30 = (v30 / ((0) | 1))
                                _ = v30
                            } else {
                                v2 -= fi6
                                for fi31 := 0; fi31 < 7; fi31++ {
                                    v32 := GenS1{F0: -692, F1: 0}
                                    switch 2147483647 {
                                    case 32767:
                                        v4 = v4
                                        v33 := GenS3{F0: "foo bar", F1: (787 > 544)}
                                        var v34 uint8 = 60
                                        _ = v33
                                        _ = v34
                                    }
                                    _ = fi31
                                    _ = v32
                                }
                            }
                            if true {
                                for fi35 := 0; fi35 < 1; fi35++ {
                                    var v36 string = "a"
                                    _ = fi35
                                    _ = v36
                                }
                                var v37 int8 = 42
                                _ = v37
                            } else {
                                v38 := GenS4{F0: v4, F1: 123, F2: -1}
                                var v39 int = fi6
                                var v40 uint32 = 1
                                _ = v38
                                _ = v39
                                _ = v40
                            }
                            _ = v29
                        }
                    }
                    _ = v19
                }
                v2 = (fi6 ^ fi6)
                v41 := GenS3{F0: "", F1: ((!v4) || true)}
                _ = v41
            } else {
                fmt.Println(v2)
            }
            var v42 int32 = 2147483647
            switch v2 {
            case -590:
                fmt.Println(v4, v42, v4)
            case 9223372036854775807:
                var v43 int = -600
                _ = v43
            default:
                var v44 float32 = (0.0 * float32(3.14159))
                _ = v44
            }
            _ = v42
        }
        v45 := GenS5{F0: 0, F1: 2}
        _ = fi6
        _ = v45
    }
    if ((float64(0.0) * -2.25) != 0.0) {
        if (42 == 1) {
            var v46 int64 = 0
            if v4 {
                if (151 == 39) {
                    v4 = v4
                    v3 = append(v3, 2147483648)
                    v46 += 9223372036854775807
                } else {
                    switch v2 {
                    case 128:
                        v46 = 9223372036854775807
                    case -409:
                        for fi47 := 0; fi47 < 4; fi47++ {
                            switch (-(v2)) {
                            case 2147483647:
                                v4 = (!v4)
                            case 128:
                                v3 = append(v3, 125)
                            case 132:
                                v46 = (v46 >> ((0) & 63))
                            }
                            if ((fi47 >> ((620) & 63)) <= fi47) {
                                v48 := GenS4{F0: v4, F1: 127, F2: (v46 | (v46 - (v46 << ((717) & 63))))}
                                _ = v48
                            } else {
                                for fi49 := 0; fi49 < 10; fi49++ {
                                    if len(v3) > 0 {
                                        v3[0] = 954
                                    }
                                    _ = fi49
                                }
                                var v50 string = "foo bar"
                                _ = v50
                            }
                            v51 := GenS6{F0: 1, F1: 1, F2: 42}
                            _ = fi47
                            _ = v51
                        }
                    case -426:
                        for fi52 := 0; fi52 < 3; fi52++ {
                            for fi53 := 0; fi53 < 20; fi53++ {
                                if (!(!(!v4))) {
                                    v46 -= (v46 ^ (^(v46)))
                                    switch fi53 {
                                    case 4294967295:
                                        switch "\n\t" {
                                        case "\n\t":
                                            var v54 int = v2
                                            _ = v54
                                        }
                                    case 145:
                                        fmt.Println(fi52)
                                    default:
                                        switch "hello" {
                                        case "a":
                                            if true {
                                                var v55 uint = 341
                                                fmt.Println(v2, v55, fi53)
                                                _ = v55
                                            } else {
                                                var v56 string = "a"
                                                _ = v56
                                            }
                                        default:
                                            v4 = v4
                                        }
                                    }
                                } else {
                                    var v57 int32 = 975
                                    switch (v2 << ((791) & 63)) {
                                    case -9223372036854775808:
                                        v57 = (v57 + (-(v57)))
                                    case 4294967295:
                                        if ((fi53 >> ((748) & 63)) != fi52) {
                                            fmt.Println(v4)
                                            var v58 int64 = v46
                                            _ = v58
                                        } else {
                                            fmt.Println(v4, v4, fi53)
                                        }
                                    default:
                                        if len(v3) > 0 {
                                            v3[0] = 42
                                        }
                                    }
                                    _ = v57
                                }
                                _ = fi53
                            }
                            _ = fi52
                        }
                    default:
                        var v59 int16 = 32767
                        _ = v59
                    }
                    if (4294967295 < 160) {
                        var v60 float64 = -2.25
                        v5[302] = true
                        _ = v60
                    }
                    for fi61 := 0; fi61 < 15; fi61++ {
                        for fi62 := 0; fi62 < 19; fi62++ {
                            var v63 int = fi61
                            _ = fi62
                            _ = v63
                        }
                        v4 = (v46 <= (v46 | v46))
                        _ = fi61
                    }
                }
                var v64 float64 = 3.14159
                if false {
                    for fi65 := 0; fi65 < 6; fi65++ {
                        v66 := GenS7{F0: 65535}
                        var v67 float32 = 1e10
                        _ = fi65
                        _ = v66
                        _ = v67
                    }
                    var v68 int = v2
                    _ = v68
                }
                _ = v64
            }
            _ = v46
        } else {
            switch (-(v2)) {
            case 127:
                var v69 int32 = -706
                _ = v69
            }
        }
        v2 = 32768
    }
    v4 = v4
    for fi70 := 0; fi70 < 4; fi70++ {
        if false {
            var v71 float32 = 0.0
            for fi72 := 0; fi72 < 4; fi72++ {
                v2 += -9223372036854775808
                v73 := GenS7{F0: 42}
                switch ("" + "") {
                case "foo bar":
                    v4 = (-630 == (v2 * (fi72 >> ((525) & 63))))
                case "":
                    var v74 uint16 = 42
                    _ = v74
                case "a":
                    v4 = v4
                default:
                    switch ("foo bar" + (((("hello" + "a") + "\n\t") + "\n\t") + ("" + "hello"))) {
                    case "\n\t":
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    case "foo bar":
                        if v4 {
                            if v4 {
                                var v75 uint = 65535
                                _ = v75
                            } else {
                                fmt.Println(fi70, v4)
                                v76 := GenS5{F0: 1, F1: 128}
                                v77 := GenS8{F0: 42, F1: -572, F2: (("" + (("" + "foo bar") + "")) != "a")}
                                _ = v76
                                _ = v77
                            }
                        } else {
                            var v78 float64 = 0.0
                            v79 := GenS1{F0: -534, F1: 32767}
                            v80 := GenS2{F0: (fi70 / (((fi70 & fi70)) | 1)), F1: "\n\t", F2: 919}
                            _ = v78
                            _ = v79
                            _ = v80
                        }
                    }
                }
                _ = fi72
                _ = v73
            }
            var v81 uint32 = 667
            _ = v71
            _ = v81
        }
        _ = fi70
    }
    switch 32767 {
    case 9223372036854775807:
        var v82 string = "\n\t"
        _ = v82
    }
    v83 := v1["\n\t"]
    v84 := GenS1{F0: 32767, F1: -32768}
    v2 += (v2 - v2)
    switch "\n\t" {
    case "a":
        switch 127 {
        case -9223372036854775808:
            v85 := GenS2{F0: (^(v2)), F1: (v83 + v83), F2: 4294967295}
            _ = v85
        case 127:
            var v86 uint = 185
            _ = v86
        }
    default:
        var v87 int = 0
        _ = v87
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v83:", v83)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v83
    _ = v84
}
