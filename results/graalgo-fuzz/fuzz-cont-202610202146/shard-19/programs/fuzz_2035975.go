package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 string }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 uint64; F1 bool; F2 int16 }
type GenS4 struct { F0 int16; F1 int64 }
type GenS5 struct { F0 uint32; F1 uint16 }
type GenS6 struct { F0 int8; F1 uint64; F2 uint64 }
type GenS7 struct { F0 int16; F1 int32 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint32; F1 bool; F2 uint32 }
func genFn0(p0 int8, p1 uint16) int64 {
    for fi4 := 0; fi4 < 2; fi4++ {
        var v5 int64 = 42
        _ = fi4
        _ = v5
    }
    var v6 string = "\n\t"
    if false {
        var v7 int64 = 42
        _ = v7
    } else {
        fmt.Println(v6, p1, p0)
        v8 := GenS1{F0: -89, F1: p1, F2: (v6 + "a")}
        _ = v8
    }
    var v9 int16 = 13
    var v10 uint = 0
    _ = p0
    _ = p1
    _ = v10
    _ = v6
    _ = v9
    return 9223372036854775807
}
func genFn1(p0 uint8, p1 uint, p2 string) uint32 {
    fmt.Println(p2, p1, p1)
    for fi11 := 0; fi11 < 11; fi11++ {
        p1 = 4294967295
        _ = fi11
    }
    var v12 bool = (true || (!((true || (-148 >= -2147483648)) || (0 < 484))))
    p0 = (p0 - 0)
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 0
}

