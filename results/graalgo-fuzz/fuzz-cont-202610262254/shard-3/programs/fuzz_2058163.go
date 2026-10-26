package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 int32; F1 uint32 }
func genFn0(p0 int16, p1 int) string {
    switch "" {
    case "a":
        var v4 float64 = float64(1.5)
        _ = v4
    case "hello":
        v5 := []int{p1}
        _ = v5
    case "\n\t":
        var v6 int64 = 1
        _ = v6
    }
    p1 = p1
    _ = p0
    _ = p1
    return ("\n\t" + ("hello" + "foo bar"))
}

func main() {
    var v1 bool = (!(-804 >= 1))
    var v2 int = -2147483648
    var v3 uint32 = 747
    v2 += v2
    fmt.Println(v1)
    v7 := []int{v2}
    var v8 uint8 = 229
    var v9 float32 = ((1e10 * (3.14159 / (1.5 - -2.25))) * 1e10)
    v2 -= v2
    v10 := GenS1{F0: 127, F1: 259}
    var v11 uint = 904
    v12 := GenS2{F0: 1, F1: (v3 + 255)}
    fmt.Println(v3, v2)
    var v13 bool = ((v11 - v11) != v11)
    switch "" {
    case "\n\t":
        fmt.Println(v8, v3)
    case "":
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
