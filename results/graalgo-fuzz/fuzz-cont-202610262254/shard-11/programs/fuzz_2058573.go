package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8; F1 int; F2 uint64 }

func main() {
    var v1 uint32 = 42
    var v2 uint8 = 17
    var v3 int32 = 1
    var v4 int16 = -18
    v5 := make(map[int]int)
    var v6 uint64 = 1
    v1 = 196
    v7 := GenS1{F0: 128}
    var v8 uint8 = (v2 & v2)
    v9 := GenS1{F0: -82}
    for fi10 := 0; fi10 < 16; fi10++ {
        _ = len(v5)
        _ = fi10
    }
    fmt.Println(v2)
    fmt.Println(v1, v4)
    v1 = v1
    v3 += (v3 + (v3 | (v3 & (-(v3)))))
    v11 := GenS2{F0: int64(v3)}
    if true {
        for fi12 := 0; fi12 < 8; fi12++ {
            for fi13 := 0; fi13 < 1; fi13++ {
                v14 := v5[4294967295]
                v6 += 606
                var v15 uint64 = v6
                _ = fi13
                _ = v14
                _ = v15
            }
            for fi16 := 0; fi16 < 1; fi16++ {
                var v17 uint64 = 4294967295
                fmt.Println(v2, v3)
                var v18 int = fi12
                _ = fi16
                _ = v17
                _ = v18
            }
            for fi19 := 0; fi19 < 13; fi19++ {
                var v20 string = "a"
                var v21 int64 = -892
                _ = fi19
                _ = v20
                _ = v21
            }
            _ = fi12
        }
        v22 := []int{126, -731}
        _ = v22
    }
    var v23 float64 = float64((-2.25 + ((0.0 * float32(float64(3.14159))) * 0.0)))
    for fi24 := 0; fi24 < 18; fi24++ {
        v6 = (v6 / (((v6 + (v6 + (v6 >> ((2) & 63))))) | 1))
        _ = fi24
    }
    v25 := []int{int(v6), 0, 1, -1}
    v26 := GenS3{F0: 127, F1: -2147483648, F2: (v6 & v6)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v23
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
