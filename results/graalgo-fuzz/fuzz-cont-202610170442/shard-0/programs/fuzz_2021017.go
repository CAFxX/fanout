package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 uint8; F1 int16; F2 uint32 }
type GenS4 struct { F0 bool; F1 bool; F2 int32 }
type GenS5 struct { F0 uint; F1 int; F2 string }
func genFn0(p0 int32, p1 string, p2 uint) uint16 {
    v7 := GenS1{F0: -1, F1: p0}
    switch "foo bar" {
    case "foo bar":
        p1 = "foo bar"
    case "":
        if false {
            if true {
                p1 = (p1 + p1)
            }
            v8 := GenS2{F0: true, F1: p1}
            v9 := []int{9223372036854775807, 838}
            _ = v8
            _ = v9
        }
    case "a":
        fmt.Println(p0, p1, p2)
    }
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 13
}
func genFn1(p0 int32, p1 bool) int8 {
    if p1 {
        p0 += 1
        switch "a" {
        case "":
            switch "hello" {
            case "\n\t":
                var v10 uint8 = 128
                _ = v10
            case "hello":
                fmt.Println(p1, p1)
            default:
                switch ("foo bar" + ("" + ("foo bar" + "hello"))) {
                case "a":
                    fmt.Println(p1, p1)
                case "foo bar":
                    p1 = p1
                case "":
                    v11 := []int{-642, -727}
                    _ = v11
                default:
                    switch "\n\t" {
                    case "":
                        for fi12 := 0; fi12 < 20; fi12++ {
                            var v13 int8 = -128
                            if true {
                                v14 := []int{(fi12 << ((947) & 63)), (fi12 | (^(fi12))), 76, fi12}
                                if len(v14) >= 2 {
                                    v14 = v14[1:len(v14)]
                                }
                                _ = v14
                            }
                            _ = fi12
                            _ = v13
                        }
                    case "\n\t":
                        for fi15 := 0; fi15 < 9; fi15++ {
                            switch "" {
                            case "\n\t":
                                var v16 float64 = ((0.1 * 3.14159) / 1e10)
                                _ = v16
                            case "foo bar":
                                var v17 float64 = (1e10 * 1e10)
                                _ = v17
                            case "a":
                                var v18 int = (^(fi15))
                                _ = v18
                            default:
                                v19 := []int{fi15, (-(fi15)), -502}
                                _ = v19
                            }
                            v20 := []int{fi15}
                            v21 := GenS3{F0: 42, F1: 1, F2: 2147483647}
                            _ = fi15
                            _ = v20
                            _ = v21
                        }
                    case "hello":
                        var v22 int64 = 0
                        _ = v22
                    default:
                        var v23 int64 = 42
                        _ = v23
                    }
                }
            }
        case "hello":
            p0 -= p0
        default:
            p0 -= (p0 & (p0 & p0))
        }
        var v24 uint = 852
        _ = v24
    }
    var v25 int = -569
    v26 := GenS1{F0: 127, F1: 1}
    var v27 float32 = (3.14159 + ((0.0 * -2.25) - 1.5))
    _ = p0
    _ = p1
    _ = v25
    _ = v26
    _ = v27
    return 127
}

func main() {
    var v1 uint16 = 177
    var v2 string = "a"
    v3 := []uint32{65535}
    var v4 float64 = (3.14159 - 0.0)
    v5 := []uint32{2147483648, 386, 1, 2147483648}
    v6 := []uint32{}
    v1 = v1
    var v28 int64 = -1
    _ = len(v3)
    var v29 string = "hello"
    var v30 float64 = 0.1
    fmt.Println(v4)
    var v31 int16 = 1
    v32 := GenS4{F0: (0 >= 255), F1: (0 <= 127), F2: -383}
    for fi33 := 0; fi33 < 7; fi33++ {
        if len(v3) > 0 {
            v3[0] = 2147483647
        }
        switch (fi33 / ((fi33) | 1)) {
        case -298:
            v30 -= (((1e10 / (v30 + float64(1.5))) * 1e10) + (1.5 / v30))
        default:
            if true {
                var v34 int16 = v31
                _ = v34
            }
        }
        _ = fi33
    }
    v31 += v31
    v4 += float64(0.1)
    v35 := GenS5{F0: 65535, F1: 2147483647, F2: ("a" + v2)}
    v28 += (v28 % (((v28 ^ v28)) | 1))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v29 = v2
    var v36 uint16 = v1
    v1 -= v36
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
