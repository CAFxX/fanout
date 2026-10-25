package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8; F1 uint64 }
type GenS4 struct { F0 int64; F1 int16 }
type GenS5 struct { F0 int8; F1 uint8; F2 int8 }
func genFn0(p0 int64, p1 bool) uint {
    p0 = 443
    fmt.Println(p0)
    v6 := GenS1{F0: 532}
    _ = p0
    _ = p1
    _ = v6
    return 42
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int = 32767
    var v3 string = "foo bar"
    v4 := []string{((v3 + "\n\t") + "a")}
    v5 := make(map[int]bool)
    var v7 uint8 = 191
    switch "foo bar" {
    case "hello":
        v2 = v2
    }
    switch "hello" {
    case "foo bar":
        v8 := GenS1{F0: 1}
        _ = v8
    }
    for fi9 := 0; fi9 < 9; fi9++ {
        for fi10 := 0; fi10 < 8; fi10++ {
            v3 = "\n\t"
            v2 = (fi10 - (^(fi10)))
            if false {
                v3 = "\n\t"
            } else {
                fmt.Println(v3, fi10, v3)
            }
            _ = fi10
        }
        _ = fi9
    }
    v3 = ""
    var v11 bool = false
    for fi12 := 0; fi12 < 3; fi12++ {
        v13 := GenS1{F0: -32768}
        for fi14 := 0; fi14 < 12; fi14++ {
            v15 := GenS1{F0: 42}
            if true {
                if (0 < 57) {
                    v11 = (!v11)
                }
                delete(v1, (fi14 >> ((2) & 63)))
            }
            v7 += v7
            _ = fi14
            _ = v15
        }
        _ = fi12
        _ = v13
    }
    v11 = ((-1 == 1) || (1 == 65535))
    fmt.Println(v7, v2)
    v7 -= 127
    v16 := GenS1{F0: 497}
    v17 := v1[v2]
    if (9223372036854775807 >= -272) {
        switch (("a" + v3) + ("hello" + "")) {
        case "":
            var v18 int = (v2 ^ v2)
            _ = v18
        case "\n\t":
            for fi19 := 0; fi19 < 4; fi19++ {
                switch "\n\t" {
                case "hello":
                    fmt.Println(v2, v2)
                case "foo bar":
                    v20 := GenS1{F0: 0}
                    _ = v20
                default:
                    v7 += v7
                }
                _ = fi19
            }
        case "hello":
            delete(v5, v2)
        }
        if v11 {
            if (float32(-2.25) <= 3.14159) {
                var v21 bool = v11
                _ = v21
            } else {
                switch (v2 % ((576) | 1)) {
                case 32768:
                    v2 += v2
                case 347:
                    v2 = (v2 - v2)
                }
                var v22 bool = (1 < 0)
                if false {
                    if v22 {
                        fmt.Println(v3, v2, v17)
                        v23 := GenS2{F0: v3}
                        if v22 {
                            delete(v5, 2147483647)
                        } else {
                            var v24 uint64 = 336
                            var v25 string = v3
                            _ = v24
                            _ = v25
                        }
                        _ = v23
                    } else {
                        delete(v5, (v2 ^ v2))
                    }
                } else {
                    var v26 uint64 = 42
                    v27 := GenS2{F0: "foo bar"}
                    for fi28 := 0; fi28 < 20; fi28++ {
                        if true {
                            var v29 bool = v22
                            _ = v29
                        }
                        v11 = v22
                        _ = fi28
                    }
                    _ = v26
                    _ = v27
                }
                _ = v22
            }
        } else {
            v30 := GenS3{F0: 171, F1: 18446744073709551615}
            if (32767 != 42) {
                var v31 uint16 = uint16(v7)
                _ = v31
            } else {
                switch -120 {
                case 0:
                    v17 = v17
                default:
                    v7 += 252
                }
                if v11 {
                    v4 = append(v4, v3)
                    fmt.Println(v2, v17, v17)
                } else {
                    if (0 == 198) {
                        switch "a" {
                        case "foo bar":
                            if (42 == 42) {
                                if (-1 == -1) {
                                    for fi32 := 0; fi32 < 15; fi32++ {
                                        var v33 string = v3
                                        if true {
                                            v34 := GenS4{F0: 1, F1: 1000}
                                            _ = v34
                                        } else {
                                            switch "foo bar" {
                                            case "a":
                                                var v35 bool = (v11 && (!((float32(0.0) + 0.1) >= 1e10)))
                                                _ = v35
                                            default:
                                                fmt.Println(v17)
                                            }
                                            switch (v2 * v2) {
                                            case -2147483648:
                                                v33 = (v3 + ((v33 + "hello") + v3))
                                            case 42:
                                                var v36 uint8 = 0
                                                _ = v36
                                            case 2147483647:
                                                var v37 uint8 = v7
                                                _ = v37
                                            default:
                                                switch "a" {
                                                case "hello":
                                                    switch "a" {
                                                    case "foo bar":
                                                        switch v33 {
                                                        case "a":
                                                            if true {
                                                                v5[fi32] = (v17 >= v17)
                                                                if (!v11) {
                                                                    v33 = v3
                                                                    switch v33 {
                                                                    case "\n\t":
                                                                        var v38 uint32 = 2147483647
                                                                        _ = v38
                                                                    case "":
                                                                        if v11 {
                                                                            for fi39 := 0; fi39 < 5; fi39++ {
                                                                                var v40 uint16 = 239
                                                                                _ = fi39
                                                                                _ = v40
                                                                            }
                                                                        } else {
                                                                            fmt.Println(v3, v11, v7)
                                                                            v41 := GenS3{F0: 229, F1: 231}
                                                                            _ = v41
                                                                        }
                                                                    default:
                                                                        if (9 > 42) {
                                                                            if v11 {
                                                                                v33 = "foo bar"
                                                                                v2 = 42
                                                                            }
                                                                        }
                                                                    }
                                                                } else {
                                                                    if v11 {
                                                                        v42 := v1[(fi32 / ((fi32) | 1))]
                                                                        if v11 {
                                                                            v42 = v17
                                                                            v43 := GenS5{F0: -1, F1: v7, F2: 100}
                                                                            _ = v43
                                                                        }
                                                                        _ = v42
                                                                    } else {
                                                                        _ = len(v5)
                                                                        v44 := GenS2{F0: v3}
                                                                        v2 += fi32
                                                                        _ = v44
                                                                    }
                                                                    v3 = (v33 + "\n\t")
                                                                    _ = len(v5)
                                                                }
                                                            }
                                                        case "\n\t":
                                                            delete(v5, (v2 + (v2 ^ v2)))
                                                        }
                                                    case "\n\t":
                                                        var v45 int64 = 9223372036854775807
                                                        _ = v45
                                                    }
                                                case "\n\t":
                                                    for fi46 := 0; fi46 < 8; fi46++ {
                                                        var v47 bool = v11
                                                        var v48 string = v3
                                                        _ = fi46
                                                        _ = v47
                                                        _ = v48
                                                    }
                                                case "a":
                                                    for fi49 := 0; fi49 < 20; fi49++ {
                                                        v11 = (4294967295 != 1)
                                                        if (!v11) {
                                                            v50 := GenS3{F0: uint8(v17), F1: 0}
                                                            fmt.Println(fi49, fi32)
                                                            _ = v50
                                                        }
                                                        _ = fi49
                                                    }
                                                }
                                            }
                                            if len(v4) > 0 {
                                                v4[0] = v3
                                            }
                                        }
                                        switch 1 {
                                        case -2147483648:
                                            var v51 int16 = 0
                                            _ = v51
                                        case 0:
                                            for fi52 := 0; fi52 < 20; fi52++ {
                                                var v53 uint32 = 0
                                                fmt.Println(v11, fi52)
                                                _ = fi52
                                                _ = v53
                                            }
                                        default:
                                            for fi54 := 0; fi54 < 20; fi54++ {
                                                var v55 bool = ((303 != 414) || ((false || v11) && (v7 != (v7 * v7))))
                                                fmt.Println(v33)
                                                var v56 float32 = float32(-2.25)
                                                _ = fi54
                                                _ = v55
                                                _ = v56
                                            }
                                        }
                                        _ = fi32
                                        _ = v33
                                    }
                                    for fi57 := 0; fi57 < 7; fi57++ {
                                        var v58 string = ("" + (v3 + (v3 + "")))
                                        v11 = false
                                        _ = fi57
                                        _ = v58
                                    }
                                    var v59 uint32 = v17
                                    _ = v59
                                } else {
                                    switch v2 {
                                    case 2147483647:
                                        switch "hello" {
                                        case "hello":
                                            for fi60 := 0; fi60 < 9; fi60++ {
                                                switch 856 {
                                                case 1:
                                                    v2 += int(v7)
                                                case 412:
                                                    for fi61 := 0; fi61 < 4; fi61++ {
                                                        delete(v1, fi60)
                                                        _ = fi61
                                                    }
                                                }
                                                fmt.Println(v3, v2, v2)
                                                _ = fi60
                                            }
                                        case "foo bar":
                                            v62 := GenS4{F0: 9223372036854775807, F1: 1000}
                                            _ = v62
                                        }
                                    case 32767:
                                        switch "\n\t" {
                                        case "a":
                                            if v11 {
                                                var v63 int8 = -128
                                                _ = v63
                                            }
                                        }
                                    case -1:
                                        v17 += 42
                                    }
                                }
                                _ = len(v1)
                            } else {
                                var v64 uint = 42
                                if (false || true) {
                                    v65 := GenS2{F0: "a"}
                                    v66 := GenS1{F0: -32768}
                                    var v67 int8 = -76
                                    _ = v65
                                    _ = v66
                                    _ = v67
                                }
                                _ = v64
                            }
                        case "\n\t":
                            v2 = (v2 + (v2 >> ((4294967295) & 63)))
                        default:
                            var v68 string = v3
                            _ = v68
                        }
                        switch "foo bar" {
                        case "":
                            v2 += (v2 | v2)
                        case "\n\t":
                            for fi69 := 0; fi69 < 2; fi69++ {
                                var v70 int32 = 42
                                _ = fi69
                                _ = v70
                            }
                        case "hello":
                            fmt.Println(v17)
                        }
                    }
                    v71 := v5[28]
                    _ = v71
                }
                switch v3 {
                case "":
                    fmt.Println(v7, v17)
                case "\n\t":
                    if (!v11) {
                        var v72 int8 = -79
                        for fi73 := 0; fi73 < 16; fi73++ {
                            if v11 {
                                switch (fi73 + (v2 << ((2) & 63))) {
                                case 1:
                                    for fi74 := 0; fi74 < 8; fi74++ {
                                        var v75 uint8 = (v7 * (v7 - v7))
                                        _ = fi74
                                        _ = v75
                                    }
                                default:
                                    for fi76 := 0; fi76 < 17; fi76++ {
                                        var v77 int64 = 9223372036854775807
                                        var v78 int16 = -1
                                        v11 = false
                                        _ = fi76
                                        _ = v77
                                        _ = v78
                                    }
                                }
                                if len(v4) > 0 {
                                    v4[0] = ""
                                }
                            }
                            _ = fi73
                        }
                        v79 := GenS2{F0: ((v3 + ((v3 + "") + v3)) + (v3 + (v3 + v3)))}
                        _ = v72
                        _ = v79
                    }
                case "a":
                    if true {
                        var v80 uint32 = (v17 + v17)
                        _ = v80
                    }
                }
            }
            v17 = v17
            _ = v30
        }
        v17 = v17
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
