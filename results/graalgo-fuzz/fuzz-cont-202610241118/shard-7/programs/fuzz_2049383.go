package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 uint8; F1 int16 }
type GenS4 struct { F0 string; F1 int8 }
func genFn0() uint {
    v7 := GenS1{F0: 0}
    if true {
        v8 := []int{0, 9223372036854775807}
        var v9 bool = false
        _ = v8
        _ = v9
    } else {
        var v10 uint16 = 65535
        _ = v10
    }
    v11 := []int{4294967295, 2147483647, 0}
    _ = v11
    _ = v7
    return 1
}

func main() {
    var v1 uint = 129
    var v2 uint = 104
    v3 := []int{128, 9223372036854775807, -9223372036854775808}
    v4 := make(map[string]int)
    v5 := []string{}
    var v6 uint64 = 1
    for fi12 := 0; fi12 < 1; fi12++ {
        v2 -= v2
        v2 += 18446744073709551615
        v1 = (v2 >> ((v1) & 63))
        _ = fi12
    }
    var v13 int = 4294967295
    var v14 int8 = 9
    v13 = v13
    v15 := GenS2{F0: 1, F1: 534}
    v16 := GenS3{F0: 128, F1: 1000}
    v17 := GenS4{F0: "foo bar", F1: (v14 * (v14 | v14))}
    if true {
        v18 := GenS4{F0: "", F1: (^(v14))}
        _ = v18
    }
    v2 = (^(v1))
    v19 := GenS2{F0: genFn0(), F1: 919}
    fmt.Println(v2, v2, v13)
    var v20 int8 = 100
    v1 = v1
    for fi21 := 0; fi21 < 13; fi21++ {
        var v22 float32 = (0.0 / (3.14159 / 0.1))
        _ = fi21
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