func main() {
    var v1 string = ("foo bar" + "a")
    var v2 uint64 = 0
    v3 := make(map[string]string)
    switch v1 {
    case "a":
        var v13 int32 = 42
        _ = v13
    case "hello":
        for fi14 := 0; fi14 < 11; fi14++ {
            var v15 int64 = 468
            v1 = (("\n\t" + "foo bar") + v1)
            var v16 int32 = -577
            _ = fi14
            _ = v15
            _ = v16
        }
    }
    v2 = 42
    v2 = 0
    var v17 int64 = 0
    var v18 uint = 0
    v17 += v17
    for fi19 := 0; fi19 < 8; fi19++ {
        v20 := GenS2{F0: 65535, F1: 0}
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 3; fi21++ {
        var v22 uint = (v18 >> ((65535) & 63))
        var v23 uint32 = 1
        v22 += v22
        _ = fi21
        _ = v22
        _ = v23
    }
    var v24 float32 = -2.25
    for fi25 := 0; fi25 < 13; fi25++ {
        for fi26 := 0; fi26 < 12; fi26++ {
            v17 = (v17 << ((v18) & 63))
            if (!true) {
                v18 = v18
                v27 := GenS1{F0: 0, F1: 905, F2: (v1 + v1)}
                _ = v27
            }
            v3[("hello" + v1)] = "foo bar"
            v28 := GenS2{F0: 2147483648, F1: 4294967295}
            _ = fi26
            _ = v28
        }
        _ = fi25
    }
    if true {
        delete(v3, (v1 + ""))
    } else {
        if true {
            switch v1 {
            case "hello":
                if false {
                    v2 = 164
                } else {
                    v29 := GenS2{F0: 141, F1: 1}
                    if true {
                        v17 = -1
                    }
                    _ = v29
                }
            default:
                switch 32767 {
                case 11:
                    switch 2147483647 {
                    case -9223372036854775808:
                        switch "a" {
                        case "foo bar":
                            var v30 float64 = 1e10
                            _ = v30
                        case "a":
                            for fi31 := 0; fi31 < 13; fi31++ {
                                if false {
                                    v32 := GenS2{F0: 2147483648, F1: v2}
                                    for fi33 := 0; fi33 < 17; fi33++ {
                                        var v34 int = fi31
                                        fmt.Println(fi33, v2, fi33)
                                        v1 = ((v1 + (v1 + v1)) + "a")
                                        _ = len(v3)
                                        v24 += -2.25
                                        _ = fi33
                                        _ = v34
                                    }
                                    _ = v32
                                }
                                v35 := GenS2{F0: 2147483648, F1: 1}
                                if (genFn1(uint8(v18), v18, v1) <= genFn1(128, 42, "")) {
                                    v36 := []int{(fi31 - (fi31 << ((v18) & 63))), fi31, (fi31 % ((0) | 1)), 42}
                                    delete(v3, "a")
                                    if true {
                                        if false {
                                            v18 = v18
                                        }
                                    }
                                    _ = v36
                                } else {
                                    v17 += -633
                                    var v37 string = ((v1 + "\n\t") + v1)
                                    _ = v37
                                }
                                _ = fi31
                                _ = v35
                            }
                        case "\n\t":
                            v2 += 42
                        default:
                            var v38 uint16 = 791
                            _ = v38
                        }
                    default:
                        var v39 uint8 = 1
                        _ = v39
                    }
                }
            }
        } else {
            v18 -= 2
            var v40 int = -53
            _ = v40
        }
    }
    var v41 uint = (v18 ^ (^(v18)))
    switch 4294967295 {
    case 32767:
        v42 := GenS2{F0: 1, F1: (v2 ^ v2)}
        _ = v42
    case -847:
        v43 := []int{32767, -9223372036854775808, 4294967295}
        _ = v43
    default:
        v44 := GenS3{F0: v2, F1: (false && false), F2: 42}
        _ = v44
    }
    var v45 int = 32767
    switch 1 {
    case 1:
        v46 := []int{(v45 ^ 32768)}
        _ = v46
    case 0:
        var v47 float32 = ((v24 - v24) * v24)
        _ = v47
    case -2147483648:
        v48 := []int{int(v18), (v45 << (((v18 << ((0) & 63))) & 63)), (v45 * 978), (^(v45))}
        _ = v48
    }
    if (!false) {
        if false {
            var v49 string = "hello"
            _ = v49
        } else {
            var v50 uint = (v18 | v18)
            v2 = (v2 + v2)
            _ = v50
        }
    } else {
        if false {
            for fi51 := 0; fi51 < 16; fi51++ {
                v52 := GenS4{F0: 1000, F1: (v17 ^ -1)}
                if false {
                    var v53 int64 = genFn0(100, 654)
                    if true {
                        if false {
                            v54 := GenS2{F0: 30, F1: v2}
                            _ = v54
                        }
                        v1 = v1
                        v55 := []int{int(v53), 830, v45, fi51}
                        _ = v55
                    }
                    _ = v53
                } else {
                    var v56 float64 = (1.5 / 1e10)
                    v57 := GenS5{F0: 287, F1: 0}
                    var v58 uint = v18
                    _ = v56
                    _ = v57
                    _ = v58
                }
                for fi59 := 0; fi59 < 8; fi59++ {
                    if true {
                        if ((!((-1 >= 0) || false)) && (-1 < 127)) {
                            v60 := GenS6{F0: -108, F1: 4294967295, F2: (^(v2))}
                            v17 += 398
                            switch (^(v45)) {
                            case -355:
                                v61 := GenS7{F0: -32768, F1: 2147483647}
                                _ = v61
                            case 313:
                                v2 -= (v2 & v2)
                                v3["a"] = ((v1 + ("hello" + "a")) + ("" + v1))
                                v62 := GenS8{F0: (^(fi51))}
                                _ = v62
                            case 21:
                                var v63 uint8 = 255
                                _ = v63
                            }
                            _ = v60
                        } else {
                            v64 := GenS5{F0: 469, F1: 313}
                            var v65 uint = 42
                            _ = v64
                            _ = v65
                        }
                        fmt.Println(v17, v1, v45)
                    } else {
                        var v66 uint8 = 128
                        fmt.Println(v17, fi51, v45)
                        _ = v66
                    }
                    var v67 uint8 = 2
                    v68 := GenS9{F0: 65535, F1: (((^(v41)) < 18446744073709551615) || false), F2: 860}
                    _ = fi59
                    _ = v67
                    _ = v68
                }
                _ = fi51
                _ = v52
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v41:", v41)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v24
    _ = v3
    _ = v41
    _ = v45
}
