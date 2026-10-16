package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
func genFn0(p0 uint16, p1 bool, p2 uint64) int16 {
    v7 := GenS1{F0: 395, F1: (!(p1 || (!false)))}
    if (2147483648 > 2) {
        p2 += (^(p2))
    }
    switch "a" {
    case "\n\t":
        var v8 int = 2147483647
        _ = v8
    default:
        v9 := GenS1{F0: -494, F1: p1}
        _ = v9
    }
    var v10 int32 = 1
    switch "" {
    case "a":
        fmt.Println(p2, p1, p1)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    return 32767
}

func main() {
    v1 := make(map[string]int)
    var v2 int8 = 42
    var v3 int16 = 1000
    var v4 uint = 4294967295
    v5 := []int{}
    v6 := []string{"\n\t", ((("hello" + "a") + "foo bar") + "\n\t"), ((((("hello" + "foo bar") + "") + ("hello" + "")) + "hello") + "")}
    fmt.Println(v3, v2, v3)
    var v11 uint16 = 32767
    var v12 uint64 = 4294967295
    var v13 uint32 = 2
    v12 = (v12 >> ((v4) & 63))
    v14 := v1["a"]
    v1[""] = (-(v14))
    fmt.Println(v11)
    fmt.Println(v2)
    v2 = v2
    v15 := v1[(((("a" + "a") + "a") + "hello") + "hello")]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
