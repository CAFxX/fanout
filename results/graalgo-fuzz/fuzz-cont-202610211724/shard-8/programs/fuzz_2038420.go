package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int; F1 int64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int; F1 int; F2 uint8 }
type GenS6 struct { F0 bool; F1 int32; F2 string }
type GenS7 struct { F0 uint64; F1 string }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 uint8; F1 int16; F2 uint64 }
type GenS10 struct { F0 uint; F1 uint16; F2 uint64 }
type GenS11 struct { F0 int16 }

func main() {
    var v1 string = "foo bar"
    var v2 float32 = 3.14159
    v3 := make(map[string]string)
    v4 := make(map[uint32]int)
    var v5 uint64 = 18446744073709551615
    var v6 float64 = 0.1
    v2 -= float32(float64(0.0))
    switch ("\n\t" + v1) {
    case "":
        for fi7 := 0; fi7 < 16; fi7++ {
            fmt.Println(v6, v5)
            _ = fi7
        }
    case "\n\t":
        for fi8 := 0; fi8 < 17; fi8++ {
            fmt.Println(v5, v1)
            v6 -= 1.5
            _ = fi8
        }
    default:
        fmt.Println(v1)
    }
    for fi9 := 0; fi9 < 10; fi9++ {
        if false {
            v10 := []int{373, (fi9 + -1), 929, fi9}
            _ = v10
        } else {
            for fi11 := 0; fi11 < 2; fi11++ {
                if true {
                    if false {
                        if false {
                            v1 = ("hello" + v1)
                        }
                    }
                    if false {
                        v2 -= v2
                    } else {
                        var v12 uint16 = 0
                        if true {
                            v6 = (0.1 + 3.14159)
                            v13 := []int{2147483647, (fi11 & -270)}
                            _ = v13
                        } else {
                            v5 += v5
                        }
                        v14 := GenS1{F0: v1, F1: ("" + (v1 + v1)), F2: 0}
                        _ = v12
                        _ = v14
                    }
                }
                v15 := []int{fi9, fi9}
                _ = fi11
                _ = v15
            }
        }
        _ = fi9
    }
    for fi16 := 0; fi16 < 19; fi16++ {
        if (!(false || (!false))) {
            var v17 bool = ((v5 + v5) != 183)
            fmt.Println(v17)
            var v18 int16 = 787
            _ = v17
            _ = v18
        } else {
            v19 := GenS2{F0: (fi16 & (fi16 % ((-585) | 1)))}
            v5 = v5
            v5 = (v5 - (^(v5)))
            _ = v19
        }
        v20 := []int{fi16}
        for fi21 := 0; fi21 < 20; fi21++ {
            var v22 int64 = -9223372036854775808
            _ = fi21
            _ = v22
        }
        _ = fi16
        _ = v20
    }
    v23 := GenS3{F0: 1, F1: -570}
    v24 := GenS3{F0: 2147483647, F1: -74}
    switch v1 {
    case "hello":
        v5 = v5
    case "a":
        v25 := GenS3{F0: 1, F1: -9223372036854775808}
        _ = v25
    default:
        if false {
            var v26 float32 = 3.14159
            _ = v26
        } else {
            if true {
                if false {
                    v2 = float32((v6 * (v6 - v6)))
                    var v27 int8 = 127
                    _ = v27
                }
                for fi28 := 0; fi28 < 18; fi28++ {
                    v2 = 1.5
                    _ = fi28
                }
            } else {
                var v29 float32 = v2
                var v30 uint = 652
                var v31 uint32 = 2147483648
                _ = v29
                _ = v30
                _ = v31
            }
            if (-2147483648 > 9223372036854775807) {
                var v32 int16 = 13
                v5 = (v5 | v5)
                v33 := GenS4{F0: 904}
                _ = v32
                _ = v33
            } else {
                v34 := GenS3{F0: 421, F1: -9223372036854775808}
                v2 = v2
                v35 := GenS2{F0: -13}
                _ = v34
                _ = v35
            }
            _ = len(v4)
        }
    }
    v36 := []int{-1, 552}
    for fi37 := 0; fi37 < 5; fi37++ {
        v5 = (v5 | v5)
        fmt.Println(fi37, fi37)
        fmt.Println(v2, v1)
        _ = fi37
    }
    v5 = (v5 % ((42) | 1))
    var v38 int32 = 0
    var v39 int64 = 0
    for fi40 := 0; fi40 < 18; fi40++ {
        var v41 uint8 = 127
        switch fi40 {
        case 410:
            if len(v36) > 0 {
                v36[0] = (fi40 ^ 42)
            }
        case 2147483647:
            if true {
                switch v1 {
                case "hello":
                    for fi42 := 0; fi42 < 17; fi42++ {
                        v43 := GenS5{F0: (fi40 ^ (fi42 % ((-1) | 1))), F1: fi40, F2: v41}
                        fmt.Println(v38, v39)
                        v36 = append(v36, (fi42 * fi40))
                        _ = fi42
                        _ = v43
                    }
                default:
                    fmt.Println(v6, v38)
                }
            } else {
                var v44 uint32 = 0
                if true {
                    if (1 != -128) {
                        if ((-(v39)) == v39) {
                            v38 = -694
                        } else {
                            var v45 uint = 65535
                            for fi46 := 0; fi46 < 13; fi46++ {
                                switch v1 {
                                case "a":
                                    var v47 float64 = ((((0.1 - v6) / 1.0) / float64((float32(v6) - v2))) - v6)
                                    _ = v47
                                case "foo bar":
                                    v45 = v45
                                default:
                                    _ = len(v36)
                                }
                                var v48 uint = 65535
                                _ = fi46
                                _ = v48
                            }
                            _ = v45
                        }
                        v39 += v39
                    }
                }
                var v49 uint16 = 666
                _ = v44
                _ = v49
            }
        case 0:
            for fi50 := 0; fi50 < 12; fi50++ {
                v36 = append(v36, int(v39))
                v38 += (v38 | -2147483648)
                v6 += (v6 - v6)
                _ = fi50
            }
        default:
            for fi51 := 0; fi51 < 13; fi51++ {
                switch (fi51 & (fi51 >> ((255) & 63))) {
                case 42:
                    if (false && false) {
                        switch 619 {
                        case 9223372036854775807:
                            switch fi40 {
                            case -1:
                                _ = len(v3)
                            case -276:
                                var v52 uint16 = 775
                                _ = v52
                            default:
                                delete(v3, "")
                            }
                        case 127:
                            delete(v3, ("foo bar" + v1))
                        case -1:
                            v2 -= (1.5 + (v2 + v2))
                        }
                        v38 -= v38
                    }
                default:
                    v53 := GenS6{F0: true, F1: (-(v38)), F2: v1}
                    switch 0 {
                    case -31:
                        fmt.Println(fi51)
                    }
                    if true {
                        v41 += (v41 * (^(v41)))
                        var v54 float64 = v6
                        switch (fi51 / (((fi51 >> ((778) & 63))) | 1)) {
                        case -9223372036854775808:
                            v55 := GenS5{F0: (-(fi40)), F1: 9223372036854775807, F2: 42}
                            _ = v55
                        case 980:
                            var v56 bool = false
                            _ = v56
                        case -597:
                            switch fi51 {
                            case -2147483648:
                                switch v1 {
                                case "a":
                                    v3[v1] = "foo bar"
                                case "hello":
                                    var v57 float64 = (float64((float32(-2.25) * float32((1e10 + v6)))) + float64(float32(1.5)))
                                    _ = v57
                                default:
                                    switch fi40 {
                                    case 32768:
                                        if false {
                                            v3[""] = v1
                                            if len(v36) > 0 {
                                                v36[0] = (fi40 ^ 0)
                                            }
                                        }
                                    case 123:
                                        v58 := GenS7{F0: v5, F1: v1}
                                        _ = v58
                                    case 9223372036854775807:
                                        v39 = (v39 >> ((261) & 63))
                                    default:
                                        switch "\n\t" {
                                        case "":
                                            fmt.Println(v54, fi51)
                                        case "hello":
                                            v38 += (^(v38))
                                        }
                                        if ((^(v41)) == 187) {
                                            fmt.Println(v5, v41, v39)
                                            v5 += (v5 - (^(v5)))
                                        }
                                    }
                                }
                            case 4294967295:
                                v59 := GenS8{F0: 772}
                                _ = v59
                            default:
                                var v60 uint = 65535
                                _ = v60
                            }
                        default:
                            switch (fi40 * (fi51 - (fi40 ^ int(v39)))) {
                            case -1:
                                v61 := GenS7{F0: (v5 << ((1) & 63)), F1: (v1 + "hello")}
                                _ = v61
                            case -602:
                                var v62 uint = 255
                                _ = v62
                            default:
                                if len(v36) > 0 {
                                    v36[0] = (^(fi51))
                                }
                            }
                        }
                        v63 := GenS6{F0: true, F1: v38, F2: ((v1 + v1) + "a")}
                        _ = v54
                        _ = v63
                    } else {
                        v64 := GenS9{F0: 160, F1: -1, F2: (^(v5))}
                        if (!(0 > 4294967295)) {
                            _ = len(v36)
                        }
                        v65 := GenS7{F0: 4294967295, F1: ("" + (v1 + v1))}
                        _ = v64
                        _ = v65
                    }
                    _ = v53
                }
                _ = fi51
            }
        }
        _ = fi40
        _ = v41
    }
    for fi66 := 0; fi66 < 9; fi66++ {
        var v67 uint64 = 522
        for fi68 := 0; fi68 < 3; fi68++ {
            fmt.Println(fi68)
            if (93 > 1) {
                fmt.Println(v67)
                var v69 bool = true
                v70 := GenS10{F0: 65535, F1: 194, F2: v5}
                _ = v69
                _ = v70
            } else {
                v2 = float32((v6 / v6))
            }
            switch fi66 {
            case 256:
                var v71 uint = uint(v5)
                _ = v71
            case -678:
                var v72 int32 = v38
                _ = v72
            default:
                _ = len(v3)
            }
            _ = fi68
        }
        _ = fi66
        _ = v67
    }
    v73 := GenS2{F0: 865}
    if (127 >= 141) {
        switch 127 {
        case 128:
            switch 42 {
            case -9223372036854775808:
                v39 -= 42
            default:
                for fi74 := 0; fi74 < 20; fi74++ {
                    v75 := GenS2{F0: 32767}
                    for fi76 := 0; fi76 < 13; fi76++ {
                        if true {
                            if len(v36) > 0 {
                                v36[0] = -434
                            }
                        }
                        var v77 int16 = 42
                        _ = fi76
                        _ = v77
                    }
                    v78 := GenS11{F0: -32768}
                    _ = fi74
                    _ = v75
                    _ = v78
                }
            }
        case -1:
            _ = len(v36)
        case -2147483648:
            var v79 uint32 = 359
            _ = v79
        }
    } else {
        v36 = append(v36, 360)
        v1 = (("hello" + "a") + "foo bar")
        var v80 int64 = v39
        _ = v80
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v36
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v73
}
