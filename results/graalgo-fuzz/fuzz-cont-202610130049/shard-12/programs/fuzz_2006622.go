package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int64; F1 uint64 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int64; F1 uint; F2 string }
type GenS7 struct { F0 uint16; F1 uint32 }
func genFn0(p0 uint64, p1 int32, p2 uint32) bool {
    p0 = (p0 - (p0 ^ p0))
    var v5 int16 = 0
    var v6 int = 42
    if false {
        fmt.Println(v6)
        v6 -= (v6 * v6)
        var v7 uint16 = 0
        _ = v7
    } else {
        v8 := GenS1{F0: 2}
        v9 := GenS2{F0: (uint(p2) == 1), F1: 127}
        _ = v8
        _ = v9
    }
    var v10 int8 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v5
    _ = v6
    return false
}
func genFn1(p0 int, p1 int, p2 uint) uint {
    for fi11 := 0; fi11 < 6; fi11++ {
        fmt.Println(fi11, p2)
        if false {
            fmt.Println(fi11, p1)
            fmt.Println(p2, p2)
        } else {
            v12 := GenS1{F0: 123}
            var v13 uint8 = 226
            if ((!false) && true) {
                var v14 uint16 = 255
                if false {
                    p0 -= p0
                    v14 -= uint16(v13)
                    v14 += v14
                } else {
                    if (int64(p1) > -195) {
                        v15 := GenS1{F0: v13}
                        v16 := GenS1{F0: v13}
                        p2 -= p2
                        _ = v15
                        _ = v16
                    }
                }
                _ = v14
            }
            _ = v12
            _ = v13
        }
        _ = fi11
    }
    fmt.Println(p0, p2, p2)
    p2 -= 944
    p0 = p1
    p0 = p1
    _ = p0
    _ = p1
    _ = p2
    return 739
}

