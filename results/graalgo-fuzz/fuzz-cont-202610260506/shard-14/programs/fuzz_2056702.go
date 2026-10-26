package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 int32; F1 bool; F2 int }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int32 }

func main() {
    var v1 uint64 = 724
    var v2 string = "hello"
    v3 := make(map[string]bool)
    v4 := make(map[int]uint32)
    delete(v4, 32768)
    if ((v1 * 0) <= 246) {
        fmt.Println(v2, v2)
        v5 := GenS1{F0: 0, F1: -800}
        if true {
            v6 := []int{-766, 2147483647, 9223372036854775807, 579}
            v2 = ((v2 + (v2 + "")) + v2)
            v7 := GenS1{F0: -9223372036854775808, F1: 9223372036854775807}
            _ = v6
            _ = v7
        } else {
            for fi8 := 0; fi8 < 17; fi8++ {
                for fi9 := 0; fi9 < 5; fi9++ {
                    v1 -= 2
                    switch fi9 {
                    case 775:
                        if (1 == -9223372036854775808) {
                            v2 = "\n\t"
                            if (1 != (v1 & 4294967295)) {
                                if (false || (639 != 2)) {
                                    _ = len(v4)
                                }
                                switch (("foo bar" + v2) + (v2 + ("foo bar" + (v2 + v2)))) {
                                case "foo bar":
                                    v10 := GenS2{F0: 120, F1: v2}
                                    _ = v10
                                case "hello":
                                    v2 = v2
                                default:
                                    var v11 bool = (-1 == 1)
                                    _ = v11
                                }
                                v1 = (v1 * v1)
                            }
                            var v12 int16 = -387
                            _ = v12
                        } else {
                            var v13 string = v2
                            _ = v13
                        }
                    case 127:
                        fmt.Println(fi9)
                    case 128:
                        v14 := GenS3{F0: 42, F1: (((fi9 ^ fi9) < (fi8 | fi8)) && false), F2: 2147483647}
                        _ = v14
                    default:
                        v15 := GenS1{F0: -9223372036854775808, F1: 42}
                        _ = v15
                    }
                    _ = fi9
                }
                var v16 string = v2
                _ = fi8
                _ = v16
            }
            fmt.Println(v2, v1)
            var v17 string = (v2 + v2)
            _ = v17
        }
        _ = v5
    } else {
        var v18 int32 = -2147483648
        v2 = v2
        v2 = v2
        _ = v18
    }
    if (!false) {
        v2 = "foo bar"
        delete(v3, "a")
        if (false || (606 > 65535)) {
            v19 := []int{32768, 32768, 42, 32768}
            for fi20 := 0; fi20 < 1; fi20++ {
                var v21 uint = 433
                for fi22 := 0; fi22 < 20; fi22++ {
                    var v23 int = fi22
                    _ = fi22
                    _ = v23
                }
                var v24 float32 = float32(-2.25)
                _ = fi20
                _ = v21
                _ = v24
            }
            var v25 uint32 = 1
            _ = v19
            _ = v25
        }
    }
    v26 := []int{}
    v27 := GenS4{F0: 100}
    for fi28 := 0; fi28 < 15; fi28++ {
        if (2147483647 == -2147483648) {
            _ = len(v26)
            var v29 uint32 = 583
            if false {
                var v30 float64 = float64(0.0)
                _ = v30
            } else {
                v29 -= v29
            }
            _ = v29
        } else {
            if (-1 <= -193) {
                for fi31 := 0; fi31 < 8; fi31++ {
                    var v32 int8 = -76
                    v33 := GenS3{F0: 0, F1: true, F2: (-(fi28))}
                    _ = fi31
                    _ = v32
                    _ = v33
                }
                fmt.Println(fi28, v2)
            }
        }
        _ = fi28
    }
    var v34 uint8 = 1
    v35 := GenS5{F0: -2147483648}
    var v36 float64 = 1.5
    v37 := GenS4{F0: -1}
    v36 = v36
    v36 += (3.14159 - ((1e10 + v36) + v36))
    var v38 float32 = 1e10
    fmt.Println(v1, v34, v1)
    _ = len(v26)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
}
