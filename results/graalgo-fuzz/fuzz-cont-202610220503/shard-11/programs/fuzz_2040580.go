package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 int }
func genFn0() int32 {
    var v7 uint32 = 2
    v8 := GenS1{F0: 42}
    _ = v7
    _ = v8
    return 944
}
func genFn1() uint {
    var v9 uint8 = 192
    var v10 float32 = 3.14159
    v11 := GenS1{F0: v9}
    v9 -= (^(v9))
    _ = v10
    _ = v11
    _ = v9
    return 18446744073709551615
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []int{42, 32768, 32768, 42}
    var v3 float32 = 3.14159
    v4 := make(map[int]int)
    v5 := make(map[int]string)
    v6 := []uint32{644, 468, 596, 2147483647}
    var v12 float32 = float32(1.5)
    var v13 uint8 = 128
    v2 = append(v2, -1)
    v13 = (v13 + 1)
    var v14 int64 = -529
    var v15 uint = 28
    v16 := GenS1{F0: (^(v13))}
    switch "a" {
    case "a":
        switch ("\n\t" + (("\n\t" + "a") + "")) {
        case "foo bar":
            v17 := GenS2{F0: v14, F1: 4294967295}
            _ = v17
        default:
            v18 := GenS2{F0: v14, F1: 715}
            _ = v18
        }
    case "\n\t":
        var v19 uint = v15
        _ = v19
    case "foo bar":
        v14 -= v14
    default:
        if (v13 != 202) {
            var v20 float64 = float64(0.1)
            var v21 bool = true
            var v22 uint64 = 878
            _ = v20
            _ = v21
            _ = v22
        } else {
            if true {
                v23 := GenS2{F0: v14, F1: 335}
                _ = v23
            }
        }
    }
    var v24 int = 32768
    v12 = v12
    var v25 bool = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
