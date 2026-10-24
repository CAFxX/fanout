package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }

func main() {
    var v1 string = "hello"
    var v2 uint32 = 825
    v3 := make(map[string]string)
    v4 := make(map[int]string)
    v5 := []int{}
    for fi6 := 0; fi6 < 11; fi6++ {
        var v7 uint = 18446744073709551615
        v7 += v7
        _ = fi6
        _ = v7
    }
    fmt.Println(v1, v1)
    v1 = "foo bar"
    for fi8 := 0; fi8 < 3; fi8++ {
        v2 = (v2 % (((v2 | (v2 % ((42) | 1)))) | 1))
        _ = fi8
    }
    var v9 uint = 255
    fmt.Println(v1, v2, v2)
    _ = len(v4)
    if (-360 == 32768) {
        if false {
            var v10 int = 1
            _ = v10
        }
    }
    v11 := GenS1{F0: 255, F1: 1000}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
