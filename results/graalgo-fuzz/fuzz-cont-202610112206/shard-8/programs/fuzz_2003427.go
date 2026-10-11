package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 int8 }
type GenS2 struct { F0 uint16; F1 string; F2 uint8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 string; F1 string; F2 int }
func genFn0(p0 int64, p1 uint) uint32 {
    v4 := GenS1{F0: -749, F1: 65535, F2: 100}
    v5 := []int{}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 0
}
func genFn1() uint64 {
    var v6 uint32 = genFn0(1, 213)
    v7 := GenS2{F0: 0, F1: "hello", F2: 255}
    _ = v6
    _ = v7
    return 439
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []uint32{42}
    var v3 int32 = 0
    v8 := v1[-2147483648]
    var v9 int32 = v3
    var v10 uint64 = 42
    v11 := GenS1{F0: 4294967295, F1: 750, F2: -1}
    var v12 string = "a"
    var v13 uint8 = 227
    switch v12 {
    case "\n\t":
        for fi14 := 0; fi14 < 15; fi14++ {
            var v15 int = int(v13)
            _ = fi14
            _ = v15
        }
    case "hello":
        v16 := GenS3{F0: 2147483647}
        _ = v16
    case "foo bar":
        _ = len(v1)
    default:
        var v17 float64 = ((3.14159 - 1.5) * -2.25)
        _ = v17
    }
    var v18 float32 = float32(-2.25)
    v19 := GenS4{F0: v12, F1: v12, F2: 167}
    var v20 uint8 = uint8(v8)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v8
    _ = v9
}
