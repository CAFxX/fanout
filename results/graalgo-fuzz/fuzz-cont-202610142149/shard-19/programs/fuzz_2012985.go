package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 uint32; F1 int8; F2 int16 }
type GenS4 struct { F0 uint64; F1 int16 }
func genFn0() int16 {
    var v7 float32 = float32((1.5 * (1.5 - 0.0)))
    v8 := []int{9223372036854775807, -2147483648, 9223372036854775807}
    v9 := GenS1{F0: 0}
    _ = v7
    _ = v8
    _ = v9
    return 1
}
func genFn1(p0 int64, p1 int32) uint8 {
    p0 -= int64(p1)
    v10 := GenS2{F0: 603, F1: 1}
    _ = p0
    _ = p1
    _ = v10
    return 120
}

func main() {
    var v1 string = "a"
    v2 := make(map[uint32]bool)
    v3 := []int{}
    v4 := make(map[string]int)
    var v5 string = v1
    var v6 int64 = 9223372036854775807
    v11 := GenS2{F0: 923, F1: 42}
    if (true && (0 <= 0)) {
        for fi12 := 0; fi12 < 1; fi12++ {
            for fi13 := 0; fi13 < 15; fi13++ {
                v6 -= v6
                var v14 int16 = 1
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi13
                _ = v14
            }
            if (!false) {
                var v15 float64 = (0.1 / (3.14159 - -2.25))
                var v16 uint64 = 4294967295
                fmt.Println(fi12, fi12, v15)
                _ = v15
                _ = v16
            } else {
                v17 := GenS2{F0: 65535, F1: 77}
                fmt.Println(v6, v5, fi12)
                _ = v17
            }
            var v18 uint32 = 717
            _ = fi12
            _ = v18
        }
        fmt.Println(v6, v1, v5)
    } else {
        var v19 uint64 = 0
        v5 = ""
        _ = v19
    }
    v5 = ""
    fmt.Println(v6, v1, v1)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v20 uint8 = 80
    v21 := GenS3{F0: 206, F1: 42, F2: genFn0()}
    v22 := GenS4{F0: 42, F1: -1}
    fmt.Println(v6)
    v5 = (v1 + v5)
    fmt.Println(v20, v1, v1)
    for fi23 := 0; fi23 < 3; fi23++ {
        v5 = ("a" + (v1 + v1))
        _ = fi23
    }
    v24 := v2[42]
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
