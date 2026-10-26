package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 uint; F1 int8; F2 uint16 }
type GenS3 struct { F0 uint32; F1 string }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int64; F1 uint64; F2 int }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]uint32)
    var v3 int32 = -990
    var v4 uint16 = 743
    v4 += v4
    fmt.Println(v3, v4)
    var v5 float32 = 0.0
    v6 := GenS1{F0: 1, F1: -313}
    var v7 float64 = 1e10
    if true {
        v8 := GenS1{F0: 42, F1: 524}
        _ = v8
    } else {
        switch "" {
        case "\n\t":
            v9 := []int{32768}
            _ = v9
        default:
            for fi10 := 0; fi10 < 19; fi10++ {
                if (true || (0.0 < 0.0)) {
                    if true {
                        v7 = float64(v5)
                        v11 := GenS2{F0: 65535, F1: -1, F2: (^(v4))}
                        if true {
                            v12 := []int{(fi10 & fi10)}
                            _ = v12
                        } else {
                            for fi13 := 0; fi13 < 10; fi13++ {
                                v2[958] = 2
                                if (-1 != (v3 >> ((42) & 63))) {
                                    v14 := []int{(fi10 & fi10), (fi10 % ((fi10) | 1)), fi13, fi13}
                                    _ = v14
                                }
                                _ = fi13
                            }
                        }
                        _ = v11
                    } else {
                        v5 = v5
                    }
                } else {
                    if true {
                        fmt.Println(v5, v5)
                        if (true || false) {
                            fmt.Println(v3)
                            v7 += v7
                        } else {
                            v3 = (v3 * (-(v3)))
                            v15 := []int{fi10, fi10}
                            switch (("foo bar" + "foo bar") + ("" + "a")) {
                            case "":
                                v16 := GenS3{F0: 2147483648, F1: "hello"}
                                _ = v16
                            case "hello":
                                fmt.Println(v3, v4)
                            default:
                                if true {
                                    var v17 float64 = v7
                                    v18 := GenS2{F0: 856, F1: 42, F2: (v4 / ((v4) | 1))}
                                    _ = v17
                                    _ = v18
                                }
                            }
                            _ = v15
                        }
                    }
                    fmt.Println(v7)
                    v19 := GenS4{F0: 32767}
                    _ = v19
                }
                v20 := []int{-164, fi10, 32768}
                _ = fi10
                _ = v20
            }
        }
        v4 = (v4 ^ (^(v4)))
    }
    v21 := GenS5{F0: 935, F1: 1, F2: 127}
    var v22 float32 = v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
