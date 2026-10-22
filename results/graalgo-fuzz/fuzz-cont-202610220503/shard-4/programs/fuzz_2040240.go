package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 int32 }
type GenS2 struct { F0 int; F1 uint8; F2 uint16 }
type GenS3 struct { F0 int8; F1 int; F2 int8 }
type GenS4 struct { F0 int; F1 int8 }
type GenS5 struct { F0 int; F1 int16; F2 uint32 }
type GenS6 struct { F0 int }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]int)
    v3 := []int{2147483647, 9223372036854775807}
    var v4 uint32 = 2
    var v5 int16 = 1000
    v6 := make(map[uint32]string)
    v7 := GenS1{F0: (("a" + ("" + (("" + "a") + ""))) + "foo bar"), F1: 2, F2: -2147483648}
    fmt.Println(v5)
    var v8 float32 = float32(0.1)
    v9 := GenS1{F0: "hello", F1: 1, F2: 0}
    v8 = (v8 / (v8 - 1e10))
    switch 2147483647 {
    case 128:
        v10 := GenS2{F0: -9223372036854775808, F1: 53, F2: uint16(v4)}
        _ = v10
    default:
        v11 := GenS3{F0: 127, F1: 999, F2: 42}
        _ = v11
    }
    v12 := GenS4{F0: 674, F1: 100}
    var v13 bool = false
    if v13 {
        v14 := GenS5{F0: -55, F1: v5, F2: (v4 >> ((343) & 63))}
        _ = v14
    }
    v5 += v5
    for fi15 := 0; fi15 < 2; fi15++ {
        if v13 {
            v3 = append(v3, fi15)
        }
        delete(v1, "\n\t")
        _ = fi15
    }
    v16 := GenS5{F0: -299, F1: (v5 - v5), F2: v4}
    v17 := GenS6{F0: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
