package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 string }

func main() {
    v1 := []string{("foo bar" + "hello")}
    var v2 uint16 = 0
    v3 := make(map[int]int)
    v4 := make(map[uint32]string)
    v5 := GenS1{F0: (v2 | (v2 + (v2 >> ((1) & 63)))), F1: 2147483647}
    if true {
        var v6 int64 = -397
        _ = v6
    }
    var v7 string = ("\n\t" + "foo bar")
    fmt.Println(v7, v7)
    if (!(true && (!(639 == v2)))) {
        fmt.Println(v2, v7)
        var v8 int8 = 1
        _ = v8
    }
    v9 := GenS2{F0: "a"}
    if true {
        var v10 float64 = (3.14159 + (-2.25 * 0.1))
        delete(v3, -2147483648)
        v1 = append(v1, ((((((v7 + "a") + (v7 + v7)) + "hello") + (v7 + "")) + v7) + v7))
        _ = v10
    }
    var v11 float32 = 3.14159
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v9
}
