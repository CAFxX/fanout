package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 string = "hello"
    var v2 float64 = 0.0
    v3 := make(map[int]bool)
    v4 := make(map[uint32]uint32)
    v5 := make(map[int]int)
    v6 := []int{1, 411}
    v7 := GenS1{F0: 213, F1: 0}
    var v8 int8 = -1
    _ = len(v4)
    if false {
        v8 = v8
        v8 = (v8 * -128)
        for fi9 := 0; fi9 < 12; fi9++ {
            if true {
                if false {
                    v8 = 127
                    v10 := GenS1{F0: 18446744073709551615, F1: -14}
                    _ = v10
                } else {
                    v1 = v1
                    var v11 float64 = float64((0.0 / 1.0))
                    _ = v11
                }
            }
            v2 = v2
            fmt.Println(v8)
            _ = fi9
        }
    }
    var v12 int = -9223372036854775808
    for fi13 := 0; fi13 < 6; fi13++ {
        var v14 int64 = -9223372036854775808
        _ = fi13
        _ = v14
    }
    if ((v8 - (v8 + (v8 ^ -1))) < v8) {
        var v15 float32 = (float32(v2) * 1e10)
        switch (v12 * -379) {
        case 613:
            v16 := GenS2{F0: 42}
            _ = v16
        case 128:
            v17 := GenS1{F0: 0, F1: 1}
            _ = v17
        case 32768:
            if true {
                var v18 uint64 = 18446744073709551615
                switch (((v1 + v1) + (v1 + "a")) + "\n\t") {
                case "hello":
                    var v19 int32 = -851
                    _ = v19
                case "foo bar":
                    v15 = v15
                case "a":
                    v4[0] = 912
                default:
                    var v20 int64 = 0
                    _ = v20
                }
                _ = v18
            }
        }
        var v21 int = v12
        _ = v15
        _ = v21
    }
    v12 = 2147483647
    v22 := v3[v12]
    var v23 int16 = 331
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
