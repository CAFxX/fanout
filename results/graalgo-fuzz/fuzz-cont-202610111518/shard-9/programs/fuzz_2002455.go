package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 uint8 }
type GenS2 struct { F0 int32; F1 uint; F2 int32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int64; F1 uint; F2 int32 }
type GenS5 struct { F0 uint }
func genFn0(p0 uint64, p1 uint) int8 {
    p0 = 884
    fmt.Println(p0, p0)
    switch 614 {
    case 760:
        if true {
            v6 := GenS1{F0: 391, F1: p0, F2: 2}
            _ = v6
        }
    case 1:
        switch "" {
        case "a":
            fmt.Println(p0)
        default:
            if false {
                v7 := []int{9223372036854775807, 2147483647}
                _ = v7
            } else {
                for fi8 := 0; fi8 < 18; fi8++ {
                    p0 = (p0 >> ((p1) & 63))
                    v9 := GenS1{F0: 1, F1: (p0 - p0), F2: 58}
                    _ = fi8
                    _ = v9
                }
            }
        }
    case 754:
        switch ("" + "hello") {
        case "":
            var v10 uint16 = 255
            _ = v10
        default:
            switch "" {
            case "hello":
                var v11 float32 = 1e10
                _ = v11
            case "":
                switch "" {
                case "a":
                    p1 += (p1 >> (((p1 >> (((p1 & p1)) & 63))) & 63))
                case "":
                    v12 := GenS2{F0: -483, F1: p1, F2: -2147483648}
                    _ = v12
                case "\n\t":
                    switch "" {
                    case "foo bar":
                        var v13 uint32 = 105
                        _ = v13
                    case "\n\t":
                        var v14 int64 = 0
                        _ = v14
                    case "":
                        v15 := GenS3{F0: (p0 & p0)}
                        _ = v15
                    }
                default:
                    fmt.Println(p0, p0)
                }
            case "\n\t":
                for fi16 := 0; fi16 < 1; fi16++ {
                    fmt.Println(fi16)
                    if false {
                        var v17 uint16 = 545
                        _ = v17
                    }
                    _ = fi16
                }
            }
        }
    default:
        if true {
            p1 += (p1 / (((p1 % ((p1) | 1))) | 1))
            var v18 int16 = -1
            v19 := GenS1{F0: -995, F1: 2, F2: 2}
            _ = v18
            _ = v19
        } else {
            p1 = 0
        }
    }
    var v20 string = ("foo bar" + "hello")
    p0 -= p0
    _ = p0
    _ = p1
    _ = v20
    return -1
}

func main() {
    v1 := []int{32768, 4294967295, 677}
    var v2 uint = 2
    var v3 float32 = 0.1
    var v4 int = 42
    var v5 bool = (false || (!true))
    var v21 int = (v4 ^ -1)
    var v22 uint32 = 796
    var v23 int16 = -32768
    if (2147483647 <= -2147483648) {
        for fi24 := 0; fi24 < 10; fi24++ {
            v25 := GenS2{F0: 0, F1: v2, F2: -1}
            v26 := GenS2{F0: 1, F1: (v2 - (v2 & v2)), F2: 1}
            v27 := GenS4{F0: -9223372036854775808, F1: v2, F2: 0}
            _ = fi24
            _ = v25
            _ = v26
            _ = v27
        }
        switch "" {
        case "\n\t":
            v2 = 520
        case "a":
            fmt.Println(v4, v2)
        case "":
            v28 := GenS5{F0: v2}
            _ = v28
        }
    }
    var v29 float32 = 1.5
    v29 = (((v29 - (-2.25 + (0.1 - (float32(-2.25) * v29)))) - (v29 + 0.1)) + float32(-2.25))
    for fi30 := 0; fi30 < 3; fi30++ {
        if v5 {
            v3 = float32(0.1)
        } else {
            v31 := GenS2{F0: 42, F1: (v2 | (^(v2))), F2: -2147483648}
            _ = len(v1)
            switch "" {
            case "\n\t":
                fmt.Println(v23, v4)
            case "a":
                var v32 uint64 = 815
                _ = v32
            default:
                if v5 {
                    v5 = true
                } else {
                    if v5 {
                        var v33 uint = (v2 / ((42) | 1))
                        _ = v33
                    }
                    fmt.Println(fi30, v2, v4)
                }
            }
            _ = v31
        }
        if true {
            var v34 int16 = -671
            var v35 int64 = 0
            _ = v34
            _ = v35
        } else {
            v1 = append(v1, v21)
        }
        var v36 uint32 = 4294967295
        _ = fi30
        _ = v36
    }
    if (0 >= 42) {
        var v37 float64 = 0.0
        fmt.Println(v22)
        var v38 uint = v2
        _ = v37
        _ = v38
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
