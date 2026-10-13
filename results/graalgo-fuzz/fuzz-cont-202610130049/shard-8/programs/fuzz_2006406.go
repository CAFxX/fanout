package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int; F1 uint8 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 bool; F1 uint16; F2 bool }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint32; F1 int64; F2 uint }
type GenS9 struct { F0 int64; F1 int8; F2 uint8 }
type GenS10 struct { F0 int16; F1 bool; F2 uint32 }
type GenS11 struct { F0 int16; F1 int32; F2 int32 }
type GenS12 struct { F0 uint8; F1 bool }
func genFn0() bool {
    switch ("" + "\n\t") {
    case "\n\t":
        switch 2147483647 {
        case 1:
            v5 := GenS1{F0: 0, F1: 4294967295}
            _ = v5
        case 616:
            var v6 int32 = -386
            _ = v6
        default:
            var v7 float64 = 0.1
            _ = v7
        }
    case "":
        v8 := GenS2{F0: 2147483647}
        _ = v8
    case "a":
        v9 := []int{}
        _ = v9
    default:
        var v10 int64 = -1
        _ = v10
    }
    if ((!(false && false)) && false) {
        var v11 uint32 = 2147483647
        _ = v11
    }
    var v12 int64 = 0
    _ = v12
    return (2 > 0)
}
func genFn1(p0 string, p1 uint8) uint64 {
    for fi13 := 0; fi13 < 10; fi13++ {
        switch 827 {
        case -9223372036854775808:
            p0 = ("foo bar" + p0)
        default:
            for fi14 := 0; fi14 < 13; fi14++ {
                v15 := []int{(fi13 / ((int(p1)) | 1)), (^(fi14)), fi13, 9223372036854775807}
                _ = len(v15)
                v16 := GenS3{F0: ((("a" + (("\n\t" + ("" + p0)) + p0)) + p0) >= p0)}
                p1 -= (p1 * 128)
                _ = fi14
                _ = v15
                _ = v16
            }
        }
        var v17 string = "\n\t"
        _ = fi13
        _ = v17
    }
    p0 = p0
    var v18 float32 = (1.5 - -2.25)
    var v19 int = 9223372036854775807
    v18 = (0.0 - (v18 + -2.25))
    _ = p0
    _ = p1
    _ = v18
    _ = v19
    return 42
}

