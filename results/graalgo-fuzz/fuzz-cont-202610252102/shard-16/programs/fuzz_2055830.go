package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0() int64 {
    v4 := []int{}
    var v5 uint64 = 0
    var v6 float64 = 0.0
    _ = v4
    _ = v5
    _ = v6
    return 9223372036854775807
}

func main() {
    var v1 int64 = 581
    v2 := make(map[int]uint32)
    v3 := make(map[string]string)
    fmt.Println(v1)
    if true {
        fmt.Println(v1)
    }
    v7 := GenS1{F0: ("hello" + "a")}
    var v8 int32 = 42
    fmt.Println(v8, v8)
    v9 := []int{}
    var v10 int = 9223372036854775807
    switch (v10 | 2147483647) {
    case 0:
        v1 += v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
