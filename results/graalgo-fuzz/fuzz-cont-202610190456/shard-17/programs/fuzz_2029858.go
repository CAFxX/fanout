package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 uint32; F2 uint }
func genFn0(p0 bool, p1 uint8) int32 {
    var v6 float32 = ((1e10 - 0.0) - 3.14159)
    switch ("foo bar" + ("foo bar" + "foo bar")) {
    case "a":
        p0 = (1.5 < 0.0)
    }
    var v7 int = 1
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return -2147483648
}
func genFn1(p0 bool, p1 uint32) uint32 {
    v8 := GenS1{F0: 128}
    var v9 int64 = 42
    var v10 uint = 998
    var v11 uint8 = 42
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v8
    _ = v9
    return (p1 & (p1 ^ p1))
}

func main() {
    var v1 float32 = 1e10
    var v2 int16 = -32768
    v3 := make(map[uint32]string)
    var v4 int8 = 42
    v5 := make(map[string]string)
    var v12 int16 = v2
    fmt.Println(v2, v4, v4)
    var v13 uint = 65535
    fmt.Println(v1, v13, v1)
    v14 := GenS2{F0: 128, F1: 943, F2: (v13 << (((v13 * (v13 + (v13 + (v13 ^ 255))))) & 63))}
    var v15 float64 = -2.25
    v1 = ((float32(v15) - 0.0) * 0.1)
    for fi16 := 0; fi16 < 16; fi16++ {
        var v17 float32 = float32((3.14159 / float64((v1 / float32(float64((((v1 - (0.1 / v1)) + -2.25) - 1.5)))))))
        _ = fi16
        _ = v17
    }
    fmt.Println(v12)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
