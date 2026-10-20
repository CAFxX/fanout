package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 uint }
type GenS2 struct { F0 uint16; F1 int16; F2 int16 }
func genFn0(p0 int64, p1 int8) int16 {
    fmt.Println(p1)
    v7 := GenS1{F0: 0, F1: ((float64((0.1 + (1e10 / (3.14159 - ((1e10 / 0.1) + -2.25))))) - 0.1) >= ((1e10 + (((3.14159 * -2.25) / 3.14159) / 1.5)) - (-2.25 / 1.5))), F2: 1}
    _ = p0
    _ = p1
    _ = v7
    return -32768
}
func genFn1(p0 int32, p1 uint8, p2 uint32) uint8 {
    p0 = (p0 << ((516) & 63))
    var v8 int16 = 644
    p0 -= (p0 - p0)
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return (p1 & p1)
}

func main() {
    v1 := []uint32{}
    var v2 int = -2147483648
    var v3 int8 = 32
    v4 := []uint32{1}
    var v5 int16 = 0
    v6 := []uint32{}
    var v9 int8 = v3
    fmt.Println(v2)
    var v10 int16 = -1
    switch ("a" + ("" + "foo bar")) {
    case "hello":
        v11 := GenS2{F0: 65535, F1: v5, F2: genFn0(int64(v9), v9)}
        _ = v11
    }
    v2 += v2
    v2 -= (v2 + (^(v2)))
    v12 := GenS2{F0: 0, F1: v10, F2: v5}
    var v13 uint32 = 0
    var v14 int = (v2 - (-(v2)))
    v13 += v13
    var v15 float64 = 3.14159
    fmt.Println(v5, v2)
    _ = len(v1)
    fmt.Println(v15, v3, v3)
    var v16 uint16 = 403
    v13 = (v13 + v13)
    var v17 uint32 = v13
    v14 += v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
