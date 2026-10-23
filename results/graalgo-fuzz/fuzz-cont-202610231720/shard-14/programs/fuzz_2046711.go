package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 uint16; F2 uint64 }
type GenS3 struct { F0 uint64; F1 int8; F2 int8 }
type GenS4 struct { F0 string; F1 int16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int16; F1 int16; F2 int64 }
type GenS7 struct { F0 uint64; F1 int32; F2 int64 }
type GenS8 struct { F0 uint8; F1 int32; F2 uint64 }
type GenS9 struct { F0 int8; F1 uint32; F2 int }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 uint8; F1 int; F2 uint16 }
type GenS13 struct { F0 int32; F1 uint32; F2 int }
func genFn0(p0 uint32, p1 bool) int16 {
    v5 := GenS1{F0: -246}
    p1 = (0 != 32767)
    var v6 uint = 4294967295
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 0
}
func genFn1(p0 int) string {
    fmt.Println(p0)
    for fi7 := 0; fi7 < 4; fi7++ {
        v8 := GenS2{F0: 512, F1: 32767, F2: 1}
        p0 = 42
        if (1 >= 83) {
            var v9 float64 = -2.25
            v10 := GenS3{F0: 150, F1: 42, F2: -128}
            _ = v10
            _ = v9
        }
        _ = fi7
        _ = v8
    }
    for fi11 := 0; fi11 < 8; fi11++ {
        p0 = (fi11 ^ 128)
        _ = fi11
    }
    _ = p0
    return (("foo bar" + "") + "foo bar")
}

