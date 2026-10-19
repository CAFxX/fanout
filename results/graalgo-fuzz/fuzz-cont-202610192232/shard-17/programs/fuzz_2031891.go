package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint8 = 144
    var v3 float64 = 0.0
    var v4 int8 = 1
    var v5 uint32 = 2
    var v6 int16 = 585
    switch 2147483647 {
    case 745:
        var v7 uint = 65535
        _ = v7
    }
    v8 := GenS1{F0: (928 < 0), F1: 42}
    var v9 uint32 = v5
    v1[("" + "a")] = (v5 >> ((4) & 63))
    var v10 string = "hello"
    if (!true) {
        var v11 int64 = 9223372036854775807
        var v12 int16 = 32767
        _ = v11
        _ = v12
    } else {
        v6 -= (v6 | -32768)
        if (true && (-2147483648 > 2147483647)) {
            fmt.Println(v6, v4, v10)
            if ((false && false) && (!false)) {
                v6 = v6
                var v13 float64 = v3
                _ = v13
            }
        }
    }
    v14 := []int{1, 2147483647, 951, -655}
    if len(v14) >= 2 {
        v14 = v14[1:len(v14)]
    }
    var v15 float32 = 0.0
    switch 4294967295 {
    case 2147483647:
        v10 = v10
    case -849:
        var v16 float32 = v15
        _ = v16
    case -1:
        var v17 string = v10
        _ = v17
    default:
        v18 := GenS1{F0: true, F1: 4294967295}
        _ = v18
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
