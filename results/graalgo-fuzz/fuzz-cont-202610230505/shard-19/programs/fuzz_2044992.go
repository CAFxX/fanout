package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 bool; F1 uint64; F2 bool }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint32 }
func genFn0(p0 uint8, p1 uint64, p2 int8) int64 {
    for fi6 := 0; fi6 < 10; fi6++ {
        v7 := []int{(fi6 & (fi6 / (((fi6 ^ (fi6 * fi6))) | 1)))}
        var v8 int64 = 42
        switch fi6 {
        case 1:
            switch "foo bar" {
            case "foo bar":
                for fi9 := 0; fi9 < 5; fi9++ {
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                    switch ("\n\t" + "a") {
                    case "foo bar":
                        v10 := GenS1{F0: 344, F1: "hello"}
                        _ = v10
                    }
                    _ = fi9
                }
            case "hello":
                fmt.Println(p1, v8, p0)
            default:
                var v11 uint8 = (p0 ^ (p0 << ((4294967295) & 63)))
                _ = v11
            }
        case 127:
            switch "\n\t" {
            case "a":
                if false {
                    for fi12 := 0; fi12 < 14; fi12++ {
                        p2 = p2
                        p0 = p0
                        _ = fi12
                    }
                    switch "foo bar" {
                    case "foo bar":
                        p1 += 42
                    case "a":
                        v13 := GenS2{F0: false, F1: (^(p1)), F2: true}
                        _ = v13
                    case "hello":
                        for fi14 := 0; fi14 < 10; fi14++ {
                            if (!false) {
                                fmt.Println(fi14, p1, p1)
                                if true {
                                    var v15 uint8 = p0
                                    _ = v15
                                }
                            }
                            switch "\n\t" {
                            case "foo bar":
                                var v16 int16 = -32768
                                _ = v16
                            default:
                                var v17 string = ""
                                _ = v17
                            }
                            p0 -= 107
                            _ = fi14
                        }
                    }
                    switch fi6 {
                    case 9223372036854775807:
                        p2 = (p2 - (p2 ^ 1))
                    case 718:
                        fmt.Println(fi6, v8, fi6)
                    default:
                        var v18 uint = 255
                        _ = v18
                    }
                }
            default:
                switch 9223372036854775807 {
                case -898:
                    v19 := GenS1{F0: 0, F1: "a"}
                    _ = v19
                case 9223372036854775807:
                    for fi20 := 0; fi20 < 6; fi20++ {
                        var v21 float64 = (0.0 / 0.1)
                        _ = fi20
                        _ = v21
                    }
                case -830:
                    _ = len(v7)
                }
            }
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    v22 := GenS1{F0: 42, F1: ("\n\t" + "hello")}
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    return -6
}
func genFn1(p0 uint16, p1 uint32, p2 uint32) string {
    p2 = (p1 << ((255) & 63))
    if (genFn0(127, 4294967295, 0) < 9223372036854775807) {
        switch 2147483647 {
        case 392:
            for fi23 := 0; fi23 < 14; fi23++ {
                for fi24 := 0; fi24 < 12; fi24++ {
                    p1 += (p2 << ((255) & 63))
                    fmt.Println(p2, p1, p1)
                    var v25 float64 = 0.1
                    p2 += 42
                    v26 := GenS2{F0: true, F1: 2, F2: false}
                    _ = fi24
                    _ = v25
                    _ = v26
                }
                _ = fi23
            }
        }
        v27 := []int{9223372036854775807, int(p0), -9223372036854775808}
        fmt.Println(p0, p1)
        _ = v27
    }
    p2 = (p2 / ((uint32(p0)) | 1))
    if (false || false) {
        for fi28 := 0; fi28 < 8; fi28++ {
            p2 += (^(p2))
            p0 += 0
            _ = fi28
        }
    } else {
        v29 := GenS2{F0: false, F1: 831, F2: false}
        _ = v29
    }
    for fi30 := 0; fi30 < 15; fi30++ {
        v31 := GenS1{F0: 65535, F1: "a"}
        _ = fi30
        _ = v31
    }
    _ = p0
    _ = p1
    _ = p2
    return "hello"
}

func main() {
    var v1 int32 = 42
    var v2 int64 = 9223372036854775807
    var v3 int16 = -441
    var v4 float64 = float64(3.14159)
    var v5 int = -1
    var v32 uint64 = 0
    if false {
        v33 := GenS3{F0: (v1 / ((v1) | 1))}
        var v34 uint8 = 2
        v34 -= v34
        _ = v33
        _ = v34
    } else {
        var v35 int64 = genFn0(255, (v32 % ((v32) | 1)), -128)
        v36 := []int{(v5 - -2147483648), (-(v5)), (^(v5)), v5}
        switch "" {
        case "a":
            v5 = v5
        }
        _ = v35
        _ = v36
    }
    switch v5 {
    case 2147483647:
        var v37 int16 = v3
        _ = v37
    case 32767:
        fmt.Println(v5, v32, v4)
    default:
        if (127 > -128) {
            var v38 float64 = ((v4 + v4) + v4)
            v2 = (^(v2))
            var v39 uint16 = 0
            _ = v38
            _ = v39
        }
    }
    if (1.5 != 3.14159) {
        v4 = v4
    } else {
        v1 = (^(v1))
        var v40 int64 = (v2 ^ (v2 | 9223372036854775807))
        _ = v40
    }
    var v41 uint8 = 1
    if false {
        var v42 string = "hello"
        var v43 uint64 = (v32 | (v32 * 742))
        _ = v42
        _ = v43
    }
    switch -2147483648 {
    case 9223372036854775807:
        var v44 string = ("a" + ("hello" + "a"))
        _ = v44
    case 0:
        if true {
            v2 = (^(v2))
        }
    case 2147483647:
        switch (v5 & (v5 ^ v5)) {
        case -479:
            if true {
                if true {
                    for fi45 := 0; fi45 < 18; fi45++ {
                        if (("" + ((("foo bar" + ("a" + "a")) + "") + ("hello" + "hello"))) >= "a") {
                            var v46 int64 = 1
                            v47 := GenS4{F0: 2147483647}
                            _ = v46
                            _ = v47
                        } else {
                            for fi48 := 0; fi48 < 9; fi48++ {
                                switch "foo bar" {
                                case "":
                                    switch "hello" {
                                    case "hello":
                                        if true {
                                            var v49 int8 = 100
                                            if false {
                                                var v50 float32 = 0.0
                                                v51 := GenS3{F0: (v1 ^ 0)}
                                                switch "foo bar" {
                                                case "foo bar":
                                                    v50 -= (v50 * v50)
                                                }
                                                _ = v50
                                                _ = v51
                                            }
                                            _ = v49
                                        }
                                    case "foo bar":
                                        var v52 float64 = (v4 * v4)
                                        _ = v52
                                    default:
                                        if true {
                                            v53 := []int{}
                                            if false {
                                                var v54 uint64 = v32
                                                v32 -= (v54 * (v54 % (((^(v54))) | 1)))
                                                _ = v54
                                            } else {
                                                v2 += (v2 << ((2) & 63))
                                                v55 := GenS2{F0: (!false), F1: 2, F2: true}
                                                _ = v55
                                            }
                                            _ = v53
                                        } else {
                                            var v56 int = 32767
                                            _ = v56
                                        }
                                    }
                                default:
                                    v57 := []int{0, (fi48 + fi48)}
                                    _ = v57
                                }
                                switch "" {
                                case "a":
                                    v58 := GenS4{F0: 42}
                                    _ = v58
                                case "hello":
                                    v3 = 0
                                }
                                var v59 float32 = float32(v4)
                                _ = fi48
                                _ = v59
                            }
                            var v60 uint = 1
                            _ = v60
                        }
                        v41 -= (^(v41))
                        _ = fi45
                    }
                    v5 = v5
                    var v61 int32 = (v1 >> ((2) & 63))
                    _ = v61
                } else {
                    for fi62 := 0; fi62 < 5; fi62++ {
                        v63 := GenS3{F0: -1}
                        v1 -= v1
                        for fi64 := 0; fi64 < 8; fi64++ {
                            v65 := []int{0, (fi62 * (-(fi64))), (^(fi62))}
                            _ = fi64
                            _ = v65
                        }
                        _ = fi62
                        _ = v63
                    }
                }
            }
        case 32767:
            for fi66 := 0; fi66 < 15; fi66++ {
                v32 = v32
                _ = fi66
            }
        default:
            if ("" != "hello") {
                v67 := []int{-2147483648, (v5 - v5), 622, (^(v5))}
                var v68 bool = false
                switch (v5 * (-(v5))) {
                case -566:
                    if len(v67) > 0 {
                        v67[0] = -1
                    }
                case 2147483647:
                    v41 = v41
                case 9223372036854775807:
                    v4 = (1e10 * (v4 / -2.25))
                default:
                    v1 -= 1
                }
                _ = v67
                _ = v68
            } else {
                if (394 >= 128) {
                    v1 = (v1 << ((uint(v32)) & 63))
                } else {
                    for fi69 := 0; fi69 < 20; fi69++ {
                        switch fi69 {
                        case 2147483647:
                            v70 := []int{(-(v5)), (fi69 % (((v5 - (v5 << ((42) & 63)))) | 1))}
                            _ = v70
                        case -502:
                            for fi71 := 0; fi71 < 11; fi71++ {
                                fmt.Println(v5, fi71)
                                v72 := GenS1{F0: 0, F1: ""}
                                var v73 int32 = int32(v3)
                                _ = fi71
                                _ = v72
                                _ = v73
                            }
                        case -88:
                            v1 -= (v1 | (v1 / ((1) | 1)))
                        default:
                            var v74 uint16 = 32767
                            _ = v74
                        }
                        v1 = v1
                        _ = fi69
                    }
                    var v75 float32 = float32(((v4 / 0.1) + float64((-2.25 + (float32(0.1) - ((0.1 / (1.5 * 3.14159)) + 0.0))))))
                    _ = v75
                }
            }
        }
    }
    v76 := GenS2{F0: (!(947 > 147)), F1: (v32 % (((v32 - (v32 + v32))) | 1)), F2: true}
    v77 := GenS4{F0: 65535}
    v2 = int64(v1)
    v41 = 78
    v78 := GenS1{F0: 4294967295, F1: ("\n\t" + (("hello" + "a") + ("hello" + "\n\t")))}
    var v79 float32 = (((1e10 + ((float32(-2.25) / 1.0) - 0.1)) * 0.1) * float32((v4 * 0.0)))
    v1 = (v1 / ((v1) | 1))
    v5 -= (-(v5))
    v80 := []int{272, (v5 << ((405) & 63)), (v5 ^ (-(v5)))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v79:", v79)
    fmt.Println("slicelen v80:", len(v80), v80)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v41
    _ = v5
    _ = v76
    _ = v77
    _ = v78
    _ = v79
    _ = v80
}
