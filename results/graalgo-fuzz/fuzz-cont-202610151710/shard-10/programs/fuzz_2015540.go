package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 uint32; F1 uint }
type GenS4 struct { F0 uint16; F1 bool; F2 uint16 }
type GenS5 struct { F0 uint32; F1 uint8 }
type GenS6 struct { F0 uint32; F1 int }
func genFn0(p0 bool) bool {
    for fi6 := 0; fi6 < 15; fi6++ {
        p0 = false
        _ = fi6
    }
    p0 = true
    for fi7 := 0; fi7 < 20; fi7++ {
        fmt.Println(p0, fi7)
        for fi8 := 0; fi8 < 15; fi8++ {
            if p0 {
                if (!(!(false && (-1 >= -21)))) {
                    if (p0 && (235 >= 155)) {
                        switch ((("foo bar" + "\n\t") + "foo bar") + ("hello" + "foo bar")) {
                        case "a":
                            v9 := []int{(-(fi8))}
                            _ = v9
                        default:
                            var v10 uint64 = 42
                            _ = v10
                        }
                    }
                    var v11 string = ""
                    fmt.Println(p0, p0)
                    _ = v11
                } else {
                    fmt.Println(fi8, fi7)
                }
                switch "\n\t" {
                case "foo bar":
                    var v12 int16 = 0
                    _ = v12
                case "hello":
                    var v13 float64 = -2.25
                    _ = v13
                default:
                    if false {
                        fmt.Println(fi8, fi8)
                    } else {
                        var v14 uint64 = 0
                        if p0 {
                            var v15 bool = p0
                            _ = v15
                        } else {
                            switch 128 {
                            case 32768:
                                switch fi7 {
                                case -9223372036854775808:
                                    v16 := []int{(^(fi7))}
                                    _ = v16
                                case 127:
                                    p0 = p0
                                case 998:
                                    v17 := GenS1{F0: 0, F1: (fi7 & fi7)}
                                    _ = v17
                                }
                                switch (fi8 + (^(fi7))) {
                                case 9223372036854775807:
                                    if (!("hello" <= "a")) {
                                        var v18 float32 = 1e10
                                        switch "hello" {
                                        case "\n\t":
                                            p0 = (2147483647 < -2147483648)
                                        case "":
                                            v19 := []int{}
                                            _ = v19
                                        case "a":
                                            v14 -= 0
                                        }
                                        v20 := GenS1{F0: -922, F1: (fi7 | (fi7 ^ (fi7 ^ 127)))}
                                        _ = v18
                                        _ = v20
                                    } else {
                                        v21 := GenS1{F0: -9223372036854775808, F1: 990}
                                        _ = v21
                                    }
                                case 128:
                                    if p0 {
                                        switch "\n\t" {
                                        case "foo bar":
                                            p0 = (0.1 > -2.25)
                                        case "\n\t":
                                            switch fi7 {
                                            case -9223372036854775808:
                                                p0 = (!(v14 == (^(v14))))
                                                if (!p0) {
                                                    var v22 int8 = 127
                                                    fmt.Println(fi7, v14)
                                                    p0 = true
                                                    switch "" {
                                                    case "foo bar":
                                                        p0 = p0
                                                        v23 := []int{(fi8 - fi7)}
                                                        _ = v23
                                                    case "hello":
                                                        p0 = (v22 < (v22 ^ v22))
                                                    case "\n\t":
                                                        v22 = (v22 * (v22 + 100))
                                                    default:
                                                        var v24 int8 = v22
                                                        _ = v24
                                                    }
                                                    fmt.Println(v14)
                                                    _ = v22
                                                } else {
                                                    v25 := []int{127, 42}
                                                    _ = v25
                                                }
                                                fmt.Println(p0, fi8, fi7)
                                            default:
                                                v26 := []int{(fi8 ^ (fi7 | (fi8 << ((758) & 63)))), fi8, (fi7 % ((fi8) | 1)), fi8}
                                                switch "\n\t" {
                                                case "\n\t":
                                                    p0 = false
                                                }
                                                _ = v26
                                            }
                                        default:
                                            var v27 int = (fi8 % ((1) | 1))
                                            _ = v27
                                        }
                                        v28 := []int{}
                                        _ = v28
                                    }
                                default:
                                    var v29 float32 = 0.1
                                    _ = v29
                                }
                                v30 := GenS1{F0: 9223372036854775807, F1: (fi8 + (fi7 | (fi7 >> ((745) & 63))))}
                                _ = v30
                            default:
                                v14 += (v14 >> ((4294967295) & 63))
                            }
                            var v31 float64 = (0.1 / -2.25)
                            var v32 uint64 = v14
                            _ = v31
                            _ = v32
                        }
                        _ = v14
                    }
                }
                v33 := GenS1{F0: 9223372036854775807, F1: fi8}
                _ = v33
            } else {
                p0 = ((fi7 * fi7) >= 32768)
            }
            v34 := []int{(fi7 >> ((1) & 63)), 32767, fi8}
            switch (fi7 << ((533) & 63)) {
            case -2147483648:
                if ("hello" == "hello") {
                    if len(v34) > 0 {
                        v34[0] = (fi8 & fi8)
                    }
                    p0 = p0
                }
            case -1:
                if false {
                    v34 = append(v34, fi8)
                    var v35 int8 = -66
                    _ = v35
                }
            }
            var v36 int16 = -1
            _ = fi8
            _ = v34
            _ = v36
        }
        _ = fi7
    }
    fmt.Println(p0)
    v37 := []int{}
    _ = p0
    _ = v37
    return p0
}
func genFn1() uint16 {
    var v38 uint64 = 308
    v38 = v38
    fmt.Println(v38)
    v39 := GenS1{F0: 974, F1: 42}
    switch ("foo bar" + "") {
    case "\n\t":
        v40 := GenS2{F0: -146, F1: -227}
        _ = v40
    case "hello":
        for fi41 := 0; fi41 < 7; fi41++ {
            var v42 uint = 65535
            switch ("\n\t" + ("a" + "")) {
            case "foo bar":
                fmt.Println(v42, v38)
            case "hello":
                var v43 uint64 = v38
                _ = v43
            case "a":
                v44 := GenS2{F0: -49, F1: 346}
                _ = v44
            }
            _ = fi41
            _ = v42
        }
    }
    _ = v38
    _ = v39
    return 255
}