func main() {
    var v1 uint64 = 4294967295
    var v2 uint32 = 42
    var v3 int64 = 1
    var v4 uint32 = 4294967295
    v12 := []int{32768}
    var v13 uint = 255
    if (!(-9223372036854775808 == 1)) {
        for fi14 := 0; fi14 < 1; fi14++ {
            var v15 uint8 = 2
            _ = len(v12)
            _ = fi14
            _ = v15
        }
        v3 = 42
        if len(v12) > 0 {
            v12[0] = 32768
        }
    } else {
        if false {
            fmt.Println(v4, v13)
        } else {
            if (((genFn0((v2 ^ 315), (!(!true))) >= -32768) || (uint16(v1) < 42)) && (false || (!false))) {
                fmt.Println(v2)
                v12 = append(v12, 0)
            } else {
                fmt.Println(v4, v13, v13)
            }
            var v16 int64 = v3
            _ = v16
        }
        v2 = v4
        var v17 string = ("foo bar" + "hello")
        _ = v17
    }
    var v18 bool = ((v1 / ((v1) | 1)) <= uint64(v4))
    if v18 {
        v3 -= (v3 * 942)
        v19 := GenS4{F0: "foo bar", F1: -1}
        _ = v19
    }
    if (!(4294967295 <= 9223372036854775807)) {
        fmt.Println(v1, v2)
        if (17 > 128) {
            for fi20 := 0; fi20 < 18; fi20++ {
                switch (fi20 << ((18446744073709551615) & 63)) {
                case 1:
                    switch fi20 {
                    case -2147483648:
                        var v21 uint32 = (v4 >> ((18446744073709551615) & 63))
                        _ = v21
                    case 822:
                        for fi22 := 0; fi22 < 20; fi22++ {
                            if (v13 >= v13) {
                                v13 -= (v13 & (v13 + (v13 + 381)))
                                v3 += (v3 << ((1) & 63))
                                var v23 string = "hello"
                                _ = v23
                            }
                            if (!true) {
                                var v24 int16 = -32768
                                _ = v24
                            }
                            if true {
                                if ((!v18) || v18) {
                                    fmt.Println(fi22, fi22, v13)
                                }
                            }
                            _ = fi22
                        }
                    case 42:
                        if true {
                            var v25 int8 = 127
                            v18 = v18
                            v2 -= 0
                            _ = v25
                        }
                    default:
                        for fi26 := 0; fi26 < 9; fi26++ {
                            switch (fi20 >> (((v13 - uint(v1))) & 63)) {
                            case 32767:
                                if ((!("" == ((("hello" + "hello") + "hello") + ""))) && false) {
                                    var v27 bool = v18
                                    _ = v27
                                }
                            default:
                                v1 += (v1 - (^(v1)))
                            }
                            v28 := GenS5{F0: 1}
                            fmt.Println(v18, v2, fi26)
                            _ = fi26
                            _ = v28
                        }
                    }
                case 250:
                    v1 -= (v1 << (((v13 ^ v13)) & 63))
                default:
                    fmt.Println(v18)
                }
                v18 = (((1e10 != ((1e10 + ((((0.0 - ((((1e10 - 0.1) - (0.0 + (float32(1e10) * (-2.25 * 3.14159)))) / (-2.25 / 1e10)) + 1.5)) * (((float32(3.14159) + -2.25) / 0.1) * (((3.14159 + (1e10 * 1.5)) + (3.14159 * (((((3.14159 - (((3.14159 * (1.5 - (0.0 + (0.0 - (((0.0 * ((0.0 / (0.1 + 1e10)) / 1e10)) * 1e10) * (1e10 * (-2.25 * 3.14159))))))) * float32(0.1)) - (1e10 - 0.1))) + (3.14159 - ((0.1 / 1.5) - (((float32((-2.25 + 3.14159)) * (1e10 * (1.5 - 1.5))) * (0.0 + 0.0)) * (0.1 * ((0.0 / 0.1) / float32(float64(-2.25)))))))) * 3.14159) * 0.1) * ((1e10 * (1.5 / float32(3.14159))) * 3.14159)))) + ((((0.1 / ((-2.25 + 0.0) - (1e10 - (-2.25 * 1e10)))) / 1.0) + (1e10 * float32(0.1))) * 1.5)))) * 1e10) + 1e10)) - 1.5)) || v18) && (4294967295 < v1))
                _ = fi20
            }
            if len(v12) > 0 {
                v12[0] = int(v3)
            }
            var v29 uint = (v13 / ((v13) | 1))
            _ = v29
        } else {
            v4 = (v2 / ((2147483647) | 1))
            switch -2147483648 {
            case 4294967295:
                if true {
                    if v18 {
                        v12 = append(v12, int(v3))
                        for fi30 := 0; fi30 < 15; fi30++ {
                            var v31 int32 = -136
                            _ = fi30
                            _ = v31
                        }
                        var v32 int16 = 0
                        _ = v32
                    } else {
                        fmt.Println(v13)
                        v33 := GenS1{F0: 42}
                        _ = v33
                    }
                }
            case 0:
                var v34 uint8 = 42
                _ = v34
            }
            if v18 {
                v35 := GenS1{F0: 42}
                switch ("hello" + "foo bar") {
                case "a":
                    v18 = v18
                case "foo bar":
                    if (false || v18) {
                        v36 := GenS2{F0: uint16(v4), F1: 912, F2: v1}
                        v4 = (v2 >> (((^(v13))) & 63))
                        _ = v36
                    }
                case "":
                    v37 := GenS6{F0: 1000, F1: -834, F2: v3}
                    _ = v37
                default:
                    for fi38 := 0; fi38 < 3; fi38++ {
                        v2 += (v4 + (v4 & v2))
                        _ = fi38
                    }
                }
                _ = v35
            }
        }
    } else {
        v39 := GenS7{F0: 2, F1: 250, F2: (v3 ^ v3)}
        v1 += 2
        fmt.Println(v1, v4)
        _ = v39
    }
    var v40 int64 = v3
    var v41 bool = true
    fmt.Println(v4, v13)
    var v42 uint = v13
    if ((v3 << ((v42) & 63)) > (v40 + (-(v3)))) {
        for fi43 := 0; fi43 < 9; fi43++ {
            v44 := GenS8{F0: 2, F1: 1, F2: 18446744073709551615}
            var v45 float32 = (float32(float64((3.14159 + 3.14159))) - float32(1e10))
            for fi46 := 0; fi46 < 13; fi46++ {
                var v47 uint64 = (v1 - 809)
                switch ((("\n\t" + ("foo bar" + (("foo bar" + "\n\t") + ("foo bar" + "a")))) + ("" + (("a" + ("foo bar" + "")) + "\n\t"))) + ("hello" + (("hello" + (("" + "\n\t") + "")) + "hello"))) {
                case "hello":
                    v18 = (float32(0.0) < v45)
                    v48 := GenS9{F0: -92, F1: v2, F2: 0}
                    v4 = 1
                    fmt.Println(v40)
                    _ = v48
                }
                _ = fi46
                _ = v47
            }
            _ = fi43
            _ = v44
            _ = v45
        }
        for fi49 := 0; fi49 < 19; fi49++ {
            var v50 string = ("foo bar" + "\n\t")
            for fi51 := 0; fi51 < 8; fi51++ {
                if false {
                    v52 := GenS7{F0: v1, F1: 42, F2: 290}
                    var v53 int32 = -705
                    v41 = (594 == genFn0(uint32(v42), v18))
                    _ = v52
                    _ = v53
                } else {
                    v54 := GenS10{F0: v42}
                    _ = v54
                }
                _ = fi51
            }
            v55 := GenS11{F0: 2147483647}
            _ = fi49
            _ = v50
            _ = v55
        }
    }
    var v56 uint64 = 4294967295
    if (42 < -1) {
        if (!false) {
            var v57 int16 = 0
            _ = v57
        } else {
            for fi58 := 0; fi58 < 12; fi58++ {
                var v59 float64 = 1e10
                if v18 {
                    switch int(v13) {
                    case 1:
                        for fi60 := 0; fi60 < 3; fi60++ {
                            if ((v42 > (v13 << ((v13) & 63))) || (1e10 <= (((0.0 * (1.5 + (3.14159 + 3.14159))) / 3.14159) * 1.5))) {
                                if v18 {
                                    var v61 float64 = 1.5
                                    _ = v61
                                }
                            } else {
                                v62 := GenS9{F0: -1, F1: v4, F2: (-(fi60))}
                                var v63 float64 = (v59 - 3.14159)
                                v2 = (v4 % ((v4) | 1))
                                _ = v62
                                _ = v63
                            }
                            if ((true || false) && (1 >= 0)) {
                                v18 = (0.1 < (-2.25 - (((3.14159 - (-2.25 / 1.5)) * (1e10 - 1.5)) + 0.0)))
                                switch fi60 {
                                case -1:
                                    if len(v12) > 0 {
                                        v12[0] = (fi58 ^ fi60)
                                    }
                                case -773:
                                    v4 -= v4
                                }
                                fmt.Println(v18, v18)
                                switch (("" + "") + ("" + "\n\t")) {
                                case "hello":
                                    var v64 int64 = (v40 & (v40 + v40))
                                    _ = v64
                                case "a":
                                    if (!true) {
                                        var v65 float64 = v59
                                        v1 = v56
                                        v41 = v41
                                        _ = v65
                                    } else {
                                        if v18 {
                                            v66 := GenS4{F0: ("\n\t" + ("a" + (("" + "\n\t") + ("" + (("" + "foo bar") + "hello"))))), F1: 0}
                                            if (1 < 0) {
                                                v67 := GenS11{F0: -2147483648}
                                                if ((^(v2)) != v2) {
                                                    if (1 > 123) {
                                                        switch 81 {
                                                        case 460:
                                                            v56 -= 779
                                                        case -272:
                                                            var v68 int = (^(fi58))
                                                            _ = v68
                                                        case 0:
                                                            fmt.Println(v4)
                                                        default:
                                                            switch fi60 {
                                                            case -9223372036854775808:
                                                                if len(v12) > 0 {
                                                                    v12[0] = fi58
                                                                }
                                                            default:
                                                                var v69 bool = v41
                                                                _ = v69
                                                            }
                                                        }
                                                        var v70 int = 128
                                                        _ = v70
                                                    }
                                                }
                                                if len(v12) >= 2 {
                                                    v12 = v12[1:len(v12)]
                                                }
                                                var v71 float32 = (3.14159 * (1e10 / -2.25))
                                                _ = v67
                                                _ = v71
                                            } else {
                                                var v72 int8 = 42
                                                if (genFn0(v4, v18) == 32767) {
                                                    v2 -= v2
                                                } else {
                                                    switch ("foo bar" + ("\n\t" + "\n\t")) {
                                                    case "":
                                                        if false {
                                                            if len(v12) >= 2 {
                                                                v12 = v12[1:len(v12)]
                                                            }
                                                            if len(v12) >= 2 {
                                                                v12 = v12[1:len(v12)]
                                                            }
                                                        }
                                                    default:
                                                        fmt.Println(v18, v13, v72)
                                                        var v73 uint64 = 18446744073709551615
                                                        if v18 {
                                                            var v74 uint64 = 1
                                                            var v75 uint64 = (^(v1))
                                                            _ = v74
                                                            _ = v75
                                                        }
                                                        var v76 float32 = (0.0 + 0.1)
                                                        _ = v73
                                                        _ = v76
                                                    }
                                                    v77 := GenS12{F0: 127, F1: (-(fi58)), F2: 32767}
                                                    v4 = (v4 - v2)
                                                    _ = v77
                                                }
                                                fmt.Println(v59, v2, v59)
                                                _ = v72
                                            }
                                            _ = v66
                                        }
                                        fmt.Println(v4)
                                        v2 -= 2
                                    }
                                case "":
                                    var v78 string = "foo bar"
                                    _ = v78
                                default:
                                    switch "foo bar" {
                                    case "hello":
                                        v18 = (((1e10 / (float32((0.0 / (v59 / (v59 / 0.1)))) / 0.1)) / -2.25) < 1e10)
                                    case "foo bar":
                                        v79 := GenS13{F0: 1, F1: 65535, F2: fi60}
                                        _ = v79
                                    case "":
                                        if (true || false) {
                                            v80 := GenS10{F0: (v42 * 0)}
                                            _ = v80
                                        } else {
                                            v81 := GenS2{F0: 42, F1: 255, F2: 4294967295}
                                            _ = v81
                                        }
                                        var v82 uint8 = 2
                                        _ = v82
                                    default:
                                        v2 = (v2 % ((v2) | 1))
                                    }
                                }
                            }
                            _ = fi60
                        }
                    case 42:
                        v40 += v40
                    default:
                        var v83 uint8 = 127
                        _ = v83
                    }
                    v84 := GenS1{F0: -1}
                    for fi85 := 0; fi85 < 8; fi85++ {
                        v18 = ((v40 != (v40 >> ((255) & 63))) || true)
                        v3 = int64(fi85)
                        _ = fi85
                    }
                    _ = v84
                }
                var v86 int8 = 1
                _ = fi58
                _ = v59
                _ = v86
            }
            fmt.Println(v1, v56)
        }
        v56 -= v56
        var v87 uint16 = 85
        _ = v87
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v56
}
