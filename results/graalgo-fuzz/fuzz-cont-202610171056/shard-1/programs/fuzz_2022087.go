package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 uint16; F1 uint16; F2 uint64 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]uint32)
    var v4 uint64 = 2
    v5 := make(map[string]bool)
    var v6 uint32 = 2147483647
    var v7 int8 = -1
    delete(v1, v6)
    if true {
        v6 = (v6 / (((v6 >> ((1) & 63))) | 1))
        if (!((1e10 / 1.0) <= 1.5)) {
            for fi8 := 0; fi8 < 16; fi8++ {
                for fi9 := 0; fi9 < 20; fi9++ {
                    v10 := []int{32767, fi8}
                    v11 := GenS1{F0: (2 >= uint8(v4)), F1: 0}
                    var v12 float32 = (1e10 / (-2.25 + 0.0))
                    delete(v3, v6)
                    _ = fi9
                    _ = v10
                    _ = v11
                    _ = v12
                }
                for fi13 := 0; fi13 < 7; fi13++ {
                    var v14 float64 = -2.25
                    v14 = (((v14 * v14) - (v14 * v14)) - (v14 + (0.1 + (((v14 - v14) + v14) / float64(1.5)))))
                    var v15 float64 = 1.5
                    _ = len(v5)
                    var v16 uint16 = 255
                    _ = fi13
                    _ = v14
                    _ = v15
                    _ = v16
                }
                _ = fi8
            }
        }
    }
    for fi17 := 0; fi17 < 10; fi17++ {
        var v18 int64 = 42
        if false {
            var v19 uint32 = v6
            _ = v19
        } else {
            switch (fi17 / ((fi17) | 1)) {
            case -746:
                var v20 uint = 18446744073709551615
                _ = v20
            case 42:
                v21 := []int{}
                _ = v21
            default:
                var v22 uint64 = 4294967295
                _ = v22
            }
        }
        v23 := GenS2{F0: 65535, F1: 2, F2: (v4 + (^(v4)))}
        _ = fi17
        _ = v18
        _ = v23
    }
    var v24 int32 = 2147483647
    v25 := GenS2{F0: 2, F1: 42, F2: (v4 * (v4 % ((0) | 1)))}
    var v26 int64 = 9223372036854775807
    var v27 string = "hello"
    v4 = v4
    v26 += v26
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
