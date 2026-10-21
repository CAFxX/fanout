package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 uint16 }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 int; F1 int }
func genFn0() uint {
    var v4 uint32 = 497
    v4 += 1
    _ = v4
    return 621
}

func main() {
    v1 := make(map[int]int)
    v2 := []int{-2147483648, 4294967295, 4294967295, 310}
    v3 := make(map[int]bool)
    if (!false) {
        switch 9223372036854775807 {
        case 32768:
            var v5 int64 = 1
            _ = v5
        case -2147483648:
            switch "a" {
            case "foo bar":
                var v6 int32 = -973
                _ = v6
            }
        case 128:
            for fi7 := 0; fi7 < 15; fi7++ {
                var v8 bool = true
                _ = fi7
                _ = v8
            }
        default:
            var v9 uint16 = 118
            _ = v9
        }
        v10 := GenS1{F0: "hello", F1: -9223372036854775808, F2: 711}
        var v11 uint32 = 2147483647
        _ = v10
        _ = v11
    } else {
        var v12 uint8 = 42
        _ = v12
    }
    switch ("\n\t" + "foo bar") {
    case "hello":
        var v13 float64 = 3.14159
        _ = v13
    case "foo bar":
        for fi14 := 0; fi14 < 17; fi14++ {
            if true {
                var v15 uint = 42
                v15 = v15
                var v16 int64 = 9223372036854775807
                _ = v15
                _ = v16
            }
            _ = fi14
        }
    case "a":
        switch -9223372036854775808 {
        case -9223372036854775808:
            var v17 float64 = ((0.0 + 0.1) - (float64(float32(0.0)) / 1e10))
            _ = v17
        default:
            if (0 != 2147483647) {
                var v18 int16 = -32768
                if false {
                    fmt.Println(v18)
                    for fi19 := 0; fi19 < 17; fi19++ {
                        if true {
                            var v20 int = (fi19 % ((fi19) | 1))
                            v21 := GenS2{F0: -869, F1: false}
                            var v22 float64 = (-2.25 + 1e10)
                            _ = v20
                            _ = v21
                            _ = v22
                        } else {
                            fmt.Println(v18, v18)
                            var v23 float64 = -2.25
                            if (9223372036854775807 >= -257) {
                                for fi24 := 0; fi24 < 4; fi24++ {
                                    v25 := v1[fi24]
                                    _ = fi24
                                    _ = v25
                                }
                            }
                            _ = v23
                        }
                        v18 = (v18 + 1)
                        var v26 int64 = 1
                        _ = fi19
                        _ = v26
                    }
                }
                _ = v18
            }
        }
    default:
        v27 := GenS2{F0: -1, F1: (((4294967295 > 4294967295) && true) && (!false))}
        _ = v27
    }
    var v28 int16 = -1
    fmt.Println(v28)
    if false {
        if false {
            var v29 bool = (1 < 2147483647)
            _ = v29
        }
        v28 = v28
        v28 -= 1000
    }
    for fi30 := 0; fi30 < 6; fi30++ {
        for fi31 := 0; fi31 < 9; fi31++ {
            v32 := GenS1{F0: "a", F1: 1, F2: 2}
            fmt.Println(fi30, fi31)
            if false {
                _ = len(v3)
                v33 := GenS2{F0: -629, F1: true}
                fmt.Println(fi30, fi31)
                _ = v33
            } else {
                v28 = v28
            }
            _ = fi31
            _ = v32
        }
        var v34 uint = 65535
        _ = fi30
        _ = v34
    }
    var v35 int32 = 2147483647
    v36 := GenS3{F0: -9223372036854775808, F1: -1}
    _ = len(v2)
    fmt.Println(v28)
    if (false || false) {
        v35 += (^(v35))
    }
    v35 += (v35 / ((v35) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v35
    _ = v36
}
