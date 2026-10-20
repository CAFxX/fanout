package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 uint64 }
func genFn0() int64 {
    var v6 string = ("foo bar" + "foo bar")
    v7 := GenS1{F0: 323, F1: 42, F2: 18446744073709551615}
    v6 = v6
    if true {
        v8 := GenS1{F0: 42, F1: -701, F2: 42}
        var v9 uint16 = 281
        _ = v8
        _ = v9
    } else {
        fmt.Println(v6)
        v6 = (v6 + v6)
        v6 = v6
    }
    v6 = "\n\t"
    _ = v6
    _ = v7
    return -464
}

func main() {
    var v1 float64 = 0.0
    v2 := []int{0}
    var v3 uint16 = 255
    v4 := make(map[uint32]bool)
    v5 := []uint32{225}
    v1 = v1
    var v10 uint64 = 709
    v11 := GenS1{F0: 0, F1: 32767, F2: (v10 + (v10 ^ v10))}
    var v12 float32 = 1.5
    fmt.Println(v3, v1)
    var v13 uint = 1
    var v14 int = 32767
    v15 := GenS1{F0: -2147483648, F1: -32768, F2: v10}
    var v16 uint = 917
    var v17 uint = (v13 & v16)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
