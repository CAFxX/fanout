package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 int16 }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 uint; F1 int32; F2 uint }
type GenS4 struct { F0 uint64; F1 int }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint; F1 uint16; F2 uint32 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint16; F1 int16 }
func genFn0(p0 uint32, p1 uint, p2 string) uint32 {
    if (231 == 184) {
        v4 := []int{}
        _ = v4
    } else {
        for fi5 := 0; fi5 < 5; fi5++ {
            if (0 >= 514) {
                if (!true) {
                    p2 = (p2 + p2)
                }
            }
            _ = fi5
        }
        for fi6 := 0; fi6 < 2; fi6++ {
            v7 := GenS1{F0: 42, F1: (^(p0)), F2: -1}
            _ = fi6
            _ = v7
        }
    }
    p1 = (p1 ^ p1)
    if (((0.0 + (1e10 / (1.5 / 1.0))) == -2.25) && true) {
        v8 := []int{int(p0), -2147483648, 1, 2147483647}
        _ = v8
    }
    if false {
        p0 += (p0 * (p0 % ((p0) | 1)))
    } else {
        for fi9 := 0; fi9 < 4; fi9++ {
            if true {
                v10 := GenS1{F0: 0, F1: 222, F2: 0}
                _ = v10
            }
            _ = fi9
        }
        if true {
            switch -284 {
            case 42:
                if false {
                    fmt.Println(p1)
                    p1 -= 2
                    v11 := GenS2{F0: ((p2 + (("" + "a") + "\n\t")) + p2), F1: 42}
                    _ = v11
                } else {
                    v12 := []int{0, 239}
                    _ = v12
                }
            case 128:
                if true {
                    switch 0 {
                    case 128:
                        var v13 uint64 = 18446744073709551615
                        _ = v13
                    case 1:
                        switch 9223372036854775807 {
                        case -1:
                            p2 = "hello"
                        default:
                            var v14 uint32 = uint32(p1)
                            _ = v14
                        }
                    case 32768:
                        var v15 uint16 = 255
                        _ = v15
                    default:
                        v16 := []int{42, 127, -2147483648}
                        _ = v16
                    }
                    p2 = ("" + "a")
                    v17 := GenS3{F0: (^(p1)), F1: 488, F2: p1}
                    _ = v17
                } else {
                    p2 = ((p2 + p2) + "\n\t")
                }
            case -2147483648:
                var v18 uint8 = 42
                _ = v18
            default:
                p2 = "a"
            }
            var v19 uint8 = 255
            if true {
                for fi20 := 0; fi20 < 13; fi20++ {
                    if true {
                        var v21 uint32 = (p0 | p0)
                        var v22 string = "\n\t"
                        var v23 int32 = 408
                        _ = v21
                        _ = v22
                        _ = v23
                    } else {
                        for fi24 := 0; fi24 < 4; fi24++ {
                            var v25 int32 = -270
                            fmt.Println(fi24, p2, p0)
                            var v26 uint32 = p0
                            _ = fi24
                            _ = v25
                            _ = v26
                        }
                    }
                    p0 = p0
                    _ = fi20
                }
                if (652 < -478) {
                    p0 += (^(p0))
                    p1 += uint(p0)
                }
                if (1 >= 0) {
                    switch -3 {
                    case 127:
                        v27 := GenS1{F0: (^(v19)), F1: 2, F2: 42}
                        _ = v27
                    case -175:
                        for fi28 := 0; fi28 < 10; fi28++ {
                            v29 := []int{fi28}
                            _ = fi28
                            _ = v29
                        }
                    }
                }
            }
            _ = v19
        } else {
            switch -9223372036854775808 {
            case 128:
                for fi30 := 0; fi30 < 1; fi30++ {
                    for fi31 := 0; fi31 < 17; fi31++ {
                        v32 := []int{(fi31 * (fi31 << ((329) & 63))), -480}
                        v32 = append(v32, 683)
                        fmt.Println(p2)
                        _ = fi31
                        _ = v32
                    }
                    var v33 float32 = 1.5
                    _ = fi30
                    _ = v33
                }
            case 32767:
                if ((65535 != 1) || false) {
                    p1 -= 986
                } else {
                    p0 = p0
                    v34 := GenS4{F0: 2, F1: 2147483647}
                    _ = v34
                }
            }
            if true {
                for fi35 := 0; fi35 < 16; fi35++ {
                    p0 -= (p0 / ((p0) | 1))
                    v36 := GenS3{F0: 65535, F1: 346, F2: p1}
                    _ = fi35
                    _ = v36
                }
                if (189 <= 829) {
                    for fi37 := 0; fi37 < 9; fi37++ {
                        if false {
                            p0 -= (p0 - (^(p0)))
                            p0 = (p0 * p0)
                            fmt.Println(fi37)
                        }
                        _ = fi37
                    }
                    var v38 bool = (48 >= 127)
                    _ = v38
                }
                fmt.Println(p2)
            } else {
                fmt.Println(p1, p2)
                v39 := GenS5{F0: 1}
                _ = v39
            }
            p2 = (p2 + (("foo bar" + "foo bar") + (p2 + ((p2 + "") + ((p2 + "hello") + p2)))))
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return (p0 % ((p0) | 1))
}

func main() {
    var v1 uint8 = 106
    var v2 float32 = float32(-2.25)
    var v3 bool = false
    var v40 int16 = -1
    v41 := []int{128}
    v3 = ((!(32768 == 32768)) || (-44 < -19))
    if (1.5 <= 1e10) {
        var v42 int8 = -128
        if (v3 || (65535 >= 4294967295)) {
            v43 := GenS6{F0: 42, F1: 65535, F2: genFn0(1, 269, "")}
            v44 := GenS7{F0: v42}
            var v45 bool = (65535 < 255)
            _ = v43
            _ = v44
            _ = v45
        }
        _ = v42
    }
    v1 = 127
    var v46 int = 0
    var v47 int32 = -2147483648
    fmt.Println(v40)
    if v3 {
        var v48 uint8 = v1
        switch -9223372036854775808 {
        case -555:
            var v49 bool = true
            _ = v49
        }
        v47 = (v47 - (v47 * v47))
        _ = v48
    }
    if v3 {
        v1 = v1
        v50 := GenS5{F0: 18446744073709551615}
        _ = v50
    } else {
        var v51 float32 = 0.0
        switch v46 {
        case -9223372036854775808:
            switch 128 {
            case 9223372036854775807:
                if (0.1 >= 0.0) {
                    switch "foo bar" {
                    case "foo bar":
                        if len(v41) > 0 {
                            v41[0] = (v46 >> ((987) & 63))
                        }
                    case "hello":
                        if ((v40 | v40) != v40) {
                            fmt.Println(v51, v47, v3)
                            fmt.Println(v40)
                        }
                    default:
                        fmt.Println(v2, v46)
                    }
                } else {
                    fmt.Println(v46, v47)
                    var v52 float32 = v51
                    _ = len(v41)
                    _ = v52
                }
            case 127:
                v53 := GenS8{F0: 65535, F1: v40}
                _ = v53
            case -2147483648:
                for fi54 := 0; fi54 < 14; fi54++ {
                    var v55 int16 = v40
                    fmt.Println(v55)
                    _ = fi54
                    _ = v55
                }
            default:
                if len(v41) > 0 {
                    v41[0] = (v46 / ((-352) | 1))
                }
            }
        case 2147483647:
            v56 := GenS7{F0: -1}
            _ = v56
        default:
            var v57 uint16 = 255
            _ = v57
        }
        fmt.Println(v47, v2)
        _ = v51
    }
    v58 := GenS1{F0: 2, F1: 920, F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v40
    _ = v41
    _ = v46
    _ = v47
    _ = v58
}
