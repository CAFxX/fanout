package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool; F1 uint64 }
type GenS3 struct { F0 bool; F1 int64; F2 string }
type GenS4 struct { F0 uint16; F1 int16 }
type GenS5 struct { F0 int; F1 int16; F2 uint32 }
type GenS6 struct { F0 int16; F1 uint32; F2 int32 }
type GenS7 struct { F0 bool; F1 int; F2 int }
type GenS8 struct { F0 string; F1 int16; F2 bool }
type GenS9 struct { F0 int32; F1 bool }
func genFn0(p0 int8) uint16 {
    if (false && (0 == p0)) {
        p0 += p0
    }
    if (-87 > 0) {
        v6 := GenS1{F0: 2}
        _ = v6
    }
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := GenS2{F0: true, F1: 42}
        if (1 >= 255) {
            for fi9 := 0; fi9 < 17; fi9++ {
                var v10 uint = 2
                _ = fi9
                _ = v10
            }
            for fi11 := 0; fi11 < 16; fi11++ {
                v12 := []int{fi11}
                switch "a" {
                case "foo bar":
                    v12 = append(v12, fi11)
                case "":
                    p0 = (-(p0))
                case "a":
                    if true {
                        var v13 uint64 = 4294967295
                        _ = v13
                    }
                    switch "foo bar" {
                    case "a":
                        v14 := GenS1{F0: 4294967295}
                        _ = v14
                    case "":
                        v15 := GenS1{F0: 557}
                        _ = v15
                    default:
                        if true {
                            if (0 < 100) {
                                v16 := GenS3{F0: false, F1: -1, F2: "hello"}
                                if len(v12) > 0 {
                                    v12[0] = (fi11 % ((fi7) | 1))
                                }
                                _ = v16
                            } else {
                                p0 = p0
                                if false {
                                    var v17 uint8 = 0
                                    if (1e10 != (1.5 * 1e10)) {
                                        switch fi7 {
                                        case 9223372036854775807:
                                            var v18 uint16 = 32767
                                            _ = v18
                                        }
                                    } else {
                                        p0 = p0
                                        if false {
                                            v19 := GenS2{F0: true, F1: 408}
                                            var v20 bool = true
                                            _ = v19
                                            _ = v20
                                        } else {
                                            var v21 float64 = float64((float32(-2.25) * float32((((0.1 + 0.0) / 1.0) * 1e10))))
                                            v21 += v21
                                            _ = v21
                                        }
                                        p0 = p0
                                    }
                                    p0 -= (p0 ^ -50)
                                    _ = v17
                                } else {
                                    v12 = append(v12, (fi7 | 9223372036854775807))
                                    var v22 bool = (true || (65535 < 2147483647))
                                    if false {
                                        p0 -= p0
                                        v23 := GenS2{F0: (255 <= 65535), F1: 4294967295}
                                        v22 = false
                                        _ = v23
                                    } else {
                                        if v22 {
                                            v22 = (0 >= 236)
                                            if (0 >= 155) {
                                                if len(v12) > 0 {
                                                    v12[0] = (^(fi11))
                                                }
                                            } else {
                                                if v22 {
                                                    v12 = append(v12, (-(fi7)))
                                                    p0 = 100
                                                    p0 = (p0 % ((p0) | 1))
                                                    if (1 != 4294967295) {
                                                        var v24 float32 = float32(0.1)
                                                        var v25 uint64 = 1
                                                        _ = v24
                                                        _ = v25
                                                    }
                                                } else {
                                                    p0 = -100
                                                    _ = len(v12)
                                                    v22 = (42 < 2147483647)
                                                    v26 := GenS3{F0: ((true || false) || false), F1: -9223372036854775808, F2: "\n\t"}
                                                    p0 -= p0
                                                    v27 := GenS1{F0: 564}
                                                    _ = v26
                                                    _ = v27
                                                }
                                                fmt.Println(fi11, fi11)
                                                switch ("a" + "foo bar") {
                                                case "a":
                                                    var v28 int8 = (p0 << ((0) & 63))
                                                    _ = v28
                                                case "foo bar":
                                                    var v29 int64 = -1
                                                    _ = v29
                                                case "hello":
                                                    fmt.Println(fi7, p0)
                                                default:
                                                    switch "hello" {
                                                    case "hello":
                                                        v30 := GenS3{F0: (1 < 249), F1: 9223372036854775807, F2: "\n\t"}
                                                        _ = v30
                                                    case "foo bar":
                                                        var v31 bool = ((fi7 | fi11) > -9223372036854775808)
                                                        _ = v31
                                                    default:
                                                        v32 := GenS1{F0: 4294967295}
                                                        _ = v32
                                                    }
                                                }
                                            }
                                        }
                                        var v33 int16 = 42
                                        var v34 bool = true
                                        if v34 {
                                            v12 = append(v12, (fi7 | fi7))
                                            fmt.Println(fi11)
                                        } else {
                                            v34 = (1e10 <= 0.1)
                                            _ = len(v12)
                                            if len(v12) >= 2 {
                                                v12 = v12[1:len(v12)]
                                            }
                                        }
                                        _ = v33
                                        _ = v34
                                    }
                                    _ = v22
                                }
                            }
                            if ("foo bar" != ((("" + "\n\t") + "") + "hello")) {
                                v12 = append(v12, 1)
                                p0 = (p0 - p0)
                                fmt.Println(p0, fi7)
                            } else {
                                p0 -= p0
                                v35 := GenS1{F0: 2}
                                if ((!true) && ((3.14159 * 3.14159) >= (1e10 * 1e10))) {
                                    if true {
                                        var v36 uint = 255
                                        var v37 float32 = (float32(float64(0.0)) - (((-2.25 * 0.1) / 0.1) + 1.5))
                                        _ = v36
                                        _ = v37
                                    }
                                    var v38 bool = false
                                    _ = v38
                                } else {
                                    p0 -= p0
                                }
                                if (true || true) {
                                    v39 := GenS4{F0: 393, F1: 0}
                                    _ = v39
                                }
                                _ = v35
                            }
                            switch fi11 {
                            case 4294967295:
                                v40 := GenS5{F0: (fi11 - (fi7 & fi7)), F1: 1000, F2: 4294967295}
                                _ = v40
                            case 9223372036854775807:
                                if false {
                                    var v41 float32 = (1e10 * float32(1e10))
                                    v12 = append(v12, (fi7 << ((1) & 63)))
                                    _ = v41
                                }
                            default:
                                if len(v12) >= 2 {
                                    v12 = v12[1:len(v12)]
                                }
                            }
                        } else {
                            p0 -= p0
                            switch "foo bar" {
                            case "":
                                p0 = p0
                            case "hello":
                                v42 := GenS2{F0: (!((-2.25 - 3.14159) != 3.14159)), F1: 2}
                                _ = v42
                            default:
                                v43 := GenS6{F0: -32768, F1: 4294967295, F2: -2147483648}
                                _ = v43
                            }
                            v44 := GenS1{F0: 365}
                            _ = v44
                        }
                    }
                    p0 = int8(fi7)
                default:
                    p0 = 1
                }
                _ = fi11
                _ = v12
            }
            v45 := []int{128, int(p0)}
            _ = v45
        }
        var v46 float64 = (1.5 * 3.14159)
        _ = fi7
        _ = v46
        _ = v8
    }
    _ = p0
    return 255
}

