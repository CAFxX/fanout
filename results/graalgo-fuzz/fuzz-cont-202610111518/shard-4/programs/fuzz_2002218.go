package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 uint16; F1 int8; F2 uint }
type GenS3 struct { F0 bool; F1 bool }
type GenS4 struct { F0 bool }

func main() {
    var v1 int = 32767
    v2 := []uint32{4294967295}
    var v3 uint = 65535
    v4 := []string{}
    v5 := make(map[string]int)
    v6 := make(map[uint32]bool)
    fmt.Println(v1, v3)
    var v7 int32 = 1
    _ = len(v4)
    v8 := GenS1{F0: 32767, F1: 42}
    v7 += (^(v7))
    v9 := GenS2{F0: 2, F1: 56, F2: 255}
    var v10 int16 = 42
    v7 -= (v7 - -1)
    v3 -= v3
    var v11 uint8 = 1
    v2 = append(v2, 4294967295)
    v1 -= v1
    v12 := GenS3{F0: (1e10 != (float32(((1.5 + float64((float32(float64(3.14159)) + 0.1))) * ((float64(float32(1.5)) / 1.0) * (0.1 + (1.5 - 0.0))))) / ((0.1 / ((float32(1e10) - -2.25) + 3.14159)) + 1e10))), F1: true}
    v13 := GenS4{F0: ("a" < ("foo bar" + "\n\t"))}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
