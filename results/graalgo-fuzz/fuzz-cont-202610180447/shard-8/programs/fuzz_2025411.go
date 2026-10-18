package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int8; F1 int; F2 int8 }
type GenS5 struct { F0 bool; F1 int32; F2 int8 }
type GenS6 struct { F0 int16; F1 string }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 bool }

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := make(map[uint32]bool)
    var v3 uint64 = v1
    v4 := make(map[int]bool)
    v5 := []string{""}
    v3 = v3
    switch "hello" {
    case "foo bar":
        if (!true) {
            switch 1 {
            case -931:
                fmt.Println(v1, v3)
            case 2147483647:
                for fi6 := 0; fi6 < 4; fi6++ {
                    fmt.Println(v1, v1, v1)
                    _ = fi6
                }
            case 32768:
                var v7 int8 = 127
                _ = v7
            default:
                if false {
                    for fi8 := 0; fi8 < 4; fi8++ {
                        var v9 uint8 = 127
                        _ = fi8
                        _ = v9
                    }
                    var v10 string = "a"
                    _ = v10
                }
            }
            v11 := GenS1{F0: -70}
            _ = v11
        } else {
            var v12 uint16 = 2
            _ = v12
        }
    case "a":
        v1 -= (v1 & (v1 - (v3 << ((4294967295) & 63))))
    case "\n\t":
        if true {
            var v13 float32 = -2.25
            if true {
                if false {
                    v3 -= (v3 & 727)
                    v14 := GenS1{F0: 42}
                    if true {
                        var v15 uint16 = 0
                        _ = v15
                    }
                    _ = v14
                } else {
                    var v16 uint = 1
                    v16 -= v16
                    _ = v16
                }
                v17 := GenS2{F0: 42}
                var v18 uint = 485
                _ = v17
                _ = v18
            }
            var v19 uint64 = 4294967295
            _ = v13
            _ = v19
        }
    }
    var v20 int16 = 42
    fmt.Println(v1, v20, v20)
    v21 := GenS3{F0: true}
    switch 32767 {
    case 32768:
        if (!((v1 % (((v1 >> ((255) & 63))) | 1)) != 18446744073709551615)) {
            for fi22 := 0; fi22 < 18; fi22++ {
                var v23 uint16 = 862
                if false {
                    switch (fi22 / ((fi22) | 1)) {
                    case 32768:
                        switch "a" {
                        case "":
                            v20 += -1
                        default:
                            fmt.Println(v20)
                        }
                    case 4294967295:
                        var v24 int = fi22
                        _ = v24
                    case -2147483648:
                        var v25 bool = false
                        _ = v25
                    }
                    v2[964] = false
                }
                _ = fi22
                _ = v23
            }
            v1 -= v1
            v26 := GenS1{F0: v20}
            _ = v26
        } else {
            if true {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
            var v27 float32 = float32((1e10 - -2.25))
            _ = v27
        }
    case -538:
        v28 := v4[504]
        _ = v28
    default:
        v1 = v1
    }
    var v29 uint8 = 42
    switch "" {
    case "a":
        fmt.Println(v20, v1)
    case "foo bar":
        v30 := GenS4{F0: 100, F1: 127, F2: -128}
        _ = v30
    }
    var v31 bool = (true && ("foo bar" > ("" + ("hello" + "\n\t"))))
    for fi32 := 0; fi32 < 1; fi32++ {
        if v31 {
            var v33 uint8 = 2
            if false {
                v34 := GenS3{F0: v31}
                var v35 int8 = -1
                v1 = 18446744073709551615
                _ = v34
                _ = v35
            } else {
                switch (fi32 << ((1) & 63)) {
                case 627:
                    v3 = 0
                default:
                    _ = len(v5)
                }
                switch (fi32 + (fi32 + fi32)) {
                case 9223372036854775807:
                    var v36 uint16 = 1
                    _ = v36
                case 368:
                    v1 = (v3 ^ v1)
                case 32768:
                    var v37 int = 2147483647
                    _ = v37
                }
                if ((true && (!false)) && v31) {
                    for fi38 := 0; fi38 < 17; fi38++ {
                        v20 += (v20 / ((v20) | 1))
                        if v31 {
                            fmt.Println(v20, v3, v29)
                            var v39 uint16 = 0
                            _ = v39
                        } else {
                            if v31 {
                                switch (fi38 - 128) {
                                case -9223372036854775808:
                                    v33 = (v29 - v29)
                                case 127:
                                    var v40 int16 = v20
                                    _ = v40
                                case 42:
                                    v3 += 0
                                default:
                                    v41 := GenS5{F0: true, F1: -1, F2: 42}
                                    fmt.Println(fi38)
                                    _ = v41
                                }
                                var v42 uint64 = v1
                                _ = v42
                            }
                        }
                        _ = fi38
                    }
                } else {
                    if v31 {
                        _ = len(v4)
                    }
                    var v43 int = 570
                    var v44 int64 = 42
                    _ = v43
                    _ = v44
                }
            }
            _ = v33
        } else {
            if true {
                switch (fi32 << ((uint(v29)) & 63)) {
                case -518:
                    var v45 uint16 = 810
                    _ = v45
                case 1:
                    v1 += (v1 | v3)
                case 128:
                    if v31 {
                        var v46 int = fi32
                        v31 = (true && v31)
                        _ = v46
                    } else {
                        for fi47 := 0; fi47 < 15; fi47++ {
                            v48 := GenS6{F0: v20, F1: ("a" + "")}
                            var v49 int64 = 1
                            v20 = v20
                            var v50 uint = 644
                            _ = fi47
                            _ = v48
                            _ = v49
                            _ = v50
                        }
                        switch fi32 {
                        case -2147483648:
                            v51 := GenS1{F0: 1000}
                            _ = v51
                        default:
                            for fi52 := 0; fi52 < 17; fi52++ {
                                delete(v2, 62)
                                _ = fi52
                            }
                        }
                    }
                }
                v53 := GenS3{F0: v31}
                _ = len(v2)
                _ = v53
            }
            if (uint(v29) > 696) {
                var v54 float32 = (0.0 - ((-2.25 - 3.14159) / 1.5))
                _ = v54
            }
        }
        var v55 int16 = v20
        for fi56 := 0; fi56 < 1; fi56++ {
            var v57 uint16 = 65535
            fmt.Println(fi32, fi56, v20)
            v58 := GenS7{F0: 2}
            _ = fi56
            _ = v57
            _ = v58
        }
        _ = fi32
        _ = v55
    }
    v59 := v4[-654]
    var v60 int64 = -9223372036854775808
    for fi61 := 0; fi61 < 12; fi61++ {
        for fi62 := 0; fi62 < 14; fi62++ {
            if false {
                delete(v4, fi61)
                if false {
                    var v63 int = 16
                    _ = v63
                }
                var v64 string = ((("hello" + "a") + "a") + "")
                switch fi62 {
                case 668:
                    v3 = 2
                case 127:
                    v64 = v64
                }
                var v65 int = 4294967295
                _ = v64
                _ = v65
            }
            fmt.Println(v20, v3, fi62)
            _ = fi62
        }
        v3 += v3
        var v66 int64 = (v60 | (v60 / (((v60 << ((uint(v3)) & 63))) | 1)))
        _ = fi61
        _ = v66
    }
    for fi67 := 0; fi67 < 20; fi67++ {
        if ((v60 + v60) < v60) {
            var v68 int16 = (v20 * v20)
            var v69 int16 = 466
            _ = v68
            _ = v69
        } else {
            v29 = (v29 >> ((42) & 63))
            if (!((("" + "hello") >= "hello") || v59)) {
                for fi70 := 0; fi70 < 19; fi70++ {
                    v71 := GenS8{F0: v59}
                    _ = fi70
                    _ = v71
                }
                if (v59 || (1 >= 1)) {
                    v72 := v4[(-(fi67))]
                    for fi73 := 0; fi73 < 10; fi73++ {
                        if ((v60 << ((65535) & 63)) > v60) {
                            if ((!((v59 && v59) || v59)) && v31) {
                                switch "foo bar" {
                                case "":
                                    v74 := GenS1{F0: (v20 >> ((2) & 63))}
                                    _ = v74
                                default:
                                    var v75 uint = 18446744073709551615
                                    _ = v75
                                }
                                var v76 int32 = 1
                                _ = v76
                            }
                            switch fi67 {
                            case 304:
                                v59 = (fi67 == fi67)
                            default:
                                v77 := GenS4{F0: 42, F1: fi73, F2: -1}
                                _ = v77
                            }
                            v60 = v60
                        } else {
                            var v78 uint64 = v1
                            _ = v78
                        }
                        v79 := GenS3{F0: ((0 >= 1) || v31)}
                        var v80 uint16 = 42
                        _ = fi73
                        _ = v79
                        _ = v80
                    }
                    _ = v72
                } else {
                    if (!v31) {
                        var v81 bool = v31
                        for fi82 := 0; fi82 < 7; fi82++ {
                            v2[65535] = v81
                            _ = fi82
                        }
                        v5 = append(v5, "a")
                        _ = v81
                    } else {
                        if true {
                            if v31 {
                                fmt.Println(v1, v59, fi67)
                                if true {
                                    _ = len(v4)
                                }
                                for fi83 := 0; fi83 < 17; fi83++ {
                                    if true {
                                        var v84 uint = uint(v3)
                                        switch (("hello" + "hello") + ("\n\t" + "")) {
                                        case "foo bar":
                                            if (!(!((((!((v29 ^ 128) > (^(v29)))) && (v20 <= 1)) || v31) && v31))) {
                                                v85 := v4[(fi67 + 32768)]
                                                _ = v85
                                            }
                                        }
                                        v84 += (v84 / ((v84) | 1))
                                        _ = v84
                                    } else {
                                        switch fi83 {
                                        case -9223372036854775808:
                                            v86 := GenS9{F0: (-733 < 128)}
                                            _ = v86
                                        }
                                        switch (("" + ("a" + "foo bar")) + "hello") {
                                        case "foo bar":
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                        }
                                    }
                                    var v87 uint64 = (v3 << ((0) & 63))
                                    v1 = (v3 | (v87 >> ((1) & 63)))
                                    _ = fi83
                                    _ = v87
                                }
                            }
                            var v88 string = (("\n\t" + "\n\t") + "")
                            _ = v88
                        }
                        switch "" {
                        case "foo bar":
                            for fi89 := 0; fi89 < 1; fi89++ {
                                v20 = (v20 % (((v20 >> ((18446744073709551615) & 63))) | 1))
                                var v90 uint64 = v1
                                v3 = 557
                                _ = fi89
                                _ = v90
                            }
                        default:
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        }
                    }
                    v3 = (v3 / ((v1) | 1))
                }
            } else {
                if len(v5) > 0 {
                    v5[0] = "hello"
                }
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
        _ = fi67
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v29
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v59
    _ = v60
}
