package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 int16, p1 bool, p2 int32) uint64 {
    p2 = (p2 ^ (-(p2)))
    v6 := []int{-631, 128, -1, 32767}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 2
}
func genFn1() int {
    var v7 uint64 = 400
    var v8 uint32 = 2
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 float32 = ((1e10 + (1.5 / 3.14159)) / (1e10 - 1.5))
    v2 := []int{32768, 450, 950}
    var v3 int16 = 1
    var v4 int32 = -2147483648
    var v5 int16 = (v3 & (v3 * v3))
    var v9 int8 = 42
    fmt.Println(v3, v4, v9)
    fmt.Println(v1, v3)
    fmt.Println(v4)
    fmt.Println(v3, v3, v5)
    switch 879 {
    case 42:
        v4 += (v4 >> ((65535) & 63))
    case 0:
        v10 := GenS1{F0: ("" + "")}
        _ = v10
    }
    switch "" {
    case "foo bar":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case "hello":
        var v11 float64 = 1e10
        _ = v11
    default:
        fmt.Println(v3, v9)
    }
    fmt.Println(v3, v4, v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
