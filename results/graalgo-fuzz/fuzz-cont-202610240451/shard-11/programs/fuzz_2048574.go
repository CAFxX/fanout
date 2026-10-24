package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 int32 }
func genFn0(p0 int16) string {
    switch ("" + "\n\t") {
    case "":
        v5 := GenS1{F0: 4294967295, F1: 1000, F2: 0}
        _ = v5
    case "a":
        p0 += (-(p0))
    case "hello":
        v6 := []int{4294967295, -277, 583}
        _ = v6
    default:
        p0 -= (p0 % (((p0 >> ((543) & 63))) | 1))
    }
    var v7 float32 = 3.14159
    v8 := GenS1{F0: 494, F1: (p0 - p0), F2: 42}
    v9 := []int{32767, 1, -9223372036854775808}
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return (("\n\t" + ("foo bar" + ("" + ("\n\t" + ("\n\t" + ("" + "a")))))) + "a")
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint8 = 2
    v3 := make(map[int]int)
    var v4 int16 = -32768
    var v10 int64 = -163
    fmt.Println(v10)
    v10 = (v10 & -1)
    fmt.Println(v4)
    var v11 uint64 = 376
    if ((^(v11)) == 0) {
        var v12 uint16 = uint16(v11)
        v11 += (v11 ^ v11)
        v13 := []int{-9223372036854775808, -735, 32767, 917}
        _ = v12
        _ = v13
    } else {
        fmt.Println(v11, v10, v10)
    }
    v11 = (v11 | (v11 / ((v11) | 1)))
    var v14 int32 = 0
    var v15 int32 = (v14 * (^(v14)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
