package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 uint64 }
type GenS2 struct { F0 int64; F1 int8; F2 int64 }

func main() {
    v1 := make(map[string]int)
    var v2 float32 = 1.5
    var v3 uint16 = 65535
    v4 := make(map[string]int)
    var v5 int64 = 1
    var v6 int64 = v5
    var v7 uint64 = 739
    var v8 int8 = 42
    var v9 int = 42
    fmt.Println(v2)
    fmt.Println(v5)
    if ((((v2 + 0.1) + v2) - v2) >= float32((float64(v2) - -2.25))) {
        for fi10 := 0; fi10 < 19; fi10++ {
            if ((v3 ^ (v3 + 65535)) > 659) {
                if true {
                    v7 += v7
                    v6 += v6
                    v2 -= 3.14159
                }
            } else {
                var v11 int32 = -2147483648
                if true {
                    var v12 uint8 = 42
                    for fi13 := 0; fi13 < 13; fi13++ {
                        v14 := GenS1{F0: v8, F1: 26, F2: v7}
                        v15 := GenS2{F0: v5, F1: -1, F2: -236}
                        var v16 int16 = 32767
                        _ = fi13
                        _ = v14
                        _ = v15
                        _ = v16
                    }
                    v2 -= v2
                    _ = v12
                } else {
                    var v17 uint = uint(v7)
                    if false {
                        v5 = (v6 ^ v6)
                        _ = len(v4)
                    } else {
                        for fi18 := 0; fi18 < 4; fi18++ {
                            v5 -= (v5 << ((v17) & 63))
                            fmt.Println(v8)
                            var v19 int64 = v5
                            _ = fi18
                            _ = v19
                        }
                        v20 := []int{fi10, fi10}
                        _ = v20
                    }
                    _ = v17
                }
                var v21 float64 = 1.5
                _ = v11
                _ = v21
            }
            for fi22 := 0; fi22 < 17; fi22++ {
                v23 := GenS2{F0: v5, F1: v8, F2: 0}
                var v24 uint64 = (v7 / ((v7) | 1))
                v25 := GenS2{F0: (v6 | (v5 & (v5 ^ v5))), F1: (v8 & v8), F2: v5}
                _ = fi22
                _ = v23
                _ = v24
                _ = v25
            }
            v26 := GenS1{F0: 80, F1: (v8 ^ (v8 | 42)), F2: 11}
            _ = fi10
            _ = v26
        }
    }
    _ = len(v1)
    v8 = (v8 - 100)
    if true {
        v27 := v4["\n\t"]
        fmt.Println(v9, v8, v2)
        v28 := GenS2{F0: -9223372036854775808, F1: 42, F2: v6}
        _ = v27
        _ = v28
    }
    v29 := []int{v9, v9}
    v30 := GenS1{F0: (v8 >> ((2) & 63)), F1: v8, F2: (^(v7))}
    var v31 float32 = v2
    var v32 uint = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
