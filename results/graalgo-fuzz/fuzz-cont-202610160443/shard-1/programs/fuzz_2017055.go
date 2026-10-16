package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }

func main() {
    var v1 int32 = -2147483648
    var v2 int64 = 30
    v3 := make(map[int]string)
    fmt.Println(v2, v1)
    v1 -= v1
    v2 = (v2 | (v2 + (v2 % ((v2) | 1))))
    v4 := GenS1{F0: 537, F1: ((("a" + "\n\t") + "a") + "hello")}
    v1 += (v1 >> ((65535) & 63))
    var v5 float64 = ((1e10 - ((3.14159 * 3.14159) * 3.14159)) + 1e10)
    v1 -= 170
    v3[4294967295] = (("\n\t" + "a") + "\n\t")
    v5 = (v5 * v5)
    fmt.Println(v1)
    var v6 float64 = v5
    v2 = v2
    v5 += (v6 * 0.0)
    v2 = (^(v2))
    fmt.Println(v5, v1, v2)
    var v7 int64 = (v2 - -9223372036854775808)
    v8 := v3[42]
    fmt.Println(v6, v6)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
