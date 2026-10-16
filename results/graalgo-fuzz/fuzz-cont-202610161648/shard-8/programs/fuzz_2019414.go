package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64; F1 string }

func main() {
    var v1 uint64 = 245
    var v2 bool = true
    v3 := make(map[string]int)
    var v4 int64 = 0
    fmt.Println(v1, v4, v4)
    switch (("" + ("hello" + ("" + ("foo bar" + "\n\t")))) + "a") {
    case "":
        for fi5 := 0; fi5 < 3; fi5++ {
            var v6 uint16 = 858
            _ = fi5
            _ = v6
        }
    case "foo bar":
        v1 -= v1
    }
    v7 := GenS1{F0: (v1 / (((^(v1))) | 1)), F1: "a", F2: 0}
    if v2 {
        if true {
            v8 := GenS1{F0: (v1 * v1), F1: "\n\t", F2: v4}
            _ = v8
        } else {
            var v9 float64 = -2.25
            if v2 {
                v2 = (42 == 812)
                fmt.Println(v1, v4)
            } else {
                for fi10 := 0; fi10 < 7; fi10++ {
                    if false {
                        fmt.Println(v4)
                        fmt.Println(v9, v9)
                        if (float32(((-2.25 - (1e10 / v9)) * 1.5)) <= 0.0) {
                            v9 -= v9
                        }
                    }
                    if (true && (255 == 0)) {
                        v2 = true
                        var v11 uint32 = 42
                        _ = v11
                    }
                    for fi12 := 0; fi12 < 1; fi12++ {
                        fmt.Println(fi10, v9)
                        v2 = (v2 && (!(!(1000 != 461))))
                        _ = fi12
                    }
                    _ = fi10
                }
                v9 -= v9
                var v13 int = -752
                _ = v13
            }
            _ = v9
        }
    }
    if v2 {
        for fi14 := 0; fi14 < 12; fi14++ {
            fmt.Println(v2, v1, fi14)
            _ = fi14
        }
        if ((0 == -201) && v2) {
            var v15 int64 = (v4 ^ v4)
            v3[((("foo bar" + "hello") + "hello") + "hello")] = -2147483648
            var v16 int16 = -1
            _ = v15
            _ = v16
        }
    }
    fmt.Println(v4, v1, v2)
    var v17 int32 = -2147483648
    fmt.Println(v4, v2)
    for fi18 := 0; fi18 < 1; fi18++ {
        var v19 int = 9223372036854775807
        var v20 uint = 4294967295
        _ = fi18
        _ = v19
        _ = v20
    }
    v21 := []int{-1, -424, -2147483648, 4294967295}
    switch ("" + ("\n\t" + "a")) {
    case "foo bar":
        var v22 int = 128
        _ = v22
    default:
        for fi23 := 0; fi23 < 9; fi23++ {
            v4 += 1
            if v2 {
                switch fi23 {
                case 32768:
                    var v24 uint32 = 2147483647
                    _ = v24
                case 9223372036854775807:
                    v25 := GenS1{F0: v1, F1: "a", F2: (-(v4))}
                    _ = v25
                default:
                    v26 := GenS2{F0: (v17 % (((v17 & (v17 % ((v17) | 1)))) | 1))}
                    _ = v26
                }
                for fi27 := 0; fi27 < 11; fi27++ {
                    v28 := GenS2{F0: v17}
                    var v29 int32 = (v17 * v17)
                    _ = fi27
                    _ = v28
                    _ = v29
                }
            } else {
                v2 = v2
                v30 := GenS3{F0: (v4 / (((v4 << ((2) & 63))) | 1)), F1: "\n\t"}
                v2 = true
                _ = v30
            }
            v17 += v17
            _ = fi23
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v7
}
