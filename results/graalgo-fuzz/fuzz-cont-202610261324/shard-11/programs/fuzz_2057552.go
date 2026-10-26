package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 int16 }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 string = (("" + "a") + "")
    var v2 uint16 = 255
    v3 := make(map[int]uint32)
    v4 := make(map[int]int)
    var v5 uint8 = 128
    var v6 float32 = (0.0 - 1.5)
    v2 -= (^(v2))
    if false {
        if ((-2.25 * (v6 + v6)) > v6) {
            v7 := v4[4294967295]
            v8 := GenS1{F0: -1, F1: v1, F2: 686}
            _ = v7
            _ = v8
        }
        v9 := GenS1{F0: 9223372036854775807, F1: v1, F2: -32768}
        v1 = "foo bar"
        _ = v9
    } else {
        v10 := GenS2{F0: 1}
        v11 := []int{-989}
        _ = v10
        _ = v11
    }
    fmt.Println(v6, v2, v6)
    switch (("foo bar" + (v1 + v1)) + ("a" + "a")) {
    case "hello":
        _ = len(v3)
    case "":
        var v12 int = 42
        _ = v12
    }
    v2 -= (v2 % ((uint16(v5)) | 1))
    v6 = ((v6 * v6) - v6)
    var v13 int16 = -32768
    var v14 int8 = -92
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
