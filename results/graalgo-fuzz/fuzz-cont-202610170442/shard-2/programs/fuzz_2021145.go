package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 int32; F1 string }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int32; F1 int16 }

func main() {
    v1 := make(map[string]int)
    var v2 float32 = -2.25
    var v3 int8 = 1
    v4 := make(map[string]bool)
    var v5 float32 = float32(1.5)
    var v6 string = ("" + "hello")
    var v7 uint = 898
    switch -399 {
    case 747:
        v5 -= v2
    default:
        if true {
            var v8 uint = v7
            v9 := GenS1{F0: 42}
            _ = v8
            _ = v9
        } else {
            v10 := GenS1{F0: 32767}
            v11 := []int{1, -511}
            v12 := GenS1{F0: 559}
            _ = v10
            _ = v11
            _ = v12
        }
    }
    var v13 int8 = v3
    var v14 int = 119
    if ((!false) || (((v5 / v2) - v2) <= float32(0.0))) {
        _ = len(v1)
        v14 -= 277
    } else {
        if ((v6 + v6) >= v6) {
            var v15 int32 = -1
            _ = v15
        }
        v16 := GenS2{F0: 111, F1: uint16(v7)}
        _ = v16
    }
    var v17 int8 = (v3 + v3)
    if (false && (v5 <= v5)) {
        var v18 int16 = 42
        fmt.Println(v5, v2)
        fmt.Println(v13, v17, v7)
        _ = v18
    }
    fmt.Println(v2, v7)
    var v19 int16 = -32768
    if false {
        for fi20 := 0; fi20 < 6; fi20++ {
            if true {
                v5 = v2
                if (((true && true) || (2 != 0)) && (!((true || true) && false))) {
                    var v21 bool = ((18446744073709551615 < 1) && (!true))
                    if true {
                        _ = len(v4)
                        for fi22 := 0; fi22 < 11; fi22++ {
                            v6 = ((v6 + v6) + v6)
                            fmt.Println(v6)
                            var v23 float32 = 1.5
                            if v21 {
                                var v24 int16 = 42
                                if v21 {
                                    v25 := GenS1{F0: fi20}
                                    _ = v25
                                }
                                if (fi20 != -9223372036854775808) {
                                    v5 -= v5
                                    v6 = ""
                                    var v26 float32 = (((float32(-2.25) / ((0.0 + v5) + (v2 * v23))) / v5) + v23)
                                    _ = v26
                                }
                                _ = v24
                            }
                            _ = fi22
                            _ = v23
                        }
                        v3 -= 100
                    }
                    var v27 float32 = float32(1e10)
                    _ = v21
                    _ = v27
                } else {
                    var v28 uint64 = 0
                    _ = v28
                }
            } else {
                v29 := GenS3{F0: -2147483648, F1: v6}
                switch (v14 >> (((^(v7))) & 63)) {
                case 32768:
                    var v30 int16 = (v19 ^ (v19 << ((v7) & 63)))
                    _ = v30
                }
                switch v6 {
                case "foo bar":
                    for fi31 := 0; fi31 < 8; fi31++ {
                        var v32 uint8 = 190
                        _ = fi31
                        _ = v32
                    }
                case "\n\t":
                    switch int(v19) {
                    case 0:
                        switch "a" {
                        case "a":
                            switch "\n\t" {
                            case "foo bar":
                                v19 = 1
                            case "a":
                                v7 += (v7 & v7)
                            }
                        case "hello":
                            var v33 int32 = 415
                            _ = v33
                        case "":
                            for fi34 := 0; fi34 < 19; fi34++ {
                                v35 := GenS4{F0: v7}
                                v3 = v17
                                _ = fi34
                                _ = v35
                            }
                        default:
                            v5 += v2
                        }
                    case 32768:
                        if true {
                            v14 = (-(fi20))
                            for fi36 := 0; fi36 < 5; fi36++ {
                                var v37 int64 = -9223372036854775808
                                _ = fi36
                                _ = v37
                            }
                        } else {
                            v14 = -17
                            for fi38 := 0; fi38 < 1; fi38++ {
                                v19 -= -321
                                v39 := GenS2{F0: v13, F1: 1}
                                fmt.Println(v14, v19)
                                v17 += 113
                                _ = fi38
                                _ = v39
                            }
                            v40 := GenS5{F0: -790, F1: v19}
                            _ = v40
                        }
                    case 2147483647:
                        if ((v19 == v19) || (!true)) {
                            v41 := []int{(v14 & (fi20 << ((2) & 63))), (fi20 / ((v14) | 1)), (v14 >> ((v7) & 63))}
                            _ = v41
                        }
                    default:
                        v6 = v6
                    }
                case "":
                    v42 := GenS2{F0: (v13 * 100), F1: 42}
                    _ = v42
                }
                _ = v29
            }
            v17 -= 100
            _ = fi20
        }
    } else {
        if true {
            v43 := []int{int(v7), -357, 42}
            var v44 int8 = v13
            v44 += v13
            _ = v43
            _ = v44
        }
    }
    v45 := []int{4294967295, 32767}
    var v46 bool = (255 != 0)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v45:", len(v45), v45)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v6
    _ = v7
}
