package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 int; F1 uint64; F2 int }
type GenS4 struct { F0 uint8; F1 int64; F2 uint16 }
type GenS5 struct { F0 int; F1 int64 }
type GenS6 struct { F0 int32; F1 uint64 }
type GenS7 struct { F0 uint8; F1 int8 }
func genFn0(p0 int64) int {
    var v5 uint32 = 2147483648
    var v6 uint64 = 523
    _ = p0
    _ = v5
    _ = v6
    return 2147483647
}
func genFn1(p0 int) int64 {
    switch ("foo bar" + "\n\t") {
    case "a":
        if true {
            var v7 int64 = 9223372036854775807
            for fi8 := 0; fi8 < 15; fi8++ {
                p0 = genFn0((v7 | 9223372036854775807))
                _ = fi8
            }
            _ = v7
        } else {
            for fi9 := 0; fi9 < 7; fi9++ {
                v10 := []int{fi9, fi9, (fi9 ^ genFn0(-110)), p0}
                if (genFn0(42) == 182) {
                    if true {
                        if len(v10) > 0 {
                            v10[0] = (fi9 >> ((2) & 63))
                        }
                        if false {
                            for fi11 := 0; fi11 < 6; fi11++ {
                                switch genFn0(0) {
                                case 9223372036854775807:
                                    if len(v10) >= 2 {
                                        v10 = v10[1:len(v10)]
                                    }
                                case 42:
                                    if (0 > 18446744073709551615) {
                                        v12 := GenS1{F0: 18446744073709551615}
                                        _ = v12
                                    } else {
                                        if (506 <= -1) {
                                            p0 = fi9
                                        }
                                    }
                                    switch ("a" + "foo bar") {
                                    case "\n\t":
                                        switch (^(fi11)) {
                                        case 32767:
                                            if len(v10) > 0 {
                                                v10[0] = genFn0(0)
                                            }
                                        default:
                                            v10 = append(v10, genFn0(-9223372036854775808))
                                        }
                                    case "a":
                                        if true {
                                            fmt.Println(fi9, p0)
                                        }
                                        var v13 int8 = 115
                                        _ = v13
                                    default:
                                        var v14 int8 = 0
                                        _ = v14
                                    }
                                    var v15 float32 = 3.14159
                                    _ = len(v10)
                                    _ = v15
                                case 2147483647:
                                    _ = len(v10)
                                }
                                fmt.Println(fi9, p0)
                                _ = fi11
                            }
                            for fi16 := 0; fi16 < 15; fi16++ {
                                fmt.Println(fi9, fi9, fi9)
                                _ = fi16
                            }
                            p0 = 1
                        } else {
                            var v17 uint32 = 734
                            _ = v17
                        }
                    } else {
                        for fi18 := 0; fi18 < 7; fi18++ {
                            p0 += (fi18 / ((fi18) | 1))
                            _ = fi18
                        }
                        for fi19 := 0; fi19 < 7; fi19++ {
                            if len(v10) > 0 {
                                v10[0] = (fi9 >> ((182) & 63))
                            }
                            if len(v10) > 0 {
                                v10[0] = 42
                            }
                            _ = fi19
                        }
                    }
                    v20 := GenS1{F0: 1}
                    if false {
                        p0 = 2147483647
                        v21 := GenS2{F0: 25, F1: -285}
                        var v22 float64 = 1e10
                        _ = v21
                        _ = v22
                    }
                    _ = v20
                } else {
                    for fi23 := 0; fi23 < 15; fi23++ {
                        fmt.Println(p0)
                        if ("foo bar" > ("\n\t" + ("" + "hello"))) {
                            v24 := GenS3{F0: genFn0(552), F1: 212, F2: genFn0(0)}
                            _ = v24
                        }
                        p0 += (fi9 >> ((2) & 63))
                        p0 -= genFn0(-1)
                        _ = fi23
                    }
                    v25 := GenS4{F0: 68, F1: 42, F2: 0}
                    var v26 int32 = -1
                    _ = v25
                    _ = v26
                }
                _ = fi9
                _ = v10
            }
        }
    }
    for fi27 := 0; fi27 < 20; fi27++ {
        fmt.Println(p0, p0)
        _ = fi27
    }
    if false {
        for fi28 := 0; fi28 < 2; fi28++ {
            fmt.Println(p0)
            switch "foo bar" {
            case "foo bar":
                p0 = (-(p0))
            case "a":
                var v29 uint32 = 2
                _ = v29
            case "\n\t":
                if (234 > 32767) {
                    p0 += 9223372036854775807
                    p0 = 0
                } else {
                    v30 := []int{991, 32768}
                    _ = v30
                }
            }
            var v31 int64 = 265
            _ = fi28
            _ = v31
        }
        if false {
            v32 := GenS5{F0: p0, F1: -1}
            if true {
                v33 := []int{(p0 - genFn0(42)), (p0 * (p0 & genFn0(-9223372036854775808))), -1, (p0 * (p0 >> ((65535) & 63)))}
                _ = v33
            }
            _ = v32
        }
    }
    _ = p0
    return -9223372036854775808
}

func main() {
    v1 := make(map[int]int)
    var v2 int16 = -1
    var v3 float32 = 0.1
    v4 := []uint32{}
    v3 = v3
    if false {
        var v34 int16 = (v2 ^ 1000)
        for fi35 := 0; fi35 < 6; fi35++ {
            v36 := GenS6{F0: 42, F1: 2}
            v37 := GenS7{F0: 1, F1: -1}
            _ = fi35
            _ = v36
            _ = v37
        }
        _ = v34
    }
    if true {
        v3 += v3
        if (455 != 2) {
            switch "\n\t" {
            case "hello":
                var v38 int = -82
                _ = v38
            default:
                var v39 int32 = -1
                _ = v39
            }
        } else {
            v3 = 1e10
        }
        for fi40 := 0; fi40 < 15; fi40++ {
            v41 := GenS2{F0: -2147483648, F1: 9223372036854775807}
            fmt.Println(v2, v2)
            _ = fi40
            _ = v41
        }
    }
    _ = len(v4)
    if false {
        v42 := GenS5{F0: genFn0(42), F1: 1}
        _ = v42
    } else {
        v2 = v2
        delete(v1, genFn0(9223372036854775807))
    }
    if ((^(v2)) < v2) {
        v3 -= 0.1
    } else {
        v43 := GenS6{F0: 42, F1: 18446744073709551615}
        v3 = 0.0
        _ = v43
    }
    v2 += v2
    for fi44 := 0; fi44 < 11; fi44++ {
        v2 = -678
        _ = fi44
    }
    if false {
        var v45 bool = true
        _ = v45
    } else {
        v46 := GenS2{F0: -1, F1: -9223372036854775808}
        for fi47 := 0; fi47 < 17; fi47++ {
            v2 += (v2 % ((v2) | 1))
            _ = fi47
        }
        v2 -= (v2 | v2)
        _ = v46
    }
    for fi48 := 0; fi48 < 3; fi48++ {
        var v49 float64 = (-2.25 * 1e10)
        _ = fi48
        _ = v49
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
