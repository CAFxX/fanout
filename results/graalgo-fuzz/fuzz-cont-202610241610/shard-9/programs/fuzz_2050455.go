package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 uint8 }
type GenS2 struct { F0 int64; F1 uint16 }
func genFn0(p0 int, p1 uint64) string {
    v7 := GenS1{F0: 127, F1: 0, F2: 42}
    p0 = (p0 + -492)
    p1 = p1
    switch "" {
    case "a":
        v8 := []int{-320, p0, 9223372036854775807, (p0 % ((-2147483648) | 1))}
        _ = v8
    case "\n\t":
        v9 := []int{}
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    return ("foo bar" + "a")
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]int)
    var v3 float64 = 3.14159
    var v4 int8 = -108
    v5 := make(map[int]string)
    v6 := []int{4294967295}
    var v10 uint8 = 127
    var v11 bool = false
    if false {
        fmt.Println(v4, v10, v10)
    }
    v4 -= v4
    var v12 int16 = -1
    fmt.Println(v3, v10, v12)
    v12 = (v12 | (v12 & v12))
    var v13 uint = 65535
    v13 -= v13
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    if v11 {
        v14 := GenS1{F0: v10, F1: 252, F2: (v10 ^ (v10 / (((v10 + (v10 ^ 110))) | 1)))}
        v11 = (0 < 249)
        _ = v14
    }
    v15 := v5[32767]
    v16 := GenS2{F0: 9223372036854775807, F1: 5}
    var v17 int64 = int64(v10)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
