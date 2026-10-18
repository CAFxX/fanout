package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 int32 }

func main() {
    var v1 int8 = -56
    var v2 int16 = 1
    var v3 int64 = 0
    var v4 uint32 = 2
    v1 += v1
    var v5 bool = (831 != 1)
    for fi6 := 0; fi6 < 5; fi6++ {
        v7 := GenS1{F0: 139, F1: 619, F2: -2147483648}
        _ = fi6
        _ = v7
    }
    var v8 int = 9223372036854775807
    if v5 {
        var v9 float64 = -2.25
        _ = v9
    }
    v10 := []int{32768, v8, 4294967295, (v8 % ((v8) | 1))}
    var v11 uint = 255
    var v12 uint32 = v4
    var v13 float64 = 3.14159
    for fi14 := 0; fi14 < 3; fi14++ {
        switch 4294967295 {
        case 127:
            if ((951 > int32(v8)) && (!((fi14 >= (v8 % (((fi14 << ((uint(v12)) & 63))) | 1))) && (32767 < 226)))) {
                var v15 int = -1
                var v16 bool = true
                _ = v15
                _ = v16
            }
        case 128:
            v13 -= ((v13 / (v13 + (v13 + v13))) / v13)
        }
        v5 = (!(true && true))
        _ = fi14
    }
    switch "foo bar" {
    case "a":
        switch (v8 + v8) {
        case -841:
            fmt.Println(v13, v8)
        }
    case "hello":
        fmt.Println(v12, v11)
    case "foo bar":
        var v17 uint64 = 1
        _ = v17
    default:
        v18 := GenS1{F0: 895, F1: v3, F2: -459}
        _ = v18
    }
    v2 -= v2
    var v19 int = 1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
