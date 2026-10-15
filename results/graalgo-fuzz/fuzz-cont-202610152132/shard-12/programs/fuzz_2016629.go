package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 uint8 }
func genFn0(p0 string) int {
    v7 := []int{208, 423, -2147483648, 9223372036854775807}
    p0 = p0
    p0 = p0
    _ = p0
    _ = v7
    return -677
}

func main() {
    var v1 uint32 = 2147483648
    var v2 uint32 = v1
    var v3 int64 = -1
    var v4 uint = 4294967295
    v5 := make(map[int]string)
    var v6 uint8 = 127
    v8 := GenS1{F0: v6, F1: (v4 / ((18446744073709551615) | 1)), F2: 127}
    var v9 uint = 65535
    v10 := []int{1, 127}
    fmt.Println(v1)
    v11 := GenS1{F0: v6, F1: v4, F2: v6}
    var v12 int8 = 42
    switch -2147483648 {
    case 4294967295:
        v1 -= v2
    case 127:
        v13 := GenS1{F0: v6, F1: v9, F2: v6}
        _ = v13
    }
    var v14 uint32 = (^(v2))
    v3 = v3
    var v15 uint32 = 2
    var v16 uint32 = (v1 % ((42) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
