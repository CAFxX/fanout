package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 uint16; F1 uint32 }
func genFn0() uint16 {
    var v7 int16 = 32767
    switch "hello" {
    case "\n\t":
        v7 += 42
    default:
        fmt.Println(v7)
    }
    _ = v7
    return 255
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]string)
    var v3 uint64 = 489
    var v4 uint32 = 2
    var v5 int16 = 1
    var v6 uint32 = v4
    var v8 uint8 = 255
    for fi9 := 0; fi9 < 16; fi9++ {
        v6 = 2147483647
        _ = fi9
    }
    v10 := GenS1{F0: uint64(v6), F1: 4294967295}
    var v11 float64 = (3.14159 / 1e10)
    var v12 int = -9223372036854775808
    v13 := []int{(v12 + 4294967295), v12, v12, v12}
    v14 := GenS1{F0: (^(v3)), F1: (v12 >> ((719) & 63))}
    delete(v1, "hello")
    var v15 float32 = 0.1
    switch ("\n\t" + "") {
    case "hello":
        v13 = append(v13, 292)
    case "foo bar":
        if len(v13) >= 2 {
            v13 = v13[1:len(v13)]
        }
    case "a":
        fmt.Println(v15)
    default:
        v13 = append(v13, (v12 % ((v12) | 1)))
    }
    v16 := GenS2{F0: 2, F1: 2}
    v17 := GenS1{F0: (v3 >> ((2) & 63)), F1: v12}
    v18 := GenS1{F0: v3, F1: 128}
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
