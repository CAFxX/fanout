package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 uint; F1 uint8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int; F1 uint }
type GenS6 struct { F0 uint16; F1 bool; F2 uint64 }
type GenS7 struct { F0 string }
func genFn0(p0 int32, p1 uint32) int8 {
    fmt.Println(p0, p1)
    fmt.Println(p0, p0)
    var v6 bool = (127 <= 127)
    _ = p0
    _ = p1
    _ = v6
    return 100
}
func genFn1() uint8 {
    var v7 bool = false
    switch "" {
    case "":
        var v8 uint16 = 65535
        _ = v8
    case "a":
        switch 32768 {
        case -1:
            fmt.Println(v7)
        case 9223372036854775807:
            fmt.Println(v7)
        }
    case "hello":
        for fi9 := 0; fi9 < 3; fi9++ {
            v10 := GenS1{F0: 128, F1: 42}
            var v11 float32 = (1.5 / (0.1 + 0.0))
            var v12 float32 = (((v11 * float32((3.14159 * (1e10 / 1.0)))) / v11) / v11)
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
    }
    var v13 bool = v7
    if ((18446744073709551615 < 0) || (65535 == 42)) {
        switch 32768 {
        case 416:
            v14 := []int{1, -2147483648, 9223372036854775807, 0}
            _ = v14
        default:
            fmt.Println(v13, v7)
        }
        v13 = v7
    }
    _ = v13
    _ = v7
    return 179
}

func main() {
    var v1 float32 = (1e10 - (-2.25 / ((3.14159 * 0.1) * 0.1)))
    var v2 int = 306
    var v3 int = (v2 ^ v2)
    var v4 bool = (false && (v2 == (v2 + v2)))
    v5 := []string{""}
    v15 := GenS1{F0: 52, F1: 62}
    _ = len(v5)
    v16 := GenS1{F0: 117, F1: 42}
    for fi17 := 0; fi17 < 8; fi17++ {
        if (883 < 978) {
            switch fi17 {
            case 4294967295:
                v2 -= (v3 / ((v3) | 1))
            case 9223372036854775807:
                v2 -= fi17
            case 42:
                v3 = -2147483648
            default:
                v1 -= ((v1 - v1) / ((((0.0 + 1.5) + (1.5 / 1e10)) / (1e10 / float32((-2.25 * (3.14159 + 3.14159))))) / (v1 * (v1 + v1))))
            }
            for fi18 := 0; fi18 < 1; fi18++ {
                v19 := GenS2{F0: 42, F1: 1}
                v1 += v1
                _ = fi18
                _ = v19
            }
        }
        fmt.Println(v1)
        for fi20 := 0; fi20 < 10; fi20++ {
            switch "hello" {
            case "foo bar":
                v5 = append(v5, ("" + "foo bar"))
            case "\n\t":
                if false {
                    v3 = (v2 + fi20)
                } else {
                    var v21 bool = (127 > -109)
                    if (!v4) {
                        if len(v5) > 0 {
                            v5[0] = "\n\t"
                        }
                        var v22 float32 = 1e10
                        _ = v22
                    } else {
                        v5 = append(v5, (((("foo bar" + "") + ("hello" + "a")) + "a") + ("\n\t" + "hello")))
                        var v23 string = ""
                        _ = v23
                    }
                    _ = v21
                }
            default:
                v4 = (42 != 127)
            }
            if (1.5 == 3.14159) {
                var v24 int = (v3 << ((18446744073709551615) & 63))
                switch fi20 {
                case -2147483648:
                    if len(v5) > 0 {
                        v5[0] = "a"
                    }
                default:
                    var v25 uint32 = 1
                    _ = v25
                }
                if ((90 >= 0) || v4) {
                    v26 := GenS1{F0: 1, F1: 685}
                    v3 = v2
                    _ = v26
                } else {
                    switch v2 {
                    case -770:
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    case 635:
                        v4 = (!(42 <= 100))
                    case 32768:
                        v4 = (((v4 || false) && (!false)) || (!(-128 <= 1)))
                    default:
                        if v4 {
                            v27 := GenS1{F0: 255, F1: 898}
                            var v28 int32 = 2147483647
                            _ = v27
                            _ = v28
                        } else {
                            if (true && v4) {
                                v29 := GenS3{F0: 65535, F1: 84}
                                _ = len(v5)
                                v30 := GenS2{F0: 182, F1: 2}
                                v5 = append(v5, "a")
                                v2 += fi20
                                _ = v29
                                _ = v30
                            }
                            v2 = (fi17 | -305)
                            v24 += (v3 - (v24 / ((fi17) | 1)))
                        }
                        switch "foo bar" {
                        case "foo bar":
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        }
                        switch (-(v3)) {
                        case -452:
                            _ = len(v5)
                        case 1:
                            fmt.Println(fi17, v1, v2)
                        case -364:
                            if (3.14159 >= (0.0 / 1.5)) {
                                if len(v5) > 0 {
                                    v5[0] = "\n\t"
                                }
                                fmt.Println(fi20)
                            } else {
                                v4 = v4
                                v31 := GenS4{F0: 4294967295}
                                if (!true) {
                                    v3 += v2
                                    var v32 int16 = -32768
                                    _ = v32
                                } else {
                                    v24 += (v2 >> ((1) & 63))
                                    if ((float64(v1) - 1e10) > -2.25) {
                                        fmt.Println(v4)
                                        v5 = append(v5, "foo bar")
                                    } else {
                                        fmt.Println(v3)
                                    }
                                }
                                v5 = append(v5, ("" + "foo bar"))
                                _ = v31
                            }
                        }
                    }
                    v5 = append(v5, "")
                    var v33 bool = true
                    _ = v33
                }
                _ = v24
            }
            if len(v5) > 0 {
                v5[0] = ""
            }
            _ = fi20
        }
        _ = fi17
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println(v4)
    if ((2 <= 170) || (!((!(!false)) && ("\n\t" >= (("foo bar" + ("foo bar" + "foo bar")) + (("\n\t" + "") + ("hello" + ""))))))) {
        switch ("" + ("hello" + "")) {
        case "\n\t":
            if true {
                if len(v5) > 0 {
                    v5[0] = "foo bar"
                }
            } else {
                switch "a" {
                case "foo bar":
                    v34 := GenS5{F0: -489, F1: 255}
                    _ = v34
                case "\n\t":
                    v1 = v1
                default:
                    var v35 int = -19
                    _ = v35
                }
            }
        }
    }
    switch v3 {
    case -1:
        var v36 bool = ((!(825 <= 2)) && v4)
        _ = v36
    case 2147483647:
        v37 := GenS5{F0: v3, F1: 18446744073709551615}
        _ = v37
    }
    v3 -= 4294967295
    if false {
        var v38 uint16 = 65535
        fmt.Println(v38, v38, v38)
        v38 += v38
        _ = v38
    } else {
        for fi39 := 0; fi39 < 17; fi39++ {
            v1 -= v1
            _ = fi39
        }
        switch "foo bar" {
        case "foo bar":
            v40 := GenS2{F0: 1000, F1: 248}
            _ = v40
        default:
            switch 127 {
            case 4294967295:
                if (58 != 42) {
                    v4 = false
                    if ((false || v4) || (1 < 2147483647)) {
                        var v41 int64 = -1
                        _ = v41
                    } else {
                        switch v3 {
                        case 32768:
                            for fi42 := 0; fi42 < 4; fi42++ {
                                if (("foo bar" + "hello") < "\n\t") {
                                    v3 += fi42
                                    if v4 {
                                        v2 -= (v2 | (v3 << ((255) & 63)))
                                        fmt.Println(v1, v3, v4)
                                        if (-128 > 40) {
                                            var v43 bool = (!(!v4))
                                            v1 += v1
                                            v4 = (!(238 == 255))
                                            _ = v43
                                        } else {
                                            if len(v5) > 0 {
                                                v5[0] = "hello"
                                            }
                                            v44 := GenS6{F0: 2, F1: false, F2: 42}
                                            var v45 uint16 = 782
                                            _ = v44
                                            _ = v45
                                        }
                                    }
                                }
                                switch (("" + "") + "foo bar") {
                                case "\n\t":
                                    v2 += fi42
                                case "a":
                                    if (!v4) {
                                        if len(v5) > 0 {
                                            v5[0] = ("a" + "\n\t")
                                        }
                                        for fi46 := 0; fi46 < 4; fi46++ {
                                            if v4 {
                                                fmt.Println(v3, fi46)
                                                v2 = (fi42 | (-(fi46)))
                                            }
                                            v3 = v3
                                            switch "a" {
                                            case "\n\t":
                                                switch fi42 {
                                                case 42:
                                                    v4 = (!v4)
                                                default:
                                                    if v4 {
                                                        v1 = ((float32((3.14159 * (1e10 * (3.14159 * -2.25)))) + (v1 / v1)) + (float32((1.5 / (1.5 * (0.1 + 1.5)))) * float32(0.0)))
                                                        v3 += 127
                                                    }
                                                }
                                            case "hello":
                                                if v4 {
                                                    var v47 int32 = -1
                                                    fmt.Println(v2, v2)
                                                    _ = v47
                                                }
                                            case "":
                                                var v48 float32 = float32(3.14159)
                                                _ = v48
                                            default:
                                                switch "foo bar" {
                                                case "\n\t":
                                                    if len(v5) > 0 {
                                                        v5[0] = ("" + ("hello" + ("" + "\n\t")))
                                                    }
                                                default:
                                                    switch "a" {
                                                    case "foo bar":
                                                        if (!("\n\t" <= ("foo bar" + "foo bar"))) {
                                                            var v49 uint16 = 0
                                                            var v50 bool = (65535 > 664)
                                                            _ = v49
                                                            _ = v50
                                                        }
                                                    case "":
                                                        fmt.Println(v3, v3, v1)
                                                    case "\n\t":
                                                        v3 -= fi46
                                                    }
                                                }
                                            }
                                            _ = fi46
                                        }
                                    } else {
                                        for fi51 := 0; fi51 < 8; fi51++ {
                                            if v4 {
                                                var v52 uint64 = 378
                                                v5 = append(v5, "foo bar")
                                                _ = v52
                                            }
                                            v3 += v2
                                            v3 -= 32768
                                            _ = fi51
                                        }
                                    }
                                default:
                                    for fi53 := 0; fi53 < 14; fi53++ {
                                        if v4 {
                                            var v54 int8 = 100
                                            v54 = (v54 * v54)
                                            _ = v54
                                        }
                                        fmt.Println(fi42, v1, fi53)
                                        _ = len(v5)
                                        _ = fi53
                                    }
                                }
                                var v55 int8 = 127
                                _ = fi42
                                _ = v55
                            }
                        }
                        _ = len(v5)
                    }
                }
            case 127:
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            case 9223372036854775807:
                v56 := GenS7{F0: "foo bar"}
                _ = v56
            }
        }
        _ = len(v5)
    }
    v57 := GenS2{F0: 1000, F1: 255}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v58 uint8 = genFn1()
    var v59 uint16 = 255
    fmt.Println(v59, v58)
    for fi60 := 0; fi60 < 20; fi60++ {
        v5 = append(v5, ("hello" + ""))
        _ = fi60
    }
    var v61 bool = (v4 || v4)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v58
    _ = v59
    _ = v61
}
