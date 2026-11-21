package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8; F1 bool; F2 int32 }
type GenS3 struct { F0 int16; F1 int32; F2 uint8 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]int)
    v3 := make(map[int]bool)
    var v4 float64 = -2.25
    var v5 uint = 42
    v4 = 1.5
    fmt.Println(v4, v5)
    v6 := v2[2]
    if true {
        var v7 int8 = 42
        for fi8 := 0; fi8 < 19; fi8++ {
            var v9 string = (("a" + "\n\t") + "hello")
            v7 = (v7 % ((84) | 1))
            if (!true) {
                v10 := GenS1{F0: 157}
                _ = v10
            } else {
                var v11 float32 = (3.14159 - (1e10 / 1.5))
                var v12 uint8 = 128
                _ = v11
                _ = v12
            }
            _ = fi8
            _ = v9
        }
        _ = v7
    }
    v6 = -2147483648
    v13 := GenS1{F0: int32(v6)}
    fmt.Println(v4)
    v5 = (^(v5))
    for fi14 := 0; fi14 < 15; fi14++ {
        for fi15 := 0; fi15 < 20; fi15++ {
            if false {
                var v16 string = ""
                if false {
                    var v17 float32 = ((3.14159 * 3.14159) / (-2.25 + (float32(1.5) + -2.25)))
                    var v18 uint = v5
                    _ = v17
                    _ = v18
                } else {
                    if false {
                        if true {
                            v1[(fi14 / ((9223372036854775807) | 1))] = v6
                            switch "foo bar" {
                            case "a":
                                v19 := GenS1{F0: 100}
                                _ = v19
                            case "\n\t":
                                v20 := GenS1{F0: 1}
                                _ = v20
                            }
                            var v21 uint32 = 4294967295
                            v22 := v2[v21]
                            v23 := GenS1{F0: 1}
                            if true {
                                switch "a" {
                                case "a":
                                    v24 := GenS2{F0: 0, F1: true, F2: -2147483648}
                                    _ = v24
                                case "foo bar":
                                    fmt.Println(v6, v22, v6)
                                default:
                                    v4 = v4
                                }
                            } else {
                                v5 -= v5
                                v25 := GenS3{F0: 32767, F1: -27, F2: 128}
                                _ = v25
                            }
                            _ = v21
                            _ = v22
                            _ = v23
                        }
                        var v26 int16 = -32768
                        _ = v26
                    } else {
                        var v27 int32 = 42
                        _ = v27
                    }
                    v28 := []int{fi14, fi15, fi14, 32768}
                    _ = v28
                }
                v5 -= (v5 ^ v5)
                _ = v16
            } else {
                fmt.Println(v5)
                var v29 int = (fi15 ^ fi15)
                var v30 uint16 = 42
                _ = v29
                _ = v30
            }
            _ = fi15
        }
        fmt.Println(v6, fi14, fi14)
        v31 := GenS1{F0: -2147483648}
        _ = fi14
        _ = v31
    }
    v32 := []int{(v6 ^ v6)}
    v33 := v2[60]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
