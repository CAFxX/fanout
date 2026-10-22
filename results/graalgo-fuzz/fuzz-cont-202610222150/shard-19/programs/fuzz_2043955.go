package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 uint8 }
type GenS2 struct { F0 int64; F1 uint64; F2 uint16 }
type GenS3 struct { F0 uint16; F1 string; F2 uint64 }
type GenS4 struct { F0 int; F1 int64; F2 uint8 }
type GenS5 struct { F0 uint16; F1 bool; F2 uint8 }
type GenS6 struct { F0 uint8; F1 uint8; F2 bool }
type GenS7 struct { F0 int64; F1 int64 }
type GenS8 struct { F0 int16; F1 int; F2 uint8 }

func main() {
    var v1 bool = (((false && (!false)) && false) || (!(true && (true || false))))
    v2 := make(map[int]bool)
    var v3 int16 = 1
    var v4 uint16 = 487
    v5 := []string{"a", "a"}
    v6 := make(map[int]bool)
    var v7 int16 = (v3 ^ v3)
    fmt.Println(v7, v1, v1)
    var v8 string = (("" + (("\n\t" + ("foo bar" + "\n\t")) + ("a" + "foo bar"))) + "a")
    if (!v1) {
        if (451 >= 42) {
            switch v8 {
            case "hello":
                v5 = append(v5, ("hello" + ((v8 + v8) + v8)))
            case "foo bar":
                v8 = "foo bar"
            default:
                v8 = v8
            }
            v1 = (((^(v4)) <= (^(v4))) || (1 < 65535))
        } else {
            for fi9 := 0; fi9 < 1; fi9++ {
                v3 = (v3 << ((222) & 63))
                fmt.Println(fi9)
                _ = fi9
            }
        }
    }
    _ = len(v6)
    fmt.Println(v1, v1, v3)
    var v10 int64 = 0
    v4 = (v4 + (^(v4)))
    for fi11 := 0; fi11 < 13; fi11++ {
        fmt.Println(v10, v7)
        var v12 uint32 = 65535
        _ = fi11
        _ = v12
    }
    var v13 bool = false
    var v14 uint64 = 0
    fmt.Println(v13, v4, v4)
    switch (v8 + ((v8 + (v8 + "")) + v8)) {
    case "hello":
        switch -9223372036854775808 {
        case 1:
            for fi15 := 0; fi15 < 11; fi15++ {
                fmt.Println(v4, v3)
                v4 += 2
                v7 = v3
                _ = fi15
            }
        }
    case "":
        if ((v4 | (v4 << ((0) & 63))) < 2) {
            fmt.Println(v4, v4)
            if false {
                v16 := GenS1{F0: 42, F1: -1, F2: 2}
                _ = v16
            } else {
                fmt.Println(v3, v13)
                v17 := GenS1{F0: 2147483647, F1: (v10 * v10), F2: 127}
                _ = v17
            }
        } else {
            fmt.Println(v8, v3, v10)
            for fi18 := 0; fi18 < 9; fi18++ {
                v13 = v1
                if (v3 != v7) {
                    switch "\n\t" {
                    case "hello":
                        fmt.Println(v14, v1)
                    default:
                        v19 := GenS1{F0: 2147483647, F1: v10, F2: 2}
                        _ = v19
                    }
                    if v1 {
                        for fi20 := 0; fi20 < 1; fi20++ {
                            if (false && v1) {
                                if (true && (42 > -313)) {
                                    v13 = (((127 >= 0) && v1) || true)
                                    var v21 uint8 = 0
                                    _ = v21
                                }
                                v22 := GenS2{F0: v10, F1: (v14 + 856), F2: v4}
                                _ = v22
                            } else {
                                switch (((v8 + v8) + ("a" + ("foo bar" + v8))) + (v8 + v8)) {
                                case "foo bar":
                                    switch (fi18 & fi20) {
                                    case 2147483647:
                                        fmt.Println(v3)
                                    case -282:
                                        if false {
                                            v4 += uint16(v14)
                                        } else {
                                            var v23 float32 = (((float32(float64(0.0)) + 0.0) * 3.14159) / 1.0)
                                            _ = v23
                                        }
                                    case 0:
                                        v8 = "\n\t"
                                    default:
                                        switch -1 {
                                        case -2147483648:
                                            if (2 != 128) {
                                                v24 := GenS2{F0: v10, F1: v14, F2: v4}
                                                _ = v24
                                            } else {
                                                var v25 uint32 = 2147483648
                                                if false {
                                                    fmt.Println(v4)
                                                } else {
                                                    v14 -= 2
                                                    if len(v5) > 0 {
                                                        v5[0] = ("" + v8)
                                                    }
                                                }
                                                var v26 string = v8
                                                if ("\n\t" == v8) {
                                                    var v27 int = fi18
                                                    _ = v27
                                                }
                                                v13 = v13
                                                _ = v25
                                                _ = v26
                                            }
                                        case 0:
                                            v28 := GenS2{F0: (v10 ^ (v10 - v10)), F1: v14, F2: 65535}
                                            v7 -= 1000
                                            if (v14 >= v14) {
                                                var v29 int16 = (v7 & v3)
                                                _ = v29
                                            } else {
                                                switch v8 {
                                                case "a":
                                                    _ = len(v5)
                                                case "\n\t":
                                                    v30 := GenS3{F0: (^(v4)), F1: "\n\t", F2: (v14 >> ((787) & 63))}
                                                    _ = v30
                                                default:
                                                    var v31 uint64 = 1
                                                    _ = v31
                                                }
                                                var v32 int8 = int8(v7)
                                                _ = v32
                                            }
                                            _ = v28
                                        case 32768:
                                            v10 = (v10 >> ((4294967295) & 63))
                                        }
                                    }
                                case "a":
                                    v14 += (^(v14))
                                default:
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                }
                            }
                            _ = fi20
                        }
                    }
                    if (((!v1) && false) && (!v13)) {
                        if false {
                            v33 := GenS4{F0: (fi18 ^ 42), F1: v10, F2: 42}
                            if false {
                                var v34 uint64 = (v14 / (((^(v14))) | 1))
                                var v35 int = fi18
                                _ = v34
                                _ = v35
                            }
                            _ = v33
                        }
                        v5 = append(v5, "\n\t")
                    }
                }
                _ = fi18
            }
        }
    }
    fmt.Println(v7)
    v3 += v7
    if v1 {
        if false {
            var v36 uint16 = (v4 % ((113) | 1))
            if (4294967295 <= 975) {
                var v37 uint64 = 681
                if (42 <= 42) {
                    v13 = v1
                }
                v7 -= (v7 - (v7 - v3))
                _ = v37
            } else {
                for fi38 := 0; fi38 < 3; fi38++ {
                    if ((-(v10)) <= (^(v10))) {
                        switch v8 {
                        case "a":
                            if v1 {
                                fmt.Println(fi38, v13, v13)
                            } else {
                                for fi39 := 0; fi39 < 8; fi39++ {
                                    var v40 uint64 = 208
                                    v41 := v6[int(v10)]
                                    if true {
                                        v8 = (("hello" + "\n\t") + v8)
                                        v40 = v40
                                        switch "foo bar" {
                                        case "hello":
                                            v42 := v6[(^(fi39))]
                                            _ = v42
                                        case "\n\t":
                                            var v43 float32 = 0.0
                                            var v44 int8 = 1
                                            _ = v43
                                            _ = v44
                                        case "a":
                                            fmt.Println(v8)
                                        }
                                    } else {
                                        if len(v5) > 0 {
                                            v5[0] = v8
                                        }
                                        switch 0 {
                                        case -9223372036854775808:
                                            var v45 int32 = 0
                                            _ = v45
                                        case 4294967295:
                                            if false {
                                                if (((0.1 + 3.14159) * 3.14159) >= 0.0) {
                                                    if v1 {
                                                        v46 := GenS5{F0: 774, F1: v41, F2: 146}
                                                        if v41 {
                                                            var v47 float64 = (-2.25 * -2.25)
                                                            var v48 float32 = 1e10
                                                            if (true && true) {
                                                                v49 := GenS3{F0: (^(v4)), F1: v8, F2: v40}
                                                                switch "hello" {
                                                                case "":
                                                                    if v13 {
                                                                        v50 := GenS1{F0: 1, F1: v10, F2: 2}
                                                                        _ = v50
                                                                    } else {
                                                                        v6[fi38] = (!v41)
                                                                        var v51 uint8 = 148
                                                                        _ = v51
                                                                    }
                                                                case "a":
                                                                    var v52 uint32 = uint32(v40)
                                                                    _ = v52
                                                                default:
                                                                    if true {
                                                                        if (32767 != v7) {
                                                                            v13 = (!false)
                                                                            v53 := GenS6{F0: 0, F1: 128, F2: ((v36 > v4) && false)}
                                                                            var v54 int64 = v10
                                                                            _ = v53
                                                                            _ = v54
                                                                        }
                                                                        if len(v5) > 0 {
                                                                            v5[0] = v8
                                                                        }
                                                                    }
                                                                }
                                                                _ = v49
                                                            } else {
                                                                var v55 int32 = -2147483648
                                                                _ = v55
                                                            }
                                                            v5 = append(v5, (((v8 + "\n\t") + v8) + "hello"))
                                                            _ = v47
                                                            _ = v48
                                                        }
                                                        _ = v46
                                                    }
                                                    fmt.Println(v7)
                                                    var v56 float64 = 3.14159
                                                    _ = v56
                                                } else {
                                                    fmt.Println(v1, v13)
                                                    if v13 {
                                                        if len(v5) > 0 {
                                                            v5[0] = ((("foo bar" + "hello") + (("hello" + "") + "a")) + "foo bar")
                                                        }
                                                        v57 := GenS6{F0: 118, F1: 163, F2: ((1.5 == (3.14159 * 1.5)) && v13)}
                                                        var v58 bool = (v41 || v13)
                                                        v59 := GenS7{F0: (v10 << ((0) & 63)), F1: v10}
                                                        _ = v57
                                                        _ = v58
                                                        _ = v59
                                                    } else {
                                                        var v60 uint = 2
                                                        switch fi39 {
                                                        case -9223372036854775808:
                                                            if ((fi39 << ((v60) & 63)) < fi39) {
                                                                var v61 int64 = (^(v10))
                                                                _ = v61
                                                            }
                                                        default:
                                                            if v41 {
                                                                if v41 {
                                                                    var v62 int64 = (-(v10))
                                                                    _ = v62
                                                                } else {
                                                                    v63 := GenS8{F0: (^(v7)), F1: (fi38 + (fi38 - fi38)), F2: 255}
                                                                    v4 = (^(v36))
                                                                    v3 = v7
                                                                    _ = v63
                                                                }
                                                                v5 = append(v5, v8)
                                                                var v64 int64 = (v10 | (-(v10)))
                                                                fmt.Println(v13, v4, v40)
                                                                _ = v64
                                                            } else {
                                                                v14 = v40
                                                            }
                                                        }
                                                        switch (^(fi39)) {
                                                        case 4294967295:
                                                            v36 = uint16(v40)
                                                        }
                                                        _ = v60
                                                    }
                                                }
                                            }
                                        case 9223372036854775807:
                                            var v65 string = v8
                                            _ = len(v6)
                                            _ = v65
                                        }
                                    }
                                    _ = fi39
                                    _ = v40
                                    _ = v41
                                }
                            }
                        case "foo bar":
                            for fi66 := 0; fi66 < 10; fi66++ {
                                v3 += int16(fi38)
                                _ = fi66
                            }
                        }
                        v3 = 0
                        var v67 uint64 = (v14 >> ((255) & 63))
                        _ = v67
                    } else {
                        for fi68 := 0; fi68 < 3; fi68++ {
                            v7 += -32768
                            _ = fi68
                        }
                        v69 := GenS2{F0: (v10 | 1), F1: v14, F2: (^(v4))}
                        _ = v69
                    }
                    var v70 int = fi38
                    _ = fi38
                    _ = v70
                }
                var v71 float64 = 3.14159
                _ = v71
            }
            _ = v36
        }
        v13 = v1
    }
    for fi72 := 0; fi72 < 15; fi72++ {
        for fi73 := 0; fi73 < 11; fi73++ {
            v8 = v8
            var v74 uint16 = v4
            _ = fi73
            _ = v74
        }
        switch -1 {
        case 1:
            var v75 int64 = v10
            _ = v75
        case 9223372036854775807:
            v76 := GenS4{F0: int(v10), F1: 0, F2: 124}
            _ = v76
        case 127:
            v10 = v10
        default:
            fmt.Println(v13)
        }
        if (!(v10 != v10)) {
            var v77 string = "a"
            fmt.Println(v8)
            var v78 uint = 255
            _ = v77
            _ = v78
        }
        _ = fi72
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
