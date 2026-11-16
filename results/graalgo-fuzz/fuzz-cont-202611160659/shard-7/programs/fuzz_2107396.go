package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0() int32 {
    var v4 bool = (-32768 <= 0)
    var v5 uint64 = 2
    _ = v4
    _ = v5
    return -620
}
func genFn1(p0 int) int16 {
    var v6 uint16 = 381
    for fi7 := 0; fi7 < 12; fi7++ {
        v8 := []int{}
        var v9 uint8 = 195
        p0 = 42
        _ = fi7
        _ = v8
        _ = v9
    }
    v6 = (^(v6))
    var v10 bool = (-1 == 0)
    _ = p0
    _ = v10
    _ = v6
    return -139
}

func main() {
    var v1 int = -2147483648
    var v2 string = "foo bar"
    var v3 uint32 = 561
    fmt.Println(v3, v3, v2)
    var v11 uint16 = 42
    v12 := []int{(v1 >> ((4294967295) & 63)), (v1 + v1), (v1 / ((32768) | 1))}
    var v13 float64 = float64(1e10)
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    var v14 int64 = -1
    var v15 uint64 = 42
    _ = len(v12)
    v15 -= 167
    v16 := GenS1{F0: int(v14)}
    v17 := GenS1{F0: v1}
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
