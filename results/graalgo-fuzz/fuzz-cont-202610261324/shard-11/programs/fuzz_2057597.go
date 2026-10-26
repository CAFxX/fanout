package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint16; F1 string }
type GenS5 struct { F0 uint; F1 uint16 }
type GenS6 struct { F0 uint32; F1 uint }
type GenS7 struct { F0 uint16; F1 bool; F2 int16 }
type GenS8 struct { F0 int32; F1 int16; F2 int16 }
type GenS9 struct { F0 int16; F1 int32 }
type GenS10 struct { F0 string }
func genFn0(p0 int, p1 uint64, p2 uint16) int16 {
    p1 = (p1 >> ((286) & 63))
    v5 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 0
}
func genFn1(p0 bool, p1 int32) int64 {
    var v6 float32 = 1.5
    switch -458 {
    case 2147483647:
        for fi7 := 0; fi7 < 14; fi7++ {
            if p0 {
                v8 := []int{fi7}
                _ = len(v8)
                switch (("foo bar" + "\n\t") + "\n\t") {
                case "":
                    v9 := GenS1{F0: 199, F1: 93}
                    _ = v9
                case "\n\t":
                    p0 = (2147483647 == 614)
                case "a":
                    var v10 int = 128
                    _ = v10
                default:
                    var v11 int32 = p1
                    _ = v11
                }
                _ = v8
            }
            _ = fi7
        }
    case 141:
        fmt.Println(p1, v6)
    case 9223372036854775807:
        switch "\n\t" {
        case "hello":
            var v12 float32 = v6
            _ = v12
        case "foo bar":
            for fi13 := 0; fi13 < 16; fi13++ {
                var v14 float32 = v6
                switch (("a" + "\n\t") + "a") {
                case "foo bar":
                    for fi15 := 0; fi15 < 12; fi15++ {
                        v14 = v6
                        v16 := []int{(fi15 + 128), -1, -215, (^(fi15))}
                        switch ("" + "\n\t") {
                        case "hello":
                            var v17 bool = (!(0 == 0))
                            _ = v17
                        }
                        var v18 int64 = -108
                        _ = fi15
                        _ = v16
                        _ = v18
                    }
                }
                p0 = (!p0)
                _ = fi13
                _ = v14
            }
        case "a":
            p1 = p1
        default:
            switch -494 {
            case -2147483648:
                v19 := []int{146, 206, 677, 32768}
                _ = v19
            case 127:
                p1 += p1
            }
        }
    }
    for fi20 := 0; fi20 < 11; fi20++ {
        v21 := []int{-9223372036854775808}
        switch "hello" {
        case "foo bar":
            v22 := GenS1{F0: 0, F1: 0}
            _ = v22
        case "\n\t":
            switch 240 {
            case 832:
                p1 = p1
            case 0:
                v23 := GenS2{F0: 0, F1: 939}
                _ = v23
            case 568:
                switch (fi20 % (((fi20 + 9223372036854775807)) | 1)) {
                case 32767:
                    for fi24 := 0; fi24 < 9; fi24++ {
                        v25 := GenS3{F0: 255}
                        _ = fi24
                        _ = v25
                    }
                case 128:
                    _ = len(v21)
                default:
                    var v26 int64 = 1
                    _ = v26
                }
            }
        default:
            _ = len(v21)
        }
        _ = fi20
        _ = v21
    }
    _ = p0
    _ = p1
    _ = v6
    return 967
}

