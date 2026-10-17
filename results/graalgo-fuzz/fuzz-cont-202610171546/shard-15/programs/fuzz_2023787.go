package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 uint = 255
    v2 := make(map[uint32]uint32)
    var v3 uint32 = 759
    v4 := make(map[int]bool)
    var v5 float32 = 1.5
    var v6 uint16 = 1
    var v7 int16 = -32768
    var v8 string = "foo bar"
    v9 := GenS1{F0: (false || (!((!(true && true)) || (1 == 32768))))}
    v10 := GenS1{F0: (true && (3.14159 < 1.5))}
    v8 = v8
    fmt.Println(v6)
    var v11 string = v8
    var v12 uint64 = 42
    var v13 uint = (v1 ^ (v1 * (v1 ^ 2)))
    v14 := []int{-9223372036854775808, 675}
    var v15 int16 = (v7 * -1)
    for fi16 := 0; fi16 < 13; fi16++ {
        var v17 float32 = 1e10
        _ = fi16
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
