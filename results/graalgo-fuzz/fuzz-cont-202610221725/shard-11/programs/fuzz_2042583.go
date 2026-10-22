package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 bool }
type GenS2 struct { F0 int16; F1 uint16; F2 bool }
type GenS3 struct { F0 uint8; F1 uint32; F2 int8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int8; F1 int32; F2 uint16 }
func genFn0() bool {
    v5 := []int{-1, -939, 9223372036854775807, 9223372036854775807}
    for fi6 := 0; fi6 < 2; fi6++ {
        fmt.Println(fi6)
        var v7 float64 = 1e10
        v8 := GenS1{F0: 9223372036854775807, F1: 10, F2: (!(!false))}
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = v5
    return false
}
func genFn1(p0 bool) int64 {
    v9 := GenS1{F0: -9223372036854775808, F1: 1, F2: true}
    var v10 int64 = -1
    v11 := GenS2{F0: 279, F1: 105, F2: p0}
    fmt.Println(p0, v10)
    _ = p0
    _ = v10
    _ = v11
    _ = v9
    return (v10 | 756)
}

func main() {
    v1 := make(map[int]bool)
    var v2 int = -2147483648
    v3 := make(map[int]int)
    var v4 int = (v2 | v2)
    v12 := GenS1{F0: -1, F1: 273, F2: (false && false)}
    for fi13 := 0; fi13 < 18; fi13++ {
        switch v2 {
        case 9223372036854775807:
            if false {
                var v14 int8 = 100
                if (!false) {
                    for fi15 := 0; fi15 < 19; fi15++ {
                        v2 = 127
                        var v16 bool = true
                        switch "foo bar" {
                        case "\n\t":
                            v16 = v16
                        }
                        _ = fi15
                        _ = v16
                    }
                    v17 := GenS3{F0: 165, F1: 130, F2: (v14 % ((v14) | 1))}
                    _ = v17
                }
                fmt.Println(v4)
                _ = v14
            }
        case 992:
            v18 := GenS1{F0: -616, F1: 999, F2: false}
            _ = v18
        default:
            v2 += v4
        }
        v19 := GenS1{F0: 1, F1: 2, F2: false}
        _ = fi13
        _ = v19
    }
    if true {
        v20 := GenS2{F0: 7, F1: 450, F2: true}
        v21 := []int{v2, 0, v2}
        v22 := GenS4{F0: "a"}
        _ = v20
        _ = v21
        _ = v22
    }
    for fi23 := 0; fi23 < 10; fi23++ {
        if false {
            for fi24 := 0; fi24 < 13; fi24++ {
                var v25 string = "a"
                _ = fi24
                _ = v25
            }
        }
        v4 = v4
        if true {
            v26 := GenS5{F0: -38, F1: -1, F2: 0}
            v4 = v2
            if true {
                v27 := []int{1, fi23, (fi23 & v4)}
                if false {
                    var v28 uint16 = 789
                    _ = v28
                } else {
                    if false {
                        v27 = append(v27, v2)
                        for fi29 := 0; fi29 < 14; fi29++ {
                            v30 := GenS3{F0: 128, F1: 65535, F2: -1}
                            _ = fi29
                            _ = v30
                        }
                    } else {
                        var v31 float32 = (-2.25 + 1e10)
                        var v32 float64 = (float64(3.14159) + 0.0)
                        _ = v31
                        _ = v32
                    }
                    var v33 string = "hello"
                    _ = len(v27)
                    _ = v33
                }
                v4 = v2
                _ = v27
            } else {
                switch "hello" {
                case "hello":
                    fmt.Println(fi23, v4)
                case "foo bar":
                    v4 = v4
                }
                var v34 uint64 = 507
                v35 := GenS2{F0: 1, F1: 675, F2: true}
                _ = v34
                _ = v35
            }
            _ = v26
        } else {
            var v36 uint8 = 127
            _ = v36
        }
        _ = fi23
    }
    if false {
        v4 += (v2 + v4)
        _ = len(v3)
    }
    var v37 uint16 = 322
    v38 := GenS2{F0: int16(v4), F1: (v37 >> ((65535) & 63)), F2: false}
    v37 += (v37 & 42)
    var v39 int16 = 42
    v40 := []int{v4, (v4 & v4), v4, -541}
    switch "a" {
    case "":
        v41 := GenS4{F0: "\n\t"}
        _ = v41
    case "\n\t":
        switch (v4 + v2) {
        case 128:
            v37 = (v37 * (^(v37)))
        case 42:
            for fi42 := 0; fi42 < 8; fi42++ {
                if (!false) {
                    _ = len(v40)
                    v39 = (v39 * v39)
                    if len(v40) >= 2 {
                        v40 = v40[1:len(v40)]
                    }
                }
                if len(v40) > 0 {
                    v40[0] = (v2 ^ (v4 ^ (v2 << ((255) & 63))))
                }
                _ = fi42
            }
        }
    case "foo bar":
        v43 := v3[(-(v2))]
        _ = v43
    }
    var v44 uint = 655
    v44 -= 470
    _ = len(v40)
    v37 += (v37 + 623)
    if len(v40) > 0 {
        v40[0] = (v4 | 2147483647)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v44
}
