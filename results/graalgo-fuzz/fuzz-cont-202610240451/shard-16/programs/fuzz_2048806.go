package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32; F1 bool }
func genFn0(p0 string, p1 int32) int {
    var v6 int8 = 1
    var v7 int = 32767
    fmt.Println(p1, v7)
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return (^(v7))
}

func main() {
    var v1 int16 = 371
    v2 := make(map[int]string)
    v3 := make(map[uint32]bool)
    var v4 float32 = ((1e10 - 1e10) * (0.0 - 0.0))
    var v5 string = ("a" + "a")
    var v8 uint8 = 1
    var v9 int8 = 32
    if (479 <= 1) {
        for fi10 := 0; fi10 < 9; fi10++ {
            var v11 int16 = v1
            v12 := GenS1{F0: 0}
            _ = fi10
            _ = v11
            _ = v12
        }
        var v13 string = ""
        _ = v13
    }
    switch (v5 + v5) {
    case "a":
        v8 += v8
    case "\n\t":
        var v14 uint64 = 1
        _ = v14
    }
    v15 := []int{929}
    v16 := GenS1{F0: 27}
    var v17 uint8 = 42
    fmt.Println(v4, v1, v8)
    var v18 uint8 = v8
    v19 := GenS2{F0: 352, F1: (!(-2147483648 >= genFn0(v5, -461)))}
    var v20 uint32 = 2
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
