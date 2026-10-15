package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
func genFn0() string {
    var v5 uint8 = 2
    var v6 uint64 = 0
    v7 := GenS1{F0: 232, F1: 4294967295}
    _ = v5
    _ = v6
    _ = v7
    return "hello"
}

func main() {
    v1 := make(map[int]int)
    var v2 uint = 4294967295
    var v3 int8 = 100
    var v4 string = (("foo bar" + "\n\t") + ("a" + ("foo bar" + "")))
    var v8 bool = (v4 <= "")
    v3 = v3
    fmt.Println(v2, v4)
    fmt.Println(v8)
    var v9 int64 = 0
    var v10 int = -1
    var v11 int = (v10 & v10)
    if v8 {
        v2 += v2
        for fi12 := 0; fi12 < 12; fi12++ {
            if v8 {
                var v13 string = v4
                _ = v13
            }
            _ = fi12
        }
    } else {
        v14 := []int{(v10 - v10), v11}
        fmt.Println(v10, v2)
        _ = v14
    }
    v15 := []int{(-(v11)), v11, v10}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
