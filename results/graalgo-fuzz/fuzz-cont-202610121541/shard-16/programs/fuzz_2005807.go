package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 uint8 }
func genFn0(p0 string, p1 int16) int {
    p0 = "foo bar"
    for fi6 := 0; fi6 < 7; fi6++ {
        switch (fi6 + fi6) {
        case 32768:
            var v7 uint16 = 65535
            _ = v7
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    return 684
}

func main() {
    var v1 uint8 = 37
    v2 := make(map[uint32]int)
    var v3 bool = false
    var v4 int64 = 9223372036854775807
    var v5 float32 = 0.0
    for fi8 := 0; fi8 < 8; fi8++ {
        fmt.Println(v4, v3, v3)
        fmt.Println(v5, v4)
        _ = fi8
    }
    fmt.Println(v5, v1, v4)
    v9 := []int{-323, -1}
    v1 = (v1 % ((v1) | 1))
    var v10 float64 = 0.0
    var v11 int8 = -1
    var v12 int = 0
    var v13 int8 = v11
    var v14 string = "hello"
    var v15 bool = (951 <= 1)
    v16 := GenS1{F0: 100, F1: 15, F2: v1}
    var v17 string = ("\n\t" + v14)
    v9 = append(v9, genFn0("", -756))
    v10 = float64(1.5)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v9:", len(v9), v9)
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
    _ = v9
}
