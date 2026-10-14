package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 uint32; F1 int64 }
type GenS3 struct { F0 int32; F1 bool; F2 bool }
type GenS4 struct { F0 string; F1 uint16; F2 int32 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint64; F1 uint16; F2 uint8 }
type GenS7 struct { F0 uint64; F1 uint64 }
type GenS8 struct { F0 int64; F1 uint; F2 int16 }
type GenS9 struct { F0 uint8; F1 string }
func genFn0() uint8 {
    var v5 int = -2147483648
    var v6 uint64 = 257
    _ = v5
    _ = v6
    return 2
}
func genFn1(p0 int16) uint {
    for fi7 := 0; fi7 < 18; fi7++ {
        switch "foo bar" {
        case "a":
            v8 := []int{(fi7 - fi7)}
            _ = v8
        case "\n\t":
            p0 = (^(p0))
        case "hello":
            for fi9 := 0; fi9 < 12; fi9++ {
                v10 := GenS1{F0: 42, F1: 1}
                p0 += (p0 | p0)
                fmt.Println(fi9, fi7, p0)
                _ = fi9
                _ = v10
            }
        default:
            var v11 bool = (false && false)
            _ = v11
        }
        _ = fi7
    }
    p0 = 1
    for fi12 := 0; fi12 < 1; fi12++ {
        v13 := GenS2{F0: 42, F1: 42}
        p0 += (p0 / ((p0) | 1))
        for fi14 := 0; fi14 < 2; fi14++ {
            v15 := GenS1{F0: 883, F1: 0}
            v16 := []int{(fi14 + fi14), (fi14 >> ((563) & 63))}
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = fi12
        _ = v13
    }
    fmt.Println(p0)
    fmt.Println(p0)
    _ = p0
    return 18446744073709551615
}

func main() {
    var v1 int32 = 669
    var v2 float64 = -2.25
    v3 := []int{}
    var v4 uint8 = 128
    _ = len(v3)
    if true {
        v17 := GenS3{F0: (v1 << ((4294967295) & 63)), F1: ("hello" == ""), F2: false}
        _ = v17
    }
    _ = len(v3)
    var v18 float64 = v2
    var v19 uint64 = 2
    switch "foo bar" {
    case "\n\t":
        if true {
            var v20 float64 = v2
            for fi21 := 0; fi21 < 8; fi21++ {
                v20 = -2.25
                switch ((("\n\t" + "hello") + "a") + "a") {
                case "\n\t":
                    switch "\n\t" {
                    case "\n\t":
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                case "foo bar":
                    v22 := GenS1{F0: (v19 % ((18446744073709551615) | 1)), F1: 1}
                    _ = v22
                case "a":
                    switch ((("a" + "foo bar") + "a") + "a") {
                    case "a":
                        if true {
                            if (v19 <= v19) {
                                for fi23 := 0; fi23 < 2; fi23++ {
                                    v4 = uint8(v19)
                                    if ((267 > 0) && (2147483647 < (v1 ^ (v1 - int32(fi23))))) {
                                        fmt.Println(fi23)
                                        var v24 int8 = 127
                                        if true {
                                            switch (fi23 - fi21) {
                                            case -9223372036854775808:
                                                switch "\n\t" {
                                                case "foo bar":
                                                    var v25 float64 = v2
                                                    _ = v25
                                                case "\n\t":
                                                    var v26 bool = false
                                                    _ = v26
                                                case "hello":
                                                    if (0.0 <= 1e10) {
                                                        if false {
                                                            var v27 int32 = v1
                                                            if ((fi23 % ((fi21) | 1)) <= (^(fi23))) {
                                                                v2 -= 0.0
                                                            }
                                                            _ = v27
                                                        }
                                                        v18 += -2.25
                                                        var v28 bool = true
                                                        _ = v28
                                                    } else {
                                                        if (167 < 65535) {
                                                            fmt.Println(fi21, v2)
                                                        } else {
                                                            _ = len(v3)
                                                        }
                                                    }
                                                }
                                            default:
                                                if true {
                                                    v20 -= ((-2.25 / v2) * (v2 * (v20 + ((v20 + float64((3.14159 - 1.5))) - 3.14159))))
                                                } else {
                                                    var v29 uint8 = v4
                                                    if (0 < genFn1(-634)) {
                                                        if (true || (fi23 < fi21)) {
                                                            v24 -= (v24 ^ v24)
                                                        } else {
                                                            v4 = (v29 % (((v4 + (v29 * v29))) | 1))
                                                            fmt.Println(v4, v2, fi23)
                                                        }
                                                        fmt.Println(v2)
                                                    }
                                                    var v30 uint8 = (v4 & genFn0())
                                                    switch "foo bar" {
                                                    case "hello":
                                                        if (true && (!(979 == uint(v19)))) {
                                                            switch "\n\t" {
                                                            case "hello":
                                                                v19 -= (v19 & v19)
                                                            case "foo bar":
                                                                fmt.Println(v19, v29, fi23)
                                                            }
                                                            v2 -= ((((float64(float32(v20)) - float64((1.5 + 0.1))) / float64(float32(v18))) / (v20 + v2)) - v20)
                                                            v31 := GenS4{F0: (("a" + "\n\t") + (("" + "a") + "")), F1: 255, F2: (v1 >> ((42) & 63))}
                                                            switch "" {
                                                            case "foo bar":
                                                                if false {
                                                                    v32 := GenS1{F0: (v19 - v19), F1: -1}
                                                                    if true {
                                                                        v4 = 0
                                                                        v33 := GenS1{F0: v19, F1: -1}
                                                                        _ = v33
                                                                    }
                                                                    _ = v32
                                                                } else {
                                                                    var v34 int16 = -371
                                                                    var v35 int = -1
                                                                    var v36 uint64 = v19
                                                                    _ = v34
                                                                    _ = v35
                                                                    _ = v36
                                                                }
                                                                v3 = append(v3, 840)
                                                                _ = len(v3)
                                                            default:
                                                                if (false || false) {
                                                                    var v37 uint = 4294967295
                                                                    var v38 uint32 = 4294967295
                                                                    if len(v3) > 0 {
                                                                        v3[0] = 32768
                                                                    }
                                                                    _ = v37
                                                                    _ = v38
                                                                }
                                                            }
                                                            v3 = append(v3, fi21)
                                                            _ = v31
                                                        }
                                                    case "":
                                                        _ = len(v3)
                                                    }
                                                    var v39 string = (("hello" + "") + (("a" + "\n\t") + "\n\t"))
                                                    _ = len(v3)
                                                    _ = v29
                                                    _ = v30
                                                    _ = v39
                                                }
                                            }
                                        } else {
                                            if (!false) {
                                                var v40 uint64 = v19
                                                switch 9223372036854775807 {
                                                case 127:
                                                    var v41 uint16 = 255
                                                    _ = v41
                                                case -331:
                                                    v1 -= (v1 * (v1 * -1))
                                                }
                                                switch 0 {
                                                case 0:
                                                    v19 += (^(v40))
                                                case 1:
                                                    if false {
                                                        var v42 int8 = (v24 ^ v24)
                                                        v43 := GenS2{F0: 807, F1: 42}
                                                        v2 -= float64(-2.25)
                                                        _ = v42
                                                        _ = v43
                                                    } else {
                                                        var v44 float32 = 3.14159
                                                        fmt.Println(v2, fi23)
                                                        _ = v44
                                                    }
                                                }
                                                switch (-(fi21)) {
                                                case 32768:
                                                    if (v1 < int32(v24)) {
                                                        v24 -= v24
                                                        v45 := GenS5{F0: v4}
                                                        var v46 uint64 = 2
                                                        _ = v45
                                                        _ = v46
                                                    }
                                                case 32767:
                                                    v47 := GenS5{F0: v4}
                                                    _ = v47
                                                default:
                                                    v20 = 0.1
                                                }
                                                _ = v40
                                            }
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            v18 -= 0.0
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            if false {
                                                if (1 == 65535) {
                                                    v48 := GenS1{F0: (v19 - (v19 - (v19 << ((425) & 63)))), F1: 9223372036854775807}
                                                    v1 -= v1
                                                    v49 := GenS2{F0: 1, F1: 1}
                                                    _ = v48
                                                    _ = v49
                                                } else {
                                                    v50 := GenS2{F0: 2147483647, F1: 9223372036854775807}
                                                    var v51 uint16 = uint16(v4)
                                                    _ = v50
                                                    _ = v51
                                                }
                                                v19 -= v19
                                                v2 = 1e10
                                            }
                                        }
                                        _ = v24
                                    } else {
                                        fmt.Println(fi21, fi21, v2)
                                    }
                                    var v52 bool = (65535 > uint16(v4))
                                    _ = fi23
                                    _ = v52
                                }
                            } else {
                                fmt.Println(v18, v18, fi21)
                            }
                        }
                    case "hello":
                        switch fi21 {
                        case -1:
                            v53 := GenS4{F0: "foo bar", F1: 32767, F2: (v1 | v1)}
                            _ = v53
                        case 42:
                            switch "a" {
                            case "":
                                if (1 >= 255) {
                                    v19 = (v19 & 18446744073709551615)
                                    if true {
                                        _ = len(v3)
                                        v18 = v20
                                        v4 = (^(v4))
                                    } else {
                                        _ = len(v3)
                                    }
                                }
                            case "a":
                                if false {
                                    v19 = v19
                                } else {
                                    var v54 uint32 = 398
                                    v20 -= 0.0
                                    _ = v54
                                }
                            case "foo bar":
                                switch 128 {
                                case 128:
                                    v20 += float64(1e10)
                                case 2147483647:
                                    switch fi21 {
                                    case 2147483647:
                                        if false {
                                            for fi55 := 0; fi55 < 5; fi55++ {
                                                switch (((("foo bar" + ("foo bar" + "hello")) + "hello") + "hello") + "foo bar") {
                                                case "hello":
                                                    v56 := GenS1{F0: (^(v19)), F1: 629}
                                                    _ = v56
                                                case "":
                                                    var v57 uint = 430
                                                    _ = v57
                                                }
                                                v1 += -2147483648
                                                v20 = (v18 - (float64(float32(v18)) * (v20 / -2.25)))
                                                _ = fi55
                                            }
                                            v58 := GenS6{F0: v19, F1: 32767, F2: (v4 + (v4 - 42))}
                                            v1 = v1
                                            _ = v58
                                        } else {
                                            v19 = 661
                                            if (1 == 86) {
                                                var v59 int = fi21
                                                var v60 int8 = 125
                                                _ = v59
                                                _ = v60
                                            } else {
                                                var v61 uint8 = v4
                                                _ = v61
                                            }
                                        }
                                    case 596:
                                        for fi62 := 0; fi62 < 3; fi62++ {
                                            fmt.Println(v18, v18)
                                            v63 := GenS7{F0: (v19 - (v19 * 4294967295)), F1: (v19 ^ 4294967295)}
                                            var v64 bool = false
                                            v4 -= (v4 & v4)
                                            v65 := GenS8{F0: 9223372036854775807, F1: genFn1(-1), F2: -235}
                                            _ = fi62
                                            _ = v63
                                            _ = v64
                                            _ = v65
                                        }
                                    default:
                                        v66 := GenS5{F0: (v4 << ((18446744073709551615) & 63))}
                                        _ = v66
                                    }
                                case 42:
                                    v2 -= v18
                                }
                            }
                        default:
                            var v67 int32 = (-(v1))
                            _ = v67
                        }
                    default:
                        v68 := GenS5{F0: genFn0()}
                        _ = v68
                    }
                }
                v20 += -2.25
                _ = fi21
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v20
        }
    case "hello":
        v18 = v2
    case "":
        v18 = v18
    default:
        _ = len(v3)
    }
    for fi69 := 0; fi69 < 20; fi69++ {
        if true {
            if (v18 != 3.14159) {
                switch "a" {
                case "":
                    v4 = uint8(v19)
                case "hello":
                    var v70 int16 = -1
                    _ = v70
                case "a":
                    if ((v4 << ((65535) & 63)) != (^(v4))) {
                        if (418 < uint32(v4)) {
                            for fi71 := 0; fi71 < 13; fi71++ {
                                var v72 uint16 = 32767
                                fmt.Println(v72, v72, v2)
                                _ = fi71
                                _ = v72
                            }
                            if true {
                                var v73 int32 = v1
                                _ = v73
                            }
                        }
                        v19 = v19
                        if ((0 != 4294967295) && (v1 >= (-(v1)))) {
                            v19 -= (v19 & (v19 * v19))
                            for fi74 := 0; fi74 < 18; fi74++ {
                                var v75 uint8 = (^(v4))
                                v1 = (v1 ^ v1)
                                v19 -= (v19 * v19)
                                fmt.Println(v1, v18, v4)
                                _ = fi74
                                _ = v75
                            }
                        }
                    }
                }
                var v76 float64 = -2.25
                for fi77 := 0; fi77 < 18; fi77++ {
                    v19 += 415
                    if false {
                        if false {
                            var v78 uint = 152
                            _ = v78
                        }
                        switch "a" {
                        case "\n\t":
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        case "a":
                            var v79 bool = true
                            _ = v79
                        default:
                            if len(v3) > 0 {
                                v3[0] = (fi77 | (fi77 / ((-2147483648) | 1)))
                            }
                        }
                        if (42 <= genFn1(1)) {
                            var v80 int16 = 1
                            v18 = -2.25
                            v1 = v1
                            _ = v80
                        }
                        fmt.Println(fi69, v4, v1)
                    } else {
                        v81 := GenS9{F0: v4, F1: "foo bar"}
                        switch fi77 {
                        case 25:
                            fmt.Println(v19, v18)
                        }
                        _ = v81
                    }
                    switch -80 {
                    case -1:
                        var v82 uint64 = 42
                        _ = v82
                    case 128:
                        var v83 int16 = 32767
                        _ = v83
                    default:
                        v19 = 4294967295
                    }
                    var v84 float64 = float64(1e10)
                    _ = fi77
                    _ = v84
                }
                _ = v76
            }
            v18 = (v18 / v2)
            var v85 uint16 = 889
            _ = v85
        }
        _ = fi69
    }
    var v86 int32 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v86
}
