package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint }

func main() {
    v1 := make(map[string]bool)
    var v2 bool = false
    var v3 int = -2147483648
    var v4 uint = 65535
    var v5 string = ("a" + ("a" + ((("\n\t" + "a") + "a") + "")))
    v6 := GenS1{F0: false}
    v7 := []int{v3, int(v4)}
    v8 := GenS2{F0: v4}
    if len(v7) > 0 {
        v7[0] = (-(v3))
    }
    v9 := v1[(v5 + (v5 + v5))]
    if (!(false || v2)) {
        var v10 int16 = 1
        var v11 uint8 = 127
        _ = v10
        _ = v11
    } else {
        fmt.Println(v2, v3)
        var v12 bool = true
        _ = v12
    }
    if (1000 != -1) {
        var v13 int64 = 9223372036854775807
        v7 = append(v7, 348)
        _ = v13
    }
    var v14 uint32 = 42
    delete(v1, v5)
    v2 = (516 >= 784)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v9:", v9)
    _ = v1
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
