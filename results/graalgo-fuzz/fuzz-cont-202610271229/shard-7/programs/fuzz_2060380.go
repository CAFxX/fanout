package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 uint16; F1 uint64; F2 uint8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int8 }

func main() {
    v1 := []int{}
    var v2 int16 = 0
    var v3 string = "\n\t"
    var v4 int8 = -128
    var v5 uint8 = 128
    v6 := make(map[int]int)
    fmt.Println(v2, v4, v3)
    var v7 uint32 = 0
    v7 = (v7 | v7)
    for fi8 := 0; fi8 < 12; fi8++ {
        v9 := GenS1{F0: (fi8 / ((-825) | 1)), F1: "a"}
        _ = fi8
        _ = v9
    }
    v4 += (v4 & -128)
    var v10 bool = (0 == 481)
    var v11 uint16 = 418
    var v12 uint = 270
    if ((v12 | v12) != (v12 & 18446744073709551615)) {
        var v13 string = v3
        for fi14 := 0; fi14 < 13; fi14++ {
            v15 := GenS1{F0: fi14, F1: v13}
            if v10 {
                switch 4294967295 {
                case 440:
                    if false {
                        var v16 uint = v12
                        var v17 string = "a"
                        _ = v16
                        _ = v17
                    } else {
                        for fi18 := 0; fi18 < 6; fi18++ {
                            var v19 uint8 = (v5 / (((v5 * (v5 | (v5 | (^(v5)))))) | 1))
                            var v20 float64 = 1.5
                            _ = fi18
                            _ = v19
                            _ = v20
                        }
                        v4 = v4
                    }
                case 42:
                    v11 = v11
                default:
                    fmt.Println(v3, v2, v3)
                }
                var v21 int32 = -2147483648
                var v22 uint64 = 228
                _ = v21
                _ = v22
            } else {
                for fi23 := 0; fi23 < 8; fi23++ {
                    var v24 float32 = 1e10
                    fmt.Println(v2, v10)
                    var v25 int8 = v4
                    var v26 int64 = 1
                    v27 := GenS2{F0: (v11 - (v11 * (v11 >> ((uint(v7)) & 63)))), F1: 42, F2: (v5 + v5)}
                    v28 := v6[(fi14 * fi14)]
                    if true {
                        v10 = true
                        fmt.Println(v3, v3, fi14)
                        v29 := GenS3{F0: -1}
                        _ = v29
                    } else {
                        var v30 uint = 812
                        v31 := GenS1{F0: int(v4), F1: (("a" + "foo bar") + (((("hello" + v3) + (v13 + v13)) + "hello") + (v13 + "\n\t")))}
                        v11 = (v11 & (^(v11)))
                        v12 += (v12 * (v12 / ((v12) | 1)))
                        _ = v30
                        _ = v31
                    }
                    _ = fi23
                    _ = v24
                    _ = v25
                    _ = v26
                    _ = v27
                    _ = v28
                }
                fmt.Println(v10, v4, fi14)
            }
            _ = fi14
            _ = v15
        }
        if ("foo bar" >= "") {
            v32 := GenS4{F0: v4}
            v33 := GenS3{F0: 319}
            _ = v32
            _ = v33
        } else {
            if len(v1) > 0 {
                v1[0] = 0
            }
        }
        _ = v13
    }
    var v34 float64 = 0.1
    var v35 float32 = ((1.5 + 0.0) - (0.0 / (1e10 * 3.14159)))
    v36 := v6[-277]
    fmt.Println(v11, v34, v35)
    v37 := GenS1{F0: 127, F1: (((("hello" + ("a" + v3)) + v3) + "hello") + v3)}
    delete(v6, v36)
    fmt.Println(v34)
    for fi38 := 0; fi38 < 10; fi38++ {
        fmt.Println(v5)
        v39 := GenS3{F0: 0}
        _ = fi38
        _ = v39
    }
    for fi40 := 0; fi40 < 12; fi40++ {
        v2 = -1
        _ = len(v6)
        var v41 string = ("hello" + v3)
        _ = fi40
        _ = v41
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
