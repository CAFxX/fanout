package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 int32 }
type GenS2 struct { F0 bool; F1 uint32 }

func main() {
    v1 := make(map[string]bool)
    var v2 uint = 1
    v3 := make(map[uint32]bool)
    var v4 uint = v2
    v5 := []string{("hello" + ""), "a", ""}
    v6 := make(map[int]bool)
    var v7 uint32 = 83
    var v8 uint64 = 465
    fmt.Println(v7, v4, v2)
    v8 = v8
    switch -2147483648 {
    case 1:
        for fi9 := 0; fi9 < 4; fi9++ {
            v10 := v3[915]
            fmt.Println(v10)
            _ = fi9
            _ = v10
        }
    }
    fmt.Println(v8, v2, v7)
    v8 += v8
    if (-32768 == -365) {
        delete(v6, 42)
        v8 = 287
    } else {
        v5 = append(v5, "a")
    }
    var v11 string = "hello"
    var v12 uint32 = (^(v7))
    var v13 uint8 = 127
    if false {
        if ((v2 * v2) != 463) {
            fmt.Println(v13, v4)
            var v14 bool = (float32(1e10) != (float32((3.14159 - float64((-2.25 + (3.14159 + 0.0))))) + (1e10 / ((float32(1e10) - float32(-2.25)) / 1.5))))
            var v15 uint = (v4 + 1)
            _ = v14
            _ = v15
        }
        v4 += uint(v7)
        var v16 int64 = int64(v7)
        _ = v16
    } else {
        var v17 uint8 = v13
        var v18 bool = true
        v19 := GenS1{F0: 0, F1: true, F2: -2147483648}
        _ = v17
        _ = v18
        _ = v19
    }
    var v20 float64 = 0.1
    var v21 uint = v4
    v22 := GenS2{F0: true, F1: (v7 - 65535)}
    v23 := GenS1{F0: -469, F1: false, F2: -2147483648}
    switch v11 {
    case "":
        var v24 int64 = 1
        _ = v24
    case "hello":
        _ = len(v6)
    default:
        var v25 uint64 = 4294967295
        _ = v25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
