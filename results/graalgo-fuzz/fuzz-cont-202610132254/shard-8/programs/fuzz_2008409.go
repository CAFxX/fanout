package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 uint8 }
func genFn0(p0 string, p1 int64) uint {
    v6 := []int{}
    p0 = p0
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = v6
    return 42
}

func main() {
    var v1 int32 = 0
    v2 := make(map[uint32]int)
    v3 := make(map[int]int)
    var v4 float32 = 1e10
    v5 := []string{"hello", ("foo bar" + ("hello" + ("hello" + "\n\t"))), "hello"}
    switch ("hello" + ("a" + (("foo bar" + ("hello" + ("\n\t" + "hello"))) + "hello"))) {
    case "hello":
        v7 := GenS1{F0: 255}
        _ = v7
    }
    if true {
        var v8 string = "hello"
        v9 := GenS2{F0: -32768, F1: 90}
        _ = v8
        _ = v9
    }
    switch ("" + ("\n\t" + "hello")) {
    case "":
        v4 = v4
    case "\n\t":
        fmt.Println(v4, v1)
    case "foo bar":
        fmt.Println(v1, v1)
    default:
        fmt.Println(v4, v4)
    }
    if true {
        fmt.Println(v1)
        v4 = float32(-2.25)
    }
    v1 -= v1
    fmt.Println(v4, v4)
    fmt.Println(v4, v1)
    fmt.Println(v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
