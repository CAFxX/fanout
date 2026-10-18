package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
func genFn0() uint64 {
    v4 := []int{702, 127, 5, 2147483647}
    if ((42 <= 206) || true) {
        var v5 uint64 = 0
        switch "a" {
        case "a":
            var v6 bool = (-32768 < 1)
            _ = v6
        case "":
            var v7 float64 = 1.5
            _ = v7
        case "foo bar":
            v5 = (v5 | v5)
        }
        _ = v5
    }
    _ = v4
    return 710
}
func genFn1() uint {
    v8 := GenS1{F0: 63, F1: 2147483647}
    var v9 uint16 = 1
    v10 := []int{-119}
    fmt.Println(v9)
    var v11 string = (("a" + "\n\t") + "a")
    _ = v10
    _ = v11
    _ = v8
    _ = v9
    return 977
}

func main() {
    var v1 string = "\n\t"
    var v2 int8 = -109
    v3 := make(map[int]string)
    var v12 int8 = v2
    v13 := []int{32767}
    if (1 >= -1) {
        var v14 uint16 = 255
        _ = v14
    }
    switch "a" {
    case "":
        v15 := GenS1{F0: 42, F1: 0}
        _ = v15
    case "\n\t":
        delete(v3, -758)
    }
    fmt.Println(v2)
    var v16 int64 = 42
    if true {
        v12 = int8(v16)
        v17 := GenS1{F0: 87, F1: 128}
        for fi18 := 0; fi18 < 8; fi18++ {
            var v19 uint32 = 2147483647
            _ = fi18
            _ = v19
        }
        _ = v17
    }
    var v20 uint = 0
    var v21 uint16 = 255
    v16 = v16
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
