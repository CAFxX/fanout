package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 uint64 }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 uint; F1 uint32 }
type GenS4 struct { F0 uint; F1 int }
func genFn0(p0 int) uint8 {
    v7 := GenS1{F0: -227, F1: 4294967295, F2: 745}
    v8 := []int{-404}
    _ = p0
    _ = v7
    _ = v8
    return 2
}
func genFn1(p0 uint64, p1 int64, p2 int32) int64 {
    p0 += 2
    v9 := GenS2{F0: -1, F1: 979}
    v10 := GenS3{F0: 1, F1: 65535}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return p1
}

func main() {
    v1 := []int{42}
    var v2 int16 = -146
    var v3 int = 947
    v4 := make(map[int]uint32)
    var v5 int64 = 976
    var v6 int64 = 0
    if (0 > 0) {
        _ = len(v4)
    }
    _ = len(v4)
    v6 += genFn1(2, v6, 0)
    var v11 uint32 = 2
    fmt.Println(v11)
    v11 = v11
    v12 := GenS4{F0: 186, F1: (v3 * (v3 + v3))}
    if true {
        var v13 uint64 = 2
        var v14 float64 = 1e10
        _ = v13
        _ = v14
    }
    var v15 int = (v3 * -9223372036854775808)
    var v16 string = "hello"
    for fi17 := 0; fi17 < 15; fi17++ {
        fmt.Println(v15)
        switch v15 {
        case -637:
            var v18 float64 = 1.5
            _ = v18
        case -9223372036854775808:
            v2 -= (^(v2))
        case -1:
            for fi19 := 0; fi19 < 13; fi19++ {
                var v20 float32 = 1e10
                switch fi17 {
                case 399:
                    if false {
                        switch fi17 {
                        case 88:
                            _ = len(v1)
                        default:
                            fmt.Println(v11)
                        }
                        v11 = v11
                    } else {
                        fmt.Println(v6, fi17, v5)
                        if false {
                            if true {
                                var v21 int8 = -128
                                var v22 uint32 = 2147483648
                                fmt.Println(v20, v16, v11)
                                _ = v21
                                _ = v22
                            }
                            var v23 bool = false
                            _ = v23
                        } else {
                            var v24 uint32 = v11
                            var v25 int64 = v6
                            _ = v24
                            _ = v25
                        }
                    }
                default:
                    var v26 float32 = (v20 - 1e10)
                    _ = v26
                }
                switch (fi19 | fi19) {
                case -2147483648:
                    var v27 uint64 = 0
                    _ = v27
                case 0:
                    switch "a" {
                    case "foo bar":
                        v28 := GenS2{F0: (v6 / ((v5) | 1)), F1: 18446744073709551615}
                        _ = v28
                    case "a":
                        v1 = append(v1, (-(fi19)))
                    case "":
                        var v29 float32 = v20
                        _ = v29
                    default:
                        v3 = v15
                    }
                case -334:
                    switch v16 {
                    case "a":
                        var v30 uint8 = uint8(v11)
                        _ = v30
                    case "foo bar":
                        fmt.Println(fi19, v15)
                    default:
                        var v31 int64 = -1
                        _ = v31
                    }
                default:
                    var v32 uint64 = 2
                    _ = v32
                }
                _ = fi19
                _ = v20
            }
        default:
            var v33 int16 = v2
            _ = v33
        }
        v16 = ((v16 + v16) + "a")
        _ = fi17
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
