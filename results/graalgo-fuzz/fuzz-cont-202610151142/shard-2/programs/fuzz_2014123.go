package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 uint32 }
func genFn0() string {
    var v6 uint = 18446744073709551615
    v7 := []int{}
    _ = v6
    _ = v7
    return "hello"
}
func genFn1(p0 int64, p1 uint16) uint64 {
    p0 -= 0
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    return 1
}

func main() {
    var v1 string = ""
    var v2 int8 = 0
    var v3 string = "hello"
    v4 := make(map[uint32]uint32)
    var v5 int8 = v2
    var v8 int64 = 9223372036854775807
    v9 := []int{532}
    v1 = v3
    v10 := GenS1{F0: 0, F1: 4294967295, F2: 2}
    var v11 uint64 = 368
    if len(v9) > 0 {
        v9[0] = -351
    }
    fmt.Println(v8, v8, v3)
    var v12 float32 = -2.25
    fmt.Println(v3, v2, v8)
    var v13 uint16 = 2
    if (-2147483648 != 1) {
        if len(v9) > 0 {
            v9[0] = -9223372036854775808
        }
    }
    var v14 uint64 = 1
    if true {
        if (382 <= int32(v8)) {
            var v15 bool = true
            _ = v15
        }
    } else {
        var v16 uint64 = 644
        v9 = append(v9, 32768)
        _ = v16
    }
    var v17 uint64 = 381
    v5 += v5
    fmt.Println(v1, v5, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
