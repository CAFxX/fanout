package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
func genFn0(p0 int32) uint {
    var v6 uint16 = 1
    v7 := []int{-9223372036854775808, 1, -1, 128}
    for fi8 := 0; fi8 < 16; fi8++ {
        var v9 float32 = 1.5
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = v6
    _ = v7
    return 1
}

func main() {
    v1 := []int{}
    var v2 string = "foo bar"
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]string)
    v5 := []uint32{1, 820, 42}
    if true {
        v2 = v2
        var v10 uint16 = 65535
        v11 := GenS1{F0: (811 >= genFn0(-2147483648)), F1: v2}
        _ = v10
        _ = v11
    }
    v2 = ""
    var v12 int16 = -32768
    v12 += 285
    fmt.Println(v12, v2)
    fmt.Println(v12, v12)
    var v13 bool = false
    _ = len(v1)
    v14 := GenS1{F0: (160 != 0), F1: ("hello" + "\n\t")}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
