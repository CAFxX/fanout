package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 uint; F1 int32 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint8 }

func main() {
    v1 := make(map[int]string)
    var v2 bool = (!false)
    v3 := make(map[string]int)
    var v4 uint8 = 127
    if ((3.14159 >= ((1e10 + float64(1e10)) / 1.0)) && v2) {
        var v5 bool = (!(-1 > -144))
        _ = v5
    } else {
        v6 := GenS1{F0: (((1e10 + 3.14159) > 3.14159) && true), F1: v4, F2: ((-1 < 9223372036854775807) && false)}
        _ = v6
    }
    v7 := GenS2{F0: 0}
    v4 -= 1
    if (0 > 2) {
        for fi8 := 0; fi8 < 2; fi8++ {
            fmt.Println(fi8, v4, fi8)
            _ = fi8
        }
        v9 := GenS3{F0: 790, F1: -894}
        if false {
            v10 := GenS4{F0: 2, F1: 1}
            _ = v10
        }
        _ = v9
    } else {
        switch -2147483648 {
        case 1:
            for fi11 := 0; fi11 < 1; fi11++ {
                for fi12 := 0; fi12 < 4; fi12++ {
                    fmt.Println(v4, v4, fi12)
                    _ = fi12
                }
                if true {
                    for fi13 := 0; fi13 < 19; fi13++ {
                        v14 := v1[815]
                        if (!v2) {
                            if false {
                                v15 := []int{}
                                _ = v15
                            } else {
                                if true {
                                    if (3.14159 > 0.0) {
                                        v2 = (-1 <= 979)
                                        var v16 int64 = 0
                                        var v17 int16 = 1
                                        _ = v16
                                        _ = v17
                                    }
                                    if false {
                                        v2 = (v2 && (!true))
                                        v18 := GenS5{F0: 1}
                                        _ = v18
                                    }
                                    v2 = false
                                }
                            }
                            v4 += 128
                        } else {
                            if (-802 >= fi13) {
                                _ = len(v3)
                                var v19 int = (fi11 << ((827) & 63))
                                v3[v14] = -180
                                _ = v19
                            } else {
                                var v20 int = fi13
                                var v21 float32 = 1.5
                                _ = v20
                                _ = v21
                            }
                            v4 -= 31
                            if false {
                                switch ((v14 + "hello") + v14) {
                                case "hello":
                                    v2 = v2
                                }
                            }
                        }
                        _ = fi13
                        _ = v14
                    }
                } else {
                    switch "" {
                    case "hello":
                        v22 := []int{(fi11 / (((fi11 << ((18446744073709551615) & 63))) | 1)), 2147483647, fi11, -440}
                        _ = v22
                    default:
                        switch (fi11 & (fi11 >> ((uint(v4)) & 63))) {
                        case 1:
                            v2 = v2
                        }
                    }
                    v1[fi11] = "hello"
                }
                v23 := GenS6{F0: (v4 * v4)}
                _ = fi11
                _ = v23
            }
        default:
            var v24 uint32 = 367
            _ = v24
        }
        var v25 bool = v2
        _ = v25
    }
    var v26 uint16 = 42
    var v27 int16 = 0
    _ = len(v3)
    if (!(v27 <= -839)) {
        var v28 int8 = 1
        for fi29 := 0; fi29 < 4; fi29++ {
            var v30 int = -9223372036854775808
            _ = fi29
            _ = v30
        }
        if (-9223372036854775808 <= 32767) {
            for fi31 := 0; fi31 < 16; fi31++ {
                fmt.Println(v4, v27)
                _ = fi31
            }
            delete(v3, "hello")
        } else {
            switch ("foo bar" + "\n\t") {
            case "a":
                var v32 uint = uint(v26)
                _ = v32
            default:
                _ = len(v1)
            }
            fmt.Println(v2, v4, v2)
        }
        _ = v28
    } else {
        v33 := []int{128, 0, -2147483648, 32767}
        if (-1 < 762) {
            _ = len(v3)
        }
        v4 += (v4 / ((v4) | 1))
        _ = v33
    }
    for fi34 := 0; fi34 < 5; fi34++ {
        fmt.Println(fi34, v2)
        _ = fi34
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v7
}
