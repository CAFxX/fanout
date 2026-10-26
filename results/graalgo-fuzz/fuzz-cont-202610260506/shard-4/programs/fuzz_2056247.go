package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
func genFn0(p0 int32, p1 string, p2 string) int8 {
    v4 := GenS1{F0: p2, F1: ((p1 + p1) + p2)}
    if (!true) {
        v5 := []int{1}
        var v6 int32 = (p0 & -2147483648)
        _ = v5
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return -125
}

func main() {
    var v1 uint8 = 128
    var v2 uint8 = v1
    var v3 int16 = 0
    v3 -= (v3 % ((-32768) | 1))
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 uint64 = 18446744073709551615
        var v9 int64 = 0
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 string = "foo bar"
    v10 = v10
    var v11 uint = 524
    v12 := GenS1{F0: v10, F1: ((v10 + v10) + v10)}
    fmt.Println(v10)
    var v13 bool = ("\n\t" != "a")
    var v14 int16 = -1
    var v15 int32 = 1
    v10 = v10
    var v16 uint64 = 957
    var v17 float32 = 3.14159
    v18 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
