package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
func genFn0(p0 int8, p1 uint32, p2 int8) uint32 {
    var v7 int64 = 1
    if (882 <= 255) {
        v8 := []int{32768, 9223372036854775807}
        _ = v8
    } else {
        p2 += p2
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 0
}

func main() {
    v1 := []int{32767, 107, -533, 128}
    var v2 uint8 = 42
    v3 := []uint32{}
    var v4 string = ""
    var v5 float64 = -2.25
    v6 := []string{v4}
    var v9 int32 = 2147483647
    v6 = append(v6, v4)
    v1 = append(v1, int(v2))
    v4 = "hello"
    fmt.Println(v9, v4)
    v10 := GenS1{F0: 453, F1: 39}
    var v11 string = (v4 + ("foo bar" + v4))
    v6 = append(v6, v4)
    fmt.Println(v5, v9)
    if false {
        if (618 != 4294967295) {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            fmt.Println(v2, v11, v4)
            v11 = v4
        }
        v11 = (("" + (("a" + (v11 + "hello")) + v4)) + v4)
        v5 = v5
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
