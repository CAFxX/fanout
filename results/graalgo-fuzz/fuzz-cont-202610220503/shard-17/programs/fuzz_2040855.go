package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8; F1 uint8; F2 uint16 }
type GenS4 struct { F0 uint8; F1 int64 }
type GenS5 struct { F0 int16; F1 uint64; F2 int8 }
type GenS6 struct { F0 uint8; F1 string; F2 uint32 }
type GenS7 struct { F0 int; F1 uint }
func genFn0(p0 bool, p1 bool, p2 string) int8 {
    if ((p0 && (p0 && false)) && p0) {
        var v4 float32 = 0.0
        var v5 uint32 = 1
        for fi6 := 0; fi6 < 8; fi6++ {
            v7 := []int{127, (fi6 | fi6), (fi6 - fi6)}
            switch ("" + p2) {
            case "foo bar":
                var v8 int = fi6
                _ = v8
            case "hello":
                var v9 string = "a"
                _ = v9
            }
            _ = fi6
            _ = v7
        }
        _ = v4
        _ = v5
    }
    var v10 uint16 = 1
    switch (((p2 + "") + "\n\t") + "\n\t") {
    case "foo bar":
        v11 := GenS1{F0: 2147483648, F1: 776}
        _ = v11
    case "hello":
        for fi12 := 0; fi12 < 4; fi12++ {
            for fi13 := 0; fi13 < 8; fi13++ {
                switch (^(fi13)) {
                case -242:
                    v14 := GenS2{F0: 1}
                    _ = v14
                default:
                    var v15 int32 = -662
                    switch p2 {
                    case "hello":
                        v16 := []int{-9223372036854775808, (fi12 ^ 32768), 2147483647}
                        _ = v16
                    case "":
                        if (!true) {
                            v15 -= -313
                            p1 = (p1 && (!p1))
                            v17 := GenS2{F0: 332}
                            _ = v17
                        } else {
                            v18 := []int{(fi13 & (-(fi12)))}
                            v19 := GenS3{F0: 127, F1: 13, F2: (v10 - 0)}
                            if (1 >= 101) {
                                var v20 uint8 = 203
                                _ = v20
                            } else {
                                switch "hello" {
                                case "":
                                    fmt.Println(p1)
                                    if len(v18) >= 2 {
                                        v18 = v18[1:len(v18)]
                                    }
                                default:
                                    if p0 {
                                        fmt.Println(v10, p1)
                                    } else {
                                        p0 = (!(990 <= 42))
                                        if len(v18) >= 2 {
                                            v18 = v18[1:len(v18)]
                                        }
                                        v21 := GenS4{F0: 240, F1: -628}
                                        _ = v21
                                    }
                                }
                                if (65535 <= 2) {
                                    switch (fi12 >> ((4294967295) & 63)) {
                                    case 680:
                                        switch -9223372036854775808 {
                                        case -912:
                                            var v22 float64 = float64(3.14159)
                                            var v23 uint32 = 0
                                            if len(v18) > 0 {
                                                v18[0] = (fi13 * (fi13 + fi12))
                                            }
                                            _ = v22
                                            _ = v23
                                        case -2147483648:
                                            switch fi12 {
                                            case -1:
                                                if len(v18) >= 2 {
                                                    v18 = v18[1:len(v18)]
                                                }
                                            case 4294967295:
                                                if p1 {
                                                    p2 = p2
                                                    v24 := GenS5{F0: 1, F1: 14, F2: 100}
                                                    _ = v24
                                                }
                                            default:
                                                if p1 {
                                                    v25 := GenS2{F0: 113}
                                                    _ = v25
                                                } else {
                                                    if (!p1) {
                                                        switch (fi13 % ((fi12) | 1)) {
                                                        case -1:
                                                            p2 = p2
                                                        }
                                                        if ("foo bar" != ("hello" + "\n\t")) {
                                                            p2 = (p2 + p2)
                                                            var v26 uint64 = 18446744073709551615
                                                            v15 = v15
                                                            _ = v26
                                                        } else {
                                                            switch 128 {
                                                            case 128:
                                                                var v27 float64 = 1e10
                                                                _ = v27
                                                            case 9223372036854775807:
                                                                v18 = append(v18, (-(fi13)))
                                                            case 551:
                                                                v28 := GenS6{F0: 0, F1: p2, F2: 4294967295}
                                                                _ = v28
                                                            }
                                                            p1 = (!(42 >= 4294967295))
                                                        }
                                                        fmt.Println(fi13, p1, v15)
                                                    }
                                                    p0 = (!true)
                                                    fmt.Println(v15, v10, fi13)
                                                }
                                                fmt.Println(v10, p0, fi12)
                                            }
                                        case 525:
                                            v15 += -2147483648
                                            v29 := GenS7{F0: (fi12 + 128), F1: 0}
                                            _ = v29
                                        default:
                                            v15 = v15
                                        }
                                    default:
                                        var v30 string = "\n\t"
                                        fmt.Println(v15, p0, v30)
                                        p0 = false
                                        _ = v30
                                    }
                                    var v31 uint16 = (v10 | 255)
                                    if (3.14159 <= 0.1) {
                                        fmt.Println(fi12)
                                        var v32 bool = (580 > 156)
                                        _ = v32
                                    }
                                    p1 = (42 > 4294967295)
                                    var v33 uint64 = 1
                                    fmt.Println(fi13, fi12)
                                    _ = v31
                                    _ = v33
                                }
                            }
                            _ = v18
                            _ = v19
                        }
                    }
                    fmt.Println(fi12, v10, p2)
                    _ = v15
                }
                v10 = (v10 << ((65535) & 63))
                var v34 float64 = float64(0.0)
                _ = fi13
                _ = v34
            }
            _ = fi12
        }
    case "":
        fmt.Println(v10, p0)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 100
}
func genFn1(p0 uint32) int64 {
    if true {
        v35 := []int{}
        _ = v35
    } else {
        switch -2147483648 {
        case -2147483648:
            for fi36 := 0; fi36 < 15; fi36++ {
                var v37 int16 = 0
                _ = fi36
                _ = v37
            }
        case 4294967295:
            var v38 int64 = 0
            _ = v38
        }
        p0 -= p0
    }
    var v39 float64 = float64(float32(float64(1e10)))
    v39 = v39
    _ = p0
    _ = v39
    return 9223372036854775807
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int64 = -9223372036854775808
    v3 := make(map[string]string)
    switch "" {
    case "a":
        var v40 int = 0
        _ = v40
    case "foo bar":
        if true {
            var v41 int16 = 1000
            for fi42 := 0; fi42 < 2; fi42++ {
                fmt.Println(fi42, v2, fi42)
                var v43 string = "hello"
                v2 += (v2 << ((255) & 63))
                _ = fi42
                _ = v43
            }
            _ = v41
        }
    default:
        if false {
            v2 -= v2
            switch 952 {
            case 9223372036854775807:
                v2 -= v2
            default:
                _ = len(v1)
            }
            var v44 int32 = 42
            _ = v44
        }
    }
    for fi45 := 0; fi45 < 3; fi45++ {
        v46 := []int{fi45, (fi45 * (fi45 + (^(fi45)))), (fi45 + -2147483648), 597}
        _ = fi45
        _ = v46
    }
    for fi47 := 0; fi47 < 20; fi47++ {
        v48 := []int{-990, int(v2)}
        _ = fi47
        _ = v48
    }
    switch ("foo bar" + "foo bar") {
    case "foo bar":
        v49 := v3[((("foo bar" + "hello") + ("" + "\n\t")) + "foo bar")]
        _ = v49
    case "":
        fmt.Println(v2)
    case "a":
        v50 := []int{32768, 339}
        _ = v50
    default:
        v2 += v2
    }
    fmt.Println(v2)
    var v51 int8 = 127
    if (0 < 66) {
        v52 := []int{597, 32767, -600, 290}
        _ = v52
    } else {
        v53 := []int{}
        var v54 float32 = 1.5
        _ = v53
        _ = v54
    }
    fmt.Println(v2, v2)
    v55 := GenS2{F0: 2147483648}
    fmt.Println(v51, v51)
    if ("\n\t" == "") {
        v56 := GenS7{F0: -776, F1: 0}
        v2 = (v2 + v2)
        _ = v56
    }
    for fi57 := 0; fi57 < 7; fi57++ {
        switch "" {
        case "hello":
            if true {
                v58 := []int{fi57, (fi57 ^ (-(fi57)))}
                fmt.Println(v51, fi57)
                _ = v58
            } else {
                v2 -= (v2 << ((970) & 63))
                v59 := []int{(-(fi57)), -818, (fi57 & (-(fi57)))}
                for fi60 := 0; fi60 < 11; fi60++ {
                    if ((-512 == 244) && true) {
                        v2 = (-(v2))
                        v51 -= (v51 | genFn0(true, true, "a"))
                        v2 = v2
                        if ((881 >= 1) && true) {
                            v51 += v51
                        } else {
                            v61 := GenS7{F0: fi57, F1: 65535}
                            _ = v61
                        }
                        if true {
                            v51 = genFn0(true, (v2 > genFn1(0)), "")
                        }
                        v62 := GenS6{F0: 0, F1: ("foo bar" + "hello"), F2: 2147483648}
                        fmt.Println(v2, fi57)
                        _ = v62
                    }
                    var v63 bool = ((-(fi60)) <= (fi57 * 128))
                    if false {
                        fmt.Println(fi60, v63)
                        var v64 bool = v63
                        var v65 uint8 = 102
                        _ = v64
                        _ = v65
                    }
                    v51 += v51
                    _ = fi60
                    _ = v63
                }
                _ = v59
            }
        case "\n\t":
            if (0.0 != 1.5) {
                var v66 int8 = (v51 % (((v51 - v51)) | 1))
                v51 -= genFn0((false || false), false, "\n\t")
                for fi67 := 0; fi67 < 6; fi67++ {
                    var v68 int32 = 1
                    _ = fi67
                    _ = v68
                }
                _ = v66
            }
        }
        _ = len(v3)
        _ = fi57
    }
    fmt.Println(v2, v51)
    v69 := GenS6{F0: 255, F1: "hello", F2: 42}
    var v70 bool = (!(42 >= 0))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v51:", v51)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v3
    _ = v51
    _ = v55
    _ = v69
    _ = v70
}
