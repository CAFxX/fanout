package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    v1 := []string{("foo bar" + ("" + ("foo bar" + ("" + "\n\t")))), "\n\t", ""}
    var v2 float32 = 0.1
    v3 := []int{}
    var v4 uint8 = 1
    v4 = 161
    v2 = float32(float64((v2 * v2)))
    var v5 int64 = 9223372036854775807
    if false {
        if true {
            v1 = append(v1, "\n\t")
        } else {
            v3 = append(v3, -1)
            v4 += 75
            v5 = v5
        }
        var v6 uint8 = (v4 ^ v4)
        v5 += (v5 & (v5 / ((v5) | 1)))
        _ = v6
    }
    var v7 string = ("foo bar" + "a")
    v8 := GenS1{F0: 42}
    v4 = (v4 / ((0) | 1))
    v2 = (((v2 / ((float32((((0.0 - 3.14159) - 1.5) - (3.14159 * 0.0))) + v2) * (v2 * -2.25))) + float32(3.14159)) / 1.5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
