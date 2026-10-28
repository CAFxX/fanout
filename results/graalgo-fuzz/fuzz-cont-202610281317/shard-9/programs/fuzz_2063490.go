package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint16; F1 int }
type GenS5 struct { F0 bool; F1 uint8 }
type GenS6 struct { F0 int16; F1 int16; F2 uint8 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int16 }
func genFn0(p0 int8) uint8 {
    var v6 int32 = 811
    v7 := []int{-2147483648, -189}
    _ = p0
    _ = v6
    _ = v7
    return 2
}
func genFn1() uint16 {
    if true {
        var v8 bool = (-1 >= -128)
        v8 = true
        _ = v8
    } else {
        v9 := []int{32767, 32767, 751}
        var v10 string = "a"
        var v11 uint8 = 42
        _ = v10
        _ = v11
        _ = v9
    }
    for fi12 := 0; fi12 < 2; fi12++ {
        v13 := GenS1{F0: 42, F1: 0}
        v14 := []int{-1}
        fmt.Println(fi12)
        _ = fi12
        _ = v13
        _ = v14
    }
    var v15 float64 = 3.14159
    _ = v15
    return 688
}

func main() {
    var v1 uint64 = 585
    var v2 bool = false
    var v3 int32 = -750
    v4 := make(map[string]int)
    var v5 uint16 = 255
    for fi16 := 0; fi16 < 14; fi16++ {
        v17 := GenS1{F0: v1, F1: 0}
        _ = fi16
        _ = v17
    }
    v18 := GenS2{F0: 206}
    if v2 {
        if (100 >= 0) {
            v3 = (v3 >> ((4294967295) & 63))
            v19 := []int{32767, -669, -2147483648, 127}
            _ = v19
        }
    } else {
        v20 := GenS2{F0: 105}
        v3 -= 86
        _ = v20
    }
    var v21 float64 = -2.25
    switch "foo bar" {
    case "\n\t":
        var v22 int64 = 0
        _ = v22
    case "a":
        delete(v4, (("hello" + "") + "foo bar"))
    default:
        var v23 uint8 = 127
        _ = v23
    }
    switch 32768 {
    case 127:
        var v24 uint32 = 0
        _ = v24
    case 2147483647:
        if (1 <= 2) {
            var v25 float32 = 3.14159
            fmt.Println(v3, v3, v5)
            _ = v25
        } else {
            for fi26 := 0; fi26 < 15; fi26++ {
                switch 2147483647 {
                case -1:
                    v2 = false
                case 4294967295:
                    var v27 int = fi26
                    _ = v27
                default:
                    switch "\n\t" {
                    case "a":
                        if true {
                            var v28 float32 = 3.14159
                            for fi29 := 0; fi29 < 18; fi29++ {
                                delete(v4, ("foo bar" + "\n\t"))
                                v21 = v21
                                v21 = (v21 + v21)
                                _ = fi29
                            }
                            _ = v28
                        }
                    case "\n\t":
                        if (1e10 >= 0.0) {
                            var v30 float32 = 1.5
                            if (!(4294967295 <= 0)) {
                                var v31 int64 = 1
                                if (uint64(v5) > (v1 ^ v1)) {
                                    var v32 int64 = (^(v31))
                                    _ = len(v4)
                                    _ = v32
                                } else {
                                    v1 += v1
                                }
                                _ = v31
                            } else {
                                v33 := []int{}
                                _ = v33
                            }
                            fmt.Println(fi26, v5, v2)
                            _ = v30
                        }
                    }
                }
                v34 := GenS2{F0: 42}
                v35 := []int{-602}
                _ = fi26
                _ = v34
                _ = v35
            }
            fmt.Println(v2, v1, v5)
        }
    }
    var v36 float32 = 0.1
    v37 := []int{49, 9223372036854775807, 127}
    v38 := GenS1{F0: v1, F1: 42}
    fmt.Println(v5, v2)
    v21 = (v21 - v21)
    if false {
        switch 32768 {
        case 702:
            v39 := GenS1{F0: 2, F1: -344}
            _ = v39
        case 99:
            if (((float32((v21 + float64(1e10))) / 1.0) == v36) || ((9223372036854775807 < 128) && (-2147483648 < v3))) {
                v40 := GenS2{F0: 56}
                v37 = append(v37, 2147483647)
                _ = v40
            } else {
                v41 := GenS2{F0: 209}
                if len(v37) > 0 {
                    v37[0] = 42
                }
                _ = v41
            }
        default:
            switch 2147483647 {
            case 1:
                fmt.Println(v1)
            case 2147483647:
                switch -2147483648 {
                case 4294967295:
                    var v42 int = 127
                    _ = v42
                case 144:
                    fmt.Println(v36)
                default:
                    if ((^(v1)) > 0) {
                        v43 := GenS1{F0: 1, F1: -32768}
                        _ = v43
                    }
                }
            case -1:
                if (!(0 != -9)) {
                    var v44 int64 = -9223372036854775808
                    if ((((float32(3.14159) / (v36 - (float32(0.0) - v36))) + v36) / v36) == (v36 + (v36 - v36))) {
                        switch ((("hello" + ((("hello" + (("\n\t" + "\n\t") + ("" + ("\n\t" + (("foo bar" + "") + ("\n\t" + ("foo bar" + "a"))))))) + ("hello" + ((("hello" + "a") + "") + "hello"))) + "a")) + (("hello" + ("hello" + "\n\t")) + "\n\t")) + (("a" + ("" + ("\n\t" + (("foo bar" + "") + (("\n\t" + ("a" + "\n\t")) + (("foo bar" + "\n\t") + "\n\t")))))) + ("hello" + "\n\t"))) {
                        case "a":
                            v45 := GenS3{F0: 0}
                            _ = v45
                        default:
                            v37 = append(v37, 4294967295)
                        }
                        if ("foo bar" > "") {
                            var v46 string = "\n\t"
                            if v2 {
                                var v47 float32 = 1e10
                                switch v46 {
                                case "foo bar":
                                    for fi48 := 0; fi48 < 4; fi48++ {
                                        var v49 uint = 18446744073709551615
                                        fmt.Println(v36, v49, v47)
                                        var v50 int8 = 100
                                        _ = fi48
                                        _ = v49
                                        _ = v50
                                    }
                                default:
                                    if true {
                                        if (float32(v21) > 3.14159) {
                                            v44 = -161
                                            var v51 int64 = v44
                                            _ = v51
                                        } else {
                                            v21 = ((3.14159 * v21) - v21)
                                            v52 := GenS2{F0: 204}
                                            _ = v52
                                        }
                                        v53 := GenS4{F0: v5, F1: 301}
                                        _ = v53
                                    } else {
                                        var v54 bool = v2
                                        _ = v54
                                    }
                                }
                                _ = v47
                            }
                            if false {
                                switch -547 {
                                case 0:
                                    var v55 int64 = (v44 | v44)
                                    _ = v55
                                case 128:
                                    if ((v2 && (3.14159 > v21)) && false) {
                                        if (0 == 32767) {
                                            for fi56 := 0; fi56 < 3; fi56++ {
                                                switch "a" {
                                                case "\n\t":
                                                    fmt.Println(v21, v21, v21)
                                                }
                                                v57 := GenS2{F0: 128}
                                                _ = fi56
                                                _ = v57
                                            }
                                            var v58 uint = 1
                                            _ = v58
                                        }
                                    }
                                }
                            } else {
                                if len(v37) >= 2 {
                                    v37 = v37[1:len(v37)]
                                }
                            }
                            _ = v46
                        }
                    }
                    _ = v44
                } else {
                    if false {
                        v59 := GenS5{F0: v2, F1: 103}
                        _ = v59
                    } else {
                        switch "\n\t" {
                        case "":
                            var v60 int = 0
                            _ = v60
                        case "a":
                            for fi61 := 0; fi61 < 15; fi61++ {
                                v2 = true
                                _ = fi61
                            }
                        default:
                            if len(v37) >= 2 {
                                v37 = v37[1:len(v37)]
                            }
                        }
                        v2 = (v36 > (0.1 - float32((((v21 + v21) + ((float64(v36) / float64(v36)) / 1e10)) - v21))))
                    }
                    v62 := GenS4{F0: v5, F1: 4294967295}
                    _ = v62
                }
            default:
                v5 = (v5 + v5)
            }
        }
    } else {
        for fi63 := 0; fi63 < 5; fi63++ {
            v64 := GenS4{F0: v5, F1: fi63}
            _ = fi63
            _ = v64
        }
    }
    var v65 bool = (("a" + "hello") > ("" + "foo bar"))
    switch "foo bar" {
    case "hello":
        if v2 {
            fmt.Println(v2, v5, v36)
        } else {
            v4["a"] = -1
        }
    case "a":
        v66 := GenS5{F0: v65, F1: 128}
        _ = v66
    case "foo bar":
        if (!v65) {
            for fi67 := 0; fi67 < 20; fi67++ {
                v37 = append(v37, 1)
                var v68 bool = false
                for fi69 := 0; fi69 < 6; fi69++ {
                    v70 := GenS5{F0: (!(32767 < int16(v5))), F1: 128}
                    v71 := GenS3{F0: 65535}
                    v1 -= 420
                    v68 = (0 != (v3 / (((-(v3))) | 1)))
                    v2 = (true && true)
                    if len(v37) > 0 {
                        v37[0] = (-(fi69))
                    }
                    var v72 uint16 = (^(v5))
                    _ = fi69
                    _ = v70
                    _ = v71
                    _ = v72
                }
                _ = fi67
                _ = v68
            }
        }
    }
    for fi73 := 0; fi73 < 1; fi73++ {
        v74 := GenS6{F0: -1, F1: 32767, F2: 255}
        if v2 {
            v75 := v4["foo bar"]
            v76 := GenS7{F0: 4294967295}
            _ = v75
            _ = v76
        } else {
            v1 += 943
            if ((557 != 23) || v65) {
                if false {
                    v36 += (v36 - v36)
                    switch ("hello" + ((("foo bar" + "hello") + "\n\t") + ("a" + ("foo bar" + "a")))) {
                    case "foo bar":
                        switch "foo bar" {
                        case "a":
                            fmt.Println(v65, v2, v1)
                        case "":
                            v77 := GenS8{F0: ("foo bar" >= (("a" + "foo bar") + "foo bar"))}
                            _ = v77
                        }
                    case "hello":
                        var v78 float64 = ((v21 * v21) * (v21 - (-2.25 * float64(v36))))
                        _ = v78
                    case "\n\t":
                        v1 += 830
                    default:
                        v65 = true
                    }
                    var v79 uint = 255
                    _ = v79
                } else {
                    v65 = (2 < 1)
                }
            } else {
                var v80 int32 = (v3 >> ((411) & 63))
                for fi81 := 0; fi81 < 15; fi81++ {
                    switch 9223372036854775807 {
                    case 127:
                        if v2 {
                            fmt.Println(v3, v21, v2)
                            if v2 {
                                v80 = (v3 / (((v3 ^ (v3 << ((4294967295) & 63)))) | 1))
                            }
                        } else {
                            if (1 < 1) {
                                v3 = v80
                                fmt.Println(v80, v5, fi81)
                                if len(v37) >= 2 {
                                    v37 = v37[1:len(v37)]
                                }
                            } else {
                                if false {
                                    switch fi81 {
                                    case 824:
                                        var v82 float64 = v21
                                        _ = v82
                                    }
                                    var v83 uint = 2
                                    if (!(v2 && (((v36 + 0.1) * 0.0) > 0.1))) {
                                        v1 -= v1
                                    }
                                    _ = v83
                                }
                                var v84 uint16 = (v5 << ((4294967295) & 63))
                                _ = v84
                            }
                            if v65 {
                                v80 -= v80
                                var v85 uint64 = 526
                                _ = v85
                            } else {
                                fmt.Println(v3, v21)
                                v21 += 1.5
                            }
                        }
                    }
                    _ = fi81
                }
                for fi86 := 0; fi86 < 19; fi86++ {
                    if (!(-1 == fi73)) {
                        v87 := GenS9{F0: 790}
                        fmt.Println(fi73, fi86)
                        _ = len(v37)
                        _ = len(v37)
                        fmt.Println(fi73, v36)
                        _ = v87
                    } else {
                        if v2 {
                            v65 = v65
                            v5 -= 32767
                        } else {
                            switch "foo bar" {
                            case "foo bar":
                                var v88 uint64 = uint64(v5)
                                _ = v88
                            case "hello":
                                fmt.Println(v5)
                            }
                            fmt.Println(v80, v1)
                            v2 = (v36 >= v36)
                        }
                        if v65 {
                            if (!false) {
                                v2 = (1 >= 240)
                                if (!v65) {
                                    v3 = (v3 << ((861) & 63))
                                    v3 += v80
                                }
                            }
                            v5 = v5
                        } else {
                            var v89 uint16 = 42
                            _ = v89
                        }
                        var v90 float32 = 3.14159
                        _ = v90
                    }
                    if len(v37) >= 2 {
                        v37 = v37[1:len(v37)]
                    }
                    var v91 string = "hello"
                    _ = fi86
                    _ = v91
                }
                _ = v80
            }
        }
        _ = fi73
        _ = v74
    }
    fmt.Println(v5, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v65
}
