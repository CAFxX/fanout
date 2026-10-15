package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 int32; F1 uint16; F2 uint8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int16; F1 int }
type GenS6 struct { F0 int32; F1 int16; F2 uint32 }
type GenS7 struct { F0 int; F1 uint32 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int8; F1 int16; F2 int16 }
type GenS10 struct { F0 uint64; F1 uint64 }
type GenS11 struct { F0 uint16; F1 uint64; F2 bool }
type GenS12 struct { F0 int8 }

func main() {
    var v1 uint16 = 65535
    var v2 int64 = -466
    var v3 int32 = -2147483648
    var v4 int16 = -32768
    var v5 uint32 = 1
    v6 := make(map[uint32]uint32)
    v5 += 974
    _ = len(v6)
    v7 := []int{32767, -441}
    v8 := GenS1{F0: v2, F1: 255}
    v9 := GenS2{F0: false, F1: v1}
    if false {
        for fi10 := 0; fi10 < 14; fi10++ {
            v11 := GenS2{F0: false, F1: (v1 & 1)}
            _ = len(v7)
            switch (-(fi10)) {
            case 42:
                if false {
                    v12 := GenS3{F0: (v3 & -2147483648), F1: v1, F2: 43}
                    _ = v12
                } else {
                    switch 4294967295 {
                    case 127:
                        _ = len(v7)
                    case 206:
                        v3 -= (v3 - v3)
                    case 32767:
                        var v13 uint16 = 42
                        _ = v13
                    }
                    for fi14 := 0; fi14 < 20; fi14++ {
                        switch "" {
                        case "":
                            if (v4 == 226) {
                                fmt.Println(v3, v3)
                                var v15 float64 = (float64(1.5) - (((-2.25 + 0.1) + (-2.25 / 1.0)) * 1.5))
                                v2 = (v2 & (v2 >> ((1) & 63)))
                                _ = v15
                            }
                        default:
                            var v16 uint16 = 42
                            _ = v16
                        }
                        if len(v7) > 0 {
                            v7[0] = fi14
                        }
                        _ = fi14
                    }
                    fmt.Println(v1)
                }
            case 2147483647:
                v3 = v3
            default:
                for fi17 := 0; fi17 < 10; fi17++ {
                    var v18 uint = 18446744073709551615
                    if false {
                        var v19 float64 = (-2.25 - 0.0)
                        v4 -= v4
                        v2 += v2
                        if (!((1e10 == 1e10) && ((v5 != 2147483647) && false))) {
                            var v20 float32 = 1e10
                            v21 := GenS4{F0: v3}
                            var v22 string = "hello"
                            _ = v20
                            _ = v21
                            _ = v22
                        }
                        v18 += uint(v1)
                        v23 := v6[(v5 >> (((v18 + 914)) & 63))]
                        _ = v19
                        _ = v23
                    }
                    if false {
                        fmt.Println(v18, fi17)
                        var v24 string = ""
                        v25 := GenS4{F0: 42}
                        _ = v24
                        _ = v25
                    }
                    _ = fi17
                    _ = v18
                }
            }
            _ = fi10
            _ = v11
        }
        switch ("a" + "") {
        case "foo bar":
            var v26 uint64 = 18446744073709551615
            _ = v26
        case "\n\t":
            if (2147483647 == v3) {
                switch 42 {
                case 4294967295:
                    switch "a" {
                    case "":
                        v27 := GenS1{F0: -1, F1: 1}
                        _ = v27
                    case "hello":
                        var v28 float64 = 1e10
                        _ = v28
                    }
                case 9223372036854775807:
                    var v29 bool = (!true)
                    _ = v29
                default:
                    if (!(42 <= 523)) {
                        v30 := GenS1{F0: (v2 >> ((398) & 63)), F1: 0}
                        v2 = 42
                        _ = v30
                    } else {
                        if (!true) {
                            var v31 uint64 = 42
                            _ = v31
                        }
                    }
                }
                var v32 uint32 = v5
                if ((v2 % ((v2) | 1)) >= (v2 << ((778) & 63))) {
                    var v33 int32 = 2147483647
                    fmt.Println(v33)
                    _ = v33
                }
                _ = v32
            }
        default:
            v1 += v1
        }
    } else {
        v5 = 42
        if true {
            var v34 float32 = 1.5
            _ = v34
        }
    }
    for fi35 := 0; fi35 < 19; fi35++ {
        v36 := GenS4{F0: -940}
        if false {
            for fi37 := 0; fi37 < 14; fi37++ {
                switch (fi35 + fi37) {
                case 42:
                    v38 := GenS5{F0: v4, F1: fi35}
                    _ = v38
                }
                var v39 int8 = 59
                _ = fi37
                _ = v39
            }
            _ = len(v7)
        } else {
            v6[42] = (v5 - v5)
            if (false && (4294967295 != 867)) {
                if (65535 == v1) {
                    v40 := GenS6{F0: v3, F1: 1000, F2: v5}
                    _ = v40
                }
            }
        }
        v4 -= int16(fi35)
        _ = fi35
        _ = v36
    }
    if true {
        if true {
            for fi41 := 0; fi41 < 20; fi41++ {
                v42 := GenS7{F0: -1, F1: v5}
                v3 = v3
                for fi43 := 0; fi43 < 2; fi43++ {
                    var v44 int32 = (v3 >> ((2) & 63))
                    var v45 uint8 = 1
                    v45 = v45
                    _ = fi43
                    _ = v44
                    _ = v45
                }
                _ = fi41
                _ = v42
            }
            if (v1 != (v1 % (((^(v1))) | 1))) {
                switch ("" + "hello") {
                case "hello":
                    if (v2 <= v2) {
                        _ = len(v7)
                    }
                }
            }
            if (127 > 127) {
                for fi46 := 0; fi46 < 2; fi46++ {
                    fmt.Println(fi46)
                    var v47 uint64 = 42
                    _ = fi46
                    _ = v47
                }
                if ((v4 << ((1) & 63)) != v4) {
                    var v48 int64 = int64(v3)
                    _ = v48
                } else {
                    if (!true) {
                        for fi49 := 0; fi49 < 8; fi49++ {
                            if false {
                                v5 = v5
                            }
                            v2 = 0
                            var v50 float32 = 0.0
                            _ = fi49
                            _ = v50
                        }
                        v51 := GenS8{F0: ("a" + "foo bar")}
                        _ = v51
                    }
                    var v52 bool = true
                    v3 += v3
                    _ = v52
                }
            } else {
                v3 = v3
                if false {
                    for fi53 := 0; fi53 < 9; fi53++ {
                        v5 -= 434
                        v54 := GenS9{F0: 42, F1: (v4 - (v4 << ((4294967295) & 63))), F2: (v4 / (((^(v4))) | 1))}
                        var v55 uint = 18446744073709551615
                        _ = fi53
                        _ = v54
                        _ = v55
                    }
                    if true {
                        var v56 bool = true
                        _ = v56
                    }
                } else {
                    for fi57 := 0; fi57 < 9; fi57++ {
                        var v58 string = "foo bar"
                        _ = fi57
                        _ = v58
                    }
                    if false {
                        v59 := GenS10{F0: 18446744073709551615, F1: 2}
                        v60 := v6[(^(v5))]
                        var v61 float64 = 1.5
                        _ = v59
                        _ = v60
                        _ = v61
                    }
                }
            }
        } else {
            v7 = append(v7, -614)
            v62 := GenS6{F0: v3, F1: -728, F2: (v5 ^ (v5 | (v5 - 2147483647)))}
            _ = v62
        }
        v2 = (v2 - (v2 - (^(v2))))
    } else {
        _ = len(v7)
        v3 += (v3 + (v3 ^ v3))
        var v63 string = ("foo bar" + ("a" + "a"))
        _ = v63
    }
    v64 := v6[v5]
    fmt.Println(v5, v4, v1)
    v65 := GenS8{F0: ("a" + ("foo bar" + ("hello" + "\n\t")))}
    fmt.Println(v5)
    fmt.Println(v4)
    switch (("\n\t" + ("\n\t" + "a")) + "") {
    case "\n\t":
        for fi66 := 0; fi66 < 5; fi66++ {
            if false {
                var v67 float32 = 1.5
                _ = v67
            } else {
                fmt.Println(v2, v1)
                v68 := GenS7{F0: fi66, F1: (v5 - v5)}
                var v69 int16 = 1000
                _ = v68
                _ = v69
            }
            v70 := GenS8{F0: (("" + "hello") + "foo bar")}
            v71 := GenS7{F0: (^(fi66)), F1: 1}
            _ = fi66
            _ = v70
            _ = v71
        }
    }
    if (v3 > (v3 * 0)) {
        for fi72 := 0; fi72 < 17; fi72++ {
            fmt.Println(v2)
            v1 = v1
            for fi73 := 0; fi73 < 10; fi73++ {
                if ((-(v4)) != int16(v3)) {
                    var v74 uint32 = (v64 % ((v5) | 1))
                    var v75 uint16 = v1
                    _ = len(v7)
                    _ = v74
                    _ = v75
                }
                if (fi73 >= fi72) {
                    v5 -= v64
                    var v76 bool = (10 > 127)
                    if ((!v76) || v76) {
                        var v77 bool = (v76 || v76)
                        v6[1] = (^(v5))
                        v78 := GenS11{F0: 2, F1: 18446744073709551615, F2: (!(32767 <= 2))}
                        _ = v77
                        _ = v78
                    }
                    _ = v76
                }
                fmt.Println(v2, v64)
                _ = fi73
            }
            _ = fi72
        }
        _ = len(v6)
        var v79 int16 = (v4 >> ((42) & 63))
        _ = v79
    }
    for fi80 := 0; fi80 < 13; fi80++ {
        for fi81 := 0; fi81 < 7; fi81++ {
            switch (fi81 / ((fi81) | 1)) {
            case 127:
                v64 += v5
            }
            fmt.Println(fi81)
            if ((v4 & (v4 + (v4 + v4))) == v4) {
                fmt.Println(v4)
                var v82 uint8 = uint8(v1)
                delete(v6, 0)
                if (!(42 >= 120)) {
                    var v83 string = "foo bar"
                    if (true || false) {
                        v84 := GenS1{F0: int64(fi80), F1: (v82 - v82)}
                        _ = v84
                    } else {
                        var v85 uint = 255
                        if false {
                            v3 -= v3
                            switch v83 {
                            case "":
                                switch v83 {
                                case "\n\t":
                                    switch fi80 {
                                    case -9223372036854775808:
                                        if (-1 < v4) {
                                            var v86 float64 = 1.5
                                            _ = v86
                                        }
                                    case 9223372036854775807:
                                        v2 -= v2
                                    case 32768:
                                        if ((v85 << ((341) & 63)) != 65535) {
                                            fmt.Println(v64, fi80)
                                        } else {
                                            v2 = v2
                                            fmt.Println(fi81, v1)
                                        }
                                        if len(v7) > 0 {
                                            v7[0] = fi81
                                        }
                                        switch (fi80 - 42) {
                                        case 4294967295:
                                            var v87 uint = v85
                                            _ = v87
                                        }
                                    }
                                default:
                                    v88 := GenS12{F0: 0}
                                    _ = v88
                                }
                            case "foo bar":
                                var v89 uint16 = 0
                                _ = v89
                            case "\n\t":
                                var v90 uint64 = 18446744073709551615
                                _ = v90
                            }
                            fmt.Println(v5, v83, fi81)
                            if len(v7) > 0 {
                                v7[0] = fi81
                            }
                        }
                        _ = v85
                    }
                    v4 -= v4
                    _ = v83
                }
                _ = v82
            }
            _ = fi81
        }
        _ = fi80
    }
    var v91 string = ("\n\t" + "")
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v64:", v64)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v64
    _ = v65
    _ = v7
    _ = v8
    _ = v9
    _ = v91
}
