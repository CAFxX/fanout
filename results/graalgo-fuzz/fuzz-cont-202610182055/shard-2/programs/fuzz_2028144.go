package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 int32; F1 uint16; F2 bool }
type GenS3 struct { F0 uint8; F1 uint32 }
func genFn0() uint8 {
    v7 := GenS1{F0: "\n\t", F1: -73}
    var v8 int64 = 1
    v9 := GenS1{F0: "hello", F1: -128}
    _ = v7
    _ = v8
    _ = v9
    return 1
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []string{"\n\t", "\n\t"}
    var v3 bool = false
    var v4 uint64 = 1
    v5 := []int{127, -1, -391}
    var v6 int64 = -187
    delete(v1, "foo bar")
    v4 = 48
    v3 = v3
    _ = len(v1)
    if len(v5) > 0 {
        v5[0] = 0
    }
    v6 = 9223372036854775807
    var v10 int = 4294967295
    v11 := GenS2{F0: 651, F1: 42, F2: v3}
    for fi12 := 0; fi12 < 2; fi12++ {
        v13 := GenS3{F0: 0, F1: 2}
        _ = fi12
        _ = v13
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
