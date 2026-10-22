package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }

func main() {
    var v1 bool = false
    var v2 uint16 = 2
    v3 := make(map[string]int)
    var v4 int16 = -32768
    v5 := make(map[string]bool)
    var v6 int64 = 9223372036854775807
    v2 += (v2 % (((v2 >> ((18446744073709551615) & 63))) | 1))
    v7 := GenS1{F0: (v4 | v4), F1: 255}
    v8 := GenS1{F0: (^(v4)), F1: v2}
    var v9 float32 = 1.5
    v2 += (v2 | v2)
    for fi10 := 0; fi10 < 2; fi10++ {
        var v11 float64 = 1e10
        _ = fi10
        _ = v11
    }
    var v12 int64 = (-(v6))
    v4 += -1
    var v13 bool = false
    var v14 float64 = float64(1e10)
    v13 = ((v12 ^ -1) <= 139)
    var v15 bool = v13
    v16 := []int{-626, 128, 4294967295, 42}
    switch 0 {
    case -2147483648:
        var v17 int = 32767
        _ = v17
    case 127:
        fmt.Println(v2, v1, v2)
    default:
        v13 = (3.14159 >= v9)
    }
    if false {
        var v18 int32 = -2147483648
        v1 = true
        _ = v18
    } else {
        fmt.Println(v1)
        if v15 {
            var v19 uint32 = 337
            var v20 float64 = v14
            var v21 string = "a"
            _ = v19
            _ = v20
            _ = v21
        }
        var v22 string = "a"
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
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
    _ = v7
    _ = v8
    _ = v9
}
