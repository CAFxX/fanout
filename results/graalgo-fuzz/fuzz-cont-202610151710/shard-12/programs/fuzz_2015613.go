package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    var v1 float32 = 0.1
    var v2 float32 = float32((3.14159 + 0.1))
    var v3 float32 = (v1 * float32((((0.0 / (3.14159 / float64(v1))) + 3.14159) * 0.0)))
    var v4 bool = (!(false || false))
    var v5 float64 = -2.25
    var v6 uint8 = 128
    v7 := []int{4294967295, 9223372036854775807, 2147483647}
    v2 = v3
    var v8 string = "\n\t"
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    v1 -= v3
    var v9 uint = 0
    fmt.Println(v3, v5)
    v10 := GenS1{F0: 889}
    fmt.Println(v2)
    for fi11 := 0; fi11 < 5; fi11++ {
        var v12 bool = v4
        var v13 uint16 = 2
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 uint8 = (v6 % ((42) | 1))
    var v15 int64 = 406
    var v16 uint64 = 0
    var v17 int = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
