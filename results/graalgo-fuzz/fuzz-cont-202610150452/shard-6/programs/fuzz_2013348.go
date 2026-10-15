package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 string }
type GenS5 struct { F0 bool }

func main() {
    var v1 int64 = 1
    var v2 int32 = 42
    var v3 string = ("foo bar" + "foo bar")
    v4 := make(map[int]int)
    v5 := make(map[int]string)
    var v6 float64 = 0.0
    v7 := GenS1{F0: 1, F1: 125}
    for fi8 := 0; fi8 < 13; fi8++ {
        if (((fi8 & (-(fi8))) == fi8) || true) {
            var v9 int64 = -1
            delete(v4, int(v2))
            v1 += v9
            _ = v9
        } else {
            switch v3 {
            case "foo bar":
                if (true || (!true)) {
                    v2 += (v2 & 2147483647)
                    switch 32767 {
                    case 2147483647:
                        v10 := GenS1{F0: 678, F1: 1}
                        _ = v10
                    case -9223372036854775808:
                        v6 = ((float64(3.14159) - 1.5) + 1.5)
                    case 859:
                        if ((v1 << ((2) & 63)) <= 461) {
                            var v11 uint32 = 1
                            v3 = "\n\t"
                            v12 := GenS1{F0: -32768, F1: 709}
                            _ = v11
                            _ = v12
                        }
                    }
                    v6 -= v6
                } else {
                    v2 = (v2 ^ 2147483647)
                }
            case "\n\t":
                v1 = int64(v2)
            default:
                fmt.Println(fi8)
            }
            v1 = (-(v1))
        }
        _ = fi8
    }
    switch 42 {
    case 32768:
        var v13 string = v3
        _ = v13
    case 128:
        if (-2147483648 > 930) {
            var v14 int32 = (v2 + (v2 % (((^(v2))) | 1)))
            fmt.Println(v6, v3)
            _ = v14
        } else {
            for fi15 := 0; fi15 < 13; fi15++ {
                if true {
                    v16 := GenS2{F0: -1}
                    _ = v16
                } else {
                    var v17 uint = 7
                    _ = v17
                }
                for fi18 := 0; fi18 < 18; fi18++ {
                    v19 := []int{}
                    v3 = (("hello" + v3) + v3)
                    _ = fi18
                    _ = v19
                }
                for fi20 := 0; fi20 < 12; fi20++ {
                    switch -373 {
                    case 127:
                        v21 := GenS2{F0: 42}
                        _ = v21
                    case 4294967295:
                        v6 -= (v6 / (float64(-2.25) - 3.14159))
                    case -2147483648:
                        v22 := []int{fi15, (fi15 + (fi20 & fi20)), 128}
                        _ = v22
                    default:
                        if true {
                            var v23 float32 = 0.0
                            if false {
                                v3 = ("a" + "foo bar")
                                v3 = v3
                            } else {
                                v24 := GenS3{F0: 1}
                                _ = v24
                            }
                            var v25 float32 = 1.5
                            _ = v23
                            _ = v25
                        }
                    }
                    v1 = (-(v1))
                    if (!(false || false)) {
                        var v26 int8 = -128
                        v27 := GenS4{F0: v3}
                        _ = v26
                        _ = v27
                    }
                    _ = fi20
                }
                _ = fi15
            }
        }
    }
    v5[632] = v3
    switch "a" {
    case "\n\t":
        v28 := []int{728}
        _ = v28
    case "hello":
        v29 := GenS5{F0: false}
        _ = v29
    default:
        v30 := []int{98, int(v1)}
        _ = v30
    }
    _ = len(v5)
    fmt.Println(v1, v2)
    v3 = v3
    var v31 uint32 = 0
    var v32 float64 = (0.0 + (v6 * float64(0.1)))
    fmt.Println(v31, v2)
    v33 := []int{-1, 2147483647, 136, -329}
    v34 := GenS2{F0: 1000}
    for fi35 := 0; fi35 < 20; fi35++ {
        var v36 bool = (false && (!false))
        var v37 int16 = 42
        _ = fi35
        _ = v36
        _ = v37
    }
    var v38 uint64 = 0
    v32 = v32
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
