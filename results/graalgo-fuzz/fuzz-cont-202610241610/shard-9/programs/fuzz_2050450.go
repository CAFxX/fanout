package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 string; F2 bool }
type GenS3 struct { F0 int8; F1 uint64; F2 bool }

func main() {
    var v1 uint = 255
    var v2 int32 = -960
    var v3 uint = (v1 - 683)
    var v4 uint16 = 32767
    var v5 bool = false
    fmt.Println(v1)
    fmt.Println(v2, v4, v4)
    v6 := GenS1{F0: (^(v3))}
    v7 := GenS2{F0: v5, F1: (("foo bar" + "foo bar") + "a"), F2: v5}
    var v8 bool = v5
    v3 = (v3 << ((42) & 63))
    var v9 float32 = 3.14159
    v10 := GenS3{F0: -14, F1: 1, F2: (32767 <= -1)}
    for fi11 := 0; fi11 < 5; fi11++ {
        v12 := GenS2{F0: v5, F1: "hello", F2: true}
        _ = fi11
        _ = v12
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
