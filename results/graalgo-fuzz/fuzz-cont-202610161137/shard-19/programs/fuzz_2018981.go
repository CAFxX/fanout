package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int; F1 int32; F2 string }
type GenS3 struct { F0 bool; F1 string; F2 uint64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint8; F1 int64; F2 int8 }
func genFn0(p0 int8, p1 uint64) uint {
    if false {
        if true {
            if false {
                for fi5 := 0; fi5 < 20; fi5++ {
                    for fi6 := 0; fi6 < 16; fi6++ {
                        p1 += (p1 ^ p1)
                        _ = fi6
                    }
                    _ = fi5
                }
            }
            v7 := []int{812, -948, 744, -1}
            _ = v7
        } else {
            for fi8 := 0; fi8 < 1; fi8++ {
                var v9 int8 = (p0 % ((p0) | 1))
                p0 = 1
                _ = fi8
                _ = v9
            }
            for fi10 := 0; fi10 < 14; fi10++ {
                switch "hello" {
                case "foo bar":
                    v11 := GenS1{F0: 1}
                    _ = v11
                case "a":
                    switch fi10 {
                    case -9223372036854775808:
                        var v12 uint32 = 2147483648
                        _ = v12
                    case 835:
                        v13 := []int{-287}
                        _ = v13
                    default:
                        p0 += p0
                    }
                default:
                    switch 2147483647 {
                    case -1:
                        var v14 bool = false
                        _ = v14
                    case 42:
                        var v15 float32 = 0.0
                        _ = v15
                    default:
                        if false {
                            fmt.Println(fi10, fi10)
                            p1 = p1
                            v16 := []int{fi10, (fi10 - fi10), 2147483647, fi10}
                            _ = v16
                        }
                    }
                }
                fmt.Println(p1, p0)
                _ = fi10
            }
        }
        v17 := []int{42, 128, 620, 852}
        if len(v17) > 0 {
            v17[0] = -1
        }
        _ = v17
    }
    p0 = p0
    for fi18 := 0; fi18 < 11; fi18++ {
        for fi19 := 0; fi19 < 15; fi19++ {
            fmt.Println(fi18, fi18)
            p0 = p0
            _ = fi19
        }
        var v20 float32 = 0.0
        v21 := GenS2{F0: (fi18 / ((fi18) | 1)), F1: -162, F2: "\n\t"}
        _ = fi18
        _ = v20
        _ = v21
    }
    v22 := []int{32768, 870, 42, -2147483648}
    _ = p0
    _ = p1
    _ = v22
    return 77
}

func main() {
    v1 := []uint32{1, 761, 2}
    var v2 string = "\n\t"
    var v3 string = v2
    var v4 uint16 = 1
    if (true || false) {
        switch -4 {
        case -72:
            var v23 uint16 = (^(v4))
            _ = v23
        default:
            var v24 uint16 = (v4 + (v4 & 173))
            _ = v24
        }
        var v25 int16 = 32767
        v26 := GenS3{F0: false, F1: v3, F2: 244}
        _ = v25
        _ = v26
    }
    v27 := GenS1{F0: (^(v4))}
    var v28 string = v2
    if true {
        if (!(120 <= 128)) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            for fi29 := 0; fi29 < 6; fi29++ {
                var v30 uint8 = 0
                if (!("\n\t" <= v28)) {
                    switch v2 {
                    case "hello":
                        for fi31 := 0; fi31 < 10; fi31++ {
                            if (!(-1 <= 42)) {
                                v32 := GenS3{F0: (18446744073709551615 <= 702), F1: v28, F2: 0}
                                var v33 float64 = (1e10 - 1.5)
                                var v34 float32 = 0.1
                                _ = v32
                                _ = v33
                                _ = v34
                            } else {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                v35 := GenS4{F0: v2}
                                v36 := GenS5{F0: (2 <= 4294967295)}
                                _ = v35
                                _ = v36
                            }
                            if ((false && (((!false) || true) && ((fi31 + 32768) != -9223372036854775808))) && true) {
                                if false {
                                    switch v2 {
                                    case "a":
                                        _ = len(v1)
                                        if (8 >= 710) {
                                            var v37 uint64 = 2
                                            v38 := GenS6{F0: (v30 >> ((255) & 63)), F1: 1, F2: 63}
                                            fmt.Println(v2, v30)
                                            _ = len(v1)
                                            _ = v37
                                            _ = v38
                                        } else {
                                            v39 := GenS3{F0: (false || false), F1: v2, F2: uint64(v4)}
                                            _ = v39
                                        }
                                    case "":
                                        v30 = 42
                                    default:
                                        v3 = (v3 + v2)
                                    }
                                } else {
                                    var v40 uint64 = 4294967295
                                    v1 = append(v1, 1)
                                    _ = v40
                                }
                                v2 = ("\n\t" + "")
                                var v41 int = fi31
                                fmt.Println(fi29, v30)
                                _ = v41
                            }
                            _ = fi31
                        }
                    }
                }
                _ = fi29
                _ = v30
            }
            if len(v1) > 0 {
                v1[0] = 2
            }
        }
        v4 += (^(v4))
    }
    v4 = (v4 + (v4 + (v4 >> ((genFn0(-15, 4294967295)) & 63))))
    for fi42 := 0; fi42 < 12; fi42++ {
        v3 = ("a" + "")
        fmt.Println(v3)
        for fi43 := 0; fi43 < 3; fi43++ {
            var v44 bool = true
            if true {
                var v45 uint64 = 2
                v1 = append(v1, 42)
                _ = v45
            }
            _ = fi43
            _ = v44
        }
        _ = fi42
    }
    if len(v1) > 0 {
        v1[0] = 72
    }
    v2 = "hello"
    v46 := GenS4{F0: "a"}
    fmt.Println(v4)
    var v47 float32 = -2.25
    var v48 uint16 = v4
    fmt.Println(v3, v47, v28)
    switch v2 {
    case "hello":
        v3 = v2
    case "foo bar":
        for fi49 := 0; fi49 < 4; fi49++ {
            var v50 bool = false
            _ = fi49
            _ = v50
        }
    case "\n\t":
        v47 = v47
    }
    var v51 int16 = -32768
    v47 = v47
    var v52 float64 = float64(v47)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v48
    _ = v51
    _ = v52
}
