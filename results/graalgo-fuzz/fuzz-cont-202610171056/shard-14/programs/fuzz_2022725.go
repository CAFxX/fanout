package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 uint64 }
type GenS2 struct { F0 uint32; F1 uint; F2 bool }
type GenS3 struct { F0 uint; F1 int64; F2 int }
type GenS4 struct { F0 uint32; F1 int32; F2 string }
type GenS5 struct { F0 uint16; F1 uint16 }

func main() {
    var v1 uint = 306
    v2 := make(map[string]uint32)
    var v3 int32 = 1
    var v4 float32 = -2.25
    v5 := make(map[int]int)
    var v6 int32 = v3
    var v7 bool = (true || true)
    var v8 int16 = 32767
    for fi9 := 0; fi9 < 8; fi9++ {
        for fi10 := 0; fi10 < 11; fi10++ {
            if (!(65 <= 284)) {
                if v7 {
                    v11 := GenS1{F0: "hello", F1: (-2.25 > float32(0.0)), F2: 4294967295}
                    v12 := GenS2{F0: 2147483648, F1: 42, F2: (fi10 == 32768)}
                    v4 = v4
                    v13 := GenS2{F0: 2147483648, F1: v1, F2: v7}
                    _ = v11
                    _ = v12
                    _ = v13
                }
                if ((-(v3)) <= (^(v3))) {
                    v7 = (v1 != 1)
                    switch "foo bar" {
                    case "foo bar":
                        var v14 uint16 = 0
                        _ = v14
                    case "a":
                        v2["\n\t"] = 2147483648
                    }
                    v15 := []int{32768}
                    if v7 {
                        v5[fi9] = 127
                    } else {
                        switch "foo bar" {
                        case "foo bar":
                            var v16 uint8 = 2
                            _ = v16
                        case "hello":
                            v7 = (-167 < -9223372036854775808)
                        default:
                            if len(v15) > 0 {
                                v15[0] = fi9
                            }
                        }
                        if (v8 != (-(v8))) {
                            var v17 int8 = -128
                            _ = v17
                        } else {
                            _ = len(v15)
                        }
                    }
                    v4 = ((v4 / 1.0) * 1.5)
                    fmt.Println(fi10, fi9, fi10)
                    _ = v15
                }
                var v18 int8 = 95
                _ = v18
            } else {
                v19 := GenS2{F0: 65535, F1: (v1 >> ((v1) & 63)), F2: ((true || v7) && (true || v7))}
                fmt.Println(v8, v1, v6)
                _ = v19
            }
            if (42 > 32767) {
                switch (fi9 << (((v1 + v1)) & 63)) {
                case -9223372036854775808:
                    fmt.Println(fi10)
                case 9223372036854775807:
                    switch fi10 {
                    case 665:
                        if (false && false) {
                            delete(v2, "hello")
                            fmt.Println(fi9)
                        } else {
                            v20 := []int{}
                            var v21 float32 = 0.1
                            _ = v20
                            _ = v21
                        }
                    case 35:
                        var v22 int = (fi10 * (fi10 % ((0) | 1)))
                        _ = v22
                    case 1:
                        var v23 float64 = 1.5
                        _ = v23
                    default:
                        v24 := []int{}
                        _ = v24
                    }
                default:
                    v3 = v3
                }
                v1 = 4294967295
                var v25 float64 = 3.14159
                _ = v25
            } else {
                var v26 int32 = -1
                var v27 bool = (103 <= 54)
                v7 = true
                _ = v26
                _ = v27
            }
            v3 = (v6 & (-(v3)))
            _ = fi10
        }
        for fi28 := 0; fi28 < 4; fi28++ {
            v29 := GenS3{F0: (v1 * v1), F1: 0, F2: fi9}
            _ = fi28
            _ = v29
        }
        _ = fi9
    }
    v30 := GenS4{F0: 21, F1: v6, F2: ("hello" + "foo bar")}
    _ = len(v2)
    fmt.Println(v4, v7)
    if (1 <= -240) {
        v31 := []int{4294967295, -2147483648, 9223372036854775807, 0}
        for fi32 := 0; fi32 < 18; fi32++ {
            var v33 uint16 = 583
            _ = fi32
            _ = v33
        }
        var v34 uint8 = 1
        _ = v31
        _ = v34
    }
    delete(v2, "foo bar")
    v8 += (v8 >> (((v1 & v1)) & 63))
    v35 := []int{-1, 900}
    v6 += (v3 * (^(v6)))
    v36 := GenS5{F0: 255, F1: 316}
    for fi37 := 0; fi37 < 7; fi37++ {
        for fi38 := 0; fi38 < 7; fi38++ {
            v35 = append(v35, fi37)
            _ = fi38
        }
        if v7 {
            v8 -= v8
        }
        var v39 int16 = v8
        _ = fi37
        _ = v39
    }
    _ = len(v2)
    for fi40 := 0; fi40 < 6; fi40++ {
        var v41 int64 = 1
        _ = fi40
        _ = v41
    }
    for fi42 := 0; fi42 < 2; fi42++ {
        for fi43 := 0; fi43 < 17; fi43++ {
            fmt.Println(v3, v6)
            var v44 uint32 = 705
            var v45 bool = false
            var v46 uint8 = 255
            _ = fi43
            _ = v44
            _ = v45
            _ = v46
        }
        if v7 {
            var v47 uint8 = 0
            v48 := GenS3{F0: v1, F1: 9223372036854775807, F2: fi42}
            _ = v47
            _ = v48
        } else {
            v4 -= v4
        }
        if (!false) {
            var v49 int64 = -9223372036854775808
            switch "hello" {
            case "a":
                switch "\n\t" {
                case "a":
                    if len(v35) > 0 {
                        v35[0] = (fi42 - (-(fi42)))
                    }
                default:
                    v50 := GenS1{F0: (("hello" + "") + "a"), F1: false, F2: 0}
                    _ = v50
                }
            case "":
                if (v1 > v1) {
                    if ((v8 >> ((2) & 63)) < (v8 ^ (v8 << ((v1) & 63)))) {
                        var v51 bool = (38 != 18446744073709551615)
                        _ = v51
                    }
                    for fi52 := 0; fi52 < 8; fi52++ {
                        _ = len(v5)
                        v53 := GenS2{F0: 348, F1: (v1 - (v1 / ((v1) | 1))), F2: (3.14159 < 1e10)}
                        _ = fi52
                        _ = v53
                    }
                    for fi54 := 0; fi54 < 1; fi54++ {
                        v6 = 0
                        _ = fi54
                    }
                }
            default:
                if (!false) {
                    _ = len(v2)
                }
            }
            _ = v49
        }
        _ = fi42
    }
    v35 = append(v35, -1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
