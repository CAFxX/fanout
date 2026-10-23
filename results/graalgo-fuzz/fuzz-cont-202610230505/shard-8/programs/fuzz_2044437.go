package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 int }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[uint32]uint32)
    var v3 int16 = 0
    v4 := make(map[uint32]string)
    var v5 string = "hello"
    var v6 string = (((v5 + (("a" + v5) + (v5 + (v5 + ("" + v5))))) + "\n\t") + "foo bar")
    if (1 > 1) {
        var v7 int8 = -101
        var v8 uint64 = 0
        _ = v7
        _ = v8
    } else {
        v6 = "foo bar"
    }
    _ = len(v2)
    v9 := []int{-798, 127, 508}
    fmt.Println(v3, v3, v3)
    switch 2147483647 {
    case 677:
        if len(v9) > 0 {
            v9[0] = -98
        }
    case -1:
        v10 := GenS1{F0: (-2.25 < 1e10), F1: (^(v3)), F2: -9223372036854775808}
        _ = v10
    default:
        var v11 int16 = (v3 + (v3 + v3))
        _ = v11
    }
    v12 := GenS1{F0: ((1.5 == (-2.25 - float32(1e10))) || true), F1: (-(v3)), F2: int(v3)}
    v13 := GenS1{F0: ((!(2 <= 42)) || true), F1: (v3 % ((-32768) | 1)), F2: 32767}
    var v14 uint64 = 4294967295
    var v15 int = -884
    v16 := GenS1{F0: true, F1: (v3 % (((-(v3))) | 1)), F2: (^(v15))}
    if len(v9) > 0 {
        v9[0] = v15
    }
    if (255 > 78) {
        fmt.Println(v5, v6)
    } else {
        delete(v1, "\n\t")
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
