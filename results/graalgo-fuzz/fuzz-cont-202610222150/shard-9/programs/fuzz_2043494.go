package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 bool; F1 uint64 }
type GenS4 struct { F0 int32; F1 uint64 }
type GenS5 struct { F0 uint16; F1 uint }
type GenS6 struct { F0 uint; F1 uint16; F2 int8 }
func genFn0(p0 bool, p1 uint32, p2 uint16) string {
    var v4 float64 = 3.14159
    switch 1 {
    case -1:
        v5 := GenS1{F0: -2147483648, F1: 998}
        _ = v5
    case 128:
        fmt.Println(p2, p2, p0)
    }
    var v6 int8 = 99
    if p0 {
        if p0 {
            var v7 float64 = 1.5
            _ = v7
        } else {
            var v8 int8 = int8(p1)
            if p0 {
                for fi9 := 0; fi9 < 3; fi9++ {
                    var v10 int64 = 644
                    _ = fi9
                    _ = v10
                }
                v6 = v8
            }
            v11 := []int{-1}
            _ = v11
            _ = v8
        }
        p1 -= 2147483648
        if (!p0) {
            var v12 float64 = float64(3.14159)
            for fi13 := 0; fi13 < 4; fi13++ {
                v12 += float64(0.1)
                _ = fi13
            }
            _ = v12
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v6
    return ""
}
func genFn1(p0 bool, p1 int32, p2 int) int32 {
    switch "" {
    case "\n\t":
        var v14 uint8 = 200
        _ = v14
    case "hello":
        switch ("\n\t" + "hello") {
        case "a":
            var v15 int8 = -128
            _ = v15
        case "\n\t":
            p1 -= 42
        case "foo bar":
            v16 := []int{(p2 >> ((162) & 63)), 4294967295}
            _ = v16
        default:
            v17 := GenS1{F0: (p2 << ((65535) & 63)), F1: 255}
            _ = v17
        }
    case "foo bar":
        v18 := GenS2{F0: "hello", F1: ((p2 << ((65535) & 63)) == -284)}
        _ = v18
    default:
        p0 = p0
    }
    p2 = 0
    v19 := []int{p2, (p2 & p2), 101, (p2 & (p2 | 32768))}
    for fi20 := 0; fi20 < 5; fi20++ {
        if true {
            var v21 float32 = 3.14159
            fmt.Println(v21, p0, v21)
            if p0 {
                var v22 uint64 = 0
                _ = v22
            }
            _ = v21
        } else {
            var v23 string = ("foo bar" + ("foo bar" + "a"))
            _ = len(v19)
            _ = v23
        }
        p0 = true
        _ = fi20
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    return p1
}

func main() {
    v1 := make(map[string]int)
    var v2 int = 322
    var v3 string = ((("foo bar" + "\n\t") + ("\n\t" + "hello")) + "hello")
    var v24 int32 = 42
    switch v2 {
    case -1:
        v25 := GenS2{F0: (v3 + "hello"), F1: false}
        _ = v25
    }
    var v26 float32 = -2.25
    if (2147483647 > 562) {
        v1[v3] = (v2 << ((150) & 63))
        fmt.Println(v26, v2)
    }
    v27 := GenS3{F0: (0 != 255), F1: 2}
    v2 = (v2 << ((4294967295) & 63))
    fmt.Println(v2, v26)
    v28 := GenS3{F0: (!(88 >= 67)), F1: 4294967295}
    if (int8(v2) < 127) {
        v29 := GenS2{F0: "foo bar", F1: (42 < 220)}
        _ = v29
    } else {
        var v30 int32 = v24
        var v31 int32 = v30
        _ = v30
        _ = v31
    }
    fmt.Println(v24)
    v26 -= float32(0.0)
    if true {
        var v32 float64 = (1.5 * float64(v26))
        _ = v32
    }
    v33 := GenS4{F0: genFn1(false, v24, (v2 >> ((657) & 63))), F1: 42}
    switch "foo bar" {
    case "foo bar":
        if (true || false) {
            if true {
                v24 -= 1
            } else {
                if true {
                    v24 += 0
                    var v34 int = v2
                    v34 -= v2
                    _ = v34
                } else {
                    fmt.Println(v2)
                    v35 := []int{}
                    delete(v1, "\n\t")
                    _ = v35
                }
                v36 := GenS4{F0: (v24 ^ 1), F1: 2}
                if false {
                    v37 := GenS1{F0: v2, F1: 42}
                    _ = v37
                } else {
                    switch v2 {
                    case 581:
                        switch v2 {
                        case -649:
                            v38 := GenS3{F0: (float32((((((((((((3.14159 + 1e10) * 0.0) + 0.0) + (float64(float32(0.1)) * (-2.25 - 0.0))) + -2.25) + 0.1) + float64(v26)) - float64(v26)) / -2.25) + 1.5) * float64((((((v26 * v26) + (v26 + float32((3.14159 / 1e10)))) * 0.1) + 1.5) - -2.25)))) == ((v26 / v26) - float32(1e10))), F1: 0}
                            _ = v38
                        case 1:
                            v24 -= 2147483647
                        case -50:
                            var v39 uint16 = 2
                            _ = v39
                        }
                    case -964:
                        var v40 uint8 = 165
                        _ = v40
                    case 128:
                        v24 += (v24 / (((^(v24))) | 1))
                    }
                    v2 = v2
                    var v41 int8 = 1
                    _ = v41
                }
                _ = v36
            }
            var v42 int16 = 1
            if ((v26 >= float32(((((0.1 - (0.0 / 1.0)) + 1.5) / ((((((0.1 * (((1e10 + (0.0 + (-2.25 + 1e10))) * (0.1 + -2.25)) + 0.0)) - -2.25) * 1e10) / 1.0) + ((3.14159 + 1.5) + -2.25)) + (-2.25 + 0.0))) * 3.14159))) && (!(!(true || (531 <= v24))))) {
                for fi43 := 0; fi43 < 14; fi43++ {
                    fmt.Println(v2)
                    var v44 uint16 = 32767
                    _ = fi43
                    _ = v44
                }
                fmt.Println(v26, v42)
                switch v3 {
                case "":
                    for fi45 := 0; fi45 < 8; fi45++ {
                        if true {
                            for fi46 := 0; fi46 < 19; fi46++ {
                                v47 := GenS1{F0: (-(fi45)), F1: 601}
                                _ = fi46
                                _ = v47
                            }
                            fmt.Println(v2)
                            switch v3 {
                            case "hello":
                                if ((v42 >> ((4294967295) & 63)) > v42) {
                                    var v48 int64 = 0
                                    v49 := []int{fi45, (fi45 * (fi45 / ((v2) | 1))), fi45, -902}
                                    _ = v48
                                    _ = v49
                                } else {
                                    var v50 uint16 = 2
                                    var v51 int = 2147483647
                                    _ = v50
                                    _ = v51
                                }
                            case "a":
                                switch 9223372036854775807 {
                                case -376:
                                    v26 = v26
                                case -2147483648:
                                    fmt.Println(v24)
                                case 27:
                                    var v52 bool = true
                                    _ = v52
                                }
                            case "\n\t":
                                v42 += 1000
                            default:
                                v53 := GenS5{F0: 1, F1: 389}
                                _ = v53
                            }
                        }
                        fmt.Println(v42, fi45, v26)
                        _ = fi45
                    }
                case "hello":
                    var v54 uint8 = 120
                    _ = v54
                }
            } else {
                if true {
                    var v55 string = v3
                    v26 = 3.14159
                    _ = v55
                } else {
                    var v56 uint64 = 4294967295
                    v26 -= 0.0
                    _ = v56
                }
                fmt.Println(v3)
                for fi57 := 0; fi57 < 17; fi57++ {
                    v26 -= (1.5 * v26)
                    v26 += v26
                    _ = fi57
                }
            }
            _ = v42
        } else {
            var v58 int = 798
            var v59 string = "hello"
            _ = v58
            _ = v59
        }
    case "a":
        v60 := []int{v2, v2, v2}
        _ = v60
    case "":
        switch v2 {
        case -32:
            fmt.Println(v26)
        case -2147483648:
            for fi61 := 0; fi61 < 7; fi61++ {
                fmt.Println(v3, fi61)
                switch (^(v2)) {
                case 32768:
                    delete(v1, "\n\t")
                default:
                    if false {
                        v2 = v2
                        v24 -= 607
                    } else {
                        fmt.Println(fi61, v2)
                        if (181 != 42) {
                            fmt.Println(v2, fi61)
                        } else {
                            v62 := []int{}
                            _ = v62
                        }
                        for fi63 := 0; fi63 < 8; fi63++ {
                            fmt.Println(v2)
                            _ = fi63
                        }
                    }
                }
                _ = fi61
            }
        default:
            if (4294967295 > 42) {
                v64 := GenS6{F0: 835, F1: 0, F2: 0}
                _ = v64
            }
        }
    default:
        switch (v2 << ((2) & 63)) {
        case -312:
            if (65535 >= 393) {
                var v65 uint16 = 274
                _ = v65
            }
        case -9223372036854775808:
            if ((!false) && (!(!(42 == 1)))) {
                v66 := GenS2{F0: "", F1: (-1 > -703)}
                for fi67 := 0; fi67 < 8; fi67++ {
                    v68 := GenS4{F0: genFn1(true, v24, (v2 | v2)), F1: 1}
                    var v69 int8 = 1
                    var v70 int = fi67
                    _ = fi67
                    _ = v68
                    _ = v69
                    _ = v70
                }
                _ = v66
            }
        default:
            v71 := GenS1{F0: v2, F1: 1}
            _ = v71
        }
    }
    if (-660 != 127) {
        for fi72 := 0; fi72 < 4; fi72++ {
            switch "" {
            case "\n\t":
                var v73 int = (v2 - 0)
                _ = v73
            case "foo bar":
                var v74 float32 = ((3.14159 - 0.1) + float32(1.5))
                _ = v74
            default:
                var v75 uint16 = 13
                _ = v75
            }
            _ = fi72
        }
        switch v2 {
        case 32768:
            v2 = -9223372036854775808
        case 9223372036854775807:
            v26 += v26
        default:
            v3 = "hello"
        }
        v2 -= 42
    }
    v26 += (1.5 + (0.0 / v26))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v24
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v33
}
