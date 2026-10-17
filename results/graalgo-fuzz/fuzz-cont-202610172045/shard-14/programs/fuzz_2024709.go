package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int; F1 string; F2 uint }
func genFn0(p0 int8) int {
    v5 := []int{1, -1, 48, 83}
    v6 := GenS1{F0: 374}
    var v7 uint16 = 649
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 127
}
func genFn1(p0 uint16, p1 bool, p2 int64) uint32 {
    v8 := []int{}
    v9 := GenS1{F0: 435}
    if true {
        fmt.Println(p1)
        var v10 string = "\n\t"
        switch v10 {
        case "\n\t":
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
        case "foo bar":
            if p1 {
                p1 = p1
                v10 = "\n\t"
            }
        }
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 4294967295
}

func main() {
    var v1 int8 = 87
    var v2 int8 = 127
    var v3 bool = true
    var v4 uint = 65535
    fmt.Println(v2, v4, v2)
    v11 := []int{988}
    var v12 uint8 = 1
    var v13 uint64 = 2
    v13 = v13
    v14 := GenS2{F0: 0, F1: ("" + ""), F2: 255}
    var v15 int8 = v1
    v12 = 86
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
