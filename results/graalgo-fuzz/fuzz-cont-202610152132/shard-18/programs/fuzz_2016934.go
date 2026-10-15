package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 uint64, p1 string, p2 int64) int64 {
    var v4 int32 = 482
    var v5 uint8 = 127
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return p2
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int8 = -22
    var v3 int8 = v2
    var v6 float32 = 1e10
    v3 = (v2 | (v2 * (v2 << ((20) & 63))))
    v7 := []int{}
    switch (("" + "hello") + "\n\t") {
    case "foo bar":
        v7 = append(v7, 32768)
    case "hello":
        v8 := GenS1{F0: (^(v2))}
        _ = v8
    default:
        v3 += -53
    }
    switch 9223372036854775807 {
    case 1:
        v6 = float32(((0.0 - (0.0 / float64(float32((1e10 / 1e10))))) + float64(v6)))
    case 630:
        v3 = v3
    }
    v6 += (v6 / v6)
    v6 -= 0.0
    var v9 string = ((("hello" + "") + "foo bar") + "a")
    fmt.Println(v2, v3)
    var v10 uint16 = 65535
    v11 := GenS1{F0: (-(v2))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v9
}
