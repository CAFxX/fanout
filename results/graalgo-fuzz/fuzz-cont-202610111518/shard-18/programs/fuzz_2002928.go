package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 int16 }

func main() {
    v1 := make(map[uint32]int)
    var v2 float32 = float32((1.5 * (0.0 / 1.0)))
    var v3 int32 = -257
    var v4 int = -457
    v5 := []int{(-(v4)), (-(v4))}
    v6 := []uint32{0, 881, 2, 638}
    _ = len(v1)
    var v7 int64 = 947
    fmt.Println(v7)
    var v8 int8 = 1
    var v9 int32 = int32(v4)
    for fi10 := 0; fi10 < 15; fi10++ {
        switch (v4 | v4) {
        case -9223372036854775808:
            v9 = -122
        case -995:
            v3 = v9
        }
        if (("a" + "hello") < ("foo bar" + "\n\t")) {
            v1[779] = v4
        }
        var v11 float32 = v2
        _ = fi10
        _ = v11
    }
    v12 := GenS1{F0: -32768, F1: "hello", F2: 1}
    var v13 uint16 = 163
    var v14 uint8 = 255
    var v15 float64 = (0.0 * 1e10)
    var v16 string = ("a" + "foo bar")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
