package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
func genFn0() int64 {
    v7 := []int{-199, 979, 878}
    var v8 bool = false
    var v9 uint32 = 2147483647
    var v10 bool = false
    v10 = ((3.14159 == (3.14159 * ((0.0 - (1e10 / 0.1)) - ((1e10 * 1e10) - (-2.25 + 0.0))))) || (!v10))
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return -151
}

func main() {
    var v1 float32 = 0.0
    v2 := make(map[uint32]uint32)
    var v3 bool = (true && (0 < 2147483648))
    v4 := make(map[string]uint32)
    v5 := make(map[string]int)
    var v6 int16 = 1
    v3 = v3
    if false {
        v11 := []int{2147483647, 4294967295, -1, 32768}
        var v12 uint8 = 127
        _ = v11
        _ = v12
    }
    var v13 float32 = (v1 - (0.0 + (v1 - 1.5)))
    v13 -= v13
    if false {
        v14 := []int{990}
        v14 = append(v14, -9223372036854775808)
        _ = v14
    } else {
        var v15 float32 = float32(-2.25)
        var v16 float32 = (0.0 - -2.25)
        _ = v15
        _ = v16
    }
    v17 := v5[("foo bar" + (("\n\t" + "a") + "hello"))]
    v18 := GenS1{F0: 2, F1: -135}
    if (909 < -375) {
        var v19 uint64 = 1
        _ = v19
    } else {
        var v20 uint = 0
        _ = v20
    }
    v1 += 3.14159
    v21 := GenS1{F0: 4294967295, F1: -681}
    if (65535 != 2) {
        fmt.Println(v3)
    }
    v13 = float32(0.0)
    v6 = -1
    var v22 float32 = (v1 - (((v13 - 3.14159) - v13) * v1))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
