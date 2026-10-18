package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 uint64; F1 uint16; F2 int64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint8; F1 int }
type GenS6 struct { F0 uint16; F1 int8 }
func genFn0(p0 int64, p1 int32) uint8 {
    v5 := []int{1}
    p0 = (p0 + -1)
    fmt.Println(p0)
    for fi6 := 0; fi6 < 16; fi6++ {
        v7 := GenS1{F0: 18446744073709551615}
        if len(v5) > 0 {
            v5[0] = -157
        }
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v5
    return 200
}
func genFn1(p0 int, p1 uint32, p2 string) int {
    var v8 string = "hello"
    switch p2 {
    case "\n\t":
        fmt.Println(v8)
    case "":
        p1 = (p1 << ((255) & 63))
    default:
        v9 := []int{}
        _ = v9
    }
    switch "hello" {
    case "\n\t":
        var v10 uint64 = 42
        _ = v10
    default:
        for fi11 := 0; fi11 < 3; fi11++ {
            var v12 int = (p0 * fi11)
            _ = fi11
            _ = v12
        }
    }
    p1 += 381
    var v13 int64 = 9223372036854775807
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v8
    return (p0 ^ (-(p0)))
}

func main() {
    v1 := []int{0, 128}
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]bool)
    v4 := make(map[string]string)
    v14 := GenS2{F0: 4294967295, F1: "foo bar"}
    if (!false) {
        var v15 uint64 = 4294967295
        if true {
            switch -2147483648 {
            case 0:
                if ("a" != "") {
                    v15 = v15
                    v16 := GenS2{F0: 627, F1: ("hello" + (("\n\t" + "a") + "foo bar"))}
                    _ = v16
                } else {
                    switch 32768 {
                    case -440:
                        fmt.Println(v15)
                    case -6:
                        v15 -= v15
                    case 2147483647:
                        switch "hello" {
                        case "hello":
                            for fi17 := 0; fi17 < 7; fi17++ {
                                v18 := GenS1{F0: 4294967295}
                                var v19 int32 = -2147483648
                                v4[""] = "a"
                                _ = fi17
                                _ = v18
                                _ = v19
                            }
                        case "\n\t":
                            switch -2147483648 {
                            case -9223372036854775808:
                                fmt.Println(v15)
                            default:
                                v15 -= (v15 + 983)
                            }
                        case "a":
                            v15 += v15
                        }
                    }
                    for fi20 := 0; fi20 < 2; fi20++ {
                        v15 += (v15 | v15)
                        switch "foo bar" {
                        case "":
                            fmt.Println(fi20)
                        case "a":
                            var v21 int32 = -374
                            _ = v21
                        case "foo bar":
                            v22 := v3[42]
                            _ = v22
                        }
                        v4[(("foo bar" + "foo bar") + "hello")] = ""
                        _ = fi20
                    }
                    v23 := GenS1{F0: v15}
                    _ = v23
                }
            case 693:
                for fi24 := 0; fi24 < 19; fi24++ {
                    var v25 int64 = -1
                    _ = fi24
                    _ = v25
                }
            case 1:
                for fi26 := 0; fi26 < 20; fi26++ {
                    v27 := GenS1{F0: (v15 / ((v15) | 1))}
                    if true {
                        fmt.Println(fi26)
                        var v28 int8 = 42
                        if (65535 >= 255) {
                            v15 += v15
                            var v29 uint8 = 1
                            _ = v29
                        } else {
                            v30 := GenS1{F0: 69}
                            v28 += v28
                            _ = v30
                        }
                        _ = v28
                    } else {
                        switch (fi26 << ((0) & 63)) {
                        case -982:
                            v31 := GenS3{F0: v15, F1: 255, F2: -161}
                            _ = v31
                        case -1:
                            for fi32 := 0; fi32 < 20; fi32++ {
                                v33 := GenS1{F0: (v15 | 0)}
                                if true {
                                    v1 = append(v1, genFn1((fi32 + -1), 2147483648, "a"))
                                    v15 -= (v15 % (((v15 << ((255) & 63))) | 1))
                                } else {
                                    v34 := GenS4{F0: 128}
                                    v35 := GenS5{F0: 127, F1: 0}
                                    var v36 string = "foo bar"
                                    _ = v34
                                    _ = v35
                                    _ = v36
                                }
                                fmt.Println(fi32, fi32)
                                fmt.Println(v15)
                                v15 += v15
                                var v37 int32 = 42
                                var v38 float64 = 0.1
                                _ = fi32
                                _ = v33
                                _ = v37
                                _ = v38
                            }
                        default:
                            switch "a" {
                            case "a":
                                if false {
                                    delete(v3, 42)
                                    for fi39 := 0; fi39 < 13; fi39++ {
                                        var v40 uint = 18446744073709551615
                                        v41 := GenS3{F0: 42, F1: 778, F2: 1}
                                        _ = fi39
                                        _ = v40
                                        _ = v41
                                    }
                                } else {
                                    v42 := GenS5{F0: 42, F1: (fi26 >> ((4294967295) & 63))}
                                    delete(v2, fi26)
                                    for fi43 := 0; fi43 < 9; fi43++ {
                                        fmt.Println(fi26, v15, fi26)
                                        fmt.Println(v15)
                                        var v44 int8 = 100
                                        _ = fi43
                                        _ = v44
                                    }
                                    _ = v42
                                }
                            }
                        }
                    }
                    fmt.Println(v15, fi26)
                    _ = fi26
                    _ = v27
                }
            default:
                v45 := GenS6{F0: 255, F1: 125}
                _ = v45
            }
            var v46 uint64 = 2
            v47 := GenS3{F0: (v15 % (((v15 << ((255) & 63))) | 1)), F1: 255, F2: -1}
            _ = v46
            _ = v47
        }
        _ = v15
    }
    if (42 < 973) {
        if (!true) {
            if (!false) {
                v48 := GenS1{F0: 906}
                var v49 int32 = -588
                _ = v48
                _ = v49
            }
            switch 1 {
            case 127:
                switch -9223372036854775808 {
                case 1:
                    v1 = append(v1, 32767)
                }
            }
        } else {
            for fi50 := 0; fi50 < 11; fi50++ {
                for fi51 := 0; fi51 < 5; fi51++ {
                    var v52 int8 = -108
                    _ = fi51
                    _ = v52
                }
                _ = fi50
            }
        }
        var v53 int8 = 0
        v4["\n\t"] = ("" + "")
        _ = v53
    }
    var v54 int16 = 1000
    v54 += (v54 & (v54 * 1))
    v54 = (v54 * v54)
    v54 -= v54
    var v55 uint64 = 724
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v54
    _ = v55
}
