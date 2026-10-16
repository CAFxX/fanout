package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 int }
func genFn0(p0 bool, p1 int16) int64 {
    switch "a" {
    case "hello":
        var v6 bool = p0
        _ = v6
    case "a":
        p1 -= p1
    }
    var v7 string = (("hello" + "foo bar") + (("a" + ("a" + (("" + "foo bar") + "foo bar"))) + "\n\t"))
    fmt.Println(v7)
    _ = p0
    _ = p1
    _ = v7
    return 9223372036854775807
}

func main() {
    var v1 uint64 = 2
    var v2 uint8 = 42
    var v3 uint = uint(v1)
    v4 := make(map[string]bool)
    v5 := make(map[int]string)
    var v8 int8 = 100
    var v9 float32 = -2.25
    v5[-2147483648] = (("a" + "a") + "\n\t")
    v3 = (v3 - (v3 - 65535))
    v10 := []int{4294967295, 42, -1}
    v11 := GenS1{F0: (v8 % (((v8 / ((-117) | 1))) | 1)), F1: 1000, F2: 1}
    fmt.Println(v3, v2, v1)
    fmt.Println(v3)
    var v12 int8 = v8
    if (1.5 == 1e10) {
        var v13 float64 = ((((float64(((v9 - (v9 / (0.0 - v9))) + 1e10)) + ((1e10 + 1.5) - (3.14159 - 1e10))) / 1e10) + 0.1) / -2.25)
        _ = v13
    }
    v2 -= (v2 % ((255) | 1))
    v2 += 202
    var v14 string = "\n\t"
    _ = len(v10)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
