package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 uint32 }
func genFn0() int16 {
    var v5 uint64 = 4294967295
    var v6 int8 = 48
    if (1 < 313) {
        fmt.Println(v6)
        if (0 != 255) {
            var v7 int16 = -1
            _ = v7
        } else {
            v8 := GenS1{F0: 42, F1: false, F2: 1}
            var v9 int = 127
            _ = v8
            _ = v9
        }
        v10 := []int{-9223372036854775808, 32767, -556}
        _ = v10
    }
    var v11 int32 = 42
    _ = v11
    _ = v5
    _ = v6
    return 42
}

func main() {
    var v1 uint8 = 2
    v2 := make(map[uint32]int)
    var v3 int = 4294967295
    var v4 int64 = 664
    for fi12 := 0; fi12 < 15; fi12++ {
        var v13 float32 = (0.1 + ((0.0 * (-2.25 + 3.14159)) / 1e10))
        fmt.Println(v3)
        var v14 float64 = 1e10
        _ = fi12
        _ = v13
        _ = v14
    }
    var v15 int = 32768
    var v16 int = v15
    v15 -= 605
    if (((4294967295 < 2147483648) && false) && true) {
        var v17 int16 = 1000
        var v18 uint16 = 0
        _ = v17
        _ = v18
    }
    v1 += (^(v1))
    var v19 float64 = 0.0
    v19 = v19
    fmt.Println(v1)
    v4 = v4
    var v20 bool = (v3 != 1)
    v1 += v1
    v2[2147483647] = v15
    v4 = v4
    var v21 uint16 = 1
    _ = len(v2)
    v1 -= 128
    var v22 int32 = -768
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
