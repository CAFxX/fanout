package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 int8 }
type GenS2 struct { F0 uint16; F1 int64; F2 int16 }
type GenS3 struct { F0 uint8; F1 int32; F2 bool }
type GenS4 struct { F0 int; F1 int32; F2 int8 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 bool; F1 uint8; F2 int32 }
type GenS7 struct { F0 uint64; F1 bool }
type GenS8 struct { F0 uint64; F1 int; F2 int32 }
type GenS9 struct { F0 uint }
func genFn0(p0 uint16, p1 bool) int16 {
    var v6 bool = (32767 != 0)
    if (32767 < -211) {
        var v7 uint8 = 80
        _ = v7
    }
    var v8 uint8 = uint8(p0)
    _ = p0
    _ = p1
    _ = v6
    _ = v8
    return -32768
}

func main() {
    var v1 float64 = -2.25
    v2 := make(map[int]int)
    var v3 int64 = -935
    var v4 int16 = 32767
    v5 := make(map[uint32]bool)
    var v9 int32 = 0
    fmt.Println(v4, v3, v3)
    var v10 string = "foo bar"
    if ((42 == 9223372036854775807) || (!false)) {
        var v11 int = -546
        fmt.Println(v4, v9)
        _ = v11
    } else {
        fmt.Println(v10, v3)
        if false {
            var v12 uint32 = 453
            v13 := GenS1{F0: 18446744073709551615, F1: v3, F2: 100}
            _ = v12
            _ = v13
        }
        if false {
            v14 := GenS2{F0: 42, F1: v3, F2: (v4 % ((v4) | 1))}
            var v15 uint64 = 947
            fmt.Println(v9)
            _ = v14
            _ = v15
        } else {
            v10 = (v10 + "a")
            if (((((0.1 - ((3.14159 - 0.1) * float32(float64(0.0)))) != (-2.25 + 1.5)) || (false || (false && false))) && ((128 > 0) && false)) && false) {
                v16 := []int{0, 4294967295}
                _ = v16
            } else {
                switch -1 {
                case -298:
                    switch (v10 + v10) {
                    case "foo bar":
                        v17 := GenS3{F0: 2, F1: 0, F2: (2 > 251)}
                        _ = v17
                    case "hello":
                        if true {
                            v18 := GenS4{F0: 470, F1: (v9 >> ((287) & 63)), F2: int8(v4)}
                            if (0 < 2) {
                                v19 := v5[239]
                                _ = v19
                            } else {
                                v20 := v5[42]
                                v21 := []int{1}
                                if v20 {
                                    var v22 int32 = 42
                                    _ = v22
                                } else {
                                    switch ((v10 + v10) + (v10 + v10)) {
                                    case "foo bar":
                                        switch 1 {
                                        case 9223372036854775807:
                                            if (2 < 2) {
                                                var v23 uint8 = 255
                                                _ = v23
                                            }
                                        case 127:
                                            if ((!(v1 > 3.14159)) || (1e10 < (3.14159 - float32(v1)))) {
                                                if v20 {
                                                    if (v9 < (v9 << ((442) & 63))) {
                                                        fmt.Println(v9, v4, v20)
                                                    } else {
                                                        var v24 uint = 42
                                                        v9 = -1
                                                        _ = v24
                                                    }
                                                } else {
                                                    delete(v2, 367)
                                                }
                                                fmt.Println(v20, v10)
                                            } else {
                                                v5[2] = ((true || true) && (1e10 != 3.14159))
                                                v25 := GenS4{F0: -9223372036854775808, F1: 1, F2: 1}
                                                _ = v25
                                            }
                                        case -143:
                                            if (true || (v3 <= v3)) {
                                                var v26 int32 = (v9 / ((v9) | 1))
                                                v1 -= float64(1e10)
                                                _ = v26
                                            } else {
                                                v21 = append(v21, 354)
                                                if false {
                                                    var v27 float32 = (1e10 + 0.0)
                                                    v21 = append(v21, 226)
                                                    v28 := GenS5{F0: 18446744073709551615}
                                                    _ = v27
                                                    _ = v28
                                                }
                                            }
                                        default:
                                            v29 := GenS6{F0: v20, F1: 128, F2: v9}
                                            _ = v29
                                        }
                                    case "":
                                        if (!v20) {
                                            if len(v21) > 0 {
                                                v21[0] = 870
                                            }
                                        } else {
                                            var v30 uint32 = 4294967295
                                            for fi31 := 0; fi31 < 19; fi31++ {
                                                var v32 bool = v20
                                                switch "a" {
                                                case "":
                                                    switch "foo bar" {
                                                    case "hello":
                                                        v21 = append(v21, 127)
                                                    case "":
                                                        if len(v21) > 0 {
                                                            v21[0] = fi31
                                                        }
                                                    case "\n\t":
                                                        var v33 bool = false
                                                        _ = v33
                                                    }
                                                }
                                                var v34 string = "hello"
                                                _ = fi31
                                                _ = v32
                                                _ = v34
                                            }
                                            v3 = (v3 - v3)
                                            _ = v30
                                        }
                                    case "a":
                                        v1 = (((float64((((0.1 + 0.1) / (0.1 + -2.25)) / 0.1)) / 1.5) * v1) / (-2.25 + float64(0.1)))
                                    default:
                                        var v35 uint32 = 949
                                        _ = v35
                                    }
                                    for fi36 := 0; fi36 < 7; fi36++ {
                                        for fi37 := 0; fi37 < 8; fi37++ {
                                            if (!(v20 && v20)) {
                                                v38 := GenS7{F0: 18446744073709551615, F1: v20}
                                                var v39 int64 = v3
                                                v9 = v9
                                                _ = v38
                                                _ = v39
                                            }
                                            v9 += (v9 >> ((65535) & 63))
                                            v40 := GenS1{F0: 897, F1: v3, F2: 55}
                                            _ = fi37
                                            _ = v40
                                        }
                                        _ = fi36
                                    }
                                }
                                _ = v20
                                _ = v21
                            }
                            _ = v18
                        } else {
                            if ((!(true || (1.5 == ((float32((float64(0.0) * 3.14159)) - 0.0) / 3.14159)))) && (true && (9223372036854775807 != 32768))) {
                                switch (("a" + ("foo bar" + ("hello" + ("foo bar" + v10)))) + "hello") {
                                case "foo bar":
                                    v9 = v9
                                case "":
                                    switch "" {
                                    case "a":
                                        var v41 int64 = (^(v3))
                                        _ = v41
                                    case "foo bar":
                                        v3 -= v3
                                    default:
                                        if false {
                                            if (float64((0.0 * (3.14159 + (((0.1 * (1.5 * 1.5)) - float32(-2.25)) - 1.5)))) < v1) {
                                                switch -733 {
                                                case 1:
                                                    fmt.Println(v1, v1)
                                                default:
                                                    var v42 int8 = -1
                                                    _ = v42
                                                }
                                                var v43 int16 = int16(v3)
                                                switch "a" {
                                                case "foo bar":
                                                    for fi44 := 0; fi44 < 10; fi44++ {
                                                        v45 := []int{-346, (fi44 & 128)}
                                                        _ = fi44
                                                        _ = v45
                                                    }
                                                case "a":
                                                    v46 := GenS6{F0: (((908 >= -9223372036854775808) || false) && true), F1: 196, F2: (v9 * v9)}
                                                    _ = v46
                                                }
                                                _ = v43
                                            } else {
                                                if (true || (42 < 1)) {
                                                    var v47 uint = 65535
                                                    var v48 int = 127
                                                    _ = v47
                                                    _ = v48
                                                }
                                                v49 := GenS5{F0: 0}
                                                _ = v49
                                            }
                                        }
                                    }
                                case "a":
                                    v50 := GenS8{F0: 1, F1: 127, F2: (^(v9))}
                                    _ = v50
                                default:
                                    switch v10 {
                                    case "foo bar":
                                        for fi51 := 0; fi51 < 16; fi51++ {
                                            var v52 uint8 = 42
                                            var v53 float64 = v1
                                            _ = fi51
                                            _ = v52
                                            _ = v53
                                        }
                                    case "\n\t":
                                        v54 := []int{2147483647, 42}
                                        _ = v54
                                    case "hello":
                                        if (!true) {
                                            v55 := GenS2{F0: 255, F1: 42, F2: 1}
                                            _ = v55
                                        }
                                    }
                                }
                                for fi56 := 0; fi56 < 14; fi56++ {
                                    for fi57 := 0; fi57 < 1; fi57++ {
                                        var v58 int32 = -768
                                        _ = fi57
                                        _ = v58
                                    }
                                    _ = fi56
                                }
                                v1 = (0.1 - (v1 + v1))
                            }
                            var v59 bool = false
                            _ = v59
                        }
                    default:
                        fmt.Println(v1, v4, v1)
                    }
                default:
                    for fi60 := 0; fi60 < 16; fi60++ {
                        fmt.Println(v4)
                        _ = fi60
                    }
                }
            }
            switch -490 {
            case 529:
                v61 := GenS9{F0: 18446744073709551615}
                _ = v61
            default:
                var v62 uint8 = 2
                _ = v62
            }
        }
    }
    if false {
        for fi63 := 0; fi63 < 19; fi63++ {
            switch "a" {
            case "foo bar":
                v3 += (v3 % ((-555) | 1))
            case "hello":
                fmt.Println(v1, v9)
            case "a":
                v4 = int16(fi63)
            default:
                if (1 == 255) {
                    v64 := GenS4{F0: (-(fi63)), F1: (v9 + 42), F2: 100}
                    var v65 bool = false
                    if false {
                        var v66 float32 = 1e10
                        v9 = v9
                        v67 := GenS5{F0: 349}
                        _ = v66
                        _ = v67
                    }
                    _ = v64
                    _ = v65
                } else {
                    v68 := GenS4{F0: (fi63 & -534), F1: (-(v9)), F2: 35}
                    _ = v68
                }
            }
            v69 := []int{(^(fi63)), (fi63 / (((fi63 | (fi63 >> ((675) & 63)))) | 1))}
            _ = fi63
            _ = v69
        }
    }
    v4 += -32768
    fmt.Println(v4, v3)
    v1 += 3.14159
    fmt.Println(v10, v1)
    fmt.Println(v9, v4)
    v70 := []int{42, 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v70:", len(v70), v70)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v70
    _ = v9
}
