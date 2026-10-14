package main

import (
    "fmt"
)


func main() {
    var v1 string = ("" + "")
    v2 := []string{("" + ((("foo bar" + "hello") + (v1 + v1)) + v1)), ((v1 + (v1 + v1)) + ""), "a"}
    var v3 uint = 2
    v4 := make(map[string]string)
    v5 := make(map[int]int)
    var v6 float64 = (1.5 / 1.0)
    fmt.Println(v3, v1)
    v2 = append(v2, v1)
    var v7 uint16 = 0
    var v8 uint = 65535
    v4[v1] = "hello"
    var v9 int64 = 586
    v3 = 1
    v9 += v9
    v1 = v1
    delete(v4, v1)
    fmt.Println(v9, v7, v9)
    switch "a" {
    case "foo bar":
        if (((0 != 956) && true) || (((^(v7)) < v7) || (((1.5 * ((0.0 + v6) - v6)) < v6) || (42 == 42)))) {
            var v10 float32 = -2.25
            _ = v10
        }
    case "hello":
        v1 = v1
    default:
        v3 -= (v3 | (v8 >> ((351) & 63)))
    }
    var v11 uint16 = (v7 ^ 440)
    v9 = v9
    var v12 float64 = (1.5 / v6)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
