package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 int }
func genFn0() uint {
    var v5 int16 = 0
    v5 = -32768
    _ = v5
    return 1
}
func genFn1(p0 uint32, p1 bool) int32 {
    switch ("foo bar" + "a") {
    case "hello":
        v6 := []int{128, -1}
        _ = v6
    case "\n\t":
        v7 := GenS1{F0: 2147483647}
        _ = v7
    }
    v8 := []int{2147483647, 128, 951}
    v8 = append(v8, 32767)
    var v9 uint16 = 42
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 2147483647
}

func main() {
    v1 := []uint32{42, 0}
    var v2 int8 = -83
    var v3 int64 = 924
    v4 := make(map[int]bool)
    var v10 int = -565
    v11 := GenS1{F0: -2147483648}
    v2 += 127
    var v12 int32 = -2147483648
    var v13 float32 = 1.5
    fmt.Println(v3, v12)
    v12 -= 36
    v2 -= -1
    v14 := GenS1{F0: 201}
    var v15 string = "hello"
    v1 = append(v1, 2)
    if len(v1) > 0 {
        v1[0] = 42
    }
    fmt.Println(v13, v2, v12)
    v16 := GenS1{F0: v12}
    var v17 string = "foo bar"
    if true {
        var v18 float32 = v13
        _ = v18
    } else {
        var v19 float32 = v13
        v20 := GenS2{F0: (v2 << ((255) & 63)), F1: (-(v10))}
        _ = v19
        _ = v20
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
