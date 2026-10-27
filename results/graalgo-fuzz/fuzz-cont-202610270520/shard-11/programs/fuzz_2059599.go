package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0() uint8 {
    if true {
        if (!true) {
            var v6 uint16 = 39
            _ = v6
        }
        v7 := []int{-807, 127, -2147483648}
        var v8 float32 = float32(1e10)
        _ = v7
        _ = v8
    }
    var v9 float64 = (1.5 * 1e10)
    switch 4294967295 {
    case 4294967295:
        var v10 int8 = 0
        _ = v10
    case 147:
        switch "hello" {
        case "foo bar":
            var v11 int64 = 9223372036854775807
            _ = v11
        }
    }
    fmt.Println(v9)
    _ = v9
    return 102
}

func main() {
    var v1 uint64 = 1
    var v2 int32 = 0
    var v3 uint32 = 0
    v4 := []string{}
    var v5 bool = (-1 <= 1)
    var v12 uint64 = (v1 / ((42) | 1))
    fmt.Println(v1, v12, v1)
    var v13 int16 = 32767
    var v14 uint8 = 74
    v3 = v3
    v4 = append(v4, "foo bar")
    if (!(v5 && v5)) {
        var v15 float32 = float32(1e10)
        v16 := GenS1{F0: ""}
        _ = v15
        _ = v16
    }
    v3 = v3
    v13 = (v13 >> ((728) & 63))
    var v17 float64 = ((float64((0.1 * 0.1)) - 1e10) / 1.0)
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
