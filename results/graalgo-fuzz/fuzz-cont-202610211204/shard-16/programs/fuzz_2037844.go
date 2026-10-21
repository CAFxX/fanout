package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 int8 = 127
    var v2 uint8 = 99
    v3 := make(map[string]string)
    var v4 bool = (0 >= 2)
    v5 := make(map[int]int)
    var v6 float32 = 0.1
    v7 := []int{9223372036854775807, int(v2)}
    v6 -= v6
    if true {
        var v8 int64 = 359
        fmt.Println(v4)
        var v9 uint = 18446744073709551615
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: ""}
    if ((-2.25 / 0.1) == float64(0.1)) {
        var v11 uint64 = uint64(v2)
        var v12 int16 = 42
        v13 := GenS1{F0: (("foo bar" + ("a" + ("hello" + ""))) + ("foo bar" + ""))}
        _ = v11
        _ = v12
        _ = v13
    }
    var v14 int = 9223372036854775807
    var v15 float64 = (3.14159 * ((0.1 - float64(((v6 * ((((v6 + 3.14159) / v6) * v6) * v6)) * v6))) * ((float64(float32(float64(v6))) - 0.1) + (-2.25 / 1.5))))
    var v16 uint8 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
