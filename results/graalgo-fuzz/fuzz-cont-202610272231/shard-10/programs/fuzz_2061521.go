package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 string; F1 string }
func genFn0() string {
    v7 := GenS1{F0: 691, F1: 1}
    var v8 int16 = -146
    v8 = v8
    v9 := GenS2{F0: 0}
    _ = v7
    _ = v8
    _ = v9
    return "hello"
}

func main() {
    var v1 float32 = float32(1.5)
    v2 := make(map[string]bool)
    v3 := make(map[uint32]bool)
    var v4 string = "hello"
    var v5 int16 = -402
    var v6 int = -9223372036854775808
    var v10 uint16 = 65535
    v5 += (v5 >> ((255) & 63))
    v2[v4] = false
    delete(v3, 2)
    for fi11 := 0; fi11 < 14; fi11++ {
        v1 += v1
        _ = fi11
    }
    v12 := GenS3{F0: "foo bar", F1: (v4 + v4)}
    fmt.Println(v4, v6)
    var v13 float64 = 1e10
    v14 := []int{(v6 ^ 761)}
    _ = len(v3)
    v13 = (1e10 * (v13 * v13))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