func main() {
    var v1 uint32 = 467
    var v2 int64 = 894
    var v3 bool = false
    v4 := []uint32{(^(v1)), (^(v1))}
    v5 := make(map[int]int)
    if v3 {
        fmt.Println(v1, v2, v1)
    }
    for fi47 := 0; fi47 < 7; fi47++ {
        v48 := v5[(fi47 % (((^(fi47))) | 1))]
        switch "foo bar" {
        case "":
            v49 := GenS7{F0: (fi47 < (v48 * (fi47 - -843))), F1: (v48 & (-(fi47))), F2: (^(v48))}
            _ = v49
        default:
            for fi50 := 0; fi50 < 3; fi50++ {
                if v3 {
                    delete(v5, 32768)
                } else {
                    v4 = append(v4, 1)
                    v48 += (fi50 + 1)
                    v51 := GenS8{F0: "a", F1: int16(v1), F2: false}
                    v48 += fi47
                    _ = v51
                }
                if len(v4) > 0 {
                    v4[0] = (v1 | 44)
                }
                v3 = v3
                _ = fi50
            }
        }
        var v52 float32 = (0.1 + 1e10)
        _ = fi47
        _ = v48
        _ = v52
    }
    v2 -= (^(v2))
    for fi53 := 0; fi53 < 11; fi53++ {
        v4 = append(v4, v1)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi53
    }
    switch ("foo bar" + "foo bar") {
    case "hello":
        if len(v4) > 0 {
            v4[0] = v1
        }
    case "":
        var v54 uint16 = 858
        _ = v54
    }
    if len(v4) > 0 {
        v4[0] = v1
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v55 uint16 = 534
    switch "a" {
    case "hello":
        var v56 int = 4294967295
        _ = v56
    case "\n\t":
        fmt.Println(v2)
    }
    if v3 {
        switch "\n\t" {
        case "foo bar":
            if false {
                v1 = (^(v1))
                _ = len(v4)
            } else {
                var v57 int16 = 1
                _ = v57
            }
        default:
            v1 = (v1 & (^(v1)))
        }
    }
    if (2 > 128) {
        delete(v5, -690)
    }
    fmt.Println(v3)
    if true {
        fmt.Println(v3, v1)
        v58 := GenS9{F0: 42, F1: true}
        v2 -= 42
        _ = v58
    } else {
        fmt.Println(v1)
        fmt.Println(v2)
    }
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v55:", v55)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
}
