package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 uint16 = 1
    var v2 int32 = 2147483647
    var v3 float32 = 1.5
    v1 = (v1 % ((0) | 1))
    var v4 float32 = (0.0 + v3)
    v2 -= 371
    switch ("a" + ("foo bar" + "\n\t")) {
    case "\n\t":
        v4 = float32(0.0)
    case "hello":
        v3 = float32((0.0 / (1.5 * (0.0 - 1e10))))
    case "":
        var v5 float64 = (3.14159 + -2.25)
        _ = v5
    }
    var v6 float64 = 0.1
    fmt.Println(v4)
    v7 := GenS1{F0: "foo bar"}
    if true {
        v3 += ((((((v4 * float32(0.0)) - v3) * 0.1) / 3.14159) / (1e10 / ((float32(((0.1 / float64(float32(0.1))) * v6)) / 1e10) / (float32(3.14159) + (((v3 + v3) / 1.5) * float32(0.1)))))) + (v4 / 3.14159))
        v1 += v1
        v3 = v3
    }
    if (!((v1 >= (v1 << ((316) & 63))) && true)) {
        var v8 int = 2147483647
        v2 -= v2
        _ = v8
    }
    v9 := []int{-2147483648, 127}
    fmt.Println(v3, v1)
    v2 = (v2 >> ((575) & 63))
    v10 := GenS1{F0: "foo bar"}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v9
}
