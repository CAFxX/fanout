package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint; F1 uint16; F2 int32 }
type GenS6 struct { F0 uint; F1 uint16; F2 int16 }
func genFn0(p0 bool) int32 {
    v6 := GenS1{F0: 128}
    var v7 uint8 = 42
    if (0 == 424) {
        var v8 bool = p0
        var v9 float32 = 0.0
        if false {
            v8 = false
            fmt.Println(v8)
            fmt.Println(p0, v8, v8)
        } else {
            if false {
                v10 := []int{0, 42, 42}
                for fi11 := 0; fi11 < 17; fi11++ {
                    v12 := GenS1{F0: fi11}
                    for fi13 := 0; fi13 < 7; fi13++ {
                        v9 += 0.1
                        v10 = append(v10, (fi13 * fi11))
                        v14 := GenS1{F0: 42}
                        _ = fi13
                        _ = v14
                    }
                    switch "a" {
                    case "\n\t":
                        for fi15 := 0; fi15 < 17; fi15++ {
                            fmt.Println(p0, v8)
                            _ = fi15
                        }
                    case "hello":
                        for fi16 := 0; fi16 < 12; fi16++ {
                            switch "" {
                            case "foo bar":
                                fmt.Println(v7)
                            }
                            _ = fi16
                        }
                    default:
                        _ = len(v10)
                    }
                    _ = fi11
                    _ = v12
                }
                _ = v10
            }
        }
        _ = v8
        _ = v9
    } else {
        v17 := []int{127, 42}
        v18 := GenS2{F0: 0, F1: 18446744073709551615}
        var v19 bool = (2147483647 >= 9223372036854775807)
        _ = v17
        _ = v18
        _ = v19
    }
    _ = p0
    _ = v6
    _ = v7
    return -2147483648
}
func genFn1(p0 uint, p1 int32, p2 string) uint {
    p0 = p0
    switch "" {
    case "hello":
        if ((!(true && (42 <= 797))) && (2147483648 < 0)) {
            fmt.Println(p1)
            var v20 int8 = -1
            v20 = 1
            _ = v20
        }
    case "":
        switch "" {
        case "":
            v21 := []int{1}
            _ = v21
        case "a":
            v22 := GenS3{F0: 4294967295}
            _ = v22
        case "\n\t":
            if (-9223372036854775808 < 42) {
                if false {
                    fmt.Println(p2)
                } else {
                    var v23 bool = (61 < 1)
                    if (v23 && (939 != 32767)) {
                        var v24 uint = (p0 >> (((p0 / (((^(p0))) | 1))) & 63))
                        _ = v24
                    }
                    _ = v23
                }
            }
        default:
            switch 857 {
            case -897:
                switch ((p2 + "a") + (p2 + p2)) {
                case "":
                    for fi25 := 0; fi25 < 3; fi25++ {
                        p0 -= p0
                        for fi26 := 0; fi26 < 11; fi26++ {
                            p2 = p2
                            switch "a" {
                            case "a":
                                p0 += (p0 + p0)
                                switch fi26 {
                                case 42:
                                    if false {
                                        if true {
                                            var v27 int32 = -848
                                            _ = v27
                                        }
                                        var v28 int8 = 95
                                        _ = v28
                                    } else {
                                        v29 := []int{}
                                        _ = v29
                                    }
                                default:
                                    var v30 uint32 = 4294967295
                                    _ = v30
                                }
                                var v31 float64 = 1.5
                                _ = v31
                            case "":
                                p2 = "\n\t"
                            default:
                                if ((!(true || true)) && (!(true && false))) {
                                    if false {
                                        var v32 bool = false
                                        p2 = (p2 + "")
                                        _ = v32
                                    }
                                } else {
                                    fmt.Println(fi26, p0)
                                    v33 := []int{128, -2147483648, fi25, (fi25 | (fi25 & 128))}
                                    var v34 int16 = 0
                                    _ = v33
                                    _ = v34
                                }
                                p2 = p2
                                var v35 float32 = float32((0.1 / (((-2.25 / 1e10) / 3.14159) + 0.1)))
                                _ = v35
                            }
                            p2 = p2
                            _ = fi26
                        }
                        _ = fi25
                    }
                case "hello":
                    if ("foo bar" < ("\n\t" + (p2 + p2))) {
                        for fi36 := 0; fi36 < 6; fi36++ {
                            p2 = "a"
                            _ = fi36
                        }
                    } else {
                        var v37 uint64 = 1
                        switch 2147483647 {
                        case 4294967295:
                            p2 = "foo bar"
                        case -977:
                            v38 := []int{}
                            _ = v38
                        case 42:
                            if (1 > 397) {
                                if (128 < 1) {
                                    var v39 int8 = -128
                                    var v40 uint8 = 255
                                    _ = v39
                                    _ = v40
                                } else {
                                    var v41 int = -1
                                    _ = v41
                                }
                                var v42 string = ((((p2 + p2) + p2) + p2) + "foo bar")
                                _ = v42
                            }
                        default:
                            p1 = (p1 << ((p0) & 63))
                        }
                        _ = v37
                    }
                case "a":
                    if ((p0 >> ((p0) & 63)) > (^(p0))) {
                        var v43 float64 = -2.25
                        v44 := GenS2{F0: 42, F1: 242}
                        v43 += v43
                        _ = v43
                        _ = v44
                    }
                default:
                    var v45 string = p2
                    _ = v45
                }
            case 0:
                fmt.Println(p2, p1)
            case -2147483648:
                switch ((p2 + ("hello" + p2)) + ((p2 + "") + p2)) {
                case "\n\t":
                    switch p2 {
                    case "\n\t":
                        for fi46 := 0; fi46 < 8; fi46++ {
                            var v47 int16 = -994
                            switch fi46 {
                            case 332:
                                var v48 int64 = 0
                                _ = v48
                            case 127:
                                var v49 uint8 = 128
                                _ = v49
                            }
                            _ = fi46
                            _ = v47
                        }
                    default:
                        var v50 int8 = 42
                        _ = v50
                    }
                case "foo bar":
                    if (!true) {
                        var v51 int8 = -1
                        if (genFn0(true) <= genFn0(false)) {
                            p2 = p2
                        }
                        var v52 float64 = float64(3.14159)
                        _ = v51
                        _ = v52
                    } else {
                        v53 := GenS1{F0: -2147483648}
                        fmt.Println(p1, p2, p0)
                        var v54 bool = false
                        _ = v53
                        _ = v54
                    }
                case "a":
                    v55 := GenS3{F0: 18446744073709551615}
                    _ = v55
                default:
                    var v56 float32 = 0.0
                    _ = v56
                }
            default:
                if false {
                    v57 := GenS4{F0: "hello"}
                    var v58 int16 = 0
                    _ = v57
                    _ = v58
                }
            }
        }
    case "\n\t":
        var v59 int8 = -128
        _ = v59
    }
    var v60 float64 = 1e10
    _ = p0
    _ = p1
    _ = p2
    _ = v60
    return (p0 / ((p0) | 1))
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]bool)
    var v3 float64 = 3.14159
    v4 := []string{}
    var v5 string = "hello"
    var v61 string = "a"
    fmt.Println(v5, v5)
    if true {
        v61 = ""
        fmt.Println(v5, v61)
        var v62 uint16 = 203
        _ = v62
    } else {
        fmt.Println(v61, v5)
    }
    fmt.Println(v5, v61, v3)
    var v63 uint16 = 2
    if true {
        v5 = v61
    }
    if false {
        switch 32768 {
        case -395:
            v5 = (("\n\t" + "\n\t") + (v61 + ("hello" + v61)))
        case 385:
            if true {
                if true {
                    v61 = (v5 + v5)
                }
                switch (((v61 + v5) + (("a" + v61) + v61)) + v61) {
                case "a":
                    var v64 int = 550
                    _ = v64
                default:
                    v65 := GenS5{F0: 873, F1: (^(v63)), F2: -58}
                    _ = v65
                }
                if (true || true) {
                    v61 = "foo bar"
                    _ = len(v2)
                }
            }
        }
    }
    v5 = (v5 + (v61 + v61))
    for fi66 := 0; fi66 < 16; fi66++ {
        for fi67 := 0; fi67 < 20; fi67++ {
            _ = len(v2)
            if len(v4) > 0 {
                v4[0] = v61
            }
            switch (fi66 << ((4294967295) & 63)) {
            case 1:
                fmt.Println(fi67, v61, v61)
            default:
                v68 := GenS3{F0: 0}
                _ = v68
            }
            _ = fi67
        }
        if ("" <= v61) {
            var v69 uint = 0
            _ = v69
        } else {
            v4 = append(v4, ("" + v5))
        }
        _ = fi66
    }
    switch -9223372036854775808 {
    case -255:
        fmt.Println(v5, v63, v61)
    default:
        var v70 int = 647
        _ = v70
    }
    if false {
        for fi71 := 0; fi71 < 14; fi71++ {
            v72 := GenS6{F0: 42, F1: 1, F2: -32768}
            fmt.Println(v5, v61, fi71)
            _ = fi71
            _ = v72
        }
        if len(v4) > 0 {
            v4[0] = "foo bar"
        }
        var v73 uint = 18446744073709551615
        _ = v73
    }
    var v74 float32 = ((float32((-2.25 - ((3.14159 / (((v3 * v3) / (v3 + float64(-2.25))) * v3)) - -2.25))) - 0.1) * 0.0)
    v3 += v3
    v5 = v5
    fmt.Println(v3)
    if true {
        v74 += 3.14159
    }
    fmt.Println(v3, v63, v74)
    v75 := GenS6{F0: 18446744073709551615, F1: v63, F2: 32767}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v61:", v61)
    fmt.Println("v63:", v63)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v61
    _ = v63
    _ = v74
    _ = v75
}
