package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 int8 }
type GenS2 struct { F0 int64; F1 int8 }

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[int]bool)
    var v3 int16 = 0
    for fi4 := 0; fi4 < 13; fi4++ {
        fmt.Println(fi4)
        _ = fi4
    }
    if (-32768 == v3) {
        v5 := GenS1{F0: 1, F1: (v1 * 42), F2: 42}
        v3 += (v3 + v3)
        fmt.Println(v1, v3)
        _ = v5
    }
    var v6 float32 = 3.14159
    var v7 int = 1
    if false {
        var v8 int = 618
        var v9 float64 = -2.25
        if (-2147483648 > v1) {
            v6 = (v6 + 1e10)
        } else {
            switch v8 {
            case 32767:
                fmt.Println(v7, v1)
            case 926:
                v10 := []int{v7, 1, (v7 % (((v8 & (v7 & (v8 & v7)))) | 1))}
                _ = v10
            case -38:
                var v11 float32 = v6
                _ = v11
            }
        }
        _ = v8
        _ = v9
    }
    var v12 int = -9223372036854775808
    v6 = v6
    var v13 int32 = v1
    v14 := GenS2{F0: -1, F1: 95}
    v15 := []int{(v7 * 127), 42}
    var v16 float64 = ((-2.25 / 3.14159) * 1e10)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
