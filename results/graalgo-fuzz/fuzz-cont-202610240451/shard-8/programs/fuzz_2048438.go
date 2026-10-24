package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 uint; F1 int16; F2 uint32 }
type GenS3 struct { F0 int16; F1 int16; F2 uint32 }
func genFn0(p0 int64, p1 int) uint64 {
    v4 := GenS1{F0: 128, F1: 128}
    p1 += 796
    v5 := GenS2{F0: 0, F1: -32768, F2: 2147483648}
    p0 = 9223372036854775807
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 4294967295
}

func main() {
    var v1 uint = 65535
    var v2 uint8 = 2
    var v3 int32 = -1
    v1 = v1
    v3 += (v3 ^ v3)
    var v6 float32 = (0.0 / (0.1 / -2.25))
    fmt.Println(v3)
    v2 = v2
    v7 := GenS3{F0: 478, F1: 1, F2: 2147483648}
    fmt.Println(v2)
    v6 = (((float32(3.14159) + 1e10) * v6) + (v6 + (v6 + ((0.0 / (v6 - ((float32(3.14159) / (0.1 / 1e10)) / (((float32(float64(v6)) - (((v6 + v6) * 1.5) * 0.1)) - 3.14159) / v6)))) / 0.1))))
    v8 := []int{-793, 128}
    var v9 string = ""
    v10 := GenS2{F0: (v1 % ((v1) | 1)), F1: 1000, F2: 370}
    var v11 string = ((v9 + ("foo bar" + ("a" + ""))) + v9)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
