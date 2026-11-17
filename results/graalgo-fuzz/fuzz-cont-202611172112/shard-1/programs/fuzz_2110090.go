package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 uint16 }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]uint32)
    var v3 string = "a"
    v4 := []int{-545, 128}
    v3 = (v3 + v3)
    v4 = append(v4, 32768)
    for fi5 := 0; fi5 < 15; fi5++ {
        for fi6 := 0; fi6 < 11; fi6++ {
            var v7 uint = 18446744073709551615
            var v8 uint32 = 415
            v9 := GenS1{F0: int16(v8), F1: true, F2: 65535}
            _ = fi6
            _ = v7
            _ = v8
            _ = v9
        }
        if (true && true) {
            var v10 bool = true
            var v11 uint = 42
            _ = v10
            _ = v11
        }
        _ = fi5
    }
    if true {
        var v12 uint = 90
        v12 = (v12 | v12)
        _ = v12
    }
    if false {
        fmt.Println(v3)
        v3 = v3
    } else {
        v13 := GenS1{F0: 1, F1: false, F2: 939}
        _ = v13
    }
    v3 = ((v3 + "") + v3)
    v3 = "hello"
    v3 = "\n\t"
    v14 := v1[v3]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
