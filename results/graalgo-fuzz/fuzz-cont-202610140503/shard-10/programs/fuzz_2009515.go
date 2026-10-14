package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }

func main() {
    v1 := []string{"", ("foo bar" + ("hello" + "hello")), "\n\t"}
    v2 := make(map[uint32]uint32)
    var v3 string = ("hello" + "")
    v4 := []uint32{0, 2}
    var v5 float32 = 1.5
    var v6 uint32 = 4294967295
    fmt.Println(v5, v3)
    fmt.Println(v6)
    v7 := GenS1{F0: -253, F1: 9223372036854775807}
    var v8 uint32 = v6
    v9 := GenS1{F0: -498, F1: 453}
    if (((0.0 - ((0.0 / 1.5) - 0.0)) + ((1.5 * 3.14159) + float64((-2.25 * v5)))) >= 1.5) {
        if false {
            var v10 int64 = 0
            _ = v10
        } else {
            if (4294967295 <= -555) {
                var v11 bool = false
                if (!v11) {
                    v3 = v3
                    v8 = 685
                }
                _ = v11
            }
        }
        v8 -= (v8 + v8)
        v2[v8] = v8
    }
    var v12 uint16 = 65535
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