func main() {
    var v1 int16 = -1
    v2 := make(map[uint32]uint32)
    var v3 uint32 = 4294967295
    var v4 int = 128
    v20 := GenS2{F0: (v3 >> ((4294967295) & 63))}
    if false {
        v21 := []int{}
        if false {
            v22 := GenS4{F0: (v4 + (v4 / ((v4) | 1))), F1: 255}
            _ = v22
        } else {
            var v23 string = "a"
            v21 = append(v21, v4)
            _ = v23
        }
        _ = v21
    } else {
        v24 := []int{(v4 >> ((2) & 63)), v4, 32768}
        for fi25 := 0; fi25 < 11; fi25++ {
            var v26 int64 = -9223372036854775808
            switch fi25 {
            case 32767:
                switch "foo bar" {
                case "\n\t":
                    v3 -= 0
                case "foo bar":
                    fmt.Println(v3)
                }
            case 9223372036854775807:
                switch fi25 {
                case 2147483647:
                    if ((!true) && ("foo bar" != ("" + ("foo bar" + "a")))) {
                        if (v3 >= 42) {
                            v3 += (v3 ^ (v3 + v3))
                        } else {
                            if true {
                                var v27 int16 = v1
                                v28 := GenS1{F0: (v26 - (-(v26))), F1: 4294967295}
                                _ = v27
                                _ = v28
                            } else {
                                for fi29 := 0; fi29 < 3; fi29++ {
                                    switch fi29 {
                                    case 334:
                                        if ((104 != -14) && (255 <= 64)) {
                                            v30 := GenS5{F0: -468}
                                            v26 = (v26 + v26)
                                            var v31 int16 = v1
                                            _ = v30
                                            _ = v31
                                        }
                                    case 9223372036854775807:
                                        var v32 int32 = -176
                                        _ = v32
                                    case 4294967295:
                                        if true {
                                            v24 = append(v24, 128)
                                            if len(v24) > 0 {
                                                v24[0] = (fi25 % ((v4) | 1))
                                            }
                                        } else {
                                            switch "" {
                                            case "hello":
                                                v3 = (^(v3))
                                                var v33 uint64 = 1
                                                _ = v33
                                            case "foo bar":
                                                var v34 float32 = 1.5
                                                fmt.Println(fi29, v26)
                                                _ = v34
                                            }
                                        }
                                    }
                                    var v35 int = (v4 & (v4 ^ 32767))
                                    _ = fi29
                                    _ = v35
                                }
                                var v36 int = 127
                                _ = v36
                            }
                            v37 := GenS6{F0: false, F1: 0, F2: false}
                            _ = v37
                        }
                        v38 := v2[2147483647]
                        v38 = v38
                        _ = v38
                    } else {
                        v24 = append(v24, (v4 & int(v26)))
                        if true {
                            v39 := GenS7{F0: 0}
                            if len(v24) > 0 {
                                v24[0] = -9223372036854775808
                            }
                            v40 := GenS6{F0: false, F1: 0, F2: (!false)}
                            _ = v39
                            _ = v40
                        }
                        v41 := GenS6{F0: false, F1: 42, F2: (genFn1("a", 127) >= 144)}
                        _ = v41
                    }
                case 488:
                    var v42 uint8 = 255
                    _ = v42
                default:
                    var v43 bool = true
                    _ = v43
                }
            }
            _ = fi25
            _ = v26
        }
        v24 = append(v24, (v4 + v4))
        _ = v24
    }
    v1 = v1
    switch ("" + (("foo bar" + "foo bar") + "")) {
    case "hello":
        v1 = (v1 / ((v1) | 1))
    case "foo bar":
        fmt.Println(v4, v1)
    default:
        if (42 == 4294967295) {
            fmt.Println(v3, v1, v1)
            if true {
                fmt.Println(v1, v4, v1)
                var v44 int8 = 100
                _ = v44
            } else {
                v45 := []int{}
                _ = v45
            }
            if false {
                switch "\n\t" {
                case "foo bar":
                    v1 = 1
                case "hello":
                    var v46 int64 = 657
                    _ = v46
                }
                var v47 bool = (true && false)
                switch v4 {
                case 1:
                    v48 := GenS8{F0: (^(v3)), F1: 0, F2: 1}
                    _ = v48
                default:
                    v4 += v4
                }
                _ = v47
            }
        } else {
            switch v4 {
            case 1:
                if (!false) {
                    fmt.Println(v3, v4, v1)
                }
            case 2147483647:
                fmt.Println(v3, v4)
            default:
                var v49 uint64 = 0
                _ = v49
            }
            v3 -= 1
            fmt.Println(v4)
        }
    }
    switch ("foo bar" + "foo bar") {
    case "\n\t":
        v50 := GenS9{F0: -9223372036854775808, F1: 127, F2: 32}
        _ = v50
    case "foo bar":
        var v51 uint = 866
        _ = v51
    case "a":
        v4 += 128
    default:
        var v52 int16 = (v1 * 42)
        _ = v52
    }
    var v53 int8 = -128
    switch "hello" {
    case "\n\t":
        v3 += 804
    }
    if true {
        switch v4 {
        case -944:
            v54 := []int{v4, (v4 / ((v4) | 1)), v4}
            _ = v54
        case 4294967295:
            var v55 float32 = (((1.5 - float32((3.14159 / float64(1.5)))) / 0.1) + -2.25)
            _ = v55
        case 32767:
            if true {
                if false {
                    var v56 float64 = (1.5 + -2.25)
                    switch ("\n\t" + "hello") {
                    case "foo bar":
                        _ = len(v2)
                    default:
                        var v57 string = "a"
                        _ = v57
                    }
                    _ = v56
                }
                v58 := []int{v4, (v4 ^ 9223372036854775807), 127, 9223372036854775807}
                _ = v58
            } else {
                v59 := []int{}
                var v60 string = (("\n\t" + "hello") + (("\n\t" + (("" + ("" + (("foo bar" + "foo bar") + ("foo bar" + "hello")))) + ("\n\t" + "\n\t"))) + ("a" + "foo bar")))
                if false {
                    switch (v4 - -2147483648) {
                    case 2147483647:
                        var v61 string = (v60 + "a")
                        _ = v61
                    case 535:
                        v62 := GenS5{F0: 42}
                        _ = v62
                    }
                } else {
                    if len(v59) > 0 {
                        v59[0] = v4
                    }
                    v63 := GenS7{F0: 42}
                    _ = v63
                }
                _ = v59
                _ = v60
            }
        }
        v64 := GenS10{F0: v1, F1: (!(!true)), F2: (v3 - (v3 - v3))}
        for fi65 := 0; fi65 < 14; fi65++ {
            switch ("" + "\n\t") {
            case "foo bar":
                switch ("" + ("" + ("foo bar" + ("a" + "")))) {
                case "\n\t":
                    for fi66 := 0; fi66 < 6; fi66++ {
                        fmt.Println(v53)
                        v67 := GenS11{F0: (^(v1)), F1: 42, F2: int32(fi66)}
                        v4 = fi66
                        _ = fi66
                        _ = v67
                    }
                }
            }
            if ((42 >= 1) || false) {
                v68 := GenS5{F0: (-(v1))}
                _ = v68
            } else {
                fmt.Println(v3, v1, v3)
                switch int(v53) {
                case -2147483648:
                    switch "hello" {
                    case "foo bar":
                        v53 = 0
                    }
                case -9223372036854775808:
                    v53 = (v53 & (v53 / ((v53) | 1)))
                }
                for fi69 := 0; fi69 < 4; fi69++ {
                    _ = len(v2)
                    fmt.Println(v3, v1)
                    _ = fi69
                }
            }
            _ = fi65
        }
        _ = v64
    }
    if (9223372036854775807 != 533) {
        v70 := []int{v4, v4, v4, 1}
        v53 = 1
        var v71 bool = (v1 != (v1 / (((v1 / (((-(v1))) | 1))) | 1)))
        _ = v70
        _ = v71
    } else {
        var v72 float64 = 3.14159
        v73 := GenS12{F0: 127, F1: true}
        _ = v72
        _ = v73
    }
    if (!(true && false)) {
        v1 += -854
        v4 -= (^(v4))
        fmt.Println(v4, v1)
    }
    fmt.Println(v4, v4)
    var v74 string = "a"
    switch v4 {
    case 32767:
        v75 := GenS12{F0: 0, F1: (4294967295 > 42)}
        _ = v75
    case 40:
        v76 := []int{-986, v4, (v4 % ((508) | 1)), v4}
        _ = v76
    default:
        v77 := []int{32768, (v4 >> ((65535) & 63)), v4, v4}
        _ = v77
    }
    v1 -= int16(v53)
    switch (v4 >> ((2) & 63)) {
    case -1:
        var v78 bool = false
        _ = v78
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v53:", v53)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v53
    _ = v74
}
