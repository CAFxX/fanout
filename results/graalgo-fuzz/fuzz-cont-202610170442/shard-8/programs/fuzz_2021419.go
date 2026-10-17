package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8; F1 int; F2 int }
type GenS4 struct { F0 string; F1 uint16; F2 uint }
type GenS5 struct { F0 string; F1 int16 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint16; F1 bool; F2 int32 }
type GenS9 struct { F0 uint32; F1 bool }
func genFn0() uint8 {
    v6 := []int{9223372036854775807}
    var v7 uint64 = 190
    var v8 uint64 = 4294967295
    for fi9 := 0; fi9 < 4; fi9++ {
        var v10 uint = 255
        _ = fi9
        _ = v10
    }
    _ = v6
    _ = v7
    _ = v8
    return 128
}
func genFn1(p0 bool, p1 bool) int64 {
    var v11 uint16 = 255
    var v12 bool = p0
    var v13 uint16 = 255
    fmt.Println(p1)
    if false {
        for fi14 := 0; fi14 < 3; fi14++ {
            if p0 {
                if false {
                    p1 = p0
                    for fi15 := 0; fi15 < 2; fi15++ {
                        if ((1e10 + 1.5) == 1.5) {
                            v16 := GenS1{F0: 0}
                            v17 := GenS1{F0: 100}
                            if true {
                                if (v11 <= (v13 >> ((101) & 63))) {
                                    switch fi14 {
                                    case 32767:
                                        var v18 int32 = -2147483648
                                        _ = v18
                                    case 128:
                                        v19 := GenS2{F0: 429}
                                        _ = v19
                                    case 2147483647:
                                        v20 := []int{}
                                        _ = v20
                                    }
                                    v21 := GenS1{F0: -118}
                                    p0 = ((true || (168 != -736)) || (true || true))
                                    v22 := []int{}
                                    _ = v21
                                    _ = v22
                                }
                            }
                            v23 := GenS1{F0: -128}
                            var v24 uint64 = 2
                            var v25 float64 = 1e10
                            v26 := []int{}
                            _ = v16
                            _ = v17
                            _ = v23
                            _ = v24
                            _ = v25
                            _ = v26
                        } else {
                            v27 := GenS2{F0: 743}
                            var v28 uint8 = 42
                            _ = v27
                            _ = v28
                        }
                        p0 = p0
                        v29 := []int{}
                        _ = fi15
                        _ = v29
                    }
                    v30 := GenS2{F0: 1}
                    _ = v30
                } else {
                    var v31 float32 = 3.14159
                    var v32 float64 = (((1e10 * 0.1) / -2.25) * -2.25)
                    _ = v31
                    _ = v32
                }
            } else {
                var v33 int16 = 1000
                v34 := []int{(fi14 - -781), (fi14 - fi14)}
                _ = v33
                _ = v34
            }
            var v35 string = (("a" + "") + ("foo bar" + "\n\t"))
            _ = fi14
            _ = v35
        }
        v36 := GenS3{F0: 11, F1: 412, F2: -2147483648}
        _ = v36
    } else {
        fmt.Println(v12, v11, p1)
        fmt.Println(v12, v13, v13)
    }
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v13
    return 551
}

