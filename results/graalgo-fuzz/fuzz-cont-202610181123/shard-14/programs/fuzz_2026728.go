package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int; F1 uint64; F2 uint32 }
type GenS4 struct { F0 int32; F1 uint64; F2 uint32 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint; F1 bool }
func genFn0(p0 bool, p1 int32) int64 {
    var v5 int32 = p1
    v6 := []int{-814, -325, 1}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return -9223372036854775808
}
func genFn1(p0 bool, p1 int16) int {
    var v7 float64 = 0.0
    var v8 int32 = 808
    for fi9 := 0; fi9 < 19; fi9++ {
        for fi10 := 0; fi10 < 19; fi10++ {
            v7 = v7
            v11 := GenS1{F0: 2}
            var v12 uint = 487
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return -289
}

func main() {
    v1 := []string{("a" + (("\n\t" + "foo bar") + ("a" + "a")))}
    var v2 float64 = ((1e10 - ((((0.0 / 1.0) + -2.25) - (0.1 - ((3.14159 - ((0.0 / (-2.25 * ((float64(3.14159) * -2.25) - (((((1e10 / 3.14159) / (0.0 - 0.1)) / (3.14159 / 0.1)) - (((0.1 + (1e10 + 1e10)) - 3.14159) - -2.25)) - (0.1 * (1e10 * 3.14159)))))) - (1e10 / 1.0))) - (float64(-2.25) * 1e10)))) * -2.25)) * ((0.0 * 1.5) / (float64(0.0) + float64(3.14159))))
    v3 := []string{("a" + "\n\t"), "hello", "foo bar", ("" + "foo bar")}
    var v4 int64 = -566
    v4 = v4
    for fi13 := 0; fi13 < 3; fi13++ {
        v14 := GenS1{F0: 2}
        for fi15 := 0; fi15 < 3; fi15++ {
            fmt.Println(fi15, v2)
            v2 += (1e10 + (v2 - (((((0.0 * 0.1) + (v2 * float64((-2.25 * 3.14159)))) + (3.14159 - (1e10 / (v2 / (v2 + v2))))) + v2) * -2.25)))
            if true {
                var v16 uint = 1
                var v17 int64 = (v4 << (((v16 * v16)) & 63))
                _ = v16
                _ = v17
            } else {
                var v18 bool = true
                v18 = ((v18 || ((-(fi15)) <= (fi13 & genFn1(true, -200)))) && v18)
                _ = v18
            }
            _ = fi15
        }
        for fi19 := 0; fi19 < 20; fi19++ {
            v20 := GenS1{F0: 341}
            var v21 float64 = 0.0
            var v22 int64 = -9223372036854775808
            _ = fi19
            _ = v20
            _ = v21
            _ = v22
        }
        _ = fi13
        _ = v14
    }
    var v23 int8 = -73
    switch 4294967295 {
    case 2147483647:
        var v24 uint = 1
        _ = v24
    }
    var v25 int64 = 0
    if (-1 > 32768) {
        var v26 bool = false
        v27 := GenS2{F0: 42}
        _ = v26
        _ = v27
    } else {
        var v28 int16 = 39
        _ = v28
    }
    _ = len(v1)
    var v29 int64 = genFn0(((1 >= -2147483648) && false), -175)
    for fi30 := 0; fi30 < 14; fi30++ {
        switch (^(fi30)) {
        case 9223372036854775807:
            v29 += v29
        default:
            v4 += v25
        }
        var v31 int32 = 2147483647
        var v32 int64 = (v4 >> ((255) & 63))
        _ = fi30
        _ = v31
        _ = v32
    }
    v33 := GenS1{F0: 259}
    var v34 int8 = v23
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v35 := GenS1{F0: 32767}
    if true {
        switch 194 {
        case 127:
            fmt.Println(v4, v2)
        }
    } else {
        if (1 < -1) {
            for fi36 := 0; fi36 < 17; fi36++ {
                v2 = v2
                fmt.Println(v2, v29)
                _ = fi36
            }
            for fi37 := 0; fi37 < 9; fi37++ {
                var v38 int8 = v23
                v34 = v38
                v25 += (^(v4))
                _ = fi37
                _ = v38
            }
        } else {
            if ((!false) || false) {
                for fi39 := 0; fi39 < 13; fi39++ {
                    v40 := GenS2{F0: 806}
                    _ = fi39
                    _ = v40
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            } else {
                for fi41 := 0; fi41 < 20; fi41++ {
                    v1 = append(v1, "hello")
                    switch (("hello" + (("foo bar" + "a") + "a")) + "hello") {
                    case "":
                        if true {
                            if (v29 >= v29) {
                                var v42 float32 = 0.0
                                _ = v42
                            }
                            for fi43 := 0; fi43 < 6; fi43++ {
                                if (-32768 <= -540) {
                                    fmt.Println(v29, v2, v29)
                                } else {
                                    fmt.Println(v23)
                                    v44 := GenS3{F0: 128, F1: 4294967295, F2: 2147483648}
                                    _ = v44
                                }
                                switch (fi43 / ((fi41) | 1)) {
                                case 0:
                                    if false {
                                        var v45 string = "\n\t"
                                        _ = v45
                                    } else {
                                        v2 = v2
                                        if false {
                                            switch ("foo bar" + (("a" + "a") + "a")) {
                                            case "":
                                                var v46 uint8 = 255
                                                _ = v46
                                            default:
                                                if len(v3) > 0 {
                                                    v3[0] = "hello"
                                                }
                                            }
                                        } else {
                                            var v47 uint64 = 265
                                            var v48 int32 = -2147483648
                                            _ = v47
                                            _ = v48
                                        }
                                    }
                                case 9223372036854775807:
                                    switch (fi41 << ((18446744073709551615) & 63)) {
                                    case 2147483647:
                                        _ = len(v1)
                                    case 32767:
                                        var v49 bool = (-895 <= 990)
                                        _ = v49
                                    }
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                default:
                                    var v50 uint16 = 255
                                    _ = len(v1)
                                    var v51 int16 = -855
                                    _ = v50
                                    _ = v51
                                }
                                _ = fi43
                            }
                            _ = len(v1)
                        }
                    default:
                        var v52 uint16 = 0
                        _ = v52
                    }
                    _ = fi41
                }
            }
        }
        for fi53 := 0; fi53 < 8; fi53++ {
            v25 = v25
            for fi54 := 0; fi54 < 16; fi54++ {
                if (-1 != -1) {
                    v29 -= int64(v23)
                    fmt.Println(v29)
                }
                v29 -= genFn0((4294967295 >= 1), -985)
                var v55 int32 = 1
                _ = fi54
                _ = v55
            }
            _ = fi53
        }
        v34 = (v34 ^ v34)
    }
    if ((-1 != 311) && (!false)) {
        var v56 int32 = -683
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        switch "\n\t" {
        case "a":
            v4 -= (-(v25))
        case "":
            if false {
                v4 -= (v4 / ((genFn0((!(878 >= 218)), v56)) | 1))
                switch 4294967295 {
                case 4294967295:
                    switch 4294967295 {
                    case -2147483648:
                        fmt.Println(v29)
                    case -3:
                        for fi57 := 0; fi57 < 15; fi57++ {
                            if (!false) {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                v58 := GenS4{F0: (v56 / (((-(v56))) | 1)), F1: 18446744073709551615, F2: 1}
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = v58
                            }
                            for fi59 := 0; fi59 < 20; fi59++ {
                                var v60 float32 = 1.5
                                if len(v3) > 0 {
                                    v3[0] = ("\n\t" + "\n\t")
                                }
                                _ = fi59
                                _ = v60
                            }
                            _ = fi57
                        }
                    }
                default:
                    var v61 float32 = 1.5
                    _ = v61
                }
                var v62 bool = (!(((v2 / ((((v2 - (float64(1.5) + ((v2 / v2) / v2))) * (3.14159 * v2)) - float64(3.14159)) / float64(((0.1 / 1.0) + float32(v2))))) * v2) <= v2))
                _ = v62
            }
        case "hello":
            if (((1.5 <= (0.1 - 3.14159)) && (127 == 255)) && true) {
                for fi63 := 0; fi63 < 9; fi63++ {
                    var v64 bool = false
                    v65 := GenS1{F0: 0}
                    _ = fi63
                    _ = v64
                    _ = v65
                }
                for fi66 := 0; fi66 < 6; fi66++ {
                    var v67 int8 = v34
                    v4 += (v25 % ((genFn0((true && ((v23 % ((127) | 1)) == (v34 << ((421) & 63)))), -2147483648)) | 1))
                    switch "a" {
                    case "\n\t":
                        v34 = v23
                    case "hello":
                        switch "hello" {
                        case "\n\t":
                            var v68 uint32 = 790
                            _ = v68
                        case "":
                            _ = len(v1)
                        case "hello":
                            switch ("\n\t" + "a") {
                            case "":
                                v69 := GenS2{F0: 1}
                                _ = v69
                            default:
                                var v70 uint64 = 806
                                _ = v70
                            }
                        }
                    }
                    _ = fi66
                    _ = v67
                }
            } else {
                if len(v1) > 0 {
                    v1[0] = "a"
                }
                switch genFn1(("foo bar" > "\n\t"), -1) {
                case 636:
                    v71 := GenS2{F0: (v56 + (^(v56)))}
                    _ = v71
                default:
                    v72 := GenS2{F0: v56}
                    _ = v72
                }
            }
        default:
            v73 := GenS5{F0: ("a" + "hello")}
            _ = v73
        }
        _ = v56
    }
    if true {
        v74 := GenS6{F0: 0, F1: (32767 != -786)}
        var v75 float64 = ((float64(1e10) - 1e10) - 0.1)
        if len(v3) > 0 {
            v3[0] = "hello"
        }
        _ = v74
        _ = v75
    } else {
        var v76 bool = false
        _ = v76
    }
    switch genFn1(false, 0) {
    case 42:
        var v77 uint32 = 1
        _ = v77
    case 4294967295:
        v4 = 42
    default:
        _ = len(v1)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v25
    _ = v29
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
}
