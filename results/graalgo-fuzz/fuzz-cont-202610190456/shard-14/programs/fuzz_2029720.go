package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 int8 }

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 float64 = (0.0 + 1e10)
    var v3 int32 = 1
    v4 := make(map[uint32]uint32)
    v5 := []int{128, 7}
    delete(v4, 2147483648)
    var v6 bool = (("\n\t" + "\n\t") <= "foo bar")
    var v7 int64 = 9223372036854775807
    if v6 {
        if false {
            v1 = 42
        } else {
            for fi8 := 0; fi8 < 2; fi8++ {
                v9 := GenS1{F0: 2147483648}
                v1 = 18446744073709551615
                for fi10 := 0; fi10 < 3; fi10++ {
                    _ = len(v5)
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    v1 = 323
                    if v6 {
                        var v11 float32 = float32(v2)
                        _ = v11
                    } else {
                        v2 = (v2 - v2)
                    }
                    fmt.Println(v1, v2)
                    _ = fi10
                }
                _ = fi8
                _ = v9
            }
            var v12 uint64 = (v1 - 18446744073709551615)
            _ = v12
        }
        var v13 int = 9223372036854775807
        _ = v13
    }
    var v14 float64 = -2.25
    var v15 float64 = float64(-2.25)
    switch 9223372036854775807 {
    case 127:
        var v16 float32 = (1.5 - 1.5)
        _ = v16
    }
    v17 := GenS1{F0: uint32(v1)}
    if len(v5) > 0 {
        v5[0] = -904
    }
    v18 := GenS2{F0: 2147483647, F1: 42}
    v6 = (2 >= 2)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