func main() {
    v1 := []int{127, -2147483648, -92}
    var v2 uint16 = 547
    var v3 int32 = 534
    v4 := make(map[uint32]uint32)
    var v27 string = "foo bar"
    v28 := GenS1{F0: 1, F1: 42}
    fmt.Println(v2)
    var v29 int8 = 71
    fmt.Println(v29, v29)
    _ = len(v1)
    if (!false) {
        switch v27 {
        case "\n\t":
            var v30 int = 1
            _ = v30
        case "a":
            if false {
                var v31 uint = 147
                v3 += 0
                var v32 uint32 = 2147483648
                _ = v31
                _ = v32
            } else {
                var v33 uint64 = 305
                if false {
                    for fi34 := 0; fi34 < 9; fi34++ {
                        v3 = (v3 / ((v3) | 1))
                        var v35 uint32 = 2
                        _ = fi34
                        _ = v35
                    }
                    v2 = v2
                }
                _ = v33
            }
        default:
            v36 := GenS4{F0: 0, F1: ("a" + v27)}
            _ = v36
        }
        if (false && true) {
            _ = len(v4)
            v37 := GenS5{F0: 65535, F1: (v2 >> ((2) & 63))}
            _ = v37
        } else {
            v27 = (v27 + (v27 + ""))
            v38 := GenS3{F0: 1}
            if false {
                v2 -= (v2 >> ((18446744073709551615) & 63))
            } else {
                var v39 int8 = v29
                switch "foo bar" {
                case "":
                    if false {
                        var v40 string = ("" + ("hello" + v27))
                        v39 = v39
                        fmt.Println(v2, v2)
                        _ = v40
                    }
                case "hello":
                    switch 128 {
                    case 42:
                        var v41 uint32 = 2
                        _ = v41
                    default:
                        v42 := GenS6{F0: 647, F1: 639}
                        _ = v42
                    }
                }
                v43 := GenS7{F0: (v2 ^ v2), F1: true, F2: 0}
                _ = v39
                _ = v43
            }
            _ = v38
        }
        var v44 int16 = 1000
        _ = v44
    } else {
        switch 4294967295 {
        case -9223372036854775808:
            v3 += (v3 % ((v3) | 1))
        case 0:
            v1 = append(v1, -9223372036854775808)
        default:
            for fi45 := 0; fi45 < 10; fi45++ {
                var v46 int = fi45
                var v47 int = (v46 << ((4294967295) & 63))
                _ = fi45
                _ = v46
                _ = v47
            }
        }
    }
    for fi48 := 0; fi48 < 17; fi48++ {
        v3 += (v3 + (-(v3)))
        if true {
            v29 += (v29 * v29)
        }
        _ = fi48
    }
    switch -798 {
    case -2147483648:
        var v49 int16 = -99
        _ = v49
    default:
        for fi50 := 0; fi50 < 17; fi50++ {
            var v51 uint = 2
            _ = fi50
            _ = v51
        }
    }
    switch "foo bar" {
    case "hello":
        v3 = int32(v2)
    case "a":
        v3 = 2147483647
    case "":
        switch v27 {
        case "":
            if (v29 <= (v29 | v29)) {
                var v52 string = v27
                var v53 int32 = v3
                _ = v52
                _ = v53
            }
        default:
            switch 32768 {
            case -191:
                v54 := GenS8{F0: v3, F1: 0, F2: genFn0(1, 0, v2)}
                _ = v54
            case -776:
                v55 := GenS8{F0: (v3 << ((2) & 63)), F1: -32768, F2: -32768}
                _ = v55
            case 0:
                v56 := GenS1{F0: 2, F1: 309}
                _ = v56
            }
        }
    }
    for fi57 := 0; fi57 < 14; fi57++ {
        v58 := GenS9{F0: 0, F1: (v3 & (^(v3)))}
        _ = fi57
        _ = v58
    }
    v1 = append(v1, -2147483648)
    v27 = "foo bar"
    var v59 bool = false
    if v59 {
        if v59 {
            var v60 uint16 = (^(v2))
            if true {
                var v61 int = 42
                if true {
                    switch v27 {
                    case "a":
                        v27 = v27
                    default:
                        switch v27 {
                        case "\n\t":
                            switch (^(v61)) {
                            case 127:
                                fmt.Println(v29)
                            case -680:
                                var v62 float64 = (1.5 / 0.1)
                                _ = v62
                            }
                        default:
                            if v59 {
                                if true {
                                    for fi63 := 0; fi63 < 15; fi63++ {
                                        v59 = false
                                        fmt.Println(v61, v27)
                                        _ = fi63
                                    }
                                    v60 = v2
                                    var v64 int8 = int8(v3)
                                    _ = v64
                                }
                                var v65 uint32 = 333
                                for fi66 := 0; fi66 < 11; fi66++ {
                                    if v59 {
                                        var v67 int64 = -588
                                        v68 := GenS10{F0: "a"}
                                        _ = v67
                                        _ = v68
                                    }
                                    _ = len(v1)
                                    _ = fi66
                                }
                                _ = v65
                            } else {
                                v27 = v27
                            }
                        }
                    }
                } else {
                    fmt.Println(v61, v2, v27)
                    v69 := GenS7{F0: v60, F1: (v29 != v29), F2: 32767}
                    _ = len(v1)
                    _ = v69
                }
                _ = v61
            }
            _ = v60
        }
        var v70 string = "hello"
        _ = v70
    }
    switch "a" {
    case "foo bar":
        if (!(v2 <= (v2 * v2))) {
            for fi71 := 0; fi71 < 4; fi71++ {
                var v72 float32 = float32((((-2.25 - -2.25) * ((3.14159 * (0.0 + 3.14159)) * 0.1)) - (((3.14159 * (((-2.25 - float64(0.0)) + -2.25) / 1.0)) * (((1e10 + (-2.25 * 0.1)) / 1.0) / 0.1)) - float64(1.5))))
                if len(v1) > 0 {
                    v1[0] = fi71
                }
                fmt.Println(v3, v59)
                _ = fi71
                _ = v72
            }
        }
    case "\n\t":
        v1 = append(v1, -2147483648)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v59
}
