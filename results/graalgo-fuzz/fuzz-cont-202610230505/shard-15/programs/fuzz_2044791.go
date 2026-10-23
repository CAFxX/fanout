package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 int8; F2 int64 }
type GenS3 struct { F0 int16; F1 uint }
type GenS4 struct { F0 int16; F1 bool }
type GenS5 struct { F0 uint32; F1 int32; F2 int8 }
type GenS6 struct { F0 int16; F1 uint8; F2 uint }
func genFn0(p0 bool) uint16 {
    v7 := GenS1{F0: 0}
    var v8 int64 = 1
    fmt.Println(p0, p0)
    switch ((("\n\t" + "\n\t") + ("a" + "hello")) + "") {
    case "foo bar":
        v9 := GenS2{F0: 0, F1: -41, F2: (v8 * 547)}
        _ = v9
    case "\n\t":
        switch "foo bar" {
        case "\n\t":
            var v10 int32 = 42
            _ = v10
        case "":
            p0 = p0
        case "a":
            v11 := GenS1{F0: 34}
            _ = v11
        default:
            for fi12 := 0; fi12 < 8; fi12++ {
                var v13 uint = 381
                _ = fi12
                _ = v13
            }
        }
    case "":
        if p0 {
            v14 := []int{9223372036854775807, 42}
            _ = v14
        } else {
            v8 -= v8
        }
    }
    v8 = v8
    _ = p0
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 int16 = 32767
    var v2 uint = 2
    var v3 uint64 = 1
    var v4 uint16 = 42
    v5 := make(map[int]int)
    v6 := []uint32{}
    if (!((v1 >= v1) || false)) {
        for fi15 := 0; fi15 < 18; fi15++ {
            v16 := GenS2{F0: 2147483647, F1: 42, F2: 42}
            for fi17 := 0; fi17 < 1; fi17++ {
                if ((0.1 == 1.5) || (2147483647 >= 0)) {
                    var v18 int = -101
                    var v19 int = fi17
                    _ = v18
                    _ = v19
                }
                _ = fi17
            }
            _ = fi15
            _ = v16
        }
    }
    switch "hello" {
    case "":
        if true {
            for fi20 := 0; fi20 < 12; fi20++ {
                if true {
                    for fi21 := 0; fi21 < 17; fi21++ {
                        v1 = v1
                        _ = fi21
                    }
                }
                _ = fi20
            }
            if (773 == 32767) {
                v22 := GenS2{F0: 455, F1: -128, F2: 0}
                var v23 uint16 = (v4 ^ v4)
                _ = v22
                _ = v23
            }
        } else {
            switch "hello" {
            case "a":
                fmt.Println(v2)
            case "\n\t":
                v24 := GenS3{F0: (-(v1)), F1: v2}
                _ = v24
            case "foo bar":
                if ((!(2 <= 1)) || ("a" != "a")) {
                    v4 -= (v4 - (v4 & v4))
                    var v25 float32 = -2.25
                    _ = v25
                }
            default:
                v26 := GenS4{F0: (v1 & v1), F1: (!(-32768 == -325))}
                _ = v26
            }
            if true {
                if false {
                    for fi27 := 0; fi27 < 15; fi27++ {
                        v3 -= 2
                        for fi28 := 0; fi28 < 3; fi28++ {
                            v29 := v5[fi28]
                            var v30 string = "hello"
                            var v31 uint8 = 127
                            _ = fi28
                            _ = v29
                            _ = v30
                            _ = v31
                        }
                        v32 := GenS1{F0: v2}
                        _ = fi27
                        _ = v32
                    }
                } else {
                    v33 := GenS2{F0: 2147483648, F1: -40, F2: 42}
                    switch 32767 {
                    case 470:
                        switch ("a" + ("" + "\n\t")) {
                        case "foo bar":
                            var v34 uint16 = 2
                            _ = v34
                        case "a":
                            var v35 uint32 = 396
                            _ = v35
                        default:
                            if false {
                                var v36 uint16 = 32767
                                _ = v36
                            } else {
                                if (1e10 > 1e10) {
                                    if true {
                                        v37 := GenS5{F0: 4294967295, F1: 42, F2: 127}
                                        var v38 int32 = -363
                                        if (127 == -207) {
                                            v3 -= 18446744073709551615
                                            if true {
                                                v39 := GenS3{F0: (v1 & v1), F1: (v2 - v2)}
                                                v40 := GenS3{F0: -1, F1: (v2 - (v2 + (v2 >> (((v2 % ((v2) | 1))) & 63))))}
                                                _ = v39
                                                _ = v40
                                            } else {
                                                v2 += 768
                                                switch -9223372036854775808 {
                                                case -762:
                                                    switch 105 {
                                                    case 2147483647:
                                                        delete(v5, 127)
                                                    }
                                                case -9223372036854775808:
                                                    var v41 int16 = (v1 << (((v2 - v2)) & 63))
                                                    _ = v41
                                                case 2147483647:
                                                    if false {
                                                        var v42 float32 = 0.1
                                                        fmt.Println(v3, v4)
                                                        _ = v42
                                                    }
                                                default:
                                                    v5[128] = 4294967295
                                                }
                                            }
                                            fmt.Println(v4, v3, v2)
                                        } else {
                                            v4 = v4
                                            fmt.Println(v1)
                                        }
                                        _ = v37
                                        _ = v38
                                    }
                                } else {
                                    for fi43 := 0; fi43 < 19; fi43++ {
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                        var v44 int8 = 127
                                        _ = fi43
                                        _ = v44
                                    }
                                    if true {
                                        v3 += v3
                                    } else {
                                        var v45 int = -873
                                        _ = v45
                                    }
                                    switch -2147483648 {
                                    case 128:
                                        for fi46 := 0; fi46 < 9; fi46++ {
                                            v3 = v3
                                            v2 = v2
                                            _ = fi46
                                        }
                                    }
                                }
                                switch "\n\t" {
                                case "foo bar":
                                    for fi47 := 0; fi47 < 7; fi47++ {
                                        var v48 int = fi47
                                        var v49 int8 = 1
                                        _ = fi47
                                        _ = v48
                                        _ = v49
                                    }
                                case "":
                                    v2 = (v2 % ((v2) | 1))
                                case "hello":
                                    if true {
                                        v2 -= 42
                                    }
                                }
                            }
                        }
                    default:
                        v1 = v1
                    }
                    _ = v33
                }
                switch "a" {
                case "foo bar":
                    var v50 float64 = 0.1
                    _ = v50
                case "a":
                    var v51 int32 = 1
                    _ = v51
                case "\n\t":
                    v52 := GenS5{F0: 2, F1: 0, F2: 0}
                    _ = v52
                }
            }
            v2 = (^(v2))
        }
    case "a":
        var v53 uint8 = 0
        _ = v53
    case "hello":
        fmt.Println(v1)
    default:
        v6 = append(v6, 2147483647)
    }
    var v54 uint8 = 128
    if true {
        var v55 uint8 = (v54 >> ((uint(v4)) & 63))
        _ = v55
    } else {
        if ((-2.25 * -2.25) == (((float64(float32((float64(float32(-2.25)) * (float64(1e10) + 3.14159)))) + -2.25) / 1e10) - 0.1)) {
            fmt.Println(v4, v3, v3)
        } else {
            v54 = 128
            delete(v5, -9223372036854775808)
            if true {
                if (false && false) {
                    if (673 == 9223372036854775807) {
                        v4 -= (^(v4))
                        var v56 uint16 = (v4 << (((v2 % ((v2) | 1))) & 63))
                        _ = v56
                    }
                    v57 := GenS6{F0: 0, F1: v54, F2: 42}
                    _ = v57
                }
                var v58 int64 = 0
                var v59 int64 = (v58 - (v58 & int64(v3)))
                _ = v58
                _ = v59
            } else {
                fmt.Println(v2)
            }
        }
        var v60 uint32 = 0
        fmt.Println(v60, v60)
        _ = v60
    }
    var v61 string = ("foo bar" + "a")
    v3 -= v3
    fmt.Println(v4)
    v1 = v1
    var v62 string = "hello"
    var v63 int = 9223372036854775807
    _ = len(v6)
    var v64 uint = v2
    var v65 uint = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v54:", v54)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v6
    _ = v61
    _ = v62
    _ = v63
    _ = v64
    _ = v65
}