func main() {
    var v1 float32 = ((float32((0.0 + float64(float32(float64(float32(0.0)))))) + 1e10) - float32(1e10))
    var v2 int = 273
    var v3 float64 = float64(v1)
    v4 := make(map[uint32]uint32)
    v5 := make(map[int]uint32)
    var v45 uint16 = 65535
    var v46 uint8 = 44
    switch "hello" {
    case "\n\t":
        switch v2 {
        case -1:
            for fi47 := 0; fi47 < 9; fi47++ {
                if (!true) {
                    for fi48 := 0; fi48 < 3; fi48++ {
                        var v49 int8 = 62
                        _ = fi48
                        _ = v49
                    }
                    v2 = (fi47 / ((v2) | 1))
                } else {
                    var v50 int8 = 42
                    if true {
                        for fi51 := 0; fi51 < 20; fi51++ {
                            if (0 > 9223372036854775807) {
                                var v52 int64 = 42
                                if false {
                                    switch (("a" + "hello") + ("a" + "hello")) {
                                    case "\n\t":
                                        var v53 uint16 = (^(v45))
                                        _ = v53
                                    case "foo bar":
                                        v50 = v50
                                        fmt.Println(v45, v1)
                                        var v54 int16 = 1
                                        _ = v54
                                    case "hello":
                                        delete(v4, 2147483647)
                                    default:
                                        v55 := GenS1{F0: -1, F1: v2}
                                        _ = v55
                                    }
                                    switch "" {
                                    case "foo bar":
                                        v56 := v5[(v2 ^ -594)]
                                        _ = v56
                                    case "":
                                        var v57 int16 = -32768
                                        _ = v57
                                    }
                                }
                                switch "\n\t" {
                                case "hello":
                                    var v58 float32 = v1
                                    _ = v58
                                case "a":
                                    v2 = (fi51 << ((453) & 63))
                                default:
                                    v2 += 9223372036854775807
                                }
                                _ = v52
                            } else {
                                var v59 uint32 = 210
                                _ = v59
                            }
                            _ = fi51
                        }
                    }
                    var v60 uint16 = (^(v45))
                    _ = v50
                    _ = v60
                }
                _ = fi47
            }
        default:
            v61 := GenS1{F0: 9223372036854775807, F1: 0}
            _ = v61
        }
    }
    var v62 bool = false
    v2 = (v2 & v2)
    v45 = genFn1()
    v2 += v2
    v63 := GenS3{F0: 2147483648, F1: 580}
    if false {
        switch "hello" {
        case "\n\t":
            var v64 float64 = 1e10
            _ = v64
        case "":
            v65 := GenS2{F0: 2147483647, F1: -1}
            _ = v65
        }
        switch v2 {
        case 2147483647:
            for fi66 := 0; fi66 < 20; fi66++ {
                v67 := GenS4{F0: (^(v45)), F1: false, F2: v45}
                _ = fi66
                _ = v67
            }
        case 9223372036854775807:
            if (!true) {
                for fi68 := 0; fi68 < 10; fi68++ {
                    v69 := GenS5{F0: 990, F1: 128}
                    _ = fi68
                    _ = v69
                }
                v70 := []int{(-(v2)), v2}
                _ = v70
            } else {
                v71 := GenS1{F0: 49, F1: v2}
                _ = v71
            }
        case 127:
            var v72 int16 = -429
            _ = v72
        default:
            v2 += v2
        }
    }
    if v62 {
        var v73 int64 = 9223372036854775807
        if v62 {
            for fi74 := 0; fi74 < 6; fi74++ {
                var v75 int = fi74
                v76 := []int{-28, v75}
                _ = fi74
                _ = v75
                _ = v76
            }
            v46 = v46
            fmt.Println(v2, v3)
        } else {
            for fi77 := 0; fi77 < 9; fi77++ {
                fmt.Println(v3, v46)
                if (!((v2 != 42) && false)) {
                    if (v2 != fi77) {
                        fmt.Println(v45, v62, v2)
                        fmt.Println(v62, fi77)
                    }
                    for fi78 := 0; fi78 < 1; fi78++ {
                        var v79 int32 = -2147483648
                        _ = fi78
                        _ = v79
                    }
                } else {
                    for fi80 := 0; fi80 < 14; fi80++ {
                        delete(v5, fi80)
                        var v81 int = (^(v2))
                        var v82 string = "hello"
                        _ = fi80
                        _ = v81
                        _ = v82
                    }
                    v83 := []int{(fi77 >> ((5) & 63)), v2}
                    _ = v83
                }
                _ = fi77
            }
            if v62 {
                delete(v5, v2)
            } else {
                if v62 {
                    var v84 int = v2
                    var v85 uint32 = 795
                    if ("foo bar" != "hello") {
                        var v86 uint16 = uint16(v85)
                        var v87 int = -680
                        _ = v86
                        _ = v87
                    } else {
                        v1 = 0.0
                        if ((v2 | 42) >= v2) {
                            for fi88 := 0; fi88 < 10; fi88++ {
                                switch (v84 * (v2 % (((fi88 & (fi88 | (v84 & 4294967295)))) | 1))) {
                                case 42:
                                    switch v84 {
                                    case 9223372036854775807:
                                        v62 = ((v85 ^ (v85 - (^(v85)))) != 2)
                                    case 576:
                                        var v89 uint32 = v85
                                        _ = v89
                                    }
                                case 470:
                                    fmt.Println(v2, v1, v85)
                                case 9223372036854775807:
                                    v46 += v46
                                default:
                                    v90 := GenS6{F0: (v85 % ((v85) | 1)), F1: -9223372036854775808}
                                    _ = v90
                                }
                                _ = fi88
                            }
                            var v91 string = "\n\t"
                            var v92 string = ""
                            _ = v91
                            _ = v92
                        }
                    }
                    _ = v84
                    _ = v85
                }
            }
        }
        _ = v73
    }
    var v93 uint32 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v62:", v62)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v62
    _ = v63
    _ = v93
}
