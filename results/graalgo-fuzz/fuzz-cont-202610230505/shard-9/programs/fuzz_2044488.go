package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 string; F1 uint8 }

func main() {
    var v1 int8 = 1
    v2 := make(map[uint32]bool)
    var v3 float32 = 3.14159
    var v4 int8 = -128
    _ = len(v2)
    var v5 int = -992
    for fi6 := 0; fi6 < 17; fi6++ {
        v1 += -1
        _ = fi6
    }
    var v7 string = (("a" + (("hello" + "a") + "a")) + (("hello" + "a") + "hello"))
    switch "" {
    case "a":
        v8 := GenS1{F0: int32(v5), F1: 4294967295}
        _ = v8
    }
    v9 := GenS2{F0: v7, F1: 2}
    fmt.Println(v3)
    if false {
        var v10 uint16 = 886
        _ = v10
    }
    v5 -= 32767
    if true {
        var v11 int = v5
        v12 := []int{}
        _ = v11
        _ = v12
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v9
}
