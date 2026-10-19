package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 bool }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 bool; F1 uint16 }
type GenS4 struct { F0 uint8; F1 uint8 }
type GenS5 struct { F0 string; F1 bool; F2 int16 }
type GenS6 struct { F0 int32; F1 uint64 }
type GenS7 struct { F0 int64; F1 int; F2 int16 }
type GenS8 struct { F0 int16; F1 int }
type GenS9 struct { F0 uint64; F1 string; F2 int64 }
func genFn0() int {
    var v7 bool = (-1 <= 1)
    if (v7 && ((!(!true)) || false)) {
        for fi8 := 0; fi8 < 2; fi8++ {
            v9 := GenS1{F0: (fi8 ^ (fi8 % ((fi8) | 1))), F1: (100 > 42), F2: ((((fi8 * 32767) >= (fi8 >> ((255) & 63))) && (63 >= 13)) || (float32((0.0 * 0.1)) != 3.14159))}
            for fi10 := 0; fi10 < 7; fi10++ {
                if v7 {
                    v7 = true
                } else {
                    if ((false && (1000 >= -1)) && v7) {
                        if (1 >= 1) {
                            fmt.Println(v7, v7)
                        } else {
                            fmt.Println(fi8, v7, fi8)
                            v11 := []int{(^(fi8)), (-(fi8)), 0}
                            v7 = (615 != 0)
                            _ = v11
                        }
                        v7 = true
                        var v12 uint = 255
                        _ = v12
                    }
                    if (!v7) {
                        v13 := []int{-9223372036854775808, fi8, -438}
                        v14 := GenS1{F0: fi10, F1: v7, F2: (((!(true || (v7 && v7))) && false) && (!v7))}
                        var v15 uint64 = 4294967295
                        _ = v13
                        _ = v14
                        _ = v15
                    }
                    v16 := []int{fi10, -2147483648, (-(fi8)), -2147483648}
                    _ = v16
                }
                fmt.Println(v7)
                v7 = v7
                _ = fi10
            }
            v7 = (!(65535 >= 174))
            _ = fi8
            _ = v9
        }
        var v17 int32 = -379
        _ = v17
    }
    fmt.Println(v7)
    if (v7 && (909 > 0)) {
        v18 := GenS2{F0: 2147483647, F1: 255}
        _ = v18
    } else {
        v19 := GenS2{F0: 1, F1: 255}
        _ = v19
    }
    _ = v7
    return 128
}
func genFn1(p0 bool, p1 int32, p2 bool) uint32 {
    var v20 float64 = 0.0
    v21 := GenS3{F0: (p2 || (false && (221 == 1))), F1: 0}
    fmt.Println(p0)
    fmt.Println(p2, p0, p0)
    fmt.Println(v20, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    _ = v21
    return 1
}

func main() {
    var v1 uint64 = 1
    var v2 int = -1
    v3 := make(map[string]int)
    v4 := make(map[string]int)
    v5 := []uint32{65535, 780, 0, 536}
    v6 := []int{}
    switch "a" {
    case "a":
        v1 = 0
    case "hello":
        _ = len(v3)
    case "foo bar":
        v1 = v1
    default:
        switch -328 {
        case 32767:
            for fi22 := 0; fi22 < 10; fi22++ {
                fmt.Println(fi22, v2, v1)
                _ = fi22
            }
        default:
            if (0 < 32767) {
                var v23 bool = (!true)
                _ = v23
            } else {
                fmt.Println(v1)
                var v24 uint8 = 2
                if (255 > 255) {
                    var v25 uint32 = 125
                    v26 := GenS2{F0: 42, F1: 255}
                    v27 := GenS4{F0: (v24 & (v24 ^ v24)), F1: 255}
                    _ = v25
                    _ = v26
                    _ = v27
                } else {
                    v1 = 42
                    if (511 < genFn1(true, -2147483648, false)) {
                        v24 = (v24 * v24)
                        switch "\n\t" {
                        case "":
                            var v28 uint = 0
                            _ = v28
                        case "foo bar":
                            for fi29 := 0; fi29 < 11; fi29++ {
                                if false {
                                    fmt.Println(v24, fi29, fi29)
                                    _ = len(v5)
                                }
                                v24 = 1
                                _ = fi29
                            }
                        default:
                            fmt.Println(v2)
                        }
                        v30 := GenS4{F0: (v24 * (v24 / (((v24 + (v24 & (v24 << ((318) & 63))))) | 1))), F1: (v24 - v24)}
                        _ = v30
                    }
                }
                _ = v24
            }
        }
    }
    switch genFn0() {
    case -1:
        var v31 int64 = 42
        _ = v31
    case 143:
        if (2 >= 127) {
            var v32 int16 = -527
            v1 = v1
            _ = v32
        }
    case 2147483647:
        if true {
            for fi33 := 0; fi33 < 6; fi33++ {
                v1 += v1
                _ = len(v6)
                v34 := GenS5{F0: "\n\t", F1: false, F2: 891}
                _ = fi33
                _ = v34
            }
        } else {
            switch "foo bar" {
            case "foo bar":
                var v35 uint64 = (v1 * v1)
                _ = v35
            case "hello":
                if true {
                    for fi36 := 0; fi36 < 17; fi36++ {
                        v2 += v2
                        fmt.Println(v1, fi36)
                        _ = fi36
                    }
                }
            case "":
                if false {
                    fmt.Println(v2, v2)
                    v37 := GenS6{F0: 278, F1: 18446744073709551615}
                    v1 -= (v1 << ((148) & 63))
                    _ = v37
                }
            default:
                v2 += v2
            }
        }
    }
    if false {
        v1 = (^(v1))
        switch ("foo bar" + "hello") {
        case "\n\t":
            if ((100 > 0) || true) {
                v2 = (v2 >> ((977) & 63))
                var v38 int16 = 177
                v39 := GenS7{F0: -414, F1: 127, F2: (v38 << ((0) & 63))}
                _ = v38
                _ = v39
            }
        case "":
            var v40 float64 = 0.1
            _ = v40
        default:
            switch ("hello" + "") {
            case "foo bar":
                v1 = v1
            case "a":
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            case "":
                if false {
                    var v41 int64 = -578
                    _ = len(v6)
                    _ = v41
                } else {
                    v42 := GenS7{F0: -1, F1: 9223372036854775807, F2: 1}
                    var v43 string = "hello"
                    _ = v42
                    _ = v43
                }
            default:
                v44 := GenS8{F0: -32768, F1: (^(v2))}
                _ = v44
            }
        }
        var v45 uint = 4294967295
        _ = v45
    }
    for fi46 := 0; fi46 < 14; fi46++ {
        v2 -= (^(v2))
        var v47 string = ((("foo bar" + "\n\t") + "foo bar") + ("\n\t" + ((("" + "\n\t") + ("" + "\n\t")) + ("\n\t" + "hello"))))
        v1 -= (v1 | (^(v1)))
        _ = fi46
        _ = v47
    }
    if (65535 > 2147483647) {
        v48 := GenS9{F0: v1, F1: (("\n\t" + "") + ("foo bar" + "a")), F2: 386}
        for fi49 := 0; fi49 < 1; fi49++ {
            v50 := GenS7{F0: 1, F1: genFn0(), F2: -32768}
            var v51 bool = ((^(v1)) > v1)
            v3["a"] = 903
            _ = fi49
            _ = v50
            _ = v51
        }
        _ = v48
    } else {
        if len(v6) > 0 {
            v6[0] = (v2 >> ((0) & 63))
        }
    }
    _ = len(v4)
    var v52 uint32 = 1
    for fi53 := 0; fi53 < 18; fi53++ {
        var v54 int8 = -48
        switch "a" {
        case "foo bar":
            _ = len(v6)
        default:
            v54 = (v54 & -1)
        }
        var v55 uint64 = v1
        _ = fi53
        _ = v54
        _ = v55
    }
    for fi56 := 0; fi56 < 5; fi56++ {
        delete(v4, "a")
        v2 = (fi56 >> ((4294967295) & 63))
        var v57 int16 = 32767
        _ = fi56
        _ = v57
    }
    var v58 float64 = float64(1.5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v52:", v52)
    fmt.Println("v58:", v58)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v58
    _ = v6
}