func main() {
    var v1 uint = 1
    v2 := []string{}
    v3 := []int{}
    v4 := make(map[uint32]string)
    fmt.Println(v1)
    v17 := GenS2{F0: false, F1: 2}
    v1 += v1
    var v18 int16 = -396
    switch (("" + "") + "") {
    case "a":
        if (!false) {
            switch ("\n\t" + "foo bar") {
            case "":
                var v19 float32 = 0.0
                _ = v19
            default:
                v4[2147483647] = ("foo bar" + "")
            }
            for fi20 := 0; fi20 < 16; fi20++ {
                var v21 uint8 = 42
                _ = fi20
                _ = v21
            }
            v1 = v1
        } else {
            if false {
                _ = len(v4)
                fmt.Println(v18, v1)
            }
            var v22 int8 = -90
            _ = v22
        }
    case "\n\t":
        v23 := GenS3{F0: 127}
        _ = v23
    case "foo bar":
        v24 := GenS4{F0: -126, F1: 2}
        _ = v24
    default:
        switch "foo bar" {
        case "foo bar":
            switch "foo bar" {
            case "":
                _ = len(v2)
            case "hello":
                _ = len(v4)
            case "\n\t":
                var v25 int8 = 100
                _ = v25
            default:
                for fi26 := 0; fi26 < 3; fi26++ {
                    v18 += v18
                    for fi27 := 0; fi27 < 2; fi27++ {
                        var v28 int32 = 1
                        _ = fi27
                        _ = v28
                    }
                    if (1.5 > 3.14159) {
                        var v29 uint8 = 42
                        for fi30 := 0; fi30 < 4; fi30++ {
                            v29 = v29
                            v31 := GenS5{F0: 71}
                            _ = fi30
                            _ = v31
                        }
                        _ = v29
                    }
                    _ = fi26
                }
            }
        case "hello":
            fmt.Println(v18, v18)
        case "a":
            var v32 uint32 = 42
            _ = v32
        }
    }
    v1 = (v1 % ((genFn1(32767, 4294967295, v1)) | 1))
    switch 32767 {
    case 127:
        fmt.Println(v18, v18)
    case -564:
        fmt.Println(v1, v18)
    case 42:
        v33 := GenS5{F0: -1}
        _ = v33
    default:
        if false {
            switch 1 {
            case -671:
                switch ("hello" + "\n\t") {
                case "\n\t":
                    v18 -= v18
                case "foo bar":
                    if (-40 <= -128) {
                        v18 += 0
                        v1 += v1
                        v34 := GenS1{F0: 249}
                        _ = v34
                    }
                case "a":
                    for fi35 := 0; fi35 < 4; fi35++ {
                        fmt.Println(v18, fi35, v18)
                        var v36 float64 = 1e10
                        _ = fi35
                        _ = v36
                    }
                default:
                    _ = len(v4)
                }
            default:
                for fi37 := 0; fi37 < 7; fi37++ {
                    for fi38 := 0; fi38 < 18; fi38++ {
                        switch "\n\t" {
                        case "a":
                            v18 -= 32767
                        case "\n\t":
                            v1 -= v1
                        case "hello":
                            v39 := GenS6{F0: 990, F1: v1, F2: "a"}
                            _ = v39
                        default:
                            var v40 int64 = 9223372036854775807
                            _ = v40
                        }
                        fmt.Println(v18, v1)
                        switch fi38 {
                        case 871:
                            _ = len(v4)
                        case 9223372036854775807:
                            v4[42] = "a"
                        default:
                            if ((fi37 - (fi38 / ((791) | 1))) != (fi38 << (((v1 >> (((v1 * 233)) & 63))) & 63))) {
                                var v41 bool = true
                                v18 = v18
                                switch ("foo bar" + ("foo bar" + "hello")) {
                                case "a":
                                    var v42 string = ""
                                    _ = v42
                                case "hello":
                                    v2 = append(v2, "a")
                                case "\n\t":
                                    v41 = (2 != 128)
                                }
                                var v43 float32 = 1e10
                                if true {
                                    var v44 float64 = float64(v43)
                                    v18 += (^(v18))
                                    if (65535 > 2147483647) {
                                        v44 = (3.14159 + ((v44 * 0.1) * 1.5))
                                        if ((fi38 % ((32768) | 1)) != int(v1)) {
                                            v18 -= 32767
                                            v45 := GenS1{F0: 155}
                                            v1 = v1
                                            switch ("a" + "foo bar") {
                                            case "":
                                                if ((0.1 != v44) && false) {
                                                    fmt.Println(fi38)
                                                } else {
                                                    v46 := GenS3{F0: 128}
                                                    var v47 bool = true
                                                    var v48 int = 2147483647
                                                    _ = v46
                                                    _ = v47
                                                    _ = v48
                                                }
                                            case "a":
                                                var v49 float32 = ((v43 * v43) / ((-2.25 / v43) * v43))
                                                _ = v49
                                            default:
                                                fmt.Println(v43, v44)
                                            }
                                            v1 = 42
                                            switch "\n\t" {
                                            case "\n\t":
                                                v1 += v1
                                            case "a":
                                                v50 := GenS3{F0: 42}
                                                _ = v50
                                            case "":
                                                v18 += v18
                                            default:
                                                v1 += (v1 ^ (v1 + (v1 << (((^(v1))) & 63))))
                                            }
                                            _ = v45
                                        } else {
                                            v4[2] = "hello"
                                            fmt.Println(v44, v18, v1)
                                        }
                                        fmt.Println(fi38, fi37)
                                    }
                                    if false {
                                        var v51 uint32 = 2147483648
                                        _ = v51
                                    }
                                    var v52 uint8 = 1
                                    _ = v44
                                    _ = v52
                                } else {
                                    if v41 {
                                        v53 := GenS1{F0: 1}
                                        v18 += -215
                                        if len(v2) > 0 {
                                            v2[0] = "a"
                                        }
                                        var v54 int = 850
                                        if (127 != 100) {
                                            v55 := GenS6{F0: 734, F1: 713, F2: "foo bar"}
                                            v43 = v43
                                            fmt.Println(fi37, fi37)
                                            _ = v55
                                        }
                                        _ = v53
                                        _ = v54
                                    } else {
                                        if v41 {
                                            if false {
                                                switch "\n\t" {
                                                case "hello":
                                                    v4[4294967295] = ""
                                                case "":
                                                    v41 = v41
                                                case "\n\t":
                                                    v43 += (v43 / v43)
                                                }
                                                if (472 < (fi37 % (((fi38 - fi37)) | 1))) {
                                                    if v41 {
                                                        v56 := GenS3{F0: 1}
                                                        var v57 bool = false
                                                        _ = v56
                                                        _ = v57
                                                    } else {
                                                        v41 = false
                                                    }
                                                    if (!v41) {
                                                        if ((v18 / (((v18 | v18)) | 1)) != (v18 / (((v18 / ((v18) | 1))) | 1))) {
                                                            switch (^(fi37)) {
                                                            case 853:
                                                                v4[436] = ("hello" + "")
                                                            case 32767:
                                                                fmt.Println(v1, v18, v18)
                                                            default:
                                                                delete(v4, 306)
                                                            }
                                                            fmt.Println(v18, fi37)
                                                        } else {
                                                            fmt.Println(v1, v41)
                                                            if true {
                                                                if v41 {
                                                                    var v58 string = "a"
                                                                    v58 = ("a" + (v58 + v58))
                                                                    if (false && true) {
                                                                        v3 = append(v3, fi37)
                                                                    }
                                                                    _ = v58
                                                                } else {
                                                                    var v59 uint16 = 42
                                                                    _ = v59
                                                                }
                                                                if (((!(((!((fi37 + fi37) <= -2147483648)) && (!(-1 == -910))) || (((((v43 - 0.0) / 1.5) * (3.14159 * (v43 / (v43 + v43)))) * 0.1) != float32(float64(float32(float64(v43))))))) || (1 < 127)) || v41) {
                                                                    fmt.Println(v43)
                                                                }
                                                                v43 += (v43 + float32(0.1))
                                                            }
                                                            if ((((true && v41) && (65535 < 65535)) || v41) || (!((("" + ("hello" + "")) + (("a" + "\n\t") + (("\n\t" + ("" + "\n\t")) + ("" + "foo bar")))) <= (((("\n\t" + "a") + "foo bar") + ("foo bar" + "")) + ((("a" + "a") + "hello") + ("foo bar" + "a")))))) {
                                                                v60 := v4[4294967295]
                                                                v60 = v60
                                                                switch 9223372036854775807 {
                                                                case 32767:
                                                                    v18 += (-(v18))
                                                                case -9223372036854775808:
                                                                    var v61 uint8 = 0
                                                                    _ = v61
                                                                case 65:
                                                                    var v62 uint32 = 2147483648
                                                                    _ = v62
                                                                default:
                                                                    v63 := v4[uint32(v1)]
                                                                    _ = v63
                                                                }
                                                                v64 := GenS6{F0: 1, F1: (v1 ^ 65535), F2: v60}
                                                                _ = v60
                                                                _ = v64
                                                            }
                                                        }
                                                    }
                                                } else {
                                                    switch ("foo bar" + "hello") {
                                                    case "a":
                                                        v43 += -2.25
                                                    default:
                                                        fmt.Println(fi37, v41)
                                                        fmt.Println(v18, v18)
                                                    }
                                                }
                                            }
                                        }
                                        v41 = (32767 < 42)
                                        v65 := GenS4{F0: 0, F1: 2}
                                        _ = v65
                                    }
                                    v66 := GenS5{F0: 42}
                                    _ = v66
                                }
                                _ = v41
                                _ = v43
                            } else {
                                delete(v4, 2147483648)
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                var v67 uint8 = 237
                                _ = v67
                            }
                        }
                        _ = fi38
                    }
                    _ = fi37
                }
            }
        }
    }
    for fi68 := 0; fi68 < 11; fi68++ {
        for fi69 := 0; fi69 < 20; fi69++ {
            v18 += (v18 % (((v18 | (v18 - v18))) | 1))
            _ = len(v3)
            var v70 float64 = ((1.5 + -2.25) - 0.1)
            _ = fi69
            _ = v70
        }
        _ = fi68
    }
    var v71 int = 0
    if true {
        v18 += 32767
    }
    var v72 int32 = -30
    v73 := GenS7{F0: 387, F1: 2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v71:", v71)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v71
    _ = v72
    _ = v73
}
