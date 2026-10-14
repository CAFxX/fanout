package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 int8 }

func main() {
    var v1 string = ""
    var v2 string = v1
    v3 := make(map[uint32]uint32)
    var v4 int16 = 31
    var v5 int8 = 1
    v6 := make(map[uint32]int)
    v7 := []int{-2147483648}
    var v8 bool = false
    v2 = ""
    v4 = (v4 & v4)
    for fi9 := 0; fi9 < 2; fi9++ {
        v10 := GenS1{F0: 626, F1: (v4 >> ((1) & 63)), F2: v5}
        v8 = true
        var v11 uint = 471
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = len(v7)
    var v12 int8 = v5
    v2 = v1
    v2 = ("hello" + "\n\t")
    fmt.Println(v8, v5, v8)
    v12 = v5
    v8 = (255 <= 255)
    var v13 int = 831
    for fi14 := 0; fi14 < 4; fi14++ {
        var v15 float32 = (float32(((0.0 - 1.5) - (((1e10 / 1.0) / 3.14159) - 1.5))) * 1.5)
        for fi16 := 0; fi16 < 10; fi16++ {
            fmt.Println(v4)
            _ = fi16
        }
        _ = fi14
        _ = v15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
