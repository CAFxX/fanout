package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 string }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64 }

func main() {
    var v1 bool = (0 >= 32767)
    v2 := make(map[uint32]int)
    var v3 float32 = 3.14159
    var v4 int8 = 0
    var v5 int16 = 0
    v5 = 42
    var v6 uint8 = 1
    var v7 int = 588
    switch (("foo bar" + (("" + "a") + "foo bar")) + "hello") {
    case "\n\t":
        v8 := GenS1{F0: 65535, F1: -1, F2: (("" + ("hello" + "a")) + "a")}
        _ = v8
    default:
        v7 = v7
    }
    v3 = (v3 - ((0.0 * (v3 + (((v3 + v3) / 3.14159) + ((v3 / float32(float64(-2.25))) + 0.1)))) / float32(3.14159)))
    v9 := GenS2{F0: 0}
    v6 = (v6 & v6)
    v5 -= v5
    v10 := GenS1{F0: 0, F1: v4, F2: ""}
    v4 -= v4
    v11 := []int{v7, -2147483648, -242, (^(v7))}
    v12 := GenS2{F0: 2}
    v13 := GenS3{F0: 1}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
