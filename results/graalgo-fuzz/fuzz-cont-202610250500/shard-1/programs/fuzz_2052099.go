package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 uint32 }
type GenS2 struct { F0 int; F1 uint64 }

func main() {
    var v1 bool = true
    var v2 float32 = 3.14159
    v3 := []uint32{1}
    var v4 uint = 65535
    var v5 int8 = 42
    var v6 uint64 = 664
    v7 := GenS1{F0: 65535, F1: 0, F2: 573}
    var v8 int = -210
    if (107 < 2) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v6 = (^(v6))
    }
    v9 := GenS2{F0: (v8 & 128), F1: (v6 >> ((42) & 63))}
    fmt.Println(v5, v2)
    switch ("a" + ("hello" + "hello")) {
    case "\n\t":
        v4 += 65535
    case "a":
        if ((!true) && (255 != 17)) {
            fmt.Println(v8, v4, v4)
            v10 := GenS2{F0: 1, F1: v6}
            _ = v10
        }
    }
    var v11 uint16 = 805
    v5 = (-(v5))
    v4 += 955
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
