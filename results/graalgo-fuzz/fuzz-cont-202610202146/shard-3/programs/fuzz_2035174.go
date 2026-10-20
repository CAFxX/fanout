package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 int64; F1 uint8; F2 int64 }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]string)
    var v3 float32 = ((float32(-2.25) - float32((1.5 / 3.14159))) * -2.25)
    var v4 uint16 = 138
    for fi5 := 0; fi5 < 6; fi5++ {
        fmt.Println(v3, fi5)
        _ = fi5
    }
    v6 := GenS1{F0: 1, F1: 2}
    var v7 uint8 = 42
    switch -844 {
    case 506:
        for fi8 := 0; fi8 < 6; fi8++ {
            if true {
                if true {
                    switch fi8 {
                    case -559:
                        for fi9 := 0; fi9 < 8; fi9++ {
                            switch "hello" {
                            case "\n\t":
                                switch "foo bar" {
                                case "":
                                    var v10 uint8 = uint8(v4)
                                    _ = v10
                                case "foo bar":
                                    v11 := GenS1{F0: 0, F1: 49}
                                    _ = v11
                                case "a":
                                    v7 += (v7 ^ v7)
                                }
                            }
                            var v12 int64 = 42
                            fmt.Println(v12, fi8, v4)
                            _ = fi9
                            _ = v12
                        }
                    case 983:
                        v13 := []int{fi8, fi8, (fi8 / ((fi8) | 1)), 841}
                        _ = v13
                    case 754:
                        var v14 bool = true
                        _ = v14
                    default:
                        switch "hello" {
                        case "a":
                            for fi15 := 0; fi15 < 15; fi15++ {
                                if false {
                                    if ((v3 + v3) > (-2.25 / v3)) {
                                        fmt.Println(v4, v3)
                                        fmt.Println(v3, fi8, v4)
                                        fmt.Println(v4, fi8, v4)
                                    }
                                    v16 := GenS1{F0: 582, F1: 42}
                                    _ = v16
                                }
                                switch fi15 {
                                case 2147483647:
                                    v1[("" + ("hello" + ("" + "\n\t")))] = true
                                default:
                                    switch "\n\t" {
                                    case "foo bar":
                                        v2[""] = "a"
                                    case "":
                                        v17 := []int{(^(fi15))}
                                        v17 = append(v17, (-(fi8)))
                                        switch 127 {
                                        case 32767:
                                            v18 := GenS1{F0: 748, F1: uint(v4)}
                                            _ = v18
                                        default:
                                            v7 = v7
                                        }
                                        _ = v17
                                    }
                                }
                                v19 := GenS1{F0: 1, F1: 1}
                                var v20 uint64 = 663
                                var v21 int64 = 627
                                _ = fi15
                                _ = v19
                                _ = v20
                                _ = v21
                            }
                        default:
                            v3 += -2.25
                        }
                    }
                    v22 := GenS2{F0: 316, F1: (v7 & (v7 | 1)), F2: -9223372036854775808}
                    _ = v22
                }
                var v23 uint64 = 2
                _ = v23
            }
            fmt.Println(v3, fi8, v3)
            _ = fi8
        }
    }
    if false {
        v4 = (v4 - v4)
        v4 -= v4
    }
    if ((v3 / 1e10) == float32(float64((3.14159 + v3)))) {
        v3 += (float32(0.1) / (3.14159 + v3))
    }
    v4 = v4
    fmt.Println(v3, v7)
    v4 -= 255
    v7 = v7
    v24 := GenS2{F0: 0, F1: (v7 << ((2) & 63)), F2: -195}
    v25 := []int{0, 2147483647, 25}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
