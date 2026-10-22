package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 int32 = 817
    var v2 bool = true
    var v3 string = "hello"
    v4 := []int{42, 434, 0, 127}
    var v5 float64 = 0.0
    v6 := GenS1{F0: 4294967295}
    var v7 bool = true
    fmt.Println(v5, v7)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if v2 {
        var v8 int64 = 0
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v8
    } else {
        v3 = v3
        fmt.Println(v5, v5)
    }
    var v9 uint64 = 93
    var v10 int16 = 89
    v11 := GenS1{F0: 42}
    var v12 float64 = float64(float32(-2.25))
    _ = len(v4)
    _ = len(v4)
    var v13 uint8 = 1
    switch 4294967295 {
    case 32767:
        v5 = (v12 + -2.25)
    case 1:
        var v14 int = -2147483648
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
