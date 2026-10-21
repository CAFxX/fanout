package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32; F1 int; F2 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int64; F1 uint8 }
type GenS5 struct { F0 string; F1 int16; F2 uint64 }
type GenS6 struct { F0 uint16; F1 uint8 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint32; F1 bool; F2 int32 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 int; F1 int32 }
type GenS11 struct { F0 uint32; F1 bool; F2 uint8 }
type GenS12 struct { F0 bool; F1 int64; F2 uint64 }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint16 = 255
    var v3 string = ("hello" + ("\n\t" + "foo bar"))
    v4 := GenS1{F0: -692}
    v5 := GenS1{F0: 2147483647}
    v6 := []int{}
    var v7 int8 = 102
    var v8 int = 127
    for fi9 := 0; fi9 < 13; fi9++ {
        for fi10 := 0; fi10 < 2; fi10++ {
            if false {
                v11 := GenS1{F0: 0}
                switch -980 {
                case -836:
                    v2 = 65535
                default:
                    if false {
                        _ = len(v1)
                        fmt.Println(v8, v2, v2)
                        v1[2] = (^(fi10))
                    }
                }
                var v12 uint8 = 255
                _ = v11
                _ = v12
            } else {
                if false {
                    var v13 uint32 = 42
                    _ = v13
                }
                _ = len(v1)
                var v14 bool = (255 <= v2)
                _ = v14
            }
            v15 := GenS1{F0: 800}
            v16 := GenS2{F0: -2147483648, F1: fi9, F2: 18446744073709551615}
            v7 = (v7 - (v7 / ((v7) | 1)))
            if false {
                switch (-(fi10)) {
                case 42:
                    var v17 float64 = ((-2.25 + 0.0) / 1.0)
                    _ = v17
                case 127:
                    if false {
                        v2 -= 235
                    }
                case 953:
                    var v18 uint = 892
                    _ = v18
                default:
                    v3 = ""
                }
            } else {
                _ = len(v6)
                var v19 float32 = 1e10
                _ = v19
            }
            _ = fi10
            _ = v15
            _ = v16
        }
        _ = len(v1)
        _ = fi9
    }
    delete(v1, 423)
    v7 += v7
    fmt.Println(v3)
    if (false && (false || true)) {
        _ = len(v1)
    }
    for fi20 := 0; fi20 < 20; fi20++ {
        var v21 uint16 = 1
        _ = fi20
        _ = v21
    }
    v1[42] = v8
    v8 = (v8 >> ((974) & 63))
    fmt.Println(v8)
    v22 := GenS2{F0: -614, F1: (v8 + v8), F2: 177}
    if (!true) {
        for fi23 := 0; fi23 < 6; fi23++ {
            var v24 uint32 = 42
            v3 = ("" + "hello")
            if (false || false) {
                switch "a" {
                case "hello":
                    fmt.Println(v7, v3, v2)
                }
            } else {
                var v25 bool = (-911 == 42)
                _ = v25
            }
            _ = fi23
            _ = v24
        }
    } else {
        fmt.Println(v3, v2)
    }
    v2 -= v2
    for fi26 := 0; fi26 < 14; fi26++ {
        for fi27 := 0; fi27 < 4; fi27++ {
            var v28 uint32 = uint32(v2)
            if true {
                switch (-(v8)) {
                case -22:
                    var v29 uint = 65535
                    _ = v29
                case 444:
                    var v30 uint16 = (^(v2))
                    _ = v30
                default:
                    if false {
                        var v31 int64 = 42
                        var v32 int8 = int8(v2)
                        _ = v31
                        _ = v32
                    } else {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        if true {
                            v8 += (fi26 ^ 963)
                            if false {
                                v2 -= (v2 / ((v2) | 1))
                            }
                            v28 += v28
                        } else {
                            _ = len(v1)
                            var v33 uint64 = 2
                            if false {
                                v7 -= -1
                                v1[v28] = fi26
                            } else {
                                v33 -= 651
                                if false {
                                    v34 := GenS3{F0: 12}
                                    v33 += 1
                                    switch "" {
                                    case "a":
                                        if (9223372036854775807 < -586) {
                                            _ = len(v6)
                                        }
                                    case "foo bar":
                                        fmt.Println(v8, v3, v8)
                                    default:
                                        v28 = 2147483647
                                    }
                                    _ = v34
                                } else {
                                    if true {
                                        if (!("" >= v3)) {
                                            var v35 int = -9223372036854775808
                                            _ = v35
                                        } else {
                                            var v36 int16 = -32768
                                            _ = v36
                                        }
                                    } else {
                                        delete(v1, (v28 ^ v28))
                                    }
                                    _ = len(v1)
                                    if ((!false) && (18 <= 1)) {
                                        v33 += v33
                                    }
                                    switch v3 {
                                    case "":
                                        switch fi26 {
                                        case 2147483647:
                                            v37 := GenS2{F0: -2147483648, F1: -2147483648, F2: 0}
                                            _ = v37
                                        case 42:
                                            if true {
                                                if (-32768 > -255) {
                                                    v2 = v2
                                                    v33 = v33
                                                }
                                            }
                                        case 233:
                                            if (v3 != v3) {
                                                if (!true) {
                                                    v38 := GenS1{F0: -1}
                                                    v7 -= (v7 / ((-13) | 1))
                                                    if true {
                                                        v39 := GenS4{F0: -1, F1: 128}
                                                        _ = v39
                                                    }
                                                    _ = v38
                                                } else {
                                                    v40 := GenS2{F0: 42, F1: (fi26 | v8), F2: (v33 ^ v33)}
                                                    _ = v40
                                                }
                                            }
                                        default:
                                            switch fi27 {
                                            case 127:
                                                var v41 int8 = 100
                                                _ = v41
                                            case 9223372036854775807:
                                                v2 = 65535
                                            case 1:
                                                var v42 int16 = 32767
                                                _ = v42
                                            }
                                        }
                                    case "foo bar":
                                        if false {
                                            fmt.Println(v8, fi27)
                                            switch ("a" + v3) {
                                            case "hello":
                                                switch v3 {
                                                case "foo bar":
                                                    v43 := GenS5{F0: ("hello" + ("foo bar" + (v3 + v3))), F1: 901, F2: 2}
                                                    _ = v43
                                                case "hello":
                                                    if false {
                                                        fmt.Println(v2, v3)
                                                        if (!(false && (-2.25 > (3.14159 / (1.5 / 0.1))))) {
                                                            switch v3 {
                                                            case "foo bar":
                                                                var v44 float64 = 0.1
                                                                _ = v44
                                                            case "":
                                                                v8 = -973
                                                            case "a":
                                                                v45 := GenS6{F0: (v2 ^ v2), F1: 2}
                                                                _ = v45
                                                            }
                                                            switch (("a" + v3) + ((v3 + v3) + (v3 + (v3 + v3)))) {
                                                            case "\n\t":
                                                                v2 = (v2 + (^(v2)))
                                                            case "hello":
                                                                fmt.Println(v33)
                                                                if false {
                                                                    v2 = v2
                                                                    if ((v33 % (((v33 >> ((2) & 63))) | 1)) != 3) {
                                                                        v2 = v2
                                                                        fmt.Println(v7, fi26, v33)
                                                                        if len(v6) > 0 {
                                                                            v6[0] = (fi26 ^ (fi27 + (fi27 % ((32767) | 1))))
                                                                        }
                                                                        v46 := GenS6{F0: uint16(v28), F1: 127}
                                                                        v47 := GenS7{F0: 123}
                                                                        _ = v46
                                                                        _ = v47
                                                                    } else {
                                                                        fmt.Println(v3, v7)
                                                                        v1[(v28 >> ((0) & 63))] = fi26
                                                                    }
                                                                    fmt.Println(v8, v8)
                                                                }
                                                                if true {
                                                                    v28 = (v28 | v28)
                                                                }
                                                                fmt.Println(v2, v28, v7)
                                                                v48 := GenS8{F0: (^(v28)), F1: (-1 > 32767), F2: 1}
                                                                _ = v48
                                                            default:
                                                                var v49 uint64 = v33
                                                                _ = v49
                                                            }
                                                            var v50 uint8 = 42
                                                            _ = v50
                                                        } else {
                                                            v51 := GenS1{F0: -2147483648}
                                                            var v52 string = v3
                                                            _ = v51
                                                            _ = v52
                                                        }
                                                    } else {
                                                        if (1 <= uint(v33)) {
                                                            var v53 uint32 = 2147483647
                                                            _ = v53
                                                        }
                                                    }
                                                default:
                                                    v3 = "a"
                                                    if false {
                                                        switch 128 {
                                                        case 357:
                                                            var v54 int8 = v7
                                                            _ = v54
                                                        default:
                                                            v55 := GenS9{F0: v7}
                                                            _ = v55
                                                        }
                                                        v56 := GenS10{F0: fi27, F1: 786}
                                                        v57 := v1[(^(v28))]
                                                        var v58 int32 = 1
                                                        _ = v56
                                                        _ = v57
                                                        _ = v58
                                                    } else {
                                                        switch (fi27 & 336) {
                                                        case -85:
                                                            fmt.Println(v8, v3)
                                                        case -150:
                                                            v7 += v7
                                                        case 378:
                                                            v59 := GenS1{F0: -912}
                                                            _ = v59
                                                        }
                                                    }
                                                }
                                            case "a":
                                                if false {
                                                    v60 := GenS6{F0: 255, F1: 243}
                                                    if (-1 == 0) {
                                                        if (((0.1 / 1e10) * 1e10) <= (0.0 - 3.14159)) {
                                                            var v61 uint64 = (^(v33))
                                                            _ = v61
                                                        } else {
                                                            switch fi27 {
                                                            case 127:
                                                                v8 += v8
                                                            case 32767:
                                                                fmt.Println(v8, v28)
                                                            default:
                                                                fmt.Println(v2, v2)
                                                            }
                                                        }
                                                    }
                                                    var v62 uint64 = (v33 % ((v33) | 1))
                                                    var v63 int = (fi27 | fi27)
                                                    var v64 float64 = 0.0
                                                    v62 = 1
                                                    v33 -= v33
                                                    _ = v60
                                                    _ = v62
                                                    _ = v63
                                                    _ = v64
                                                }
                                            case "":
                                                v65 := GenS11{F0: (^(v28)), F1: (!false), F2: 128}
                                                _ = v65
                                            default:
                                                var v66 int16 = 1000
                                                _ = v66
                                            }
                                            var v67 int64 = 42
                                            _ = v67
                                        } else {
                                            fmt.Println(v8)
                                        }
                                    case "\n\t":
                                        v3 = (v3 + "hello")
                                    }
                                }
                            }
                            _ = v33
                        }
                        v3 = (v3 + v3)
                        v68 := GenS12{F0: (true && true), F1: 496, F2: 1}
                        _ = v68
                    }
                }
                v3 = v3
                v1[(v28 << ((1) & 63))] = (^(fi27))
                var v69 int8 = (-(v7))
                _ = v69
            } else {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                v28 += v28
                var v70 uint8 = uint8(v2)
                _ = v70
            }
            _ = fi27
            _ = v28
        }
        _ = fi26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
