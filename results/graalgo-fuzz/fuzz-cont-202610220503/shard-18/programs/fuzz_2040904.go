package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 int) bool {
    p0 += p0
    if (!(!true)) {
        v7 := GenS1{F0: 2, F1: -4}
        _ = v7
    }
    _ = p0
    return (972 != 704)
}
func genFn1(p0 int16) int16 {
    var v8 string = ("foo bar" + ("hello" + (("a" + "hello") + ((("" + "foo bar") + "hello") + "hello"))))
    v9 := GenS2{F0: 65535}
    _ = p0
    _ = v8
    _ = v9
    return (p0 / ((-32768) | 1))
}

func main() {
    var v1 float32 = 0.0
    var v2 float32 = (((v1 + v1) * -2.25) + 0.0)
    v3 := make(map[int]string)
    v4 := make(map[string]uint32)
    v5 := make(map[uint32]string)
    var v6 int64 = -498
    var v10 uint8 = 83
    var v11 uint32 = 4294967295
    v10 += 255
    fmt.Println(v2, v11, v2)
    v2 -= float32((0.1 * (0.1 + 0.1)))
    v12 := []int{175, 623, -749, 834}
    fmt.Println(v2)
    var v13 string = ("" + "\n\t")
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
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
}
