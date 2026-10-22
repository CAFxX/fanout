package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 int16 }
type GenS2 struct { F0 int64; F1 uint; F2 int16 }
func genFn0(p0 uint8) int32 {
    var v7 string = ("a" + ("\n\t" + ""))
    v7 = v7
    v8 := GenS1{F0: true, F1: 0, F2: int16(p0)}
    _ = p0
    _ = v7
    _ = v8
    return 598
}
func genFn1(p0 uint64, p1 uint, p2 int) uint64 {
    p0 = (p0 << ((4294967295) & 63))
    v9 := GenS2{F0: -9223372036854775808, F1: 8, F2: 633}
    var v10 int32 = int32(p0)
    fmt.Println(p2)
    var v11 string = "\n\t"
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v9
    return 0
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint8 = 255
    var v3 uint8 = v2
    v4 := make(map[uint32]uint32)
    var v5 int8 = 86
    v6 := make(map[string]string)
    fmt.Println(v3)
    var v12 bool = false
    v2 -= (v3 + v2)
    var v13 int16 = 42
    switch "" {
    case "":
        var v14 int16 = v13
        _ = v14
    case "foo bar":
        v5 = 42
    case "hello":
        v3 -= (v3 - v3)
    default:
        var v15 int16 = (v13 | -1)
        _ = v15
    }
    var v16 int16 = 1
    var v17 uint32 = 42
    v12 = v12
    var v18 bool = (1.5 != (((((0.1 / 1.0) - float64(0.0)) / 0.1) / 0.1) - (((3.14159 * -2.25) - 3.14159) / float64(1e10))))
    v5 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
