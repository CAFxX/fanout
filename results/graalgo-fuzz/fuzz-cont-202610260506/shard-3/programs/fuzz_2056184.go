package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 int16) int8 {
    p0 += (p0 + (p0 % ((p0) | 1)))
    v6 := []int{-229}
    _ = p0
    _ = v6
    return -88
}

func main() {
    v1 := []int{-79, -9223372036854775808}
    v2 := make(map[uint32]uint32)
    var v3 uint64 = 674
    var v4 uint = 18446744073709551615
    var v5 uint16 = 42
    v3 -= (v3 >> ((18446744073709551615) & 63))
    fmt.Println(v5)
    var v7 int16 = 0
    var v8 uint16 = 1
    v9 := GenS1{F0: (v7 >> (((v4 - v4)) & 63))}
    v10 := GenS1{F0: int16(v3)}
    if false {
        v5 = v8
        v4 += (v4 - v4)
        switch "a" {
        case "foo bar":
            v7 += (v7 % (((v7 | (v7 | int16(v3)))) | 1))
        default:
            if ("foo bar" > (("\n\t" + "hello") + "a")) {
                var v11 float32 = float32(0.1)
                _ = v11
            }
        }
    } else {
        v7 -= v7
        if len(v1) > 0 {
            v1[0] = 127
        }
    }
    v4 = (v4 ^ 0)
    var v12 int8 = 127
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
