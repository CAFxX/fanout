package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64; F1 int64; F2 int8 }
type GenS3 struct { F0 bool; F1 uint64; F2 uint }
type GenS4 struct { F0 int64; F1 int32; F2 uint8 }
type GenS5 struct { F0 uint; F1 int; F2 int32 }
func genFn0(p0 uint16) uint32 {
    fmt.Println(p0)
    for fi5 := 0; fi5 < 14; fi5++ {
        v6 := GenS1{F0: true}
        for fi7 := 0; fi7 < 6; fi7++ {
            p0 -= 2
            fmt.Println(fi7, fi7, p0)
            switch "\n\t" {
            case "foo bar":
                switch "" {
                case "\n\t":
                    p0 = p0
                case "a":
                    v8 := GenS1{F0: (false && false)}
                    _ = v8
                }
            case "":
                p0 += 2
            }
            _ = fi7
        }
        if false {
            if false {
                v9 := GenS2{F0: 42, F1: 1, F2: 0}
                v10 := []int{288, -771, fi5, fi5}
                if len(v10) > 0 {
                    v10[0] = 2147483647
                }
                _ = v10
                _ = v9
            }
            for fi11 := 0; fi11 < 16; fi11++ {
                fmt.Println(fi5)
                if true {
                    var v12 float32 = float32(1.5)
                    p0 = (p0 & p0)
                    _ = v12
                } else {
                    var v13 uint8 = 111
                    _ = v13
                }
                switch fi11 {
                case 128:
                    v14 := []int{fi5, fi5}
                    p0 = p0
                    _ = v14
                case 242:
                    fmt.Println(p0, fi11, p0)
                default:
                    fmt.Println(fi11, fi5, p0)
                }
                if (!(false && (1e10 == ((-2.25 * (3.14159 / 1e10)) / 3.14159)))) {
                    p0 = (p0 * p0)
                    v15 := []int{(fi5 << ((1) & 63))}
                    p0 = (p0 >> ((2) & 63))
                    _ = v15
                }
                _ = fi11
            }
            for fi16 := 0; fi16 < 14; fi16++ {
                var v17 bool = (3.14159 == float32(0.1))
                switch "\n\t" {
                case "foo bar":
                    var v18 uint16 = 65535
                    switch fi5 {
                    case -747:
                        var v19 int32 = -1
                        _ = v19
                    case -991:
                        var v20 float64 = float64(1e10)
                        _ = v20
                    case 32767:
                        var v21 float32 = 3.14159
                        if (103 < 42) {
                            fmt.Println(v18, p0)
                            v22 := GenS3{F0: (false && (9223372036854775807 == 912)), F1: 465, F2: uint(v18)}
                            var v23 int32 = -1
                            _ = v22
                            _ = v23
                        } else {
                            switch "hello" {
                            case "foo bar":
                                var v24 string = "a"
                                _ = v24
                            }
                            var v25 int32 = 1
                            var v26 int64 = -9223372036854775808
                            _ = v25
                            _ = v26
                        }
                        fmt.Println(v17, fi5)
                        _ = v21
                    }
                    _ = v18
                }
                p0 = (p0 / ((32767) | 1))
                _ = fi16
                _ = v17
            }
        } else {
            for fi27 := 0; fi27 < 5; fi27++ {
                switch fi5 {
                case -9223372036854775808:
                    v28 := GenS1{F0: (false && ((128 >= 128) || false))}
                    _ = v28
                }
                fmt.Println(fi27)
                if true {
                    var v29 float32 = -2.25
                    v30 := GenS1{F0: true}
                    switch fi27 {
                    case 32768:
                        fmt.Println(fi5)
                    case -9223372036854775808:
                        v29 = 0.1
                    default:
                        v31 := GenS3{F0: false, F1: 380, F2: 4294967295}
                        _ = v31
                    }
                    _ = v29
                    _ = v30
                } else {
                    switch "foo bar" {
                    case "a":
                        p0 -= (^(p0))
                    case "foo bar":
                        switch ((("a" + "a") + ("foo bar" + "a")) + "foo bar") {
                        case "a":
                            if (float32((1e10 / 3.14159)) < 1e10) {
                                p0 -= (p0 / (((p0 | (^(p0)))) | 1))
                                v32 := GenS1{F0: true}
                                p0 = p0
                                _ = v32
                            } else {
                                v33 := []int{(fi5 << ((193) & 63)), -504, (-(fi5)), fi5}
                                if (629 == 2147483647) {
                                    p0 -= 255
                                } else {
                                    if len(v33) > 0 {
                                        v33[0] = (fi5 % (((-(fi27))) | 1))
                                    }
                                }
                                p0 += 731
                                p0 = (p0 ^ (^(p0)))
                                p0 -= p0
                                if (false && ((1 == -47) && (false && (!false)))) {
                                    var v34 uint32 = 2
                                    _ = v34
                                }
                                _ = v33
                            }
                        case "foo bar":
                            p0 -= p0
                        case "hello":
                            p0 -= p0
                            fmt.Println(p0, p0)
                            if (42 <= 4294967295) {
                                var v35 int16 = 0
                                _ = v35
                            } else {
                                v36 := GenS2{F0: 48, F1: -593, F2: -128}
                                var v37 uint = 1
                                _ = v36
                                _ = v37
                            }
                        default:
                            v38 := GenS2{F0: -9223372036854775808, F1: -1, F2: 127}
                            _ = v38
                        }
                    case "":
                        var v39 float64 = (3.14159 + float64(0.0))
                        _ = v39
                    }
                }
                _ = fi27
            }
        }
        _ = fi5
        _ = v6
    }
    for fi40 := 0; fi40 < 3; fi40++ {
        v41 := GenS4{F0: 9223372036854775807, F1: 0, F2: 42}
        var v42 uint32 = 0
        _ = fi40
        _ = v41
        _ = v42
    }
    var v43 uint = 990
    _ = p0
    _ = v43
    return 2
}

func main() {
    var v1 uint8 = 226
    v2 := make(map[uint32]int)
    var v3 uint16 = 24
    var v4 int16 = 0
    var v44 int32 = 42
    v4 = v4
    v1 += 242
    var v45 uint8 = (v1 | v1)
    v45 -= 241
    v4 = (v4 % (((^(v4))) | 1))
    var v46 bool = false
    fmt.Println(v44, v4)
    v47 := GenS5{F0: 1, F1: 9223372036854775807, F2: v44}
    fmt.Println(v1, v46, v4)
    var v48 uint32 = 4294967295
    v49 := []int{-364}
    if v46 {
        fmt.Println(v45, v3)
    } else {
        switch "a" {
        case "foo bar":
            fmt.Println(v48)
        }
        if len(v49) > 0 {
            v49[0] = 42
        }
    }
    fmt.Println(v48, v4)
    for fi50 := 0; fi50 < 15; fi50++ {
        var v51 float32 = (3.14159 * 1e10)
        fmt.Println(v48)
        v52 := v2[v48]
        _ = fi50
        _ = v51
        _ = v52
    }
    var v53 int32 = v44
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v48:", v48)
    fmt.Println("slicelen v49:", len(v49), v49)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v53
}
