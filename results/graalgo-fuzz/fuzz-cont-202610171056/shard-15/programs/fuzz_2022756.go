package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0() int8 {
    var v7 float32 = 3.14159
    for fi8 := 0; fi8 < 5; fi8++ {
        for fi9 := 0; fi9 < 1; fi9++ {
            fmt.Println(fi9, fi8, fi8)
            v7 = v7
            _ = fi9
        }
        _ = fi8
    }
    switch -2147483648 {
    case 2147483647:
        var v10 uint8 = 244
        _ = v10
    case 9223372036854775807:
        fmt.Println(v7)
    case 32767:
        var v11 uint = 596
        _ = v11
    }
    _ = v7
    return -121
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]uint32)
    var v3 float32 = 1.5
    var v4 float64 = float64(0.0)
    v5 := make(map[string]string)
    var v6 string = "hello"
    var v12 int8 = genFn0()
    v6 = v6
    delete(v1, 288)
    var v13 float64 = (v4 - 1e10)
    fmt.Println(v12, v3, v3)
    v14 := GenS1{F0: (v6 + v6)}
    var v15 int64 = 9223372036854775807
    var v16 float64 = 0.1
    v3 = float32((1.5 - float64((-2.25 - (((0.1 - (v3 + -2.25)) - v3) + v3)))))
    fmt.Println(v15, v6, v16)
    v1[-2147483648] = 839
    if (2147483648 <= 287) {
        v12 -= genFn0()
        fmt.Println(v15)
    }
    fmt.Println(v13, v16)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
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
}
