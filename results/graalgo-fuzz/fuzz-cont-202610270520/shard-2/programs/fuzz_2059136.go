package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 int32; F1 int32; F2 int64 }
type GenS4 struct { F0 uint64; F1 uint8 }
type GenS5 struct { F0 uint16; F1 uint8; F2 int8 }

func main() {
    var v1 uint = 4294967295
    var v2 int8 = -124
    v3 := make(map[uint32]uint32)
    var v4 int = 4294967295
    var v5 int64 = 927
    v6 := make(map[int]uint32)
    v5 = 1
    v7 := GenS1{F0: true}
    switch (v4 & -662) {
    case -1:
        delete(v3, 604)
    }
    var v8 int = (v4 + v4)
    v5 += v5
    var v9 uint8 = 116
    fmt.Println(v2, v9)
    fmt.Println(v8)
    v2 = 40
    if (786 >= 2) {
        for fi10 := 0; fi10 < 6; fi10++ {
            var v11 int8 = (-(v2))
            _ = fi10
            _ = v11
        }
        if ("" < "\n\t") {
            for fi12 := 0; fi12 < 2; fi12++ {
                fmt.Println(v1, v4, v9)
                switch "foo bar" {
                case "hello":
                    v13 := []int{v4, v8, v4, (v4 | 144)}
                    _ = v13
                }
                switch "" {
                case "":
                    var v14 int32 = 2147483647
                    _ = v14
                case "foo bar":
                    v4 -= (fi12 >> ((945) & 63))
                case "\n\t":
                    for fi15 := 0; fi15 < 13; fi15++ {
                        if false {
                            var v16 int32 = 1
                            fmt.Println(v5, v9)
                            _ = v16
                        }
                        v17 := GenS2{F0: true, F1: 2}
                        v18 := []int{v8}
                        v19 := GenS2{F0: true, F1: 0}
                        var v20 int32 = 828
                        v8 = fi12
                        if false {
                            var v21 float64 = float64((0.0 + 0.0))
                            _ = v21
                        } else {
                            var v22 float32 = 0.1
                            _ = v22
                        }
                        v8 -= 4294967295
                        v3[2147483647] = 1
                        _ = fi15
                        _ = v17
                        _ = v18
                        _ = v19
                        _ = v20
                    }
                default:
                    for fi23 := 0; fi23 < 16; fi23++ {
                        v2 += (v2 * v2)
                        v24 := GenS3{F0: 2147483647, F1: 42, F2: v5}
                        _ = fi23
                        _ = v24
                    }
                }
                _ = fi12
            }
        } else {
            var v25 uint64 = 756
            switch "a" {
            case "":
                var v26 uint8 = 127
                _ = v26
            case "a":
                v6[988] = 4294967295
            case "\n\t":
                v27 := GenS1{F0: ((0 < -242) && true)}
                _ = v27
            }
            v28 := GenS4{F0: (v25 + (v25 + v25)), F1: (v9 * v9)}
            _ = v25
            _ = v28
        }
        switch "\n\t" {
        case "hello":
            v29 := GenS5{F0: 42, F1: v9, F2: (v2 + -1)}
            _ = v29
        case "a":
            v30 := v3[uint32(v1)]
            _ = v30
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
