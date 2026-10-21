package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int; F1 int8; F2 int32 }

func main() {
    v1 := []string{"", (("\n\t" + "a") + ("hello" + (("\n\t" + "\n\t") + "hello"))), "hello"}
    v2 := []string{"foo bar"}
    var v3 uint8 = 1
    v4 := make(map[int]uint32)
    v5 := make(map[int]string)
    var v6 int32 = -1
    fmt.Println(v6, v3)
    v7 := GenS1{F0: true}
    if true {
        v8 := GenS2{F0: -9223372036854775808, F1: -104, F2: -206}
        _ = v8
    }
    v9 := GenS1{F0: (true || true)}
    for fi10 := 0; fi10 < 7; fi10++ {
        if len(v1) > 0 {
            v1[0] = "a"
        }
        _ = fi10
    }
    fmt.Println(v3)
    var v11 float32 = -2.25
    v6 -= v6
    fmt.Println(v3, v3)
    var v12 uint64 = 446
    v3 = v3
    var v13 uint = 2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
