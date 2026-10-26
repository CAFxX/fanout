package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 int8 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint32 = 42
    var v3 uint64 = 42
    v4 := []int{32768, -1, -288, 42}
    v2 = (v2 + (^(v2)))
    var v5 float32 = ((0.0 + 0.1) + 1.5)
    for fi6 := 0; fi6 < 1; fi6++ {
        v7 := GenS1{F0: 18446744073709551615, F1: ""}
        v8 := GenS1{F0: 594, F1: "\n\t"}
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 string = ("hello" + "foo bar")
    v9 = v9
    for fi10 := 0; fi10 < 7; fi10++ {
        if true {
            v3 = v3
        }
        _ = fi10
    }
    fmt.Println(v5, v9)
    var v11 string = ((v9 + "foo bar") + v9)
    v12 := GenS2{F0: 100}
    v4 = append(v4, 127)
    _ = len(v1)
    var v13 int64 = 430
    switch -1 {
    case 4294967295:
        fmt.Println(v3)
    }
    v13 += int64(v2)
    var v14 string = "hello"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
