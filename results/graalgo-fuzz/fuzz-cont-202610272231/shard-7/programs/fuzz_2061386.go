package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 int; F1 uint8 }
type GenS4 struct { F0 bool; F1 int8 }
type GenS5 struct { F0 string; F1 int64; F2 int64 }
type GenS6 struct { F0 bool; F1 uint64; F2 int64 }
type GenS7 struct { F0 uint; F1 bool }
type GenS8 struct { F0 int32; F1 string }

func main() {
    v1 := []uint32{2}
    var v2 float64 = 0.0
    v3 := []int{42, 2147483647, 239}
    if false {
        var v4 uint32 = 42
        if (-1 < 1) {
            fmt.Println(v4, v4)
        } else {
            if (!("hello" <= "")) {
                v2 -= 1.5
                switch "" {
                case "hello":
                    if false {
                        switch "a" {
                        case "\n\t":
                            fmt.Println(v4, v2)
                        default:
                            var v5 uint = 255
                            _ = v5
                        }
                        var v6 uint = 642
                        _ = v6
                    }
                }
                v4 -= v4
            } else {
                if true {
                    if (822 > 127) {
                        if false {
                            v4 = (v4 + 213)
                            var v7 string = ("hello" + ("\n\t" + "a"))
                            _ = v7
                        } else {
                            var v8 int32 = -2147483648
                            var v9 float64 = (v2 + (-2.25 - 3.14159))
                            _ = v8
                            _ = v9
                        }
                    } else {
                        if (255 != 255) {
                            for fi10 := 0; fi10 < 14; fi10++ {
                                fmt.Println(fi10, v4, v4)
                                for fi11 := 0; fi11 < 9; fi11++ {
                                    v4 = v4
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    if true {
                                        var v12 float32 = 1.5
                                        var v13 uint16 = 255
                                        if false {
                                            var v14 int8 = -1
                                            v12 = (0.1 / (v12 - (v12 * 3.14159)))
                                            _ = v14
                                        } else {
                                            v13 -= v13
                                        }
                                        _ = v12
                                        _ = v13
                                    } else {
                                        v15 := GenS1{F0: fi10}
                                        _ = v15
                                    }
                                    _ = fi11
                                }
                                for fi16 := 0; fi16 < 5; fi16++ {
                                    v4 = v4
                                    var v17 uint64 = 2
                                    if false {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        v2 = v2
                                        switch fi16 {
                                        case 2147483647:
                                            var v18 uint32 = 1
                                            _ = v18
                                        default:
                                            switch "a" {
                                            case "":
                                                var v19 int32 = 1
                                                _ = v19
                                            case "hello":
                                                v17 -= (^(v17))
                                            default:
                                                var v20 int16 = 1000
                                                v17 += (v17 + 954)
                                                _ = v20
                                            }
                                        }
                                    } else {
                                        var v21 float64 = float64(((0.0 / 1.5) * (0.0 * 1e10)))
                                        v21 += ((-2.25 + 1e10) + v21)
                                        v3 = append(v3, fi16)
                                        _ = v21
                                    }
                                    v2 = v2
                                    _ = fi16
                                    _ = v17
                                }
                                _ = fi10
                            }
                            var v22 int16 = -33
                            _ = v22
                        }
                        var v23 uint8 = 128
                        _ = v23
                    }
                    v24 := GenS1{F0: 1}
                    for fi25 := 0; fi25 < 1; fi25++ {
                        if (128 != 1) {
                            if true {
                                v26 := GenS1{F0: -539}
                                fmt.Println(v2, v2)
                                _ = v26
                            } else {
                                fmt.Println(v4, v2, v4)
                                var v27 int32 = 0
                                _ = v27
                            }
                            if ((65535 > 2) && (!true)) {
                                if (!true) {
                                    if false {
                                        fmt.Println(fi25, v4, v2)
                                    }
                                    var v28 int32 = 0
                                    v2 += v2
                                    _ = v28
                                }
                                fmt.Println(v4, v4)
                            } else {
                                switch 127 {
                                case 32768:
                                    switch (fi25 * fi25) {
                                    case 815:
                                        v4 = (v4 ^ v4)
                                    case 2147483647:
                                        var v29 int8 = 25
                                        _ = v29
                                    case 845:
                                        var v30 int8 = -18
                                        _ = v30
                                    }
                                case -641:
                                    v4 += v4
                                }
                            }
                            var v31 int64 = 42
                            _ = v31
                        }
                        var v32 int32 = 825
                        _ = fi25
                        _ = v32
                    }
                    _ = v24
                } else {
                    v4 = 2147483647
                    if true {
                        v2 += v2
                        var v33 string = "hello"
                        _ = v33
                    }
                    var v34 uint32 = v4
                    _ = v34
                }
                v35 := GenS1{F0: -2147483648}
                _ = v35
            }
            fmt.Println(v2)
            switch 127 {
            case 1:
                v2 = 0.1
            default:
                var v36 int = 9223372036854775807
                _ = v36
            }
        }
        for fi37 := 0; fi37 < 10; fi37++ {
            var v38 float64 = -2.25
            _ = fi37
            _ = v38
        }
        _ = v4
    }
    v2 = ((v2 * 3.14159) - 1.5)
    switch "hello" {
    case "":
        v39 := GenS1{F0: 32768}
        _ = v39
    case "a":
        for fi40 := 0; fi40 < 16; fi40++ {
            switch "hello" {
            case "foo bar":
                var v41 uint16 = 968
                _ = v41
            case "a":
                for fi42 := 0; fi42 < 20; fi42++ {
                    v2 -= 0.1
                    _ = fi42
                }
            case "hello":
                switch 1 {
                case 127:
                    v2 = v2
                case 32767:
                    for fi43 := 0; fi43 < 16; fi43++ {
                        v2 = ((1e10 / 1.0) - v2)
                        var v44 uint16 = 101
                        var v45 uint64 = 1
                        _ = fi43
                        _ = v44
                        _ = v45
                    }
                case -1:
                    var v46 float32 = -2.25
                    _ = v46
                }
            default:
                switch ("foo bar" + ("" + "a")) {
                case "\n\t":
                    var v47 string = "foo bar"
                    _ = v47
                case "foo bar":
                    var v48 uint16 = 255
                    _ = v48
                default:
                    v2 = (0.1 + v2)
                }
            }
            v49 := GenS1{F0: (^(fi40))}
            if false {
                for fi50 := 0; fi50 < 16; fi50++ {
                    v2 = 1.5
                    v3 = append(v3, (^(fi50)))
                    _ = fi50
                }
                var v51 uint16 = 880
                _ = v51
            } else {
                switch fi40 {
                case -1:
                    v2 -= float64(3.14159)
                default:
                    v52 := GenS2{F0: 42, F1: 665}
                    _ = v52
                }
                if true {
                    var v53 int = 2147483647
                    switch (^(fi40)) {
                    case 9223372036854775807:
                        v54 := GenS3{F0: (v53 / ((v53) | 1)), F1: 120}
                        _ = v54
                    default:
                        switch v53 {
                        case 1:
                            var v55 int8 = 127
                            _ = v55
                        case -578:
                            if true {
                                fmt.Println(v53, fi40, fi40)
                                if true {
                                    if true {
                                        if (true || true) {
                                            fmt.Println(v53)
                                            v53 += v53
                                            v56 := GenS2{F0: 0, F1: 1}
                                            _ = v56
                                        }
                                        if (v2 < v2) {
                                            v57 := GenS4{F0: true, F1: 100}
                                            var v58 int8 = 119
                                            _ = v57
                                            _ = v58
                                        } else {
                                            var v59 string = "foo bar"
                                            v60 := GenS4{F0: true, F1: 0}
                                            for fi61 := 0; fi61 < 12; fi61++ {
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                _ = fi61
                                            }
                                            _ = v59
                                            _ = v60
                                        }
                                    } else {
                                        var v62 uint32 = 2
                                        var v63 int64 = 0
                                        _ = v62
                                        _ = v63
                                    }
                                } else {
                                    v64 := GenS5{F0: "", F1: -375, F2: 0}
                                    if true {
                                        for fi65 := 0; fi65 < 14; fi65++ {
                                            switch ("hello" + "a") {
                                            case "hello":
                                                v66 := GenS6{F0: false, F1: 745, F2: -1}
                                                _ = v66
                                            case "":
                                                var v67 bool = false
                                                _ = v67
                                            case "foo bar":
                                                var v68 uint8 = 190
                                                _ = v68
                                            }
                                            _ = fi65
                                        }
                                    } else {
                                        v3 = append(v3, (v53 + fi40))
                                        v69 := GenS7{F0: 2, F1: true}
                                        _ = v69
                                    }
                                    _ = v64
                                }
                                var v70 uint32 = 0
                                _ = v70
                            } else {
                                var v71 uint64 = 2
                                _ = v71
                            }
                        default:
                            var v72 int64 = -9223372036854775808
                            _ = v72
                        }
                    }
                    var v73 uint64 = 42
                    _ = v53
                    _ = v73
                }
                for fi74 := 0; fi74 < 8; fi74++ {
                    v75 := GenS7{F0: 261, F1: (1 <= 2)}
                    fmt.Println(fi74, fi74, fi74)
                    _ = fi74
                    _ = v75
                }
            }
            _ = fi40
            _ = v49
        }
    default:
        var v76 int16 = 557
        _ = v76
    }
    v2 -= (v2 / (((((float64(1e10) / (v2 / v2)) * ((((v2 * v2) / v2) * v2) + (float64(((-2.25 + (3.14159 / 1.0)) - 1.5)) + v2))) - ((((v2 / ((v2 - float64(-2.25)) - v2)) - 1.5) + ((0.1 / v2) + float64(1.5))) * v2)) - v2) + 1e10))
    fmt.Println(v2)
    if len(v1) > 0 {
        v1[0] = 609
    }
    v77 := GenS8{F0: -1, F1: ("hello" + "a")}
    var v78 float64 = -2.25
    v79 := GenS4{F0: ((1.5 + (((((3.14159 - 0.0) * float32(0.1)) - -2.25) - (0.1 - 1.5)) / 1.0)) > 1.5), F1: 100}
    v2 += -2.25
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v80 string = (("foo bar" + ("a" + ("a" + "foo bar"))) + "")
    if false {
        var v81 float32 = (1e10 / 1.5)
        var v82 int8 = -116
        _ = v81
        _ = v82
    }
    v78 -= (v2 - v78)
    v83 := GenS5{F0: ((v80 + (v80 + v80)) + "\n\t"), F1: 525, F2: 1}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if (!(v2 != ((1e10 / float64((1.5 - ((((-2.25 * 0.0) / 3.14159) / 0.1) - 0.1)))) / 1.5))) {
        var v84 int8 = 127
        _ = v84
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v78:", v78)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v3
    _ = v77
    _ = v78
    _ = v79
    _ = v80
    _ = v83
}
