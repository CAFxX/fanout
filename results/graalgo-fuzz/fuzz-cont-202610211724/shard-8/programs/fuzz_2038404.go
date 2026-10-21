package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool; F1 uint64 }
type GenS4 struct { F0 int16; F1 string; F2 int64 }
type GenS5 struct { F0 int; F1 int }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint64; F1 int8; F2 string }
type GenS8 struct { F0 int64; F1 bool; F2 bool }
func genFn0(p0 uint16, p1 int32) uint32 {
    p1 += 2147483647
    if false {
        var v6 int32 = 0
        for fi7 := 0; fi7 < 19; fi7++ {
            if true {
                switch "\n\t" {
                case "\n\t":
                    v8 := GenS1{F0: 0, F1: false}
                    _ = v8
                default:
                    v9 := GenS2{F0: 4294967295}
                    _ = v9
                }
                switch "\n\t" {
                case "foo bar":
                    switch "hello" {
                    case "":
                        fmt.Println(p1)
                    default:
                        p0 += (^(p0))
                    }
                case "hello":
                    v10 := []int{fi7, fi7, 154, fi7}
                    _ = v10
                case "\n\t":
                    v6 = (v6 >> ((255) & 63))
                default:
                    p1 += p1
                }
                v11 := []int{(fi7 >> ((908) & 63)), fi7, fi7}
                _ = v11
            } else {
                if true {
                    v6 -= (v6 << ((0) & 63))
                    switch 2147483647 {
                    case -2147483648:
                        switch "hello" {
                        case "foo bar":
                            var v12 uint16 = p0
                            _ = v12
                        case "a":
                            v13 := GenS3{F0: true, F1: uint64(p0)}
                            _ = v13
                        }
                    case 32768:
                        v14 := []int{}
                        _ = v14
                    default:
                        for fi15 := 0; fi15 < 4; fi15++ {
                            v16 := []int{}
                            _ = fi15
                            _ = v16
                        }
                    }
                }
                switch fi7 {
                case 9223372036854775807:
                    v17 := []int{(fi7 / ((fi7) | 1))}
                    _ = v17
                case 4294967295:
                    var v18 float64 = 1.5
                    _ = v18
                default:
                    for fi19 := 0; fi19 < 5; fi19++ {
                        var v20 uint32 = 4294967295
                        _ = fi19
                        _ = v20
                    }
                }
                if true {
                    var v21 uint8 = 254
                    _ = v21
                }
            }
            _ = fi7
        }
        _ = v6
    }
    for fi22 := 0; fi22 < 19; fi22++ {
        v23 := GenS1{F0: 850, F1: ("" != (("hello" + ((("hello" + "") + ("foo bar" + "hello")) + "a")) + "a"))}
        switch "" {
        case "\n\t":
            fmt.Println(fi22, p0)
        case "hello":
            v24 := GenS3{F0: (!false), F1: 2}
            _ = v24
        default:
            p1 = (p1 + 1)
        }
        var v25 uint = 0
        _ = fi22
        _ = v23
        _ = v25
    }
    var v26 float64 = 1e10
    for fi27 := 0; fi27 < 2; fi27++ {
        if true {
            p1 -= -1
        }
        _ = fi27
    }
    _ = p0
    _ = p1
    _ = v26
    return 65535
}
func genFn1(p0 int16, p1 int64, p2 uint8) int8 {
    v28 := GenS2{F0: 1}
    p2 = (^(p2))
    _ = p0
    _ = p1
    _ = p2
    _ = v28
    return 127
}

func main() {
    var v1 int = 1
    var v2 bool = true
    v3 := make(map[int]uint32)
    var v4 int64 = 9223372036854775807
    v5 := make(map[int]bool)
    v4 -= (v4 + (^(v4)))
    v29 := []int{-994, v1, v1, -2147483648}
    var v30 int = 42
    v1 += -189
    var v31 uint8 = 161
    v32 := GenS1{F0: 879, F1: false}
    switch "a" {
    case "a":
        for fi33 := 0; fi33 < 7; fi33++ {
            v30 = (fi33 + v1)
            _ = fi33
        }
    }
    v2 = v2
    _ = len(v5)
    if v2 {
        switch "foo bar" {
        case "hello":
            var v34 float64 = 1.5
            _ = v34
        }
    } else {
        if v2 {
            v35 := GenS4{F0: 42, F1: "a", F2: v4}
            fmt.Println(v2, v1, v1)
            _ = v35
        }
        for fi36 := 0; fi36 < 14; fi36++ {
            v37 := GenS5{F0: v1, F1: (fi36 * (fi36 / ((v1) | 1)))}
            for fi38 := 0; fi38 < 3; fi38++ {
                delete(v3, 4294967295)
                var v39 uint32 = genFn0(347, -2147483648)
                _ = fi38
                _ = v39
            }
            v40 := GenS6{F0: 42}
            _ = fi36
            _ = v37
            _ = v40
        }
        switch "\n\t" {
        case "hello":
            fmt.Println(v2, v1, v30)
        case "a":
            var v41 uint8 = v31
            _ = v41
        default:
            if v2 {
                for fi42 := 0; fi42 < 13; fi42++ {
                    fmt.Println(v1, v1)
                    v43 := GenS7{F0: 0, F1: genFn1(-171, v4, v31), F2: "hello"}
                    _ = fi42
                    _ = v43
                }
                fmt.Println(v4)
                v29 = append(v29, (v1 >> ((49) & 63)))
            } else {
                v44 := GenS8{F0: (v4 >> ((4294967295) & 63)), F1: (0 > -1), F2: (32767 > 0)}
                v45 := GenS1{F0: 332, F1: ((v1 << ((18446744073709551615) & 63)) == v1)}
                _ = v44
                _ = v45
            }
        }
    }
    v31 = 0
    v46 := GenS3{F0: true, F1: 4294967295}
    var v47 uint32 = 2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v46
    _ = v47
    _ = v5
}
