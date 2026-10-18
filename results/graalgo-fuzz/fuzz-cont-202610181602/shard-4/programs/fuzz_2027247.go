package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 int16 }
type GenS2 struct { F0 uint64; F1 int8; F2 bool }
type GenS3 struct { F0 int32; F1 uint8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int16; F1 uint8 }
type GenS6 struct { F0 int64; F1 int64; F2 uint16 }
type GenS7 struct { F0 uint; F1 int32 }
func genFn0(p0 uint64) bool {
    for fi4 := 0; fi4 < 2; fi4++ {
        p0 = p0
        _ = fi4
    }
    fmt.Println(p0)
    _ = p0
    return (-128 < 100)
}
func genFn1(p0 int8, p1 uint16) uint8 {
    var v5 uint32 = 2147483647
    for fi6 := 0; fi6 < 17; fi6++ {
        v7 := GenS1{F0: 4294967295, F1: -32768, F2: 726}
        _ = fi6
        _ = v7
    }
    if true {
        for fi8 := 0; fi8 < 10; fi8++ {
            for fi9 := 0; fi9 < 1; fi9++ {
                if (uint32(p1) > (^(v5))) {
                    if (499 > (p1 ^ 0)) {
                        var v10 uint8 = 208
                        if (2 <= v5) {
                            var v11 int16 = 0
                            p0 += (^(p0))
                            _ = v11
                        } else {
                            v12 := GenS2{F0: 18446744073709551615, F1: p0, F2: true}
                            var v13 uint = 2
                            v14 := GenS2{F0: 4294967295, F1: (p0 << ((4294967295) & 63)), F2: (!((!(p0 >= (p0 / ((p0) | 1)))) || (fi9 == fi8)))}
                            _ = v12
                            _ = v13
                            _ = v14
                        }
                        _ = v10
                    }
                    var v15 uint8 = 2
                    _ = v15
                }
                switch ("" + "hello") {
                case "a":
                    v5 += (^(v5))
                case "\n\t":
                    if true {
                        var v16 float64 = 0.1
                        v16 += 1.5
                        _ = v16
                    }
                }
                v5 += (^(v5))
                _ = fi9
            }
            v17 := []int{}
            var v18 uint32 = v5
            _ = fi8
            _ = v17
            _ = v18
        }
        var v19 string = "foo bar"
        _ = v19
    } else {
        for fi20 := 0; fi20 < 8; fi20++ {
            if (!(((3.14159 + 1.5) != 0.1) && false)) {
                v21 := []int{}
                for fi22 := 0; fi22 < 13; fi22++ {
                    if len(v21) > 0 {
                        v21[0] = fi22
                    }
                    v23 := GenS3{F0: -2147483648, F1: 2}
                    var v24 bool = (!true)
                    v25 := GenS2{F0: 46, F1: p0, F2: v24}
                    _ = fi22
                    _ = v23
                    _ = v24
                    _ = v25
                }
                v21 = append(v21, 9223372036854775807)
                _ = v21
            }
            for fi26 := 0; fi26 < 9; fi26++ {
                switch "hello" {
                case "\n\t":
                    v5 = 65535
                case "a":
                    switch "hello" {
                    case "hello":
                        fmt.Println(p1)
                    default:
                        v27 := GenS3{F0: -1, F1: 0}
                        _ = v27
                    }
                    var v28 int = (fi20 % (((fi26 * (fi20 % ((fi20) | 1)))) | 1))
                    _ = v28
                default:
                    var v29 int8 = -1
                    _ = v29
                }
                var v30 bool = true
                _ = fi26
                _ = v30
            }
            v31 := GenS1{F0: 208, F1: 1000, F2: -270}
            _ = fi20
            _ = v31
        }
        if (127 == 48) {
            p1 -= p1
        }
        switch ("\n\t" + "") {
        case "a":
            v5 += v5
        case "":
            p0 = (p0 + (p0 | p0))
        default:
            var v32 float64 = 1.5
            _ = v32
        }
    }
    _ = p0
    _ = p1
    _ = v5
    return 0
}

func main() {
    v1 := make(map[string]int)
    var v2 uint16 = 42
    var v3 uint = 0
    v33 := GenS4{F0: 0}
    if (false && (!true)) {
        if (!false) {
            for fi34 := 0; fi34 < 20; fi34++ {
                if ((!true) || (-1 >= 820)) {
                    v3 += v3
                } else {
                    v35 := GenS3{F0: 163, F1: uint8(v2)}
                    switch 509 {
                    case 32768:
                        for fi36 := 0; fi36 < 10; fi36++ {
                            v2 -= v2
                            switch "" {
                            case "a":
                                var v37 uint16 = (^(v2))
                                _ = v37
                            case "\n\t":
                                v3 = v3
                            case "hello":
                                v2 += (^(v2))
                            default:
                                if false {
                                    v38 := GenS2{F0: 521, F1: 0, F2: (!false)}
                                    _ = v38
                                }
                            }
                            _ = fi36
                        }
                    }
                    fmt.Println(v2)
                    _ = v35
                }
                var v39 int16 = 0
                v40 := GenS2{F0: 4294967295, F1: -76, F2: true}
                _ = fi34
                _ = v39
                _ = v40
            }
        } else {
            fmt.Println(v3)
            if (((!true) || (2147483648 == 2147483647)) || (v3 <= v3)) {
                var v41 int64 = 9223372036854775807
                v2 = v2
                if true {
                    var v42 uint32 = uint32(v2)
                    switch "foo bar" {
                    case "hello":
                        var v43 float64 = (float64(1e10) / (1e10 + -2.25))
                        _ = v43
                    }
                    v41 = v41
                    _ = v42
                }
                _ = v41
            } else {
                for fi44 := 0; fi44 < 14; fi44++ {
                    v2 -= 2
                    v45 := GenS5{F0: 32767, F1: 1}
                    fmt.Println(fi44, v3)
                    _ = fi44
                    _ = v45
                }
            }
            var v46 uint = v3
            _ = v46
        }
    } else {
        var v47 bool = true
        fmt.Println(v47, v2)
        fmt.Println(v47)
        _ = v47
    }
    var v48 uint32 = 2147483648
    v49 := GenS6{F0: 42, F1: 70, F2: 65535}
    var v50 uint = 1
    var v51 uint8 = 101
    var v52 uint = (v3 & (v3 * (v3 << ((v50) & 63))))
    v53 := []int{32767}
    for fi54 := 0; fi54 < 8; fi54++ {
        v55 := GenS5{F0: 1000, F1: v51}
        _ = fi54
        _ = v55
    }
    fmt.Println(v48, v3, v50)
    switch -2147483648 {
    case 1:
        fmt.Println(v52)
    case 0:
        var v56 int16 = 1000
        _ = v56
    default:
        v57 := GenS7{F0: 322, F1: -1}
        _ = v57
    }
    if false {
        _ = len(v1)
    }
    _ = len(v53)
    v51 += v51
    var v58 uint16 = v2
    _ = len(v53)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v48:", v48)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("slicelen v53:", len(v53), v53)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v48
    _ = v49
    _ = v50
    _ = v51
    _ = v52
    _ = v53
    _ = v58
}
