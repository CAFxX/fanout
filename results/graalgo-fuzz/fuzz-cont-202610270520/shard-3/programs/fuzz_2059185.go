package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 int64; F1 int; F2 uint }
type GenS3 struct { F0 int8; F1 string; F2 int32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int8; F1 int64 }
func genFn0(p0 uint8, p1 int16) int16 {
    v7 := []int{128, -684, 477, 128}
    v8 := GenS1{F0: int8(p0), F1: 128}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return (-(p1))
}
func genFn1() uint32 {
    for fi9 := 0; fi9 < 14; fi9++ {
        var v10 string = ("a" + ("a" + ("foo bar" + "a")))
        var v11 int16 = 1
        v11 -= v11
        _ = fi9
        _ = v10
        _ = v11
    }
    switch "a" {
    case "\n\t":
        var v12 int64 = 86
        _ = v12
    case "foo bar":
        v13 := GenS1{F0: 1, F1: 1}
        _ = v13
    case "a":
        switch 48 {
        case 1:
            var v14 int32 = -829
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 2; fi15++ {
        var v16 uint64 = 0
        if true {
            for fi17 := 0; fi17 < 8; fi17++ {
                var v18 int16 = 42
                var v19 uint64 = 165
                _ = fi17
                _ = v18
                _ = v19
            }
            var v20 uint16 = uint16(v16)
            var v21 bool = false
            _ = v20
            _ = v21
        }
        v22 := GenS1{F0: 1, F1: fi15}
        _ = fi15
        _ = v16
        _ = v22
    }
    return 2
}

func main() {
    var v1 string = "a"
    v2 := make(map[string]uint32)
    var v3 uint8 = 41
    var v4 int64 = 0
    var v5 string = v1
    v6 := []int{542, 9223372036854775807, 4294967295, -9223372036854775808}
    v5 = ((v5 + ((v1 + v5) + (v1 + v1))) + v1)
    if false {
        var v23 float32 = (((((0.0 - 1e10) - -2.25) - 1.5) + -2.25) - 3.14159)
        if (42 != 113) {
            var v24 int32 = 1
            v25 := GenS1{F0: -1, F1: 0}
            if true {
                var v26 int = -514
                var v27 int8 = 127
                _ = v26
                _ = v27
            } else {
                if (3.14159 != ((0.0 + v23) / v23)) {
                    var v28 int16 = -1
                    _ = v28
                }
                var v29 uint16 = 32767
                v30 := GenS2{F0: 42, F1: int(v29), F2: 42}
                _ = v29
                _ = v30
            }
            _ = v24
            _ = v25
        } else {
            if (8 != -713) {
                v4 = v4
                switch 0 {
                case 9223372036854775807:
                    fmt.Println(v4, v23)
                case -1:
                    var v31 string = v5
                    _ = v31
                case 1:
                    if false {
                        for fi32 := 0; fi32 < 5; fi32++ {
                            v5 = ((("" + ((v1 + "a") + v1)) + "\n\t") + ("hello" + (v1 + (v5 + "hello"))))
                            switch fi32 {
                            case -2147483648:
                                for fi33 := 0; fi33 < 7; fi33++ {
                                    fmt.Println(fi32)
                                    if false {
                                        fmt.Println(v23, fi33)
                                        v4 -= (v4 ^ 9223372036854775807)
                                        var v34 float64 = -2.25
                                        _ = v34
                                    }
                                    switch fi33 {
                                    case 32767:
                                        v5 = v1
                                    default:
                                        var v35 int16 = -558
                                        _ = v35
                                    }
                                    v23 = v23
                                    v23 = (((((v23 - float32(3.14159)) + ((v23 + (((v23 - v23) - v23) / (v23 / (1e10 * -2.25)))) + 1e10)) - v23) * v23) - ((v23 / 1e10) * (v23 - float32((float64((v23 / float32((3.14159 / ((1.5 - (0.0 * 3.14159)) - 3.14159))))) / 1.5)))))
                                    _ = fi33
                                }
                            case 1:
                                v36 := GenS1{F0: 42, F1: (fi32 / ((fi32) | 1))}
                                _ = v36
                            }
                            if (false && false) {
                                var v37 int64 = v4
                                var v38 uint16 = 534
                                _ = v37
                                _ = v38
                            } else {
                                delete(v2, (v1 + (v1 + v1)))
                                if len(v6) > 0 {
                                    v6[0] = fi32
                                }
                                fmt.Println(v5, fi32)
                            }
                            _ = fi32
                        }
                        fmt.Println(v5, v4, v1)
                        var v39 uint16 = 65535
                        _ = v39
                    }
                default:
                    for fi40 := 0; fi40 < 13; fi40++ {
                        v5 = v5
                        v3 = (v3 / ((v3) | 1))
                        v5 = (((("foo bar" + v5) + v1) + v5) + (v5 + ("a" + v1)))
                        _ = fi40
                    }
                }
                v4 = v4
            } else {
                var v41 float64 = 0.0
                v42 := GenS1{F0: 1, F1: 32768}
                _ = v41
                _ = v42
            }
        }
        _ = v23
    }
    var v43 int32 = 1
    switch 32768 {
    case 57:
        v4 += (v4 & v4)
    case -9223372036854775808:
        v44 := GenS1{F0: -1, F1: -1}
        _ = v44
    case 671:
        for fi45 := 0; fi45 < 15; fi45++ {
            var v46 int8 = 1
            if (0 >= 4294967295) {
                switch 1 {
                case 359:
                    for fi47 := 0; fi47 < 19; fi47++ {
                        if true {
                            fmt.Println(v3)
                            var v48 float64 = 0.0
                            fmt.Println(v1)
                            _ = v48
                        } else {
                            _ = len(v6)
                            if ((42 >= 2) && (!false)) {
                                v46 -= (-(v46))
                            } else {
                                fmt.Println(v43, v43, v3)
                                v49 := GenS3{F0: 1, F1: (v1 + ((((v5 + ("hello" + v5)) + v5) + (v5 + (((v1 + (((v5 + (v1 + ("hello" + (v1 + v1)))) + v1) + (v1 + ""))) + "hello") + v1))) + "foo bar")), F2: v43}
                                _ = v49
                            }
                            if false {
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                switch (v5 + v1) {
                                case "foo bar":
                                    v5 = v5
                                case "":
                                    v3 -= 74
                                case "\n\t":
                                    var v50 uint64 = 0
                                    _ = v50
                                }
                                var v51 string = "a"
                                _ = v51
                            }
                        }
                        switch v5 {
                        case "\n\t":
                            switch (("\n\t" + "\n\t") + (((v5 + "\n\t") + (v1 + (v1 + v5))) + v5)) {
                            case "hello":
                                _ = len(v6)
                            }
                        case "hello":
                            delete(v2, "hello")
                        case "foo bar":
                            fmt.Println(v3, v43)
                        }
                        _ = fi47
                    }
                case 2147483647:
                    for fi52 := 0; fi52 < 15; fi52++ {
                        v4 = v4
                        _ = fi52
                    }
                }
                var v53 int8 = (v46 ^ v46)
                _ = v53
            } else {
                if (true && true) {
                    switch (fi45 * fi45) {
                    case -2147483648:
                        fmt.Println(v3, v46, v1)
                    default:
                        var v54 uint16 = 1
                        _ = v54
                    }
                    var v55 bool = (true && false)
                    if (273 == 42) {
                        var v56 int32 = v43
                        delete(v2, "")
                        for fi57 := 0; fi57 < 20; fi57++ {
                            var v58 float64 = (1e10 * 0.0)
                            v59 := GenS2{F0: v4, F1: fi57, F2: 1}
                            switch (fi57 / ((127) | 1)) {
                            case -9223372036854775808:
                                if (42 == 1) {
                                    var v60 int16 = -1
                                    if v55 {
                                        switch v1 {
                                        case "foo bar":
                                            var v61 float32 = (-2.25 + 3.14159)
                                            _ = v61
                                        }
                                        if len(v6) > 0 {
                                            v6[0] = (fi45 << ((2) & 63))
                                        }
                                    }
                                    _ = v60
                                } else {
                                    fmt.Println(v3, v58, v1)
                                }
                            default:
                                if v55 {
                                    switch v1 {
                                    case "\n\t":
                                        var v62 int32 = -1
                                        _ = v62
                                    }
                                    v63 := GenS2{F0: (v4 * (v4 & v4)), F1: fi45, F2: 65535}
                                    v43 -= (-(v56))
                                    fmt.Println(v43)
                                    v43 += (v56 * v56)
                                    v4 += (-(v4))
                                    _ = v63
                                } else {
                                    if v55 {
                                        _ = len(v6)
                                    } else {
                                        v64 := GenS4{F0: (v46 - 1)}
                                        var v65 float64 = ((-2.25 - v58) / v58)
                                        _ = v64
                                        _ = v65
                                    }
                                    if v55 {
                                        var v66 uint64 = 0
                                        var v67 uint64 = (v66 - v66)
                                        var v68 uint8 = (^(v3))
                                        _ = v66
                                        _ = v67
                                        _ = v68
                                    }
                                }
                            }
                            _ = fi57
                            _ = v58
                            _ = v59
                        }
                        _ = v56
                    }
                    _ = v55
                } else {
                    var v69 uint32 = 258
                    _ = v69
                }
            }
            if (false && true) {
                if len(v6) > 0 {
                    v6[0] = (fi45 - fi45)
                }
            } else {
                v2[v5] = 65535
                if len(v6) > 0 {
                    v6[0] = (fi45 ^ (fi45 & (fi45 / ((127) | 1))))
                }
            }
            _ = fi45
            _ = v46
        }
    }
    v3 += 127
    fmt.Println(v4, v43)
    for fi70 := 0; fi70 < 5; fi70++ {
        v71 := GenS5{F0: 100, F1: (v4 + (v4 | v4))}
        var v72 float64 = float64(0.0)
        for fi73 := 0; fi73 < 12; fi73++ {
            fmt.Println(v1, v72, fi73)
            if false {
                fmt.Println(v43)
                v43 -= 0
                v5 = v5
            }
            _ = fi73
        }
        _ = fi70
        _ = v71
        _ = v72
    }
    v43 += (v43 ^ -61)
    if true {
        switch -995 {
        case 474:
            v43 -= v43
        case -1:
            var v74 uint = 0
            _ = v74
        case 401:
            switch "a" {
            case "\n\t":
                v1 = "foo bar"
            case "":
                var v75 int = 4294967295
                _ = v75
            }
        }
    } else {
        v5 = "\n\t"
        delete(v2, v5)
        var v76 float64 = float64(3.14159)
        _ = v76
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}
