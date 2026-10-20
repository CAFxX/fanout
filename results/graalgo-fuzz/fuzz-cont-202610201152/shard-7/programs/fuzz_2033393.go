package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 int32 }
type GenS2 struct { F0 uint; F1 int8; F2 uint64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int; F1 int32; F2 int }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int16; F1 int8; F2 int64 }
type GenS7 struct { F0 string; F1 int32; F2 int }

func main() {
    v1 := make(map[uint32]bool)
    var v2 string = ("\n\t" + "\n\t")
    v3 := make(map[int]bool)
    var v4 string = ((("\n\t" + ((((v2 + "foo bar") + ((v2 + (v2 + v2)) + ((v2 + ((v2 + ((v2 + "hello") + (v2 + v2))) + v2)) + v2))) + (v2 + v2)) + v2)) + "") + ((v2 + v2) + ((("\n\t" + v2) + v2) + v2)))
    if false {
        switch "a" {
        case "\n\t":
            v5 := GenS1{F0: 0, F1: 18446744073709551615, F2: -1}
            _ = v5
        case "":
            if (((-2.25 * 3.14159) + 0.0) <= -2.25) {
                v2 = v2
                v6 := GenS1{F0: 717, F1: 982, F2: 1}
                _ = v6
            } else {
                if (false && (!false)) {
                    switch v2 {
                    case "\n\t":
                        var v7 int64 = 0
                        _ = v7
                    case "foo bar":
                        var v8 bool = true
                        _ = v8
                    case "a":
                        v9 := []int{1}
                        _ = v9
                    }
                    if true {
                        var v10 uint8 = 192
                        fmt.Println(v10, v10, v2)
                        _ = v10
                    }
                } else {
                    v4 = v2
                    switch (v4 + (v2 + ("foo bar" + ("a" + v2)))) {
                    case "a":
                        v11 := GenS2{F0: 255, F1: 127, F2: 64}
                        _ = v11
                    default:
                        for fi12 := 0; fi12 < 8; fi12++ {
                            v13 := GenS2{F0: 255, F1: -1, F2: 2}
                            _ = fi12
                            _ = v13
                        }
                    }
                }
                fmt.Println(v2)
                if (false && true) {
                    delete(v1, 2147483647)
                    for fi14 := 0; fi14 < 5; fi14++ {
                        var v15 uint16 = 2
                        v3[32767] = ((fi14 + fi14) <= -927)
                        if false {
                            v15 -= (v15 + v15)
                            var v16 int8 = 42
                            _ = v16
                        }
                        _ = fi14
                        _ = v15
                    }
                    v17 := GenS2{F0: 65535, F1: 1, F2: 4294967295}
                    _ = v17
                } else {
                    v4 = v4
                    switch "\n\t" {
                    case "a":
                        _ = len(v3)
                    case "foo bar":
                        var v18 string = "hello"
                        _ = v18
                    case "hello":
                        v3[32767] = false
                    }
                }
            }
        default:
            switch 32768 {
            case -619:
                for fi19 := 0; fi19 < 9; fi19++ {
                    if true {
                        var v20 float32 = -2.25
                        var v21 string = v2
                        _ = v20
                        _ = v21
                    } else {
                        v22 := GenS2{F0: 255, F1: -14, F2: 4294967295}
                        switch v4 {
                        case "":
                            v4 = "hello"
                        case "foo bar":
                            v23 := GenS1{F0: 18446744073709551615, F1: 1, F2: 0}
                            _ = v23
                        case "hello":
                            v2 = ("hello" + "")
                        default:
                            v1[0] = (!true)
                        }
                        _ = len(v1)
                        _ = v22
                    }
                    var v24 uint = 431
                    _ = fi19
                    _ = v24
                }
            case 1:
                var v25 bool = (1e10 <= 0.0)
                _ = v25
            default:
                v4 = "foo bar"
            }
        }
    }
    v2 = v4
    for fi26 := 0; fi26 < 4; fi26++ {
        v27 := GenS3{F0: 2}
        _ = fi26
        _ = v27
    }
    if true {
        switch 0 {
        case -9223372036854775808:
            for fi28 := 0; fi28 < 12; fi28++ {
                if (true || (true && true)) {
                    v29 := []int{(fi28 * 32767), fi28, (-(fi28)), 32767}
                    v2 = v4
                    if (2 >= 497) {
                        v30 := v1[788]
                        var v31 uint32 = 636
                        _ = v30
                        _ = v31
                    } else {
                        var v32 int32 = 1
                        for fi33 := 0; fi33 < 9; fi33++ {
                            v34 := GenS1{F0: 18446744073709551615, F1: 2, F2: -843}
                            _ = fi33
                            _ = v34
                        }
                        _ = v32
                    }
                    _ = v29
                } else {
                    var v35 uint64 = 833
                    if true {
                        v36 := GenS3{F0: 29}
                        fmt.Println(fi28, fi28)
                        for fi37 := 0; fi37 < 11; fi37++ {
                            var v38 float64 = float64(1e10)
                            v39 := []int{}
                            _ = fi37
                            _ = v38
                            _ = v39
                        }
                        _ = v36
                    }
                    v4 = (("" + "hello") + v2)
                    _ = v35
                }
                v2 = "a"
                _ = fi28
            }
        default:
            v40 := GenS2{F0: 2, F1: -90, F2: 391}
            _ = v40
        }
        var v41 string = ("\n\t" + v4)
        _ = v41
    }
    if (!(!(!true))) {
        switch 32767 {
        case 32767:
            var v42 uint = 4294967295
            _ = v42
        }
        for fi43 := 0; fi43 < 8; fi43++ {
            for fi44 := 0; fi44 < 19; fi44++ {
                switch "a" {
                case "foo bar":
                    var v45 uint8 = 0
                    switch fi44 {
                    case -579:
                        v2 = (v2 + (((("hello" + "") + v4) + "hello") + "a"))
                    default:
                        v46 := GenS1{F0: 450, F1: 2, F2: 437}
                        _ = v46
                    }
                    fmt.Println(v45, fi43, fi44)
                    _ = v45
                case "a":
                    var v47 uint16 = 77
                    _ = v47
                case "hello":
                    var v48 int64 = 0
                    _ = v48
                default:
                    var v49 int8 = 1
                    _ = v49
                }
                var v50 int = 779
                var v51 int16 = 348
                switch (-(fi44)) {
                case 127:
                    var v52 uint32 = 0
                    v4 = ("foo bar" + (v2 + (v2 + "hello")))
                    var v53 uint = 244
                    _ = v52
                    _ = v53
                }
                v2 = ""
                var v54 int64 = -9223372036854775808
                _ = fi44
                _ = v50
                _ = v51
                _ = v54
            }
            var v55 float32 = ((((1.5 * 0.0) / 1.0) / 0.1) + float32(0.0))
            for fi56 := 0; fi56 < 1; fi56++ {
                if (!(0 == 42)) {
                    v55 = v55
                    if true {
                        delete(v3, (fi43 << ((0) & 63)))
                        if (true && (false && (((v4 + v2) + v2) >= ""))) {
                            v57 := []int{fi56, fi43}
                            if (9223372036854775807 < 254) {
                                fmt.Println(v4, v55, v4)
                                var v58 bool = (((v2 + "") <= v4) && false)
                                _ = v58
                            } else {
                                if false {
                                    switch "" {
                                    case "hello":
                                        if (false && false) {
                                            v57 = append(v57, 1)
                                            v55 += 3.14159
                                        }
                                    default:
                                        v4 = (v2 + v2)
                                    }
                                    v59 := GenS4{F0: -9223372036854775808, F1: -1, F2: (fi56 ^ (fi43 / ((-2147483648) | 1)))}
                                    var v60 int32 = 1
                                    v55 -= float32(1.5)
                                    var v61 bool = (false && false)
                                    _ = v59
                                    _ = v60
                                    _ = v61
                                } else {
                                    v62 := GenS5{F0: 32767}
                                    _ = v62
                                }
                                v57 = append(v57, (fi56 & fi43))
                                _ = len(v57)
                            }
                            var v63 int8 = 0
                            v64 := GenS2{F0: 18446744073709551615, F1: (^(v63)), F2: 42}
                            _ = v57
                            _ = v63
                            _ = v64
                        }
                    }
                    var v65 int32 = 2147483647
                    _ = v65
                }
                _ = fi56
            }
            _ = fi43
            _ = v55
        }
    } else {
        var v66 uint8 = 127
        if false {
            v67 := GenS6{F0: 32767, F1: 42, F2: 511}
            var v68 bool = (!(-1 <= 84))
            _ = v67
            _ = v68
        } else {
            switch 457 {
            case 894:
                var v69 float64 = (1e10 / 1.0)
                _ = v69
            case -1:
                var v70 int64 = 0
                _ = v70
            case 32767:
                fmt.Println(v66, v66)
            }
        }
        _ = v66
    }
    var v71 uint = 450
    var v72 int16 = 1
    for fi73 := 0; fi73 < 17; fi73++ {
        v74 := GenS7{F0: "", F1: -1, F2: 1}
        _ = fi73
        _ = v74
    }
    var v75 uint16 = 0
    var v76 string = ((v4 + "\n\t") + ("foo bar" + (("hello" + v4) + v2)))
    var v77 uint16 = (v75 | uint16(v71))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v71:", v71)
    fmt.Println("v72:", v72)
    fmt.Println("v75:", v75)
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v71
    _ = v72
    _ = v75
    _ = v76
    _ = v77
}
