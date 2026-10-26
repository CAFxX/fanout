package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 string; F1 int64; F2 int32 }
type GenS3 struct { F0 int32; F1 uint16; F2 uint32 }
type GenS4 struct { F0 int32; F1 uint16; F2 int16 }
type GenS5 struct { F0 int16; F1 int32; F2 int }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint8 }

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[int]bool)
    v3 := make(map[string]int)
    v4 := make(map[string]bool)
    var v5 int8 = -53
    if false {
        if (4294967295 != 42) {
            var v6 int8 = v5
            switch ("hello" + ("" + "hello")) {
            case "":
                if (v1 > (v1 >> ((0) & 63))) {
                    fmt.Println(v5)
                    if ((((float32(1.5) * (float32((0.1 / 3.14159)) * 0.0)) + 1e10) - (float32(0.0) / (3.14159 * 1.5))) < -2.25) {
                        var v7 uint8 = 2
                        _ = v7
                    }
                }
            case "\n\t":
                v6 -= (v6 % ((84) | 1))
            }
            _ = v6
        } else {
            v8 := []int{-1, 679}
            if (false && true) {
                switch "foo bar" {
                case "a":
                    if false {
                        v1 = v1
                    } else {
                        delete(v2, -985)
                        v1 -= v1
                        if false {
                            v5 = (^(v5))
                            v1 = v1
                            switch "a" {
                            case "a":
                                var v9 int16 = 42
                                _ = v9
                            default:
                                switch int(v1) {
                                case -665:
                                    if len(v8) > 0 {
                                        v8[0] = 9223372036854775807
                                    }
                                default:
                                    fmt.Println(v1)
                                }
                            }
                        }
                    }
                case "foo bar":
                    for fi10 := 0; fi10 < 4; fi10++ {
                        for fi11 := 0; fi11 < 10; fi11++ {
                            var v12 int16 = 973
                            var v13 uint8 = 29
                            _ = fi11
                            _ = v12
                            _ = v13
                        }
                        v1 -= v1
                        var v14 int16 = 42
                        _ = fi10
                        _ = v14
                    }
                default:
                    var v15 uint = 65535
                    _ = v15
                }
                fmt.Println(v5, v5)
                switch 128 {
                case -1:
                    if true {
                        var v16 bool = (!true)
                        fmt.Println(v1, v1)
                        _ = v16
                    } else {
                        v1 -= 2147483647
                        v1 = int32(v5)
                        var v17 int64 = 42
                        _ = v17
                    }
                default:
                    v1 -= (-(v1))
                }
            } else {
                var v18 uint8 = 255
                for fi19 := 0; fi19 < 13; fi19++ {
                    var v20 uint32 = 2147483647
                    _ = fi19
                    _ = v20
                }
                switch ("hello" + ("foo bar" + (((("a" + (("hello" + "\n\t") + "hello")) + "\n\t") + "\n\t") + "\n\t"))) {
                case "a":
                    if ((((("\n\t" + "") + "foo bar") + "hello") + "hello") < ("a" + ("foo bar" + "hello"))) {
                        if (true && true) {
                            var v21 int32 = (v1 << ((42) & 63))
                            _ = v21
                        } else {
                            v22 := GenS1{F0: v1, F1: ((!false) && ((!true) || (!true)))}
                            _ = v22
                        }
                    }
                default:
                    v5 += v5
                }
                _ = v18
            }
            _ = v8
        }
        v23 := GenS1{F0: v1, F1: ("foo bar" < "")}
        _ = v23
    } else {
        switch 0 {
        case 1:
            v24 := GenS2{F0: "\n\t", F1: 9223372036854775807, F2: -2147483648}
            _ = v24
        default:
            fmt.Println(v5)
        }
        v5 -= (v5 >> ((255) & 63))
        v25 := GenS1{F0: 2147483647, F1: false}
        _ = v25
    }
    for fi26 := 0; fi26 < 7; fi26++ {
        for fi27 := 0; fi27 < 8; fi27++ {
            v28 := GenS1{F0: (v1 % ((v1) | 1)), F1: ((float32(1e10) + 3.14159) > 0.1)}
            _ = fi27
            _ = v28
        }
        v29 := GenS2{F0: "", F1: -862, F2: v1}
        fmt.Println(v1, fi26)
        _ = fi26
        _ = v29
    }
    if false {
        v30 := []int{42, 1, -766, 4294967295}
        fmt.Println(v1, v5)
        _ = v30
    } else {
        var v31 int16 = 690
        var v32 float64 = (1e10 + float64(3.14159))
        _ = v31
        _ = v32
    }
    switch 268 {
    case -27:
        if true {
            fmt.Println(v1, v1)
            if true {
                fmt.Println(v1, v5)
                v1 = 636
                var v33 int = 128
                _ = v33
            }
            if ((!true) && (float32(0.0) != 1.5)) {
                var v34 uint = 199
                _ = v34
            } else {
                v1 = (v1 * (v1 << ((4294967295) & 63)))
            }
        }
    }
    switch ("\n\t" + "a") {
    case "a":
        fmt.Println(v1)
    case "foo bar":
        v5 -= 64
    }
    if false {
        var v35 int = -9223372036854775808
        _ = v35
    } else {
        if true {
            v5 -= (^(v5))
        } else {
            v36 := GenS3{F0: (v1 << ((4294967295) & 63)), F1: 843, F2: 4294967295}
            v37 := GenS2{F0: "\n\t", F1: 1, F2: v1}
            _ = v36
            _ = v37
        }
        var v38 float64 = 0.1
        _ = v38
    }
    if false {
        if (-2.25 != (3.14159 * -2.25)) {
            var v39 uint64 = 2
            v40 := GenS4{F0: (v1 + v1), F1: 0, F2: 653}
            _ = v39
            _ = v40
        } else {
            switch -714 {
            case 128:
                for fi41 := 0; fi41 < 20; fi41++ {
                    if (42 != 133) {
                        v42 := GenS4{F0: (v1 | v1), F1: 957, F2: 1}
                        _ = v42
                    }
                    v43 := []int{}
                    if len(v43) >= 2 {
                        v43 = v43[1:len(v43)]
                    }
                    _ = fi41
                    _ = v43
                }
            case 127:
                v44 := []int{-9223372036854775808, 128, 32768}
                _ = v44
            default:
                var v45 uint32 = 1
                _ = v45
            }
        }
        var v46 uint64 = 1
        switch 4294967295 {
        case 9223372036854775807:
            v1 = (-(v1))
        case -335:
            fmt.Println(v1, v5)
        case 32768:
            if false {
                var v47 bool = ((822 == 2) || (2147483647 <= 2))
                var v48 uint = 455
                v5 = v5
                _ = v47
                _ = v48
            }
        }
        _ = v46
    } else {
        v49 := GenS2{F0: "hello", F1: 9223372036854775807, F2: v1}
        if (128 <= 42) {
            v3["\n\t"] = -1
            v50 := []int{128, 32768, -332, 2147483647}
            v5 -= (v5 % (((v5 >> ((2) & 63))) | 1))
            _ = v50
        } else {
            v3[(((("hello" + (("a" + ((("\n\t" + "hello") + ((("" + "\n\t") + ((("a" + (("a" + "") + "a")) + "hello") + "")) + "foo bar")) + "")) + "\n\t")) + ("" + "foo bar")) + "\n\t") + ("" + "\n\t"))] = 127
            if (true || true) {
                v51 := GenS5{F0: -646, F1: v1, F2: 0}
                _ = v51
            }
            switch int(v5) {
            case 127:
                if true {
                    switch -2147483648 {
                    case -1:
                        switch 128 {
                        case 42:
                            var v52 int32 = 42
                            _ = v52
                        }
                    case 128:
                        var v53 int8 = (v5 + v5)
                        _ = v53
                    case 0:
                        v1 = (v1 >> ((255) & 63))
                    default:
                        v54 := []int{9223372036854775807, -419, 375, 4294967295}
                        _ = v54
                    }
                    v4["hello"] = false
                    v5 = 1
                } else {
                    switch "\n\t" {
                    case "\n\t":
                        switch -9223372036854775808 {
                        case 2147483647:
                            delete(v4, "")
                        case -9223372036854775808:
                            switch "a" {
                            case "foo bar":
                                v5 += 1
                            case "hello":
                                var v55 string = "foo bar"
                                _ = v55
                            default:
                                v56 := GenS2{F0: ("\n\t" + "a"), F1: 9223372036854775807, F2: v1}
                                _ = v56
                            }
                        default:
                            v57 := GenS6{F0: "a"}
                            _ = v57
                        }
                    default:
                        fmt.Println(v5, v5)
                    }
                    fmt.Println(v1, v5)
                }
            case 128:
                var v58 int64 = 9223372036854775807
                _ = v58
            case 1:
                for fi59 := 0; fi59 < 17; fi59++ {
                    switch fi59 {
                    case -1:
                        v5 += (v5 & v5)
                    }
                    v1 += (-(v1))
                    switch (fi59 | (fi59 & 128)) {
                    case 4294967295:
                        for fi60 := 0; fi60 < 13; fi60++ {
                            v5 -= v5
                            var v61 float64 = float64(0.0)
                            switch "foo bar" {
                            case "hello":
                                v62 := GenS1{F0: (v1 | v1), F1: false}
                                _ = v62
                            default:
                                v61 += v61
                            }
                            _ = fi60
                            _ = v61
                        }
                    case 127:
                        for fi63 := 0; fi63 < 17; fi63++ {
                            v64 := GenS7{F0: 1}
                            v65 := []int{(fi63 << ((4294967295) & 63)), fi63, -2147483648, -9223372036854775808}
                            v1 = -1
                            _ = fi63
                            _ = v64
                            _ = v65
                        }
                    }
                    _ = fi59
                }
            default:
                v5 = 100
            }
        }
        _ = v49
    }
    v66 := []int{}
    var v67 int32 = (v1 / ((v1) | 1))
    v68 := GenS1{F0: v67, F1: true}
    if false {
        var v69 uint8 = 42
        var v70 bool = true
        _ = v69
        _ = v70
    } else {
        var v71 uint64 = 627
        if (false && true) {
            v67 = (v67 / ((1) | 1))
            var v72 uint = 2
            _ = v72
        } else {
            _ = len(v66)
            _ = len(v66)
            v66 = append(v66, 2147483647)
        }
        _ = v71
    }
    var v73 int = 4294967295
    fmt.Println(v5)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v66:", len(v66), v66)
    fmt.Println("v67:", v67)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v66
    _ = v67
    _ = v68
    _ = v73
}
