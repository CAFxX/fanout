package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 uint8; F2 uint64 }

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[uint32]int)
    var v3 float32 = 1.5
    var v4 uint8 = 48
    var v5 int16 = 42
    var v6 string = "a"
    v7 := []int{-238, 0, -9223372036854775808}
    switch 0 {
    case 128:
        v5 = -32768
    }
    switch 32767 {
    case -9223372036854775808:
        var v8 int16 = v5
        _ = v8
    }
    v6 = "hello"
    var v9 int8 = 127
    v10 := GenS1{F0: 255}
    for fi11 := 0; fi11 < 1; fi11++ {
        switch ("hello" + v6) {
        case "hello":
            v5 = (v5 + v5)
        }
        _ = fi11
    }
    v12 := GenS2{F0: v6, F1: 2, F2: 4294967295}
    for fi13 := 0; fi13 < 11; fi13++ {
        var v14 int8 = (-(v9))
        fmt.Println(v5)
        if (fi13 >= 9223372036854775807) {
            v15 := GenS2{F0: v6, F1: v4, F2: 42}
            v1 = v1
            v14 = v14
            _ = v15
        }
        _ = fi13
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
