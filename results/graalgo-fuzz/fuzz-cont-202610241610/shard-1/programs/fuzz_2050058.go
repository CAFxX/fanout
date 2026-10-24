package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
func genFn0() uint {
    v6 := GenS1{F0: 255, F1: 255}
    var v7 uint32 = 65535
    var v8 uint = 2
    _ = v6
    _ = v7
    _ = v8
    return v8
}

func main() {
    v1 := []int{564, 929, -2147483648, 9223372036854775807}
    v2 := make(map[string]string)
    v3 := make(map[string]int)
    var v4 int8 = 1
    var v5 int16 = -696
    var v9 int64 = 282
    for fi10 := 0; fi10 < 8; fi10++ {
        v1 = append(v1, -2147483648)
        _ = fi10
    }
    var v11 float32 = 0.0
    v12 := GenS1{F0: 943, F1: 806}
    var v13 int32 = 1
    for fi14 := 0; fi14 < 14; fi14++ {
        v13 += v13
        v15 := GenS1{F0: 32767, F1: 125}
        v11 -= v11
        _ = fi14
        _ = v15
    }
    v13 = 1
    fmt.Println(v5)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v16 int32 = (^(v13))
    fmt.Println(v11)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
