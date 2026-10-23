package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 int16 }
type GenS2 struct { F0 uint64; F1 int64; F2 bool }
func genFn0() uint {
    var v5 string = ""
    v5 = v5
    v5 = v5
    _ = v5
    return 39
}

func main() {
    v1 := []uint32{}
    v2 := []int{32768}
    v3 := make(map[int]string)
    var v4 string = ""
    if false {
        v6 := v3[32768]
        v4 = v4
        var v7 bool = false
        _ = v6
        _ = v7
    }
    switch "foo bar" {
    case "\n\t":
        _ = len(v2)
    }
    fmt.Println(v4)
    v8 := GenS1{F0: -2147483648, F1: 32767, F2: -42}
    v9 := GenS2{F0: 4294967295, F1: 389, F2: (127 >= 127)}
    v1 = append(v1, 789)
    var v10 int16 = 1
    v11 := v3[9223372036854775807]
    if len(v2) > 0 {
        v2[0] = 32767
    }
    v2 = append(v2, 398)
    delete(v3, 42)
    var v12 uint64 = 1
    if true {
        var v13 int32 = 0
        switch 32767 {
        case 32768:
            fmt.Println(v10)
        default:
            v4 = "hello"
        }
        _ = v13
    } else {
        var v14 uint32 = 65535
        _ = v14
    }
    var v15 float64 = 0.1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
