package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 int64 }
type GenS2 struct { F0 string; F1 int64; F2 int64 }

func main() {
    v1 := make(map[int]int)
    var v2 int16 = -32768
    var v3 uint64 = 4294967295
    v4 := []uint32{}
    var v5 uint16 = 42
    var v6 bool = false
    v7 := GenS1{F0: 9223372036854775807, F1: 255, F2: -185}
    var v8 int8 = 77
    var v9 uint64 = v3
    var v10 float64 = 1e10
    switch ("\n\t" + ((("a" + "hello") + ("hello" + "hello")) + "\n\t")) {
    case "hello":
        v11 := GenS2{F0: "a", F1: -1, F2: 9223372036854775807}
        _ = v11
    default:
        v12 := GenS2{F0: "hello", F1: 9223372036854775807, F2: 0}
        _ = v12
    }
    var v13 string = "a"
    fmt.Println(v8, v2)
    switch -9223372036854775808 {
    case 4294967295:
        v9 -= (^(v9))
    case 28:
        v2 += 1
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v14 uint32 = 4294967295
    v15 := GenS1{F0: 0, F1: 175, F2: 643}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
