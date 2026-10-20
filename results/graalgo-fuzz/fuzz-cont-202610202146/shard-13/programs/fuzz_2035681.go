package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint64; F1 bool }
type GenS4 struct { F0 int32; F1 int32; F2 uint16 }
type GenS5 struct { F0 uint32 }
func genFn0() int8 {
    var v5 uint = 2
    for fi6 := 0; fi6 < 8; fi6++ {
        var v7 uint8 = 42
        var v8 int32 = 1
        _ = fi6
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 20; fi9++ {
        var v10 string = "foo bar"
        fmt.Println(v10)
        _ = fi9
        _ = v10
    }
    v11 := []int{}
    if false {
        if false {
            v5 = (v5 / (((^(v5))) | 1))
            var v12 uint = (v5 % (((^(v5))) | 1))
            _ = v12
        }
        for fi13 := 0; fi13 < 14; fi13++ {
            v14 := GenS1{F0: 18446744073709551615, F1: 642, F2: true}
            var v15 string = "a"
            v15 = (v15 + v15)
            _ = fi13
            _ = v14
            _ = v15
        }
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
    }
    _ = v11
    _ = v5
    return -12
}
func genFn1(p0 uint32) uint64 {
    v16 := GenS2{F0: 0}
    p0 = 2
    v17 := GenS2{F0: -1}
    v18 := GenS1{F0: 732, F1: -460, F2: false}
    _ = p0
    _ = v16
    _ = v17
    _ = v18
    return 0
}

func main() {
    var v1 int32 = 839
    var v2 float32 = 0.1
    var v3 int = 576
    v4 := make(map[int]int)
    switch (^(v3)) {
    case -718:
        v1 += 1
    default:
        v3 -= v3
    }
    var v19 float32 = (-2.25 + ((float32((1e10 + (((0.0 * 0.0) / 0.1) - (((1e10 / 1.0) / 1.0) / float64((float32(1.5) + v2)))))) + v2) / ((v2 + 3.14159) - v2)))
    v20 := []int{}
    switch (v3 | (v3 | (^(v3)))) {
    case -2147483648:
        for fi21 := 0; fi21 < 14; fi21++ {
            if true {
                var v22 bool = true
                _ = v22
            } else {
                if ("" > ("" + "foo bar")) {
                    for fi23 := 0; fi23 < 11; fi23++ {
                        var v24 uint64 = 2
                        _ = fi23
                        _ = v24
                    }
                    var v25 uint64 = 2
                    _ = v25
                }
                switch "\n\t" {
                case "a":
                    if len(v20) >= 2 {
                        v20 = v20[1:len(v20)]
                    }
                case "\n\t":
                    var v26 int64 = 9223372036854775807
                    _ = v26
                default:
                    v3 -= (v3 % (((v3 & (v3 >> ((355) & 63)))) | 1))
                }
                if false {
                    v27 := GenS2{F0: v1}
                    _ = v27
                } else {
                    for fi28 := 0; fi28 < 16; fi28++ {
                        delete(v4, 1)
                        switch "\n\t" {
                        case "foo bar":
                            var v29 float64 = (0.0 / (3.14159 / -2.25))
                            _ = v29
                        case "hello":
                            v30 := GenS3{F0: 18446744073709551615, F1: ((3.14159 >= 1e10) || true)}
                            _ = v30
                        default:
                            v3 -= (^(v3))
                        }
                        v31 := GenS4{F0: 1, F1: 1, F2: 255}
                        _ = fi28
                        _ = v31
                    }
                    v4[fi21] = fi21
                }
            }
            v32 := v4[(fi21 * v3)]
            _ = fi21
            _ = v32
        }
    default:
        v2 = v19
    }
    for fi33 := 0; fi33 < 5; fi33++ {
        v2 -= (0.1 / v19)
        var v34 uint8 = 42
        _ = fi33
        _ = v34
    }
    v20 = append(v20, -2147483648)
    var v35 uint = 309
    var v36 int = (v3 / ((0) | 1))
    v37 := GenS5{F0: 2147483647}
    var v38 float32 = v19
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
}
