package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]int)
    var v3 bool = true
    v4 := []string{"foo bar", ("hello" + (("hello" + ("a" + ((("" + "foo bar") + "hello") + "a"))) + "\n\t")), "hello", ((("a" + ("\n\t" + "a")) + "foo bar") + "a")}
    v5 := []int{-666}
    v6 := make(map[string]bool)
    fmt.Println(v3)
    fmt.Println(v3)
    v7 := GenS1{F0: 725}
    var v8 int64 = -1
    v9 := GenS1{F0: 2147483647}
    var v10 int64 = 360
    _ = len(v5)
    _ = len(v5)
    var v11 float64 = 3.14159
    var v12 int16 = 0
    v5 = append(v5, -836)
    if true {
        var v13 uint32 = 44
        v11 -= 0.0
        _ = v13
    }
    fmt.Println(v8, v12)
    v8 = v10
    fmt.Println(v11, v10, v11)
    var v14 int8 = -128
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
