package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 bool }
func genFn0() uint32 {
    if (42 <= 255) {
        v7 := []int{-1, -243, -93, 32767}
        _ = v7
    }
    for fi8 := 0; fi8 < 3; fi8++ {
        for fi9 := 0; fi9 < 10; fi9++ {
            var v10 uint32 = 771
            v10 += (v10 ^ v10)
            var v11 int64 = -665
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = fi8
    }
    if false {
        v12 := GenS1{F0: 1}
        v13 := GenS2{F0: 255, F1: (true || (18446744073709551615 > 155))}
        _ = v12
        _ = v13
    }
    var v14 string = ((("hello" + "hello") + "foo bar") + "\n\t")
    _ = v14
    return 853
}

func main() {
    v1 := []uint32{1, 490, 190, 2}
    var v2 uint = 18446744073709551615
    var v3 bool = true
    v4 := make(map[int]bool)
    var v5 uint8 = 42
    var v6 int64 = 0
    v15 := GenS1{F0: 128}
    v2 += v2
    var v16 uint64 = 18446744073709551615
    v2 = (v2 & (v2 | (v2 >> (((v2 ^ uint(v5))) & 63))))
    for fi17 := 0; fi17 < 9; fi17++ {
        if (!(100 >= -1)) {
            var v18 int16 = -1
            _ = v18
        }
        _ = fi17
    }
    v3 = v3
    v1 = append(v1, 42)
    var v19 int = 2147483647
    switch "foo bar" {
    case "":
        var v20 uint64 = 42
        _ = v20
    case "a":
        var v21 uint16 = 1
        _ = v21
    default:
        fmt.Println(v16)
    }
    if v3 {
        v22 := GenS1{F0: (^(v19))}
        var v23 float32 = 0.0
        _ = v22
        _ = v23
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
