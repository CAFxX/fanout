package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32; F1 uint16; F2 int16 }
type GenS4 struct { F0 string; F1 uint }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int32 }
func genFn0() int64 {
    v6 := []int{127, 9223372036854775807, 9223372036854775807, -9223372036854775808}
    v6 = append(v6, 0)
    if (true || (0 == 2147483647)) {
        var v7 float64 = -2.25
        _ = v7
    }
    _ = v6
    return 1
}
func genFn1() int32 {
    var v8 uint32 = 2147483647
    var v9 uint8 = 127
    _ = v8
    _ = v9
    return -1
}

func main() {
    v1 := []int{}
    var v2 int8 = 42
    var v3 int32 = -2147483648
    v4 := make(map[uint32]int)
    v5 := []string{"a", "a"}
    for fi10 := 0; fi10 < 16; fi10++ {
        switch ("hello" + "\n\t") {
        case "":
            var v11 uint64 = 249
            _ = v11
        case "foo bar":
            if false {
                if false {
                    switch "a" {
                    case "hello":
                        var v12 float32 = float32(0.0)
                        _ = v12
                    case "a":
                        fmt.Println(v2, fi10)
                    case "":
                        if false {
                            switch fi10 {
                            case 825:
                                v3 -= -1
                            case -9223372036854775808:
                                v2 += v2
                            }
                            if (809 >= 1) {
                                if true {
                                    fmt.Println(fi10, fi10, v3)
                                    switch "a" {
                                    case "hello":
                                        v1 = append(v1, int(v3))
                                    case "foo bar":
                                        for fi13 := 0; fi13 < 6; fi13++ {
                                            v14 := GenS1{F0: 61}
                                            var v15 int64 = 24
                                            v16 := GenS1{F0: 2147483647}
                                            _ = fi13
                                            _ = v14
                                            _ = v15
                                            _ = v16
                                        }
                                    default:
                                        if true {
                                            v17 := GenS1{F0: 2147483647}
                                            fmt.Println(v2, fi10)
                                            _ = v17
                                        }
                                    }
                                } else {
                                    switch "foo bar" {
                                    case "foo bar":
                                        v18 := GenS1{F0: 630}
                                        _ = v18
                                    case "\n\t":
                                        for fi19 := 0; fi19 < 6; fi19++ {
                                            v2 = 100
                                            _ = fi19
                                        }
                                    default:
                                        if (42 < 42) {
                                            for fi20 := 0; fi20 < 15; fi20++ {
                                                v1 = append(v1, fi10)
                                                if true {
                                                    v3 += 0
                                                } else {
                                                    if len(v1) > 0 {
                                                        v1[0] = fi10
                                                    }
                                                }
                                                var v21 int16 = 1000
                                                _ = fi20
                                                _ = v21
                                            }
                                            switch 127 {
                                            case -1:
                                                switch ((("foo bar" + "hello") + "a") + "") {
                                                case "hello":
                                                    for fi22 := 0; fi22 < 17; fi22++ {
                                                        v2 -= (^(v2))
                                                        v3 += (-(v3))
                                                        var v23 uint16 = 32767
                                                        _ = fi22
                                                        _ = v23
                                                    }
                                                case "a":
                                                    v24 := GenS1{F0: 42}
                                                    _ = v24
                                                case "":
                                                    var v25 float32 = -2.25
                                                    _ = v25
                                                }
                                            case -2147483648:
                                                v3 = (v3 - v3)
                                            default:
                                                var v26 uint8 = 0
                                                _ = v26
                                            }
                                            v27 := GenS1{F0: 239}
                                            _ = v27
                                        }
                                    }
                                    var v28 uint8 = 33
                                    _ = len(v4)
                                    _ = v28
                                }
                                if false {
                                    var v29 uint32 = 836
                                    v29 = v29
                                    v2 = (^(v2))
                                    _ = v29
                                } else {
                                    var v30 int16 = 1
                                    for fi31 := 0; fi31 < 16; fi31++ {
                                        var v32 int16 = 998
                                        _ = len(v4)
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        _ = fi31
                                        _ = v32
                                    }
                                    _ = v30
                                }
                            } else {
                                v33 := GenS1{F0: 555}
                                _ = v33
                            }
                        } else {
                            v3 = v3
                        }
                    default:
                        var v34 uint32 = 65535
                        _ = v34
                    }
                    switch "foo bar" {
                    case "hello":
                        _ = len(v4)
                    case "foo bar":
                        v2 -= v2
                    default:
                        v35 := GenS2{F0: 240}
                        _ = v35
                    }
                    if false {
                        v36 := GenS2{F0: 32767}
                        switch "\n\t" {
                        case "\n\t":
                            fmt.Println(fi10, fi10)
                        default:
                            var v37 int32 = (-(v3))
                            _ = v37
                        }
                        _ = v36
                    }
                }
                v4[2] = (fi10 / (((fi10 >> ((1) & 63))) | 1))
                var v38 float64 = 1e10
                _ = v38
            } else {
                v2 -= (v2 >> ((38) & 63))
            }
        }
        v39 := GenS3{F0: 26, F1: 42, F2: 1}
        _ = fi10
        _ = v39
    }
    v3 += v3
    v3 = (v3 << ((282) & 63))
    switch -335 {
    case 9223372036854775807:
        if false {
            if ((^(v3)) == v3) {
                var v40 string = "\n\t"
                v41 := GenS2{F0: -32768}
                var v42 bool = ((false && true) && (!(0 == 687)))
                _ = v40
                _ = v41
                _ = v42
            } else {
                if (702 > 65535) {
                    v43 := GenS2{F0: 32767}
                    v44 := GenS4{F0: (("a" + "") + ("" + ("a" + ("a" + "a")))), F1: 765}
                    _ = v43
                    _ = v44
                }
            }
            switch -1 {
            case 127:
                switch "" {
                case "foo bar":
                    v3 = 2147483647
                case "":
                    for fi45 := 0; fi45 < 2; fi45++ {
                        if false {
                            _ = len(v1)
                            v5 = append(v5, "")
                            for fi46 := 0; fi46 < 7; fi46++ {
                                switch fi45 {
                                case 2147483647:
                                    v3 = (^(v3))
                                case -2147483648:
                                    if false {
                                        v47 := GenS5{F0: (v2 << ((1) & 63))}
                                        if true {
                                            var v48 int64 = 42
                                            _ = v48
                                        }
                                        _ = v47
                                    } else {
                                        delete(v4, 0)
                                    }
                                case 4294967295:
                                    var v49 bool = false
                                    _ = v49
                                }
                                v2 -= (^(v2))
                                _ = fi46
                            }
                        }
                        v1 = append(v1, -357)
                        v50 := GenS5{F0: 42}
                        _ = fi45
                        _ = v50
                    }
                default:
                    fmt.Println(v2, v2)
                }
            case 128:
                _ = len(v1)
            }
        }
    case -9223372036854775808:
        if (false || ((513 > 2) && false)) {
            switch "" {
            case "":
                switch 561 {
                case 1:
                    v2 -= v2
                default:
                    switch 621 {
                    case -1:
                        for fi51 := 0; fi51 < 8; fi51++ {
                            var v52 int32 = (v3 + (v3 % ((v3) | 1)))
                            _ = fi51
                            _ = v52
                        }
                    case 128:
                        if true {
                            fmt.Println(v3, v3)
                        }
                    }
                }
            case "\n\t":
                v2 += v2
            case "a":
                if (!false) {
                    if true {
                        v53 := GenS1{F0: 1}
                        v54 := GenS3{F0: 2147483648, F1: 42, F2: 42}
                        _ = len(v5)
                        _ = v53
                        _ = v54
                    } else {
                        var v55 int16 = 0
                        var v56 uint = 4294967295
                        v57 := v4[2147483647]
                        _ = v55
                        _ = v56
                        _ = v57
                    }
                } else {
                    var v58 int64 = -9223372036854775808
                    _ = v58
                }
            default:
                for fi59 := 0; fi59 < 15; fi59++ {
                    v3 = 42
                    if false {
                        for fi60 := 0; fi60 < 18; fi60++ {
                            v2 = (v2 ^ v2)
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            _ = fi60
                        }
                        v4[70] = 201
                        v2 = (-(v2))
                    }
                    _ = fi59
                }
            }
            v3 = v3
            var v61 float32 = ((0.0 - 1.5) + ((1.5 - 3.14159) * (((-2.25 / -2.25) + 0.0) + 1.5)))
            _ = v61
        }
    default:
        fmt.Println(v2, v3)
    }
    v62 := GenS1{F0: 65535}
    for fi63 := 0; fi63 < 8; fi63++ {
        v2 -= int8(fi63)
        var v64 int64 = -9223372036854775808
        switch (fi63 >> ((2) & 63)) {
        case -2147483648:
            v4[2] = fi63
        case -648:
            fmt.Println(v3, v64, fi63)
        }
        _ = fi63
        _ = v64
    }
    var v65 uint = 2
    fmt.Println(v2, v3, v65)
    for fi66 := 0; fi66 < 12; fi66++ {
        if len(v5) > 0 {
            v5[0] = "\n\t"
        }
        var v67 uint16 = 0
        _ = fi66
        _ = v67
    }
    var v68 uint16 = 255
    for fi69 := 0; fi69 < 11; fi69++ {
        fmt.Println(v3)
        v3 = v3
        v70 := GenS6{F0: v3}
        _ = fi69
        _ = v70
    }
    _ = len(v4)
    v65 = (v65 + (v65 ^ (v65 >> ((42) & 63))))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v65:", v65)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v62
    _ = v65
    _ = v68
}
