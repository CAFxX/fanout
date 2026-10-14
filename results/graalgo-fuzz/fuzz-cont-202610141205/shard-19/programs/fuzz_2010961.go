package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint32; F1 uint8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 string; F1 uint64 }
func genFn0() uint32 {
    var v6 float64 = (3.14159 + float64(0.0))
    var v7 bool = true
    v7 = (1.5 >= -2.25)
    for fi8 := 0; fi8 < 12; fi8++ {
        var v9 uint32 = 65535
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: 1}
    _ = v10
    _ = v6
    _ = v7
    return 2147483648
}
func genFn1(p0 uint16, p1 int, p2 int64) uint {
    p0 += p0
    if false {
        p0 = 873
        v11 := GenS1{F0: -1}
        _ = v11
    }
    if (true || true) {
        v12 := []int{42}
        _ = v12
    } else {
        if (p1 > p1) {
            v13 := []int{32768}
            for fi14 := 0; fi14 < 11; fi14++ {
                if false {
                    v15 := GenS2{F0: 854}
                    for fi16 := 0; fi16 < 11; fi16++ {
                        switch fi16 {
                        case -2147483648:
                            var v17 float32 = (0.0 / 1.5)
                            _ = v17
                        default:
                            p2 = p2
                        }
                        if false {
                            switch (-(p1)) {
                            case -9223372036854775808:
                                fmt.Println(fi14, p1)
                            case -102:
                                if false {
                                    var v18 float64 = (0.1 - 1e10)
                                    _ = v18
                                } else {
                                    var v19 int64 = 42
                                    _ = v19
                                }
                            case 214:
                                fmt.Println(p0, p2)
                            }
                            if len(v13) >= 2 {
                                v13 = v13[1:len(v13)]
                            }
                        }
                        switch 4294967295 {
                        case 81:
                            var v20 uint8 = 123
                            _ = v20
                        case 405:
                            if (-2.25 > 0.0) {
                                p0 += 32767
                            } else {
                                fmt.Println(p1, p1)
                            }
                            p2 = -9223372036854775808
                            var v21 float32 = 1.5
                            _ = v21
                        default:
                            v22 := GenS2{F0: 18446744073709551615}
                            _ = v22
                        }
                        _ = fi16
                    }
                    for fi23 := 0; fi23 < 11; fi23++ {
                        switch "\n\t" {
                        case "":
                            v24 := GenS3{F0: 65535, F1: 57}
                            _ = v24
                        default:
                            p1 += 2147483647
                            var v25 uint16 = 481
                            if len(v13) > 0 {
                                v13[0] = 32767
                            }
                            _ = v25
                        }
                        if true {
                            p2 = (-(p2))
                        } else {
                            p1 -= (fi14 * (p1 - 42))
                            p2 = (p2 - p2)
                        }
                        if len(v13) >= 2 {
                            v13 = v13[1:len(v13)]
                        }
                        _ = fi23
                    }
                    _ = v15
                } else {
                    for fi26 := 0; fi26 < 20; fi26++ {
                        p2 += p2
                        switch "foo bar" {
                        case "\n\t":
                            switch "\n\t" {
                            case "":
                                p1 = fi26
                            case "\n\t":
                                v27 := GenS4{F0: 127}
                                _ = v27
                            case "hello":
                                fmt.Println(p1)
                            default:
                                if false {
                                    var v28 uint8 = 128
                                    _ = v28
                                } else {
                                    v29 := GenS3{F0: 42, F1: 1}
                                    _ = v29
                                }
                            }
                        case "hello":
                            p0 = 0
                        default:
                            if len(v13) >= 2 {
                                v13 = v13[1:len(v13)]
                            }
                        }
                        _ = fi26
                    }
                    if (1 == 127) {
                        if true {
                            v30 := GenS5{F0: -58}
                            _ = v30
                        }
                        p2 = (p2 * (p2 | p2))
                        p1 += -596
                    }
                    for fi31 := 0; fi31 < 8; fi31++ {
                        p1 += (fi31 - fi14)
                        p2 += (p2 % ((p2) | 1))
                        p2 -= (p2 - -9223372036854775808)
                        fmt.Println(p2)
                        if true {
                            switch "hello" {
                            case "foo bar":
                                var v32 float32 = 1e10
                                _ = v32
                            case "hello":
                                var v33 int = (fi14 << ((573) & 63))
                                _ = v33
                            default:
                                p0 = 207
                            }
                            var v34 uint16 = (p0 ^ p0)
                            _ = v34
                        } else {
                            fmt.Println(fi14, fi31, fi14)
                        }
                        _ = fi31
                    }
                }
                if (4294967295 != 65535) {
                    fmt.Println(fi14, p1)
                }
                var v35 uint32 = 2147483647
                _ = fi14
                _ = v35
            }
            _ = v13
        }
        var v36 uint32 = 865
        var v37 uint = 0
        _ = v36
        _ = v37
    }
    v38 := []int{p1, 1, p1, 32768}
    _ = p0
    _ = p1
    _ = p2
    _ = v38
    return 15
}

func main() {
    var v1 bool = true
    var v2 string = "foo bar"
    v3 := make(map[int]bool)
    v4 := make(map[uint32]int)
    v5 := []int{-1, 127, 4294967295, -711}
    v39 := GenS6{F0: (v2 + v2)}
    v40 := GenS2{F0: 4294967295}
    var v41 int64 = -260
    if false {
        var v42 int64 = 913
        _ = v42
    }
    v5 = append(v5, 9223372036854775807)
    var v43 string = v2
    switch ("hello" + (v2 + ("\n\t" + "a"))) {
    case "":
        for fi44 := 0; fi44 < 11; fi44++ {
            var v45 int16 = -748
            _ = fi44
            _ = v45
        }
    case "\n\t":
        var v46 int8 = 0
        _ = v46
    default:
        v47 := v4[580]
        _ = v47
    }
    v48 := GenS7{F0: (v43 + v43), F1: 2}
    v49 := v3[102]
    v43 = (v2 + v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v43
    _ = v48
    _ = v49
    _ = v5
}
