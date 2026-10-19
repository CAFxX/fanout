package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool; F1 string }
type GenS4 struct { F0 uint; F1 int8 }
type GenS5 struct { F0 int8; F1 uint16; F2 int64 }
type GenS6 struct { F0 uint64; F1 int32; F2 int16 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int32; F1 uint32 }
type GenS10 struct { F0 int; F1 bool }
type GenS11 struct { F0 int }

func main() {
    v1 := []int{}
    var v2 int32 = -2147483648
    var v3 uint64 = 1
    v4 := make(map[string]bool)
    v5 := make(map[int]int)
    if false {
        if ((0 < 0) && true) {
            fmt.Println(v3, v3)
            fmt.Println(v3, v2)
            if (127 <= 2) {
                v3 += 4294967295
            } else {
                fmt.Println(v2)
            }
        }
        v6 := GenS1{F0: 42, F1: 1, F2: -32768}
        v3 -= 485
        _ = v6
    }
    switch "foo bar" {
    case "a":
        fmt.Println(v2)
    }
    var v7 float32 = 1.5
    v7 = v7
    v3 = (v3 & (v3 >> ((226) & 63)))
    if true {
        var v8 uint32 = 0
        for fi9 := 0; fi9 < 14; fi9++ {
            if true {
                if true {
                    switch "foo bar" {
                    case "":
                        switch "hello" {
                        case "":
                            var v10 float64 = 0.1
                            _ = v10
                        case "hello":
                            v8 = v8
                        case "foo bar":
                            for fi11 := 0; fi11 < 17; fi11++ {
                                switch (fi11 & (fi9 * int(v8))) {
                                case -2147483648:
                                    var v12 int64 = -1
                                    _ = v12
                                case -1:
                                    v3 = (v3 ^ (^(v3)))
                                default:
                                    v13 := GenS2{F0: int8(v3)}
                                    _ = v13
                                }
                                v2 = v2
                                v14 := GenS3{F0: false, F1: "foo bar"}
                                _ = fi11
                                _ = v14
                            }
                        }
                    }
                    var v15 float64 = 1e10
                    v15 -= 1e10
                    _ = v15
                } else {
                    var v16 uint8 = 2
                    v17 := GenS4{F0: 948, F1: -104}
                    _ = v16
                    _ = v17
                }
            }
            v8 = v8
            v7 = float32(((1e10 / ((0.0 - ((1e10 / ((0.0 + 1.5) + 0.1)) / (((float64((((0.1 / 3.14159) - v7) + (v7 / ((v7 - ((v7 - (-2.25 + (0.1 + ((v7 - v7) - float32(-2.25))))) * 1e10)) + 3.14159)))) + float64(v7)) / float64(((-2.25 / float32(1e10)) * 3.14159))) * 3.14159))) - 0.0)) + 0.1))
            _ = fi9
        }
        _ = v8
    } else {
        v2 -= (v2 - v2)
        if (!(-437 == (-(v2)))) {
            if false {
                v3 = (^(v3))
                v3 -= v3
                var v18 int32 = v2
                _ = v18
            }
            v7 -= (v7 * (1.5 / v7))
            if (!(-1 <= -95)) {
                for fi19 := 0; fi19 < 15; fi19++ {
                    v3 += v3
                    for fi20 := 0; fi20 < 2; fi20++ {
                        if false {
                            switch fi19 {
                            case -9223372036854775808:
                                switch (fi20 / (((fi20 * (-(fi19)))) | 1)) {
                                case 0:
                                    var v21 uint16 = 42
                                    _ = v21
                                case -870:
                                    var v22 int16 = 42
                                    _ = v22
                                case 2147483647:
                                    v3 = v3
                                }
                            case -397:
                                v23 := GenS5{F0: -119, F1: 42, F2: -1}
                                _ = v23
                            default:
                                switch "" {
                                case "foo bar":
                                    var v24 int16 = 457
                                    _ = v24
                                case "":
                                    if (v2 <= 42) {
                                        if false {
                                            v25 := GenS6{F0: v3, F1: v2, F2: 42}
                                            v7 += v7
                                            switch 128 {
                                            case 32767:
                                                if false {
                                                    var v26 int32 = v2
                                                    _ = v26
                                                } else {
                                                    v2 += (v2 >> ((uint(v3)) & 63))
                                                    var v27 float32 = float32(1.5)
                                                    var v28 int32 = 42
                                                    _ = v27
                                                    _ = v28
                                                }
                                            case -1:
                                                fmt.Println(v2, v2, fi19)
                                                v1 = append(v1, 4294967295)
                                            default:
                                                var v29 float64 = 1e10
                                                _ = v29
                                            }
                                            v30 := GenS7{F0: v3}
                                            _ = v25
                                            _ = v30
                                        } else {
                                            if (42 == -491) {
                                                v31 := GenS8{F0: 9223372036854775807}
                                                if (false || (2 <= 127)) {
                                                    var v32 uint = 42
                                                    v3 = v3
                                                    v33 := GenS8{F0: (fi20 * 946)}
                                                    _ = v32
                                                    _ = v33
                                                }
                                                v34 := GenS7{F0: (v3 ^ (^(v3)))}
                                                _ = v31
                                                _ = v34
                                            } else {
                                                var v35 int8 = 0
                                                _ = v35
                                            }
                                            fmt.Println(fi19)
                                            var v36 string = "a"
                                            _ = v36
                                        }
                                        if ((fi20 | fi20) <= (fi19 >> ((42) & 63))) {
                                            v7 = v7
                                            v37 := GenS2{F0: 0}
                                            var v38 int = fi20
                                            _ = v37
                                            _ = v38
                                        }
                                        v2 = (^(v2))
                                    }
                                case "hello":
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                default:
                                    fmt.Println(v2, v2, v2)
                                }
                            }
                            v39 := GenS9{F0: -2147483648, F1: 2147483648}
                            switch ("a" + "a") {
                            case "\n\t":
                                v2 += (v2 * (^(v2)))
                            case "foo bar":
                                var v40 uint = 4294967295
                                _ = v40
                            default:
                                v7 = v7
                            }
                            v7 = ((v7 + v7) + 0.0)
                            _ = v39
                        }
                        var v41 uint32 = 1
                        v42 := GenS1{F0: uint32(v3), F1: 0, F2: 809}
                        _ = fi20
                        _ = v41
                        _ = v42
                    }
                    fmt.Println(v2)
                    _ = fi19
                }
                for fi43 := 0; fi43 < 1; fi43++ {
                    var v44 float64 = float64((float32(-2.25) + v7))
                    _ = fi43
                    _ = v44
                }
            }
        }
        fmt.Println(v2)
    }
    for fi45 := 0; fi45 < 8; fi45++ {
        var v46 bool = false
        var v47 bool = false
        var v48 bool = v47
        _ = fi45
        _ = v46
        _ = v47
        _ = v48
    }
    v4["foo bar"] = (!false)
    for fi49 := 0; fi49 < 6; fi49++ {
        switch (fi49 << ((2) & 63)) {
        case 600:
            if true {
                for fi50 := 0; fi50 < 10; fi50++ {
                    var v51 bool = false
                    v2 = v2
                    _ = fi50
                    _ = v51
                }
                fmt.Println(fi49, v7, v3)
                v52 := GenS10{F0: (fi49 * 9223372036854775807), F1: true}
                _ = v52
            }
        case 32767:
            if true {
                v53 := v5[2147483647]
                _ = v53
            } else {
                fmt.Println(v3)
            }
        case -2147483648:
            fmt.Println(v3)
        }
        var v54 float64 = 1e10
        _ = fi49
        _ = v54
    }
    if ((float64(v7) == (1.5 + 0.0)) && (!false)) {
        if true {
            if false {
                v7 -= float32((3.14159 + (1e10 - 1.5)))
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                fmt.Println(v7)
            } else {
                var v55 int64 = int64(v2)
                _ = v55
            }
            if true {
                for fi56 := 0; fi56 < 19; fi56++ {
                    var v57 int = (fi56 ^ fi56)
                    v58 := GenS4{F0: 1, F1: 62}
                    v57 = -2147483648
                    _ = fi56
                    _ = v57
                    _ = v58
                }
                var v59 uint16 = 580
                var v60 uint32 = 42
                _ = v59
                _ = v60
            } else {
                for fi61 := 0; fi61 < 5; fi61++ {
                    var v62 float32 = v7
                    fmt.Println(v7, v2)
                    switch (-(fi61)) {
                    case 2147483647:
                        var v63 int8 = -65
                        _ = v63
                    }
                    _ = fi61
                    _ = v62
                }
            }
        }
    } else {
        _ = len(v5)
        v64 := GenS3{F0: ("foo bar" < ("\n\t" + "foo bar")), F1: "hello"}
        var v65 int = -563
        _ = v64
        _ = v65
    }
    _ = len(v1)
    if true {
        v66 := GenS7{F0: (v3 << ((940) & 63))}
        _ = v66
    } else {
        v67 := GenS11{F0: 32767}
        if true {
            fmt.Println(v3, v2)
            if true {
                v7 += v7
                v7 = (1e10 + float32((((((((-2.25 - 1e10) / 3.14159) - (0.0 + 3.14159)) * 3.14159) - -2.25) * 3.14159) / 3.14159)))
            }
            v7 = 3.14159
        } else {
            for fi68 := 0; fi68 < 9; fi68++ {
                v7 -= (float32(float64(v7)) + v7)
                fmt.Println(fi68, v2)
                var v69 float64 = 0.1
                _ = fi68
                _ = v69
            }
        }
        _ = v67
    }
    if (-9223372036854775808 < -785) {
        v2 = v2
    } else {
        if false {
            switch "\n\t" {
            case "\n\t":
                _ = len(v1)
            }
            if (true && false) {
                var v70 string = ("hello" + "\n\t")
                _ = v70
            } else {
                var v71 int = 127
                _ = v71
            }
            v3 = (v3 & v3)
        }
        v2 = 2147483647
    }
    switch "" {
    case "\n\t":
        v72 := GenS11{F0: -9223372036854775808}
        _ = v72
    case "foo bar":
        v3 = (v3 - 2)
    }
    v73 := GenS4{F0: 4294967295, F1: 0}
    v7 -= (1e10 / float32(((-2.25 + 0.0) * 1.5)))
    for fi74 := 0; fi74 < 7; fi74++ {
        var v75 uint16 = 32767
        _ = fi74
        _ = v75
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v73
}