func main() {
    v1 := []uint32{0}
    var v2 int = -11
    v3 := []uint32{4294967295}
    var v4 int32 = -702
    v5 := make(map[uint32]bool)
    for fi37 := 0; fi37 < 11; fi37++ {
        switch "hello" {
        case "":
            switch ("a" + "\n\t") {
            case "a":
                for fi38 := 0; fi38 < 6; fi38++ {
                    if false {
                        var v39 int32 = (v4 * (v4 / (((v4 >> ((20) & 63))) | 1)))
                        _ = v39
                    } else {
                        if len(v1) > 0 {
                            v1[0] = 295
                        }
                        v40 := GenS4{F0: ("foo bar" + ""), F1: 32767, F2: 738}
                        _ = v40
                    }
                    _ = fi38
                }
            case "foo bar":
                switch fi37 {
                case -222:
                    switch (^(fi37)) {
                    case 2147483647:
                        v2 += v2
                    case -9223372036854775808:
                        var v41 int32 = (v4 >> ((2) & 63))
                        _ = v41
                    default:
                        fmt.Println(v4)
                    }
                case -881:
                    switch v2 {
                    case -2147483648:
                        if (!(4294967295 == 4294967295)) {
                            v4 = 1
                            if true {
                                v4 -= (v4 + v4)
                            }
                            fmt.Println(v2, v4, v2)
                        }
                    case 4294967295:
                        if true {
                            v4 = 42
                        } else {
                            if (0 <= 127) {
                                var v42 bool = (!(genFn1((-2.25 >= (-2.25 * 0.0)), (255 <= 18446744073709551615)) > genFn1(true, (!false))))
                                _ = v42
                            }
                        }
                    case 0:
                        v43 := GenS5{F0: "a", F1: 42}
                        _ = v43
                    }
                }
            default:
                fmt.Println(fi37, v4)
            }
        case "a":
            v4 = v4
        case "\n\t":
            v2 = 2147483647
        }
        var v44 int16 = 1
        v2 -= v2
        _ = fi37
        _ = v44
    }
    switch "a" {
    case "":
        switch (v2 + 1) {
        case 9223372036854775807:
            switch ("\n\t" + "a") {
            case "a":
                fmt.Println(v2, v4)
            case "\n\t":
                fmt.Println(v4, v2)
            case "":
                var v45 int16 = -1
                _ = v45
            default:
                var v46 uint8 = 2
                _ = v46
            }
        case 128:
            v2 = v2
        }
    case "a":
        switch v2 {
        case -2147483648:
            var v47 int = 2147483647
            _ = v47
        }
    case "\n\t":
        switch "\n\t" {
        case "foo bar":
            v2 = v2
        case "\n\t":
            var v48 string = "foo bar"
            _ = v48
        case "hello":
            v49 := GenS6{F0: 32767}
            _ = v49
        default:
            for fi50 := 0; fi50 < 13; fi50++ {
                v2 += fi50
                var v51 float32 = -2.25
                if true {
                    v52 := GenS6{F0: 32767}
                    _ = v52
                } else {
                    switch "foo bar" {
                    case "":
                        for fi53 := 0; fi53 < 16; fi53++ {
                            fmt.Println(v2)
                            switch (fi50 * (v2 >> ((1) & 63))) {
                            case 4294967295:
                                if false {
                                    v54 := GenS7{F0: false}
                                    _ = v54
                                }
                            }
                            _ = fi53
                        }
                    case "a":
                        v4 = v4
                    case "\n\t":
                        var v55 string = "foo bar"
                        _ = v55
                    }
                }
                _ = fi50
                _ = v51
            }
        }
    }
    v2 = v2
    var v56 uint64 = 2
    if true {
        v57 := v5[753]
        _ = len(v3)
        _ = v57
    }
    delete(v5, 2147483648)
    var v58 uint8 = 1
    for fi59 := 0; fi59 < 20; fi59++ {
        if false {
            var v60 uint = 132
            for fi61 := 0; fi61 < 16; fi61++ {
                switch int(v56) {
                case 9223372036854775807:
                    var v62 uint64 = v56
                    _ = v62
                case 670:
                    fmt.Println(v4, v60, v60)
                case -1:
                    var v63 int32 = v4
                    _ = v63
                default:
                    v64 := GenS6{F0: 65535}
                    _ = v64
                }
                v65 := GenS7{F0: (uint8(v56) < 42)}
                v4 = v4
                fmt.Println(fi61, v2, v56)
                if true {
                    v3 = append(v3, 2)
                } else {
                    if false {
                        v4 = (v4 - 1)
                        var v66 float32 = 1.5
                        v67 := GenS6{F0: 65535}
                        _ = v66
                        _ = v67
                    }
                    var v68 int = (v2 - fi61)
                    var v69 float64 = 1.5
                    _ = v68
                    _ = v69
                }
                _ = fi61
                _ = v65
            }
            if ((1.5 <= 1.5) && ("hello" == "")) {
                v70 := GenS1{F0: 100}
                _ = v70
            }
            _ = v60
        } else {
            if false {
                v56 = 4294967295
                switch v2 {
                case 533:
                    if (18446744073709551615 > 254) {
                        for fi71 := 0; fi71 < 17; fi71++ {
                            var v72 int8 = -1
                            v73 := GenS1{F0: v72}
                            _ = fi71
                            _ = v72
                            _ = v73
                        }
                    } else {
                        v74 := GenS8{F0: 42, F1: false, F2: (v4 << ((390) & 63))}
                        _ = v74
                    }
                }
                for fi75 := 0; fi75 < 2; fi75++ {
                    fmt.Println(fi75, v2)
                    fmt.Println(fi75)
                    _ = fi75
                }
            }
            for fi76 := 0; fi76 < 13; fi76++ {
                var v77 bool = true
                fmt.Println(v56)
                var v78 int = (fi59 / ((fi76) | 1))
                _ = fi76
                _ = v77
                _ = v78
            }
        }
        for fi79 := 0; fi79 < 18; fi79++ {
            v80 := GenS3{F0: -108, F1: v2, F2: fi79}
            v3 = append(v3, 705)
            _ = fi79
            _ = v80
        }
        v4 = v4
        _ = fi59
    }
    v81 := GenS7{F0: true}
    switch "" {
    case "hello":
        fmt.Println(v2)
    default:
        var v82 uint64 = 826
        _ = v82
    }
    v83 := GenS9{F0: 406, F1: false}
    v58 += 2
    for fi84 := 0; fi84 < 12; fi84++ {
        v56 = (v56 % ((v56) | 1))
        v2 += -933
        _ = fi84
    }
    var v85 uint8 = (v58 & (v58 - genFn0()))
    var v86 int64 = 42
    var v87 int = 128
    for fi88 := 0; fi88 < 20; fi88++ {
        var v89 int64 = (^(v86))
        fmt.Println(v87)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi88
        _ = v89
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v56:", v56)
    fmt.Println("v58:", v58)
    fmt.Println("v85:", v85)
    fmt.Println("v86:", v86)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v58
    _ = v81
    _ = v83
    _ = v85
    _ = v86
    _ = v87
}
