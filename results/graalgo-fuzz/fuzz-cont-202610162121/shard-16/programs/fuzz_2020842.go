package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 float64 = 0.1
    var v2 int8 = 100
    v3 := make(map[uint32]uint32)
    v4 := make(map[uint32]string)
    v5 := []uint32{2, 0}
    _ = len(v5)
    switch 127 {
    case 9223372036854775807:
        v6 := GenS1{F0: 4294967295}
        _ = v6
    default:
        v1 -= 1.5
    }
    var v7 int8 = (-(v2))
    var v8 uint = 255
    v9 := v3[2147483648]
    fmt.Println(v7, v9, v7)
    v10 := v3[1]
    v7 = v2
    v8 -= v8
    var v11 bool = (32767 > -2147483648)
    v10 += v9
    v5 = append(v5, v9)
    v8 -= v8
    fmt.Println(v11, v11)
    v8 = v8
    v11 = ((-416 == 1) && (("foo bar" + "a") < (("a" + "a") + "hello")))
    if (1 >= 167) {
        delete(v3, (v10 / ((v10) | 1)))
    } else {
        v7 += (v2 * v2)
        v11 = v11
    }
    var v12 uint16 = 599
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
