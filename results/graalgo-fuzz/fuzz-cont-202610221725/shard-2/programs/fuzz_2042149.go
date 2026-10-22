package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 uint32; F1 string; F2 int8 }
type GenS3 struct { F0 int; F1 int8 }
func genFn0(p0 int16) int8 {
    var v5 int64 = 1
    var v6 bool = (88 <= 45)
    v7 := []int{int(p0)}
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return -89
}

func main() {
    v1 := make(map[uint32]string)
    var v2 bool = false
    var v3 float32 = float32(-2.25)
    v4 := make(map[int]int)
    var v8 int64 = 0
    v9 := GenS1{F0: 2147483647, F1: 1}
    fmt.Println(v8, v2, v8)
    for fi10 := 0; fi10 < 13; fi10++ {
        _ = len(v4)
        _ = fi10
    }
    var v11 int16 = -32768
    v3 = ((v3 / v3) / v3)
    v12 := GenS2{F0: 2147483648, F1: "hello", F2: 1}
    v3 -= v3
    var v13 float32 = 0.0
    var v14 uint32 = 553
    v14 = (v14 ^ v14)
    v15 := GenS3{F0: 312, F1: 42}
    v8 = v8
    var v16 float64 = (((1e10 / (float64((v13 + float32((float64(v13) + 1.5)))) + 1e10)) * 1e10) / 3.14159)
    fmt.Println(v14, v13, v11)
    var v17 int16 = v11
    var v18 int32 = 2147483647
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
